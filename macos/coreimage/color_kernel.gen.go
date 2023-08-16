// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorKernel] class.
var ColorKernelClass = _ColorKernelClass{objc.GetClass("CIColorKernel")}

type _ColorKernelClass struct {
	objc.Class
}

// An interface definition for the [ColorKernel] class.
type IColorKernel interface {
	IKernel
	ApplyWithExtentArguments(extent coregraphics.Rect, args []objc.IObject) Image
}

// A GPU-based image processing routine that processes only the color information in images, used to create custom Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorkernel?language=objc
type ColorKernel struct {
	Kernel
}

func ColorKernelFrom(ptr unsafe.Pointer) ColorKernel {
	return ColorKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (cc _ColorKernelClass) Alloc() ColorKernel {
	rv := objc.Call[ColorKernel](cc, objc.Sel("alloc"))
	return rv
}

func ColorKernel_Alloc() ColorKernel {
	return ColorKernelClass.Alloc()
}

func (cc _ColorKernelClass) New() ColorKernel {
	rv := objc.Call[ColorKernel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorKernel() ColorKernel {
	return ColorKernelClass.New()
}

func (c_ ColorKernel) Init() ColorKernel {
	rv := objc.Call[ColorKernel](c_, objc.Sel("init"))
	return rv
}

func (cc _ColorKernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) ColorKernel {
	rv := objc.Call[ColorKernel](cc, objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, objc.Ptr(error))
	return rv
}

// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/2880195-kernelwithfunctionname?language=objc
func ColorKernel_KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) ColorKernel {
	return ColorKernelClass.KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name, data, format, error)
}

// Creates a new image using the kernel and specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorkernel/1438110-applywithextent?language=objc
func (c_ ColorKernel) ApplyWithExtentArguments(extent coregraphics.Rect, args []objc.IObject) Image {
	rv := objc.Call[Image](c_, objc.Sel("applyWithExtent:arguments:"), extent, args)
	return rv
}
