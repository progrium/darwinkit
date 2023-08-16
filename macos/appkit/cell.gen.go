// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Cell] class.
var CellClass = _CellClass{objc.GetClass("NSCell")}

type _CellClass struct {
	objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	objc.IObject
	ContinueTrackingAtInView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool
	DraggingImageComponentsWithFrameInView(frame foundation.Rect, view IView) []DraggingImageComponent
	SetNextState()
	StopTrackingAtInViewMouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool)
	FieldEditorForView(controlView IView) TextView
	TakeObjectValueFrom(sender objc.IObject)
	SetUpFieldEditorAttributes(textObj IText) Text
	ResetCursorRectInView(cellFrame foundation.Rect, controlView IView)
	DrawWithFrameInView(cellFrame foundation.Rect, controlView IView)
	MenuForEventInRectOfView(event IEvent, cellFrame foundation.Rect, view IView) Menu
	StartTrackingAtInView(startPoint foundation.Point, controlView IView) bool
	TakeStringValueFrom(sender objc.IObject)
	FocusRingMaskBoundsForFrameInView(cellFrame foundation.Rect, controlView IView) foundation.Rect
	DrawWithExpansionFrameInView(cellFrame foundation.Rect, view IView)
	HitTestForEventInRectOfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult
	TakeFloatValueFrom(sender objc.IObject)
	TakeDoubleValueFrom(sender objc.IObject)
	DrawInteriorWithFrameInView(cellFrame foundation.Rect, controlView IView)
	EditWithFrameInViewEditorDelegateEvent(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent)
	TitleRectForBounds(rect foundation.Rect) foundation.Rect
	CellSizeForBounds(rect foundation.Rect) foundation.Size
	TakeIntegerValueFrom(sender objc.IObject)
	EndEditing(textObj IText)
	DrawingRectForBounds(rect foundation.Rect) foundation.Rect
	DrawFocusRingMaskWithFrameInView(cellFrame foundation.Rect, controlView IView)
	CellAttribute(parameter CellAttribute) int
	CalcDrawInfo(rect foundation.Rect)
	PerformClick(sender objc.IObject)
	SetCellAttributeTo(parameter CellAttribute, value int)
	TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool
	ImageRectForBounds(rect foundation.Rect) foundation.Rect
	HighlightColorWithFrameInView(cellFrame foundation.Rect, controlView IView) Color
	HighlightWithFrameInView(flag bool, cellFrame foundation.Rect, controlView IView)
	GetPeriodicDelayInterval(delay *float64, interval *float64)
	TakeIntValueFrom(sender objc.IObject)
	SendActionOn(mask EventMask) int
	Compare(otherCell objc.IObject) foundation.ComparisonResult
	ExpansionFrameWithFrameInView(cellFrame foundation.Rect, view IView) foundation.Rect
	SelectWithFrameInViewEditorDelegateStartLength(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int)
	IntegerValue() int
	SetIntegerValue(value int)
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	KeyEquivalent() string
	MouseDownFlags() int
	IsHighlighted() bool
	SetHighlighted(value bool)
	IntValue() int
	SetIntValue(value int)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() ControlStateValue
	SetState(value ControlStateValue)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	IsBordered() bool
	SetBordered(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	IsEditable() bool
	SetEditable(value bool)
	IsContinuous() bool
	SetContinuous(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.IAttributedString)
	StringValue() string
	SetStringValue(value string)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.IFormatter)
	NextState() int
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	HasValidObjectValue() bool
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	InteriorBackgroundStyle() BackgroundStyle
	IsSelectable() bool
	SetSelectable(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	IsOpaque() bool
	DoubleValue() float64
	SetDoubleValue(value float64)
	WantsNotificationForMarkedText() bool
	Font() Font
	SetFont(value IFont)
	CellSize() foundation.Size
	Menu() Menu
	SetMenu(value IMenu)
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	Wraps() bool
	SetWraps(value bool)
	AcceptsFirstResponder() bool
	ShowsFirstResponder() bool
	SetShowsFirstResponder(value bool)
	IsScrollable() bool
	SetScrollable(value bool)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	Tag() int
	SetTag(value int)
	Title() string
	SetTitle(value string)
	ControlView() View
	SetControlView(value IView)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	Type() CellType
	SetType(value CellType)
	FloatValue() float64
	SetFloatValue(value float64)
	IsEnabled() bool
	SetEnabled(value bool)
	Image() Image
	SetImage(value IImage)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	SendsActionOnEndEditing() bool
	SetSendsActionOnEndEditing(value bool)
}

