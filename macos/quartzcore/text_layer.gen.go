// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextLayer] class.
var TextLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}

type _TextLayerClass struct {
	objc.Class
}

// An interface definition for the [TextLayer] class.
type ITextLayer interface {
	ILayer
	FontSize() float64
	SetFontSize(value float64)
	IsWrapped() bool
	SetWrapped(value bool)
	AlignmentMode() TextLayerAlignmentMode
	SetAlignmentMode(value TextLayerAlignmentMode)
	ForegroundColor() coregraphics.ColorRef
	SetForegroundColor(value coregraphics.ColorRef)
	String() objc.Object
	SetString(value objc.IObject)
	Font() corefoundation.TypeRef
	SetFont(value corefoundation.TypeRef)
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)
	TruncationMode() TextLayerTruncationMode
	SetTruncationMode(value TextLayerTruncationMode)
}

// A layer that provides simple text layout and rendering of plain or attributed strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer?language=objc
type TextLayer struct {
	Layer
}

func TextLayerFrom(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: LayerFrom(ptr),
	}
}

func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.Call[TextLayer](tc, objc.Sel("alloc"))
	return rv
}

func TextLayer_Alloc() TextLayer {
	return TextLayerClass.Alloc()
}

func (tc _TextLayerClass) New() TextLayer {
	rv := objc.Call[TextLayer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayer() TextLayer {
	return TextLayerClass.New()
}

func (t_ TextLayer) Init() TextLayer {
	rv := objc.Call[TextLayer](t_, objc.Sel("init"))
	return rv
}

func (tc _TextLayerClass) Layer() TextLayer {
	rv := objc.Call[TextLayer](tc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func TextLayer_Layer() TextLayer {
	return TextLayerClass.Layer()
}

func (t_ TextLayer) InitWithLayer(layer objc.IObject) TextLayer {
	rv := objc.Call[TextLayer](t_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewTextLayerWithLayer(layer objc.IObject) TextLayer {
	instance := TextLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (t_ TextLayer) ModelLayer() TextLayer {
	rv := objc.Call[TextLayer](t_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func TextLayer_ModelLayer() TextLayer {
	instance := TextLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (t_ TextLayer) PresentationLayer() TextLayer {
	rv := objc.Call[TextLayer](t_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func TextLayer_PresentationLayer() TextLayer {
	instance := TextLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// The font size used to render the receiver’s text. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515290-fontsize?language=objc
func (t_ TextLayer) FontSize() float64 {
	rv := objc.Call[float64](t_, objc.Sel("fontSize"))
	return rv
}

// The font size used to render the receiver’s text. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515290-fontsize?language=objc
func (t_ TextLayer) SetFontSize(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setFontSize:"), value)
}

// Determines whether the text is wrapped to fit within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515302-wrapped?language=objc
func (t_ TextLayer) IsWrapped() bool {
	rv := objc.Call[bool](t_, objc.Sel("isWrapped"))
	return rv
}

// Determines whether the text is wrapped to fit within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515302-wrapped?language=objc
func (t_ TextLayer) SetWrapped(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setWrapped:"), value)
}

// Determines how individual lines of text are horizontally aligned within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515301-alignmentmode?language=objc
func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode {
	rv := objc.Call[TextLayerAlignmentMode](t_, objc.Sel("alignmentMode"))
	return rv
}

// Determines how individual lines of text are horizontally aligned within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515301-alignmentmode?language=objc
func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode) {
	objc.Call[objc.Void](t_, objc.Sel("setAlignmentMode:"), value)
}

// The color used to render the receiver’s text. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515305-foregroundcolor?language=objc
func (t_ TextLayer) ForegroundColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](t_, objc.Sel("foregroundColor"))
	return rv
}

// The color used to render the receiver’s text. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515305-foregroundcolor?language=objc
func (t_ TextLayer) SetForegroundColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](t_, objc.Sel("setForegroundColor:"), value)
}

// The text to be rendered by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515295-string?language=objc
func (t_ TextLayer) String() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("string"))
	return rv
}

// The text to be rendered by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515295-string?language=objc
func (t_ TextLayer) SetString(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setString:"), value)
}

// The font used to render the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515303-font?language=objc
func (t_ TextLayer) Font() corefoundation.TypeRef {
	rv := objc.Call[corefoundation.TypeRef](t_, objc.Sel("font"))
	return rv
}

// The font used to render the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515303-font?language=objc
func (t_ TextLayer) SetFont(value corefoundation.TypeRef) {
	objc.Call[objc.Void](t_, objc.Sel("setFont:"), value)
}

// Determines whether to allow subpixel quantization for the graphics context used for text rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515300-allowsfontsubpixelquantization?language=objc
func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsFontSubpixelQuantization"))
	return rv
}

// Determines whether to allow subpixel quantization for the graphics context used for text rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515300-allowsfontsubpixelquantization?language=objc
func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsFontSubpixelQuantization:"), value)
}

// Determines how the text is truncated to fit within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515296-truncationmode?language=objc
func (t_ TextLayer) TruncationMode() TextLayerTruncationMode {
	rv := objc.Call[TextLayerTruncationMode](t_, objc.Sel("truncationMode"))
	return rv
}

// Determines how the text is truncated to fit within the receiver’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/1515296-truncationmode?language=objc
func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode) {
	objc.Call[objc.Void](t_, objc.Sel("setTruncationMode:"), value)
}
