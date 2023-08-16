// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GraphicsContext] class.
var GraphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}

type _GraphicsContextClass struct {
	objc.Class
}

// An interface definition for the [GraphicsContext] class.
type IGraphicsContext interface {
	objc.IObject
	FlushGraphics()
	PatternPhase() foundation.Point
	SetPatternPhase(value foundation.Point)
	CIContext() coreimage.Context
	ImageInterpolation() ImageInterpolation
	SetImageInterpolation(value ImageInterpolation)
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	CGContext() coregraphics.ContextRef
	CompositingOperation() CompositingOperation
	SetCompositingOperation(value CompositingOperation)
	ColorRenderingIntent() ColorRenderingIntent
	SetColorRenderingIntent(value ColorRenderingIntent)
	IsDrawingToScreen() bool
	IsFlipped() bool
	Attributes() map[GraphicsContextAttributeKey]objc.Object
}

// An object that represents a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext?language=objc
type GraphicsContext struct {
	objc.Object
}

func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("alloc"))
	return rv
}

func GraphicsContext_Alloc() GraphicsContext {
	return GraphicsContextClass.Alloc()
}

func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGraphicsContext() GraphicsContext {
	return GraphicsContextClass.New()
}

func (g_ GraphicsContext) Init() GraphicsContext {
	rv := objc.Call[GraphicsContext](g_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the current context is drawing to the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1525944-currentcontextdrawingtoscreen?language=objc
func (gc _GraphicsContextClass) CurrentContextDrawingToScreen() bool {
	rv := objc.Call[bool](gc, objc.Sel("currentContextDrawingToScreen"))
	return rv
}

// Returns a Boolean value that indicates whether the current context is drawing to the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1525944-currentcontextdrawingtoscreen?language=objc
func GraphicsContext_CurrentContextDrawingToScreen() bool {
	return GraphicsContextClass.CurrentContextDrawingToScreen()
}

// Creates a new graphics context for drawing into a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1532883-graphicscontextwithwindow?language=objc
func (gc _GraphicsContextClass) GraphicsContextWithWindow(window IWindow) GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("graphicsContextWithWindow:"), objc.Ptr(window))
	return rv
}

// Creates a new graphics context for drawing into a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1532883-graphicscontextwithwindow?language=objc
func GraphicsContext_GraphicsContextWithWindow(window IWindow) GraphicsContext {
	return GraphicsContextClass.GraphicsContextWithWindow(window)
}

// Forces any buffered operations or data to be sent to the graphics context’s destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1527919-flushgraphics?language=objc
func (g_ GraphicsContext) FlushGraphics() {
	objc.Call[objc.Void](g_, objc.Sel("flushGraphics"))
}

// Creates a new graphics context from the specified Core Graphics context and the initial flipped state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535380-graphicscontextwithcgcontext?language=objc
func (gc _GraphicsContextClass) GraphicsContextWithCGContextFlipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("graphicsContextWithCGContext:flipped:"), graphicsPort, initialFlippedState)
	return rv
}

// Creates a new graphics context from the specified Core Graphics context and the initial flipped state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535380-graphicscontextwithcgcontext?language=objc
func GraphicsContext_GraphicsContextWithCGContextFlipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	return GraphicsContextClass.GraphicsContextWithCGContextFlipped(graphicsPort, initialFlippedState)
}

// Pops a graphics context from the per-thread stack, makes it current, and sends the context a restore graphics state message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524294-restoregraphicsstate?language=objc
func (gc _GraphicsContextClass) RestoreGraphicsState() {
	objc.Call[objc.Void](gc, objc.Sel("restoreGraphicsState"))
}

// Pops a graphics context from the per-thread stack, makes it current, and sends the context a restore graphics state message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524294-restoregraphicsstate?language=objc
func GraphicsContext_RestoreGraphicsState() {
	GraphicsContextClass.RestoreGraphicsState()
}

// Creates a new graphics context using the specified bipmap image representation object as the context destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529827-graphicscontextwithbitmapimagere?language=objc
func (gc _GraphicsContextClass) GraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("graphicsContextWithBitmapImageRep:"), objc.Ptr(bitmapRep))
	return rv
}

// Creates a new graphics context using the specified bipmap image representation object as the context destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529827-graphicscontextwithbitmapimagere?language=objc
func GraphicsContext_GraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) GraphicsContext {
	return GraphicsContextClass.GraphicsContextWithBitmapImageRep(bitmapRep)
}

// Saves the graphics state of the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1526017-savegraphicsstate?language=objc
func (gc _GraphicsContextClass) SaveGraphicsState() {
	objc.Call[objc.Void](gc, objc.Sel("saveGraphicsState"))
}

// Saves the graphics state of the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1526017-savegraphicsstate?language=objc
func GraphicsContext_SaveGraphicsState() {
	GraphicsContextClass.SaveGraphicsState()
}

