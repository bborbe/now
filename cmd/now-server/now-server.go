package main

import (
	"fmt"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/now"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"net/http"
	"runtime"
)

var (
	portPtr = flag.Int("port", 8080, "port")
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
				resp.Header().Set("Content-Type", "text/plain")
				now.Write(resp)
			}),
		},
	)
}
