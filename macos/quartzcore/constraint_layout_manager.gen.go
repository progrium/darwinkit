// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ConstraintLayoutManager] class.
var ConstraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}

type _ConstraintLayoutManagerClass struct {
	objc.Class
}

// An interface definition for the [ConstraintLayoutManager] class.
type IConstraintLayoutManager interface {
	objc.IObject
}

// An object that provides a constraint-based layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraintlayoutmanager?language=objc
type ConstraintLayoutManager struct {
	objc.Object
}

func ConstraintLayoutManagerFrom(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ConstraintLayoutManagerClass) LayoutManager() ConstraintLayoutManager {
	rv := objc.Call[ConstraintLayoutManager](cc, objc.Sel("layoutManager"))
	return rv
}

// Returns the shared layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraintlayoutmanager/1574940-layoutmanager?language=objc
func ConstraintLayoutManager_LayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.LayoutManager()
}

func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := objc.Call[ConstraintLayoutManager](cc, objc.Sel("alloc"))
	return rv
}

func ConstraintLayoutManager_Alloc() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.Alloc()
}

func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := objc.Call[ConstraintLayoutManager](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConstraintLayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.New()
}

func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := objc.Call[ConstraintLayoutManager](c_, objc.Sel("init"))
	return rv
}
