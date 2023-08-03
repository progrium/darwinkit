// auto-generated code, do not modify
package quartzcore

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

type ILayoutManager interface {
	ImplementsInvalidateLayoutOfLayer() bool
	// optional
	InvalidateLayoutOfLayer(layer Layer)
	ImplementsLayoutSublayersOfLayer() bool
	// optional
	LayoutSublayersOfLayer(layer Layer)
	ImplementsPreferredSizeOfLayer() bool
	// optional
	PreferredSizeOfLayer(layer Layer) coregraphics.Size
}

type LayoutManagerWrapper struct {
	objc.Object
}

func (l_ LayoutManagerWrapper) ImplementsInvalidateLayoutOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("invalidateLayoutOfLayer:"))
}

func (l_ LayoutManagerWrapper) InvalidateLayoutOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateLayoutOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayoutManagerWrapper) ImplementsLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutSublayersOfLayer:"))
}

func (l_ LayoutManagerWrapper) LayoutSublayersOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutSublayersOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayoutManagerWrapper) ImplementsPreferredSizeOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("preferredSizeOfLayer:"))
}

func (l_ LayoutManagerWrapper) PreferredSizeOfLayer(layer ILayer) coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, objc.GetSelector("preferredSizeOfLayer:"), objc.ExtractPtr(layer))
	return rv
}
