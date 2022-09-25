module grpc-example

go 1.15

require (
	github.com/WeCanRun/gin-blog v0.0.0-20220919113150-348471f76543
	github.com/docker/docker v20.10.5+incompatible
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	google.golang.org/genproto v0.0.0-20220920201722-2b89144ce006
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.0
