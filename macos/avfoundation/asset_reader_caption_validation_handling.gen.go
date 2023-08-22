// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods for caption validation events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadercaptionvalidationhandling?language=objc
type PAssetReaderCaptionValidationHandling interface {
	// optional
	CaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements(adaptor AssetReaderOutputCaptionAdaptor, caption Caption, syntaxElements []string)
	HasCaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements() bool
}

// A concrete type wrapper for the [PAssetReaderCaptionValidationHandling] protocol.
type AssetReaderCaptionValidationHandlingWrapper struct {
	objc.Object
}

func (a_ AssetReaderCaptionValidationHandlingWrapper) HasCaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements() bool {
	return a_.RespondsToSelector(objc.Sel("captionAdaptor:didVendCaption:skippingUnsupportedSourceSyntaxElements:"))
}

// Tells the delegate that the adaptor ignored one or more syntax elements when it created the caption object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadercaptionvalidationhandling/3752792-captionadaptor?language=objc
func (a_ AssetReaderCaptionValidationHandlingWrapper) CaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements(adaptor IAssetReaderOutputCaptionAdaptor, caption ICaption, syntaxElements []string) {
	objc.Call[objc.Void](a_, objc.Sel("captionAdaptor:didVendCaption:skippingUnsupportedSourceSyntaxElements:"), objc.Ptr(adaptor), objc.Ptr(caption), syntaxElements)
}
