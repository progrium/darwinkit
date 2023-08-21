// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Set] class.
var SetClass = _SetClass{objc.GetClass("NSSet")}

type _SetClass struct {
	objc.Class
}

// An interface definition for the [Set] class.
type ISet interface {
	objc.IObject
	ObjectEnumerator() Enumerator
	ObjectsPassingTest(predicate func(obj objc.Object, stop *bool) bool) Set
	EnumerateObjectsUsingBlock(block func(obj objc.Object, stop *bool))
	SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object
	FilteredSetUsingPredicate(predicate IPredicate) Set
	Member(object objc.IObject) objc.Object
	IsEqualToSet(otherSet ISet) bool
	SetByAddingObjectsFromArray(other []objc.IObject) Set
	EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, stop *bool))
	EnumerateIndexPathsWithOptionsUsingBlock(opts EnumerationOptions, block func(indexPath IndexPath, stop *bool))
	SetByAddingObject(anObject objc.IObject) Set
	ContainsObject(anObject objc.IObject) bool
	IntersectsSet(otherSet ISet) bool
	MakeObjectsPerformSelector(aSelector objc.Selector)
	ObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, stop *bool) bool) Set
	IsSubsetOfSet(otherSet ISet) bool
	SetByAddingObjectsFromSet(other ISet) Set
	RemoveObserverForKeyPathContext(observer objc.IObject, keyPath string, context unsafe.Pointer)
	DescriptionWithLocale(locale objc.IObject) string
	AnyObject() objc.Object
	SetValueForKey(value objc.IObject, key string)
	AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer)
	ValueForKey(key string) objc.Object
	Description() string
	Count() uint
	AllObjects() []objc.Object
}

// A static, unordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset?language=objc
type Set struct {
	objc.Object
}

