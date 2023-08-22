// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderManager] class.
var FileProviderManagerClass = _FileProviderManagerClass{objc.GetClass("NSFileProviderManager")}

type _FileProviderManagerClass struct {
	objc.Class
}

// An interface definition for the [FileProviderManager] class.
type IFileProviderManager interface {
	objc.IObject
	ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error))
	SignalErrorResolvedCompletionHandler(error foundation.IError, completionHandler func(error foundation.Error))
	WaitForStabilizationWithCompletionHandler(completionHandler func(error foundation.Error))
	EvictItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error))
	RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.IURLSessionTask, identifier FileProviderItemIdentifier, completion func(error foundation.Error))
	DisconnectWithReasonOptionsCompletionHandler(localizedReason string, options FileProviderManagerDisconnectionOptions, completionHandler func(error foundation.Error))
	SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier FileProviderItemIdentifier, completion func(error foundation.Error))
	ListAvailableTestingOperationsWithError(error foundation.IError) []FileProviderTestingOperationWrapper
	WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error))
	GlobalProgressForKind(kind foundation.ProgressFileOperationKind) foundation.Progress
	EnumeratorForMaterializedItems() FileProviderEnumeratorWrapper
	GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(userVisibleFile foundation.URL, error foundation.Error))
	ReconnectWithCompletionHandler(completionHandler func(error foundation.Error))
	TemporaryDirectoryURLWithError(error foundation.IError) foundation.URL
	RunTestingOperationsError(operations []PFileProviderTestingOperation, error foundation.IError) foundation.Dictionary
	EnumeratorForPendingItems() FileProviderPendingSetEnumeratorWrapper
}

// A manager object that you use to communicate with the file provider from either your app or your File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager?language=objc
type FileProviderManager struct {
	objc.Object
}

func FileProviderManagerFrom(ptr unsafe.Pointer) FileProviderManager {
	return FileProviderManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileProviderManagerClass) ManagerForDomain(domain IFileProviderDomain) FileProviderManager {
	rv := objc.Call[FileProviderManager](fc, objc.Sel("managerForDomain:"), objc.Ptr(domain))
	return rv
}

// Returns a newly created file provider manager for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2882047-managerfordomain?language=objc
func FileProviderManager_ManagerForDomain(domain IFileProviderDomain) FileProviderManager {
	return FileProviderManagerClass.ManagerForDomain(domain)
}

func (fc _FileProviderManagerClass) Alloc() FileProviderManager {
	rv := objc.Call[FileProviderManager](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderManager_Alloc() FileProviderManager {
	return FileProviderManagerClass.Alloc()
}

func (fc _FileProviderManagerClass) New() FileProviderManager {
	rv := objc.Call[FileProviderManager](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderManager() FileProviderManager {
	return FileProviderManagerClass.New()
}

func (f_ FileProviderManager) Init() FileProviderManager {
	rv := objc.Call[FileProviderManager](f_, objc.Sel("init"))
	return rv
}

// Tells the system to reimport the item and its content recursively. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3181165-reimportitemsbelowitemwithidenti?language=objc
func (f_ FileProviderManager) ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("reimportItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Indicates a resolved error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3656534-signalerrorresolved?language=objc
func (f_ FileProviderManager) SignalErrorResolvedCompletionHandler(error foundation.IError, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("signalErrorResolved:completionHandler:"), objc.Ptr(error), completionHandler)
}

// Requests a notification after the domain stabilizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3567073-waitforstabilizationwithcompleti?language=objc
func (f_ FileProviderManager) WaitForStabilizationWithCompletionHandler(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("waitForStabilizationWithCompletionHandler:"), completionHandler)
}

// Returns all of the File Provider extension's domains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2882045-getdomainswithcompletionhandler?language=objc
func (fc _FileProviderManagerClass) GetDomainsWithCompletionHandler(completionHandler func(domains []FileProviderDomain, error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("getDomainsWithCompletionHandler:"), completionHandler)
}

// Returns all of the File Provider extension's domains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2882045-getdomainswithcompletionhandler?language=objc
func FileProviderManager_GetDomainsWithCompletionHandler(completionHandler func(domains []FileProviderDomain, error foundation.Error)) {
	FileProviderManagerClass.GetDomainsWithCompletionHandler(completionHandler)
}

// Removes all domains from the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2882044-removealldomainswithcompletionha?language=objc
func (fc _FileProviderManagerClass) RemoveAllDomainsWithCompletionHandler(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("removeAllDomainsWithCompletionHandler:"), completionHandler)
}

// Removes all domains from the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2882044-removealldomainswithcompletionha?language=objc
func FileProviderManager_RemoveAllDomainsWithCompletionHandler(completionHandler func(error foundation.Error)) {
	FileProviderManagerClass.RemoveAllDomainsWithCompletionHandler(completionHandler)
}

// Asks the system to remove an item from its cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3191974-evictitemwithidentifier?language=objc
func (f_ FileProviderManager) EvictItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("evictItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Registers the URL session task responsible for the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890932-registerurlsessiontask?language=objc
func (f_ FileProviderManager) RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.IURLSessionTask, identifier FileProviderItemIdentifier, completion func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("registerURLSessionTask:forItemWithIdentifier:completionHandler:"), objc.Ptr(task), identifier, completion)
}

// Disconnects the domain from the extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3603576-disconnectwithreason?language=objc
func (f_ FileProviderManager) DisconnectWithReasonOptionsCompletionHandler(localizedReason string, options FileProviderManagerDisconnectionOptions, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("disconnectWithReason:options:completionHandler:"), localizedReason, options, completionHandler)
}

