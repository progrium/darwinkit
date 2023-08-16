// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface you implement that represents an abstract location inside your document’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlocation?language=objc
type PTextLocation interface {
	// optional
	Compare(location TextLocationWrapper) foundation.ComparisonResult
	HasCompare() bool
}

// A concrete type wrapper for the [PTextLocation] protocol.
type TextLocationWrapper struct {
	objc.Object
}

func (t_ TextLocationWrapper) HasCompare() bool {
	return t_.RespondsToSelector(objc.Sel("compare:"))
}

// Compares and returns the logical ordering to location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlocation/3801801-compare?language=objc
func (t_ TextLocationWrapper) Compare(location PTextLocation) foundation.ComparisonResult {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[foundation.ComparisonResult](t_, objc.Sel("compare:"), po0)
	return rv
}
