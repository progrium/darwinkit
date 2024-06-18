package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// todo: generate?
type AccessibilityLoadingToken unsafe.Pointer

const VariableStatusItemLength float64 = -1

// for some reason not in symbols.zip?
type FontWidth float64

// currently this is deprecated but is still needed pre-mac14
// todo: remove when deprecated are added back (but also under build tag per version)

func (a_ Application) ActivateIgnoringOtherApps(flag bool) {
	objc.Call[objc.Void](a_, objc.Sel("activateIgnoringOtherApps:"), flag)
}

// helpers?

func DisableAutoresizingTranslate[T IView](v T) T {
	v.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return v
}

// NewMenuItem create a new menu item, with selector
func NewMenuItemWithSelector(title string, charCode string, selector objc.Selector) MenuItem {
	return NewMenuItemWithTitleActionKeyEquivalent(title, selector, charCode)
}

// NewMenuItemWithAction create a new menu item with action
func NewMenuItemWithAction(title string, charCode string, handler action.Handler) MenuItem {
	item := NewMenuItemWithTitleActionKeyEquivalent(title, objc.Selector{}, charCode)
	action.Set(item, handler)
	return item
}

// NewSubMenuItem create a menu item that hold a sub menu
func NewSubMenuItem(menu IMenu) MenuItem {
	item := NewMenuItemWithTitleActionKeyEquivalent("", objc.Selector{}, "")
	item.SetSubmenu(menu)
	return item
}

// NewButton return a new common used Button
func NewButtonWithTitle(title string) Button {
	return ButtonClass.ButtonWithTitleTargetAction(title, objc.Object{}, objc.Selector{})
}

// NewButtonWithImage return a new common used Button with image
func NewButtonWithImage(image Image) Button {
	return ButtonClass.ButtonWithImageTargetAction(image, objc.Object{}, objc.Selector{})
}

// NewCheckBox return a new common used checkbox Button
func NewCheckBox(title string) Button {
	return ButtonClass.CheckboxWithTitleTargetAction(title, objc.Object{}, objc.Selector{})
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(title string) Button {
	return ButtonClass.RadioButtonWithTitleTargetAction(title, objc.Object{}, objc.Selector{})
}

// NewPushButton return a button that switches between on and off states with each click.
func NewPushButton(title string) Button {
	btn := NewButton()
	btn.SetButtonType(ButtonTypePushOnPushOff)
	btn.SetTitle(title)
	return btn
}

// NewVerticalStackView return a new vertical StackView
func NewVerticalStackView() StackView {
	sv := NewStackView()
	sv.SetOrientation(UserInterfaceLayoutOrientationVertical)
	return sv
}

// NewHorizontalStackView return a new horizontal StackView
func NewHorizontalStackView() StackView {
	sv := NewStackView()
	sv.SetOrientation(UserInterfaceLayoutOrientationHorizontal)
	return sv
}

// NewLabel create a text field, which looks like a Label
func NewLabel(title string) TextField {
	return TextFieldClass.LabelWithString(title)
}

// ScrollableTextView is ScrollView that contains a TextView
type ScrollableTextView struct {
	ScrollView
}

// ContentTextView return the inner TextView
func (t *ScrollableTextView) ContentTextView() TextView {
	return TextViewFrom(t.DocumentView().Ptr())
}

// NewScrollableTextView create and return new scrollable text view.
func NewScrollableTextView() ScrollableTextView {
	stv := TextView_ScrollableTextView()
	stv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := TextViewFrom(stv.DocumentView().Ptr())
	tv.SetAutomaticQuoteSubstitutionEnabled(false)
	tv.SetAutomaticDashSubstitutionEnabled(false)
	tv.SetAllowsUndo(true)
	return ScrollableTextView{ScrollView: stv}
}

// NewWindowWithSize create a common window with close/minimize buttons
func NewWindowWithSize(width, height float64) Window {
	return NewWindowWithContentRectStyleMaskBackingDefer(
		foundation.Rect{Size: foundation.Size{Width: width, Height: height}},
		WindowStyleMaskTitled|WindowStyleMaskClosable|WindowStyleMaskResizable|WindowStyleMaskMiniaturizable,
		BackingStoreBuffered,
		false,
	)
}

// NewWindowWithSizeAndStyle create a common window with styles
func NewWindowWithSizeAndStyle(width, height float64, style WindowStyleMask) Window {
	return NewWindowWithContentRectStyleMaskBackingDefer(
		foundation.Rect{Size: foundation.Size{Width: width, Height: height}},
		style,
		BackingStoreBuffered,
		false,
	)
}
