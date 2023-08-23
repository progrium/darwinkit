// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A File Provider extension in which the system replicates the contents on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension?language=objc
type PFileProviderReplicatedExtension interface {
	// optional
	CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate objc.Object, fields FileProviderItemFields, url foundation.URL, options FileProviderCreateItemOptions, request FileProviderRequest, completionHandler func(createdItem objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.IProgress
	HasCreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler() bool

	// optional
	ImportDidFinishWithCompletionHandler(completionHandler func())
	HasImportDidFinishWithCompletionHandler() bool

	// optional
	MaterializedItemsDidChangeWithCompletionHandler(completionHandler func())
	HasMaterializedItemsDidChangeWithCompletionHandler() bool

	// optional
	DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier FileProviderItemIdentifier, version FileProviderItemVersion, options FileProviderDeleteItemOptions, request FileProviderRequest, completionHandler func(arg0 foundation.Error)) foundation.IProgress
	HasDeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler() bool

	// optional
	ItemForIdentifierRequestCompletionHandler(identifier FileProviderItemIdentifier, request FileProviderRequest, completionHandler func(arg0 objc.Object, arg1 foundation.Error)) foundation.IProgress
	HasItemForIdentifierRequestCompletionHandler() bool

	// optional
	InitWithDomain(domain FileProviderDomain) objc.IObject
	HasInitWithDomain() bool

	// optional
	PendingItemsDidChangeWithCompletionHandler(completionHandler func())
	HasPendingItemsDidChangeWithCompletionHandler() bool

	// optional
	Invalidate()
	HasInvalidate() bool

	// optional
	ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item objc.Object, version FileProviderItemVersion, changedFields FileProviderItemFields, newContents foundation.URL, options FileProviderModifyItemOptions, request FileProviderRequest, completionHandler func(item objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.IProgress
	HasModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler() bool

	// optional
	FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion FileProviderItemVersion, request FileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.IProgress
	HasFetchContentsForItemWithIdentifierVersionRequestCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderReplicatedExtension] protocol.
type FileProviderReplicatedExtensionWrapper struct {
	objc.Object
}

func (f_ FileProviderReplicatedExtensionWrapper) HasCreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("createItemBasedOnTemplate:fields:contents:options:request:completionHandler:"))
}

// Tells the file provider to create or import an item based on a template. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656549-createitembasedontemplate?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate objc.IObject, fields FileProviderItemFields, url foundation.IURL, options FileProviderCreateItemOptions, request IFileProviderRequest, completionHandler func(createdItem objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("createItemBasedOnTemplate:fields:contents:options:request:completionHandler:"), objc.Ptr(itemTemplate), fields, objc.Ptr(url), options, objc.Ptr(request), completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionWrapper) HasImportDidFinishWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("importDidFinishWithCompletionHandler:"))
}

// Tells the File Provider extension that the system finished importing items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553304-importdidfinishwithcompletionhan?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) ImportDidFinishWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("importDidFinishWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionWrapper) HasMaterializedItemsDidChangeWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("materializedItemsDidChangeWithCompletionHandler:"))
}

// Tells the file provider that the set of materialized items changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553308-materializeditemsdidchangewithco?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) MaterializedItemsDidChangeWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("materializedItemsDidChangeWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionWrapper) HasDeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("deleteItemWithIdentifier:baseVersion:options:request:completionHandler:"))
}

// Tells the file provider to delete an item forever. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656550-deleteitemwithidentifier?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier FileProviderItemIdentifier, version IFileProviderItemVersion, options FileProviderDeleteItemOptions, request IFileProviderRequest, completionHandler func(arg0 foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("deleteItemWithIdentifier:baseVersion:options:request:completionHandler:"), identifier, objc.Ptr(version), options, objc.Ptr(request), completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionWrapper) HasItemForIdentifierRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("itemForIdentifier:request:completionHandler:"))
}

// Asks the file provider for the metadata of the provided item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656551-itemforidentifier?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) ItemForIdentifierRequestCompletionHandler(identifier FileProviderItemIdentifier, request IFileProviderRequest, completionHandler func(arg0 objc.Object, arg1 foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("itemForIdentifier:request:completionHandler:"), identifier, objc.Ptr(request), completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionWrapper) HasInitWithDomain() bool {
	return f_.RespondsToSelector(objc.Sel("initWithDomain:"))
}

// Creates an instance of the file provider for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553305-initwithdomain?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) InitWithDomain(domain IFileProviderDomain) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("initWithDomain:"), objc.Ptr(domain))
	return rv
}

func (f_ FileProviderReplicatedExtensionWrapper) HasPendingItemsDidChangeWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("pendingItemsDidChangeWithCompletionHandler:"))
}

// Tells the file provider extension that the set of pending items has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3727821-pendingitemsdidchangewithcomplet?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) PendingItemsDidChangeWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("pendingItemsDidChangeWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionWrapper) HasInvalidate() bool {
	return f_.RespondsToSelector(objc.Sel("invalidate"))
}

// Tells the file provider to perform any necessary cleanup so that the system can deallocate it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553306-invalidate?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) Invalidate() {
	objc.Call[objc.Void](f_, objc.Sel("invalidate"))
}

func (f_ FileProviderReplicatedExtensionWrapper) HasModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("modifyItem:baseVersion:changedFields:contents:options:request:completionHandler:"))
}

// Tells the file provider that an item’s content or metadata changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656552-modifyitem?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item objc.IObject, version IFileProviderItemVersion, changedFields FileProviderItemFields, newContents foundation.IURL, options FileProviderModifyItemOptions, request IFileProviderRequest, completionHandler func(item objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("modifyItem:baseVersion:changedFields:contents:options:request:completionHandler:"), objc.Ptr(item), objc.Ptr(version), changedFields, objc.Ptr(newContents), options, objc.Ptr(request), completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionWrapper) HasFetchContentsForItemWithIdentifierVersionRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchContentsForItemWithIdentifier:version:request:completionHandler:"))
}

// Tells the file provider to download the requested item from remote storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553303-fetchcontentsforitemwithidentifi?language=objc
func (f_ FileProviderReplicatedExtensionWrapper) FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion IFileProviderItemVersion, request IFileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchContentsForItemWithIdentifier:version:request:completionHandler:"), itemIdentifier, objc.Ptr(requestedVersion), objc.Ptr(request), completionHandler)
	return rv
}
