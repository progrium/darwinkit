package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSNib struct {
	objc.Object
}

func NSNib_InitWithNibNamed_Bundle(name string, bundle NSBundle) NSNib {
	return NSNib{objc.Get("NSNib").Alloc().Send("initWithNibNamed:bundle:",
		core.String(name), bundle)}
}

func NSNib_InitWithNibData_Bundle(data core.NSData, bundle NSBundle) NSNib {
	return NSNib{objc.Get("NSNib").Alloc().Send("initWithNibData:bundle:",
		data, bundle)}
}

func (nib NSNib) InstantiateWithOwner_TopLevelObjects(owner objc.Object) (core.NSArray, bool) {
	var ptr uintptr
	ok := nib.Send("instantiateWithOwner:topLevelObjects:", owner, &ptr).Bool()
	return core.NSArray{Object: objc.ObjectPtr(ptr)}, ok
}

// Deprecated: use InstantiateWithOwner_TopLevelObjects
func (nib NSNib) InstantiateWithOwner(owner objc.Object) {
	nib.Send("instantiateNibWithOwner:topLevelObjects:", owner, nil)
}
