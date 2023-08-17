package dispatch

// #import <stdlib.h>
// #import <stdint.h>
//
// void* Dispatch_Get_Main_Queue();
// void* Dispatch_Get_Global_Queue(intptr_t identifier, uintptr_t flags);
// void Dispatch_Retain(void* queue);
// void Dispatch_Release(void* queue);
// void Dispatch_Async(void* queue, uintptr_t task);
// void Dispatch_Sync(void* queue, uintptr_t task);
import "C"
import (
	"math"
	"runtime/cgo"
	"unsafe"
)

// Queue is for dispatch_queue_t
type Queue struct {
	ptr unsafe.Pointer
}

type QueuePriority int

const (
	QueuePriorityDefault    QueuePriority = 0
	QueuePriorityHigh       QueuePriority = 2
	QueuePriorityLow        QueuePriority = -2
	QueuePriorityBackground QueuePriority = math.MinInt16
)

// Returns the serial dispatch queue associated with the application’s main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452921-dispatch_get_main_queue?language=objc
func MainQueue() Queue {
	p := C.Dispatch_Get_Main_Queue()
	return Queue{p}
}

// Returns a system-defined global concurrent queue with the specified quality-of-service class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452927-dispatch_get_global_queue?language=objc
func GlobalQueue(identifier QueuePriority, flags uintptr) Queue {
	p := C.Dispatch_Get_Global_Queue(C.intptr_t(identifier), C.uintptr_t(flags))
	return Queue{p}
}

func (q Queue) Retain() {
	C.Dispatch_Retain(q.ptr)
}

func (q Queue) Release() {
	C.Dispatch_Release(q.ptr)
}

// Submits a block for asynchronous execution on a dispatch queue and returns immediately.
func (q Queue) DispatchAsync(task func()) {
	id := cgo.NewHandle(task)
	C.Dispatch_Async(q.ptr, C.uintptr_t(id))
}

// Submits a block object for execution and returns after that block finishes executing.
func (q Queue) DispatchSync(task func()) {
	id := cgo.NewHandle(task)
	C.Dispatch_Sync(q.ptr, C.uintptr_t(id))
}

//export runQueueTask
func runQueueTask(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}
