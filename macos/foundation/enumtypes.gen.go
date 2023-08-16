// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

// Option flags used with beginActivityWithOptions:reason: and performActivityWithOptions:reason:usingBlock:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsactivityoptions?language=objc
type ActivityOptions uint64

const (
	ActivityAutomaticTerminationDisabled         ActivityOptions = 32768
	ActivityBackground                           ActivityOptions = 255
	ActivityIdleDisplaySleepDisabled             ActivityOptions = 1099511627776
	ActivityIdleSystemSleepDisabled              ActivityOptions = 1048576
	ActivityLatencyCritical                      ActivityOptions = 1095216660480
	ActivitySuddenTerminationDisabled            ActivityOptions = 16384
	ActivityUserInitiated                        ActivityOptions = 16777215
	ActivityUserInitiatedAllowingIdleSystemSleep ActivityOptions = 15728639
)

// Values representing alignment operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsalignmentoptions?language=objc
type AlignmentOptions int64

const (
	AlignAllEdgesInward  AlignmentOptions = 15
	AlignAllEdgesNearest AlignmentOptions = 983040
	AlignAllEdgesOutward AlignmentOptions = 3840
	AlignHeightInward    AlignmentOptions = 32
	AlignHeightNearest   AlignmentOptions = 2097152
	AlignHeightOutward   AlignmentOptions = 8192
	AlignMaxXInward      AlignmentOptions = 4
	AlignMaxXNearest     AlignmentOptions = 262144
	AlignMaxXOutward     AlignmentOptions = 1024
	AlignMaxYInward      AlignmentOptions = 8
	AlignMaxYNearest     AlignmentOptions = 524288
	AlignMaxYOutward     AlignmentOptions = 2048
	AlignMinXInward      AlignmentOptions = 1
	AlignMinXNearest     AlignmentOptions = 65536
	AlignMinXOutward     AlignmentOptions = 256
	AlignMinYInward      AlignmentOptions = 2
	AlignMinYNearest     AlignmentOptions = 131072
	AlignMinYOutward     AlignmentOptions = 512
	AlignRectFlipped     AlignmentOptions = -9223372036854775808
	AlignWidthInward     AlignmentOptions = 16
	AlignWidthNearest    AlignmentOptions = 1048576
	AlignWidthOutward    AlignmentOptions = 4096
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventsendoptions?language=objc
type AppleEventSendOptions uint

const (
	AppleEventSendAlwaysInteract AppleEventSendOptions = 48
	AppleEventSendCanInteract    AppleEventSendOptions = 32
	AppleEventSendCanSwitchLayer AppleEventSendOptions = 64
	AppleEventSendDefaultOptions AppleEventSendOptions = 35
	AppleEventSendDontAnnotate   AppleEventSendOptions = 65536
	AppleEventSendDontExecute    AppleEventSendOptions = 8192
	AppleEventSendDontRecord     AppleEventSendOptions = 4096
	AppleEventSendNeverInteract  AppleEventSendOptions = 16
	AppleEventSendNoReply        AppleEventSendOptions = 1
	AppleEventSendQueueReply     AppleEventSendOptions = 2
	AppleEventSendWaitForReply   AppleEventSendOptions = 3
)

// Options for enumerating attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringenumerationoptions?language=objc
type AttributedStringEnumerationOptions uint

const (
	AttributedStringEnumerationLongestEffectiveRangeNotRequired AttributedStringEnumerationOptions = 1048576
	AttributedStringEnumerationReverse                          AttributedStringEnumerationOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringformattingoptions?language=objc
type AttributedStringFormattingOptions uint

const (
	AttributedStringFormattingApplyReplacementIndexAttribute         AttributedStringFormattingOptions = 2
	AttributedStringFormattingInsertArgumentAttributesWithoutMerging AttributedStringFormattingOptions = 1
)

// Attributes that you can apply to text in an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringkey?language=objc
type AttributedStringKey string

const (
	AlternateDescriptionAttributeName     AttributedStringKey = "NSAlternateDescription"
	ImageURLAttributeName                 AttributedStringKey = "NSImageURL"
	InflectionAlternativeAttributeName    AttributedStringKey = "NSInflectionAlternative"
	InflectionRuleAttributeName           AttributedStringKey = "NSInflect"
	InlinePresentationIntentAttributeName AttributedStringKey = "NSInlinePresentationIntent"
	LanguageIdentifierAttributeName       AttributedStringKey = "NSLanguage"
	MorphologyAttributeName               AttributedStringKey = "NSMorphology"
	PresentationIntentAttributeName       AttributedStringKey = "NSPresentationIntent"
	ReplacementIndexAttributeName         AttributedStringKey = "NSReplacementIndex"
)

// A type that represents the syntax for intepreting a Markdown string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdowninterpretedsyntax?language=objc
type AttributedStringMarkdownInterpretedSyntax int

const (
	AttributedStringMarkdownInterpretedSyntaxFull                           AttributedStringMarkdownInterpretedSyntax = 0
	AttributedStringMarkdownInterpretedSyntaxInlineOnly                     AttributedStringMarkdownInterpretedSyntax = 1
	AttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace AttributedStringMarkdownInterpretedSyntax = 2
)

// A type that represents policies for handling parsing failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingfailurepolicy?language=objc
type AttributedStringMarkdownParsingFailurePolicy int

const (
	AttributedStringMarkdownParsingFailureReturnError                     AttributedStringMarkdownParsingFailurePolicy = 0
	AttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible AttributedStringMarkdownParsingFailurePolicy = 1
)

// These constants indicate whether background activity has been completed successfully or whether additional processing should be deferred until a more optimal time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityresult?language=objc
type BackgroundActivityResult int

const (
	BackgroundActivityResultDeferred BackgroundActivityResult = 2
	BackgroundActivityResultFinished BackgroundActivityResult = 1
)

// Options for searches and insertions using indexOfObject:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbinarysearchingoptions?language=objc
type BinarySearchingOptions uint

const (
	BinarySearchingFirstEqual     BinarySearchingOptions = 256
	BinarySearchingInsertionIndex BinarySearchingOptions = 1024
	BinarySearchingLastEqual      BinarySearchingOptions = 512
)

// Specifies display of file or storage byte counts. The display style is platform specific. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformattercountstyle?language=objc
type ByteCountFormatterCountStyle int

const (
	ByteCountFormatterCountStyleBinary  ByteCountFormatterCountStyle = 3
	ByteCountFormatterCountStyleDecimal ByteCountFormatterCountStyle = 2
	ByteCountFormatterCountStyleFile    ByteCountFormatterCountStyle = 0
	ByteCountFormatterCountStyleMemory  ByteCountFormatterCountStyle = 1
)

// Specifies the units appropriate for the formatter to display. Specifying any units explicitly causes just those units to be used in showing the number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatterunits?language=objc
type ByteCountFormatterUnits uint

const (
	ByteCountFormatterUseAll        ByteCountFormatterUnits = 65535
	ByteCountFormatterUseBytes      ByteCountFormatterUnits = 1
	ByteCountFormatterUseDefault    ByteCountFormatterUnits = 0
	ByteCountFormatterUseEB         ByteCountFormatterUnits = 64
	ByteCountFormatterUseGB         ByteCountFormatterUnits = 8
	ByteCountFormatterUseKB         ByteCountFormatterUnits = 2
	ByteCountFormatterUseMB         ByteCountFormatterUnits = 4
	ByteCountFormatterUsePB         ByteCountFormatterUnits = 32
	ByteCountFormatterUseTB         ByteCountFormatterUnits = 16
	ByteCountFormatterUseYBOrHigher ByteCountFormatterUnits = 65280
	ByteCountFormatterUseZB         ByteCountFormatterUnits = 128
)

// Calculation error constants used to describe an error in exceptionDuringOperation:error:leftOperand:rightOperand:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalculationerror?language=objc
type CalculationError uint

const (
	CalculationDivideByZero    CalculationError = 4
	CalculationLossOfPrecision CalculationError = 1
	CalculationNoError         CalculationError = 0
	CalculationOverflow        CalculationError = 3
	CalculationUnderflow       CalculationError = 2
)

// The supported calendar types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendaridentifier?language=objc
type CalendarIdentifier string

const (
	CalendarIdentifierBuddhist            CalendarIdentifier = "buddhist"
	CalendarIdentifierChinese             CalendarIdentifier = "chinese"
	CalendarIdentifierCoptic              CalendarIdentifier = "coptic"
	CalendarIdentifierEthiopicAmeteAlem   CalendarIdentifier = "ethiopic-amete-alem"
	CalendarIdentifierEthiopicAmeteMihret CalendarIdentifier = "ethiopic"
	CalendarIdentifierGregorian           CalendarIdentifier = "gregorian"
	CalendarIdentifierHebrew              CalendarIdentifier = "hebrew"
	CalendarIdentifierISO8601             CalendarIdentifier = "iso8601"
	CalendarIdentifierIndian              CalendarIdentifier = "indian"
	CalendarIdentifierIslamic             CalendarIdentifier = "islamic"
	CalendarIdentifierIslamicCivil        CalendarIdentifier = "islamic-civil"
	CalendarIdentifierIslamicTabular      CalendarIdentifier = "islamic-tbla"
	CalendarIdentifierIslamicUmmAlQura    CalendarIdentifier = "islamic-umalqura"
	CalendarIdentifierJapanese            CalendarIdentifier = "japanese"
	CalendarIdentifierPersian             CalendarIdentifier = "persian"
	CalendarIdentifierRepublicOfChina     CalendarIdentifier = "roc"
)

// The options for arithmetic operations involving calendars. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendaroptions?language=objc
type CalendarOptions uint

const (
	CalendarMatchFirst                              CalendarOptions = 4096
	CalendarMatchLast                               CalendarOptions = 8192
	CalendarMatchNextTime                           CalendarOptions = 1024
	CalendarMatchNextTimePreservingSmallerUnits     CalendarOptions = 512
	CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 256
	CalendarMatchStrictly                           CalendarOptions = 2
	CalendarSearchBackwards                         CalendarOptions = 4
	CalendarWrapComponents                          CalendarOptions = 1
)

// Calendrical units such as year, month, day and hour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendarunit?language=objc
type CalendarUnit uint

const (
	CalendarCalendarUnit          CalendarUnit = 1048576
	CalendarUnitCalendar          CalendarUnit = 1048576
	CalendarUnitDay               CalendarUnit = 16
	CalendarUnitEra               CalendarUnit = 2
	CalendarUnitHour              CalendarUnit = 32
	CalendarUnitMinute            CalendarUnit = 64
	CalendarUnitMonth             CalendarUnit = 8
	CalendarUnitNanosecond        CalendarUnit = 32768
	CalendarUnitQuarter           CalendarUnit = 2048
	CalendarUnitSecond            CalendarUnit = 128
	CalendarUnitTimeZone          CalendarUnit = 2097152
	CalendarUnitWeekOfMonth       CalendarUnit = 4096
	CalendarUnitWeekOfYear        CalendarUnit = 8192
	CalendarUnitWeekday           CalendarUnit = 512
	CalendarUnitWeekdayOrdinal    CalendarUnit = 1024
	CalendarUnitYear              CalendarUnit = 4
	CalendarUnitYearForWeekOfYear CalendarUnit = 16384
	DayCalendarUnit               CalendarUnit = 16
	EraCalendarUnit               CalendarUnit = 2
	HourCalendarUnit              CalendarUnit = 32
	MinuteCalendarUnit            CalendarUnit = 64
	MonthCalendarUnit             CalendarUnit = 8
	QuarterCalendarUnit           CalendarUnit = 2048
	SecondCalendarUnit            CalendarUnit = 128
	TimeZoneCalendarUnit          CalendarUnit = 2097152
	WeekCalendarUnit              CalendarUnit = 256
	WeekOfMonthCalendarUnit       CalendarUnit = 4096
	WeekOfYearCalendarUnit        CalendarUnit = 8192
	WeekdayCalendarUnit           CalendarUnit = 512
	WeekdayOrdinalCalendarUnit    CalendarUnit = 1024
	YearCalendarUnit              CalendarUnit = 4
	YearForWeekOfYearCalendarUnit CalendarUnit = 16384
)

// The type of change represented in computing the difference of an ordered collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscollectionchangetype?language=objc
type CollectionChangeType int

const (
	CollectionChangeInsert CollectionChangeType = 0
	CollectionChangeRemove CollectionChangeType = 1
)

// Constants that describe the possible types of modifier for a comparison predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicatemodifier?language=objc
type ComparisonPredicateModifier uint

const (
	AllPredicateModifier    ComparisonPredicateModifier = 1
	AnyPredicateModifier    ComparisonPredicateModifier = 2
	DirectPredicateModifier ComparisonPredicateModifier = 0
)

// Constants that describe the possible types of string comparison for comparison predicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicateoptions?language=objc
type ComparisonPredicateOptions uint

const (
	CaseInsensitivePredicateOption      ComparisonPredicateOptions = 1
	DiacriticInsensitivePredicateOption ComparisonPredicateOptions = 2
	NormalizedPredicateOption           ComparisonPredicateOptions = 4
)

// Constants that indicate sort order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonresult?language=objc
type ComparisonResult int

const (
	OrderedAscending  ComparisonResult = -1
	OrderedDescending ComparisonResult = 1
	OrderedSame       ComparisonResult = 0
)

// Constants that describe the possible types of a compound predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicatetype?language=objc
type CompoundPredicateType uint

const (
	AndPredicateType CompoundPredicateType = 1
	NotPredicateType CompoundPredicateType = 0
	OrPredicateType  CompoundPredicateType = 2
)

// Options to modify the decoding algorithm used to decode Base64 encoded data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatabase64decodingoptions?language=objc
type DataBase64DecodingOptions uint

const (
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1
)

// Options for methods used to Base64 encode data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatabase64encodingoptions?language=objc
type DataBase64EncodingOptions uint

const (
	DataBase64Encoding64CharacterLineLength     DataBase64EncodingOptions = 1
	DataBase64Encoding76CharacterLineLength     DataBase64EncodingOptions = 2
	DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 16
	DataBase64EncodingEndLineWithLineFeed       DataBase64EncodingOptions = 32
)

// An algorithm that indicates how to compress or decompress data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatacompressionalgorithm?language=objc
type DataCompressionAlgorithm int

const (
	DataCompressionAlgorithmLZ4   DataCompressionAlgorithm = 1
	DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
	DataCompressionAlgorithmLZMA  DataCompressionAlgorithm = 2
	DataCompressionAlgorithmZlib  DataCompressionAlgorithm = 3
)

// Options for methods used to read data objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatareadingoptions?language=objc
type DataReadingOptions uint

const (
	DataReadingMapped       DataReadingOptions = 1
	DataReadingMappedAlways DataReadingOptions = 8
	DataReadingMappedIfSafe DataReadingOptions = 1
	DataReadingUncached     DataReadingOptions = 2
	MappedRead              DataReadingOptions = 1
	UncachedRead            DataReadingOptions = 2
)

// Options for method used to search data objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatasearchoptions?language=objc
type DataSearchOptions uint

const (
	DataSearchAnchored  DataSearchOptions = 2
	DataSearchBackwards DataSearchOptions = 1
)

// Options for methods used to write data objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatawritingoptions?language=objc
type DataWritingOptions uint

const (
	AtomicWrite                                                   DataWritingOptions = 1
	DataWritingAtomic                                             DataWritingOptions = 1
	DataWritingFileProtectionComplete                             DataWritingOptions = 536870912
	DataWritingFileProtectionCompleteUnlessOpen                   DataWritingOptions = 805306368
	DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 1073741824
	DataWritingFileProtectionMask                                 DataWritingOptions = 4026531840
	DataWritingFileProtectionNone                                 DataWritingOptions = 268435456
	DataWritingWithoutOverwriting                                 DataWritingOptions = 2
)

// Constants for specifying how to represent quantities of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatterunitsstyle?language=objc
type DateComponentsFormatterUnitsStyle int

const (
	DateComponentsFormatterUnitsStyleAbbreviated DateComponentsFormatterUnitsStyle = 1
	DateComponentsFormatterUnitsStyleBrief       DateComponentsFormatterUnitsStyle = 5
	DateComponentsFormatterUnitsStyleFull        DateComponentsFormatterUnitsStyle = 3
	DateComponentsFormatterUnitsStylePositional  DateComponentsFormatterUnitsStyle = 0
	DateComponentsFormatterUnitsStyleShort       DateComponentsFormatterUnitsStyle = 2
	DateComponentsFormatterUnitsStyleSpellOut    DateComponentsFormatterUnitsStyle = 4
)

// Formatting constants for when values contain zeroes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatterzeroformattingbehavior?language=objc
type DateComponentsFormatterZeroFormattingBehavior uint

const (
	DateComponentsFormatterZeroFormattingBehaviorDefault      DateComponentsFormatterZeroFormattingBehavior = 1
	DateComponentsFormatterZeroFormattingBehaviorDropAll      DateComponentsFormatterZeroFormattingBehavior = 14
	DateComponentsFormatterZeroFormattingBehaviorDropLeading  DateComponentsFormatterZeroFormattingBehavior = 2
	DateComponentsFormatterZeroFormattingBehaviorDropMiddle   DateComponentsFormatterZeroFormattingBehavior = 4
	DateComponentsFormatterZeroFormattingBehaviorDropTrailing DateComponentsFormatterZeroFormattingBehavior = 8
	DateComponentsFormatterZeroFormattingBehaviorNone         DateComponentsFormatterZeroFormattingBehavior = 0
	DateComponentsFormatterZeroFormattingBehaviorPad          DateComponentsFormatterZeroFormattingBehavior = 65536
)

// Constants that specify the behavior NSDateFormatter should exhibit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatterbehavior?language=objc
type DateFormatterBehavior uint

const (
	DateFormatterBehavior10_0    DateFormatterBehavior = 1000
	DateFormatterBehavior10_4    DateFormatterBehavior = 1040
	DateFormatterBehaviorDefault DateFormatterBehavior = 0
)

// The following constants specify predefined format styles for dates and times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatterstyle?language=objc
type DateFormatterStyle uint

const (
	DateFormatterFullStyle   DateFormatterStyle = 4
	DateFormatterLongStyle   DateFormatterStyle = 3
	DateFormatterMediumStyle DateFormatterStyle = 2
	DateFormatterNoStyle     DateFormatterStyle = 0
	DateFormatterShortStyle  DateFormatterStyle = 1
)

// Formatting styles for individual date and time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatterstyle?language=objc
type DateIntervalFormatterStyle uint

const (
	DateIntervalFormatterFullStyle   DateIntervalFormatterStyle = 4
	DateIntervalFormatterLongStyle   DateIntervalFormatterStyle = 3
	DateIntervalFormatterMediumStyle DateIntervalFormatterStyle = 2
	DateIntervalFormatterNoStyle     DateIntervalFormatterStyle = 0
	DateIntervalFormatterShortStyle  DateIntervalFormatterStyle = 1
)

// Policies describing the action the coder should take when encountering decode failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecodingfailurepolicy?language=objc
type DecodingFailurePolicy int

const (
	DecodingFailurePolicyRaiseException    DecodingFailurePolicy = 0
	DecodingFailurePolicySetErrorAndReturn DecodingFailurePolicy = 1
)

// Options for enumerating the contents of directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerationoptions?language=objc
type DirectoryEnumerationOptions uint

const (
	DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 8
	DirectoryEnumerationProducesRelativePathURLs     DirectoryEnumerationOptions = 16
	DirectoryEnumerationSkipsHiddenFiles             DirectoryEnumerationOptions = 4
	DirectoryEnumerationSkipsPackageDescendants      DirectoryEnumerationOptions = 2
	DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1
)

// This constant specifies the notification center type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcentertype?language=objc
type DistributedNotificationCenterType string

const (
	LocalNotificationCenterType DistributedNotificationCenterType = "_NSLocalNotificationCenter"
)

// These constants specify the behavior of notifications posted using the postNotificationName:object: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationoptions?language=objc
type DistributedNotificationOptions uint

const (
	DistributedNotificationDeliverImmediately DistributedNotificationOptions = 1
	DistributedNotificationPostToAllSessions  DistributedNotificationOptions = 2
	NotificationDeliverImmediately            DistributedNotificationOptions = 1
	NotificationPostToAllSessions             DistributedNotificationOptions = 2
)

// The units supported by the NSEnergyFormatter class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatterunit?language=objc
type EnergyFormatterUnit int

const (
	EnergyFormatterUnitCalorie     EnergyFormatterUnit = 1793
	EnergyFormatterUnitJoule       EnergyFormatterUnit = 11
	EnergyFormatterUnitKilocalorie EnergyFormatterUnit = 1794
	EnergyFormatterUnitKilojoule   EnergyFormatterUnit = 14
)

// Options for block enumeration operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenumerationoptions?language=objc
type EnumerationOptions uint

const (
	EnumerationConcurrent EnumerationOptions = 1
	EnumerationReverse    EnumerationOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserrordomain?language=objc
type ErrorDomain string

const (
	CocoaErrorDomain           ErrorDomain = "NSCocoaErrorDomain"
	MachErrorDomain            ErrorDomain = "NSMachErrorDomain"
	NetServicesErrorDomain     ErrorDomain = "NSNetServicesErrorDomain"
	OSStatusErrorDomain        ErrorDomain = "NSOSStatusErrorDomain"
	POSIXErrorDomain           ErrorDomain = "NSPOSIXErrorDomain"
	StreamSOCKSErrorDomain     ErrorDomain = "NSStreamSOCKSErrorDomain"
	StreamSocketSSLErrorDomain ErrorDomain = "NSStreamSocketSSLErrorDomain"
	URLErrorDomain             ErrorDomain = "NSURLErrorDomain"
	XMLParserErrorDomain       ErrorDomain = "NSXMLParserErrorDomain"
)

// These keys may exist in the user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserroruserinfokey?language=objc
type ErrorUserInfoKey string

const (
	DebugDescriptionErrorKey            ErrorUserInfoKey = "NSDebugDescription"
	FilePathErrorKey                    ErrorUserInfoKey = "NSFilePath"
	HelpAnchorErrorKey                  ErrorUserInfoKey = "NSHelpAnchor"
	LocalizedDescriptionKey             ErrorUserInfoKey = "NSLocalizedDescription"
	LocalizedFailureErrorKey            ErrorUserInfoKey = "NSLocalizedFailure"
	LocalizedFailureReasonErrorKey      ErrorUserInfoKey = "NSLocalizedFailureReason"
	LocalizedRecoveryOptionsErrorKey    ErrorUserInfoKey = "NSLocalizedRecoveryOptions"
	LocalizedRecoverySuggestionErrorKey ErrorUserInfoKey = "NSLocalizedRecoverySuggestion"
	MultipleUnderlyingErrorsKey         ErrorUserInfoKey = "NSMultipleUnderlyingErrorsKey"
	RecoveryAttempterErrorKey           ErrorUserInfoKey = "NSRecoveryAttempter"
	StringEncodingErrorKey              ErrorUserInfoKey = "NSStringEncoding"
	URLErrorKey                         ErrorUserInfoKey = "NSURL"
	URLErrorNetworkUnavailableReasonKey ErrorUserInfoKey = "NSURLErrorNetworkUnavailableReasonKey"
	UnderlyingErrorKey                  ErrorUserInfoKey = "NSUnderlyingError"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexceptionname?language=objc
type ExceptionName string

const (
	CharacterConversionException           ExceptionName = "NSCharacterConversionException"
	DecimalNumberDivideByZeroException     ExceptionName = "NSDecimalNumberDivideByZeroException"
	DecimalNumberExactnessException        ExceptionName = "NSDecimalNumberExactnessException"
	DecimalNumberOverflowException         ExceptionName = "NSDecimalNumberOverflowException"
	DecimalNumberUnderflowException        ExceptionName = "NSDecimalNumberUnderflowException"
	DestinationInvalidException            ExceptionName = "NSDestinationInvalidException"
	FileHandleOperationException           ExceptionName = "NSFileHandleOperationException"
	GenericException                       ExceptionName = "NSGenericException"
	InconsistentArchiveException           ExceptionName = "NSArchiverArchiveInconsistency"
	InternalInconsistencyException         ExceptionName = "NSInternalInconsistencyException"
	InvalidArchiveOperationException       ExceptionName = "NSInvalidArchiveOperationException"
	InvalidArgumentException               ExceptionName = "NSInvalidArgumentException"
	InvalidReceivePortException            ExceptionName = "NSInvalidReceivePortException"
	InvalidSendPortException               ExceptionName = "NSInvalidSendPortException"
	InvalidUnarchiveOperationException     ExceptionName = "NSInvalidUnarchiveOperationException"
	InvocationOperationCancelledException  ExceptionName = "NSInvocationOperationCancelledException"
	InvocationOperationVoidResultException ExceptionName = "NSInvocationOperationVoidResultException"
	MallocException                        ExceptionName = "NSMallocException"
	ObjectInaccessibleException            ExceptionName = "NSObjectInaccessibleException"
	ObjectNotAvailableException            ExceptionName = "NSObjectNotAvailableException"
	OldStyleException                      ExceptionName = "NSOldStyleException"
	ParseErrorException                    ExceptionName = "NSParseErrorException"
	PortReceiveException                   ExceptionName = "NSPortReceiveException"
	PortSendException                      ExceptionName = "NSPortSendException"
	PortTimeoutException                   ExceptionName = "NSPortTimeoutException"
	RangeException                         ExceptionName = "NSRangeException"
	UndefinedKeyException                  ExceptionName = "NSUnknownKeyException"
)

// Defines the possible types of an expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpressiontype?language=objc
type ExpressionType uint

const (
	AggregateExpressionType       ExpressionType = 14
	AnyKeyExpressionType          ExpressionType = 15
	BlockExpressionType           ExpressionType = 19
	ConditionalExpressionType     ExpressionType = 20
	ConstantValueExpressionType   ExpressionType = 0
	EvaluatedObjectExpressionType ExpressionType = 1
	FunctionExpressionType        ExpressionType = 4
	IntersectSetExpressionType    ExpressionType = 6
	KeyPathExpressionType         ExpressionType = 3
	MinusSetExpressionType        ExpressionType = 7
	SubqueryExpressionType        ExpressionType = 13
	UnionSetExpressionType        ExpressionType = 5
	VariableExpressionType        ExpressionType = 2
)

// Keys in dictionaries used to get and set file attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileattributekey?language=objc
type FileAttributeKey string

const (
	FileAppendOnly            FileAttributeKey = "NSFileAppendOnly"
	FileBusy                  FileAttributeKey = "NSFileBusy"
	FileCreationDate          FileAttributeKey = "NSFileCreationDate"
	FileDeviceIdentifier      FileAttributeKey = "NSFileDeviceIdentifier"
	FileExtensionHidden       FileAttributeKey = "NSFileExtensionHidden"
	FileGroupOwnerAccountID   FileAttributeKey = "NSFileGroupOwnerAccountID"
	FileGroupOwnerAccountName FileAttributeKey = "NSFileGroupOwnerAccountName"
	FileHFSCreatorCode        FileAttributeKey = "NSFileHFSCreatorCode"
	FileHFSTypeCode           FileAttributeKey = "NSFileHFSTypeCode"
	FileImmutable             FileAttributeKey = "NSFileImmutable"
	FileModificationDate      FileAttributeKey = "NSFileModificationDate"
	FileOwnerAccountID        FileAttributeKey = "NSFileOwnerAccountID"
	FileOwnerAccountName      FileAttributeKey = "NSFileOwnerAccountName"
	FilePosixPermissions      FileAttributeKey = "NSFilePosixPermissions"
	FileProtectionKey         FileAttributeKey = "NSFileProtectionKey"
	FileReferenceCount        FileAttributeKey = "NSFileReferenceCount"
	FileSize                  FileAttributeKey = "NSFileSize"
	FileSystemFileNumber      FileAttributeKey = "NSFileSystemFileNumber"
	FileSystemFreeNodes       FileAttributeKey = "NSFileSystemFreeNodes"
	FileSystemFreeSize        FileAttributeKey = "NSFileSystemFreeSize"
	FileSystemNodes           FileAttributeKey = "NSFileSystemNodes"
	FileSystemNumber          FileAttributeKey = "NSFileSystemNumber"
	FileSystemSize            FileAttributeKey = "NSFileSystemSize"
	FileType                  FileAttributeKey = "NSFileType"
)

// Values representing a file’s type attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileattributetype?language=objc
type FileAttributeType string

const (
	FileTypeBlockSpecial     FileAttributeType = "NSFileTypeBlockSpecial"
	FileTypeCharacterSpecial FileAttributeType = "NSFileTypeCharacterSpecial"
	FileTypeDirectory        FileAttributeType = "NSFileTypeDirectory"
	FileTypeRegular          FileAttributeType = "NSFileTypeRegular"
	FileTypeSocket           FileAttributeType = "NSFileTypeSocket"
	FileTypeSymbolicLink     FileAttributeType = "NSFileTypeSymbolicLink"
	FileTypeUnknown          FileAttributeType = "NSFileTypeUnknown"
)

// Options to use when reading the contents or attributes of a file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinatorreadingoptions?language=objc
type FileCoordinatorReadingOptions uint

const (
	FileCoordinatorReadingForUploading                     FileCoordinatorReadingOptions = 8
	FileCoordinatorReadingImmediatelyAvailableMetadataOnly FileCoordinatorReadingOptions = 4
	FileCoordinatorReadingResolvesSymbolicLink             FileCoordinatorReadingOptions = 2
	FileCoordinatorReadingWithoutChanges                   FileCoordinatorReadingOptions = 1
)

// Options to use when changing the contents or attributes of a file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinatorwritingoptions?language=objc
type FileCoordinatorWritingOptions uint

const (
	FileCoordinatorWritingContentIndependentMetadataOnly FileCoordinatorWritingOptions = 16
	FileCoordinatorWritingForDeleting                    FileCoordinatorWritingOptions = 1
	FileCoordinatorWritingForMerging                     FileCoordinatorWritingOptions = 4
	FileCoordinatorWritingForMoving                      FileCoordinatorWritingOptions = 2
	FileCoordinatorWritingForReplacing                   FileCoordinatorWritingOptions = 8
)

// Options for specifying the behavior of file replacement operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanageritemreplacementoptions?language=objc
type FileManagerItemReplacementOptions uint

const (
	FileManagerItemReplacementUsingNewMetadataOnly      FileManagerItemReplacementOptions = 1
	FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 2
)

// Options that specify the behavior of an unmount operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerunmountoptions?language=objc
type FileManagerUnmountOptions uint

const (
	FileManagerUnmountAllPartitionsAndEjectDisk FileManagerUnmountOptions = 1
	FileManagerUnmountWithoutUI                 FileManagerUnmountOptions = 2
)

// Protection level values that can be associated with a file attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileprotectiontype?language=objc
type FileProtectionType string

const (
	FileProtectionComplete                             FileProtectionType = "NSFileProtectionComplete"
	FileProtectionCompleteUnlessOpen                   FileProtectionType = "NSFileProtectionCompleteUnlessOpen"
	FileProtectionCompleteUntilFirstUserAuthentication FileProtectionType = "NSFileProtectionCompleteUntilFirstUserAuthentication"
	FileProtectionNone                                 FileProtectionType = "NSFileProtectionNone"
)

// The name used to identify a File Provider service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileproviderservicename?language=objc
type FileProviderServiceName string

// Options for adding a new file version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversionaddingoptions?language=objc
type FileVersionAddingOptions uint

const (
	FileVersionAddingByMoving FileVersionAddingOptions = 1
)

// Options for replacing a file version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversionreplacingoptions?language=objc
type FileVersionReplacingOptions uint

const (
	FileVersionReplacingByMoving FileVersionReplacingOptions = 1
)

// Reading options that can be set by the initWithURL:options:error: and readFromURL:options:error: methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapperreadingoptions?language=objc
type FileWrapperReadingOptions uint

const (
	FileWrapperReadingImmediate      FileWrapperReadingOptions = 1
	FileWrapperReadingWithoutMapping FileWrapperReadingOptions = 2
)

// Writing options that can be set by the writeToURL:options:originalContentsURL:error: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapperwritingoptions?language=objc
type FileWrapperWritingOptions uint

const (
	FileWrapperWritingAtomic           FileWrapperWritingOptions = 1
	FileWrapperWritingWithNameUpdating FileWrapperWritingOptions = 2
)

// The formatting context for a formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformattingcontext?language=objc
type FormattingContext int

const (
	FormattingContextBeginningOfSentence FormattingContext = 4
	FormattingContextDynamic             FormattingContext = 1
	FormattingContextListItem            FormattingContext = 3
	FormattingContextMiddleOfSentence    FormattingContext = 5
	FormattingContextStandalone          FormattingContext = 2
	FormattingContextUnknown             FormattingContext = 0
)

// Specifies the width of the unit, determining the textual representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformattingunitstyle?language=objc
type FormattingUnitStyle int

const (
	FormattingUnitStyleLong   FormattingUnitStyle = 3
	FormattingUnitStyleMedium FormattingUnitStyle = 2
	FormattingUnitStyleShort  FormattingUnitStyle = 1
)

// A representation of grammatical gender, used for inflecting strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsgrammaticalgender?language=objc
type GrammaticalGender int

const (
	GrammaticalGenderFeminine  GrammaticalGender = 1
	GrammaticalGenderMasculine GrammaticalGender = 2
	GrammaticalGenderNeuter    GrammaticalGender = 3
	GrammaticalGenderNotSet    GrammaticalGender = 0
)

// A representation of grammatical number, used for inflecting strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsgrammaticalnumber?language=objc
type GrammaticalNumber int

const (
	GrammaticalNumberNotSet     GrammaticalNumber = 0
	GrammaticalNumberPlural     GrammaticalNumber = 3
	GrammaticalNumberPluralFew  GrammaticalNumber = 5
	GrammaticalNumberPluralMany GrammaticalNumber = 6
	GrammaticalNumberPluralTwo  GrammaticalNumber = 4
	GrammaticalNumberSingular   GrammaticalNumber = 1
	GrammaticalNumberZero       GrammaticalNumber = 2
)

// A representation of grammatical parts of speech, used for inflecting strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsgrammaticalpartofspeech?language=objc
type GrammaticalPartOfSpeech int

const (
	GrammaticalPartOfSpeechAbbreviation GrammaticalPartOfSpeech = 14
	GrammaticalPartOfSpeechAdjective    GrammaticalPartOfSpeech = 6
	GrammaticalPartOfSpeechAdposition   GrammaticalPartOfSpeech = 7
	GrammaticalPartOfSpeechAdverb       GrammaticalPartOfSpeech = 4
	GrammaticalPartOfSpeechConjunction  GrammaticalPartOfSpeech = 10
	GrammaticalPartOfSpeechDeterminer   GrammaticalPartOfSpeech = 1
	GrammaticalPartOfSpeechInterjection GrammaticalPartOfSpeech = 12
	GrammaticalPartOfSpeechLetter       GrammaticalPartOfSpeech = 3
	GrammaticalPartOfSpeechNotSet       GrammaticalPartOfSpeech = 0
	GrammaticalPartOfSpeechNoun         GrammaticalPartOfSpeech = 9
	GrammaticalPartOfSpeechNumeral      GrammaticalPartOfSpeech = 11
	GrammaticalPartOfSpeechParticle     GrammaticalPartOfSpeech = 5
	GrammaticalPartOfSpeechPreposition  GrammaticalPartOfSpeech = 13
	GrammaticalPartOfSpeechPronoun      GrammaticalPartOfSpeech = 2
	GrammaticalPartOfSpeechVerb         GrammaticalPartOfSpeech = 8
)

// Cookie acceptance policies implemented by the NSHTTPCookieStorage class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookieacceptpolicy?language=objc
type HTTPCookieAcceptPolicy uint

const (
	HTTPCookieAcceptPolicyAlways                     HTTPCookieAcceptPolicy = 0
	HTTPCookieAcceptPolicyNever                      HTTPCookieAcceptPolicy = 1
	HTTPCookieAcceptPolicyOnlyFromMainDocumentDomain HTTPCookieAcceptPolicy = 2
)

// Constants that define the supported keys in a cookie attributes dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiepropertykey?language=objc
type HTTPCookiePropertyKey string

const (
	HTTPCookieComment        HTTPCookiePropertyKey = "Comment"
	HTTPCookieCommentURL     HTTPCookiePropertyKey = "CommentURL"
	HTTPCookieDiscard        HTTPCookiePropertyKey = "Discard"
	HTTPCookieDomain         HTTPCookiePropertyKey = "Domain"
	HTTPCookieExpires        HTTPCookiePropertyKey = "Expires"
	HTTPCookieMaximumAge     HTTPCookiePropertyKey = "Max-Age"
	HTTPCookieName           HTTPCookiePropertyKey = "Name"
	HTTPCookieOriginURL      HTTPCookiePropertyKey = "OriginURL"
	HTTPCookiePath           HTTPCookiePropertyKey = "Path"
	HTTPCookiePort           HTTPCookiePropertyKey = "Port"
	HTTPCookieSameSitePolicy HTTPCookiePropertyKey = "SameSite"
	HTTPCookieSecure         HTTPCookiePropertyKey = "Secure"
	HTTPCookieValue          HTTPCookiePropertyKey = "Value"
	HTTPCookieVersion        HTTPCookiePropertyKey = "Version"
)

// Values that indicate whether to restrict the cookie to requests sent back to the same site that created it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestringpolicy?language=objc
type HTTPCookieStringPolicy string

const (
	HTTPCookieSameSiteLax    HTTPCookieStringPolicy = "lax"
	HTTPCookieSameSiteStrict HTTPCookieStringPolicy = "strict"
)

// Components in a bit-field to specify the behavior of elements in an NSHashTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtableoptions?language=objc
type HashTableOptions uint

// Options used to generate and parse ISO 8601 date representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatoptions?language=objc
type ISO8601DateFormatOptions uint

const (
	ISO8601DateFormatWithColonSeparatorInTime     ISO8601DateFormatOptions = 512
	ISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 1024
	ISO8601DateFormatWithDashSeparatorInDate      ISO8601DateFormatOptions = 256
	ISO8601DateFormatWithDay                      ISO8601DateFormatOptions = 16
	ISO8601DateFormatWithFractionalSeconds        ISO8601DateFormatOptions = 2048
	ISO8601DateFormatWithFullDate                 ISO8601DateFormatOptions = 275
	ISO8601DateFormatWithFullTime                 ISO8601DateFormatOptions = 1632
	ISO8601DateFormatWithInternetDateTime         ISO8601DateFormatOptions = 1907
	ISO8601DateFormatWithMonth                    ISO8601DateFormatOptions = 2
	ISO8601DateFormatWithSpaceBetweenDateAndTime  ISO8601DateFormatOptions = 128
	ISO8601DateFormatWithTime                     ISO8601DateFormatOptions = 32
	ISO8601DateFormatWithTimeZone                 ISO8601DateFormatOptions = 64
	ISO8601DateFormatWithWeekOfYear               ISO8601DateFormatOptions = 4
	ISO8601DateFormatWithYear                     ISO8601DateFormatOptions = 1
)

// A type that defines presentation intent for runs of characters for traits like emphasis, strikethrough, and code voice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinlinepresentationintent?language=objc
type InlinePresentationIntent uint

const (
	InlinePresentationIntentBlockHTML          InlinePresentationIntent = 512
	InlinePresentationIntentCode               InlinePresentationIntent = 4
	InlinePresentationIntentEmphasized         InlinePresentationIntent = 1
	InlinePresentationIntentInlineHTML         InlinePresentationIntent = 256
	InlinePresentationIntentLineBreak          InlinePresentationIntent = 128
	InlinePresentationIntentSoftBreak          InlinePresentationIntent = 64
	InlinePresentationIntentStrikethrough      InlinePresentationIntent = 32
	InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2
)

// The following constants are defined by NSPositionalSpecifier to specify an insertion position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinsertionposition?language=objc
type InsertionPosition uint

const (
	PositionAfter     InsertionPosition = 0
	PositionBefore    InsertionPosition = 1
	PositionBeginning InsertionPosition = 2
	PositionEnd       InsertionPosition = 3
	PositionReplace   InsertionPosition = 4
)

// The error codes that describe problems with consuming data from an item provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovidererrorcode?language=objc
type ItemProviderErrorCode int

const (
	ItemProviderItemUnavailableError      ItemProviderErrorCode = -1000
	ItemProviderUnavailableCoercionError  ItemProviderErrorCode = -1200
	ItemProviderUnexpectedValueClassError ItemProviderErrorCode = -1100
	ItemProviderUnknownError              ItemProviderErrorCode = -1
)

// Data-access specifications that declare how to handle items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderfileoptions?language=objc
type ItemProviderFileOptions int

const (
	ItemProviderFileOptionOpenInPlace ItemProviderFileOptions = 1
)

// Specifications that control which categories of processes can see an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderrepresentationvisibility?language=objc
type ItemProviderRepresentationVisibility int

const (
	ItemProviderRepresentationVisibilityAll        ItemProviderRepresentationVisibility = 0
	ItemProviderRepresentationVisibilityGroup      ItemProviderRepresentationVisibility = 2
	ItemProviderRepresentationVisibilityOwnProcess ItemProviderRepresentationVisibility = 3
)

// Options used when creating Foundation objects from JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonreadingoptions?language=objc
type JSONReadingOptions uint

const (
	JSONReadingAllowFragments            JSONReadingOptions = 4
	JSONReadingFragmentsAllowed          JSONReadingOptions = 4
	JSONReadingJSON5Allowed              JSONReadingOptions = 8
	JSONReadingMutableContainers         JSONReadingOptions = 1
	JSONReadingMutableLeaves             JSONReadingOptions = 2
	JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 16
)

// Options for writing JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonwritingoptions?language=objc
type JSONWritingOptions uint

const (
	JSONWritingFragmentsAllowed       JSONWritingOptions = 4
	JSONWritingPrettyPrinted          JSONWritingOptions = 1
	JSONWritingSortedKeys             JSONWritingOptions = 2
	JSONWritingWithoutEscapingSlashes JSONWritingOptions = 8
)

// The kinds of changes that can be observed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyvaluechange?language=objc
type KeyValueChange uint

const (
	KeyValueChangeInsertion   KeyValueChange = 2
	KeyValueChangeRemoval     KeyValueChange = 3
	KeyValueChangeReplacement KeyValueChange = 4
	KeyValueChangeSetting     KeyValueChange = 1
)

// The keys that can appear in the change dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyvaluechangekey?language=objc
type KeyValueChangeKey string

const (
	KeyValueChangeIndexesKey             KeyValueChangeKey = "indexes"
	KeyValueChangeKindKey                KeyValueChangeKey = "kind"
	KeyValueChangeNewKey                 KeyValueChangeKey = "new"
	KeyValueChangeNotificationIsPriorKey KeyValueChangeKey = "notificationIsPrior"
	KeyValueChangeOldKey                 KeyValueChangeKey = "old"
)

// The values that can be returned in a change dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyvalueobservingoptions?language=objc
type KeyValueObservingOptions uint

const (
	KeyValueObservingOptionInitial KeyValueObservingOptions = 4
	KeyValueObservingOptionNew     KeyValueObservingOptions = 1
	KeyValueObservingOptionOld     KeyValueObservingOptions = 2
	KeyValueObservingOptionPrior   KeyValueObservingOptions = 8
)

// These constants define the available array operators. See [devLink-1699119] for more information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyvalueoperator?language=objc
type KeyValueOperator string

const (
	AverageKeyValueOperator                KeyValueOperator = "avg"
	CountKeyValueOperator                  KeyValueOperator = "count"
	DistinctUnionOfArraysKeyValueOperator  KeyValueOperator = "distinctUnionOfArrays"
	DistinctUnionOfObjectsKeyValueOperator KeyValueOperator = "distinctUnionOfObjects"
	DistinctUnionOfSetsKeyValueOperator    KeyValueOperator = "distinctUnionOfSets"
	MaximumKeyValueOperator                KeyValueOperator = "max"
	MinimumKeyValueOperator                KeyValueOperator = "min"
	SumKeyValueOperator                    KeyValueOperator = "sum"
	UnionOfArraysKeyValueOperator          KeyValueOperator = "unionOfArrays"
	UnionOfObjectsKeyValueOperator         KeyValueOperator = "unionOfObjects"
	UnionOfSetsKeyValueOperator            KeyValueOperator = "unionOfSets"
)

// The kinds of mutation that you can make to an unordered collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyvaluesetmutationkind?language=objc
type KeyValueSetMutationKind uint

const (
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
	KeyValueMinusSetMutation     KeyValueSetMutationKind = 2
	KeyValueSetSetMutation       KeyValueSetMutationKind = 4
	KeyValueUnionSetMutation     KeyValueSetMutationKind = 1
)

// The units supported by the NSLengthFormatter class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatterunit?language=objc
type LengthFormatterUnit int

const (
	LengthFormatterUnitCentimeter LengthFormatterUnit = 9
	LengthFormatterUnitFoot       LengthFormatterUnit = 1282
	LengthFormatterUnitInch       LengthFormatterUnit = 1281
	LengthFormatterUnitKilometer  LengthFormatterUnit = 14
	LengthFormatterUnitMeter      LengthFormatterUnit = 11
	LengthFormatterUnitMile       LengthFormatterUnit = 1284
	LengthFormatterUnitMillimeter LengthFormatterUnit = 8
	LengthFormatterUnitYard       LengthFormatterUnit = 1283
)

// A token, lexical class, name, lemma, language, or script returned by a linguistic tagger for natural language text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictag?language=objc
type LinguisticTag string

const (
	LinguisticTagAdjective          LinguisticTag = "Adjective"
	LinguisticTagAdverb             LinguisticTag = "Adverb"
	LinguisticTagClassifier         LinguisticTag = "Classifier"
	LinguisticTagCloseParenthesis   LinguisticTag = "CloseParenthesis"
	LinguisticTagCloseQuote         LinguisticTag = "CloseQuote"
	LinguisticTagConjunction        LinguisticTag = "Conjunction"
	LinguisticTagDash               LinguisticTag = "Dash"
	LinguisticTagDeterminer         LinguisticTag = "Determiner"
	LinguisticTagIdiom              LinguisticTag = "Idiom"
	LinguisticTagInterjection       LinguisticTag = "Interjection"
	LinguisticTagNoun               LinguisticTag = "Noun"
	LinguisticTagNumber             LinguisticTag = "Number"
	LinguisticTagOpenParenthesis    LinguisticTag = "OpenParenthesis"
	LinguisticTagOpenQuote          LinguisticTag = "OpenQuote"
	LinguisticTagOrganizationName   LinguisticTag = "OrganizationName"
	LinguisticTagOther              LinguisticTag = "Other"
	LinguisticTagOtherPunctuation   LinguisticTag = "Punctuation"
	LinguisticTagOtherWhitespace    LinguisticTag = "Whitespace"
	LinguisticTagOtherWord          LinguisticTag = "OtherWord"
	LinguisticTagParagraphBreak     LinguisticTag = "ParagraphBreak"
	LinguisticTagParticle           LinguisticTag = "Particle"
	LinguisticTagPersonalName       LinguisticTag = "PersonalName"
	LinguisticTagPlaceName          LinguisticTag = "PlaceName"
	LinguisticTagPreposition        LinguisticTag = "Preposition"
	LinguisticTagPronoun            LinguisticTag = "Pronoun"
	LinguisticTagPunctuation        LinguisticTag = "Punctuation"
	LinguisticTagSentenceTerminator LinguisticTag = "SentenceTerminator"
	LinguisticTagVerb               LinguisticTag = "Verb"
	LinguisticTagWhitespace         LinguisticTag = "Whitespace"
	LinguisticTagWord               LinguisticTag = "Word"
	LinguisticTagWordJoiner         LinguisticTag = "WordJoiner"
)

// Constants for the tag schemes specified when initializing a linguistic tagger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagscheme?language=objc
type LinguisticTagScheme string

const (
	LinguisticTagSchemeLanguage               LinguisticTagScheme = "Language"
	LinguisticTagSchemeLemma                  LinguisticTagScheme = "Lemma"
	LinguisticTagSchemeLexicalClass           LinguisticTagScheme = "LexicalClass"
	LinguisticTagSchemeNameType               LinguisticTagScheme = "NameType"
	LinguisticTagSchemeNameTypeOrLexicalClass LinguisticTagScheme = "NameTypeOrLexicalClass"
	LinguisticTagSchemeScript                 LinguisticTagScheme = "Script"
	LinguisticTagSchemeTokenType              LinguisticTagScheme = "TokenType"
)

// Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictaggeroptions?language=objc
type LinguisticTaggerOptions uint

const (
	LinguisticTaggerJoinNames       LinguisticTaggerOptions = 16
	LinguisticTaggerOmitOther       LinguisticTaggerOptions = 8
	LinguisticTaggerOmitPunctuation LinguisticTaggerOptions = 2
	LinguisticTaggerOmitWhitespace  LinguisticTaggerOptions = 4
	LinguisticTaggerOmitWords       LinguisticTaggerOptions = 1
)

// Constants representing linguistic units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictaggerunit?language=objc
type LinguisticTaggerUnit int

const (
	LinguisticTaggerUnitDocument  LinguisticTaggerUnit = 3
	LinguisticTaggerUnitParagraph LinguisticTaggerUnit = 2
	LinguisticTaggerUnitSentence  LinguisticTaggerUnit = 1
	LinguisticTaggerUnitWord      LinguisticTaggerUnit = 0
)

// The keys used to access components of a locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocalekey?language=objc
type LocaleKey string

const (
	LocaleAlternateQuotationBeginDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationBeginDelimiterKey"
	LocaleAlternateQuotationEndDelimiterKey   LocaleKey = "kCFLocaleAlternateQuotationEndDelimiterKey"
	LocaleCalendar                            LocaleKey = "kCFLocaleCalendarKey"
	LocaleCollationIdentifier                 LocaleKey = "collation"
	LocaleCollatorIdentifier                  LocaleKey = "kCFLocaleCollatorIdentifierKey"
	LocaleCountryCode                         LocaleKey = "kCFLocaleCountryCodeKey"
	LocaleCurrencyCode                        LocaleKey = "currency"
	LocaleCurrencySymbol                      LocaleKey = "kCFLocaleCurrencySymbolKey"
	LocaleDecimalSeparator                    LocaleKey = "kCFLocaleDecimalSeparatorKey"
	LocaleExemplarCharacterSet                LocaleKey = "kCFLocaleExemplarCharacterSetKey"
	LocaleGroupingSeparator                   LocaleKey = "kCFLocaleGroupingSeparatorKey"
	LocaleIdentifier                          LocaleKey = "kCFLocaleIdentifierKey"
	LocaleLanguageCode                        LocaleKey = "kCFLocaleLanguageCodeKey"
	LocaleMeasurementSystem                   LocaleKey = "kCFLocaleMeasurementSystemKey"
	LocaleQuotationBeginDelimiterKey          LocaleKey = "kCFLocaleQuotationBeginDelimiterKey"
	LocaleQuotationEndDelimiterKey            LocaleKey = "kCFLocaleQuotationEndDelimiterKey"
	LocaleScriptCode                          LocaleKey = "kCFLocaleScriptCodeKey"
	LocaleUsesMetricSystem                    LocaleKey = "kCFLocaleUsesMetricSystemKey"
	LocaleVariantCode                         LocaleKey = "kCFLocaleVariantCodeKey"
)

// The directions that a language may take across a page of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocalelanguagedirection?language=objc
type LocaleLanguageDirection uint

const (
	LocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 4
	LocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 1
	LocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 2
	LocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 3
	LocaleLanguageDirectionUnknown     LocaleLanguageDirection = 0
)

// Used to remove access rights to a mach port when the NSMachPort object is invalidated or destroyed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachportoptions?language=objc
type MachPortOptions uint

const (
	MachPortDeallocateNone         MachPortOptions = 0
	MachPortDeallocateReceiveRight MachPortOptions = 2
	MachPortDeallocateSendRight    MachPortOptions = 1
)

// Constants used as components in a bitfield to specify the behavior of elements (keys and values) in an NSMapTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptableoptions?language=objc
type MapTableOptions uint

// The units supported by the NSMassFormatter class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatterunit?language=objc
type MassFormatterUnit int

const (
	MassFormatterUnitGram     MassFormatterUnit = 11
	MassFormatterUnitKilogram MassFormatterUnit = 14
	MassFormatterUnitOunce    MassFormatterUnit = 1537
	MassFormatterUnitPound    MassFormatterUnit = 1538
	MassFormatterUnitStone    MassFormatterUnit = 1539
)

// Set by the Block as the matching progresses, completes, or fails. Used by the method enumerateMatchesInString:options:range:usingBlock:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmatchingflags?language=objc
type MatchingFlags uint

const (
	MatchingCompleted     MatchingFlags = 2
	MatchingHitEnd        MatchingFlags = 4
	MatchingInternalError MatchingFlags = 16
	MatchingProgress      MatchingFlags = 1
	MatchingRequiredEnd   MatchingFlags = 8
)

// The matching options constants specify the reporting, completion and matching rules to the expression matching methods. These constants are used by all methods that search for, or replace values, using a regular expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmatchingoptions?language=objc
type MatchingOptions uint

const (
	MatchingAnchored               MatchingOptions = 4
	MatchingReportCompletion       MatchingOptions = 2
	MatchingReportProgress         MatchingOptions = 1
	MatchingWithTransparentBounds  MatchingOptions = 8
	MatchingWithoutAnchoringBounds MatchingOptions = 16
)

// Measurement formatter options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatterunitoptions?language=objc
type MeasurementFormatterUnitOptions uint

const (
	MeasurementFormatterUnitOptionsNaturalScale           MeasurementFormatterUnitOptions = 2
	MeasurementFormatterUnitOptionsProvidedUnit           MeasurementFormatterUnitOptions = 1
	MeasurementFormatterUnitOptionsTemperatureWithoutUnit MeasurementFormatterUnitOptions = 4
)

// These constants specify options for a network service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetserviceoptions?language=objc
type NetServiceOptions uint

const (
	NetServiceListenForConnections NetServiceOptions = 2
	NetServiceNoAutoRename         NetServiceOptions = 1
)

// These constants identify errors that can occur when accessing net services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetserviceserror?language=objc
type NetServicesError int

const (
	NetServicesActivityInProgress                NetServicesError = -72003
	NetServicesBadArgumentError                  NetServicesError = -72004
	NetServicesCancelledError                    NetServicesError = -72005
	NetServicesCollisionError                    NetServicesError = -72001
	NetServicesInvalidError                      NetServicesError = -72006
	NetServicesMissingRequiredConfigurationError NetServicesError = -72008
	NetServicesNotFoundError                     NetServicesError = -72002
	NetServicesTimeoutError                      NetServicesError = -72007
	NetServicesUnknownError                      NetServicesError = -72000
)

// The constants that specify how notifications are coalesced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcoalescing?language=objc
type NotificationCoalescing uint

const (
	NotificationCoalescingOnName   NotificationCoalescing = 1
	NotificationCoalescingOnSender NotificationCoalescing = 2
	NotificationNoCoalescing       NotificationCoalescing = 0
)

// A structure that defines the name of a notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationname?language=objc
type NotificationName string

const (
	AppleEventManagerWillProcessFirstEventNotification     NotificationName = "NSAppleEventManagerWillProcessFirstEvent"
	BundleDidLoadNotification                              NotificationName = "NSBundleDidLoadNotification"
	CalendarDayChangedNotification                         NotificationName = "NSCalendarDayChangedNotification"
	ClassDescriptionNeededForClassNotification             NotificationName = "NSClassDescriptionNeededForClassNotification"
	CurrentLocaleDidChangeNotification                     NotificationName = "kCFLocaleCurrentLocaleDidChangeNotification"
	DidBecomeSingleThreadedNotification                    NotificationName = "NSDidBecomeSingleThreadedNotification"
	FileHandleConnectionAcceptedNotification               NotificationName = "NSFileHandleConnectionAcceptedNotification"
	FileHandleDataAvailableNotification                    NotificationName = "NSFileHandleDataAvailableNotification"
	FileHandleReadCompletionNotification                   NotificationName = "NSFileHandleReadCompletionNotification"
	FileHandleReadToEndOfFileCompletionNotification        NotificationName = "NSFileHandleReadToEndOfFileCompletionNotification"
	HTTPCookieManagerAcceptPolicyChangedNotification       NotificationName = "com.apple.Foundation.NSHTTPCookieManagerAcceptPolicyChanged"
	HTTPCookieManagerCookiesChangedNotification            NotificationName = "NSHTTPCookieManagerCookiesChangedNotification"
	MetadataQueryDidFinishGatheringNotification            NotificationName = "NSMetadataQueryDidFinishGatheringNotification"
	MetadataQueryDidStartGatheringNotification             NotificationName = "NSMetadataQueryDidStartGatheringNotification"
	MetadataQueryDidUpdateNotification                     NotificationName = "NSMetadataQueryDidUpdateNotification"
	MetadataQueryGatheringProgressNotification             NotificationName = "NSMetadataQueryGatheringProgressNotification"
	PortDidBecomeInvalidNotification                       NotificationName = "NSPortDidBecomeInvalidNotification"
	ProcessInfoPowerStateDidChangeNotification             NotificationName = "NSProcessInfoPowerStateDidChangeNotification"
	ProcessInfoThermalStateDidChangeNotification           NotificationName = "NSProcessInfoThermalStateDidChangeNotification"
	SystemClockDidChangeNotification                       NotificationName = "NSSystemClockDidChangeNotification"
	SystemTimeZoneDidChangeNotification                    NotificationName = "kCFTimeZoneSystemTimeZoneDidChangeNotification"
	TaskDidTerminateNotification                           NotificationName = "NSTaskDidTerminateNotification"
	ThreadWillExitNotification                             NotificationName = "NSThreadWillExitNotification"
	URLCredentialStorageChangedNotification                NotificationName = "NSURLCredentialStorageChangedNotification"
	UbiquitousKeyValueStoreDidChangeExternallyNotification NotificationName = "NSUbiquitousKeyValueStoreDidChangeExternallyNotification"
	UbiquityIdentityDidChangeNotification                  NotificationName = "NSUbiquityIdentityDidChangeNotification"
	UndoManagerCheckpointNotification                      NotificationName = "NSUndoManagerCheckpointNotification"
	UndoManagerDidCloseUndoGroupNotification               NotificationName = "NSUndoManagerDidCloseUndoGroupNotification"
	UndoManagerDidOpenUndoGroupNotification                NotificationName = "NSUndoManagerDidOpenUndoGroupNotification"
	UndoManagerDidRedoChangeNotification                   NotificationName = "NSUndoManagerDidRedoChangeNotification"
	UndoManagerDidUndoChangeNotification                   NotificationName = "NSUndoManagerDidUndoChangeNotification"
	UndoManagerWillCloseUndoGroupNotification              NotificationName = "NSUndoManagerWillCloseUndoGroupNotification"
	UndoManagerWillRedoChangeNotification                  NotificationName = "NSUndoManagerWillRedoChangeNotification"
	UndoManagerWillUndoChangeNotification                  NotificationName = "NSUndoManagerWillUndoChangeNotification"
	UserDefaultsDidChangeNotification                      NotificationName = "NSUserDefaultsDidChangeNotification"
	WillBecomeMultiThreadedNotification                    NotificationName = "NSWillBecomeMultiThreadedNotification"
)

// These constants specify the types of notification delivery suspension behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationsuspensionbehavior?language=objc
type NotificationSuspensionBehavior uint

const (
	NotificationSuspensionBehaviorCoalesce           NotificationSuspensionBehavior = 2
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 4
	NotificationSuspensionBehaviorDrop               NotificationSuspensionBehavior = 1
	NotificationSuspensionBehaviorHold               NotificationSuspensionBehavior = 3
)

// These constants specify the behavior of a number formatter. These constants are returned by the defaultFormatterBehavior class method and the formatterBehavior property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatterbehavior?language=objc
type NumberFormatterBehavior uint

const (
	NumberFormatterBehavior10_0    NumberFormatterBehavior = 1000
	NumberFormatterBehavior10_4    NumberFormatterBehavior = 1040
	NumberFormatterBehaviorDefault NumberFormatterBehavior = 0
)

// These constants are used to specify how numbers should be padded. These constants are used by the paddingPosition property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatterpadposition?language=objc
type NumberFormatterPadPosition uint

const (
	NumberFormatterPadAfterPrefix  NumberFormatterPadPosition = 1
	NumberFormatterPadAfterSuffix  NumberFormatterPadPosition = 3
	NumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
	NumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 2
)

// These constants are used to specify how numbers should be rounded. These constants are used by the roundingMode property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatterroundingmode?language=objc
type NumberFormatterRoundingMode uint

const (
	NumberFormatterRoundCeiling  NumberFormatterRoundingMode = 0
	NumberFormatterRoundDown     NumberFormatterRoundingMode = 2
	NumberFormatterRoundFloor    NumberFormatterRoundingMode = 1
	NumberFormatterRoundHalfDown NumberFormatterRoundingMode = 5
	NumberFormatterRoundHalfEven NumberFormatterRoundingMode = 4
	NumberFormatterRoundHalfUp   NumberFormatterRoundingMode = 6
	NumberFormatterRoundUp       NumberFormatterRoundingMode = 3
)

// The predefined number format styles used by the numberStyle property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatterstyle?language=objc
type NumberFormatterStyle uint

const (
	NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 10
	NumberFormatterCurrencyISOCodeStyle    NumberFormatterStyle = 8
	NumberFormatterCurrencyPluralStyle     NumberFormatterStyle = 9
	NumberFormatterCurrencyStyle           NumberFormatterStyle = 2
	NumberFormatterDecimalStyle            NumberFormatterStyle = 1
	NumberFormatterNoStyle                 NumberFormatterStyle = 0
	NumberFormatterOrdinalStyle            NumberFormatterStyle = 6
	NumberFormatterPercentStyle            NumberFormatterStyle = 3
	NumberFormatterScientificStyle         NumberFormatterStyle = 4
	NumberFormatterSpellOutStyle           NumberFormatterStyle = 5
)

// These constants let you prioritize the order in which operations execute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueuepriority?language=objc
type OperationQueuePriority int

const (
	OperationQueuePriorityHigh     OperationQueuePriority = 4
	OperationQueuePriorityLow      OperationQueuePriority = -4
	OperationQueuePriorityNormal   OperationQueuePriority = 0
	OperationQueuePriorityVeryHigh OperationQueuePriority = 8
	OperationQueuePriorityVeryLow  OperationQueuePriority = -8
)

// Constants that specify the options to use when creating an ordered collection difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifferencecalculationoptions?language=objc
type OrderedCollectionDifferenceCalculationOptions uint

const (
	OrderedCollectionDifferenceCalculationInferMoves          OrderedCollectionDifferenceCalculationOptions = 4
	OrderedCollectionDifferenceCalculationOmitInsertedObjects OrderedCollectionDifferenceCalculationOptions = 1
	OrderedCollectionDifferenceCalculationOmitRemovedObjects  OrderedCollectionDifferenceCalculationOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatteroptions?language=objc
type PersonNameComponentsFormatterOptions uint

const (
	PersonNameComponentsFormatterPhonetic PersonNameComponentsFormatterOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatterstyle?language=objc
type PersonNameComponentsFormatterStyle int

const (
	PersonNameComponentsFormatterStyleAbbreviated PersonNameComponentsFormatterStyle = 4
	PersonNameComponentsFormatterStyleDefault     PersonNameComponentsFormatterStyle = 0
	PersonNameComponentsFormatterStyleLong        PersonNameComponentsFormatterStyle = 3
	PersonNameComponentsFormatterStyleMedium      PersonNameComponentsFormatterStyle = 2
	PersonNameComponentsFormatterStyleShort       PersonNameComponentsFormatterStyle = 1
)

// Defines the memory and personality options for an NSPointerFunctions object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctionsoptions?language=objc
type PointerFunctionsOptions uint

const (
	HashTableCopyIn                          PointerFunctionsOptions = 65536
	HashTableObjectPointerPersonality        PointerFunctionsOptions = 512
	HashTableStrongMemory                    PointerFunctionsOptions = 0
	HashTableWeakMemory                      PointerFunctionsOptions = 5
	HashTableZeroingWeakMemory               PointerFunctionsOptions = 1
	MapTableCopyIn                           PointerFunctionsOptions = 65536
	MapTableObjectPointerPersonality         PointerFunctionsOptions = 512
	MapTableStrongMemory                     PointerFunctionsOptions = 0
	MapTableWeakMemory                       PointerFunctionsOptions = 5
	MapTableZeroingWeakMemory                PointerFunctionsOptions = 1
	PointerFunctionsCStringPersonality       PointerFunctionsOptions = 768
	PointerFunctionsCopyIn                   PointerFunctionsOptions = 65536
	PointerFunctionsIntegerPersonality       PointerFunctionsOptions = 1280
	PointerFunctionsMachVirtualMemory        PointerFunctionsOptions = 4
	PointerFunctionsMallocMemory             PointerFunctionsOptions = 3
	PointerFunctionsObjectPersonality        PointerFunctionsOptions = 0
	PointerFunctionsObjectPointerPersonality PointerFunctionsOptions = 512
	PointerFunctionsOpaqueMemory             PointerFunctionsOptions = 2
	PointerFunctionsOpaquePersonality        PointerFunctionsOptions = 256
	PointerFunctionsStrongMemory             PointerFunctionsOptions = 0
	PointerFunctionsStructPersonality        PointerFunctionsOptions = 1024
	PointerFunctionsWeakMemory               PointerFunctionsOptions = 5
	PointerFunctionsZeroingWeakMemory        PointerFunctionsOptions = 1
)

// The constants that specify when notifications are posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspostingstyle?language=objc
type PostingStyle uint

const (
	PostASAP     PostingStyle = 2
	PostNow      PostingStyle = 3
	PostWhenIdle PostingStyle = 1
)

// Defines the type of comparison for a comparison predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicateoperatortype?language=objc
type PredicateOperatorType uint

const (
	BeginsWithPredicateOperatorType           PredicateOperatorType = 8
	BetweenPredicateOperatorType              PredicateOperatorType = 100
	ContainsPredicateOperatorType             PredicateOperatorType = 99
	CustomSelectorPredicateOperatorType       PredicateOperatorType = 11
	EndsWithPredicateOperatorType             PredicateOperatorType = 9
	EqualToPredicateOperatorType              PredicateOperatorType = 4
	GreaterThanOrEqualToPredicateOperatorType PredicateOperatorType = 3
	GreaterThanPredicateOperatorType          PredicateOperatorType = 2
	InPredicateOperatorType                   PredicateOperatorType = 10
	LessThanOrEqualToPredicateOperatorType    PredicateOperatorType = 1
	LessThanPredicateOperatorType             PredicateOperatorType = 0
	LikePredicateOperatorType                 PredicateOperatorType = 7
	MatchesPredicateOperatorType              PredicateOperatorType = 6
	NotEqualToPredicateOperatorType           PredicateOperatorType = 5
)

// An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintentkind?language=objc
type PresentationIntentKind int

const (
	PresentationIntentKindBlockQuote     PresentationIntentKind = 6
	PresentationIntentKindCodeBlock      PresentationIntentKind = 5
	PresentationIntentKindHeader         PresentationIntentKind = 1
	PresentationIntentKindListItem       PresentationIntentKind = 4
	PresentationIntentKindOrderedList    PresentationIntentKind = 2
	PresentationIntentKindParagraph      PresentationIntentKind = 0
	PresentationIntentKindTable          PresentationIntentKind = 8
	PresentationIntentKindTableCell      PresentationIntentKind = 11
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
	PresentationIntentKindTableRow       PresentationIntentKind = 10
	PresentationIntentKindThematicBreak  PresentationIntentKind = 7
	PresentationIntentKindUnorderedList  PresentationIntentKind = 3
)

// An enumeration of values for aligning the contents of table columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintenttablecolumnalignment?language=objc
type PresentationIntentTableColumnAlignment int

const (
	PresentationIntentTableColumnAlignmentCenter PresentationIntentTableColumnAlignment = 1
	PresentationIntentTableColumnAlignmentLeft   PresentationIntentTableColumnAlignment = 0
	PresentationIntentTableColumnAlignmentRight  PresentationIntentTableColumnAlignment = 2
)

// Values used to indicate the system’s thermal state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfothermalstate?language=objc
type ProcessInfoThermalState int

const (
	ProcessInfoThermalStateCritical ProcessInfoThermalState = 3
	ProcessInfoThermalStateFair     ProcessInfoThermalState = 1
	ProcessInfoThermalStateNominal  ProcessInfoThermalState = 0
	ProcessInfoThermalStateSerious  ProcessInfoThermalState = 2
)

// The kind of file operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogressfileoperationkind?language=objc
type ProgressFileOperationKind string

const (
	ProgressFileOperationKindCopying                       ProgressFileOperationKind = "NSProgressFileOperationKindCopying"
	ProgressFileOperationKindDecompressingAfterDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDecompressingAfterDownloading"
	ProgressFileOperationKindDownloading                   ProgressFileOperationKind = "NSProgressFileOperationKindDownloading"
	ProgressFileOperationKindDuplicating                   ProgressFileOperationKind = "NSProgressFileOperationKindDuplicating"
	ProgressFileOperationKindReceiving                     ProgressFileOperationKind = "NSProgressFileOperationKindReceiving"
	ProgressFileOperationKindUploading                     ProgressFileOperationKind = "NSProgressFileOperationKindUploading"
)

// An object that represents the kind of progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogresskind?language=objc
type ProgressKind string

const (
	ProgressKindFile ProgressKind = "NSProgressKindFile"
)

// Keys for the user info dictionary that affect the autogenerated localized additional description string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogressuserinfokey?language=objc
type ProgressUserInfoKey string

const (
	ProgressEstimatedTimeRemainingKey         ProgressUserInfoKey = "NSProgressEstimatedTimeRemainingKey"
	ProgressFileAnimationImageKey             ProgressUserInfoKey = "NSProgressFlyToImageKey"
	ProgressFileAnimationImageOriginalRectKey ProgressUserInfoKey = "NSProgressFileAnimationImageOriginalRectKey"
	ProgressFileCompletedCountKey             ProgressUserInfoKey = "NSProgressFileCompletedCountKey"
	ProgressFileIconKey                       ProgressUserInfoKey = "NSProgressFileIconKey"
	ProgressFileOperationKindKey              ProgressUserInfoKey = "NSProgressFileOperationKindKey"
	ProgressFileTotalCountKey                 ProgressUserInfoKey = "NSProgressFileTotalCountKey"
	ProgressFileURLKey                        ProgressUserInfoKey = "NSProgressFileURLKey"
	ProgressThroughputKey                     ProgressUserInfoKey = "NSProgressThroughputKey"
)

// These constants are used to specify a property list serialization format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistformat?language=objc
type PropertyListFormat uint

const (
	PropertyListBinaryFormat_v1_0 PropertyListFormat = 200
	PropertyListOpenStepFormat    PropertyListFormat = 1
	PropertyListXMLFormat_v1_0    PropertyListFormat = 100
)

// These constants specify mutability options in property lists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistmutabilityoptions?language=objc
type PropertyListMutabilityOptions uint

const (
	PropertyListImmutable                  PropertyListMutabilityOptions = 0
	PropertyListMutableContainers          PropertyListMutabilityOptions = 1
	PropertyListMutableContainersAndLeaves PropertyListMutabilityOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwriteoptions?language=objc
type PropertyListWriteOptions uint

// Constants that indicate the nature and importance of work to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsqualityofservice?language=objc
type QualityOfService int

const (
	QualityOfServiceBackground      QualityOfService = 9
	QualityOfServiceDefault         QualityOfService = -1
	QualityOfServiceUserInitiated   QualityOfService = 25
	QualityOfServiceUserInteractive QualityOfService = 33
	QualityOfServiceUtility         QualityOfService = 17
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrectedge?language=objc
type RectEdge uint

const (
	RectEdgeMaxX RectEdge = 2
	RectEdgeMaxY RectEdge = 3
	RectEdgeMinX RectEdge = 0
	RectEdgeMinY RectEdge = 1
)

// These constants define the regular expression options. These constants are used by the property options, regularExpressionWithPattern:options:error:, and initWithPattern:options:error:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpressionoptions?language=objc
type RegularExpressionOptions uint

const (
	RegularExpressionAllowCommentsAndWhitespace RegularExpressionOptions = 2
	RegularExpressionAnchorsMatchLines          RegularExpressionOptions = 16
	RegularExpressionCaseInsensitive            RegularExpressionOptions = 1
	RegularExpressionDotMatchesLineSeparators   RegularExpressionOptions = 8
	RegularExpressionIgnoreMetacharacters       RegularExpressionOptions = 4
	RegularExpressionUseUnicodeWordBoundaries   RegularExpressionOptions = 64
	RegularExpressionUseUnixLineSeparators      RegularExpressionOptions = 32
)

// A type that represents the style to use when formatting relative dates, such as “1 week ago” or “last week”. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatterstyle?language=objc
type RelativeDateTimeFormatterStyle int

const (
	RelativeDateTimeFormatterStyleNamed   RelativeDateTimeFormatterStyle = 1
	RelativeDateTimeFormatterStyleNumeric RelativeDateTimeFormatterStyle = 0
)

// A type that represents the style to use when formatting the units of relative dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatterunitsstyle?language=objc
type RelativeDateTimeFormatterUnitsStyle int

const (
	RelativeDateTimeFormatterUnitsStyleAbbreviated RelativeDateTimeFormatterUnitsStyle = 3
	RelativeDateTimeFormatterUnitsStyleFull        RelativeDateTimeFormatterUnitsStyle = 0
	RelativeDateTimeFormatterUnitsStyleShort       RelativeDateTimeFormatterUnitsStyle = 2
	RelativeDateTimeFormatterUnitsStyleSpellOut    RelativeDateTimeFormatterUnitsStyle = 1
)

// These constants are used by relativePosition and relativePosition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativeposition?language=objc
type RelativePosition uint

const (
	RelativeAfter  RelativePosition = 0
	RelativeBefore RelativePosition = 1
)

// These constants specify rounding behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsroundingmode?language=objc
type RoundingMode uint

const (
	RoundBankers RoundingMode = 3
	RoundDown    RoundingMode = 1
	RoundPlain   RoundingMode = 0
	RoundUp      RoundingMode = 2
)

// Modes that a run loop operates in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloopmode?language=objc
type RunLoopMode string

const (
	DefaultRunLoopMode RunLoopMode = "kCFRunLoopDefaultMode"
	RunLoopCommonModes RunLoopMode = "kCFRunLoopCommonModes"
)

// The saveOptions method returns one of the following constants to indicate how to deal with saving any modified documents: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssaveoptions?language=objc
type SaveOptions uint

const (
	SaveOptionsAsk SaveOptions = 2
	SaveOptionsNo  SaveOptions = 1
	SaveOptionsYes SaveOptions = 0
)

// The location of significant directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssearchpathdirectory?language=objc
type SearchPathDirectory uint

const (
	AdminApplicationDirectory     SearchPathDirectory = 4
	AllApplicationsDirectory      SearchPathDirectory = 100
	AllLibrariesDirectory         SearchPathDirectory = 101
	ApplicationDirectory          SearchPathDirectory = 1
	ApplicationScriptsDirectory   SearchPathDirectory = 23
	ApplicationSupportDirectory   SearchPathDirectory = 14
	AutosavedInformationDirectory SearchPathDirectory = 11
	CachesDirectory               SearchPathDirectory = 13
	CoreServiceDirectory          SearchPathDirectory = 10
	DemoApplicationDirectory      SearchPathDirectory = 2
	DesktopDirectory              SearchPathDirectory = 12
	DeveloperApplicationDirectory SearchPathDirectory = 3
	DeveloperDirectory            SearchPathDirectory = 6
	DocumentDirectory             SearchPathDirectory = 9
	DocumentationDirectory        SearchPathDirectory = 8
	DownloadsDirectory            SearchPathDirectory = 15
	InputMethodsDirectory         SearchPathDirectory = 16
	ItemReplacementDirectory      SearchPathDirectory = 99
	LibraryDirectory              SearchPathDirectory = 5
	MoviesDirectory               SearchPathDirectory = 17
	MusicDirectory                SearchPathDirectory = 18
	PicturesDirectory             SearchPathDirectory = 19
	PreferencePanesDirectory      SearchPathDirectory = 22
	PrinterDescriptionDirectory   SearchPathDirectory = 20
	SharedPublicDirectory         SearchPathDirectory = 21
	TrashDirectory                SearchPathDirectory = 102
	UserDirectory                 SearchPathDirectory = 7
)

// Domain constants specifying base locations to use when you search for significant directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssearchpathdomainmask?language=objc
type SearchPathDomainMask uint

const (
	AllDomainsMask    SearchPathDomainMask = 65535
	LocalDomainMask   SearchPathDomainMask = 2
	NetworkDomainMask SearchPathDomainMask = 4
	SystemDomainMask  SearchPathDomainMask = 8
	UserDomainMask    SearchPathDomainMask = 1
)

// Type for the platform-specific native socket handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketnativehandle?language=objc
type SocketNativeHandle int

// Options for block sorting operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortoptions?language=objc
type SortOptions uint

const (
	SortConcurrent SortOptions = 1
	SortStable     SortOptions = 16
)

// Describes the constants that may be sent to the delegate as a bit field in the second parameter of stream:handleEvent: to specify the kind of stream event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamevent?language=objc
type StreamEvent uint

const (
	StreamEventEndEncountered    StreamEvent = 16
	StreamEventErrorOccurred     StreamEvent = 8
	StreamEventHasBytesAvailable StreamEvent = 2
	StreamEventHasSpaceAvailable StreamEvent = 4
	StreamEventNone              StreamEvent = 0
	StreamEventOpenCompleted     StreamEvent = 1
)

// NSStream defines these string constants for specifying the service type of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamnetworkservicetypevalue?language=objc
type StreamNetworkServiceTypeValue string

const (
	StreamNetworkServiceTypeBackground    StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeBackground"
	StreamNetworkServiceTypeCallSignaling StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeCallSignaling"
	StreamNetworkServiceTypeVideo         StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVideo"
	StreamNetworkServiceTypeVoIP          StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVoIP"
	StreamNetworkServiceTypeVoice         StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVoice"
)

// NSStream defines these string constants as keys for accessing stream properties using propertyForKey: and setting properties with setProperty:forKey:: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreampropertykey?language=objc
type StreamPropertyKey string

const (
	StreamDataWrittenToMemoryStreamKey StreamPropertyKey = "kCFStreamPropertyDataWritten"
	StreamFileCurrentOffsetKey         StreamPropertyKey = "kCFStreamPropertyFileCurrentOffset"
	StreamNetworkServiceType           StreamPropertyKey = "kCFStreamNetworkServiceType"
	StreamSOCKSProxyConfigurationKey   StreamPropertyKey = "kCFStreamPropertySOCKSProxy"
	StreamSocketSecurityLevelKey       StreamPropertyKey = "kCFStreamPropertySocketSecurityLevel"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocksproxyconfiguration?language=objc
type StreamSOCKSProxyConfiguration string

const (
	StreamSOCKSProxyHostKey     StreamSOCKSProxyConfiguration = "SOCKSProxy"
	StreamSOCKSProxyPasswordKey StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSPassword"
	StreamSOCKSProxyPortKey     StreamSOCKSProxyConfiguration = "SOCKSPort"
	StreamSOCKSProxyUserKey     StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSUser"
	StreamSOCKSProxyVersionKey  StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSVersion"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocksproxyversion?language=objc
type StreamSOCKSProxyVersion string

const (
	StreamSOCKSProxyVersion4 StreamSOCKSProxyVersion = "kCFStreamSocketSOCKSVersion4"
	StreamSOCKSProxyVersion5 StreamSOCKSProxyVersion = "kCFStreamSocketSOCKSVersion5"
)

// NSStream defines these string constants for specifying the secure-socket layer (SSL) security level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsecuritylevel?language=objc
type StreamSocketSecurityLevel string

const (
	StreamSocketSecurityLevelNegotiatedSSL StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelNegotiatedSSL"
	StreamSocketSecurityLevelNone          StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelNone"
	StreamSocketSecurityLevelSSLv2         StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelSSLv2"
	StreamSocketSecurityLevelSSLv3         StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelSSLv3"
	StreamSocketSecurityLevelTLSv1         StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelTLSv1"
)

// The type declared for the constants listed in [foundation/nsstream/stream_status_constants]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamstatus?language=objc
type StreamStatus uint

const (
	StreamStatusAtEnd   StreamStatus = 5
	StreamStatusClosed  StreamStatus = 6
	StreamStatusError   StreamStatus = 7
	StreamStatusNotOpen StreamStatus = 0
	StreamStatusOpen    StreamStatus = 2
	StreamStatusOpening StreamStatus = 1
	StreamStatusReading StreamStatus = 3
	StreamStatusWriting StreamStatus = 4
)

// These values represent the options available to many of the string classes’ search and comparison methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringcompareoptions?language=objc
type StringCompareOptions uint

const (
	AnchoredSearch             StringCompareOptions = 8
	BackwardsSearch            StringCompareOptions = 4
	CaseInsensitiveSearch      StringCompareOptions = 1
	DiacriticInsensitiveSearch StringCompareOptions = 128
	ForcedOrderingSearch       StringCompareOptions = 512
	LiteralSearch              StringCompareOptions = 2
	NumericSearch              StringCompareOptions = 64
	RegularExpressionSearch    StringCompareOptions = 1024
	WidthInsensitiveSearch     StringCompareOptions = 256
)

// The following constants are provided by NSString as possible string encodings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringencoding?language=objc
type StringEncoding uint

// Options for converting string encodings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringencodingconversionoptions?language=objc
type StringEncodingConversionOptions uint

const (
	StringEncodingConversionAllowLossy             StringEncodingConversionOptions = 1
	StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringencodingdetectionoptionskey?language=objc
type StringEncodingDetectionOptionsKey string

const (
	StringEncodingDetectionAllowLossyKey                StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionAllowLossyKey"
	StringEncodingDetectionDisallowedEncodingsKey       StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionDisallowedEncodingsKey"
	StringEncodingDetectionFromWindowsKey               StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionFromWindowsKey"
	StringEncodingDetectionLikelyLanguageKey            StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionLikelyLanguageKey"
	StringEncodingDetectionLossySubstitutionKey         StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionLossySubstitutionKey"
	StringEncodingDetectionSuggestedEncodingsKey        StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionSuggestedEncodingsKey"
	StringEncodingDetectionUseOnlySuggestedEncodingsKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionUseOnlySuggestedEncodingsKey"
)

// Constants to specify kinds of substrings and styles of enumeration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringenumerationoptions?language=objc
type StringEnumerationOptions uint

const (
	StringEnumerationByCaretPositions             StringEnumerationOptions = 5
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
	StringEnumerationByDeletionClusters           StringEnumerationOptions = 6
	StringEnumerationByLines                      StringEnumerationOptions = 0
	StringEnumerationByParagraphs                 StringEnumerationOptions = 1
	StringEnumerationBySentences                  StringEnumerationOptions = 4
	StringEnumerationByWords                      StringEnumerationOptions = 3
	StringEnumerationLocalized                    StringEnumerationOptions = 1024
	StringEnumerationReverse                      StringEnumerationOptions = 256
	StringEnumerationSubstringNotRequired         StringEnumerationOptions = 512
)

// Constants representing an ICU string transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstringtransform?language=objc
type StringTransform string

const (
	StringTransformFullwidthToHalfwidth StringTransform = ")kCFStringTransformFullwidthHalfwidth"
	StringTransformHiraganaToKatakana   StringTransform = ")kCFStringTransformHiraganaKatakana"
	StringTransformLatinToArabic        StringTransform = ")kCFStringTransformLatinArabic"
	StringTransformLatinToCyrillic      StringTransform = ")kCFStringTransformLatinCyrillic"
	StringTransformLatinToGreek         StringTransform = ")kCFStringTransformLatinGreek"
	StringTransformLatinToHangul        StringTransform = ")kCFStringTransformLatinHangul"
	StringTransformLatinToHebrew        StringTransform = ")kCFStringTransformLatinHebrew"
	StringTransformLatinToHiragana      StringTransform = ")kCFStringTransformLatinHiragana"
	StringTransformLatinToKatakana      StringTransform = ")kCFStringTransformLatinKatakana"
	StringTransformLatinToThai          StringTransform = ")kCFStringTransformLatinThai"
	StringTransformMandarinToLatin      StringTransform = ")kCFStringTransformMandarinLatin"
	StringTransformStripCombiningMarks  StringTransform = ")kCFStringTransformStripCombiningMarks"
	StringTransformStripDiacritics      StringTransform = ")kCFStringTransformStripDiacritics"
	StringTransformToLatin              StringTransform = ")kCFStringTransformToLatin"
	StringTransformToUnicodeName        StringTransform = ")kCFStringTransformToUnicodeName"
	StringTransformToXMLHex             StringTransform = ")kCFStringTransformToXMLHex"
)

// Constants that specify the termination reason values that the system returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstaskterminationreason?language=objc
type TaskTerminationReason int

const (
	TaskTerminationReasonExit           TaskTerminationReason = 1
	TaskTerminationReasonUncaughtSignal TaskTerminationReason = 2
)

// These are passed to  initWithObjectSpecifier:comparisonOperator:testObject: to specify the comparison operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstestcomparisonoperation?language=objc
type TestComparisonOperation uint

const (
	BeginsWithComparison           TestComparisonOperation = 5
	ContainsComparison             TestComparisonOperation = 7
	EndsWithComparison             TestComparisonOperation = 6
	EqualToComparison              TestComparisonOperation = 0
	GreaterThanComparison          TestComparisonOperation = 4
	GreaterThanOrEqualToComparison TestComparisonOperation = 3
	LessThanComparison             TestComparisonOperation = 2
	LessThanOrEqualToComparison    TestComparisonOperation = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingkey?language=objc
type TextCheckingKey string

const (
	TextCheckingAirlineKey      TextCheckingKey = "Airline"
	TextCheckingCityKey         TextCheckingKey = "City"
	TextCheckingCountryKey      TextCheckingKey = "Country"
	TextCheckingFlightKey       TextCheckingKey = "Flight"
	TextCheckingJobTitleKey     TextCheckingKey = "JobTitle"
	TextCheckingNameKey         TextCheckingKey = "Name"
	TextCheckingOrganizationKey TextCheckingKey = "Organization"
	TextCheckingPhoneKey        TextCheckingKey = "Phone"
	TextCheckingStateKey        TextCheckingKey = "State"
	TextCheckingStreetKey       TextCheckingKey = "Street"
	TextCheckingZIPKey          TextCheckingKey = "ZIP"
)

// These constants specify the type of checking the methods should do. They are returned by resultType. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingtype?language=objc
type TextCheckingType uint64

const (
	TextCheckingTypeAddress            TextCheckingType = 16
	TextCheckingTypeCorrection         TextCheckingType = 512
	TextCheckingTypeDash               TextCheckingType = 128
	TextCheckingTypeDate               TextCheckingType = 8
	TextCheckingTypeGrammar            TextCheckingType = 4
	TextCheckingTypeLink               TextCheckingType = 32
	TextCheckingTypeOrthography        TextCheckingType = 1
	TextCheckingTypePhoneNumber        TextCheckingType = 2048
	TextCheckingTypeQuote              TextCheckingType = 64
	TextCheckingTypeRegularExpression  TextCheckingType = 1024
	TextCheckingTypeReplacement        TextCheckingType = 256
	TextCheckingTypeSpelling           TextCheckingType = 2
	TextCheckingTypeTransitInformation TextCheckingType = 4096
)

// Defines the types of checking that are available. These values can be combined using the C-bitwise OR operator. The system supports its own internal types, and the user can extend those types by subclassing NSTextCheckingResult and adding their own custom types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingtypes?language=objc
type TextCheckingTypes uint64

// A number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimeinterval?language=objc
type TimeInterval float64

// Constants you use to specify a style when presenting time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezonenamestyle?language=objc
type TimeZoneNameStyle int

const (
	TimeZoneNameStyleDaylightSaving      TimeZoneNameStyle = 2
	TimeZoneNameStyleGeneric             TimeZoneNameStyle = 4
	TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
	TimeZoneNameStyleShortGeneric        TimeZoneNameStyle = 5
	TimeZoneNameStyleShortStandard       TimeZoneNameStyle = 1
	TimeZoneNameStyleStandard            TimeZoneNameStyle = 0
)

// Options used when creating bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlbookmarkcreationoptions?language=objc
type URLBookmarkCreationOptions uint

const (
	URLBookmarkCreationMinimalBookmark                  URLBookmarkCreationOptions = 512
	URLBookmarkCreationPreferFileIDResolution           URLBookmarkCreationOptions = 256
	URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 4096
	URLBookmarkCreationSuitableForBookmarkFile          URLBookmarkCreationOptions = 1024
	URLBookmarkCreationWithSecurityScope                URLBookmarkCreationOptions = 2048
	URLBookmarkCreationWithoutImplicitSecurityScope     URLBookmarkCreationOptions = 536870912
)

// Options used when creating file bookmark data [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlbookmarkfilecreationoptions?language=objc
type URLBookmarkFileCreationOptions uint

// Options used when resolving bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlbookmarkresolutionoptions?language=objc
type URLBookmarkResolutionOptions uint

const (
	URLBookmarkResolutionWithSecurityScope             URLBookmarkResolutionOptions = 1024
	URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 32768
	URLBookmarkResolutionWithoutMounting               URLBookmarkResolutionOptions = 512
	URLBookmarkResolutionWithoutUI                     URLBookmarkResolutionOptions = 256
)

// These constants specify the caching strategy used by an NSCachedURLResponse object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcachestoragepolicy?language=objc
type URLCacheStoragePolicy uint

const (
	URLCacheStorageAllowed             URLCacheStoragePolicy = 0
	URLCacheStorageAllowedInMemoryOnly URLCacheStoragePolicy = 1
	URLCacheStorageNotAllowed          URLCacheStoragePolicy = 2
)

// Constants that specify how long the credential will be kept. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialpersistence?language=objc
type URLCredentialPersistence uint

const (
	URLCredentialPersistenceForSession     URLCredentialPersistence = 1
	URLCredentialPersistenceNone           URLCredentialPersistence = 0
	URLCredentialPersistencePermanent      URLCredentialPersistence = 2
	URLCredentialPersistenceSynchronizable URLCredentialPersistence = 3
)

// An enumeration of reasons why a task couldn’t satisfy networking constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlerrornetworkunavailablereason?language=objc
type URLErrorNetworkUnavailableReason int

const (
	URLErrorNetworkUnavailableReasonCellular    URLErrorNetworkUnavailableReason = 0
	URLErrorNetworkUnavailableReasonConstrained URLErrorNetworkUnavailableReason = 2
	URLErrorNetworkUnavailableReasonExpensive   URLErrorNetworkUnavailableReason = 1
)

// Protection-level values for a URL resource key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlfileprotectiontype?language=objc
type URLFileProtectionType string

const (
	URLFileProtectionComplete                             URLFileProtectionType = "NSURLFileProtectionComplete"
	URLFileProtectionCompleteUnlessOpen                   URLFileProtectionType = "NSURLFileProtectionCompleteUnlessOpen"
	URLFileProtectionCompleteUntilFirstUserAuthentication URLFileProtectionType = "NSURLFileProtectionCompleteUntilFirstUserAuthentication"
	URLFileProtectionNone                                 URLFileProtectionType = "NSURLFileProtectionNone"
)

// Possible values for the type of file resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlfileresourcetype?language=objc
type URLFileResourceType string

const (
	URLFileResourceTypeBlockSpecial     URLFileResourceType = "NSURLFileResourceTypeBlockSpecial"
	URLFileResourceTypeCharacterSpecial URLFileResourceType = "NSURLFileResourceTypeCharacterSpecial"
	URLFileResourceTypeDirectory        URLFileResourceType = "NSURLFileResourceTypeDirectory"
	URLFileResourceTypeNamedPipe        URLFileResourceType = "NSURLFileResourceTypeNamedPipe"
	URLFileResourceTypeRegular          URLFileResourceType = "NSURLFileResourceTypeRegular"
	URLFileResourceTypeSocket           URLFileResourceType = "NSURLFileResourceTypeSocket"
	URLFileResourceTypeSymbolicLink     URLFileResourceType = "NSURLFileResourceTypeSymbolicLink"
	URLFileResourceTypeUnknown          URLFileResourceType = "NSURLFileResourceTypeUnknown"
)

// These following constants are defined by NSURLHandle and are returned by status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlhandlestatus?language=objc
type URLHandleStatus uint

const (
	URLHandleLoadFailed     URLHandleStatus = 3
	URLHandleLoadInProgress URLHandleStatus = 2
	URLHandleLoadSucceeded  URLHandleStatus = 1
	URLHandleNotLoaded      URLHandleStatus = 0
)

// Constants indicating the relationship between a directory and an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrelationship?language=objc
type URLRelationship int

const (
	URLRelationshipContains URLRelationship = 0
	URLRelationshipOther    URLRelationship = 2
	URLRelationshipSame     URLRelationship = 1
)

// The entities that can make a network request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequestattribution?language=objc
type URLRequestAttribution uint

const (
	URLRequestAttributionDeveloper URLRequestAttribution = 0
	URLRequestAttributionUser      URLRequestAttribution = 1
)

// The constants used to specify interaction with the cached responses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequestcachepolicy?language=objc
type URLRequestCachePolicy uint

const (
	URLRequestReloadIgnoringCacheData               URLRequestCachePolicy = 1
	URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
	URLRequestReloadIgnoringLocalCacheData          URLRequestCachePolicy = 1
	URLRequestReloadRevalidatingCacheData           URLRequestCachePolicy = 5
	URLRequestReturnCacheDataDontLoad               URLRequestCachePolicy = 3
	URLRequestReturnCacheDataElseLoad               URLRequestCachePolicy = 2
	URLRequestUseProtocolCachePolicy                URLRequestCachePolicy = 0
)

// Constants that specify how a request uses network resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequestnetworkservicetype?language=objc
type URLRequestNetworkServiceType uint

const (
	URLNetworkServiceTypeAVStreaming    URLRequestNetworkServiceType = 8
	URLNetworkServiceTypeBackground     URLRequestNetworkServiceType = 3
	URLNetworkServiceTypeCallSignaling  URLRequestNetworkServiceType = 11
	URLNetworkServiceTypeDefault        URLRequestNetworkServiceType = 0
	URLNetworkServiceTypeResponsiveAV   URLRequestNetworkServiceType = 9
	URLNetworkServiceTypeResponsiveData URLRequestNetworkServiceType = 6
	URLNetworkServiceTypeVideo          URLRequestNetworkServiceType = 2
	URLNetworkServiceTypeVoIP           URLRequestNetworkServiceType = 1
	URLNetworkServiceTypeVoice          URLRequestNetworkServiceType = 4
)

// Keys that apply to file system URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresourcekey?language=objc
type URLResourceKey string

const (
	URLAddedToDirectoryDateKey                               URLResourceKey = "NSURLAddedToDirectoryDateKey"
	URLApplicationIsScriptableKey                            URLResourceKey = "NSURLApplicationIsScriptableKey"
	URLAttributeModificationDateKey                          URLResourceKey = "NSURLAttributeModificationDateKey"
	URLCanonicalPathKey                                      URLResourceKey = "NSURLCanonicalPathKey"
	URLContentAccessDateKey                                  URLResourceKey = "NSURLContentAccessDateKey"
	URLContentModificationDateKey                            URLResourceKey = "NSURLContentModificationDateKey"
	URLContentTypeKey                                        URLResourceKey = "NSURLContentTypeKey"
	URLCreationDateKey                                       URLResourceKey = "NSURLCreationDateKey"
	URLCustomIconKey                                         URLResourceKey = "NSURLCustomIconKey"
	URLDocumentIdentifierKey                                 URLResourceKey = "NSURLDocumentIdentifierKey"
	URLEffectiveIconKey                                      URLResourceKey = "NSURLEffectiveIconKey"
	URLFileAllocatedSizeKey                                  URLResourceKey = "NSURLFileAllocatedSizeKey"
	URLFileContentIdentifierKey                              URLResourceKey = "NSURLFileContentIdentifierKey"
	URLFileProtectionKey                                     URLResourceKey = "NSURLFileProtectionKey"
	URLFileResourceIdentifierKey                             URLResourceKey = "NSURLFileResourceIdentifierKey"
	URLFileResourceTypeKey                                   URLResourceKey = "NSURLFileResourceTypeKey"
	URLFileSecurityKey                                       URLResourceKey = "NSURLFileSecurityKey"
	URLFileSizeKey                                           URLResourceKey = "NSURLFileSizeKey"
	URLGenerationIdentifierKey                               URLResourceKey = "NSURLGenerationIdentifierKey"
	URLHasHiddenExtensionKey                                 URLResourceKey = "NSURLHasHiddenExtensionKey"
	URLIsAliasFileKey                                        URLResourceKey = "NSURLIsAliasFileKey"
	URLIsApplicationKey                                      URLResourceKey = "_NSURLIsApplicationKey"
	URLIsDirectoryKey                                        URLResourceKey = "NSURLIsDirectoryKey"
	URLIsExcludedFromBackupKey                               URLResourceKey = "NSURLIsExcludedFromBackupKey"
	URLIsExecutableKey                                       URLResourceKey = "NSURLIsExecutableKey"
	URLIsHiddenKey                                           URLResourceKey = "NSURLIsHiddenKey"
	URLIsMountTriggerKey                                     URLResourceKey = "NSURLIsMountTriggerKey"
	URLIsPackageKey                                          URLResourceKey = "NSURLIsPackageKey"
	URLIsPurgeableKey                                        URLResourceKey = "NSURLIsPurgeableKey"
	URLIsReadableKey                                         URLResourceKey = "NSURLIsReadableKey"
	URLIsRegularFileKey                                      URLResourceKey = "NSURLIsRegularFileKey"
	URLIsSparseKey                                           URLResourceKey = "NSURLIsSparseKey"
	URLIsSymbolicLinkKey                                     URLResourceKey = "NSURLIsSymbolicLinkKey"
	URLIsSystemImmutableKey                                  URLResourceKey = "NSURLIsSystemImmutableKey"
	URLIsUbiquitousItemKey                                   URLResourceKey = "NSURLIsUbiquitousItemKey"
	URLIsUserImmutableKey                                    URLResourceKey = "NSURLIsUserImmutableKey"
	URLIsVolumeKey                                           URLResourceKey = "NSURLIsVolumeKey"
	URLIsWritableKey                                         URLResourceKey = "NSURLIsWritableKey"
	URLKeysOfUnsetValuesKey                                  URLResourceKey = "NSURLKeysOfUnsetValuesKey"
	URLLabelColorKey                                         URLResourceKey = "NSURLLabelColorKey"
	URLLabelNumberKey                                        URLResourceKey = "NSURLLabelNumberKey"
	URLLinkCountKey                                          URLResourceKey = "NSURLLinkCountKey"
	URLLocalizedLabelKey                                     URLResourceKey = "NSURLLocalizedLabelKey"
	URLLocalizedNameKey                                      URLResourceKey = "NSURLLocalizedNameKey"
	URLLocalizedTypeDescriptionKey                           URLResourceKey = "NSURLLocalizedTypeDescriptionKey"
	URLMayHaveExtendedAttributesKey                          URLResourceKey = "NSURLMayHaveExtendedAttributesKey"
	URLMayShareFileContentKey                                URLResourceKey = "NSURLMayShareFileContentKey"
	URLNameKey                                               URLResourceKey = "NSURLNameKey"
	URLParentDirectoryURLKey                                 URLResourceKey = "NSURLParentDirectoryURLKey"
	URLPathKey                                               URLResourceKey = "_NSURLPathKey"
	URLPreferredIOBlockSizeKey                               URLResourceKey = "NSURLPreferredIOBlockSizeKey"
	URLQuarantinePropertiesKey                               URLResourceKey = "NSURLQuarantinePropertiesKey"
	URLTagNamesKey                                           URLResourceKey = "NSURLTagNamesKey"
	URLThumbnailDictionaryKey                                URLResourceKey = "NSURLThumbnailDictionaryKey"
	URLThumbnailKey                                          URLResourceKey = "NSURLThumbnailKey"
	URLTotalFileAllocatedSizeKey                             URLResourceKey = "NSURLTotalFileAllocatedSizeKey"
	URLTotalFileSizeKey                                      URLResourceKey = "NSURLTotalFileSizeKey"
	URLTypeIdentifierKey                                     URLResourceKey = "NSURLTypeIdentifierKey"
	URLUbiquitousItemContainerDisplayNameKey                 URLResourceKey = "NSURLUbiquitousItemContainerDisplayNameKey"
	URLUbiquitousItemDownloadRequestedKey                    URLResourceKey = "NSURLUbiquitousItemDownloadRequestedKey"
	URLUbiquitousItemDownloadingErrorKey                     URLResourceKey = "NSURLUbiquitousItemDownloadingErrorKey"
	URLUbiquitousItemDownloadingStatusKey                    URLResourceKey = "NSURLUbiquitousItemDownloadingStatusKey"
	URLUbiquitousItemHasUnresolvedConflictsKey               URLResourceKey = "NSURLUbiquitousItemHasUnresolvedConflictsKey"
	URLUbiquitousItemIsDownloadedKey                         URLResourceKey = "NSURLUbiquitousItemIsDownloadedKey"
	URLUbiquitousItemIsDownloadingKey                        URLResourceKey = "NSURLUbiquitousItemIsDownloadingKey"
	URLUbiquitousItemIsExcludedFromSyncKey                   URLResourceKey = "NSURLUbiquitousItemIsExcludedFromSyncKey"
	URLUbiquitousItemIsSharedKey                             URLResourceKey = "NSURLUbiquitousItemIsSharedKey"
	URLUbiquitousItemIsUploadedKey                           URLResourceKey = "NSURLUbiquitousItemIsUploadedKey"
	URLUbiquitousItemIsUploadingKey                          URLResourceKey = "NSURLUbiquitousItemIsUploadingKey"
	URLUbiquitousItemPercentDownloadedKey                    URLResourceKey = "NSURLUbiquitousItemPercentDownloadedKey"
	URLUbiquitousItemPercentUploadedKey                      URLResourceKey = "NSURLUbiquitousItemPercentUploadedKey"
	URLUbiquitousItemUploadingErrorKey                       URLResourceKey = "NSURLUbiquitousItemUploadingErrorKey"
	URLUbiquitousSharedItemCurrentUserPermissionsKey         URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserPermissionsKey"
	URLUbiquitousSharedItemCurrentUserRoleKey                URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserRoleKey"
	URLUbiquitousSharedItemMostRecentEditorNameComponentsKey URLResourceKey = "NSURLUbiquitousSharedItemMostRecentEditorNameComponentsKey"
	URLUbiquitousSharedItemOwnerNameComponentsKey            URLResourceKey = "NSURLUbiquitousSharedItemOwnerNameComponentsKey"
	URLVolumeAvailableCapacityForImportantUsageKey           URLResourceKey = "NSURLVolumeAvailableCapacityForImportantUsageKey"
	URLVolumeAvailableCapacityForOpportunisticUsageKey       URLResourceKey = "NSURLVolumeAvailableCapacityForOpportunisticUsageKey"
	URLVolumeAvailableCapacityKey                            URLResourceKey = "NSURLVolumeAvailableCapacityKey"
	URLVolumeCreationDateKey                                 URLResourceKey = "NSURLVolumeCreationDateKey"
	URLVolumeIdentifierKey                                   URLResourceKey = "NSURLVolumeIdentifierKey"
	URLVolumeIsAutomountedKey                                URLResourceKey = "NSURLVolumeIsAutomountedKey"
	URLVolumeIsBrowsableKey                                  URLResourceKey = "NSURLVolumeIsBrowsableKey"
	URLVolumeIsEjectableKey                                  URLResourceKey = "NSURLVolumeIsEjectableKey"
	URLVolumeIsEncryptedKey                                  URLResourceKey = "NSURLVolumeIsEncryptedKey"
	URLVolumeIsInternalKey                                   URLResourceKey = "NSURLVolumeIsInternalKey"
	URLVolumeIsJournalingKey                                 URLResourceKey = "NSURLVolumeIsJournalingKey"
	URLVolumeIsLocalKey                                      URLResourceKey = "NSURLVolumeIsLocalKey"
	URLVolumeIsReadOnlyKey                                   URLResourceKey = "NSURLVolumeIsReadOnlyKey"
	URLVolumeIsRemovableKey                                  URLResourceKey = "NSURLVolumeIsRemovableKey"
	URLVolumeIsRootFileSystemKey                             URLResourceKey = "NSURLVolumeIsRootFileSystemKey"
	URLVolumeLocalizedFormatDescriptionKey                   URLResourceKey = "NSURLVolumeLocalizedFormatDescriptionKey"
	URLVolumeLocalizedNameKey                                URLResourceKey = "NSURLVolumeLocalizedNameKey"
	URLVolumeMaximumFileSizeKey                              URLResourceKey = "NSURLVolumeMaximumFileSizeKey"
	URLVolumeNameKey                                         URLResourceKey = "NSURLVolumeNameKey"
	URLVolumeResourceCountKey                                URLResourceKey = "NSURLVolumeResourceCountKey"
	URLVolumeSupportsAccessPermissionsKey                    URLResourceKey = "NSURLVolumeSupportsAccessPermissionsKey"
	URLVolumeSupportsAdvisoryFileLockingKey                  URLResourceKey = "NSURLVolumeSupportsAdvisoryFileLockingKey"
	URLVolumeSupportsCasePreservedNamesKey                   URLResourceKey = "NSURLVolumeSupportsCasePreservedNamesKey"
	URLVolumeSupportsCaseSensitiveNamesKey                   URLResourceKey = "NSURLVolumeSupportsCaseSensitiveNamesKey"
	URLVolumeSupportsCompressionKey                          URLResourceKey = "NSURLVolumeSupportsCompressionKey"
	URLVolumeSupportsExclusiveRenamingKey                    URLResourceKey = "NSURLVolumeSupportsExclusiveRenamingKey"
	URLVolumeSupportsExtendedSecurityKey                     URLResourceKey = "NSURLVolumeSupportsExtendedSecurityKey"
	URLVolumeSupportsFileCloningKey                          URLResourceKey = "NSURLVolumeSupportsFileCloningKey"
	URLVolumeSupportsFileProtectionKey                       URLResourceKey = "NSURLVolumeSupportsFileProtectionKey"
	URLVolumeSupportsHardLinksKey                            URLResourceKey = "NSURLVolumeSupportsHardLinksKey"
	URLVolumeSupportsImmutableFilesKey                       URLResourceKey = "NSURLVolumeSupportsImmutableFilesKey"
	URLVolumeSupportsJournalingKey                           URLResourceKey = "NSURLVolumeSupportsJournalingKey"
	URLVolumeSupportsPersistentIDsKey                        URLResourceKey = "NSURLVolumeSupportsPersistentIDsKey"
	URLVolumeSupportsRenamingKey                             URLResourceKey = "NSURLVolumeSupportsRenamingKey"
	URLVolumeSupportsRootDirectoryDatesKey                   URLResourceKey = "NSURLVolumeSupportsRootDirectoryDatesKey"
	URLVolumeSupportsSparseFilesKey                          URLResourceKey = "NSURLVolumeSupportsSparseFilesKey"
	URLVolumeSupportsSwapRenamingKey                         URLResourceKey = "NSURLVolumeSupportsSwapRenamingKey"
	URLVolumeSupportsSymbolicLinksKey                        URLResourceKey = "NSURLVolumeSupportsSymbolicLinksKey"
	URLVolumeSupportsVolumeSizesKey                          URLResourceKey = "NSURLVolumeSupportsVolumeSizesKey"
	URLVolumeSupportsZeroRunsKey                             URLResourceKey = "NSURLVolumeSupportsZeroRunsKey"
	URLVolumeTotalCapacityKey                                URLResourceKey = "NSURLVolumeTotalCapacityKey"
	URLVolumeURLForRemountingKey                             URLResourceKey = "NSURLVolumeURLForRemountingKey"
	URLVolumeURLKey                                          URLResourceKey = "NSURLVolumeURLKey"
	URLVolumeUUIDStringKey                                   URLResourceKey = "NSURLVolumeUUIDStringKey"
)

// Constants passed by session or task delegates to the provided continuation block in response to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionauthchallengedisposition?language=objc
type URLSessionAuthChallengeDisposition int

const (
	URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionAuthChallengeDisposition = 2
	URLSessionAuthChallengePerformDefaultHandling        URLSessionAuthChallengeDisposition = 1
	URLSessionAuthChallengeRejectProtectionSpace         URLSessionAuthChallengeDisposition = 3
	URLSessionAuthChallengeUseCredential                 URLSessionAuthChallengeDisposition = 0
)

// The action to take on a delayed URL session task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelayedrequestdisposition?language=objc
type URLSessionDelayedRequestDisposition int

const (
	URLSessionDelayedRequestCancel          URLSessionDelayedRequestDisposition = 2
	URLSessionDelayedRequestContinueLoading URLSessionDelayedRequestDisposition = 0
	URLSessionDelayedRequestUseNewRequest   URLSessionDelayedRequestDisposition = 1
)

// Constants indicating how a data or upload session should proceed after receiving the initial headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionresponsedisposition?language=objc
type URLSessionResponseDisposition int

const (
	URLSessionResponseAllow          URLSessionResponseDisposition = 1
	URLSessionResponseBecomeDownload URLSessionResponseDisposition = 2
	URLSessionResponseBecomeStream   URLSessionResponseDisposition = 3
	URLSessionResponseCancel         URLSessionResponseDisposition = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetricsdomainresolutionprotocol?language=objc
type URLSessionTaskMetricsDomainResolutionProtocol int

const (
	URLSessionTaskMetricsDomainResolutionProtocolHTTPS   URLSessionTaskMetricsDomainResolutionProtocol = 4
	URLSessionTaskMetricsDomainResolutionProtocolTCP     URLSessionTaskMetricsDomainResolutionProtocol = 2
	URLSessionTaskMetricsDomainResolutionProtocolTLS     URLSessionTaskMetricsDomainResolutionProtocol = 3
	URLSessionTaskMetricsDomainResolutionProtocolUDP     URLSessionTaskMetricsDomainResolutionProtocol = 1
	URLSessionTaskMetricsDomainResolutionProtocolUnknown URLSessionTaskMetricsDomainResolutionProtocol = 0
)

// The manner in which a resource is fetched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetricsresourcefetchtype?language=objc
type URLSessionTaskMetricsResourceFetchType int

const (
	URLSessionTaskMetricsResourceFetchTypeLocalCache  URLSessionTaskMetricsResourceFetchType = 3
	URLSessionTaskMetricsResourceFetchTypeNetworkLoad URLSessionTaskMetricsResourceFetchType = 1
	URLSessionTaskMetricsResourceFetchTypeServerPush  URLSessionTaskMetricsResourceFetchType = 2
	URLSessionTaskMetricsResourceFetchTypeUnknown     URLSessionTaskMetricsResourceFetchType = 0
)

// Constants for determining the current state of a task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskstate?language=objc
type URLSessionTaskState int

const (
	URLSessionTaskStateCanceling URLSessionTaskState = 2
	URLSessionTaskStateCompleted URLSessionTaskState = 3
	URLSessionTaskStateRunning   URLSessionTaskState = 0
	URLSessionTaskStateSuspended URLSessionTaskState = 1
)

// A code that indicates why a WebSocket connection closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketclosecode?language=objc
type URLSessionWebSocketCloseCode int

const (
	URLSessionWebSocketCloseCodeAbnormalClosure           URLSessionWebSocketCloseCode = 1006
	URLSessionWebSocketCloseCodeGoingAway                 URLSessionWebSocketCloseCode = 1001
	URLSessionWebSocketCloseCodeInternalServerError       URLSessionWebSocketCloseCode = 1011
	URLSessionWebSocketCloseCodeInvalid                   URLSessionWebSocketCloseCode = 0
	URLSessionWebSocketCloseCodeInvalidFramePayloadData   URLSessionWebSocketCloseCode = 1007
	URLSessionWebSocketCloseCodeMandatoryExtensionMissing URLSessionWebSocketCloseCode = 1010
	URLSessionWebSocketCloseCodeMessageTooBig             URLSessionWebSocketCloseCode = 1009
	URLSessionWebSocketCloseCodeNoStatusReceived          URLSessionWebSocketCloseCode = 1005
	URLSessionWebSocketCloseCodeNormalClosure             URLSessionWebSocketCloseCode = 1000
	URLSessionWebSocketCloseCodePolicyViolation           URLSessionWebSocketCloseCode = 1008
	URLSessionWebSocketCloseCodeProtocolError             URLSessionWebSocketCloseCode = 1002
	URLSessionWebSocketCloseCodeTLSHandshakeFailure       URLSessionWebSocketCloseCode = 1015
	URLSessionWebSocketCloseCodeUnsupportedData           URLSessionWebSocketCloseCode = 1003
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessagetype?language=objc
type URLSessionWebSocketMessageType int

const (
	URLSessionWebSocketMessageTypeData   URLSessionWebSocketMessageType = 0
	URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1
)

// Possible keys for the NSURLThumbnailDictionaryKey dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlthumbnaildictionaryitem?language=objc
type URLThumbnailDictionaryItem string

const (
	Thumbnail1024x1024SizeKey URLThumbnailDictionaryItem = "NSThumbnail1024x1024SizeKey"
)

// Values that describe the iCloud storage state of a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlubiquitousitemdownloadingstatus?language=objc
type URLUbiquitousItemDownloadingStatus string

const (
	URLUbiquitousItemDownloadingStatusCurrent       URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusCurrent"
	URLUbiquitousItemDownloadingStatusDownloaded    URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusDownloaded"
	URLUbiquitousItemDownloadingStatusNotDownloaded URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusNotDownloaded"
)

// The key for the permissions of a shared item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlubiquitousshareditempermissions?language=objc
type URLUbiquitousSharedItemPermissions string

const (
	URLUbiquitousSharedItemPermissionsReadOnly  URLUbiquitousSharedItemPermissions = "NSURLUbiquitousSharedItemPermissionsReadOnly"
	URLUbiquitousSharedItemPermissionsReadWrite URLUbiquitousSharedItemPermissions = "NSURLUbiquitousSharedItemPermissionsReadWrite"
)

// The key for the role of a shared item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlubiquitousshareditemrole?language=objc
type URLUbiquitousSharedItemRole string

const (
	URLUbiquitousSharedItemRoleOwner       URLUbiquitousSharedItemRole = "NSURLUbiquitousSharedItemRoleOwner"
	URLUbiquitousSharedItemRoleParticipant URLUbiquitousSharedItemRole = "NSURLUbiquitousSharedItemRoleParticipant"
)

// The type that defines a persistent identifier value for a user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitypersistentidentifier?language=objc
type UserActivityPersistentIdentifier string

// These constants describe how the user notification was activated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationactivationtype?language=objc
type UserNotificationActivationType int

const (
	UserNotificationActivationTypeActionButtonClicked     UserNotificationActivationType = 2
	UserNotificationActivationTypeAdditionalActionClicked UserNotificationActivationType = 4
	UserNotificationActivationTypeContentsClicked         UserNotificationActivationType = 1
	UserNotificationActivationTypeNone                    UserNotificationActivationType = 0
	UserNotificationActivationTypeReplied                 UserNotificationActivationType = 3
)

// Named value transformers defined by NSValueTransformer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvaluetransformername?language=objc
type ValueTransformerName string

const (
	IsNilTransformerName                   ValueTransformerName = "NSIsNil"
	IsNotNilTransformerName                ValueTransformerName = "NSIsNotNil"
	KeyedUnarchiveFromDataTransformerName  ValueTransformerName = "NSKeyedUnarchiveFromData"
	NegateBooleanTransformerName           ValueTransformerName = "NSNegateBoolean"
	SecureUnarchiveFromDataTransformerName ValueTransformerName = "NSSecureUnarchiveFromData"
	UnarchiveFromDataTransformerName       ValueTransformerName = "NSUnarchiveFromData"
)

// Options for enumerating mounted volumes with the mountedVolumeURLsIncludingResourceValuesForKeys:options: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvolumeenumerationoptions?language=objc
type VolumeEnumerationOptions uint

const (
	VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 4
	VolumeEnumerationSkipHiddenVolumes        VolumeEnumerationOptions = 2
)

// NSWhoseSpecifier uses these constants to specify sub-elements within the collection of objects being tested that pass the specifier’s test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nswhosesubelementidentifier?language=objc
type WhoseSubelementIdentifier uint

const (
	EverySubelement  WhoseSubelementIdentifier = 1
	IndexSubelement  WhoseSubelementIdentifier = 0
	MiddleSubelement WhoseSubelementIdentifier = 2
	NoSubelement     WhoseSubelementIdentifier = 4
	RandomSubelement WhoseSubelementIdentifier = 3
)

// The type defined for the constants that specify the kind and subkind of DTD declaration represented by an NSXMLDTDNode object. You set the DTD-node kind using the setDTDKind: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnodekind?language=objc
type XMLDTDNodeKind uint

const (
	XMLAttributeCDATAKind              XMLDTDNodeKind = 6
	XMLAttributeEntitiesKind           XMLDTDNodeKind = 11
	XMLAttributeEntityKind             XMLDTDNodeKind = 10
	XMLAttributeEnumerationKind        XMLDTDNodeKind = 14
	XMLAttributeIDKind                 XMLDTDNodeKind = 7
	XMLAttributeIDRefKind              XMLDTDNodeKind = 8
	XMLAttributeIDRefsKind             XMLDTDNodeKind = 9
	XMLAttributeNMTokenKind            XMLDTDNodeKind = 12
	XMLAttributeNMTokensKind           XMLDTDNodeKind = 13
	XMLAttributeNotationKind           XMLDTDNodeKind = 15
	XMLElementDeclarationAnyKind       XMLDTDNodeKind = 18
	XMLElementDeclarationElementKind   XMLDTDNodeKind = 20
	XMLElementDeclarationEmptyKind     XMLDTDNodeKind = 17
	XMLElementDeclarationMixedKind     XMLDTDNodeKind = 19
	XMLElementDeclarationUndefinedKind XMLDTDNodeKind = 16
	XMLEntityGeneralKind               XMLDTDNodeKind = 1
	XMLEntityParameterKind             XMLDTDNodeKind = 4
	XMLEntityParsedKind                XMLDTDNodeKind = 2
	XMLEntityPredefined                XMLDTDNodeKind = 5
	XMLEntityUnparsedKind              XMLDTDNodeKind = 3
)

// Type used to define the kind of document content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocumentcontentkind?language=objc
type XMLDocumentContentKind uint

const (
	XMLDocumentHTMLKind  XMLDocumentContentKind = 2
	XMLDocumentTextKind  XMLDocumentContentKind = 3
	XMLDocumentXHTMLKind XMLDocumentContentKind = 1
	XMLDocumentXMLKind   XMLDocumentContentKind = 0
)

// NSXMLNode declares the following constants of type NSXMLNodeKind for specifying a node’s kind in the initializer methods initWithKind:options: and initWithKind:options:: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnodekind?language=objc
type XMLNodeKind uint

const (
	XMLAttributeDeclarationKind  XMLNodeKind = 10
	XMLAttributeKind             XMLNodeKind = 3
	XMLCommentKind               XMLNodeKind = 6
	XMLDTDKind                   XMLNodeKind = 8
	XMLDocumentKind              XMLNodeKind = 1
	XMLElementDeclarationKind    XMLNodeKind = 11
	XMLElementKind               XMLNodeKind = 2
	XMLEntityDeclarationKind     XMLNodeKind = 9
	XMLInvalidKind               XMLNodeKind = 0
	XMLNamespaceKind             XMLNodeKind = 4
	XMLNotationDeclarationKind   XMLNodeKind = 12
	XMLProcessingInstructionKind XMLNodeKind = 5
	XMLTextKind                  XMLNodeKind = 7
)

// These constants are input and output options for all NSXMLNode objects (unless otherwise indicated), including NSXMLDocument objects. You can specify these options in the NSXMLNode methods initWithKind:options: and XMLStringWithOptions:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnodeoptions?language=objc
type XMLNodeOptions uint

const (
	XMLDocumentIncludeContentTypeDeclaration  XMLNodeOptions = 262144
	XMLDocumentTidyHTML                       XMLNodeOptions = 512
	XMLDocumentTidyXML                        XMLNodeOptions = 1024
	XMLDocumentValidate                       XMLNodeOptions = 8192
	XMLDocumentXInclude                       XMLNodeOptions = 65536
	XMLNodeCompactEmptyElement                XMLNodeOptions = 4
	XMLNodeExpandEmptyElement                 XMLNodeOptions = 2
	XMLNodeIsCDATA                            XMLNodeOptions = 1
	XMLNodeLoadExternalEntitiesAlways         XMLNodeOptions = 16384
	XMLNodeLoadExternalEntitiesNever          XMLNodeOptions = 524288
	XMLNodeLoadExternalEntitiesSameOriginOnly XMLNodeOptions = 32768
	XMLNodeNeverEscapeContents                XMLNodeOptions = 32
	XMLNodeOptionsNone                        XMLNodeOptions = 0
	XMLNodePreserveAll                        XMLNodeOptions = 4293918750
	XMLNodePreserveAttributeOrder             XMLNodeOptions = 2097152
	XMLNodePreserveCDATA                      XMLNodeOptions = 16777216
	XMLNodePreserveCharacterReferences        XMLNodeOptions = 134217728
	XMLNodePreserveDTD                        XMLNodeOptions = 67108864
	XMLNodePreserveEmptyElements              XMLNodeOptions = 6
	XMLNodePreserveEntities                   XMLNodeOptions = 4194304
	XMLNodePreserveNamespaceOrder             XMLNodeOptions = 1048576
	XMLNodePreservePrefixes                   XMLNodeOptions = 8388608
	XMLNodePreserveQuotes                     XMLNodeOptions = 24
	XMLNodePreserveWhitespace                 XMLNodeOptions = 33554432
	XMLNodePrettyPrint                        XMLNodeOptions = 131072
	XMLNodePromoteSignificantWhitespace       XMLNodeOptions = 268435456
	XMLNodeUseDoubleQuotes                    XMLNodeOptions = 16
	XMLNodeUseSingleQuotes                    XMLNodeOptions = 8
)

// The following error codes are defined by NSXMLParser. For error codes not listed here, see the <libxml/xmlerror.h> header file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparsererror?language=objc
type XMLParserError int

const (
	XMLParserAttributeHasNoValueError              XMLParserError = 41
	XMLParserAttributeListNotFinishedError         XMLParserError = 51
	XMLParserAttributeListNotStartedError          XMLParserError = 50
	XMLParserAttributeNotFinishedError             XMLParserError = 40
	XMLParserAttributeNotStartedError              XMLParserError = 39
	XMLParserAttributeRedefinedError               XMLParserError = 42
	XMLParserCDATANotFinishedError                 XMLParserError = 63
	XMLParserCharacterRefAtEOFError                XMLParserError = 10
	XMLParserCharacterRefInDTDError                XMLParserError = 13
	XMLParserCharacterRefInEpilogError             XMLParserError = 12
	XMLParserCharacterRefInPrologError             XMLParserError = 11
	XMLParserCommentContainsDoubleHyphenError      XMLParserError = 80
	XMLParserCommentNotFinishedError               XMLParserError = 45
	XMLParserConditionalSectionNotFinishedError    XMLParserError = 59
	XMLParserConditionalSectionNotStartedError     XMLParserError = 58
	XMLParserDOCTYPEDeclNotFinishedError           XMLParserError = 61
	XMLParserDelegateAbortedParseError             XMLParserError = 512
	XMLParserDocumentStartError                    XMLParserError = 3
	XMLParserElementContentDeclNotFinishedError    XMLParserError = 55
	XMLParserElementContentDeclNotStartedError     XMLParserError = 54
	XMLParserEmptyDocumentError                    XMLParserError = 4
	XMLParserEncodingNotSupportedError             XMLParserError = 32
	XMLParserEntityBoundaryError                   XMLParserError = 90
	XMLParserEntityIsExternalError                 XMLParserError = 29
	XMLParserEntityIsParameterError                XMLParserError = 30
	XMLParserEntityNotFinishedError                XMLParserError = 37
	XMLParserEntityNotStartedError                 XMLParserError = 36
	XMLParserEntityRefAtEOFError                   XMLParserError = 14
	XMLParserEntityRefInDTDError                   XMLParserError = 17
	XMLParserEntityRefInEpilogError                XMLParserError = 16
	XMLParserEntityRefInPrologError                XMLParserError = 15
	XMLParserEntityRefLoopError                    XMLParserError = 89
	XMLParserEntityReferenceMissingSemiError       XMLParserError = 23
	XMLParserEntityReferenceWithoutNameError       XMLParserError = 22
	XMLParserEntityValueRequiredError              XMLParserError = 84
	XMLParserEqualExpectedError                    XMLParserError = 75
	XMLParserExternalStandaloneEntityError         XMLParserError = 82
	XMLParserExternalSubsetNotFinishedError        XMLParserError = 60
	XMLParserExtraContentError                     XMLParserError = 86
	XMLParserGTRequiredError                       XMLParserError = 73
	XMLParserInternalError                         XMLParserError = 1
	XMLParserInvalidCharacterError                 XMLParserError = 9
	XMLParserInvalidCharacterInEntityError         XMLParserError = 87
	XMLParserInvalidCharacterRefError              XMLParserError = 8
	XMLParserInvalidConditionalSectionError        XMLParserError = 83
	XMLParserInvalidDecimalCharacterRefError       XMLParserError = 7
	XMLParserInvalidEncodingError                  XMLParserError = 81
	XMLParserInvalidEncodingNameError              XMLParserError = 79
	XMLParserInvalidHexCharacterRefError           XMLParserError = 6
	XMLParserInvalidURIError                       XMLParserError = 91
	XMLParserLTRequiredError                       XMLParserError = 72
	XMLParserLTSlashRequiredError                  XMLParserError = 74
	XMLParserLessThanSymbolInAttributeError        XMLParserError = 38
	XMLParserLiteralNotFinishedError               XMLParserError = 44
	XMLParserLiteralNotStartedError                XMLParserError = 43
	XMLParserMisplacedCDATAEndStringError          XMLParserError = 62
	XMLParserMisplacedXMLDeclarationError          XMLParserError = 64
	XMLParserMixedContentDeclNotFinishedError      XMLParserError = 53
	XMLParserMixedContentDeclNotStartedError       XMLParserError = 52
	XMLParserNAMERequiredError                     XMLParserError = 68
	XMLParserNMTOKENRequiredError                  XMLParserError = 67
	XMLParserNamespaceDeclarationError             XMLParserError = 35
	XMLParserNoDTDError                            XMLParserError = 94
	XMLParserNotWellBalancedError                  XMLParserError = 85
	XMLParserNotationNotFinishedError              XMLParserError = 49
	XMLParserNotationNotStartedError               XMLParserError = 48
	XMLParserOutOfMemoryError                      XMLParserError = 2
	XMLParserPCDATARequiredError                   XMLParserError = 69
	XMLParserParsedEntityRefAtEOFError             XMLParserError = 18
	XMLParserParsedEntityRefInEpilogError          XMLParserError = 20
	XMLParserParsedEntityRefInInternalError        XMLParserError = 88
	XMLParserParsedEntityRefInInternalSubsetError  XMLParserError = 21
	XMLParserParsedEntityRefInPrologError          XMLParserError = 19
	XMLParserParsedEntityRefMissingSemiError       XMLParserError = 25
	XMLParserParsedEntityRefNoNameError            XMLParserError = 24
	XMLParserPrematureDocumentEndError             XMLParserError = 5
	XMLParserProcessingInstructionNotFinishedError XMLParserError = 47
	XMLParserProcessingInstructionNotStartedError  XMLParserError = 46
	XMLParserPublicIdentifierRequiredError         XMLParserError = 71
	XMLParserSeparatorRequiredError                XMLParserError = 66
	XMLParserSpaceRequiredError                    XMLParserError = 65
	XMLParserStandaloneValueError                  XMLParserError = 78
	XMLParserStringNotClosedError                  XMLParserError = 34
	XMLParserStringNotStartedError                 XMLParserError = 33
	XMLParserTagNameMismatchError                  XMLParserError = 76
	XMLParserURIFragmentError                      XMLParserError = 92
	XMLParserURIRequiredError                      XMLParserError = 70
	XMLParserUndeclaredEntityError                 XMLParserError = 26
	XMLParserUnfinishedTagError                    XMLParserError = 77
	XMLParserUnknownEncodingError                  XMLParserError = 31
	XMLParserUnparsedEntityError                   XMLParserError = 28
	XMLParserXMLDeclNotFinishedError               XMLParserError = 57
	XMLParserXMLDeclNotStartedError                XMLParserError = 56
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserexternalentityresolvingpolicy?language=objc
type XMLParserExternalEntityResolvingPolicy uint

const (
	XMLParserResolveExternalEntitiesAlways         XMLParserExternalEntityResolvingPolicy = 3
	XMLParserResolveExternalEntitiesNever          XMLParserExternalEntityResolvingPolicy = 0
	XMLParserResolveExternalEntitiesNoNetwork      XMLParserExternalEntityResolvingPolicy = 1
	XMLParserResolveExternalEntitiesSameOriginOnly XMLParserExternalEntityResolvingPolicy = 2
)

// Options that you can pass to a connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionoptions?language=objc
type XPCConnectionOptions uint

const (
	XPCConnectionPrivileged XPCConnectionOptions = 4096
)

// Type for UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/unichar?language=objc
type Unichar int
