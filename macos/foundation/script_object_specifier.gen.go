// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptObjectSpecifier] class.
var ScriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}

type _ScriptObjectSpecifierClass struct {
	objc.Class
}

// An interface definition for the [ScriptObjectSpecifier] class.
type IScriptObjectSpecifier interface {
	objc.IObject
	ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object
	IndicesOfObjectsByEvaluatingWithContainerCount(container objc.IObject, count *int) *int
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
	ContainerSpecifier() ScriptObjectSpecifier
	SetContainerSpecifier(value IScriptObjectSpecifier)
	ContainerIsRangeContainerObject() bool
	SetContainerIsRangeContainerObject(value bool)
	Key() string
	SetKey(value string)
	KeyClassDescription() ScriptClassDescription
	ContainerClassDescription() ScriptClassDescription
	SetContainerClassDescription(value IScriptClassDescription)
	EvaluationErrorSpecifier() ScriptObjectSpecifier
	ObjectsByEvaluatingSpecifier() objc.Object
	ChildSpecifier() ScriptObjectSpecifier
	SetChildSpecifier(value IScriptObjectSpecifier)
	Descriptor() AppleEventDescriptor
	EvaluationErrorNumber() int
	SetEvaluationErrorNumber(value int)
}

// An abstract class used to represent natural language expressions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier?language=objc
type ScriptObjectSpecifier struct {
	objc.Object
}

func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScriptObjectSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	instance := ScriptObjectSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
	instance.Autorelease()
	return instance
}

func (s_ ScriptObjectSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NewScriptObjectSpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	instance := ScriptObjectSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
	instance.Autorelease()
	return instance
}

func (sc _ScriptObjectSpecifierClass) Alloc() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](sc, objc.Sel("alloc"))
	return rv
}

func ScriptObjectSpecifier_Alloc() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.Alloc()
}

func (sc _ScriptObjectSpecifierClass) New() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptObjectSpecifier() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.New()
}

func (s_ ScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("init"))
	return rv
}

// Returns the actual object or objects specified by the receiver as evaluated in the context of given container object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409842-objectsbyevaluatingwithcontainer?language=objc
func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("objectsByEvaluatingWithContainers:"), containers)
	return rv
}

// Returns a new object specifier for an Apple event descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409779-objectspecifierwithdescriptor?language=objc
func (sc _ScriptObjectSpecifierClass) ObjectSpecifierWithDescriptor(descriptor IAppleEventDescriptor) ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](sc, objc.Sel("objectSpecifierWithDescriptor:"), objc.Ptr(descriptor))
	return rv
}

// Returns a new object specifier for an Apple event descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409779-objectspecifierwithdescriptor?language=objc
func ScriptObjectSpecifier_ObjectSpecifierWithDescriptor(descriptor IAppleEventDescriptor) ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.ObjectSpecifierWithDescriptor(descriptor)
}

// This primitive method must be overridden by subclasses to return a pointer to an array of indices identifying objects in the key of a given container that are identified by the receiver of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1415107-indicesofobjectsbyevaluatingwith?language=objc
func (s_ ScriptObjectSpecifier) IndicesOfObjectsByEvaluatingWithContainerCount(container objc.IObject, count *int) *int {
	rv := objc.Call[*int](s_, objc.Sel("indicesOfObjectsByEvaluatingWithContainer:count:"), container, count)
	return rv
}

// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410887-containerisobjectbeingtested?language=objc
func (s_ ScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	rv := objc.Call[bool](s_, objc.Sel("containerIsObjectBeingTested"))
	return rv
}

// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410887-containerisobjectbeingtested?language=objc
func (s_ ScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerIsObjectBeingTested:"), value)
}

// Sets the container specifier of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1414424-containerspecifier?language=objc
func (s_ ScriptObjectSpecifier) ContainerSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("containerSpecifier"))
	return rv
}

// Sets the container specifier of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1414424-containerspecifier?language=objc
func (s_ ScriptObjectSpecifier) SetContainerSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerSpecifier:"), objc.Ptr(value))
}

// Sets whether the receiver’s container is to be the container for a range specifier or a top-level object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1416507-containerisrangecontainerobject?language=objc
func (s_ ScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	rv := objc.Call[bool](s_, objc.Sel("containerIsRangeContainerObject"))
	return rv
}

// Sets whether the receiver’s container is to be the container for a range specifier or a top-level object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1416507-containerisrangecontainerobject?language=objc
func (s_ ScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerIsRangeContainerObject:"), value)
}

// Sets the key of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1412986-key?language=objc
func (s_ ScriptObjectSpecifier) Key() string {
	rv := objc.Call[string](s_, objc.Sel("key"))
	return rv
}

// Sets the key of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1412986-key?language=objc
func (s_ ScriptObjectSpecifier) SetKey(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setKey:"), value)
}

// Returns the class description of the objects specified by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1417974-keyclassdescription?language=objc
func (s_ ScriptObjectSpecifier) KeyClassDescription() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("keyClassDescription"))
	return rv
}

// Sets the class description of the receiver’s container specifier to a given specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1413179-containerclassdescription?language=objc
func (s_ ScriptObjectSpecifier) ContainerClassDescription() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("containerClassDescription"))
	return rv
}

// Sets the class description of the receiver’s container specifier to a given specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1413179-containerclassdescription?language=objc
func (s_ ScriptObjectSpecifier) SetContainerClassDescription(value IScriptClassDescription) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerClassDescription:"), objc.Ptr(value))
}

// Returns the object specifier in which an evaluation error occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1416385-evaluationerrorspecifier?language=objc
func (s_ ScriptObjectSpecifier) EvaluationErrorSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("evaluationErrorSpecifier"))
	return rv
}

// Returns the actual object represented by the nested series of object specifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1413391-objectsbyevaluatingspecifier?language=objc
func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("objectsByEvaluatingSpecifier"))
	return rv
}

// Sets the receiver’s child reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409882-childspecifier?language=objc
func (s_ ScriptObjectSpecifier) ChildSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("childSpecifier"))
	return rv
}

// Sets the receiver’s child reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409882-childspecifier?language=objc
func (s_ ScriptObjectSpecifier) SetChildSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](s_, objc.Sel("setChildSpecifier:"), objc.Ptr(value))
}

// Returns an Apple event descriptor that represents the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410018-descriptor?language=objc
func (s_ ScriptObjectSpecifier) Descriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](s_, objc.Sel("descriptor"))
	return rv
}

// Sets the value of the evaluation error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1416938-evaluationerrornumber?language=objc
func (s_ ScriptObjectSpecifier) EvaluationErrorNumber() int {
	rv := objc.Call[int](s_, objc.Sel("evaluationErrorNumber"))
	return rv
}

// Sets the value of the evaluation error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1416938-evaluationerrornumber?language=objc
func (s_ ScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setEvaluationErrorNumber:"), value)
}
