// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BezierPath] class.
var BezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}

type _BezierPathClass struct {
	objc.Class
}

// An interface definition for the [BezierPath] class.
type IBezierPath interface {
	objc.IObject
	AppendBezierPathWithCGGlyphsCountInFont(glyphs *coregraphics.Glyph, count int, font IFont)
	ElementAtIndex(index int) BezierPathElement
	SetClip()
	AddClip()
	AppendBezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64)
	TransformUsingAffineTransform(transform foundation.IAffineTransform)
	AppendBezierPathWithOvalInRect(rect foundation.Rect)
	ContainsPoint(point foundation.Point) bool
	RemoveAllPoints()
	AppendBezierPathWithRect(rect foundation.Rect)
	RelativeCurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	LineToPoint(point foundation.Point)
	GetLineDashCountPhase(pattern *float64, count *int, phase *float64)
	RelativeMoveToPoint(point foundation.Point)
	AppendBezierPathWithArcFromPointToPointRadius(point1 foundation.Point, point2 foundation.Point, radius float64)
	CurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	AppendBezierPath(path IBezierPath)
	Stroke()
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64)
	ClosePath()
	RelativeLineToPoint(point foundation.Point)
	MoveToPoint(point foundation.Point)
	AppendBezierPathWithPointsCount(points foundation.PointArray, count int)
	SetLineDashCountPhase(pattern *float64, count int, phase float64)
	SetAssociatedPointsAtIndex(points foundation.PointArray, index int)
	Fill()
	AppendBezierPathWithCGGlyphInFont(glyph coregraphics.Glyph, font IFont)
	LineWidth() float64
	SetLineWidth(value float64)
	IsEmpty() bool
	BezierPathByReversingPath() BezierPath
	MiterLimit() float64
	SetMiterLimit(value float64)
	Flatness() float64
	SetFlatness(value float64)
	LineJoinStyle() LineJoinStyle
	SetLineJoinStyle(value LineJoinStyle)
	Bounds() foundation.Rect
	LineCapStyle() LineCapStyle
	SetLineCapStyle(value LineCapStyle)
	ControlPointBounds() foundation.Rect
	ElementCount() int
	BezierPathByFlatteningPath() BezierPath
	CurrentPoint() foundation.Point
	WindingRule() WindingRule
	SetWindingRule(value WindingRule)
}

// An object that can create paths using PostScript-style commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath?language=objc
type BezierPath struct {
	objc.Object
}

func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("alloc"))
	return rv
}

func BezierPath_Alloc() BezierPath {
	return BezierPathClass.Alloc()
}

func (bc _BezierPathClass) New() BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBezierPath() BezierPath {
	return BezierPathClass.New()
}

func (b_ BezierPath) Init() BezierPath {
	rv := objc.Call[BezierPath](b_, objc.Sel("init"))
	return rv
}

// Appends the outlines of the specified glyphs to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/2887165-appendbezierpathwithcgglyphs?language=objc
func (b_ BezierPath) AppendBezierPathWithCGGlyphsCountInFont(glyphs *coregraphics.Glyph, count int, font IFont) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithCGGlyphs:count:inFont:"), glyphs, count, objc.Ptr(font))
}

// Returns the type of path element at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520751-elementatindex?language=objc
func (b_ BezierPath) ElementAtIndex(index int) BezierPathElement {
	rv := objc.Call[BezierPathElement](b_, objc.Sel("elementAtIndex:"), index)
	return rv
}

// Creates and returns a new Bézier path object initialized with an oval path inscribed in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520647-bezierpathwithovalinrect?language=objc
func (bc _BezierPathClass) BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("bezierPathWithOvalInRect:"), rect)
	return rv
}

// Creates and returns a new Bézier path object initialized with an oval path inscribed in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520647-bezierpathwithovalinrect?language=objc
func BezierPath_BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	return BezierPathClass.BezierPathWithOvalInRect(rect)
}

// Replaces the clipping path of the current graphics context with the area inside the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520704-setclip?language=objc
func (b_ BezierPath) SetClip() {
	objc.Call[objc.Void](b_, objc.Sel("setClip"))
}

// Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520634-addclip?language=objc
func (b_ BezierPath) AddClip() {
	objc.Call[objc.Void](b_, objc.Sel("addClip"))
}

// Creates and returns a new Bézier path object initialized with a rounded rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520638-bezierpathwithroundedrect?language=objc
func (bc _BezierPathClass) BezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return rv
}

// Creates and returns a new Bézier path object initialized with a rounded rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520638-bezierpathwithroundedrect?language=objc
func BezierPath_BezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	return BezierPathClass.BezierPathWithRoundedRectXRadiusYRadius(rect, xRadius, yRadius)
}

