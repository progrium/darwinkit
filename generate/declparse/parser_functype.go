package declparse

import (
	"fmt"

	"github.com/progrium/macdriver/generate/declparse/lexer"
)

func (p *Parser) expectFuncType(returnType *TypeInfo, isTopLevel bool) (fn *FunctionDecl, err error) {
	fn = &FunctionDecl{ReturnType: TypeInfo{
		Name:     returnType.Name,
		Annots:   returnType.Annots,
		IsPtr:    returnType.IsPtr,
		IsPtrPtr: returnType.IsPtrPtr,
		Params:   returnType.Params,
	}}

	var lit string
	if tok, _, _ := p.tb.Scan(); tok == lexer.LPAREN {
		tok, _, lit = p.tb.Scan()
		switch tok {
		case lexer.XOR: // ^
			fn.IsBlock = true
		case lexer.MUL: // *
			fn.IsPtr = true
		default:
			return nil, fmt.Errorf("found %q, expected ^ for block or * for func ptr", lit)
		}
	} else {
		p.tb.Unscan()
	}

	if fn.Name, err = p.expectIdent(); err != nil {
		p.tb.Unscan()
	}

	if fn.IsBlock || fn.IsPtr {
		if err := p.expectToken(lexer.RPAREN); err != nil {
			return nil, err
		}
	}

	if err := p.expectToken(lexer.LPAREN); err != nil {
		return nil, err
	}

	for {
		arg := ArgInfo{}

		// Peek at the next token to check for '...'
		if tok, _, _ := p.tb.Scan(); tok == lexer.VARARG {
			fn.Variadic = true
			break
		} else {
			p.tb.Unscan()
		}

		typ, err := p.expectType(false)
		if err != nil {
			return nil, err
		}
		arg.Type = *typ

		// If type is "void", we break the loop without appending arg
		if isTopLevel && arg.Type.Name == "void" {
			break
		}

		if arg.Name, err = p.expectIdent(); err != nil {
			p.tb.Unscan()
		}

		fn.Args = append(fn.Args, arg)

		if tok, _, _ := p.tb.Scan(); tok != lexer.COMMA {
			p.tb.Unscan()
			break
		}
	}

	if err := p.expectToken(lexer.RPAREN); err != nil {
		return nil, err
	}

	return fn, nil
}
