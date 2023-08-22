// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNYOLOLoss] class.
var CNNYOLOLossClass = _CNNYOLOLossClass{objc.GetClass("MPSCNNYOLOLoss")}

type _CNNYOLOLossClass struct {
	objc.Class
}

// An interface definition for the [CNNYOLOLoss] class.
type ICNNYOLOLoss interface {
	ICNNKernel
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesLabels(commandBufferObject objc.IObject, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array
	EncodeToCommandBufferSourceImageLabels(commandBuffer metal.PCommandBuffer, sourceImage IImage, labels ICNNLossLabels) Image
	EncodeToCommandBufferObjectSourceImageLabels(commandBufferObject objc.IObject, sourceImage IImage, labels ICNNLossLabels) Image
	LossClasses() CNNLoss
	ScaleWH() float64
	MinIOUForObjectPresence() float64
	ScaleXY() float64
	ScaleObject() float64
	MaxIOUForObjectAbsence() float64
	AnchorBoxes() []byte
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	ScaleClass() float64
	LossConfidence() CNNLoss
	ScaleNoObject() float64
	LossXY() CNNLoss
	LossWH() CNNLoss
	NumberOfAnchorBoxes() uint
}

// A kernel that computes the YOLO loss and loss gradient between specified predictions and labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss?language=objc
type CNNYOLOLoss struct {
	CNNKernel
}

func CNNYOLOLossFrom(ptr unsafe.Pointer) CNNYOLOLoss {
	return CNNYOLOLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNYOLOLoss) InitWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNYOLOLossDescriptor) CNNYOLOLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNYOLOLoss](c_, objc.Sel("initWithDevice:lossDescriptor:"), po0, objc.Ptr(lossDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976481-initwithdevice?language=objc
func NewCNNYOLOLossWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNYOLOLossDescriptor) CNNYOLOLoss {
	instance := CNNYOLOLossClass.Alloc().InitWithDeviceLossDescriptor(device, lossDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNYOLOLossClass) Alloc() CNNYOLOLoss {
	rv := objc.Call[CNNYOLOLoss](cc, objc.Sel("alloc"))
	return rv
}

func CNNYOLOLoss_Alloc() CNNYOLOLoss {
	return CNNYOLOLossClass.Alloc()
}

func (cc _CNNYOLOLossClass) New() CNNYOLOLoss {
	rv := objc.Call[CNNYOLOLoss](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNYOLOLoss() CNNYOLOLoss {
	return CNNYOLOLossClass.New()
}

func (c_ CNNYOLOLoss) Init() CNNYOLOLoss {
	rv := objc.Call[CNNYOLOLoss](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNYOLOLoss) InitWithDevice(device metal.PDevice) CNNYOLOLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNYOLOLoss](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNYOLOLossWithDevice(device metal.PDevice) CNNYOLOLoss {
	instance := CNNYOLOLossClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNYOLOLoss) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNYOLOLoss {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNYOLOLoss](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNYOLOLoss_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNYOLOLoss {
	instance := CNNYOLOLossClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encodebatchtocommandbuffer?language=objc
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), po0, sourceImage, labels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encodebatchtocommandbuffer?language=objc
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferObjectSourceImagesLabels(commandBufferObject objc.IObject, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), objc.Ptr(commandBufferObject), sourceImage, labels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976478-encodetocommandbuffer?language=objc
func (c_ CNNYOLOLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer metal.PCommandBuffer, sourceImage IImage, labels ICNNLossLabels) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), po0, objc.Ptr(sourceImage), objc.Ptr(labels))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976478-encodetocommandbuffer?language=objc
func (c_ CNNYOLOLoss) EncodeToCommandBufferObjectSourceImageLabels(commandBufferObject objc.IObject, sourceImage IImage, labels ICNNLossLabels) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(labels))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976482-lossclasses?language=objc
func (c_ CNNYOLOLoss) LossClasses() CNNLoss {
	rv := objc.Call[CNNLoss](c_, objc.Sel("lossClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976493-scalewh?language=objc
func (c_ CNNYOLOLoss) ScaleWH() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleWH"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976487-miniouforobjectpresence?language=objc
func (c_ CNNYOLOLoss) MinIOUForObjectPresence() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minIOUForObjectPresence"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976494-scalexy?language=objc
func (c_ CNNYOLOLoss) ScaleXY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleXY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976492-scaleobject?language=objc
func (c_ CNNYOLOLoss) ScaleObject() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleObject"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976486-maxiouforobjectabsence?language=objc
func (c_ CNNYOLOLoss) MaxIOUForObjectAbsence() float64 {
	rv := objc.Call[float64](c_, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976475-anchorboxes?language=objc
func (c_ CNNYOLOLoss) AnchorBoxes() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("anchorBoxes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976489-reductiontype?language=objc
func (c_ CNNYOLOLoss) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](c_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/3547983-reduceacrossbatch?language=objc
func (c_ CNNYOLOLoss) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976490-scaleclass?language=objc
func (c_ CNNYOLOLoss) ScaleClass() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleClass"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976483-lossconfidence?language=objc
func (c_ CNNYOLOLoss) LossConfidence() CNNLoss {
	rv := objc.Call[CNNLoss](c_, objc.Sel("lossConfidence"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976491-scalenoobject?language=objc
func (c_ CNNYOLOLoss) ScaleNoObject() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleNoObject"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976485-lossxy?language=objc
func (c_ CNNYOLOLoss) LossXY() CNNLoss {
	rv := objc.Call[CNNLoss](c_, objc.Sel("lossXY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976484-losswh?language=objc
func (c_ CNNYOLOLoss) LossWH() CNNLoss {
	rv := objc.Call[CNNLoss](c_, objc.Sel("lossWH"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976488-numberofanchorboxes?language=objc
func (c_ CNNYOLOLoss) NumberOfAnchorBoxes() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfAnchorBoxes"))
	return rv
}
