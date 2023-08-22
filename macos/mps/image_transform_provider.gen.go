// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

// A general interface for objects that provide image resampling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetransformprovider?language=objc
type PImageTransformProvider interface {
	// optional
	TransformForSourceImageHandle(image Image, handle HandleWrapper) ScaleTransform
	HasTransformForSourceImageHandle() bool
}

// A concrete type wrapper for the [PImageTransformProvider] protocol.
type ImageTransformProviderWrapper struct {
	objc.Object
}

func (i_ ImageTransformProviderWrapper) HasTransformForSourceImageHandle() bool {
	return i_.RespondsToSelector(objc.Sel("transformForSourceImage:handle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetransformprovider/2915282-transformforsourceimage?language=objc
func (i_ ImageTransformProviderWrapper) TransformForSourceImageHandle(image IImage, handle PHandle) ScaleTransform {
	po1 := objc.WrapAsProtocol("MPSHandle", handle)
	rv := objc.Call[ScaleTransform](i_, objc.Sel("transformForSourceImage:handle:"), objc.Ptr(image), po1)
	return rv
}
