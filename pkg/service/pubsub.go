package service

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc/metadata"
	"grpc-example/global/errcode"
	pb "grpc-example/proto"
	"log"
	"strings"
	"time"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSub() pb.PubSubServiceServer {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Microsecond, 10),
	}
}

func (s *PubSubService) Publish(ctx context.Context, req *pb.PubRequest) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Publish metadata: %#v", md)
	s.pub.Publish(req.GetPublish())
	return errcode.Success.Response(&pb.PubRequest{Publish: req.GetPublish()})
}

func (s *PubSubService) Subscribe(req *pb.SubRequest, stream pb.PubSubService_SubscribeServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	log.Printf("Subscribe metadata: %#v", md)
	ch := s.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasSuffix(key, req.GetSubscribe()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		response, _ := errcode.Success.Response(&pb.SubResponse{Response: v.(string)})
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return nil
}
