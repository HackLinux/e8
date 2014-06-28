package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"e8vm.net/e8/img"
	"e8vm.net/e8/inst"
	"e8vm.net/e8/mem"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("todo: intro")
		return
	}

	cmd := args[1]
	args = args[2:]

	switch cmd {
	case "dasm":
		mainDasm(args)
	default:
		fmt.Fprintf(os.Stderr, "e8: unknown subcommand %q.\n", cmd)
		fmt.Fprintf(os.Stderr, "Run 'e8 help' for usage.\n")
	}

	// we need several sub commands here
	/*
		we need a bunch of sub command here
		- help
		- asm
		- dasm
		- run
	*/
}

func printError(e error) {
	fmt.Fprintf(os.Stderr, "error: %s", e)
}

func isCode(s uint32) bool {
	return s >= mem.SegCode && s-mem.SegCode < mem.SegSize
}

func mainDasm(args []string) {
	for _, f := range args {
		fin, e := os.Open(f)
		if e != nil {
			printError(e)
			continue
		}

		for {
			header, bytes, e := img.Read(fin)
			if e == io.EOF {
				break
			}
			if e != nil {
				printError(e)
				break
			}

			start := header.Start()

			if isCode(start) {
				dumpCodeSeg(start, bytes)
			} else {
				dumpDataSeg(start, bytes)
			}
		}

		fin.Close()
	}
}

func makeInst(buf []byte) inst.Inst {
	ret := uint32(buf[0])
	ret |= uint32(buf[1]) << 8
	ret |= uint32(buf[2]) << 16
	ret |= uint32(buf[3]) << 24

	return inst.Inst(ret)
}

func dumpCodeSeg(start uint32, b []byte) {
	n := len(b)
	fmt.Printf("[code] // %08x - %08x, %d bytes\n",
		start, start+uint32(n), n,
	)

	for i := 0; i < n; i += 4 {
		fmt.Printf("%04x:%04x :", uint16(i>>16), uint16(i))

		b := b[i : i+4]
		for _, bt := range b {
			fmt.Printf(" %02x", bt)
		}

		if len(b) != 4 {
			line := makeInst(b[i : i+4])
			fmt.Printf(" // %s", line.String())
		}

		fmt.Println()
	}
	fmt.Println()
}

func dumpDataSeg(start uint32, b []byte) {
	n := len(b)
	fmt.Println("[data] // %08x - %08x, %d bytes\n",
		start, start+uint32(n), n,
	)

	dumper := hex.Dumper(os.Stdout)
	dumper.Write(b)
	dumper.Close()
	fmt.Println()
}
