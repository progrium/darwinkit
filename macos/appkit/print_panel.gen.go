// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PrintPanel] class.
var PrintPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}

type _PrintPanelClass struct {
	objc.Class
}

// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	objc.IObject
	RemoveAccessoryController(accessoryController IViewController)
	RunModal() int
	AddAccessoryController(accessoryController IViewController)
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	DefaultButtonTitle() string
	SetDefaultButtonTitle(defaultButtonTitle string)
	PrintInfo() PrintInfo
	AccessoryControllers() []ViewController
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	JobStyleHint() PrintPanelJobStyleHint
	SetJobStyleHint(value PrintPanelJobStyleHint)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
}

// The Print panel that queries the user for information about a print job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel?language=objc
type PrintPanel struct {
	objc.Object
}

func PrintPanelFrom(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.Call[PrintPanel](pc, objc.Sel("alloc"))
	return rv
}

func PrintPanel_Alloc() PrintPanel {
	return PrintPanelClass.Alloc()
}

func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.Call[PrintPanel](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.Call[PrintPanel](p_, objc.Sel("init"))
	return rv
}

// Removes the specified controller and accessory view from the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490527-removeaccessorycontroller?language=objc
func (p_ PrintPanel) RemoveAccessoryController(accessoryController IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("removeAccessoryController:"), objc.Ptr(accessoryController))
}

// Displays the Print panel and begins the modal loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490523-runmodal?language=objc
func (p_ PrintPanel) RunModal() int {
	rv := objc.Call[int](p_, objc.Sel("runModal"))
	return rv
}

// Adds a custom controller to the Print panel to manage an accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490510-addaccessorycontroller?language=objc
func (p_ PrintPanel) AddAccessoryController(accessoryController IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("addAccessoryController:"), objc.Ptr(accessoryController))
}

// Displays the Print panel and runs the modal loop using the specified printing information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490539-runmodalwithprintinfo?language=objc
func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.Call[int](p_, objc.Sel("runModalWithPrintInfo:"), objc.Ptr(printInfo))
	return rv
}

// Returns a new print panel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490531-printpanel?language=objc
func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := objc.Call[PrintPanel](pc, objc.Sel("printPanel"))
	return rv
}

// Returns a new print panel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490531-printpanel?language=objc
func PrintPanel_PrintPanel() PrintPanel {
	return PrintPanelClass.PrintPanel()
}

// Returns the title of the Print panel’s default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490548-defaultbuttontitle?language=objc
func (p_ PrintPanel) DefaultButtonTitle() string {
	rv := objc.Call[string](p_, objc.Sel("defaultButtonTitle"))
	return rv
}

// Sets the title of the Print panel’s default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490525-setdefaultbuttontitle?language=objc
func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.Call[objc.Void](p_, objc.Sel("setDefaultButtonTitle:"), defaultButtonTitle)
}

// The information associated with the running Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490513-printinfo?language=objc
func (p_ PrintPanel) PrintInfo() PrintInfo {
	rv := objc.Call[PrintInfo](p_, objc.Sel("printInfo"))
	return rv
}

// The array of controller objects that manage the Print panel’s accessory views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490543-accessorycontrollers?language=objc
func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.Call[[]ViewController](p_, objc.Sel("accessoryControllers"))
	return rv
}

// The current configuration options for the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490542-options?language=objc
func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.Call[PrintPanelOptions](p_, objc.Sel("options"))
	return rv
}

// The current configuration options for the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490542-options?language=objc
func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.Call[objc.Void](p_, objc.Sel("setOptions:"), value)
}

// The type of settings that the print panel displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490529-jobstylehint?language=objc
func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := objc.Call[PrintPanelJobStyleHint](p_, objc.Sel("jobStyleHint"))
	return rv
}

// The type of settings that the print panel displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490529-jobstylehint?language=objc
func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	objc.Call[objc.Void](p_, objc.Sel("setJobStyleHint:"), value)
}

// The HTML help anchor associated with the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490537-helpanchor?language=objc
func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := objc.Call[HelpAnchorName](p_, objc.Sel("helpAnchor"))
	return rv
}

// The HTML help anchor associated with the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanel/1490537-helpanchor?language=objc
func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName) {
	objc.Call[objc.Void](p_, objc.Sel("setHelpAnchor:"), value)
}
