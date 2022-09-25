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

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		newCtx := client(ctx, method)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
		streamer grpc.Streamer, opts ...grpc.CallOption) (stream grpc.ClientStream, err error) {

		newCtx := client(ctx, method)
		return streamer(newCtx, desc, cc, method, opts...)
	}
}

func client(ctx context.Context, method string) context.Context {
	var parentCtx opentracing.SpanContext
	var spanOpts []opentracing.StartSpanOption
	var parentSpan = opentracing.SpanFromContext(ctx)

	if parentSpan != nil {
		parentCtx = parentSpan.Context()
		spanOpts = append(spanOpts, opentracing.ChildOf(parentCtx))
	}

	opts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCClient,
	}
	spanOpts = append(spanOpts, opts...)

	span := global.Tracer.StartSpan(method, spanOpts...)
	defer span.Finish()

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	_ = global.Tracer.Inject(span.Context(), opentracing.TextMap, metatext.MetaDataTextMap{md})

	return opentracing.ContextWithSpan(ctx, span)
}
