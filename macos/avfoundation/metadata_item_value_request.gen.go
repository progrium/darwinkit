// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataItemValueRequest] class.
var MetadataItemValueRequestClass = _MetadataItemValueRequestClass{objc.GetClass("AVMetadataItemValueRequest")}

type _MetadataItemValueRequestClass struct {
	objc.Class
}

// An interface definition for the [MetadataItemValueRequest] class.
type IMetadataItemValueRequest interface {
	objc.IObject
	RespondWithValue(value objc.IObject)
	RespondWithError(error foundation.IError)
	MetadataItem() MetadataItem
}

// An object that responds to a request to load the value of a metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemvaluerequest?language=objc
type MetadataItemValueRequest struct {
	objc.Object
}

func MetadataItemValueRequestFrom(ptr unsafe.Pointer) MetadataItemValueRequest {
	return MetadataItemValueRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataItemValueRequestClass) Alloc() MetadataItemValueRequest {
	rv := objc.Call[MetadataItemValueRequest](mc, objc.Sel("alloc"))
	return rv
}

func MetadataItemValueRequest_Alloc() MetadataItemValueRequest {
	return MetadataItemValueRequestClass.Alloc()
}

func (mc _MetadataItemValueRequestClass) New() MetadataItemValueRequest {
	rv := objc.Call[MetadataItemValueRequest](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataItemValueRequest() MetadataItemValueRequest {
	return MetadataItemValueRequestClass.New()
}

func (m_ MetadataItemValueRequest) Init() MetadataItemValueRequest {
	rv := objc.Call[MetadataItemValueRequest](m_, objc.Sel("init"))
	return rv
}

// Returns the metadata item’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemvaluerequest/1386820-respondwithvalue?language=objc
func (m_ MetadataItemValueRequest) RespondWithValue(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("respondWithValue:"), value)
}

// Returns an error when the system fails to load the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemvaluerequest/1390783-respondwitherror?language=objc
func (m_ MetadataItemValueRequest) RespondWithError(error foundation.IError) {
	objc.Call[objc.Void](m_, objc.Sel("respondWithError:"), objc.Ptr(error))
}

// The metadata item to request a value for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemvaluerequest/1388069-metadataitem?language=objc
func (m_ MetadataItemValueRequest) MetadataItem() MetadataItem {
	rv := objc.Call[MetadataItem](m_, objc.Sel("metadataItem"))
	return rv
}
