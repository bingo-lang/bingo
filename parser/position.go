package parser

func (p *Parser) advance() {
	p.token = p.source.NextToken()
}
