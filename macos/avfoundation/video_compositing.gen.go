// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods custom video compositors must implement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing?language=objc
type PVideoCompositing interface {
	// optional
	RenderContextChanged(newRenderContext VideoCompositionRenderContext)
	HasRenderContextChanged() bool

	// optional
	AnticipateRenderingUsingHint(renderHint VideoCompositionRenderHint)
	HasAnticipateRenderingUsingHint() bool

	// optional
	PrerollForRenderingUsingHint(renderHint VideoCompositionRenderHint)
	HasPrerollForRenderingUsingHint() bool

	// optional
	CancelAllPendingVideoCompositionRequests()
	HasCancelAllPendingVideoCompositionRequests() bool

	// optional
	StartVideoCompositionRequest(asyncVideoCompositionRequest AsynchronousVideoCompositionRequest)
	HasStartVideoCompositionRequest() bool

	// optional
	SupportsHDRSourceFrames() bool
	HasSupportsHDRSourceFrames() bool

	// optional
	SupportsWideColorSourceFrames() bool
	HasSupportsWideColorSourceFrames() bool

	// optional
	RequiredPixelBufferAttributesForRenderContext() map[string]objc.IObject
	HasRequiredPixelBufferAttributesForRenderContext() bool

	// optional
	CanConformColorOfSourceFrames() bool
	HasCanConformColorOfSourceFrames() bool

	// optional
	SourcePixelBufferAttributes() map[string]objc.IObject
	HasSourcePixelBufferAttributes() bool
}

// A concrete type wrapper for the [PVideoCompositing] protocol.
type VideoCompositingWrapper struct {
	objc.Object
}

func (v_ VideoCompositingWrapper) HasRenderContextChanged() bool {
	return v_.RespondsToSelector(objc.Sel("renderContextChanged:"))
}

// Tells the compositor that the composition changed render contexts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1390363-rendercontextchanged?language=objc
func (v_ VideoCompositingWrapper) RenderContextChanged(newRenderContext IVideoCompositionRenderContext) {
	objc.Call[objc.Void](v_, objc.Sel("renderContextChanged:"), objc.Ptr(newRenderContext))
}

func (v_ VideoCompositingWrapper) HasAnticipateRenderingUsingHint() bool {
	return v_.RespondsToSelector(objc.Sel("anticipateRenderingUsingHint:"))
}

// Informs a custom video compositor about upcoming rendering requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/3227885-anticipaterenderingusinghint?language=objc
func (v_ VideoCompositingWrapper) AnticipateRenderingUsingHint(renderHint IVideoCompositionRenderHint) {
	objc.Call[objc.Void](v_, objc.Sel("anticipateRenderingUsingHint:"), objc.Ptr(renderHint))
}

func (v_ VideoCompositingWrapper) HasPrerollForRenderingUsingHint() bool {
	return v_.RespondsToSelector(objc.Sel("prerollForRenderingUsingHint:"))
}

// Tells a custom video compositor to perform any work in the prerolling phase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/3227886-prerollforrenderingusinghint?language=objc
func (v_ VideoCompositingWrapper) PrerollForRenderingUsingHint(renderHint IVideoCompositionRenderHint) {
	objc.Call[objc.Void](v_, objc.Sel("prerollForRenderingUsingHint:"), objc.Ptr(renderHint))
}

func (v_ VideoCompositingWrapper) HasCancelAllPendingVideoCompositionRequests() bool {
	return v_.RespondsToSelector(objc.Sel("cancelAllPendingVideoCompositionRequests"))
}

// Directs a custom video compositor object to cancel or finish all pending video composition requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1390659-cancelallpendingvideocomposition?language=objc
func (v_ VideoCompositingWrapper) CancelAllPendingVideoCompositionRequests() {
	objc.Call[objc.Void](v_, objc.Sel("cancelAllPendingVideoCompositionRequests"))
}

func (v_ VideoCompositingWrapper) HasStartVideoCompositionRequest() bool {
	return v_.RespondsToSelector(objc.Sel("startVideoCompositionRequest:"))
}

// Directs a custom video compositor object to create a new pixel buffer composed asynchronously from a collection of sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1388894-startvideocompositionrequest?language=objc
func (v_ VideoCompositingWrapper) StartVideoCompositionRequest(asyncVideoCompositionRequest IAsynchronousVideoCompositionRequest) {
	objc.Call[objc.Void](v_, objc.Sel("startVideoCompositionRequest:"), objc.Ptr(asyncVideoCompositionRequest))
}

func (v_ VideoCompositingWrapper) HasSupportsHDRSourceFrames() bool {
	return v_.RespondsToSelector(objc.Sel("supportsHDRSourceFrames"))
}

// A Boolean value that indicates whether the compositor handles source frames that contain high dynamic range (HDR) properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/3626026-supportshdrsourceframes?language=objc
func (v_ VideoCompositingWrapper) SupportsHDRSourceFrames() bool {
	rv := objc.Call[bool](v_, objc.Sel("supportsHDRSourceFrames"))
	return rv
}

func (v_ VideoCompositingWrapper) HasSupportsWideColorSourceFrames() bool {
	return v_.RespondsToSelector(objc.Sel("supportsWideColorSourceFrames"))
}

// A Boolean value that indicates whether the compositor handles source frames that contains wide color properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1643657-supportswidecolorsourceframes?language=objc
func (v_ VideoCompositingWrapper) SupportsWideColorSourceFrames() bool {
	rv := objc.Call[bool](v_, objc.Sel("supportsWideColorSourceFrames"))
	return rv
}

func (v_ VideoCompositingWrapper) HasRequiredPixelBufferAttributesForRenderContext() bool {
	return v_.RespondsToSelector(objc.Sel("requiredPixelBufferAttributesForRenderContext"))
}

// The pixel buffer attributes that the compositor requires for pixel buffers that it creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1386414-requiredpixelbufferattributesfor?language=objc
func (v_ VideoCompositingWrapper) RequiredPixelBufferAttributesForRenderContext() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](v_, objc.Sel("requiredPixelBufferAttributesForRenderContext"))
	return rv
}

func (v_ VideoCompositingWrapper) HasCanConformColorOfSourceFrames() bool {
	return v_.RespondsToSelector(objc.Sel("canConformColorOfSourceFrames"))
}

// A Boolean value that indicates whether the compositor conforms the color space of source frames to the composition color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/3750314-canconformcolorofsourceframes?language=objc
func (v_ VideoCompositingWrapper) CanConformColorOfSourceFrames() bool {
	rv := objc.Call[bool](v_, objc.Sel("canConformColorOfSourceFrames"))
	return rv
}

func (v_ VideoCompositingWrapper) HasSourcePixelBufferAttributes() bool {
	return v_.RespondsToSelector(objc.Sel("sourcePixelBufferAttributes"))
}

// The pixel buffer attributes that the compositor accepts for source frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositing/1388610-sourcepixelbufferattributes?language=objc
func (v_ VideoCompositingWrapper) SourcePixelBufferAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](v_, objc.Sel("sourcePixelBufferAttributes"))
	return rv
}
