// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UpdateTask] class.
var UpdateTaskClass = _UpdateTaskClass{objc.GetClass("MLUpdateTask")}

type _UpdateTaskClass struct {
	objc.Class
}

// An interface definition for the [UpdateTask] class.
type IUpdateTask interface {
	ITask
	ResumeWithParameters(updateParameters foundation.Dictionary)
}

// A task that updates a model with additional training data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatetask?language=objc
type UpdateTask struct {
	Task
}

func UpdateTaskFrom(ptr unsafe.Pointer) UpdateTask {
	return UpdateTask{
		Task: TaskFrom(ptr),
	}
}

func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL foundation.IURL, trainingData PBatchProvider, progressHandlers IUpdateProgressHandlers, error foundation.IError) UpdateTask {
	po1 := objc.WrapAsProtocol("MLBatchProvider", trainingData)
	rv := objc.Call[UpdateTask](uc, objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), objc.Ptr(modelURL), po1, objc.Ptr(progressHandlers), objc.Ptr(error))
	return rv
}

// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatetask/3547163-updatetaskformodelaturl?language=objc
func UpdateTask_UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL foundation.IURL, trainingData PBatchProvider, progressHandlers IUpdateProgressHandlers, error foundation.IError) UpdateTask {
	return UpdateTaskClass.UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL, trainingData, progressHandlers, error)
}

func (uc _UpdateTaskClass) Alloc() UpdateTask {
	rv := objc.Call[UpdateTask](uc, objc.Sel("alloc"))
	return rv
}

func UpdateTask_Alloc() UpdateTask {
	return UpdateTaskClass.Alloc()
}

func (uc _UpdateTaskClass) New() UpdateTask {
	rv := objc.Call[UpdateTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUpdateTask() UpdateTask {
	return UpdateTaskClass.New()
}

func (u_ UpdateTask) Init() UpdateTask {
	rv := objc.Call[UpdateTask](u_, objc.Sel("init"))
	return rv
}

// Resumes a model update with updated parameter values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatetask/3180109-resumewithparameters?language=objc
func (u_ UpdateTask) ResumeWithParameters(updateParameters foundation.Dictionary) {
	objc.Call[objc.Void](u_, objc.Sel("resumeWithParameters:"), updateParameters)
}
