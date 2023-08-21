// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ConstraintConflict] class.
var ConstraintConflictClass = _ConstraintConflictClass{objc.GetClass("NSConstraintConflict")}

type _ConstraintConflictClass struct {
	objc.Class
}

// An interface definition for the [ConstraintConflict] class.
type IConstraintConflict interface {
	objc.IObject
	ConstraintValues() map[string]objc.Object
	DatabaseObject() ManagedObject
	Constraint() []string
	DatabaseSnapshot() map[string]objc.Object
	ConflictingSnapshots() []foundation.Dictionary
	ConflictingObjects() []ManagedObject
}

// An encapsulation of conflicts that occur during an attempt to save a managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict?language=objc
type ConstraintConflict struct {
	objc.Object
}

func ConstraintConflictFrom(ptr unsafe.Pointer) ConstraintConflict {
	return ConstraintConflict{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ConstraintConflict) InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint []string, databaseObject IManagedObject, databaseSnapshot foundation.Dictionary, conflictingObjects []IManagedObject, conflictingSnapshots []objc.IObject) ConstraintConflict {
	rv := objc.Call[ConstraintConflict](c_, objc.Sel("initWithConstraint:databaseObject:databaseSnapshot:conflictingObjects:conflictingSnapshots:"), contraint, objc.Ptr(databaseObject), databaseSnapshot, conflictingObjects, conflictingSnapshots)
	return rv
}

// Initializes a constraint conflict. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506668-initwithconstraint?language=objc
func NewConstraintConflictWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint []string, databaseObject IManagedObject, databaseSnapshot foundation.Dictionary, conflictingObjects []IManagedObject, conflictingSnapshots []objc.IObject) ConstraintConflict {
	instance := ConstraintConflictClass.Alloc().InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint, databaseObject, databaseSnapshot, conflictingObjects, conflictingSnapshots)
	instance.Autorelease()
	return instance
}

func (cc _ConstraintConflictClass) Alloc() ConstraintConflict {
	rv := objc.Call[ConstraintConflict](cc, objc.Sel("alloc"))
	return rv
}

func ConstraintConflict_Alloc() ConstraintConflict {
	return ConstraintConflictClass.Alloc()
}

func (cc _ConstraintConflictClass) New() ConstraintConflict {
	rv := objc.Call[ConstraintConflict](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConstraintConflict() ConstraintConflict {
	return ConstraintConflictClass.New()
}

func (c_ ConstraintConflict) Init() ConstraintConflict {
	rv := objc.Call[ConstraintConflict](c_, objc.Sel("init"))
	return rv
}

// The values that the conflicting objects had when the conflict was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506399-constraintvalues?language=objc
func (c_ ConstraintConflict) ConstraintValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("constraintValues"))
	return rv
}

// The object whose database row is using constraint values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506665-databaseobject?language=objc
func (c_ ConstraintConflict) DatabaseObject() ManagedObject {
	rv := objc.Call[ManagedObject](c_, objc.Sel("databaseObject"))
	return rv
}

// The constraint that has been violated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506902-constraint?language=objc
func (c_ ConstraintConflict) Constraint() []string {
	rv := objc.Call[[]string](c_, objc.Sel("constraint"))
	return rv
}

// The values currently stored in the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506687-databasesnapshot?language=objc
func (c_ ConstraintConflict) DatabaseSnapshot() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("databaseSnapshot"))
	return rv
}

// The original property values of objects in violation of the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506774-conflictingsnapshots?language=objc
func (c_ ConstraintConflict) ConflictingSnapshots() []foundation.Dictionary {
	rv := objc.Call[[]foundation.Dictionary](c_, objc.Sel("conflictingSnapshots"))
	return rv
}

// The managed objects that are in conflict. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/1506707-conflictingobjects?language=objc
func (c_ ConstraintConflict) ConflictingObjects() []ManagedObject {
	rv := objc.Call[[]ManagedObject](c_, objc.Sel("conflictingObjects"))
	return rv
}
