// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var IndexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}

type _IndexSetClass struct {
	objc.Class
}

type IIndexSet interface {
	objc.IObject
	ContainsIndex(value uint) bool
	ContainsIndexes(indexSet IIndexSet) bool
	ContainsIndexesInRange(range_ Range) bool
	IntersectsIndexesInRange(range_ Range) bool
	CountOfIndexesInRange(range_ Range) uint
	IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint
	IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet
	IndexWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	IndexInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	EnumerateRangesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool))
	EnumerateRangesUsingBlock(block func(range_ Range, stop *bool))
	EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool))
	IsEqualToIndexSet(indexSet IIndexSet) bool
	IndexLessThanIndex(value uint) uint
	IndexLessThanOrEqualToIndex(value uint) uint
	IndexGreaterThanOrEqualToIndex(value uint) uint
	IndexGreaterThanIndex(value uint) uint
	GetIndexesMaxCountInIndexRange(indexBuffer *uint, bufferSize uint, range_ *Range) uint
	EnumerateIndexesUsingBlock(block func(idx uint, stop *bool))
	EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool))
	EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool))
	Count() uint
	FirstIndex() uint
	LastIndex() uint
}

type IndexSet struct {
	objc.Object
}

func MakeIndexSet(ptr unsafe.Pointer) IndexSet {
	return IndexSet{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _IndexSetClass) IndexSet() IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSet"))
	return rv
}

func IndexSet_IndexSet() IndexSet {
	return IndexSetClass.IndexSet()
}

func (ic _IndexSetClass) IndexSetWithIndex(value uint) IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSetWithIndex:"), value)
	return rv
}

func IndexSet_IndexSetWithIndex(value uint) IndexSet {
	return IndexSetClass.IndexSetWithIndex(value)
}

func (ic _IndexSetClass) IndexSetWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSetWithIndexesInRange:"), range_)
	return rv
}

func IndexSet_IndexSetWithIndexesInRange(range_ Range) IndexSet {
	return IndexSetClass.IndexSetWithIndexesInRange(range_)
}

func (i_ IndexSet) InitWithIndex(value uint) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndex:"), value)
	return rv
}

func IndexSet_InitWithIndex(value uint) IndexSet {
	return IndexSetClass.Alloc().InitWithIndex(value)
}

func (i_ IndexSet) InitWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndexesInRange:"), range_)
	return rv
}

func IndexSet_InitWithIndexesInRange(range_ Range) IndexSet {
	return IndexSetClass.Alloc().InitWithIndexesInRange(range_)
}

func (i_ IndexSet) InitWithIndexSet(indexSet IIndexSet) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndexSet:"), objc.ExtractPtr(indexSet))
	return rv
}

func IndexSet_InitWithIndexSet(indexSet IIndexSet) IndexSet {
	return IndexSetClass.Alloc().InitWithIndexSet(indexSet)
}

func (ic _IndexSetClass) Alloc() IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("alloc"))
	return rv
}

func IndexSet_Alloc() IndexSet {
	return IndexSetClass.Alloc()
}

func (ic _IndexSetClass) New() IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewIndexSet() IndexSet {
	return IndexSetClass.New()
}

func IndexSet_New() IndexSet {
	return IndexSetClass.New()
}

func (i_ IndexSet) Init() IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("init"))
	return rv
}

func IndexSet_Init() IndexSet {
	return IndexSetClass.Alloc().Init()
}

func (i_ IndexSet) ContainsIndex(value uint) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndex:"), value)
	return rv
}

func (i_ IndexSet) ContainsIndexes(indexSet IIndexSet) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndexes:"), objc.ExtractPtr(indexSet))
	return rv
}

func (i_ IndexSet) ContainsIndexesInRange(range_ Range) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) IntersectsIndexesInRange(range_ Range) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("intersectsIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) CountOfIndexesInRange(range_ Range) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("countOfIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexPassingTest:"), predicate)
	return rv
}

func (i_ IndexSet) IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesPassingTest:"), predicate)
	return rv
}

func (i_ IndexSet) IndexWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexWithOptions:passingTest:"), opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesWithOptions:passingTest:"), opts, predicate)
	return rv
}

func (i_ IndexSet) IndexInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

func (i_ IndexSet) EnumerateRangesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesInRange:options:usingBlock:"), range_, opts, block)
}

func (i_ IndexSet) EnumerateRangesUsingBlock(block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesUsingBlock:"), block)
}

func (i_ IndexSet) EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesWithOptions:usingBlock:"), opts, block)
}

func (i_ IndexSet) IsEqualToIndexSet(indexSet IIndexSet) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isEqualToIndexSet:"), objc.ExtractPtr(indexSet))
	return rv
}

func (i_ IndexSet) IndexLessThanIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexLessThanIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexLessThanOrEqualToIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexGreaterThanOrEqualToIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexGreaterThanIndex:"), value)
	return rv
}

func (i_ IndexSet) GetIndexesMaxCountInIndexRange(indexBuffer *uint, bufferSize uint, range_ *Range) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("getIndexes:maxCount:inIndexRange:"), indexBuffer, bufferSize, range_)
	return rv
}

func (i_ IndexSet) EnumerateIndexesUsingBlock(block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesUsingBlock:"), block)
}

func (i_ IndexSet) EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}

func (i_ IndexSet) EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}

func (i_ IndexSet) Count() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("count"))
	return rv
}

func (i_ IndexSet) FirstIndex() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("firstIndex"))
	return rv
}

func (i_ IndexSet) LastIndex() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("lastIndex"))
	return rv
}
