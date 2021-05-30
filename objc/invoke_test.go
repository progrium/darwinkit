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

	i2 := newInvocation(sigID, "getArgumentTypeAtIndex:")
	idx := 1
	// this works, but may not be safe since this is holding onto a pointer to the Go memory
	// unless it makes a copy?
	// or pass everything we need into a single C call that performs the call and the releases it?
	// or do alloc in C?
	// maybe set retainArguments to make sure it makes a copy, and we can
	// then manually "release" when we're done?
	// Based on this it looks like it's copying the value into the target frame:
	// https://github.com/microsoft/WinObjC/blob/ea3f7983803fa02309f974ff878b6c9ecc72c7c4/Frameworks/Foundation/_NSInvocation.x86.mm
	setArgumentAtIndex(i2, unsafe.Pointer(&idx), 2)
	var arg uintptr
	invoke(i2, unsafe.Pointer(&arg))
	t.Errorf("argType: %v", cstring(arg))
}