// A mechanism for displaying text or images in a view object without the overhead of a full NSView subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell?language=objc
type Cell struct {
	objc.Object
}

func CellFrom(ptr unsafe.Pointer) Cell {
	return Cell{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Cell) InitImageCell(image IImage) Cell {
	rv := objc.Call[Cell](c_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func Cell_InitImageCell(image IImage) Cell {
	return CellClass.Alloc().InitImageCell(image)
}

func (c_ Cell) InitTextCell(string_ string) Cell {
	rv := objc.Call[Cell](c_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func Cell_InitTextCell(string_ string) Cell {
	return CellClass.Alloc().InitTextCell(string_)
}

func (c_ Cell) Init() Cell {
	rv := objc.Call[Cell](c_, objc.Sel("init"))
	return rv
}

func (cc _CellClass) Alloc() Cell {
	rv := objc.Call[Cell](cc, objc.Sel("alloc"))
	return rv
}

func Cell_Alloc() Cell {
	return CellClass.Alloc()
}

func (cc _CellClass) New() Cell {
	rv := objc.Call[Cell](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCell() Cell {
	return CellClass.New()
}

// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535599-continuetracking?language=objc
func (c_ Cell) ContinueTrackingAtInView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool {
	rv := objc.Call[bool](c_, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, objc.Ptr(controlView))
	return rv
}

// Generates dragging image components with the specified frame in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532987-draggingimagecomponentswithframe?language=objc
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame foundation.Rect, view IView) []DraggingImageComponent {
	rv := objc.Call[[]DraggingImageComponent](c_, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, objc.Ptr(view))
	return rv
}

// Changes cell’s state to the next value in the sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533557-setnextstate?language=objc
func (c_ Cell) SetNextState() {
	objc.Call[objc.Void](c_, objc.Sel("setNextState"))
}

// Stops tracking mouse events within the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534650-stoptracking?language=objc
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool) {
	objc.Call[objc.Void](c_, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, objc.Ptr(controlView), flag)
}

// Returns a custom field editor for editing in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532763-fieldeditorforview?language=objc
func (c_ Cell) FieldEditorForView(controlView IView) TextView {
	rv := objc.Call[TextView](c_, objc.Sel("fieldEditorForView:"), objc.Ptr(controlView))
	return rv
}

// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529915-takeobjectvaluefrom?language=objc
func (c_ Cell) TakeObjectValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeObjectValueFrom:"), sender)
}

// Configures the textual and background attributes of the receiver's field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535864-setupfieldeditorattributes?language=objc
func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) Text {
	rv := objc.Call[Text](c_, objc.Sel("setUpFieldEditorAttributes:"), objc.Ptr(textObj))
	return rv
}

// Sets the receiver to show the I-beam cursor while it tracks the mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524496-resetcursorrect?language=objc
func (c_ Cell) ResetCursorRectInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](c_, objc.Sel("resetCursorRect:inView:"), cellFrame, objc.Ptr(controlView))
}

// Draws the receiver’s border and then draws the interior of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535830-drawwithframe?language=objc
func (c_ Cell) DrawWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](c_, objc.Sel("drawWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Returns the menu associated with the cell and related to the specified event and frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535924-menuforevent?language=objc
func (c_ Cell) MenuForEventInRectOfView(event IEvent, cellFrame foundation.Rect, view IView) Menu {
	rv := objc.Call[Menu](c_, objc.Sel("menuForEvent:inRect:ofView:"), objc.Ptr(event), cellFrame, objc.Ptr(view))
	return rv
}

// Begins tracking mouse events within the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526663-starttrackingat?language=objc
func (c_ Cell) StartTrackingAtInView(startPoint foundation.Point, controlView IView) bool {
	rv := objc.Call[bool](c_, objc.Sel("startTrackingAt:inView:"), startPoint, objc.Ptr(controlView))
	return rv
}

// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527148-takestringvaluefrom?language=objc
func (c_ Cell) TakeStringValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeStringValueFrom:"), sender)
}

// Returns the bounds of the focus ring mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534929-focusringmaskboundsforframe?language=objc
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, objc.Ptr(controlView))
	return rv
}

