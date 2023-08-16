// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageProcessorKernel] class.
var ImageProcessorKernelClass = _ImageProcessorKernelClass{objc.GetClass("CIImageProcessorKernel")}

type _ImageProcessorKernelClass struct {
	objc.Class
}

// An interface definition for the [ImageProcessorKernel] class.
type IImageProcessorKernel interface {
	objc.IObject
}

// The abstract class you extend to create custom image processors that can integrate with Core Image workflows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel?language=objc
type ImageProcessorKernel struct {
	objc.Object
}

func ImageProcessorKernelFrom(ptr unsafe.Pointer) ImageProcessorKernel {
	return ImageProcessorKernel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageProcessorKernelClass) Alloc() ImageProcessorKernel {
	rv := objc.Call[ImageProcessorKernel](ic, objc.Sel("alloc"))
	return rv
}

func ImageProcessorKernel_Alloc() ImageProcessorKernel {
	return ImageProcessorKernelClass.Alloc()
}

func (ic _ImageProcessorKernelClass) New() ImageProcessorKernel {
	rv := objc.Call[ImageProcessorKernel](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageProcessorKernel() ImageProcessorKernel {
	return ImageProcessorKernelClass.New()
}

func (i_ ImageProcessorKernel) Init() ImageProcessorKernel {
	rv := objc.Call[ImageProcessorKernel](i_, objc.Sel("init"))
	return rv
}

// Method to override for determining specific region of input image required to process in rendering a specified region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138287-roiforinput?language=objc
func (ic _ImageProcessorKernelClass) RoiForInputArgumentsOutputRect(input int, arguments map[string]objc.IObject, outputRect coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](ic, objc.Sel("roiForInput:arguments:outputRect:"), input, arguments, outputRect)
	return rv
}

// Method to override for determining specific region of input image required to process in rendering a specified region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138287-roiforinput?language=objc
func ImageProcessorKernel_RoiForInputArgumentsOutputRect(input int, arguments map[string]objc.IObject, outputRect coregraphics.Rect) coregraphics.Rect {
	return ImageProcessorKernelClass.RoiForInputArgumentsOutputRect(input, arguments, outputRect)
}

// Method to override for customizing the kernel's image processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138290-processwithinputs?language=objc
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputError(inputs []PImageProcessorInput, arguments map[string]objc.IObject, output PImageProcessorOutput, error foundation.IError) bool {
	po2 := objc.WrapAsProtocol("CIImageProcessorOutput", output)
	rv := objc.Call[bool](ic, objc.Sel("processWithInputs:arguments:output:error:"), inputs, arguments, po2, objc.Ptr(error))
	return rv
}

// Method to override for customizing the kernel's image processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138290-processwithinputs?language=objc
func ImageProcessorKernel_ProcessWithInputsArgumentsOutputError(inputs []PImageProcessorInput, arguments map[string]objc.IObject, output PImageProcessorOutput, error foundation.IError) bool {
	return ImageProcessorKernelClass.ProcessWithInputsArgumentsOutputError(inputs, arguments, output, error)
}

// Method to override for customizing the kernel's image processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138290-processwithinputs?language=objc
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputObjectError(inputs []PImageProcessorInput, arguments map[string]objc.IObject, outputObject objc.IObject, error foundation.IError) bool {
	rv := objc.Call[bool](ic, objc.Sel("processWithInputs:arguments:output:error:"), inputs, arguments, objc.Ptr(outputObject), objc.Ptr(error))
	return rv
}

// Method to override for customizing the kernel's image processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138290-processwithinputs?language=objc
func ImageProcessorKernel_ProcessWithInputsArgumentsOutputObjectError(inputs []PImageProcessorInput, arguments map[string]objc.IObject, outputObject objc.IObject, error foundation.IError) bool {
	return ImageProcessorKernelClass.ProcessWithInputsArgumentsOutputObjectError(inputs, arguments, outputObject, error)
}

// Method to override for returning the image processing kernel's input pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138289-formatforinputatindex?language=objc
func (ic _ImageProcessorKernelClass) FormatForInputAtIndex(input int) Format {
	rv := objc.Call[Format](ic, objc.Sel("formatForInputAtIndex:"), input)
	return rv
}

// Method to override for returning the image processing kernel's input pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138289-formatforinputatindex?language=objc
func ImageProcessorKernel_FormatForInputAtIndex(input int) Format {
	return ImageProcessorKernelClass.FormatForInputAtIndex(input)
}

// Method to override when applying a custom image processor kernel to an image and returning the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138284-applywithextent?language=objc
func (ic _ImageProcessorKernelClass) ApplyWithExtentInputsArgumentsError(extent coregraphics.Rect, inputs []IImage, args map[string]objc.IObject, error foundation.IError) Image {
	rv := objc.Call[Image](ic, objc.Sel("applyWithExtent:inputs:arguments:error:"), extent, inputs, args, objc.Ptr(error))
	return rv
}

// Method to override when applying a custom image processor kernel to an image and returning the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2138284-applywithextent?language=objc
func ImageProcessorKernel_ApplyWithExtentInputsArgumentsError(extent coregraphics.Rect, inputs []IImage, args map[string]objc.IObject, error foundation.IError) Image {
	return ImageProcessorKernelClass.ApplyWithExtentInputsArgumentsError(extent, inputs, args, error)
}

// The processor's output pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2143065-outputformat?language=objc
func (ic _ImageProcessorKernelClass) OutputFormat() Format {
	rv := objc.Call[Format](ic, objc.Sel("outputFormat"))
	return rv
}

// The processor's output pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2143065-outputformat?language=objc
func ImageProcessorKernel_OutputFormat() Format {
	return ImageProcessorKernelClass.OutputFormat()
}

// Tells whether or not processor input should be synchronized for CPU access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2143066-synchronizeinputs?language=objc
func (ic _ImageProcessorKernelClass) SynchronizeInputs() bool {
	rv := objc.Call[bool](ic, objc.Sel("synchronizeInputs"))
	return rv
}

// Tells whether or not processor input should be synchronized for CPU access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2143066-synchronizeinputs?language=objc
func ImageProcessorKernel_SynchronizeInputs() bool {
	return ImageProcessorKernelClass.SynchronizeInputs()
}

// Boolean determining whether or not processor outputs an opaque image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2867389-outputisopaque?language=objc
func (ic _ImageProcessorKernelClass) OutputIsOpaque() bool {
	rv := objc.Call[bool](ic, objc.Sel("outputIsOpaque"))
	return rv
}

// Boolean determining whether or not processor outputs an opaque image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorkernel/2867389-outputisopaque?language=objc
func ImageProcessorKernel_OutputIsOpaque() bool {
	return ImageProcessorKernelClass.OutputIsOpaque()
}
