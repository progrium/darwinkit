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

//export monitorReentry
func monitorReentry(event unsafe.Pointer) {
	if event == nil {
		return
	}
	obj := objc.ObjectPtr(uintptr(event))
	obj.Retain()
	monitorCh <- NSEvent_fromRef(obj)
}
