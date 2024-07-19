package main

import (
	"github.com/atomicptr/tmplr/pkg/cli"
)

func main() {
	err := cli.Run()
	if err != nil {
		panic(err)
	}
}
