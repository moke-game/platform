// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: matchmaking/matchmaking.proto

/*
Package matchmaking is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package matchmaking

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_MatchService_Match_0(ctx context.Context, marshaler runtime.Marshaler, client MatchServiceClient, req *http.Request, pathParams map[string]string) (MatchService_MatchClient, runtime.ServerMetadata, error) {
	var protoReq MatchRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.Match(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterMatchServiceHandlerServer registers the http handlers for service MatchService to "mux".
// UnaryRPC     :call MatchServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMatchServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterMatchServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MatchServiceServer) error {

	mux.Handle("POST", pattern_MatchService_Match_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterMatchServiceHandlerFromEndpoint is same as RegisterMatchServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMatchServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMatchServiceHandler(ctx, mux, conn)
}

// RegisterMatchServiceHandler registers the http handlers for service MatchService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMatchServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMatchServiceHandlerClient(ctx, mux, NewMatchServiceClient(conn))
}

// RegisterMatchServiceHandlerClient registers the http handlers for service MatchService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MatchServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MatchServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MatchServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterMatchServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MatchServiceClient) error {

	mux.Handle("POST", pattern_MatchService_Match_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/matchmaking.v1.MatchService/Match", runtime.WithHTTPPathPattern("/matchmaking.v1.MatchService/Match"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MatchService_Match_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MatchService_Match_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MatchService_Match_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"matchmaking.v1.MatchService", "Match"}, ""))
)

var (
	forward_MatchService_Match_0 = runtime.ForwardResponseStream
)
