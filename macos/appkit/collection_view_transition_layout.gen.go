// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionViewTransitionLayoutClass = _CollectionViewTransitionLayoutClass{objc.GetClass("NSCollectionViewTransitionLayout")}

type _CollectionViewTransitionLayoutClass struct {
	objc.Class
}

type ICollectionViewTransitionLayout interface {
	ICollectionViewLayout
	UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey)
	ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64
	TransitionProgress() float64
	SetTransitionProgress(value float64)
	CurrentLayout() CollectionViewLayout
	NextLayout() CollectionViewLayout
}

type CollectionViewTransitionLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewTransitionLayout(ptr unsafe.Pointer) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (c_ CollectionViewTransitionLayout) InitWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, objc.GetSelector("initWithCurrentLayout:nextLayout:"), objc.ExtractPtr(currentLayout), objc.ExtractPtr(newLayout))
	return rv
}

func CollectionViewTransitionLayout_InitWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.Alloc().InitWithCurrentLayoutNextLayout(currentLayout, newLayout)
}

func (cc _CollectionViewTransitionLayoutClass) Alloc() CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewTransitionLayout_Alloc() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.Alloc()
}

func (cc _CollectionViewTransitionLayoutClass) New() CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewTransitionLayout() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.New()
}

func CollectionViewTransitionLayout_New() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.New()
}

func (c_ CollectionViewTransitionLayout) Init() CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewTransitionLayout_Init() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.Alloc().Init()
}

func (c_ CollectionViewTransitionLayout) UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("updateValue:forAnimatedKey:"), value, key)
}

func (c_ CollectionViewTransitionLayout) ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("valueForAnimatedKey:"), key)
	return rv
}

func (c_ CollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("transitionProgress"))
	return rv
}

func (c_ CollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTransitionProgress:"), value)
}

func (c_ CollectionViewTransitionLayout) CurrentLayout() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](c_, objc.GetSelector("currentLayout"))
	return rv
}

func (c_ CollectionViewTransitionLayout) NextLayout() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](c_, objc.GetSelector("nextLayout"))
	return rv
}
