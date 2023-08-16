// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Task] class.
var TaskClass = _TaskClass{objc.GetClass("MLTask")}

type _TaskClass struct {
	objc.Class
}

// An interface definition for the [Task] class.
type ITask interface {
	objc.IObject
	Resume()
	Cancel()
	Error() foundation.Error
	State() TaskState
	TaskIdentifier() string
}

// An abstract base class for machine learning tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask?language=objc
type Task struct {
	objc.Object
}

func TaskFrom(ptr unsafe.Pointer) Task {
	return Task{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TaskClass) Alloc() Task {
	rv := objc.Call[Task](tc, objc.Sel("alloc"))
	return rv
}

func Task_Alloc() Task {
	return TaskClass.Alloc()
}

func (tc _TaskClass) New() Task {
	rv := objc.Call[Task](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTask() Task {
	return TaskClass.New()
}

func (t_ Task) Init() Task {
	rv := objc.Call[Task](t_, objc.Sel("init"))
	return rv
}

// Begins or resumes a machine learning task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask/3180080-resume?language=objc
func (t_ Task) Resume() {
	objc.Call[objc.Void](t_, objc.Sel("resume"))
}

// Cancels a machine learning task before it completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask/3180078-cancel?language=objc
func (t_ Task) Cancel() {
	objc.Call[objc.Void](t_, objc.Sel("cancel"))
}

// The underlying error if the task is in a failed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask/3180079-error?language=objc
func (t_ Task) Error() foundation.Error {
	rv := objc.Call[foundation.Error](t_, objc.Sel("error"))
	return rv
}

// The current state of the machine learning task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask/3180081-state?language=objc
func (t_ Task) State() TaskState {
	rv := objc.Call[TaskState](t_, objc.Sel("state"))
	return rv
}

// A unique name of the task to distinguish it from all other tasks at runtime. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltask/3180082-taskidentifier?language=objc
func (t_ Task) TaskIdentifier() string {
	rv := objc.Call[string](t_, objc.Sel("taskIdentifier"))
	return rv
}
