// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureSession] class.
var CaptureSessionClass = _CaptureSessionClass{objc.GetClass("AVCaptureSession")}

type _CaptureSessionClass struct {
	objc.Class
}

// An interface definition for the [CaptureSession] class.
type ICaptureSession interface {
	objc.IObject
	StopRunning()
	CanSetSessionPreset(preset CaptureSessionPreset) bool
	RemoveInput(input ICaptureInput)
	AddConnection(connection ICaptureConnection)
	BeginConfiguration()
	StartRunning()
	AddInput(input ICaptureInput)
	RemoveConnection(connection ICaptureConnection)
	CanAddInput(input ICaptureInput) bool
	AddInputWithNoConnections(input ICaptureInput)
	AddOutput(output ICaptureOutput)
	CommitConfiguration()
	RemoveOutput(output ICaptureOutput)
	CanAddOutput(output ICaptureOutput) bool
	AddOutputWithNoConnections(output ICaptureOutput)
	CanAddConnection(connection ICaptureConnection) bool
	Connections() []CaptureConnection
	IsRunning() bool
	SessionPreset() CaptureSessionPreset
	SetSessionPreset(value CaptureSessionPreset)
	Outputs() []CaptureOutput
	SynchronizationClock() coremedia.ClockRef
	Inputs() []CaptureInput
}

// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession?language=objc
type CaptureSession struct {
	objc.Object
}

func CaptureSessionFrom(ptr unsafe.Pointer) CaptureSession {
	return CaptureSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureSessionClass) Alloc() CaptureSession {
	rv := objc.Call[CaptureSession](cc, objc.Sel("alloc"))
	return rv
}

func CaptureSession_Alloc() CaptureSession {
	return CaptureSessionClass.Alloc()
}

func (cc _CaptureSessionClass) New() CaptureSession {
	rv := objc.Call[CaptureSession](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureSession() CaptureSession {
	return CaptureSessionClass.New()
}

func (c_ CaptureSession) Init() CaptureSession {
	rv := objc.Call[CaptureSession](c_, objc.Sel("init"))
	return rv
}

// Stops the flow of data through the capture pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1385661-stoprunning?language=objc
func (c_ CaptureSession) StopRunning() {
	objc.Call[objc.Void](c_, objc.Sel("stopRunning"))
}

// Determines whether you can configure a capture session with the specified preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389824-cansetsessionpreset?language=objc
func (c_ CaptureSession) CanSetSessionPreset(preset CaptureSessionPreset) bool {
	rv := objc.Call[bool](c_, objc.Sel("canSetSessionPreset:"), preset)
	return rv
}

// Removes an input from the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388073-removeinput?language=objc
func (c_ CaptureSession) RemoveInput(input ICaptureInput) {
	objc.Call[objc.Void](c_, objc.Sel("removeInput:"), objc.Ptr(input))
}

// Adds a connection to the capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389687-addconnection?language=objc
func (c_ CaptureSession) AddConnection(connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("addConnection:"), objc.Ptr(connection))
}

// Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389174-beginconfiguration?language=objc
func (c_ CaptureSession) BeginConfiguration() {
	objc.Call[objc.Void](c_, objc.Sel("beginConfiguration"))
}

// Starts the flow of data through the capture pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388185-startrunning?language=objc
func (c_ CaptureSession) StartRunning() {
	objc.Call[objc.Void](c_, objc.Sel("startRunning"))
}

// Adds a capture input to the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1387239-addinput?language=objc
func (c_ CaptureSession) AddInput(input ICaptureInput) {
	objc.Call[objc.Void](c_, objc.Sel("addInput:"), objc.Ptr(input))
}

// Removes a capture connection from the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1390041-removeconnection?language=objc
func (c_ CaptureSession) RemoveConnection(connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("removeConnection:"), objc.Ptr(connection))
}

