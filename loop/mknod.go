package loop

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
)

func NewLoopDeviceFile(id int64) {
	deviceID := strconv.FormatInt(id, 10)
	deviceName := "/dev/loop" + deviceID

	cmd := exec.Command("mknod", "-m", "0660", deviceName, "b", "7", deviceID)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Printf("Executing cmd [%s]", cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Print(out.String())
		log.Fatal(err)
	}

	log.Printf("Loop device [%s] has been created", deviceName)
}
