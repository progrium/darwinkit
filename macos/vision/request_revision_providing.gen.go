// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for specifying the revision number of Vision algorithms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestrevisionproviding?language=objc
type PRequestRevisionProviding interface {
	// optional
	RequestRevision() uint
	HasRequestRevision() bool
}

// A concrete type wrapper for the [PRequestRevisionProviding] protocol.
type RequestRevisionProvidingWrapper struct {
	objc.Object
}

func (r_ RequestRevisionProvidingWrapper) HasRequestRevision() bool {
	return r_.RespondsToSelector(objc.Sel("requestRevision"))
}

// The revision of the VNRequest subclass used to generate the implementing object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestrevisionproviding/2967114-requestrevision?language=objc
func (r_ RequestRevisionProvidingWrapper) RequestRevision() uint {
	rv := objc.Call[uint](r_, objc.Sel("requestRevision"))
	return rv
}
