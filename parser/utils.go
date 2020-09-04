package parser

import (
	"github.com/bingo-lang/bingo/token"
)

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

func (p *Parser) assertTokenIsBinaryOperator() bool {
	if p.tokenIsBinaryOperator() {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) assertTokenIsUnaryOperator() bool {
	if p.tokenIsUnaryOperator() {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) assertTokenIsBoolean() bool {
	return p.assertTokenIs(token.BOOLEAN)
}

func (p *Parser) assertTokenIsInteger() bool {
	return p.assertTokenIs(token.INTEGER)
}

func (p *Parser) assertTokenIsLParen() bool {
	return p.assertTokenIs(token.LPAREN)
}

func (p *Parser) assertTokenIsRParen() bool {
	return p.assertTokenIs(token.RPAREN)
}

func (p *Parser) tokenIsBinaryOperator() bool {
	switch p.token.Type {
	case token.PLUS, token.MINUS, token.ASTERISK, token.SLASH,
		token.GT, token.GTE, token.LT, token.LTE, token.EQUAL,
		token.OR, token.AND:
		return true
	default:
		return false
	}
}

func (p *Parser) tokenIsUnaryOperator() bool {
	switch p.token.Type {
	case token.PLUS, token.MINUS,
		token.BANG:
		return true
	default:
		return false
	}
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

func (p *Parser) assertTokenIs(typ token.Type) bool {
	if p.tokenIs(typ) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) tokenIs(typ token.Type) bool {
	return p.token.Type == typ
}
