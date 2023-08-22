// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInputMetadataAdaptor] class.
var AssetWriterInputMetadataAdaptorClass = _AssetWriterInputMetadataAdaptorClass{objc.GetClass("AVAssetWriterInputMetadataAdaptor")}

type _AssetWriterInputMetadataAdaptorClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInputMetadataAdaptor] class.
type IAssetWriterInputMetadataAdaptor interface {
	objc.IObject
	AppendTimedMetadataGroup(timedMetadataGroup ITimedMetadataGroup) bool
	AssetWriterInput() AssetWriterInput
}

// An object that appends timed metadata groups to an asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor?language=objc
type AssetWriterInputMetadataAdaptor struct {
	objc.Object
}

func AssetWriterInputMetadataAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetWriterInputMetadataAdaptor) InitWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputMetadataAdaptor {
	rv := objc.Call[AssetWriterInputMetadataAdaptor](a_, objc.Sel("initWithAssetWriterInput:"), objc.Ptr(input))
	return rv
}

// Creates a metadata group adaptor to append timed metadata groups to write to an output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/1389706-initwithassetwriterinput?language=objc
func NewAssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputMetadataAdaptor {
	instance := AssetWriterInputMetadataAdaptorClass.Alloc().InitWithAssetWriterInput(input)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterInputMetadataAdaptorClass) AssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputMetadataAdaptor {
	rv := objc.Call[AssetWriterInputMetadataAdaptor](ac, objc.Sel("assetWriterInputMetadataAdaptorWithAssetWriterInput:"), objc.Ptr(input))
	return rv
}

// Returns a new metadata adaptor to append timed metadata groups to write to an output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/1449094-assetwriterinputmetadataadaptorw?language=objc
func AssetWriterInputMetadataAdaptor_AssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptorClass.AssetWriterInputMetadataAdaptorWithAssetWriterInput(input)
}

func (ac _AssetWriterInputMetadataAdaptorClass) Alloc() AssetWriterInputMetadataAdaptor {
	rv := objc.Call[AssetWriterInputMetadataAdaptor](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInputMetadataAdaptor_Alloc() AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptorClass.Alloc()
}

func (ac _AssetWriterInputMetadataAdaptorClass) New() AssetWriterInputMetadataAdaptor {
	rv := objc.Call[AssetWriterInputMetadataAdaptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInputMetadataAdaptor() AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptorClass.New()
}

func (a_ AssetWriterInputMetadataAdaptor) Init() AssetWriterInputMetadataAdaptor {
	rv := objc.Call[AssetWriterInputMetadataAdaptor](a_, objc.Sel("init"))
	return rv
}

// Appends a timed metadata group to the adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/1389014-appendtimedmetadatagroup?language=objc
func (a_ AssetWriterInputMetadataAdaptor) AppendTimedMetadataGroup(timedMetadataGroup ITimedMetadataGroup) bool {
	rv := objc.Call[bool](a_, objc.Sel("appendTimedMetadataGroup:"), objc.Ptr(timedMetadataGroup))
	return rv
}

// The input for the metadata adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/1386633-assetwriterinput?language=objc
func (a_ AssetWriterInputMetadataAdaptor) AssetWriterInput() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("assetWriterInput"))
	return rv
}
