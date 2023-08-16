package declparse

func parseFunction(p *Parser) (next stateFn, node Node, err error) {
	typ, err := p.expectType(false)
	if err != nil {
		return nil, nil, err
	}

	decl, err := p.expectFuncType(typ, true)
	if err != nil {
		return nil, nil, err
	}
	return nil, decl, nil
}
