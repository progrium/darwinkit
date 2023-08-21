// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Control] class.
var ControlClass = _ControlClass{objc.GetClass("NSControl")}

type _ControlClass struct {
	objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	IView
	SendActionTo(action objc.Selector, target objc.IObject) bool
	SizeThatFits(size foundation.Size) foundation.Size
	TakeObjectValueFrom(sender objc.IObject)
	InvalidateIntrinsicContentSizeForCell(cell ICell)
	ValidateEditing()
	TakeStringValueFrom(sender objc.IObject)
	CurrentEditor() Text
	DrawWithExpansionFrameInView(contentFrame foundation.Rect, view IView)
	AbortEditing() bool
	TakeFloatValueFrom(sender objc.IObject)
	TakeDoubleValueFrom(sender objc.IObject)
	EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent)
	TakeIntegerValueFrom(sender objc.IObject)
	EndEditing(textObj IText)
	PerformClick(sender objc.IObject)
	SizeToFit()
	TakeIntValueFrom(sender objc.IObject)
	SendActionOn(mask EventMask) int
	ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect
	SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int)
	IntegerValue() int
	SetIntegerValue(value int)
	IsHighlighted() bool
	SetHighlighted(value bool)
	IntValue() int
	SetIntValue(value int)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	Target() objc.Object
	SetTarget(value objc.IObject)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	IsContinuous() bool
	SetContinuous(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.IAttributedString)
	StringValue() string
	SetStringValue(value string)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.IFormatter)
	DoubleValue() float64
	SetDoubleValue(value float64)
	Font() Font
	SetFont(value IFont)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	SetTag(value int)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	FloatValue() float64
	SetFloatValue(value float64)
	IsEnabled() bool
	SetEnabled(value bool)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
}

// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol?language=objc
type Control struct {
	View
}

func ControlFrom(ptr unsafe.Pointer) Control {
	return Control{
		View: ViewFrom(ptr),
	}
}

func (c_ Control) InitWithFrame(frameRect foundation.Rect) Control {
	rv := objc.Call[Control](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewControlWithFrame(frameRect foundation.Rect) Control {
	instance := ControlClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (cc _ControlClass) Alloc() Control {
	rv := objc.Call[Control](cc, objc.Sel("alloc"))
	return rv
}

func Control_Alloc() Control {
	return ControlClass.Alloc()
}

func (cc _ControlClass) New() Control {
	rv := objc.Call[Control](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewControl() Control {
	return ControlClass.New()
}

func (c_ Control) Init() Control {
	rv := objc.Call[Control](c_, objc.Sel("init"))
	return rv
}

// Causes the specified action to be sent to the target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428851-sendaction?language=objc
func (c_ Control) SendActionTo(action objc.Selector, target objc.IObject) bool {
	rv := objc.Call[bool](c_, objc.Sel("sendAction:to:"), action, target)
	return rv
}

// Asks the control to calculate and return the size that best fits the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428902-sizethatfits?language=objc
func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("sizeThatFits:"), size)
	return rv
}

// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428853-takeobjectvaluefrom?language=objc
func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeObjectValueFrom:"), sender)
}

// Notifies the control that the intrinsic content size for its cell is no longer valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1526876-invalidateintrinsiccontentsizefo?language=objc
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateIntrinsicContentSizeForCell:"), objc.Ptr(cell))
}

// Validates changes to any user-typed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428855-validateediting?language=objc
func (c_ Control) ValidateEditing() {
	objc.Call[objc.Void](c_, objc.Sel("validateEditing"))
}

// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428912-takestringvaluefrom?language=objc
func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeStringValueFrom:"), sender)
}

// Returns the current field editor for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428980-currenteditor?language=objc
func (c_ Control) CurrentEditor() Text {
	rv := objc.Call[Text](c_, objc.Sel("currentEditor"))
	return rv
}

// Performs custom expansion tool tip drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428895-drawwithexpansionframe?language=objc
func (c_ Control) DrawWithExpansionFrameInView(contentFrame foundation.Rect, view IView) {
	objc.Call[objc.Void](c_, objc.Sel("drawWithExpansionFrame:inView:"), contentFrame, objc.Ptr(view))
}

