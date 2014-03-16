package main

import (
	// "fmt"

	"github.com/h8liu/e8/stay/module"
)

func main() {
	list, _ := module.ListModules()
	maker := module.NewMaker()

	for _, m := range list {
		_, e := maker.Open(m)
		if e != nil {
			panic(e)
		}
	}
}
