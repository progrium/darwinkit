package coremidi

import "github.com/progrium/darwinkit/macos/corefoundation"

// A 96-bit MIDI message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midimessage_96?language=objc
type Message_96 struct {
	Word0 uint32
	Word1 uint32
	Word2 uint32
}

// A series of simultaneous MIDI events in Universal MIDI Packets (UMP) format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midieventpacket?language=objc
type EventPacket struct {
	TimeStamp uint64
	WordCount uint32
	Words     [64]uint32
}

// A custom lookup table to transform MIDI 7-bit values, as contained in note numbers, velocities, control values, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midivaluemap?language=objc
type ValueMap struct {
	Value [128]uint8
}

// The interface to a MIDI driver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/mididriverinterface?language=objc
type DriverInterface struct {
	X_reserved     uintptr
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	FindDevices    uintptr
	Start          uintptr
	Stop           uintptr
	Configure      uintptr
	Send           uintptr
	EnableSource   uintptr
	Flush          uintptr
	Monitor        uintptr
	SendPackets    uintptr
	MonitorEvents  uintptr
}

// A structure that describes a MIDI-CI device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceidentification?language=objc
type CIDeviceIdentification struct {
	Manufacturer  [3]uint8
	Family        [2]uint8
	ModelNumber   [2]uint8
	RevisionLevel [4]uint8
	Reserved      [5]uint8
}

// The transformation of a single type of MIDI event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/miditransform?language=objc
type Transform struct {
	Transform uint16
	Param     int16
}

// A 64-bit MIDI message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midimessage_64?language=objc
type Message_64 struct {
	Word0 uint32
	Word1 uint32
}

// A request to asynchronously transmit a single System Exclusive (SysEx) event to a destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midisysexsendrequest?language=objc
type SysexSendRequest struct {
	Destination      uint32
	Data             *uint8
	BytesToSend      uint32
	Complete         uint8
	Reserved         [3]uint8
	CompletionProc   uintptr
	CompletionRefCon uintptr
}

// A message that describes the change to an object property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiobjectpropertychangenotification?language=objc
type ObjectPropertyChangeNotification struct {
	MessageID    uint32
	MessageSize  uint32
	Object       uint32
	ObjectType   int32
	PropertyName corefoundation.StringRef //*_Ctype_struct___CFString
}

// A message that describes a system state change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinotification?language=objc
type Notification struct {
	MessageID   uint32
	MessageSize uint32
}

// A general I/O error notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiioerrornotification?language=objc
type IOErrorNotification struct {
	MessageID    uint32
	MessageSize  uint32
	DriverDevice uint32
	ErrorCode    int32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiuniversalmessage?language=objc
type UniversalMessage struct {
	Type      uint32
	Group     uint8
	Reserved  [3]uint8
	Utility   uintptr // _Ctype_struct___0
	Pad_cgo_0 [12]byte
}

// A set of MIDI routings and transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midithruconnectionparams?language=objc
type ThruConnectionParams struct {
	Version              uint32
	NumSources           uint32
	Sources              [8]ThruConnectionEndpoint
	NumDestinations      uint32
	Destinations         [8]ThruConnectionEndpoint
	ChannelMap           [16]uint8
	LowVelocity          uint8
	HighVelocity         uint8
	LowNote              uint8
	HighNote             uint8
	NoteNumber           Transform
	Velocity             Transform
	KeyPressure          Transform
	ChannelPressure      Transform
	ProgramChange        Transform
	PitchBend            Transform
	FilterOutSysEx       uint8
	FilterOutMTC         uint8
	FilterOutBeatClock   uint8
	FilterOutTuneRequest uint8
	Reserved2            [3]uint8
	FilterOutAllControls uint8
	NumControlTransforms uint16
	NumMaps              uint16
	Reserved3            [4]uint16
}

// A source or destination in a MIDI thru connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midithruconnectionendpoint?language=objc
type ThruConnectionEndpoint struct {
	EndpointRef uint32
	UniqueID    int32
}

// A list of MIDI events the system sends to or receives from an endpoint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midipacketlist?language=objc
type PacketList struct {
	NumPackets uint32
	Pad_cgo_0  [268]byte
}

// A 128-bit MIDI message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midimessage_128?language=objc
type Message_128 struct {
	Word0 uint32
	Word1 uint32
	Word2 uint32
	Word3 uint32
}

// A collection of simultaneous MIDI events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midipacket?language=objc
type Packet struct {
	TimeStamp uint64
	Length    uint16
	Data      [256]uint8
	Pad_cgo_0 [2]byte
}

// A variable-length list of MIDI event packets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midieventlist?language=objc
type EventList struct {
	Protocol   uint32
	NumPackets uint32
	Packet     [1]EventPacket
}

// A structure that describes the transformation of MIDI control change events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicontroltransform?language=objc
type ControlTransform struct {
	ControlType         uint8
	RemappedControlType uint8
	ControlNumber       uint16
	Transform           uint16
	Param               int16
}

// A message that describes the addition or removal of an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiobjectaddremovenotification?language=objc
type ObjectAddRemoveNotification struct {
	MessageID   uint32
	MessageSize uint32
	Parent      uint32
	ParentType  int32
	Child       uint32
	ChildType   int32
}
