package coremediaio

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodeviceavccommand?language=objc
type DeviceAVCCommand struct {
	MCommand        *uint8
	MCommandLength  uint32
	Pad_cgo_0       [8]byte
	MResponseLength uint32
	MResponseUsed   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicers422command?language=objc
type DeviceRS422Command struct {
	MCommand        *uint8
	MCommandLength  uint32
	Pad_cgo_0       [8]byte
	MResponseLength uint32
	MResponseUsed   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicesmptetimecallback?language=objc
type DeviceSMPTETimeCallback struct {
	MGetSMPTETimeProc uintptr
	MRefCon           uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiostreamscheduledoutputnotificationprocandrefcon?language=objc
type StreamScheduledOutputNotificationProcAndRefCon struct {
	ScheduledOutputNotificationProc   uintptr
	ScheduledOutputNotificationRefCon uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicestreamconfiguration?language=objc
type DeviceStreamConfiguration struct {
	MNumberStreams uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioobjectpropertyaddress?language=objc
type ObjectPropertyAddress struct {
	MSelector uint32
	MScope    uint32
	MElement  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiostreamdeck?language=objc
type StreamDeck struct {
	MStatus uint32
	MState  uint32
	MState2 uint32
}

// TODO (unable to generate):
// CMIOHardwarePlugInInterface
