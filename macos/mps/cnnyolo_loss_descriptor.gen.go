// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNYOLOLossDescriptor] class.
var CNNYOLOLossDescriptorClass = _CNNYOLOLossDescriptorClass{objc.GetClass("MPSCNNYOLOLossDescriptor")}

type _CNNYOLOLossDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNYOLOLossDescriptor] class.
type ICNNYOLOLossDescriptor interface {
	objc.IObject
	ScaleWH() float64
	SetScaleWH(value float64)
	MinIOUForObjectPresence() float64
	SetMinIOUForObjectPresence(value float64)
	ScaleXY() float64
	SetScaleXY(value float64)
	ScaleObject() float64
	SetScaleObject(value float64)
	MaxIOUForObjectAbsence() float64
	SetMaxIOUForObjectAbsence(value float64)
	AnchorBoxes() []byte
	SetAnchorBoxes(value []byte)
	Rescore() bool
	SetRescore(value bool)
	ReductionType() CNNReductionType
	SetReductionType(value CNNReductionType)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	ConfidenceLossDescriptor() CNNLossDescriptor
	SetConfidenceLossDescriptor(value ICNNLossDescriptor)
	WHLossDescriptor() CNNLossDescriptor
	SetWHLossDescriptor(value ICNNLossDescriptor)
	ScaleClass() float64
	SetScaleClass(value float64)
	XYLossDescriptor() CNNLossDescriptor
	SetXYLossDescriptor(value ICNNLossDescriptor)
	ScaleNoObject() float64
	SetScaleNoObject(value float64)
	ClassesLossDescriptor() CNNLossDescriptor
	SetClassesLossDescriptor(value ICNNLossDescriptor)
	NumberOfAnchorBoxes() uint
	SetNumberOfAnchorBoxes(value uint)
}

// An object that specifies properties used by a YOLO loss kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor?language=objc
type CNNYOLOLossDescriptor struct {
	objc.Object
}

