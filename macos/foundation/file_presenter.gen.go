// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface a file coordinator uses to inform an object presenting a file about changes to that file made elsewhere in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter?language=objc
type PFilePresenter interface {
	// optional
	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler func(errorOrNil Error))
	HasAccommodatePresentedItemDeletionWithCompletionHandler() bool

	// optional
	PresentedItemDidMoveToURL(newURL URL)
	HasPresentedItemDidMoveToURL() bool

	// optional
	PresentedSubitemAtURLDidGainVersion(url URL, version FileVersion)
	HasPresentedSubitemAtURLDidGainVersion() bool

	// optional
	SavePresentedItemChangesWithCompletionHandler(completionHandler func(errorOrNil Error))
	HasSavePresentedItemChangesWithCompletionHandler() bool

	// optional
	PresentedItemDidChange()
	HasPresentedItemDidChange() bool

	// optional
	PresentedItemDidLoseVersion(version FileVersion)
	HasPresentedItemDidLoseVersion() bool

	// optional
	RelinquishPresentedItemToReader(reader func(arg0 func()))
	HasRelinquishPresentedItemToReader() bool

	// optional
	PresentedItemDidGainVersion(version FileVersion)
	HasPresentedItemDidGainVersion() bool

	// optional
	PresentedItemDidResolveConflictVersion(version FileVersion)
	HasPresentedItemDidResolveConflictVersion() bool

	// optional
	PresentedItemDidChangeUbiquityAttributes(attributes Set)
	HasPresentedItemDidChangeUbiquityAttributes() bool

	// optional
	RelinquishPresentedItemToWriter(writer func(arg0 func()))
	HasRelinquishPresentedItemToWriter() bool

	// optional
	PresentedSubitemDidAppearAtURL(url URL)
	HasPresentedSubitemDidAppearAtURL() bool

	// optional
	PresentedSubitemDidChangeAtURL(url URL)
	HasPresentedSubitemDidChangeAtURL() bool

	// optional
	AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url URL, completionHandler func(errorOrNil Error))
	HasAccommodatePresentedSubitemDeletionAtURLCompletionHandler() bool

	// optional
	PresentedItemOperationQueue() IOperationQueue
	HasPresentedItemOperationQueue() bool

	// optional
	PresentedItemURL() IURL
	HasPresentedItemURL() bool

	// optional
	ObservedPresentedItemUbiquityAttributes() ISet
	HasObservedPresentedItemUbiquityAttributes() bool

	// optional
	PrimaryPresentedItemURL() IURL
	HasPrimaryPresentedItemURL() bool
}

// A concrete type wrapper for the [PFilePresenter] protocol.
type FilePresenterWrapper struct {
	objc.Object
}

func (f_ FilePresenterWrapper) HasAccommodatePresentedItemDeletionWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"))
}

// Tells your object that its presented item is about to be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414732-accommodatepresenteditemdeletion?language=objc
func (f_ FilePresenterWrapper) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), completionHandler)
}

func (f_ FilePresenterWrapper) HasPresentedItemDidMoveToURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidMoveToURL:"))
}

// Tells your object that the presented item moved or was renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1417861-presenteditemdidmovetourl?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidMoveToURL(newURL IURL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidMoveToURL:"), objc.Ptr(newURL))
}

func (f_ FilePresenterWrapper) HasPresentedSubitemAtURLDidGainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemAtURL:didGainVersion:"))
}

// Tells the delegate that the item inside the presented directory gained a new version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415472-presentedsubitematurl?language=objc
func (f_ FilePresenterWrapper) PresentedSubitemAtURLDidGainVersion(url IURL, version IFileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemAtURL:didGainVersion:"), objc.Ptr(url), objc.Ptr(version))
}

func (f_ FilePresenterWrapper) HasSavePresentedItemChangesWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("savePresentedItemChangesWithCompletionHandler:"))
}

// Tells your object to save any unsaved changes for the presented item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414407-savepresenteditemchangeswithcomp?language=objc
func (f_ FilePresenterWrapper) SavePresentedItemChangesWithCompletionHandler(completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), completionHandler)
}

func (f_ FilePresenterWrapper) HasPresentedItemDidChange() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidChange"))
}

// Tells your object that the presented item’s contents or attributes changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1416103-presenteditemdidchange?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidChange() {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidChange"))
}

func (f_ FilePresenterWrapper) HasPresentedItemDidLoseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidLoseVersion:"))
}

// Tells the delegate that a version of the file or file package was removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1417258-presenteditemdidloseversion?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidLoseVersion(version IFileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidLoseVersion:"), objc.Ptr(version))
}

func (f_ FilePresenterWrapper) HasRelinquishPresentedItemToReader() bool {
	return f_.RespondsToSelector(objc.Sel("relinquishPresentedItemToReader:"))
}

