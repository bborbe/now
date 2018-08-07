package now

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"io"
)

type Now struct {
	Locations []string
}

func (n *Now) Write(writer io.Writer) {
	glog.V(4).Infof("write now ...")
	for _, location := range n.Locations {
		loadLocation, err := time.LoadLocation(location)
		if err != nil {
			glog.Exitf("load location %s failed", location)
		}
		now := time.Now().In(loadLocation)
		fmt.Fprintf(writer, "%s %s\n", now.Format("2006-01-02T15:04:05"), location)
	}
}
