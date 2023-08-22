// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemMetadataCollector] class.
var PlayerItemMetadataCollectorClass = _PlayerItemMetadataCollectorClass{objc.GetClass("AVPlayerItemMetadataCollector")}

type _PlayerItemMetadataCollectorClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemMetadataCollector] class.
type IPlayerItemMetadataCollector interface {
	IPlayerItemMediaDataCollector
	SetDelegateQueue(delegate PPlayerItemMetadataCollectorPushDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	Delegate() PlayerItemMetadataCollectorPushDelegateWrapper
	DelegateQueue() dispatch.Queue
}

// An object used to capture the date range metadata defined for an HTTP Live Streaming asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector?language=objc
type PlayerItemMetadataCollector struct {
	PlayerItemMediaDataCollector
}

func PlayerItemMetadataCollectorFrom(ptr unsafe.Pointer) PlayerItemMetadataCollector {
	return PlayerItemMetadataCollector{
		PlayerItemMediaDataCollector: PlayerItemMediaDataCollectorFrom(ptr),
	}
}

func (p_ PlayerItemMetadataCollector) InitWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) PlayerItemMetadataCollector {
	rv := objc.Call[PlayerItemMetadataCollector](p_, objc.Sel("initWithIdentifiers:classifyingLabels:"), identifiers, classifyingLabels)
	return rv
}

// Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector/1617191-initwithidentifiers?language=objc
func NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) PlayerItemMetadataCollector {
	instance := PlayerItemMetadataCollectorClass.Alloc().InitWithIdentifiersClassifyingLabels(identifiers, classifyingLabels)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemMetadataCollectorClass) Alloc() PlayerItemMetadataCollector {
	rv := objc.Call[PlayerItemMetadataCollector](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemMetadataCollector_Alloc() PlayerItemMetadataCollector {
	return PlayerItemMetadataCollectorClass.Alloc()
}

func (pc _PlayerItemMetadataCollectorClass) New() PlayerItemMetadataCollector {
	rv := objc.Call[PlayerItemMetadataCollector](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemMetadataCollector() PlayerItemMetadataCollector {
	return PlayerItemMetadataCollectorClass.New()
}

func (p_ PlayerItemMetadataCollector) Init() PlayerItemMetadataCollector {
	rv := objc.Call[PlayerItemMetadataCollector](p_, objc.Sel("init"))
	return rv
}

// Sets the delegate and a dispatch queue on which the delegate will be called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector/1617195-setdelegate?language=objc
func (p_ PlayerItemMetadataCollector) SetDelegateQueue(delegate PPlayerItemMetadataCollectorPushDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVPlayerItemMetadataCollectorPushDelegate", delegate)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the delegate and a dispatch queue on which the delegate will be called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector/1617195-setdelegate?language=objc
func (p_ PlayerItemMetadataCollector) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// Accesses the metadata collector’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector/1617196-delegate?language=objc
func (p_ PlayerItemMetadataCollector) Delegate() PlayerItemMetadataCollectorPushDelegateWrapper {
	rv := objc.Call[PlayerItemMetadataCollectorPushDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The dispatch queue on which the delegate’s methods are called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollector/1617192-delegatequeue?language=objc
func (p_ PlayerItemMetadataCollector) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](p_, objc.Sel("delegateQueue"))
	return rv
}
