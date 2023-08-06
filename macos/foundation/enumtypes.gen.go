// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import "math"

type AlignmentOptions uint64

const AlignMinXInward AlignmentOptions = 1
const AlignMinYInward AlignmentOptions = 2
const AlignMaxXInward AlignmentOptions = 4
const AlignMaxYInward AlignmentOptions = 8
const AlignWidthInward AlignmentOptions = 16
const AlignHeightInward AlignmentOptions = 32
const AlignMinXOutward AlignmentOptions = 256
const AlignMinYOutward AlignmentOptions = 512
const AlignMaxXOutward AlignmentOptions = 1024
const AlignMaxYOutward AlignmentOptions = 2048
const AlignWidthOutward AlignmentOptions = 4096
const AlignHeightOutward AlignmentOptions = 8192
const AlignMinXNearest AlignmentOptions = 65536
const AlignMinYNearest AlignmentOptions = 131072
const AlignMaxXNearest AlignmentOptions = 262144
const AlignMaxYNearest AlignmentOptions = 524288
const AlignWidthNearest AlignmentOptions = 1048576
const AlignHeightNearest AlignmentOptions = 2097152
const AlignRectFlipped AlignmentOptions = 9223372036854775808
const AlignAllEdgesInward AlignmentOptions = 15
const AlignAllEdgesOutward AlignmentOptions = 3840
const AlignAllEdgesNearest AlignmentOptions = 983040

type AppleEventSendOptions uint

const AppleEventSendAlwaysInteract AppleEventSendOptions = 48
const AppleEventSendCanInteract AppleEventSendOptions = 32
const AppleEventSendCanSwitchLayer AppleEventSendOptions = 64
const AppleEventSendDefaultOptions AppleEventSendOptions = 35
const AppleEventSendDontAnnotate AppleEventSendOptions = 65536
const AppleEventSendDontExecute AppleEventSendOptions = 8192
const AppleEventSendDontRecord AppleEventSendOptions = 4096
const AppleEventSendNeverInteract AppleEventSendOptions = 16
const AppleEventSendNoReply AppleEventSendOptions = 1
const AppleEventSendQueueReply AppleEventSendOptions = 2
const AppleEventSendWaitForReply AppleEventSendOptions = 3

type AttributedStringEnumerationOptions uint

const AttributedStringEnumerationReverse AttributedStringEnumerationOptions = 2
const AttributedStringEnumerationLongestEffectiveRangeNotRequired AttributedStringEnumerationOptions = 1048576

type AttributedStringFormattingOptions uint

const AttributedStringFormattingApplyReplacementIndexAttribute AttributedStringFormattingOptions = 2
const AttributedStringFormattingInsertArgumentAttributesWithoutMerging AttributedStringFormattingOptions = 1

type AttributedStringKey string

const MarkdownSourcePositionAttributeName AttributedStringKey = "NSMarkdownSourcePosition"
const AlternateDescriptionAttributeName AttributedStringKey = "NSAlternateDescription"
const ImageURLAttributeName AttributedStringKey = "NSImageURL"
const InflectionAlternativeAttributeName AttributedStringKey = "NSInflectionAlternative"
const InflectionRuleAttributeName AttributedStringKey = "NSInflect"
const InlinePresentationIntentAttributeName AttributedStringKey = "NSInlinePresentationIntent"
const LanguageIdentifierAttributeName AttributedStringKey = "NSLanguage"
const MorphologyAttributeName AttributedStringKey = "NSMorphology"
const PresentationIntentAttributeName AttributedStringKey = "NSPresentationIntent"
const ReplacementIndexAttributeName AttributedStringKey = "NSReplacementIndex"

type ByteCountFormatterCountStyle int

const ByteCountFormatterCountStyleFile ByteCountFormatterCountStyle = 0
const ByteCountFormatterCountStyleMemory ByteCountFormatterCountStyle = 1
const ByteCountFormatterCountStyleDecimal ByteCountFormatterCountStyle = 2
const ByteCountFormatterCountStyleBinary ByteCountFormatterCountStyle = 3

type ByteCountFormatterUnits uint

const ByteCountFormatterUseDefault ByteCountFormatterUnits = 0
const ByteCountFormatterUseBytes ByteCountFormatterUnits = 1
const ByteCountFormatterUseKB ByteCountFormatterUnits = 2
const ByteCountFormatterUseMB ByteCountFormatterUnits = 4
const ByteCountFormatterUseGB ByteCountFormatterUnits = 8
const ByteCountFormatterUseTB ByteCountFormatterUnits = 16
const ByteCountFormatterUsePB ByteCountFormatterUnits = 32
const ByteCountFormatterUseEB ByteCountFormatterUnits = 64
const ByteCountFormatterUseZB ByteCountFormatterUnits = 128
const ByteCountFormatterUseYBOrHigher ByteCountFormatterUnits = 65280
const ByteCountFormatterUseAll ByteCountFormatterUnits = 65535

type CalendarIdentifier string

const CalendarIdentifierGregorian CalendarIdentifier = "gregorian"
const CalendarIdentifierISO8601 CalendarIdentifier = "iso8601"
const CalendarIdentifierBuddhist CalendarIdentifier = "buddhist"
const CalendarIdentifierChinese CalendarIdentifier = "chinese"
const CalendarIdentifierCoptic CalendarIdentifier = "coptic"
const CalendarIdentifierEthiopicAmeteAlem CalendarIdentifier = "ethiopic-amete-alem"
const CalendarIdentifierEthiopicAmeteMihret CalendarIdentifier = "ethiopic"
const CalendarIdentifierHebrew CalendarIdentifier = "hebrew"
const CalendarIdentifierIndian CalendarIdentifier = "indian"
const CalendarIdentifierIslamic CalendarIdentifier = "islamic"
const CalendarIdentifierIslamicCivil CalendarIdentifier = "islamic-civil"
const CalendarIdentifierIslamicTabular CalendarIdentifier = "islamic-tbla"
const CalendarIdentifierIslamicUmmAlQura CalendarIdentifier = "islamic-umalqura"
const CalendarIdentifierJapanese CalendarIdentifier = "japanese"
const CalendarIdentifierPersian CalendarIdentifier = "persian"
const CalendarIdentifierRepublicOfChina CalendarIdentifier = "roc"

type CalendarOptions uint

