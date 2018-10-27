package main

import (
	"os"
	"runtime"

	"github.com/bborbe/now"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	now.Write(os.Stdout)
}
