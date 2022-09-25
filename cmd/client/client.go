package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/WeCanRun/gin-blog/global"
	log "github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/tracer"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-example/pkg/client/auth"
	"grpc-example/pkg/middleware/tracing"
	"grpc-example/pkg/server"
	pb "grpc-example/proto"
	"io"
	"time"
)

var ctx = metadata.AppendToOutgoingContext(context.Background(), "key", "value")

func CallChannel(c pb.SearchServiceClient) {
	stream, err := c.Channel(ctx)
	if err != nil {
		log.Infof("call channel, err: %v: ", err.Error())
	}

	go func() {
		for {
			if err := stream.Send(&pb.SearchRequest{Request: "call chanel"}); err != nil {
				log.Info("send: " + err.Error())
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		recv, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Info("recv: " + err.Error())
			return
		}
		log.Info("resp: " + recv.Response)
	}
}

func Publish(c pb.PubSubServiceClient, publish string) {
	_, err := c.Publish(ctx, &pb.PubRequest{Publish: publish})
	if err != nil {
		log.Info(err)
	}
	log.Info(fmt.Sprintf("published %s ...", publish))
}

func SubscribeTopic(c pb.PubSubServiceClient) {
	stream, err := c.Subscribe(ctx, &pb.SubRequest{Subscribe: "golang:"})
	if err != nil {
		log.Info(err)
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Info(err)
		}
		log.Info("subscribe: " + recv.String())
	}
}

func main() {
	s := setting.Setup("")
	global.Setting = s

	log.Setup()

	tracer.Setup("example-client", ":6831")

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(&auth.Auth{
			AppKey:    auth.GetAppKey(),
			AppSecret: auth.GetAppSecret(),
		}),

		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			tracing.UnaryClientInterceptor(),
		)),

		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			tracing.StreamClientInterceptor(),
		)),
	}

	conn, err := grpc.Dial(":"+server.Port, opts...)

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{
		Request: "gRPC ",
	})

	if err != nil {
		log.Infof("client.Search err: %v", err)
	}

	log.Infof("resp: %s", resp.String())

	go CallChannel(client)

	svcClient := pb.NewPubSubServiceClient(conn)
	go func() {
		for {
			Publish(svcClient, "golang: hello go")
			Publish(svcClient, "docker: hello go")
			time.Sleep(5 * time.Second)
		}
	}()

	SubscribeTopic(svcClient)

	select {}
}