const CalendarWrapComponents CalendarOptions = 1
const CalendarMatchStrictly CalendarOptions = 2
const CalendarSearchBackwards CalendarOptions = 4
const CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 256
const CalendarMatchNextTimePreservingSmallerUnits CalendarOptions = 512
const CalendarMatchNextTime CalendarOptions = 1024
const CalendarMatchFirst CalendarOptions = 4096
const CalendarMatchLast CalendarOptions = 8192

type CalendarUnit uint

const CalendarUnitEra CalendarUnit = 2
const CalendarUnitYear CalendarUnit = 4
const CalendarUnitYearForWeekOfYear CalendarUnit = 16384
const CalendarUnitQuarter CalendarUnit = 2048
const CalendarUnitMonth CalendarUnit = 8
const CalendarUnitWeekOfYear CalendarUnit = 8192
const CalendarUnitWeekOfMonth CalendarUnit = 4096
const CalendarUnitWeekday CalendarUnit = 512
const CalendarUnitWeekdayOrdinal CalendarUnit = 1024
const CalendarUnitDay CalendarUnit = 16
const CalendarUnitHour CalendarUnit = 32
const CalendarUnitMinute CalendarUnit = 64
const CalendarUnitSecond CalendarUnit = 128
const CalendarUnitNanosecond CalendarUnit = 32768
const CalendarUnitCalendar CalendarUnit = 1048576
const CalendarUnitTimeZone CalendarUnit = 2097152

type ComparisonPredicateModifier uint

const DirectPredicateModifier ComparisonPredicateModifier = 0
const AllPredicateModifier ComparisonPredicateModifier = 1
const AnyPredicateModifier ComparisonPredicateModifier = 2

type ComparisonResult int

const OrderedAscending ComparisonResult = -1
const OrderedSame ComparisonResult = 0
const OrderedDescending ComparisonResult = 1

type DateComponentsFormatterUnitsStyle int

const DateComponentsFormatterUnitsStyleSpellOut DateComponentsFormatterUnitsStyle = 4
const DateComponentsFormatterUnitsStyleFull DateComponentsFormatterUnitsStyle = 3
const DateComponentsFormatterUnitsStyleShort DateComponentsFormatterUnitsStyle = 2
const DateComponentsFormatterUnitsStyleBrief DateComponentsFormatterUnitsStyle = 5
const DateComponentsFormatterUnitsStyleAbbreviated DateComponentsFormatterUnitsStyle = 1
const DateComponentsFormatterUnitsStylePositional DateComponentsFormatterUnitsStyle = 0

type DateComponentsFormatterZeroFormattingBehavior uint

const DateComponentsFormatterZeroFormattingBehaviorNone DateComponentsFormatterZeroFormattingBehavior = 0
const DateComponentsFormatterZeroFormattingBehaviorDefault DateComponentsFormatterZeroFormattingBehavior = 1
const DateComponentsFormatterZeroFormattingBehaviorDropLeading DateComponentsFormatterZeroFormattingBehavior = 2
const DateComponentsFormatterZeroFormattingBehaviorDropMiddle DateComponentsFormatterZeroFormattingBehavior = 4
const DateComponentsFormatterZeroFormattingBehaviorDropTrailing DateComponentsFormatterZeroFormattingBehavior = 8
const DateComponentsFormatterZeroFormattingBehaviorDropAll DateComponentsFormatterZeroFormattingBehavior = 14
const DateComponentsFormatterZeroFormattingBehaviorPad DateComponentsFormatterZeroFormattingBehavior = 65536

type DateFormatterBehavior uint

const DateFormatterBehaviorDefault DateFormatterBehavior = 0
const DateFormatterBehavior10_0 DateFormatterBehavior = 1000
const DateFormatterBehavior10_4 DateFormatterBehavior = 1040

type DateFormatterStyle uint

const DateFormatterNoStyle DateFormatterStyle = 0
const DateFormatterShortStyle DateFormatterStyle = 1
const DateFormatterMediumStyle DateFormatterStyle = 2
const DateFormatterLongStyle DateFormatterStyle = 3
const DateFormatterFullStyle DateFormatterStyle = 4

type DateIntervalFormatterStyle uint

const DateIntervalFormatterNoStyle DateIntervalFormatterStyle = 0
const DateIntervalFormatterShortStyle DateIntervalFormatterStyle = 1
const DateIntervalFormatterMediumStyle DateIntervalFormatterStyle = 2
const DateIntervalFormatterLongStyle DateIntervalFormatterStyle = 3
const DateIntervalFormatterFullStyle DateIntervalFormatterStyle = 4

type EnumerationOptions uint

const EnumerationConcurrent EnumerationOptions = 1
const EnumerationReverse EnumerationOptions = 2

type ErrorDomain string

type ErrorUserInfoKey string

const URLErrorKey ErrorUserInfoKey = "NSURL"
const FilePathErrorKey ErrorUserInfoKey = "NSFilePath"
const HelpAnchorErrorKey ErrorUserInfoKey = "NSHelpAnchor"
const LocalizedDescriptionKey ErrorUserInfoKey = "NSLocalizedDescription"
const LocalizedFailureErrorKey ErrorUserInfoKey = "NSLocalizedFailure"
const LocalizedFailureReasonErrorKey ErrorUserInfoKey = "NSLocalizedFailureReason"
const LocalizedRecoveryOptionsErrorKey ErrorUserInfoKey = "NSLocalizedRecoveryOptions"
const LocalizedRecoverySuggestionErrorKey ErrorUserInfoKey = "NSLocalizedRecoverySuggestion"
const RecoveryAttempterErrorKey ErrorUserInfoKey = "NSRecoveryAttempter"
const StringEncodingErrorKey ErrorUserInfoKey = "NSStringEncoding"
const UnderlyingErrorKey ErrorUserInfoKey = "NSUnderlyingError"
const DebugDescriptionErrorKey ErrorUserInfoKey = "NSDebugDescription"
const MultipleUnderlyingErrorsKey ErrorUserInfoKey = "NSMultipleUnderlyingErrorsKey"

type ExceptionName string

