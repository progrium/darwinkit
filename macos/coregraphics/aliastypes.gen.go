// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/iosurface"
	"github.com/progrium/macdriver/objc"
)

// A callback function that returns a generic pointer to the provider data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovidergetbytepointercallback?language=objc
type DataProviderGetBytePointerCallback = func(info unsafe.Pointer) unsafe.Pointer

// Performs custom tasks at the end of a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterenddocumentcallback?language=objc
type PSConverterEndDocumentCallback = func(info unsafe.Pointer, success bool)

// A client-supplied callback function that’s invoked when an area of the display is modified or refreshed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenrefreshcallback?language=objc
type ScreenRefreshCallback = func(count uint32, rects *Rect, userInfo unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionaryapplierblock?language=objc
type PDFDictionaryApplierBlock = func(key *uint8, value unsafe.Pointer, info unsafe.Pointer) bool

// Performs custom tasks at the beginning of a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterbegindocumentcallback?language=objc
type PSConverterBeginDocumentCallback = func(info unsafe.Pointer)

// Performs custom processing for PDF operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfoperatorcallback?language=objc
type PDFOperatorCallback = func(scanner unsafe.Pointer, info unsafe.Pointer)

// A callback function that releases the pointer Core Graphics obtained by calling CGDataProviderGetBytePointerCallback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderreleasebytepointercallback?language=objc
type DataProviderReleaseBytePointerCallback = func(info unsafe.Pointer, pointer unsafe.Pointer)

// Draws a pattern cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatterndrawpatterncallback?language=objc
type PatternDrawPatternCallback = func(info unsafe.Pointer, context ContextRef)

// Passes messages generated during a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconvertermessagecallback?language=objc
type PSConverterMessageCallback = func(info unsafe.Pointer, message corefoundation.StringRef)

// A client-supplied callback function that’s invoked whenever the configuration of a local display is changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayreconfigurationcallback?language=objc
type DisplayReconfigurationCallBack = func(display DirectDisplayID, flags DisplayChangeSummaryFlags, userInfo unsafe.Pointer)

// Reports progress periodically during a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterprogresscallback?language=objc
type PSConverterProgressCallback = func(info unsafe.Pointer)

// A client-supplied callback function that’s invoked whenever an associated event tap receives a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtapcallback?language=objc
type EventTapCallBack = func(proxy unsafe.Pointer, type_ EventType, event EventRef, userInfo unsafe.Pointer) EventRef

// Copies data from a Core Graphics-supplied buffer into a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumerputbytescallback?language=objc
type DataConsumerPutBytesCallback = func(info unsafe.Pointer, buffer unsafe.Pointer, count uint) uint

// Defines a callback function that can view an element in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathapplierfunction?language=objc
type PathApplierFunction = func(info unsafe.Pointer, element *PathElement)

// Performs custom processing on a key-value pair from a PDF dictionary, using optional contextual information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionaryapplierfunction?language=objc
type PDFDictionaryApplierFunction = func(key *uint8, value unsafe.Pointer, info unsafe.Pointer)

// Release private data or resources associated with the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatternreleaseinfocallback?language=objc
type PatternReleaseInfoCallback = func(info unsafe.Pointer)

// Releases any private data or resources associated with the data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumerreleaseinfocallback?language=objc
type DataConsumerReleaseInfoCallback = func(info unsafe.Pointer)

// A client-supplied callback function invoked when an area of the display is moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenupdatemovecallback?language=objc
type ScreenUpdateMoveCallback = func(delta ScreenUpdateMoveDelta, count uint, rects *Rect, userInfo unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathapplyblock?language=objc
type PathApplyBlock = func(element *PathElement)

// Performs custom tasks when a PostScript converter is released. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterreleaseinfocallback?language=objc
type PSConverterReleaseInfoCallback = func(info unsafe.Pointer)

// Performs custom tasks at the beginning of each page in a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterbeginpagecallback?language=objc
type PSConverterBeginPageCallback = func(info unsafe.Pointer, pageNumber uint, pageInfo corefoundation.DictionaryRef)

// Performs custom clean-up tasks when Core Graphics deallocates a CGFunctionRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctionreleaseinfocallback?language=objc
type FunctionReleaseInfoCallback = func(info unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarrayapplierblock?language=objc
type PDFArrayApplierBlock = func(index uint, value unsafe.Pointer, info unsafe.Pointer) bool

// A callback function that releases data you supply to the function CGDataProviderCreateWithData. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderreleasedatacallback?language=objc
type DataProviderReleaseDataCallback = func(info unsafe.Pointer, data unsafe.Pointer, size uint)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgerrorcallback?language=objc
type ErrorCallback = func()

// Performs custom tasks at the end of each page of a PostScript conversion process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterendpagecallback?language=objc
type PSConverterEndPageCallback = func(info unsafe.Pointer, pageNumber uint, pageInfo corefoundation.DictionaryRef)

// A block called when a data stream has a new frame event to process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamframeavailablehandler?language=objc
type DisplayStreamFrameAvailableHandler = func(status DisplayStreamFrameStatus, displayTime uint64, frameSurface iosurface.Ref, updateRef DisplayStreamUpdateRef)

// Performs custom operations on the supplied input data to produce output data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctionevaluatecallback?language=objc
type FunctionEvaluateCallback = func(info unsafe.Pointer, arg1 *float64, arg2 *float64)

// A callback function that moves the current position in the data stream back to the beginning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderrewindcallback?language=objc
type DataProviderRewindCallback = func(info unsafe.Pointer)

// A callback function used to release data associate with the bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgbitmapcontextreleasedatacallback?language=objc
type BitmapContextReleaseDataCallback = func(releaseInfo unsafe.Pointer, data unsafe.Pointer)

// A callback function that copies from a provider data stream into a Core Graphics buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovidergetbytescallback?language=objc
type DataProviderGetBytesCallback = func(info unsafe.Pointer, buffer unsafe.Pointer, count uint) uint

// A callback function that releases any private data or resources associated with the data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderreleaseinfocallback?language=objc
type DataProviderReleaseInfoCallback = func(info unsafe.Pointer)

// A callback function that advances the current position in the data stream supplied by the provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderskipforwardcallback?language=objc
type DataProviderSkipForwardCallback = func(info unsafe.Pointer, count objc.Object) objc.Object
