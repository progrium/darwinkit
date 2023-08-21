// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileCoordinator] class.
var FileCoordinatorClass = _FileCoordinatorClass{objc.GetClass("NSFileCoordinator")}

type _FileCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [FileCoordinator] class.
type IFileCoordinator interface {
	objc.IObject
	PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []IURL, readingOptions FileCoordinatorReadingOptions, writingURLs []IURL, writingOptions FileCoordinatorWritingOptions, outError IError, batchAccessor func(arg0 func()))
	CoordinateAccessWithIntentsQueueByAccessor(intents []IFileAccessIntent, queue IOperationQueue, accessor func(error Error))
	CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 IURL, options1 FileCoordinatorWritingOptions, url2 IURL, options2 FileCoordinatorWritingOptions, outError IError, writer func(newURL1 URL, newURL2 URL))
	CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL IURL, readingOptions FileCoordinatorReadingOptions, writingURL IURL, writingOptions FileCoordinatorWritingOptions, outError IError, readerWriter func(newReadingURL URL, newWritingURL URL))
	ItemAtURLDidChangeUbiquityAttributes(url IURL, attributes ISet)
	Cancel()
	PurposeIdentifier() string
	SetPurposeIdentifier(value string)
}

// An object that coordinates the reading and writing of files and directories among file presenters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator?language=objc
type FileCoordinator struct {
	objc.Object
}

func FileCoordinatorFrom(ptr unsafe.Pointer) FileCoordinator {
	return FileCoordinator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FileCoordinator) InitWithFilePresenter(filePresenterOrNil PFilePresenter) FileCoordinator {
	po0 := objc.WrapAsProtocol("NSFilePresenter", filePresenterOrNil)
	rv := objc.Call[FileCoordinator](f_, objc.Sel("initWithFilePresenter:"), po0)
	return rv
}

// Initializes and returns a file coordinator object using the specified file presenter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1416795-initwithfilepresenter?language=objc
func NewFileCoordinatorWithFilePresenter(filePresenterOrNil PFilePresenter) FileCoordinator {
	instance := FileCoordinatorClass.Alloc().InitWithFilePresenter(filePresenterOrNil)
	instance.Autorelease()
	return instance
}

func (fc _FileCoordinatorClass) Alloc() FileCoordinator {
	rv := objc.Call[FileCoordinator](fc, objc.Sel("alloc"))
	return rv
}

func FileCoordinator_Alloc() FileCoordinator {
	return FileCoordinatorClass.Alloc()
}

