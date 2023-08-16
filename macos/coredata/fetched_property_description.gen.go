// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchedPropertyDescription] class.
var FetchedPropertyDescriptionClass = _FetchedPropertyDescriptionClass{objc.GetClass("NSFetchedPropertyDescription")}

type _FetchedPropertyDescriptionClass struct {
	objc.Class
}

// An interface definition for the [FetchedPropertyDescription] class.
type IFetchedPropertyDescription interface {
	IPropertyDescription
	FetchRequest() FetchRequest
	SetFetchRequest(value IFetchRequest)
}

// A description object used to define which properties are fetched from Core Data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedpropertydescription?language=objc
type FetchedPropertyDescription struct {
	PropertyDescription
}

func FetchedPropertyDescriptionFrom(ptr unsafe.Pointer) FetchedPropertyDescription {
	return FetchedPropertyDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

func (fc _FetchedPropertyDescriptionClass) Alloc() FetchedPropertyDescription {
	rv := objc.Call[FetchedPropertyDescription](fc, objc.Sel("alloc"))
	return rv
}

func FetchedPropertyDescription_Alloc() FetchedPropertyDescription {
	return FetchedPropertyDescriptionClass.Alloc()
}

func (fc _FetchedPropertyDescriptionClass) New() FetchedPropertyDescription {
	rv := objc.Call[FetchedPropertyDescription](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchedPropertyDescription() FetchedPropertyDescription {
	return FetchedPropertyDescriptionClass.New()
}

func (f_ FetchedPropertyDescription) Init() FetchedPropertyDescription {
	rv := objc.Call[FetchedPropertyDescription](f_, objc.Sel("init"))
	return rv
}

// The fetch request of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedpropertydescription/1494679-fetchrequest?language=objc
func (f_ FetchedPropertyDescription) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](f_, objc.Sel("fetchRequest"))
	return rv
}

// The fetch request of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedpropertydescription/1494679-fetchrequest?language=objc
func (f_ FetchedPropertyDescription) SetFetchRequest(value IFetchRequest) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchRequest:"), objc.Ptr(value))
}