// Appends a rounded rectangular path to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520705-appendbezierpathwithroundedrect?language=objc
func (b_ BezierPath) AppendBezierPathWithRoundedRectXRadiusYRadius(rect foundation.Rect, xRadius float64, yRadius float64) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
}

// Transforms all points in the path using the specified transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520635-transformusingaffinetransform?language=objc
func (b_ BezierPath) TransformUsingAffineTransform(transform foundation.IAffineTransform) {
	objc.Call[objc.Void](b_, objc.Sel("transformUsingAffineTransform:"), objc.Ptr(transform))
}

// Appends an oval path to the path, inscribing the oval in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520718-appendbezierpathwithovalinrect?language=objc
func (b_ BezierPath) AppendBezierPathWithOvalInRect(rect foundation.Rect) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithOvalInRect:"), rect)
}

// Returns a Boolean value that indicates whether the path contains the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520716-containspoint?language=objc
func (b_ BezierPath) ContainsPoint(point foundation.Point) bool {
	rv := objc.Call[bool](b_, objc.Sel("containsPoint:"), point)
	return rv
}

// Removes all path elements from the path, effectively clearing the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520668-removeallpoints?language=objc
func (b_ BezierPath) RemoveAllPoints() {
	objc.Call[objc.Void](b_, objc.Sel("removeAllPoints"))
}

// Appends a rectangular path to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520670-appendbezierpathwithrect?language=objc
func (b_ BezierPath) AppendBezierPathWithRect(rect foundation.Rect) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithRect:"), rect)
}

// Creates and returns a new Bézier path object initialized with a rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520664-bezierpathwithrect?language=objc
func (bc _BezierPathClass) BezierPathWithRect(rect foundation.Rect) BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("bezierPathWithRect:"), rect)
	return rv
}

// Creates and returns a new Bézier path object initialized with a rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520664-bezierpathwithrect?language=objc
func BezierPath_BezierPathWithRect(rect foundation.Rect) BezierPath {
	return BezierPathClass.BezierPathWithRect(rect)
}

// Adds a Bezier cubic curve to the path from the current point to a new location, which is specified as a relative distance from the current point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520714-relativecurvetopoint?language=objc
func (b_ BezierPath) RelativeCurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("relativeCurveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

// Appends a straight line to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520742-linetopoint?language=objc
func (b_ BezierPath) LineToPoint(point foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("lineToPoint:"), point)
}

// Returns the line-stroking pattern for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520636-getlinedash?language=objc
func (b_ BezierPath) GetLineDashCountPhase(pattern *float64, count *int, phase *float64) {
	objc.Call[objc.Void](b_, objc.Sel("getLineDash:count:phase:"), pattern, count, phase)
}

// Moves the path’s current point to a new point whose location is the specified distance from the current point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520724-relativemovetopoint?language=objc
func (b_ BezierPath) RelativeMoveToPoint(point foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("relativeMoveToPoint:"), point)
}

// Appends an arc to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520737-appendbezierpathwitharcfrompoint?language=objc
func (b_ BezierPath) AppendBezierPathWithArcFromPointToPointRadius(point1 foundation.Point, point2 foundation.Point, radius float64) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithArcFromPoint:toPoint:radius:"), point1, point2, radius)
}

// Creates and returns a new Bézier path object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520652-bezierpath?language=objc
func (bc _BezierPathClass) BezierPath() BezierPath {
	rv := objc.Call[BezierPath](bc, objc.Sel("bezierPath"))
	return rv
}

// Creates and returns a new Bézier path object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520652-bezierpath?language=objc
func BezierPath_BezierPath() BezierPath {
	return BezierPathClass.BezierPath()
}

// Adds a Bezier cubic curve to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520628-curvetopoint?language=objc
func (b_ BezierPath) CurveToPointControlPoint1ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("curveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

// Appends the contents of the specified path object to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520688-appendbezierpath?language=objc
func (b_ BezierPath) AppendBezierPath(path IBezierPath) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPath:"), objc.Ptr(path))
}

// Draws a line along the path using the current stroke color and drawing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520739-stroke?language=objc
func (b_ BezierPath) Stroke() {
	objc.Call[objc.Void](b_, objc.Sel("stroke"))
}

// Appends an arc of a circle to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520680-appendbezierpathwitharcwithcente?language=objc
func (b_ BezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:"), center, radius, startAngle, endAngle)
}

// Closes the most recently added subpath. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520640-closepath?language=objc
func (b_ BezierPath) ClosePath() {
	objc.Call[objc.Void](b_, objc.Sel("closePath"))
}

