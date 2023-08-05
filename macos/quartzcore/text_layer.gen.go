// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

var TextLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}

type _TextLayerClass struct {
	objc.Class
}

type ITextLayer interface {
	ILayer
	String() objc.Object
	SetString(value objc.IObject)
	FontSize() float64
	SetFontSize(value float64)
	ForegroundColor() coregraphics.ColorRef
	SetForegroundColor(value coregraphics.ColorRef)
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)
	IsWrapped() bool
	SetWrapped(value bool)
	AlignmentMode() TextLayerAlignmentMode
	SetAlignmentMode(value TextLayerAlignmentMode)
	TruncationMode() TextLayerTruncationMode
	SetTruncationMode(value TextLayerTruncationMode)
}

type TextLayer struct {
	Layer
}

func MakeTextLayer(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: MakeLayer(ptr),
	}
}

func (tc _TextLayerClass) Layer() TextLayer {
	rv := objc.CallMethod[TextLayer](tc, objc.GetSelector("layer"))
	return rv
}

func TextLayer_Layer() TextLayer {
	return TextLayerClass.Layer()
}

func (t_ TextLayer) Init() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, objc.GetSelector("init"))
	return rv
}

func TextLayer_Init() TextLayer {
	return TextLayerClass.Alloc().Init()
}

func (t_ TextLayer) InitWithLayer(layer objc.IObject) TextLayer {
	rv := objc.CallMethod[TextLayer](t_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func TextLayer_InitWithLayer(layer objc.IObject) TextLayer {
	return TextLayerClass.Alloc().InitWithLayer(layer)
}

func (t_ TextLayer) PresentationLayer() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, objc.GetSelector("presentationLayer"))
	return rv
}

func TextLayer_PresentationLayer() TextLayer {
	return TextLayerClass.Alloc().PresentationLayer()
}

func (t_ TextLayer) ModelLayer() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, objc.GetSelector("modelLayer"))
	return rv
}

func TextLayer_ModelLayer() TextLayer {
	return TextLayerClass.Alloc().ModelLayer()
}

func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.CallMethod[TextLayer](tc, objc.GetSelector("alloc"))
	return rv
}

func TextLayer_Alloc() TextLayer {
	return TextLayerClass.Alloc()
}

func (tc _TextLayerClass) New() TextLayer {
	rv := objc.CallMethod[TextLayer](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayer() TextLayer {
	return TextLayerClass.New()
}

func TextLayer_New() TextLayer {
	return TextLayerClass.New()
}

func (t_ TextLayer) String() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("string"))
	return rv
}

func (t_ TextLayer) SetString(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setString:"), objc.ExtractPtr(value))
}

func (t_ TextLayer) FontSize() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("fontSize"))
	return rv
}

func (t_ TextLayer) SetFontSize(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFontSize:"), value)
}

func (t_ TextLayer) ForegroundColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](t_, objc.GetSelector("foregroundColor"))
	return rv
}

func (t_ TextLayer) SetForegroundColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setForegroundColor:"), value)
}

func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsFontSubpixelQuantization"))
	return rv
}

func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsFontSubpixelQuantization:"), value)
}

func (t_ TextLayer) IsWrapped() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isWrapped"))
	return rv
}

func (t_ TextLayer) SetWrapped(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWrapped:"), value)
}

func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode {
	rv := objc.CallMethod[TextLayerAlignmentMode](t_, objc.GetSelector("alignmentMode"))
	return rv
}

func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAlignmentMode:"), value)
}

func (t_ TextLayer) TruncationMode() TextLayerTruncationMode {
	rv := objc.CallMethod[TextLayerTruncationMode](t_, objc.GetSelector("truncationMode"))
	return rv
}

func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTruncationMode:"), value)
}
