package parser

func (p *Parser) advance() {
	p.buffer = p.source.NextToken()
}
