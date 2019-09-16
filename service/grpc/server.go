package grpc

import (
	"context"
)

type Server struct {
}

func (s Server) New(context.Context, *NewRequest) (*OKResponse, error) {
	panic("implement me")
}

func (s Server) Delete(context.Context, *DeleteRequest) (*OKResponse, error) {
	panic("implement me")
}
