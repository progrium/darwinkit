// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SortDescriptor] class.
var SortDescriptorClass = _SortDescriptorClass{objc.GetClass("NSSortDescriptor")}

type _SortDescriptorClass struct {
	objc.Class
}

// An interface definition for the [SortDescriptor] class.
type ISortDescriptor interface {
	objc.IObject
	AllowEvaluation()
	CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult
	Comparator() Comparator
	Key() string
	Selector() objc.Selector
	ReversedSortDescriptor() objc.Object
	Ascending() bool
}

// An immutable description of how to order a collection of objects according to a property common to all the objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor?language=objc
type SortDescriptor struct {
	objc.Object
}

func SortDescriptorFrom(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr Comparator) SortDescriptor {
	rv := objc.Call[SortDescriptor](sc, objc.Sel("sortDescriptorWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

// Creates and returns a sort descriptor initialized with the specified key path and ordering, and a comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1503734-sortdescriptorwithkey?language=objc
func SortDescriptor_SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr Comparator) SortDescriptor {
	return SortDescriptorClass.SortDescriptorWithKeyAscendingComparator(key, ascending, cmptr)
}

func (s_ SortDescriptor) InitWithKeyAscending(key string, ascending bool) SortDescriptor {
	rv := objc.Call[SortDescriptor](s_, objc.Sel("initWithKey:ascending:"), key, ascending)
	return rv
}

// Creates a sort descriptor with a specified string key path and sort order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1413572-initwithkey?language=objc
func NewSortDescriptorWithKeyAscending(key string, ascending bool) SortDescriptor {
	instance := SortDescriptorClass.Alloc().InitWithKeyAscending(key, ascending)
	instance.Autorelease()
	return instance
}

func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.Call[SortDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func SortDescriptor_Alloc() SortDescriptor {
	return SortDescriptorClass.Alloc()
}

func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := objc.Call[SortDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSortDescriptor() SortDescriptor {
	return SortDescriptorClass.New()
}

func (s_ SortDescriptor) Init() SortDescriptor {
	rv := objc.Call[SortDescriptor](s_, objc.Sel("init"))
	return rv
}

// Forces a securely decoded sort descriptor to allow evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1412371-allowevaluation?language=objc
func (s_ SortDescriptor) AllowEvaluation() {
	objc.Call[objc.Void](s_, objc.Sel("allowEvaluation"))
}

// Returns a comparison result value that indicates the sort order of two objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1411557-compareobject?language=objc
func (s_ SortDescriptor) CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("compareObject:toObject:"), object1, object2)
	return rv
}

// The comparator for the sort descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1411426-comparator?language=objc
func (s_ SortDescriptor) Comparator() Comparator {
	rv := objc.Call[Comparator](s_, objc.Sel("comparator"))
	return rv
}

// The key that specifies the property to compare during sorting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1415022-key?language=objc
func (s_ SortDescriptor) Key() string {
	rv := objc.Call[string](s_, objc.Sel("key"))
	return rv
}

// The selector for comparing objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1418337-selector?language=objc
func (s_ SortDescriptor) Selector() objc.Selector {
	rv := objc.Call[objc.Selector](s_, objc.Sel("selector"))
	return rv
}

// Returns a sort descriptor that reverses the sort order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1407712-reversedsortdescriptor?language=objc
func (s_ SortDescriptor) ReversedSortDescriptor() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("reversedSortDescriptor"))
	return rv
}

// A Boolean value that indicates whether the receiver specifies sorting in ascending order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1408931-ascending?language=objc
func (s_ SortDescriptor) Ascending() bool {
	rv := objc.Call[bool](s_, objc.Sel("ascending"))
	return rv
}
