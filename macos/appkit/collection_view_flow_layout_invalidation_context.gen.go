// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewFlowLayoutInvalidationContext] class.
var CollectionViewFlowLayoutInvalidationContextClass = _CollectionViewFlowLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewFlowLayoutInvalidationContext")}

type _CollectionViewFlowLayoutInvalidationContextClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewFlowLayoutInvalidationContext] class.
type ICollectionViewFlowLayoutInvalidationContext interface {
	ICollectionViewLayoutInvalidationContext
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

// An object that identifies the portions of a flow layout object that need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayoutinvalidationcontext?language=objc
type CollectionViewFlowLayoutInvalidationContext struct {
	CollectionViewLayoutInvalidationContext
}

func CollectionViewFlowLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContext{
		CollectionViewLayoutInvalidationContext: CollectionViewLayoutInvalidationContextFrom(ptr),
	}
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) Alloc() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Call[CollectionViewFlowLayoutInvalidationContext](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewFlowLayoutInvalidationContext_Alloc() CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContextClass.Alloc()
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) New() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Call[CollectionViewFlowLayoutInvalidationContext](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayoutInvalidationContext() CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewFlowLayoutInvalidationContext) Init() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Call[CollectionViewFlowLayoutInvalidationContext](c_, objc.Sel("init"))
	return rv
}

// A Boolean value indicating whether the flow layout object should invalidate its current attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayoutinvalidationcontext/1402870-invalidateflowlayoutattributes?language=objc
func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	rv := objc.Call[bool](c_, objc.Sel("invalidateFlowLayoutAttributes"))
	return rv
}

// A Boolean value indicating whether the flow layout object should invalidate its current attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayoutinvalidationcontext/1402870-invalidateflowlayoutattributes?language=objc
func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setInvalidateFlowLayoutAttributes:"), value)
}

// A Boolean value indicating whether the flow layout object should fetch new size information from its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayoutinvalidationcontext/1402902-invalidateflowlayoutdelegatemetr?language=objc
func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	rv := objc.Call[bool](c_, objc.Sel("invalidateFlowLayoutDelegateMetrics"))
	return rv
}

// A Boolean value indicating whether the flow layout object should fetch new size information from its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayoutinvalidationcontext/1402902-invalidateflowlayoutdelegatemetr?language=objc
func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setInvalidateFlowLayoutDelegateMetrics:"), value)
}
