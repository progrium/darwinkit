// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var LayerClass = _LayerClass{objc.GetClass("CALayer")}

type _LayerClass struct {
	objc.Class
}

type ILayer interface {
	objc.IObject
	Display()
	DrawInContext(ctx coregraphics.ContextRef)
	ContentsAreFlipped() bool
	RenderInContext(ctx coregraphics.ContextRef)
	AffineTransform() coregraphics.AffineTransform
	SetAffineTransform(m coregraphics.AffineTransform)
	AddSublayer(layer ILayer)
	RemoveFromSuperlayer()
	InsertSublayerAtIndex(layer ILayer, idx uint32)
	InsertSublayerBelow(layer ILayer, sibling ILayer)
	InsertSublayerAbove(layer ILayer, sibling ILayer)
	ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer)
	SetNeedsDisplay()
	SetNeedsDisplayInRect(r coregraphics.Rect)
	DisplayIfNeeded()
	NeedsDisplay() bool
	AddAnimationForKey(anim IAnimation, key string)
	AnimationForKey(key string) Animation
	RemoveAllAnimations()
	RemoveAnimationForKey(key string)
	AnimationKeys() []string
	SetNeedsLayout()
	LayoutSublayers()
	LayoutIfNeeded()
	NeedsLayout() bool
	ResizeWithOldSuperlayerSize(size coregraphics.Size)
	ResizeSublayersWithOldSize(size coregraphics.Size)
	PreferredFrameSize() coregraphics.Size
	AddConstraint(c IConstraint)
	ActionForKey(event string) ActionWrapper
	ConvertPointFromLayer(p coregraphics.Point, l ILayer) coregraphics.Point
	ConvertPointToLayer(p coregraphics.Point, l ILayer) coregraphics.Point
	ConvertRectFromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect
	ConvertRectToLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect
	HitTest(p coregraphics.Point) Layer
	ContainsPoint(p coregraphics.Point) bool
	ScrollPoint(p coregraphics.Point)
	ScrollRectToVisible(r coregraphics.Rect)
	ShouldArchiveValueForKey(key string) bool
	Delegate() LayerDelegateWrapper
	SetDelegate(value ILayerDelegate)
	SetDelegate0(value objc.IObject)
	Contents() objc.Object
	SetContents(value objc.IObject)
	ContentsRect() coregraphics.Rect
	SetContentsRect(value coregraphics.Rect)
	ContentsCenter() coregraphics.Rect
	SetContentsCenter(value coregraphics.Rect)
	ContentsGravity() LayerContentsGravity
	SetContentsGravity(value LayerContentsGravity)
	Opacity() float32
	SetOpacity(value float32)
	IsHidden() bool
	SetHidden(value bool)
	MasksToBounds() bool
	SetMasksToBounds(value bool)
	Mask() Layer
	SetMask(value ILayer)
	IsDoubleSided() bool
	SetDoubleSided(value bool)
	CornerRadius() float64
	SetCornerRadius(value float64)
	MaskedCorners() CornerMask
	SetMaskedCorners(value CornerMask)
	BorderWidth() float64
	SetBorderWidth(value float64)
	BorderColor() coregraphics.ColorRef
	SetBorderColor(value coregraphics.ColorRef)
	BackgroundColor() coregraphics.ColorRef
	SetBackgroundColor(value coregraphics.ColorRef)
	ShadowOpacity() float32
	SetShadowOpacity(value float32)
	ShadowRadius() float64
	SetShadowRadius(value float64)
	ShadowOffset() coregraphics.Size
	SetShadowOffset(value coregraphics.Size)
	ShadowColor() coregraphics.ColorRef
	SetShadowColor(value coregraphics.ColorRef)
	ShadowPath() coregraphics.PathRef
	SetShadowPath(value coregraphics.PathRef)
	AllowsEdgeAntialiasing() bool
	SetAllowsEdgeAntialiasing(value bool)
	AllowsGroupOpacity() bool
	SetAllowsGroupOpacity(value bool)
	Filters() []objc.Object
	SetFilters(value []objc.IObject)
	CompositingFilter() objc.Object
	SetCompositingFilter(value objc.IObject)
	BackgroundFilters() []objc.Object
	SetBackgroundFilters(value []objc.IObject)
	MinificationFilter() LayerContentsFilter
	SetMinificationFilter(value LayerContentsFilter)
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	MagnificationFilter() LayerContentsFilter
	SetMagnificationFilter(value LayerContentsFilter)
	IsOpaque() bool
	SetOpaque(value bool)
	EdgeAntialiasingMask() EdgeAntialiasingMask
	SetEdgeAntialiasingMask(value EdgeAntialiasingMask)
	IsGeometryFlipped() bool
	SetGeometryFlipped(value bool)
	DrawsAsynchronously() bool
	SetDrawsAsynchronously(value bool)
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	RasterizationScale() float64
	SetRasterizationScale(value float64)
	ContentsFormat() LayerContentsFormat
	SetContentsFormat(value LayerContentsFormat)
	Frame() coregraphics.Rect
	SetFrame(value coregraphics.Rect)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	Position() coregraphics.Point
	SetPosition(value coregraphics.Point)
	ZPosition() float64
	SetZPosition(value float64)
	AnchorPointZ() float64
	SetAnchorPointZ(value float64)
	AnchorPoint() coregraphics.Point
	SetAnchorPoint(value coregraphics.Point)
	ContentsScale() float64
	SetContentsScale(value float64)
	Transform() Transform3D
	SetTransform(value Transform3D)
	SublayerTransform() Transform3D
	SetSublayerTransform(value Transform3D)
	Sublayers() []Layer
	SetSublayers(value []ILayer)
	Superlayer() Layer
	NeedsDisplayOnBoundsChange() bool
	SetNeedsDisplayOnBoundsChange(value bool)
	LayoutManager() LayoutManagerWrapper
	SetLayoutManager(value ILayoutManager)
	SetLayoutManager0(value objc.IObject)
	AutoresizingMask() AutoresizingMask
	SetAutoresizingMask(value AutoresizingMask)
	Constraints() []Constraint
	SetConstraints(value []IConstraint)
	Actions() foundation.Dictionary
	SetActions(value foundation.Dictionary)
	VisibleRect() coregraphics.Rect
	Name() string
	SetName(value string)
	CornerCurve() LayerCornerCurve
	SetCornerCurve(value LayerCornerCurve)
}

