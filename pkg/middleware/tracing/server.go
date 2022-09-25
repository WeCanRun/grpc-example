package tracing

import (
	"context"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-example/pkg/metatext"
)

func UnaryServerInterceptor(enable bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if enable {
			ctx = server(ctx, info.FullMethod)
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor(enable bool) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		if enable {
			ctx = server(ctx, info.FullMethod)
		}
		return handler(ctx, ss)
	}
}

func server(ctx context.Context, method string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	spanCtx, _ := global.Tracer.Extract(opentracing.TextMap, metatext.MetaDataTextMap{md})
	spanOpts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCServer,
		ext.RPCServerOption(spanCtx),
	}

	span := global.Tracer.StartSpan(method, spanOpts...)
	defer span.Finish()

	return opentracing.ContextWithSpan(ctx, span)
}
