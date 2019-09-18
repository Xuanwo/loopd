package loop

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// Setup will associate a loop device to an image file.
func Setup(deviceName, fileName string) (err error) {
	cmd := exec.Command("losetup", deviceName, fileName)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Print(out.String())
		return
	}

	log.Printf("Loop device [%s] has been setup to file [%s]", deviceName, fileName)
	return nil
}

// TearDown will disassociate a loop device.
func TearDown(deviceName string) (err error) {
	cmd := exec.Command("losetup", "-d", deviceName)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Print(out.String())
		return
	}

	log.Printf("Loop device [%s] has been teardown", deviceName)
	return nil
}

// NextFreeDevice will get next available device.
func NextFreeDevice() (device string, err error) {
	cmd := exec.Command("losetup", "-f")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Print(out.String())
		log.Fatal(err)
	}
	return strings.TrimSpace(out.String()), nil
}
