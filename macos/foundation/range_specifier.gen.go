// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RangeSpecifier] class.
var RangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}

type _RangeSpecifierClass struct {
	objc.Class
}

// An interface definition for the [RangeSpecifier] class.
type IRangeSpecifier interface {
	IScriptObjectSpecifier
	StartSpecifier() ScriptObjectSpecifier
	SetStartSpecifier(value IScriptObjectSpecifier)
	EndSpecifier() ScriptObjectSpecifier
	SetEndSpecifier(value IScriptObjectSpecifier)
}

// A specifier for a range of objects in a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier?language=objc
type RangeSpecifier struct {
	ScriptObjectSpecifier
}

func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (r_ RangeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, startSpec IScriptObjectSpecifier, endSpec IScriptObjectSpecifier) RangeSpecifier {
	rv := objc.Call[RangeSpecifier](r_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:startSpecifier:endSpecifier:"), objc.Ptr(classDesc), objc.Ptr(container), property, objc.Ptr(startSpec), objc.Ptr(endSpec))
	return rv
}

// Returns a range specifier initialized with the given properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/1409215-initwithcontainerclassdescriptio?language=objc
func NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, startSpec IScriptObjectSpecifier, endSpec IScriptObjectSpecifier) RangeSpecifier {
	instance := RangeSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc, container, property, startSpec, endSpec)
	instance.Autorelease()
	return instance
}

func (rc _RangeSpecifierClass) Alloc() RangeSpecifier {
	rv := objc.Call[RangeSpecifier](rc, objc.Sel("alloc"))
	return rv
}

func RangeSpecifier_Alloc() RangeSpecifier {
	return RangeSpecifierClass.Alloc()
}

func (rc _RangeSpecifierClass) New() RangeSpecifier {
	rv := objc.Call[RangeSpecifier](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRangeSpecifier() RangeSpecifier {
	return RangeSpecifierClass.New()
}

func (r_ RangeSpecifier) Init() RangeSpecifier {
	rv := objc.Call[RangeSpecifier](r_, objc.Sel("init"))
	return rv
}

func (r_ RangeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RangeSpecifier {
	rv := objc.Call[RangeSpecifier](r_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RangeSpecifier {
	instance := RangeSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
	instance.Autorelease()
	return instance
}

func (r_ RangeSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RangeSpecifier {
	rv := objc.Call[RangeSpecifier](r_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NewRangeSpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RangeSpecifier {
	instance := RangeSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
	instance.Autorelease()
	return instance
}

// Returns the object specifier representing the first object of the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/1418394-startspecifier?language=objc
func (r_ RangeSpecifier) StartSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](r_, objc.Sel("startSpecifier"))
	return rv
}

// Returns the object specifier representing the first object of the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/1418394-startspecifier?language=objc
func (r_ RangeSpecifier) SetStartSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](r_, objc.Sel("setStartSpecifier:"), objc.Ptr(value))
}

// Sets the object specifier representing the last object of the range to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/1418470-endspecifier?language=objc
func (r_ RangeSpecifier) EndSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](r_, objc.Sel("endSpecifier"))
	return rv
}

// Sets the object specifier representing the last object of the range to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/1418470-endspecifier?language=objc
func (r_ RangeSpecifier) SetEndSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](r_, objc.Sel("setEndSpecifier:"), objc.Ptr(value))
}
