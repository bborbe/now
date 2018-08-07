package main

import (
	"github.com/bborbe/now"
	"os"
	"time"
	"github.com/golang/glog"
)

func main() {
	names := []string{
		"Atlantic/Azores",
		"UTC",
		"Atlantic/Canary",
		"Europe/Berlin",
	}
	var locations []*time.Location
	for _, name := range names {
		location, err := time.LoadLocation(name)
		if err != nil {
			glog.Exitf("location not found %s", name)
		}
		locations = append(locations, location)

	}
	n := now.Now{
		Locations: locations,
	}
	n.Write(os.Stdout)
}
