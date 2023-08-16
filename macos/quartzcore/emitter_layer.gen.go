// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EmitterLayer] class.
var EmitterLayerClass = _EmitterLayerClass{objc.GetClass("CAEmitterLayer")}

type _EmitterLayerClass struct {
	objc.Class
}

// An interface definition for the [EmitterLayer] class.
type IEmitterLayer interface {
	ILayer
	Scale() float64
	SetScale(value float64)
	EmitterMode() EmitterLayerEmitterMode
	SetEmitterMode(value EmitterLayerEmitterMode)
	BirthRate() float64
	SetBirthRate(value float64)
	EmitterZPosition() float64
	SetEmitterZPosition(value float64)
	EmitterDepth() float64
	SetEmitterDepth(value float64)
	EmitterSize() coregraphics.Size
	SetEmitterSize(value coregraphics.Size)
	EmitterPosition() coregraphics.Point
	SetEmitterPosition(value coregraphics.Point)
	Lifetime() float64
	SetLifetime(value float64)
	EmitterCells() []EmitterCell
	SetEmitterCells(value []IEmitterCell)
	Spin() float64
	SetSpin(value float64)
	Velocity() float64
	SetVelocity(value float64)
	PreservesDepth() bool
	SetPreservesDepth(value bool)
	RenderMode() EmitterLayerRenderMode
	SetRenderMode(value EmitterLayerRenderMode)
	Seed() int
	SetSeed(value int)
	EmitterShape() EmitterLayerEmitterShape
	SetEmitterShape(value EmitterLayerEmitterShape)
}

// A layer that emits, animates, and renders a particle system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer?language=objc
type EmitterLayer struct {
	Layer
}

func EmitterLayerFrom(ptr unsafe.Pointer) EmitterLayer {
	return EmitterLayer{
		Layer: LayerFrom(ptr),
	}
}

func (ec _EmitterLayerClass) Alloc() EmitterLayer {
	rv := objc.Call[EmitterLayer](ec, objc.Sel("alloc"))
	return rv
}

func EmitterLayer_Alloc() EmitterLayer {
	return EmitterLayerClass.Alloc()
}

