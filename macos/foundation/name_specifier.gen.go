// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NameSpecifier] class.
var NameSpecifierClass = _NameSpecifierClass{objc.GetClass("NSNameSpecifier")}

type _NameSpecifierClass struct {
	objc.Class
}

// An interface definition for the [NameSpecifier] class.
type INameSpecifier interface {
	IScriptObjectSpecifier
	Name() string
	SetName(value string)
}

// A specifier for an object in a collection (or container) by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnamespecifier?language=objc
type NameSpecifier struct {
	ScriptObjectSpecifier
}

func NameSpecifierFrom(ptr unsafe.Pointer) NameSpecifier {
	return NameSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (n_ NameSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyName(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, name string) NameSpecifier {
	rv := objc.Call[NameSpecifier](n_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:name:"), objc.Ptr(classDesc), objc.Ptr(container), property, name)
	return rv
}

// Invokes the super class’s initWithContainerClassDescription:containerSpecifier:key: method and then sets the name instance variable to name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnamespecifier/1408615-initwithcontainerclassdescriptio?language=objc
func NameSpecifier_InitWithContainerClassDescriptionContainerSpecifierKeyName(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, name string) NameSpecifier {
	return NameSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyName(classDesc, container, property, name)
}

func (nc _NameSpecifierClass) Alloc() NameSpecifier {
	rv := objc.Call[NameSpecifier](nc, objc.Sel("alloc"))
	return rv
}

func NameSpecifier_Alloc() NameSpecifier {
	return NameSpecifierClass.Alloc()
}

func (nc _NameSpecifierClass) New() NameSpecifier {
	rv := objc.Call[NameSpecifier](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNameSpecifier() NameSpecifier {
	return NameSpecifierClass.New()
}

func (n_ NameSpecifier) Init() NameSpecifier {
	rv := objc.Call[NameSpecifier](n_, objc.Sel("init"))
	return rv
}

func (n_ NameSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) NameSpecifier {
	rv := objc.Call[NameSpecifier](n_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NameSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) NameSpecifier {
	return NameSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (n_ NameSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) NameSpecifier {
	rv := objc.Call[NameSpecifier](n_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NameSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) NameSpecifier {
	return NameSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}

// Sets the name encapsulated with the receiver for the specified object in the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnamespecifier/1407411-name?language=objc
func (n_ NameSpecifier) Name() string {
	rv := objc.Call[string](n_, objc.Sel("name"))
	return rv
}

// Sets the name encapsulated with the receiver for the specified object in the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnamespecifier/1407411-name?language=objc
func (n_ NameSpecifier) SetName(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setName:"), value)
}
