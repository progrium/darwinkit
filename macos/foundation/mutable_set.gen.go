// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableSet] class.
var MutableSetClass = _MutableSetClass{objc.GetClass("NSMutableSet")}

type _MutableSetClass struct {
	objc.Class
}

// An interface definition for the [MutableSet] class.
type IMutableSet interface {
	ISet
	AddObjectsFromArray(array []objc.IObject)
	IntersectSet(otherSet ISet)
	FilterUsingPredicate(predicate IPredicate)
	SetSet(otherSet ISet)
	AddObject(object objc.IObject)
	RemoveObject(object objc.IObject)
	MinusSet(otherSet ISet)
	UnionSet(otherSet ISet)
	RemoveAllObjects()
}

// A dynamic unordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset?language=objc
type MutableSet struct {
	Set
}

func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		Set: SetFrom(ptr),
	}
}

func (m_ MutableSet) InitWithCapacity(numItems uint) MutableSet {
	rv := objc.Call[MutableSet](m_, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Returns an initialized mutable set with a given initial capacity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1411953-initwithcapacity?language=objc
func NewMutableSetWithCapacity(numItems uint) MutableSet {
	instance := MutableSetClass.Alloc().InitWithCapacity(numItems)
	instance.Autorelease()
	return instance
}

func (mc _MutableSetClass) SetWithCapacity(numItems uint) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithCapacity:"), numItems)
	return rv
}

// Creates and returns a mutable set with a given initial capacity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1574820-setwithcapacity?language=objc
func MutableSet_SetWithCapacity(numItems uint) MutableSet {
	return MutableSetClass.SetWithCapacity(numItems)
}

func (m_ MutableSet) Init() MutableSet {
	rv := objc.Call[MutableSet](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableSetClass) Alloc() MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("alloc"))
	return rv
}

func MutableSet_Alloc() MutableSet {
	return MutableSetClass.Alloc()
}

func (mc _MutableSetClass) New() MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableSet() MutableSet {
	return MutableSetClass.New()
}

func (mc _MutableSetClass) Set() MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("set"))
	return rv
}

// Creates and returns an empty set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574818-set?language=objc
func MutableSet_Set() MutableSet {
	return MutableSetClass.Set()
}

func (m_ MutableSet) InitWithObjects(firstObj objc.IObject, args ...any) MutableSet {
	rv := objc.Call[MutableSet](m_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574822-initwithobjects?language=objc
func NewMutableSetWithObjects(firstObj objc.IObject, args ...any) MutableSet {
	instance := MutableSetClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (mc _MutableSetClass) SetWithSet(set ISet) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns a set containing the objects from another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574817-setwithset?language=objc
func MutableSet_SetWithSet(set ISet) MutableSet {
	return MutableSetClass.SetWithSet(set)
}

func (mc _MutableSetClass) SetWithObject(object objc.IObject) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithObject:"), objc.Ptr(object))
	return rv
}

// Creates and returns a set that contains a single given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1415878-setwithobject?language=objc
func MutableSet_SetWithObject(object objc.IObject) MutableSet {
	return MutableSetClass.SetWithObject(object)
}

func (m_ MutableSet) InitWithArray(array []objc.IObject) MutableSet {
	rv := objc.Call[MutableSet](m_, objc.Sel("initWithArray:"), array)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1408152-initwitharray?language=objc
func NewMutableSetWithArray(array []objc.IObject) MutableSet {
	instance := MutableSetClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (mc _MutableSetClass) SetWithArray(array []objc.IObject) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithArray:"), array)
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574823-setwitharray?language=objc
func MutableSet_SetWithArray(array []objc.IObject) MutableSet {
	return MutableSetClass.SetWithArray(array)
}

func (mc _MutableSetClass) SetWithCollectionViewIndexPath(indexPath IIndexPath) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithCollectionViewIndexPath:"), objc.Ptr(indexPath))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528161-setwithcollectionviewindexpath?language=objc
func MutableSet_SetWithCollectionViewIndexPath(indexPath IIndexPath) MutableSet {
	return MutableSetClass.SetWithCollectionViewIndexPath(indexPath)
}

func (mc _MutableSetClass) SetWithObjectsCount(objects objc.IObject, cnt uint) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithObjects:count:"), objc.Ptr(objects), cnt)
	return rv
}

// Creates and returns a set containing a specified number of objects from a given C array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574824-setwithobjects?language=objc
func MutableSet_SetWithObjectsCount(objects objc.IObject, cnt uint) MutableSet {
	return MutableSetClass.SetWithObjectsCount(objects, cnt)
}

func (mc _MutableSetClass) SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) MutableSet {
	rv := objc.Call[MutableSet](mc, objc.Sel("setWithCollectionViewIndexPaths:"), indexPaths)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528255-setwithcollectionviewindexpaths?language=objc
func MutableSet_SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) MutableSet {
	return MutableSetClass.SetWithCollectionViewIndexPaths(indexPaths)
}

func (m_ MutableSet) InitWithSet(set ISet) MutableSet {
	rv := objc.Call[MutableSet](m_, objc.Sel("initWithSet:"), objc.Ptr(set))
	return rv
}

// Initializes a newly allocated set and adds to it objects from another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1410612-initwithset?language=objc
func NewMutableSetWithSet(set ISet) MutableSet {
	instance := MutableSetClass.Alloc().InitWithSet(set)
	instance.Autorelease()
	return instance
}

// Adds to the set each object contained in a given array that is not already a member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1414345-addobjectsfromarray?language=objc
func (m_ MutableSet) AddObjectsFromArray(array []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObjectsFromArray:"), array)
}

// Removes from the receiving set each object that isn’t a member of another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1407231-intersectset?language=objc
func (m_ MutableSet) IntersectSet(otherSet ISet) {
	objc.Call[objc.Void](m_, objc.Sel("intersectSet:"), objc.Ptr(otherSet))
}

// Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1407868-filterusingpredicate?language=objc
func (m_ MutableSet) FilterUsingPredicate(predicate IPredicate) {
	objc.Call[objc.Void](m_, objc.Sel("filterUsingPredicate:"), objc.Ptr(predicate))
}

// Empties the receiving set, then adds each object contained in another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1416405-setset?language=objc
func (m_ MutableSet) SetSet(otherSet ISet) {
	objc.Call[objc.Void](m_, objc.Sel("setSet:"), objc.Ptr(otherSet))
}

// Adds a given object to the set, if it is not already a member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1407460-addobject?language=objc
func (m_ MutableSet) AddObject(object objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObject:"), objc.Ptr(object))
}

// Removes a given object from the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1416085-removeobject?language=objc
func (m_ MutableSet) RemoveObject(object objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObject:"), objc.Ptr(object))
}

// Removes each object in another given set from the receiving set, if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1416710-minusset?language=objc
func (m_ MutableSet) MinusSet(otherSet ISet) {
	objc.Call[objc.Void](m_, objc.Sel("minusSet:"), objc.Ptr(otherSet))
}

// Adds each object in another given set to the receiving set, if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1413187-unionset?language=objc
func (m_ MutableSet) UnionSet(otherSet ISet) {
	objc.Call[objc.Void](m_, objc.Sel("unionSet:"), objc.Ptr(otherSet))
}

// Empties the set of all of its members. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1417497-removeallobjects?language=objc
func (m_ MutableSet) RemoveAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllObjects"))
}
