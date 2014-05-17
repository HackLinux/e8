package ir1

type Expr interface {
	Type() Type
	String() string
}

type VarExpr struct {
	*Var
}

func Ve(v *Var) *VarExpr {
	if v == nil {
		panic("bug")
	}

	return &VarExpr{v}
}

func (self *VarExpr) Type() Type {
	return self.Var.Type
}

func (self *VarExpr) String() string {
	return self.Var.Name
}

type BinExpr struct {
	V1, V2 *Var
	Op Op
}

type UnaExpr struct {
	V *Var
	Op Op
}

type Op int

