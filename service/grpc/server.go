package grpc

import (
	"context"
)

type Server struct {
}

func (s Server) Setup(context.Context, *SetupRequest) (*SetupResponse, error) {
	panic("implement me")
}

func (s Server) Teardown(context.Context, *TeardownRequest) (*EmptyResponse, error) {
	panic("implement me")
}
