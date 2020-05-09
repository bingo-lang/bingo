package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/scanner"
	"github.com/bingo-lang/bingo/token"
	"io"
)

type Parser struct {
	source *scanner.Scanner
	buffer token.Token
}

func New(source io.RuneReader) *Parser {
	s := scanner.New(source)
	parser := &Parser{source: s}
	return parser
}

func (p *Parser) Parse() *ast.Program {
	program := ast.New()
	for ; !p.source.EOF(); p.advance() {
		statement := p.parseStatement()
		program.Statements = append(program.Statements, statement)
	}
	return program
}
