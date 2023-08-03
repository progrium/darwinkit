package declparse

import (
	"github.com/progrium/macdriver/generate/declparse/lexer"
)

func (p *Parser) expectType(parens bool) (ti *TypeInfo, err error) {
	ti = &TypeInfo{Annots: make(map[TypeAnnotation]bool)}

	if parens {
		if err := p.expectToken(lexer.LPAREN); err != nil {
			return nil, err
		}
	}

	// type annotations before name
	for {
		tok, _, lit := p.tb.Scan()
		if lit == "" {
			lit = tok.String()
		}

		if annot, ok := isTypeAnnot(lit); ok {
			ti.Annots[annot] = true
		} else {
			p.tb.Unscan()
			break
		}
	}

	// type name
	if ti.Name, err = p.expectIdent(); err != nil {
		return nil, err
	}
	if ti.Name == "long" {
		if _, _, lit := p.tb.Scan(); lit == "long" {
			ti.Name += " long"
		} else {
			p.tb.Unscan()
		}
	}

	// type params
	if tok, _, _ := p.tb.Scan(); tok == lexer.LT {
		p.tb.OneRuneOperators(true)

		for {
			typ, err := p.expectType(false)
			if err != nil {
				return nil, err
			}
			ti.Params = append(ti.Params, *typ)

			if tok, _, _ := p.tb.Scan(); tok != lexer.COMMA {
				p.tb.Unscan()
				break
			}
		}

		if err := p.expectToken(lexer.GT); err != nil {
			return nil, err
		}

		p.tb.OneRuneOperators(false)
	} else {
		p.tb.Unscan()
	}

	// pointer
	if tok, _, _ := p.tb.Scan(); tok == lexer.MUL {
		ti.IsPtr = true
	} else if tok == lexer.POW {
		ti.IsPtr = true
		ti.IsPtrPtr = true
	} else {
		p.tb.Unscan()
	}

	// type annotations after name
	for {
		tok, _, lit := p.tb.Scan()
		if lit == "" {
			lit = tok.String()
		}

		if annot, ok := isTypeAnnot(lit); ok {
			ti.Annots[annot] = true
		} else {
			p.tb.Unscan()
			break
		}
	}

	// detect function type
	if tok, _, _ := p.tb.Scan(); tok == lexer.LPAREN {
		p.tb.Unscan()
		ti.Func, err = p.expectFuncType(ti, false)
		if err != nil {
			return nil, err
		}
		ti.Name = ""
		ti.IsPtr = false
		ti.IsPtrPtr = false
		ti.Annots = nil
	} else {
		p.tb.Unscan()
	}

	// pointer of pointer
	if tok, _, _ := p.tb.Scan(); tok == lexer.MUL {
		ti.IsPtrPtr = true
	} else {
		p.tb.Unscan()
	}

	if parens {
		if err := p.expectToken(lexer.RPAREN); err != nil {
			return nil, err
		}
	}

	return ti, nil
}
