package ast

type Program struct {
	Statements []Statement
	Errors     []*StatementError
}

func NewProgram() *Program {
	program := &Program{}
	program.Statements = make([]Statement, 0)
	program.Errors = make([]*StatementError, 0)
	return program
}
