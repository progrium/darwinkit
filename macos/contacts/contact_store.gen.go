// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactStore] class.
var ContactStoreClass = _ContactStoreClass{objc.GetClass("CNContactStore")}

type _ContactStoreClass struct {
	objc.Class
}

// An interface definition for the [ContactStore] class.
type IContactStore interface {
	objc.IObject
	EnumerateContactsWithFetchRequestErrorUsingBlock(fetchRequest IContactFetchRequest, error foundation.IError, block func(contact Contact, stop *bool)) bool
	RequestAccessForEntityTypeCompletionHandler(entityType EntityType, completionHandler func(granted bool, error foundation.Error))
	UnifiedContactWithIdentifierKeysToFetchError(identifier string, keys []objc.IObject, error foundation.IError) Contact
	EnumeratorForContactFetchRequestError(request IContactFetchRequest, error foundation.IError) FetchResult
	ContainersMatchingPredicateError(predicate foundation.IPredicate, error foundation.IError) []Container
	ExecuteSaveRequestError(saveRequest ISaveRequest, error foundation.IError) bool
	EnumeratorForChangeHistoryFetchRequestError(request IChangeHistoryFetchRequest, error foundation.IError) FetchResult
	GroupsMatchingPredicateError(predicate foundation.IPredicate, error foundation.IError) []Group
	DefaultContainerIdentifier() string
	UnifiedMeContactWithKeysToFetchError(keys []objc.IObject, error foundation.IError) Contact
	UnifiedContactsMatchingPredicateKeysToFetchError(predicate foundation.IPredicate, keys []objc.IObject, error foundation.IError) []Contact
	CurrentHistoryToken() []byte
}

// The object that fetches and saves contacts, groups, and containers from the user's Contacts database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore?language=objc
type ContactStore struct {
	objc.Object
}

func ContactStoreFrom(ptr unsafe.Pointer) ContactStore {
	return ContactStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactStoreClass) Alloc() ContactStore {
	rv := objc.Call[ContactStore](cc, objc.Sel("alloc"))
	return rv
}

func ContactStore_Alloc() ContactStore {
	return ContactStoreClass.Alloc()
}

func (cc _ContactStoreClass) New() ContactStore {
	rv := objc.Call[ContactStore](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactStore() ContactStore {
	return ContactStoreClass.New()
}

func (c_ ContactStore) Init() ContactStore {
	rv := objc.Call[ContactStore](c_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the enumeration of all contacts matching a contact fetch request executes successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1402849-enumeratecontactswithfetchreques?language=objc
func (c_ ContactStore) EnumerateContactsWithFetchRequestErrorUsingBlock(fetchRequest IContactFetchRequest, error foundation.IError, block func(contact Contact, stop *bool)) bool {
	rv := objc.Call[bool](c_, objc.Sel("enumerateContactsWithFetchRequest:error:usingBlock:"), objc.Ptr(fetchRequest), objc.Ptr(error), block)
	return rv
}

// Requests access to the user's contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1402873-requestaccessforentitytype?language=objc
func (c_ ContactStore) RequestAccessForEntityTypeCompletionHandler(entityType EntityType, completionHandler func(granted bool, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("requestAccessForEntityType:completionHandler:"), entityType, completionHandler)
}

// Fetches a unified contact for the specified contact identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403256-unifiedcontactwithidentifier?language=objc
func (c_ ContactStore) UnifiedContactWithIdentifierKeysToFetchError(identifier string, keys []objc.IObject, error foundation.IError) Contact {
	rv := objc.Call[Contact](c_, objc.Sel("unifiedContactWithIdentifier:keysToFetch:error:"), identifier, keys, objc.Ptr(error))
	return rv
}

// Enumerates a contact fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/3294191-enumeratorforcontactfetchrequest?language=objc
func (c_ ContactStore) EnumeratorForContactFetchRequestError(request IContactFetchRequest, error foundation.IError) FetchResult {
	rv := objc.Call[FetchResult](c_, objc.Sel("enumeratorForContactFetchRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Fetches all containers matching the specified predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403238-containersmatchingpredicate?language=objc
func (c_ ContactStore) ContainersMatchingPredicateError(predicate foundation.IPredicate, error foundation.IError) []Container {
	rv := objc.Call[[]Container](c_, objc.Sel("containersMatchingPredicate:error:"), objc.Ptr(predicate), objc.Ptr(error))
	return rv
}

// Executes a save request and returns success or failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403230-executesaverequest?language=objc
func (c_ ContactStore) ExecuteSaveRequestError(saveRequest ISaveRequest, error foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("executeSaveRequest:error:"), objc.Ptr(saveRequest), objc.Ptr(error))
	return rv
}

// Enumerates a change history fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/3294190-enumeratorforchangehistoryfetchr?language=objc
func (c_ ContactStore) EnumeratorForChangeHistoryFetchRequestError(request IChangeHistoryFetchRequest, error foundation.IError) FetchResult {
	rv := objc.Call[FetchResult](c_, objc.Sel("enumeratorForChangeHistoryFetchRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Returns the current authorization status to access the contact data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403143-authorizationstatusforentitytype?language=objc
func (cc _ContactStoreClass) AuthorizationStatusForEntityType(entityType EntityType) AuthorizationStatus {
	rv := objc.Call[AuthorizationStatus](cc, objc.Sel("authorizationStatusForEntityType:"), entityType)
	return rv
}

// Returns the current authorization status to access the contact data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403143-authorizationstatusforentitytype?language=objc
func ContactStore_AuthorizationStatusForEntityType(entityType EntityType) AuthorizationStatus {
	return ContactStoreClass.AuthorizationStatusForEntityType(entityType)
}

// Fetches all groups matching the specified predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403121-groupsmatchingpredicate?language=objc
func (c_ ContactStore) GroupsMatchingPredicateError(predicate foundation.IPredicate, error foundation.IError) []Group {
	rv := objc.Call[[]Group](c_, objc.Sel("groupsMatchingPredicate:error:"), objc.Ptr(predicate), objc.Ptr(error))
	return rv
}

// Returns the identifier of the default container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403159-defaultcontaineridentifier?language=objc
func (c_ ContactStore) DefaultContainerIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("defaultContainerIdentifier"))
	return rv
}

// Fetches the unified contact that’s the me card. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1402987-unifiedmecontactwithkeystofetch?language=objc
func (c_ ContactStore) UnifiedMeContactWithKeysToFetchError(keys []objc.IObject, error foundation.IError) Contact {
	rv := objc.Call[Contact](c_, objc.Sel("unifiedMeContactWithKeysToFetch:error:"), keys, objc.Ptr(error))
	return rv
}

// Fetches all unified contacts matching the specified predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/1403266-unifiedcontactsmatchingpredicate?language=objc
func (c_ ContactStore) UnifiedContactsMatchingPredicateKeysToFetchError(predicate foundation.IPredicate, keys []objc.IObject, error foundation.IError) []Contact {
	rv := objc.Call[[]Contact](c_, objc.Sel("unifiedContactsMatchingPredicate:keysToFetch:error:"), objc.Ptr(predicate), keys, objc.Ptr(error))
	return rv
}

// The current history token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactstore/3113739-currenthistorytoken?language=objc
func (c_ ContactStore) CurrentHistoryToken() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("currentHistoryToken"))
	return rv
}
