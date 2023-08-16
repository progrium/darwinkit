// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Storyboard] class.
var StoryboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}

type _StoryboardClass struct {
	objc.Class
}

// An interface definition for the [Storyboard] class.
type IStoryboard interface {
	objc.IObject
	InstantiateControllerWithIdentifierCreator(identifier StoryboardSceneIdentifier, block StoryboardControllerCreator) objc.Object
	InstantiateInitialController() objc.Object
	InstantiateInitialControllerWithCreator(block StoryboardControllerCreator) objc.Object
}

// An encapsulation of the design-time view controller and window controller graph represented in an Interface Builder storyboard resource file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard?language=objc
type Storyboard struct {
	objc.Object
}

func StoryboardFrom(ptr unsafe.Pointer) Storyboard {
	return Storyboard{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StoryboardClass) StoryboardWithNameBundle(name StoryboardName, storyboardBundleOrNil foundation.IBundle) Storyboard {
	rv := objc.Call[Storyboard](sc, objc.Sel("storyboardWithName:bundle:"), name, objc.Ptr(storyboardBundleOrNil))
	return rv
}

// Creates a storyboard based on the named storyboard file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/1426547-storyboardwithname?language=objc
func Storyboard_StoryboardWithNameBundle(name StoryboardName, storyboardBundleOrNil foundation.IBundle) Storyboard {
	return StoryboardClass.StoryboardWithNameBundle(name, storyboardBundleOrNil)
}

func (sc _StoryboardClass) Alloc() Storyboard {
	rv := objc.Call[Storyboard](sc, objc.Sel("alloc"))
	return rv
}

func Storyboard_Alloc() Storyboard {
	return StoryboardClass.Alloc()
}

func (sc _StoryboardClass) New() Storyboard {
	rv := objc.Call[Storyboard](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStoryboard() Storyboard {
	return StoryboardClass.New()
}

func (s_ Storyboard) Init() Storyboard {
	rv := objc.Call[Storyboard](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/3174922-instantiatecontrollerwithidentif?language=objc
func (s_ Storyboard) InstantiateControllerWithIdentifierCreator(identifier StoryboardSceneIdentifier, block StoryboardControllerCreator) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("instantiateControllerWithIdentifier:creator:"), identifier, block)
	return rv
}

// Creates the initial view controller or window controller from a storyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/1426545-instantiateinitialcontroller?language=objc
func (s_ Storyboard) InstantiateInitialController() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("instantiateInitialController"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/3174923-instantiateinitialcontrollerwith?language=objc
func (s_ Storyboard) InstantiateInitialControllerWithCreator(block StoryboardControllerCreator) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("instantiateInitialControllerWithCreator:"), block)
	return rv
}

// The app's main storyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/2881236-mainstoryboard?language=objc
func (sc _StoryboardClass) MainStoryboard() Storyboard {
	rv := objc.Call[Storyboard](sc, objc.Sel("mainStoryboard"))
	return rv
}

// The app's main storyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboard/2881236-mainstoryboard?language=objc
func Storyboard_MainStoryboard() Storyboard {
	return StoryboardClass.MainStoryboard()
}
