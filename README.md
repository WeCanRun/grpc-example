# grpc-example

## 快速开始
1. 运行代码
    ```shell script
    git clone git@github.com:WeCanRun/grpc-example.git
    make && make run  
    ```

2. 访问 `swagger-ui` 首页
    ```shell script
    http://localhost:9001/swagger-ui/
    ```
3.  输入`swagger`文档地址, 点击 `Explore`
    ```shell script
    http://localhost:9001/swagger/all.swagger.json
    ```

## 步骤
0. 版本信息
   ```shell script
   libprotoc 3.21.5
   protoc-gen-go v1.28.1
   protoc-gen-go-grpc 1.1.0
   ```
1. 定义 `proto` 文件 `proto/*.proto`
2. 生成 `grpc` 代码
    ```shell script
    protoc -I . -I ../../googleapis/googleapis  --go_out ./ --go_opt paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./ --go-grpc_opt paths=source_relative  proto/*.proto
    ```
3. 创建 `grpc` 服务器，并注册相关服务 `cmd/server/server.go`
4. 修改 `proto` 文件，加入 `grpc-gatway` 相关配置
5. 生成 `grpc-gateway` 相关代码
    ```shell script
    protoc -I . -I ../../googleapis/googleapis  --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true  proto/*.proto
    ```
6. 使用 `protoc-gen-openapiv2` 生成 `swagger` 文档
    ```shell script
    protoc -I . -I ../../googleapis/googleapis --openapiv2_out ./docs --openapiv2_opt use_go_templates=true  --openapiv2_opt logtostderr=true  --openapiv2_opt  allow_merge=true,merge_file_name=all   proto/*.proto
    ``` 
7. 增加 `swagger` 路由
   ```shell script
   mkdir third_patry && cd third_patry && git clone git@github.com:swagger-api/swagger-ui.git
   mkdir swagger && cp -r swagger-ui/dist/* swagger/ && rm -rf swagger-ui/
   
   go get -u github.com/go-bindata/go-bindata/...
   go get -u github.com/elazarl/go-bindata-assetfs/...

   go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger/...
   ```
   ```go
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
   ```
        
      

## 参考
https://github.com/grpc-ecosystem/grpc-gateway

https://www.cnblogs.com/yisany/p/14875488.html
