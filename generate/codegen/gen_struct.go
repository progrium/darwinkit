package codegen

import (
	"github.com/progrium/darwinkit/generate/typing"
)

// Struct is code generator for objective-c struct
type Struct struct {
	Type        typing.Type
	Name        string // the first part of objc function name
	GoName      string
	Deprecated  bool // if has been deprecated
	Suffix      bool // GoName conflicts so add suffix to this function
	Description string
	DocURL      string

	goFuncName string
	identifier string
}
