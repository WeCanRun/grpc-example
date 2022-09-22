package service

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
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
	md, _ := metadata.FromIncomingContext(ctx)

	log.Printf("Request: %#v, metadata: %#v", r, md)
	resp := pb.SearchResponse{Response: r.Request}

	return errcode.ErrorNotExistTag.Response(&resp)
}

func (s *SearchService) Channel(stream pb.SearchService_ChannelServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	log.Printf("Metadata: %#v", md)
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
