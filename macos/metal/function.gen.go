// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object that represents a public shader function in a Metal library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction?language=objc
type PFunction interface {
	// optional
	NewArgumentEncoderWithBufferIndex(bufferIndex uint) PArgumentEncoder
	HasNewArgumentEncoderWithBufferIndex() bool

	// optional
	PatchControlPointCount() int
	HasPatchControlPointCount() bool

	// optional
	PatchType() PatchType
	HasPatchType() bool

	// optional
	Options() FunctionOptions
	HasOptions() bool

	// optional
	Name() string
	HasName() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	StageInputAttributes() []IAttribute
	HasStageInputAttributes() bool

	// optional
	VertexAttributes() []IVertexAttribute
	HasVertexAttributes() bool

	// optional
	FunctionConstantsDictionary() map[string]IFunctionConstant
	HasFunctionConstantsDictionary() bool

	// optional
	FunctionType() FunctionType
	HasFunctionType() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PFunction] protocol.
type FunctionWrapper struct {
	objc.Object
}

func (f_ FunctionWrapper) HasNewArgumentEncoderWithBufferIndex() bool {
	return f_.RespondsToSelector(objc.Sel("newArgumentEncoderWithBufferIndex:"))
}

// Creates an argument encoder for an argument buffer that’s one of this function's arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/2915781-newargumentencoderwithbufferinde?language=objc
func (f_ FunctionWrapper) NewArgumentEncoderWithBufferIndex(bufferIndex uint) ArgumentEncoderWrapper {
	rv := objc.Call[ArgumentEncoderWrapper](f_, objc.Sel("newArgumentEncoderWithBufferIndex:"), bufferIndex)
	rv.Autorelease()
	return rv
}

func (f_ FunctionWrapper) HasPatchControlPointCount() bool {
	return f_.RespondsToSelector(objc.Sel("patchControlPointCount"))
}

// The number of patch control points in the post-tessellation vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1639890-patchcontrolpointcount?language=objc
func (f_ FunctionWrapper) PatchControlPointCount() int {
	rv := objc.Call[int](f_, objc.Sel("patchControlPointCount"))
	return rv
}

func (f_ FunctionWrapper) HasPatchType() bool {
	return f_.RespondsToSelector(objc.Sel("patchType"))
}

// The tessellation patch type of a post-tessellation vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1639909-patchtype?language=objc
func (f_ FunctionWrapper) PatchType() PatchType {
	rv := objc.Call[PatchType](f_, objc.Sel("patchType"))
	return rv
}

func (f_ FunctionWrapper) HasOptions() bool {
	return f_.RespondsToSelector(objc.Sel("options"))
}

// The options that Metal used to compile this function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/3566558-options?language=objc
func (f_ FunctionWrapper) Options() FunctionOptions {
	rv := objc.Call[FunctionOptions](f_, objc.Sel("options"))
	return rv
}

func (f_ FunctionWrapper) HasName() bool {
	return f_.RespondsToSelector(objc.Sel("name"))
}

// The function’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1515424-name?language=objc
func (f_ FunctionWrapper) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

func (f_ FunctionWrapper) HasDevice() bool {
	return f_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1515496-device?language=objc
func (f_ FunctionWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](f_, objc.Sel("device"))
	return rv
}

func (f_ FunctionWrapper) HasStageInputAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("stageInputAttributes"))
}

// An array that describes the input attributes to the function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/2097159-stageinputattributes?language=objc
func (f_ FunctionWrapper) StageInputAttributes() []Attribute {
	rv := objc.Call[[]Attribute](f_, objc.Sel("stageInputAttributes"))
	return rv
}

func (f_ FunctionWrapper) HasVertexAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("vertexAttributes"))
}

// An array that describes the vertex input attributes to a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1515944-vertexattributes?language=objc
func (f_ FunctionWrapper) VertexAttributes() []VertexAttribute {
	rv := objc.Call[[]VertexAttribute](f_, objc.Sel("vertexAttributes"))
	return rv
}

func (f_ FunctionWrapper) HasFunctionConstantsDictionary() bool {
	return f_.RespondsToSelector(objc.Sel("functionConstantsDictionary"))
}

// A dictionary of function constants for a specialized function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/2314777-functionconstantsdictionary?language=objc
func (f_ FunctionWrapper) FunctionConstantsDictionary() map[string]FunctionConstant {
	rv := objc.Call[map[string]FunctionConstant](f_, objc.Sel("functionConstantsDictionary"))
	return rv
}

func (f_ FunctionWrapper) HasFunctionType() bool {
	return f_.RespondsToSelector(objc.Sel("functionType"))
}

// The shader function’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1516042-functiontype?language=objc
func (f_ FunctionWrapper) FunctionType() FunctionType {
	rv := objc.Call[FunctionType](f_, objc.Sel("functionType"))
	return rv
}

func (f_ FunctionWrapper) HasSetLabel() bool {
	return f_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1640034-label?language=objc
func (f_ FunctionWrapper) SetLabel(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setLabel:"), value)
}

func (f_ FunctionWrapper) HasLabel() bool {
	return f_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/1640034-label?language=objc
func (f_ FunctionWrapper) Label() string {
	rv := objc.Call[string](f_, objc.Sel("label"))
	return rv
}
