// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// An optional layer delegate method for handling resolution changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate?language=objc
type PViewLayerContentScaleDelegate interface {
	// optional
	LayerShouldInheritContentsScaleFromWindow(layer quartzcore.Layer, newScale float64, window Window) bool
	HasLayerShouldInheritContentsScaleFromWindow() bool
}

// A delegate implementation builder for the [PViewLayerContentScaleDelegate] protocol.
type ViewLayerContentScaleDelegate struct {
	_LayerShouldInheritContentsScaleFromWindow func(layer quartzcore.Layer, newScale float64, window Window) bool
}

func (di *ViewLayerContentScaleDelegate) HasLayerShouldInheritContentsScaleFromWindow() bool {
	return di._LayerShouldInheritContentsScaleFromWindow != nil
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (di *ViewLayerContentScaleDelegate) SetLayerShouldInheritContentsScaleFromWindow(f func(layer quartzcore.Layer, newScale float64, window Window) bool) {
	di._LayerShouldInheritContentsScaleFromWindow = f
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (di *ViewLayerContentScaleDelegate) LayerShouldInheritContentsScaleFromWindow(layer quartzcore.Layer, newScale float64, window Window) bool {
	return di._LayerShouldInheritContentsScaleFromWindow(layer, newScale, window)
}

// A concrete type wrapper for the [PViewLayerContentScaleDelegate] protocol.
type ViewLayerContentScaleDelegateWrapper struct {
	objc.Object
}

func (v_ ViewLayerContentScaleDelegateWrapper) HasLayerShouldInheritContentsScaleFromWindow() bool {
	return v_.RespondsToSelector(objc.Sel("layer:shouldInheritContentsScale:fromWindow:"))
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (v_ ViewLayerContentScaleDelegateWrapper) LayerShouldInheritContentsScaleFromWindow(layer quartzcore.ILayer, newScale float64, window IWindow) bool {
	rv := objc.Call[bool](v_, objc.Sel("layer:shouldInheritContentsScale:fromWindow:"), objc.Ptr(layer), newScale, objc.Ptr(window))
	return rv
}
