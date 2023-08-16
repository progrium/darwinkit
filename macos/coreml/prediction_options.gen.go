// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PredictionOptions] class.
var PredictionOptionsClass = _PredictionOptionsClass{objc.GetClass("MLPredictionOptions")}

type _PredictionOptionsClass struct {
	objc.Class
}

// An interface definition for the [PredictionOptions] class.
type IPredictionOptions interface {
	objc.IObject
	OutputBackings() map[string]objc.Object
	SetOutputBackings(value map[string]objc.IObject)
}

// The options available when making a prediction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlpredictionoptions?language=objc
type PredictionOptions struct {
	objc.Object
}

func PredictionOptionsFrom(ptr unsafe.Pointer) PredictionOptions {
	return PredictionOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PredictionOptionsClass) Alloc() PredictionOptions {
	rv := objc.Call[PredictionOptions](pc, objc.Sel("alloc"))
	return rv
}

func PredictionOptions_Alloc() PredictionOptions {
	return PredictionOptionsClass.Alloc()
}

func (pc _PredictionOptionsClass) New() PredictionOptions {
	rv := objc.Call[PredictionOptions](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPredictionOptions() PredictionOptions {
	return PredictionOptionsClass.New()
}

func (p_ PredictionOptions) Init() PredictionOptions {
	rv := objc.Call[PredictionOptions](p_, objc.Sel("init"))
	return rv
}

// A dictionary of feature names and client-allocated buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlpredictionoptions/3882837-outputbackings?language=objc
func (p_ PredictionOptions) OutputBackings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("outputBackings"))
	return rv
}

// A dictionary of feature names and client-allocated buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlpredictionoptions/3882837-outputbackings?language=objc
func (p_ PredictionOptions) SetOutputBackings(value map[string]objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setOutputBackings:"), value)
}
