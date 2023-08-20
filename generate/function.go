package generate

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

func (db *Generator) ToFunctionGen(fw string, sym Symbol) *codegen.Function {
	if db.genCache == nil {
		db.genCache = make(map[string]codegen.CodeGen)
	}
	key := fmt.Sprintf("%s.%s", fw, sym.Name)
	fg, ok := db.genCache[key]
	fmt.Printf("function: %s\n", key, fg, ok)

	type_ := &typing.FunctionType{
		Name:   sym.Name,
		GName:  modules.TrimPrefix(sym.Name),
		Module: modules.Get(fw),
	}
	functionGen := &codegen.Function{
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        type_,
	}

	db.genCache[key] = functionGen
	return functionGen

}
