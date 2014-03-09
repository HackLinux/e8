package ast

type FuncDecl struct {
	Name    string
	DeclPos uint32
	Body    *BlockStmt
}
