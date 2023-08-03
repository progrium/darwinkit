package declparse

import (
	"fmt"
	"strings"
)

func (s Statement) String() string {
	b := &strings.Builder{}
	if s.Typedef != "" {
		b.WriteString("typedef ")
	}
	if s.Method != nil {
		b.WriteString(s.Method.String())
	}
	if s.Property != nil {
		b.WriteString(s.Property.String())
	}
	if s.Interface != nil {
		b.WriteString(s.Interface.String())
	}
	if s.Protocol != nil {
		b.WriteString(s.Protocol.String())
	}
	if s.Function != nil {
		b.WriteString(s.Function.String())
	}
	if s.Variable != nil {
		b.WriteString(s.Variable.String())
	}
	if s.Enum != nil {
		b.WriteString(s.Enum.String())
	}
	if s.Struct != nil {
		b.WriteString(s.Struct.String())
	}
	if s.Typedef != "" {
		fmt.Fprintf(b, " %s", s.Typedef)
	}
	b.WriteString(";")
	return b.String()
}

func (i ProtocolDecl) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprintf(b, "@protocol %s", i.Name)
	if i.SuperName != "" {
		_, _ = fmt.Fprintf(b, " : %s", i.SuperName)
	}
	return b.String()
}

func (i InterfaceDecl) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprintf(b, "@interface %s", i.Name)
	if i.SuperName != "" {
		_, _ = fmt.Fprintf(b, " : %s", i.SuperName)
	}
	return b.String()
}

func (p PropertyDecl) String() string {
	b := &strings.Builder{}
	b.WriteString("@property")
	var attrs []string
	for _, attr := range PropAttrs() {
		val, ok := p.Attrs[attr]
		if !ok {
			continue
		}
		if val == "" {
			attrs = append(attrs, attr.String())
		} else {
			attrs = append(attrs, fmt.Sprintf("%s=%s", attr, val))
		}
	}
	if len(attrs) != 0 {
		fmt.Fprintf(b, "(%s)", strings.Join(attrs, ", "))
	}
	b.WriteString(" ")
	if p.IsOutlet {
		b.WriteString("IBOutlet ")
	}
	typ := p.Type.String()
	b.WriteString(typ)
	if typ[len(typ)-1] != '*' {
		b.WriteString(" ")
	}
	b.WriteString(p.Name)
	return b.String()
}

func (args FuncArgs) String() string {
	if len(args) == 0 {
		return "void"
	} else if len(args) == 1 && args[0].Type.Name == "void" {
		return "void"
	} else {
		var str []string
		for _, arg := range args {
			str = append(str, strings.Trim(fmt.Sprintf("%s %s", arg.Type, arg.Name), " "))
		}
		return strings.Join(str, ", ")
	}
}

func (f FunctionDecl) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%s %s(%s", f.ReturnType, f.Ident(), f.Args)
	if f.Variadic {
		if len(f.Args) > 0 {
			b.WriteString(", ")
		}
		b.WriteString("...")
	}
	b.WriteString(")")
	return b.String()
}

func (m MethodDecl) String() string {
	b := &strings.Builder{}
	prefix := "-"
	if m.TypeMethod {
		prefix = "+"
	}
	if len(m.Args) == 0 {
		fmt.Fprintf(b, "%s (%s)%s", prefix, m.ReturnType, m.Name())
	} else {
		var parts []string
		for arg, part := range m.NameParts {
			parts = append(parts, fmt.Sprintf("%s:%s", part, m.Args[arg]))
		}
		fmt.Fprintf(b, "%s (%s)%s", prefix, m.ReturnType, strings.Join(parts, " "))
	}
	if m.Variadic {
		if len(m.Args) > 0 {
			b.WriteString(", ")
		}
		b.WriteString("...")
	}
	if m.Unavailable {
		b.WriteString(" NS_UNAVAILABLE")
	}
	return b.String()
}

func (t TypeInfo) String() string {
	if t.Func != nil {
		return t.Func.String()
	}
	params := ""
	if len(t.Params) > 0 {
		var p []string
		for _, param := range t.Params {
			p = append(p, param.String())
		}
		params = fmt.Sprintf("<%s>", strings.Join(p, ", "))
	}
	ptr := ""
	if t.IsPtr {
		ptr = "*"
	}
	str := strings.Trim(fmt.Sprintf("%s%s %s", t.Name, params, ptr), " ")
	for annot, ok := range t.Annots {
		if !ok {
			continue
		}
		str = fmt.Sprintf(annot.Format(), str)
	}
	b := &strings.Builder{}
	b.WriteString(str)
	if t.IsPtrPtr {
		if str[len(str)-1] != '*' {
			b.WriteString(" ")
		}
		b.WriteString("*")
	}
	return b.String()
}

func (arg ArgInfo) String() string {
	return fmt.Sprintf("(%s)%s", arg.Type, arg.Name)
}

func (v VariableDecl) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%s %s", v.Type, v.Name)
	if v.Value != "" {
		fmt.Fprintf(b, " = %s", v.Value)
	}
	return b.String()
}

func (e EnumDecl) String() string {
	b := &strings.Builder{}
	if e.Name != "" {
		if e.Type.Name != "" {
			fmt.Fprintf(b, "enum %s : %s { ", e.Name, e.Type.String())
		} else {
			fmt.Fprintf(b, "enum %s { ", e.Name)
		}
	} else {
		b.WriteString("enum { ")
	}
	for idx, c := range e.Cases {
		if c.Value != "" {
			fmt.Fprintf(b, "%s = %s", c.Name, c.Value)
		} else {
			fmt.Fprintf(b, "%s", c.Name)
		}
		if idx == len(e.Cases)-1 {
			b.WriteString(" ")
		} else {
			b.WriteString(", ")
		}
	}
	b.WriteString("}")
	return b.String()
}

func (s StructDecl) String() string {
	b := &strings.Builder{}
	if s.Name != "" {
		fmt.Fprintf(b, "struct %s { ", s.Name)
	} else {
		b.WriteString("struct { ")
	}
	// for idx, c := range s.Fields {
	// 	// TODO
	// }
	b.WriteString("}")
	return b.String()
}
