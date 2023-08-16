// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a spot color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor?language=objc
type PSpotColor interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetContrast2(value float64)
	HasSetContrast2() bool

	// optional
	Contrast2() float64
	HasContrast2() bool

	// optional
	SetCenterColor3(value Color)
	HasSetCenterColor3() bool

	// optional
	CenterColor3() IColor
	HasCenterColor3() bool

	// optional
	SetCenterColor2(value Color)
	HasSetCenterColor2() bool

	// optional
	CenterColor2() IColor
	HasCenterColor2() bool

	// optional
	SetContrast3(value float64)
	HasSetContrast3() bool

	// optional
	Contrast3() float64
	HasContrast3() bool

	// optional
	SetCloseness2(value float64)
	HasSetCloseness2() bool

	// optional
	Closeness2() float64
	HasCloseness2() bool

	// optional
	SetReplacementColor2(value Color)
	HasSetReplacementColor2() bool

	// optional
	ReplacementColor2() IColor
	HasReplacementColor2() bool

	// optional
	SetReplacementColor3(value Color)
	HasSetReplacementColor3() bool

	// optional
	ReplacementColor3() IColor
	HasReplacementColor3() bool

	// optional
	SetCloseness3(value float64)
	HasSetCloseness3() bool

	// optional
	Closeness3() float64
	HasCloseness3() bool

	// optional
	SetCloseness1(value float64)
	HasSetCloseness1() bool

	// optional
	Closeness1() float64
	HasCloseness1() bool

	// optional
	SetReplacementColor1(value Color)
	HasSetReplacementColor1() bool

	// optional
	ReplacementColor1() IColor
	HasReplacementColor1() bool

	// optional
	SetCenterColor1(value Color)
	HasSetCenterColor1() bool

	// optional
	CenterColor1() IColor
	HasCenterColor1() bool

	// optional
	SetContrast1(value float64)
	HasSetContrast1() bool

	// optional
	Contrast1() float64
	HasContrast1() bool
}

// A concrete type wrapper for the [PSpotColor] protocol.
type SpotColorWrapper struct {
	objc.Object
}

func (s_ SpotColorWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228737-inputimage?language=objc
func (s_ SpotColorWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228737-inputimage?language=objc
func (s_ SpotColorWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SpotColorWrapper) HasSetContrast2() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast2:"))
}

// The contrast of the second replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228735-contrast2?language=objc
func (s_ SpotColorWrapper) SetContrast2(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast2:"), value)
}

func (s_ SpotColorWrapper) HasContrast2() bool {
	return s_.RespondsToSelector(objc.Sel("contrast2"))
}

// The contrast of the second replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228735-contrast2?language=objc
func (s_ SpotColorWrapper) Contrast2() float64 {
	rv := objc.Call[float64](s_, objc.Sel("contrast2"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCenterColor3() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor3:"))
}

// The center value of the third color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228730-centercolor3?language=objc
func (s_ SpotColorWrapper) SetCenterColor3(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor3:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasCenterColor3() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor3"))
}

// The center value of the third color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228730-centercolor3?language=objc
func (s_ SpotColorWrapper) CenterColor3() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor3"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCenterColor2() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor2:"))
}

// The center value of the second color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228729-centercolor2?language=objc
func (s_ SpotColorWrapper) SetCenterColor2(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor2:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasCenterColor2() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor2"))
}

// The center value of the second color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228729-centercolor2?language=objc
func (s_ SpotColorWrapper) CenterColor2() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor2"))
	return rv
}

func (s_ SpotColorWrapper) HasSetContrast3() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast3:"))
}

// The contrast of the third replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228736-contrast3?language=objc
func (s_ SpotColorWrapper) SetContrast3(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast3:"), value)
}

func (s_ SpotColorWrapper) HasContrast3() bool {
	return s_.RespondsToSelector(objc.Sel("contrast3"))
}

