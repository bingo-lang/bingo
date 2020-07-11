package ast

type Program struct {
	Statements []Statement
}

func NewProgram() Program {
	program := Program{}
	program.Statements = make([]Statement, 0)
	return program
}

func (p Program) HasErrors() bool {
	for _, statement := range p.Statements {
		if statement.HasErrors() {
			return true
		}
	}
	return false
}

func (p Program) String() string {
	errors := ""
	for _, statement := range p.Statements {
		if statement.HasErrors() {
			errors += "\n" + statement.String()
		}
	}
	return errors
}
