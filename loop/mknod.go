package loop

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/sys/unix"
)

// ref: https://github.com/torvalds/linux/blob/master/include/uapi/linux/loop.h
// #define LOOP_CTL_GET_FREE	0x4C82
func NextFreeDeviceID() (id int, err error) {
	f, err := os.OpenFile("/dev/loop-control", os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}

	x, err := unix.IoctlGetInt(int(f.Fd()), 0x4C82)
	if err != nil {
		log.Fatal(err)
	}
	return x, nil
}

func NewLoopDevice(id int) (deviceName string, err error) {
	deviceID := strconv.FormatInt(int64(id), 10)
	deviceName = "/dev/loop" + deviceID

	_, err = os.Stat(deviceName)
	if err == nil {
		return deviceName, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	cmd := exec.Command("mknod", "-m", "0660", deviceName, "b", "7", deviceID)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Print(out.String())
		log.Fatal(err)
	}

	log.Printf("Loop device [%s] has been created", deviceName)
	return deviceName, nil
}