type Layer struct {
	objc.Object
}

func MakeLayer(ptr unsafe.Pointer) Layer {
	return Layer{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayerClass) Layer() Layer {
	rv := objc.CallMethod[Layer](lc, objc.GetSelector("layer"))
	return rv
}

func Layer_Layer() Layer {
	return LayerClass.Layer()
}

func (l_ Layer) Init() Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("init"))
	return rv
}

func Layer_Init() Layer {
	return LayerClass.Alloc().Init()
}

func (l_ Layer) InitWithLayer(layer objc.IObject) Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func Layer_InitWithLayer(layer objc.IObject) Layer {
	return LayerClass.Alloc().InitWithLayer(layer)
}

func (l_ Layer) PresentationLayer() Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("presentationLayer"))
	return rv
}

func Layer_PresentationLayer() Layer {
	return LayerClass.Alloc().PresentationLayer()
}

func (l_ Layer) ModelLayer() Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("modelLayer"))
	return rv
}

func Layer_ModelLayer() Layer {
	return LayerClass.Alloc().ModelLayer()
}

func (lc _LayerClass) Alloc() Layer {
	rv := objc.CallMethod[Layer](lc, objc.GetSelector("alloc"))
	return rv
}

func Layer_Alloc() Layer {
	return LayerClass.Alloc()
}

func (lc _LayerClass) New() Layer {
	rv := objc.CallMethod[Layer](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayer() Layer {
	return LayerClass.New()
}

func Layer_New() Layer {
	return LayerClass.New()
}

func (lc _LayerClass) LayerWithRemoteClientId(client_id uint32) Layer {
	rv := objc.CallMethod[Layer](lc, objc.GetSelector("layerWithRemoteClientId:"), client_id)
	return rv
}

func Layer_LayerWithRemoteClientId(client_id uint32) Layer {
	return LayerClass.LayerWithRemoteClientId(client_id)
}

func (l_ Layer) Display() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("display"))
}

func (l_ Layer) DrawInContext(ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawInContext:"), ctx)
}

func (l_ Layer) ContentsAreFlipped() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("contentsAreFlipped"))
	return rv
}

func (l_ Layer) RenderInContext(ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("renderInContext:"), ctx)
}

func (l_ Layer) AffineTransform() coregraphics.AffineTransform {
	rv := objc.CallMethod[coregraphics.AffineTransform](l_, objc.GetSelector("affineTransform"))
	return rv
}

