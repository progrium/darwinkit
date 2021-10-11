package dispatch

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation -framework CoreFoundation -framework WebKit
#include <Foundation/Foundation.h>
#include <dispatch/dispatch.h>

void dispatcher();

void dispatch_async_signal(void *queue) {
	dispatch_async_f((dispatch_queue_t)queue, NULL, (dispatch_function_t)dispatcher);
}

*/
import "C"
import (
	"unsafe"
)

var dispatchQueue chan func()

func init() {
	dispatchQueue = make(chan func(), 1)
}

type Queue unsafe.Pointer

func MainQueue() Queue {
	return Queue(C.dispatch_get_main_queue())
}

func Async(queue Queue, fn func()) {
	dispatchQueue <- fn
	C.dispatch_async_signal(unsafe.Pointer(queue))
}

func Sync(queue Queue, fn func()) {
	done := make(chan struct{})
	Async(queue, func() {
		fn()
		done <- struct{}{}
	})
	<-done
}

type Dispatched struct {
	err chan error
}

func (d Dispatched) Wait() error {
	return <-d.err
}

func Do(queue Queue, fn func() error) Dispatched {
	d := Dispatched{err: make(chan error, 1)}
	Async(queue, func() {
		d.err <- fn()
	})
	return d
}
