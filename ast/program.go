package ast

type Program struct {
	Statements []Statement
}

func NewProgram() Program {
	program := Program{}
	program.Statements = make([]Statement, 0)
	return program
}
