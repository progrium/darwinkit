// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WindowController] class.
var WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}

type _WindowControllerClass struct {
	objc.Class
}

// An interface definition for the [WindowController] class.
type IWindowController interface {
	IResponder
	SetDocumentEdited(dirtyFlag bool)
	WindowDidLoad()
	ShowWindow(sender objc.IObject) objc.Object
	DismissController(sender objc.IObject) objc.Object
	Close()
	WindowTitleForDocumentDisplayName(displayName string) string
	SynchronizeWindowTitleWithDocumentName()
	LoadWindow()
	WindowWillLoad()
	WindowFrameAutosaveName() WindowFrameAutosaveName
	SetWindowFrameAutosaveName(value WindowFrameAutosaveName)
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)
	Owner() objc.Object
	Document() objc.Object
	SetDocument(value objc.IObject)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	IsWindowLoaded() bool
	Storyboard() Storyboard
	WindowNibName() NibName
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	Window() Window
	SetWindow(value IWindow)
	WindowNibPath() string
}

// A controller that manages a window, usually a window stored in a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller?language=objc
type WindowController struct {
	Responder
}

func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: ResponderFrom(ptr),
	}
}

func (w_ WindowController) InitWithWindow(window IWindow) WindowController {
	rv := objc.Call[WindowController](w_, objc.Sel("initWithWindow:"), objc.Ptr(window))
	return rv
}

// Returns a window controller initialized with a given window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1533442-initwithwindow?language=objc
func NewWindowControllerWithWindow(window IWindow) WindowController {
	instance := WindowControllerClass.Alloc().InitWithWindow(window)
	instance.Autorelease()
	return instance
}

func (w_ WindowController) InitWithWindowNibPathOwner(windowNibPath string, owner objc.IObject) WindowController {
	rv := objc.Call[WindowController](w_, objc.Sel("initWithWindowNibPath:owner:"), windowNibPath, owner)
	return rv
}

// Returns a window controller initialized with a nib file at an absolute path and a specified owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1532584-initwithwindownibpath?language=objc
func NewWindowControllerWithWindowNibPathOwner(windowNibPath string, owner objc.IObject) WindowController {
	instance := WindowControllerClass.Alloc().InitWithWindowNibPathOwner(windowNibPath, owner)
	instance.Autorelease()
	return instance
}

func (w_ WindowController) InitWithWindowNibNameOwner(windowNibName NibName, owner objc.IObject) WindowController {
	rv := objc.Call[WindowController](w_, objc.Sel("initWithWindowNibName:owner:"), windowNibName, owner)
	return rv
}

// Returns a window controller initialized with a nib file and a specified owner for that nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1535239-initwithwindownibname?language=objc
func NewWindowControllerWithWindowNibNameOwner(windowNibName NibName, owner objc.IObject) WindowController {
	instance := WindowControllerClass.Alloc().InitWithWindowNibNameOwner(windowNibName, owner)
	instance.Autorelease()
	return instance
}

func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.Call[WindowController](wc, objc.Sel("alloc"))
	return rv
}

func WindowController_Alloc() WindowController {
	return WindowControllerClass.Alloc()
}

func (wc _WindowControllerClass) New() WindowController {
	rv := objc.Call[WindowController](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

func (w_ WindowController) Init() WindowController {
	rv := objc.Call[WindowController](w_, objc.Sel("init"))
	return rv
}

// Sets the document edited flag for the window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1534261-setdocumentedited?language=objc
func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setDocumentEdited:"), dirtyFlag)
}

// Sent after the window owned by the receiver has been loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1534205-windowdidload?language=objc
func (w_ WindowController) WindowDidLoad() {
	objc.Call[objc.Void](w_, objc.Sel("windowDidLoad"))
}

// Displays the window associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1534037-showwindow?language=objc
func (w_ WindowController) ShowWindow(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("showWindow:"), sender)
	return rv
}

// Dismisses the window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1531963-dismisscontroller?language=objc
func (w_ WindowController) DismissController(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("dismissController:"), sender)
	return rv
}

