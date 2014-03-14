package main

import (
	"fmt"

	"github.com/h8liu/e8/stay/module"
)

func main() {
	fmt.Println(module.StayPath())
	
	/*
	// scan all the modules and return in dependency order
	module.ScanModules()

	// Scan a module and return the modules based on dependency order
	// 
	mods, e := module.ScanModule("m")

	for _, mod := range mods {
		// Check if the module signature changed
		// the signature is a hash of the list of files with 
		// the files last update time
		if mod.Changed() {
			mod.Compile()
		}
	}



	*/
}