func (l_ Layer) SetAffineTransform(m coregraphics.AffineTransform) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAffineTransform:"), m)
}

func (l_ Layer) AddSublayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addSublayer:"), objc.ExtractPtr(layer))
}

func (l_ Layer) RemoveFromSuperlayer() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeFromSuperlayer"))
}

func (l_ Layer) InsertSublayerAtIndex(layer ILayer, idx uint32) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertSublayer:atIndex:"), objc.ExtractPtr(layer), idx)
}

func (l_ Layer) InsertSublayerBelow(layer ILayer, sibling ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertSublayer:below:"), objc.ExtractPtr(layer), objc.ExtractPtr(sibling))
}

func (l_ Layer) InsertSublayerAbove(layer ILayer, sibling ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertSublayer:above:"), objc.ExtractPtr(layer), objc.ExtractPtr(sibling))
}

func (l_ Layer) ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("replaceSublayer:with:"), objc.ExtractPtr(oldLayer), objc.ExtractPtr(newLayer))
}

func (l_ Layer) SetNeedsDisplay() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNeedsDisplay"))
}

func (l_ Layer) SetNeedsDisplayInRect(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNeedsDisplayInRect:"), r)
}

func (l_ Layer) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("displayIfNeeded"))
}

func (l_ Layer) NeedsDisplay() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("needsDisplay"))
	return rv
}

func (lc _LayerClass) NeedsDisplayForKey(key string) bool {
	rv := objc.CallMethod[bool](lc, objc.GetSelector("needsDisplayForKey:"), key)
	return rv
}

func Layer_NeedsDisplayForKey(key string) bool {
	return LayerClass.NeedsDisplayForKey(key)
}

func (l_ Layer) AddAnimationForKey(anim IAnimation, key string) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addAnimation:forKey:"), objc.ExtractPtr(anim), key)
}

func (l_ Layer) AnimationForKey(key string) Animation {
	rv := objc.CallMethod[Animation](l_, objc.GetSelector("animationForKey:"), key)
	return rv
}

func (l_ Layer) RemoveAllAnimations() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeAllAnimations"))
}

func (l_ Layer) RemoveAnimationForKey(key string) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeAnimationForKey:"), key)
}

func (l_ Layer) AnimationKeys() []string {
	rv := objc.CallMethod[[]string](l_, objc.GetSelector("animationKeys"))
	return rv
}

func (l_ Layer) SetNeedsLayout() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNeedsLayout"))
}

func (l_ Layer) LayoutSublayers() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutSublayers"))
}

func (l_ Layer) LayoutIfNeeded() {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutIfNeeded"))
}

func (l_ Layer) NeedsLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("needsLayout"))
	return rv
}

func (l_ Layer) ResizeWithOldSuperlayerSize(size coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("resizeWithOldSuperlayerSize:"), size)
}

func (l_ Layer) ResizeSublayersWithOldSize(size coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("resizeSublayersWithOldSize:"), size)
}

func (l_ Layer) PreferredFrameSize() coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, objc.GetSelector("preferredFrameSize"))
	return rv
}

func (l_ Layer) AddConstraint(c IConstraint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addConstraint:"), objc.ExtractPtr(c))
}

func (l_ Layer) ActionForKey(event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](l_, objc.GetSelector("actionForKey:"), event)
	return rv
}

func (lc _LayerClass) DefaultActionForKey(event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](lc, objc.GetSelector("defaultActionForKey:"), event)
	return rv
}

func Layer_DefaultActionForKey(event string) ActionWrapper {
	return LayerClass.DefaultActionForKey(event)
}

func (l_ Layer) ConvertPointFromLayer(p coregraphics.Point, l ILayer) coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, objc.GetSelector("convertPoint:fromLayer:"), p, objc.ExtractPtr(l))
	return rv
}

func (l_ Layer) ConvertPointToLayer(p coregraphics.Point, l ILayer) coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, objc.GetSelector("convertPoint:toLayer:"), p, objc.ExtractPtr(l))
	return rv
}

func (l_ Layer) ConvertRectFromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("convertRect:fromLayer:"), r, objc.ExtractPtr(l))
	return rv
}

func (l_ Layer) ConvertRectToLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("convertRect:toLayer:"), r, objc.ExtractPtr(l))
	return rv
}

