package iobluetooth

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciextendedinquiryresult?language=objc
type HCIExtendedInquiryResult struct {
	NumberOfReponses        uint8
	DeviceAddress           DeviceAddress
	PageScanRepetitionMode  uint8
	Reserved                uint8
	ClassOfDevice           uint32
	ClockOffset             uint16
	RSSIValue               int8
	ExtendedInquiryResponse HCIExtendedInquiryResponse
	Pad_cgo_0               [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventconnectionrequestresults?language=objc
type HCIEventConnectionRequestResults struct {
	DeviceAddress DeviceAddress
	ClassOfDevice uint32
	LinkType      uint8
	Pad_cgo_0     [3]byte
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexputcommandresponsedata?language=objc
type OBEXPutCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventleconnectionupdatecompleteresults?language=objc
type HCIEventLEConnectionUpdateCompleteResults struct {
	ConnectionHandle   uint16
	ConnInterval       uint16
	ConnLatency        uint16
	SupervisionTimeout uint16
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexconnectcommandresponsedata?language=objc
type OBEXConnectCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
	MaxPacketSize        uint16
	Version              uint8
	Flags                uint8
	Pad_cgo_0            [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothuserconfirmationrequest?language=objc
type UserConfirmationRequest struct {
	DeviceAddress DeviceAddress
	NumericValue  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventconnectionpackettyperesults?language=objc
type HCIEventConnectionPacketTypeResults struct {
	ConnectionHandle uint16
	PacketType       uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcilinkqualityinfo?language=objc
type HCILinkQualityInfo struct {
	Handle       uint16
	QualityValue uint8
	Pad_cgo_0    [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcifailedcontactinfo?language=objc
type HCIFailedContactInfo struct {
	Count  uint16
	Handle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcicurrentinquiryaccesscodesforwrite?language=objc
type HCICurrentInquiryAccessCodesForWrite struct {
	Count uint8
	Codes [192]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothiocapabilityresponse?language=objc
type IOCapabilityResponse struct {
	DeviceAddress              DeviceAddress
	IoCapability               uint8
	OOBDataPresence            uint8
	AuthenticationRequirements uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcirequestcallbackinfo?language=objc
type HCIRequestCallbackInfo struct {
	UserCallback   uint64
	UserRefCon     uint64
	InternalRefCon uint64
	AsyncIDRefCon  uint64
	Reserved       uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcisupportedcommands?language=objc
type HCISupportedCommands struct {
	Data [64]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcilinkpolicysettingsinfo?language=objc
type HCILinkPolicySettingsInfo struct {
	Settings uint16
	Handle   uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventlereadremoteusedfeaturescompleteresults?language=objc
type HCIEventLEReadRemoteUsedFeaturesCompleteResults struct {
	ConnectionHandle uint16
	UsedFeatures     HCISupportedFeatures
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciversioninfo?language=objc
type HCIVersionInfo struct {
	ManufacturerName uint16
	LmpVersion       uint8
	LmpSubVersion    uint16
	HciVersion       uint8
	HciRevision      uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventconnectioncompleteresults?language=objc
type HCIEventConnectionCompleteResults struct {
	ConnectionHandle uint16
	DeviceAddress    DeviceAddress
	LinkType         uint8
	EncryptionMode   uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventmasterlinkkeycompleteresults?language=objc
type HCIEventMasterLinkKeyCompleteResults struct {
	ConnectionHandle uint16
	KeyFlag          uint8
	Pad_cgo_0        [1]byte
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexputcommanddata?language=objc
type OBEXPutCommandData struct {
	HeaderDataPtr      uintptr
	HeaderDataLength   uint64
	BodyDataLeftToSend uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcibuffersize?language=objc
type HCIBufferSize struct {
	ACLDataPacketLength    uint16
	SCODataPacketLength    uint8
	TotalNumACLDataPackets uint16
	TotalNumSCODataPackets uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciextendedinquiryresponse?language=objc
type HCIExtendedInquiryResponse struct {
	Data [240]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventdisconnectioncompleteresults?language=objc
type HCIEventDisconnectionCompleteResults struct {
	ConnectionHandle uint16
	Reason           uint8
	Pad_cgo_0        [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcisetupsynchronousconnectionparams?language=objc
type HCISetupSynchronousConnectionParams struct {
	TransmitBandwidth    uint32
	ReceiveBandwidth     uint32
	MaxLatency           uint16
	VoiceSetting         uint16
	RetransmissionEffort uint8
	PacketType           uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventpagescanrepetitionmodechangeresults?language=objc
type HCIEventPageScanRepetitionModeChangeResults struct {
	DeviceAddress          DeviceAddress
	PageScanRepetitionMode uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcicurrentinquiryaccesscodes?language=objc
type HCICurrentInquiryAccessCodes struct {
	Count uint8
	Codes *HCIInquiryAccessCode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventvendorspecificresults?language=objc
type HCIEventVendorSpecificResults struct {
	Length uint8
	Data   [255]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventqossetupcompleteresults?language=objc
type HCIEventQoSSetupCompleteResults struct {
	ConnectionHandle uint16
	SetupParams      HCIQualityOfServiceSetupParams
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciencryptionkeysizeinfo?language=objc
type HCIEncryptionKeySizeInfo struct {
	Handle    uint16
	KeySize   uint8
	Pad_cgo_0 [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothl2capchannelevent?language=objc
type L2CAPChannelEvent struct {
	EventType uint32
	Pad_cgo_0 [4]byte
	U         [32]byte
	Status    int32
	Pad_cgo_1 [4]byte
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexsetpathcommanddata?language=objc
type OBEXSetPathCommandData struct {
	HeaderDataPtr    uintptr
	HeaderDataLength uint64
	Flags            uint8
	Constants        uint8
	Pad_cgo_0        [6]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothremotehostsupportedfeaturesnotification?language=objc
type RemoteHostSupportedFeaturesNotification struct {
	DeviceAddress         DeviceAddress
	HostSupportedFeatures HCISupportedFeatures
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventflowspecificationdata?language=objc
type HCIEventFlowSpecificationData struct {
	ConnectionHandle uint16
	Flags            uint8
	FlowDirection    uint8
	ServiceType      uint8
	TokenRate        uint32
	TokenBucketSize  uint32
	PeakBandwidth    uint32
	AccessLatency    uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothirk?language=objc
type IRK struct {
	Data [16]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothkeypressnotification?language=objc
type KeypressNotification struct {
	DeviceAddress    DeviceAddress
	NotificationType uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciinquiryaccesscode?language=objc
type HCIInquiryAccessCode struct {
	Data [3]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventchangeconnectionlinkkeycompleteresults?language=objc
type HCIEventChangeConnectionLinkKeyCompleteResults struct {
	ConnectionHandle uint16
}

// Bits to determine what Bluetooth devices to search for [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevicesearchdeviceattributes?language=objc
type DeviceSearchDeviceAttributes struct {
	Address           DeviceAddress
	Name              [248]uint8
	ServiceClassMajor uint32
	DeviceClassMajor  uint32
	DeviceClassMinor  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventmaxslotschangeresults?language=objc
type HCIEventMaxSlotsChangeResults struct {
	ConnectionHandle uint16
	MaxSlots         uint8
	Pad_cgo_0        [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcilebuffersize?language=objc
type HCILEBufferSize struct {
	ACLDataPacketLength    uint16
	TotalNumACLDataPackets uint8
	Pad_cgo_0              [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadextendedfeaturesresults?language=objc
type HCIEventReadExtendedFeaturesResults struct {
	ConnectionHandle      uint16
	SupportedFeaturesInfo HCIExtendedFeaturesInfo
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciscanactivity?language=objc
type HCIScanActivity struct {
	ScanInterval uint16
	ScanWindow   uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcienhancedsetupsynchronousconnectionparams?language=objc
type HCIEnhancedSetupSynchronousConnectionParams struct {
	TransmitBandwidth                 uint32
	ReceiveBandwidth                  uint32
	TransmitCodingFormat              uint64
	ReceiveCodingFormat               uint64
	TransmitCodecFrameSize            uint16
	ReceiveCodecFrameSize             uint16
	InputBandwidth                    uint32
	OutputBandwidth                   uint32
	InputCodingFormat                 uint64
	OutputCodingFormat                uint64
	InputCodedDataSize                uint16
	OutputCodedDataSize               uint16
	InputPCMDataFormat                uint8
	OutputPCMDataFormat               uint8
	InputPCMSamplePayloadMSBPosition  uint8
	OutputPCMSamplePayloadMSBPosition uint8
	InputDataPath                     uint8
	OutputDataPath                    uint8
	InputTransportUnitSize            uint8
	OutputTransportUnitSize           uint8
	MaxLatency                        uint16
	PacketType                        uint16
	RetransmissionEffort              uint8
	Pad_cgo_0                         [7]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothl2capqualityofserviceoptions?language=objc
type L2CAPQualityOfServiceOptions struct {
	Flags           uint8
	ServiceType     uint8
	TokenRate       uint32
	TokenBucketSize uint32
	PeakBandwidth   uint32
	Latency         uint32
	DelayVariation  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcisupportedfeatures?language=objc
type HCISupportedFeatures struct {
	Data [8]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcistoredlinkkeysinfo?language=objc
type HCIStoredLinkKeysInfo struct {
	NumLinkKeysRead               uint16
	MaxNumLinkKeysAllowedInDevice uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventleenhancedconnectioncompleteresults?language=objc
type HCIEventLEEnhancedConnectionCompleteResults struct {
	ConnectionHandle              uint16
	Role                          uint8
	PeerAddressType               uint8
	PeerAddress                   DeviceAddress
	LocalResolvablePrivateAddress DeviceAddress
	PeerResolvablePrivateAddress  DeviceAddress
	ConnInterval                  uint16
	ConnLatency                   uint16
	SupervisionTimeout            uint16
	MasterClockAccuracy           uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventremotenamerequestresults?language=objc
type HCIEventRemoteNameRequestResults struct {
	DeviceAddress DeviceAddress
	DeviceName    [248]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventsynchronousconnectioncompleteresults?language=objc
type HCIEventSynchronousConnectionCompleteResults struct {
	ConnectionHandle     uint16
	DeviceAddress        DeviceAddress
	LinkType             uint8
	TransmissionInterval uint8
	RetransmissionWindow uint8
	ReceivePacketLength  uint16
	TransmitPacketLength uint16
	AirMode              uint8
	Pad_cgo_0            [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcisimplepairingoobdata?language=objc
type HCISimplePairingOOBData struct {
	Data [16]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventlemetaresults?language=objc
type HCIEventLEMetaResults struct {
	Length uint8
	Data   [255]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventsniffsubratingresults?language=objc
type HCIEventSniffSubratingResults struct {
	ConnectionHandle   uint16
	MaxTransmitLatency uint16
	MaxReceiveLatency  uint16
	MinRemoteTimeout   uint16
	MinLocalTimeout    uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadsupportedfeaturesresults?language=objc
type HCIEventReadSupportedFeaturesResults struct {
	ConnectionHandle  uint16
	SupportedFeatures HCISupportedFeatures
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothseteventmask?language=objc
type SetEventMask struct {
	Data [8]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadremoteextendedfeaturesresults?language=objc
type HCIEventReadRemoteExtendedFeaturesResults struct {
	Error            uint8
	ConnectionHandle uint16
	Page             uint8
	MaxPage          uint8
	LmpFeatures      HCISupportedFeatures
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventsimplepairingcompleteresults?language=objc
type HCIEventSimplePairingCompleteResults struct {
	DeviceAddress DeviceAddress
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothuserpasskeynotification?language=objc
type UserPasskeyNotification struct {
	DeviceAddress DeviceAddress
	Passkey       uint32
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexabortcommanddata?language=objc
type OBEXAbortCommandData struct {
	HeaderDataPtr    uintptr
	HeaderDataLength uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothdeviceaddress?language=objc
type DeviceAddress struct {
	Data [6]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcireadextendedinquiryresponseresults?language=objc
type HCIReadExtendedInquiryResponseResults struct {
	OutFECRequired          uint8
	ExtendedInquiryResponse HCIExtendedInquiryResponse
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexconnectcommanddata?language=objc
type OBEXConnectCommandData struct {
	HeaderDataPtr    uintptr
	HeaderDataLength uint64
	MaxPacketSize    uint16
	Version          uint8
	Flags            uint8
	Pad_cgo_0        [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciroleinfo?language=objc
type HCIRoleInfo struct {
	Role   uint8
	Handle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevicesearchattributes?language=objc
type DeviceSearchAttributes struct {
	Options              uint32
	MaxResults           uint32
	DeviceAttributeCount uint32
	AttributeList        *DeviceSearchDeviceAttributes
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventpagescanmodechangeresults?language=objc
type HCIEventPageScanModeChangeResults struct {
	DeviceAddress DeviceAddress
	PageScanMode  uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcienhancedacceptsynchronousconnectionrequestparams?language=objc
type HCIEnhancedAcceptSynchronousConnectionRequestParams struct {
	TransmitBandwidth                 uint32
	ReceiveBandwidth                  uint32
	TransmitCodingFormat              uint64
	ReceiveCodingFormat               uint64
	TransmitCodecFrameSize            uint16
	ReceiveCodecFrameSize             uint16
	InputBandwidth                    uint32
	OutputBandwidth                   uint32
	InputCodingFormat                 uint64
	OutputCodingFormat                uint64
	InputCodedDataSize                uint16
	OutputCodedDataSize               uint16
	InputPCMDataFormat                uint8
	OutputPCMDataFormat               uint8
	InputPCMSamplePayloadMSBPosition  uint8
	OutputPCMSamplePayloadMSBPosition uint8
	InputDataPath                     uint8
	OutputDataPath                    uint8
	InputTransportUnitSize            uint8
	OutputTransportUnitSize           uint8
	MaxLatency                        uint16
	PacketType                        uint16
	RetransmissionEffort              uint8
	Pad_cgo_0                         [7]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothreadclockinfo?language=objc
type ReadClockInfo struct {
	Handle    uint16
	Clock     uint32
	Accuracy  uint16
	Pad_cgo_0 [2]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventauthenticationcompleteresults?language=objc
type HCIEventAuthenticationCompleteResults struct {
	ConnectionHandle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothenhancedsynchronousconnectioninfo?language=objc
type EnhancedSynchronousConnectionInfo struct {
	TransmitBandWidth                 uint32
	ReceiveBandWidth                  uint32
	TransmitCodingFormat              uint64
	ReceiveCodingFormat               uint64
	TransmitCodecFrameSize            uint16
	ReceiveCodecFrameSize             uint16
	InputBandwidth                    uint32
	OutputBandwidth                   uint32
	InputCodingFormat                 uint64
	OutputCodingFormat                uint64
	InputCodedDataSize                uint16
	OutputCodedDataSize               uint16
	InputPCMDataFormat                uint8
	OutputPCMDataFormat               uint8
	InputPCMSampelPayloadMSBPosition  uint8
	OutputPCMSampelPayloadMSBPosition uint8
	InputDataPath                     uint8
	OutputDataPath                    uint8
	InputTransportUnitSize            uint8
	OutputTransportUnitSize           uint8
	MaxLatency                        uint16
	VoiceSetting                      uint16
	RetransmissionEffort              uint8
	PacketType                        uint16
	Pad_cgo_0                         [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadclockoffsetresults?language=objc
type HCIEventReadClockOffsetResults struct {
	ConnectionHandle uint16
	ClockOffset      uint16
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexabortcommandresponsedata?language=objc
type OBEXAbortCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothl2capchanneldatablock?language=objc
type L2CAPChannelDataBlock struct {
	DataPtr  uintptr
	DataSize uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothsynchronousconnectioninfo?language=objc
type SynchronousConnectionInfo struct {
	TransmitBandWidth    uint32
	ReceiveBandWidth     uint32
	MaxLatency           uint16
	VoiceSetting         uint16
	RetransmissionEffort uint8
	PacketType           uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothpincode?language=objc
type PINCode struct {
	Data [16]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciinquirywithrssiresults?language=objc
type HCIInquiryWithRSSIResults struct {
	Results [50]HCIInquiryWithRSSIResult
	Count   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcitransmitpowerlevelinfo?language=objc
type HCITransmitPowerLevelInfo struct {
	Handle    uint16
	Level     int8
	Pad_cgo_0 [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciqualityofservicesetupparams?language=objc
type HCIQualityOfServiceSetupParams struct {
	Flags          uint8
	ServiceType    uint8
	TokenRate      uint32
	PeakBandwidth  uint32
	Latency        uint32
	DelayVariation uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciacceptsynchronousconnectionrequestparams?language=objc
type HCIAcceptSynchronousConnectionRequestParams struct {
	TransmitBandwidth    uint32
	ReceiveBandwidth     uint32
	MaxLatency           uint16
	ContentFormat        uint16
	RetransmissionEffort uint8
	PacketType           uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciinquiryresults?language=objc
type HCIInquiryResults struct {
	Results [50]HCIInquiryResult
	Count   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexsessionevent?language=objc
type OBEXSessionEvent struct {
	Type uint32
	//TODO: Session			*_Ctype_struct_OpaqueOBEXSessionRef
	RefCon           uintptr
	IsEndOfEventData uint8
	Reserved1        uintptr
	Reserved2        uintptr
	U                [32]byte
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexgetcommandresponsedata?language=objc
type OBEXGetCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciextendedfeaturesinfo?language=objc
type HCIExtendedFeaturesInfo struct {
	Page    uint8
	MaxPage uint8
	Data    [8]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventhardwareerrorresults?language=objc
type HCIEventHardwareErrorResults struct {
	Error uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciinquirywithrssiresult?language=objc
type HCIInquiryWithRSSIResult struct {
	DeviceAddress          DeviceAddress
	PageScanRepetitionMode uint8
	Reserved               uint8
	ClassOfDevice          uint32
	ClockOffset            uint16
	RSSIValue              int8
	Pad_cgo_0              [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventflushoccurredresults?language=objc
type HCIEventFlushOccurredResults struct {
	ConnectionHandle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadremotesupportedfeaturesresults?language=objc
type HCIEventReadRemoteSupportedFeaturesResults struct {
	Error            uint8
	ConnectionHandle uint16
	LmpFeatures      HCISupportedFeatures
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcireadlocaloobdataresults?language=objc
type HCIReadLocalOOBDataResults struct {
	Hash       HCISimplePairingOOBData
	Randomizer HCISimplePairingOOBData
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexsetpathcommandresponsedata?language=objc
type OBEXSetPathCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
	Flags                uint8
	Constants            uint8
	Pad_cgo_0            [6]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventsynchronousconnectionchangedresults?language=objc
type HCIEventSynchronousConnectionChangedResults struct {
	ConnectionHandle     uint16
	TransmissionInterval uint8
	RetransmissionWindow uint8
	ReceivePacketLength  uint16
	TransmitPacketLength uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventlelongtermkeyrequestresults?language=objc
type HCIEventLELongTermKeyRequestResults struct {
	ConnectionHandle uint16
	RandomNumber     [8]uint8
	Ediv             uint16
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexerrordata?language=objc
type OBEXErrorData struct {
	Error      int32
	DataPtr    uintptr
	DataLength uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obextransportevent?language=objc
type OBEXTransportEvent struct {
	Type       uint32
	Status     int32
	DataPtr    uintptr
	DataLength uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothtransportinfo?language=objc
type TransportInfo struct {
	ProductID              uint32
	VendorID               uint32
	Type                   uint32
	ProductName            [35]int8
	VendorName             [35]int8
	TotalDataBytesSent     uint64
	TotalSCOBytesSent      uint64
	TotalDataBytesReceived uint64
	TotalSCOBytesReceived  uint64
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexdisconnectcommandresponsedata?language=objc
type OBEXDisconnectCommandResponseData struct {
	ServerResponseOpCode uint8
	HeaderDataPtr        uintptr
	HeaderDataLength     uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciautomaticflushtimeoutinfo?language=objc
type HCIAutomaticFlushTimeoutInfo struct {
	Handle  uint16
	Timeout uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreturnlinkkeysresults?language=objc
type HCIEventReturnLinkKeysResults struct {
	NumLinkKeys uint8
	// TODO: LinkKeys	[1]_Ctype_struct___0
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcilinksupervisiontimeout?language=objc
type HCILinkSupervisionTimeout struct {
	Handle  uint16
	Timeout uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventmodechangeresults?language=objc
type HCIEventModeChangeResults struct {
	ConnectionHandle uint16
	Mode             uint8
	ModeInterval     uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcirssiinfo?language=objc
type HCIRSSIInfo struct {
	Handle    uint16
	RSSIValue int8
	Pad_cgo_0 [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventleconnectioncompleteresults?language=objc
type HCIEventLEConnectionCompleteResults struct {
	ConnectionHandle    uint16
	Role                uint8
	PeerAddressType     uint8
	PeerAddress         DeviceAddress
	ConnInterval        uint16
	ConnLatency         uint16
	SupervisionTimeout  uint16
	MasterClockAccuracy uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventreadremoteversioninforesults?language=objc
type HCIEventReadRemoteVersionInfoResults struct {
	ConnectionHandle uint16
	LmpVersion       uint8
	ManufacturerName uint16
	LmpSubversion    uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcireadlmphandleresults?language=objc
type HCIReadLMPHandleResults struct {
	Handle     uint16
	Lmp_handle uint8
	Reserved   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventlinkkeynotificationresults?language=objc
type HCIEventLinkKeyNotificationResults struct {
	DeviceAddress DeviceAddress
	LinkKey       Key
	KeyType       uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventqosviolationresults?language=objc
type HCIEventQoSViolationResults struct {
	ConnectionHandle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetootheventfiltercondition?language=objc
type EventFilterCondition struct {
	Data [7]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothafhhostchannelclassification?language=objc
type AFHHostChannelClassification struct {
	Data [10]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothafhresults?language=objc
type AFHResults struct {
	Handle    uint16
	Mode      uint8
	AfhMap    [10]uint8
	Pad_cgo_0 [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventdatabufferoverflowresults?language=objc
type HCIEventDataBufferOverflowResults struct {
	LinkType uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventencryptionchangeresults?language=objc
type HCIEventEncryptionChangeResults struct {
	ConnectionHandle uint16
	Enable           uint8
	Pad_cgo_0        [1]byte
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexgetcommanddata?language=objc
type OBEXGetCommandData struct {
	HeaderDataPtr    uintptr
	HeaderDataLength uint64
}

// Part of the OBEXSessionEvent structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexdisconnectcommanddata?language=objc
type OBEXDisconnectCommandData struct {
	HeaderDataPtr    uintptr
	HeaderDataLength uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventencryptionkeyrefreshcompleteresults?language=objc
type HCIEventEncryptionKeyRefreshCompleteResults struct {
	ConnectionHandle uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhciinquiryresult?language=objc
type HCIInquiryResult struct {
	DeviceAddress          DeviceAddress
	PageScanRepetitionMode uint8
	PageScanPeriodMode     uint8
	PageScanMode           uint8
	ClassOfDevice          uint32
	ClockOffset            uint16
	Pad_cgo_0              [2]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothkey?language=objc
type Key struct {
	Data [16]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothhcieventrolechangeresults?language=objc
type HCIEventRoleChangeResults struct {
	ConnectionHandle uint16
	DeviceAddress    DeviceAddress
	Role             uint8
	Pad_cgo_0        [1]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/bluetoothl2capretransmissionandflowcontroloptions?language=objc
type L2CAPRetransmissionAndFlowControlOptions struct {
	Flags                 uint8
	TxWindowSize          uint8
	MaxTransmit           uint8
	RetransmissionTimeout uint16
	MonitorTimeout        uint16
	MaxPDUPayloadSize     uint16
}

// TODO (unable to generate):
// BluetoothHCILinkPolicySettingsValues BluetoothHCIAFHChannelAssessmentModes FTSFileType OBEXRealmValues BluetoothLinkTypes BluetoothHCIGeneralFlowControlStates BluetoothHCIExtendedInquiryResponseDataTypes OBEXErrorCodes OBEXTransportEventTypes BluetoothOOBDataPresenceValues OBEXPutFlagValues BluetoothHCIDeleteStoredLinkKeyFlags BluetoothHCIEncryptionModes OBEXNonceFlagValues BluetoothHCIRoles BluetoothCompanyIdentifers BluetoothHCIPageScanEnableStates BluetoothHCITimeoutValues BluetoothSimplePairingDebugModes OBEXOpCodeResponseValues BluetoothHCIInquiryScanTypes BluetoothKeypressNotificationTypes SDPServiceClasses OBEXOpCodeCommandValues BluetoothHCIFECRequiredValues BluetoothLEFeatureBits BluetoothHCIRetransmissionEffortTypes ProtocolParameters OBEXSessionEventTypes BluetoothTransportTypes BluetoothHCISimplePairingModes BluetoothFeatureBits BluetoothHCIReadStoredLinkKeysFlags IOBluetoothDeviceSearchOptionsBits BluetoothHCIPageScanPeriodModes OBEXSessionParameterTags OBEXVersions BluetoothHCIAuthentionEnableModes BluetoothIOCapabilities BluetoothAuthenticationRequirementsValues BluetoothHCIConnectionModes OBEXConnectFlagValues SDPAttributeIdentifierCodes BluetoothLESecurityManagerKeyDistributionFormat IOBluetoothDeviceSearchTypesBits BluetoothHCIInquiryModes BluetoothHCIPageScanTypes BluetoothHCITransmitReadPowerLevelTypes OBEXOpCodeSessionValues BluetoothHCIPageScanModes BluetoothHCISCOFlowControlStates SDPAttributeDeviceIdentificationRecord BluetoothHCIHoldModeActivityStates OBEXHeaderIdentifiers
