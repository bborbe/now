package now

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"io"
)

type Now struct {
	Locations []*time.Location
}

func (n *Now) Write(writer io.Writer) {
	glog.V(4).Infof("write now ...")
	for _, location := range n.Locations {
		now := time.Now().In(location)
		fmt.Fprintf(writer, "%s %s\n", now.Format("2006-01-02T15:04:05"), location.String())
	}
}
