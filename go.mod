module grpc-example

go 1.15

require (
	github.com/docker/docker v20.10.5+incompatible
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e
	google.golang.org/genproto v0.0.0-20220916172020-2692e8806bfa
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.0