const CharacterConversionException ExceptionName = "NSCharacterConversionException"
const DecimalNumberDivideByZeroException ExceptionName = "NSDecimalNumberDivideByZeroException"
const DecimalNumberExactnessException ExceptionName = "NSDecimalNumberExactnessException"
const DecimalNumberOverflowException ExceptionName = "NSDecimalNumberOverflowException"
const DecimalNumberUnderflowException ExceptionName = "NSDecimalNumberUnderflowException"
const DestinationInvalidException ExceptionName = "NSDestinationInvalidException"
const FileHandleOperationException ExceptionName = "NSFileHandleOperationException"
const GenericException ExceptionName = "NSGenericException"
const InternalInconsistencyException ExceptionName = "NSInternalInconsistencyException"
const InvalidArchiveOperationException ExceptionName = "NSInvalidArchiveOperationException"
const InvalidArgumentException ExceptionName = "NSInvalidArgumentException"
const InvalidReceivePortException ExceptionName = "NSInvalidReceivePortException"
const InvalidSendPortException ExceptionName = "NSInvalidSendPortException"
const InvalidUnarchiveOperationException ExceptionName = "NSInvalidUnarchiveOperationException"
const InvocationOperationCancelledException ExceptionName = "NSInvocationOperationCancelledException"
const InvocationOperationVoidResultException ExceptionName = "NSInvocationOperationVoidResultException"
const MallocException ExceptionName = "NSMallocException"
const ObjectInaccessibleException ExceptionName = "NSObjectInaccessibleException"
const ObjectNotAvailableException ExceptionName = "NSObjectNotAvailableException"
const OldStyleException ExceptionName = "NSOldStyleException"
const ParseErrorException ExceptionName = "NSParseErrorException"
const PortReceiveException ExceptionName = "NSPortReceiveException"
const PortSendException ExceptionName = "NSPortSendException"
const PortTimeoutException ExceptionName = "NSPortTimeoutException"
const RangeException ExceptionName = "NSRangeException"
const UndefinedKeyException ExceptionName = "NSUnknownKeyException"
const InconsistentArchiveException ExceptionName = "NSArchiverArchiveInconsistency"

type ExpressionType uint

const ConstantValueExpressionType ExpressionType = 0
const EvaluatedObjectExpressionType ExpressionType = 1
const VariableExpressionType ExpressionType = 2
const KeyPathExpressionType ExpressionType = 3
const FunctionExpressionType ExpressionType = 4
const UnionSetExpressionType ExpressionType = 5
const IntersectSetExpressionType ExpressionType = 6
const MinusSetExpressionType ExpressionType = 7
const SubqueryExpressionType ExpressionType = 13
const AggregateExpressionType ExpressionType = 14
const AnyKeyExpressionType ExpressionType = 15
const BlockExpressionType ExpressionType = 19
const ConditionalExpressionType ExpressionType = 20

type FileWrapperReadingOptions uint

const FileWrapperReadingImmediate FileWrapperReadingOptions = 1
const FileWrapperReadingWithoutMapping FileWrapperReadingOptions = 2

type FileWrapperWritingOptions uint

const FileWrapperWritingAtomic FileWrapperWritingOptions = 1
const FileWrapperWritingWithNameUpdating FileWrapperWritingOptions = 2

type FormattingContext int

const FormattingContextUnknown FormattingContext = 0
const FormattingContextDynamic FormattingContext = 1
const FormattingContextStandalone FormattingContext = 2
const FormattingContextListItem FormattingContext = 3
const FormattingContextBeginningOfSentence FormattingContext = 4
const FormattingContextMiddleOfSentence FormattingContext = 5

type FormattingUnitStyle int

const FormattingUnitStyleShort FormattingUnitStyle = 1
const FormattingUnitStyleMedium FormattingUnitStyle = 2
const FormattingUnitStyleLong FormattingUnitStyle = 3

type HTTPCookiePropertyKey string

const HTTPCookieComment HTTPCookiePropertyKey = "Comment"
const HTTPCookieCommentURL HTTPCookiePropertyKey = "CommentURL"
const HTTPCookieDiscard HTTPCookiePropertyKey = "Discard"
const HTTPCookieDomain HTTPCookiePropertyKey = "Domain"
const HTTPCookieExpires HTTPCookiePropertyKey = "Expires"
const HTTPCookieMaximumAge HTTPCookiePropertyKey = "Max-Age"
const HTTPCookieName HTTPCookiePropertyKey = "Name"
const HTTPCookieOriginURL HTTPCookiePropertyKey = "OriginURL"
const HTTPCookiePath HTTPCookiePropertyKey = "Path"
const HTTPCookiePort HTTPCookiePropertyKey = "Port"
const HTTPCookieSameSitePolicy HTTPCookiePropertyKey = "SameSite"
const HTTPCookieSecure HTTPCookiePropertyKey = "Secure"
const HTTPCookieValue HTTPCookiePropertyKey = "Value"
const HTTPCookieVersion HTTPCookiePropertyKey = "Version"

type HTTPCookieStringPolicy string

const HTTPCookieSameSiteStrict HTTPCookieStringPolicy = "strict"
const HTTPCookieSameSiteLax HTTPCookieStringPolicy = "lax"

type LocaleKey string

const LocaleIdentifier LocaleKey = "kCFLocaleIdentifierKey"
const LocaleCountryCode LocaleKey = "kCFLocaleCountryCodeKey"
const LocaleLanguageCode LocaleKey = "kCFLocaleLanguageCodeKey"
const LocaleScriptCode LocaleKey = "kCFLocaleScriptCodeKey"
const LocaleVariantCode LocaleKey = "kCFLocaleVariantCodeKey"
const LocaleExemplarCharacterSet LocaleKey = "kCFLocaleExemplarCharacterSetKey"
const LocaleCalendar LocaleKey = "kCFLocaleCalendarKey"
const LocaleCollationIdentifier LocaleKey = "collation"
const LocaleCollatorIdentifier LocaleKey = "kCFLocaleCollatorIdentifierKey"
const LocaleUsesMetricSystem LocaleKey = "kCFLocaleUsesMetricSystemKey"
const LocaleMeasurementSystem LocaleKey = "kCFLocaleMeasurementSystemKey"
const LocaleDecimalSeparator LocaleKey = "kCFLocaleDecimalSeparatorKey"
const LocaleGroupingSeparator LocaleKey = "kCFLocaleGroupingSeparatorKey"
const LocaleCurrencySymbol LocaleKey = "kCFLocaleCurrencySymbolKey"
const LocaleCurrencyCode LocaleKey = "currency"
const LocaleQuotationEndDelimiterKey LocaleKey = "kCFLocaleQuotationEndDelimiterKey"
const LocaleQuotationBeginDelimiterKey LocaleKey = "kCFLocaleQuotationBeginDelimiterKey"
const LocaleAlternateQuotationEndDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationEndDelimiterKey"
const LocaleAlternateQuotationBeginDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationBeginDelimiterKey"

