package parser

func (self *Parser) scanProgram() {
	self.scanImports()

	for !self.s.Closed() {
		self.s.Next()
	}
}