func (l_ Layer) HitTest(p coregraphics.Point) Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("hitTest:"), p)
	return rv
}

func (l_ Layer) ContainsPoint(p coregraphics.Point) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("containsPoint:"), p)
	return rv
}

func (l_ Layer) ScrollPoint(p coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("scrollPoint:"), p)
}

func (l_ Layer) ScrollRectToVisible(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("scrollRectToVisible:"), r)
}

func (l_ Layer) ShouldArchiveValueForKey(key string) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("shouldArchiveValueForKey:"), key)
	return rv
}

func (lc _LayerClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.CallMethod[objc.Object](lc, objc.GetSelector("defaultValueForKey:"), key)
	return rv
}

func Layer_DefaultValueForKey(key string) objc.Object {
	return LayerClass.DefaultValueForKey(key)
}

func (lc _LayerClass) CornerCurveExpansionFactor(curve LayerCornerCurve) float64 {
	rv := objc.CallMethod[float64](lc, objc.GetSelector("cornerCurveExpansionFactor:"), curve)
	return rv
}

func Layer_CornerCurveExpansionFactor(curve LayerCornerCurve) float64 {
	return LayerClass.CornerCurveExpansionFactor(curve)
}

func (l_ Layer) Delegate() LayerDelegateWrapper {
	rv := objc.CallMethod[LayerDelegateWrapper](l_, objc.GetSelector("delegate"))
	return rv
}

func (l_ Layer) SetDelegate(value ILayerDelegate) {
	po := objc.WrapAsProtocol("CALayerDelegate", value)
	objc.SetAssociatedObject(l_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), po)
}

func (l_ Layer) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (l_ Layer) Contents() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("contents"))
	return rv
}

func (l_ Layer) SetContents(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContents:"), objc.ExtractPtr(value))
}

func (l_ Layer) ContentsRect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("contentsRect"))
	return rv
}

func (l_ Layer) SetContentsRect(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContentsRect:"), value)
}

func (l_ Layer) ContentsCenter() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("contentsCenter"))
	return rv
}

func (l_ Layer) SetContentsCenter(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContentsCenter:"), value)
}

func (l_ Layer) ContentsGravity() LayerContentsGravity {
	rv := objc.CallMethod[LayerContentsGravity](l_, objc.GetSelector("contentsGravity"))
	return rv
}

func (l_ Layer) SetContentsGravity(value LayerContentsGravity) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContentsGravity:"), value)
}

func (l_ Layer) Opacity() float32 {
	rv := objc.CallMethod[float32](l_, objc.GetSelector("opacity"))
	return rv
}

func (l_ Layer) SetOpacity(value float32) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setOpacity:"), value)
}

func (l_ Layer) IsHidden() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isHidden"))
	return rv
}

func (l_ Layer) SetHidden(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setHidden:"), value)
}

func (l_ Layer) MasksToBounds() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("masksToBounds"))
	return rv
}

func (l_ Layer) SetMasksToBounds(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMasksToBounds:"), value)
}

func (l_ Layer) Mask() Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("mask"))
	return rv
}

func (l_ Layer) SetMask(value ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMask:"), objc.ExtractPtr(value))
}

func (l_ Layer) IsDoubleSided() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isDoubleSided"))
	return rv
}

func (l_ Layer) SetDoubleSided(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDoubleSided:"), value)
}

func (l_ Layer) CornerRadius() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("cornerRadius"))
	return rv
}

func (l_ Layer) SetCornerRadius(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCornerRadius:"), value)
}

func (l_ Layer) MaskedCorners() CornerMask {
	rv := objc.CallMethod[CornerMask](l_, objc.GetSelector("maskedCorners"))
	return rv
}

func (l_ Layer) SetMaskedCorners(value CornerMask) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMaskedCorners:"), value)
}

func (l_ Layer) BorderWidth() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("borderWidth"))
	return rv
}

func (l_ Layer) SetBorderWidth(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBorderWidth:"), value)
}

func (l_ Layer) BorderColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, objc.GetSelector("borderColor"))
	return rv
}

func (l_ Layer) SetBorderColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBorderColor:"), value)
}

func (l_ Layer) BackgroundColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, objc.GetSelector("backgroundColor"))
	return rv
}

func (l_ Layer) SetBackgroundColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBackgroundColor:"), value)
}

func (l_ Layer) ShadowOpacity() float32 {
	rv := objc.CallMethod[float32](l_, objc.GetSelector("shadowOpacity"))
	return rv
}

