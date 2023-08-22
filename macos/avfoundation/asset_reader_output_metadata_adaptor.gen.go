// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderOutputMetadataAdaptor] class.
var AssetReaderOutputMetadataAdaptorClass = _AssetReaderOutputMetadataAdaptorClass{objc.GetClass("AVAssetReaderOutputMetadataAdaptor")}

type _AssetReaderOutputMetadataAdaptorClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderOutputMetadataAdaptor] class.
type IAssetReaderOutputMetadataAdaptor interface {
	objc.IObject
	NextTimedMetadataGroup() TimedMetadataGroup
	AssetReaderTrackOutput() AssetReaderTrackOutput
}

// An object that creates timed metadata group objects for an asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputmetadataadaptor?language=objc
type AssetReaderOutputMetadataAdaptor struct {
	objc.Object
}

func AssetReaderOutputMetadataAdaptorFrom(ptr unsafe.Pointer) AssetReaderOutputMetadataAdaptor {
	return AssetReaderOutputMetadataAdaptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetReaderOutputMetadataAdaptor) InitWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputMetadataAdaptor {
	rv := objc.Call[AssetReaderOutputMetadataAdaptor](a_, objc.Sel("initWithAssetReaderTrackOutput:"), objc.Ptr(trackOutput))
	return rv
}

// Creates an object that reads timed metadata groups from an asset reader output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputmetadataadaptor/1388009-initwithassetreadertrackoutput?language=objc
func NewAssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputMetadataAdaptor {
	instance := AssetReaderOutputMetadataAdaptorClass.Alloc().InitWithAssetReaderTrackOutput(trackOutput)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderOutputMetadataAdaptorClass) AssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputMetadataAdaptor {
	rv := objc.Call[AssetReaderOutputMetadataAdaptor](ac, objc.Sel("assetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput:"), objc.Ptr(trackOutput))
	return rv
}

// Returns a new object that reads timed metadata groups from an asset reader output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputmetadataadaptor/1490330-assetreaderoutputmetadataadaptor?language=objc
func AssetReaderOutputMetadataAdaptor_AssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputMetadataAdaptor {
	return AssetReaderOutputMetadataAdaptorClass.AssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput)
}

func (ac _AssetReaderOutputMetadataAdaptorClass) Alloc() AssetReaderOutputMetadataAdaptor {
	rv := objc.Call[AssetReaderOutputMetadataAdaptor](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderOutputMetadataAdaptor_Alloc() AssetReaderOutputMetadataAdaptor {
	return AssetReaderOutputMetadataAdaptorClass.Alloc()
}

func (ac _AssetReaderOutputMetadataAdaptorClass) New() AssetReaderOutputMetadataAdaptor {
	rv := objc.Call[AssetReaderOutputMetadataAdaptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderOutputMetadataAdaptor() AssetReaderOutputMetadataAdaptor {
	return AssetReaderOutputMetadataAdaptorClass.New()
}

func (a_ AssetReaderOutputMetadataAdaptor) Init() AssetReaderOutputMetadataAdaptor {
	rv := objc.Call[AssetReaderOutputMetadataAdaptor](a_, objc.Sel("init"))
	return rv
}

// Returns the next timed metadata group for the asset reader output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputmetadataadaptor/1390008-nexttimedmetadatagroup?language=objc
func (a_ AssetReaderOutputMetadataAdaptor) NextTimedMetadataGroup() TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](a_, objc.Sel("nextTimedMetadataGroup"))
	return rv
}

// The asset reader track output that provides the timed metadata groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputmetadataadaptor/1388330-assetreadertrackoutput?language=objc
func (a_ AssetReaderOutputMetadataAdaptor) AssetReaderTrackOutput() AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](a_, objc.Sel("assetReaderTrackOutput"))
	return rv
}
