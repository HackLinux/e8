package reporter

type Error struct {
	Line, Col int
	E         error
}

func (self *Error) Error() string {
	return self.E.Error()
}
