// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HapticFeedbackManager] class.
var HapticFeedbackManagerClass = _HapticFeedbackManagerClass{objc.GetClass("NSHapticFeedbackManager")}

type _HapticFeedbackManagerClass struct {
	objc.Class
}

// An interface definition for the [HapticFeedbackManager] class.
type IHapticFeedbackManager interface {
	objc.IObject
}

// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackmanager?language=objc
type HapticFeedbackManager struct {
	objc.Object
}

func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HapticFeedbackManagerClass) Alloc() HapticFeedbackManager {
	rv := objc.Call[HapticFeedbackManager](hc, objc.Sel("alloc"))
	return rv
}

func HapticFeedbackManager_Alloc() HapticFeedbackManager {
	return HapticFeedbackManagerClass.Alloc()
}

func (hc _HapticFeedbackManagerClass) New() HapticFeedbackManager {
	rv := objc.Call[HapticFeedbackManager](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHapticFeedbackManager() HapticFeedbackManager {
	return HapticFeedbackManagerClass.New()
}

func (h_ HapticFeedbackManager) Init() HapticFeedbackManager {
	rv := objc.Call[HapticFeedbackManager](h_, objc.Sel("init"))
	return rv
}

// Requests a haptic feedback performer object that is based on the current input device, accessibility settings, and user preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackmanager/1441752-defaultperformer?language=objc
func (hc _HapticFeedbackManagerClass) DefaultPerformer() HapticFeedbackPerformerWrapper {
	rv := objc.Call[HapticFeedbackPerformerWrapper](hc, objc.Sel("defaultPerformer"))
	return rv
}

// Requests a haptic feedback performer object that is based on the current input device, accessibility settings, and user preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackmanager/1441752-defaultperformer?language=objc
func HapticFeedbackManager_DefaultPerformer() HapticFeedbackPerformerWrapper {
	return HapticFeedbackManagerClass.DefaultPerformer()
}
