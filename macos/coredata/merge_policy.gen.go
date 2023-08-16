// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MergePolicy] class.
var MergePolicyClass = _MergePolicyClass{objc.GetClass("NSMergePolicy")}

type _MergePolicyClass struct {
	objc.Class
}

// An interface definition for the [MergePolicy] class.
type IMergePolicy interface {
	objc.IObject
	ResolveConstraintConflictsError(list []IConstraintConflict, error foundation.IError) bool
	ResolveConflictsError(list []objc.IObject, error foundation.IError) bool
	InitWithMergeType(ty MergePolicyType) objc.Object
	ResolveOptimisticLockingVersionConflictsError(list []IMergeConflict, error foundation.IError) bool
	MergeType() MergePolicyType
}

// A policy object that you use to resolve conflicts between the persistent store and in-memory versions of managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy?language=objc
type MergePolicy struct {
	objc.Object
}

func MergePolicyFrom(ptr unsafe.Pointer) MergePolicy {
	return MergePolicy{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MergePolicyClass) Alloc() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("alloc"))
	return rv
}

func MergePolicy_Alloc() MergePolicy {
	return MergePolicyClass.Alloc()
}

func (mc _MergePolicyClass) New() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMergePolicy() MergePolicy {
	return MergePolicyClass.New()
}

func (m_ MergePolicy) Init() MergePolicy {
	rv := objc.Call[MergePolicy](m_, objc.Sel("init"))
	return rv
}

// Resolves the conflicts in a given list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1506678-resolveconstraintconflicts?language=objc
func (m_ MergePolicy) ResolveConstraintConflictsError(list []IConstraintConflict, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("resolveConstraintConflicts:error:"), list, objc.Ptr(error))
	return rv
}

// Resolves the conflicts in a given list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1506253-resolveconflicts?language=objc
func (m_ MergePolicy) ResolveConflictsError(list []objc.IObject, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("resolveConflicts:error:"), list, objc.Ptr(error))
	return rv
}

// Returns a merge policy initialized with a given policy type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1506763-initwithmergetype?language=objc
func (m_ MergePolicy) InitWithMergeType(ty MergePolicyType) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("initWithMergeType:"), ty)
	return rv
}

// Resolves the conflicts in a given list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1506787-resolveoptimisticlockingversionc?language=objc
func (m_ MergePolicy) ResolveOptimisticLockingVersionConflictsError(list []IMergeConflict, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("resolveOptimisticLockingVersionConflicts:error:"), list, objc.Ptr(error))
	return rv
}

// A property-based merge policy that applies in-memory changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690607-mergebypropertyobjecttrumpmergep?language=objc
func (mc _MergePolicyClass) MergeByPropertyObjectTrumpMergePolicy() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("mergeByPropertyObjectTrumpMergePolicy"))
	return rv
}

// A property-based merge policy that applies in-memory changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690607-mergebypropertyobjecttrumpmergep?language=objc
func MergePolicy_MergeByPropertyObjectTrumpMergePolicy() MergePolicy {
	return MergePolicyClass.MergeByPropertyObjectTrumpMergePolicy()
}

// The merge type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1506675-mergetype?language=objc
func (m_ MergePolicy) MergeType() MergePolicyType {
	rv := objc.Call[MergePolicyType](m_, objc.Sel("mergeType"))
	return rv
}

// A property-based merge policy that applies external changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690609-mergebypropertystoretrumpmergepo?language=objc
func (mc _MergePolicyClass) MergeByPropertyStoreTrumpMergePolicy() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("mergeByPropertyStoreTrumpMergePolicy"))
	return rv
}

// A property-based merge policy that applies external changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690609-mergebypropertystoretrumpmergepo?language=objc
func MergePolicy_MergeByPropertyStoreTrumpMergePolicy() MergePolicy {
	return MergePolicyClass.MergeByPropertyStoreTrumpMergePolicy()
}

// A merge policy that discards unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690613-rollbackmergepolicy?language=objc
func (mc _MergePolicyClass) RollbackMergePolicy() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("rollbackMergePolicy"))
	return rv
}

// A merge policy that discards unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690613-rollbackmergepolicy?language=objc
func MergePolicy_RollbackMergePolicy() MergePolicy {
	return MergePolicyClass.RollbackMergePolicy()
}

// The default merge policy for all managed object contexts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690612-errormergepolicy?language=objc
func (mc _MergePolicyClass) ErrorMergePolicy() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("errorMergePolicy"))
	return rv
}

// The default merge policy for all managed object contexts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690612-errormergepolicy?language=objc
func MergePolicy_ErrorMergePolicy() MergePolicy {
	return MergePolicyClass.ErrorMergePolicy()
}

// A merge policy that overwrites the entire stored object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690610-overwritemergepolicy?language=objc
func (mc _MergePolicyClass) OverwriteMergePolicy() MergePolicy {
	rv := objc.Call[MergePolicy](mc, objc.Sel("overwriteMergePolicy"))
	return rv
}

// A merge policy that overwrites the entire stored object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicy/1690610-overwritemergepolicy?language=objc
func MergePolicy_OverwriteMergePolicy() MergePolicy {
	return MergePolicyClass.OverwriteMergePolicy()
}
