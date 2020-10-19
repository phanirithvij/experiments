package main

import (
	"github.com/phanirithvij/experiments/goexps/cmd"
)

//go:generate go-bindata -o experiments/bindata.go -pkg experiments assets/

func main() {
	cmd.Execute()
}
