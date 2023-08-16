// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorSampler] class.
var ColorSamplerClass = _ColorSamplerClass{objc.GetClass("NSColorSampler")}

type _ColorSamplerClass struct {
	objc.Class
}

// An interface definition for the [ColorSampler] class.
type IColorSampler interface {
	objc.IObject
	ShowSamplerWithSelectionHandler(selectionHandler func(selectedColor Color))
}

// An object that displays the system's color-sampling interface and returns the selected color to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorsampler?language=objc
type ColorSampler struct {
	objc.Object
}

func ColorSamplerFrom(ptr unsafe.Pointer) ColorSampler {
	return ColorSampler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ColorSamplerClass) Alloc() ColorSampler {
	rv := objc.Call[ColorSampler](cc, objc.Sel("alloc"))
	return rv
}

func ColorSampler_Alloc() ColorSampler {
	return ColorSamplerClass.Alloc()
}

func (cc _ColorSamplerClass) New() ColorSampler {
	rv := objc.Call[ColorSampler](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorSampler() ColorSampler {
	return ColorSamplerClass.New()
}

func (c_ ColorSampler) Init() ColorSampler {
	rv := objc.Call[ColorSampler](c_, objc.Sel("init"))
	return rv
}

// Displays the system color-sampling interface asynchronously and reports the selected color back to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorsampler/3237187-showsamplerwithselectionhandler?language=objc
func (c_ ColorSampler) ShowSamplerWithSelectionHandler(selectionHandler func(selectedColor Color)) {
	objc.Call[objc.Void](c_, objc.Sel("showSamplerWithSelectionHandler:"), selectionHandler)
}
