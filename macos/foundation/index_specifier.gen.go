// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IndexSpecifier] class.
var IndexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}

type _IndexSpecifierClass struct {
	objc.Class
}

// An interface definition for the [IndexSpecifier] class.
type IIndexSpecifier interface {
	IScriptObjectSpecifier
	Index() int
	SetIndex(value int)
}

// A specifier representing an object in a collection (or container) with an index number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier?language=objc
type IndexSpecifier struct {
	ScriptObjectSpecifier
}

func IndexSpecifierFrom(ptr unsafe.Pointer) IndexSpecifier {
	return IndexSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (i_ IndexSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, index int) IndexSpecifier {
	rv := objc.Call[IndexSpecifier](i_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:index:"), objc.Ptr(classDesc), objc.Ptr(container), property, index)
	return rv
}

// Initializes an allocated NSIndexSpecifier object with a class description, container specifier, collection key, and object index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier/1407502-initwithcontainerclassdescriptio?language=objc
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, index int) IndexSpecifier {
	instance := IndexSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc, container, property, index)
	instance.Autorelease()
	return instance
}

func (ic _IndexSpecifierClass) Alloc() IndexSpecifier {
	rv := objc.Call[IndexSpecifier](ic, objc.Sel("alloc"))
	return rv
}

func IndexSpecifier_Alloc() IndexSpecifier {
	return IndexSpecifierClass.Alloc()
}

func (ic _IndexSpecifierClass) New() IndexSpecifier {
	rv := objc.Call[IndexSpecifier](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIndexSpecifier() IndexSpecifier {
	return IndexSpecifierClass.New()
}

func (i_ IndexSpecifier) Init() IndexSpecifier {
	rv := objc.Call[IndexSpecifier](i_, objc.Sel("init"))
	return rv
}

func (i_ IndexSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) IndexSpecifier {
	rv := objc.Call[IndexSpecifier](i_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) IndexSpecifier {
	instance := IndexSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
	instance.Autorelease()
	return instance
}

func (i_ IndexSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) IndexSpecifier {
	rv := objc.Call[IndexSpecifier](i_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func NewIndexSpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) IndexSpecifier {
	instance := IndexSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
	instance.Autorelease()
	return instance
}

// Sets the value of the receiver’s index property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier/1408567-index?language=objc
func (i_ IndexSpecifier) Index() int {
	rv := objc.Call[int](i_, objc.Sel("index"))
	return rv
}

// Sets the value of the receiver’s index property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier/1408567-index?language=objc
func (i_ IndexSpecifier) SetIndex(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setIndex:"), value)
}
