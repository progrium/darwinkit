package codegen

import (
	"fmt"
	"strings"

	"github.com/progrium/darwinkit/internal/set"
	"github.com/progrium/darwinkit/internal/stringx"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

// Function is code generator for objective-c function
type Function struct {
	Type        *typing.FunctionType
	Name        string // the first part of objc function name
	GoName      string
	Params      []*Param
	ReturnType  typing.Type
	Deprecated  bool // if has been deprecated
	Suffix      bool // GoName conflicts so add suffix to this function
	Description string
	DocURL      string

	goFuncName string
	identifier string
}

func (f *Function) Init() {
}

func (f *Function) WriteGoCode(cw *CodeWriter) {
	panic("implement me")
}

// Copy for copy fetcher cache value
func (f *Function) Copy() CodeGen {
	if f == nil {
		return nil
	}
	return &Function{
		Type:        f.Type,
		Name:        f.Name,
		GoName:      f.GoName,
		Params:      f.Params,
		ReturnType:  f.ReturnType,
		Deprecated:  f.Deprecated,
		Suffix:      f.Suffix,
		Description: f.Description,
		DocURL:      f.DocURL,
		goFuncName:  f.goFuncName,
		identifier:  f.identifier,
	}
}

func (m *Function) needRelease() bool {
	switch m.ReturnType.(type) {
	case *typing.PrimitiveType, *typing.StringType:
		return false
	}
	return strings.HasPrefix(m.Name, "new") || !strings.HasPrefix(m.Name, "init") && strings.HasPrefix(m.Name, "Initial") ||
		strings.HasPrefix(m.Name, "copy") || strings.HasPrefix(m.Name, "mutableCopy")
}

// Selector return full Objc function name
func (m *Function) Selector() string {
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

func (m *Function) String() string {
	return m.Selector()
}

// NormalizeInstanceTypeFunction return new init function.
func (m *Function) NormalizeInstanceTypeFunction(returnType *typing.ClassType) *Function {
	nm := &Function{
		Name:       m.Name,
		GoName:     m.GoName,
		Params:     m.Params,
		ReturnType: returnType,
		goFuncName: m.goFuncName,
		Suffix:     m.Suffix,
	}
	return nm
}

// WriteGoCallCode generate go function code to call c wrapper code
func (m *Function) WriteGoCallCode(currentModule *modules.Module, typeName string, cw *CodeWriter) {
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
	receiver = strings.ToLower(typeName[0:1] + "_")
	cw.WriteLine("func (" + receiver + " " + typeName + ") " + funcDeclare + " {")

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
	var sb strings.Builder
	for idx, p := range m.Params {
		sb.WriteString(", ")
		switch tt := p.Type.(type) {
		case *typing.ClassType:
			sb.WriteString("objc.Ptr(" + p.GoName() + ")")
		case *typing.ProtocolType:
			pvar := fmt.Sprintf("po%d", idx)
			cw.WriteLineF("%s := objc.WrapAsProtocol(\"%s\", %s)", pvar, tt.Name, p.GoName())
			sb.WriteString(pvar)
		case *typing.PointerType:
			switch tt.Type.(type) {
			case *typing.ClassType: //object pointer convert to unsafe.Pointer, avoiding ffi treat it as PointerHolder
				sb.WriteString(fmt.Sprintf("unsafe.Pointer(%s)", p.GoName()))
			default:
				sb.WriteString(p.GoName())
			}
		default:
			sb.WriteString(p.GoName())
		}
	}
	callCode += sb.String() + ")"

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

// WriteGoInterfaceCode generate go interface function signature code
func (f *Function) WriteGoInterfaceCode(currentModule *modules.Module, classType *typing.ClassType, w *CodeWriter) {
	if f.Deprecated {
		return
		w.WriteLine("// deprecated")
	}
	funcDeclare := f.GoFuncDeclare(currentModule, classType.GName)
	w.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (f *Function) GoFuncDeclare(currentModule *modules.Module, goTypeName string) string {
	var paramStrs []string
	for _, p := range f.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, false))
	}

	var returnType = f.ReturnType.GoName(currentModule, true)
	return f.GoFuncName() + "(" + strings.Join(paramStrs, ", ") + ")" + " " + returnType
}

// GoFuncName return go func name
func (f *Function) GoFuncName() string {
	if f.goFuncName == "" {
		var sb strings.Builder
		name := f.GoName
		if len(f.Params) == 0 {
			sb.WriteString(stringx.Capitalize(name))
		}

		for _, p := range f.Params {
			sb.WriteString(stringx.Capitalize(p.FieldName))
			if p.Object {
				sb.WriteString("Object")
			}
		}

		f.goFuncName = sb.String()
	}
	if f.Suffix || f.goFuncName == "Object" {
		return f.goFuncName + "_"
	}
	return f.goFuncName
}

// ProtocolGoFuncFieldType generate go function declaration for protocol struct impl field
func (f *Function) ProtocolGoFuncFieldType(currentModule *modules.Module) string {
	var paramStrs []string
	for _, p := range f.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, true))
	}

	return "(" + strings.Join(paramStrs, ", ") + ")" + " " + f.ReturnType.GoName(currentModule, false)
}

// ProtocolGoFuncName return go protocol func name
func (f *Function) ProtocolGoFuncName() string {
	if f.goFuncName == "" {
		var sb strings.Builder
		sb.WriteString(stringx.Capitalize(f.Name))
		for idx, p := range f.Params {
			if idx == 0 {
				continue
			}
			sb.WriteString(stringx.Capitalize(p.FieldName))
			if p.Object {
				sb.WriteString("Object")
			}
		}

		f.goFuncName = sb.String()
	}
	if f.Suffix {
		return f.goFuncName + "_"
	}
	return f.goFuncName
}

// GoImports return all imports for go file
func (f *Function) GoImports() set.Set[string] {
	var imports = set.New("github.com/progrium/darwinkit/objc")
	for _, param := range f.Params {
		imports.AddSet(param.Type.GoImports())
	}
	if f.ReturnType != nil {
		imports.AddSet(f.ReturnType.GoImports())
	}
	return imports
}

func (f *Function) HasProtocolParam() bool {
	for _, p := range f.Params {
		switch p.Type.(type) {
		case *typing.ProtocolType:
			return true
		}
	}
	return false
}

func (f *Function) ToProtocolParamAsObjectFunction() *Function {
	var newParams = make([]*Param, len(f.Params))
	for i, p := range f.Params {
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
	return &Function{
		Name:        f.Name,
		GoName:      f.GoName,
		Params:      newParams,
		Suffix:      f.Suffix,
		ReturnType:  f.ReturnType,
		Deprecated:  f.Deprecated,
		Description: f.Description,
		DocURL:      f.DocURL,
	}
}
