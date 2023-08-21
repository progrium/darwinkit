// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Sampler] class.
var SamplerClass = _SamplerClass{objc.GetClass("CISampler")}

type _SamplerClass struct {
	objc.Class
}

// An interface definition for the [Sampler] class.
type ISampler interface {
	objc.IObject
	Extent() coregraphics.Rect
	Definition() FilterShape
}

// An object that retrieves pixel samples for processing by a filter kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisampler?language=objc
type Sampler struct {
	objc.Object
}

func SamplerFrom(ptr unsafe.Pointer) Sampler {
	return Sampler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SamplerClass) SamplerWithImage(im IImage) Sampler {
	rv := objc.Call[Sampler](sc, objc.Sel("samplerWithImage:"), objc.Ptr(im))
	return rv
}

// Creates and returns a sampler that references an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisampler/1555075-samplerwithimage?language=objc
func Sampler_SamplerWithImage(im IImage) Sampler {
	return SamplerClass.SamplerWithImage(im)
}

func (s_ Sampler) InitWithImage(im IImage) Sampler {
	rv := objc.Call[Sampler](s_, objc.Sel("initWithImage:"), objc.Ptr(im))
	return rv
}

// Initializes a sampler with an image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisampler/1438117-initwithimage?language=objc
func NewSamplerWithImage(im IImage) Sampler {
	instance := SamplerClass.Alloc().InitWithImage(im)
	instance.Autorelease()
	return instance
}

func (sc _SamplerClass) Alloc() Sampler {
	rv := objc.Call[Sampler](sc, objc.Sel("alloc"))
	return rv
}

func Sampler_Alloc() Sampler {
	return SamplerClass.Alloc()
}

func (sc _SamplerClass) New() Sampler {
	rv := objc.Call[Sampler](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampler() Sampler {
	return SamplerClass.New()
}

func (s_ Sampler) Init() Sampler {
	rv := objc.Call[Sampler](s_, objc.Sel("init"))
	return rv
}

// The rectangle that specifies the extent of the sampler [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisampler/1437872-extent?language=objc
func (s_ Sampler) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](s_, objc.Sel("extent"))
	return rv
}

// The domain of definition (DOD) of the sampler [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisampler/1437877-definition?language=objc
func (s_ Sampler) Definition() FilterShape {
	rv := objc.Call[FilterShape](s_, objc.Sel("definition"))
	return rv
}
