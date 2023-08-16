// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans?language=objc
type PKMeans interface {
	// optional
	SetInputMeans(value Image)
	HasSetInputMeans() bool

	// optional
	InputMeans() IImage
	HasInputMeans() bool

	// optional
	SetCount(value int)
	HasSetCount() bool

	// optional
	Count() int
	HasCount() bool

	// optional
	SetPerceptual(value bool)
	HasSetPerceptual() bool

	// optional
	Perceptual() bool
	HasPerceptual() bool

	// optional
	SetPasses(value float64)
	HasSetPasses() bool

	// optional
	Passes() float64
	HasPasses() bool
}

// A concrete type wrapper for the [PKMeans] protocol.
type KMeansWrapper struct {
	objc.Object
}

func (k_ KMeansWrapper) HasSetInputMeans() bool {
	return k_.RespondsToSelector(objc.Sel("setInputMeans:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547131-inputmeans?language=objc
func (k_ KMeansWrapper) SetInputMeans(value IImage) {
	objc.Call[objc.Void](k_, objc.Sel("setInputMeans:"), objc.Ptr(value))
}

func (k_ KMeansWrapper) HasInputMeans() bool {
	return k_.RespondsToSelector(objc.Sel("inputMeans"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547131-inputmeans?language=objc
func (k_ KMeansWrapper) InputMeans() Image {
	rv := objc.Call[Image](k_, objc.Sel("inputMeans"))
	return rv
}

func (k_ KMeansWrapper) HasSetCount() bool {
	return k_.RespondsToSelector(objc.Sel("setCount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547130-count?language=objc
func (k_ KMeansWrapper) SetCount(value int) {
	objc.Call[objc.Void](k_, objc.Sel("setCount:"), value)
}

func (k_ KMeansWrapper) HasCount() bool {
	return k_.RespondsToSelector(objc.Sel("count"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547130-count?language=objc
func (k_ KMeansWrapper) Count() int {
	rv := objc.Call[int](k_, objc.Sel("count"))
	return rv
}

func (k_ KMeansWrapper) HasSetPerceptual() bool {
	return k_.RespondsToSelector(objc.Sel("setPerceptual:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547133-perceptual?language=objc
func (k_ KMeansWrapper) SetPerceptual(value bool) {
	objc.Call[objc.Void](k_, objc.Sel("setPerceptual:"), value)
}

func (k_ KMeansWrapper) HasPerceptual() bool {
	return k_.RespondsToSelector(objc.Sel("perceptual"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547133-perceptual?language=objc
func (k_ KMeansWrapper) Perceptual() bool {
	rv := objc.Call[bool](k_, objc.Sel("perceptual"))
	return rv
}

func (k_ KMeansWrapper) HasSetPasses() bool {
	return k_.RespondsToSelector(objc.Sel("setPasses:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547132-passes?language=objc
func (k_ KMeansWrapper) SetPasses(value float64) {
	objc.Call[objc.Void](k_, objc.Sel("setPasses:"), value)
}

func (k_ KMeansWrapper) HasPasses() bool {
	return k_.RespondsToSelector(objc.Sel("passes"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikmeans/3547132-passes?language=objc
func (k_ KMeansWrapper) Passes() float64 {
	rv := objc.Call[float64](k_, objc.Sel("passes"))
	return rv
}
