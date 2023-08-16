// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that the destination object (or recipient) of a dragged object can implement to support spring-loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination?language=objc
type PSpringLoadingDestination interface {
	// optional
	SpringLoadingUpdated(draggingInfo DraggingInfoWrapper) SpringLoadingOptions
	HasSpringLoadingUpdated() bool

	// optional
	SpringLoadingEntered(draggingInfo DraggingInfoWrapper) SpringLoadingOptions
	HasSpringLoadingEntered() bool

	// optional
	DraggingEnded(draggingInfo DraggingInfoWrapper)
	HasDraggingEnded() bool

	// optional
	SpringLoadingHighlightChanged(draggingInfo DraggingInfoWrapper)
	HasSpringLoadingHighlightChanged() bool

	// optional
	SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo DraggingInfoWrapper)
	HasSpringLoadingActivatedDraggingInfo() bool

	// optional
	SpringLoadingExited(draggingInfo DraggingInfoWrapper)
	HasSpringLoadingExited() bool
}

// A concrete type wrapper for the [PSpringLoadingDestination] protocol.
type SpringLoadingDestinationWrapper struct {
	objc.Object
}

func (s_ SpringLoadingDestinationWrapper) HasSpringLoadingUpdated() bool {
	return s_.RespondsToSelector(objc.Sel("springLoadingUpdated:"))
}

// Returns whether to enable or disable spring-loading as a drag moves within the bounds of the spring-loading destination or draggingInfo changes during the drag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1415976-springloadingupdated?language=objc
func (s_ SpringLoadingDestinationWrapper) SpringLoadingUpdated(draggingInfo PDraggingInfo) SpringLoadingOptions {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.Call[SpringLoadingOptions](s_, objc.Sel("springLoadingUpdated:"), po0)
	return rv
}

func (s_ SpringLoadingDestinationWrapper) HasSpringLoadingEntered() bool {
	return s_.RespondsToSelector(objc.Sel("springLoadingEntered:"))
}

// Returns whether to enable or disable spring-loading when a drag enters the bounds of the spring-loading destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1415996-springloadingentered?language=objc
func (s_ SpringLoadingDestinationWrapper) SpringLoadingEntered(draggingInfo PDraggingInfo) SpringLoadingOptions {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.Call[SpringLoadingOptions](s_, objc.Sel("springLoadingEntered:"), po0)
	return rv
}

func (s_ SpringLoadingDestinationWrapper) HasDraggingEnded() bool {
	return s_.RespondsToSelector(objc.Sel("draggingEnded:"))
}

// Responds to the end of a drag operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1416047-draggingended?language=objc
func (s_ SpringLoadingDestinationWrapper) DraggingEnded(draggingInfo PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.Call[objc.Void](s_, objc.Sel("draggingEnded:"), po0)
}

func (s_ SpringLoadingDestinationWrapper) HasSpringLoadingHighlightChanged() bool {
	return s_.RespondsToSelector(objc.Sel("springLoadingHighlightChanged:"))
}

// Updates the destination’s user interface to display a new highlighting style during a spring-loading operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1416062-springloadinghighlightchanged?language=objc
func (s_ SpringLoadingDestinationWrapper) SpringLoadingHighlightChanged(draggingInfo PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.Call[objc.Void](s_, objc.Sel("springLoadingHighlightChanged:"), po0)
}

func (s_ SpringLoadingDestinationWrapper) HasSpringLoadingActivatedDraggingInfo() bool {
	return s_.RespondsToSelector(objc.Sel("springLoadingActivated:draggingInfo:"))
}

// Responds to the activation or deactivation of spring-loading on a destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1416106-springloadingactivated?language=objc
func (s_ SpringLoadingDestinationWrapper) SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo PDraggingInfo) {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.Call[objc.Void](s_, objc.Sel("springLoadingActivated:draggingInfo:"), activated, po1)
}

func (s_ SpringLoadingDestinationWrapper) HasSpringLoadingExited() bool {
	return s_.RespondsToSelector(objc.Sel("springLoadingExited:"))
}

// Responds when a drag exits the bounds of the spring-loading destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingdestination/1415972-springloadingexited?language=objc
func (s_ SpringLoadingDestinationWrapper) SpringLoadingExited(draggingInfo PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.Call[objc.Void](s_, objc.Sel("springLoadingExited:"), po0)
}