type LocaleLanguageDirection uint

const LocaleLanguageDirectionUnknown LocaleLanguageDirection = 0
const LocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 1
const LocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 2
const LocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 3
const LocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 4

type MatchingFlags uint

const MatchingProgress MatchingFlags = 1
const MatchingCompleted MatchingFlags = 2
const MatchingHitEnd MatchingFlags = 4
const MatchingRequiredEnd MatchingFlags = 8
const MatchingInternalError MatchingFlags = 16

type MatchingOptions uint

const MatchingReportProgress MatchingOptions = 1
const MatchingReportCompletion MatchingOptions = 2
const MatchingAnchored MatchingOptions = 4
const MatchingWithTransparentBounds MatchingOptions = 8
const MatchingWithoutAnchoringBounds MatchingOptions = 16

type MeasurementFormatterUnitOptions uint

const MeasurementFormatterUnitOptionsProvidedUnit MeasurementFormatterUnitOptions = 1
const MeasurementFormatterUnitOptionsNaturalScale MeasurementFormatterUnitOptions = 2
const MeasurementFormatterUnitOptionsTemperatureWithoutUnit MeasurementFormatterUnitOptions = 4

type NotificationName string

const ClassDescriptionNeededForClassNotification NotificationName = "NSClassDescriptionNeededForClassNotification"
const UbiquityIdentityDidChangeNotification NotificationName = "NSUbiquityIdentityDidChangeNotification"
const AppleEventManagerWillProcessFirstEventNotification NotificationName = "NSAppleEventManagerWillProcessFirstEvent"
const UndoManagerCheckpointNotification NotificationName = "NSUndoManagerCheckpointNotification"
const UndoManagerDidCloseUndoGroupNotification NotificationName = "NSUndoManagerDidCloseUndoGroupNotification"
const UndoManagerDidOpenUndoGroupNotification NotificationName = "NSUndoManagerDidOpenUndoGroupNotification"
const UndoManagerDidRedoChangeNotification NotificationName = "NSUndoManagerDidRedoChangeNotification"
const UndoManagerDidUndoChangeNotification NotificationName = "NSUndoManagerDidUndoChangeNotification"
const UndoManagerWillCloseUndoGroupNotification NotificationName = "NSUndoManagerWillCloseUndoGroupNotification"
const UndoManagerWillRedoChangeNotification NotificationName = "NSUndoManagerWillRedoChangeNotification"
const UndoManagerWillUndoChangeNotification NotificationName = "NSUndoManagerWillUndoChangeNotification"
const WillBecomeMultiThreadedNotification NotificationName = "NSWillBecomeMultiThreadedNotification"
const BundleResourceRequestLowDiskSpaceNotification NotificationName = "NSBundleResourceRequestLowDiskSpaceNotification"
const CalendarDayChangedNotification NotificationName = "NSCalendarDayChangedNotification"
const DidBecomeSingleThreadedNotification NotificationName = "NSDidBecomeSingleThreadedNotification"
const ExtensionHostDidBecomeActiveNotification NotificationName = "NSExtensionHostDidBecomeActiveNotification"
const ExtensionHostDidEnterBackgroundNotification NotificationName = "NSExtensionHostDidEnterBackgroundNotification"
const ExtensionHostWillEnterForegroundNotification NotificationName = "NSExtensionHostWillEnterForegroundNotification"
const ExtensionHostWillResignActiveNotification NotificationName = "NSExtensionHostWillResignActiveNotification"
const FileHandleConnectionAcceptedNotification NotificationName = "NSFileHandleConnectionAcceptedNotification"
const FileHandleDataAvailableNotification NotificationName = "NSFileHandleDataAvailableNotification"
const FileHandleReadToEndOfFileCompletionNotification NotificationName = "NSFileHandleReadToEndOfFileCompletionNotification"
const HTTPCookieManagerAcceptPolicyChangedNotification NotificationName = "com.apple.Foundation.NSHTTPCookieManagerAcceptPolicyChanged"
const HTTPCookieManagerCookiesChangedNotification NotificationName = "NSHTTPCookieManagerCookiesChangedNotification"
const MetadataQueryDidFinishGatheringNotification NotificationName = "NSMetadataQueryDidFinishGatheringNotification"
const MetadataQueryDidStartGatheringNotification NotificationName = "NSMetadataQueryDidStartGatheringNotification"
const MetadataQueryDidUpdateNotification NotificationName = "NSMetadataQueryDidUpdateNotification"
const MetadataQueryGatheringProgressNotification NotificationName = "NSMetadataQueryGatheringProgressNotification"
const ProcessInfoPowerStateDidChangeNotification NotificationName = "NSProcessInfoPowerStateDidChangeNotification"
const SystemClockDidChangeNotification NotificationName = "NSSystemClockDidChangeNotification"
const SystemTimeZoneDidChangeNotification NotificationName = "kCFTimeZoneSystemTimeZoneDidChangeNotification"
const ThreadWillExitNotification NotificationName = "NSThreadWillExitNotification"
const URLCredentialStorageChangedNotification NotificationName = "NSURLCredentialStorageChangedNotification"

type NumberFormatterBehavior uint

const NumberFormatterBehaviorDefault NumberFormatterBehavior = 0
const NumberFormatterBehavior10_0 NumberFormatterBehavior = 1000
const NumberFormatterBehavior10_4 NumberFormatterBehavior = 1040

type NumberFormatterPadPosition uint

const NumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
const NumberFormatterPadAfterPrefix NumberFormatterPadPosition = 1
const NumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 2
const NumberFormatterPadAfterSuffix NumberFormatterPadPosition = 3

type NumberFormatterRoundingMode uint

const NumberFormatterRoundCeiling NumberFormatterRoundingMode = 0
const NumberFormatterRoundFloor NumberFormatterRoundingMode = 1
const NumberFormatterRoundDown NumberFormatterRoundingMode = 2
const NumberFormatterRoundUp NumberFormatterRoundingMode = 3
const NumberFormatterRoundHalfEven NumberFormatterRoundingMode = 4
const NumberFormatterRoundHalfDown NumberFormatterRoundingMode = 5
const NumberFormatterRoundHalfUp NumberFormatterRoundingMode = 6

