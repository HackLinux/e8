package main

import (
	"fmt"

	"github.com/h8liu/e8/stay/module"
)

func main() {
	list, _ := module.ListModules()
	for _, m := range list {
		fmt.Println(m)
	}

	/*
		// scan all the modules and return in dependency order
		mods := module.ScanModules()

		// Scan a module and return the modules based on dependency order
		//
		mods := module.ScanModule("m")

		for _, mod := range mods {
			// Check if the module signature changed
			// the signature is a hash of the list of files with
			// the files last update time
			e = mod.Error()
			if e != nil {
				fmt.Println(e)
			}

			if mod.DepError() {
				continue
			}
			if !mod.Changed() && !mod.DepChanged() {
				continue
			}

			mod.Compile()
		}

	*/
}
