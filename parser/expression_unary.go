package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionUnary() (expressionUnary ast.ExpressionUnary, err error) {
	operator := p.advance()
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return
	}
	expressionUnary = ast.ExpressionUnary{
		Operator:   operator,
		Expression: expression,
	}
	return
}
