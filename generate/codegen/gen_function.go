package codegen

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/darwinkit/internal/set"

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

var typeMap = map[string]string{
	"*kernel.Boolean_t": "*int",
	"*kernel.UniChar":   "*uint16",
	"kernel.Boolean_t":  "int",
	"kernel.Pid_t":      "int32",
}

// GoArgs return go function args
func (f *Function) GoArgs(currentModule *modules.Module) string {
	log.Println("rendering function", f.Name)
	var args []string
	var blankArgCounter = 0
	for _, p := range f.Parameters {
		log.Println("rendering function", f.Name, p.Name, p.Type)
		log.Printf("rendering function ptype: %T", p.Type)
		// if is reserved word, add _ suffix
		if p.Name == "" {
			p.Name = fmt.Sprintf("arg%d", blankArgCounter)
			blankArgCounter++
		}
		if _, ok := reservedWords[p.Name]; ok {
			p.Name = p.Name + "_"
		}
		typ := p.Type.GoName(currentModule, true)
		if v, ok := typeMap[typ]; ok {
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
	// log.Printf("rendering GoReturn function return: %s %T", f.ReturnType, f.ReturnType)
	typ := f.ReturnType.GoName(currentModule, true)
	if v, ok := typeMap[typ]; ok {
		typ = v
	}
	return typ
}

// CArgs return go function args
func (f *Function) CArgs(currentModule *modules.Module) string {
	// log.Println("rendering function", f.Name)
	var args []string
	for _, p := range f.Parameters {
		// log.Printf("rendering cfunction arg: %s %s %T", p.Name, p.Type, p.Type)
		typ := p.Type.CName()
		if cs, ok := p.Type.(CSignatureer); ok {
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

	returnTypeStr := f.GoReturn(currentModule)

	callCode := fmt.Sprintf("C.%s(\n", f.GoName)
	var sb strings.Builder
	for _, p := range f.Parameters {
		// cast to C type
		sb.WriteString(fmt.Sprintf(cw.IndentStr+" // %T\n", p.Type))
		typ := p.Type
		switch tt := typ.(type) {
		case *typing.AliasType:
			sb.WriteString(fmt.Sprintf(cw.IndentStr+" // %T\n", tt.Type))
			if _, ok := tt.Type.(*typing.VoidPointerType); ok {
				sb.WriteString(cw.IndentStr + fmt.Sprintf("  unsafe.Pointer(%s)", p.GoName()))
			} else if _, ok := tt.Type.(*typing.ClassType); ok {
				sb.WriteString(cw.IndentStr + fmt.Sprintf("  unsafe.Pointer(%s)", p.GoName()))
			} else {
				sb.WriteString(cw.IndentStr + fmt.Sprintf("(C.%s)(%s)", tt.CName(), p.GoName()))
			}
		case *typing.RefType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  unsafe.Pointer(%s)", p.GoName()))
		case *typing.StructType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  *(*C.%s)(unsafe.Pointer(&%s))", tt.CName(), p.GoName()))
		case *typing.PrimitiveType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  C.%s(%s)", tt.CName(), p.GoName()))
		case *typing.PointerType:
			sb.WriteString(cw.IndentStr + fmt.Sprintf("  (*C.%s)(unsafe.Pointer(%s))", tt.Type.ObjcName(), p.GoName()))
		default:
			sb.WriteString(cw.IndentStr + p.GoName())
		}
		sb.WriteString(",\n")
	}
	callCode += sb.String() + cw.IndentStr + ")"

	if returnTypeStr == "" {
		cw.WriteLine(callCode)
	} else {
		var resultName = "rv"
		cw.WriteLine(resultName + " := " + callCode)
		cw.WriteLineF("// %T", f.ReturnType)
		switch tt := f.ReturnType.(type) {
		case *typing.StructType, *typing.PointerType:
			cw.WriteLineF("return *(*%s)(unsafe.Pointer(&%s))", tt.GoName(currentModule, true), resultName)
		default:
			cw.WriteLineF("return %s(%s)", returnTypeStr, resultName)
		}
	}
	cw.UnIndent()
	cw.WriteLine("}")
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
	if cs, ok := f.Type.ReturnType.(CSignatureer); ok {
		returnTypeStr = cs.CSignature()
	}
	cw.WriteLineF("%v %v(%v) {", returnTypeStr, f.GoName, f.CArgs(currentModule))
	cw.Indent()
	var args []string
	for _, p := range f.Parameters {
		args = append(args, p.Name)
	}
	cw.WriteLineF("return %v(%v);", f.Type.Name, strings.Join(args, ", "))
	cw.UnIndent()
	cw.WriteLine("}")
}

type CSignatureer interface {
	CSignature() string
}

func (f *Function) WriteCSignature(currentModule *modules.Module, cw *CodeWriter) {
	var returnTypeStr string
	rt := f.Type.ReturnType
	returnTypeStr = rt.CName()
	if v, ok := map[string]string{
		"NSInteger":  "int",
		"NSUInteger": "uint",
		"BOOL":       "bool",
	}[returnTypeStr]; ok {
		returnTypeStr = v
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
