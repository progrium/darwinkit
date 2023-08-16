// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchIndexDescription] class.
var FetchIndexDescriptionClass = _FetchIndexDescriptionClass{objc.GetClass("NSFetchIndexDescription")}

type _FetchIndexDescriptionClass struct {
	objc.Class
}

// An interface definition for the [FetchIndexDescription] class.
type IFetchIndexDescription interface {
	objc.IObject
	Entity() EntityDescription
	Name() string
	SetName(value string)
	Elements() []FetchIndexElementDescription
	SetElements(value []IFetchIndexElementDescription)
	PartialIndexPredicate() foundation.Predicate
	SetPartialIndexPredicate(value foundation.IPredicate)
}

// The description of the index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription?language=objc
type FetchIndexDescription struct {
	objc.Object
}

func FetchIndexDescriptionFrom(ptr unsafe.Pointer) FetchIndexDescription {
	return FetchIndexDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FetchIndexDescription) InitWithNameElements(name string, elements []IFetchIndexElementDescription) FetchIndexDescription {
	rv := objc.Call[FetchIndexDescription](f_, objc.Sel("initWithName:elements:"), name, elements)
	return rv
}

// Creates a fetch index description using the specified name and element descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887053-initwithname?language=objc
func FetchIndexDescription_InitWithNameElements(name string, elements []IFetchIndexElementDescription) FetchIndexDescription {
	return FetchIndexDescriptionClass.Alloc().InitWithNameElements(name, elements)
}

func (fc _FetchIndexDescriptionClass) Alloc() FetchIndexDescription {
	rv := objc.Call[FetchIndexDescription](fc, objc.Sel("alloc"))
	return rv
}

func FetchIndexDescription_Alloc() FetchIndexDescription {
	return FetchIndexDescriptionClass.Alloc()
}

func (fc _FetchIndexDescriptionClass) New() FetchIndexDescription {
	rv := objc.Call[FetchIndexDescription](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchIndexDescription() FetchIndexDescription {
	return FetchIndexDescriptionClass.New()
}

func (f_ FetchIndexDescription) Init() FetchIndexDescription {
	rv := objc.Call[FetchIndexDescription](f_, objc.Sel("init"))
	return rv
}

// The entity description for the fetch index description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887055-entity?language=objc
func (f_ FetchIndexDescription) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](f_, objc.Sel("entity"))
	return rv
}

// The name of the fetch index description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887056-name?language=objc
func (f_ FetchIndexDescription) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// The name of the fetch index description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887056-name?language=objc
func (f_ FetchIndexDescription) SetName(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setName:"), value)
}

// An array of fetch index element descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887054-elements?language=objc
func (f_ FetchIndexDescription) Elements() []FetchIndexElementDescription {
	rv := objc.Call[[]FetchIndexElementDescription](f_, objc.Sel("elements"))
	return rv
}

// An array of fetch index element descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887054-elements?language=objc
func (f_ FetchIndexDescription) SetElements(value []IFetchIndexElementDescription) {
	objc.Call[objc.Void](f_, objc.Sel("setElements:"), value)
}

// A predicate that selects rows for indexing, if the index is a partial index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887058-partialindexpredicate?language=objc
func (f_ FetchIndexDescription) PartialIndexPredicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](f_, objc.Sel("partialIndexPredicate"))
	return rv
}

// A predicate that selects rows for indexing, if the index is a partial index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexdescription/2887058-partialindexpredicate?language=objc
func (f_ FetchIndexDescription) SetPartialIndexPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](f_, objc.Sel("setPartialIndexPredicate:"), objc.Ptr(value))
}
