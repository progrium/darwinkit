// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OrderedSet] class.
var OrderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}

type _OrderedSetClass struct {
	objc.Class
}

// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objc.IObject
	ObjectEnumerator() Enumerator
	IsEqualToOrderedSet(other IOrderedSet) bool
	SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object
	ObjectsAtIndexes(indexes IIndexSet) []objc.Object
	EnumerateObjectsUsingBlock(block func(obj objc.Object, idx uint, stop *bool))
	SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object
	IsSubsetOfOrderedSet(other IOrderedSet) bool
	OrderedSetByApplyingDifference(difference IOrderedCollectionDifference) OrderedSet
	IntersectsOrderedSet(other IOrderedSet) bool
	DifferenceFromOrderedSet(other IOrderedSet) OrderedCollectionDifference
	IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	IndexOfObjectPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	ObjectAtIndex(idx uint) objc.Object
	ObjectAtIndexedSubscript(idx uint) objc.Object
	EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool))
	GetObjectsRange(objects objc.IObject, range_ Range)
	ContainsObject(object objc.IObject) bool
	IntersectsSet(set ISet) bool
	FilteredOrderedSetUsingPredicate(p IPredicate) OrderedSet
	ReverseObjectEnumerator() Enumerator
	IsSubsetOfSet(set ISet) bool
	IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	RemoveObserverForKeyPath(observer objc.IObject, keyPath string)
	DescriptionWithLocale(locale objc.IObject) string
	SetValueForKey(value objc.IObject, key string)
	IndexesOfObjectsPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	IndexOfObject(object objc.IObject) uint
	AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer)
	IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool))
	ValueForKey(key string) objc.Object
	IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	SortedArrayUsingComparator(cmptr Comparator) []objc.Object
	Set() Set
	LastObject() objc.Object
	Array() []objc.Object
	Description() string
	Count() uint
	ReversedOrderedSet() OrderedSet
	FirstObject() objc.Object
}

// A static, ordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset?language=objc
type OrderedSet struct {
	objc.Object
}

func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OrderedSetClass) OrderedSetWithOrderedSet(set IOrderedSet) OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSetWithOrderedSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns an ordered set containing the objects from another ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543280-orderedsetwithorderedset?language=objc
func OrderedSet_OrderedSetWithOrderedSet(set IOrderedSet) OrderedSet {
	return OrderedSetClass.OrderedSetWithOrderedSet(set)
}

func (o_ OrderedSet) InitWithObjects(firstObj objc.IObject, args ...any) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543287-initwithobjects?language=objc
func NewOrderedSetWithObjects(firstObj objc.IObject, args ...any) OrderedSet {
	instance := OrderedSetClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (oc _OrderedSetClass) OrderedSetWithObjects(firstObj objc.IObject, args ...any) OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSetWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Creates and returns a ordered set containing the objects in a given argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543312-orderedsetwithobjects?language=objc
func OrderedSet_OrderedSetWithObjects(firstObj objc.IObject, args ...any) OrderedSet {
	return OrderedSetClass.OrderedSetWithObjects(firstObj, args...)
}

func (o_ OrderedSet) InitWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("initWithOrderedSet:range:copyItems:"), objc.Ptr(set), range_, flag)
	return rv
}

// Initializes a new ordered set with the contents of an ordered set, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1417751-initwithorderedset?language=objc
func NewOrderedSetWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) OrderedSet {
	instance := OrderedSetClass.Alloc().InitWithOrderedSetRangeCopyItems(set, range_, flag)
	instance.Autorelease()
	return instance
}

func (o_ OrderedSet) InitWithObject(object objc.IObject) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("initWithObject:"), objc.Ptr(object))
	return rv
}

// Initializes a new ordered set with the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413883-initwithobject?language=objc
func NewOrderedSetWithObject(object objc.IObject) OrderedSet {
	instance := OrderedSetClass.Alloc().InitWithObject(object)
	instance.Autorelease()
	return instance
}

func (o_ OrderedSet) InitWithArrayCopyItems(set []objc.IObject, flag bool) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("initWithArray:copyItems:"), set, flag)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1418006-initwitharray?language=objc
func NewOrderedSetWithArrayCopyItems(set []objc.IObject, flag bool) OrderedSet {
	instance := OrderedSetClass.Alloc().InitWithArrayCopyItems(set, flag)
	instance.Autorelease()
	return instance
}

func (oc _OrderedSetClass) OrderedSetWithArray(array []objc.IObject) OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSetWithArray:"), array)
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543310-orderedsetwitharray?language=objc
func OrderedSet_OrderedSetWithArray(array []objc.IObject) OrderedSet {
	return OrderedSetClass.OrderedSetWithArray(array)
}

func (oc _OrderedSetClass) OrderedSetWithObject(object objc.IObject) OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSetWithObject:"), objc.Ptr(object))
	return rv
}

// Creates and returns a ordered set that contains a single given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543339-orderedsetwithobject?language=objc
func OrderedSet_OrderedSetWithObject(object objc.IObject) OrderedSet {
	return OrderedSetClass.OrderedSetWithObject(object)
}

