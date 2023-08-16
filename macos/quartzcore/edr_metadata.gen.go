// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EDRMetadata] class.
var EDRMetadataClass = _EDRMetadataClass{objc.GetClass("CAEDRMetadata")}

type _EDRMetadataClass struct {
	objc.Class
}

// An interface definition for the [EDRMetadata] class.
type IEDRMetadata interface {
	objc.IObject
}

// Metadata describing how extended dynamic range (EDR) values should be tone mapped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata?language=objc
type EDRMetadata struct {
	objc.Object
}

func EDRMetadataFrom(ptr unsafe.Pointer) EDRMetadata {
	return EDRMetadata{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EDRMetadataClass) Alloc() EDRMetadata {
	rv := objc.Call[EDRMetadata](ec, objc.Sel("alloc"))
	return rv
}

func EDRMetadata_Alloc() EDRMetadata {
	return EDRMetadataClass.Alloc()
}

func (ec _EDRMetadataClass) New() EDRMetadata {
	rv := objc.Call[EDRMetadata](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEDRMetadata() EDRMetadata {
	return EDRMetadataClass.New()
}

func (e_ EDRMetadata) Init() EDRMetadata {
	rv := objc.Call[EDRMetadata](e_, objc.Sel("init"))
	return rv
}

// Creates EDR metadata for HDR10 content based on the luminance characteristics of a mastering display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194383-hdr10metadatawithminluminance?language=objc
func (ec _EDRMetadataClass) HDR10MetadataWithMinLuminanceMaxLuminanceOpticalOutputScale(minNits float64, maxNits float64, scale float64) EDRMetadata {
	rv := objc.Call[EDRMetadata](ec, objc.Sel("HDR10MetadataWithMinLuminance:maxLuminance:opticalOutputScale:"), minNits, maxNits, scale)
	return rv
}

// Creates EDR metadata for HDR10 content based on the luminance characteristics of a mastering display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194383-hdr10metadatawithminluminance?language=objc
func EDRMetadata_HDR10MetadataWithMinLuminanceMaxLuminanceOpticalOutputScale(minNits float64, maxNits float64, scale float64) EDRMetadata {
	return EDRMetadataClass.HDR10MetadataWithMinLuminanceMaxLuminanceOpticalOutputScale(minNits, maxNits, scale)
}

// Creates EDR metadata for HDR10 content based on mastering display color information and content light levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194382-hdr10metadatawithdisplayinfo?language=objc
func (ec _EDRMetadataClass) HDR10MetadataWithDisplayInfoContentInfoOpticalOutputScale(displayData []byte, contentData []byte, scale float64) EDRMetadata {
	rv := objc.Call[EDRMetadata](ec, objc.Sel("HDR10MetadataWithDisplayInfo:contentInfo:opticalOutputScale:"), displayData, contentData, scale)
	return rv
}

// Creates EDR metadata for HDR10 content based on mastering display color information and content light levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194382-hdr10metadatawithdisplayinfo?language=objc
func EDRMetadata_HDR10MetadataWithDisplayInfoContentInfoOpticalOutputScale(displayData []byte, contentData []byte, scale float64) EDRMetadata {
	return EDRMetadataClass.HDR10MetadataWithDisplayInfoContentInfoOpticalOutputScale(displayData, contentData, scale)
}

// Extended dynamic range (EDR) metadata for the Hybrid Log-Gamma (HLG) transfer function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194384-hlgmetadata?language=objc
func (ec _EDRMetadataClass) HLGMetadata() EDRMetadata {
	rv := objc.Call[EDRMetadata](ec, objc.Sel("HLGMetadata"))
	return rv
}

// Extended dynamic range (EDR) metadata for the Hybrid Log-Gamma (HLG) transfer function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedrmetadata/3194384-hlgmetadata?language=objc
func EDRMetadata_HLGMetadata() EDRMetadata {
	return EDRMetadataClass.HLGMetadata()
}
