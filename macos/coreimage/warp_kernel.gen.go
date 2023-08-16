// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WarpKernel] class.
var WarpKernelClass = _WarpKernelClass{objc.GetClass("CIWarpKernel")}

type _WarpKernelClass struct {
	objc.Class
}

// An interface definition for the [WarpKernel] class.
type IWarpKernel interface {
	IKernel
	ApplyWithExtentRoiCallbackInputImageArguments(extent coregraphics.Rect, callback KernelROICallback, image IImage, args []objc.IObject) Image
}

// A GPU-based image processing routine that processes only the geometry information in an image, used to create custom Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwarpkernel?language=objc
type WarpKernel struct {
	Kernel
}

func WarpKernelFrom(ptr unsafe.Pointer) WarpKernel {
	return WarpKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (wc _WarpKernelClass) Alloc() WarpKernel {
	rv := objc.Call[WarpKernel](wc, objc.Sel("alloc"))
	return rv
}

func WarpKernel_Alloc() WarpKernel {
	return WarpKernelClass.Alloc()
}

func (wc _WarpKernelClass) New() WarpKernel {
	rv := objc.Call[WarpKernel](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWarpKernel() WarpKernel {
	return WarpKernelClass.New()
}

func (w_ WarpKernel) Init() WarpKernel {
	rv := objc.Call[WarpKernel](w_, objc.Sel("init"))
	return rv
}

func (wc _WarpKernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) WarpKernel {
	rv := objc.Call[WarpKernel](wc, objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, objc.Ptr(error))
	return rv
}

// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/2880195-kernelwithfunctionname?language=objc
func WarpKernel_KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) WarpKernel {
	return WarpKernelClass.KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name, data, format, error)
}

// Creates a new image using the kernel and the specified input image and arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwarpkernel/1437798-applywithextent?language=objc
func (w_ WarpKernel) ApplyWithExtentRoiCallbackInputImageArguments(extent coregraphics.Rect, callback KernelROICallback, image IImage, args []objc.IObject) Image {
	rv := objc.Call[Image](w_, objc.Sel("applyWithExtent:roiCallback:inputImage:arguments:"), extent, callback, objc.Ptr(image), args)
	return rv
}
