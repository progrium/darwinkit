// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Lock] class.
var LockClass = _LockClass{objc.GetClass("NSLock")}

type _LockClass struct {
	objc.Class
}

// An interface definition for the [Lock] class.
type ILock interface {
	objc.IObject
	TryLock() bool
	LockBeforeDate(limit IDate) bool
	Name() string
	SetName(value string)
}

// An object that coordinates the operation of multiple threads of execution within the same application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslock?language=objc
type Lock struct {
	objc.Object
}

func LockFrom(ptr unsafe.Pointer) Lock {
	return Lock{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LockClass) Alloc() Lock {
	rv := objc.Call[Lock](lc, objc.Sel("alloc"))
	return rv
}

func Lock_Alloc() Lock {
	return LockClass.Alloc()
}

func (lc _LockClass) New() Lock {
	rv := objc.Call[Lock](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLock() Lock {
	return LockClass.New()
}

func (l_ Lock) Init() Lock {
	rv := objc.Call[Lock](l_, objc.Sel("init"))
	return rv
}

// Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslock/1418105-trylock?language=objc
func (l_ Lock) TryLock() bool {
	rv := objc.Call[bool](l_, objc.Sel("tryLock"))
	return rv
}

// Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslock/1411133-lockbeforedate?language=objc
func (l_ Lock) LockBeforeDate(limit IDate) bool {
	rv := objc.Call[bool](l_, objc.Sel("lockBeforeDate:"), objc.Ptr(limit))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslock/1412568-name?language=objc
func (l_ Lock) Name() string {
	rv := objc.Call[string](l_, objc.Sel("name"))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslock/1412568-name?language=objc
func (l_ Lock) SetName(value string) {
	objc.Call[objc.Void](l_, objc.Sel("setName:"), value)
}
