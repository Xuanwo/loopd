package main

import (
	"log"

	"github.com/Xuanwo/loopd/loop"
)

func main() {
	deviceName, err := loop.NextFreeDevice()
	if err != nil {
		log.Panic(err)
	}

	err = loop.NewLoopDevice(deviceName)
	if err != nil {
		log.Panic(err)
	}
}