type NumberFormatterStyle uint

const NumberFormatterNoStyle NumberFormatterStyle = 0
const NumberFormatterDecimalStyle NumberFormatterStyle = 1
const NumberFormatterPercentStyle NumberFormatterStyle = 3
const NumberFormatterScientificStyle NumberFormatterStyle = 4
const NumberFormatterSpellOutStyle NumberFormatterStyle = 5
const NumberFormatterOrdinalStyle NumberFormatterStyle = 6
const NumberFormatterCurrencyStyle NumberFormatterStyle = 2
const NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 10
const NumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 8
const NumberFormatterCurrencyPluralStyle NumberFormatterStyle = 9

type OperationQueuePriority int

const OperationQueuePriorityVeryLow OperationQueuePriority = -8
const OperationQueuePriorityLow OperationQueuePriority = -4
const OperationQueuePriorityNormal OperationQueuePriority = 0
const OperationQueuePriorityHigh OperationQueuePriority = 4
const OperationQueuePriorityVeryHigh OperationQueuePriority = 8

type PersonNameComponentsFormatterOptions uint

const PersonNameComponentsFormatterPhonetic PersonNameComponentsFormatterOptions = 2

type PersonNameComponentsFormatterStyle int

const PersonNameComponentsFormatterStyleDefault PersonNameComponentsFormatterStyle = 0
const PersonNameComponentsFormatterStyleShort PersonNameComponentsFormatterStyle = 1
const PersonNameComponentsFormatterStyleMedium PersonNameComponentsFormatterStyle = 2
const PersonNameComponentsFormatterStyleLong PersonNameComponentsFormatterStyle = 3
const PersonNameComponentsFormatterStyleAbbreviated PersonNameComponentsFormatterStyle = 4

type ProgressFileOperationKind string

const ProgressFileOperationKindCopying ProgressFileOperationKind = "NSProgressFileOperationKindCopying"
const ProgressFileOperationKindDecompressingAfterDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDecompressingAfterDownloading"
const ProgressFileOperationKindDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDownloading"
const ProgressFileOperationKindUploading ProgressFileOperationKind = "NSProgressFileOperationKindUploading"
const ProgressFileOperationKindReceiving ProgressFileOperationKind = "NSProgressFileOperationKindReceiving"
const ProgressFileOperationKindDuplicating ProgressFileOperationKind = "NSProgressFileOperationKindDuplicating"

type ProgressKind string

const ProgressKindFile ProgressKind = "NSProgressKindFile"

type ProgressUserInfoKey string

const ProgressEstimatedTimeRemainingKey ProgressUserInfoKey = "NSProgressEstimatedTimeRemainingKey"
const ProgressThroughputKey ProgressUserInfoKey = "NSProgressThroughputKey"
const ProgressFileAnimationImageKey ProgressUserInfoKey = "NSProgressFlyToImageKey"
const ProgressFileAnimationImageOriginalRectKey ProgressUserInfoKey = "NSProgressFileAnimationImageOriginalRectKey"
const ProgressFileCompletedCountKey ProgressUserInfoKey = "NSProgressFileCompletedCountKey"
const ProgressFileIconKey ProgressUserInfoKey = "NSProgressFileIconKey"
const ProgressFileOperationKindKey ProgressUserInfoKey = "NSProgressFileOperationKindKey"
const ProgressFileTotalCountKey ProgressUserInfoKey = "NSProgressFileTotalCountKey"
const ProgressFileURLKey ProgressUserInfoKey = "NSProgressFileURLKey"

type QualityOfService int

const QualityOfServiceUserInteractive QualityOfService = 33
const QualityOfServiceUserInitiated QualityOfService = 25
const QualityOfServiceUtility QualityOfService = 17
const QualityOfServiceBackground QualityOfService = 9
const QualityOfServiceDefault QualityOfService = -1

type RectEdge uint

const MaxXEdge RectEdge = 2
const MaxYEdge RectEdge = 3
const MinXEdge RectEdge = 0
const MinYEdge RectEdge = 1
const RectEdgeMaxX RectEdge = 2
const RectEdgeMaxY RectEdge = 3
const RectEdgeMinX RectEdge = 0
const RectEdgeMinY RectEdge = 1

type RegularExpressionOptions uint

const RegularExpressionCaseInsensitive RegularExpressionOptions = 1
const RegularExpressionAllowCommentsAndWhitespace RegularExpressionOptions = 2
const RegularExpressionIgnoreMetacharacters RegularExpressionOptions = 4
const RegularExpressionDotMatchesLineSeparators RegularExpressionOptions = 8
const RegularExpressionAnchorsMatchLines RegularExpressionOptions = 16
const RegularExpressionUseUnixLineSeparators RegularExpressionOptions = 32
const RegularExpressionUseUnicodeWordBoundaries RegularExpressionOptions = 64

type RoundingMode uint

const RoundPlain RoundingMode = 0
const RoundDown RoundingMode = 1
const RoundUp RoundingMode = 2
const RoundBankers RoundingMode = 3

type RunLoopMode string

const RunLoopCommonModes RunLoopMode = "kCFRunLoopCommonModes"
const DefaultRunLoopMode RunLoopMode = "kCFRunLoopDefaultMode"

type SaveOptions uint

const SaveOptionsYes SaveOptions = 0
const SaveOptionsNo SaveOptions = 1
const SaveOptionsAsk SaveOptions = 2

type StreamEvent uint

const StreamEventNone StreamEvent = 0
const StreamEventOpenCompleted StreamEvent = 1
const StreamEventHasBytesAvailable StreamEvent = 2
const StreamEventHasSpaceAvailable StreamEvent = 4
const StreamEventErrorOccurred StreamEvent = 8
const StreamEventEndEncountered StreamEvent = 16

type StreamPropertyKey string

const StreamDataWrittenToMemoryStreamKey StreamPropertyKey = "kCFStreamPropertyDataWritten"
const StreamFileCurrentOffsetKey StreamPropertyKey = "kCFStreamPropertyFileCurrentOffset"
const StreamNetworkServiceType StreamPropertyKey = "kCFStreamNetworkServiceType"
const StreamSocketSecurityLevelKey StreamPropertyKey = "kCFStreamPropertySocketSecurityLevel"
const StreamSOCKSProxyConfigurationKey StreamPropertyKey = "kCFStreamPropertySOCKSProxy"

type StreamStatus uint

