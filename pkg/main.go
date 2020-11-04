package main

import (
	"fmt"
	"github.com/cambricon/cambricon-k8s-device-plugin/pkg/dcmi"
	"log"
	"os"
)

func main() {
	if err := dcmi.Init(); err != nil {
		log.Printf("Failed to initialize DCMI: %s.", err)
		os.Exit(1)
	}
	n, err := dcmi.GetDeviceCount()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	dcmi.Shutdown()
}
