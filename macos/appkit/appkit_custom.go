package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// todo: generate
type ModalSession unsafe.Pointer
type AccessibilityLoadingToken unsafe.Pointer

// todo: generate
const VariableStatusItemLength float64 = -1

// for some reason not in symbols.zip?
type FontWidth float64

// symbols db is missing a few common methods

func (l_ LayoutAnchor) ConstraintEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToAnchor:constant:"), anchor, c)
	return rv
}

// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483640-addsubview?language=objc
func (v_ View) AddSubviewPositionedRelativeTo(view IView, place WindowOrderingMode, otherView IView) {
	objc.Call[objc.Void](v_, objc.Sel("addSubview:positioned:relativeTo:"), view, place, otherView)
}

func (fc _FontClass) FontWithNameSize(fontName string, fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("fontWithName:size:"), fontName, fontSize)
	return rv
}

// Creates a font object for the specified font name and font size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1525977-fontwithname/
func Font_FontWithNameSize(fontName string, fontSize float64) Font {
	return FontClass.FontWithNameSize(fontName, fontSize)
}

func (ic _MenuItemClass) SeparatorItem() MenuItem {
	rv := objc.Call[MenuItem](ic, objc.Sel("separatorItem"))
	return rv
}

// Returns a menu item that is used to separate logical groups of menu commands.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514838-separatoritem
func MenuItem_SeparatorItem() MenuItem {
	return MenuItemClass.SeparatorItem()
}

func (w_ Window) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	rv := objc.Call[Window](w_, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

// Initializes the window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419477-initwithcontentrect?language=objc
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	instance := WindowClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
	instance.Autorelease()
	return instance
}

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
