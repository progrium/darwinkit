// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PatchController] class.
var PatchControllerClass = _PatchControllerClass{objc.GetClass("QCPatchController")}

type _PatchControllerClass struct {
	objc.Class
}

// An interface definition for the [PatchController] class.
type IPatchController interface {
	appkit.IController
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcpatchcontroller?language=objc
type PatchController struct {
	appkit.Controller
}

func PatchControllerFrom(ptr unsafe.Pointer) PatchController {
	return PatchController{
		Controller: appkit.ControllerFrom(ptr),
	}
}

func (pc _PatchControllerClass) Alloc() PatchController {
	rv := objc.Call[PatchController](pc, objc.Sel("alloc"))
	return rv
}

func PatchController_Alloc() PatchController {
	return PatchControllerClass.Alloc()
}

func (pc _PatchControllerClass) New() PatchController {
	rv := objc.Call[PatchController](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPatchController() PatchController {
	return PatchControllerClass.New()
}

func (p_ PatchController) Init() PatchController {
	rv := objc.Call[PatchController](p_, objc.Sel("init"))
	return rv
}
