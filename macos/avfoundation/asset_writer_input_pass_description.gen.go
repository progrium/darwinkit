// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInputPassDescription] class.
var AssetWriterInputPassDescriptionClass = _AssetWriterInputPassDescriptionClass{objc.GetClass("AVAssetWriterInputPassDescription")}

type _AssetWriterInputPassDescriptionClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInputPassDescription] class.
type IAssetWriterInputPassDescription interface {
	objc.IObject
	SourceTimeRanges() []foundation.Value
}

// An object that defines the interface to query for the requirements of the current pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpassdescription?language=objc
type AssetWriterInputPassDescription struct {
	objc.Object
}

func AssetWriterInputPassDescriptionFrom(ptr unsafe.Pointer) AssetWriterInputPassDescription {
	return AssetWriterInputPassDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetWriterInputPassDescriptionClass) Alloc() AssetWriterInputPassDescription {
	rv := objc.Call[AssetWriterInputPassDescription](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInputPassDescription_Alloc() AssetWriterInputPassDescription {
	return AssetWriterInputPassDescriptionClass.Alloc()
}

func (ac _AssetWriterInputPassDescriptionClass) New() AssetWriterInputPassDescription {
	rv := objc.Call[AssetWriterInputPassDescription](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInputPassDescription() AssetWriterInputPassDescription {
	return AssetWriterInputPassDescriptionClass.New()
}

func (a_ AssetWriterInputPassDescription) Init() AssetWriterInputPassDescription {
	rv := objc.Call[AssetWriterInputPassDescription](a_, objc.Sel("init"))
	return rv
}

// An array of time ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpassdescription/1388732-sourcetimeranges?language=objc
func (a_ AssetWriterInputPassDescription) SourceTimeRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](a_, objc.Sel("sourceTimeRanges"))
	return rv
}
