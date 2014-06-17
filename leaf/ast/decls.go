package ast

type Func struct {
	Name  string
	Args  []*FuncArg
	Rets  []*FuncArg
	Block *Block
}

type FuncArg struct {
	Name string
	Type Node
}
