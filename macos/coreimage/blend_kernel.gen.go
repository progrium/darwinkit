// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BlendKernel] class.
var BlendKernelClass = _BlendKernelClass{objc.GetClass("CIBlendKernel")}

type _BlendKernelClass struct {
	objc.Class
}

// An interface definition for the [BlendKernel] class.
type IBlendKernel interface {
	IColorKernel
	ApplyWithForegroundBackgroundColorSpace(foreground IImage, background IImage, colorSpace coregraphics.ColorSpaceRef) Image
}

// A GPU-based image processing routine that is optimized for blending two images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel?language=objc
type BlendKernel struct {
	ColorKernel
}

func BlendKernelFrom(ptr unsafe.Pointer) BlendKernel {
	return BlendKernel{
		ColorKernel: ColorKernelFrom(ptr),
	}
}

func (bc _BlendKernelClass) Alloc() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("alloc"))
	return rv
}

func BlendKernel_Alloc() BlendKernel {
	return BlendKernelClass.Alloc()
}

func (bc _BlendKernelClass) New() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBlendKernel() BlendKernel {
	return BlendKernelClass.New()
}

func (b_ BlendKernel) Init() BlendKernel {
	rv := objc.Call[BlendKernel](b_, objc.Sel("init"))
	return rv
}

func (bc _BlendKernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, objc.Ptr(error))
	return rv
}

// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/2880195-kernelwithfunctionname?language=objc
func BlendKernel_KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) BlendKernel {
	return BlendKernelClass.KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name, data, format, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/3152403-applywithforeground?language=objc
func (b_ BlendKernel) ApplyWithForegroundBackgroundColorSpace(foreground IImage, background IImage, colorSpace coregraphics.ColorSpaceRef) Image {
	rv := objc.Call[Image](b_, objc.Sel("applyWithForeground:background:colorSpace:"), objc.Ptr(foreground), objc.Ptr(background), colorSpace)
	return rv
}

// A blend kernel that uses the foreground image to define what to take out of the background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867415-sourceout?language=objc
func (bc _BlendKernelClass) SourceOut() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("sourceOut"))
	return rv
}

// A blend kernel that uses the foreground image to define what to take out of the background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867415-sourceout?language=objc
func BlendKernel_SourceOut() BlendKernel {
	return BlendKernelClass.SourceOut()
}

// A blend kernel that conditionally replaces background image samples with source image samples depending on the brightness of the source image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867420-pinlight?language=objc
func (bc _BlendKernelClass) PinLight() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("pinLight"))
	return rv
}

// A blend kernel that conditionally replaces background image samples with source image samples depending on the brightness of the source image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867420-pinlight?language=objc
func BlendKernel_PinLight() BlendKernel {
	return BlendKernelClass.PinLight()
}

// A blend kernel that uses the luminance values of the background with the hue and saturation values of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867350-color?language=objc
func (bc _BlendKernelClass) Color() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("color"))
	return rv
}

// A blend kernel that uses the luminance values of the background with the hue and saturation values of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867350-color?language=objc
func BlendKernel_Color() BlendKernel {
	return BlendKernelClass.Color()
}

// A blend kernel that creates an image using the maximum values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867433-componentmax?language=objc
func (bc _BlendKernelClass) ComponentMax() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("componentMax"))
	return rv
}

// A blend kernel that creates an image using the maximum values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867433-componentmax?language=objc
func BlendKernel_ComponentMax() BlendKernel {
	return BlendKernelClass.ComponentMax()
}

// A blend kernel that places the foreground image over the input background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867413-sourceover?language=objc
func (bc _BlendKernelClass) SourceOver() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("sourceOver"))
	return rv
}

// A blend kernel that places the foreground image over the input background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867413-sourceover?language=objc
func BlendKernel_SourceOver() BlendKernel {
	return BlendKernelClass.SourceOver()
}

// A blend kernel that multiplies the inverse of the foreground image samples with the inverse of the background image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867356-screen?language=objc
func (bc _BlendKernelClass) Screen() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("screen"))
	return rv
}

// A blend kernel that multiplies the inverse of the foreground image samples with the inverse of the background image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867356-screen?language=objc
func BlendKernel_Screen() BlendKernel {
	return BlendKernelClass.Screen()
}

// A blend kernel that creates an image using the darker values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867348-darken?language=objc
func (bc _BlendKernelClass) Darken() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("darken"))
	return rv
}

