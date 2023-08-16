// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecursiveLock] class.
var RecursiveLockClass = _RecursiveLockClass{objc.GetClass("NSRecursiveLock")}

type _RecursiveLockClass struct {
	objc.Class
}

// An interface definition for the [RecursiveLock] class.
type IRecursiveLock interface {
	objc.IObject
	TryLock() bool
	LockBeforeDate(limit IDate) bool
	Name() string
	SetName(value string)
}

// A lock that may be acquired multiple times by the same thread without causing a deadlock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecursivelock?language=objc
type RecursiveLock struct {
	objc.Object
}

func RecursiveLockFrom(ptr unsafe.Pointer) RecursiveLock {
	return RecursiveLock{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RecursiveLockClass) Alloc() RecursiveLock {
	rv := objc.Call[RecursiveLock](rc, objc.Sel("alloc"))
	return rv
}

func RecursiveLock_Alloc() RecursiveLock {
	return RecursiveLockClass.Alloc()
}

func (rc _RecursiveLockClass) New() RecursiveLock {
	rv := objc.Call[RecursiveLock](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecursiveLock() RecursiveLock {
	return RecursiveLockClass.New()
}

func (r_ RecursiveLock) Init() RecursiveLock {
	rv := objc.Call[RecursiveLock](r_, objc.Sel("init"))
	return rv
}

// Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecursivelock/1411547-trylock?language=objc
func (r_ RecursiveLock) TryLock() bool {
	rv := objc.Call[bool](r_, objc.Sel("tryLock"))
	return rv
}

// Attempts to acquire a lock before a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecursivelock/1417151-lockbeforedate?language=objc
func (r_ RecursiveLock) LockBeforeDate(limit IDate) bool {
	rv := objc.Call[bool](r_, objc.Sel("lockBeforeDate:"), objc.Ptr(limit))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecursivelock/1416232-name?language=objc
func (r_ RecursiveLock) Name() string {
	rv := objc.Call[string](r_, objc.Sel("name"))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecursivelock/1416232-name?language=objc
func (r_ RecursiveLock) SetName(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setName:"), value)
}
