// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var BezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}

type _BezierPathClass struct {
	objc.Class
}

type IBezierPath interface {
	objc.IObject
	MoveToPoint(point foundation.Point)
	LineToPoint(point foundation.Point)
	CurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	ClosePath()
	RelativeMoveToPoint(point foundation.Point)
	RelativeLineToPoint(point foundation.Point)
	RelativeCurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	AppendBezierPath(path IBezierPath)
	AppendBezierPathWithPointsCount(points *foundation.Point, count int)
	AppendBezierPathWithOvalInRect(rect foundation.Rect)
	AppendBezierPathWithArcFromPointToPointRadius(point1 foundation.Point, point2 foundation.Point, radius float64)
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64)
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center foundation.Point, radius float64, startAngle float64, endAngle float64, clockwise bool)
	AppendBezierPathWithRect(rect foundation.Rect)
	AppendBezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64)
	AppendBezierPathWithCGGlyphInFont(glyph coregraphics.Glyph, font IFont)
	AppendBezierPathWithCGGlyphsCountInFont(glyphs *coregraphics.Glyph, count int, font IFont)
	GetLineDashCountPhase(pattern *float64, count *int, phase *float64)
	SetLineDashCountPhase(pattern *float64, count int, phase float64)
	Stroke()
	Fill()
	AddClip()
	SetClip()
	ContainsPoint(point foundation.Point) bool
	TransformUsingAffineTransform(transform foundation.IAffineTransform)
	ElementAtIndex(index int) BezierPathElement
	ElementAtIndexAssociatedPoints(index int, points *foundation.Point) BezierPathElement
	RemoveAllPoints()
	SetAssociatedPointsAtIndex(points *foundation.Point, index int)
	BezierPathByFlatteningPath() BezierPath
	BezierPathByReversingPath() BezierPath
	WindingRule() WindingRule
	SetWindingRule(value WindingRule)
	LineCapStyle() LineCapStyle
	SetLineCapStyle(value LineCapStyle)
	LineJoinStyle() LineJoinStyle
	SetLineJoinStyle(value LineJoinStyle)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	Flatness() float64
	SetFlatness(value float64)
	Bounds() foundation.Rect
	ControlPointBounds() foundation.Rect
	CurrentPoint() foundation.Point
	IsEmpty() bool
	ElementCount() int
}

type BezierPath struct {
	objc.Object
}

func MakeBezierPath(ptr unsafe.Pointer) BezierPath {
	return BezierPath{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("alloc"))
	return rv
}

func BezierPath_Alloc() BezierPath {
	return BezierPathClass.Alloc()
}

func (bc _BezierPathClass) New() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBezierPath() BezierPath {
	return BezierPathClass.New()
}

func BezierPath_New() BezierPath {
	return BezierPathClass.New()
}

func (b_ BezierPath) Init() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.GetSelector("init"))
	return rv
}

func BezierPath_Init() BezierPath {
	return BezierPathClass.Alloc().Init()
}

func (bc _BezierPathClass) BezierPath() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("bezierPath"))
	return rv
}

func BezierPath_BezierPath() BezierPath {
	return BezierPathClass.BezierPath()
}

func (bc _BezierPathClass) BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("bezierPathWithOvalInRect:"), rect)
	return rv
}

func BezierPath_BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	return BezierPathClass.BezierPathWithOvalInRect(rect)
}

func (bc _BezierPathClass) BezierPathWithRect(rect foundation.Rect) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("bezierPathWithRect:"), rect)
	return rv
}

func BezierPath_BezierPathWithRect(rect foundation.Rect) BezierPath {
	return BezierPathClass.BezierPathWithRect(rect)
}

func (bc _BezierPathClass) BezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.GetSelector("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return rv
}

func BezierPath_BezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	return BezierPathClass.BezierPathWithRoundedRectXRadiusYRadius(rect, xRadius, yRadius)
}

func (b_ BezierPath) MoveToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("moveToPoint:"), point)
}

func (b_ BezierPath) LineToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("lineToPoint:"), point)
}

func (b_ BezierPath) CurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("curveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) ClosePath() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("closePath"))
}

func (b_ BezierPath) RelativeMoveToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("relativeMoveToPoint:"), point)
}

func (b_ BezierPath) RelativeLineToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("relativeLineToPoint:"), point)
}

func (b_ BezierPath) RelativeCurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("relativeCurveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) AppendBezierPath(path IBezierPath) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPath:"), objc.ExtractPtr(path))
}

func (b_ BezierPath) AppendBezierPathWithPointsCount(points *foundation.Point, count int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithPoints:count:"), points, count)
}

func (b_ BezierPath) AppendBezierPathWithOvalInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithOvalInRect:"), rect)
}

