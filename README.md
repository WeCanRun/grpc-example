# grpc-example

## 步骤
1. 定义 proto 文件
2. 生成 grpc 代码
    ```shell script
    protoc -I . --go_out=plugins=grpc:.  proto/*.proto
    ```
3. 创建 grpc 服务器，并注册相关服务
4. 修改 proto 文件，加入 grpc-gatway 相关配置
5. 生成 grpc-gateway 相关代码
    ```shell script
    protoc -I ./ -I ../../googleapis/googleapis  --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_optpaths=source_relative --grpc-gateway_opt generate_unbound_methods=true    proto/*.proto
    ```


## 参考
https://github.com/grpc-ecosystem/grpc-gateway
https://www.cnblogs.com/yisany/p/14875488.html