// Instructs the receiver to draw in an expansion frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1528566-drawwithexpansionframe?language=objc
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame foundation.Rect, view IView) {
	objc.Call[objc.Void](c_, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, objc.Ptr(view))
}

// Returns hit testing information for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529601-hittestforevent?language=objc
func (c_ Cell) HitTestForEventInRectOfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult {
	rv := objc.Call[CellHitResult](c_, objc.Sel("hitTestForEvent:inRect:ofView:"), objc.Ptr(event), cellFrame, objc.Ptr(controlView))
	return rv
}

// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535839-takefloatvaluefrom?language=objc
func (c_ Cell) TakeFloatValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeFloatValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526375-takedoublevaluefrom?language=objc
func (c_ Cell) TakeDoubleValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531274-drawinteriorwithframe?language=objc
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](c_, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Begins editing of the receiver’s text using the specified field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533600-editwithframe?language=objc
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent) {
	objc.Call[objc.Void](c_, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, objc.Ptr(controlView), objc.Ptr(textObj), delegate, objc.Ptr(event))
}

// Returns the rectangle in which the receiver draws its title text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531281-titlerectforbounds?language=objc
func (c_ Cell) TitleRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("titleRectForBounds:"), rect)
	return rv
}

// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524792-cellsizeforbounds?language=objc
func (c_ Cell) CellSizeForBounds(rect foundation.Rect) foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("cellSizeForBounds:"), rect)
	return rv
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534969-takeintegervaluefrom?language=objc
func (c_ Cell) TakeIntegerValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Ends the editing of text in the receiver using the specified field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535574-endediting?language=objc
func (c_ Cell) EndEditing(textObj IText) {
	objc.Call[objc.Void](c_, objc.Sel("endEditing:"), objc.Ptr(textObj))
}

// Returns the rectangle within which the receiver draws itself [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526266-drawingrectforbounds?language=objc
func (c_ Cell) DrawingRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("drawingRectForBounds:"), rect)
	return rv
}

// Draws the focus ring for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524608-drawfocusringmaskwithframe?language=objc
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](c_, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Returns the value for the specified cell attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530877-cellattribute?language=objc
func (c_ Cell) CellAttribute(parameter CellAttribute) int {
	rv := objc.Call[int](c_, objc.Sel("cellAttribute:"), parameter)
	return rv
}

// Recalculates the cell geometry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533752-calcdrawinfo?language=objc
func (c_ Cell) CalcDrawInfo(rect foundation.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("calcDrawInfo:"), rect)
}

// Simulates a single mouse click on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534984-performclick?language=objc
func (c_ Cell) PerformClick(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("performClick:"), sender)
}

// Sets the value for the specified cell attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531257-setcellattribute?language=objc
func (c_ Cell) SetCellAttributeTo(parameter CellAttribute, value int) {
	objc.Call[objc.Void](c_, objc.Sel("setCellAttribute:to:"), parameter, value)
}

