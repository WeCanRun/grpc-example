package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-example/global/constants"
	"grpc-example/global/errcode"
	"grpc-example/pkg/client/auth"
)

func AuthUnaryServer(enable bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if enable {
			if err := _auth(ctx); err != nil {
				return nil, err
			}
		}

		resp, err = handler(ctx, req)
		return resp, err
	}
}

func AuthStreamServerInterceptor(enable bool) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if enable {
			if err := _auth(ss.Context()); err != nil {
				return err
			}
		}

		err := handler(srv, ss)
		return err
	}
}

func _auth(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var key, secret string
	if value, ok := md[constants.AppKey]; ok {
		key = value[0]
	}

	if value, ok := md[constants.AppSecret]; ok {
		secret = value[0]
	}

	if key != auth.GetAppKey() || secret != auth.GetAppSecret() {
		return errors.New(errcode.Unauthorized.Msg())
	}

	return nil
}
