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
	cards, err := dcmi.GetCardList()
	if err != nil {
		fmt.Println(err)
	}
	for _, card := range cards {
		device, err := dcmi.GetDeviceInCard(card)
		if err != nil {
			fmt.Println(err)
		}
		info, err := dcmi.GetDeviceHBMInfo(card)
		health, err := dcmi.GetDeviceHealth(card)
		fmt.Printf("card: %d, device: %d, MemUsage: %v, MemFree: %v, Health: %v\n", card, device, *info.Used, *info.Free, health)
	}

	dcmi.Shutdown()
}
