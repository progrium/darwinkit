// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var UniqueIDSpecifierClass = _UniqueIDSpecifierClass{objc.GetClass("NSUniqueIDSpecifier")}

type _UniqueIDSpecifierClass struct {
	objc.Class
}

// An interface definition for the [UniqueIDSpecifier] class.
type IUniqueIDSpecifier interface {
	IScriptObjectSpecifier
	UniqueID() objc.Object
	SetUniqueID(value objc.IObject)
}

// A specifier for an object in a collection (or container) by unique ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier?language=objc
type UniqueIDSpecifier struct {
	ScriptObjectSpecifier
}

func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (u_ UniqueIDSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, uniqueID objc.IObject) UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](u_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:uniqueID:"), objc.Ptr(classDesc), objc.Ptr(container), property, uniqueID)
	return rv
}

// Returns an NSUniqueIDSpecifier object, initialized with the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier/1416055-initwithcontainerclassdescriptio?language=objc
func NewUniqueIDSpecifierWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, uniqueID objc.IObject) UniqueIDSpecifier {
	instance := UniqueIDSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc, container, property, uniqueID)
	instance.Autorelease()
	return instance
}

func (uc _UniqueIDSpecifierClass) Alloc() UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](uc, objc.Sel("alloc"))
	return rv
}

func UniqueIDSpecifier_Alloc() UniqueIDSpecifier {
	return UniqueIDSpecifierClass.Alloc()
}

func (uc _UniqueIDSpecifierClass) New() UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUniqueIDSpecifier() UniqueIDSpecifier {
	return UniqueIDSpecifierClass.New()
}

func (u_ UniqueIDSpecifier) Init() UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](u_, objc.Sel("init"))
	return rv
}

func (u_ UniqueIDSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](u_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NewUniqueIDSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) UniqueIDSpecifier {
	instance := UniqueIDSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
	instance.Autorelease()
	return instance
}

func (u_ UniqueIDSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) UniqueIDSpecifier {
	rv := objc.Call[UniqueIDSpecifier](u_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NewUniqueIDSpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) UniqueIDSpecifier {
	instance := UniqueIDSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
	instance.Autorelease()
	return instance
}

// Returns the ID encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier/1415634-uniqueid?language=objc
func (u_ UniqueIDSpecifier) UniqueID() objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("uniqueID"))
	return rv
}

// Returns the ID encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier/1415634-uniqueid?language=objc
func (u_ UniqueIDSpecifier) SetUniqueID(value objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setUniqueID:"), value)
}
