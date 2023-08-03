package codegen

import "C"
import (
	"fmt"
	"strings"

	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/set"
)

var _ CodeGen = (*Protocol)(nil)

// Protocol is code generator for objc protocol.
type Protocol struct {
	Type       *typing.ProtocolType
	Parents    []*Protocol
	Methods    []*Method
	Properties []*Property

	init bool
}

// Copy implements CodeGen
func (p *Protocol) Copy() CodeGen {
	parents := make([]*Protocol, len(p.Parents))
	for i := 0; i < len(p.Parents); i++ {
		parents[i] = p.Parents[i].Copy().(*Protocol)
	}
	return &Protocol{
		Type:       p.Type,
		Parents:    p.Parents,
		Methods:    p.Methods,
		Properties: p.Properties,
	}
}

func (p *Protocol) String() string {
	return p.Type.Name
}

func (p *Protocol) Init() {
	if p.init {
		return
	}

	for _, m := range p.Methods {
		if _, ok := m.ReturnType.(*typing.InstanceType); ok {
			m.ReturnType = typing.Object
		}
	}

	for _, parent := range p.Parents {
		if parent != nil && parent.Type.Name != "NSObject" {
			parent.Init()
		}
	}

	var methodSet = set.New[string]()

	var methods []*Method
	for _, m := range p.Methods {
		if methodSet.Contains(m.Selector()) {
			continue
		}
		methodSet.Add(m.Selector())
		methods = append(methods, m)
	}
	p.Methods = methods

	var propertySet = set.New[string]()

	var properties []*Property
	for _, pp := range p.Properties {
		if propertySet.Contains(pp.Name) {
			continue
		}
		propertySet.Add(pp.Name)
		properties = append(properties, pp)
	}
	p.Properties = properties

	p.init = true
}

func (p *Protocol) GoImports() set.Set[string] {
	imports := set.New("github.com/progrium/macdriver/objc")
	for _, parent := range p.Parents {
		imports.Add("github.com/progrium/macdriver/macos/" + parent.Type.Module.Package)
	}
	for _, m := range p.Methods {
		imports.AddSet(m.GoImports())
	}
	for _, p := range p.Properties {
		imports.AddSet(p.getter().GoImports())
	}
	return imports
}

func (p *Protocol) WriteGoCode(w *CodeWriter) {
	w.WriteLine("")

	// Protocol Interface
	p.writeProtocolInterface(w)

	// Delegate Prototol impl struct
	w.WriteLine("")
	if strings.Contains(p.Type.GName, "Delegate") {
		p.writeDelegateStruct(w)
	}

	// Protocol Wrapper
	w.WriteLine("")
	p.writeProtocolWrapperStruct(w)
}

func (p *Protocol) writeProtocolInterface(w *CodeWriter) {
	w.WriteLine("type " + p.Type.GoInterfaceName() + " interface {")
	w.Indent()
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			w.WriteLine(pp.Type.GoName(p.Type.Module, false))
		}
	}
	for _, m := range p.allMethods() {
		if !m.Required {
			w.WriteLine("Implements" + m.ProtocolGoFuncName() + "() bool")
			w.WriteLine("// optional")
		} else {
			w.WriteLine("// required")
		}
		if m.Deprecated {
			w.WriteLine("// deprecated")
		}
		w.WriteLine(m.ProtocolGoFuncName() + m.ProtocolGoFuncFieldType(p.Type.Module))
	}
	w.UnIndent()
	w.WriteLine("}")
}

func (p *Protocol) writeDelegateStruct(w *CodeWriter) {
	implStructName := p.Type.GoStructName()
	w.WriteLine("type " + implStructName + " struct {")
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			w.WriteLine(pp.Type.GoStructName())
		}
	}
	w.Indent()
	for _, m := range p.allMethods() {
		w.WriteLine(fmt.Sprintf("_%s func %s", m.ProtocolGoFuncName(), m.ProtocolGoFuncFieldType(p.Type.Module)))
	}
	w.UnIndent()
	w.WriteLine("}")

	receiver := "di"
	for _, m := range p.allMethods() {
		if !m.Required {
			w.WriteLine(fmt.Sprintf("func (%s *%s) Implements%s() bool {", receiver, implStructName, m.ProtocolGoFuncName()))
			w.WriteLine(fmt.Sprintf("\t return %s._%s != nil", receiver, m.ProtocolGoFuncName()))
			w.WriteLine("}")
		}

		w.WriteLine("")
		if m.Deprecated {
			w.WriteLine("// deprecated")
		}
		w.WriteLine(fmt.Sprintf("func (%s *%s) Set%s(f func %s) {", receiver, implStructName, m.ProtocolGoFuncName(), m.ProtocolGoFuncFieldType(p.Type.Module)))
		w.WriteLine(fmt.Sprintf("\t%s._%s = f", receiver, m.ProtocolGoFuncName()))
		w.WriteLine("}")

		if m.Required {
			w.WriteLine("// required")
		}
		w.WriteLine("")
		if m.Deprecated {
			w.WriteLine("// deprecated")
		}
		w.WriteLine(fmt.Sprintf("func (%s *%s) %s%s {", receiver, implStructName, m.ProtocolGoFuncName(), m.ProtocolGoFuncFieldType(p.Type.Module)))
		var sb strings.Builder
		for i, p := range m.Params {
			if i != 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(p.GoName())
		}
		callArgs := sb.String()
		if _, ok := m.ReturnType.(*typing.VoidType); ok {
			w.WriteLine(fmt.Sprintf("\t%s._%s(%s)", receiver, m.ProtocolGoFuncName(), callArgs))
		} else {
			w.WriteLine(fmt.Sprintf("\treturn %s._%s(%s)", receiver, m.ProtocolGoFuncName(), callArgs))
		}
		w.WriteLine("}")
	}
}

func (p *Protocol) writeProtocolWrapperStruct(w *CodeWriter) {
	typeName := p.Type.GoWrapperName()
	w.WriteLine("type " + typeName + " struct {")
	w.Indent()
	if len(p.Parents) == 0 {
		w.WriteLine("objc.Object")
	}
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			w.WriteLine(pp.Type.GoName(p.Type.Module, true))
		} else {
			w.WriteLine("objc.Object")
		}
	}
	w.UnIndent()
	w.WriteLine("}")

	for _, m := range p.allMethods() {
		w.WriteLine("")
		if !m.Required {
			receiver := strings.ToLower(typeName[0:1] + "_")
			w.WriteLine(fmt.Sprintf("func (%s %s) Implements%s() bool {", receiver, typeName, m.ProtocolGoFuncName()))
			w.WriteLine(fmt.Sprintf("\t return %s.RespondsToSelector(objc.GetSelector(\"%s\"))", receiver, m.Selector()))
			w.WriteLine("}")
		}
		w.WriteLine("")
		m.WriteGoCallCode(p.Type.Module, typeName, w)
	}

}

// include property getter/setter
func (p *Protocol) allMethods() []*Method {
	var allMethods = make([]*Method, len(p.Methods))
	copy(allMethods, p.Methods)
	for _, pp := range p.Properties {
		if !pp.ReadOnly {
			allMethods = append(allMethods, (*Method)(pp.setter()))
		}

		allMethods = append(allMethods, (*Method)(pp.getter()))
	}
	return allMethods
}
