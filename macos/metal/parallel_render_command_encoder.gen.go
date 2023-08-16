// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object that splits up a single render pass so that it can be simultaneously encoded from multiple threads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder?language=objc
type PParallelRenderCommandEncoder interface {
	// optional
	SetDepthStoreAction(storeAction StoreAction)
	HasSetDepthStoreAction() bool

	// optional
	SetStencilStoreAction(storeAction StoreAction)
	HasSetStencilStoreAction() bool

	// optional
	SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint)
	HasSetColorStoreActionOptionsAtIndex() bool

	// optional
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)
	HasSetColorStoreActionAtIndex() bool

	// optional
	SetStencilStoreActionOptions(storeActionOptions StoreActionOptions)
	HasSetStencilStoreActionOptions() bool

	// optional
	RenderCommandEncoder() PRenderCommandEncoder
	HasRenderCommandEncoder() bool

	// optional
	SetDepthStoreActionOptions(storeActionOptions StoreActionOptions)
	HasSetDepthStoreActionOptions() bool
}

// A concrete type wrapper for the [PParallelRenderCommandEncoder] protocol.
type ParallelRenderCommandEncoderWrapper struct {
	objc.Object
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetDepthStoreAction() bool {
	return p_.RespondsToSelector(objc.Sel("setDepthStoreAction:"))
}

// Specifies a known store action to replace the initial [metal/mtlstoreaction/mtlstoreactionunknown] value specified for a given depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/1639937-setdepthstoreaction?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetDepthStoreAction(storeAction StoreAction) {
	objc.Call[objc.Void](p_, objc.Sel("setDepthStoreAction:"), storeAction)
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetStencilStoreAction() bool {
	return p_.RespondsToSelector(objc.Sel("setStencilStoreAction:"))
}

// Specifies a known store action to replace the initial [metal/mtlstoreaction/mtlstoreactionunknown] value specified for a given stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/1640016-setstencilstoreaction?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetStencilStoreAction(storeAction StoreAction) {
	objc.Call[objc.Void](p_, objc.Sel("setStencilStoreAction:"), storeAction)
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetColorStoreActionOptionsAtIndex() bool {
	return p_.RespondsToSelector(objc.Sel("setColorStoreActionOptions:atIndex:"))
}

// Specifies known store action options for a given color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/2919777-setcolorstoreactionoptions?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint) {
	objc.Call[objc.Void](p_, objc.Sel("setColorStoreActionOptions:atIndex:"), storeActionOptions, colorAttachmentIndex)
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetColorStoreActionAtIndex() bool {
	return p_.RespondsToSelector(objc.Sel("setColorStoreAction:atIndex:"))
}

// Specifies a known store action to replace the initial [metal/mtlstoreaction/mtlstoreactionunknown] value specified for a given color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/1639891-setcolorstoreaction?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint) {
	objc.Call[objc.Void](p_, objc.Sel("setColorStoreAction:atIndex:"), storeAction, colorAttachmentIndex)
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetStencilStoreActionOptions() bool {
	return p_.RespondsToSelector(objc.Sel("setStencilStoreActionOptions:"))
}

// Specifies known store action options for a given stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/2919778-setstencilstoreactionoptions?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetStencilStoreActionOptions(storeActionOptions StoreActionOptions) {
	objc.Call[objc.Void](p_, objc.Sel("setStencilStoreActionOptions:"), storeActionOptions)
}

func (p_ ParallelRenderCommandEncoderWrapper) HasRenderCommandEncoder() bool {
	return p_.RespondsToSelector(objc.Sel("renderCommandEncoder"))
}

// Create an object that encodes commands that perform graphics rendering operations and may be assigned to a different thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/1515911-rendercommandencoder?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) RenderCommandEncoder() RenderCommandEncoderWrapper {
	rv := objc.Call[RenderCommandEncoderWrapper](p_, objc.Sel("renderCommandEncoder"))
	return rv
}

func (p_ ParallelRenderCommandEncoderWrapper) HasSetDepthStoreActionOptions() bool {
	return p_.RespondsToSelector(objc.Sel("setDepthStoreActionOptions:"))
}

// Specifies known store action options for a given depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlparallelrendercommandencoder/2919779-setdepthstoreactionoptions?language=objc
func (p_ ParallelRenderCommandEncoderWrapper) SetDepthStoreActionOptions(storeActionOptions StoreActionOptions) {
	objc.Call[objc.Void](p_, objc.Sel("setDepthStoreActionOptions:"), storeActionOptions)
}