// A blend kernel that creates an image using the darker values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867348-darken?language=objc
func BlendKernel_Darken() BlendKernel {
	return BlendKernelClass.Darken()
}

// A blend kernel that creates an image using the minimum values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867425-componentmin?language=objc
func (bc _BlendKernelClass) ComponentMin() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("componentMin"))
	return rv
}

// A blend kernel that creates an image using the minimum values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867425-componentmin?language=objc
func BlendKernel_ComponentMin() BlendKernel {
	return BlendKernelClass.ComponentMin()
}

// A blend kernel that uses the hue and saturation of the background image with the luminance of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867423-luminosity?language=objc
func (bc _BlendKernelClass) Luminosity() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("luminosity"))
	return rv
}

// A blend kernel that uses the hue and saturation of the background image with the luminance of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867423-luminosity?language=objc
func BlendKernel_Luminosity() BlendKernel {
	return BlendKernelClass.Luminosity()
}

// A blend kernel that produces an effect similar to difference blending but with lower contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867343-exclusion?language=objc
func (bc _BlendKernelClass) Exclusion() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("exclusion"))
	return rv
}

// A blend kernel that produces an effect similar to difference blending but with lower contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867343-exclusion?language=objc
func BlendKernel_Exclusion() BlendKernel {
	return BlendKernelClass.Exclusion()
}

// A blend kernel that uses the luminance and hue values of the background image with the saturation of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867431-saturation?language=objc
func (bc _BlendKernelClass) Saturation() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("saturation"))
	return rv
}

// A blend kernel that uses the luminance and hue values of the background image with the saturation of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867431-saturation?language=objc
func BlendKernel_Saturation() BlendKernel {
	return BlendKernelClass.Saturation()
}

// A blend kernel that returns the foreground input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867407-source?language=objc
func (bc _BlendKernelClass) Source() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("source"))
	return rv
}

// A blend kernel that returns the foreground input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867407-source?language=objc
func BlendKernel_Source() BlendKernel {
	return BlendKernelClass.Source()
}

// A blend kernel that either multiplies or screens the foreground image samples with the background image samples, depending on the background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867411-overlay?language=objc
func (bc _BlendKernelClass) Overlay() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("overlay"))
	return rv
}

// A blend kernel that either multiplies or screens the foreground image samples with the background image samples, depending on the background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867411-overlay?language=objc
func BlendKernel_Overlay() BlendKernel {
	return BlendKernelClass.Overlay()
}

// A blend kernel that places the background image over the input foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867432-destinationover?language=objc
func (bc _BlendKernelClass) DestinationOver() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("destinationOver"))
	return rv
}

// A blend kernel that places the background image over the input foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867432-destinationover?language=objc
func BlendKernel_DestinationOver() BlendKernel {
	return BlendKernelClass.DestinationOver()
}

// A blend kernel that returns a clear color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867388-clear?language=objc
func (bc _BlendKernelClass) Clear() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("clear"))
	return rv
}

// A blend kernel that returns a clear color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867388-clear?language=objc
func BlendKernel_Clear() BlendKernel {
	return BlendKernelClass.Clear()
}

// A blend kernel that creates an image using the lighter color of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867427-lightercolor?language=objc
func (bc _BlendKernelClass) LighterColor() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("lighterColor"))
	return rv
}

// A blend kernel that creates an image using the lighter color of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867427-lightercolor?language=objc
func BlendKernel_LighterColor() BlendKernel {
	return BlendKernelClass.LighterColor()
}

// A blend kernel that divides the background image sample color with the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867410-divide?language=objc
func (bc _BlendKernelClass) Divide() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("divide"))
	return rv
}

// A blend kernel that divides the background image sample color with the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867410-divide?language=objc
func BlendKernel_Divide() BlendKernel {
	return BlendKernelClass.Divide()
}

// A blend kernel that burns or dodges colors by changing brightness, depending on the blend color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867435-linearlight?language=objc
func (bc _BlendKernelClass) LinearLight() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("linearLight"))
	return rv
}

// A blend kernel that burns or dodges colors by changing brightness, depending on the blend color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867435-linearlight?language=objc
func BlendKernel_LinearLight() BlendKernel {
	return BlendKernelClass.LinearLight()
}

// A blend kernel that either darkens or lightens colors, depending on the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867434-softlight?language=objc
func (bc _BlendKernelClass) SoftLight() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("softLight"))
	return rv
}

