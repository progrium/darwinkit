// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsheapprovider?language=objc
type PHeapProvider interface {
	// optional
	RetireHeapCacheDelay(heap metal.HeapWrapper, seconds float64)
	HasRetireHeapCacheDelay() bool

	// optional
	NewHeapWithDescriptor(descriptor metal.HeapDescriptor) metal.PHeap
	HasNewHeapWithDescriptor() bool
}

// A concrete type wrapper for the [PHeapProvider] protocol.
type HeapProviderWrapper struct {
	objc.Object
}

func (h_ HeapProviderWrapper) HasRetireHeapCacheDelay() bool {
	return h_.RespondsToSelector(objc.Sel("retireHeap:cacheDelay:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsheapprovider/3229862-retireheap?language=objc
func (h_ HeapProviderWrapper) RetireHeapCacheDelay(heap metal.PHeap, seconds float64) {
	po0 := objc.WrapAsProtocol("MTLHeap", heap)
	objc.Call[objc.Void](h_, objc.Sel("retireHeap:cacheDelay:"), po0, seconds)
}

func (h_ HeapProviderWrapper) HasNewHeapWithDescriptor() bool {
	return h_.RespondsToSelector(objc.Sel("newHeapWithDescriptor:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsheapprovider/3229861-newheapwithdescriptor?language=objc
func (h_ HeapProviderWrapper) NewHeapWithDescriptor(descriptor metal.IHeapDescriptor) metal.HeapWrapper {
	rv := objc.Call[metal.HeapWrapper](h_, objc.Sel("newHeapWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}
