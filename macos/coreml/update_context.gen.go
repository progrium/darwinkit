// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UpdateContext] class.
var UpdateContextClass = _UpdateContextClass{objc.GetClass("MLUpdateContext")}

type _UpdateContextClass struct {
	objc.Class
}

// An interface definition for the [UpdateContext] class.
type IUpdateContext interface {
	objc.IObject
	Metrics() foundation.Dictionary
	Model() Model
	Parameters() foundation.Dictionary
	Task() UpdateTask
	Event() UpdateProgressEvent
}

// The context an update task provides to your app’s completion and update progress handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext?language=objc
type UpdateContext struct {
	objc.Object
}

func UpdateContextFrom(ptr unsafe.Pointer) UpdateContext {
	return UpdateContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UpdateContextClass) Alloc() UpdateContext {
	rv := objc.Call[UpdateContext](uc, objc.Sel("alloc"))
	return rv
}

func UpdateContext_Alloc() UpdateContext {
	return UpdateContextClass.Alloc()
}

func (uc _UpdateContextClass) New() UpdateContext {
	rv := objc.Call[UpdateContext](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUpdateContext() UpdateContext {
	return UpdateContextClass.New()
}

func (u_ UpdateContext) Init() UpdateContext {
	rv := objc.Call[UpdateContext](u_, objc.Sel("init"))
	return rv
}

// The training metrics of the model for the update task, contained in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/3180095-metrics?language=objc
func (u_ UpdateContext) Metrics() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](u_, objc.Sel("metrics"))
	return rv
}

// The underlying Core ML model stored in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/3180096-model?language=objc
func (u_ UpdateContext) Model() Model {
	rv := objc.Call[Model](u_, objc.Sel("model"))
	return rv
}

// The parameters for the update task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/3180097-parameters?language=objc
func (u_ UpdateContext) Parameters() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](u_, objc.Sel("parameters"))
	return rv
}

// The update task that generated the update context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/3180098-task?language=objc
func (u_ UpdateContext) Task() UpdateTask {
	rv := objc.Call[UpdateTask](u_, objc.Sel("task"))
	return rv
}

// The event type that triggered an update task to notify your app’s completion and update progress handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/3180094-event?language=objc
func (u_ UpdateContext) Event() UpdateProgressEvent {
	rv := objc.Call[UpdateProgressEvent](u_, objc.Sel("event"))
	return rv
}
