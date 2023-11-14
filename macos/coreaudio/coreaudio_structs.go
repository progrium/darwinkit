package coreaudio

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertyaddress?language=objc
type ObjectPropertyAddress struct {
	MSelector uint32
	MScope    uint32
	MElement  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiostreamrangeddescription?language=objc
type StreamRangedDescription struct {
	MFormat          uintptr // _Ctype_struct_AudioStreamBasicDescription
	MSampleRateRange uintptr // _Ctype_struct_AudioValueRange
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiohardwareioprocstreamusage?language=objc
type HardwareIOProcStreamUsage struct {
	MIOProc        uintptr
	MNumberStreams uint32
	MStreamIsOn    [1]uint32
}

// TODO (unable to generate):
// AudioServerPlugInClientInfo AudioDriverPlugInHostInfo AudioServerPlugInHostInterface AudioServerPlugInIOCycleInfo AudioServerPlugInCustomPropertyInfo AudioServerPlugInDriverInterface
