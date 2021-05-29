package objc

import (
	"testing"
	"unsafe"
)

func TestInvocation(t *testing.T) {
	rect := NSRect{
		NSPoint{
			X: 0,
			Y: 0,
		},
		NSSize{
			Width:  500,
			Height: 500,
		},
	}

	obj := GetClass("NSValue").Send("valueWithRect:", rect)
	sigID := methodSignatureForSelector(obj.Pointer(), "description")

	sig := ObjectPtr(sigID)
	t.Errorf("sel: %v", sig)

	i := newInvocation(obj.Pointer(), "description")
	var out uintptr = 1
	t.Errorf("out: %v", out)
	invoke(i, unsafe.Pointer(&out))

	t.Errorf("out: %v", out)
	outObj := ObjectPtr(out)
	t.Errorf("outObj: %v", outObj)
}
