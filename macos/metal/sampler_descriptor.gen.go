// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SamplerDescriptor] class.
var SamplerDescriptorClass = _SamplerDescriptorClass{objc.GetClass("MTLSamplerDescriptor")}

type _SamplerDescriptorClass struct {
	objc.Class
}

// An interface definition for the [SamplerDescriptor] class.
type ISamplerDescriptor interface {
	objc.IObject
	NormalizedCoordinates() bool
	SetNormalizedCoordinates(value bool)
	MagFilter() SamplerMinMagFilter
	SetMagFilter(value SamplerMinMagFilter)
	LodMinClamp() float64
	SetLodMinClamp(value float64)
	MipFilter() SamplerMipFilter
	SetMipFilter(value SamplerMipFilter)
	MinFilter() SamplerMinMagFilter
	SetMinFilter(value SamplerMinMagFilter)
	TAddressMode() SamplerAddressMode
	SetTAddressMode(value SamplerAddressMode)
	SAddressMode() SamplerAddressMode
	SetSAddressMode(value SamplerAddressMode)
	LodMaxClamp() float64
	SetLodMaxClamp(value float64)
	RAddressMode() SamplerAddressMode
	SetRAddressMode(value SamplerAddressMode)
	BorderColor() SamplerBorderColor
	SetBorderColor(value SamplerBorderColor)
	CompareFunction() CompareFunction
	SetCompareFunction(value CompareFunction)
	Label() string
	SetLabel(value string)
	LodAverage() bool
	SetLodAverage(value bool)
	MaxAnisotropy() uint
	SetMaxAnisotropy(value uint)
	SupportArgumentBuffers() bool
	SetSupportArgumentBuffers(value bool)
}

// An object that you use to configure a texture sampler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor?language=objc
type SamplerDescriptor struct {
	objc.Object
}

