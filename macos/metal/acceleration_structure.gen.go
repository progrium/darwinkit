// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A collection of model data for GPU-accelerated intersection of rays with the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructure?language=objc
type PAccelerationStructure interface {
	// optional
	Size() uint
	HasSize() bool
}

// A concrete type wrapper for the [PAccelerationStructure] protocol.
type AccelerationStructureWrapper struct {
	objc.Object
}

func (a_ AccelerationStructureWrapper) HasSize() bool {
	return a_.RespondsToSelector(objc.Sel("size"))
}

// The size of the acceleration structure’s memory allocation, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructure/3553858-size?language=objc
func (a_ AccelerationStructureWrapper) Size() uint {
	rv := objc.Call[uint](a_, objc.Sel("size"))
	return rv
}
