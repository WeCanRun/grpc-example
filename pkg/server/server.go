package server

import (
	"context"
	"flag"
	log "github.com/WeCanRun/gin-blog/pkg/logging"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-example/pkg/middleware"
	"grpc-example/pkg/middleware/tracing"
	"grpc-example/pkg/service"
	"grpc-example/pkg/swagger"
	pb "grpc-example/proto"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

var svr *http.Server
var Port string

func init() {
	flag.StringVar(&Port, "port", "9001", "启动端口号")
}

func New() *http.Server {
	httpSvr := NewHttpServer()
	grpcSvr := NewGrpcServer()
	gateway := NewGrpcGateway(Port)

	httpSvr.Handle("/", gateway)

	log.Info("Server is running...")
	svr = &http.Server{
		Addr:    ":" + Port,
		Handler: grpcHandlerFunc(grpcSvr, httpSvr),
	}
	return svr
}

func Run() error {
	if svr == nil {
		New()
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Infof("ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	// 阻塞、等待终止信号
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		log.Infof("Shutdown err: %v", err)
	}

	log.Info("Server is shutdowned")

	return nil
}

func NewGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			middleware.AccessUnaryServer(true),
			//middleware.AuthUnaryServer(true),
			tracing.UnaryServerInterceptor(true),
		)),

		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(),
			middleware.AccessStreamServer(true),
			//middleware.AuthStreamServerInterceptor(true),
			//tracing.StreamServerInterceptor(true),
		)),
	}

	server := grpc.NewServer(opts...)

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

func NewGrpcGateway(port string) *runtime.ServeMux {
	ctx := context.Background()
	endporint := ":" + port
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 注册服务
	pb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, endporint, opts)
	pb.RegisterPubSubServiceHandlerFromEndpoint(ctx, mux, endporint, opts)

	return mux
}

func grpcHandlerFunc(grpcSvr *grpc.Server, httpSvr http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-type"), "application/grpc") {
			log.Info("Receive grpc request")
			grpcSvr.ServeHTTP(w, r)
		} else {
			log.Info("Receive http request")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Info("Body: ", string(body))

			httpSvr.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func serveSwaggerFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start serveSwaggerFile")

		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			log.Infof("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("docs/", p)

		log.Infof("Serving swagger-file: %s", p)

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
