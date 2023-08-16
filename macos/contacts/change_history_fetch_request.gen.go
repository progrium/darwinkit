// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryFetchRequest] class.
var ChangeHistoryFetchRequestClass = _ChangeHistoryFetchRequestClass{objc.GetClass("CNChangeHistoryFetchRequest")}

type _ChangeHistoryFetchRequestClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryFetchRequest] class.
type IChangeHistoryFetchRequest interface {
	IFetchRequest
	MutableObjects() bool
	SetMutableObjects(value bool)
	ExcludedTransactionAuthors() []string
	SetExcludedTransactionAuthors(value []string)
	ShouldUnifyResults() bool
	SetShouldUnifyResults(value bool)
	StartingToken() []byte
	SetStartingToken(value []byte)
	AdditionalContactKeyDescriptors() []objc.Object
	SetAdditionalContactKeyDescriptors(value []objc.IObject)
	IncludeGroupChanges() bool
	SetIncludeGroupChanges(value bool)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest?language=objc
type ChangeHistoryFetchRequest struct {
	FetchRequest
}

func ChangeHistoryFetchRequestFrom(ptr unsafe.Pointer) ChangeHistoryFetchRequest {
	return ChangeHistoryFetchRequest{
		FetchRequest: FetchRequestFrom(ptr),
	}
}

func (cc _ChangeHistoryFetchRequestClass) Alloc() ChangeHistoryFetchRequest {
	rv := objc.Call[ChangeHistoryFetchRequest](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryFetchRequest_Alloc() ChangeHistoryFetchRequest {
	return ChangeHistoryFetchRequestClass.Alloc()
}

func (cc _ChangeHistoryFetchRequestClass) New() ChangeHistoryFetchRequest {
	rv := objc.Call[ChangeHistoryFetchRequest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryFetchRequest() ChangeHistoryFetchRequest {
	return ChangeHistoryFetchRequestClass.New()
}

func (c_ ChangeHistoryFetchRequest) Init() ChangeHistoryFetchRequest {
	rv := objc.Call[ChangeHistoryFetchRequest](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294187-mutableobjects?language=objc
func (c_ ChangeHistoryFetchRequest) MutableObjects() bool {
	rv := objc.Call[bool](c_, objc.Sel("mutableObjects"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294187-mutableobjects?language=objc
func (c_ ChangeHistoryFetchRequest) SetMutableObjects(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setMutableObjects:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294185-excludedtransactionauthors?language=objc
func (c_ ChangeHistoryFetchRequest) ExcludedTransactionAuthors() []string {
	rv := objc.Call[[]string](c_, objc.Sel("excludedTransactionAuthors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294185-excludedtransactionauthors?language=objc
func (c_ ChangeHistoryFetchRequest) SetExcludedTransactionAuthors(value []string) {
	objc.Call[objc.Void](c_, objc.Sel("setExcludedTransactionAuthors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294188-shouldunifyresults?language=objc
func (c_ ChangeHistoryFetchRequest) ShouldUnifyResults() bool {
	rv := objc.Call[bool](c_, objc.Sel("shouldUnifyResults"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294188-shouldunifyresults?language=objc
func (c_ ChangeHistoryFetchRequest) SetShouldUnifyResults(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setShouldUnifyResults:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294189-startingtoken?language=objc
func (c_ ChangeHistoryFetchRequest) StartingToken() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("startingToken"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294189-startingtoken?language=objc
func (c_ ChangeHistoryFetchRequest) SetStartingToken(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setStartingToken:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294184-additionalcontactkeydescriptors?language=objc
func (c_ ChangeHistoryFetchRequest) AdditionalContactKeyDescriptors() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("additionalContactKeyDescriptors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294184-additionalcontactkeydescriptors?language=objc
func (c_ ChangeHistoryFetchRequest) SetAdditionalContactKeyDescriptors(value []objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setAdditionalContactKeyDescriptors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294186-includegroupchanges?language=objc
func (c_ ChangeHistoryFetchRequest) IncludeGroupChanges() bool {
	rv := objc.Call[bool](c_, objc.Sel("includeGroupChanges"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryfetchrequest/3294186-includegroupchanges?language=objc
func (c_ ChangeHistoryFetchRequest) SetIncludeGroupChanges(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setIncludeGroupChanges:"), value)
}
