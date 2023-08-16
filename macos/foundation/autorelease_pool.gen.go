// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AutoreleasePool] class.
var AutoreleasePoolClass = _AutoreleasePoolClass{objc.GetClass("NSAutoreleasePool")}

type _AutoreleasePoolClass struct {
	objc.Class
}

// An interface definition for the [AutoreleasePool] class.
type IAutoreleasePool interface {
	objc.IObject
	Drain()
	AddObject(anObject objc.IObject)
}

// An object that supports Cocoa’s reference-counted memory management system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsautoreleasepool?language=objc
type AutoreleasePool struct {
	objc.Object
}

func AutoreleasePoolFrom(ptr unsafe.Pointer) AutoreleasePool {
	return AutoreleasePool{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AutoreleasePoolClass) Alloc() AutoreleasePool {
	rv := objc.Call[AutoreleasePool](ac, objc.Sel("alloc"))
	return rv
}

func AutoreleasePool_Alloc() AutoreleasePool {
	return AutoreleasePoolClass.Alloc()
}

func (ac _AutoreleasePoolClass) New() AutoreleasePool {
	rv := objc.Call[AutoreleasePool](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAutoreleasePool() AutoreleasePool {
	return AutoreleasePoolClass.New()
}

func (a_ AutoreleasePool) Init() AutoreleasePool {
	rv := objc.Call[AutoreleasePool](a_, objc.Sel("init"))
	return rv
}

// In a reference-counted environment, releases and pops the receiver; in a garbage-collected environment, triggers garbage collection if the memory allocated since the last collection is greater than the current threshold. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsautoreleasepool/1520553-drain?language=objc
func (a_ AutoreleasePool) Drain() {
	objc.Call[objc.Void](a_, objc.Sel("drain"))
}

// Adds a given object to the receiver [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsautoreleasepool/1520555-addobject?language=objc
func (a_ AutoreleasePool) AddObject(anObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("addObject:"), anObject)
}

// Displays the state of the current thread's autorelease pool stack to stderr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsautoreleasepool/1539124-showpools?language=objc
func (ac _AutoreleasePoolClass) ShowPools() {
	objc.Call[objc.Void](ac, objc.Sel("showPools"))
}

// Displays the state of the current thread's autorelease pool stack to stderr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsautoreleasepool/1539124-showpools?language=objc
func AutoreleasePool_ShowPools() {
	AutoreleasePoolClass.ShowPools()
}
