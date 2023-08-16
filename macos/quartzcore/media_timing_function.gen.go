// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaTimingFunction] class.
var MediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}

type _MediaTimingFunctionClass struct {
	objc.Class
}

// An interface definition for the [MediaTimingFunction] class.
type IMediaTimingFunction interface {
	objc.IObject
	GetControlPointAtIndexValues(idx uint, ptr *float64)
}

// A function that defines the pacing of an animation as a timing curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunction?language=objc
type MediaTimingFunction struct {
	objc.Object
}

func MediaTimingFunctionFrom(ptr unsafe.Pointer) MediaTimingFunction {
	return MediaTimingFunction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MediaTimingFunctionClass) FunctionWithName(name MediaTimingFunctionName) MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](mc, objc.Sel("functionWithName:"), name)
	return rv
}

// Creates and returns a new instance of CAMediaTimingFunction configured with the predefined timing function specified by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunction/1521979-functionwithname?language=objc
func MediaTimingFunction_FunctionWithName(name MediaTimingFunctionName) MediaTimingFunction {
	return MediaTimingFunctionClass.FunctionWithName(name)
}

func (m_ MediaTimingFunction) InitWithControlPoints(c1x float64) MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](m_, objc.Sel("initWithControlPoints:"), c1x)
	return rv
}

// Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunction/1522235-initwithcontrolpoints?language=objc
func MediaTimingFunction_InitWithControlPoints(c1x float64) MediaTimingFunction {
	return MediaTimingFunctionClass.Alloc().InitWithControlPoints(c1x)
}

func (mc _MediaTimingFunctionClass) FunctionWithControlPoints(c1x float64) MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](mc, objc.Sel("functionWithControlPoints:"), c1x)
	return rv
}

// Creates and returns a new instance of CAMediaTimingFunction timing function modeled as a cubic Bézier curve using the specified control points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunction/1574338-functionwithcontrolpoints?language=objc
func MediaTimingFunction_FunctionWithControlPoints(c1x float64) MediaTimingFunction {
	return MediaTimingFunctionClass.FunctionWithControlPoints(c1x)
}

func (mc _MediaTimingFunctionClass) Alloc() MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](mc, objc.Sel("alloc"))
	return rv
}

func MediaTimingFunction_Alloc() MediaTimingFunction {
	return MediaTimingFunctionClass.Alloc()
}

func (mc _MediaTimingFunctionClass) New() MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaTimingFunction() MediaTimingFunction {
	return MediaTimingFunctionClass.New()
}

func (m_ MediaTimingFunction) Init() MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](m_, objc.Sel("init"))
	return rv
}

// Returns the control point for the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunction/1522057-getcontrolpointatindex?language=objc
func (m_ MediaTimingFunction) GetControlPointAtIndexValues(idx uint, ptr *float64) {
	objc.Call[objc.Void](m_, objc.Sel("getControlPointAtIndex:values:"), idx, ptr)
}
