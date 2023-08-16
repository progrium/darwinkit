// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that allows objects to respond to actions triggered by a CALayer change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caaction?language=objc
type PAction interface {
	// optional
	RunActionForKeyObjectArguments(event string, anObject objc.Object, dict foundation.Dictionary)
	HasRunActionForKeyObjectArguments() bool
}

// A concrete type wrapper for the [PAction] protocol.
type ActionWrapper struct {
	objc.Object
}

func (a_ ActionWrapper) HasRunActionForKeyObjectArguments() bool {
	return a_.RespondsToSelector(objc.Sel("runActionForKey:object:arguments:"))
}

// Called to trigger the action specified by the identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caaction/1410806-runactionforkey?language=objc
func (a_ ActionWrapper) RunActionForKeyObjectArguments(event string, anObject objc.IObject, dict foundation.Dictionary) {
	objc.Call[objc.Void](a_, objc.Sel("runActionForKey:object:arguments:"), event, anObject, dict)
}
