package parser

import (
	"github.com/h8liu/e8/leaf/ast"
	t "github.com/h8liu/e8/leaf/token"
)

func (p *Parser) parseStmt() {

}

func (p *Parser) parseBlock() *ast.Block {
	p.push("block-stmt")
	defer p.pop()

	ret := new(ast.Block)

	if !p.expect(t.Lbrace) {
		return ret
	}

	for !p.ahead(t.Rbrace) {
		p.parseStmt()
	}

	p.expect(t.Rbrace)
	return ret
}
