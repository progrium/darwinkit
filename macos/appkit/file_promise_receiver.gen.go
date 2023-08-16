// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilePromiseReceiver] class.
var FilePromiseReceiverClass = _FilePromiseReceiverClass{objc.GetClass("NSFilePromiseReceiver")}

type _FilePromiseReceiverClass struct {
	objc.Class
}

// An interface definition for the [FilePromiseReceiver] class.
type IFilePromiseReceiver interface {
	objc.IObject
	ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir foundation.IURL, options foundation.Dictionary, operationQueue foundation.IOperationQueue, reader func(fileURL foundation.URL, errorOrNil foundation.Error))
	FileTypes() []string
	FileNames() []string
}

// An object that receives a file promise from the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver?language=objc
type FilePromiseReceiver struct {
	objc.Object
}

func FilePromiseReceiverFrom(ptr unsafe.Pointer) FilePromiseReceiver {
	return FilePromiseReceiver{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FilePromiseReceiverClass) Alloc() FilePromiseReceiver {
	rv := objc.Call[FilePromiseReceiver](fc, objc.Sel("alloc"))
	return rv
}

func FilePromiseReceiver_Alloc() FilePromiseReceiver {
	return FilePromiseReceiverClass.Alloc()
}

func (fc _FilePromiseReceiverClass) New() FilePromiseReceiver {
	rv := objc.Call[FilePromiseReceiver](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilePromiseReceiver() FilePromiseReceiver {
	return FilePromiseReceiverClass.New()
}

func (f_ FilePromiseReceiver) Init() FilePromiseReceiver {
	rv := objc.Call[FilePromiseReceiver](f_, objc.Sel("init"))
	return rv
}

// Fulfills the promises at the specified destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/1642138-receivepromisedfilesatdestinatio?language=objc
func (f_ FilePromiseReceiver) ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir foundation.IURL, options foundation.Dictionary, operationQueue foundation.IOperationQueue, reader func(fileURL foundation.URL, errorOrNil foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("receivePromisedFilesAtDestination:options:operationQueue:reader:"), objc.Ptr(destinationDir), options, objc.Ptr(operationQueue), reader)
}

// An array containing types of the promised files being written to the destination location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/1642141-filetypes?language=objc
func (f_ FilePromiseReceiver) FileTypes() []string {
	rv := objc.Call[[]string](f_, objc.Sel("fileTypes"))
	return rv
}

// An array containing names of the promised files being written to the destination location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/1642142-filenames?language=objc
func (f_ FilePromiseReceiver) FileNames() []string {
	rv := objc.Call[[]string](f_, objc.Sel("fileNames"))
	return rv
}

// An array containing dragged file types that are readable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/1642140-readabledraggedtypes?language=objc
func (fc _FilePromiseReceiverClass) ReadableDraggedTypes() []string {
	rv := objc.Call[[]string](fc, objc.Sel("readableDraggedTypes"))
	return rv
}

// An array containing dragged file types that are readable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/1642140-readabledraggedtypes?language=objc
func FilePromiseReceiver_ReadableDraggedTypes() []string {
	return FilePromiseReceiverClass.ReadableDraggedTypes()
}