func (l_ Layer) SetShadowOpacity(value float32) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShadowOpacity:"), value)
}

func (l_ Layer) ShadowRadius() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("shadowRadius"))
	return rv
}

func (l_ Layer) SetShadowRadius(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShadowRadius:"), value)
}

func (l_ Layer) ShadowOffset() coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, objc.GetSelector("shadowOffset"))
	return rv
}

func (l_ Layer) SetShadowOffset(value coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShadowOffset:"), value)
}

func (l_ Layer) ShadowColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, objc.GetSelector("shadowColor"))
	return rv
}

func (l_ Layer) SetShadowColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShadowColor:"), value)
}

func (l_ Layer) ShadowPath() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](l_, objc.GetSelector("shadowPath"))
	return rv
}

func (l_ Layer) SetShadowPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShadowPath:"), value)
}

func (l_ Layer) AllowsEdgeAntialiasing() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("allowsEdgeAntialiasing"))
	return rv
}

func (l_ Layer) SetAllowsEdgeAntialiasing(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAllowsEdgeAntialiasing:"), value)
}

func (l_ Layer) AllowsGroupOpacity() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("allowsGroupOpacity"))
	return rv
}

func (l_ Layer) SetAllowsGroupOpacity(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAllowsGroupOpacity:"), value)
}

func (l_ Layer) Filters() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](l_, objc.GetSelector("filters"))
	// TODO: convert slice items...
	return rv
}

func (l_ Layer) SetFilters(value []objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setFilters:"), value)
}

func (l_ Layer) CompositingFilter() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("compositingFilter"))
	return rv
}

func (l_ Layer) SetCompositingFilter(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCompositingFilter:"), objc.ExtractPtr(value))
}

func (l_ Layer) BackgroundFilters() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](l_, objc.GetSelector("backgroundFilters"))
	// TODO: convert slice items...
	return rv
}

func (l_ Layer) SetBackgroundFilters(value []objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBackgroundFilters:"), value)
}

func (l_ Layer) MinificationFilter() LayerContentsFilter {
	rv := objc.CallMethod[LayerContentsFilter](l_, objc.GetSelector("minificationFilter"))
	return rv
}

func (l_ Layer) SetMinificationFilter(value LayerContentsFilter) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMinificationFilter:"), value)
}

func (l_ Layer) MinificationFilterBias() float32 {
	rv := objc.CallMethod[float32](l_, objc.GetSelector("minificationFilterBias"))
	return rv
}

func (l_ Layer) SetMinificationFilterBias(value float32) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMinificationFilterBias:"), value)
}

func (l_ Layer) MagnificationFilter() LayerContentsFilter {
	rv := objc.CallMethod[LayerContentsFilter](l_, objc.GetSelector("magnificationFilter"))
	return rv
}

func (l_ Layer) SetMagnificationFilter(value LayerContentsFilter) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMagnificationFilter:"), value)
}

func (l_ Layer) IsOpaque() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isOpaque"))
	return rv
}

func (l_ Layer) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setOpaque:"), value)
}

func (l_ Layer) EdgeAntialiasingMask() EdgeAntialiasingMask {
	rv := objc.CallMethod[EdgeAntialiasingMask](l_, objc.GetSelector("edgeAntialiasingMask"))
	return rv
}

func (l_ Layer) SetEdgeAntialiasingMask(value EdgeAntialiasingMask) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setEdgeAntialiasingMask:"), value)
}

func (l_ Layer) IsGeometryFlipped() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isGeometryFlipped"))
	return rv
}

func (l_ Layer) SetGeometryFlipped(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setGeometryFlipped:"), value)
}

func (l_ Layer) DrawsAsynchronously() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("drawsAsynchronously"))
	return rv
}

func (l_ Layer) SetDrawsAsynchronously(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDrawsAsynchronously:"), value)
}

func (l_ Layer) ShouldRasterize() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("shouldRasterize"))
	return rv
}

func (l_ Layer) SetShouldRasterize(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShouldRasterize:"), value)
}

func (l_ Layer) RasterizationScale() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("rasterizationScale"))
	return rv
}

func (l_ Layer) SetRasterizationScale(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setRasterizationScale:"), value)
}

func (l_ Layer) ContentsFormat() LayerContentsFormat {
	rv := objc.CallMethod[LayerContentsFormat](l_, objc.GetSelector("contentsFormat"))
	return rv
}

