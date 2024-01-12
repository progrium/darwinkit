package codegen

import (
	"fmt"
	"strings"

	"github.com/progrium/macdriver/internal/set"
	"github.com/progrium/macdriver/internal/stringx"

	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

// Method is code generator for objective-c method
type Method struct {
	Name         string // the first part of objc method name
	GoName       string
	Params       []*Param
	ReturnType   typing.Type
	ClassMethod  bool // true if is class method
	WeakProperty bool // if is a weak property setter
	Deprecated   bool // if has been deprecated
	Required     bool // If this method is required. only for protocol method.
	InitMethod   bool // method that return instancetype
	Suffix       bool // GoName conflicts so add suffix to this method
	Variadic     bool
	Description  string
	DocURL       string
	Protocol     bool

	goFuncName string
	identifier string
}

func (m *Method) needRelease() bool {
	switch m.ReturnType.(type) {
	case *typing.PrimitiveType, *typing.StringType:
		return false
	}
	//!strings.HasPrefix(m.Name, "init") && strings.HasPrefix(m.Name, "Initial")
	return m.Name == "new" || m.Name == "copy" || m.Name == "mutableCopy"
}

// Selector return full Objc function name
func (m *Method) Selector() string {
	if m.identifier == "" {
		var sb strings.Builder
		sb.WriteString(m.Name)
		for idx, p := range m.Params {
			if idx > 0 {
				sb.WriteString(p.FieldName)
			}
			sb.WriteString(":")
		}
		m.identifier = sb.String()
	}
	return m.identifier
}

func (m *Method) String() string {
	return m.Selector()
}

// NormalizeInstanceTypeMethod return new init method.
func (m *Method) NormalizeInstanceTypeMethod(returnType *typing.ClassType) *Method {
	nm := &Method{
		Name:        m.Name,
		GoName:      m.GoName,
		Params:      m.Params,
		ReturnType:  returnType,
		ClassMethod: m.ClassMethod,
		InitMethod:  m.InitMethod,
		goFuncName:  m.goFuncName,
		Suffix:      m.Suffix,
		Variadic:    m.Variadic,
		Protocol:    m.Protocol,
	}
	return nm
}

// WriteGoCallCode generate go method code to call c wrapper code
func (m *Method) WriteGoCallCode(currentModule *modules.Module, typeName string, cw *CodeWriter) {
	funcDeclare := m.GoFuncDeclare(currentModule, typeName)

	if m.Deprecated {
		return
		cw.WriteLine("// deprecated")
	}

	if m.DocURL != "" {
		cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", m.Description))
		cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", m.DocURL))
	}

	var receiver string
	if m.ClassMethod {
		receiver = strings.ToLower(typeName[0:1]) + "c"
		cw.WriteLine("func (" + receiver + " _" + typeName + "Class) " + funcDeclare + " {")
	} else {
		receiver = strings.ToLower(typeName[0:1] + "_")
		cw.WriteLine("func (" + receiver + " " + typeName + ") " + funcDeclare + " {")
	}

	cw.Indent()

	var returnTypeStr string
	rt := typing.UnwrapAlias(m.ReturnType)
	switch rt.(type) {
	case *typing.VoidType:
		returnTypeStr = "objc.Void"
	default:
		returnTypeStr = m.ReturnType.GoName(currentModule, true)
	}
	callCode := fmt.Sprintf("objc.Call[%s](%s, objc.Sel(\"%s\")", returnTypeStr, receiver, m.Selector())
	var args []string
	for idx, p := range m.Params {
		switch tt := p.Type.(type) {
		case *typing.ClassType:
			args = append(args, p.GoName())
		case *typing.ProtocolType:
			pvar := fmt.Sprintf("po%d", idx)
			cw.WriteLineF("%s := objc.WrapAsProtocol(\"%s\", %s)", pvar, tt.Name, p.GoName())
			if m.WeakProperty { // weak property setter
				cw.WriteLineF("objc.SetAssociatedObject(%s, objc.AssociationKey(\"%s\"), %s, objc.ASSOCIATION_RETAIN)",
					receiver, m.GoName, pvar)
			}
			args = append(args, pvar)
		case *typing.PointerType:
			switch tt.Type.(type) {
			case *typing.ClassType: //object pointer convert to unsafe.Pointer, avoiding ffi treat it as PointerHolder
				args = append(args, fmt.Sprintf("unsafe.Pointer(%s)", p.GoName()))
			default:
				args = append(args, p.GoName())
			}
		default:
			args = append(args, p.GoName())
		}
	}
	if m.Variadic {
		callCode += fmt.Sprintf(", append([]any{%s}, args...)...)", strings.Join(args, ", "))
	} else {
		callCode += fmt.Sprintf(", %s)", strings.Join(args, ", "))
	}

	switch rt.(type) {
	case *typing.VoidType:
		cw.WriteLine(callCode)
	default:
		var resultName = "rv"
		cw.WriteLine(resultName + " := " + callCode)
		if m.needRelease() {
			cw.WriteLine(resultName + ".Autorelease()")
		}
		cw.WriteLine("return " + resultName)
	}
	cw.UnIndent()
	cw.WriteLine("}")
}

// WriteGoInterfaceCode generate go interface method signature code
func (m *Method) WriteGoInterfaceCode(currentModule *modules.Module, classType *typing.ClassType, w *CodeWriter) {
	if m.ClassMethod {
		return
	}
	if m.Deprecated {
		return
		w.WriteLine("// deprecated")
	}
	funcDeclare := m.GoFuncDeclare(currentModule, classType.GName)
	w.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (m *Method) GoFuncDeclare(currentModule *modules.Module, goTypeName string) string {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, m.Protocol))
	}
	if m.Variadic {
		paramStrs = append(paramStrs, "args ...any")
	}

	var returnType = m.ReturnType.GoName(currentModule, true)
	return m.GoFuncName() + "(" + strings.Join(paramStrs, ", ") + ")" + " " + returnType
}

