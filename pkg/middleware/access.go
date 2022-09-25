package middleware

import (
	"context"
	log "github.com/WeCanRun/gin-blog/pkg/logging"
	"google.golang.org/grpc"
	"time"
)

func AccessUnaryServer(enable bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if enable {
			begin := beforeAccess(req, info.FullMethod)
			resp, err := handler(ctx, req)
			AfterAccess(info.FullMethod, begin, resp)
			return resp, err
		}

		return handler(ctx, req)
	}
}

func AccessStreamServer(enable bool) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if enable {
			begin := beforeAccess(ss.Context(), info.FullMethod)
			err := handler(srv, ss)
			AfterAccess(info.FullMethod, begin, nil)
			return err
		}

		return handler(srv, ss)
	}
}

func beforeAccess(req interface{}, method string) int64 {
	requestLog := "access request log: method: %s, begin_time: %d, request: %v"
	begin := time.Now().Unix()
	log.Infof(requestLog, method, begin, req)

	return begin
}

func AfterAccess(method string, begin int64, resp interface{}) {
	end := time.Now().Unix()
	respLog := "access response log: method: %s, begin_time: %d, end_time: %d, spend_time: %d, response: %v"
	log.Infof(respLog, method, begin, end, end-begin, resp)
}
