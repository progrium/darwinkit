// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentHistoryToken] class.
var PersistentHistoryTokenClass = _PersistentHistoryTokenClass{objc.GetClass("NSPersistentHistoryToken")}

type _PersistentHistoryTokenClass struct {
	objc.Class
}

// An interface definition for the [PersistentHistoryToken] class.
type IPersistentHistoryToken interface {
	objc.IObject
}

// A bookmark for keeping track the most recent history that you’ve processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytoken?language=objc
type PersistentHistoryToken struct {
	objc.Object
}

func PersistentHistoryTokenFrom(ptr unsafe.Pointer) PersistentHistoryToken {
	return PersistentHistoryToken{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentHistoryTokenClass) Alloc() PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](pc, objc.Sel("alloc"))
	return rv
}

func PersistentHistoryToken_Alloc() PersistentHistoryToken {
	return PersistentHistoryTokenClass.Alloc()
}

func (pc _PersistentHistoryTokenClass) New() PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentHistoryToken() PersistentHistoryToken {
	return PersistentHistoryTokenClass.New()
}

func (p_ PersistentHistoryToken) Init() PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](p_, objc.Sel("init"))
	return rv
}
