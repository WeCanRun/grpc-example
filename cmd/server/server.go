package main

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/global"
	log "github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/tracer"
	"grpc-example/pkg/server"
)

func main() {
	//加载配置文件
	s := setting.Setup("")
	global.Setting = s

	log.Setup()
	tracer.Setup("example", fmt.Sprintf(":%d", 6831))

	log.Info("Server is starting...")

	_ = server.Run()

}
