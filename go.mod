module grpc-example

go 1.15

require (
	github.com/docker/docker v20.10.5+incompatible
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	google.golang.org/genproto v0.0.0-20220916172020-2692e8806bfa // indirect
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.0
