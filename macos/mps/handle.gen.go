// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

// The protocol that provides resource identification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpshandle?language=objc
type PHandle interface {
	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PHandle] protocol.
type HandleWrapper struct {
	objc.Object
}

func (h_ HandleWrapper) HasLabel() bool {
	return h_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpshandle/2866414-label?language=objc
func (h_ HandleWrapper) Label() string {
	rv := objc.Call[string](h_, objc.Sel("label"))
	return rv
}
