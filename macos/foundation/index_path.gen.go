// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var IndexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}

type _IndexPathClass struct {
	objc.Class
}

type IIndexPath interface {
	objc.IObject
	IndexPathByAddingIndex(index uint) IndexPath
	IndexPathByRemovingLastIndex() IndexPath
	Compare(otherObject IIndexPath) ComparisonResult
	IndexAtPosition(position uint) uint
	GetIndexesRange(indexes *uint, positionRange Range)
	Length() uint
}

type IndexPath struct {
	objc.Object
}

func MakeIndexPath(ptr unsafe.Pointer) IndexPath {
	return IndexPath{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _IndexPathClass) IndexPathWithIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](ic, objc.GetSelector("indexPathWithIndex:"), index)
	return rv
}

func IndexPath_IndexPathWithIndex(index uint) IndexPath {
	return IndexPathClass.IndexPathWithIndex(index)
}

func (ic _IndexPathClass) IndexPathWithIndexesLength(indexes *uint, length uint) IndexPath {
	rv := objc.CallMethod[IndexPath](ic, objc.GetSelector("indexPathWithIndexes:length:"), indexes, length)
	return rv
}

func IndexPath_IndexPathWithIndexesLength(indexes *uint, length uint) IndexPath {
	return IndexPathClass.IndexPathWithIndexesLength(indexes, length)
}

func (i_ IndexPath) InitWithIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, objc.GetSelector("initWithIndex:"), index)
	return rv
}

func IndexPath_InitWithIndex(index uint) IndexPath {
	return IndexPathClass.Alloc().InitWithIndex(index)
}

func (i_ IndexPath) InitWithIndexesLength(indexes *uint, length uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, objc.GetSelector("initWithIndexes:length:"), indexes, length)
	return rv
}

func IndexPath_InitWithIndexesLength(indexes *uint, length uint) IndexPath {
	return IndexPathClass.Alloc().InitWithIndexesLength(indexes, length)
}

func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.CallMethod[IndexPath](ic, objc.GetSelector("alloc"))
	return rv
}

func IndexPath_Alloc() IndexPath {
	return IndexPathClass.Alloc()
}

func (ic _IndexPathClass) New() IndexPath {
	rv := objc.CallMethod[IndexPath](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewIndexPath() IndexPath {
	return IndexPathClass.New()
}

func IndexPath_New() IndexPath {
	return IndexPathClass.New()
}

func (i_ IndexPath) Init() IndexPath {
	rv := objc.CallMethod[IndexPath](i_, objc.GetSelector("init"))
	return rv
}

func IndexPath_Init() IndexPath {
	return IndexPathClass.Alloc().Init()
}

func (i_ IndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, objc.GetSelector("indexPathByAddingIndex:"), index)
	return rv
}

func (i_ IndexPath) IndexPathByRemovingLastIndex() IndexPath {
	rv := objc.CallMethod[IndexPath](i_, objc.GetSelector("indexPathByRemovingLastIndex"))
	return rv
}

func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](i_, objc.GetSelector("compare:"), objc.ExtractPtr(otherObject))
	return rv
}

func (i_ IndexPath) IndexAtPosition(position uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexAtPosition:"), position)
	return rv
}

func (i_ IndexPath) GetIndexesRange(indexes *uint, positionRange Range) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("getIndexes:range:"), indexes, positionRange)
}

func (i_ IndexPath) Length() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("length"))
	return rv
}
