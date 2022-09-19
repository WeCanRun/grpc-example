package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-example/proto"
	"grpc-example/service"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

var port string

func init() {
	flag.StringVar(&port, "port", "9001", "启动端口号")
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

func RunServer(port string) error {
	httpSvr := NewHttpServer()
	grpcSvr := NewGrpcServer()
	gateway := RunGrpcGateway(port)

	httpSvr.Handle("/", gateway)

	log.Println("server is running...")
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcSvr, httpSvr))
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

func main() {
	log.Println("server is starting...")
	_ = RunServer(port)
}
