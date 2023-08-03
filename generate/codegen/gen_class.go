package codegen

import (
	"fmt"
	"strings"

	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/set"
)

var _ CodeGen = (*Class)(nil)

// Class is code generator for objc class
type Class struct {
	Type                *typing.ClassType
	Parent              *Class
	Properties          []*Property
	InstanceTypeMethods []*Method // methods that return instance type
	Methods             []*Method

	init              bool
	methodIdentifiers set.Set[string]
}

// Copy for copy fetcher cache value
func (c *Class) Copy() CodeGen {
	if c == nil {
		return nil
	}
	var parent = c.Parent
	if parent != nil {
		parent = parent.Copy().(*Class)
	}
	return &Class{
		Type:                c.Type,
		Parent:              parent,
		Properties:          c.Properties,
		InstanceTypeMethods: c.InstanceTypeMethods,
		Methods:             c.Methods,
	}
}

func (c *Class) String() string {
	return c.Type.Name
}

func (c *Class) Init() {
	if c.init {
		return
	}
	c.methodIdentifiers = set.New[string]()
	if c.Parent != nil {
		c.Parent.Init()
		if c.Parent.Type.Name != "NSObject" {
			c.methodIdentifiers.AddSet(c.Parent.methodIdentifiers)
		}
	}
	c.init = true

	declaredMethods := c.Methods

	// properties to methods
	for _, p := range c.Properties {
		declaredMethods = append(declaredMethods, p.getter())
		if !p.ReadOnly {
			declaredMethods = append(declaredMethods, p.setter())
		}
	}

	// methods
	var finalMethods []*Method
	var finalInitMethods []*Method
	for _, m := range declaredMethods {
		if _, ok := m.ReturnType.(*typing.InstanceType); ok {
			if m.Name != "initWithCoder" {
				m.InitMethod = true
				c.InstanceTypeMethods = append(c.InstanceTypeMethods, m)
			}
		} else {
			if c.methodIdentifiers.Contains(m.Selector()) {
				continue
			}
			c.methodIdentifiers.Add(m.Selector())
			finalMethods = append(finalMethods, m)
			if m.HasProtocolParam() {
				finalMethods = append(finalMethods, m.ToProtocolParamAsObjectMethod())
			}
		}
	}

	var imSet = set.New[string]()
	for _, im := range c.InstanceTypeMethods {
		if imSet.Contains(im.Selector()) {
			continue
		}
		finalInitMethods = append(finalInitMethods, im)
		imSet.Add(im.Selector())
	}
	// parent init methods
	if c.Parent != nil {
		for _, im := range c.Parent.InstanceTypeMethods {
			if imSet.Contains(im.Selector()) {
				continue
			}
			finalInitMethods = append(finalInitMethods, im)
			imSet.Add(im.Selector())
		}
	}
	c.InstanceTypeMethods = finalInitMethods

	c.Methods = finalMethods

}

func (c *Class) GoImports() set.Set[string] {
	imports := set.New("github.com/progrium/macdriver/objc")
	if c.Parent != nil {
		imports.Add("github.com/progrium/macdriver/macos/" + c.Parent.Type.Module.Package)
	}
	for _, m := range c.InstanceTypeMethods {
		im := m.NormalizeInstanceTypeMethod(c.Type)
		imports.AddSet(im.GoImports())
	}
	for _, m := range c.Methods {
		imports.AddSet(m.GoImports())
	}
	return imports
}

func (c *Class) WriteGoCode(cw *CodeWriter) {
	cw.WriteLine("")
	c.writeClassDef(cw)
	cw.WriteLine("")
	c.writeGoInterface(cw)
	cw.WriteLine("")
	c.writeGoStruct(cw)
}

func (c *Class) writeClassDef(cw *CodeWriter) {
	cw.WriteLine(fmt.Sprintf("var %sClass = _%sClass{objc.GetClass(\"%s\")}", c.Type.GName, c.Type.GName, c.Type.Name))
	cw.WriteLine(fmt.Sprintf("type _%sClass struct {", c.Type.GName))
	cw.Indent()
	cw.WriteLine("objc.Class")
	cw.UnIndent()
	cw.WriteLine("}")
}