func CNNYOLOLossDescriptorFrom(ptr unsafe.Pointer) CNNYOLOLossDescriptor {
	return CNNYOLOLossDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CNNYOLOLossDescriptorClass) Alloc() CNNYOLOLossDescriptor {
	rv := objc.Call[CNNYOLOLossDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CNNYOLOLossDescriptor_Alloc() CNNYOLOLossDescriptor {
	return CNNYOLOLossDescriptorClass.Alloc()
}

func (cc _CNNYOLOLossDescriptorClass) New() CNNYOLOLossDescriptor {
	rv := objc.Call[CNNYOLOLossDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNYOLOLossDescriptor() CNNYOLOLossDescriptor {
	return CNNYOLOLossDescriptorClass.New()
}

func (c_ CNNYOLOLossDescriptor) Init() CNNYOLOLossDescriptor {
	rv := objc.Call[CNNYOLOLossDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976500-cnnlossdescriptorwithxylosstype?language=objc
func (cc _CNNYOLOLossDescriptorClass) CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes(XYLossType CNNLossType, WHLossType CNNLossType, confidenceLossType CNNLossType, classesLossType CNNLossType, reductionType CNNReductionType, anchorBoxes []byte, numberOfAnchorBoxes uint) CNNYOLOLossDescriptor {
	rv := objc.Call[CNNYOLOLossDescriptor](cc, objc.Sel("cnnLossDescriptorWithXYLossType:WHLossType:confidenceLossType:classesLossType:reductionType:anchorBoxes:numberOfAnchorBoxes:"), XYLossType, WHLossType, confidenceLossType, classesLossType, reductionType, anchorBoxes, numberOfAnchorBoxes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976500-cnnlossdescriptorwithxylosstype?language=objc
func CNNYOLOLossDescriptor_CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes(XYLossType CNNLossType, WHLossType CNNLossType, confidenceLossType CNNLossType, classesLossType CNNLossType, reductionType CNNReductionType, anchorBoxes []byte, numberOfAnchorBoxes uint) CNNYOLOLossDescriptor {
	return CNNYOLOLossDescriptorClass.CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes(XYLossType, WHLossType, confidenceLossType, classesLossType, reductionType, anchorBoxes, numberOfAnchorBoxes)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976510-scalewh?language=objc
func (c_ CNNYOLOLossDescriptor) ScaleWH() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleWH"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976510-scalewh?language=objc
func (c_ CNNYOLOLossDescriptor) SetScaleWH(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleWH:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976503-miniouforobjectpresence?language=objc
func (c_ CNNYOLOLossDescriptor) MinIOUForObjectPresence() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minIOUForObjectPresence"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976503-miniouforobjectpresence?language=objc
func (c_ CNNYOLOLossDescriptor) SetMinIOUForObjectPresence(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinIOUForObjectPresence:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976511-scalexy?language=objc
func (c_ CNNYOLOLossDescriptor) ScaleXY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleXY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976511-scalexy?language=objc
func (c_ CNNYOLOLossDescriptor) SetScaleXY(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleXY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976509-scaleobject?language=objc
func (c_ CNNYOLOLossDescriptor) ScaleObject() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleObject"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976509-scaleobject?language=objc
func (c_ CNNYOLOLossDescriptor) SetScaleObject(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleObject:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976502-maxiouforobjectabsence?language=objc
func (c_ CNNYOLOLossDescriptor) MaxIOUForObjectAbsence() float64 {
	rv := objc.Call[float64](c_, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976502-maxiouforobjectabsence?language=objc
func (c_ CNNYOLOLossDescriptor) SetMaxIOUForObjectAbsence(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxIOUForObjectAbsence:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976498-anchorboxes?language=objc
func (c_ CNNYOLOLossDescriptor) AnchorBoxes() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("anchorBoxes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976498-anchorboxes?language=objc
func (c_ CNNYOLOLossDescriptor) SetAnchorBoxes(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setAnchorBoxes:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976506-rescore?language=objc
func (c_ CNNYOLOLossDescriptor) Rescore() bool {
	rv := objc.Call[bool](c_, objc.Sel("rescore"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976506-rescore?language=objc
func (c_ CNNYOLOLossDescriptor) SetRescore(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setRescore:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976505-reductiontype?language=objc
func (c_ CNNYOLOLossDescriptor) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](c_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976505-reductiontype?language=objc
func (c_ CNNYOLOLossDescriptor) SetReductionType(value CNNReductionType) {
	objc.Call[objc.Void](c_, objc.Sel("setReductionType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/3547984-reduceacrossbatch?language=objc
func (c_ CNNYOLOLossDescriptor) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/3547984-reduceacrossbatch?language=objc
func (c_ CNNYOLOLossDescriptor) SetReduceAcrossBatch(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setReduceAcrossBatch:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976501-confidencelossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) ConfidenceLossDescriptor() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](c_, objc.Sel("confidenceLossDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976501-confidencelossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) SetConfidenceLossDescriptor(value ICNNLossDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setConfidenceLossDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976496-whlossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) WHLossDescriptor() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](c_, objc.Sel("WHLossDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976496-whlossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) SetWHLossDescriptor(value ICNNLossDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setWHLossDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976507-scaleclass?language=objc
func (c_ CNNYOLOLossDescriptor) ScaleClass() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleClass"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976507-scaleclass?language=objc
func (c_ CNNYOLOLossDescriptor) SetScaleClass(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleClass:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976497-xylossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) XYLossDescriptor() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](c_, objc.Sel("XYLossDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976497-xylossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) SetXYLossDescriptor(value ICNNLossDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setXYLossDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976508-scalenoobject?language=objc
func (c_ CNNYOLOLossDescriptor) ScaleNoObject() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleNoObject"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976508-scalenoobject?language=objc
func (c_ CNNYOLOLossDescriptor) SetScaleNoObject(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleNoObject:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976499-classeslossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) ClassesLossDescriptor() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](c_, objc.Sel("classesLossDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976499-classeslossdescriptor?language=objc
func (c_ CNNYOLOLossDescriptor) SetClassesLossDescriptor(value ICNNLossDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setClassesLossDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976504-numberofanchorboxes?language=objc
func (c_ CNNYOLOLossDescriptor) NumberOfAnchorBoxes() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfAnchorBoxes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976504-numberofanchorboxes?language=objc
func (c_ CNNYOLOLossDescriptor) SetNumberOfAnchorBoxes(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfAnchorBoxes:"), value)
}
