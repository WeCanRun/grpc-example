package main

import (
	"flag"
	"google.golang.org/grpc"
	pb "grpc-example/proto"
	"grpc-example/service"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
)

var port string

func init() {
	flag.StringVar(&port, "port", "9001", "启动端口号")
}

func NewTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func NewGrcpSrever() *grpc.Server {
	server := grpc.NewServer()

	// 注册服务
	pb.RegisterSearchServiceServer(server, service.NewSearch())
	pb.RegisterPubSubServiceServer(server, service.NewPubSub())
	return server
}
func NewHttpServer(port string) *http.ServeMux {
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
	
	return mux
}

func RunServer(port string) error  {
	httpSvr := NewHttpServer(port)
	grpcSvr := NewGrcpSrever()
	gateway := RunGrpcGateway(port)
	httpSvr.Handle("/", gateway)
	return http.ListenAndServe(":"+port, grpcHandlefFunc(grpcSvr, httpSvr))
}

func grpcHandlefFunc(grpcSvr *grpc.Server, httpSvr http.Handler) http.Handler {
	return h2c.NewHanler(http.HandleFunc(func(w http.ResponseWriter, r http.Request	) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-type")
		, "application/grpc") {
			grpcSvr.ServeHTTP(w, r)
		} else {
			httpSvr.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func main() {
	log.Println("server is starting...")
	lis, err := NewTcpServer(port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	svr := NewGrcpSrever()


	log.Println("server is running...")
	svr.Serve(lis)
}