// Determines whether you can add an input to a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1387180-canaddinput?language=objc
func (c_ CaptureSession) CanAddInput(input ICaptureInput) bool {
	rv := objc.Call[bool](c_, objc.Sel("canAddInput:"), objc.Ptr(input))
	return rv
}

// Adds a capture input to a session without forming any connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1390383-addinputwithnoconnections?language=objc
func (c_ CaptureSession) AddInputWithNoConnections(input ICaptureInput) {
	objc.Call[objc.Void](c_, objc.Sel("addInputWithNoConnections:"), objc.Ptr(input))
}

// Adds an output to the capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1387325-addoutput?language=objc
func (c_ CaptureSession) AddOutput(output ICaptureOutput) {
	objc.Call[objc.Void](c_, objc.Sel("addOutput:"), objc.Ptr(output))
}

// Commits one or more changes to a running capture session’s configuration in a single atomic update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388173-commitconfiguration?language=objc
func (c_ CaptureSession) CommitConfiguration() {
	objc.Call[objc.Void](c_, objc.Sel("commitConfiguration"))
}

// Removes an output from a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1385688-removeoutput?language=objc
func (c_ CaptureSession) RemoveOutput(output ICaptureOutput) {
	objc.Call[objc.Void](c_, objc.Sel("removeOutput:"), objc.Ptr(output))
}

// Determines whether you can add an output to a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388944-canaddoutput?language=objc
func (c_ CaptureSession) CanAddOutput(output ICaptureOutput) bool {
	rv := objc.Call[bool](c_, objc.Sel("canAddOutput:"), objc.Ptr(output))
	return rv
}

// Adds a capture output to the session without forming any connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388709-addoutputwithnoconnections?language=objc
func (c_ CaptureSession) AddOutputWithNoConnections(output ICaptureOutput) {
	objc.Call[objc.Void](c_, objc.Sel("addOutputWithNoConnections:"), objc.Ptr(output))
}

// Determines whether a you can add a connection to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389596-canaddconnection?language=objc
func (c_ CaptureSession) CanAddConnection(connection ICaptureConnection) bool {
	rv := objc.Call[bool](c_, objc.Sel("canAddConnection:"), objc.Ptr(connection))
	return rv
}

// The connections between inputs and outputs that a capture session contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/3153020-connections?language=objc
func (c_ CaptureSession) Connections() []CaptureConnection {
	rv := objc.Call[[]CaptureConnection](c_, objc.Sel("connections"))
	return rv
}

// A Boolean value that indicates whether the capture session is in a running state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1388133-running?language=objc
func (c_ CaptureSession) IsRunning() bool {
	rv := objc.Call[bool](c_, objc.Sel("isRunning"))
	return rv
}

// A preset value that indicates the quality level or bit rate of the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389696-sessionpreset?language=objc
func (c_ CaptureSession) SessionPreset() CaptureSessionPreset {
	rv := objc.Call[CaptureSessionPreset](c_, objc.Sel("sessionPreset"))
	return rv
}

// A preset value that indicates the quality level or bit rate of the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1389696-sessionpreset?language=objc
func (c_ CaptureSession) SetSessionPreset(value CaptureSessionPreset) {
	objc.Call[objc.Void](c_, objc.Sel("setSessionPreset:"), value)
}

// The output destinations to which a captures session sends its data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1387621-outputs?language=objc
func (c_ CaptureSession) Outputs() []CaptureOutput {
	rv := objc.Call[[]CaptureOutput](c_, objc.Sel("outputs"))
	return rv
}

// A clock to use for output synchronization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/3915813-synchronizationclock?language=objc
func (c_ CaptureSession) SynchronizationClock() coremedia.ClockRef {
	rv := objc.Call[coremedia.ClockRef](c_, objc.Sel("synchronizationClock"))
	return rv
}

// The inputs that provide media data to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/1390865-inputs?language=objc
func (c_ CaptureSession) Inputs() []CaptureInput {
	rv := objc.Call[[]CaptureInput](c_, objc.Sel("inputs"))
	return rv
}
