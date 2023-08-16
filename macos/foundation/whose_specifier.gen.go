// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WhoseSpecifier] class.
var WhoseSpecifierClass = _WhoseSpecifierClass{objc.GetClass("NSWhoseSpecifier")}

type _WhoseSpecifierClass struct {
	objc.Class
}

// An interface definition for the [WhoseSpecifier] class.
type IWhoseSpecifier interface {
	IScriptObjectSpecifier
	StartSubelementIndex() int
	SetStartSubelementIndex(value int)
	EndSubelementIdentifier() WhoseSubelementIdentifier
	SetEndSubelementIdentifier(value WhoseSubelementIdentifier)
	Test() ScriptWhoseTest
	SetTest(value IScriptWhoseTest)
	EndSubelementIndex() int
	SetEndSubelementIndex(value int)
	StartSubelementIdentifier() WhoseSubelementIdentifier
	SetStartSubelementIdentifier(value WhoseSubelementIdentifier)
}

// A specifier that indicates every object in a collection matching a condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier?language=objc
type WhoseSpecifier struct {
	ScriptObjectSpecifier
}

func WhoseSpecifierFrom(ptr unsafe.Pointer) WhoseSpecifier {
	return WhoseSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

func (w_ WhoseSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, test IScriptWhoseTest) WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](w_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:test:"), objc.Ptr(classDesc), objc.Ptr(container), property, objc.Ptr(test))
	return rv
}

// Returns an NSWhoseSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1412173-initwithcontainerclassdescriptio?language=objc
func WhoseSpecifier_InitWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, test IScriptWhoseTest) WhoseSpecifier {
	return WhoseSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc, container, property, test)
}

func (wc _WhoseSpecifierClass) Alloc() WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](wc, objc.Sel("alloc"))
	return rv
}

func WhoseSpecifier_Alloc() WhoseSpecifier {
	return WhoseSpecifierClass.Alloc()
}

func (wc _WhoseSpecifierClass) New() WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWhoseSpecifier() WhoseSpecifier {
	return WhoseSpecifierClass.New()
}

func (w_ WhoseSpecifier) Init() WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](w_, objc.Sel("init"))
	return rv
}

func (w_ WhoseSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](w_, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), objc.Ptr(classDesc), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1410480-initwithcontainerclassdescriptio?language=objc
func WhoseSpecifier_InitWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) WhoseSpecifier {
	return WhoseSpecifierClass.Alloc().InitWithContainerClassDescriptionContainerSpecifierKey(classDesc, container, property)
}

func (w_ WhoseSpecifier) InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) WhoseSpecifier {
	rv := objc.Call[WhoseSpecifier](w_, objc.Sel("initWithContainerSpecifier:key:"), objc.Ptr(container), property)
	return rv
}

// Returns an NSScriptObjectSpecifier object initialized with a given container specifier  and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/1409384-initwithcontainerspecifier?language=objc
func WhoseSpecifier_InitWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) WhoseSpecifier {
	return WhoseSpecifierClass.Alloc().InitWithContainerSpecifierKey(container, property)
}

// Returns the index position of the first sub-element within the range of objects being tested that pass the receiver's test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1417856-startsubelementindex?language=objc
func (w_ WhoseSpecifier) StartSubelementIndex() int {
	rv := objc.Call[int](w_, objc.Sel("startSubelementIndex"))
	return rv
}

// Returns the index position of the first sub-element within the range of objects being tested that pass the receiver's test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1417856-startsubelementindex?language=objc
func (w_ WhoseSpecifier) SetStartSubelementIndex(value int) {
	objc.Call[objc.Void](w_, objc.Sel("setStartSubelementIndex:"), value)
}

// Sets the end sub-element identifier for the specifier to the value of a given sub-element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1407761-endsubelementidentifier?language=objc
func (w_ WhoseSpecifier) EndSubelementIdentifier() WhoseSubelementIdentifier {
	rv := objc.Call[WhoseSubelementIdentifier](w_, objc.Sel("endSubelementIdentifier"))
	return rv
}

// Sets the end sub-element identifier for the specifier to the value of a given sub-element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1407761-endsubelementidentifier?language=objc
func (w_ WhoseSpecifier) SetEndSubelementIdentifier(value WhoseSubelementIdentifier) {
	objc.Call[objc.Void](w_, objc.Sel("setEndSubelementIdentifier:"), value)
}

// Returns the test object encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1412482-test?language=objc
func (w_ WhoseSpecifier) Test() ScriptWhoseTest {
	rv := objc.Call[ScriptWhoseTest](w_, objc.Sel("test"))
	return rv
}

// Returns the test object encapsulated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1412482-test?language=objc
func (w_ WhoseSpecifier) SetTest(value IScriptWhoseTest) {
	objc.Call[objc.Void](w_, objc.Sel("setTest:"), objc.Ptr(value))
}

// Sets the index position of the last sub-element within the range of objects being tested that pass the specifier’s test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1416686-endsubelementindex?language=objc
func (w_ WhoseSpecifier) EndSubelementIndex() int {
	rv := objc.Call[int](w_, objc.Sel("endSubelementIndex"))
	return rv
}

// Sets the index position of the last sub-element within the range of objects being tested that pass the specifier’s test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1416686-endsubelementindex?language=objc
func (w_ WhoseSpecifier) SetEndSubelementIndex(value int) {
	objc.Call[objc.Void](w_, objc.Sel("setEndSubelementIndex:"), value)
}

// Returns the start sub-element identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1413408-startsubelementidentifier?language=objc
func (w_ WhoseSpecifier) StartSubelementIdentifier() WhoseSubelementIdentifier {
	rv := objc.Call[WhoseSubelementIdentifier](w_, objc.Sel("startSubelementIdentifier"))
	return rv
}

// Returns the start sub-element identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosespecifier/1413408-startsubelementidentifier?language=objc
func (w_ WhoseSpecifier) SetStartSubelementIdentifier(value WhoseSubelementIdentifier) {
	objc.Call[objc.Void](w_, objc.Sel("setStartSubelementIdentifier:"), value)
}