// Alerts the system to changes in the specified folder’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890931-signalenumeratorforcontaineritem?language=objc
func (f_ FileProviderManager) SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier FileProviderItemIdentifier, completion func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("signalEnumeratorForContainerItemIdentifier:completionHandler:"), containerItemIdentifier, completion)
}

// Lists all the operations that are ready for scheduling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3727823-listavailabletestingoperationswi?language=objc
func (f_ FileProviderManager) ListAvailableTestingOperationsWithError(error foundation.IError) []FileProviderTestingOperationWrapper {
	rv := objc.Call[[]FileProviderTestingOperationWrapper](f_, objc.Sel("listAvailableTestingOperationsWithError:"), objc.Ptr(error))
	return rv
}

// Requests a notification after the system completes all the specified changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3585198-waitforchangesonitemsbelowitemwi?language=objc
func (f_ FileProviderManager) WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("waitForChangesOnItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3731236-globalprogressforkind?language=objc
func (f_ FileProviderManager) GlobalProgressForKind(kind foundation.ProgressFileOperationKind) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("globalProgressForKind:"), kind)
	return rv
}

// Adds a domain to the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890934-adddomain?language=objc
func (fc _FileProviderManagerClass) AddDomainCompletionHandler(domain IFileProviderDomain, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("addDomain:completionHandler:"), objc.Ptr(domain), completionHandler)
}

// Adds a domain to the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890934-adddomain?language=objc
func FileProviderManager_AddDomainCompletionHandler(domain IFileProviderDomain, completionHandler func(error foundation.Error)) {
	FileProviderManagerClass.AddDomainCompletionHandler(domain, completionHandler)
}

// Returns the identifier and domain for a user-visible URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3074519-getidentifierforuservisiblefilea?language=objc
func (fc _FileProviderManagerClass) GetIdentifierForUserVisibleFileAtURLCompletionHandler(url foundation.IURL, completionHandler func(itemIdentifier FileProviderItemIdentifier, domainIdentifier FileProviderDomainIdentifier, error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("getIdentifierForUserVisibleFileAtURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

// Returns the identifier and domain for a user-visible URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3074519-getidentifierforuservisiblefilea?language=objc
func FileProviderManager_GetIdentifierForUserVisibleFileAtURLCompletionHandler(url foundation.IURL, completionHandler func(itemIdentifier FileProviderItemIdentifier, domainIdentifier FileProviderDomainIdentifier, error foundation.Error)) {
	FileProviderManagerClass.GetIdentifierForUserVisibleFileAtURLCompletionHandler(url, completionHandler)
}

// Returns an enumerator for all the items the system currently stores on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3181167-enumeratorformaterializeditems?language=objc
func (f_ FileProviderManager) EnumeratorForMaterializedItems() FileProviderEnumeratorWrapper {
	rv := objc.Call[FileProviderEnumeratorWrapper](f_, objc.Sel("enumeratorForMaterializedItems"))
	return rv
}

// Creates a new domain that takes ownership of on-disk data that your app previously managed without a file provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3181164-importdomain?language=objc
func (fc _FileProviderManagerClass) ImportDomainFromDirectoryAtURLCompletionHandler(domain IFileProviderDomain, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("importDomain:fromDirectoryAtURL:completionHandler:"), objc.Ptr(domain), objc.Ptr(url), completionHandler)
}

// Creates a new domain that takes ownership of on-disk data that your app previously managed without a file provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3181164-importdomain?language=objc
func FileProviderManager_ImportDomainFromDirectoryAtURLCompletionHandler(domain IFileProviderDomain, url foundation.IURL, completionHandler func(error foundation.Error)) {
	FileProviderManagerClass.ImportDomainFromDirectoryAtURLCompletionHandler(domain, url, completionHandler)
}

// Returns the user-visible URL for an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3074520-getuservisibleurlforitemidentifi?language=objc
func (f_ FileProviderManager) GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(userVisibleFile foundation.URL, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("getUserVisibleURLForItemIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Reconnects the domain with the extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3603577-reconnectwithcompletionhandler?language=objc
func (f_ FileProviderManager) ReconnectWithCompletionHandler(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("reconnectWithCompletionHandler:"), completionHandler)
}

// Removes a domain from the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890933-removedomain?language=objc
func (fc _FileProviderManagerClass) RemoveDomainCompletionHandler(domain IFileProviderDomain, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](fc, objc.Sel("removeDomain:completionHandler:"), objc.Ptr(domain), completionHandler)
}

// Removes a domain from the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/2890933-removedomain?language=objc
func FileProviderManager_RemoveDomainCompletionHandler(domain IFileProviderDomain, completionHandler func(error foundation.Error)) {
	FileProviderManagerClass.RemoveDomainCompletionHandler(domain, completionHandler)
}

// Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3656535-temporarydirectoryurlwitherror?language=objc
func (f_ FileProviderManager) TemporaryDirectoryURLWithError(error foundation.IError) foundation.URL {
	rv := objc.Call[foundation.URL](f_, objc.Sel("temporaryDirectoryURLWithError:"), objc.Ptr(error))
	return rv
}

// Asks the system to schedule and execute the specified operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3727824-runtestingoperations?language=objc
func (f_ FileProviderManager) RunTestingOperationsError(operations []PFileProviderTestingOperation, error foundation.IError) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("runTestingOperations:error:"), operations, objc.Ptr(error))
	return rv
}

// Returns an enumerator for the set of pending items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/3727809-enumeratorforpendingitems?language=objc
func (f_ FileProviderManager) EnumeratorForPendingItems() FileProviderPendingSetEnumeratorWrapper {
	rv := objc.Call[FileProviderPendingSetEnumeratorWrapper](f_, objc.Sel("enumeratorForPendingItems"))
	return rv
}
