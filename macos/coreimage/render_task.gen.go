// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderTask] class.
var RenderTaskClass = _RenderTaskClass{objc.GetClass("CIRenderTask")}

type _RenderTaskClass struct {
	objc.Class
}

// An interface definition for the [RenderTask] class.
type IRenderTask interface {
	objc.IObject
	WaitUntilCompletedAndReturnError(error foundation.IError) RenderInfo
}

// A single render task issued in conjunction with CIRenderDestination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirendertask?language=objc
type RenderTask struct {
	objc.Object
}

func RenderTaskFrom(ptr unsafe.Pointer) RenderTask {
	return RenderTask{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderTaskClass) Alloc() RenderTask {
	rv := objc.Call[RenderTask](rc, objc.Sel("alloc"))
	return rv
}

func RenderTask_Alloc() RenderTask {
	return RenderTaskClass.Alloc()
}

func (rc _RenderTaskClass) New() RenderTask {
	rv := objc.Call[RenderTask](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderTask() RenderTask {
	return RenderTaskClass.New()
}

func (r_ RenderTask) Init() RenderTask {
	rv := objc.Call[RenderTask](r_, objc.Sel("init"))
	return rv
}

// Waits until the CIRenderTask finishes and returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirendertask/2881294-waituntilcompletedandreturnerror?language=objc
func (r_ RenderTask) WaitUntilCompletedAndReturnError(error foundation.IError) RenderInfo {
	rv := objc.Call[RenderInfo](r_, objc.Sel("waitUntilCompletedAndReturnError:"), objc.Ptr(error))
	return rv
}
