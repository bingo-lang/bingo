package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/scanner"
	"github.com/bingo-lang/bingo/token"
	"go.uber.org/multierr"
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

func (p *Parser) ParseProgram() (program ast.Program, err error) {
	program = ast.NewProgram()
	for !p.tokenIsEOF() {
		if statement, statementError := p.parseStatement(); statementError == nil {
			program.Statements = append(program.Statements, statement)
		} else {
			err = multierr.Combine(err, statementError)
		}
	}
	return
}

func (p *Parser) IsEOF() bool {
	return p.tokenIsEOF()
}