func SamplerDescriptorFrom(ptr unsafe.Pointer) SamplerDescriptor {
	return SamplerDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SamplerDescriptorClass) Alloc() SamplerDescriptor {
	rv := objc.Call[SamplerDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func SamplerDescriptor_Alloc() SamplerDescriptor {
	return SamplerDescriptorClass.Alloc()
}

func (sc _SamplerDescriptorClass) New() SamplerDescriptor {
	rv := objc.Call[SamplerDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSamplerDescriptor() SamplerDescriptor {
	return SamplerDescriptorClass.New()
}

func (s_ SamplerDescriptor) Init() SamplerDescriptor {
	rv := objc.Call[SamplerDescriptor](s_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether texture coordinates are normalized to the range [0.0, 1.0]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516289-normalizedcoordinates?language=objc
func (s_ SamplerDescriptor) NormalizedCoordinates() bool {
	rv := objc.Call[bool](s_, objc.Sel("normalizedCoordinates"))
	return rv
}

// A Boolean value that indicates whether texture coordinates are normalized to the range [0.0, 1.0]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516289-normalizedcoordinates?language=objc
func (s_ SamplerDescriptor) SetNormalizedCoordinates(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setNormalizedCoordinates:"), value)
}

// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515926-magfilter?language=objc
func (s_ SamplerDescriptor) MagFilter() SamplerMinMagFilter {
	rv := objc.Call[SamplerMinMagFilter](s_, objc.Sel("magFilter"))
	return rv
}

// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515926-magfilter?language=objc
func (s_ SamplerDescriptor) SetMagFilter(value SamplerMinMagFilter) {
	objc.Call[objc.Void](s_, objc.Sel("setMagFilter:"), value)
}

// The minimum level of detail (LOD) to use when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515629-lodminclamp?language=objc
func (s_ SamplerDescriptor) LodMinClamp() float64 {
	rv := objc.Call[float64](s_, objc.Sel("lodMinClamp"))
	return rv
}

// The minimum level of detail (LOD) to use when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515629-lodminclamp?language=objc
func (s_ SamplerDescriptor) SetLodMinClamp(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLodMinClamp:"), value)
}

// The filtering option for combining pixels between two mipmap levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515553-mipfilter?language=objc
func (s_ SamplerDescriptor) MipFilter() SamplerMipFilter {
	rv := objc.Call[SamplerMipFilter](s_, objc.Sel("mipFilter"))
	return rv
}

// The filtering option for combining pixels between two mipmap levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515553-mipfilter?language=objc
func (s_ SamplerDescriptor) SetMipFilter(value SamplerMipFilter) {
	objc.Call[objc.Void](s_, objc.Sel("setMipFilter:"), value)
}

// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515792-minfilter?language=objc
func (s_ SamplerDescriptor) MinFilter() SamplerMinMagFilter {
	rv := objc.Call[SamplerMinMagFilter](s_, objc.Sel("minFilter"))
	return rv
}

// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515792-minfilter?language=objc
func (s_ SamplerDescriptor) SetMinFilter(value SamplerMinMagFilter) {
	objc.Call[objc.Void](s_, objc.Sel("setMinFilter:"), value)
}

// The address mode for the texture height (t) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515900-taddressmode?language=objc
func (s_ SamplerDescriptor) TAddressMode() SamplerAddressMode {
	rv := objc.Call[SamplerAddressMode](s_, objc.Sel("tAddressMode"))
	return rv
}

// The address mode for the texture height (t) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515900-taddressmode?language=objc
func (s_ SamplerDescriptor) SetTAddressMode(value SamplerAddressMode) {
	objc.Call[objc.Void](s_, objc.Sel("setTAddressMode:"), value)
}

// The address mode for the texture width (s) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515779-saddressmode?language=objc
func (s_ SamplerDescriptor) SAddressMode() SamplerAddressMode {
	rv := objc.Call[SamplerAddressMode](s_, objc.Sel("sAddressMode"))
	return rv
}

// The address mode for the texture width (s) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515779-saddressmode?language=objc
func (s_ SamplerDescriptor) SetSAddressMode(value SamplerAddressMode) {
	objc.Call[objc.Void](s_, objc.Sel("setSAddressMode:"), value)
}

// The maximum level of detail (LOD) to use when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516234-lodmaxclamp?language=objc
func (s_ SamplerDescriptor) LodMaxClamp() float64 {
	rv := objc.Call[float64](s_, objc.Sel("lodMaxClamp"))
	return rv
}

// The maximum level of detail (LOD) to use when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516234-lodmaxclamp?language=objc
func (s_ SamplerDescriptor) SetLodMaxClamp(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLodMaxClamp:"), value)
}

// The address mode for the texture depth (r) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515466-raddressmode?language=objc
func (s_ SamplerDescriptor) RAddressMode() SamplerAddressMode {
	rv := objc.Call[SamplerAddressMode](s_, objc.Sel("rAddressMode"))
	return rv
}

// The address mode for the texture depth (r) coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515466-raddressmode?language=objc
func (s_ SamplerDescriptor) SetRAddressMode(value SamplerAddressMode) {
	objc.Call[objc.Void](s_, objc.Sel("setRAddressMode:"), value)
}

// The border color for clamped texture values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/2092299-bordercolor?language=objc
func (s_ SamplerDescriptor) BorderColor() SamplerBorderColor {
	rv := objc.Call[SamplerBorderColor](s_, objc.Sel("borderColor"))
	return rv
}

// The border color for clamped texture values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/2092299-bordercolor?language=objc
func (s_ SamplerDescriptor) SetBorderColor(value SamplerBorderColor) {
	objc.Call[objc.Void](s_, objc.Sel("setBorderColor:"), value)
}

// The sampler comparison function used when performing a sample compare operation on a depth texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516001-comparefunction?language=objc
func (s_ SamplerDescriptor) CompareFunction() CompareFunction {
	rv := objc.Call[CompareFunction](s_, objc.Sel("compareFunction"))
	return rv
}

// The sampler comparison function used when performing a sample compare operation on a depth texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516001-comparefunction?language=objc
func (s_ SamplerDescriptor) SetCompareFunction(value CompareFunction) {
	objc.Call[objc.Void](s_, objc.Sel("setCompareFunction:"), value)
}

// A string that identifies the sampler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515771-label?language=objc
func (s_ SamplerDescriptor) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}

// A string that identifies the sampler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1515771-label?language=objc
func (s_ SamplerDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setLabel:"), value)
}

// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1615844-lodaverage?language=objc
func (s_ SamplerDescriptor) LodAverage() bool {
	rv := objc.Call[bool](s_, objc.Sel("lodAverage"))
	return rv
}

// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1615844-lodaverage?language=objc
func (s_ SamplerDescriptor) SetLodAverage(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setLodAverage:"), value)
}

// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516164-maxanisotropy?language=objc
func (s_ SamplerDescriptor) MaxAnisotropy() uint {
	rv := objc.Call[uint](s_, objc.Sel("maxAnisotropy"))
	return rv
}

// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/1516164-maxanisotropy?language=objc
func (s_ SamplerDescriptor) SetMaxAnisotropy(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxAnisotropy:"), value)
}

// A Boolean value that specifies whether the sampler can be encoded into an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/2915782-supportargumentbuffers?language=objc
func (s_ SamplerDescriptor) SupportArgumentBuffers() bool {
	rv := objc.Call[bool](s_, objc.Sel("supportArgumentBuffers"))
	return rv
}

// A Boolean value that specifies whether the sampler can be encoded into an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/2915782-supportargumentbuffers?language=objc
func (s_ SamplerDescriptor) SetSupportArgumentBuffers(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSupportArgumentBuffers:"), value)
}