// Initiates the mouse tracking behavior in a cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533606-trackmouse?language=objc
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool {
	rv := objc.Call[bool](c_, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), objc.Ptr(event), cellFrame, objc.Ptr(controlView), flag)
	return rv
}

// Returns the rectangle in which the receiver draws its image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533408-imagerectforbounds?language=objc
func (c_ Cell) ImageRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("imageRectForBounds:"), rect)
	return rv
}

// Returns the color the receiver uses when drawing the selection highlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534018-highlightcolorwithframe?language=objc
func (c_ Cell) HighlightColorWithFrameInView(cellFrame foundation.Rect, controlView IView) Color {
	rv := objc.Call[Color](c_, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
	return rv
}

// Redraws the receiver with the specified highlight setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533933-highlight?language=objc
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](c_, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, objc.Ptr(controlView))
}

// Returns the initial delay and repeat values for continuous sending of action messages to target objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535611-getperiodicdelay?language=objc
func (c_ Cell) GetPeriodicDelayInterval(delay *float64, interval *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533920-takeintvaluefrom?language=objc
func (c_ Cell) TakeIntValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1528114-sendactionon?language=objc
func (c_ Cell) SendActionOn(mask EventMask) int {
	rv := objc.Call[int](c_, objc.Sel("sendActionOn:"), mask)
	return rv
}

// Compares the string values of the receiver another cell, disregarding case. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527644-compare?language=objc
func (c_ Cell) Compare(otherCell objc.IObject) foundation.ComparisonResult {
	rv := objc.Call[foundation.ComparisonResult](c_, objc.Sel("compare:"), otherCell)
	return rv
}

// Returns the expansion cell frame for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526362-expansionframewithframe?language=objc
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame foundation.Rect, view IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, objc.Ptr(view))
	return rv
}

// Selects the specified text range in the cell's field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527438-selectwithframe?language=objc
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.Call[objc.Void](c_, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, objc.Ptr(controlView), objc.Ptr(textObj), delegate, selStart, selLength)
}

// The cell’s value as an NSInteger type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527783-integervalue?language=objc
func (c_ Cell) IntegerValue() int {
	rv := objc.Call[int](c_, objc.Sel("integerValue"))
	return rv
}

// The cell’s value as an NSInteger type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527783-integervalue?language=objc
func (c_ Cell) SetIntegerValue(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setIntegerValue:"), value)
}

// A Boolean value indicating whether the cell supports three states instead of two. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531578-allowsmixedstate?language=objc
func (c_ Cell) AllowsMixedState() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsMixedState"))
	return rv
}

// A Boolean value indicating whether the cell supports three states instead of two. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531578-allowsmixedstate?language=objc
func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsMixedState:"), value)
}

// The key equivalent associated with clicking the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532218-keyequivalent?language=objc
func (c_ Cell) KeyEquivalent() string {
	rv := objc.Call[string](c_, objc.Sel("keyEquivalent"))
	return rv
}

// The modifier flags for the last (left) mouse-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527798-mousedownflags?language=objc
func (c_ Cell) MouseDownFlags() int {
	rv := objc.Call[int](c_, objc.Sel("mouseDownFlags"))
	return rv
}

// A Boolean value indicating whether the cell has a highlighted appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530864-highlighted?language=objc
func (c_ Cell) IsHighlighted() bool {
	rv := objc.Call[bool](c_, objc.Sel("isHighlighted"))
	return rv
}

// A Boolean value indicating whether the cell has a highlighted appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530864-highlighted?language=objc
func (c_ Cell) SetHighlighted(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHighlighted:"), value)
}

// Returns the default menu for instances of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531847-defaultmenu?language=objc
func (cc _CellClass) DefaultMenu() Menu {
	rv := objc.Call[Menu](cc, objc.Sel("defaultMenu"))
	return rv
}

