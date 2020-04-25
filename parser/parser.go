package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/scanner"
	"io"
)

type Parser struct {
	source *scanner.Scanner
}

func New(source io.RuneReader) *Parser {
	s := scanner.New(source)
	parser := &Parser{s}
	return parser
}

func (p *Parser) Parse() *ast.Program {
	return &ast.Program{}
}
