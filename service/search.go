package service

import (
	"context"
	"errors"
	"grpc-example/global/errcode"
	pb "grpc-example/proto"
	"io"
	"log"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
}

func NewSearch() pb.SearchServiceServer {
	return &SearchService{}
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.Response, error) {
	log.Printf("Request: %#v", r)
	resp := pb.SearchResponse{Response: r.Request}

	return errcode.ErrorNotExistTag.ToResponse(&resp)
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

		err = stream.Send(&pb.SearchResponse{Response: args.Request})
		if err != nil {
			return err
		}
	}
}