// The contrast of the third replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228736-contrast3?language=objc
func (s_ SpotColorWrapper) Contrast3() float64 {
	rv := objc.Call[float64](s_, objc.Sel("contrast3"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCloseness2() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness2:"))
}

// A value that indicates how closely the second color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228732-closeness2?language=objc
func (s_ SpotColorWrapper) SetCloseness2(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness2:"), value)
}

func (s_ SpotColorWrapper) HasCloseness2() bool {
	return s_.RespondsToSelector(objc.Sel("closeness2"))
}

// A value that indicates how closely the second color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228732-closeness2?language=objc
func (s_ SpotColorWrapper) Closeness2() float64 {
	rv := objc.Call[float64](s_, objc.Sel("closeness2"))
	return rv
}

func (s_ SpotColorWrapper) HasSetReplacementColor2() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor2:"))
}

// A replacement color for the second color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228739-replacementcolor2?language=objc
func (s_ SpotColorWrapper) SetReplacementColor2(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor2:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasReplacementColor2() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor2"))
}

// A replacement color for the second color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228739-replacementcolor2?language=objc
func (s_ SpotColorWrapper) ReplacementColor2() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor2"))
	return rv
}

func (s_ SpotColorWrapper) HasSetReplacementColor3() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor3:"))
}

// A replacement color for the third color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228740-replacementcolor3?language=objc
func (s_ SpotColorWrapper) SetReplacementColor3(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor3:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasReplacementColor3() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor3"))
}

// A replacement color for the third color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228740-replacementcolor3?language=objc
func (s_ SpotColorWrapper) ReplacementColor3() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor3"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCloseness3() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness3:"))
}

// A value that indicates how closely the third color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228733-closeness3?language=objc
func (s_ SpotColorWrapper) SetCloseness3(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness3:"), value)
}

func (s_ SpotColorWrapper) HasCloseness3() bool {
	return s_.RespondsToSelector(objc.Sel("closeness3"))
}

// A value that indicates how closely the third color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228733-closeness3?language=objc
func (s_ SpotColorWrapper) Closeness3() float64 {
	rv := objc.Call[float64](s_, objc.Sel("closeness3"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCloseness1() bool {
	return s_.RespondsToSelector(objc.Sel("setCloseness1:"))
}

// A value that indicates how closely the first color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228731-closeness1?language=objc
func (s_ SpotColorWrapper) SetCloseness1(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCloseness1:"), value)
}

func (s_ SpotColorWrapper) HasCloseness1() bool {
	return s_.RespondsToSelector(objc.Sel("closeness1"))
}

// A value that indicates how closely the first color must match before it’s replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228731-closeness1?language=objc
func (s_ SpotColorWrapper) Closeness1() float64 {
	rv := objc.Call[float64](s_, objc.Sel("closeness1"))
	return rv
}

func (s_ SpotColorWrapper) HasSetReplacementColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setReplacementColor1:"))
}

// A replacement color for the first color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228738-replacementcolor1?language=objc
func (s_ SpotColorWrapper) SetReplacementColor1(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setReplacementColor1:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasReplacementColor1() bool {
	return s_.RespondsToSelector(objc.Sel("replacementColor1"))
}

// A replacement color for the first color range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228738-replacementcolor1?language=objc
func (s_ SpotColorWrapper) ReplacementColor1() Color {
	rv := objc.Call[Color](s_, objc.Sel("replacementColor1"))
	return rv
}

func (s_ SpotColorWrapper) HasSetCenterColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterColor1:"))
}

// The center value of the first color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228728-centercolor1?language=objc
func (s_ SpotColorWrapper) SetCenterColor1(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterColor1:"), objc.Ptr(value))
}

func (s_ SpotColorWrapper) HasCenterColor1() bool {
	return s_.RespondsToSelector(objc.Sel("centerColor1"))
}

// The center value of the first color range to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228728-centercolor1?language=objc
func (s_ SpotColorWrapper) CenterColor1() Color {
	rv := objc.Call[Color](s_, objc.Sel("centerColor1"))
	return rv
}

func (s_ SpotColorWrapper) HasSetContrast1() bool {
	return s_.RespondsToSelector(objc.Sel("setContrast1:"))
}

// The contrast of the first replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228734-contrast1?language=objc
func (s_ SpotColorWrapper) SetContrast1(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setContrast1:"), value)
}

func (s_ SpotColorWrapper) HasContrast1() bool {
	return s_.RespondsToSelector(objc.Sel("contrast1"))
}

// The contrast of the first replacement color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotcolor/3228734-contrast1?language=objc
func (s_ SpotColorWrapper) Contrast1() float64 {
	rv := objc.Call[float64](s_, objc.Sel("contrast1"))
	return rv
}
