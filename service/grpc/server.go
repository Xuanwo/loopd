package grpc

import (
	"context"
	"log"
	"strconv"

	"github.com/Xuanwo/loopd/loop"
)

type Server struct {
	devices map[int]string
}

func NewServer() *Server {
	return &Server{devices: make(map[int]string)}
}

func (s Server) Setup(ctx context.Context, req *SetupRequest) (resp *SetupResponse, err error) {
	resp = &SetupResponse{}

	id, err := loop.NextFreeDeviceID()
	if err != nil {
		return
	}

	deviceName, err := loop.NewLoopDevice(id)
	if err != nil {
		return
	}

	err = loop.Setup(deviceName, req.Filename)
	if err != nil {
		return
	}

	s.devices[id] = req.Filename
	return
}

func (s Server) Teardown(ctx context.Context, req *TeardownRequest) (resp *EmptyResponse, err error) {
	resp = &EmptyResponse{}

	if _, ok := s.devices[int(req.Id)]; !ok {
		log.Printf("Device [%d] is not found in db.", req.Id)
	}

	deviceName := "/dev/loop" + strconv.Itoa(int(req.Id))

	err = loop.TearDown(deviceName)
	if err != nil {
		return
	}

	delete(s.devices, int(req.Id))
	return
}
