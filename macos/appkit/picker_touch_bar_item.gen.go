// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PickerTouchBarItem] class.
var PickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}

type _PickerTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [PickerTouchBarItem] class.
type IPickerTouchBarItem interface {
	ITouchBarItem
	SetLabelAtIndex(label string, index int)
	ImageAtIndex(index int) Image
	SetImageAtIndex(image IImage, index int)
	LabelAtIndex(index int) string
	SetEnabledAtIndex(enabled bool, index int)
	IsEnabledAtIndex(index int) bool
	CollapsedRepresentationImage() Image
	SetCollapsedRepresentationImage(value IImage)
	Target() objc.Object
	SetTarget(value objc.IObject)
	SetCustomizationLabel(value string)
	Action() objc.Selector
	SetAction(value objc.Selector)
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)
	SelectionColor() Color
	SetSelectionColor(value IColor)
	SelectionMode() PickerTouchBarItemSelectionMode
	SetSelectionMode(value PickerTouchBarItemSelectionMode)
	NumberOfOptions() int
	SetNumberOfOptions(value int)
	ControlRepresentation() PickerTouchBarItemControlRepresentation
	SetControlRepresentation(value PickerTouchBarItemControlRepresentation)
	IsEnabled() bool
	SetEnabled(value bool)
	SelectedIndex() int
	SetSelectedIndex(value int)
}

// A bar item that provides a picker control with multiple options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem?language=objc
type PickerTouchBarItem struct {
	TouchBarItem
}

func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (pc _PickerTouchBarItemClass) PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier TouchBarItemIdentifier, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.Selector) PickerTouchBarItem {
	rv := objc.Call[PickerTouchBarItem](pc, objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237204-pickertouchbaritemwithidentifier?language=objc
func PickerTouchBarItem_PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier TouchBarItemIdentifier, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.Selector) PickerTouchBarItem {
	return PickerTouchBarItemClass.PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier, labels, selectionMode, target, action)
}

func (pc _PickerTouchBarItemClass) Alloc() PickerTouchBarItem {
	rv := objc.Call[PickerTouchBarItem](pc, objc.Sel("alloc"))
	return rv
}

func PickerTouchBarItem_Alloc() PickerTouchBarItem {
	return PickerTouchBarItemClass.Alloc()
}

func (pc _PickerTouchBarItemClass) New() PickerTouchBarItem {
	rv := objc.Call[PickerTouchBarItem](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPickerTouchBarItem() PickerTouchBarItem {
	return PickerTouchBarItemClass.New()
}

func (p_ PickerTouchBarItem) Init() PickerTouchBarItem {
	rv := objc.Call[PickerTouchBarItem](p_, objc.Sel("init"))
	return rv
}

func (p_ PickerTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) PickerTouchBarItem {
	rv := objc.Call[PickerTouchBarItem](p_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewPickerTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) PickerTouchBarItem {
	instance := PickerTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237209-setlabel?language=objc
func (p_ PickerTouchBarItem) SetLabelAtIndex(label string, index int) {
	objc.Call[objc.Void](p_, objc.Sel("setLabel:atIndex:"), label, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237199-imageatindex?language=objc
func (p_ PickerTouchBarItem) ImageAtIndex(index int) Image {
	rv := objc.Call[Image](p_, objc.Sel("imageAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237208-setimage?language=objc
func (p_ PickerTouchBarItem) SetImageAtIndex(image IImage, index int) {
	objc.Call[objc.Void](p_, objc.Sel("setImage:atIndex:"), objc.Ptr(image), index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237201-labelatindex?language=objc
func (p_ PickerTouchBarItem) LabelAtIndex(index int) string {
	rv := objc.Call[string](p_, objc.Sel("labelAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3353057-setenabled?language=objc
func (p_ PickerTouchBarItem) SetEnabledAtIndex(enabled bool, index int) {
	objc.Call[objc.Void](p_, objc.Sel("setEnabled:atIndex:"), enabled, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3353056-isenabledatindex?language=objc
func (p_ PickerTouchBarItem) IsEnabledAtIndex(index int) bool {
	rv := objc.Call[bool](p_, objc.Sel("isEnabledAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237195-collapsedrepresentationimage?language=objc
func (p_ PickerTouchBarItem) CollapsedRepresentationImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("collapsedRepresentationImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237195-collapsedrepresentationimage?language=objc
func (p_ PickerTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setCollapsedRepresentationImage:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237210-target?language=objc
func (p_ PickerTouchBarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("target"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237210-target?language=objc
func (p_ PickerTouchBarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setTarget:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237198-customizationlabel?language=objc
func (p_ PickerTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setCustomizationLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237194-action?language=objc
func (p_ PickerTouchBarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](p_, objc.Sel("action"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237194-action?language=objc
func (p_ PickerTouchBarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](p_, objc.Sel("setAction:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237196-collapsedrepresentationlabel?language=objc
func (p_ PickerTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Call[string](p_, objc.Sel("collapsedRepresentationLabel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237196-collapsedrepresentationlabel?language=objc
func (p_ PickerTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setCollapsedRepresentationLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237206-selectioncolor?language=objc
func (p_ PickerTouchBarItem) SelectionColor() Color {
	rv := objc.Call[Color](p_, objc.Sel("selectionColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237206-selectioncolor?language=objc
func (p_ PickerTouchBarItem) SetSelectionColor(value IColor) {
	objc.Call[objc.Void](p_, objc.Sel("setSelectionColor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237207-selectionmode?language=objc
func (p_ PickerTouchBarItem) SelectionMode() PickerTouchBarItemSelectionMode {
	rv := objc.Call[PickerTouchBarItemSelectionMode](p_, objc.Sel("selectionMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237207-selectionmode?language=objc
func (p_ PickerTouchBarItem) SetSelectionMode(value PickerTouchBarItemSelectionMode) {
	objc.Call[objc.Void](p_, objc.Sel("setSelectionMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237202-numberofoptions?language=objc
func (p_ PickerTouchBarItem) NumberOfOptions() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfOptions"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237202-numberofoptions?language=objc
func (p_ PickerTouchBarItem) SetNumberOfOptions(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setNumberOfOptions:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237197-controlrepresentation?language=objc
func (p_ PickerTouchBarItem) ControlRepresentation() PickerTouchBarItemControlRepresentation {
	rv := objc.Call[PickerTouchBarItemControlRepresentation](p_, objc.Sel("controlRepresentation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237197-controlrepresentation?language=objc
func (p_ PickerTouchBarItem) SetControlRepresentation(value PickerTouchBarItemControlRepresentation) {
	objc.Call[objc.Void](p_, objc.Sel("setControlRepresentation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3353055-enabled?language=objc
func (p_ PickerTouchBarItem) IsEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3353055-enabled?language=objc
func (p_ PickerTouchBarItem) SetEnabled(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setEnabled:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237205-selectedindex?language=objc
func (p_ PickerTouchBarItem) SelectedIndex() int {
	rv := objc.Call[int](p_, objc.Sel("selectedIndex"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/3237205-selectedindex?language=objc
func (p_ PickerTouchBarItem) SetSelectedIndex(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setSelectedIndex:"), value)
}
