// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadTask] class.
var AssetDownloadTaskClass = _AssetDownloadTaskClass{objc.GetClass("AVAssetDownloadTask")}

type _AssetDownloadTaskClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadTask] class.
type IAssetDownloadTask interface {
	foundation.IURLSessionTask
	Options() map[string]objc.Object
	LoadedTimeRanges() []foundation.Value
	URLAsset() URLAsset
}

// A session used to download HTTP Live Streaming assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask?language=objc
type AssetDownloadTask struct {
	foundation.URLSessionTask
}

func AssetDownloadTaskFrom(ptr unsafe.Pointer) AssetDownloadTask {
	return AssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

func (ac _AssetDownloadTaskClass) Alloc() AssetDownloadTask {
	rv := objc.Call[AssetDownloadTask](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadTask_Alloc() AssetDownloadTask {
	return AssetDownloadTaskClass.Alloc()
}

func (ac _AssetDownloadTaskClass) New() AssetDownloadTask {
	rv := objc.Call[AssetDownloadTask](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadTask() AssetDownloadTask {
	return AssetDownloadTaskClass.New()
}

func (a_ AssetDownloadTask) Init() AssetDownloadTask {
	rv := objc.Call[AssetDownloadTask](a_, objc.Sel("init"))
	return rv
}

// The configuration options for the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/1621014-options?language=objc
func (a_ AssetDownloadTask) Options() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("options"))
	return rv
}

// The time ranges of the downloaded media that are ready for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/1621026-loadedtimeranges?language=objc
func (a_ AssetDownloadTask) LoadedTimeRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](a_, objc.Sel("loadedTimeRanges"))
	return rv
}

// The asset that this task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/1621024-urlasset?language=objc
func (a_ AssetDownloadTask) URLAsset() URLAsset {
	rv := objc.Call[URLAsset](a_, objc.Sel("URLAsset"))
	return rv
}
