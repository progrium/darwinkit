// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableArray] class.
var MutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}

type _MutableArrayClass struct {
	objc.Class
}

// An interface definition for the [MutableArray] class.
type IMutableArray interface {
	IArray
	SortUsingDescriptors(sortDescriptors []ISortDescriptor)
	AddObjectsFromArray(otherArray []objc.IObject)
	RemoveObjectIdenticalTo(anObject objc.IObject)
	InitWithContentsOfFile(path string) MutableArray
	SortUsingFunctionContext(compare func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer)
	RemoveLastObject()
	FilterUsingPredicate(predicate IPredicate)
	ReplaceObjectsInRangeWithObjectsFromArray(range_ Range, otherArray []objc.IObject)
	AddObject(anObject objc.IObject)
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	InsertObjectAtIndex(anObject objc.IObject, index uint)
	ApplyDifference(difference IOrderedCollectionDifference)
	SetObjectAtIndexedSubscript(obj objc.IObject, idx uint)
	SortUsingSelector(comparator objc.Selector)
	RemoveObject(anObject objc.IObject)
	RemoveObjectAtIndex(index uint)
	RemoveObjectsInArray(otherArray []objc.IObject)
	SetArray(otherArray []objc.IObject)
	InitWithContentsOfURL(url IURL) MutableArray
	RemoveObjectsInRange(range_ Range)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.IObject)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	SortUsingComparator(cmptr Comparator)
	InsertObjectsAtIndexes(objects []objc.IObject, indexes IIndexSet)
	SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator)
	RemoveAllObjects()
	ReplaceObjectAtIndexWithObject(index uint, anObject objc.IObject)
}

// A dynamic ordered collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray?language=objc
type MutableArray struct {
	Array
}

func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: ArrayFrom(ptr),
	}
}

func (mc _MutableArrayClass) ArrayWithCapacity(numItems uint) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithCapacity:"), numItems)
	return rv
}

// Creates and returns an NSMutableArray object with enough allocated memory to initially hold a given number of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460057-arraywithcapacity?language=objc
func MutableArray_ArrayWithCapacity(numItems uint) MutableArray {
	return MutableArrayClass.ArrayWithCapacity(numItems)
}

func (m_ MutableArray) InitWithCapacity(numItems uint) MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Returns an array, initialized with enough memory to initially hold a given number of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1415811-initwithcapacity?language=objc
func NewMutableArrayWithCapacity(numItems uint) MutableArray {
	instance := MutableArrayClass.Alloc().InitWithCapacity(numItems)
	instance.Autorelease()
	return instance
}

func (m_ MutableArray) Init() MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableArrayClass) Alloc() MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("alloc"))
	return rv
}

func MutableArray_Alloc() MutableArray {
	return MutableArrayClass.Alloc()
}

func (mc _MutableArrayClass) New() MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableArray() MutableArray {
	return MutableArrayClass.New()
}

func (m_ MutableArray) InitWithObjects(firstObj objc.IObject, args ...any) MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated array by placing in it the objects in the argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460068-initwithobjects?language=objc
func NewMutableArrayWithObjects(firstObj objc.IObject, args ...any) MutableArray {
	instance := MutableArrayClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (mc _MutableArrayClass) Array() MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("array"))
	return rv
}

// Creates and returns an empty array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460120-array?language=objc
func MutableArray_Array() MutableArray {
	return MutableArrayClass.Array()
}

func (mc _MutableArrayClass) ArrayWithObjects(firstObj objc.IObject, args ...any) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Creates and returns an array containing the objects in the argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460145-arraywithobjects?language=objc
func MutableArray_ArrayWithObjects(firstObj objc.IObject, args ...any) MutableArray {
	return MutableArrayClass.ArrayWithObjects(firstObj, args...)
}

func (m_ MutableArray) InitWithArray(array []objc.IObject) MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("initWithArray:"), array)
	return rv
}

// Initializes a newly allocated array by placing in it the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412169-initwitharray?language=objc
func NewMutableArrayWithArray(array []objc.IObject) MutableArray {
	instance := MutableArrayClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (mc _MutableArrayClass) ArrayWithArray(array []objc.IObject) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithArray:"), array)
	return rv
}

// Creates and returns an array containing the objects in another given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460122-arraywitharray?language=objc
func MutableArray_ArrayWithArray(array []objc.IObject) MutableArray {
	return MutableArrayClass.ArrayWithArray(array)
}

func (mc _MutableArrayClass) ArrayWithObject(anObject objc.IObject) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithObject:"), objc.Ptr(anObject))
	return rv
}

// Creates and returns an array containing a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411981-arraywithobject?language=objc
func MutableArray_ArrayWithObject(anObject objc.IObject) MutableArray {
	return MutableArrayClass.ArrayWithObject(anObject)
}

// Sorts the receiver using a given array of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1410745-sortusingdescriptors?language=objc
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors []ISortDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}

// Adds the objects contained in another given array to the end of the receiving array’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1408963-addobjectsfromarray?language=objc
func (m_ MutableArray) AddObjectsFromArray(otherArray []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObjectsFromArray:"), otherArray)
}

// Removes all occurrences of a given object in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1417759-removeobjectidenticalto?language=objc
func (m_ MutableArray) RemoveObjectIdenticalTo(anObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectIdenticalTo:"), objc.Ptr(anObject))
}

// Initializes a newly allocated mutable array with the contents of the file specified by a given path [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1414670-initwithcontentsoffile?language=objc
func (m_ MutableArray) InitWithContentsOfFile(path string) MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("initWithContentsOfFile:"), path)
	return rv
}

// Sorts the receiver in ascending order as defined by the comparison function compare. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1408332-sortusingfunction?language=objc
func (m_ MutableArray) SortUsingFunctionContext(compare func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingFunction:context:"), compare, context)
}

