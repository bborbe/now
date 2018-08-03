package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
)

func main() {
	locations := []string{
		"Atlantic/Azores",
		"UTC",
		"Atlantic/Canary",
		"Europe/Berlin",
	}

	for _, location := range locations {
		loadLocation, err := time.LoadLocation(location)
		if err != nil {
			glog.Exitf("load location %s failed", location)
		}
		now := time.Now().In(loadLocation)
		fmt.Printf("%s %s\n", now.Format("2006-01-02T15:04:05"), location)
	}
}