// Returns the default menu for instances of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531847-defaultmenu?language=objc
func Cell_DefaultMenu() Menu {
	return CellClass.DefaultMenu()
}

// The cell’s value as an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525170-intvalue?language=objc
func (c_ Cell) IntValue() int {
	rv := objc.Call[int](c_, objc.Sel("intValue"))
	return rv
}

// The cell’s value as an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525170-intvalue?language=objc
func (c_ Cell) SetIntValue(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setIntValue:"), value)
}

// The cell’s background style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524686-backgroundstyle?language=objc
func (c_ Cell) BackgroundStyle() BackgroundStyle {
	rv := objc.Call[BackgroundStyle](c_, objc.Sel("backgroundStyle"))
	return rv
}

// The cell’s background style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524686-backgroundstyle?language=objc
func (c_ Cell) SetBackgroundStyle(value BackgroundStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundStyle:"), value)
}

// The size of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530780-controlsize?language=objc
func (c_ Cell) ControlSize() ControlSize {
	rv := objc.Call[ControlSize](c_, objc.Sel("controlSize"))
	return rv
}

// The size of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530780-controlsize?language=objc
func (c_ Cell) SetControlSize(value ControlSize) {
	objc.Call[objc.Void](c_, objc.Sel("setControlSize:"), value)
}

// The object that receives the cell’s action messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535832-target?language=objc
func (c_ Cell) Target() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("target"))
	return rv
}

// The object that receives the cell’s action messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535832-target?language=objc
func (c_ Cell) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setTarget:"), value)
}

// The cell’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527417-state?language=objc
func (c_ Cell) State() ControlStateValue {
	rv := objc.Call[ControlStateValue](c_, objc.Sel("state"))
	return rv
}

// The cell’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527417-state?language=objc
func (c_ Cell) SetState(value ControlStateValue) {
	objc.Call[objc.Void](c_, objc.Sel("setState:"), value)
}

// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525481-usessinglelinemode?language=objc
func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.Call[bool](c_, objc.Sel("usesSingleLineMode"))
	return rv
}

// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525481-usessinglelinemode?language=objc
func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setUsesSingleLineMode:"), value)
}

// A Boolean value indicating whether the cell draws itself outlined with a plain border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525990-bordered?language=objc
func (c_ Cell) IsBordered() bool {
	rv := objc.Call[bool](c_, objc.Sel("isBordered"))
	return rv
}

// A Boolean value indicating whether the cell draws itself outlined with a plain border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525990-bordered?language=objc
func (c_ Cell) SetBordered(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setBordered:"), value)
}

// The action performed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524654-action?language=objc
func (c_ Cell) Action() objc.Selector {
	rv := objc.Call[objc.Selector](c_, objc.Sel("action"))
	return rv
}

// The action performed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524654-action?language=objc
func (c_ Cell) SetAction(value objc.Selector) {
	objc.Call[objc.Void](c_, objc.Sel("setAction:"), value)
}

// A Boolean value indicating whether the cell is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535822-editable?language=objc
func (c_ Cell) IsEditable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEditable"))
	return rv
}

// A Boolean value indicating whether the cell is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535822-editable?language=objc
func (c_ Cell) SetEditable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEditable:"), value)
}

// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529179-continuous?language=objc
func (c_ Cell) IsContinuous() bool {
	rv := objc.Call[bool](c_, objc.Sel("isContinuous"))
	return rv
}

// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529179-continuous?language=objc
func (c_ Cell) SetContinuous(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setContinuous:"), value)
}

// The cell’s value as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534119-attributedstringvalue?language=objc
func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](c_, objc.Sel("attributedStringValue"))
	return rv
}

// The cell’s value as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534119-attributedstringvalue?language=objc
func (c_ Cell) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.Call[objc.Void](c_, objc.Sel("setAttributedStringValue:"), objc.Ptr(value))
}

// The cell’s value as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530915-stringvalue?language=objc
func (c_ Cell) StringValue() string {
	rv := objc.Call[string](c_, objc.Sel("stringValue"))
	return rv
}

