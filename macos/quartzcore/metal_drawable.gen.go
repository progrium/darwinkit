// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// A Metal drawable associated with a Core Animation layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldrawable?language=objc
type PMetalDrawable interface {
	// optional
	Layer() IMetalLayer
	HasLayer() bool

	// optional
	Texture() metal.PTexture
	HasTexture() bool
}

// A concrete type wrapper for the [PMetalDrawable] protocol.
type MetalDrawableWrapper struct {
	objc.Object
}

func (m_ MetalDrawableWrapper) HasLayer() bool {
	return m_.RespondsToSelector(objc.Sel("layer"))
}

// The layer that owns this drawable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldrawable/1478165-layer?language=objc
func (m_ MetalDrawableWrapper) Layer() MetalLayer {
	rv := objc.Call[MetalLayer](m_, objc.Sel("layer"))
	return rv
}

func (m_ MetalDrawableWrapper) HasTexture() bool {
	return m_.RespondsToSelector(objc.Sel("texture"))
}

// A Metal texture object that contains the drawable’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldrawable/1478159-texture?language=objc
func (m_ MetalDrawableWrapper) Texture() metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](m_, objc.Sel("texture"))
	return rv
}
