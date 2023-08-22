// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleBufferGenerator] class.
var SampleBufferGeneratorClass = _SampleBufferGeneratorClass{objc.GetClass("AVSampleBufferGenerator")}

type _SampleBufferGeneratorClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferGenerator] class.
type ISampleBufferGenerator interface {
	objc.IObject
}

// An object that creates sample buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffergenerator?language=objc
type SampleBufferGenerator struct {
	objc.Object
}

func SampleBufferGeneratorFrom(ptr unsafe.Pointer) SampleBufferGenerator {
	return SampleBufferGenerator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SampleBufferGenerator) InitWithAssetTimebase(asset IAsset, timebase coremedia.TimebaseRef) SampleBufferGenerator {
	rv := objc.Call[SampleBufferGenerator](s_, objc.Sel("initWithAsset:timebase:"), objc.Ptr(asset), timebase)
	return rv
}

// Creates a new sample buffer generator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffergenerator/1387477-initwithasset?language=objc
func NewSampleBufferGeneratorWithAssetTimebase(asset IAsset, timebase coremedia.TimebaseRef) SampleBufferGenerator {
	instance := SampleBufferGeneratorClass.Alloc().InitWithAssetTimebase(asset, timebase)
	instance.Autorelease()
	return instance
}

func (sc _SampleBufferGeneratorClass) Alloc() SampleBufferGenerator {
	rv := objc.Call[SampleBufferGenerator](sc, objc.Sel("alloc"))
	return rv
}

func SampleBufferGenerator_Alloc() SampleBufferGenerator {
	return SampleBufferGeneratorClass.Alloc()
}

func (sc _SampleBufferGeneratorClass) New() SampleBufferGenerator {
	rv := objc.Call[SampleBufferGenerator](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferGenerator() SampleBufferGenerator {
	return SampleBufferGeneratorClass.New()
}

func (s_ SampleBufferGenerator) Init() SampleBufferGenerator {
	rv := objc.Call[SampleBufferGenerator](s_, objc.Sel("init"))
	return rv
}

// Notifies the sample buffer generator when data is ready for the sample buffer reference or an error has occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffergenerator/1387295-notifyofdatareadyforsamplebuffer?language=objc
func (sc _SampleBufferGeneratorClass) NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf coremedia.SampleBufferRef, completionHandler func(dataReady bool, error foundation.Error)) {
	objc.Call[objc.Void](sc, objc.Sel("notifyOfDataReadyForSampleBuffer:completionHandler:"), sbuf, completionHandler)
}

// Notifies the sample buffer generator when data is ready for the sample buffer reference or an error has occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffergenerator/1387295-notifyofdatareadyforsamplebuffer?language=objc
func SampleBufferGenerator_NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf coremedia.SampleBufferRef, completionHandler func(dataReady bool, error foundation.Error)) {
	SampleBufferGeneratorClass.NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf, completionHandler)
}
