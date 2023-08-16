// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An individual counter a GPU device lists within one of its counter sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounter?language=objc
type PCounter interface {
	// optional
	Name() string
	HasName() bool
}

// A concrete type wrapper for the [PCounter] protocol.
type CounterWrapper struct {
	objc.Object
}

func (c_ CounterWrapper) HasName() bool {
	return c_.RespondsToSelector(objc.Sel("name"))
}

// The name of a GPU’s counter instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounter/3081701-name?language=objc
func (c_ CounterWrapper) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}
