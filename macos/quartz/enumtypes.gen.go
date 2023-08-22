// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

// These constants specify the display mode used by the camera view. These constants are used by mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdisplaymode?language=objc
type CameraDeviceViewDisplayMode int

const (
	CameraDeviceViewDisplayModeIcon  CameraDeviceViewDisplayMode = 1
	CameraDeviceViewDisplayModeNone  CameraDeviceViewDisplayMode = -1
	CameraDeviceViewDisplayModeTable CameraDeviceViewDisplayMode = 0
)

// These constants specify the transfer mode used by the camera view. These constants are used by mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewtransfermode?language=objc
type CameraDeviceViewTransferMode int

const (
	CameraDeviceViewTransferModeFileBased   CameraDeviceViewTransferMode = 0
	CameraDeviceViewTransferModeMemoryBased CameraDeviceViewTransferMode = 1
)

// These constants specify the display mode of the device browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserviewdisplaymode?language=objc
type DeviceBrowserViewDisplayMode int

const (
	DeviceBrowserViewDisplayModeIcon    DeviceBrowserViewDisplayMode = 2
	DeviceBrowserViewDisplayModeOutline DeviceBrowserViewDisplayMode = 1
	DeviceBrowserViewDisplayModeTable   DeviceBrowserViewDisplayMode = 0
)

// The possible states for the browser cell. These values are used by the cellState method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercellstate?language=objc
type ImageBrowserCellState int

const (
	ImageStateInvalid ImageBrowserCellState = 1
	ImageStateNoImage ImageBrowserCellState = 0
	ImageStateReady   ImageBrowserCellState = 2
)

// These constants specify the locations for dropping items onto the browser view. Used by the method setDropIndex:dropOperation:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserdropoperation?language=objc
type ImageBrowserDropOperation int

const (
	ImageBrowserDropBefore ImageBrowserDropOperation = 1
	ImageBrowserDropOn     ImageBrowserDropOperation = 0
)

// These constants specify the display mode the scanner view will use. They are used by the mode property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdisplaymode?language=objc
type ScannerDeviceViewDisplayMode int

const (
	ScannerDeviceViewDisplayModeAdvanced ScannerDeviceViewDisplayMode = 1
	ScannerDeviceViewDisplayModeNone     ScannerDeviceViewDisplayMode = -1
	ScannerDeviceViewDisplayModeSimple   ScannerDeviceViewDisplayMode = 0
)

// These constants determine how the scanner data is returned to the delegate. They are used by the transferMode property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewtransfermode?language=objc
type ScannerDeviceViewTransferMode int

const (
	ScannerDeviceViewTransferModeFileBased   ScannerDeviceViewTransferMode = 0
	ScannerDeviceViewTransferModeMemoryBased ScannerDeviceViewTransferMode = 1
)

// Execution modes for custom patches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcpluginexecutionmode?language=objc
type PlugInExecutionMode int

const (
	KPlugInExecutionModeConsumer  PlugInExecutionMode = 3
	KPlugInExecutionModeProcessor PlugInExecutionMode = 2
	KPlugInExecutionModeProvider  PlugInExecutionMode = 1
)

// Time modes for custom patches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcplugintimemode?language=objc
type PlugInTimeMode int

const (
	KPlugInTimeModeIdle     PlugInTimeMode = 1
	KPlugInTimeModeNone     PlugInTimeMode = 0
	KPlugInTimeModeTimeBase PlugInTimeMode = 2
)