// A blend kernel that either darkens or lightens colors, depending on the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867434-softlight?language=objc
func BlendKernel_SoftLight() BlendKernel {
	return BlendKernelClass.SoftLight()
}

// A blend kernel that uses the background image to define what to take out of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867368-destinationout?language=objc
func (bc _BlendKernelClass) DestinationOut() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("destinationOut"))
	return rv
}

// A blend kernel that uses the background image to define what to take out of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867368-destinationout?language=objc
func BlendKernel_DestinationOut() BlendKernel {
	return BlendKernelClass.DestinationOut()
}

// A blend kernel that creates an image using the lighter values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867424-lighten?language=objc
func (bc _BlendKernelClass) Lighten() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("lighten"))
	return rv
}

// A blend kernel that creates an image using the lighter values of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867424-lighten?language=objc
func BlendKernel_Lighten() BlendKernel {
	return BlendKernelClass.Lighten()
}

// A blend kernel that brightens the background image samples to reflect the foreground image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867417-colordodge?language=objc
func (bc _BlendKernelClass) ColorDodge() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("colorDodge"))
	return rv
}

// A blend kernel that brightens the background image samples to reflect the foreground image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867417-colordodge?language=objc
func BlendKernel_ColorDodge() BlendKernel {
	return BlendKernelClass.ColorDodge()
}

// A blend kernel that adds two images together, setting each color channel value to either 0 or 1. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867347-hardmix?language=objc
func (bc _BlendKernelClass) HardMix() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("hardMix"))
	return rv
}

// A blend kernel that adds two images together, setting each color channel value to either 0 or 1. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867347-hardmix?language=objc
func BlendKernel_HardMix() BlendKernel {
	return BlendKernelClass.HardMix()
}

// A blend kernel that subtracts the background image sample color from the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867370-subtract?language=objc
func (bc _BlendKernelClass) Subtract() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("subtract"))
	return rv
}

// A blend kernel that subtracts the background image sample color from the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867370-subtract?language=objc
func BlendKernel_Subtract() BlendKernel {
	return BlendKernelClass.Subtract()
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples while also increasing contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867409-linearburn?language=objc
func (bc _BlendKernelClass) LinearBurn() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("linearBurn"))
	return rv
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples while also increasing contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867409-linearburn?language=objc
func BlendKernel_LinearBurn() BlendKernel {
	return BlendKernelClass.LinearBurn()
}

// A blend kernel that places the foreground over the background and crops based on the visibility of both. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867428-sourcein?language=objc
func (bc _BlendKernelClass) SourceIn() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("sourceIn"))
	return rv
}

// A blend kernel that places the foreground over the background and crops based on the visibility of both. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867428-sourcein?language=objc
func BlendKernel_SourceIn() BlendKernel {
	return BlendKernelClass.SourceIn()
}

// A blend kernel that returns the background input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867422-destination?language=objc
func (bc _BlendKernelClass) Destination() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("destination"))
	return rv
}

// A blend kernel that returns the background input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867422-destination?language=objc
func BlendKernel_Destination() BlendKernel {
	return BlendKernelClass.Destination()
}

// A blend kernel that multiplies the color components of its input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867406-componentmultiply?language=objc
func (bc _BlendKernelClass) ComponentMultiply() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("componentMultiply"))
	return rv
}

// A blend kernel that multiplies the color components of its input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867406-componentmultiply?language=objc
func BlendKernel_ComponentMultiply() BlendKernel {
	return BlendKernelClass.ComponentMultiply()
}

// A blend kernel that creates an image using the difference between the background and foreground images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867416-difference?language=objc
func (bc _BlendKernelClass) Difference() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("difference"))
	return rv
}

// A blend kernel that creates an image using the difference between the background and foreground images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867416-difference?language=objc
func BlendKernel_Difference() BlendKernel {
	return BlendKernelClass.Difference()
}

// A blend kernel that uses the luminance and saturation values of the background image with the hue of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867408-hue?language=objc
func (bc _BlendKernelClass) Hue() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("hue"))
	return rv
}

// A blend kernel that uses the luminance and saturation values of the background image with the hue of the foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867408-hue?language=objc
func BlendKernel_Hue() BlendKernel {
	return BlendKernelClass.Hue()
}