func (l_ Layer) SetContentsFormat(value LayerContentsFormat) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContentsFormat:"), value)
}

func (l_ Layer) Frame() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("frame"))
	return rv
}

func (l_ Layer) SetFrame(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setFrame:"), value)
}

func (l_ Layer) Bounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("bounds"))
	return rv
}

func (l_ Layer) SetBounds(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBounds:"), value)
}

func (l_ Layer) Position() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, objc.GetSelector("position"))
	return rv
}

func (l_ Layer) SetPosition(value coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setPosition:"), value)
}

func (l_ Layer) ZPosition() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("zPosition"))
	return rv
}

func (l_ Layer) SetZPosition(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setZPosition:"), value)
}

func (l_ Layer) AnchorPointZ() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("anchorPointZ"))
	return rv
}

func (l_ Layer) SetAnchorPointZ(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAnchorPointZ:"), value)
}

func (l_ Layer) AnchorPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, objc.GetSelector("anchorPoint"))
	return rv
}

func (l_ Layer) SetAnchorPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAnchorPoint:"), value)
}

func (l_ Layer) ContentsScale() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("contentsScale"))
	return rv
}

func (l_ Layer) SetContentsScale(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setContentsScale:"), value)
}

func (l_ Layer) Transform() Transform3D {
	rv := objc.CallMethod[Transform3D](l_, objc.GetSelector("transform"))
	return rv
}

func (l_ Layer) SetTransform(value Transform3D) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTransform:"), value)
}

func (l_ Layer) SublayerTransform() Transform3D {
	rv := objc.CallMethod[Transform3D](l_, objc.GetSelector("sublayerTransform"))
	return rv
}

func (l_ Layer) SetSublayerTransform(value Transform3D) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setSublayerTransform:"), value)
}

func (l_ Layer) Sublayers() []Layer {
	rv := objc.CallMethod[[]Layer](l_, objc.GetSelector("sublayers"))
	// TODO: convert slice items...
	return rv
}

func (l_ Layer) SetSublayers(value []ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setSublayers:"), value)
}

func (l_ Layer) Superlayer() Layer {
	rv := objc.CallMethod[Layer](l_, objc.GetSelector("superlayer"))
	return rv
}

func (l_ Layer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("needsDisplayOnBoundsChange"))
	return rv
}

func (l_ Layer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNeedsDisplayOnBoundsChange:"), value)
}

func (l_ Layer) LayoutManager() LayoutManagerWrapper {
	rv := objc.CallMethod[LayoutManagerWrapper](l_, objc.GetSelector("layoutManager"))
	return rv
}

func (l_ Layer) SetLayoutManager(value ILayoutManager) {
	po := objc.WrapAsProtocol("CALayoutManager", value)
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLayoutManager:"), po)
}

func (l_ Layer) SetLayoutManager0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLayoutManager:"), objc.ExtractPtr(value))
}

func (l_ Layer) AutoresizingMask() AutoresizingMask {
	rv := objc.CallMethod[AutoresizingMask](l_, objc.GetSelector("autoresizingMask"))
	return rv
}

func (l_ Layer) SetAutoresizingMask(value AutoresizingMask) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAutoresizingMask:"), value)
}

func (l_ Layer) Constraints() []Constraint {
	rv := objc.CallMethod[[]Constraint](l_, objc.GetSelector("constraints"))
	// TODO: convert slice items...
	return rv
}

func (l_ Layer) SetConstraints(value []IConstraint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setConstraints:"), value)
}

func (l_ Layer) Actions() foundation.Dictionary {
	rv := objc.CallMethod[foundation.Dictionary](l_, objc.GetSelector("actions"))
	return rv
}

func (l_ Layer) SetActions(value foundation.Dictionary) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setActions:"), value)
}

func (l_ Layer) VisibleRect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, objc.GetSelector("visibleRect"))
	return rv
}

func (l_ Layer) Name() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("name"))
	return rv
}

func (l_ Layer) SetName(value string) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setName:"), value)
}

func (l_ Layer) CornerCurve() LayerCornerCurve {
	rv := objc.CallMethod[LayerCornerCurve](l_, objc.GetSelector("cornerCurve"))
	return rv
}

func (l_ Layer) SetCornerCurve(value LayerCornerCurve) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCornerCurve:"), value)
}