func (b_ BezierPath) AppendBezierPathWithArcFromPointToPointRadius(point1 foundation.Point, point2 foundation.Point, radius float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithArcFromPoint:toPoint:radius:"), point1, point2, radius)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:"), center, radius, startAngle, endAngle)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center foundation.Point, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:clockwise:"), center, radius, startAngle, endAngle, clockwise)
}

func (b_ BezierPath) AppendBezierPathWithRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithRect:"), rect)
}

func (b_ BezierPath) AppendBezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
}

func (b_ BezierPath) AppendBezierPathWithCGGlyphInFont(glyph coregraphics.Glyph, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithCGGlyph:inFont:"), glyph, objc.ExtractPtr(font))
}

func (b_ BezierPath) AppendBezierPathWithCGGlyphsCountInFont(glyphs *coregraphics.Glyph, count int, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("appendBezierPathWithCGGlyphs:count:inFont:"), glyphs, count, objc.ExtractPtr(font))
}

func (b_ BezierPath) GetLineDashCountPhase(pattern *float64, count *int, phase *float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getLineDash:count:phase:"), pattern, count, phase)
}

func (b_ BezierPath) SetLineDashCountPhase(pattern *float64, count int, phase float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLineDash:count:phase:"), pattern, count, phase)
}

func (b_ BezierPath) Stroke() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("stroke"))
}

func (b_ BezierPath) Fill() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("fill"))
}

func (bc _BezierPathClass) FillRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("fillRect:"), rect)
}

func BezierPath_FillRect(rect foundation.Rect) {
	BezierPathClass.FillRect(rect)
}

func (bc _BezierPathClass) StrokeRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("strokeRect:"), rect)
}

func BezierPath_StrokeRect(rect foundation.Rect) {
	BezierPathClass.StrokeRect(rect)
}

func (bc _BezierPathClass) StrokeLineFromPointToPoint(point1 foundation.Point, point2 foundation.Point) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("strokeLineFromPoint:toPoint:"), point1, point2)
}

func BezierPath_StrokeLineFromPointToPoint(point1 foundation.Point, point2 foundation.Point) {
	BezierPathClass.StrokeLineFromPointToPoint(point1, point2)
}

func (bc _BezierPathClass) DrawPackedGlyphsAtPoint(packedGlyphs *byte, point foundation.Point) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("drawPackedGlyphs:atPoint:"), packedGlyphs, point)
}

func BezierPath_DrawPackedGlyphsAtPoint(packedGlyphs *byte, point foundation.Point) {
	BezierPathClass.DrawPackedGlyphsAtPoint(packedGlyphs, point)
}

func (b_ BezierPath) AddClip() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("addClip"))
}

func (b_ BezierPath) SetClip() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setClip"))
}

func (bc _BezierPathClass) ClipRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("clipRect:"), rect)
}

func BezierPath_ClipRect(rect foundation.Rect) {
	BezierPathClass.ClipRect(rect)
}

func (b_ BezierPath) ContainsPoint(point foundation.Point) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("containsPoint:"), point)
	return rv
}

func (b_ BezierPath) TransformUsingAffineTransform(transform foundation.IAffineTransform) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("transformUsingAffineTransform:"), objc.ExtractPtr(transform))
}

func (b_ BezierPath) ElementAtIndex(index int) BezierPathElement {
	rv := objc.CallMethod[BezierPathElement](b_, objc.GetSelector("elementAtIndex:"), index)
	return rv
}

func (b_ BezierPath) ElementAtIndexAssociatedPoints(index int, points *foundation.Point) BezierPathElement {
	rv := objc.CallMethod[BezierPathElement](b_, objc.GetSelector("elementAtIndex:associatedPoints:"), index, points)
	return rv
}

func (b_ BezierPath) RemoveAllPoints() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("removeAllPoints"))
}

func (b_ BezierPath) SetAssociatedPointsAtIndex(points *foundation.Point, index int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAssociatedPoints:atIndex:"), points, index)
}

func (b_ BezierPath) BezierPathByFlatteningPath() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.GetSelector("bezierPathByFlatteningPath"))
	return rv
}

func (b_ BezierPath) BezierPathByReversingPath() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.GetSelector("bezierPathByReversingPath"))
	return rv
}

func (b_ BezierPath) WindingRule() WindingRule {
	rv := objc.CallMethod[WindingRule](b_, objc.GetSelector("windingRule"))
	return rv
}

func (b_ BezierPath) SetWindingRule(value WindingRule) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setWindingRule:"), value)
}

func (b_ BezierPath) LineCapStyle() LineCapStyle {
	rv := objc.CallMethod[LineCapStyle](b_, objc.GetSelector("lineCapStyle"))
	return rv
}

func (b_ BezierPath) SetLineCapStyle(value LineCapStyle) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLineCapStyle:"), value)
}

