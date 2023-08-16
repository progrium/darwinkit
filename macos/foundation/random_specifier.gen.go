// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RandomSpecifier] class.
var RandomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}

type _RandomSpecifierClass struct {
	objc.Class
}

// An interface definition for the [RandomSpecifier] class.
type IRandomSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrandomspecifier?language=objc
type RandomSpecifier struct {
	ScriptObjectSpecifier
}

func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (rc _RandomSpecifierClass) Alloc() RandomSpecifier {
	rv := objc.Call[RandomSpecifier](rc, objc.Sel("alloc"))
	return rv
}

func RandomSpecifier_Alloc() RandomSpecifier {
	return RandomSpecifierClass.Alloc()
}

func (rc _RandomSpecifierClass) New() RandomSpecifier {
	rv := objc.Call[RandomSpecifier](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRandomSpecifier() RandomSpecifier {
	return RandomSpecifierClass.New()
}

func (r_ RandomSpecifier) Init() RandomSpecifier {
	rv := objc.Call[RandomSpecifier](r_, objc.Sel("init"))
	return rv
}

func (r_ RandomSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RandomSpecifier {
	rv := objc.Call[RandomSpecifier](r_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func RandomSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) RandomSpecifier {
	return RandomSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (r_ RandomSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RandomSpecifier {
	rv := objc.Call[RandomSpecifier](r_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func RandomSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) RandomSpecifier {
	return RandomSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}