// GoFuncName return go func name
func (m *Method) GoFuncName() string {
	if m.goFuncName == "" {
		var sb strings.Builder
		name := m.GoName
		if len(m.Params) == 0 {
			sb.WriteString(stringx.Capitalize(name))
		}

		for _, p := range m.Params {
			sb.WriteString(stringx.Capitalize(p.FieldName))
			if p.Object {
				sb.WriteString("Object")
			}
		}

		m.goFuncName = sb.String()
	}
	if m.Suffix || m.goFuncName == "Object" {
		return m.goFuncName + "_"
	}
	return m.goFuncName
}

// ProtocolGoFuncFieldType generate go function declaration for protocol struct impl field
func (m *Method) ProtocolGoFuncFieldType(currentModule *modules.Module) string {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, true))
	}
	if m.Variadic {
		paramStrs = append(paramStrs, "args ...any")
	}

	return "(" + strings.Join(paramStrs, ", ") + ")" + " " + m.ReturnType.GoName(currentModule, true)
}

// ProtocolGoFuncName return go protocol func name
func (m *Method) ProtocolGoFuncName() string {
	if m.goFuncName == "" {
		var sb strings.Builder
		sb.WriteString(stringx.Capitalize(m.Name))
		for idx, p := range m.Params {
			if idx == 0 {
				continue
			}
			sb.WriteString(stringx.Capitalize(p.FieldName))
			if p.Object {
				sb.WriteString("Object")
			}
		}

		m.goFuncName = sb.String()
	}
	if m.Suffix {
		return m.goFuncName + "_"
	}
	return m.goFuncName
}

// GoImports return all imports for go file
func (m *Method) GoImports() set.Set[string] {
	var imports = set.New("github.com/progrium/macdriver/objc")
	for _, param := range m.Params {
		imports.AddSet(param.Type.GoImports())
	}
	if m.ReturnType != nil {
		imports.AddSet(m.ReturnType.GoImports())
	}
	return imports
}

func (m *Method) HasProtocolParam() bool {
	for _, p := range m.Params {
		switch p.Type.(type) {
		case *typing.ProtocolType:
			return true
		}
	}
	return false
}

func (m *Method) ToProtocolParamAsObjectMethod() *Method {
	var newParams = make([]*Param, len(m.Params))
	for i, p := range m.Params {
		switch p.Type.(type) {
		case *typing.ProtocolType:
			newParams[i] = &Param{
				Name:      p.Name,
				Type:      typing.Object,
				FieldName: p.FieldName,
				Object:    true,
			}
		default:
			newParams[i] = p
		}
	}
	return &Method{
		Name:         m.Name,
		GoName:       m.GoName,
		Params:       newParams,
		Suffix:       m.Suffix,
		ReturnType:   m.ReturnType,
		ClassMethod:  m.ClassMethod,
		WeakProperty: m.WeakProperty,
		Deprecated:   m.Deprecated,
		Required:     m.Required,
		Description:  m.Description,
		DocURL:       m.DocURL,
		Variadic:     m.Variadic,
		Protocol:     m.Protocol,
	}
}
