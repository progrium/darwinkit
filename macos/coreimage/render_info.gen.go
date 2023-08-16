// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderInfo] class.
var RenderInfoClass = _RenderInfoClass{objc.GetClass("CIRenderInfo")}

type _RenderInfoClass struct {
	objc.Class
}

// An interface definition for the [RenderInfo] class.
type IRenderInfo interface {
	objc.IObject
	KernelExecutionTime() foundation.TimeInterval
	PixelsProcessed() int
	PassCount() int
}

// An encapsulation of a render task's timing, passes, and pixels processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderinfo?language=objc
type RenderInfo struct {
	objc.Object
}

func RenderInfoFrom(ptr unsafe.Pointer) RenderInfo {
	return RenderInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderInfoClass) Alloc() RenderInfo {
	rv := objc.Call[RenderInfo](rc, objc.Sel("alloc"))
	return rv
}

func RenderInfo_Alloc() RenderInfo {
	return RenderInfoClass.Alloc()
}

func (rc _RenderInfoClass) New() RenderInfo {
	rv := objc.Call[RenderInfo](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderInfo() RenderInfo {
	return RenderInfoClass.New()
}

func (r_ RenderInfo) Init() RenderInfo {
	rv := objc.Call[RenderInfo](r_, objc.Sel("init"))
	return rv
}

// The amount of time a render spent executing kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderinfo/2875453-kernelexecutiontime?language=objc
func (r_ RenderInfo) KernelExecutionTime() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](r_, objc.Sel("kernelExecutionTime"))
	return rv
}

// The number of pixels the render produced executing kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderinfo/2919725-pixelsprocessed?language=objc
func (r_ RenderInfo) PixelsProcessed() int {
	rv := objc.Call[int](r_, objc.Sel("pixelsProcessed"))
	return rv
}

// The number of passes the render took. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderinfo/2875446-passcount?language=objc
func (r_ RenderInfo) PassCount() int {
	rv := objc.Call[int](r_, objc.Sel("passCount"))
	return rv
}
