// AUTO-GENERATED CODE, DO NOT MODIFY

package coreaudio

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioclassid?language=objc
type ClassID uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodeviceclockalgorithmselector?language=objc
type DeviceClockAlgorithmSelector uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodeviceid?language=objc
type DeviceID ObjectID

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodevicepropertyid?language=objc
type DevicePropertyID ObjectPropertySelector

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiohardwarepowerhint?language=objc
type HardwarePowerHint uint32

const (
	KHardwarePowerHintFavorSavingPower HardwarePowerHint = 1
	KHardwarePowerHintNone             HardwarePowerHint = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiohardwarepropertyid?language=objc
type HardwarePropertyID ObjectPropertySelector

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiolevelcontroltransferfunction?language=objc
type LevelControlTransferFunction uint32

const (
	KLevelControlTranferFunction10Over1 LevelControlTransferFunction = 13
	KLevelControlTranferFunction11Over1 LevelControlTransferFunction = 14
	KLevelControlTranferFunction12Over1 LevelControlTransferFunction = 15
	KLevelControlTranferFunction1Over2  LevelControlTransferFunction = 2
	KLevelControlTranferFunction1Over3  LevelControlTransferFunction = 1
	KLevelControlTranferFunction2Over1  LevelControlTransferFunction = 5
	KLevelControlTranferFunction3Over1  LevelControlTransferFunction = 6
	KLevelControlTranferFunction3Over2  LevelControlTransferFunction = 4
	KLevelControlTranferFunction3Over4  LevelControlTransferFunction = 3
	KLevelControlTranferFunction4Over1  LevelControlTransferFunction = 7
	KLevelControlTranferFunction5Over1  LevelControlTransferFunction = 8
	KLevelControlTranferFunction6Over1  LevelControlTransferFunction = 9
	KLevelControlTranferFunction7Over1  LevelControlTransferFunction = 10
	KLevelControlTranferFunction8Over1  LevelControlTransferFunction = 11
	KLevelControlTranferFunction9Over1  LevelControlTransferFunction = 12
	KLevelControlTranferFunctionLinear  LevelControlTransferFunction = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectid?language=objc
type ObjectID uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertyelement?language=objc
type ObjectPropertyElement uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertyscope?language=objc
type ObjectPropertyScope uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertyselector?language=objc
type ObjectPropertySelector uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioserverplugincustompropertydatatype?language=objc
type ServerPlugInCustomPropertyDataType uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioserverpluginiooperation?language=objc
type ServerPlugInIOOperation uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiostreamid?language=objc
type StreamID ObjectID
