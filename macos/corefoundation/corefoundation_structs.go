package corefoundation

// A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopTimer’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunlooptimercontext?language=objc
type RunLoopTimerContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// Defines the buffer and related fields used for in-line buffer access of characters in CFString objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringinlinebuffer?language=objc
type StringInlineBuffer struct {
	Buffer              [64]uint16
	TheString           uintptr
	DirectUniCharBuffer *uint16
	DirectCStringBuffer *int8
	RangeToBuffer       Range
	BufferedRangeStart  int64
	BufferedRangeEnd    int64
}

// A reference to a CFNull object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnullref?language=objc
type NullRef uintptr

// A reference to a CFBoolean object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbooleanref?language=objc
type BooleanRef uintptr

// Structure used to represent a point in time using the Gregorian calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfgregoriandate?language=objc
type GregorianDate struct {
	Year   int32
	Month  int8
	Day    int8
	Hour   int8
	Minute int8
	Second float64
}

// Contains information about an element attribute definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlattributedeclarationinfo?language=objc
type XMLAttributeDeclarationInfo struct {
	AttributeName uintptr
	TypeString    uintptr
	DefaultString uintptr
}

// Contains version information and function pointers to callbacks needed when parsing XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparsercallbacks?language=objc
type XMLParserCallBacks struct {
	Version               int64
	CreateXMLStructure    uintptr
	AddChild              uintptr
	EndXMLStructure       uintptr
	ResolveExternalEntity uintptr
	HandleError           uintptr
}

// Contains the system and public IDs for an external entity reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlexternalid?language=objc
type XMLExternalID struct {
	SystemID uintptr
	PublicID uintptr
}

// A reference to a CFStringTokenizer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringtokenizerref?language=objc
type StringTokenizerRef uintptr

// Defines a structure for the context of a CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffiledescriptorcontext?language=objc
type FileDescriptorContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A reference to a CFCalendar object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcalendarref?language=objc
type CalendarRef uintptr

// Structure used to represent a time interval in Gregorian units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfgregorianunits?language=objc
type GregorianUnits struct {
	Years   int32
	Months  int32
	Days    int32
	Hours   int32
	Minutes int32
	Seconds float64
}

// A reference to a CFAttributedString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfattributedstringref?language=objc
type AttributedStringRef uintptr

// A reference to a CFURL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlref?language=objc
type URLRef uintptr

// A reference to an CFFileDescriptor object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffiledescriptorref?language=objc
type FileDescriptorRef uintptr

// A reference to an immutable character set object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcharactersetref?language=objc
type CharacterSetRef uintptr

// Encapsulates a file system object’s security information in a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffilesecurityref?language=objc
type FileSecurityRef uintptr

// Contains the source URL and text encoding information for the XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmldocumentinfo?language=objc
type XMLDocumentInfo struct {
	SourceURL uintptr
	Encoding  uint32
	Pad_cgo_0 [4]byte
}

// A reference to a CFTree object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreeref?language=objc
type TreeRef uintptr

// A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopObserver object’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopobservercontext?language=objc
type RunLoopObserverContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A structure that fully specifies the communication protocol and connection address of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketsignature?language=objc
type SocketSignature struct {
	ProtocolFamily int32
	SocketType     int32
	Protocol       int32
	Address        uintptr
}

// Structure holding a 64-bit float value in a platform-independentbyte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfswappedfloat64?language=objc
type SwappedFloat64 struct {
	V uint64
}

// Contains the text of the processing instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlprocessinginstructioninfo?language=objc
type XMLProcessingInstructionInfo struct {
	DataString uintptr
}

// A reference to an immutable set object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetref?language=objc
type SetRef uintptr

// Contains a description of the element type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlelementtypedeclarationinfo?language=objc
type XMLElementTypeDeclarationInfo struct {
	ContentDescription uintptr
}

// Structure containing program-defined data and callbacks for a CFTree object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreecontext?language=objc
type TreeContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// Contains information describing an XML entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlentityinfo?language=objc
type XMLEntityInfo struct {
	EntityType      uint64
	ReplacementText uintptr
	EntityID        XMLExternalID
	NotationName    uintptr
}

// Not used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbinaryheapcomparecontext?language=objc
type BinaryHeapCompareContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// Structure containing the callbacks of a CFArray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarraycallbacks?language=objc
type ArrayCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
}

// A structure that defines the context or operating environment for an allocator (CFAllocator) object. Every Core Foundation allocator object must have a context defined for it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorcontext?language=objc
type AllocatorContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Allocate        uintptr
	Reallocate      uintptr
	Deallocate      uintptr
	PreferredSize   uintptr
}

// A reference to a user notification object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfusernotificationref?language=objc
type UserNotificationRef uintptr

// A reference to an immutable CFDate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdateref?language=objc
type DateRef uintptr

// A reference to a CFString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringref?language=objc
type StringRef uintptr

// A reference to a run loop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopref?language=objc
type RunLoopRef uintptr

// The structure returned by [corefoundation/cfreadstreamgeterror] and [corefoundation/cfwritestreamgeterror]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstreamerror?language=objc
type StreamError struct {
	Domain    int64
	Error     int32
	Pad_cgo_0 [4]byte
}

// Structure holding a 32-bit float value in a platform-independentbyte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfswappedfloat32?language=objc
type SwappedFloat32 struct {
	V uint32
}

// This structure contains the callbacks used to retain, release, describe, and compare the values of a CFSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetcallbacks?language=objc
type SetCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
	Hash            uintptr
}

// A reference to a writable stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfwritestreamref?language=objc
type WriteStreamRef uintptr

// Contains the external ID of the notation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlnotationinfo?language=objc
type XMLNotationInfo struct {
	ExternalID XMLExternalID
}

