package codegen

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

// AliasInfo enum type and values
type AliasInfo struct {
	typing.AliasType
	Description string
	DocURL      string
	Values      []*EnumValue
}

// IsString return if is a string type Enum
func (e *AliasInfo) IsString() bool {
	t := e.Type
	for {
		at, ok := t.(*typing.AliasType)
		if !ok {
			break
		}
		t = at.Type
	}
	_, ok := t.(*typing.StringType)
	return ok
}

// EnumValue the enum name and value
type EnumValue struct {
	Name       string          // the objc enum name
	GoName     string          // the go name of enum
	Value      string          // the value(amd64 arch)
	Arm64Value string          // the value for arm64. if is empty, use Value
	Module     *modules.Module // the module enum value defined in
}
