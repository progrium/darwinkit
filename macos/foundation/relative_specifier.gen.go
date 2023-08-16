// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RelativeSpecifier] class.
var RelativeSpecifierClass = _RelativeSpecifierClass{objc.GetClass("NSRelativeSpecifier")}

type _RelativeSpecifierClass struct {
	objc.Class
}

// An interface definition for the [RelativeSpecifier] class.
type IRelativeSpecifier interface {
	IScriptObjectSpecifier
	RelativePosition() RelativePosition
	SetRelativePosition(value RelativePosition)
	BaseSpecifier() ScriptObjectSpecifier
	SetBaseSpecifier(value IScriptObjectSpecifier)
}

// A specifier that indicates an object in a collection by its position relative to another object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier?language=objc
type RelativeSpecifier struct {
	ScriptObjectSpecifier
}

func RelativeSpecifierFrom(ptr unsafe.Pointer) RelativeSpecifier {
	return RelativeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (r_ RelativeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, relPos RelativePosition, baseSpecifier IScriptObjectSpecifier) RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](r_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:relativePosition:baseSpecifier:"), objc.Ptr(classDesc), objc.Ptr(container), property, relPos, objc.Ptr(baseSpecifier))
	return rv
}

// Invokes the super class’s initWithContainerClassDescription:containerSpecifier:key: method and initializes the relative position and base specifier to relPos and baseSpecifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier/1409205-initwithcontainerclassdescriptio?language=objc
func RelativeSpecifier_InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, relPos RelativePosition, baseSpecifier IScriptObjectSpecifier) RelativeSpecifier {
	return RelativeSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc, container, property, relPos, baseSpecifier)
}

func (rc _RelativeSpecifierClass) Alloc() RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](rc, objc.Sel("alloc"))
	return rv
}

func RelativeSpecifier_Alloc() RelativeSpecifier {
	return RelativeSpecifierClass.Alloc()
}

func (rc _RelativeSpecifierClass) New() RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRelativeSpecifier() RelativeSpecifier {
	return RelativeSpecifierClass.New()
}

func (r_ RelativeSpecifier) Init() RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](r_, objc.Sel("init"))
	return rv
}

func (r_ RelativeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](r_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func RelativeSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RelativeSpecifier {
	return RelativeSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (r_ RelativeSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RelativeSpecifier {
	rv := objc.Call[RelativeSpecifier](r_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func RelativeSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RelativeSpecifier {
	return RelativeSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}

// Sets the relative position encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier/1416001-relativeposition?language=objc
func (r_ RelativeSpecifier) RelativePosition() RelativePosition {
	rv := objc.Call[RelativePosition](r_, objc.Sel("relativePosition"))
	return rv
}

// Sets the relative position encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier/1416001-relativeposition?language=objc
func (r_ RelativeSpecifier) SetRelativePosition(value RelativePosition) {
	objc.Call[objc.Void](r_, objc.Sel("setRelativePosition:"), value)
}

// Sets the specifier for the base object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier/1409071-basespecifier?language=objc
func (r_ RelativeSpecifier) BaseSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](r_, objc.Sel("baseSpecifier"))
	return rv
}

// Sets the specifier for the base object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativespecifier/1409071-basespecifier?language=objc
func (r_ RelativeSpecifier) SetBaseSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](r_, objc.Sel("setBaseSpecifier:"), objc.Ptr(value))
}
