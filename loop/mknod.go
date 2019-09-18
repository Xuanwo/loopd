package loop

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
)

// NewLoopDevice will create a new special file.
func NewLoopDevice(deviceName string) (err error) {
	_, err = os.Stat(deviceName)
	if err == nil {
		return nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// "/dev/loopXXXX"
	cmd := exec.Command("mknod", "-m", "0660", deviceName, "b", "7", deviceName[9:])

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Print(out.String())
		return
	}

	log.Printf("Loop device [%s] has been created", deviceName)
	return nil
}