func (oc _OrderedSetClass) OrderedSet() OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSet"))
	return rv
}

// Creates and returns an empty ordered set [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543313-orderedset?language=objc
func OrderedSet_OrderedSet() OrderedSet {
	return OrderedSetClass.OrderedSet()
}

func (o_ OrderedSet) InitWithSet(set ISet) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("initWithSet:"), objc.Ptr(set))
	return rv
}

// Initializes a new ordered set with the contents of a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1416344-initwithset?language=objc
func NewOrderedSetWithSet(set ISet) OrderedSet {
	instance := OrderedSetClass.Alloc().InitWithSet(set)
	instance.Autorelease()
	return instance
}

func (o_ OrderedSet) Init() OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("init"))
	return rv
}

func (oc _OrderedSetClass) OrderedSetWithSet(set ISet) OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("orderedSetWithSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns an ordered set with the contents of a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543298-orderedsetwithset?language=objc
func OrderedSet_OrderedSetWithSet(set ISet) OrderedSet {
	return OrderedSetClass.OrderedSetWithSet(set)
}

func (oc _OrderedSetClass) Alloc() OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("alloc"))
	return rv
}

func OrderedSet_Alloc() OrderedSet {
	return OrderedSetClass.Alloc()
}

func (oc _OrderedSetClass) New() OrderedSet {
	rv := objc.Call[OrderedSet](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOrderedSet() OrderedSet {
	return OrderedSetClass.New()
}

// Returns an enumerator object that lets you access each object in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409430-objectenumerator?language=objc
func (o_ OrderedSet) ObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](o_, objc.Sel("objectEnumerator"))
	return rv
}

// Compares the receiving ordered set to another ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408049-isequaltoorderedset?language=objc
func (o_ OrderedSet) IsEqualToOrderedSet(other IOrderedSet) bool {
	rv := objc.Call[bool](o_, objc.Sel("isEqualToOrderedSet:"), objc.Ptr(other))
	return rv
}

// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given NSComparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1414806-sortedarraywithoptions?language=objc
func (o_ OrderedSet) SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return rv
}

// Returns the objects in the ordered set at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1414943-objectsatindexes?language=objc
func (o_ OrderedSet) ObjectsAtIndexes(indexes IIndexSet) []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("objectsAtIndexes:"), objc.Ptr(indexes))
	return rv
}

// Executes a given block using each object in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413531-enumerateobjectsusingblock?language=objc
func (o_ OrderedSet) EnumerateObjectsUsingBlock(block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](o_, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Returns an array of the ordered set’s elements sorted as specified by a given array of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409953-sortedarrayusingdescriptors?language=objc
func (o_ OrderedSet) SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}

// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411496-issubsetoforderedset?language=objc
func (o_ OrderedSet) IsSubsetOfOrderedSet(other IOrderedSet) bool {
	rv := objc.Call[bool](o_, objc.Sel("isSubsetOfOrderedSet:"), objc.Ptr(other))
	return rv
}

// Creates a new ordered set by applying a difference object to an existing ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/3152198-orderedsetbyapplyingdifference?language=objc
func (o_ OrderedSet) OrderedSetByApplyingDifference(difference IOrderedCollectionDifference) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("orderedSetByApplyingDifference:"), objc.Ptr(difference))
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1414364-intersectsorderedset?language=objc
func (o_ OrderedSet) IntersectsOrderedSet(other IOrderedSet) bool {
	rv := objc.Call[bool](o_, objc.Sel("intersectsOrderedSet:"), objc.Ptr(other))
	return rv
}

// Compares two ordered sets to create a difference object that represents the changes between them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/3152195-differencefromorderedset?language=objc
func (o_ OrderedSet) DifferenceFromOrderedSet(other IOrderedSet) OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("differenceFromOrderedSet:"), objc.Ptr(other))
	return rv
}

// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1415944-indexesofobjectswithoptions?language=objc
func (o_ OrderedSet) IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](o_, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the index of the object in the ordered set that passes a test in a given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413003-indexofobjectpassingtest?language=objc
func (o_ OrderedSet) IndexOfObjectPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](o_, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}

// Returns the object at the specified index of the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1414734-objectatindex?language=objc
func (o_ OrderedSet) ObjectAtIndex(idx uint) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("objectAtIndex:"), idx)
	return rv
}

