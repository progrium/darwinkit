// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an Aztec code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator?language=objc
type PAztecCodeGenerator interface {
	// optional
	SetLayers(value float64)
	HasSetLayers() bool

	// optional
	Layers() float64
	HasLayers() bool

	// optional
	SetCompactStyle(value float64)
	HasSetCompactStyle() bool

	// optional
	CompactStyle() float64
	HasCompactStyle() bool

	// optional
	SetMessage(value []byte)
	HasSetMessage() bool

	// optional
	Message() []byte
	HasMessage() bool

	// optional
	SetCorrectionLevel(value float64)
	HasSetCorrectionLevel() bool

	// optional
	CorrectionLevel() float64
	HasCorrectionLevel() bool
}

// A concrete type wrapper for the [PAztecCodeGenerator] protocol.
type AztecCodeGeneratorWrapper struct {
	objc.Object
}

func (a_ AztecCodeGeneratorWrapper) HasSetLayers() bool {
	return a_.RespondsToSelector(objc.Sel("setLayers:"))
}

// The number of Aztec layers, a value from 1 to 32. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228065-layers?language=objc
func (a_ AztecCodeGeneratorWrapper) SetLayers(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setLayers:"), value)
}

func (a_ AztecCodeGeneratorWrapper) HasLayers() bool {
	return a_.RespondsToSelector(objc.Sel("layers"))
}

// The number of Aztec layers, a value from 1 to 32. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228065-layers?language=objc
func (a_ AztecCodeGeneratorWrapper) Layers() float64 {
	rv := objc.Call[float64](a_, objc.Sel("layers"))
	return rv
}

func (a_ AztecCodeGeneratorWrapper) HasSetCompactStyle() bool {
	return a_.RespondsToSelector(objc.Sel("setCompactStyle:"))
}

// A Boolean that specifies whether to force a compact style Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228063-compactstyle?language=objc
func (a_ AztecCodeGeneratorWrapper) SetCompactStyle(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setCompactStyle:"), value)
}

func (a_ AztecCodeGeneratorWrapper) HasCompactStyle() bool {
	return a_.RespondsToSelector(objc.Sel("compactStyle"))
}

// A Boolean that specifies whether to force a compact style Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228063-compactstyle?language=objc
func (a_ AztecCodeGeneratorWrapper) CompactStyle() float64 {
	rv := objc.Call[float64](a_, objc.Sel("compactStyle"))
	return rv
}

func (a_ AztecCodeGeneratorWrapper) HasSetMessage() bool {
	return a_.RespondsToSelector(objc.Sel("setMessage:"))
}

// The message to encode in the Aztec barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228066-message?language=objc
func (a_ AztecCodeGeneratorWrapper) SetMessage(value []byte) {
	objc.Call[objc.Void](a_, objc.Sel("setMessage:"), value)
}

func (a_ AztecCodeGeneratorWrapper) HasMessage() bool {
	return a_.RespondsToSelector(objc.Sel("message"))
}

// The message to encode in the Aztec barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228066-message?language=objc
func (a_ AztecCodeGeneratorWrapper) Message() []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("message"))
	return rv
}

func (a_ AztecCodeGeneratorWrapper) HasSetCorrectionLevel() bool {
	return a_.RespondsToSelector(objc.Sel("setCorrectionLevel:"))
}

// The Aztec error correction, a value from 5 to 95. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228064-correctionlevel?language=objc
func (a_ AztecCodeGeneratorWrapper) SetCorrectionLevel(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setCorrectionLevel:"), value)
}

func (a_ AztecCodeGeneratorWrapper) HasCorrectionLevel() bool {
	return a_.RespondsToSelector(objc.Sel("correctionLevel"))
}

// The Aztec error correction, a value from 5 to 95. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodegenerator/3228064-correctionlevel?language=objc
func (a_ AztecCodeGeneratorWrapper) CorrectionLevel() float64 {
	rv := objc.Call[float64](a_, objc.Sel("correctionLevel"))
	return rv
}
