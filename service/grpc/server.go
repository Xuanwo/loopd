package grpc

import (
	"context"
	"log"

	"github.com/Xuanwo/loopd/loop"
)

// Server will implements LoopdServer interface.
type Server struct {
	devices map[string]string
}

// NewServer will create a new server.
func NewServer() *Server {
	return &Server{devices: make(map[string]string)}
}

// Setup implements LoopdServer.Setup.
func (s Server) Setup(ctx context.Context, req *SetupRequest) (resp *SetupResponse, err error) {
	resp = &SetupResponse{}

	deviceName, err := loop.NextFreeDevice()
	if err != nil {
		return
	}

	err = loop.NewLoopDevice(deviceName)
	if err != nil {
		return
	}

	err = loop.Setup(deviceName, req.Filename)
	if err != nil {
		return
	}

	s.devices[deviceName] = req.Filename
	return
}

// Teardown implements LoopdServer.Teardown.
func (s Server) Teardown(ctx context.Context, req *TeardownRequest) (resp *EmptyResponse, err error) {
	resp = &EmptyResponse{}

	if _, ok := s.devices[req.DeviceName]; !ok {
		log.Printf("Device [%s] is not found in db.", req.DeviceName)
	}

	err = loop.TearDown(req.DeviceName)
	if err != nil {
		return
	}

	delete(s.devices, req.DeviceName)
	return
}