// Terminates the current editing operation and discards any edited text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428867-abortediting?language=objc
func (c_ Control) AbortEditing() bool {
	rv := objc.Call[bool](c_, objc.Sel("abortEditing"))
	return rv
}

// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428938-takefloatvaluefrom?language=objc
func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeFloatValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428958-takedoublevaluefrom?language=objc
func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Begins editing of the receiver’s text using the specified field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428919-editwithframe?language=objc
func (c_ Control) EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent) {
	objc.Call[objc.Void](c_, objc.Sel("editWithFrame:editor:delegate:event:"), rect, objc.Ptr(textObj), delegate, objc.Ptr(event))
}

// Sets the value of the receiver’s cell to an NSInteger value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428875-takeintegervaluefrom?language=objc
func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Ends the editing of text in the receiver using the specified field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428936-endediting?language=objc
func (c_ Control) EndEditing(textObj IText) {
	objc.Call[objc.Void](c_, objc.Sel("endEditing:"), objc.Ptr(textObj))
}

// Simulates a single mouse click on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428974-performclick?language=objc
func (c_ Control) PerformClick(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("performClick:"), sender)
}

// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428877-sizetofit?language=objc
func (c_ Control) SizeToFit() {
	objc.Call[objc.Void](c_, objc.Sel("sizeToFit"))
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428859-takeintvaluefrom?language=objc
func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428972-sendactionon?language=objc
func (c_ Control) SendActionOn(mask EventMask) int {
	rv := objc.Call[int](c_, objc.Sel("sendActionOn:"), mask)
	return rv
}

// The frame in which a tool tip can be displayed, if needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428932-expansionframewithframe?language=objc
func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("expansionFrameWithFrame:"), contentFrame)
	return rv
}

// Selects the specified text range in the receiver's field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428968-selectwithframe?language=objc
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.Call[objc.Void](c_, objc.Sel("selectWithFrame:editor:delegate:start:length:"), rect, objc.Ptr(textObj), delegate, selStart, selLength)
}

// The value of the receiver’s cell as an NSInteger value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc
func (c_ Control) IntegerValue() int {
	rv := objc.Call[int](c_, objc.Sel("integerValue"))
	return rv
}

// The value of the receiver’s cell as an NSInteger value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc
func (c_ Control) SetIntegerValue(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setIntegerValue:"), value)
}

// A Boolean value that indicates whether the cell is highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc
func (c_ Control) IsHighlighted() bool {
	rv := objc.Call[bool](c_, objc.Sel("isHighlighted"))
	return rv
}

// A Boolean value that indicates whether the cell is highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc
func (c_ Control) SetHighlighted(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHighlighted:"), value)
}

// The value of the receiver’s cell as an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc
func (c_ Control) IntValue() int {
	rv := objc.Call[int](c_, objc.Sel("intValue"))
	return rv
}

// The value of the receiver’s cell as an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc
func (c_ Control) SetIntValue(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setIntValue:"), value)
}

// The size of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428871-controlsize?language=objc
func (c_ Control) ControlSize() ControlSize {
	rv := objc.Call[ControlSize](c_, objc.Sel("controlSize"))
	return rv
}

// The size of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428871-controlsize?language=objc
func (c_ Control) SetControlSize(value ControlSize) {
	objc.Call[objc.Void](c_, objc.Sel("setControlSize:"), value)
}

// The target object that receives action messages from the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc
func (c_ Control) Target() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("target"))
	return rv
}

// The target object that receives action messages from the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc
func (c_ Control) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setTarget:"), value)
}

// A Boolean value that indicates whether the text in the control’s cell uses single line mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc
func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.Call[bool](c_, objc.Sel("usesSingleLineMode"))
	return rv
}

// A Boolean value that indicates whether the text in the control’s cell uses single line mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc
func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setUsesSingleLineMode:"), value)
}

// The default action-message selector associated with the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc
func (c_ Control) Action() objc.Selector {
	rv := objc.Call[objc.Selector](c_, objc.Sel("action"))
	return rv
}

// The default action-message selector associated with the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc
func (c_ Control) SetAction(value objc.Selector) {
	objc.Call[objc.Void](c_, objc.Sel("setAction:"), value)
}

// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc
func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsExpansionToolTips"))
	return rv
}

// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc
func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsExpansionToolTips:"), value)
}

// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc
func (c_ Control) IsContinuous() bool {
	rv := objc.Call[bool](c_, objc.Sel("isContinuous"))
	return rv
}

// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc
func (c_ Control) SetContinuous(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setContinuous:"), value)
}

// The value of the receiver’s cell as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc
func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](c_, objc.Sel("attributedStringValue"))
	return rv
}

// The value of the receiver’s cell as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc
func (c_ Control) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.Call[objc.Void](c_, objc.Sel("setAttributedStringValue:"), objc.Ptr(value))
}

// The value of the receiver’s cell as an NSString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc
func (c_ Control) StringValue() string {
	rv := objc.Call[string](c_, objc.Sel("stringValue"))
	return rv
}

// The value of the receiver’s cell as an NSString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc
func (c_ Control) SetStringValue(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setStringValue:"), value)
}

// The alignment mode of the text in the receiver’s cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428906-alignment?language=objc
func (c_ Control) Alignment() TextAlignment {
	rv := objc.Call[TextAlignment](c_, objc.Sel("alignment"))
	return rv
}

// The alignment mode of the text in the receiver’s cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428906-alignment?language=objc
func (c_ Control) SetAlignment(value TextAlignment) {
	objc.Call[objc.Void](c_, objc.Sel("setAlignment:"), value)
}

// The receiver’s formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428887-formatter?language=objc
func (c_ Control) Formatter() foundation.Formatter {
	rv := objc.Call[foundation.Formatter](c_, objc.Sel("formatter"))
	return rv
}

// The receiver’s formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428887-formatter?language=objc
func (c_ Control) SetFormatter(value foundation.IFormatter) {
	objc.Call[objc.Void](c_, objc.Sel("setFormatter:"), objc.Ptr(value))
}

// The value of the receiver’s cell as a double-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428942-doublevalue?language=objc
func (c_ Control) DoubleValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("doubleValue"))
	return rv
}

// The value of the receiver’s cell as a double-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428942-doublevalue?language=objc
func (c_ Control) SetDoubleValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDoubleValue:"), value)
}

// The font used to draw text in the receiver’s cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc
func (c_ Control) Font() Font {
	rv := objc.Call[Font](c_, objc.Sel("font"))
	return rv
}

// The font used to draw text in the receiver’s cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc
func (c_ Control) SetFont(value IFont) {
	objc.Call[objc.Void](c_, objc.Sel("setFont:"), objc.Ptr(value))
}

// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc
func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.Call[bool](c_, objc.Sel("ignoresMultiClick"))
	return rv
}

// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc
func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setIgnoresMultiClick:"), value)
}

// The value of the receiver’s cell as an Objective-C object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc
func (c_ Control) ObjectValue() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("objectValue"))
	return rv
}

// The value of the receiver’s cell as an Objective-C object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc
func (c_ Control) SetObjectValue(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setObjectValue:"), value)
}

// The tag identifying the receiver (not the tag of the receiver’s cell). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428910-tag?language=objc
func (c_ Control) SetTag(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setTag:"), value)
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428921-basewritingdirection?language=objc
func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := objc.Call[WritingDirection](c_, objc.Sel("baseWritingDirection"))
	return rv
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428921-basewritingdirection?language=objc
func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	objc.Call[objc.Void](c_, objc.Sel("setBaseWritingDirection:"), value)
}

// The value of the receiver’s cell as a single-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428889-floatvalue?language=objc
func (c_ Control) FloatValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("floatValue"))
	return rv
}

// The value of the receiver’s cell as a single-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428889-floatvalue?language=objc
func (c_ Control) SetFloatValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setFloatValue:"), value)
}

// A Boolean value that indicates whether the receiver reacts to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc
func (c_ Control) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the receiver reacts to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc
func (c_ Control) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}

// A Boolean value indicating whether the receiver refuses the first responder role. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc
func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.Call[bool](c_, objc.Sel("refusesFirstResponder"))
	return rv
}

// A Boolean value indicating whether the receiver refuses the first responder role. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc
func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setRefusesFirstResponder:"), value)
}