func (fc _FileCoordinatorClass) New() FileCoordinator {
	rv := objc.Call[FileCoordinator](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileCoordinator() FileCoordinator {
	return FileCoordinatorClass.New()
}

func (f_ FileCoordinator) Init() FileCoordinator {
	rv := objc.Call[FileCoordinator](f_, objc.Sel("init"))
	return rv
}

// Prepare to read or write from multiple files in a single batch operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1412420-prepareforreadingitemsaturls?language=objc
func (f_ FileCoordinator) PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []IURL, readingOptions FileCoordinatorReadingOptions, writingURLs []IURL, writingOptions FileCoordinatorWritingOptions, outError IError, batchAccessor func(arg0 func())) {
	objc.Call[objc.Void](f_, objc.Sel("prepareForReadingItemsAtURLs:options:writingItemsAtURLs:options:error:byAccessor:"), readingURLs, readingOptions, writingURLs, writingOptions, objc.Ptr(outError), batchAccessor)
}

// Performs a number of coordinated-read or -write operations asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1411533-coordinateaccesswithintents?language=objc
func (f_ FileCoordinator) CoordinateAccessWithIntentsQueueByAccessor(intents []IFileAccessIntent, queue IOperationQueue, accessor func(error Error)) {
	objc.Call[objc.Void](f_, objc.Sel("coordinateAccessWithIntents:queue:byAccessor:"), intents, objc.Ptr(queue), accessor)
}

// Initiates a write operation that involves a secondary write operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1408970-coordinatewritingitematurl?language=objc
func (f_ FileCoordinator) CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 IURL, options1 FileCoordinatorWritingOptions, url2 IURL, options2 FileCoordinatorWritingOptions, outError IError, writer func(newURL1 URL, newURL2 URL)) {
	objc.Call[objc.Void](f_, objc.Sel("coordinateWritingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), objc.Ptr(url1), options1, objc.Ptr(url2), options2, objc.Ptr(outError), writer)
}

// Initiates a read operation that contains a follow-up write operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1413385-coordinatereadingitematurl?language=objc
func (f_ FileCoordinator) CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL IURL, readingOptions FileCoordinatorReadingOptions, writingURL IURL, writingOptions FileCoordinatorWritingOptions, outError IError, readerWriter func(newReadingURL URL, newWritingURL URL)) {
	objc.Call[objc.Void](f_, objc.Sel("coordinateReadingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), objc.Ptr(readingURL), readingOptions, objc.Ptr(writingURL), writingOptions, objc.Ptr(outError), readerWriter)
}

// Tells observing file providers that the item's ubiquity attributes have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/2909000-itematurl?language=objc
func (f_ FileCoordinator) ItemAtURLDidChangeUbiquityAttributes(url IURL, attributes ISet) {
	objc.Call[objc.Void](f_, objc.Sel("itemAtURL:didChangeUbiquityAttributes:"), objc.Ptr(url), objc.Ptr(attributes))
}

// Unregisters the specified file presenter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1414924-removefilepresenter?language=objc
func (fc _FileCoordinatorClass) RemoveFilePresenter(filePresenter PFilePresenter) {
	po0 := objc.WrapAsProtocol("NSFilePresenter", filePresenter)
	objc.Call[objc.Void](fc, objc.Sel("removeFilePresenter:"), po0)
}

// Unregisters the specified file presenter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1414924-removefilepresenter?language=objc
func FileCoordinator_RemoveFilePresenter(filePresenter PFilePresenter) {
	FileCoordinatorClass.RemoveFilePresenter(filePresenter)
}

// Registers the specified file presenter object so that it can receive notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1417120-addfilepresenter?language=objc
func (fc _FileCoordinatorClass) AddFilePresenter(filePresenter PFilePresenter) {
	po0 := objc.WrapAsProtocol("NSFilePresenter", filePresenter)
	objc.Call[objc.Void](fc, objc.Sel("addFilePresenter:"), po0)
}

// Registers the specified file presenter object so that it can receive notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1417120-addfilepresenter?language=objc
func FileCoordinator_AddFilePresenter(filePresenter PFilePresenter) {
	FileCoordinatorClass.AddFilePresenter(filePresenter)
}

// Cancels any active file coordination calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1418457-cancel?language=objc
func (f_ FileCoordinator) Cancel() {
	objc.Call[objc.Void](f_, objc.Sel("cancel"))
}

// A string that uniquely identifies the file access that was performed by this file coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1411868-purposeidentifier?language=objc
func (f_ FileCoordinator) PurposeIdentifier() string {
	rv := objc.Call[string](f_, objc.Sel("purposeIdentifier"))
	return rv
}

// A string that uniquely identifies the file access that was performed by this file coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1411868-purposeidentifier?language=objc
func (f_ FileCoordinator) SetPurposeIdentifier(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setPurposeIdentifier:"), value)
}

// Returns an array containing the currently registered file presenter objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1407685-filepresenters?language=objc
func (fc _FileCoordinatorClass) FilePresenters() []FilePresenterWrapper {
	rv := objc.Call[[]FilePresenterWrapper](fc, objc.Sel("filePresenters"))
	return rv
}

// Returns an array containing the currently registered file presenter objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/1407685-filepresenters?language=objc
func FileCoordinator_FilePresenters() []FilePresenterWrapper {
	return FileCoordinatorClass.FilePresenters()
}
