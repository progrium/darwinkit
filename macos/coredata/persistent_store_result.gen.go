// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentStoreResult] class.
var PersistentStoreResultClass = _PersistentStoreResultClass{objc.GetClass("NSPersistentStoreResult")}

type _PersistentStoreResultClass struct {
	objc.Class
}

// An interface definition for the [PersistentStoreResult] class.
type IPersistentStoreResult interface {
	objc.IObject
}

// The abstract base class for results returned from a persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreresult?language=objc
type PersistentStoreResult struct {
	objc.Object
}

func PersistentStoreResultFrom(ptr unsafe.Pointer) PersistentStoreResult {
	return PersistentStoreResult{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentStoreResultClass) Alloc() PersistentStoreResult {
	rv := objc.Call[PersistentStoreResult](pc, objc.Sel("alloc"))
	return rv
}

func PersistentStoreResult_Alloc() PersistentStoreResult {
	return PersistentStoreResultClass.Alloc()
}

func (pc _PersistentStoreResultClass) New() PersistentStoreResult {
	rv := objc.Call[PersistentStoreResult](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStoreResult() PersistentStoreResult {
	return PersistentStoreResultClass.New()
}

func (p_ PersistentStoreResult) Init() PersistentStoreResult {
	rv := objc.Call[PersistentStoreResult](p_, objc.Sel("init"))
	return rv
}
