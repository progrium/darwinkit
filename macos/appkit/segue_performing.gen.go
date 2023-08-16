// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that support the mediation of a custom segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming?language=objc
type PSeguePerforming interface {
	// optional
	PerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object)
	HasPerformSegueWithIdentifierSender() bool

	// optional
	ShouldPerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.Object) bool
	HasShouldPerformSegueWithIdentifierSender() bool

	// optional
	PrepareForSegueSender(segue StoryboardSegue, sender objc.Object)
	HasPrepareForSegueSender() bool
}

// A concrete type wrapper for the [PSeguePerforming] protocol.
type SeguePerformingWrapper struct {
	objc.Object
}

func (s_ SeguePerformingWrapper) HasPerformSegueWithIdentifierSender() bool {
	return s_.RespondsToSelector(objc.Sel("performSegueWithIdentifier:sender:"))
}

// Performs the specified segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409583-performseguewithidentifier?language=objc
func (s_ SeguePerformingWrapper) PerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("performSegueWithIdentifier:sender:"), identifier, sender)
}

func (s_ SeguePerformingWrapper) HasShouldPerformSegueWithIdentifierSender() bool {
	return s_.RespondsToSelector(objc.Sel("shouldPerformSegueWithIdentifier:sender:"))
}

// Called immediately prior to the performance of a storyboard segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409574-shouldperformseguewithidentifier?language=objc
func (s_ SeguePerformingWrapper) ShouldPerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objc.IObject) bool {
	rv := objc.Call[bool](s_, objc.Sel("shouldPerformSegueWithIdentifier:sender:"), identifier, sender)
	return rv
}

func (s_ SeguePerformingWrapper) HasPrepareForSegueSender() bool {
	return s_.RespondsToSelector(objc.Sel("prepareForSegue:sender:"))
}

// Called when a segue is about to be performed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegueperforming/1409580-prepareforsegue?language=objc
func (s_ SeguePerformingWrapper) PrepareForSegueSender(segue IStoryboardSegue, sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("prepareForSegue:sender:"), objc.Ptr(segue), sender)
}
