package ir1

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