func (ec _EmitterLayerClass) New() EmitterLayer {
	rv := objc.Call[EmitterLayer](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEmitterLayer() EmitterLayer {
	return EmitterLayerClass.New()
}

func (e_ EmitterLayer) Init() EmitterLayer {
	rv := objc.Call[EmitterLayer](e_, objc.Sel("init"))
	return rv
}

func (ec _EmitterLayerClass) Layer() EmitterLayer {
	rv := objc.Call[EmitterLayer](ec, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func EmitterLayer_Layer() EmitterLayer {
	return EmitterLayerClass.Layer()
}

func (e_ EmitterLayer) InitWithLayer(layer objc.IObject) EmitterLayer {
	rv := objc.Call[EmitterLayer](e_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func EmitterLayer_InitWithLayer(layer objc.IObject) EmitterLayer {
	return EmitterLayerClass.Alloc().InitWithLayer(layer)
}

func (e_ EmitterLayer) ModelLayer() EmitterLayer {
	rv := objc.Call[EmitterLayer](e_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func EmitterLayer_ModelLayer() EmitterLayer {
	return EmitterLayerClass.Alloc().ModelLayer()
}

func (e_ EmitterLayer) PresentationLayer() EmitterLayer {
	rv := objc.Call[EmitterLayer](e_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func EmitterLayer_PresentationLayer() EmitterLayer {
	return EmitterLayerClass.Alloc().PresentationLayer()
}

// Defines a multiplier applied to the cell-defined particle scale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521841-scale?language=objc
func (e_ EmitterLayer) Scale() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scale"))
	return rv
}

// Defines a multiplier applied to the cell-defined particle scale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521841-scale?language=objc
func (e_ EmitterLayer) SetScale(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setScale:"), value)
}

// Specifies the emitter mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522128-emittermode?language=objc
func (e_ EmitterLayer) EmitterMode() EmitterLayerEmitterMode {
	rv := objc.Call[EmitterLayerEmitterMode](e_, objc.Sel("emitterMode"))
	return rv
}

// Specifies the emitter mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522128-emittermode?language=objc
func (e_ EmitterLayer) SetEmitterMode(value EmitterLayerEmitterMode) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterMode:"), value)
}

// Defines a multiplier that is applied to the cell-defined birth rate. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521976-birthrate?language=objc
func (e_ EmitterLayer) BirthRate() float64 {
	rv := objc.Call[float64](e_, objc.Sel("birthRate"))
	return rv
}

// Defines a multiplier that is applied to the cell-defined birth rate. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521976-birthrate?language=objc
func (e_ EmitterLayer) SetBirthRate(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setBirthRate:"), value)
}

// Specifies the center of the particle emitter shape along the z-axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522169-emitterzposition?language=objc
func (e_ EmitterLayer) EmitterZPosition() float64 {
	rv := objc.Call[float64](e_, objc.Sel("emitterZPosition"))
	return rv
}

// Specifies the center of the particle emitter shape along the z-axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522169-emitterzposition?language=objc
func (e_ EmitterLayer) SetEmitterZPosition(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterZPosition:"), value)
}

// Determines the depth of the emitter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521844-emitterdepth?language=objc
func (e_ EmitterLayer) EmitterDepth() float64 {
	rv := objc.Call[float64](e_, objc.Sel("emitterDepth"))
	return rv
}

// Determines the depth of the emitter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521844-emitterdepth?language=objc
func (e_ EmitterLayer) SetEmitterDepth(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterDepth:"), value)
}

// Determines the size of the particle emitter shape. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521869-emittersize?language=objc
func (e_ EmitterLayer) EmitterSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](e_, objc.Sel("emitterSize"))
	return rv
}

// Determines the size of the particle emitter shape. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521869-emittersize?language=objc
func (e_ EmitterLayer) SetEmitterSize(value coregraphics.Size) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterSize:"), value)
}

// The position of the center of the particle emitter. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522289-emitterposition?language=objc
func (e_ EmitterLayer) EmitterPosition() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](e_, objc.Sel("emitterPosition"))
	return rv
}

// The position of the center of the particle emitter. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522289-emitterposition?language=objc
func (e_ EmitterLayer) SetEmitterPosition(value coregraphics.Point) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterPosition:"), value)
}

// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522144-lifetime?language=objc
func (e_ EmitterLayer) Lifetime() float64 {
	rv := objc.Call[float64](e_, objc.Sel("lifetime"))
	return rv
}

// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522144-lifetime?language=objc
func (e_ EmitterLayer) SetLifetime(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setLifetime:"), value)
}

// The array emitter cells attached to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521923-emittercells?language=objc
func (e_ EmitterLayer) EmitterCells() []EmitterCell {
	rv := objc.Call[[]EmitterCell](e_, objc.Sel("emitterCells"))
	return rv
}

// The array emitter cells attached to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521923-emittercells?language=objc
func (e_ EmitterLayer) SetEmitterCells(value []IEmitterCell) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterCells:"), value)
}

// Defines a multiplier applied to the cell-defined particle spin. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521861-spin?language=objc
func (e_ EmitterLayer) Spin() float64 {
	rv := objc.Call[float64](e_, objc.Sel("spin"))
	return rv
}

// Defines a multiplier applied to the cell-defined particle spin. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521861-spin?language=objc
func (e_ EmitterLayer) SetSpin(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setSpin:"), value)
}

// Defines a multiplier applied to the cell-defined particle velocity. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522015-velocity?language=objc
func (e_ EmitterLayer) Velocity() float64 {
	rv := objc.Call[float64](e_, objc.Sel("velocity"))
	return rv
}

// Defines a multiplier applied to the cell-defined particle velocity. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522015-velocity?language=objc
func (e_ EmitterLayer) SetVelocity(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setVelocity:"), value)
}

// Defines whether the layer flattens the particles into its plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521872-preservesdepth?language=objc
func (e_ EmitterLayer) PreservesDepth() bool {
	rv := objc.Call[bool](e_, objc.Sel("preservesDepth"))
	return rv
}

// Defines whether the layer flattens the particles into its plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521872-preservesdepth?language=objc
func (e_ EmitterLayer) SetPreservesDepth(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setPreservesDepth:"), value)
}

// Defines how particle cells are rendered into the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522104-rendermode?language=objc
func (e_ EmitterLayer) RenderMode() EmitterLayerRenderMode {
	rv := objc.Call[EmitterLayerRenderMode](e_, objc.Sel("renderMode"))
	return rv
}

// Defines how particle cells are rendered into the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522104-rendermode?language=objc
func (e_ EmitterLayer) SetRenderMode(value EmitterLayerRenderMode) {
	objc.Call[objc.Void](e_, objc.Sel("setRenderMode:"), value)
}

// Specifies the seed used to initialize the random number generator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522079-seed?language=objc
func (e_ EmitterLayer) Seed() int {
	rv := objc.Call[int](e_, objc.Sel("seed"))
	return rv
}

// Specifies the seed used to initialize the random number generator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1522079-seed?language=objc
func (e_ EmitterLayer) SetSeed(value int) {
	objc.Call[objc.Void](e_, objc.Sel("setSeed:"), value)
}

// Specifies the emitter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521919-emittershape?language=objc
func (e_ EmitterLayer) EmitterShape() EmitterLayerEmitterShape {
	rv := objc.Call[EmitterLayerEmitterShape](e_, objc.Sel("emitterShape"))
	return rv
}

// Specifies the emitter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/1521919-emittershape?language=objc
func (e_ EmitterLayer) SetEmitterShape(value EmitterLayerEmitterShape) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterShape:"), value)
}
