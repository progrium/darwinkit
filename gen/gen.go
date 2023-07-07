package gen

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/progrium/macschema/schema"
)

func isVoid(dt schema.DataType) bool {
	// NOTE: IBAction is a special hint type for Interface Builder, but for our
	// purposes should be `void`:
	// #define IBAction void
	// See https://nshipster.com/ibaction-iboutlet-iboutletcollection/
	if dt.Name == "IBAction" {
		return true
	}
	return dt.Name == "void" && !dt.IsPtr && !dt.IsPtrPtr
}

func msgSendFuncName(generatedNames map[string]string, cls schema.Class, selector string, isTypeMethod bool) string {
	target := "inst"
	if isTypeMethod {
		target = "type"
	}
	return fmt.Sprintf("%s_%s_%s", cls.Name, target, selectorNameToGoIdent(generatedNames, selector))
}

func selectorNameToGoIdent(generatedNames map[string]string, sel string) string {
	if ident, ok := generatedNames[sel]; ok {
		return ident
	}
	// convert to idiomatic Go name (e.g. "sendAction:to:from" -> "SendActionToFrom")
	var ident string

	// for every colon, capitalize the next letter
	// walk runes:
	// - if rune is a colon, skip it and capitalize the next letter
	// - else, add it to the ident
	capNext := true
	for _, r := range sel {
		if r == ':' {
			capNext = true
			continue
		}
		if capNext {
			ident += string(unicode.ToUpper(r))
			capNext = false
		} else {
			ident += string(r)
		}
	}
	if capNext {
		ident += "_"
	}

	if strings.HasSuffix(sel, ":") {
		trimmedSel := strings.TrimSuffix(sel, ":")
		if _, ok := generatedNames[trimmedSel]; !ok {
			ident = strings.TrimSuffix(ident, "_")
			generatedNames[sel] = ident
			return ident
		}
	}
	generatedNames[sel] = ident
	return ident
}

// Objective-C properties are syntactic sugar for getter/setter methods, this
// translates the property definition into those methods.
func propertyMethods(p schema.Property) []schema.Method {
	readonly := false
	for k, v := range p.Attrs {
		switch k {
		case "readonly":
			readonly = v.(bool)
		case "copy", "strong", "nullable", "class", "getter", "nonatomic", "assign",
			"retain", "weak":
			// no-op
		default:
			panic(unimplemented("unsupported property attr: %s", k))
		}
	}
	r := []schema.Method{
		propertyReadMethod(p),
	}
	if !readonly {
		r = append(r, propertyWriteMethod(p))
	}
	return r
}

func propertyReadMethod(p schema.Property) schema.Method {
	m := schema.Method{
		Name:        p.Name,
		Description: p.Description,
		Declaration: p.Declaration,
		Return:      p.Type,
		Args:        nil,
		Deprecated:  p.Deprecated,
		TopicURL:    p.TopicURL,
	}
	if getter, ok := p.Attrs["getter"]; ok {
		m.Name = getter.(string)
	}
	return m
}

func propertyWriteMethod(p schema.Property) schema.Method {
	return schema.Method{
		Name:        "set" + toExportedName(p.Name) + ":", // TODO rename the function?
		Description: p.Description,
		Declaration: p.Declaration,
		Return:      schema.DataType{Name: "void"},
		Args: []schema.Arg{
			{Name: "value", Type: p.Type},
		},
		Deprecated: p.Deprecated,
		TopicURL:   p.TopicURL,
	}
}

// grabbed from go/token
func toExportedName(name string) string {
	ch, n := utf8.DecodeRuneInString(name)
	if unicode.IsUpper(ch) {
		return name
	}
	ch = unicode.ToUpper(ch)
	// TODO confirm that there is actually an upper-case mapping? or we would need
	// to skip this method, or prepend a letter here to make it a valid exported
	// Go method name.
	return string(ch) + name[n:]
}

type declaration struct {
	Name string
	Base string
}

func parseClassDeclaration(decl string) (declaration, error) {
	decl = strings.TrimPrefix(decl, "@interface ")
	if decl == "NSObject" {
		return declaration{
			Name: "NSObject",
			Base: "",
		}, nil
	}
	parts := strings.Split(decl, ":")
	if len(parts) != 2 {
		//panic(fmt.Errorf("unable to parse: %q", decl))
		return declaration{}, fmt.Errorf("unable to parse: %q (%d parts)", decl, len(parts))

	}
	return declaration{
		Name: strings.TrimSpace(parts[0]),
		Base: strings.TrimSpace(parts[1]),
	}, nil
}
