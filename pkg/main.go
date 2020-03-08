package main

import (
	"fmt"
	"github.com/cambricon/cambricon-k8s-device-plugin/pkg/cndev"
	"log"
)

func main() {
	if err := cndev.Init(); err != nil {
		log.Printf("Failed to initialize CNDEV: %s.", err)

	}
	n, err := cndev.GetDeviceCount()
	if err !=nil{
		fmt.Println(err)
	}
	for i := uint(0); i < n; i++ {
		d, err := cndev.NewDeviceLite(i)
		status, err := d.Status()
		fmt.Println(status)

		processInfos := d.DeviceAllRunningProcessInfo()
		fmt.Println(processInfos)

		if err !=nil{
			fmt.Println(err)
		}
		fmt.Println(d.UUID)

	}
	

	cndev.Shutdown()
}
