// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewTransitionLayout] class.
var CollectionViewTransitionLayoutClass = _CollectionViewTransitionLayoutClass{objc.GetClass("NSCollectionViewTransitionLayout")}

type _CollectionViewTransitionLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewTransitionLayout] class.
type ICollectionViewTransitionLayout interface {
	ICollectionViewLayout
	ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64
	UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey)
	CurrentLayout() CollectionViewLayout
	TransitionProgress() float64
	SetTransitionProgress(value float64)
	NextLayout() CollectionViewLayout
}

// An object that implements custom behaviors when changing from one layout to another in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout?language=objc
type CollectionViewTransitionLayout struct {
	CollectionViewLayout
}

func CollectionViewTransitionLayoutFrom(ptr unsafe.Pointer) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

func (c_ CollectionViewTransitionLayout) InitWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.Call[CollectionViewTransitionLayout](c_, objc.Sel("initWithCurrentLayout:nextLayout:"), objc.Ptr(currentLayout), objc.Ptr(newLayout))
	return rv
}

// Initializes and returns the transition layout object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524431-initwithcurrentlayout?language=objc
func CollectionViewTransitionLayout_InitWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.Alloc().InitWithCurrentLayoutNextLayout(currentLayout, newLayout)
}

func (cc _CollectionViewTransitionLayoutClass) Alloc() CollectionViewTransitionLayout {
	rv := objc.Call[CollectionViewTransitionLayout](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewTransitionLayout_Alloc() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.Alloc()
}

func (cc _CollectionViewTransitionLayoutClass) New() CollectionViewTransitionLayout {
	rv := objc.Call[CollectionViewTransitionLayout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewTransitionLayout() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.New()
}

func (c_ CollectionViewTransitionLayout) Init() CollectionViewTransitionLayout {
	rv := objc.Call[CollectionViewTransitionLayout](c_, objc.Sel("init"))
	return rv
}

// Returns the most recently set value for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524429-valueforanimatedkey?language=objc
func (c_ CollectionViewTransitionLayout) ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64 {
	rv := objc.Call[float64](c_, objc.Sel("valueForAnimatedKey:"), key)
	return rv
}

// Sets the value of a key whose value you use during the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524303-updatevalue?language=objc
func (c_ CollectionViewTransitionLayout) UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey) {
	objc.Call[objc.Void](c_, objc.Sel("updateValue:forAnimatedKey:"), value, key)
}

// The collection view’s current layout object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524425-currentlayout?language=objc
func (c_ CollectionViewTransitionLayout) CurrentLayout() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](c_, objc.Sel("currentLayout"))
	return rv
}

// The completion percentage of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524424-transitionprogress?language=objc
func (c_ CollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := objc.Call[float64](c_, objc.Sel("transitionProgress"))
	return rv
}

// The completion percentage of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524424-transitionprogress?language=objc
func (c_ CollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setTransitionProgress:"), value)
}

// The collection view’s new layout object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayout/1524426-nextlayout?language=objc
func (c_ CollectionViewTransitionLayout) NextLayout() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](c_, objc.Sel("nextLayout"))
	return rv
}
