package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/scanner"
	"github.com/bingo-lang/bingo/token"
	"io"
)

type Parser struct {
	source *scanner.Scanner
	token  token.Token
}

func New(source io.RuneReader) *Parser {
	s := scanner.New(source)
	parser := &Parser{source: s}
	parser.advance()
	return parser
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()
	for {
		if p.IsEOF() {
			return program
		}
		statement := p.parseStatement()
		if err, ok := statement.(*ast.StatementError); !ok {
			program.Statements = append(program.Statements, statement)
		} else {
			program.Errors = append(program.Errors, err)
		}
	}
}

func (p *Parser) IsEOF() bool {
	return p.tokenIsEOF()
}
