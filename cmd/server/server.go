package main

import (
	"context"
	"flag"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-example/pkg/swagger"
	pb "grpc-example/proto"
	"grpc-example/service"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
)

var port string

func init() {
	flag.StringVar(&port, "port", "9001", "启动端口号")
}

func main() {
	log.Println("server is starting...")
	svr := NewServer(port)

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from ", err)
		}
	}()

	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	//log.Println("Begin shutdown server")
	//if err := svr.Shutdown(context.Background()); err != nil {
	//	log.Fatal(err)
	//}

}

func NewTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}
func NewGrpcServer() *grpc.Server {
	server := grpc.NewServer()

	// 注册服务
	pb.RegisterSearchServiceServer(server, service.NewSearch())
	pb.RegisterPubSubServiceServer(server, service.NewPubSub())
	reflection.Register(server)

	return server
}

func NewHttpServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})
	mux.HandleFunc("/swagger/", serveSwaggerFile())

	serveSwaggerUI(mux)

	return mux
}

func RunGrpcGateway(port string) *runtime.ServeMux {
	endporint := ":" + port
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 注册服务
	pb.RegisterSearchServiceHandlerFromEndpoint(context.Background(), mux, endporint, opts)
	pb.RegisterPubSubServiceHandlerFromEndpoint(context.Background(), mux, endporint, opts)

	return mux
}

func NewServer(port string) *http.Server {
	httpSvr := NewHttpServer()
	grpcSvr := NewGrpcServer()
	gateway := RunGrpcGateway(port)

	httpSvr.Handle("/", gateway)

	log.Println("server is running...")
	return &http.Server{
		Addr:    ":" + port,
		Handler: grpcHandlerFunc(grpcSvr, httpSvr),
	}
}

func grpcHandlerFunc(grpcSvr *grpc.Server, httpSvr http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-type"), "application/grpc") {
			log.Println("Receive grpc request")
			grpcSvr.ServeHTTP(w, r)
		} else {
			log.Println("Receive http request")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Body: ", string(body))

			httpSvr.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func serveSwaggerFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("start serveSwaggerFile")

		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			log.Printf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("docs/", p)

		log.Printf("Serving swagger-file: %s", p)

		http.ServeFile(w, r, p)
	}
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
