package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-example/pkg/client/auth"
	pb "grpc-example/proto"
	"io"
	"log"
	"time"
)

const PORT = "9001"

var ctx = metadata.AppendToOutgoingContext(context.Background(), "key", "value")

func CallChannel(c pb.SearchServiceClient) {
	stream, err := c.Channel(ctx)
	if err != nil {
		log.Printf("call channel, err: %v: ", err.Error())
	}

	go func() {
		for {
			if err := stream.Send(&pb.SearchRequest{Request: "call chanel"}); err != nil {
				log.Println("send: " + err.Error())
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
			log.Println("recv: " + err.Error())
		}
		log.Println("resp: " + recv.Response)
	}
}

func Publish(c pb.PubSubServiceClient, publish string) {
	_, err := c.Publish(ctx, &pb.PubRequest{Publish: publish})
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("published %s ...", publish))
}

func SubscribeTopic(c pb.PubSubServiceClient) {
	stream, err := c.Subscribe(ctx, &pb.SubRequest{Subscribe: "golang:"})
	if err != nil {
		log.Println(err)
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println(err)
		}
		log.Println("subscribe: " + recv.String())
	}
}

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth.Auth{
		AppKey:    auth.GetAppKey(),
		AppSecret: auth.GetAppSecret(),
	})}

	conn, err := grpc.Dial(":"+PORT, opts...)

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC ",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.String())

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