// The cell’s value as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530915-stringvalue?language=objc
func (c_ Cell) SetStringValue(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setStringValue:"), value)
}

// The alignment of the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534359-alignment?language=objc
func (c_ Cell) Alignment() TextAlignment {
	rv := objc.Call[TextAlignment](c_, objc.Sel("alignment"))
	return rv
}

// The alignment of the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534359-alignment?language=objc
func (c_ Cell) SetAlignment(value TextAlignment) {
	objc.Call[objc.Void](c_, objc.Sel("setAlignment:"), value)
}

// The type of focus ring to use with the associated view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534132-focusringtype?language=objc
func (c_ Cell) FocusRingType() FocusRingType {
	rv := objc.Call[FocusRingType](c_, objc.Sel("focusRingType"))
	return rv
}

// The type of focus ring to use with the associated view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534132-focusringtype?language=objc
func (c_ Cell) SetFocusRingType(value FocusRingType) {
	objc.Call[objc.Void](c_, objc.Sel("setFocusRingType:"), value)
}

// The cell’s formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531115-formatter?language=objc
func (c_ Cell) Formatter() foundation.Formatter {
	rv := objc.Call[foundation.Formatter](c_, objc.Sel("formatter"))
	return rv
}

// The cell’s formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531115-formatter?language=objc
func (c_ Cell) SetFormatter(value foundation.IFormatter) {
	objc.Call[objc.Void](c_, objc.Sel("setFormatter:"), objc.Ptr(value))
}

// The cell’s next state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531235-nextstate?language=objc
func (c_ Cell) NextState() int {
	rv := objc.Call[int](c_, objc.Sel("nextState"))
	return rv
}

// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535675-allowseditingtextattributes?language=objc
func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsEditingTextAttributes"))
	return rv
}

// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535675-allowseditingtextattributes?language=objc
func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsEditingTextAttributes:"), value)
}

// A Boolean value that indicates whether the cell has a valid object value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534009-hasvalidobjectvalue?language=objc
func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasValidObjectValue"))
	return rv
}

// The layout direction of the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529213-userinterfacelayoutdirection?language=objc
func (c_ Cell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](c_, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}

// The layout direction of the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529213-userinterfacelayoutdirection?language=objc
func (c_ Cell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Call[objc.Void](c_, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// The cell’s interior background style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526141-interiorbackgroundstyle?language=objc
func (c_ Cell) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Call[BackgroundStyle](c_, objc.Sel("interiorBackgroundStyle"))
	return rv
}

// A Boolean value indicating whether the cell’s text can be selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529411-selectable?language=objc
func (c_ Cell) IsSelectable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSelectable"))
	return rv
}

// A Boolean value indicating whether the cell’s text can be selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1529411-selectable?language=objc
func (c_ Cell) SetSelectable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSelectable:"), value)
}

// A Boolean value indicating whether the cell supports the importation of images into its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532380-importsgraphics?language=objc
func (c_ Cell) ImportsGraphics() bool {
	rv := objc.Call[bool](c_, objc.Sel("importsGraphics"))
	return rv
}

// A Boolean value indicating whether the cell supports the importation of images into its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532380-importsgraphics?language=objc
func (c_ Cell) SetImportsGraphics(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setImportsGraphics:"), value)
}

// A Boolean value indicating whether the cell is completely opaque. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531610-opaque?language=objc
func (c_ Cell) IsOpaque() bool {
	rv := objc.Call[bool](c_, objc.Sel("isOpaque"))
	return rv
}

// The cell’s value as a double-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534765-doublevalue?language=objc
func (c_ Cell) DoubleValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("doubleValue"))
	return rv
}

// The cell’s value as a double-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534765-doublevalue?language=objc
func (c_ Cell) SetDoubleValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDoubleValue:"), value)
}

