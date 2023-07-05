package gen

import (
	"embed"
	"io"
	"text/template"
)

//go:embed "template/*.tmpl"
var templateFS embed.FS

var templates = template.Must(template.ParseFS(templateFS, "template/*.tmpl"))

type Import struct {
	Path  string
	Alias string
}

type Primitive struct {
	ctype    string
	toCGoFmt string
}

var (
	PrimitiveULong  = Primitive{"unsigned long", "C.ulong(%s)"}
	PrimitiveLong   = Primitive{"long", "C.long(%s)"}
	PrimitiveUShort = Primitive{"unsigned short", "C.ushort(%s)"}
	PrimitiveDouble = Primitive{"double", "C.double(%s)"}
)

// PackageContents describes the Objective-C wrapper types found in a Go wrapper
// package. It's used to locate where to import a type name in the Go code, as
// well as whether it's a class, or another C type.
type PackageContents struct {
	// Import can be `nil` if this is representing the contents of the package
	// currently being generated, which don't need to be "imported".
	Import     *Import
	Classes    map[string]bool
	Primitives map[string]Primitive
	// TODO structs
}

type Arg struct {
	Name string
	Type string
}

type SelectorPart struct {
	Key string
	Arg *Arg
}

type CGoMsgSend struct {
	Name     string
	Class    string
	Selector []SelectorPart
	Return   string
}

func (m CGoMsgSend) HasReturn() bool {
	return m.Return != "void"
}

type CGoWrapperArg struct {
	Name     string
	Type     string
	ToCGoFmt string
}

type CGoWrapperReturn struct {
	Type       string
	FromCGoFmt string
}

type CGoWrapperFunc struct {
	Description string
	Name        string
	Args        []CGoWrapperArg
	Returns     []CGoWrapperReturn
}

func (f CGoWrapperFunc) HasReturn() bool {
	return len(f.Returns) > 0
}

type MethodDef struct {
	Description string
	Name        string
	WrappedFunc CGoWrapperFunc
}

type ClassDef struct {
	Name            string
	Base            string
	ClassMethods    []MethodDef
	InstanceMethods []MethodDef
}

type StructDef struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name string
	Type string
}

type PackageDescription struct {
	Name           string
	LinkFrameworks []string
	CIncludes      []string
}

type GoPackage struct {
	PackageDescription
	Imports              []Import
	ClassMsgSendWrappers []CGoMsgSend
	MsgSendWrappers      []CGoMsgSend
	CGoWrapperFuncs      []MethodDef
	Classes              []ClassDef
	Structs              []StructDef
}

func (p *GoPackage) Generate(w io.Writer) error {
	return templates.ExecuteTemplate(w, "main", p)
}
