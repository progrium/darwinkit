// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StoryboardSegue] class.
var StoryboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}

type _StoryboardSegueClass struct {
	objc.Class
}

// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	objc.IObject
	Perform()
	SourceController() objc.Object
	DestinationController() objc.Object
	Identifier() StoryboardSegueIdentifier
}

// A transition or containment relationship between two scenes in a storyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue?language=objc
type StoryboardSegue struct {
	objc.Object
}

func StoryboardSegueFrom(ptr unsafe.Pointer) StoryboardSegue {
	return StoryboardSegue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ StoryboardSegue) InitWithIdentifierSourceDestination(identifier StoryboardSegueIdentifier, sourceController objc.IObject, destinationController objc.IObject) StoryboardSegue {
	rv := objc.Call[StoryboardSegue](s_, objc.Sel("initWithIdentifier:source:destination:"), identifier, sourceController, destinationController)
	return rv
}

// The designated initializer for a storyboard segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409572-initwithidentifier?language=objc
func StoryboardSegue_InitWithIdentifierSourceDestination(identifier StoryboardSegueIdentifier, sourceController objc.IObject, destinationController objc.IObject) StoryboardSegue {
	return StoryboardSegueClass.Alloc().InitWithIdentifierSourceDestination(identifier, sourceController, destinationController)
}

func (sc _StoryboardSegueClass) SegueWithIdentifierSourceDestinationPerformHandler(identifier StoryboardSegueIdentifier, sourceController objc.IObject, destinationController objc.IObject, performHandler func()) StoryboardSegue {
	rv := objc.Call[StoryboardSegue](sc, objc.Sel("segueWithIdentifier:source:destination:performHandler:"), identifier, sourceController, destinationController, performHandler)
	return rv
}

// Creates a storyboard segue and a block used when the segue is performed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409576-seguewithidentifier?language=objc
func StoryboardSegue_SegueWithIdentifierSourceDestinationPerformHandler(identifier StoryboardSegueIdentifier, sourceController objc.IObject, destinationController objc.IObject, performHandler func()) StoryboardSegue {
	return StoryboardSegueClass.SegueWithIdentifierSourceDestinationPerformHandler(identifier, sourceController, destinationController, performHandler)
}

func (sc _StoryboardSegueClass) Alloc() StoryboardSegue {
	rv := objc.Call[StoryboardSegue](sc, objc.Sel("alloc"))
	return rv
}

func StoryboardSegue_Alloc() StoryboardSegue {
	return StoryboardSegueClass.Alloc()
}

func (sc _StoryboardSegueClass) New() StoryboardSegue {
	rv := objc.Call[StoryboardSegue](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStoryboardSegue() StoryboardSegue {
	return StoryboardSegueClass.New()
}

func (s_ StoryboardSegue) Init() StoryboardSegue {
	rv := objc.Call[StoryboardSegue](s_, objc.Sel("init"))
	return rv
}

// Performs a visual transition from one controller to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409587-perform?language=objc
func (s_ StoryboardSegue) Perform() {
	objc.Call[objc.Void](s_, objc.Sel("perform"))
}

// The starting/containing view controller or window controller for the storyboard segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409582-sourcecontroller?language=objc
func (s_ StoryboardSegue) SourceController() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("sourceController"))
	return rv
}

// The ending/contained view controller or window controller for the storyboard segue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409586-destinationcontroller?language=objc
func (s_ StoryboardSegue) DestinationController() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("destinationController"))
	return rv
}

// An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegue/1409578-identifier?language=objc
func (s_ StoryboardSegue) Identifier() StoryboardSegueIdentifier {
	rv := objc.Call[StoryboardSegueIdentifier](s_, objc.Sel("identifier"))
	return rv
}
