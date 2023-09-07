package codegen

import (
	"fmt"
	"strings"

	"github.com/progrium/darwinkit/internal/set"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

// Function is code generator for objective-c (and c) functions.
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

var reservedWords = map[string]bool{
	"func":   true,
	"map":    true,
	"new":    true,
	"var":    true,
	"len":    true,
	"copy":   true,
	"range":  true,
	"type":   true,
	"string": true,
}

// list of fixups for types that are not properly mapped
// ideally we shorten this list over time
var goTypeFixupMap = map[string]string{
	"*kernel.Boolean_t": "*int",
	"*kernel.Mode_t":    "*int",
	"*kernel.Uid_t":     "*int",
	"*kernel.Gid_t":     "*int",
	"*kernel.UniChar":   "*uint16",
	"CGFloat":           "float64",
	"kernel.Boolean_t":  "int",
	"kernel.Cpu_type_t": "int",
	"kernel.Gid_t":      "int",
	"kernel.Mode_t":     "int",
	"kernel.Pid_t":      "int32",
	"kernel.Uid_t":      "int",
	"kernel.UniChar":    "uint16",
	"uint8_t":           "byte",
}

// GoArgs return go function args
func (f *Function) GoArgs(currentModule *modules.Module) string {
	var args []string
	var blankArgCounter = 0
	for _, p := range f.Parameters {
		// if is reserved word, add _ suffix
		if p.Name == "" {
			p.Name = fmt.Sprintf("arg%d", blankArgCounter)
			blankArgCounter++
		}
		if _, ok := reservedWords[p.Name]; ok {
			p.Name = p.Name + "_"
		}
		typ := p.Type.GoName(currentModule, true)
		if v, ok := goTypeFixupMap[typ]; ok {
			typ = v
		}
		args = append(args, fmt.Sprintf("%s %s", p.Name, typ))
	}
	return strings.Join(args, ", ")
}

// GoReturn return go function return
func (f *Function) GoReturn(currentModule *modules.Module) string {
	if f.ReturnType == nil {
		return ""
	}
	typ := f.ReturnType.GoName(currentModule, true)
	if v, ok := goTypeFixupMap[typ]; ok {
		typ = v
	}
	return typ
}

// CArgs return go function args
func (f *Function) CArgs(currentModule *modules.Module) string {
	var args []string
	for _, p := range f.Parameters {
		typ := p.Type.CName()
		if cs, ok := p.Type.(hasCSignature); ok {
			typ = cs.CSignature()
		}
		// check reserved words
		if _, ok := reservedWords[p.Name]; ok {
			p.Name = p.Name + "_"
		}
		args = append(args, fmt.Sprintf("%s %s", typ, p.Name))

	}
	return strings.Join(args, ", ")
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

// WriteGoCallCode generate go function code to call c wrapper code
func (f *Function) WriteGoCallCode(currentModule *modules.Module, cw *CodeWriter) {
	funcDeclare := f.GoFuncDeclare(currentModule)

	if f.Deprecated {
		cw.WriteLine("// deprecated")
		return
	}

	if hasBlockParam(f.Parameters) {
		cw.WriteLineF("// // TODO: %v not implemented (missing block param support)", f.Name)
		return
	}

	if f.DocURL != "" {
		cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", f.Description))
		cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", f.DocURL))
	}

	cw.WriteLine("func " + funcDeclare + " {")
	cw.Indent()

	f.writeGoCallParameterPrep(currentModule, cw)

	callCode := fmt.Sprintf("C.%s(\n", f.GoName)
	var sb strings.Builder
	for _, p := range f.Parameters {
		// cast to C type
		sb.WriteString(fmt.Sprintf(cw.IndentStr+" // %T\n", p.Type))
		typ := p.Type
		switch tt := typ.(type) {
		case *typing.AliasType:
			sb.WriteString(fmt.Sprintf(cw.IndentStr+" // %T\n", tt.Type))
			sb.WriteString(cw.IndentStr + fmt.Sprintf("(C.%s)(%s)", tt.CName(), p.GoName()))
		case *typing.CStringType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  %vVal", p.GoName()))
		case *typing.RefType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  unsafe.Pointer(%s)", p.GoName()))
		case *typing.StructType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  *(*C.%s)(unsafe.Pointer(&%s))", tt.CName(), p.GoName()))
		case *typing.PrimitiveType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  C.%s(%s)", tt.CName(), p.GoName()))
		case *typing.PointerType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  (*C.%s)(unsafe.Pointer(&%s))", tt.CName(), p.GoName()))
		case *typing.DispatchType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  (*C.%s)(unsafe.Pointer(&%s))", tt.CName(), p.GoName()))
		case *typing.IDType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  %s.Ptr()", p.GoName()))
		case *typing.ClassType, *typing.ProtocolType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  unsafe.Pointer(&%s)", p.GoName()))
		default:
			sb.WriteString(cw.IndentStr + p.GoName())
		}
		sb.WriteString(",\n")
	}
	callCode += sb.String() + cw.IndentStr + ")"

	returnTypeStr := f.GoReturn(currentModule)
	if returnTypeStr == "" {
		cw.WriteLine(callCode)
	} else {
		var resultName = "rv"
		cw.WriteLine(resultName + " := " + callCode)
		cw.WriteLineF("// %T", f.ReturnType)
		switch tt := f.ReturnType.(type) {
		case *typing.StructType, *typing.PointerType:
			cw.WriteLineF("return *(*%s)(unsafe.Pointer(&%s))", tt.GoName(currentModule, true), resultName)
		case *typing.CStringType:
			cw.WriteLineF("return C.GoString(%s)", resultName)
		case *typing.ProtocolType:
			cw.WriteLineF("return %s{objc.ObjectFrom(%s)}", returnTypeStr, resultName)
		case *typing.AliasType:
			cw.WriteLineF("return *(*%s)(unsafe.Pointer(&%s))", returnTypeStr, resultName)
		default:
			cw.WriteLineF("return %s(%s)", returnTypeStr, resultName)
		}
	}
	cw.UnIndent()
	cw.WriteLine("}")
}

// writeGoCallParameterPrep generate go code to prepare parameters for c function call
func (f *Function) writeGoCallParameterPrep(currentModule *modules.Module, cw *CodeWriter) {
	for _, p := range f.Parameters {
		switch p.Type.(type) {
		default:
			continue
		case *typing.CStringType:
			cw.WriteLineF("%sVal := C.CString(%v)", p.GoName(), p.GoName())
			cw.WriteLineF("defer C.free(unsafe.Pointer(%sVal))", p.GoName())
		}
	}
}

func hasBlockParam(params []*Param) bool {
	for _, p := range params {
		if _, ok := p.Type.(*typing.BlockType); ok {
			return true
		}
		if pt, ok := p.Type.(*typing.AliasType); ok {
			t := typing.UnwrapAlias(pt.Type)
			if _, ok := t.(*typing.BlockType); ok {
				return true
			}
		}
	}
	return false
}

// WriteObjcWrapper generate objc wrapper code that maps between C and ObjC.
func (f *Function) WriteObjcWrapper(currentModule *modules.Module, cw *CodeWriter) {
	if f.Deprecated {
		return
		cw.WriteLine("// deprecated")
	}
	if hasBlockParam(f.Parameters) {
		cw.WriteLineF("// // TODO: %v not implemented (missing block param support)", f.Name)
		return
	}
	returnTypeStr := f.Type.ReturnType.CName()
	if cs, ok := f.Type.ReturnType.(hasCSignature); ok {
		returnTypeStr = cs.CSignature()
	}
	cw.WriteLineF("%v %v(%v) {", returnTypeStr, f.GoName, f.CArgs(currentModule))
	cw.Indent()
	cw.WriteLineF("return (%v)%v(", returnTypeStr, f.Type.Name)
	cw.Indent()

	for idx, p := range f.Parameters {
		cw.WriteLineF("// %T", p.Type)

		var conv string
		switch tt := p.Type.(type) {
		case *typing.PointerType:
			cw.WriteLineF("// -> %T", tt.Type)
			conv = tt.ObjcName()
		// case *typing.AliasType:
		// 	conv = tt.ObjcName() + "*"
		default:
			conv = tt.ObjcName()
		}
		// get conversion to C type
		arg := fmt.Sprintf("(%v)%v", conv, p.Name)
		if idx < len(f.Parameters)-1 {
			arg += ","
		}
		cw.WriteLineF("%v", arg)
	}
	//cw.WriteLineF("return (%v)%v(%v);", returnTypeStr, f.Type.Name, strings.Join(args, ", "))
	cw.UnIndent()
	cw.WriteLine(");")
	cw.UnIndent()
	cw.WriteLine("}")
}

type hasCSignature interface {
	CSignature() string
}

func (f *Function) WriteCSignature(currentModule *modules.Module, cw *CodeWriter) {
	var returnTypeStr string
	rt := f.Type.ReturnType
	returnTypeStr = rt.CName()
	// check for CSignature:
	if cs, ok := rt.(hasCSignature); ok {
		returnTypeStr = cs.CSignature()
	}

	if hasBlockParam(f.Parameters) {
		cw.WriteLineF("// // TODO: %v not implemented (missing block param support)", f.Name)
		return
	}
	cw.WriteLineF("// %v %v(%v); ", returnTypeStr, f.GoName, f.CArgs(currentModule))
}

// WriteGoInterfaceCode generate go interface function signature code
func (f *Function) WriteGoInterfaceCode(currentModule *modules.Module, classType *typing.ClassType, w *CodeWriter) {
	if f.Deprecated {
		return
		w.WriteLine("// deprecated")
	}
	funcDeclare := f.GoFuncDeclare(currentModule)
	w.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (f *Function) GoFuncDeclare(currentModule *modules.Module) string {
	var returnType = f.GoReturn(currentModule)
	return f.Type.GoName(currentModule, true) + "(" + f.GoArgs(currentModule) + ") " + returnType
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
