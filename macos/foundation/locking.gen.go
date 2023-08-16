// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The elementary methods adopted by classes that define lock objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocking?language=objc
type PLocking interface {
	// optional
	Lock()
	HasLock() bool

	// optional
	Unlock()
	HasUnlock() bool
}

// A concrete type wrapper for the [PLocking] protocol.
type LockingWrapper struct {
	objc.Object
}

func (l_ LockingWrapper) HasLock() bool {
	return l_.RespondsToSelector(objc.Sel("lock"))
}

// Attempts to acquire a lock, blocking a thread’s execution until the lock can be acquired. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocking/1416318-lock?language=objc
func (l_ LockingWrapper) Lock() {
	objc.Call[objc.Void](l_, objc.Sel("lock"))
}

func (l_ LockingWrapper) HasUnlock() bool {
	return l_.RespondsToSelector(objc.Sel("unlock"))
}

// Relinquishes a previously acquired lock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocking/1418241-unlock?language=objc
func (l_ LockingWrapper) Unlock() {
	objc.Call[objc.Void](l_, objc.Sel("unlock"))
}
