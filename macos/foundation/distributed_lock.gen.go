// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DistributedLock] class.
var DistributedLockClass = _DistributedLockClass{objc.GetClass("NSDistributedLock")}

type _DistributedLockClass struct {
	objc.Class
}

// An interface definition for the [DistributedLock] class.
type IDistributedLock interface {
	objc.IObject
	TryLock() bool
	Unlock()
	BreakLock()
	LockDate() Date
}

// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock?language=objc
type DistributedLock struct {
	objc.Object
}

func DistributedLockFrom(ptr unsafe.Pointer) DistributedLock {
	return DistributedLock{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DistributedLock) InitWithPath(path string) DistributedLock {
	rv := objc.Call[DistributedLock](d_, objc.Sel("initWithPath:"), path)
	return rv
}

// Initializes an NSDistributedLock object to use as the lock the file-system entry specified by a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1410387-initwithpath?language=objc
func DistributedLock_InitWithPath(path string) DistributedLock {
	return DistributedLockClass.Alloc().InitWithPath(path)
}

func (dc _DistributedLockClass) Alloc() DistributedLock {
	rv := objc.Call[DistributedLock](dc, objc.Sel("alloc"))
	return rv
}

func DistributedLock_Alloc() DistributedLock {
	return DistributedLockClass.Alloc()
}

func (dc _DistributedLockClass) New() DistributedLock {
	rv := objc.Call[DistributedLock](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDistributedLock() DistributedLock {
	return DistributedLockClass.New()
}

func (d_ DistributedLock) Init() DistributedLock {
	rv := objc.Call[DistributedLock](d_, objc.Sel("init"))
	return rv
}

// Returns an NSDistributedLock object initialized to use as the locking object the file-system entry specified by a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1552203-lockwithpath?language=objc
func (dc _DistributedLockClass) LockWithPath(path string) DistributedLock {
	rv := objc.Call[DistributedLock](dc, objc.Sel("lockWithPath:"), path)
	return rv
}

// Returns an NSDistributedLock object initialized to use as the locking object the file-system entry specified by a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1552203-lockwithpath?language=objc
func DistributedLock_LockWithPath(path string) DistributedLock {
	return DistributedLockClass.LockWithPath(path)
}

// Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1412293-trylock?language=objc
func (d_ DistributedLock) TryLock() bool {
	rv := objc.Call[bool](d_, objc.Sel("tryLock"))
	return rv
}

// Relinquishes the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1417238-unlock?language=objc
func (d_ DistributedLock) Unlock() {
	objc.Call[objc.Void](d_, objc.Sel("unlock"))
}

// Forces the lock to be relinquished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1413425-breaklock?language=objc
func (d_ DistributedLock) BreakLock() {
	objc.Call[objc.Void](d_, objc.Sel("breakLock"))
}

// Returns the time the receiver was acquired by any of the NSDistributedLock objects using the same path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributedlock/1413773-lockdate?language=objc
func (d_ DistributedLock) LockDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("lockDate"))
	return rv
}