// Returns the object at the specified index of the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1414253-objectatindexedsubscript?language=objc
func (o_ OrderedSet) ObjectAtIndexedSubscript(idx uint) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Executes a given block using each object in the set, using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409354-enumerateobjectswithoptions?language=objc
func (o_ OrderedSet) EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](o_, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Copies the objects contained in the ordered set that fall within the specified range to objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411401-getobjects?language=objc
func (o_ OrderedSet) GetObjectsRange(objects objc.IObject, range_ Range) {
	objc.Call[objc.Void](o_, objc.Sel("getObjects:range:"), objc.Ptr(objects), range_)
}

// Returns a Boolean value that indicates whether a given object is present in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408681-containsobject?language=objc
func (o_ OrderedSet) ContainsObject(object objc.IObject) bool {
	rv := objc.Call[bool](o_, objc.Sel("containsObject:"), objc.Ptr(object))
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408625-intersectsset?language=objc
func (o_ OrderedSet) IntersectsSet(set ISet) bool {
	rv := objc.Call[bool](o_, objc.Sel("intersectsSet:"), objc.Ptr(set))
	return rv
}

// Evaluates a given predicate against each object in the receiving ordered set and returns a new ordered set containing the objects for which the predicate returns true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1415807-filteredorderedsetusingpredicate?language=objc
func (o_ OrderedSet) FilteredOrderedSetUsingPredicate(p IPredicate) OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("filteredOrderedSetUsingPredicate:"), objc.Ptr(p))
	return rv
}

// Returns an enumerator object that lets you access each object in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1407607-reverseobjectenumerator?language=objc
func (o_ OrderedSet) ReverseObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](o_, objc.Sel("reverseObjectEnumerator"))
	return rv
}

// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1418464-issubsetofset?language=objc
func (o_ OrderedSet) IsSubsetOfSet(set ISet) bool {
	rv := objc.Call[bool](o_, objc.Sel("isSubsetOfSet:"), objc.Ptr(set))
	return rv
}

// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413586-indexesofobjectsatindexes?language=objc
func (o_ OrderedSet) IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](o_, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), objc.Ptr(s), opts, predicate)
	return rv
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1412955-removeobserver?language=objc
func (o_ OrderedSet) RemoveObserverForKeyPath(observer objc.IObject, keyPath string) {
	objc.Call[objc.Void](o_, objc.Sel("removeObserver:forKeyPath:"), objc.Ptr(observer), keyPath)
}

// Returns a string that represents the contents of the ordered set, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1417325-descriptionwithlocale?language=objc
func (o_ OrderedSet) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](o_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Invokes setValue:forKey: on each of the receiver's members using the specified value and key [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413118-setvalue?language=objc
func (o_ OrderedSet) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](o_, objc.Sel("setValue:forKey:"), value, key)
}

// Returns the index of the object in the ordered set that passes a test in a given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411331-indexesofobjectspassingtest?language=objc
func (o_ OrderedSet) IndexesOfObjectsPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](o_, objc.Sel("indexesOfObjectsPassingTest:"), predicate)
	return rv
}

// Returns the index of the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411856-indexofobject?language=objc
func (o_ OrderedSet) IndexOfObject(object objc.IObject) uint {
	rv := objc.Call[uint](o_, objc.Sel("indexOfObject:"), objc.Ptr(object))
	return rv
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408740-addobserver?language=objc
func (o_ OrderedSet) AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer) {
	objc.Call[objc.Void](o_, objc.Sel("addObserver:forKeyPath:options:context:"), objc.Ptr(observer), keyPath, options, context)
}

// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408700-indexofobjectwithoptions?language=objc
func (o_ OrderedSet) IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](o_, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Executes a given block using the objects in the ordered set at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1412332-enumerateobjectsatindexes?language=objc
func (o_ OrderedSet) EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](o_, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), objc.Ptr(s), opts, block)
}

// Returns an ordered set containing the results of invoking valueForKey: using key on each of the ordered set’s objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409378-valueforkey?language=objc
func (o_ OrderedSet) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("valueForKey:"), key)
	return rv
}

// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1417531-indexofobjectatindexes?language=objc
func (o_ OrderedSet) IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](o_, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), objc.Ptr(s), opts, predicate)
	return rv
}

// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given NSComparator block [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413383-sortedarrayusingcomparator?language=objc
func (o_ OrderedSet) SortedArrayUsingComparator(cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return rv
}

// A representation of the set containing the contents of the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413944-set?language=objc
func (o_ OrderedSet) Set() Set {
	rv := objc.Call[Set](o_, objc.Sel("set"))
	return rv
}

// The last object in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409143-lastobject?language=objc
func (o_ OrderedSet) LastObject() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("lastObject"))
	return rv
}

// A representation of the ordered set as an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411531-array?language=objc
func (o_ OrderedSet) Array() []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("array"))
	return rv
}

// A string that represents the contents of the ordered set, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1415872-description?language=objc
func (o_ OrderedSet) Description() string {
	rv := objc.Call[string](o_, objc.Sel("description"))
	return rv
}

// The number of members in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1410106-count?language=objc
func (o_ OrderedSet) Count() uint {
	rv := objc.Call[uint](o_, objc.Sel("count"))
	return rv
}

// An ordered set in the reverse order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411022-reversedorderedset?language=objc
func (o_ OrderedSet) ReversedOrderedSet() OrderedSet {
	rv := objc.Call[OrderedSet](o_, objc.Sel("reversedOrderedSet"))
	return rv
}

// The first object in the ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409765-firstobject?language=objc
func (o_ OrderedSet) FirstObject() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("firstObject"))
	return rv
}
