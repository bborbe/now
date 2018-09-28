package main

import (
	"github.com/bborbe/now"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	now.Write(os.Stdout)
}
