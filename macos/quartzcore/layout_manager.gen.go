// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// Methods that allow an object to manage the layout of a layer and its sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayoutmanager?language=objc
type PLayoutManager interface {
	// optional
	LayoutSublayersOfLayer(layer Layer)
	HasLayoutSublayersOfLayer() bool

	// optional
	InvalidateLayoutOfLayer(layer Layer)
	HasInvalidateLayoutOfLayer() bool

	// optional
	PreferredSizeOfLayer(layer Layer) coregraphics.Size
	HasPreferredSizeOfLayer() bool
}

// A concrete type wrapper for the [PLayoutManager] protocol.
type LayoutManagerWrapper struct {
	objc.Object
}

func (l_ LayoutManagerWrapper) HasLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.Sel("layoutSublayersOfLayer:"))
}

// Override to customize layout of sublayers whenever the layer needs redrawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayoutmanager/2097260-layoutsublayersoflayer?language=objc
func (l_ LayoutManagerWrapper) LayoutSublayersOfLayer(layer ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("layoutSublayersOfLayer:"), objc.Ptr(layer))
}

func (l_ LayoutManagerWrapper) HasInvalidateLayoutOfLayer() bool {
	return l_.RespondsToSelector(objc.Sel("invalidateLayoutOfLayer:"))
}

// Invalidates the layout of a layer so it knows to refresh its content on the next frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayoutmanager/2097258-invalidatelayoutoflayer?language=objc
func (l_ LayoutManagerWrapper) InvalidateLayoutOfLayer(layer ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("invalidateLayoutOfLayer:"), objc.Ptr(layer))
}

func (l_ LayoutManagerWrapper) HasPreferredSizeOfLayer() bool {
	return l_.RespondsToSelector(objc.Sel("preferredSizeOfLayer:"))
}

// Override to customize layer size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayoutmanager/2097256-preferredsizeoflayer?language=objc
func (l_ LayoutManagerWrapper) PreferredSizeOfLayer(layer ILayer) coregraphics.Size {
	rv := objc.Call[coregraphics.Size](l_, objc.Sel("preferredSizeOfLayer:"), objc.Ptr(layer))
	return rv
}
