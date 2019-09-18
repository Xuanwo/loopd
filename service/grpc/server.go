package grpc

import (
	"context"
	"log"

	"github.com/Xuanwo/loopd/loop"
)

type Server struct {
	devices map[string]string
}

func NewServer() *Server {
	return &Server{devices: make(map[string]string)}
}

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

func (s Server) Teardown(ctx context.Context, req *TeardownRequest) (resp *EmptyResponse, err error) {
	resp = &EmptyResponse{}

	if _, ok := s.devices[req.DeviceName]; !ok {
		log.Printf("Device [%d] is not found in db.", req.DeviceName)
	}

	err = loop.TearDown(req.DeviceName)
	if err != nil {
		return
	}

	delete(s.devices, req.DeviceName)
	return
}
