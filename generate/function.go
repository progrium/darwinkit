package generate

import (
	"fmt"
	"log"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

func (db *Generator) ToFunction(fw string, sym Symbol) *codegen.Function {
	// these functions have known declparse failures
	knownIssues := map[string]bool{}
	if knownIssues[sym.Name] {
		_, err := sym.Parse(db.Platform)
		log.Printf("skipping function %s %s because of known issue: decl='%s' err='%v'\n", fw, sym.Name, sym.Declaration, err)
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	fntyp, ok := typ.(*typing.FunctionType)
	if !ok {
		return nil
	}
	fn := &codegen.Function{
		Name:        sym.Name,
		Deprecated:  sym.Deprecated,
		GoName:      modules.TrimPrefix(sym.Name),
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        fntyp,
	}
	// temporary skip for things deprecated in 14.0
	// check if macOS platform is DeprecatedAt 14.0
	for _, p := range sym.Platforms {
		if p.Name == "macOS" && p.Deprecated {
			fn.Deprecated = true
		}
	}

	// populate params:
	log.Printf("decl: %v %s\n", sym.Name, sym.Declaration)
	for i, p := range fntyp.Parameters {
		if p.Type != nil {
			log.Printf(" param %#v: %v %+v\n", i, p.Name, p.Type.ObjcName())
		}
	}
	for _, p := range fntyp.Parameters {
		if p.Type == nil {
			fmt.Printf("skipping %s: %s because of nil type\n", sym.Name, p.Name)
			return nil
		}
		// skip pointers to ref types (for now)
		if pt, ok := p.Type.(*typing.PointerType); ok {
			if _, ok := pt.Type.(*typing.RefType); ok {
				fmt.Printf("skipping %s: %s because of pointer to ref type\n", sym.Name, p.Name)
				return nil
			}
		}
		fn.Parameters = append(fn.Parameters, &codegen.Param{
			Name: p.Name,
			Type: p.Type,
		})
	}
	// populate return type
	if fntyp.ReturnType != nil {
		fn.ReturnType = fntyp.ReturnType
	}

	return fn

}

/*
 */
