// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ScriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}

type _ScriptClassDescriptionClass struct {
	objc.Class
}

type IScriptClassDescription interface {
	IClassDescription
	ClassDescriptionForKey(key string) ScriptClassDescription
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	HasOrderedToManyRelationshipForKey(key string) bool
	HasPropertyForKey(key string) bool
	HasReadablePropertyForKey(key string) bool
	HasWritablePropertyForKey(key string) bool
	TypeForKey(key string) string
	SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector
	SupportsCommand(commandDescription IScriptCommandDescription) bool
	SuperclassDescription() ScriptClassDescription
	ClassName() string
	DefaultSubcontainerAttributeKey() string
	ImplementationClassName() string
	SuiteName() string
}

type ScriptClassDescription struct {
	ClassDescription
}

func MakeScriptClassDescription(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: MakeClassDescription(ptr),
	}
}

func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](sc, objc.GetSelector("alloc"))
	return rv
}

func ScriptClassDescription_Alloc() ScriptClassDescription {
	return ScriptClassDescriptionClass.Alloc()
}

func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScriptClassDescription() ScriptClassDescription {
	return ScriptClassDescriptionClass.New()
}

func ScriptClassDescription_New() ScriptClassDescription {
	return ScriptClassDescriptionClass.New()
}

func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.GetSelector("init"))
	return rv
}

func ScriptClassDescription_Init() ScriptClassDescription {
	return ScriptClassDescriptionClass.Alloc().Init()
}

func (s_ ScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.GetSelector("classDescriptionForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isLocationRequiredToCreateForKey:"), toManyRelationshipKey)
	return rv
}

func (s_ ScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasOrderedToManyRelationshipForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasPropertyForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasReadablePropertyForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasWritablePropertyForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) TypeForKey(key string) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("typeForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.GetSelector("selectorForCommand:"), objc.ExtractPtr(commandDescription))
	return rv
}

func (s_ ScriptClassDescription) SupportsCommand(commandDescription IScriptCommandDescription) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("supportsCommand:"), objc.ExtractPtr(commandDescription))
	return rv
}

func (s_ ScriptClassDescription) SuperclassDescription() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.GetSelector("superclassDescription"))
	return rv
}

func (s_ ScriptClassDescription) ClassName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("className"))
	return rv
}

func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("defaultSubcontainerAttributeKey"))
	return rv
}

func (s_ ScriptClassDescription) ImplementationClassName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("implementationClassName"))
	return rv
}

func (s_ ScriptClassDescription) SuiteName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("suiteName"))
	return rv
}
