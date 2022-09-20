all: gen run

gen:
	protoc -I . -I ../../googleapis/googleapis  --go_out ./ --go_opt paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./ --go-grpc_opt paths=source_relative  proto/*.proto
	protoc -I . -I ../../googleapis/googleapis  --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true  proto/*.proto
	protoc -I . -I ../../googleapis/googleapis --openapiv2_out ./docs --openapiv2_opt use_go_templates=true  --openapiv2_opt logtostderr=true  --openapiv2_opt  allow_merge=true,merge_file_name=all   proto/*.proto

	go generate .

build:
	go build -v -o bin/server cmd/server/server.go

run: build
	./bin/server

tool:
	go tool vet . |& grep -v vendor; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf bin/*
	go clean -i .


