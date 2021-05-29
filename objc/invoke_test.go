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
	sigID := MethodSignatureForSelector(obj.Pointer(), "description")

	sig := ObjectPtr(sigID)
	t.Errorf("sel: %v", sig)

	// inv := GetClass("NSInvocation").Send("invocationWithMethodSignature:", sig)
	// inv.Set("target:", obj)
	// inv.Set("selector:", GetSelector("description"))
	i := NewInvocation(obj.Pointer(), "description")
	inv := ObjectPtr(i)
	t.Errorf("inv: %v", inv)
	// inv.Send("invoke")
	var out uintptr = 1
	t.Errorf("out: %v", out)
	invoke(i, unsafe.Pointer(&out))

	// t.Errorf("i: %v", i)
	// inv.Send("getReturnValue:", uintptr(unsafe.Pointer(&out)))
	t.Errorf("out: %v", out)
	outObj := ObjectPtr(out)
	t.Errorf("outObj: %v", outObj)
	// utf8Bytes := outObj.Send("UTF8String")
	// s := C.GoString((*C.char)(unsafe.Pointer(utf8Bytes.Pointer())))

}
