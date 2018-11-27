// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"runtime"

	"github.com/bborbe/now"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	now.DefaultLocations.Write(os.Stdout)
}
