// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for providing progress information on long-running tasks in Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestprogressproviding?language=objc
type PRequestProgressProviding interface {
	// optional
	SetProgressHandler(value RequestProgressHandler)
	HasSetProgressHandler() bool

	// optional
	ProgressHandler() RequestProgressHandler
	HasProgressHandler() bool

	// optional
	Indeterminate() bool
	HasIndeterminate() bool
}

// A concrete type wrapper for the [PRequestProgressProviding] protocol.
type RequestProgressProvidingWrapper struct {
	objc.Object
}

func (r_ RequestProgressProvidingWrapper) HasSetProgressHandler() bool {
	return r_.RespondsToSelector(objc.Sel("setProgressHandler:"))
}

// A block of code executed periodically during a Vision request to report progress on long-running tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestprogressproviding/3152652-progresshandler?language=objc
func (r_ RequestProgressProvidingWrapper) SetProgressHandler(value RequestProgressHandler) {
	objc.Call[objc.Void](r_, objc.Sel("setProgressHandler:"), value)
}

func (r_ RequestProgressProvidingWrapper) HasProgressHandler() bool {
	return r_.RespondsToSelector(objc.Sel("progressHandler"))
}

// A block of code executed periodically during a Vision request to report progress on long-running tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestprogressproviding/3152652-progresshandler?language=objc
func (r_ RequestProgressProvidingWrapper) ProgressHandler() RequestProgressHandler {
	rv := objc.Call[RequestProgressHandler](r_, objc.Sel("progressHandler"))
	return rv
}

func (r_ RequestProgressProvidingWrapper) HasIndeterminate() bool {
	return r_.RespondsToSelector(objc.Sel("indeterminate"))
}

// A Boolean set to true when a request can't determine its progress in fractions completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestprogressproviding/3152651-indeterminate?language=objc
func (r_ RequestProgressProvidingWrapper) Indeterminate() bool {
	rv := objc.Call[bool](r_, objc.Sel("indeterminate"))
	return rv
}
