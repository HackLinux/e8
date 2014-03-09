package ast

type ImportDecl struct {
	As   string
	Path string
	Pos  uint32
}