// Strokes a line between two points using the current stroke color and the default drawing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520626-strokelinefrompoint?language=objc
func (bc _BezierPathClass) StrokeLineFromPointToPoint(point1 foundation.Point, point2 foundation.Point) {
	objc.Call[objc.Void](bc, objc.Sel("strokeLineFromPoint:toPoint:"), point1, point2)
}

// Strokes a line between two points using the current stroke color and the default drawing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520626-strokelinefrompoint?language=objc
func BezierPath_StrokeLineFromPointToPoint(point1 foundation.Point, point2 foundation.Point) {
	BezierPathClass.StrokeLineFromPointToPoint(point1, point2)
}

// Appends a straight line segment to the path starting at the current point and moving towards the specified point, relative to the current location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520709-relativelinetopoint?language=objc
func (b_ BezierPath) RelativeLineToPoint(point foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("relativeLineToPoint:"), point)
}

// Moves the path’s current point to the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520684-movetopoint?language=objc
func (b_ BezierPath) MoveToPoint(point foundation.Point) {
	objc.Call[objc.Void](b_, objc.Sel("moveToPoint:"), point)
}

// Appends a series of line segments to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520749-appendbezierpathwithpoints?language=objc
func (b_ BezierPath) AppendBezierPathWithPointsCount(points foundation.PointArray, count int) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithPoints:count:"), points, count)
}

// Fills the specified rectangular path with the current fill color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520747-fillrect?language=objc
func (bc _BezierPathClass) FillRect(rect foundation.Rect) {
	objc.Call[objc.Void](bc, objc.Sel("fillRect:"), rect)
}

// Fills the specified rectangular path with the current fill color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520747-fillrect?language=objc
func BezierPath_FillRect(rect foundation.Rect) {
	BezierPathClass.FillRect(rect)
}

// Sets the line-stroking pattern for the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520730-setlinedash?language=objc
func (b_ BezierPath) SetLineDashCountPhase(pattern *float64, count int, phase float64) {
	objc.Call[objc.Void](b_, objc.Sel("setLineDash:count:phase:"), pattern, count, phase)
}

// Changes the points associated with the specified path element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520671-setassociatedpoints?language=objc
func (b_ BezierPath) SetAssociatedPointsAtIndex(points foundation.PointArray, index int) {
	objc.Call[objc.Void](b_, objc.Sel("setAssociatedPoints:atIndex:"), points, index)
}

// Strokes the path of the specified rectangle using the current stroke color and the default drawing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520741-strokerect?language=objc
func (bc _BezierPathClass) StrokeRect(rect foundation.Rect) {
	objc.Call[objc.Void](bc, objc.Sel("strokeRect:"), rect)
}

// Strokes the path of the specified rectangle using the current stroke color and the default drawing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520741-strokerect?language=objc
func BezierPath_StrokeRect(rect foundation.Rect) {
	BezierPathClass.StrokeRect(rect)
}

// Intersects the specified rectangle with the clipping path of the current graphics context and makes the resulting shape the current clipping path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520694-cliprect?language=objc
func (bc _BezierPathClass) ClipRect(rect foundation.Rect) {
	objc.Call[objc.Void](bc, objc.Sel("clipRect:"), rect)
}

// Intersects the specified rectangle with the clipping path of the current graphics context and makes the resulting shape the current clipping path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520694-cliprect?language=objc
func BezierPath_ClipRect(rect foundation.Rect) {
	BezierPathClass.ClipRect(rect)
}

// Paints the region enclosed by the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520700-fill?language=objc
func (b_ BezierPath) Fill() {
	objc.Call[objc.Void](b_, objc.Sel("fill"))
}

// Draws a set of packed glyphs at the specified point in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520630-drawpackedglyphs?language=objc
func (bc _BezierPathClass) DrawPackedGlyphsAtPoint(packedGlyphs *uint8, point foundation.Point) {
	objc.Call[objc.Void](bc, objc.Sel("drawPackedGlyphs:atPoint:"), packedGlyphs, point)
}

// Draws a set of packed glyphs at the specified point in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520630-drawpackedglyphs?language=objc
func BezierPath_DrawPackedGlyphsAtPoint(packedGlyphs *uint8, point foundation.Point) {
	BezierPathClass.DrawPackedGlyphsAtPoint(packedGlyphs, point)
}

// Appends an outline of the specified glyph to the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/2887184-appendbezierpathwithcgglyph?language=objc
func (b_ BezierPath) AppendBezierPathWithCGGlyphInFont(glyph coregraphics.Glyph, font IFont) {
	objc.Call[objc.Void](b_, objc.Sel("appendBezierPathWithCGGlyph:inFont:"), glyph, objc.Ptr(font))
}

