package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-example/proto"
	"io"
	"log"
	"time"
)

const PORT = "9001"

func CallChannel(c pb.SearchServiceClient) {
	stream, err := c.Channel(context.Background())
	if err != nil {
		log.Fatal("call channel: " + err.Error())
	}

	go func() {
		for {
			if err := stream.Send(&pb.SearchRequest{Request: "call chanel"}); err != nil {
				log.Fatal("send: " + err.Error())
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
			log.Fatal("recv: " + err.Error())
		}
		log.Println("resp: " + recv.Response)
	}
}

func Publish(c pb.PubSubServiceClient, publish string) {
	_, err := c.Publish(context.Background(), &pb.PubRequest{Publish: publish})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("published %s ...", publish))
}

func SubscribeTopic(c pb.PubSubServiceClient) {
	stream, err := c.Subscribe(context.Background(), &pb.SubRequest{Subscribe: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		log.Println("subscribe: " + recv.Value)
	}
}

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
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

	log.Printf("resp: %s", resp.GetResponse())

	go CallChannel(client)

	svcClient := pb.NewPubSubServiceClient(conn)
	go func() {
		for  {
			Publish(svcClient, "golang: hello go")
			Publish(svcClient, "docker: hello go")
		}
	}()
	SubscribeTopic(svcClient)
}
