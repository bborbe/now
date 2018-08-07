package main

import (
	"github.com/bborbe/now"
	"os"
)

func main() {
	n := now.Now{
		Locations: []string{
			"Atlantic/Azores",
			"UTC",
			"Atlantic/Canary",
			"Europe/Berlin",
		},
	}
	n.Write(os.Stdout)
}
