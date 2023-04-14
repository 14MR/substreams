// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sf/substreams/rpc/v2/service.proto

package pbsubstreamsrpcconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v2 "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StreamName is the fully-qualified name of the Stream service.
	StreamName = "sf.substreams.rpc.v2.Stream"
)

// StreamClient is a client for the sf.substreams.rpc.v2.Stream service.
type StreamClient interface {
	Blocks(context.Context, *connect_go.Request[v2.Request]) (*connect_go.ServerStreamForClient[v2.Response], error)
}

// NewStreamClient constructs a client for the sf.substreams.rpc.v2.Stream service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStreamClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StreamClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &streamClient{
		blocks: connect_go.NewClient[v2.Request, v2.Response](
			httpClient,
			baseURL+"/sf.substreams.rpc.v2.Stream/Blocks",
			opts...,
		),
	}
}

// streamClient implements StreamClient.
type streamClient struct {
	blocks *connect_go.Client[v2.Request, v2.Response]
}

// Blocks calls sf.substreams.rpc.v2.Stream.Blocks.
func (c *streamClient) Blocks(ctx context.Context, req *connect_go.Request[v2.Request]) (*connect_go.ServerStreamForClient[v2.Response], error) {
	return c.blocks.CallServerStream(ctx, req)
}

// StreamHandler is an implementation of the sf.substreams.rpc.v2.Stream service.
type StreamHandler interface {
	Blocks(context.Context, *connect_go.Request[v2.Request], *connect_go.ServerStream[v2.Response]) error
}

// NewStreamHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStreamHandler(svc StreamHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sf.substreams.rpc.v2.Stream/Blocks", connect_go.NewServerStreamHandler(
		"/sf.substreams.rpc.v2.Stream/Blocks",
		svc.Blocks,
		opts...,
	))
	return "/sf.substreams.rpc.v2.Stream/", mux
}

// UnimplementedStreamHandler returns CodeUnimplemented from all methods.
type UnimplementedStreamHandler struct{}

func (UnimplementedStreamHandler) Blocks(context.Context, *connect_go.Request[v2.Request], *connect_go.ServerStream[v2.Response]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.rpc.v2.Stream.Blocks is not implemented"))
}