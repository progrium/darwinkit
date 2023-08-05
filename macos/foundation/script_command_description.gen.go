// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ScriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}

type _ScriptCommandDescriptionClass struct {
	objc.Class
}

type IScriptCommandDescription interface {
	objc.IObject
	IsOptionalArgumentWithName(argumentName string) bool
	TypeForArgumentWithName(argumentName string) string
	CreateCommandInstance() ScriptCommand
	CommandClassName() string
	CommandName() string
	SuiteName() string
	ArgumentNames() []string
	ReturnType() string
}

type ScriptCommandDescription struct {
	objc.Object
}

func MakeScriptCommandDescription(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](sc, objc.GetSelector("alloc"))
	return rv
}

func ScriptCommandDescription_Alloc() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.Alloc()
}

func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCommandDescription() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.New()
}

func ScriptCommandDescription_New() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.New()
}

func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](s_, objc.GetSelector("init"))
	return rv
}

func ScriptCommandDescription_Init() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.Alloc().Init()
}

func (s_ ScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isOptionalArgumentWithName:"), argumentName)
	return rv
}

func (s_ ScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("typeForArgumentWithName:"), argumentName)
	return rv
}

func (s_ ScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](s_, objc.GetSelector("createCommandInstance"))
	return rv
}

func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("commandClassName"))
	return rv
}

func (s_ ScriptCommandDescription) CommandName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("commandName"))
	return rv
}

func (s_ ScriptCommandDescription) SuiteName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("suiteName"))
	return rv
}

func (s_ ScriptCommandDescription) ArgumentNames() []string {
	rv := objc.CallMethod[[]string](s_, objc.GetSelector("argumentNames"))
	return rv
}

func (s_ ScriptCommandDescription) ReturnType() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("returnType"))
	return rv
}
