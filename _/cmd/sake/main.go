package main

import (
	"flag"
	// "fmt"

	"e8vm.net/e8/leaf/module"
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
		_, e := maker.Add(m)
		if e != nil {
			panic(e)
		}
	}
}
