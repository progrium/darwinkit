// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MergeConflict] class.
var MergeConflictClass = _MergeConflictClass{objc.GetClass("NSMergeConflict")}

type _MergeConflictClass struct {
	objc.Class
}

// An interface definition for the [MergeConflict] class.
type IMergeConflict interface {
	objc.IObject
	CachedSnapshot() map[string]objc.Object
	OldVersionNumber() uint
	SourceObject() ManagedObject
	NewVersionNumber() uint
	ObjectSnapshot() map[string]objc.Object
	PersistedSnapshot() map[string]objc.Object
}

// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict?language=objc
type MergeConflict struct {
	objc.Object
}

func MergeConflictFrom(ptr unsafe.Pointer) MergeConflict {
	return MergeConflict{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MergeConflict) InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject IManagedObject, newvers uint, oldvers uint, cachesnap map[string]objc.IObject, persnap map[string]objc.IObject) MergeConflict {
	rv := objc.Call[MergeConflict](m_, objc.Sel("initWithSource:newVersion:oldVersion:cachedSnapshot:persistedSnapshot:"), objc.Ptr(srcObject), newvers, oldvers, cachesnap, persnap)
	return rv
}

// Initializes a merge conflict. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506216-initwithsource?language=objc
func MergeConflict_InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject IManagedObject, newvers uint, oldvers uint, cachesnap map[string]objc.IObject, persnap map[string]objc.IObject) MergeConflict {
	return MergeConflictClass.Alloc().InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject, newvers, oldvers, cachesnap, persnap)
}

func (mc _MergeConflictClass) Alloc() MergeConflict {
	rv := objc.Call[MergeConflict](mc, objc.Sel("alloc"))
	return rv
}

func MergeConflict_Alloc() MergeConflict {
	return MergeConflictClass.Alloc()
}

func (mc _MergeConflictClass) New() MergeConflict {
	rv := objc.Call[MergeConflict](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMergeConflict() MergeConflict {
	return MergeConflictClass.New()
}

func (m_ MergeConflict) Init() MergeConflict {
	rv := objc.Call[MergeConflict](m_, objc.Sel("init"))
	return rv
}

// A dictionary containing the values of the source object held in the persistent store coordinator layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506685-cachedsnapshot?language=objc
func (m_ MergeConflict) CachedSnapshot() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("cachedSnapshot"))
	return rv
}

// The old version number for the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506271-oldversionnumber?language=objc
func (m_ MergeConflict) OldVersionNumber() uint {
	rv := objc.Call[uint](m_, objc.Sel("oldVersionNumber"))
	return rv
}

// The source object for the conflict. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506809-sourceobject?language=objc
func (m_ MergeConflict) SourceObject() ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("sourceObject"))
	return rv
}

// The new version number for the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506190-newversionnumber?language=objc
func (m_ MergeConflict) NewVersionNumber() uint {
	rv := objc.Call[uint](m_, objc.Sel("newVersionNumber"))
	return rv
}

// A dictionary containing the values of the source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506454-objectsnapshot?language=objc
func (m_ MergeConflict) ObjectSnapshot() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("objectSnapshot"))
	return rv
}

// A dictionary containing the values of the source object held in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/1506412-persistedsnapshot?language=objc
func (m_ MergeConflict) PersistedSnapshot() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("persistedSnapshot"))
	return rv
}
