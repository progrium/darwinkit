// AUTO-GENERATED CODE, DO NOT MODIFY

package corevideo

import (
	"unsafe"
)

// A type for a display link callback function that the system invokes when it’s time for the app to output a video frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvdisplaylinkoutputcallback?language=objc
type DisplayLinkOutputCallback = func(displayLink DisplayLinkRef, inNow *TimeStamp, inOutputTime *TimeStamp, flagsIn OptionFlags, flagsOut *OptionFlags, displayLinkContext unsafe.Pointer) Return

// Defines a pointer to a pixel buffer release callback function, which is called when a pixel buffer created by CVPixelBufferCreateWithPlanarBytes is released. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvpixelbufferreleaseplanarbytescallback?language=objc
type PixelBufferReleasePlanarBytesCallback = func(releaseRefCon unsafe.Pointer, dataPtr unsafe.Pointer, dataSize uint, numberOfPlanes uint, planeAddresses unsafe.Pointer)

// A type that defines a release callback function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvpixelbufferreleasebytescallback?language=objc
type PixelBufferReleaseBytesCallback = func(releaseRefCon unsafe.Pointer, baseAddress unsafe.Pointer)

// Defines a pointer to a custom extended pixel-fill function, which is called whenever the system needs to pad a buffer holding your custom pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvfillextendedpixelscallback?language=objc
type FillExtendedPixelsCallBack = func(pixelBuffer PixelBufferRef, refCon unsafe.Pointer) bool

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvdisplaylinkoutputhandler?language=objc
type DisplayLinkOutputHandler = func(displayLink DisplayLinkRef, inNow *TimeStamp, inOutputTime *TimeStamp, flagsIn OptionFlags, flagsOut *OptionFlags) Return