// Creates a graphics context using the specified attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524838-graphicscontextwithattributes?language=objc
func (gc _GraphicsContextClass) GraphicsContextWithAttributes(attributes map[GraphicsContextAttributeKey]objc.IObject) GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("graphicsContextWithAttributes:"), attributes)
	return rv
}

// Creates a graphics context using the specified attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524838-graphicscontextwithattributes?language=objc
func GraphicsContext_GraphicsContextWithAttributes(attributes map[GraphicsContextAttributeKey]objc.IObject) GraphicsContext {
	return GraphicsContextClass.GraphicsContextWithAttributes(attributes)
}

// The amount to offset the pattern color when filling the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1527510-patternphase?language=objc
func (g_ GraphicsContext) PatternPhase() foundation.Point {
	rv := objc.Call[foundation.Point](g_, objc.Sel("patternPhase"))
	return rv
}

// The amount to offset the pattern color when filling the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1527510-patternphase?language=objc
func (g_ GraphicsContext) SetPatternPhase(value foundation.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setPatternPhase:"), value)
}

// Returns the current graphics context of the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535352-currentcontext?language=objc
func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := objc.Call[GraphicsContext](gc, objc.Sel("currentContext"))
	return rv
}

// Returns the current graphics context of the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535352-currentcontext?language=objc
func GraphicsContext_CurrentContext() GraphicsContext {
	return GraphicsContextClass.CurrentContext()
}

// Returns the current graphics context of the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535352-currentcontext?language=objc
func (gc _GraphicsContextClass) SetCurrentContext(value IGraphicsContext) {
	objc.Call[objc.Void](gc, objc.Sel("setCurrentContext:"), objc.Ptr(value))
}

// Returns the current graphics context of the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1535352-currentcontext?language=objc
func GraphicsContext_SetCurrentContext(value IGraphicsContext) {
	GraphicsContextClass.SetCurrentContext(value)
}

// A context for Core Image objects that you can use to render into the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1534326-cicontext?language=objc
func (g_ GraphicsContext) CIContext() coreimage.Context {
	rv := objc.Call[coreimage.Context](g_, objc.Sel("CIContext"))
	return rv
}

// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529711-imageinterpolation?language=objc
func (g_ GraphicsContext) ImageInterpolation() ImageInterpolation {
	rv := objc.Call[ImageInterpolation](g_, objc.Sel("imageInterpolation"))
	return rv
}

// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529711-imageinterpolation?language=objc
func (g_ GraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	objc.Call[objc.Void](g_, objc.Sel("setImageInterpolation:"), value)
}

// A Boolean value that indicates whether the graphics context uses antialiasing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529486-shouldantialias?language=objc
func (g_ GraphicsContext) ShouldAntialias() bool {
	rv := objc.Call[bool](g_, objc.Sel("shouldAntialias"))
	return rv
}

// A Boolean value that indicates whether the graphics context uses antialiasing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1529486-shouldantialias?language=objc
func (g_ GraphicsContext) SetShouldAntialias(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setShouldAntialias:"), value)
}

// The Core Graphics context, which is a low-level, platform-specific graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524878-cgcontext?language=objc
func (g_ GraphicsContext) CGContext() coregraphics.ContextRef {
	rv := objc.Call[coregraphics.ContextRef](g_, objc.Sel("CGContext"))
	return rv
}

// The graphics context’s global compositing operation setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1525984-compositingoperation?language=objc
func (g_ GraphicsContext) CompositingOperation() CompositingOperation {
	rv := objc.Call[CompositingOperation](g_, objc.Sel("compositingOperation"))
	return rv
}

// The graphics context’s global compositing operation setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1525984-compositingoperation?language=objc
func (g_ GraphicsContext) SetCompositingOperation(value CompositingOperation) {
	objc.Call[objc.Void](g_, objc.Sel("setCompositingOperation:"), value)
}

// The color rendering intent in the graphics context’s graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1533527-colorrenderingintent?language=objc
func (g_ GraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	rv := objc.Call[ColorRenderingIntent](g_, objc.Sel("colorRenderingIntent"))
	return rv
}

// The color rendering intent in the graphics context’s graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1533527-colorrenderingintent?language=objc
func (g_ GraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	objc.Call[objc.Void](g_, objc.Sel("setColorRenderingIntent:"), value)
}

// A Boolean value that indicates whether the drawing destination is the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1524673-drawingtoscreen?language=objc
func (g_ GraphicsContext) IsDrawingToScreen() bool {
	rv := objc.Call[bool](g_, objc.Sel("isDrawingToScreen"))
	return rv
}

// A Boolean value that indicates the graphics context’s flipped state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1531568-flipped?language=objc
func (g_ GraphicsContext) IsFlipped() bool {
	rv := objc.Call[bool](g_, objc.Sel("isFlipped"))
	return rv
}

// The attributes used to create this instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/1528254-attributes?language=objc
func (g_ GraphicsContext) Attributes() map[GraphicsContextAttributeKey]objc.Object {
	rv := objc.Call[map[GraphicsContextAttributeKey]objc.Object](g_, objc.Sel("attributes"))
	return rv
}