// A Boolean value indicating whether the cell’s field editor should post text change notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535894-wantsnotificationformarkedtext?language=objc
func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.Call[bool](c_, objc.Sel("wantsNotificationForMarkedText"))
	return rv
}

// The font that the cell uses to display text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526710-font?language=objc
func (c_ Cell) Font() Font {
	rv := objc.Call[Font](c_, objc.Sel("font"))
	return rv
}

// The font that the cell uses to display text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526710-font?language=objc
func (c_ Cell) SetFont(value IFont) {
	objc.Call[objc.Void](c_, objc.Sel("setFont:"), objc.Ptr(value))
}

// Returns the default type of focus ring for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525184-defaultfocusringtype?language=objc
func (cc _CellClass) DefaultFocusRingType() FocusRingType {
	rv := objc.Call[FocusRingType](cc, objc.Sel("defaultFocusRingType"))
	return rv
}

// Returns the default type of focus ring for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525184-defaultfocusringtype?language=objc
func Cell_DefaultFocusRingType() FocusRingType {
	return CellClass.DefaultFocusRingType()
}

// The minimum size needed to display the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532056-cellsize?language=objc
func (c_ Cell) CellSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("cellSize"))
	return rv
}

// The cell’s contextual menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530019-menu?language=objc
func (c_ Cell) Menu() Menu {
	rv := objc.Call[Menu](c_, objc.Sel("menu"))
	return rv
}

// The cell’s contextual menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530019-menu?language=objc
func (c_ Cell) SetMenu(value IMenu) {
	objc.Call[objc.Void](c_, objc.Sel("setMenu:"), objc.Ptr(value))
}

// The cell’s value as an Objective-C object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530936-objectvalue?language=objc
func (c_ Cell) ObjectValue() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("objectValue"))
	return rv
}

// The cell’s value as an Objective-C object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530936-objectvalue?language=objc
func (c_ Cell) SetObjectValue(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setObjectValue:"), value)
}

// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527479-wraps?language=objc
func (c_ Cell) Wraps() bool {
	rv := objc.Call[bool](c_, objc.Sel("wraps"))
	return rv
}

// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1527479-wraps?language=objc
func (c_ Cell) SetWraps(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setWraps:"), value)
}

// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530790-preferstrackinguntilmouseup?language=objc
func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Call[bool](cc, objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}

// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530790-preferstrackinguntilmouseup?language=objc
func Cell_PrefersTrackingUntilMouseUp() bool {
	return CellClass.PrefersTrackingUntilMouseUp()
}

// A Boolean value indicating whether the cell accepts first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525710-acceptsfirstresponder?language=objc
func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.Call[bool](c_, objc.Sel("acceptsFirstResponder"))
	return rv
}

// A Boolean value indicating whether the cell provides a visual indication that it is the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532415-showsfirstresponder?language=objc
func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.Call[bool](c_, objc.Sel("showsFirstResponder"))
	return rv
}

// A Boolean value indicating whether the cell provides a visual indication that it is the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532415-showsfirstresponder?language=objc
func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setShowsFirstResponder:"), value)
}

// A Boolean value indicating whether excess text scrolls past the cell’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534125-scrollable?language=objc
func (c_ Cell) IsScrollable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isScrollable"))
	return rv
}

// A Boolean value indicating whether excess text scrolls past the cell’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534125-scrollable?language=objc
func (c_ Cell) SetScrollable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setScrollable:"), value)
}

// The object represented by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533116-representedobject?language=objc
func (c_ Cell) RepresentedObject() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("representedObject"))
	return rv
}

// The object represented by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533116-representedobject?language=objc
func (c_ Cell) SetRepresentedObject(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setRepresentedObject:"), value)
}

// A tag for identifying the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532348-tag?language=objc
func (c_ Cell) Tag() int {
	rv := objc.Call[int](c_, objc.Sel("tag"))
	return rv
}