// Removes the object with the highest-valued index in the array [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1409708-removelastobject?language=objc
func (m_ MutableArray) RemoveLastObject() {
	objc.Call[objc.Void](m_, objc.Sel("removeLastObject"))
}

// Evaluates a given predicate against the array’s content and leaves only objects that match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1412085-filterusingpredicate?language=objc
func (m_ MutableArray) FilterUsingPredicate(predicate IPredicate) {
	objc.Call[objc.Void](m_, objc.Sel("filterUsingPredicate:"), objc.Ptr(predicate))
}

// Replaces the objects in the receiving array specified by a given range with all of the objects from a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1416902-replaceobjectsinrange?language=objc
func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ Range, otherArray []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectsInRange:withObjectsFromArray:"), range_, otherArray)
}

// Inserts a given object at the end of the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1411274-addobject?language=objc
func (m_ MutableArray) AddObject(anObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObject:"), objc.Ptr(anObject))
}

// Exchanges the objects in the array at given indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1411160-exchangeobjectatindex?language=objc
func (m_ MutableArray) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Call[objc.Void](m_, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}

// Inserts a given object into the array’s contents at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1416682-insertobject?language=objc
func (m_ MutableArray) InsertObjectAtIndex(anObject objc.IObject, index uint) {
	objc.Call[objc.Void](m_, objc.Sel("insertObject:atIndex:"), objc.Ptr(anObject), index)
}

// Creates and returns a mutable array containing the contents specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460070-arraywithcontentsofurl?language=objc
func (mc _MutableArrayClass) ArrayWithContentsOfURL(url IURL) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns a mutable array containing the contents specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460070-arraywithcontentsofurl?language=objc
func MutableArray_ArrayWithContentsOfURL(url IURL) MutableArray {
	return MutableArrayClass.ArrayWithContentsOfURL(url)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/3152169-applydifference?language=objc
func (m_ MutableArray) ApplyDifference(difference IOrderedCollectionDifference) {
	objc.Call[objc.Void](m_, objc.Sel("applyDifference:"), objc.Ptr(difference))
}

// Replaces the object at the index with the new object, possibly adding the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460093-setobject?language=objc
func (m_ MutableArray) SetObjectAtIndexedSubscript(obj objc.IObject, idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(obj), idx)
}

// Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1412273-sortusingselector?language=objc
func (m_ MutableArray) SortUsingSelector(comparator objc.Selector) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingSelector:"), comparator)
}

// Removes all occurrences in the array of a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1410689-removeobject?language=objc
func (m_ MutableArray) RemoveObject(anObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObject:"), objc.Ptr(anObject))
}

// Removes the object at index . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1416616-removeobjectatindex?language=objc
func (m_ MutableArray) RemoveObjectAtIndex(index uint) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectAtIndex:"), index)
}

// Removes from the receiving array the objects in another given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1413942-removeobjectsinarray?language=objc
func (m_ MutableArray) RemoveObjectsInArray(otherArray []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsInArray:"), otherArray)
}

// Sets the receiving array’s elements to those in another given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1417821-setarray?language=objc
func (m_ MutableArray) SetArray(otherArray []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setArray:"), otherArray)
}

// Initialized a newly allocated mutable array with the contents of the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1411688-initwithcontentsofurl?language=objc
func (m_ MutableArray) InitWithContentsOfURL(url IURL) MutableArray {
	rv := objc.Call[MutableArray](m_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Removes from the array each of the objects within a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1407586-removeobjectsinrange?language=objc
func (m_ MutableArray) RemoveObjectsInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsInRange:"), range_)
}

// Replaces the objects in the receiving array at locations specified with the objects from a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1418287-replaceobjectsatindexes?language=objc
func (m_ MutableArray) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectsAtIndexes:withObjects:"), objc.Ptr(indexes), objects)
}

// Removes the objects at the specified indexes from the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1410154-removeobjectsatindexes?language=objc
func (m_ MutableArray) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsAtIndexes:"), objc.Ptr(indexes))
}

// Sorts the receiver in ascending order using the comparison method specified by a given NSComparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1414904-sortusingcomparator?language=objc
func (m_ MutableArray) SortUsingComparator(cmptr Comparator) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingComparator:"), cmptr)
}

// Inserts the objects in the provided array into the receiving array at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1416482-insertobjects?language=objc
func (m_ MutableArray) InsertObjectsAtIndexes(objects []objc.IObject, indexes IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("insertObjects:atIndexes:"), objects, objc.Ptr(indexes))
}

// Creates and returns a mutable array containing the contents of the file specified by the given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460079-arraywithcontentsoffile?language=objc
func (mc _MutableArrayClass) ArrayWithContentsOfFile(path string) MutableArray {
	rv := objc.Call[MutableArray](mc, objc.Sel("arrayWithContentsOfFile:"), path)
	return rv
}

// Creates and returns a mutable array containing the contents of the file specified by the given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1460079-arraywithcontentsoffile?language=objc
func MutableArray_ArrayWithContentsOfFile(path string) MutableArray {
	return MutableArrayClass.ArrayWithContentsOfFile(path)
}

// Sorts the receiver in ascending order using the specified options and the comparison method specified by a given NSComparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1414396-sortwithoptions?language=objc
func (m_ MutableArray) SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) {
	objc.Call[objc.Void](m_, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}

// Empties the array of all its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1410618-removeallobjects?language=objc
func (m_ MutableArray) RemoveAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllObjects"))
}

// Replaces the object at index with anObject. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablearray/1414510-replaceobjectatindex?language=objc
func (m_ MutableArray) ReplaceObjectAtIndexWithObject(index uint, anObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectAtIndex:withObject:"), index, objc.Ptr(anObject))
}