const StreamStatusAtEnd StreamStatus = 5
const StreamStatusClosed StreamStatus = 6
const StreamStatusError StreamStatus = 7
const StreamStatusNotOpen StreamStatus = 0
const StreamStatusOpen StreamStatus = 2
const StreamStatusOpening StreamStatus = 1
const StreamStatusReading StreamStatus = 3
const StreamStatusWriting StreamStatus = 4

type StringEncoding uint

const ASCIIStringEncoding StringEncoding = 1
const NEXTSTEPStringEncoding StringEncoding = 2
const JapaneseEUCStringEncoding StringEncoding = 3
const UTF8StringEncoding StringEncoding = 4
const ISOLatin1StringEncoding StringEncoding = 5
const SymbolStringEncoding StringEncoding = 6
const NonLossyASCIIStringEncoding StringEncoding = 7
const ShiftJISStringEncoding StringEncoding = 8
const ISOLatin2StringEncoding StringEncoding = 9
const UnicodeStringEncoding StringEncoding = 10
const WindowsCP1251StringEncoding StringEncoding = 11
const WindowsCP1252StringEncoding StringEncoding = 12
const WindowsCP1253StringEncoding StringEncoding = 13
const WindowsCP1254StringEncoding StringEncoding = 14
const WindowsCP1250StringEncoding StringEncoding = 15
const ISO2022JPStringEncoding StringEncoding = 21
const MacOSRomanStringEncoding StringEncoding = 30
const UTF16StringEncoding StringEncoding = 10
const UTF16BigEndianStringEncoding StringEncoding = 2415919360
const UTF16LittleEndianStringEncoding StringEncoding = 2483028224
const UTF32StringEncoding StringEncoding = 2348810496
const UTF32BigEndianStringEncoding StringEncoding = 2550137088
const UTF32LittleEndianStringEncoding StringEncoding = 2617245952
const ProprietaryStringEncoding StringEncoding = 65536

type TextCheckingKey string

const TextCheckingAirlineKey TextCheckingKey = "Airline"
const TextCheckingCityKey TextCheckingKey = "City"
const TextCheckingCountryKey TextCheckingKey = "Country"
const TextCheckingFlightKey TextCheckingKey = "Flight"
const TextCheckingJobTitleKey TextCheckingKey = "JobTitle"
const TextCheckingNameKey TextCheckingKey = "Name"
const TextCheckingOrganizationKey TextCheckingKey = "Organization"
const TextCheckingPhoneKey TextCheckingKey = "Phone"
const TextCheckingStateKey TextCheckingKey = "State"
const TextCheckingStreetKey TextCheckingKey = "Street"
const TextCheckingZIPKey TextCheckingKey = "ZIP"

type TextCheckingType uint64

const TextCheckingTypeOrthography TextCheckingType = 1
const TextCheckingTypeSpelling TextCheckingType = 2
const TextCheckingTypeGrammar TextCheckingType = 4
const TextCheckingTypeDate TextCheckingType = 8
const TextCheckingTypeAddress TextCheckingType = 16
const TextCheckingTypeLink TextCheckingType = 32
const TextCheckingTypeQuote TextCheckingType = 64
const TextCheckingTypeDash TextCheckingType = 128
const TextCheckingTypeReplacement TextCheckingType = 256
const TextCheckingTypeCorrection TextCheckingType = 512
const TextCheckingTypeRegularExpression TextCheckingType = 1024
const TextCheckingTypePhoneNumber TextCheckingType = 2048
const TextCheckingTypeTransitInformation TextCheckingType = 4096

type TextCheckingTypes uint64

const TextCheckingAllSystemTypes TextCheckingTypes = 4294967295
const TextCheckingAllCustomTypes TextCheckingTypes = 4294967296
const TextCheckingAllTypes TextCheckingTypes = math.MaxUint

type TimeInterval float64

type TimeZoneNameStyle int

const TimeZoneNameStyleStandard TimeZoneNameStyle = 0
const TimeZoneNameStyleShortStandard TimeZoneNameStyle = 1
const TimeZoneNameStyleDaylightSaving TimeZoneNameStyle = 2
const TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
const TimeZoneNameStyleGeneric TimeZoneNameStyle = 4
const TimeZoneNameStyleShortGeneric TimeZoneNameStyle = 5

type URLBookmarkCreationOptions uint

const URLBookmarkCreationMinimalBookmark URLBookmarkCreationOptions = 512
const URLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 1024
const URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 2048
const URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 4096
const URLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 536870912

type URLBookmarkFileCreationOptions uint

type URLBookmarkResolutionOptions uint

const URLBookmarkResolutionWithoutUI URLBookmarkResolutionOptions = 256
const URLBookmarkResolutionWithoutMounting URLBookmarkResolutionOptions = 512
const URLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 1024
const URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 32768

type URLRequestAttribution uint

const URLRequestAttributionDeveloper URLRequestAttribution = 0
const URLRequestAttributionUser URLRequestAttribution = 1

type URLRequestCachePolicy uint

const URLRequestUseProtocolCachePolicy URLRequestCachePolicy = 0
const URLRequestReloadIgnoringLocalCacheData URLRequestCachePolicy = 1
const URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
const URLRequestReloadIgnoringCacheData URLRequestCachePolicy = 1
const URLRequestReturnCacheDataElseLoad URLRequestCachePolicy = 2
const URLRequestReturnCacheDataDontLoad URLRequestCachePolicy = 3
const URLRequestReloadRevalidatingCacheData URLRequestCachePolicy = 5

type URLRequestNetworkServiceType uint

const URLNetworkServiceTypeDefault URLRequestNetworkServiceType = 0
const URLNetworkServiceTypeVideo URLRequestNetworkServiceType = 2
const URLNetworkServiceTypeBackground URLRequestNetworkServiceType = 3
const URLNetworkServiceTypeVoice URLRequestNetworkServiceType = 4
const URLNetworkServiceTypeCallSignaling URLRequestNetworkServiceType = 11
const URLNetworkServiceTypeResponsiveData URLRequestNetworkServiceType = 6
const URLNetworkServiceTypeAVStreaming URLRequestNetworkServiceType = 8
const URLNetworkServiceTypeResponsiveAV URLRequestNetworkServiceType = 9

type URLResourceKey string

