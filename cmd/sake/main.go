package main

import (
	"flag"
	// "fmt"

	"github.com/h8liu/e8/stay/module"
)

var (
	forceRebuild = flag.Bool("B", false, "rebuild all")
)

func main() {
	flag.Parse()

	list, _ := module.ListModules()
	maker := module.NewMaker()
	maker.ForceRebuild = *forceRebuild

	for _, m := range list {
		_, e := maker.Open(m)
		if e != nil {
			panic(e)
		}
	}
}
