// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Alert] class.
var AlertClass = _AlertClass{objc.GetClass("NSAlert")}

type _AlertClass struct {
	objc.Class
}

// An interface definition for the [Alert] class.
type IAlert interface {
	objc.IObject
	RunModal() ModalResponse
	Layout()
	AddButtonWithTitle(title string) Button
	BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	SuppressionButton() Button
	MessageText() string
	SetMessageText(value string)
	InformativeText() string
	SetInformativeText(value string)
	Delegate() AlertDelegateWrapper
	SetDelegate(value PAlertDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ShowsHelp() bool
	SetShowsHelp(value bool)
	Icon() Image
	SetIcon(value IImage)
	ShowsSuppressionButton() bool
	SetShowsSuppressionButton(value bool)
	AccessoryView() View
	SetAccessoryView(value IView)
	AlertStyle() AlertStyle
	SetAlertStyle(value AlertStyle)
	Window() Window
	Buttons() []Button
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
}

// A modal dialog or sheet attached to a document window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert?language=objc
type Alert struct {
	objc.Object
}

func AlertFrom(ptr unsafe.Pointer) Alert {
	return Alert{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AlertClass) Alloc() Alert {
	rv := objc.Call[Alert](ac, objc.Sel("alloc"))
	return rv
}

func Alert_Alloc() Alert {
	return AlertClass.Alloc()
}

func (ac _AlertClass) New() Alert {
	rv := objc.Call[Alert](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAlert() Alert {
	return AlertClass.New()
}

func (a_ Alert) Init() Alert {
	rv := objc.Call[Alert](a_, objc.Sel("init"))
	return rv
}

// Runs the alert as an app-modal dialog and returns the constant that identifies the button clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1535441-runmodal?language=objc
func (a_ Alert) RunModal() ModalResponse {
	rv := objc.Call[ModalResponse](a_, objc.Sel("runModal"))
	return rv
}

// Specifies that the alert must do immediate layout instead of lazily just before display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1526747-layout?language=objc
func (a_ Alert) Layout() {
	objc.Call[objc.Void](a_, objc.Sel("layout"))
}

// Adds a button with a given title to the alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1524532-addbuttonwithtitle?language=objc
func (a_ Alert) AddButtonWithTitle(title string) Button {
	rv := objc.Call[Button](a_, objc.Sel("addButtonWithTitle:"), title)
	return rv
}

// Returns an alert initialized from information in an error object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1531823-alertwitherror?language=objc
func (ac _AlertClass) AlertWithError(error foundation.IError) Alert {
	rv := objc.Call[Alert](ac, objc.Sel("alertWithError:"), objc.Ptr(error))
	return rv
}

// Returns an alert initialized from information in an error object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1531823-alertwitherror?language=objc
func Alert_AlertWithError(error foundation.IError) Alert {
	return AlertClass.AlertWithError(error)
}

// Runs the alert modally as a sheet attached to the specified window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1524296-beginsheetmodalforwindow?language=objc
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.Call[objc.Void](a_, objc.Sel("beginSheetModalForWindow:completionHandler:"), objc.Ptr(sheetWindow), handler)
}

// The alert’s suppression checkbox. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1532209-suppressionbutton?language=objc
func (a_ Alert) SuppressionButton() Button {
	rv := objc.Call[Button](a_, objc.Sel("suppressionButton"))
	return rv
}

// The alert’s message text or title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1532498-messagetext?language=objc
func (a_ Alert) MessageText() string {
	rv := objc.Call[string](a_, objc.Sel("messageText"))
	return rv
}

// The alert’s message text or title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1532498-messagetext?language=objc
func (a_ Alert) SetMessageText(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setMessageText:"), value)
}

// The alert’s informative text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1529629-informativetext?language=objc
func (a_ Alert) InformativeText() string {
	rv := objc.Call[string](a_, objc.Sel("informativeText"))
	return rv
}

// The alert’s informative text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1529629-informativetext?language=objc
func (a_ Alert) SetInformativeText(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setInformativeText:"), value)
}

// The alert’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1534327-delegate?language=objc
func (a_ Alert) Delegate() AlertDelegateWrapper {
	rv := objc.Call[AlertDelegateWrapper](a_, objc.Sel("delegate"))
	return rv
}

// The alert’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1534327-delegate?language=objc
func (a_ Alert) SetDelegate(value PAlertDelegate) {
	po0 := objc.WrapAsProtocol("NSAlertDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), po0)
}

// The alert’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1534327-delegate?language=objc
func (a_ Alert) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Specifies whether the alert has a help button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1535856-showshelp?language=objc
func (a_ Alert) ShowsHelp() bool {
	rv := objc.Call[bool](a_, objc.Sel("showsHelp"))
	return rv
}

// Specifies whether the alert has a help button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1535856-showshelp?language=objc
func (a_ Alert) SetShowsHelp(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setShowsHelp:"), value)
}

// The custom icon displayed in the alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1531688-icon?language=objc
func (a_ Alert) Icon() Image {
	rv := objc.Call[Image](a_, objc.Sel("icon"))
	return rv
}

// The custom icon displayed in the alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1531688-icon?language=objc
func (a_ Alert) SetIcon(value IImage) {
	objc.Call[objc.Void](a_, objc.Sel("setIcon:"), objc.Ptr(value))
}

// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1535196-showssuppressionbutton?language=objc
func (a_ Alert) ShowsSuppressionButton() bool {
	rv := objc.Call[bool](a_, objc.Sel("showsSuppressionButton"))
	return rv
}

// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1535196-showssuppressionbutton?language=objc
func (a_ Alert) SetShowsSuppressionButton(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setShowsSuppressionButton:"), value)
}

// The alert’s accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1530575-accessoryview?language=objc
func (a_ Alert) AccessoryView() View {
	rv := objc.Call[View](a_, objc.Sel("accessoryView"))
	return rv
}

// The alert’s accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1530575-accessoryview?language=objc
func (a_ Alert) SetAccessoryView(value IView) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// Indicates the alert’s severity level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1528506-alertstyle?language=objc
func (a_ Alert) AlertStyle() AlertStyle {
	rv := objc.Call[AlertStyle](a_, objc.Sel("alertStyle"))
	return rv
}

// Indicates the alert’s severity level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1528506-alertstyle?language=objc
func (a_ Alert) SetAlertStyle(value AlertStyle) {
	objc.Call[objc.Void](a_, objc.Sel("setAlertStyle:"), value)
}

// The app-modal panel or document-modal sheet that corresponds to the alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1526566-window?language=objc
func (a_ Alert) Window() Window {
	rv := objc.Call[Window](a_, objc.Sel("window"))
	return rv
}

// The array of response buttons for the alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1532992-buttons?language=objc
func (a_ Alert) Buttons() []Button {
	rv := objc.Call[[]Button](a_, objc.Sel("buttons"))
	return rv
}

// The alert’s HTML help anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1534314-helpanchor?language=objc
func (a_ Alert) HelpAnchor() HelpAnchorName {
	rv := objc.Call[HelpAnchorName](a_, objc.Sel("helpAnchor"))
	return rv
}

// The alert’s HTML help anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalert/1534314-helpanchor?language=objc
func (a_ Alert) SetHelpAnchor(value HelpAnchorName) {
	objc.Call[objc.Void](a_, objc.Sel("setHelpAnchor:"), value)
}
