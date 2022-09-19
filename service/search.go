package service

import (
	"context"
	"errors"
	"fmt"
	pb "grpc-example/proto"
	"io"
	"log"
)

type SearchService struct{}

func NewSearch() pb.SearchServiceServer {
	return &SearchService{}
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Printf("Request: %#v", r)
	value := ctx.Value("request")
	log.Println("value: ", value)
	return &pb.SearchResponse{Response: r.GetRequest() + " Server response"}, nil
}

func (s *SearchService) Channel(stream pb.SearchService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		err = stream.Send(&pb.SearchResponse{
			Response: fmt.Sprintf("server recv [%s] ", args.GetRequest()),
		})
		if err != nil {
			return err
		}
	}
}
