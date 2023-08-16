// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a mesh generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator?language=objc
type PMeshGenerator interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetMesh(value []objc.Object)
	HasSetMesh() bool

	// optional
	Mesh() []objc.IObject
	HasMesh() bool
}

// A concrete type wrapper for the [PMeshGenerator] protocol.
type MeshGeneratorWrapper struct {
	objc.Object
}

func (m_ MeshGeneratorWrapper) HasSetColor() bool {
	return m_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the rendered mesh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228559-color?language=objc
func (m_ MeshGeneratorWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](m_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (m_ MeshGeneratorWrapper) HasColor() bool {
	return m_.RespondsToSelector(objc.Sel("color"))
}

// The color of the rendered mesh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228559-color?language=objc
func (m_ MeshGeneratorWrapper) Color() Color {
	rv := objc.Call[Color](m_, objc.Sel("color"))
	return rv
}

func (m_ MeshGeneratorWrapper) HasSetWidth() bool {
	return m_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228561-width?language=objc
func (m_ MeshGeneratorWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setWidth:"), value)
}

func (m_ MeshGeneratorWrapper) HasWidth() bool {
	return m_.RespondsToSelector(objc.Sel("width"))
}

// The width of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228561-width?language=objc
func (m_ MeshGeneratorWrapper) Width() float64 {
	rv := objc.Call[float64](m_, objc.Sel("width"))
	return rv
}

func (m_ MeshGeneratorWrapper) HasSetMesh() bool {
	return m_.RespondsToSelector(objc.Sel("setMesh:"))
}

// An array that describes the mesh to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228560-mesh?language=objc
func (m_ MeshGeneratorWrapper) SetMesh(value []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setMesh:"), value)
}

func (m_ MeshGeneratorWrapper) HasMesh() bool {
	return m_.RespondsToSelector(objc.Sel("mesh"))
}

// An array that describes the mesh to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimeshgenerator/3228560-mesh?language=objc
func (m_ MeshGeneratorWrapper) Mesh() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("mesh"))
	return rv
}
