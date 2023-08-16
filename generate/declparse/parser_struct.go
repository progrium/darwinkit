package declparse

import (
	"github.com/progrium/macdriver/generate/declparse/keywords"
	"github.com/progrium/macdriver/generate/declparse/lexer"
)

func parseStruct(p *Parser) (next stateFn, node Node, err error) {
	decl := &StructDecl{}

	if err := p.expectToken(keywords.STRUCT); err != nil {
		return nil, nil, err
	}

	decl.Name, err = p.expectIdent()
	if err != nil {
		p.tb.Unscan()
	}

	if err := p.expectToken(lexer.LCURLY); err != nil {
		return nil, decl, nil
	}

	if err := p.expectToken(lexer.VARARG); err != nil {
		p.tb.Unscan()

		// TODO: fields

		// for {
		// 	enum := VariableDecl{}

		// 	if enum.Name, err = p.expectIdent(); err != nil {
		// 		p.tb.Unscan()
		// 	}

		// 	if err := p.expectToken(lexer.EQ); err == nil {
		// 		tok, pos, lit := p.tb.Scan()
		// 		if tok != lexer.INTEGER && tok != lexer.DECIMAL {
		// 			return nil, nil, fmt.Errorf("found %q, expected numeric at %v", lit, pos)
		// 		}
		// 		enum.Value = lit
		// 	} else {
		// 		p.tb.Unscan()
		// 	}

		// 	decl.Cases = append(decl.Cases, enum)

		// 	if tok, _, _ := p.tb.Scan(); tok != lexer.COMMA {
		// 		p.tb.Unscan()
		// 		break
		// 	}
		// }
	}

	if err := p.expectToken(lexer.RCURLY); err != nil {
		return nil, nil, err
	}

	return nil, decl, nil
}
