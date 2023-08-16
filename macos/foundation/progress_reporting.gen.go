// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// An interface for objects that report progress using a single progress instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogressreporting?language=objc
type PProgressReporting interface {
	// optional
	Progress() IProgress
	HasProgress() bool
}

// A concrete type wrapper for the [PProgressReporting] protocol.
type ProgressReportingWrapper struct {
	objc.Object
}

func (p_ ProgressReportingWrapper) HasProgress() bool {
	return p_.RespondsToSelector(objc.Sel("progress"))
}

// The progress object returned by the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogressreporting/1412781-progress?language=objc
func (p_ ProgressReportingWrapper) Progress() Progress {
	rv := objc.Call[Progress](p_, objc.Sel("progress"))
	return rv
}
