// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlugInViewController] class.
var PlugInViewControllerClass = _PlugInViewControllerClass{objc.GetClass("QCPlugInViewController")}

type _PlugInViewControllerClass struct {
	objc.Class
}

// An interface definition for the [PlugInViewController] class.
type IPlugInViewController interface {
	appkit.IViewController
}

// The QCPlugInViewController class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the QCPlugInViewController class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcpluginviewcontroller?language=objc
type PlugInViewController struct {
	appkit.ViewController
}

func PlugInViewControllerFrom(ptr unsafe.Pointer) PlugInViewController {
	return PlugInViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

func (pc _PlugInViewControllerClass) Alloc() PlugInViewController {
	rv := objc.Call[PlugInViewController](pc, objc.Sel("alloc"))
	return rv
}

func PlugInViewController_Alloc() PlugInViewController {
	return PlugInViewControllerClass.Alloc()
}

func (pc _PlugInViewControllerClass) New() PlugInViewController {
	rv := objc.Call[PlugInViewController](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlugInViewController() PlugInViewController {
	return PlugInViewControllerClass.New()
}

func (p_ PlugInViewController) Init() PlugInViewController {
	rv := objc.Call[PlugInViewController](p_, objc.Sel("init"))
	return rv
}

func (p_ PlugInViewController) InitWithNibNameBundle(nibNameOrNil appkit.NibName, nibBundleOrNil foundation.IBundle) PlugInViewController {
	rv := objc.Call[PlugInViewController](p_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func NewPlugInViewControllerWithNibNameBundle(nibNameOrNil appkit.NibName, nibBundleOrNil foundation.IBundle) PlugInViewController {
	instance := PlugInViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
	instance.Autorelease()
	return instance
}
