package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Xuanwo/loopd/loop"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Args is too short")
	}

	id, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	loop.NewLoopDevice(int(id))
}