// Notifies your object that another object or process wants to read the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1410743-relinquishpresenteditemtoreader?language=objc
func (f_ FilePresenterWrapper) RelinquishPresentedItemToReader(reader func(arg0 func())) {
	objc.Call[objc.Void](f_, objc.Sel("relinquishPresentedItemToReader:"), reader)
}

func (f_ FilePresenterWrapper) HasPresentedItemDidGainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidGainVersion:"))
}

// Tells the delegate that a new version of the file or file package was added. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415018-presenteditemdidgainversion?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidGainVersion(version IFileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidGainVersion:"), objc.Ptr(version))
}

func (f_ FilePresenterWrapper) HasPresentedItemDidResolveConflictVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidResolveConflictVersion:"))
}

// Tells the delegate that some other entity resolved a version conflict for the presenter’s file or file package. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1418445-presenteditemdidresolveconflictv?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidResolveConflictVersion(version IFileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidResolveConflictVersion:"), objc.Ptr(version))
}

func (f_ FilePresenterWrapper) HasPresentedItemDidChangeUbiquityAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidChangeUbiquityAttributes:"))
}

// Tells your object that the file or file package's ubiquity attributes have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/2909021-presenteditemdidchangeubiquityat?language=objc
func (f_ FilePresenterWrapper) PresentedItemDidChangeUbiquityAttributes(attributes ISet) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), objc.Ptr(attributes))
}

func (f_ FilePresenterWrapper) HasRelinquishPresentedItemToWriter() bool {
	return f_.RespondsToSelector(objc.Sel("relinquishPresentedItemToWriter:"))
}

// Notifies your object that another object or process wants to write to the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1413688-relinquishpresenteditemtowriter?language=objc
func (f_ FilePresenterWrapper) RelinquishPresentedItemToWriter(writer func(arg0 func())) {
	objc.Call[objc.Void](f_, objc.Sel("relinquishPresentedItemToWriter:"), writer)
}

func (f_ FilePresenterWrapper) HasPresentedSubitemDidAppearAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemDidAppearAtURL:"))
}

// Tells the delegate that an item was added to the presented directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1408642-presentedsubitemdidappearaturl?language=objc
func (f_ FilePresenterWrapper) PresentedSubitemDidAppearAtURL(url IURL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemDidAppearAtURL:"), objc.Ptr(url))
}

func (f_ FilePresenterWrapper) HasPresentedSubitemDidChangeAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemDidChangeAtURL:"))
}

// Tells the delegate that the contents or attributes of the specified item changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1411135-presentedsubitemdidchangeaturl?language=objc
func (f_ FilePresenterWrapper) PresentedSubitemDidChangeAtURL(url IURL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemDidChangeAtURL:"), objc.Ptr(url))
}

func (f_ FilePresenterWrapper) HasAccommodatePresentedSubitemDeletionAtURLCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("accommodatePresentedSubitemDeletionAtURL:completionHandler:"))
}

// Tells the delegate that some entity wants to delete an item that is inside of a presented directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415657-accommodatepresentedsubitemdelet?language=objc
func (f_ FilePresenterWrapper) AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url IURL, completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("accommodatePresentedSubitemDeletionAtURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

func (f_ FilePresenterWrapper) HasPresentedItemOperationQueue() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemOperationQueue"))
}

// The operation queue in which to execute presenter-related messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415250-presenteditemoperationqueue?language=objc
func (f_ FilePresenterWrapper) PresentedItemOperationQueue() OperationQueue {
	rv := objc.Call[OperationQueue](f_, objc.Sel("presentedItemOperationQueue"))
	return rv
}

func (f_ FilePresenterWrapper) HasPresentedItemURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemURL"))
}

// The URL of the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414861-presenteditemurl?language=objc
func (f_ FilePresenterWrapper) PresentedItemURL() URL {
	rv := objc.Call[URL](f_, objc.Sel("presentedItemURL"))
	return rv
}

func (f_ FilePresenterWrapper) HasObservedPresentedItemUbiquityAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("observedPresentedItemUbiquityAttributes"))
}

// A list of ubiquity attributes used to generate and send notifications whenever an attribute in the list changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/2909022-observedpresenteditemubiquityatt?language=objc
func (f_ FilePresenterWrapper) ObservedPresentedItemUbiquityAttributes() Set {
	rv := objc.Call[Set](f_, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return rv
}

func (f_ FilePresenterWrapper) HasPrimaryPresentedItemURL() bool {
	return f_.RespondsToSelector(objc.Sel("primaryPresentedItemURL"))
}

// The URL of a secondary item’s primary presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415415-primarypresenteditemurl?language=objc
func (f_ FilePresenterWrapper) PrimaryPresentedItemURL() URL {
	rv := objc.Call[URL](f_, objc.Sel("primaryPresentedItemURL"))
	return rv
}
