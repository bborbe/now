// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/now"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
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
				now.DefaultLocations.Write(resp)
			}),
		},
	)
}
