// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram?language=objc
type PAreaHistogram interface {
	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

	// optional
	SetCount(value int)
	HasSetCount() bool

	// optional
	Count() int
	HasCount() bool
}

// A concrete type wrapper for the [PAreaHistogram] protocol.
type AreaHistogramWrapper struct {
	objc.Object
}

func (a_ AreaHistogramWrapper) HasSetScale() bool {
	return a_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547093-scale?language=objc
func (a_ AreaHistogramWrapper) SetScale(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setScale:"), value)
}

func (a_ AreaHistogramWrapper) HasScale() bool {
	return a_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547093-scale?language=objc
func (a_ AreaHistogramWrapper) Scale() float64 {
	rv := objc.Call[float64](a_, objc.Sel("scale"))
	return rv
}

func (a_ AreaHistogramWrapper) HasSetCount() bool {
	return a_.RespondsToSelector(objc.Sel("setCount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547092-count?language=objc
func (a_ AreaHistogramWrapper) SetCount(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setCount:"), value)
}

func (a_ AreaHistogramWrapper) HasCount() bool {
	return a_.RespondsToSelector(objc.Sel("count"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareahistogram/3547092-count?language=objc
func (a_ AreaHistogramWrapper) Count() int {
	rv := objc.Call[int](a_, objc.Sel("count"))
	return rv
}
