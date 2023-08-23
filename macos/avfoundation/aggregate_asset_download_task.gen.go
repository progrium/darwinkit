// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AggregateAssetDownloadTask] class.
var AggregateAssetDownloadTaskClass = _AggregateAssetDownloadTaskClass{objc.GetClass("AVAggregateAssetDownloadTask")}

type _AggregateAssetDownloadTaskClass struct {
	objc.Class
}

// An interface definition for the [AggregateAssetDownloadTask] class.
type IAggregateAssetDownloadTask interface {
	foundation.IURLSessionTask
	URLAsset() URLAsset
}

// A task that downloads multiple media selections for an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaggregateassetdownloadtask?language=objc
type AggregateAssetDownloadTask struct {
	foundation.URLSessionTask
}

func AggregateAssetDownloadTaskFrom(ptr unsafe.Pointer) AggregateAssetDownloadTask {
	return AggregateAssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

func (ac _AggregateAssetDownloadTaskClass) Alloc() AggregateAssetDownloadTask {
	rv := objc.Call[AggregateAssetDownloadTask](ac, objc.Sel("alloc"))
	return rv
}

func AggregateAssetDownloadTask_Alloc() AggregateAssetDownloadTask {
	return AggregateAssetDownloadTaskClass.Alloc()
}

func (ac _AggregateAssetDownloadTaskClass) New() AggregateAssetDownloadTask {
	rv := objc.Call[AggregateAssetDownloadTask](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAggregateAssetDownloadTask() AggregateAssetDownloadTask {
	return AggregateAssetDownloadTaskClass.New()
}

func (a_ AggregateAssetDownloadTask) Init() AggregateAssetDownloadTask {
	rv := objc.Call[AggregateAssetDownloadTask](a_, objc.Sel("init"))
	return rv
}

// The asset the parent task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaggregateassetdownloadtask/2897238-urlasset?language=objc
func (a_ AggregateAssetDownloadTask) URLAsset() URLAsset {
	rv := objc.Call[URLAsset](a_, objc.Sel("URLAsset"))
	return rv
}
