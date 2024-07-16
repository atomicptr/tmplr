package main

import (
	"github.com/atomicptr/tplr/pkg/cli"
)

func main() {
	err := cli.Run()
	if err != nil {
		panic(err)
	}
}
