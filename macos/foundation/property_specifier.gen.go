// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PropertySpecifier] class.
var PropertySpecifierClass = _PropertySpecifierClass{objc.GetClass("NSPropertySpecifier")}

type _PropertySpecifierClass struct {
	objc.Class
}

// An interface definition for the [PropertySpecifier] class.
type IPropertySpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertyspecifier?language=objc
type PropertySpecifier struct {
	ScriptObjectSpecifier
}

func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (pc _PropertySpecifierClass) Alloc() PropertySpecifier {
	rv := objc.Call[PropertySpecifier](pc, objc.Sel("alloc"))
	return rv
}

func PropertySpecifier_Alloc() PropertySpecifier {
	return PropertySpecifierClass.Alloc()
}

func (pc _PropertySpecifierClass) New() PropertySpecifier {
	rv := objc.Call[PropertySpecifier](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPropertySpecifier() PropertySpecifier {
	return PropertySpecifierClass.New()
}

func (p_ PropertySpecifier) Init() PropertySpecifier {
	rv := objc.Call[PropertySpecifier](p_, objc.Sel("init"))
	return rv
}

func (p_ PropertySpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) PropertySpecifier {
	rv := objc.Call[PropertySpecifier](p_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NewPropertySpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) PropertySpecifier {
	instance := PropertySpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
	instance.Autorelease()
	return instance
}

func (p_ PropertySpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) PropertySpecifier {
	rv := objc.Call[PropertySpecifier](p_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NewPropertySpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) PropertySpecifier {
	instance := PropertySpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
	instance.Autorelease()
	return instance
}
