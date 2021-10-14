package cocoa

import (
	"unsafe"

	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSNib struct {
	gen_NSNib
}

func NSNib_InitWithNibNamed_Bundle(name string, bundle NSBundle) NSNib {
	return NSNib_fromRef(NSNib_alloc().Send("initWithNibNamed:bundle:", core.String(name), bundle))
}

func NSNib_InitWithNibData_Bundle(data core.NSDataRef, bundle NSBundleRef) NSNib {
	return NSNib_alloc().InitWithNibData_bundle__asNSNib(data, bundle)
}

func (nib NSNib) InstantiateWithOwner_TopLevelObjects(owner objc.Object) (core.NSArray, bool) {
	var ptr uintptr
	ok := nib.Send("instantiateWithOwner:topLevelObjects:", owner, &ptr).Bool()
	return core.NSArray_fromPointer(unsafe.Pointer(ptr)), ok
}

// Deprecated: use InstantiateWithOwner_TopLevelObjects
func (nib NSNib) InstantiateWithOwner(owner objc.Object) {
	nib.Send("instantiateNibWithOwner:topLevelObjects:", owner, nil)
}
