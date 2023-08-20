package generate

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

func (db *Generator) ToFunction(fw string, sym Symbol) *codegen.Function {
	typ := db.TypeFromSymbol(sym)
	fmt.Println("typ:", typ)
	type_ := &typing.FunctionType{
		Name:   sym.Name,
		GName:  modules.TrimPrefix(sym.Name),
		Module: modules.Get(fw),
	}
	fn := &codegen.Function{
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        type_,
	}
	for _, arg := range sym.Parameters {
		fn.Params = append(fn.Params, &codegen.Param{
			Name: arg.Name,
		})
	}
	return fn

}
