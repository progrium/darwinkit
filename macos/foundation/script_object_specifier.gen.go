// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ScriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}

type _ScriptObjectSpecifierClass struct {
	objc.Class
}

type IScriptObjectSpecifier interface {
	objc.IObject
	IndicesOfObjectsByEvaluatingWithContainerCount(container objc.IObject, count *int) *int
	ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object
	ObjectsByEvaluatingSpecifier() objc.Object
	ContainerClassDescription() ScriptClassDescription
	SetContainerClassDescription(value IScriptClassDescription)
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
	ContainerIsRangeContainerObject() bool
	SetContainerIsRangeContainerObject(value bool)
	ContainerSpecifier() ScriptObjectSpecifier
	SetContainerSpecifier(value IScriptObjectSpecifier)
	ChildSpecifier() ScriptObjectSpecifier
	SetChildSpecifier(value IScriptObjectSpecifier)
	Key() string
	SetKey(value string)
	KeyClassDescription() ScriptClassDescription
	EvaluationErrorSpecifier() ScriptObjectSpecifier
	EvaluationErrorNumber() int
	SetEvaluationErrorNumber(value int)
	Descriptor() AppleEventDescriptor
}

type ScriptObjectSpecifier struct {
	objc.Object
}

func MakeScriptObjectSpecifier(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ ScriptObjectSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("initWithContainerClassDescription:containerSpecifier:key:"), objc.ExtractPtr(classDesc), objc.ExtractPtr(container), property)
	return rv
}

func ScriptObjectSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (s_ ScriptObjectSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("initWithContainerSpecifier:key:"), objc.ExtractPtr(container), property)
	return rv
}

func ScriptObjectSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}

func (sc _ScriptObjectSpecifierClass) Alloc() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](sc, objc.GetSelector("alloc"))
	return rv
}

func ScriptObjectSpecifier_Alloc() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.Alloc()
}

func (sc _ScriptObjectSpecifierClass) New() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScriptObjectSpecifier() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.New()
}

func ScriptObjectSpecifier_New() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.New()
}

func (s_ ScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("init"))
	return rv
}

func ScriptObjectSpecifier_Init() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.Alloc().Init()
}

func (sc _ScriptObjectSpecifierClass) ObjectSpecifierWithDescriptor(descriptor IAppleEventDescriptor) ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](sc, objc.GetSelector("objectSpecifierWithDescriptor:"), objc.ExtractPtr(descriptor))
	return rv
}

func ScriptObjectSpecifier_ObjectSpecifierWithDescriptor(descriptor IAppleEventDescriptor) ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.ObjectSpecifierWithDescriptor(descriptor)
}

func (s_ ScriptObjectSpecifier) IndicesOfObjectsByEvaluatingWithContainerCount(container objc.IObject, count *int) *int {
	rv := objc.CallMethod[*int](s_, objc.GetSelector("indicesOfObjectsByEvaluatingWithContainer:count:"), objc.ExtractPtr(container), count)
	return rv
}

func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("objectsByEvaluatingWithContainers:"), objc.ExtractPtr(containers))
	return rv
}

func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("objectsByEvaluatingSpecifier"))
	return rv
}

func (s_ ScriptObjectSpecifier) ContainerClassDescription() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.GetSelector("containerClassDescription"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerClassDescription(value IScriptClassDescription) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContainerClassDescription:"), objc.ExtractPtr(value))
}

func (s_ ScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("containerIsObjectBeingTested"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContainerIsObjectBeingTested:"), value)
}

func (s_ ScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("containerIsRangeContainerObject"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContainerIsRangeContainerObject:"), value)
}

func (s_ ScriptObjectSpecifier) ContainerSpecifier() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("containerSpecifier"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerSpecifier(value IScriptObjectSpecifier) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContainerSpecifier:"), objc.ExtractPtr(value))
}

func (s_ ScriptObjectSpecifier) ChildSpecifier() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("childSpecifier"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetChildSpecifier(value IScriptObjectSpecifier) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setChildSpecifier:"), objc.ExtractPtr(value))
}

func (s_ ScriptObjectSpecifier) Key() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("key"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetKey(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setKey:"), value)
}

func (s_ ScriptObjectSpecifier) KeyClassDescription() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.GetSelector("keyClassDescription"))
	return rv
}

func (s_ ScriptObjectSpecifier) EvaluationErrorSpecifier() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.GetSelector("evaluationErrorSpecifier"))
	return rv
}

func (s_ ScriptObjectSpecifier) EvaluationErrorNumber() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("evaluationErrorNumber"))
	return rv
}

func (s_ ScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEvaluationErrorNumber:"), value)
}

func (s_ ScriptObjectSpecifier) Descriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](s_, objc.GetSelector("descriptor"))
	return rv
}
