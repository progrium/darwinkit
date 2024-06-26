// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods and constants that a haptic feedback performer implements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackperformer?language=objc
type PHapticFeedbackPerformer interface {
	// optional
	PerformFeedbackPatternPerformanceTime(pattern HapticFeedbackPattern, performanceTime HapticFeedbackPerformanceTime)
	HasPerformFeedbackPatternPerformanceTime() bool
}

// ensure impl type implements protocol interface
var _ PHapticFeedbackPerformer = (*HapticFeedbackPerformerObject)(nil)

// A concrete type for the [PHapticFeedbackPerformer] protocol.
type HapticFeedbackPerformerObject struct {
	objc.Object
}

func (h_ HapticFeedbackPerformerObject) HasPerformFeedbackPatternPerformanceTime() bool {
	return h_.RespondsToSelector(objc.Sel("performFeedbackPattern:performanceTime:"))
}

// Initiates a specific pattern of haptic feedback to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackperformer/1441738-performfeedbackpattern?language=objc
func (h_ HapticFeedbackPerformerObject) PerformFeedbackPatternPerformanceTime(pattern HapticFeedbackPattern, performanceTime HapticFeedbackPerformanceTime) {
	objc.Call[objc.Void](h_, objc.Sel("performFeedbackPattern:performanceTime:"), pattern, performanceTime)
}
