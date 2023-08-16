// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an attributed-text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciattributedtextimagegenerator?language=objc
type PAttributedTextImageGenerator interface {
	// optional
	SetScaleFactor(value float64)
	HasSetScaleFactor() bool

	// optional
	ScaleFactor() float64
	HasScaleFactor() bool

	// optional
	SetText(value foundation.AttributedString)
	HasSetText() bool

	// optional
	Text() foundation.IAttributedString
	HasText() bool
}

// A concrete type wrapper for the [PAttributedTextImageGenerator] protocol.
type AttributedTextImageGeneratorWrapper struct {
	objc.Object
}

func (a_ AttributedTextImageGeneratorWrapper) HasSetScaleFactor() bool {
	return a_.RespondsToSelector(objc.Sel("setScaleFactor:"))
}

// The scale at which to render the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciattributedtextimagegenerator/3228060-scalefactor?language=objc
func (a_ AttributedTextImageGeneratorWrapper) SetScaleFactor(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setScaleFactor:"), value)
}

func (a_ AttributedTextImageGeneratorWrapper) HasScaleFactor() bool {
	return a_.RespondsToSelector(objc.Sel("scaleFactor"))
}

// The scale at which to render the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciattributedtextimagegenerator/3228060-scalefactor?language=objc
func (a_ AttributedTextImageGeneratorWrapper) ScaleFactor() float64 {
	rv := objc.Call[float64](a_, objc.Sel("scaleFactor"))
	return rv
}

func (a_ AttributedTextImageGeneratorWrapper) HasSetText() bool {
	return a_.RespondsToSelector(objc.Sel("setText:"))
}

// The text to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciattributedtextimagegenerator/3228061-text?language=objc
func (a_ AttributedTextImageGeneratorWrapper) SetText(value foundation.IAttributedString) {
	objc.Call[objc.Void](a_, objc.Sel("setText:"), objc.Ptr(value))
}

func (a_ AttributedTextImageGeneratorWrapper) HasText() bool {
	return a_.RespondsToSelector(objc.Sel("text"))
}

// The text to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciattributedtextimagegenerator/3228061-text?language=objc
func (a_ AttributedTextImageGeneratorWrapper) Text() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](a_, objc.Sel("text"))
	return rv
}
