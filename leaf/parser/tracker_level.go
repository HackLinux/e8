package parser

import (
	"github.com/h8liu/e8/leaf/lexer"
	"github.com/h8liu/e8/prt"
)

type trackLevel struct {
	name string
	subs []trackNode
}

func (self *trackLevel) add(n trackNode) {
	self.subs = append(self.subs, n)
}

func (self *trackLevel) PrintTo(p prt.Iface) {
	p.Printf("+ %s:", self.name)
	p.ShiftIn()
	for _, sub := range self.subs {
		level, isLevel := sub.(*trackLevel)
		if isLevel {
			level.PrintTo(p)
		} else {
			tok := sub.(*lexer.Token)
			p.Print("- ", tok)
		}
	}
	p.ShiftOut()
}
