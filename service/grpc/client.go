package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type Client struct {
}

func (c Client) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*SetupResponse, error) {
	panic("implement me")
}

func (c Client) Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	panic("implement me")
}
