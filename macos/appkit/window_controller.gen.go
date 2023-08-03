// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}

type _WindowControllerClass struct {
	objc.Class
}

type IWindowController interface {
	IResponder
	LoadWindow()
	ShowWindow(sender objc.IObject)
	WindowDidLoad()
	WindowWillLoad()
	SetDocumentEdited(dirtyFlag bool)
	Close()
	SynchronizeWindowTitleWithDocumentName()
	WindowTitleForDocumentDisplayName(displayName string) string
	DismissController(sender objc.IObject)
	IsWindowLoaded() bool
	Window() Window
	SetWindow(value IWindow)
	Document() objc.Object
	SetDocument(value objc.IObject)
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)
	Owner() objc.Object
	Storyboard() Storyboard
	WindowNibName() NibName
	WindowNibPath() string
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	WindowFrameAutosaveName() WindowFrameAutosaveName
	SetWindowFrameAutosaveName(value WindowFrameAutosaveName)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
}

type WindowController struct {
	Responder
}

func MakeWindowController(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: MakeResponder(ptr),
	}
}

func (w_ WindowController) InitWithWindow(window IWindow) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("initWithWindow:"), objc.ExtractPtr(window))
	return rv
}

func WindowController_InitWithWindow(window IWindow) WindowController {
	return WindowControllerClass.Alloc().InitWithWindow(window)
}

func (w_ WindowController) InitWithWindowNibName(windowNibName NibName) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("initWithWindowNibName:"), windowNibName)
	return rv
}

func WindowController_InitWithWindowNibName(windowNibName NibName) WindowController {
	return WindowControllerClass.Alloc().InitWithWindowNibName(windowNibName)
}

func (w_ WindowController) InitWithWindowNibNameOwner(windowNibName NibName, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("initWithWindowNibName:owner:"), windowNibName, objc.ExtractPtr(owner))
	return rv
}

func WindowController_InitWithWindowNibNameOwner(windowNibName NibName, owner objc.IObject) WindowController {
	return WindowControllerClass.Alloc().InitWithWindowNibNameOwner(windowNibName, owner)
}

func (w_ WindowController) InitWithWindowNibPathOwner(windowNibPath string, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("initWithWindowNibPath:owner:"), windowNibPath, objc.ExtractPtr(owner))
	return rv
}

func WindowController_InitWithWindowNibPathOwner(windowNibPath string, owner objc.IObject) WindowController {
	return WindowControllerClass.Alloc().InitWithWindowNibPathOwner(windowNibPath, owner)
}

func (w_ WindowController) Init() WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("init"))
	return rv
}

func WindowController_Init() WindowController {
	return WindowControllerClass.Alloc().Init()
}

func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.CallMethod[WindowController](wc, objc.GetSelector("alloc"))
	return rv
}

func WindowController_Alloc() WindowController {
	return WindowControllerClass.Alloc()
}

func (wc _WindowControllerClass) New() WindowController {
	rv := objc.CallMethod[WindowController](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

func WindowController_New() WindowController {
	return WindowControllerClass.New()
}

func (w_ WindowController) LoadWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("loadWindow"))
}

func (w_ WindowController) ShowWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("showWindow:"), objc.ExtractPtr(sender))
}

func (w_ WindowController) WindowDidLoad() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidLoad"))
}

func (w_ WindowController) WindowWillLoad() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillLoad"))
}

func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDocumentEdited:"), dirtyFlag)
}

func (w_ WindowController) Close() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("close"))
}

func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("synchronizeWindowTitleWithDocumentName"))
}

func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("windowTitleForDocumentDisplayName:"), displayName)
	return rv
}

func (w_ WindowController) DismissController(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("dismissController:"), objc.ExtractPtr(sender))
}

func (w_ WindowController) IsWindowLoaded() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isWindowLoaded"))
	return rv
}

func (w_ WindowController) Window() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("window"))
	return rv
}

func (w_ WindowController) SetWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setWindow:"), objc.ExtractPtr(value))
}

func (w_ WindowController) Document() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("document"))
	return rv
}

func (w_ WindowController) SetDocument(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDocument:"), objc.ExtractPtr(value))
}

func (w_ WindowController) ShouldCloseDocument() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("shouldCloseDocument"))
	return rv
}

func (w_ WindowController) SetShouldCloseDocument(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setShouldCloseDocument:"), value)
}

func (w_ WindowController) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("owner"))
	return rv
}

func (w_ WindowController) Storyboard() Storyboard {
	rv := objc.CallMethod[Storyboard](w_, objc.GetSelector("storyboard"))
	return rv
}

func (w_ WindowController) WindowNibName() NibName {
	rv := objc.CallMethod[NibName](w_, objc.GetSelector("windowNibName"))
	return rv
}

func (w_ WindowController) WindowNibPath() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("windowNibPath"))
	return rv
}

func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("shouldCascadeWindows"))
	return rv
}

func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setShouldCascadeWindows:"), value)
}

func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, objc.GetSelector("windowFrameAutosaveName"))
	return rv
}

func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setWindowFrameAutosaveName:"), value)
}

func (w_ WindowController) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, objc.GetSelector("contentViewController"))
	return rv
}

func (w_ WindowController) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentViewController:"), objc.ExtractPtr(value))
}