// Closes the window if it was loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1535390-close?language=objc
func (w_ WindowController) Close() {
	objc.Call[objc.Void](w_, objc.Sel("close"))
}

// Returns the window title to be used for a given document display name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1528112-windowtitlefordocumentdisplaynam?language=objc
func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := objc.Call[string](w_, objc.Sel("windowTitleForDocumentDisplayName:"), displayName)
	return rv
}

// Synchronizes the displayed window title and the represented filename with the information in the associated document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1524667-synchronizewindowtitlewithdocume?language=objc
func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.Call[objc.Void](w_, objc.Sel("synchronizeWindowTitleWithDocumentName"))
}

// Loads the receiver’s window from the nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1535137-loadwindow?language=objc
func (w_ WindowController) LoadWindow() {
	objc.Call[objc.Void](w_, objc.Sel("loadWindow"))
}

// Sent before the window owned by the receiver is loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1524557-windowwillload?language=objc
func (w_ WindowController) WindowWillLoad() {
	objc.Call[objc.Void](w_, objc.Sel("windowWillLoad"))
}

// The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1528616-windowframeautosavename?language=objc
func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.Call[WindowFrameAutosaveName](w_, objc.Sel("windowFrameAutosaveName"))
	return rv
}

// The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1528616-windowframeautosavename?language=objc
func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	objc.Call[objc.Void](w_, objc.Sel("setWindowFrameAutosaveName:"), value)
}

// A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1526933-shouldclosedocument?language=objc
func (w_ WindowController) ShouldCloseDocument() bool {
	rv := objc.Call[bool](w_, objc.Sel("shouldCloseDocument"))
	return rv
}

// A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1526933-shouldclosedocument?language=objc
func (w_ WindowController) SetShouldCloseDocument(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setShouldCloseDocument:"), value)
}

// The owner of the nib file containing the window managed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1532707-owner?language=objc
func (w_ WindowController) Owner() objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("owner"))
	return rv
}

// The document associated with the window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1534220-document?language=objc
func (w_ WindowController) Document() objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("document"))
	return rv
}

// The document associated with the window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1534220-document?language=objc
func (w_ WindowController) SetDocument(value objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setDocument:"), value)
}

// The view controller for the window’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1532552-contentviewcontroller?language=objc
func (w_ WindowController) ContentViewController() ViewController {
	rv := objc.Call[ViewController](w_, objc.Sel("contentViewController"))
	return rv
}

// The view controller for the window’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1532552-contentviewcontroller?language=objc
func (w_ WindowController) SetContentViewController(value IViewController) {
	objc.Call[objc.Void](w_, objc.Sel("setContentViewController:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1527496-windowloaded?language=objc
func (w_ WindowController) IsWindowLoaded() bool {
	rv := objc.Call[bool](w_, objc.Sel("isWindowLoaded"))
	return rv
}

// The storyboard file from which the window controller was loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1527268-storyboard?language=objc
func (w_ WindowController) Storyboard() Storyboard {
	rv := objc.Call[Storyboard](w_, objc.Sel("storyboard"))
	return rv
}

// The name of the nib file that stores the window associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1527084-windownibname?language=objc
func (w_ WindowController) WindowNibName() NibName {
	rv := objc.Call[NibName](w_, objc.Sel("windowNibName"))
	return rv
}

// A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1528177-shouldcascadewindows?language=objc
func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := objc.Call[bool](w_, objc.Sel("shouldCascadeWindows"))
	return rv
}

// A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1528177-shouldcascadewindows?language=objc
func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setShouldCascadeWindows:"), value)
}

// The window owned by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1535593-window?language=objc
func (w_ WindowController) Window() Window {
	rv := objc.Call[Window](w_, objc.Sel("window"))
	return rv
}

// The window owned by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1535593-window?language=objc
func (w_ WindowController) SetWindow(value IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("setWindow:"), objc.Ptr(value))
}

// The full path of the nib file that stores the window associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcontroller/1524719-windownibpath?language=objc
func (w_ WindowController) WindowNibPath() string {
	rv := objc.Call[string](w_, objc.Sel("windowNibPath"))
	return rv
}