func (b_ BezierPath) LineJoinStyle() LineJoinStyle {
	rv := objc.CallMethod[LineJoinStyle](b_, objc.GetSelector("lineJoinStyle"))
	return rv
}

func (b_ BezierPath) SetLineJoinStyle(value LineJoinStyle) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLineJoinStyle:"), value)
}

func (b_ BezierPath) LineWidth() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("lineWidth"))
	return rv
}

func (b_ BezierPath) SetLineWidth(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLineWidth:"), value)
}

func (b_ BezierPath) MiterLimit() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("miterLimit"))
	return rv
}

func (b_ BezierPath) SetMiterLimit(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setMiterLimit:"), value)
}

func (b_ BezierPath) Flatness() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("flatness"))
	return rv
}

func (b_ BezierPath) SetFlatness(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFlatness:"), value)
}

func (bc _BezierPathClass) DefaultWindingRule() WindingRule {
	rv := objc.CallMethod[WindingRule](bc, objc.GetSelector("defaultWindingRule"))
	return rv
}

func BezierPath_DefaultWindingRule() WindingRule {
	return BezierPathClass.DefaultWindingRule()
}

func (bc _BezierPathClass) SetDefaultWindingRule(value WindingRule) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultWindingRule:"), value)
}

func BezierPath_SetDefaultWindingRule(value WindingRule) {
	BezierPathClass.SetDefaultWindingRule(value)
}

func (bc _BezierPathClass) DefaultLineCapStyle() LineCapStyle {
	rv := objc.CallMethod[LineCapStyle](bc, objc.GetSelector("defaultLineCapStyle"))
	return rv
}

func BezierPath_DefaultLineCapStyle() LineCapStyle {
	return BezierPathClass.DefaultLineCapStyle()
}

func (bc _BezierPathClass) SetDefaultLineCapStyle(value LineCapStyle) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultLineCapStyle:"), value)
}

func BezierPath_SetDefaultLineCapStyle(value LineCapStyle) {
	BezierPathClass.SetDefaultLineCapStyle(value)
}

func (bc _BezierPathClass) DefaultLineJoinStyle() LineJoinStyle {
	rv := objc.CallMethod[LineJoinStyle](bc, objc.GetSelector("defaultLineJoinStyle"))
	return rv
}

func BezierPath_DefaultLineJoinStyle() LineJoinStyle {
	return BezierPathClass.DefaultLineJoinStyle()
}

func (bc _BezierPathClass) SetDefaultLineJoinStyle(value LineJoinStyle) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultLineJoinStyle:"), value)
}

func BezierPath_SetDefaultLineJoinStyle(value LineJoinStyle) {
	BezierPathClass.SetDefaultLineJoinStyle(value)
}

func (bc _BezierPathClass) DefaultLineWidth() float64 {
	rv := objc.CallMethod[float64](bc, objc.GetSelector("defaultLineWidth"))
	return rv
}

func BezierPath_DefaultLineWidth() float64 {
	return BezierPathClass.DefaultLineWidth()
}

func (bc _BezierPathClass) SetDefaultLineWidth(value float64) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultLineWidth:"), value)
}

func BezierPath_SetDefaultLineWidth(value float64) {
	BezierPathClass.SetDefaultLineWidth(value)
}

func (bc _BezierPathClass) DefaultMiterLimit() float64 {
	rv := objc.CallMethod[float64](bc, objc.GetSelector("defaultMiterLimit"))
	return rv
}

func BezierPath_DefaultMiterLimit() float64 {
	return BezierPathClass.DefaultMiterLimit()
}

func (bc _BezierPathClass) SetDefaultMiterLimit(value float64) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultMiterLimit:"), value)
}

func BezierPath_SetDefaultMiterLimit(value float64) {
	BezierPathClass.SetDefaultMiterLimit(value)
}

func (bc _BezierPathClass) DefaultFlatness() float64 {
	rv := objc.CallMethod[float64](bc, objc.GetSelector("defaultFlatness"))
	return rv
}

func BezierPath_DefaultFlatness() float64 {
	return BezierPathClass.DefaultFlatness()
}

func (bc _BezierPathClass) SetDefaultFlatness(value float64) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("setDefaultFlatness:"), value)
}

func BezierPath_SetDefaultFlatness(value float64) {
	BezierPathClass.SetDefaultFlatness(value)
}

func (b_ BezierPath) Bounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("bounds"))
	return rv
}

func (b_ BezierPath) ControlPointBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("controlPointBounds"))
	return rv
}

func (b_ BezierPath) CurrentPoint() foundation.Point {
	rv := objc.CallMethod[foundation.Point](b_, objc.GetSelector("currentPoint"))
	return rv
}

func (b_ BezierPath) IsEmpty() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isEmpty"))
	return rv
}

func (b_ BezierPath) ElementCount() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("elementCount"))
	return rv
}
