// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ManagedObjectID] class.
var ManagedObjectIDClass = _ManagedObjectIDClass{objc.GetClass("NSManagedObjectID")}

type _ManagedObjectIDClass struct {
	objc.Class
}

// An interface definition for the [ManagedObjectID] class.
type IManagedObjectID interface {
	objc.IObject
	URIRepresentation() foundation.URL
	Entity() EntityDescription
	PersistentStore() PersistentStore
	IsTemporaryID() bool
}

// A compact, universal identifier for a managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid?language=objc
type ManagedObjectID struct {
	objc.Object
}

func ManagedObjectIDFrom(ptr unsafe.Pointer) ManagedObjectID {
	return ManagedObjectID{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ManagedObjectIDClass) Alloc() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](mc, objc.Sel("alloc"))
	return rv
}

func ManagedObjectID_Alloc() ManagedObjectID {
	return ManagedObjectIDClass.Alloc()
}

func (mc _ManagedObjectIDClass) New() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewManagedObjectID() ManagedObjectID {
	return ManagedObjectIDClass.New()
}

func (m_ ManagedObjectID) Init() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](m_, objc.Sel("init"))
	return rv
}

// Returns a URI that provides an archiveable reference to the object for the object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/1391689-urirepresentation?language=objc
func (m_ ManagedObjectID) URIRepresentation() foundation.URL {
	rv := objc.Call[foundation.URL](m_, objc.Sel("URIRepresentation"))
	return rv
}

// The entity description associated with the object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/1391684-entity?language=objc
func (m_ ManagedObjectID) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](m_, objc.Sel("entity"))
	return rv
}

// The persistent store that fetched the object for the object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/1391693-persistentstore?language=objc
func (m_ ManagedObjectID) PersistentStore() PersistentStore {
	rv := objc.Call[PersistentStore](m_, objc.Sel("persistentStore"))
	return rv
}

// A Boolean value that indicates whether the object ID is temporary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectid/1391691-temporaryid?language=objc
func (m_ ManagedObjectID) IsTemporaryID() bool {
	rv := objc.Call[bool](m_, objc.Sel("isTemporaryID"))
	return rv
}
