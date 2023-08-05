// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ConstraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}

type _ConstraintLayoutManagerClass struct {
	objc.Class
}

type IConstraintLayoutManager interface {
	objc.IObject
}

type ConstraintLayoutManager struct {
	objc.Object
}

func MakeConstraintLayoutManager(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ConstraintLayoutManagerClass) LayoutManager() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, objc.GetSelector("layoutManager"))
	return rv
}

func ConstraintLayoutManager_LayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.LayoutManager()
}

func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, objc.GetSelector("alloc"))
	return rv
}

func ConstraintLayoutManager_Alloc() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.Alloc()
}

func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewConstraintLayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.New()
}

func ConstraintLayoutManager_New() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.New()
}

func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](c_, objc.GetSelector("init"))
	return rv
}

func ConstraintLayoutManager_Init() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.Alloc().Init()
}
