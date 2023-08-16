// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UpdateProgressHandlers] class.
var UpdateProgressHandlersClass = _UpdateProgressHandlersClass{objc.GetClass("MLUpdateProgressHandlers")}

type _UpdateProgressHandlersClass struct {
	objc.Class
}

// An interface definition for the [UpdateProgressHandlers] class.
type IUpdateProgressHandlers interface {
	objc.IObject
}

// A collection of closures an update task uses to notify your app of its progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdateprogresshandlers?language=objc
type UpdateProgressHandlers struct {
	objc.Object
}

func UpdateProgressHandlersFrom(ptr unsafe.Pointer) UpdateProgressHandlers {
	return UpdateProgressHandlers{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UpdateProgressHandlers) InitForEventsProgressHandlerCompletionHandler(interestedEvents UpdateProgressEvent, progressHandler func(context UpdateContext), completionHandler func(context UpdateContext)) UpdateProgressHandlers {
	rv := objc.Call[UpdateProgressHandlers](u_, objc.Sel("initForEvents:progressHandler:completionHandler:"), interestedEvents, progressHandler, completionHandler)
	return rv
}

// Creates the collection of closures an update task uses to notify your app of its progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdateprogresshandlers/3294204-initforevents?language=objc
func UpdateProgressHandlers_InitForEventsProgressHandlerCompletionHandler(interestedEvents UpdateProgressEvent, progressHandler func(context UpdateContext), completionHandler func(context UpdateContext)) UpdateProgressHandlers {
	return UpdateProgressHandlersClass.Alloc().InitForEventsProgressHandlerCompletionHandler(interestedEvents, progressHandler, completionHandler)
}

func (uc _UpdateProgressHandlersClass) Alloc() UpdateProgressHandlers {
	rv := objc.Call[UpdateProgressHandlers](uc, objc.Sel("alloc"))
	return rv
}

func UpdateProgressHandlers_Alloc() UpdateProgressHandlers {
	return UpdateProgressHandlersClass.Alloc()
}

func (uc _UpdateProgressHandlersClass) New() UpdateProgressHandlers {
	rv := objc.Call[UpdateProgressHandlers](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUpdateProgressHandlers() UpdateProgressHandlers {
	return UpdateProgressHandlersClass.New()
}

func (u_ UpdateProgressHandlers) Init() UpdateProgressHandlers {
	rv := objc.Call[UpdateProgressHandlers](u_, objc.Sel("init"))
	return rv
}
