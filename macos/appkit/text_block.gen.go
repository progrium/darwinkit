// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextBlock] class.
var TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}

type _TextBlockClass struct {
	objc.Class
}

// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	objc.IObject
	ValueForDimension(dimension TextBlockDimension) float64
	RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	BorderColorForEdge(edge foundation.RectEdge) Color
	DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	SetBorderColor(color IColor)
	BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	WidthValueTypeForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType
	SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer)
	WidthForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) float64
	SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension)
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	SetContentWidthType(val float64, type_ TextBlockValueType)
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ContentWidthValueType() TextBlockValueType
	ContentWidth() float64
}

// A block of text laid out in a subregion of the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock?language=objc
type TextBlock struct {
	objc.Object
}

func TextBlockFrom(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextBlock) Init() TextBlock {
	rv := objc.Call[TextBlock](t_, objc.Sel("init"))
	return rv
}

func (tc _TextBlockClass) Alloc() TextBlock {
	rv := objc.Call[TextBlock](tc, objc.Sel("alloc"))
	return rv
}

func TextBlock_Alloc() TextBlock {
	return TextBlockClass.Alloc()
}

func (tc _TextBlockClass) New() TextBlock {
	rv := objc.Call[TextBlock](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextBlock() TextBlock {
	return TextBlockClass.New()
}

// Returns the value of the specified text block dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1526445-valuefordimension?language=objc
func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 {
	rv := objc.Call[float64](t_, objc.Sel("valueForDimension:"), dimension)
	return rv
}

// Returns the rectangle within which glyphs should be laid out for the specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1527965-rectforlayoutatpoint?language=objc
func (t_ TextBlock) RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("rectForLayoutAtPoint:inRect:textContainer:characterRange:"), startingPoint, rect, objc.Ptr(textContainer), charRange)
	return rv
}

// Returns the border color of the specified text block edge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1534711-bordercolorforedge?language=objc
func (t_ TextBlock) BorderColorForEdge(edge foundation.RectEdge) Color {
	rv := objc.Call[Color](t_, objc.Sel("borderColorForEdge:"), edge)
	return rv
}

// Called by the layout manager to draw any colors and other decorations before the text is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1531424-drawbackgroundwithframe?language=objc
func (t_ TextBlock) DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("drawBackgroundWithFrame:inView:characterRange:layoutManager:"), frameRect, objc.Ptr(controlView), charRange, objc.Ptr(layoutManager))
}

// Sets the color of all borders of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1531850-setbordercolor?language=objc
func (t_ TextBlock) SetBorderColor(color IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBorderColor:"), objc.Ptr(color))
}

// Returns the rectangle the text in the block actually occupies, including padding, borders, and margins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1532041-boundsrectforcontentrect?language=objc
func (t_ TextBlock) BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("boundsRectForContentRect:inRect:textContainer:characterRange:"), contentRect, rect, objc.Ptr(textContainer), charRange)
	return rv
}

// Returns the value type of an edge of a specified layer of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1533551-widthvaluetypeforlayer?language=objc
func (t_ TextBlock) WidthValueTypeForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType {
	rv := objc.Call[TextBlockValueType](t_, objc.Sel("widthValueTypeForLayer:edge:"), layer, edge)
	return rv
}

// Sets the width of all edges of a specified layer of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1535325-setwidth?language=objc
func (t_ TextBlock) SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer) {
	objc.Call[objc.Void](t_, objc.Sel("setWidth:type:forLayer:"), val, type_, layer)
}

// Returns the width of an edge of a specified layer of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1533532-widthforlayer?language=objc
func (t_ TextBlock) WidthForLayerEdge(layer TextBlockLayer, edge foundation.RectEdge) float64 {
	rv := objc.Call[float64](t_, objc.Sel("widthForLayer:edge:"), layer, edge)
	return rv
}

// Sets a dimension of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1533000-setvalue?language=objc
func (t_ TextBlock) SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension) {
	objc.Call[objc.Void](t_, objc.Sel("setValue:type:forDimension:"), val, type_, dimension)
}

// Returns the value type of the specified text block dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1530561-valuetypefordimension?language=objc
func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	rv := objc.Call[TextBlockValueType](t_, objc.Sel("valueTypeForDimension:"), dimension)
	return rv
}

// Sets the width of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1528308-setcontentwidth?language=objc
func (t_ TextBlock) SetContentWidthType(val float64, type_ TextBlockValueType) {
	objc.Call[objc.Void](t_, objc.Sel("setContentWidth:type:"), val, type_)
}

// The vertical alignment of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1533097-verticalalignment?language=objc
func (t_ TextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	rv := objc.Call[TextBlockVerticalAlignment](t_, objc.Sel("verticalAlignment"))
	return rv
}

// The vertical alignment of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1533097-verticalalignment?language=objc
func (t_ TextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	objc.Call[objc.Void](t_, objc.Sel("setVerticalAlignment:"), value)
}

// The background color of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1527300-backgroundcolor?language=objc
func (t_ TextBlock) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1527300-backgroundcolor?language=objc
func (t_ TextBlock) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The type of value stored for the text block width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1525975-contentwidthvaluetype?language=objc
func (t_ TextBlock) ContentWidthValueType() TextBlockValueType {
	rv := objc.Call[TextBlockValueType](t_, objc.Sel("contentWidthValueType"))
	return rv
}

// The width of the text block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/1532506-contentwidth?language=objc
func (t_ TextBlock) ContentWidth() float64 {
	rv := objc.Call[float64](t_, objc.Sel("contentWidth"))
	return rv
}
