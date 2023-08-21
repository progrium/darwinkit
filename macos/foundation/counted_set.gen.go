// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CountedSet] class.
var CountedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}

type _CountedSetClass struct {
	objc.Class
}

// An interface definition for the [CountedSet] class.
type ICountedSet interface {
	IMutableSet
	CountForObject(object objc.IObject) uint
}

// A mutable, unordered collection of distinct objects that may appear more than once in the collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountedset?language=objc
type CountedSet struct {
	MutableSet
}

func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		MutableSet: MutableSetFrom(ptr),
	}
}

func (c_ CountedSet) InitWithCapacity(numItems uint) CountedSet {
	rv := objc.Call[CountedSet](c_, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Returns a counted set object initialized with enough memory to hold a given number of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountedset/1415625-initwithcapacity?language=objc
func NewCountedSetWithCapacity(numItems uint) CountedSet {
	instance := CountedSetClass.Alloc().InitWithCapacity(numItems)
	instance.Autorelease()
	return instance
}

func (c_ CountedSet) InitWithArray(array []objc.IObject) CountedSet {
	rv := objc.Call[CountedSet](c_, objc.Sel("initWithArray:"), array)
	return rv
}

// Returns a counted set object initialized with the contents of a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountedset/1416767-initwitharray?language=objc
func NewCountedSetWithArray(array []objc.IObject) CountedSet {
	instance := CountedSetClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (c_ CountedSet) InitWithSet(set ISet) CountedSet {
	rv := objc.Call[CountedSet](c_, objc.Sel("initWithSet:"), objc.Ptr(set))
	return rv
}

// Returns a counted set object initialized with the contents of a given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountedset/1411730-initwithset?language=objc
func NewCountedSetWithSet(set ISet) CountedSet {
	instance := CountedSetClass.Alloc().InitWithSet(set)
	instance.Autorelease()
	return instance
}

func (cc _CountedSetClass) Alloc() CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("alloc"))
	return rv
}

func CountedSet_Alloc() CountedSet {
	return CountedSetClass.Alloc()
}

func (cc _CountedSetClass) New() CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCountedSet() CountedSet {
	return CountedSetClass.New()
}

func (c_ CountedSet) Init() CountedSet {
	rv := objc.Call[CountedSet](c_, objc.Sel("init"))
	return rv
}

func (cc _CountedSetClass) SetWithCapacity(numItems uint) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithCapacity:"), numItems)
	return rv
}

// Creates and returns a mutable set with a given initial capacity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableset/1574820-setwithcapacity?language=objc
func CountedSet_SetWithCapacity(numItems uint) CountedSet {
	return CountedSetClass.SetWithCapacity(numItems)
}

func (cc _CountedSetClass) Set() CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("set"))
	return rv
}

// Creates and returns an empty set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574818-set?language=objc
func CountedSet_Set() CountedSet {
	return CountedSetClass.Set()
}

func (c_ CountedSet) InitWithObjects(firstObj objc.IObject, args ...any) CountedSet {
	rv := objc.Call[CountedSet](c_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574822-initwithobjects?language=objc
func NewCountedSetWithObjects(firstObj objc.IObject, args ...any) CountedSet {
	instance := CountedSetClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (cc _CountedSetClass) SetWithSet(set ISet) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns a set containing the objects from another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574817-setwithset?language=objc
func CountedSet_SetWithSet(set ISet) CountedSet {
	return CountedSetClass.SetWithSet(set)
}

func (cc _CountedSetClass) SetWithObject(object objc.IObject) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithObject:"), objc.Ptr(object))
	return rv
}

// Creates and returns a set that contains a single given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1415878-setwithobject?language=objc
func CountedSet_SetWithObject(object objc.IObject) CountedSet {
	return CountedSetClass.SetWithObject(object)
}

func (cc _CountedSetClass) SetWithArray(array []objc.IObject) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithArray:"), array)
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574823-setwitharray?language=objc
func CountedSet_SetWithArray(array []objc.IObject) CountedSet {
	return CountedSetClass.SetWithArray(array)
}

func (cc _CountedSetClass) SetWithCollectionViewIndexPath(indexPath IIndexPath) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithCollectionViewIndexPath:"), objc.Ptr(indexPath))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528161-setwithcollectionviewindexpath?language=objc
func CountedSet_SetWithCollectionViewIndexPath(indexPath IIndexPath) CountedSet {
	return CountedSetClass.SetWithCollectionViewIndexPath(indexPath)
}

func (cc _CountedSetClass) SetWithObjectsCount(objects objc.IObject, cnt uint) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithObjects:count:"), objc.Ptr(objects), cnt)
	return rv
}

// Creates and returns a set containing a specified number of objects from a given C array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574824-setwithobjects?language=objc
func CountedSet_SetWithObjectsCount(objects objc.IObject, cnt uint) CountedSet {
	return CountedSetClass.SetWithObjectsCount(objects, cnt)
}

func (cc _CountedSetClass) SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) CountedSet {
	rv := objc.Call[CountedSet](cc, objc.Sel("setWithCollectionViewIndexPaths:"), indexPaths)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528255-setwithcollectionviewindexpaths?language=objc
func CountedSet_SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) CountedSet {
	return CountedSetClass.SetWithCollectionViewIndexPaths(indexPaths)
}

// Returns the count associated with a given object in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountedset/1408658-countforobject?language=objc
func (c_ CountedSet) CountForObject(object objc.IObject) uint {
	rv := objc.Call[uint](c_, objc.Sel("countForObject:"), objc.Ptr(object))
	return rv
}
