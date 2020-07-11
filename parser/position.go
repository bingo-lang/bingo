package parser

import (
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) advance() (old token.Token) {
	old, p.token = p.token, p.source.ScanToken()
	return old
}
