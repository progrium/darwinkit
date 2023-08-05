// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}

type _RulerMarkerClass struct {
	objc.Class
}

type IRulerMarker interface {
	objc.IObject
	DrawRect(rect foundation.Rect)
	TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool
	Ruler() RulerView
	Image() Image
	SetImage(value IImage)
	ImageOrigin() foundation.Point
	SetImageOrigin(value foundation.Point)
	ImageRectInRuler() foundation.Rect
	ThicknessRequiredInRuler() float64
	IsMovable() bool
	SetMovable(value bool)
	IsRemovable() bool
	SetRemovable(value bool)
	MarkerLocation() float64
	SetMarkerLocation(value float64)
	RepresentedObject() CopyingWrapper
	SetRepresentedObject(value ICopying)
	SetRepresentedObject0(value objc.IObject)
	IsDragging() bool
}

type RulerMarker struct {
	objc.Object
}

func MakeRulerMarker(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ RulerMarker) InitWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin foundation.Point) RulerMarker {
	rv := objc.CallMethod[RulerMarker](r_, objc.GetSelector("initWithRulerView:markerLocation:image:imageOrigin:"), objc.ExtractPtr(ruler), location, objc.ExtractPtr(image), imageOrigin)
	return rv
}

func RulerMarker_InitWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin foundation.Point) RulerMarker {
	return RulerMarkerClass.Alloc().InitWithRulerViewMarkerLocationImageImageOrigin(ruler, location, image, imageOrigin)
}

func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, objc.GetSelector("alloc"))
	return rv
}

func RulerMarker_Alloc() RulerMarker {
	return RulerMarkerClass.Alloc()
}

func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

func RulerMarker_New() RulerMarker {
	return RulerMarkerClass.New()
}

func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.CallMethod[RulerMarker](r_, objc.GetSelector("init"))
	return rv
}

func RulerMarker_Init() RulerMarker {
	return RulerMarkerClass.Alloc().Init()
}

func (r_ RulerMarker) DrawRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("drawRect:"), rect)
}

func (r_ RulerMarker) TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("trackMouse:adding:"), objc.ExtractPtr(mouseDownEvent), isAdding)
	return rv
}

func (r_ RulerMarker) Ruler() RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.GetSelector("ruler"))
	return rv
}

func (r_ RulerMarker) Image() Image {
	rv := objc.CallMethod[Image](r_, objc.GetSelector("image"))
	return rv
}

func (r_ RulerMarker) SetImage(value IImage) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (r_ RulerMarker) ImageOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](r_, objc.GetSelector("imageOrigin"))
	return rv
}

func (r_ RulerMarker) SetImageOrigin(value foundation.Point) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setImageOrigin:"), value)
}

func (r_ RulerMarker) ImageRectInRuler() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](r_, objc.GetSelector("imageRectInRuler"))
	return rv
}

func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("thicknessRequiredInRuler"))
	return rv
}

func (r_ RulerMarker) IsMovable() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isMovable"))
	return rv
}

func (r_ RulerMarker) SetMovable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setMovable:"), value)
}

func (r_ RulerMarker) IsRemovable() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isRemovable"))
	return rv
}

func (r_ RulerMarker) SetRemovable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRemovable:"), value)
}

func (r_ RulerMarker) MarkerLocation() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("markerLocation"))
	return rv
}

func (r_ RulerMarker) SetMarkerLocation(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setMarkerLocation:"), value)
}

func (r_ RulerMarker) RepresentedObject() CopyingWrapper {
	rv := objc.CallMethod[CopyingWrapper](r_, objc.GetSelector("representedObject"))
	return rv
}

func (r_ RulerMarker) SetRepresentedObject(value ICopying) {
	po := objc.WrapAsProtocol("NSCopying", value)
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRepresentedObject:"), po)
}

func (r_ RulerMarker) SetRepresentedObject0(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRepresentedObject:"), objc.ExtractPtr(value))
}

func (r_ RulerMarker) IsDragging() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isDragging"))
	return rv
}