// Returns the default miter limit for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520728-defaultmiterlimit?language=objc
func (bc _BezierPathClass) DefaultMiterLimit() float64 {
	rv := objc.Call[float64](bc, objc.Sel("defaultMiterLimit"))
	return rv
}

// Returns the default miter limit for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520728-defaultmiterlimit?language=objc
func BezierPath_DefaultMiterLimit() float64 {
	return BezierPathClass.DefaultMiterLimit()
}

// Returns the default miter limit for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520728-defaultmiterlimit?language=objc
func (bc _BezierPathClass) SetDefaultMiterLimit(value float64) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultMiterLimit:"), value)
}

// Returns the default miter limit for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520728-defaultmiterlimit?language=objc
func BezierPath_SetDefaultMiterLimit(value float64) {
	BezierPathClass.SetDefaultMiterLimit(value)
}

// The width of stroked path lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520655-linewidth?language=objc
func (b_ BezierPath) LineWidth() float64 {
	rv := objc.Call[float64](b_, objc.Sel("lineWidth"))
	return rv
}

// The width of stroked path lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520655-linewidth?language=objc
func (b_ BezierPath) SetLineWidth(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setLineWidth:"), value)
}

// A Boolean value that indicates whether the path is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520712-empty?language=objc
func (b_ BezierPath) IsEmpty() bool {
	rv := objc.Call[bool](b_, objc.Sel("isEmpty"))
	return rv
}

// A path containing the reversed contents of the current path object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520656-bezierpathbyreversingpath?language=objc
func (b_ BezierPath) BezierPathByReversingPath() BezierPath {
	rv := objc.Call[BezierPath](b_, objc.Sel("bezierPathByReversingPath"))
	return rv
}

// Returns the default line width for the all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520644-defaultlinewidth?language=objc
func (bc _BezierPathClass) DefaultLineWidth() float64 {
	rv := objc.Call[float64](bc, objc.Sel("defaultLineWidth"))
	return rv
}

// Returns the default line width for the all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520644-defaultlinewidth?language=objc
func BezierPath_DefaultLineWidth() float64 {
	return BezierPathClass.DefaultLineWidth()
}

// Returns the default line width for the all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520644-defaultlinewidth?language=objc
func (bc _BezierPathClass) SetDefaultLineWidth(value float64) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultLineWidth:"), value)
}

// Returns the default line width for the all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520644-defaultlinewidth?language=objc
func BezierPath_SetDefaultLineWidth(value float64) {
	BezierPathClass.SetDefaultLineWidth(value)
}

// The limit at which miter joins are converted to bevel joins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520740-miterlimit?language=objc
func (b_ BezierPath) MiterLimit() float64 {
	rv := objc.Call[float64](b_, objc.Sel("miterLimit"))
	return rv
}

// The limit at which miter joins are converted to bevel joins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520740-miterlimit?language=objc
func (b_ BezierPath) SetMiterLimit(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setMiterLimit:"), value)
}

// The accuracy with which curves are rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520676-flatness?language=objc
func (b_ BezierPath) Flatness() float64 {
	rv := objc.Call[float64](b_, objc.Sel("flatness"))
	return rv
}

// The accuracy with which curves are rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520676-flatness?language=objc
func (b_ BezierPath) SetFlatness(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setFlatness:"), value)
}

// Returns the default winding rule used to fill all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520632-defaultwindingrule?language=objc
func (bc _BezierPathClass) DefaultWindingRule() WindingRule {
	rv := objc.Call[WindingRule](bc, objc.Sel("defaultWindingRule"))
	return rv
}

// Returns the default winding rule used to fill all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520632-defaultwindingrule?language=objc
func BezierPath_DefaultWindingRule() WindingRule {
	return BezierPathClass.DefaultWindingRule()
}

// Returns the default winding rule used to fill all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520632-defaultwindingrule?language=objc
func (bc _BezierPathClass) SetDefaultWindingRule(value WindingRule) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultWindingRule:"), value)
}

// Returns the default winding rule used to fill all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520632-defaultwindingrule?language=objc
func BezierPath_SetDefaultWindingRule(value WindingRule) {
	BezierPathClass.SetDefaultWindingRule(value)
}

// The line join style for the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520726-linejoinstyle?language=objc
func (b_ BezierPath) LineJoinStyle() LineJoinStyle {
	rv := objc.Call[LineJoinStyle](b_, objc.Sel("lineJoinStyle"))
	return rv
}

// The line join style for the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520726-linejoinstyle?language=objc
func (b_ BezierPath) SetLineJoinStyle(value LineJoinStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setLineJoinStyle:"), value)
}

