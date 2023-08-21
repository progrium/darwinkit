package codegen

import (
	"fmt"
	"log"
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
	Parameters  []*Param
	ReturnType  typing.Type
	Deprecated  bool // if has been deprecated
	Suffix      bool // GoName conflicts so add suffix to this function
	Description string
	DocURL      string

	goFuncName string
	identifier string
}

// GoArgs return go function args
func (f *Function) GoArgs(currentModule *modules.Module) string {
	// log.Println("rendering function", f.Name)
	var args []string
	for _, p := range f.Parameters {
		log.Println("rendering function", f.Name, p.Name, p.Type)
		log.Printf("rendering function ptype: %T", p.Type)
		if pt, ok := p.Type.(*typing.PointerType); ok {
			log.Printf("ptr type: %T", pt.Type)
		}
		args = append(args, fmt.Sprintf("%s %s", p.Name, p.Type.GoName(currentModule, true)))
	}
	return strings.Join(args, ", ")
}

// GoReturn return go function return
func (f *Function) GoReturn(currentModule *modules.Module) string {
	if f.ReturnType == nil {
		return ""
	}
	return f.ReturnType.GoName(currentModule, true)
}

// Selector return full Objc function name
func (f *Function) Selector() string {
	if f.identifier == "" {
		var sb strings.Builder
		sb.WriteString(f.Name)
		for idx, p := range f.Parameters {
			if idx > 0 {
				sb.WriteString(p.FieldName)
			}
			sb.WriteString(":")
		}
		f.identifier = sb.String()
	}
	return f.identifier
}

func (f *Function) String() string {
	return f.Selector()
}

// NormalizeInstanceTypeFunction return new init function.
func (f *Function) NormalizeInstanceTypeFunction(returnType *typing.ClassType) *Function {
	nm := &Function{
		Name:       f.Name,
		GoName:     f.GoName,
		Parameters: f.Parameters,
		ReturnType: returnType,
		goFuncName: f.goFuncName,
		Suffix:     f.Suffix,
	}
	return nm
}

// WriteGoCallCode generate go function code to call c wrapper code
func (f *Function) WriteGoCallCode(currentModule *modules.Module, typeName string, cw *CodeWriter) {
	funcDeclare := f.GoFuncDeclare(currentModule, typeName)

	if f.Deprecated {
		return
		cw.WriteLine("// deprecated")
	}

	if f.DocURL != "" {
		cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", f.Description))
		cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", f.DocURL))
	}

	var receiver string
	receiver = strings.ToLower(typeName[0:1] + "_")
	cw.WriteLine("func (" + receiver + " " + typeName + ") " + funcDeclare + " {")

	cw.Indent()

	var returnTypeStr string
	rt := typing.UnwrapAlias(f.ReturnType)
	switch rt.(type) {
	case *typing.VoidType:
		returnTypeStr = "objc.Void"
	default:
		returnTypeStr = f.ReturnType.GoName(currentModule, true)
	}
	callCode := fmt.Sprintf("objc.Call[%s](%s, objc.Sel(\"%s\")", returnTypeStr, receiver, f.Selector())
	var sb strings.Builder
	for idx, p := range f.Parameters {
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
	for _, p := range f.Parameters {
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
		if len(f.Parameters) == 0 {
			sb.WriteString(stringx.Capitalize(name))
		}

		for _, p := range f.Parameters {
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
	for _, p := range f.Parameters {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, true))
	}

	return "(" + strings.Join(paramStrs, ", ") + ")" + " " + f.ReturnType.GoName(currentModule, false)
}

// ProtocolGoFuncName return go protocol func name
func (f *Function) ProtocolGoFuncName() string {
	if f.goFuncName == "" {
		var sb strings.Builder
		sb.WriteString(stringx.Capitalize(f.Name))
		for idx, p := range f.Parameters {
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
	for _, param := range f.Parameters {
		imports.AddSet(param.Type.GoImports())
	}
	if f.ReturnType != nil {
		imports.AddSet(f.ReturnType.GoImports())
	}
	return imports
}

func (f *Function) HasProtocolParam() bool {
	for _, p := range f.Parameters {
		switch p.Type.(type) {
		case *typing.ProtocolType:
			return true
		}
	}
	return false
}

func (f *Function) ToProtocolParamAsObjectFunction() *Function {
	var newParams = make([]*Param, len(f.Parameters))
	for i, p := range f.Parameters {
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
		Parameters:  newParams,
		Suffix:      f.Suffix,
		ReturnType:  f.ReturnType,
		Deprecated:  f.Deprecated,
		Description: f.Description,
		DocURL:      f.DocURL,
	}
}
