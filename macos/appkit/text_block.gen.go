// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}

type _TextBlockClass struct {
	objc.Class
}

type ITextBlock interface {
	objc.IObject
	SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension)
	ValueForDimension(dimension TextBlockDimension) float64
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	SetContentWidthType(val float64, type_ TextBlockValueType)
	SetWidthTypeForLayerEdge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge)
	SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer)
	WidthForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) float64
	WidthValueTypeForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType
	SetBorderColorForEdge(color IColor, edge foundation.RectEdge)
	SetBorderColor(color IColor)
	BorderColorForEdge(edge foundation.RectEdge) Color
	RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	ContentWidth() float64
	ContentWidthValueType() TextBlockValueType
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
}

type TextBlock struct {
	objc.Object
}

func MakeTextBlock(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextBlock) Init() TextBlock {
	rv := objc.CallMethod[TextBlock](t_, objc.GetSelector("init"))
	return rv
}

func TextBlock_Init() TextBlock {
	return TextBlockClass.Alloc().Init()
}

func (tc _TextBlockClass) Alloc() TextBlock {
	rv := objc.CallMethod[TextBlock](tc, objc.GetSelector("alloc"))
	return rv
}

func TextBlock_Alloc() TextBlock {
	return TextBlockClass.Alloc()
}

func (tc _TextBlockClass) New() TextBlock {
	rv := objc.CallMethod[TextBlock](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextBlock() TextBlock {
	return TextBlockClass.New()
}

func TextBlock_New() TextBlock {
	return TextBlockClass.New()
}

func (t_ TextBlock) SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setValue:type:forDimension:"), val, type_, dimension)
}

func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("valueForDimension:"), dimension)
	return rv
}

func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	rv := objc.CallMethod[TextBlockValueType](t_, objc.GetSelector("valueTypeForDimension:"), dimension)
	return rv
}

func (t_ TextBlock) SetContentWidthType(val float64, type_ TextBlockValueType) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setContentWidth:type:"), val, type_)
}

func (t_ TextBlock) SetWidthTypeForLayerEdge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWidth:type:forLayer:edge:"), val, type_, layer, edge)
}

func (t_ TextBlock) SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWidth:type:forLayer:"), val, type_, layer)
}

func (t_ TextBlock) WidthForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("widthForLayer:edge:"), layer, edge)
	return rv
}

func (t_ TextBlock) WidthValueTypeForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType {
	rv := objc.CallMethod[TextBlockValueType](t_, objc.GetSelector("widthValueTypeForLayer:edge:"), layer, edge)
	return rv
}

func (t_ TextBlock) SetBorderColorForEdge(color IColor, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBorderColor:forEdge:"), objc.ExtractPtr(color), edge)
}

func (t_ TextBlock) SetBorderColor(color IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBorderColor:"), objc.ExtractPtr(color))
}

func (t_ TextBlock) BorderColorForEdge(edge foundation.RectEdge) Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("borderColorForEdge:"), edge)
	return rv
}

func (t_ TextBlock) RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("rectForLayoutAtPoint:inRect:textContainer:characterRange:"), startingPoint, rect, objc.ExtractPtr(textContainer), charRange)
	return rv
}

func (t_ TextBlock) BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("boundsRectForContentRect:inRect:textContainer:characterRange:"), contentRect, rect, objc.ExtractPtr(textContainer), charRange)
	return rv
}

func (t_ TextBlock) DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawBackgroundWithFrame:inView:characterRange:layoutManager:"), frameRect, objc.ExtractPtr(controlView), charRange, objc.ExtractPtr(layoutManager))
}

func (t_ TextBlock) ContentWidth() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("contentWidth"))
	return rv
}

func (t_ TextBlock) ContentWidthValueType() TextBlockValueType {
	rv := objc.CallMethod[TextBlockValueType](t_, objc.GetSelector("contentWidthValueType"))
	return rv
}

func (t_ TextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	rv := objc.CallMethod[TextBlockVerticalAlignment](t_, objc.GetSelector("verticalAlignment"))
	return rv
}

func (t_ TextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVerticalAlignment:"), value)
}

func (t_ TextBlock) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TextBlock) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}
