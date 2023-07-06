package cocoa

/*
#cgo CFLAGS: -x objective-c

void monitorReentry(void *e);

*/
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var monitorCh chan NSEvent

//export monitorGlobalReentry
func monitorGlobalReentry(event unsafe.Pointer) {
	if event == nil {
		return
	}
	obj := objc.ObjectPtr(uintptr(event))
	obj.Retain()
	monitorCh <- NSEvent_FromRef(obj)
}

//export monitorLocalReentry
func monitorLocalReentry(event unsafe.Pointer) unsafe.Pointer {
	if event == nil {
		return nil
	}
	obj := objc.ObjectPtr(uintptr(event))
	obj.Retain()
	monitorCh <- NSEvent_FromRef(obj)
	return event
}
