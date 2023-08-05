// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TitlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}

type _TitlebarAccessoryViewControllerClass struct {
	objc.Class
}

type ITitlebarAccessoryViewController interface {
	IViewController
	FullScreenMinHeight() float64
	SetFullScreenMinHeight(value float64)
	LayoutAttribute() LayoutAttribute
	SetLayoutAttribute(value LayoutAttribute)
	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	IsHidden() bool
	SetHidden(value bool)
}

type TitlebarAccessoryViewController struct {
	ViewController
}

func MakeTitlebarAccessoryViewController(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ViewController: MakeViewController(ptr),
	}
}

func (t_ TitlebarAccessoryViewController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](t_, objc.GetSelector("initWithNibName:bundle:"), nibNameOrNil, objc.ExtractPtr(nibBundleOrNil))
	return rv
}

func TitlebarAccessoryViewController_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](t_, objc.GetSelector("init"))
	return rv
}

func TitlebarAccessoryViewController_Init() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.Alloc().Init()
}

func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](tc, objc.GetSelector("alloc"))
	return rv
}

func TitlebarAccessoryViewController_Alloc() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.Alloc()
}

func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.New()
}

func TitlebarAccessoryViewController_New() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.New()
}

func (t_ TitlebarAccessoryViewController) FullScreenMinHeight() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("fullScreenMinHeight"))
	return rv
}

func (t_ TitlebarAccessoryViewController) SetFullScreenMinHeight(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFullScreenMinHeight:"), value)
}

func (t_ TitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](t_, objc.GetSelector("layoutAttribute"))
	return rv
}

func (t_ TitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutAttribute:"), value)
}

func (t_ TitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("automaticallyAdjustsSize"))
	return rv
}

func (t_ TitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticallyAdjustsSize:"), value)
}

func (t_ TitlebarAccessoryViewController) IsHidden() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isHidden"))
	return rv
}

func (t_ TitlebarAccessoryViewController) SetHidden(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHidden:"), value)
}
