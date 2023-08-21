// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableIndexSet] class.
var MutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}

type _MutableIndexSetClass struct {
	objc.Class
}

// An interface definition for the [MutableIndexSet] class.
type IMutableIndexSet interface {
	IIndexSet
	ShiftIndexesStartingAtIndexBy(index uint, delta int)
	RemoveIndex(value uint)
	RemoveIndexesInRange(range_ Range)
	RemoveIndexes(indexSet IIndexSet)
	RemoveAllIndexes()
	AddIndex(value uint)
	AddIndexes(indexSet IIndexSet)
	AddIndexesInRange(range_ Range)
}

// A mutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset?language=objc
type MutableIndexSet struct {
	IndexSet
}

func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		IndexSet: IndexSetFrom(ptr),
	}
}

func (mc _MutableIndexSetClass) Alloc() MutableIndexSet {
	rv := objc.Call[MutableIndexSet](mc, objc.Sel("alloc"))
	return rv
}

func MutableIndexSet_Alloc() MutableIndexSet {
	return MutableIndexSetClass.Alloc()
}

func (mc _MutableIndexSetClass) New() MutableIndexSet {
	rv := objc.Call[MutableIndexSet](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableIndexSet() MutableIndexSet {
	return MutableIndexSetClass.New()
}

func (m_ MutableIndexSet) Init() MutableIndexSet {
	rv := objc.Call[MutableIndexSet](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableIndexSetClass) IndexSetWithIndex(value uint) MutableIndexSet {
	rv := objc.Call[MutableIndexSet](mc, objc.Sel("indexSetWithIndex:"), value)
	return rv
}

// Creates an index set with an index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427254-indexsetwithindex?language=objc
func MutableIndexSet_IndexSetWithIndex(value uint) MutableIndexSet {
	return MutableIndexSetClass.IndexSetWithIndex(value)
}

func (mc _MutableIndexSetClass) IndexSet() MutableIndexSet {
	rv := objc.Call[MutableIndexSet](mc, objc.Sel("indexSet"))
	return rv
}

// Creates an empty index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427281-indexset?language=objc
func MutableIndexSet_IndexSet() MutableIndexSet {
	return MutableIndexSetClass.IndexSet()
}

func (m_ MutableIndexSet) InitWithIndexesInRange(range_ Range) MutableIndexSet {
	rv := objc.Call[MutableIndexSet](m_, objc.Sel("initWithIndexesInRange:"), range_)
	return rv
}

// Initializes an allocated NSIndexSet object with an index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1414013-initwithindexesinrange?language=objc
func NewMutableIndexSetWithIndexesInRange(range_ Range) MutableIndexSet {
	instance := MutableIndexSetClass.Alloc().InitWithIndexesInRange(range_)
	instance.Autorelease()
	return instance
}

func (m_ MutableIndexSet) InitWithIndexSet(indexSet IIndexSet) MutableIndexSet {
	rv := objc.Call[MutableIndexSet](m_, objc.Sel("initWithIndexSet:"), objc.Ptr(indexSet))
	return rv
}

// Initializes an allocated NSIndexSet object with an index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1415602-initwithindexset?language=objc
func NewMutableIndexSetWithIndexSet(indexSet IIndexSet) MutableIndexSet {
	instance := MutableIndexSetClass.Alloc().InitWithIndexSet(indexSet)
	instance.Autorelease()
	return instance
}

func (mc _MutableIndexSetClass) IndexSetWithIndexesInRange(range_ Range) MutableIndexSet {
	rv := objc.Call[MutableIndexSet](mc, objc.Sel("indexSetWithIndexesInRange:"), range_)
	return rv
}

// Creates an index set with an index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427274-indexsetwithindexesinrange?language=objc
func MutableIndexSet_IndexSetWithIndexesInRange(range_ Range) MutableIndexSet {
	return MutableIndexSetClass.IndexSetWithIndexesInRange(range_)
}

func (m_ MutableIndexSet) InitWithIndex(value uint) MutableIndexSet {
	rv := objc.Call[MutableIndexSet](m_, objc.Sel("initWithIndex:"), value)
	return rv
}

// Initializes an allocated NSIndexSet object with an index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1416501-initwithindex?language=objc
func NewMutableIndexSetWithIndex(value uint) MutableIndexSet {
	instance := MutableIndexSetClass.Alloc().InitWithIndex(value)
	instance.Autorelease()
	return instance
}

// Shifts a group of indexes to the left or the right within the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1416951-shiftindexesstartingatindex?language=objc
func (m_ MutableIndexSet) ShiftIndexesStartingAtIndexBy(index uint, delta int) {
	objc.Call[objc.Void](m_, objc.Sel("shiftIndexesStartingAtIndex:by:"), index, delta)
}

// Removes an index from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1410650-removeindex?language=objc
func (m_ MutableIndexSet) RemoveIndex(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("removeIndex:"), value)
}

// Removes the indexes in an index range from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1415791-removeindexesinrange?language=objc
func (m_ MutableIndexSet) RemoveIndexesInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeIndexesInRange:"), range_)
}

// Removes the indexes in an index set from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1412018-removeindexes?language=objc
func (m_ MutableIndexSet) RemoveIndexes(indexSet IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("removeIndexes:"), objc.Ptr(indexSet))
}

// Removes the receiver’s indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1408738-removeallindexes?language=objc
func (m_ MutableIndexSet) RemoveAllIndexes() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllIndexes"))
}

// Adds an index  to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1410712-addindex?language=objc
func (m_ MutableIndexSet) AddIndex(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("addIndex:"), value)
}

// Adds the indexes in an index set to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1414594-addindexes?language=objc
func (m_ MutableIndexSet) AddIndexes(indexSet IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("addIndexes:"), objc.Ptr(indexSet))
}

// Adds the indexes in an index range to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableindexset/1408251-addindexesinrange?language=objc
func (m_ MutableIndexSet) AddIndexesInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("addIndexesInRange:"), range_)
}
