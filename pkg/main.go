package main

import (
	"fmt"
	"github.com/cambricon/cambricon-k8s-device-plugin/pkg/cndev"
	"log"
	"os"
)

func main() {
	if err := cndev.Init(); err != nil {
		log.Printf("Failed to initialize CNDEV: %s.", err)
		os.Exit(1)
	}
	n, err := cndev.GetDeviceCount()
	if err != nil {
		fmt.Println(err)
	}
	for i := uint(0); i < n; i++ {
		d, err := cndev.NewDeviceLite(i)
		status, err := d.Status()
		fmt.Printf("(%v/%v)/%v\n", *status.Memory.Used, *status.Memory.Free, *status.Utilization.Board)

		processInfos := d.DeviceAllRunningProcessInfo()
		for i := 0; i < len(processInfos); i++ {
			fmt.Printf("(%v/%v/%v)\n", processInfos[i].Pid, processInfos[i].PhysicalMemoryUsed, processInfos[i].Command)
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(d.UUID)

	}

	cndev.Shutdown()
}