const URLIsApplicationKey URLResourceKey = "_NSURLIsApplicationKey"
const URLApplicationIsScriptableKey URLResourceKey = "NSURLApplicationIsScriptableKey"
const URLIsDirectoryKey URLResourceKey = "NSURLIsDirectoryKey"
const URLParentDirectoryURLKey URLResourceKey = "NSURLParentDirectoryURLKey"
const URLFileAllocatedSizeKey URLResourceKey = "NSURLFileAllocatedSizeKey"
const URLFileProtectionKey URLResourceKey = "NSURLFileProtectionKey"
const URLFileContentIdentifierKey URLResourceKey = "NSURLFileContentIdentifierKey"
const URLFileResourceIdentifierKey URLResourceKey = "NSURLFileResourceIdentifierKey"
const URLFileResourceTypeKey URLResourceKey = "NSURLFileResourceTypeKey"
const URLFileSecurityKey URLResourceKey = "NSURLFileSecurityKey"
const URLFileSizeKey URLResourceKey = "NSURLFileSizeKey"
const URLIsAliasFileKey URLResourceKey = "NSURLIsAliasFileKey"
const URLIsPackageKey URLResourceKey = "NSURLIsPackageKey"
const URLIsRegularFileKey URLResourceKey = "NSURLIsRegularFileKey"
const URLIsPurgeableKey URLResourceKey = "NSURLIsPurgeableKey"
const URLIsSparseKey URLResourceKey = "NSURLIsSparseKey"
const URLMayHaveExtendedAttributesKey URLResourceKey = "NSURLMayHaveExtendedAttributesKey"
const URLMayShareFileContentKey URLResourceKey = "NSURLMayShareFileContentKey"
const URLPreferredIOBlockSizeKey URLResourceKey = "NSURLPreferredIOBlockSizeKey"
const URLTotalFileAllocatedSizeKey URLResourceKey = "NSURLTotalFileAllocatedSizeKey"
const URLTotalFileSizeKey URLResourceKey = "NSURLTotalFileSizeKey"
const URLVolumeAvailableCapacityKey URLResourceKey = "NSURLVolumeAvailableCapacityKey"
const URLVolumeAvailableCapacityForImportantUsageKey URLResourceKey = "NSURLVolumeAvailableCapacityForImportantUsageKey"
const URLVolumeAvailableCapacityForOpportunisticUsageKey URLResourceKey = "NSURLVolumeAvailableCapacityForOpportunisticUsageKey"
const URLVolumeTotalCapacityKey URLResourceKey = "NSURLVolumeTotalCapacityKey"
const URLVolumeIsAutomountedKey URLResourceKey = "NSURLVolumeIsAutomountedKey"
const URLVolumeIsBrowsableKey URLResourceKey = "NSURLVolumeIsBrowsableKey"
const URLVolumeIsEjectableKey URLResourceKey = "NSURLVolumeIsEjectableKey"
const URLVolumeIsEncryptedKey URLResourceKey = "NSURLVolumeIsEncryptedKey"
const URLVolumeIsInternalKey URLResourceKey = "NSURLVolumeIsInternalKey"
const URLVolumeIsJournalingKey URLResourceKey = "NSURLVolumeIsJournalingKey"
const URLVolumeIsLocalKey URLResourceKey = "NSURLVolumeIsLocalKey"
const URLVolumeIsReadOnlyKey URLResourceKey = "NSURLVolumeIsReadOnlyKey"
const URLVolumeIsRemovableKey URLResourceKey = "NSURLVolumeIsRemovableKey"
const URLVolumeIsRootFileSystemKey URLResourceKey = "NSURLVolumeIsRootFileSystemKey"
const URLVolumeSupportsFileProtectionKey URLResourceKey = "NSURLVolumeSupportsFileProtectionKey"
const URLIsMountTriggerKey URLResourceKey = "NSURLIsMountTriggerKey"
const URLIsVolumeKey URLResourceKey = "NSURLIsVolumeKey"
const URLVolumeCreationDateKey URLResourceKey = "NSURLVolumeCreationDateKey"
const URLVolumeIdentifierKey URLResourceKey = "NSURLVolumeIdentifierKey"
const URLVolumeLocalizedFormatDescriptionKey URLResourceKey = "NSURLVolumeLocalizedFormatDescriptionKey"
const URLVolumeLocalizedNameKey URLResourceKey = "NSURLVolumeLocalizedNameKey"
const URLVolumeMaximumFileSizeKey URLResourceKey = "NSURLVolumeMaximumFileSizeKey"
const URLVolumeNameKey URLResourceKey = "NSURLVolumeNameKey"
const URLVolumeResourceCountKey URLResourceKey = "NSURLVolumeResourceCountKey"
const URLVolumeSupportsAccessPermissionsKey URLResourceKey = "NSURLVolumeSupportsAccessPermissionsKey"
const URLVolumeSupportsAdvisoryFileLockingKey URLResourceKey = "NSURLVolumeSupportsAdvisoryFileLockingKey"
const URLVolumeSupportsCasePreservedNamesKey URLResourceKey = "NSURLVolumeSupportsCasePreservedNamesKey"
const URLVolumeSupportsCaseSensitiveNamesKey URLResourceKey = "NSURLVolumeSupportsCaseSensitiveNamesKey"
const URLVolumeSupportsCompressionKey URLResourceKey = "NSURLVolumeSupportsCompressionKey"
const URLVolumeSupportsExclusiveRenamingKey URLResourceKey = "NSURLVolumeSupportsExclusiveRenamingKey"
const URLVolumeSupportsExtendedSecurityKey URLResourceKey = "NSURLVolumeSupportsExtendedSecurityKey"
const URLVolumeSupportsFileCloningKey URLResourceKey = "NSURLVolumeSupportsFileCloningKey"
const URLVolumeSupportsHardLinksKey URLResourceKey = "NSURLVolumeSupportsHardLinksKey"
const URLVolumeSupportsImmutableFilesKey URLResourceKey = "NSURLVolumeSupportsImmutableFilesKey"
const URLVolumeSupportsJournalingKey URLResourceKey = "NSURLVolumeSupportsJournalingKey"
const URLVolumeSupportsPersistentIDsKey URLResourceKey = "NSURLVolumeSupportsPersistentIDsKey"
const URLVolumeSupportsRenamingKey URLResourceKey = "NSURLVolumeSupportsRenamingKey"
const URLVolumeSupportsRootDirectoryDatesKey URLResourceKey = "NSURLVolumeSupportsRootDirectoryDatesKey"
const URLVolumeSupportsSparseFilesKey URLResourceKey = "NSURLVolumeSupportsSparseFilesKey"
const URLVolumeSupportsSwapRenamingKey URLResourceKey = "NSURLVolumeSupportsSwapRenamingKey"
const URLVolumeSupportsSymbolicLinksKey URLResourceKey = "NSURLVolumeSupportsSymbolicLinksKey"
const URLVolumeSupportsVolumeSizesKey URLResourceKey = "NSURLVolumeSupportsVolumeSizesKey"
const URLVolumeSupportsZeroRunsKey URLResourceKey = "NSURLVolumeSupportsZeroRunsKey"
const URLVolumeURLForRemountingKey URLResourceKey = "NSURLVolumeURLForRemountingKey"
const URLVolumeURLKey URLResourceKey = "NSURLVolumeURLKey"
const URLVolumeUUIDStringKey URLResourceKey = "NSURLVolumeUUIDStringKey"
const URLIsUbiquitousItemKey URLResourceKey = "NSURLIsUbiquitousItemKey"
const URLUbiquitousSharedItemMostRecentEditorNameComponentsKey URLResourceKey = "NSURLUbiquitousSharedItemMostRecentEditorNameComponentsKey"
const URLUbiquitousItemDownloadRequestedKey URLResourceKey = "NSURLUbiquitousItemDownloadRequestedKey"
const URLUbiquitousItemIsDownloadingKey URLResourceKey = "NSURLUbiquitousItemIsDownloadingKey"
const URLUbiquitousItemDownloadingErrorKey URLResourceKey = "NSURLUbiquitousItemDownloadingErrorKey"
const URLUbiquitousItemDownloadingStatusKey URLResourceKey = "NSURLUbiquitousItemDownloadingStatusKey"
const URLUbiquitousItemIsExcludedFromSyncKey URLResourceKey = "NSURLUbiquitousItemIsExcludedFromSyncKey"
const URLUbiquitousItemIsUploadedKey URLResourceKey = "NSURLUbiquitousItemIsUploadedKey"
const URLUbiquitousItemIsUploadingKey URLResourceKey = "NSURLUbiquitousItemIsUploadingKey"
const URLUbiquitousItemUploadingErrorKey URLResourceKey = "NSURLUbiquitousItemUploadingErrorKey"
const URLUbiquitousItemHasUnresolvedConflictsKey URLResourceKey = "NSURLUbiquitousItemHasUnresolvedConflictsKey"
const URLUbiquitousItemContainerDisplayNameKey URLResourceKey = "NSURLUbiquitousItemContainerDisplayNameKey"
const URLUbiquitousSharedItemOwnerNameComponentsKey URLResourceKey = "NSURLUbiquitousSharedItemOwnerNameComponentsKey"
const URLUbiquitousSharedItemCurrentUserPermissionsKey URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserPermissionsKey"
const URLUbiquitousSharedItemCurrentUserRoleKey URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserRoleKey"
const URLUbiquitousItemIsSharedKey URLResourceKey = "NSURLUbiquitousItemIsSharedKey"
const URLKeysOfUnsetValuesKey URLResourceKey = "NSURLKeysOfUnsetValuesKey"
const URLQuarantinePropertiesKey URLResourceKey = "NSURLQuarantinePropertiesKey"
const URLAddedToDirectoryDateKey URLResourceKey = "NSURLAddedToDirectoryDateKey"
const URLAttributeModificationDateKey URLResourceKey = "NSURLAttributeModificationDateKey"
const URLContentAccessDateKey URLResourceKey = "NSURLContentAccessDateKey"
const URLContentModificationDateKey URLResourceKey = "NSURLContentModificationDateKey"
const URLCreationDateKey URLResourceKey = "NSURLCreationDateKey"
const URLCustomIconKey URLResourceKey = "NSURLCustomIconKey"
const URLDocumentIdentifierKey URLResourceKey = "NSURLDocumentIdentifierKey"
const URLEffectiveIconKey URLResourceKey = "NSURLEffectiveIconKey"
const URLGenerationIdentifierKey URLResourceKey = "NSURLGenerationIdentifierKey"
const URLHasHiddenExtensionKey URLResourceKey = "NSURLHasHiddenExtensionKey"
const URLIsExcludedFromBackupKey URLResourceKey = "NSURLIsExcludedFromBackupKey"
const URLIsExecutableKey URLResourceKey = "NSURLIsExecutableKey"
const URLIsHiddenKey URLResourceKey = "NSURLIsHiddenKey"
const URLIsReadableKey URLResourceKey = "NSURLIsReadableKey"
const URLIsSymbolicLinkKey URLResourceKey = "NSURLIsSymbolicLinkKey"
const URLIsSystemImmutableKey URLResourceKey = "NSURLIsSystemImmutableKey"
const URLIsUserImmutableKey URLResourceKey = "NSURLIsUserImmutableKey"
const URLIsWritableKey URLResourceKey = "NSURLIsWritableKey"
const URLLabelColorKey URLResourceKey = "NSURLLabelColorKey"
const URLLabelNumberKey URLResourceKey = "NSURLLabelNumberKey"
const URLLinkCountKey URLResourceKey = "NSURLLinkCountKey"
const URLLocalizedLabelKey URLResourceKey = "NSURLLocalizedLabelKey"
const URLLocalizedNameKey URLResourceKey = "NSURLLocalizedNameKey"
const URLLocalizedTypeDescriptionKey URLResourceKey = "NSURLLocalizedTypeDescriptionKey"
const URLNameKey URLResourceKey = "NSURLNameKey"
const URLPathKey URLResourceKey = "_NSURLPathKey"
const URLCanonicalPathKey URLResourceKey = "NSURLCanonicalPathKey"
const URLTagNamesKey URLResourceKey = "NSURLTagNamesKey"
const URLContentTypeKey URLResourceKey = "NSURLContentTypeKey"

type URLSessionAuthChallengeDisposition int

const URLSessionAuthChallengeUseCredential URLSessionAuthChallengeDisposition = 0
const URLSessionAuthChallengePerformDefaultHandling URLSessionAuthChallengeDisposition = 1
const URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionAuthChallengeDisposition = 2
const URLSessionAuthChallengeRejectProtectionSpace URLSessionAuthChallengeDisposition = 3

type UserActivityPersistentIdentifier string

type unichar uint16
