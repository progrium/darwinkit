package generate

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

func (db *Generator) ToClassGen(sym Symbol) *codegen.Class {
	if db.genCache == nil {
		db.genCache = make(map[string]codegen.CodeGen)
	}
	fw := sym.MainModule().Package
	key := fmt.Sprintf("%s.%s", fw, sym.Name)
	cg, ok := db.genCache[key]
	if ok {
		if ccg, ok := cg.Copy().(*codegen.Class); ok {
			return ccg
		} else {
			delete(db.genCache, key)
		}
	}

	// todo: generalize/improve special case for spritekit/3DNode
	if modules.TrimPrefix(sym.Name)[0] == '3' {
		sym.Name = strings.ReplaceAll(sym.Name, "3", "Three")
	}
	type_ := &typing.ClassType{
		Name:   sym.Name,
		GName:  modules.TrimPrefix(sym.Name),
		Module: modules.Get(fw),
	}
	classGen := &codegen.Class{
		Type:        type_,
		Description: sym.Description,
		DocURL:      sym.DocURL(),
	}

	st, err := sym.Parse()
	if err != nil {
		panic(err)
	}
	var covariantTypes []string
	if len(st.Interface.Params) > 0 {
		for _, p := range st.Interface.Params {
			covariantTypes = append(covariantTypes, p.Name)
		}
	}
	if st.Interface.SuperName != "" {
		if st.Interface.SuperName == "NSObject" {
			classGen.Super = &codegen.Class{
				Type: typing.Object,
			}
		} else {
			// todo: push into FindTypeSymbol
			if db.Platform == "macos" && strings.HasPrefix(st.Interface.SuperName, "UI") {
				st.Interface.SuperName = "NS" + strings.TrimPrefix(st.Interface.SuperName, "UI")
			}
			supersym := db.FindTypeSymbol(st.Interface.SuperName)
			if supersym == nil {
				log.Printf("cannot find super '%s' for '%s'\n", st.Interface.SuperName, st.Interface.Name)
				return nil
			}
			classGen.Super = db.ToClassGen(*supersym)
		}
	}

	classGen.Properties, classGen.Methods = db.Members(fw, sym, covariantTypes)

	if st.Interface.SuperName != "" {
		classGen.Methods = append(classGen.Methods, &codegen.Method{
			Name:        "alloc",
			GoName:      "alloc",
			ReturnType:  &typing.InstanceType{},
			ClassMethod: true,
		})
		classGen.Methods = append(classGen.Methods, &codegen.Method{
			Name:        "new",
			GoName:      "new",
			ReturnType:  &typing.InstanceType{},
			ClassMethod: true,
		})
		classGen.Methods = append(classGen.Methods, &codegen.Method{
			Name:       "init",
			GoName:     "init",
			ReturnType: &typing.InstanceType{},
		})
	}

	db.genCache[key] = classGen
	return classGen

}