func (c *Class) writeGoInterface(w *CodeWriter) {
	w.WriteLine("type " + c.Type.GoInterfaceName() + " interface {")
	w.Indent()

	if c.Parent != nil {
		w.WriteLine(c.Parent.Type.GoName(c.Type.Module, false))
	}

	for _, m := range c.Methods {
		if !m.ClassMethod {
			m.WriteGoInterfaceCode(c.Type.Module, c.Type, w)
		}
	}

	// w.WriteLine("")
	// for _, m := range c.InstanceTypeMethods {
	// 	if strings.ToLower(m.GoName) == "init" {
	// 		continue
	// 	}
	// 	im := m.NormalizeInstanceTypeMethod(c.Type)
	// 	im.WriteGoInterfaceCode(c.Type.Module, c.Type, w)
	// }

	w.UnIndent()
	w.WriteLine("}")
}

func (c *Class) writeGoStruct(w *CodeWriter) {
	w.WriteLine("type " + c.Type.GoStructName() + " struct {")
	w.Indent()
	if c.Parent != nil {
		w.WriteLine(c.Parent.Type.GoName(c.Type.Module, true))
	}
	w.UnIndent()
	w.WriteLine("}")

	// make func
	w.WriteLine("")
	w.WriteLine(fmt.Sprintf("func Make%s(ptr unsafe.Pointer) %s {", c.Type.GName, c.Type.GoStructName()))
	w.Indent()
	w.WriteLines([]string{
		fmt.Sprintf("return %s{", c.Type.GoStructName()),
	})
	if c.Parent != nil {
		pType := c.Parent.Type
		w.WriteLine(fmt.Sprintf("\t%s: %s(ptr),", pType.GName, typing.PrependPackage(*pType.Module, "Make"+pType.GName, *c.Type.Module)))
	}
	w.WriteLine("}")
	w.UnIndent()
	w.WriteLine("}")

	// methods
	for _, m := range c.InstanceTypeMethods {
		w.WriteLine("")
		im := m.NormalizeInstanceTypeMethod(c.Type)
		im.WriteGoCallCode(c.Type.Module, c.Type.GoStructName(), w)
		if im.Name == "new" {
			//add a convenient Newxxx Method
			w.WriteLine("")
			w.WriteLine(fmt.Sprintf("func New%s() %s {", c.Type.GName, c.Type.GName))
			w.Indent()
			w.WriteLine(fmt.Sprintf("return %sClass.New()", c.Type.GName))
			w.UnIndent()
			w.WriteLine("}")
		}
		if im.InitMethod || im.ClassMethod {
			//add a convenient init / class method function
			w.WriteLine("")
			funcDeclare := im.GoFuncDeclare(c.Type.Module, c.Type.GoStructName())
			w.WriteLine(fmt.Sprintf("func %s_%s {", c.Type.GName, funcDeclare))
			w.Indent()
			var params []string
			for _, p := range m.Params {
				params = append(params, p.GoName())
			}
			alloc := ".Alloc()"
			if im.ClassMethod {
				alloc = ""
			}
			w.WriteLine(fmt.Sprintf("return %sClass%s.%s(%s)",
				c.Type.GName,
				alloc,
				m.GoFuncName(c.Type.GoStructName()),
				strings.Join(params, ", ")))
			w.UnIndent()
			w.WriteLine("}")
		}
	}

	for _, m := range c.Methods {
		w.WriteLine("")
		m.WriteGoCallCode(c.Type.Module, c.Type.GoStructName(), w)
		if m.ClassMethod && !m.Deprecated {
			//add a convenient init / class method function
			w.WriteLine("")
			funcDeclare := m.GoFuncDeclare(c.Type.Module, c.Type.GoStructName())
			w.WriteLine(fmt.Sprintf("func %s_%s {", c.Type.GName, funcDeclare))
			w.Indent()
			var params []string
			for _, p := range m.Params {
				params = append(params, p.GoName())
			}
			returnKeyword := "return "
			rt := typing.UnwrapAlias(m.ReturnType)
			if _, ok := rt.(*typing.VoidType); ok {
				returnKeyword = ""
			}
			w.WriteLine(fmt.Sprintf("%s%sClass.%s(%s)",
				returnKeyword,
				c.Type.GName,
				m.GoFuncName(c.Type.GoStructName()),
				strings.Join(params, ", ")))
			w.UnIndent()
			w.WriteLine("}")
		}
	}

}