// A reference to a CFAllocator object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorref?language=objc
type AllocatorRef uintptr

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfplugininstanceref?language=objc
type PlugInInstanceRef uintptr

// Contains the external ID of the DTD. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmldocumenttypeinfo?language=objc
type XMLDocumentTypeInfo struct {
	ExternalID XMLExternalID
}

// This structure contains the callbacks used to retain, release, describe, and compare the values of a CFBag object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagcallbacks?language=objc
type BagCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
	Hash            uintptr
}

// A reference to a run loop timer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunlooptimerref?language=objc
type RunLoopTimerRef uintptr

// A reference to a CFLocale object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cflocaleref?language=objc
type LocaleRef uintptr

// A reference to an immutable bag object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagref?language=objc
type BagRef uintptr

// A reference to a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketref?language=objc
type SocketRef uintptr

// A reference to an immutable dictionary object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryref?language=objc
type DictionaryRef uintptr

// A 128-bit struct that represents a UUID as raw bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfuuidbytes?language=objc
type UUIDBytes struct {
	Byte0  uint8
	Byte1  uint8
	Byte2  uint8
	Byte3  uint8
	Byte4  uint8
	Byte5  uint8
	Byte6  uint8
	Byte7  uint8
	Byte8  uint8
	Byte9  uint8
	Byte10 uint8
	Byte11 uint8
	Byte12 uint8
	Byte13 uint8
	Byte14 uint8
	Byte15 uint8
}

// This structure contains the callbacks used to retain, release, describe, and compare the values in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryvaluecallbacks?language=objc
type DictionaryValueCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
}

// A reference to a CFDateFormatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdateformatterref?language=objc
type DateFormatterRef uintptr

// Contains version information and function pointers to callbacks used when handling a program-defined context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparsercontext?language=objc
type XMLParserContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// This structure contains the callbacks used to retain, release, describe, and compare the keys in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionarykeycallbacks?language=objc
type DictionaryKeyCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
	Hash            uintptr
}

// A structure that contains program-defined data and callbacks with which you can configure a CFMachPort object’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmachportcontext?language=objc
type MachPortContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A structure that contains program-defined data and callbacks with which you can configure a version 0 CFRunLoopSource’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopsourcecontext?language=objc
type RunLoopSourceContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
	Hash            uintptr
	Schedule        uintptr
	Cancel          uintptr
	Perform         uintptr
}

// Contains information describing an XML entity reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlentityreferenceinfo?language=objc
type XMLEntityReferenceInfo struct {
	EntityType uint64
}

// A reference to a CFNumberFormatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberformatterref?language=objc
type NumberFormatterRef uintptr

// A reference to a CFNumber object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberref?language=objc
type NumberRef uintptr

// A reference to an immutable CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdataref?language=objc
type DataRef uintptr

// A reference to a message port object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmessageportref?language=objc
type MessagePortRef uintptr

// A reference to a CFURLEnumerator object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlenumeratorref?language=objc
type URLEnumeratorRef uintptr

// The type of a reference to a CFNotificationCenter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnotificationcenterref?language=objc
type NotificationCenterRef uintptr

// Contains a list of the attributes associated with an element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlattributelistdeclarationinfo?language=objc
type XMLAttributeListDeclarationInfo struct {
	NumberOfAttributes int64
	Attributes         *XMLAttributeDeclarationInfo
}

// Structure containing the callbacks for values for a CFBinaryHeap object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbinaryheapcallbacks?language=objc
type BinaryHeapCallBacks struct {
	Version         int64
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Compare         uintptr
}

// A reference to a binary heap object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbinaryheapref?language=objc
type BinaryHeapRef uintptr

// A structure that contains program-defined data and callbacks with which you can configure a version 1 CFRunLoopSource’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopsourcecontext1?language=objc
type RunLoopSourceContext1 struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
	Equal           uintptr
	Hash            uintptr
	GetPort         uintptr
	Perform         uintptr
}

// A reference to a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmachportref?language=objc
type MachPortRef uintptr

// A reference to a CFError object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cferrorref?language=objc
type ErrorRef uintptr

// A reference to a CFUUID object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfuuidref?language=objc
type UUIDRef uintptr

// A reference to a readable stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfreadstreamref?language=objc
type ReadStreamRef uintptr

// A structure that contains program-defined data and callbacks with which you can configure a CFSocket object’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketcontext?language=objc
type SocketContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A reference to an immutable bit vector object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbitvectorref?language=objc
type BitVectorRef uintptr

// A structure that contains program-defined data and callbacks with which you can configure a stream’s client behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstreamclientcontext?language=objc
type StreamClientContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A structure that contains program-defined data and callbacks with which you can configure a CFMessagePort object’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmessageportcontext?language=objc
type MessagePortContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// Contains a list of element attributes packaged as CFDictionary key/value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlelementinfo?language=objc
type XMLElementInfo struct {
	Attributes     uintptr
	AttributeOrder uintptr
	IsEmpty        uint8
	X_reserved     [3]int8
	Pad_cgo_0      [4]byte
}

// A reference to a run loop source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopsourceref?language=objc
type RunLoopSourceRef uintptr

// A reference to a run loop observer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopobserverref?language=objc
type RunLoopObserverRef uintptr

// A reference to a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbundleref?language=objc
type BundleRef uintptr

// A reference to an immutable array object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarrayref?language=objc
type ArrayRef uintptr

// A structure representing a range of sequential items in a container, such as characters in a buffer or elements in a collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrange?language=objc
type Range struct {
	Location int64
	Length   int64
}

// A reference to a CFTimeZone object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftimezoneref?language=objc
type TimeZoneRef uintptr