// A tag for identifying the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1532348-tag?language=objc
func (c_ Cell) SetTag(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setTag:"), value)
}

// The cell’s title text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525561-title?language=objc
func (c_ Cell) Title() string {
	rv := objc.Call[string](c_, objc.Sel("title"))
	return rv
}

// The cell’s title text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525561-title?language=objc
func (c_ Cell) SetTitle(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setTitle:"), value)
}

// The view associated with the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535913-controlview?language=objc
func (c_ Cell) ControlView() View {
	rv := objc.Call[View](c_, objc.Sel("controlView"))
	return rv
}

// The view associated with the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1535913-controlview?language=objc
func (c_ Cell) SetControlView(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setControlView:"), objc.Ptr(value))
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525503-basewritingdirection?language=objc
func (c_ Cell) BaseWritingDirection() WritingDirection {
	rv := objc.Call[WritingDirection](c_, objc.Sel("baseWritingDirection"))
	return rv
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525503-basewritingdirection?language=objc
func (c_ Cell) SetBaseWritingDirection(value WritingDirection) {
	objc.Call[objc.Void](c_, objc.Sel("setBaseWritingDirection:"), value)
}

// The type of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524871-type?language=objc
func (c_ Cell) Type() CellType {
	rv := objc.Call[CellType](c_, objc.Sel("type"))
	return rv
}

// The type of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1524871-type?language=objc
func (c_ Cell) SetType(value CellType) {
	objc.Call[objc.Void](c_, objc.Sel("setType:"), value)
}

// The cell’s value as a single-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534292-floatvalue?language=objc
func (c_ Cell) FloatValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("floatValue"))
	return rv
}

// The cell’s value as a single-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1534292-floatvalue?language=objc
func (c_ Cell) SetFloatValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setFloatValue:"), value)
}

// A Boolean value indicating whether the cell is currently enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533415-enabled?language=objc
func (c_ Cell) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value indicating whether the cell is currently enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533415-enabled?language=objc
func (c_ Cell) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}

// The image displayed by the cell, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526028-image?language=objc
func (c_ Cell) Image() Image {
	rv := objc.Call[Image](c_, objc.Sel("image"))
	return rv
}

// The image displayed by the cell, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526028-image?language=objc
func (c_ Cell) SetImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setImage:"), objc.Ptr(value))
}

// A Boolean value indicating whether the cell refuses the first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525857-refusesfirstresponder?language=objc
func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.Call[bool](c_, objc.Sel("refusesFirstResponder"))
	return rv
}

// A Boolean value indicating whether the cell refuses the first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1525857-refusesfirstresponder?language=objc
func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setRefusesFirstResponder:"), value)
}

// A Boolean value indicating whether the cell assumes responsibility for undo operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531576-allowsundo?language=objc
func (c_ Cell) AllowsUndo() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsUndo"))
	return rv
}

// A Boolean value indicating whether the cell assumes responsibility for undo operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1531576-allowsundo?language=objc
func (c_ Cell) SetAllowsUndo(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsUndo:"), value)
}

// A Boolean value indicating whether the cell has a bezeled border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533376-bezeled?language=objc
func (c_ Cell) IsBezeled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isBezeled"))
	return rv
}

// A Boolean value indicating whether the cell has a bezeled border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533376-bezeled?language=objc
func (c_ Cell) SetBezeled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setBezeled:"), value)
}

// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526736-truncateslastvisibleline?language=objc
func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.Call[bool](c_, objc.Sel("truncatesLastVisibleLine"))
	return rv
}

// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526736-truncateslastvisibleline?language=objc
func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setTruncatesLastVisibleLine:"), value)
}

// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526617-sendsactiononendediting?language=objc
func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.Call[bool](c_, objc.Sel("sendsActionOnEndEditing"))
	return rv
}

// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1526617-sendsactiononendediting?language=objc
func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSendsActionOnEndEditing:"), value)
}
