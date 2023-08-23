// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInputCaptionAdaptor] class.
var AssetWriterInputCaptionAdaptorClass = _AssetWriterInputCaptionAdaptorClass{objc.GetClass("AVAssetWriterInputCaptionAdaptor")}

type _AssetWriterInputCaptionAdaptorClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInputCaptionAdaptor] class.
type IAssetWriterInputCaptionAdaptor interface {
	objc.IObject
	AppendCaptionGroup(captionGroup ICaptionGroup) bool
	AppendCaption(caption ICaption) bool
	AssetWriterInput() AssetWriterInput
}

// An object that appends captions to an asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor?language=objc
type AssetWriterInputCaptionAdaptor struct {
	objc.Object
}

func AssetWriterInputCaptionAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputCaptionAdaptor {
	return AssetWriterInputCaptionAdaptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetWriterInputCaptionAdaptor) InitWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputCaptionAdaptor {
	rv := objc.Call[AssetWriterInputCaptionAdaptor](a_, objc.Sel("initWithAssetWriterInput:"), objc.Ptr(input))
	return rv
}

// Creates a new caption adaptor that writes to the specified asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor/3752805-initwithassetwriterinput?language=objc
func NewAssetWriterInputCaptionAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputCaptionAdaptor {
	instance := AssetWriterInputCaptionAdaptorClass.Alloc().InitWithAssetWriterInput(input)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterInputCaptionAdaptorClass) AssetWriterInputCaptionAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputCaptionAdaptor {
	rv := objc.Call[AssetWriterInputCaptionAdaptor](ac, objc.Sel("assetWriterInputCaptionAdaptorWithAssetWriterInput:"), objc.Ptr(input))
	return rv
}

// A class method that creates a new caption adaptor that writes to the specified asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor/3752804-assetwriterinputcaptionadaptorwi?language=objc
func AssetWriterInputCaptionAdaptor_AssetWriterInputCaptionAdaptorWithAssetWriterInput(input IAssetWriterInput) AssetWriterInputCaptionAdaptor {
	return AssetWriterInputCaptionAdaptorClass.AssetWriterInputCaptionAdaptorWithAssetWriterInput(input)
}

func (ac _AssetWriterInputCaptionAdaptorClass) Alloc() AssetWriterInputCaptionAdaptor {
	rv := objc.Call[AssetWriterInputCaptionAdaptor](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInputCaptionAdaptor_Alloc() AssetWriterInputCaptionAdaptor {
	return AssetWriterInputCaptionAdaptorClass.Alloc()
}

func (ac _AssetWriterInputCaptionAdaptorClass) New() AssetWriterInputCaptionAdaptor {
	rv := objc.Call[AssetWriterInputCaptionAdaptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInputCaptionAdaptor() AssetWriterInputCaptionAdaptor {
	return AssetWriterInputCaptionAdaptorClass.New()
}

func (a_ AssetWriterInputCaptionAdaptor) Init() AssetWriterInputCaptionAdaptor {
	rv := objc.Call[AssetWriterInputCaptionAdaptor](a_, objc.Sel("init"))
	return rv
}

// Appends a caption group that the system writes to the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor/3752802-appendcaptiongroup?language=objc
func (a_ AssetWriterInputCaptionAdaptor) AppendCaptionGroup(captionGroup ICaptionGroup) bool {
	rv := objc.Call[bool](a_, objc.Sel("appendCaptionGroup:"), objc.Ptr(captionGroup))
	return rv
}

// Appends a caption to the writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor/3752801-appendcaption?language=objc
func (a_ AssetWriterInputCaptionAdaptor) AppendCaption(caption ICaption) bool {
	rv := objc.Call[bool](a_, objc.Sel("appendCaption:"), objc.Ptr(caption))
	return rv
}

// The associated asset writer input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputcaptionadaptor/3752803-assetwriterinput?language=objc
func (a_ AssetWriterInputCaptionAdaptor) AssetWriterInput() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("assetWriterInput"))
	return rv
}
