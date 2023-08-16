// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// A general interface for objects that produce CIFilter instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilterconstructor?language=objc
type PFilterConstructor interface {
	// optional
	FilterWithName(name string) IFilter
	HasFilterWithName() bool
}

// A concrete type wrapper for the [PFilterConstructor] protocol.
type FilterConstructorWrapper struct {
	objc.Object
}

func (f_ FilterConstructorWrapper) HasFilterWithName() bool {
	return f_.RespondsToSelector(objc.Sel("filterWithName:"))
}

// Returns a filter object specified by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilterconstructor/1438018-filterwithname?language=objc
func (f_ FilterConstructorWrapper) FilterWithName(name string) Filter {
	rv := objc.Call[Filter](f_, objc.Sel("filterWithName:"), name)
	return rv
}
