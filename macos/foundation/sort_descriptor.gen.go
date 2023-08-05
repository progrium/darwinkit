// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var SortDescriptorClass = _SortDescriptorClass{objc.GetClass("NSSortDescriptor")}

type _SortDescriptorClass struct {
	objc.Class
}

type ISortDescriptor interface {
	objc.IObject
	CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult
	AllowEvaluation()
	Ascending() bool
	Key() string
	Selector() objc.Selector
	Comparator() func(obj1 objc.Object, obj2 objc.Object) ComparisonResult
	ReversedSortDescriptor() objc.Object
}

type SortDescriptor struct {
	objc.Object
}

func MakeSortDescriptor(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _SortDescriptorClass) SortDescriptorWithKeyAscending(key string, ascending bool) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.GetSelector("sortDescriptorWithKey:ascending:"), key, ascending)
	return rv
}

func SortDescriptor_SortDescriptorWithKeyAscending(key string, ascending bool) SortDescriptor {
	return SortDescriptorClass.SortDescriptorWithKeyAscending(key, ascending)
}

func (s_ SortDescriptor) InitWithKeyAscending(key string, ascending bool) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.GetSelector("initWithKey:ascending:"), key, ascending)
	return rv
}

func SortDescriptor_InitWithKeyAscending(key string, ascending bool) SortDescriptor {
	return SortDescriptorClass.Alloc().InitWithKeyAscending(key, ascending)
}

func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingSelector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.GetSelector("sortDescriptorWithKey:ascending:selector:"), key, ascending, selector)
	return rv
}

func SortDescriptor_SortDescriptorWithKeyAscendingSelector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	return SortDescriptorClass.SortDescriptorWithKeyAscendingSelector(key, ascending, selector)
}

func (s_ SortDescriptor) InitWithKeyAscendingSelector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.GetSelector("initWithKey:ascending:selector:"), key, ascending, selector)
	return rv
}

func SortDescriptor_InitWithKeyAscendingSelector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	return SortDescriptorClass.Alloc().InitWithKeyAscendingSelector(key, ascending, selector)
}

func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.GetSelector("sortDescriptorWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

func SortDescriptor_SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	return SortDescriptorClass.SortDescriptorWithKeyAscendingComparator(key, ascending, cmptr)
}

func (s_ SortDescriptor) InitWithKeyAscendingComparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.GetSelector("initWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

func SortDescriptor_InitWithKeyAscendingComparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	return SortDescriptorClass.Alloc().InitWithKeyAscendingComparator(key, ascending, cmptr)
}

func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.GetSelector("alloc"))
	return rv
}

func SortDescriptor_Alloc() SortDescriptor {
	return SortDescriptorClass.Alloc()
}

func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSortDescriptor() SortDescriptor {
	return SortDescriptorClass.New()
}

func SortDescriptor_New() SortDescriptor {
	return SortDescriptorClass.New()
}

func (s_ SortDescriptor) Init() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.GetSelector("init"))
	return rv
}

func SortDescriptor_Init() SortDescriptor {
	return SortDescriptorClass.Alloc().Init()
}

func (s_ SortDescriptor) CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](s_, objc.GetSelector("compareObject:toObject:"), objc.ExtractPtr(object1), objc.ExtractPtr(object2))
	return rv
}

func (s_ SortDescriptor) AllowEvaluation() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("allowEvaluation"))
}

func (s_ SortDescriptor) Ascending() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("ascending"))
	return rv
}

func (s_ SortDescriptor) Key() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("key"))
	return rv
}

func (s_ SortDescriptor) Selector() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.GetSelector("selector"))
	return rv
}

func (s_ SortDescriptor) Comparator() func(obj1 objc.Object, obj2 objc.Object) ComparisonResult {
	rv := objc.CallMethod[func(obj1 objc.Object, obj2 objc.Object) ComparisonResult](s_, objc.GetSelector("comparator"))
	return rv
}

func (s_ SortDescriptor) ReversedSortDescriptor() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("reversedSortDescriptor"))
	return rv
}
