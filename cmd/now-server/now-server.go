package main

import (
	"fmt"
	"net/http"
	"runtime"
	flag "github.com/bborbe/flagenv"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"github.com/bborbe/now"
	"time"
)

const (
	DEFAULT_PORT   int = 8080
	PARAMETER_PORT     = "port"
)

var (
	portPtr = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "port")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}
}

func do() error {
	return gracehttp.Serve(
		&http.Server{
			Addr: fmt.Sprintf(":%d", *portPtr),
			Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				n := now.Now{
					Locations: []*time.Location{
						time.FixedZone("Atlantic/Azores", 0),
						time.FixedZone("UTC", 0),
						time.FixedZone("Atlantic/Canary", 3600),
						time.FixedZone("Europe/Berlin", 7200),
					},
				}
				resp.Header().Set("Content-Type", "text/plain/json")
				n.Write(resp)
			}),
		},
	)
}