func SetFrom(ptr unsafe.Pointer) Set {
	return Set{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SetClass) Set() Set {
	rv := objc.Call[Set](sc, objc.Sel("set"))
	return rv
}

// Creates and returns an empty set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574818-set?language=objc
func Set_Set() Set {
	return SetClass.Set()
}

func (s_ Set) InitWithObjects(firstObj objc.IObject, args ...any) Set {
	rv := objc.Call[Set](s_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574822-initwithobjects?language=objc
func NewSetWithObjects(firstObj objc.IObject, args ...any) Set {
	instance := SetClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (sc _SetClass) SetWithSet(set ISet) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns a set containing the objects from another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574817-setwithset?language=objc
func Set_SetWithSet(set ISet) Set {
	return SetClass.SetWithSet(set)
}

func (sc _SetClass) SetWithObject(object objc.IObject) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithObject:"), objc.Ptr(object))
	return rv
}

// Creates and returns a set that contains a single given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1415878-setwithobject?language=objc
func Set_SetWithObject(object objc.IObject) Set {
	return SetClass.SetWithObject(object)
}

func (s_ Set) InitWithArray(array []objc.IObject) Set {
	rv := objc.Call[Set](s_, objc.Sel("initWithArray:"), array)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1408152-initwitharray?language=objc
func NewSetWithArray(array []objc.IObject) Set {
	instance := SetClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (sc _SetClass) SetWithArray(array []objc.IObject) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithArray:"), array)
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574823-setwitharray?language=objc
func Set_SetWithArray(array []objc.IObject) Set {
	return SetClass.SetWithArray(array)
}

func (sc _SetClass) SetWithCollectionViewIndexPath(indexPath IIndexPath) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithCollectionViewIndexPath:"), objc.Ptr(indexPath))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528161-setwithcollectionviewindexpath?language=objc
func Set_SetWithCollectionViewIndexPath(indexPath IIndexPath) Set {
	return SetClass.SetWithCollectionViewIndexPath(indexPath)
}

func (sc _SetClass) SetWithObjectsCount(objects objc.IObject, cnt uint) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithObjects:count:"), objc.Ptr(objects), cnt)
	return rv
}

// Creates and returns a set containing a specified number of objects from a given C array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574824-setwithobjects?language=objc
func Set_SetWithObjectsCount(objects objc.IObject, cnt uint) Set {
	return SetClass.SetWithObjectsCount(objects, cnt)
}

func (sc _SetClass) SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) Set {
	rv := objc.Call[Set](sc, objc.Sel("setWithCollectionViewIndexPaths:"), indexPaths)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528255-setwithcollectionviewindexpaths?language=objc
func Set_SetWithCollectionViewIndexPaths(indexPaths []IIndexPath) Set {
	return SetClass.SetWithCollectionViewIndexPaths(indexPaths)
}

func (s_ Set) InitWithSet(set ISet) Set {
	rv := objc.Call[Set](s_, objc.Sel("initWithSet:"), objc.Ptr(set))
	return rv
}

// Initializes a newly allocated set and adds to it objects from another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1410612-initwithset?language=objc
func NewSetWithSet(set ISet) Set {
	instance := SetClass.Alloc().InitWithSet(set)
	instance.Autorelease()
	return instance
}

func (s_ Set) Init() Set {
	rv := objc.Call[Set](s_, objc.Sel("init"))
	return rv
}

func (sc _SetClass) Alloc() Set {
	rv := objc.Call[Set](sc, objc.Sel("alloc"))
	return rv
}

func Set_Alloc() Set {
	return SetClass.Alloc()
}

func (sc _SetClass) New() Set {
	rv := objc.Call[Set](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSet() Set {
	return SetClass.New()
}

// Returns an enumerator object that lets you access each object in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1412373-objectenumerator?language=objc
func (s_ Set) ObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](s_, objc.Sel("objectEnumerator"))
	return rv
}

// Returns a set of objects that pass a test in a given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1414392-objectspassingtest?language=objc
func (s_ Set) ObjectsPassingTest(predicate func(obj objc.Object, stop *bool) bool) Set {
	rv := objc.Call[Set](s_, objc.Sel("objectsPassingTest:"), predicate)
	return rv
}

// Executes a given block using each object in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1418129-enumerateobjectsusingblock?language=objc
func (s_ Set) EnumerateObjectsUsingBlock(block func(obj objc.Object, stop *bool)) {
	objc.Call[objc.Void](s_, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Returns an array of the set’s content sorted as specified by a given array of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416427-sortedarrayusingdescriptors?language=objc
func (s_ Set) SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object {
	rv := objc.Call[[]objc.Object](s_, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}

// Evaluates a given predicate against each object in the receiving set and returns a new set containing the objects for which the predicate returns true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416324-filteredsetusingpredicate?language=objc
func (s_ Set) FilteredSetUsingPredicate(predicate IPredicate) Set {
	rv := objc.Call[Set](s_, objc.Sel("filteredSetUsingPredicate:"), objc.Ptr(predicate))
	return rv
}

// Determines whether a given object is present in the set, and returns that object if it is. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1412896-member?language=objc
func (s_ Set) Member(object objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("member:"), objc.Ptr(object))
	return rv
}

// Compares the receiving set to another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1414829-isequaltoset?language=objc
func (s_ Set) IsEqualToSet(otherSet ISet) bool {
	rv := objc.Call[bool](s_, objc.Sel("isEqualToSet:"), objc.Ptr(otherSet))
	return rv
}

// Returns a new set formed by adding the objects in a given array to the receiving set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1418438-setbyaddingobjectsfromarray?language=objc
func (s_ Set) SetByAddingObjectsFromArray(other []objc.IObject) Set {
	rv := objc.Call[Set](s_, objc.Sel("setByAddingObjectsFromArray:"), other)
	return rv
}

// Executes a given block using each object in the set, using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1412024-enumerateobjectswithoptions?language=objc
func (s_ Set) EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, stop *bool)) {
	objc.Call[objc.Void](s_, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1528216-enumerateindexpathswithoptions?language=objc
func (s_ Set) EnumerateIndexPathsWithOptionsUsingBlock(opts EnumerationOptions, block func(indexPath IndexPath, stop *bool)) {
	objc.Call[objc.Void](s_, objc.Sel("enumerateIndexPathsWithOptions:usingBlock:"), opts, block)
}

// Returns a new set formed by adding a given object to the receiving set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416316-setbyaddingobject?language=objc
func (s_ Set) SetByAddingObject(anObject objc.IObject) Set {
	rv := objc.Call[Set](s_, objc.Sel("setByAddingObject:"), objc.Ptr(anObject))
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1414555-containsobject?language=objc
func (s_ Set) ContainsObject(anObject objc.IObject) bool {
	rv := objc.Call[bool](s_, objc.Sel("containsObject:"), objc.Ptr(anObject))
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the receiving set is also present in another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1417472-intersectsset?language=objc
func (s_ Set) IntersectsSet(otherSet ISet) bool {
	rv := objc.Call[bool](s_, objc.Sel("intersectsSet:"), objc.Ptr(otherSet))
	return rv
}

// Sends a message specified by a given selector to each object in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1574819-makeobjectsperformselector?language=objc
func (s_ Set) MakeObjectsPerformSelector(aSelector objc.Selector) {
	objc.Call[objc.Void](s_, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}

// Returns a set of objects that pass a test in a given block, using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416826-objectswithoptions?language=objc
func (s_ Set) ObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, stop *bool) bool) Set {
	rv := objc.Call[Set](s_, objc.Sel("objectsWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns a Boolean value that indicates whether every object in the receiving set is also present in another given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1418319-issubsetofset?language=objc
func (s_ Set) IsSubsetOfSet(otherSet ISet) bool {
	rv := objc.Call[bool](s_, objc.Sel("isSubsetOfSet:"), objc.Ptr(otherSet))
	return rv
}

// Returns a new set formed by adding the objects in a given set to the receiving set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1408217-setbyaddingobjectsfromset?language=objc
func (s_ Set) SetByAddingObjectsFromSet(other ISet) Set {
	rv := objc.Call[Set](s_, objc.Sel("setByAddingObjectsFromSet:"), objc.Ptr(other))
	return rv
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1415413-removeobserver?language=objc
func (s_ Set) RemoveObserverForKeyPathContext(observer objc.IObject, keyPath string, context unsafe.Pointer) {
	objc.Call[objc.Void](s_, objc.Sel("removeObserver:forKeyPath:context:"), objc.Ptr(observer), keyPath, context)
}

// Returns a string that represents the contents of the set, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1417205-descriptionwithlocale?language=objc
func (s_ Set) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](s_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns one of the objects in the set, or nil if the set contains no objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416575-anyobject?language=objc
func (s_ Set) AnyObject() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("anyObject"))
	return rv
}

// Invokes setValue:forKey: on each of the set’s members. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1408322-setvalue?language=objc
func (s_ Set) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](s_, objc.Sel("setValue:forKey:"), value, key)
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1414043-addobserver?language=objc
func (s_ Set) AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer) {
	objc.Call[objc.Void](s_, objc.Sel("addObserver:forKeyPath:options:context:"), objc.Ptr(observer), keyPath, options, context)
}

// Return a set containing the results of invoking valueForKey: on each of the receiving set's members. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1418386-valueforkey?language=objc
func (s_ Set) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("valueForKey:"), key)
	return rv
}

// A string that represents the contents of the set, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1418176-description?language=objc
func (s_ Set) Description() string {
	rv := objc.Call[string](s_, objc.Sel("description"))
	return rv
}

// The number of members in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1416229-count?language=objc
func (s_ Set) Count() uint {
	rv := objc.Call[uint](s_, objc.Sel("count"))
	return rv
}

// An array containing the set’s members, or an empty array if the set has no members. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/1417653-allobjects?language=objc
func (s_ Set) AllObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](s_, objc.Sel("allObjects"))
	return rv
}
