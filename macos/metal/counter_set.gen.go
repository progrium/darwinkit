// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A collection of individual counters a GPU device supports for a counter set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterset?language=objc
type PCounterSet interface {
	// optional
	Name() string
	HasName() bool

	// optional
	Counters() []PCounter
	HasCounters() bool
}

// A concrete type wrapper for the [PCounterSet] protocol.
type CounterSetWrapper struct {
	objc.Object
}

func (c_ CounterSetWrapper) HasName() bool {
	return c_.RespondsToSelector(objc.Sel("name"))
}

// The name of the GPU’s counter set instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterset/3081736-name?language=objc
func (c_ CounterSetWrapper) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

func (c_ CounterSetWrapper) HasCounters() bool {
	return c_.RespondsToSelector(objc.Sel("counters"))
}

// An array of the counter instances a GPU device supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterset/3081735-counters?language=objc
func (c_ CounterSetWrapper) Counters() []CounterWrapper {
	rv := objc.Call[[]CounterWrapper](c_, objc.Sel("counters"))
	return rv
}
