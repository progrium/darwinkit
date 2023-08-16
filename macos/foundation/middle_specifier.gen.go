// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MiddleSpecifier] class.
var MiddleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}

type _MiddleSpecifierClass struct {
	objc.Class
}

// An interface definition for the [MiddleSpecifier] class.
type IMiddleSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmiddlespecifier?language=objc
type MiddleSpecifier struct {
	ScriptObjectSpecifier
}

func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (mc _MiddleSpecifierClass) Alloc() MiddleSpecifier {
	rv := objc.Call[MiddleSpecifier](mc, objc.Sel("alloc"))
	return rv
}

func MiddleSpecifier_Alloc() MiddleSpecifier {
	return MiddleSpecifierClass.Alloc()
}

func (mc _MiddleSpecifierClass) New() MiddleSpecifier {
	rv := objc.Call[MiddleSpecifier](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMiddleSpecifier() MiddleSpecifier {
	return MiddleSpecifierClass.New()
}

func (m_ MiddleSpecifier) Init() MiddleSpecifier {
	rv := objc.Call[MiddleSpecifier](m_, objc.Sel("init"))
	return rv
}

func (m_ MiddleSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) MiddleSpecifier {
	rv := objc.Call[MiddleSpecifier](m_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func MiddleSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) MiddleSpecifier {
	return MiddleSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (m_ MiddleSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) MiddleSpecifier {
	rv := objc.Call[MiddleSpecifier](m_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func MiddleSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) MiddleSpecifier {
	return MiddleSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}