// The bounding box of the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520722-bounds?language=objc
func (b_ BezierPath) Bounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("bounds"))
	return rv
}

// Returns the default line cap style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520678-defaultlinecapstyle?language=objc
func (bc _BezierPathClass) DefaultLineCapStyle() LineCapStyle {
	rv := objc.Call[LineCapStyle](bc, objc.Sel("defaultLineCapStyle"))
	return rv
}

// Returns the default line cap style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520678-defaultlinecapstyle?language=objc
func BezierPath_DefaultLineCapStyle() LineCapStyle {
	return BezierPathClass.DefaultLineCapStyle()
}

// Returns the default line cap style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520678-defaultlinecapstyle?language=objc
func (bc _BezierPathClass) SetDefaultLineCapStyle(value LineCapStyle) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultLineCapStyle:"), value)
}

// Returns the default line cap style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520678-defaultlinecapstyle?language=objc
func BezierPath_SetDefaultLineCapStyle(value LineCapStyle) {
	BezierPathClass.SetDefaultLineCapStyle(value)
}

// Returns the default line join style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520648-defaultlinejoinstyle?language=objc
func (bc _BezierPathClass) DefaultLineJoinStyle() LineJoinStyle {
	rv := objc.Call[LineJoinStyle](bc, objc.Sel("defaultLineJoinStyle"))
	return rv
}

// Returns the default line join style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520648-defaultlinejoinstyle?language=objc
func BezierPath_DefaultLineJoinStyle() LineJoinStyle {
	return BezierPathClass.DefaultLineJoinStyle()
}

// Returns the default line join style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520648-defaultlinejoinstyle?language=objc
func (bc _BezierPathClass) SetDefaultLineJoinStyle(value LineJoinStyle) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultLineJoinStyle:"), value)
}

// Returns the default line join style for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520648-defaultlinejoinstyle?language=objc
func BezierPath_SetDefaultLineJoinStyle(value LineJoinStyle) {
	BezierPathClass.SetDefaultLineJoinStyle(value)
}

// The line cap style for the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520667-linecapstyle?language=objc
func (b_ BezierPath) LineCapStyle() LineCapStyle {
	rv := objc.Call[LineCapStyle](b_, objc.Sel("lineCapStyle"))
	return rv
}

// The line cap style for the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520667-linecapstyle?language=objc
func (b_ BezierPath) SetLineCapStyle(value LineCapStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setLineCapStyle:"), value)
}

// The default flatness value for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520720-defaultflatness?language=objc
func (bc _BezierPathClass) DefaultFlatness() float64 {
	rv := objc.Call[float64](bc, objc.Sel("defaultFlatness"))
	return rv
}

// The default flatness value for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520720-defaultflatness?language=objc
func BezierPath_DefaultFlatness() float64 {
	return BezierPathClass.DefaultFlatness()
}

// The default flatness value for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520720-defaultflatness?language=objc
func (bc _BezierPathClass) SetDefaultFlatness(value float64) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultFlatness:"), value)
}

// The default flatness value for all paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520720-defaultflatness?language=objc
func BezierPath_SetDefaultFlatness(value float64) {
	BezierPathClass.SetDefaultFlatness(value)
}

// The bounding box of the path, including any control points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520654-controlpointbounds?language=objc
func (b_ BezierPath) ControlPointBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("controlPointBounds"))
	return rv
}

// The total number of path elements in the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520645-elementcount?language=objc
func (b_ BezierPath) ElementCount() int {
	rv := objc.Call[int](b_, objc.Sel("elementCount"))
	return rv
}

// A flattened version of the path object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520733-bezierpathbyflatteningpath?language=objc
func (b_ BezierPath) BezierPathByFlatteningPath() BezierPath {
	rv := objc.Call[BezierPath](b_, objc.Sel("bezierPathByFlatteningPath"))
	return rv
}

// The current point (the trailing point or ending point in the most recently added segment). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520698-currentpoint?language=objc
func (b_ BezierPath) CurrentPoint() foundation.Point {
	rv := objc.Call[foundation.Point](b_, objc.Sel("currentPoint"))
	return rv
}

// The winding rule used to fill the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520657-windingrule?language=objc
func (b_ BezierPath) WindingRule() WindingRule {
	rv := objc.Call[WindingRule](b_, objc.Sel("windingRule"))
	return rv
}

// The winding rule used to fill the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/1520657-windingrule?language=objc
func (b_ BezierPath) SetWindingRule(value WindingRule) {
	objc.Call[objc.Void](b_, objc.Sel("setWindingRule:"), value)
}
