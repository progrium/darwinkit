// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactFetchRequest] class.
var ContactFetchRequestClass = _ContactFetchRequestClass{objc.GetClass("CNContactFetchRequest")}

type _ContactFetchRequestClass struct {
	objc.Class
}

// An interface definition for the [ContactFetchRequest] class.
type IContactFetchRequest interface {
	IFetchRequest
	MutableObjects() bool
	SetMutableObjects(value bool)
	KeysToFetch() []objc.Object
	SetKeysToFetch(value []objc.IObject)
	SortOrder() ContactSortOrder
	SetSortOrder(value ContactSortOrder)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
	UnifyResults() bool
	SetUnifyResults(value bool)
}

// An object that defines the options to use when fetching contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest?language=objc
type ContactFetchRequest struct {
	FetchRequest
}

func ContactFetchRequestFrom(ptr unsafe.Pointer) ContactFetchRequest {
	return ContactFetchRequest{
		FetchRequest: FetchRequestFrom(ptr),
	}
}

func (c_ ContactFetchRequest) InitWithKeysToFetch(keysToFetch []objc.IObject) ContactFetchRequest {
	rv := objc.Call[ContactFetchRequest](c_, objc.Sel("initWithKeysToFetch:"), keysToFetch)
	return rv
}

// Creates a fetch request for the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403020-initwithkeystofetch?language=objc
func NewContactFetchRequestWithKeysToFetch(keysToFetch []objc.IObject) ContactFetchRequest {
	instance := ContactFetchRequestClass.Alloc().InitWithKeysToFetch(keysToFetch)
	instance.Autorelease()
	return instance
}

func (cc _ContactFetchRequestClass) Alloc() ContactFetchRequest {
	rv := objc.Call[ContactFetchRequest](cc, objc.Sel("alloc"))
	return rv
}

func ContactFetchRequest_Alloc() ContactFetchRequest {
	return ContactFetchRequestClass.Alloc()
}

func (cc _ContactFetchRequestClass) New() ContactFetchRequest {
	rv := objc.Call[ContactFetchRequest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactFetchRequest() ContactFetchRequest {
	return ContactFetchRequestClass.New()
}

func (c_ ContactFetchRequest) Init() ContactFetchRequest {
	rv := objc.Call[ContactFetchRequest](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether to return mutable contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1402835-mutableobjects?language=objc
func (c_ ContactFetchRequest) MutableObjects() bool {
	rv := objc.Call[bool](c_, objc.Sel("mutableObjects"))
	return rv
}

// A Boolean value that indicates whether to return mutable contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1402835-mutableobjects?language=objc
func (c_ ContactFetchRequest) SetMutableObjects(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setMutableObjects:"), value)
}

// The properties to fetch in the returned contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403360-keystofetch?language=objc
func (c_ ContactFetchRequest) KeysToFetch() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("keysToFetch"))
	return rv
}

// The properties to fetch in the returned contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403360-keystofetch?language=objc
func (c_ ContactFetchRequest) SetKeysToFetch(value []objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setKeysToFetch:"), value)
}

// The sort order for contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403262-sortorder?language=objc
func (c_ ContactFetchRequest) SortOrder() ContactSortOrder {
	rv := objc.Call[ContactSortOrder](c_, objc.Sel("sortOrder"))
	return rv
}

// The sort order for contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403262-sortorder?language=objc
func (c_ ContactFetchRequest) SetSortOrder(value ContactSortOrder) {
	objc.Call[objc.Void](c_, objc.Sel("setSortOrder:"), value)
}

// The predicate to match contacts against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403080-predicate?language=objc
func (c_ ContactFetchRequest) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](c_, objc.Sel("predicate"))
	return rv
}

// The predicate to match contacts against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403080-predicate?language=objc
func (c_ ContactFetchRequest) SetPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](c_, objc.Sel("setPredicate:"), objc.Ptr(value))
}

// A Boolean value that indicates whether to return linked contacts as unified contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403293-unifyresults?language=objc
func (c_ ContactFetchRequest) UnifyResults() bool {
	rv := objc.Call[bool](c_, objc.Sel("unifyResults"))
	return rv
}

// A Boolean value that indicates whether to return linked contacts as unified contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactfetchrequest/1403293-unifyresults?language=objc
func (c_ ContactFetchRequest) SetUnifyResults(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setUnifyResults:"), value)
}
