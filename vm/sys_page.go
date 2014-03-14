package vm

import (
	"io"

	"github.com/h8liu/e8/vm/mem"
)

// System page is a special page that is mapped to
type SysPage struct {
	AddrError   bool  // if anything read or written to address 0-3
	Halt        bool  // if halt byte (address 4) is written
	HaltValue   uint8 // the halt value written on address 4
	StdoutError error // last IO error on flushing stdout

	stdin  chan byte
	stdout chan byte
}

// Special addresses on Sys page
const (
	// write: sets the halt value, halts the machine
	Halt = 8

	// read: if stdout is ready for output
	// write: output a byte to stdout
	Stdout = 9

	// read: if stdin is ready
	StdinReady = 10

	// read: fetch a byte if any, returns 0 if stdin is not ready
	Stdin = 11
)

var _ mem.Page = new(SysPage)

// Creates a system page
func NewSysPage() *SysPage {
	ret := new(SysPage)
	ret.stdin = make(chan byte, 32)
	ret.stdout = make(chan byte, 32)

	return ret
}

// Returns if the state is halted
func (self *SysPage) Halted() bool {
	if self == nil {
		return false
	}

	return self.Halt
}

// Clear errors, no longer halting.
func (self *SysPage) Reset() {
	if self == nil {
		return
	}
	self.AddrError = false
	self.Halt = false
}

func (self *SysPage) addrError() {
	self.AddrError = true
	self.Halt = true
	self.HaltValue = 0xff
}

// Reads a byte at address offset
func (self *SysPage) Read(offset uint32) uint8 {
	if offset < 4 {
		self.addrError()
		return 0
	}

	switch offset {
	case Stdout: // stdout ready
		if len(self.stdout) < cap(self.stdout) {
			return 0 // ready
		}
		return 1 // busy
	case StdinReady: // stdin ready
		if len(self.stdin) > 0 {
			return 0
		}
		return 1 // invalid
	case Stdin: // stdin value
		if len(self.stdin) > 0 {
			return <-self.stdin
		}
		return 0
	default:
		return 0
	}
}

// Writes a byte at address offset
func (self *SysPage) Write(offset uint32, b uint8) {
	if offset < 4 {
		self.addrError()
		return
	}

	switch offset {
	case Halt: // halt
		self.Halt = true
		self.HaltValue = b
	case Stdout: // stdout
		if len(self.stdout) < cap(self.stdout) {
			self.stdout <- b
		}
	}
}

// Flushes buffered stdout bytes to Writer
// Errors will be stored on self.IoError
func (self *SysPage) FlushStdout(w io.Writer) {
	if self == nil {
		return
	}

	for len(self.stdout) > 0 {
		b := <-self.stdout
		_, e := w.Write([]byte{b})
		if e != nil {
			self.StdoutError = e
		}
	}
}
