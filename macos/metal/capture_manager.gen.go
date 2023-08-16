// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureManager] class.
var CaptureManagerClass = _CaptureManagerClass{objc.GetClass("MTLCaptureManager")}

type _CaptureManagerClass struct {
	objc.Class
}

// An interface definition for the [CaptureManager] class.
type ICaptureManager interface {
	objc.IObject
	NewCaptureScopeWithDevice(device PDevice) CaptureScopeWrapper
	NewCaptureScopeWithDeviceObject(deviceObject objc.IObject) CaptureScopeWrapper
	StopCapture()
	StartCaptureWithDescriptorError(descriptor ICaptureDescriptor, error foundation.IError) bool
	NewCaptureScopeWithCommandQueue(commandQueue PCommandQueue) CaptureScopeWrapper
	NewCaptureScopeWithCommandQueueObject(commandQueueObject objc.IObject) CaptureScopeWrapper
	SupportsDestination(destination CaptureDestination) bool
	DefaultCaptureScope() CaptureScopeWrapper
	SetDefaultCaptureScope(value PCaptureScope)
	SetDefaultCaptureScopeObject(valueObject objc.IObject)
	IsCapturing() bool
}

// An instance you use to capture Metal command data in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager?language=objc
type CaptureManager struct {
	objc.Object
}

func CaptureManagerFrom(ptr unsafe.Pointer) CaptureManager {
	return CaptureManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureManagerClass) Alloc() CaptureManager {
	rv := objc.Call[CaptureManager](cc, objc.Sel("alloc"))
	return rv
}

func CaptureManager_Alloc() CaptureManager {
	return CaptureManagerClass.Alloc()
}

func (cc _CaptureManagerClass) New() CaptureManager {
	rv := objc.Call[CaptureManager](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureManager() CaptureManager {
	return CaptureManagerClass.New()
}

func (c_ CaptureManager) Init() CaptureManager {
	rv := objc.Call[CaptureManager](c_, objc.Sel("init"))
	return rv
}

// Creates a capture scope for commands submitted to a specific device object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869719-newcapturescopewithdevice?language=objc
func (c_ CaptureManager) NewCaptureScopeWithDevice(device PDevice) CaptureScopeWrapper {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CaptureScopeWrapper](c_, objc.Sel("newCaptureScopeWithDevice:"), po0)
	rv.Autorelease()
	return rv
}

// Creates a capture scope for commands submitted to a specific device object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869719-newcapturescopewithdevice?language=objc
func (c_ CaptureManager) NewCaptureScopeWithDeviceObject(deviceObject objc.IObject) CaptureScopeWrapper {
	rv := objc.Call[CaptureScopeWrapper](c_, objc.Sel("newCaptureScopeWithDevice:"), objc.Ptr(deviceObject))
	rv.Autorelease()
	return rv
}

// Provides the shared capture manager for your Metal app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869720-sharedcapturemanager?language=objc
func (cc _CaptureManagerClass) SharedCaptureManager() CaptureManager {
	rv := objc.Call[CaptureManager](cc, objc.Sel("sharedCaptureManager"))
	return rv
}

// Provides the shared capture manager for your Metal app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869720-sharedcapturemanager?language=objc
func CaptureManager_SharedCaptureManager() CaptureManager {
	return CaptureManagerClass.SharedCaptureManager()
}

// Stops capturing Metal commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869736-stopcapture?language=objc
func (c_ CaptureManager) StopCapture() {
	objc.Call[objc.Void](c_, objc.Sel("stopCapture"))
}

// Starts capturing any of your app’s Metal commands, with the capture session defined by a descriptor object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/3237259-startcapturewithdescriptor?language=objc
func (c_ CaptureManager) StartCaptureWithDescriptorError(descriptor ICaptureDescriptor, error foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("startCaptureWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	return rv
}

// Creates a capture scope for commands submitted to a specific command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869732-newcapturescopewithcommandqueue?language=objc
func (c_ CaptureManager) NewCaptureScopeWithCommandQueue(commandQueue PCommandQueue) CaptureScopeWrapper {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	rv := objc.Call[CaptureScopeWrapper](c_, objc.Sel("newCaptureScopeWithCommandQueue:"), po0)
	rv.Autorelease()
	return rv
}

// Creates a capture scope for commands submitted to a specific command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869732-newcapturescopewithcommandqueue?language=objc
func (c_ CaptureManager) NewCaptureScopeWithCommandQueueObject(commandQueueObject objc.IObject) CaptureScopeWrapper {
	rv := objc.Call[CaptureScopeWrapper](c_, objc.Sel("newCaptureScopeWithCommandQueue:"), objc.Ptr(commandQueueObject))
	rv.Autorelease()
	return rv
}

// Checks to see whether a particular capture destination is supported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/3237260-supportsdestination?language=objc
func (c_ CaptureManager) SupportsDestination(destination CaptureDestination) bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsDestination:"), destination)
	return rv
}

// The capture scope to use when a capture is initiated in Xcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2887120-defaultcapturescope?language=objc
func (c_ CaptureManager) DefaultCaptureScope() CaptureScopeWrapper {
	rv := objc.Call[CaptureScopeWrapper](c_, objc.Sel("defaultCaptureScope"))
	return rv
}

// The capture scope to use when a capture is initiated in Xcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2887120-defaultcapturescope?language=objc
func (c_ CaptureManager) SetDefaultCaptureScope(value PCaptureScope) {
	po0 := objc.WrapAsProtocol("MTLCaptureScope", value)
	objc.Call[objc.Void](c_, objc.Sel("setDefaultCaptureScope:"), po0)
}

// The capture scope to use when a capture is initiated in Xcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2887120-defaultcapturescope?language=objc
func (c_ CaptureManager) SetDefaultCaptureScopeObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDefaultCaptureScope:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether Metal commands are being captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/2869727-iscapturing?language=objc
func (c_ CaptureManager) IsCapturing() bool {
	rv := objc.Call[bool](c_, objc.Sel("isCapturing"))
	return rv
}
