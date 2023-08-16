package generate

import (
	"fmt"
	"log"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

func (db *Generator) ToProtocolGen(fw string, sym Symbol) *codegen.Protocol {
	if db.genCache == nil {
		db.genCache = make(map[string]codegen.CodeGen)
	}
	key := fmt.Sprintf("%s.%s", fw, sym.Name)
	cg, ok := db.genCache[key]
	if ok {
		if pcg, ok := cg.Copy().(*codegen.Protocol); ok {
			return pcg
		}
		log.Printf("protocol overlap with class: %s\n", sym.Name)
		return nil
	}

	type_ := &typing.ProtocolType{
		Name:   sym.Name,
		GName:  modules.TrimPrefix(sym.Name),
		Module: modules.Get(fw),
	}
	protocolGen := &codegen.Protocol{
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        type_,
	}

	// symbolsdb doesnt have protocol superclasses yet,
	// so these are known ones for appkit for now
	knownSupers := map[string]string{
		"NSComboBoxDelegate":    "NSTextFieldDelegate",
		"NSMatrixDelegate":      "NSControlTextEditingDelegate",
		"NSOutlineViewDelegate": "NSControlTextEditingDelegate",
		"NSSearchFieldDelegate": "NSTextFieldDelegate",
		"NSTableViewDelegate":   "NSControlTextEditingDelegate",
		"NSTextFieldDelegate":   "NSControlTextEditingDelegate",
		"NSTextViewDelegate":    "NSTextDelegate",
		"NSTokenFieldDelegate":  "NSTextFieldDelegate",
	}
	if supername, ok := knownSupers[sym.Name]; ok {
		supersym := db.FindTypeSymbol(supername)
		if supersym == nil {
			log.Printf("protocol: cannot find super '%s' for '%s'\n", supername, sym.Name)
			return nil
		}
		protocolGen.Supers = append(protocolGen.Supers, db.ToProtocolGen(fw, *supersym))
	}

	protocolGen.Properties, protocolGen.Methods = db.Members(fw, sym, nil)

	db.genCache[key] = protocolGen
	return protocolGen

}
