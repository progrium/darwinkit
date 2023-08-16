// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PointerFunctions] class.
var PointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}

type _PointerFunctionsClass struct {
	objc.Class
}

// An interface definition for the [PointerFunctions] class.
type IPointerFunctions interface {
	objc.IObject
	DescriptionFunction() func(item unsafe.Pointer) string
	SetDescriptionFunction(value func(item unsafe.Pointer) string)
	IsEqualFunction() func(item1 unsafe.Pointer, item2 unsafe.Pointer, arg2 func(item unsafe.Pointer) uint) bool
	SetIsEqualFunction(value func(item1 unsafe.Pointer, item2 unsafe.Pointer, arg2 func(item unsafe.Pointer) uint) bool)
	HashFunction() func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) uint
	SetHashFunction(value func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) uint)
	SizeFunction() func(item unsafe.Pointer) uint
	SetSizeFunction(value func(item unsafe.Pointer) uint)
	RelinquishFunction() func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint)
	SetRelinquishFunction(value func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint))
	AcquireFunction() func(src unsafe.Pointer, arg1 func(item unsafe.Pointer) uint, shouldCopy bool) unsafe.Pointer
	SetAcquireFunction(value func(src unsafe.Pointer, arg1 func(item unsafe.Pointer) uint, shouldCopy bool) unsafe.Pointer)
}

// An instance of NSPointerFunctions defines callout functions appropriate for managing a pointer reference held somewhere else. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions?language=objc
type PointerFunctions struct {
	objc.Object
}

func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PointerFunctions) InitWithOptions(options PointerFunctionsOptions) PointerFunctions {
	rv := objc.Call[PointerFunctions](p_, objc.Sel("initWithOptions:"), options)
	return rv
}

// Returns an NSPointerFunctions object initialized with the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1417715-initwithoptions?language=objc
func PointerFunctions_InitWithOptions(options PointerFunctionsOptions) PointerFunctions {
	return PointerFunctionsClass.Alloc().InitWithOptions(options)
}

func (pc _PointerFunctionsClass) Alloc() PointerFunctions {
	rv := objc.Call[PointerFunctions](pc, objc.Sel("alloc"))
	return rv
}

func PointerFunctions_Alloc() PointerFunctions {
	return PointerFunctionsClass.Alloc()
}

func (pc _PointerFunctionsClass) New() PointerFunctions {
	rv := objc.Call[PointerFunctions](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPointerFunctions() PointerFunctions {
	return PointerFunctionsClass.New()
}

func (p_ PointerFunctions) Init() PointerFunctions {
	rv := objc.Call[PointerFunctions](p_, objc.Sel("init"))
	return rv
}

// Returns a new NSPointerFunctions object initialized with the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1451770-pointerfunctionswithoptions?language=objc
func (pc _PointerFunctionsClass) PointerFunctionsWithOptions(options PointerFunctionsOptions) PointerFunctions {
	rv := objc.Call[PointerFunctions](pc, objc.Sel("pointerFunctionsWithOptions:"), options)
	return rv
}

// Returns a new NSPointerFunctions object initialized with the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1451770-pointerfunctionswithoptions?language=objc
func PointerFunctions_PointerFunctionsWithOptions(options PointerFunctionsOptions) PointerFunctions {
	return PointerFunctionsClass.PointerFunctionsWithOptions(options)
}

// The function used to describe elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1415200-descriptionfunction?language=objc
func (p_ PointerFunctions) DescriptionFunction() func(item unsafe.Pointer) string {
	rv := objc.Call[func(item unsafe.Pointer) string](p_, objc.Sel("descriptionFunction"))
	return rv
}

// The function used to describe elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1415200-descriptionfunction?language=objc
func (p_ PointerFunctions) SetDescriptionFunction(value func(item unsafe.Pointer) string) {
	objc.Call[objc.Void](p_, objc.Sel("setDescriptionFunction:"), value)
}

// The function used to compare pointers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1413473-isequalfunction?language=objc
func (p_ PointerFunctions) IsEqualFunction() func(item1 unsafe.Pointer, item2 unsafe.Pointer, arg2 func(item unsafe.Pointer) uint) bool {
	rv := objc.Call[func(item1 unsafe.Pointer, item2 unsafe.Pointer, arg2 func(item unsafe.Pointer) uint) bool](p_, objc.Sel("isEqualFunction"))
	return rv
}

// The function used to compare pointers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1413473-isequalfunction?language=objc
func (p_ PointerFunctions) SetIsEqualFunction(value func(item1 unsafe.Pointer, item2 unsafe.Pointer, arg2 func(item unsafe.Pointer) uint) bool) {
	objc.Call[objc.Void](p_, objc.Sel("setIsEqualFunction:"), value)
}

// The hash function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1415939-hashfunction?language=objc
func (p_ PointerFunctions) HashFunction() func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) uint {
	rv := objc.Call[func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) uint](p_, objc.Sel("hashFunction"))
	return rv
}

// The hash function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1415939-hashfunction?language=objc
func (p_ PointerFunctions) SetHashFunction(value func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) uint) {
	objc.Call[objc.Void](p_, objc.Sel("setHashFunction:"), value)
}

// The function used to determine the size of pointers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1408045-sizefunction?language=objc
func (p_ PointerFunctions) SizeFunction() func(item unsafe.Pointer) uint {
	rv := objc.Call[func(item unsafe.Pointer) uint](p_, objc.Sel("sizeFunction"))
	return rv
}

// The function used to determine the size of pointers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1408045-sizefunction?language=objc
func (p_ PointerFunctions) SetSizeFunction(value func(item unsafe.Pointer) uint) {
	objc.Call[objc.Void](p_, objc.Sel("setSizeFunction:"), value)
}

// The function used to relinquish memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1408565-relinquishfunction?language=objc
func (p_ PointerFunctions) RelinquishFunction() func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint) {
	rv := objc.Call[func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint)](p_, objc.Sel("relinquishFunction"))
	return rv
}

// The function used to relinquish memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1408565-relinquishfunction?language=objc
func (p_ PointerFunctions) SetRelinquishFunction(value func(item unsafe.Pointer, arg1 func(item unsafe.Pointer) uint)) {
	objc.Call[objc.Void](p_, objc.Sel("setRelinquishFunction:"), value)
}

// The function used to acquire memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1410537-acquirefunction?language=objc
func (p_ PointerFunctions) AcquireFunction() func(src unsafe.Pointer, arg1 func(item unsafe.Pointer) uint, shouldCopy bool) unsafe.Pointer {
	rv := objc.Call[func(src unsafe.Pointer, arg1 func(item unsafe.Pointer) uint, shouldCopy bool) unsafe.Pointer](p_, objc.Sel("acquireFunction"))
	return rv
}

// The function used to acquire memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/1410537-acquirefunction?language=objc
func (p_ PointerFunctions) SetAcquireFunction(value func(src unsafe.Pointer, arg1 func(item unsafe.Pointer) uint, shouldCopy bool) unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("setAcquireFunction:"), value)
}
