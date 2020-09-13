package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/token"
)

var unaryOperators = []token.Type{
	token.PLUS, token.MINUS,
	token.BANG,
}

var binaryOperators = []token.Type{
	token.PLUS, token.MINUS, token.ASTERISK, token.SLASH,
	token.GT, token.GTE, token.LT, token.LTE, token.EQUAL,
	token.OR, token.AND,
}

func (p *Parser) checkIsStatementSeparator() bool {
	if p.tokenIsStatementSeparator() {
		p.advance()
		return true
	} else {
		for ; !p.tokenIsStatementSeparator(); p.advance() {
		}
		p.advance()
		return false
	}
}

func (p *Parser) assertTokenIsAssign() (token.Token, error) {
	return p.assertTokenIs(token.ASSIGN)
}

func (p *Parser) assertTokenIsIdentifier() (token.Token, error) {
	return p.assertTokenIs(token.IDENTIFIER)
}

func (p *Parser) assertTokenIsLet() (token.Token, error) {
	return p.assertTokenIs(token.LET)
}

func (p *Parser) assertTokenIsBinaryOperator() (token.Token, error) {
	return p.assertTokenIs(binaryOperators...)

}

func (p *Parser) assertTokenIsUnaryOperator() (token.Token, error) {
	return p.assertTokenIs(unaryOperators...)
}

func (p *Parser) assertTokenIsBoolean() (token.Token, error) {
	return p.assertTokenIs(token.BOOLEAN)
}

func (p *Parser) assertTokenIsInteger() (token.Token, error) {
	return p.assertTokenIs(token.INTEGER)
}

func (p *Parser) assertTokenIsLParen() (token.Token, error) {
	return p.assertTokenIs(token.LPAREN)
}

func (p *Parser) assertTokenIsRParen() (token.Token, error) {
	return p.assertTokenIs(token.RPAREN)
}

func (p *Parser) tokenIsLet() bool {
	return p.tokenIs(token.LET)
}

func (p *Parser) tokenIsBinaryOperator() bool {
	return p.tokenIs(binaryOperators...)
}

func (p *Parser) tokenIsUnaryOperator() bool {
	return p.tokenIs(unaryOperators...)
}

func (p *Parser) tokenIsStatementSeparator() bool {
	return p.tokenIsEOF() || p.tokenIs(token.SEMICOLON)
}

func (p *Parser) tokenIsInteger() bool {
	return p.tokenIs(token.INTEGER)
}

func (p *Parser) tokenIsBoolean() bool {
	return p.tokenIs(token.BOOLEAN)
}

func (p *Parser) tokenIsLParen() bool {
	return p.tokenIs(token.LPAREN)
}

func (p *Parser) tokenIsRParen() bool {
	return p.tokenIs(token.RPAREN)
}

func (p *Parser) tokenIsEOF() bool {
	return p.tokenIs(token.EOF)
}

func (p *Parser) assertTokenIs(types ...token.Type) (tok token.Token, err error) {
	if p.tokenIs(types...) {
		tok = p.advance()
	} else {
		err = fmt.Errorf("[SyntaxError] Expecting token to be %s got %s instead", types, p.token.Type)
	}
	return
}

func (p *Parser) tokenIs(types ...token.Type) bool {
	for _, t := range types {
		if p.token.Type == t {
			return true
		}
	}
	return false
}
