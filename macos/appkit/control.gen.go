// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ControlClass = _ControlClass{objc.GetClass("NSControl")}

type _ControlClass struct {
	objc.Class
}

type IControl interface {
	IView
	TakeDoubleValueFrom(sender objc.IObject)
	TakeFloatValueFrom(sender objc.IObject)
	TakeIntValueFrom(sender objc.IObject)
	TakeIntegerValueFrom(sender objc.IObject)
	TakeObjectValueFrom(sender objc.IObject)
	TakeStringValueFrom(sender objc.IObject)
	DrawWithExpansionFrameInView(contentFrame foundation.Rect, view IView)
	ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect
	AbortEditing() bool
	CurrentEditor() Text
	ValidateEditing()
	EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent)
	EndEditing(textObj IText)
	SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int)
	SizeThatFits(size foundation.Size) foundation.Size
	SizeToFit()
	SendActionTo(action objc.Selector, target objc.IObject) bool
	SendActionOn(mask EventMask) int
	PerformClick(sender objc.IObject)
	InvalidateIntrinsicContentSizeForCell(cell ICell)
	IsEnabled() bool
	SetEnabled(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	IntValue() int32
	SetIntValue(value int32)
	IntegerValue() int
	SetIntegerValue(value int)
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	StringValue() string
	SetStringValue(value string)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.IAttributedString)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() Font
	SetFont(value IFont)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.IFormatter)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	IsHighlighted() bool
	SetHighlighted(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	IsContinuous() bool
	SetContinuous(value bool)
	SetTag(value int)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
}

type Control struct {
	View
}

func MakeControl(ptr unsafe.Pointer) Control {
	return Control{
		View: MakeView(ptr),
	}
}

func (c_ Control) InitWithFrame(frameRect foundation.Rect) Control {
	rv := objc.CallMethod[Control](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Control_InitWithFrame(frameRect foundation.Rect) Control {
	return ControlClass.Alloc().InitWithFrame(frameRect)
}

func (c_ Control) Init() Control {
	rv := objc.CallMethod[Control](c_, objc.GetSelector("init"))
	return rv
}

func Control_Init() Control {
	return ControlClass.Alloc().Init()
}

func (cc _ControlClass) Alloc() Control {
	rv := objc.CallMethod[Control](cc, objc.GetSelector("alloc"))
	return rv
}

func Control_Alloc() Control {
	return ControlClass.Alloc()
}

func (cc _ControlClass) New() Control {
	rv := objc.CallMethod[Control](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewControl() Control {
	return ControlClass.New()
}

func Control_New() Control {
	return ControlClass.New()
}

func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeDoubleValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeFloatValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeIntValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeIntegerValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeObjectValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeStringValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) DrawWithExpansionFrameInView(contentFrame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawWithExpansionFrame:inView:"), contentFrame, objc.ExtractPtr(view))
}

func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("expansionFrameWithFrame:"), contentFrame)
	return rv
}

func (c_ Control) AbortEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("abortEditing"))
	return rv
}

func (c_ Control) CurrentEditor() Text {
	rv := objc.CallMethod[Text](c_, objc.GetSelector("currentEditor"))
	return rv
}

func (c_ Control) ValidateEditing() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("validateEditing"))
}

func (c_ Control) EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("editWithFrame:editor:delegate:event:"), rect, objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), objc.ExtractPtr(event))
}

func (c_ Control) EndEditing(textObj IText) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("endEditing:"), objc.ExtractPtr(textObj))
}

func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectWithFrame:editor:delegate:start:length:"), rect, objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), selStart, selLength)
}

func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("sizeThatFits:"), size)
	return rv
}

func (c_ Control) SizeToFit() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("sizeToFit"))
}

func (c_ Control) SendActionTo(action objc.Selector, target objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("sendAction:to:"), action, objc.ExtractPtr(target))
	return rv
}

func (c_ Control) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("sendActionOn:"), mask)
	return rv
}

func (c_ Control) PerformClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("performClick:"), objc.ExtractPtr(sender))
}

func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateIntrinsicContentSizeForCell:"), objc.ExtractPtr(cell))
}

func (c_ Control) IsEnabled() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEnabled"))
	return rv
}

func (c_ Control) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEnabled:"), value)
}

func (c_ Control) DoubleValue() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("doubleValue"))
	return rv
}

func (c_ Control) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDoubleValue:"), value)
}

func (c_ Control) FloatValue() float32 {
	rv := objc.CallMethod[float32](c_, objc.GetSelector("floatValue"))
	return rv
}

func (c_ Control) SetFloatValue(value float32) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFloatValue:"), value)
}

func (c_ Control) IntValue() int32 {
	rv := objc.CallMethod[int32](c_, objc.GetSelector("intValue"))
	return rv
}

func (c_ Control) SetIntValue(value int32) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIntValue:"), value)
}

func (c_ Control) IntegerValue() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("integerValue"))
	return rv
}

func (c_ Control) SetIntegerValue(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIntegerValue:"), value)
}

func (c_ Control) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("objectValue"))
	return rv
}

func (c_ Control) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setObjectValue:"), objc.ExtractPtr(value))
}

func (c_ Control) StringValue() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("stringValue"))
	return rv
}

func (c_ Control) SetStringValue(value string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setStringValue:"), value)
}

func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](c_, objc.GetSelector("attributedStringValue"))
	return rv
}

func (c_ Control) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAttributedStringValue:"), objc.ExtractPtr(value))
}

func (c_ Control) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](c_, objc.GetSelector("alignment"))
	return rv
}

func (c_ Control) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAlignment:"), value)
}

func (c_ Control) Font() Font {
	rv := objc.CallMethod[Font](c_, objc.GetSelector("font"))
	return rv
}

func (c_ Control) SetFont(value IFont) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFont:"), objc.ExtractPtr(value))
}

func (c_ Control) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](c_, objc.GetSelector("lineBreakMode"))
	return rv
}

func (c_ Control) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setLineBreakMode:"), value)
}

func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("usesSingleLineMode"))
	return rv
}

func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setUsesSingleLineMode:"), value)
}

func (c_ Control) Formatter() foundation.Formatter {
	rv := objc.CallMethod[foundation.Formatter](c_, objc.GetSelector("formatter"))
	return rv
}

func (c_ Control) SetFormatter(value foundation.IFormatter) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFormatter:"), objc.ExtractPtr(value))
}

func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](c_, objc.GetSelector("baseWritingDirection"))
	return rv
}

func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBaseWritingDirection:"), value)
}

func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsExpansionToolTips"))
	return rv
}

func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsExpansionToolTips:"), value)
}

func (c_ Control) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](c_, objc.GetSelector("controlSize"))
	return rv
}

func (c_ Control) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setControlSize:"), value)
}

func (c_ Control) IsHighlighted() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isHighlighted"))
	return rv
}

func (c_ Control) SetHighlighted(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHighlighted:"), value)
}

func (c_ Control) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.GetSelector("action"))
	return rv
}

func (c_ Control) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAction:"), value)
}

func (c_ Control) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("target"))
	return rv
}

func (c_ Control) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

func (c_ Control) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isContinuous"))
	return rv
}

func (c_ Control) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContinuous:"), value)
}

func (c_ Control) SetTag(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTag:"), value)
}

func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("refusesFirstResponder"))
	return rv
}

func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setRefusesFirstResponder:"), value)
}

func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("ignoresMultiClick"))
	return rv
}

func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIgnoresMultiClick:"), value)
}
