// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a document enhancer filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidocumentenhancer?language=objc
type PDocumentEnhancer interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetAmount(value float64)
	HasSetAmount() bool

	// optional
	Amount() float64
	HasAmount() bool
}

// A concrete type wrapper for the [PDocumentEnhancer] protocol.
type DocumentEnhancerWrapper struct {
	objc.Object
}

func (d_ DocumentEnhancerWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidocumentenhancer/3228229-inputimage?language=objc
func (d_ DocumentEnhancerWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DocumentEnhancerWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidocumentenhancer/3228229-inputimage?language=objc
func (d_ DocumentEnhancerWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DocumentEnhancerWrapper) HasSetAmount() bool {
	return d_.RespondsToSelector(objc.Sel("setAmount:"))
}

// The amount of enhancement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidocumentenhancer/3228228-amount?language=objc
func (d_ DocumentEnhancerWrapper) SetAmount(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setAmount:"), value)
}

func (d_ DocumentEnhancerWrapper) HasAmount() bool {
	return d_.RespondsToSelector(objc.Sel("amount"))
}

// The amount of enhancement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidocumentenhancer/3228228-amount?language=objc
func (d_ DocumentEnhancerWrapper) Amount() float64 {
	rv := objc.Call[float64](d_, objc.Sel("amount"))
	return rv
}
