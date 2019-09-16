package grpc

import (
	"context"
)

type Client struct {
}

func (c Client) New(ctx context.Context, in *NewRequest, opts ...interface{}) (*OKResponse, error) {
	panic("implement me")
}

func (c Client) Delete(ctx context.Context, in *DeleteRequest, opts ...interface{}) (*OKResponse, error) {
	panic("implement me")
}
