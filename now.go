package now

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"io"
)

var Locations = []*time.Location{
	time.FixedZone("Atlantic/Azores", 0),
	time.FixedZone("UTC", 0),
	time.FixedZone("Atlantic/Canary", 3600),
	time.FixedZone("Europe/Berlin", 7200),
}

func Write(writer io.Writer) {
	glog.V(4).Infof("write now ...")
	for _, location := range Locations {
		now := time.Now().In(location)
		fmt.Fprintf(writer, "%s %s\n", now.Format("2006-01-02T15:04:05"), location.String())
	}
}