// A blend kernel that places the foreground over the background and crops based on the visibility of the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867357-sourceatop?language=objc
func (bc _BlendKernelClass) SourceAtop() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("sourceAtop"))
	return rv
}

// A blend kernel that places the foreground over the background and crops based on the visibility of the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867357-sourceatop?language=objc
func BlendKernel_SourceAtop() BlendKernel {
	return BlendKernelClass.SourceAtop()
}

// A blend kernel that returns either the foreground or background image if the other contains a clear color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867421-exclusiveor?language=objc
func (bc _BlendKernelClass) ExclusiveOr() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("exclusiveOr"))
	return rv
}

// A blend kernel that returns either the foreground or background image if the other contains a clear color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867421-exclusiveor?language=objc
func BlendKernel_ExclusiveOr() BlendKernel {
	return BlendKernelClass.ExclusiveOr()
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867391-colorburn?language=objc
func (bc _BlendKernelClass) ColorBurn() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("colorBurn"))
	return rv
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867391-colorburn?language=objc
func BlendKernel_ColorBurn() BlendKernel {
	return BlendKernelClass.ColorBurn()
}

// A blend kernel that adds color components to achieve a brightening effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867384-componentadd?language=objc
func (bc _BlendKernelClass) ComponentAdd() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("componentAdd"))
	return rv
}

// A blend kernel that adds color components to achieve a brightening effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867384-componentadd?language=objc
func BlendKernel_ComponentAdd() BlendKernel {
	return BlendKernelClass.ComponentAdd()
}

// A blend kernel that places the background over the foreground and crops based on the visibility of both. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867352-destinationin?language=objc
func (bc _BlendKernelClass) DestinationIn() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("destinationIn"))
	return rv
}

// A blend kernel that places the background over the foreground and crops based on the visibility of both. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867352-destinationin?language=objc
func BlendKernel_DestinationIn() BlendKernel {
	return BlendKernelClass.DestinationIn()
}

// A blend kernel that lightens the background image samples to reflect the foreground image samples while also increasing contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867354-lineardodge?language=objc
func (bc _BlendKernelClass) LinearDodge() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("linearDodge"))
	return rv
}

// A blend kernel that lightens the background image samples to reflect the foreground image samples while also increasing contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867354-lineardodge?language=objc
func BlendKernel_LinearDodge() BlendKernel {
	return BlendKernelClass.LinearDodge()
}

// A blend kernel that creates an image using the darker color of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867351-darkercolor?language=objc
func (bc _BlendKernelClass) DarkerColor() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("darkerColor"))
	return rv
}

// A blend kernel that creates an image using the darker color of two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867351-darkercolor?language=objc
func BlendKernel_DarkerColor() BlendKernel {
	return BlendKernelClass.DarkerColor()
}

// A blend kernel that places the background over the foreground and crops based on the visibility of the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867385-destinationatop?language=objc
func (bc _BlendKernelClass) DestinationAtop() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("destinationAtop"))
	return rv
}

// A blend kernel that places the background over the foreground and crops based on the visibility of the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867385-destinationatop?language=objc
func BlendKernel_DestinationAtop() BlendKernel {
	return BlendKernelClass.DestinationAtop()
}

// A blend kernel that either multiplies or screens colors, depending on the source image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867418-hardlight?language=objc
func (bc _BlendKernelClass) HardLight() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("hardLight"))
	return rv
}

// A blend kernel that either multiplies or screens colors, depending on the source image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867418-hardlight?language=objc
func BlendKernel_HardLight() BlendKernel {
	return BlendKernelClass.HardLight()
}

// A blend kernel that multiplies the background image sample color with the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867419-multiply?language=objc
func (bc _BlendKernelClass) Multiply() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("multiply"))
	return rv
}

// A blend kernel that multiplies the background image sample color with the foreground image sample color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867419-multiply?language=objc
func BlendKernel_Multiply() BlendKernel {
	return BlendKernelClass.Multiply()
}

// A blend kernel that burns or dodges colors by changing contrast, depending on the blend color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867358-vividlight?language=objc
func (bc _BlendKernelClass) VividLight() BlendKernel {
	rv := objc.Call[BlendKernel](bc, objc.Sel("vividLight"))
	return rv
}

// A blend kernel that burns or dodges colors by changing contrast, depending on the blend color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendkernel/2867358-vividlight?language=objc
func BlendKernel_VividLight() BlendKernel {
	return BlendKernelClass.VividLight()
}
