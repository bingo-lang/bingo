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

func (p *Parser) ParseProgram() (program ast.Program, err error) {
	program = ast.NewProgram()
	program.Statements, err = p.parseStatements()
	return
}

func (p *Parser) IsEOF() bool {
	return p.tokenIsEOF()
}
