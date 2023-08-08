// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

type ActivityOptions uint64

const ActivityAutomaticTerminationDisabled ActivityOptions = 32768
const ActivityBackground ActivityOptions = 255
const ActivityIdleDisplaySleepDisabled ActivityOptions = 1099511627776
const ActivityIdleSystemSleepDisabled ActivityOptions = 1048576
const ActivityLatencyCritical ActivityOptions = 1095216660480
const ActivitySuddenTerminationDisabled ActivityOptions = 16384
const ActivityUserInitiated ActivityOptions = 16777215
const ActivityUserInitiatedAllowingIdleSystemSleep ActivityOptions = 15728639

type AlignmentOptions int64

const AlignAllEdgesInward AlignmentOptions = 15
const AlignAllEdgesNearest AlignmentOptions = 983040
const AlignAllEdgesOutward AlignmentOptions = 3840
const AlignHeightInward AlignmentOptions = 32
const AlignHeightNearest AlignmentOptions = 2097152
const AlignHeightOutward AlignmentOptions = 8192
const AlignMaxXInward AlignmentOptions = 4
const AlignMaxXNearest AlignmentOptions = 262144
const AlignMaxXOutward AlignmentOptions = 1024
const AlignMaxYInward AlignmentOptions = 8
const AlignMaxYNearest AlignmentOptions = 524288
const AlignMaxYOutward AlignmentOptions = 2048
const AlignMinXInward AlignmentOptions = 1
const AlignMinXNearest AlignmentOptions = 65536
const AlignMinXOutward AlignmentOptions = 256
const AlignMinYInward AlignmentOptions = 2
const AlignMinYNearest AlignmentOptions = 131072
const AlignMinYOutward AlignmentOptions = 512
const AlignRectFlipped AlignmentOptions = -9223372036854775808
const AlignWidthInward AlignmentOptions = 16
const AlignWidthNearest AlignmentOptions = 1048576
const AlignWidthOutward AlignmentOptions = 4096

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

const AttributedStringEnumerationLongestEffectiveRangeNotRequired AttributedStringEnumerationOptions = 1048576
const AttributedStringEnumerationReverse AttributedStringEnumerationOptions = 2

type AttributedStringFormattingOptions uint

const AttributedStringFormattingApplyReplacementIndexAttribute AttributedStringFormattingOptions = 2
const AttributedStringFormattingInsertArgumentAttributesWithoutMerging AttributedStringFormattingOptions = 1

type AttributedStringKey string

const AlternateDescriptionAttributeName AttributedStringKey = "NSAlternateDescription"
const ImageURLAttributeName AttributedStringKey = "NSImageURL"
const InflectionAlternativeAttributeName AttributedStringKey = "NSInflectionAlternative"
const InflectionRuleAttributeName AttributedStringKey = "NSInflect"
const InlinePresentationIntentAttributeName AttributedStringKey = "NSInlinePresentationIntent"
const LanguageIdentifierAttributeName AttributedStringKey = "NSLanguage"
const MorphologyAttributeName AttributedStringKey = "NSMorphology"
const PresentationIntentAttributeName AttributedStringKey = "NSPresentationIntent"
const ReplacementIndexAttributeName AttributedStringKey = "NSReplacementIndex"

type AttributedStringMarkdownInterpretedSyntax int

const AttributedStringMarkdownInterpretedSyntaxFull AttributedStringMarkdownInterpretedSyntax = 0
const AttributedStringMarkdownInterpretedSyntaxInlineOnly AttributedStringMarkdownInterpretedSyntax = 1
const AttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace AttributedStringMarkdownInterpretedSyntax = 2

type AttributedStringMarkdownParsingFailurePolicy int

const AttributedStringMarkdownParsingFailureReturnError AttributedStringMarkdownParsingFailurePolicy = 0
const AttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible AttributedStringMarkdownParsingFailurePolicy = 1

type BackgroundActivityResult int

const BackgroundActivityResultDeferred BackgroundActivityResult = 2
const BackgroundActivityResultFinished BackgroundActivityResult = 1

type BinarySearchingOptions uint

const BinarySearchingFirstEqual BinarySearchingOptions = 256
const BinarySearchingInsertionIndex BinarySearchingOptions = 1024
const BinarySearchingLastEqual BinarySearchingOptions = 512

type ByteCountFormatterCountStyle int

const ByteCountFormatterCountStyleBinary ByteCountFormatterCountStyle = 3
const ByteCountFormatterCountStyleDecimal ByteCountFormatterCountStyle = 2
const ByteCountFormatterCountStyleFile ByteCountFormatterCountStyle = 0
const ByteCountFormatterCountStyleMemory ByteCountFormatterCountStyle = 1

type ByteCountFormatterUnits uint

const ByteCountFormatterUseAll ByteCountFormatterUnits = 65535
const ByteCountFormatterUseBytes ByteCountFormatterUnits = 1
const ByteCountFormatterUseDefault ByteCountFormatterUnits = 0
const ByteCountFormatterUseEB ByteCountFormatterUnits = 64
const ByteCountFormatterUseGB ByteCountFormatterUnits = 8
const ByteCountFormatterUseKB ByteCountFormatterUnits = 2
const ByteCountFormatterUseMB ByteCountFormatterUnits = 4
const ByteCountFormatterUsePB ByteCountFormatterUnits = 32
const ByteCountFormatterUseTB ByteCountFormatterUnits = 16
const ByteCountFormatterUseYBOrHigher ByteCountFormatterUnits = 65280
const ByteCountFormatterUseZB ByteCountFormatterUnits = 128

type CalculationError uint

const CalculationDivideByZero CalculationError = 4
const CalculationLossOfPrecision CalculationError = 1
const CalculationNoError CalculationError = 0
const CalculationOverflow CalculationError = 3
const CalculationUnderflow CalculationError = 2

type CalendarIdentifier string

const CalendarIdentifierBuddhist CalendarIdentifier = "buddhist"
const CalendarIdentifierChinese CalendarIdentifier = "chinese"
const CalendarIdentifierCoptic CalendarIdentifier = "coptic"
const CalendarIdentifierEthiopicAmeteAlem CalendarIdentifier = "ethiopic-amete-alem"
const CalendarIdentifierEthiopicAmeteMihret CalendarIdentifier = "ethiopic"
const CalendarIdentifierGregorian CalendarIdentifier = "gregorian"
const CalendarIdentifierHebrew CalendarIdentifier = "hebrew"
const CalendarIdentifierISO8601 CalendarIdentifier = "iso8601"
const CalendarIdentifierIndian CalendarIdentifier = "indian"
const CalendarIdentifierIslamic CalendarIdentifier = "islamic"
const CalendarIdentifierIslamicCivil CalendarIdentifier = "islamic-civil"
const CalendarIdentifierIslamicTabular CalendarIdentifier = "islamic-tbla"
const CalendarIdentifierIslamicUmmAlQura CalendarIdentifier = "islamic-umalqura"
const CalendarIdentifierJapanese CalendarIdentifier = "japanese"
const CalendarIdentifierPersian CalendarIdentifier = "persian"
const CalendarIdentifierRepublicOfChina CalendarIdentifier = "roc"

type CalendarOptions uint

const CalendarMatchFirst CalendarOptions = 4096
const CalendarMatchLast CalendarOptions = 8192
const CalendarMatchNextTime CalendarOptions = 1024
const CalendarMatchNextTimePreservingSmallerUnits CalendarOptions = 512
const CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 256
const CalendarMatchStrictly CalendarOptions = 2
const CalendarSearchBackwards CalendarOptions = 4
const CalendarWrapComponents CalendarOptions = 1

type CalendarUnit uint

const CalendarCalendarUnit CalendarUnit = 1048576
const CalendarUnitCalendar CalendarUnit = 1048576
const CalendarUnitDay CalendarUnit = 16
const CalendarUnitEra CalendarUnit = 2
const CalendarUnitHour CalendarUnit = 32
const CalendarUnitMinute CalendarUnit = 64
const CalendarUnitMonth CalendarUnit = 8
const CalendarUnitNanosecond CalendarUnit = 32768
const CalendarUnitQuarter CalendarUnit = 2048
const CalendarUnitSecond CalendarUnit = 128
const CalendarUnitTimeZone CalendarUnit = 2097152
const CalendarUnitWeekOfMonth CalendarUnit = 4096
const CalendarUnitWeekOfYear CalendarUnit = 8192
const CalendarUnitWeekday CalendarUnit = 512
const CalendarUnitWeekdayOrdinal CalendarUnit = 1024
const CalendarUnitYear CalendarUnit = 4
const CalendarUnitYearForWeekOfYear CalendarUnit = 16384
const DayCalendarUnit CalendarUnit = 16
const EraCalendarUnit CalendarUnit = 2
const HourCalendarUnit CalendarUnit = 32
const MinuteCalendarUnit CalendarUnit = 64
const MonthCalendarUnit CalendarUnit = 8
const QuarterCalendarUnit CalendarUnit = 2048
const SecondCalendarUnit CalendarUnit = 128
const TimeZoneCalendarUnit CalendarUnit = 2097152
const WeekCalendarUnit CalendarUnit = 256
const WeekOfMonthCalendarUnit CalendarUnit = 4096
const WeekOfYearCalendarUnit CalendarUnit = 8192
const WeekdayCalendarUnit CalendarUnit = 512
const WeekdayOrdinalCalendarUnit CalendarUnit = 1024
const YearCalendarUnit CalendarUnit = 4
const YearForWeekOfYearCalendarUnit CalendarUnit = 16384

type CollectionChangeType int

const CollectionChangeInsert CollectionChangeType = 0
const CollectionChangeRemove CollectionChangeType = 1

type ComparisonPredicateModifier uint

const AllPredicateModifier ComparisonPredicateModifier = 1
const AnyPredicateModifier ComparisonPredicateModifier = 2
const DirectPredicateModifier ComparisonPredicateModifier = 0

type ComparisonPredicateOptions uint

const CaseInsensitivePredicateOption ComparisonPredicateOptions = 1
const DiacriticInsensitivePredicateOption ComparisonPredicateOptions = 2
const NormalizedPredicateOption ComparisonPredicateOptions = 4

type ComparisonResult int

const OrderedAscending ComparisonResult = -1
const OrderedDescending ComparisonResult = 1
const OrderedSame ComparisonResult = 0

type CompoundPredicateType uint

const AndPredicateType CompoundPredicateType = 1
const NotPredicateType CompoundPredicateType = 0
const OrPredicateType CompoundPredicateType = 2

type DataBase64DecodingOptions uint

const DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1

type DataBase64EncodingOptions uint

const DataBase64Encoding64CharacterLineLength DataBase64EncodingOptions = 1
const DataBase64Encoding76CharacterLineLength DataBase64EncodingOptions = 2
const DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 16
const DataBase64EncodingEndLineWithLineFeed DataBase64EncodingOptions = 32

type DataCompressionAlgorithm int

const DataCompressionAlgorithmLZ4 DataCompressionAlgorithm = 1
const DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
const DataCompressionAlgorithmLZMA DataCompressionAlgorithm = 2
const DataCompressionAlgorithmZlib DataCompressionAlgorithm = 3

type DataReadingOptions uint

const DataReadingMapped DataReadingOptions = 1
const DataReadingMappedAlways DataReadingOptions = 8
const DataReadingMappedIfSafe DataReadingOptions = 1
const DataReadingUncached DataReadingOptions = 2
const MappedRead DataReadingOptions = 1
const UncachedRead DataReadingOptions = 2

type DataSearchOptions uint

const DataSearchAnchored DataSearchOptions = 2
const DataSearchBackwards DataSearchOptions = 1

type DataWritingOptions uint

const AtomicWrite DataWritingOptions = 1
const DataWritingAtomic DataWritingOptions = 1
const DataWritingFileProtectionComplete DataWritingOptions = 536870912
const DataWritingFileProtectionCompleteUnlessOpen DataWritingOptions = 805306368
const DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 1073741824
const DataWritingFileProtectionMask DataWritingOptions = 4026531840
const DataWritingFileProtectionNone DataWritingOptions = 268435456
const DataWritingWithoutOverwriting DataWritingOptions = 2

type DateComponentsFormatterUnitsStyle int

const DateComponentsFormatterUnitsStyleAbbreviated DateComponentsFormatterUnitsStyle = 1
const DateComponentsFormatterUnitsStyleBrief DateComponentsFormatterUnitsStyle = 5
const DateComponentsFormatterUnitsStyleFull DateComponentsFormatterUnitsStyle = 3
const DateComponentsFormatterUnitsStylePositional DateComponentsFormatterUnitsStyle = 0
const DateComponentsFormatterUnitsStyleShort DateComponentsFormatterUnitsStyle = 2
const DateComponentsFormatterUnitsStyleSpellOut DateComponentsFormatterUnitsStyle = 4

type DateComponentsFormatterZeroFormattingBehavior uint

const DateComponentsFormatterZeroFormattingBehaviorDefault DateComponentsFormatterZeroFormattingBehavior = 1
const DateComponentsFormatterZeroFormattingBehaviorDropAll DateComponentsFormatterZeroFormattingBehavior = 14
const DateComponentsFormatterZeroFormattingBehaviorDropLeading DateComponentsFormatterZeroFormattingBehavior = 2
const DateComponentsFormatterZeroFormattingBehaviorDropMiddle DateComponentsFormatterZeroFormattingBehavior = 4
const DateComponentsFormatterZeroFormattingBehaviorDropTrailing DateComponentsFormatterZeroFormattingBehavior = 8
const DateComponentsFormatterZeroFormattingBehaviorNone DateComponentsFormatterZeroFormattingBehavior = 0
const DateComponentsFormatterZeroFormattingBehaviorPad DateComponentsFormatterZeroFormattingBehavior = 65536

type DateFormatterBehavior uint

const DateFormatterBehavior10_0 DateFormatterBehavior = 1000
const DateFormatterBehavior10_4 DateFormatterBehavior = 1040
const DateFormatterBehaviorDefault DateFormatterBehavior = 0

type DateFormatterStyle uint

const DateFormatterFullStyle DateFormatterStyle = 4
const DateFormatterLongStyle DateFormatterStyle = 3
const DateFormatterMediumStyle DateFormatterStyle = 2
const DateFormatterNoStyle DateFormatterStyle = 0
const DateFormatterShortStyle DateFormatterStyle = 1

type DateIntervalFormatterStyle uint

const DateIntervalFormatterFullStyle DateIntervalFormatterStyle = 4
const DateIntervalFormatterLongStyle DateIntervalFormatterStyle = 3
const DateIntervalFormatterMediumStyle DateIntervalFormatterStyle = 2
const DateIntervalFormatterNoStyle DateIntervalFormatterStyle = 0
const DateIntervalFormatterShortStyle DateIntervalFormatterStyle = 1

type DecodingFailurePolicy int

const DecodingFailurePolicyRaiseException DecodingFailurePolicy = 0
const DecodingFailurePolicySetErrorAndReturn DecodingFailurePolicy = 1

type DirectoryEnumerationOptions uint

const DirectoryEnumerationIncludesDirectoriesPostOrder DirectoryEnumerationOptions = 8
const DirectoryEnumerationProducesRelativePathURLs DirectoryEnumerationOptions = 16
const DirectoryEnumerationSkipsHiddenFiles DirectoryEnumerationOptions = 4
const DirectoryEnumerationSkipsPackageDescendants DirectoryEnumerationOptions = 2
const DirectoryEnumerationSkipsSubdirectoryDescendants DirectoryEnumerationOptions = 1

type DistributedNotificationCenterType string

const LocalNotificationCenterType DistributedNotificationCenterType = "_NSLocalNotificationCenter"

type DistributedNotificationOptions uint

const DistributedNotificationDeliverImmediately DistributedNotificationOptions = 1
const DistributedNotificationPostToAllSessions DistributedNotificationOptions = 2
const NotificationDeliverImmediately DistributedNotificationOptions = 1
const NotificationPostToAllSessions DistributedNotificationOptions = 2

type EnergyFormatterUnit int

const EnergyFormatterUnitCalorie EnergyFormatterUnit = 1793
const EnergyFormatterUnitJoule EnergyFormatterUnit = 11
const EnergyFormatterUnitKilocalorie EnergyFormatterUnit = 1794
const EnergyFormatterUnitKilojoule EnergyFormatterUnit = 14

type EnumerationOptions uint

const EnumerationConcurrent EnumerationOptions = 1
const EnumerationReverse EnumerationOptions = 2

type ErrorDomain string

const CocoaErrorDomain ErrorDomain = "NSCocoaErrorDomain"
const MachErrorDomain ErrorDomain = "NSMachErrorDomain"
const NetServicesErrorDomain ErrorDomain = "NSNetServicesErrorDomain"
const OSStatusErrorDomain ErrorDomain = "NSOSStatusErrorDomain"
const POSIXErrorDomain ErrorDomain = "NSPOSIXErrorDomain"
const StreamSOCKSErrorDomain ErrorDomain = "NSStreamSOCKSErrorDomain"
const StreamSocketSSLErrorDomain ErrorDomain = "NSStreamSocketSSLErrorDomain"
const URLErrorDomain ErrorDomain = "NSURLErrorDomain"
const XMLParserErrorDomain ErrorDomain = "NSXMLParserErrorDomain"

type ErrorUserInfoKey string

const DebugDescriptionErrorKey ErrorUserInfoKey = "NSDebugDescription"
const FilePathErrorKey ErrorUserInfoKey = "NSFilePath"
const HelpAnchorErrorKey ErrorUserInfoKey = "NSHelpAnchor"
const LocalizedDescriptionKey ErrorUserInfoKey = "NSLocalizedDescription"
const LocalizedFailureErrorKey ErrorUserInfoKey = "NSLocalizedFailure"
const LocalizedFailureReasonErrorKey ErrorUserInfoKey = "NSLocalizedFailureReason"
const LocalizedRecoveryOptionsErrorKey ErrorUserInfoKey = "NSLocalizedRecoveryOptions"
const LocalizedRecoverySuggestionErrorKey ErrorUserInfoKey = "NSLocalizedRecoverySuggestion"
const MultipleUnderlyingErrorsKey ErrorUserInfoKey = "NSMultipleUnderlyingErrorsKey"
const RecoveryAttempterErrorKey ErrorUserInfoKey = "NSRecoveryAttempter"
const StringEncodingErrorKey ErrorUserInfoKey = "NSStringEncoding"
const URLErrorKey ErrorUserInfoKey = "NSURL"
const URLErrorNetworkUnavailableReasonKey ErrorUserInfoKey = "NSURLErrorNetworkUnavailableReasonKey"
const UnderlyingErrorKey ErrorUserInfoKey = "NSUnderlyingError"

type ExceptionName string

const CharacterConversionException ExceptionName = "NSCharacterConversionException"
const DecimalNumberDivideByZeroException ExceptionName = "NSDecimalNumberDivideByZeroException"
const DecimalNumberExactnessException ExceptionName = "NSDecimalNumberExactnessException"
const DecimalNumberOverflowException ExceptionName = "NSDecimalNumberOverflowException"
const DecimalNumberUnderflowException ExceptionName = "NSDecimalNumberUnderflowException"
const DestinationInvalidException ExceptionName = "NSDestinationInvalidException"
const FileHandleOperationException ExceptionName = "NSFileHandleOperationException"
const GenericException ExceptionName = "NSGenericException"
const InconsistentArchiveException ExceptionName = "NSArchiverArchiveInconsistency"
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

type ExpressionType uint

const AggregateExpressionType ExpressionType = 14
const AnyKeyExpressionType ExpressionType = 15
const BlockExpressionType ExpressionType = 19
const ConditionalExpressionType ExpressionType = 20
const ConstantValueExpressionType ExpressionType = 0
const EvaluatedObjectExpressionType ExpressionType = 1
const FunctionExpressionType ExpressionType = 4
const IntersectSetExpressionType ExpressionType = 6
const KeyPathExpressionType ExpressionType = 3
const MinusSetExpressionType ExpressionType = 7
const SubqueryExpressionType ExpressionType = 13
const UnionSetExpressionType ExpressionType = 5
const VariableExpressionType ExpressionType = 2

type FileAttributeKey string

const FileAppendOnly FileAttributeKey = "NSFileAppendOnly"
const FileBusy FileAttributeKey = "NSFileBusy"
const FileCreationDate FileAttributeKey = "NSFileCreationDate"
const FileDeviceIdentifier FileAttributeKey = "NSFileDeviceIdentifier"
const FileExtensionHidden FileAttributeKey = "NSFileExtensionHidden"
const FileGroupOwnerAccountID FileAttributeKey = "NSFileGroupOwnerAccountID"
const FileGroupOwnerAccountName FileAttributeKey = "NSFileGroupOwnerAccountName"
const FileHFSCreatorCode FileAttributeKey = "NSFileHFSCreatorCode"
const FileHFSTypeCode FileAttributeKey = "NSFileHFSTypeCode"
const FileImmutable FileAttributeKey = "NSFileImmutable"
const FileModificationDate FileAttributeKey = "NSFileModificationDate"
const FileOwnerAccountID FileAttributeKey = "NSFileOwnerAccountID"
const FileOwnerAccountName FileAttributeKey = "NSFileOwnerAccountName"
const FilePosixPermissions FileAttributeKey = "NSFilePosixPermissions"
const FileProtectionKey FileAttributeKey = "NSFileProtectionKey"
const FileReferenceCount FileAttributeKey = "NSFileReferenceCount"
const FileSize FileAttributeKey = "NSFileSize"
const FileSystemFileNumber FileAttributeKey = "NSFileSystemFileNumber"
const FileSystemFreeNodes FileAttributeKey = "NSFileSystemFreeNodes"
const FileSystemFreeSize FileAttributeKey = "NSFileSystemFreeSize"
const FileSystemNodes FileAttributeKey = "NSFileSystemNodes"
const FileSystemNumber FileAttributeKey = "NSFileSystemNumber"
const FileSystemSize FileAttributeKey = "NSFileSystemSize"
const FileType FileAttributeKey = "NSFileType"

type FileAttributeType string

const FileTypeBlockSpecial FileAttributeType = "NSFileTypeBlockSpecial"
const FileTypeCharacterSpecial FileAttributeType = "NSFileTypeCharacterSpecial"
const FileTypeDirectory FileAttributeType = "NSFileTypeDirectory"
const FileTypeRegular FileAttributeType = "NSFileTypeRegular"
const FileTypeSocket FileAttributeType = "NSFileTypeSocket"
const FileTypeSymbolicLink FileAttributeType = "NSFileTypeSymbolicLink"
const FileTypeUnknown FileAttributeType = "NSFileTypeUnknown"

type FileCoordinatorReadingOptions uint

const FileCoordinatorReadingForUploading FileCoordinatorReadingOptions = 8
const FileCoordinatorReadingImmediatelyAvailableMetadataOnly FileCoordinatorReadingOptions = 4
const FileCoordinatorReadingResolvesSymbolicLink FileCoordinatorReadingOptions = 2
const FileCoordinatorReadingWithoutChanges FileCoordinatorReadingOptions = 1

type FileCoordinatorWritingOptions uint

const FileCoordinatorWritingContentIndependentMetadataOnly FileCoordinatorWritingOptions = 16
const FileCoordinatorWritingForDeleting FileCoordinatorWritingOptions = 1
const FileCoordinatorWritingForMerging FileCoordinatorWritingOptions = 4
const FileCoordinatorWritingForMoving FileCoordinatorWritingOptions = 2
const FileCoordinatorWritingForReplacing FileCoordinatorWritingOptions = 8

type FileManagerItemReplacementOptions uint

const FileManagerItemReplacementUsingNewMetadataOnly FileManagerItemReplacementOptions = 1
const FileManagerItemReplacementWithoutDeletingBackupItem FileManagerItemReplacementOptions = 2

type FileManagerUnmountOptions uint

const FileManagerUnmountAllPartitionsAndEjectDisk FileManagerUnmountOptions = 1
const FileManagerUnmountWithoutUI FileManagerUnmountOptions = 2

type FileProtectionType string

const FileProtectionComplete FileProtectionType = "NSFileProtectionComplete"
const FileProtectionCompleteUnlessOpen FileProtectionType = "NSFileProtectionCompleteUnlessOpen"
const FileProtectionCompleteUntilFirstUserAuthentication FileProtectionType = "NSFileProtectionCompleteUntilFirstUserAuthentication"
const FileProtectionNone FileProtectionType = "NSFileProtectionNone"

type FileProviderServiceName string

type FileVersionAddingOptions uint

const FileVersionAddingByMoving FileVersionAddingOptions = 1

type FileVersionReplacingOptions uint

const FileVersionReplacingByMoving FileVersionReplacingOptions = 1

type FileWrapperReadingOptions uint

const FileWrapperReadingImmediate FileWrapperReadingOptions = 1
const FileWrapperReadingWithoutMapping FileWrapperReadingOptions = 2

type FileWrapperWritingOptions uint

const FileWrapperWritingAtomic FileWrapperWritingOptions = 1
const FileWrapperWritingWithNameUpdating FileWrapperWritingOptions = 2

type FormattingContext int

const FormattingContextBeginningOfSentence FormattingContext = 4
const FormattingContextDynamic FormattingContext = 1
const FormattingContextListItem FormattingContext = 3
const FormattingContextMiddleOfSentence FormattingContext = 5
const FormattingContextStandalone FormattingContext = 2
const FormattingContextUnknown FormattingContext = 0

type FormattingUnitStyle int

const FormattingUnitStyleLong FormattingUnitStyle = 3
const FormattingUnitStyleMedium FormattingUnitStyle = 2
const FormattingUnitStyleShort FormattingUnitStyle = 1

type GrammaticalGender int

const GrammaticalGenderFeminine GrammaticalGender = 1
const GrammaticalGenderMasculine GrammaticalGender = 2
const GrammaticalGenderNeuter GrammaticalGender = 3
const GrammaticalGenderNotSet GrammaticalGender = 0

type GrammaticalNumber int

const GrammaticalNumberNotSet GrammaticalNumber = 0
const GrammaticalNumberPlural GrammaticalNumber = 3
const GrammaticalNumberPluralFew GrammaticalNumber = 5
const GrammaticalNumberPluralMany GrammaticalNumber = 6
const GrammaticalNumberPluralTwo GrammaticalNumber = 4
const GrammaticalNumberSingular GrammaticalNumber = 1
const GrammaticalNumberZero GrammaticalNumber = 2

type GrammaticalPartOfSpeech int

const GrammaticalPartOfSpeechAbbreviation GrammaticalPartOfSpeech = 14
const GrammaticalPartOfSpeechAdjective GrammaticalPartOfSpeech = 6
const GrammaticalPartOfSpeechAdposition GrammaticalPartOfSpeech = 7
const GrammaticalPartOfSpeechAdverb GrammaticalPartOfSpeech = 4
const GrammaticalPartOfSpeechConjunction GrammaticalPartOfSpeech = 10
const GrammaticalPartOfSpeechDeterminer GrammaticalPartOfSpeech = 1
const GrammaticalPartOfSpeechInterjection GrammaticalPartOfSpeech = 12
const GrammaticalPartOfSpeechLetter GrammaticalPartOfSpeech = 3
const GrammaticalPartOfSpeechNotSet GrammaticalPartOfSpeech = 0
const GrammaticalPartOfSpeechNoun GrammaticalPartOfSpeech = 9
const GrammaticalPartOfSpeechNumeral GrammaticalPartOfSpeech = 11
const GrammaticalPartOfSpeechParticle GrammaticalPartOfSpeech = 5
const GrammaticalPartOfSpeechPreposition GrammaticalPartOfSpeech = 13
const GrammaticalPartOfSpeechPronoun GrammaticalPartOfSpeech = 2
const GrammaticalPartOfSpeechVerb GrammaticalPartOfSpeech = 8

type HTTPCookieAcceptPolicy uint

const HTTPCookieAcceptPolicyAlways HTTPCookieAcceptPolicy = 0
const HTTPCookieAcceptPolicyNever HTTPCookieAcceptPolicy = 1
const HTTPCookieAcceptPolicyOnlyFromMainDocumentDomain HTTPCookieAcceptPolicy = 2

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

const HTTPCookieSameSiteLax HTTPCookieStringPolicy = "lax"
const HTTPCookieSameSiteStrict HTTPCookieStringPolicy = "strict"

type HashTableOptions uint

type ISO8601DateFormatOptions uint

const ISO8601DateFormatWithColonSeparatorInTime ISO8601DateFormatOptions = 512
const ISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 1024
const ISO8601DateFormatWithDashSeparatorInDate ISO8601DateFormatOptions = 256
const ISO8601DateFormatWithDay ISO8601DateFormatOptions = 16
const ISO8601DateFormatWithFractionalSeconds ISO8601DateFormatOptions = 2048
const ISO8601DateFormatWithFullDate ISO8601DateFormatOptions = 275
const ISO8601DateFormatWithFullTime ISO8601DateFormatOptions = 1632
const ISO8601DateFormatWithInternetDateTime ISO8601DateFormatOptions = 1907
const ISO8601DateFormatWithMonth ISO8601DateFormatOptions = 2
const ISO8601DateFormatWithSpaceBetweenDateAndTime ISO8601DateFormatOptions = 128
const ISO8601DateFormatWithTime ISO8601DateFormatOptions = 32
const ISO8601DateFormatWithTimeZone ISO8601DateFormatOptions = 64
const ISO8601DateFormatWithWeekOfYear ISO8601DateFormatOptions = 4
const ISO8601DateFormatWithYear ISO8601DateFormatOptions = 1

type InlinePresentationIntent uint

const InlinePresentationIntentBlockHTML InlinePresentationIntent = 512
const InlinePresentationIntentCode InlinePresentationIntent = 4
const InlinePresentationIntentEmphasized InlinePresentationIntent = 1
const InlinePresentationIntentInlineHTML InlinePresentationIntent = 256
const InlinePresentationIntentLineBreak InlinePresentationIntent = 128
const InlinePresentationIntentSoftBreak InlinePresentationIntent = 64
const InlinePresentationIntentStrikethrough InlinePresentationIntent = 32
const InlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2

type InsertionPosition uint

const PositionAfter InsertionPosition = 0
const PositionBefore InsertionPosition = 1
const PositionBeginning InsertionPosition = 2
const PositionEnd InsertionPosition = 3
const PositionReplace InsertionPosition = 4

type ItemProviderErrorCode int

const ItemProviderItemUnavailableError ItemProviderErrorCode = -1000
const ItemProviderUnavailableCoercionError ItemProviderErrorCode = -1200
const ItemProviderUnexpectedValueClassError ItemProviderErrorCode = -1100
const ItemProviderUnknownError ItemProviderErrorCode = -1

type ItemProviderFileOptions int

const ItemProviderFileOptionOpenInPlace ItemProviderFileOptions = 1

type ItemProviderRepresentationVisibility int

const ItemProviderRepresentationVisibilityAll ItemProviderRepresentationVisibility = 0
const ItemProviderRepresentationVisibilityGroup ItemProviderRepresentationVisibility = 2
const ItemProviderRepresentationVisibilityOwnProcess ItemProviderRepresentationVisibility = 3

type JSONReadingOptions uint

const JSONReadingAllowFragments JSONReadingOptions = 4
const JSONReadingFragmentsAllowed JSONReadingOptions = 4
const JSONReadingJSON5Allowed JSONReadingOptions = 8
const JSONReadingMutableContainers JSONReadingOptions = 1
const JSONReadingMutableLeaves JSONReadingOptions = 2
const JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 16

type JSONWritingOptions uint

const JSONWritingFragmentsAllowed JSONWritingOptions = 4
const JSONWritingPrettyPrinted JSONWritingOptions = 1
const JSONWritingSortedKeys JSONWritingOptions = 2
const JSONWritingWithoutEscapingSlashes JSONWritingOptions = 8

type KeyValueChange uint

const KeyValueChangeInsertion KeyValueChange = 2
const KeyValueChangeRemoval KeyValueChange = 3
const KeyValueChangeReplacement KeyValueChange = 4
const KeyValueChangeSetting KeyValueChange = 1

type KeyValueChangeKey string

const KeyValueChangeIndexesKey KeyValueChangeKey = "indexes"
const KeyValueChangeKindKey KeyValueChangeKey = "kind"
const KeyValueChangeNewKey KeyValueChangeKey = "new"
const KeyValueChangeNotificationIsPriorKey KeyValueChangeKey = "notificationIsPrior"
const KeyValueChangeOldKey KeyValueChangeKey = "old"

type KeyValueObservingOptions uint

const KeyValueObservingOptionInitial KeyValueObservingOptions = 4
const KeyValueObservingOptionNew KeyValueObservingOptions = 1
const KeyValueObservingOptionOld KeyValueObservingOptions = 2
const KeyValueObservingOptionPrior KeyValueObservingOptions = 8

type KeyValueOperator string

const AverageKeyValueOperator KeyValueOperator = "avg"
const CountKeyValueOperator KeyValueOperator = "count"
const DistinctUnionOfArraysKeyValueOperator KeyValueOperator = "distinctUnionOfArrays"
const DistinctUnionOfObjectsKeyValueOperator KeyValueOperator = "distinctUnionOfObjects"
const DistinctUnionOfSetsKeyValueOperator KeyValueOperator = "distinctUnionOfSets"
const MaximumKeyValueOperator KeyValueOperator = "max"
const MinimumKeyValueOperator KeyValueOperator = "min"
const SumKeyValueOperator KeyValueOperator = "sum"
const UnionOfArraysKeyValueOperator KeyValueOperator = "unionOfArrays"
const UnionOfObjectsKeyValueOperator KeyValueOperator = "unionOfObjects"
const UnionOfSetsKeyValueOperator KeyValueOperator = "unionOfSets"

type KeyValueSetMutationKind uint

const KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
const KeyValueMinusSetMutation KeyValueSetMutationKind = 2
const KeyValueSetSetMutation KeyValueSetMutationKind = 4
const KeyValueUnionSetMutation KeyValueSetMutationKind = 1

type LengthFormatterUnit int

const LengthFormatterUnitCentimeter LengthFormatterUnit = 9
const LengthFormatterUnitFoot LengthFormatterUnit = 1282
const LengthFormatterUnitInch LengthFormatterUnit = 1281
const LengthFormatterUnitKilometer LengthFormatterUnit = 14
const LengthFormatterUnitMeter LengthFormatterUnit = 11
const LengthFormatterUnitMile LengthFormatterUnit = 1284
const LengthFormatterUnitMillimeter LengthFormatterUnit = 8
const LengthFormatterUnitYard LengthFormatterUnit = 1283

type LinguisticTag string

const LinguisticTagAdjective LinguisticTag = "Adjective"
const LinguisticTagAdverb LinguisticTag = "Adverb"
const LinguisticTagClassifier LinguisticTag = "Classifier"
const LinguisticTagCloseParenthesis LinguisticTag = "CloseParenthesis"
const LinguisticTagCloseQuote LinguisticTag = "CloseQuote"
const LinguisticTagConjunction LinguisticTag = "Conjunction"
const LinguisticTagDash LinguisticTag = "Dash"
const LinguisticTagDeterminer LinguisticTag = "Determiner"
const LinguisticTagIdiom LinguisticTag = "Idiom"
const LinguisticTagInterjection LinguisticTag = "Interjection"
const LinguisticTagNoun LinguisticTag = "Noun"
const LinguisticTagNumber LinguisticTag = "Number"
const LinguisticTagOpenParenthesis LinguisticTag = "OpenParenthesis"
const LinguisticTagOpenQuote LinguisticTag = "OpenQuote"
const LinguisticTagOrganizationName LinguisticTag = "OrganizationName"
const LinguisticTagOther LinguisticTag = "Other"
const LinguisticTagOtherPunctuation LinguisticTag = "Punctuation"
const LinguisticTagOtherWhitespace LinguisticTag = "Whitespace"
const LinguisticTagOtherWord LinguisticTag = "OtherWord"
const LinguisticTagParagraphBreak LinguisticTag = "ParagraphBreak"
const LinguisticTagParticle LinguisticTag = "Particle"
const LinguisticTagPersonalName LinguisticTag = "PersonalName"
const LinguisticTagPlaceName LinguisticTag = "PlaceName"
const LinguisticTagPreposition LinguisticTag = "Preposition"
const LinguisticTagPronoun LinguisticTag = "Pronoun"
const LinguisticTagPunctuation LinguisticTag = "Punctuation"
const LinguisticTagSentenceTerminator LinguisticTag = "SentenceTerminator"
const LinguisticTagVerb LinguisticTag = "Verb"
const LinguisticTagWhitespace LinguisticTag = "Whitespace"
const LinguisticTagWord LinguisticTag = "Word"
const LinguisticTagWordJoiner LinguisticTag = "WordJoiner"

type LinguisticTagScheme string

const LinguisticTagSchemeLanguage LinguisticTagScheme = "Language"
const LinguisticTagSchemeLemma LinguisticTagScheme = "Lemma"
const LinguisticTagSchemeLexicalClass LinguisticTagScheme = "LexicalClass"
const LinguisticTagSchemeNameType LinguisticTagScheme = "NameType"
const LinguisticTagSchemeNameTypeOrLexicalClass LinguisticTagScheme = "NameTypeOrLexicalClass"
const LinguisticTagSchemeScript LinguisticTagScheme = "Script"
const LinguisticTagSchemeTokenType LinguisticTagScheme = "TokenType"

type LinguisticTaggerOptions uint

const LinguisticTaggerJoinNames LinguisticTaggerOptions = 16
const LinguisticTaggerOmitOther LinguisticTaggerOptions = 8
const LinguisticTaggerOmitPunctuation LinguisticTaggerOptions = 2
const LinguisticTaggerOmitWhitespace LinguisticTaggerOptions = 4
const LinguisticTaggerOmitWords LinguisticTaggerOptions = 1

type LinguisticTaggerUnit int

const LinguisticTaggerUnitDocument LinguisticTaggerUnit = 3
const LinguisticTaggerUnitParagraph LinguisticTaggerUnit = 2
const LinguisticTaggerUnitSentence LinguisticTaggerUnit = 1
const LinguisticTaggerUnitWord LinguisticTaggerUnit = 0

type LocaleKey string

const LocaleAlternateQuotationBeginDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationBeginDelimiterKey"
const LocaleAlternateQuotationEndDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationEndDelimiterKey"
const LocaleCalendar LocaleKey = "kCFLocaleCalendarKey"
const LocaleCollationIdentifier LocaleKey = "collation"
const LocaleCollatorIdentifier LocaleKey = "kCFLocaleCollatorIdentifierKey"
const LocaleCountryCode LocaleKey = "kCFLocaleCountryCodeKey"
const LocaleCurrencyCode LocaleKey = "currency"
const LocaleCurrencySymbol LocaleKey = "kCFLocaleCurrencySymbolKey"
const LocaleDecimalSeparator LocaleKey = "kCFLocaleDecimalSeparatorKey"
const LocaleExemplarCharacterSet LocaleKey = "kCFLocaleExemplarCharacterSetKey"
const LocaleGroupingSeparator LocaleKey = "kCFLocaleGroupingSeparatorKey"
const LocaleIdentifier LocaleKey = "kCFLocaleIdentifierKey"
const LocaleLanguageCode LocaleKey = "kCFLocaleLanguageCodeKey"
const LocaleMeasurementSystem LocaleKey = "kCFLocaleMeasurementSystemKey"
const LocaleQuotationBeginDelimiterKey LocaleKey = "kCFLocaleQuotationBeginDelimiterKey"
const LocaleQuotationEndDelimiterKey LocaleKey = "kCFLocaleQuotationEndDelimiterKey"
const LocaleScriptCode LocaleKey = "kCFLocaleScriptCodeKey"
const LocaleUsesMetricSystem LocaleKey = "kCFLocaleUsesMetricSystemKey"
const LocaleVariantCode LocaleKey = "kCFLocaleVariantCodeKey"

type LocaleLanguageDirection uint

const LocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 4
const LocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 1
const LocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 2
const LocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 3
const LocaleLanguageDirectionUnknown LocaleLanguageDirection = 0

type MachPortOptions uint

const MachPortDeallocateNone MachPortOptions = 0
const MachPortDeallocateReceiveRight MachPortOptions = 2
const MachPortDeallocateSendRight MachPortOptions = 1

type MapTableOptions uint

type MassFormatterUnit int

const MassFormatterUnitGram MassFormatterUnit = 11
const MassFormatterUnitKilogram MassFormatterUnit = 14
const MassFormatterUnitOunce MassFormatterUnit = 1537
const MassFormatterUnitPound MassFormatterUnit = 1538
const MassFormatterUnitStone MassFormatterUnit = 1539

type MatchingFlags uint

const MatchingCompleted MatchingFlags = 2
const MatchingHitEnd MatchingFlags = 4
const MatchingInternalError MatchingFlags = 16
const MatchingProgress MatchingFlags = 1
const MatchingRequiredEnd MatchingFlags = 8

type MatchingOptions uint

const MatchingAnchored MatchingOptions = 4
const MatchingReportCompletion MatchingOptions = 2
const MatchingReportProgress MatchingOptions = 1
const MatchingWithTransparentBounds MatchingOptions = 8
const MatchingWithoutAnchoringBounds MatchingOptions = 16

type MeasurementFormatterUnitOptions uint

const MeasurementFormatterUnitOptionsNaturalScale MeasurementFormatterUnitOptions = 2
const MeasurementFormatterUnitOptionsProvidedUnit MeasurementFormatterUnitOptions = 1
const MeasurementFormatterUnitOptionsTemperatureWithoutUnit MeasurementFormatterUnitOptions = 4

type NetServiceOptions uint

const NetServiceListenForConnections NetServiceOptions = 2
const NetServiceNoAutoRename NetServiceOptions = 1

type NetServicesError int

const NetServicesActivityInProgress NetServicesError = -72003
const NetServicesBadArgumentError NetServicesError = -72004
const NetServicesCancelledError NetServicesError = -72005
const NetServicesCollisionError NetServicesError = -72001
const NetServicesInvalidError NetServicesError = -72006
const NetServicesMissingRequiredConfigurationError NetServicesError = -72008
const NetServicesNotFoundError NetServicesError = -72002
const NetServicesTimeoutError NetServicesError = -72007
const NetServicesUnknownError NetServicesError = -72000

type NotificationCoalescing uint

const NotificationCoalescingOnName NotificationCoalescing = 1
const NotificationCoalescingOnSender NotificationCoalescing = 2
const NotificationNoCoalescing NotificationCoalescing = 0

type NotificationName string

const AppleEventManagerWillProcessFirstEventNotification NotificationName = "NSAppleEventManagerWillProcessFirstEvent"
const BundleDidLoadNotification NotificationName = "NSBundleDidLoadNotification"
const CalendarDayChangedNotification NotificationName = "NSCalendarDayChangedNotification"
const ClassDescriptionNeededForClassNotification NotificationName = "NSClassDescriptionNeededForClassNotification"
const CurrentLocaleDidChangeNotification NotificationName = "kCFLocaleCurrentLocaleDidChangeNotification"
const DidBecomeSingleThreadedNotification NotificationName = "NSDidBecomeSingleThreadedNotification"
const FileHandleConnectionAcceptedNotification NotificationName = "NSFileHandleConnectionAcceptedNotification"
const FileHandleDataAvailableNotification NotificationName = "NSFileHandleDataAvailableNotification"
const FileHandleReadCompletionNotification NotificationName = "NSFileHandleReadCompletionNotification"
const FileHandleReadToEndOfFileCompletionNotification NotificationName = "NSFileHandleReadToEndOfFileCompletionNotification"
const HTTPCookieManagerAcceptPolicyChangedNotification NotificationName = "com.apple.Foundation.NSHTTPCookieManagerAcceptPolicyChanged"
const HTTPCookieManagerCookiesChangedNotification NotificationName = "NSHTTPCookieManagerCookiesChangedNotification"
const MetadataQueryDidFinishGatheringNotification NotificationName = "NSMetadataQueryDidFinishGatheringNotification"
const MetadataQueryDidStartGatheringNotification NotificationName = "NSMetadataQueryDidStartGatheringNotification"
const MetadataQueryDidUpdateNotification NotificationName = "NSMetadataQueryDidUpdateNotification"
const MetadataQueryGatheringProgressNotification NotificationName = "NSMetadataQueryGatheringProgressNotification"
const PortDidBecomeInvalidNotification NotificationName = "NSPortDidBecomeInvalidNotification"
const ProcessInfoPowerStateDidChangeNotification NotificationName = "NSProcessInfoPowerStateDidChangeNotification"
const ProcessInfoThermalStateDidChangeNotification NotificationName = "NSProcessInfoThermalStateDidChangeNotification"
const SystemClockDidChangeNotification NotificationName = "NSSystemClockDidChangeNotification"
const SystemTimeZoneDidChangeNotification NotificationName = "kCFTimeZoneSystemTimeZoneDidChangeNotification"
const TaskDidTerminateNotification NotificationName = "NSTaskDidTerminateNotification"
const ThreadWillExitNotification NotificationName = "NSThreadWillExitNotification"
const URLCredentialStorageChangedNotification NotificationName = "NSURLCredentialStorageChangedNotification"
const UbiquitousKeyValueStoreDidChangeExternallyNotification NotificationName = "NSUbiquitousKeyValueStoreDidChangeExternallyNotification"
const UbiquityIdentityDidChangeNotification NotificationName = "NSUbiquityIdentityDidChangeNotification"
const UndoManagerCheckpointNotification NotificationName = "NSUndoManagerCheckpointNotification"
const UndoManagerDidCloseUndoGroupNotification NotificationName = "NSUndoManagerDidCloseUndoGroupNotification"
const UndoManagerDidOpenUndoGroupNotification NotificationName = "NSUndoManagerDidOpenUndoGroupNotification"
const UndoManagerDidRedoChangeNotification NotificationName = "NSUndoManagerDidRedoChangeNotification"
const UndoManagerDidUndoChangeNotification NotificationName = "NSUndoManagerDidUndoChangeNotification"
const UndoManagerWillCloseUndoGroupNotification NotificationName = "NSUndoManagerWillCloseUndoGroupNotification"
const UndoManagerWillRedoChangeNotification NotificationName = "NSUndoManagerWillRedoChangeNotification"
const UndoManagerWillUndoChangeNotification NotificationName = "NSUndoManagerWillUndoChangeNotification"
const UserDefaultsDidChangeNotification NotificationName = "NSUserDefaultsDidChangeNotification"
const WillBecomeMultiThreadedNotification NotificationName = "NSWillBecomeMultiThreadedNotification"

type NotificationSuspensionBehavior uint

const NotificationSuspensionBehaviorCoalesce NotificationSuspensionBehavior = 2
const NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 4
const NotificationSuspensionBehaviorDrop NotificationSuspensionBehavior = 1
const NotificationSuspensionBehaviorHold NotificationSuspensionBehavior = 3

type NumberFormatterBehavior uint

const NumberFormatterBehavior10_0 NumberFormatterBehavior = 1000
const NumberFormatterBehavior10_4 NumberFormatterBehavior = 1040
const NumberFormatterBehaviorDefault NumberFormatterBehavior = 0

type NumberFormatterPadPosition uint

const NumberFormatterPadAfterPrefix NumberFormatterPadPosition = 1
const NumberFormatterPadAfterSuffix NumberFormatterPadPosition = 3
const NumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
const NumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 2

type NumberFormatterRoundingMode uint

const NumberFormatterRoundCeiling NumberFormatterRoundingMode = 0
const NumberFormatterRoundDown NumberFormatterRoundingMode = 2
const NumberFormatterRoundFloor NumberFormatterRoundingMode = 1
const NumberFormatterRoundHalfDown NumberFormatterRoundingMode = 5
const NumberFormatterRoundHalfEven NumberFormatterRoundingMode = 4
const NumberFormatterRoundHalfUp NumberFormatterRoundingMode = 6
const NumberFormatterRoundUp NumberFormatterRoundingMode = 3

type NumberFormatterStyle uint

const NumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 10
const NumberFormatterCurrencyISOCodeStyle NumberFormatterStyle = 8
const NumberFormatterCurrencyPluralStyle NumberFormatterStyle = 9
const NumberFormatterCurrencyStyle NumberFormatterStyle = 2
const NumberFormatterDecimalStyle NumberFormatterStyle = 1
const NumberFormatterNoStyle NumberFormatterStyle = 0
const NumberFormatterOrdinalStyle NumberFormatterStyle = 6
const NumberFormatterPercentStyle NumberFormatterStyle = 3
const NumberFormatterScientificStyle NumberFormatterStyle = 4
const NumberFormatterSpellOutStyle NumberFormatterStyle = 5

type OperationQueuePriority int

const OperationQueuePriorityHigh OperationQueuePriority = 4
const OperationQueuePriorityLow OperationQueuePriority = -4
const OperationQueuePriorityNormal OperationQueuePriority = 0
const OperationQueuePriorityVeryHigh OperationQueuePriority = 8
const OperationQueuePriorityVeryLow OperationQueuePriority = -8

type OrderedCollectionDifferenceCalculationOptions uint

const OrderedCollectionDifferenceCalculationInferMoves OrderedCollectionDifferenceCalculationOptions = 4
const OrderedCollectionDifferenceCalculationOmitInsertedObjects OrderedCollectionDifferenceCalculationOptions = 1
const OrderedCollectionDifferenceCalculationOmitRemovedObjects OrderedCollectionDifferenceCalculationOptions = 2

type PersonNameComponentsFormatterOptions uint

const PersonNameComponentsFormatterPhonetic PersonNameComponentsFormatterOptions = 2

type PersonNameComponentsFormatterStyle int

const PersonNameComponentsFormatterStyleAbbreviated PersonNameComponentsFormatterStyle = 4
const PersonNameComponentsFormatterStyleDefault PersonNameComponentsFormatterStyle = 0
const PersonNameComponentsFormatterStyleLong PersonNameComponentsFormatterStyle = 3
const PersonNameComponentsFormatterStyleMedium PersonNameComponentsFormatterStyle = 2
const PersonNameComponentsFormatterStyleShort PersonNameComponentsFormatterStyle = 1

type PointerFunctionsOptions uint

const HashTableCopyIn PointerFunctionsOptions = 65536
const HashTableObjectPointerPersonality PointerFunctionsOptions = 512
const HashTableStrongMemory PointerFunctionsOptions = 0
const HashTableWeakMemory PointerFunctionsOptions = 5
const HashTableZeroingWeakMemory PointerFunctionsOptions = 1
const MapTableCopyIn PointerFunctionsOptions = 65536
const MapTableObjectPointerPersonality PointerFunctionsOptions = 512
const MapTableStrongMemory PointerFunctionsOptions = 0
const MapTableWeakMemory PointerFunctionsOptions = 5
const MapTableZeroingWeakMemory PointerFunctionsOptions = 1
const PointerFunctionsCStringPersonality PointerFunctionsOptions = 768
const PointerFunctionsCopyIn PointerFunctionsOptions = 65536
const PointerFunctionsIntegerPersonality PointerFunctionsOptions = 1280
const PointerFunctionsMachVirtualMemory PointerFunctionsOptions = 4
const PointerFunctionsMallocMemory PointerFunctionsOptions = 3
const PointerFunctionsObjectPersonality PointerFunctionsOptions = 0
const PointerFunctionsObjectPointerPersonality PointerFunctionsOptions = 512
const PointerFunctionsOpaqueMemory PointerFunctionsOptions = 2
const PointerFunctionsOpaquePersonality PointerFunctionsOptions = 256
const PointerFunctionsStrongMemory PointerFunctionsOptions = 0
const PointerFunctionsStructPersonality PointerFunctionsOptions = 1024
const PointerFunctionsWeakMemory PointerFunctionsOptions = 5
const PointerFunctionsZeroingWeakMemory PointerFunctionsOptions = 1

type PostingStyle uint

const PostASAP PostingStyle = 2
const PostNow PostingStyle = 3
const PostWhenIdle PostingStyle = 1

type PredicateOperatorType uint

const BeginsWithPredicateOperatorType PredicateOperatorType = 8
const BetweenPredicateOperatorType PredicateOperatorType = 100
const ContainsPredicateOperatorType PredicateOperatorType = 99
const CustomSelectorPredicateOperatorType PredicateOperatorType = 11
const EndsWithPredicateOperatorType PredicateOperatorType = 9
const EqualToPredicateOperatorType PredicateOperatorType = 4
const GreaterThanOrEqualToPredicateOperatorType PredicateOperatorType = 3
const GreaterThanPredicateOperatorType PredicateOperatorType = 2
const InPredicateOperatorType PredicateOperatorType = 10
const LessThanOrEqualToPredicateOperatorType PredicateOperatorType = 1
const LessThanPredicateOperatorType PredicateOperatorType = 0
const LikePredicateOperatorType PredicateOperatorType = 7
const MatchesPredicateOperatorType PredicateOperatorType = 6
const NotEqualToPredicateOperatorType PredicateOperatorType = 5

type PresentationIntentKind int

const PresentationIntentKindBlockQuote PresentationIntentKind = 6
const PresentationIntentKindCodeBlock PresentationIntentKind = 5
const PresentationIntentKindHeader PresentationIntentKind = 1
const PresentationIntentKindListItem PresentationIntentKind = 4
const PresentationIntentKindOrderedList PresentationIntentKind = 2
const PresentationIntentKindParagraph PresentationIntentKind = 0
const PresentationIntentKindTable PresentationIntentKind = 8
const PresentationIntentKindTableCell PresentationIntentKind = 11
const PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
const PresentationIntentKindTableRow PresentationIntentKind = 10
const PresentationIntentKindThematicBreak PresentationIntentKind = 7
const PresentationIntentKindUnorderedList PresentationIntentKind = 3

type PresentationIntentTableColumnAlignment int

const PresentationIntentTableColumnAlignmentCenter PresentationIntentTableColumnAlignment = 1
const PresentationIntentTableColumnAlignmentLeft PresentationIntentTableColumnAlignment = 0
const PresentationIntentTableColumnAlignmentRight PresentationIntentTableColumnAlignment = 2

type ProcessInfoThermalState int

const ProcessInfoThermalStateCritical ProcessInfoThermalState = 3
const ProcessInfoThermalStateFair ProcessInfoThermalState = 1
const ProcessInfoThermalStateNominal ProcessInfoThermalState = 0
const ProcessInfoThermalStateSerious ProcessInfoThermalState = 2

type ProgressFileOperationKind string

const ProgressFileOperationKindCopying ProgressFileOperationKind = "NSProgressFileOperationKindCopying"
const ProgressFileOperationKindDecompressingAfterDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDecompressingAfterDownloading"
const ProgressFileOperationKindDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDownloading"
const ProgressFileOperationKindDuplicating ProgressFileOperationKind = "NSProgressFileOperationKindDuplicating"
const ProgressFileOperationKindReceiving ProgressFileOperationKind = "NSProgressFileOperationKindReceiving"
const ProgressFileOperationKindUploading ProgressFileOperationKind = "NSProgressFileOperationKindUploading"

type ProgressKind string

const ProgressKindFile ProgressKind = "NSProgressKindFile"

type ProgressUserInfoKey string

const ProgressEstimatedTimeRemainingKey ProgressUserInfoKey = "NSProgressEstimatedTimeRemainingKey"
const ProgressFileAnimationImageKey ProgressUserInfoKey = "NSProgressFlyToImageKey"
const ProgressFileAnimationImageOriginalRectKey ProgressUserInfoKey = "NSProgressFileAnimationImageOriginalRectKey"
const ProgressFileCompletedCountKey ProgressUserInfoKey = "NSProgressFileCompletedCountKey"
const ProgressFileIconKey ProgressUserInfoKey = "NSProgressFileIconKey"
const ProgressFileOperationKindKey ProgressUserInfoKey = "NSProgressFileOperationKindKey"
const ProgressFileTotalCountKey ProgressUserInfoKey = "NSProgressFileTotalCountKey"
const ProgressFileURLKey ProgressUserInfoKey = "NSProgressFileURLKey"
const ProgressThroughputKey ProgressUserInfoKey = "NSProgressThroughputKey"

type PropertyListFormat uint

const PropertyListBinaryFormat_v1_0 PropertyListFormat = 200
const PropertyListOpenStepFormat PropertyListFormat = 1
const PropertyListXMLFormat_v1_0 PropertyListFormat = 100

type PropertyListMutabilityOptions uint

const PropertyListImmutable PropertyListMutabilityOptions = 0
const PropertyListMutableContainers PropertyListMutabilityOptions = 1
const PropertyListMutableContainersAndLeaves PropertyListMutabilityOptions = 2

type PropertyListWriteOptions uint

type QualityOfService int

const QualityOfServiceBackground QualityOfService = 9
const QualityOfServiceDefault QualityOfService = -1
const QualityOfServiceUserInitiated QualityOfService = 25
const QualityOfServiceUserInteractive QualityOfService = 33
const QualityOfServiceUtility QualityOfService = 17

type RectEdge uint

const RectEdgeMaxX RectEdge = 2
const RectEdgeMaxY RectEdge = 3
const RectEdgeMinX RectEdge = 0
const RectEdgeMinY RectEdge = 1

type RegularExpressionOptions uint

const RegularExpressionAllowCommentsAndWhitespace RegularExpressionOptions = 2
const RegularExpressionAnchorsMatchLines RegularExpressionOptions = 16
const RegularExpressionCaseInsensitive RegularExpressionOptions = 1
const RegularExpressionDotMatchesLineSeparators RegularExpressionOptions = 8
const RegularExpressionIgnoreMetacharacters RegularExpressionOptions = 4
const RegularExpressionUseUnicodeWordBoundaries RegularExpressionOptions = 64
const RegularExpressionUseUnixLineSeparators RegularExpressionOptions = 32

type RelativeDateTimeFormatterStyle int

const RelativeDateTimeFormatterStyleNamed RelativeDateTimeFormatterStyle = 1
const RelativeDateTimeFormatterStyleNumeric RelativeDateTimeFormatterStyle = 0

type RelativeDateTimeFormatterUnitsStyle int

const RelativeDateTimeFormatterUnitsStyleAbbreviated RelativeDateTimeFormatterUnitsStyle = 3
const RelativeDateTimeFormatterUnitsStyleFull RelativeDateTimeFormatterUnitsStyle = 0
const RelativeDateTimeFormatterUnitsStyleShort RelativeDateTimeFormatterUnitsStyle = 2
const RelativeDateTimeFormatterUnitsStyleSpellOut RelativeDateTimeFormatterUnitsStyle = 1

type RelativePosition uint

const RelativeAfter RelativePosition = 0
const RelativeBefore RelativePosition = 1

type RoundingMode uint

const RoundBankers RoundingMode = 3
const RoundDown RoundingMode = 1
const RoundPlain RoundingMode = 0
const RoundUp RoundingMode = 2

type RunLoopMode string

const DefaultRunLoopMode RunLoopMode = "kCFRunLoopDefaultMode"
const RunLoopCommonModes RunLoopMode = "kCFRunLoopCommonModes"

type SaveOptions uint

const SaveOptionsAsk SaveOptions = 2
const SaveOptionsNo SaveOptions = 1
const SaveOptionsYes SaveOptions = 0

type SearchPathDirectory uint

const AdminApplicationDirectory SearchPathDirectory = 4
const AllApplicationsDirectory SearchPathDirectory = 100
const AllLibrariesDirectory SearchPathDirectory = 101
const ApplicationDirectory SearchPathDirectory = 1
const ApplicationScriptsDirectory SearchPathDirectory = 23
const ApplicationSupportDirectory SearchPathDirectory = 14
const AutosavedInformationDirectory SearchPathDirectory = 11
const CachesDirectory SearchPathDirectory = 13
const CoreServiceDirectory SearchPathDirectory = 10
const DemoApplicationDirectory SearchPathDirectory = 2
const DesktopDirectory SearchPathDirectory = 12
const DeveloperApplicationDirectory SearchPathDirectory = 3
const DeveloperDirectory SearchPathDirectory = 6
const DocumentDirectory SearchPathDirectory = 9
const DocumentationDirectory SearchPathDirectory = 8
const DownloadsDirectory SearchPathDirectory = 15
const InputMethodsDirectory SearchPathDirectory = 16
const ItemReplacementDirectory SearchPathDirectory = 99
const LibraryDirectory SearchPathDirectory = 5
const MoviesDirectory SearchPathDirectory = 17
const MusicDirectory SearchPathDirectory = 18
const PicturesDirectory SearchPathDirectory = 19
const PreferencePanesDirectory SearchPathDirectory = 22
const PrinterDescriptionDirectory SearchPathDirectory = 20
const SharedPublicDirectory SearchPathDirectory = 21
const TrashDirectory SearchPathDirectory = 102
const UserDirectory SearchPathDirectory = 7

type SearchPathDomainMask uint

const AllDomainsMask SearchPathDomainMask = 65535
const LocalDomainMask SearchPathDomainMask = 2
const NetworkDomainMask SearchPathDomainMask = 4
const SystemDomainMask SearchPathDomainMask = 8
const UserDomainMask SearchPathDomainMask = 1

type SocketNativeHandle int

type SortOptions uint

const SortConcurrent SortOptions = 1
const SortStable SortOptions = 16

type StreamEvent uint

const StreamEventEndEncountered StreamEvent = 16
const StreamEventErrorOccurred StreamEvent = 8
const StreamEventHasBytesAvailable StreamEvent = 2
const StreamEventHasSpaceAvailable StreamEvent = 4
const StreamEventNone StreamEvent = 0
const StreamEventOpenCompleted StreamEvent = 1

type StreamNetworkServiceTypeValue string

const StreamNetworkServiceTypeBackground StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeBackground"
const StreamNetworkServiceTypeCallSignaling StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeCallSignaling"
const StreamNetworkServiceTypeVideo StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVideo"
const StreamNetworkServiceTypeVoIP StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVoIP"
const StreamNetworkServiceTypeVoice StreamNetworkServiceTypeValue = "kCFStreamNetworkServiceTypeVoice"

type StreamPropertyKey string

const StreamDataWrittenToMemoryStreamKey StreamPropertyKey = "kCFStreamPropertyDataWritten"
const StreamFileCurrentOffsetKey StreamPropertyKey = "kCFStreamPropertyFileCurrentOffset"
const StreamNetworkServiceType StreamPropertyKey = "kCFStreamNetworkServiceType"
const StreamSOCKSProxyConfigurationKey StreamPropertyKey = "kCFStreamPropertySOCKSProxy"
const StreamSocketSecurityLevelKey StreamPropertyKey = "kCFStreamPropertySocketSecurityLevel"

type StreamSOCKSProxyConfiguration string

const StreamSOCKSProxyHostKey StreamSOCKSProxyConfiguration = "SOCKSProxy"
const StreamSOCKSProxyPasswordKey StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSPassword"
const StreamSOCKSProxyPortKey StreamSOCKSProxyConfiguration = "SOCKSPort"
const StreamSOCKSProxyUserKey StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSUser"
const StreamSOCKSProxyVersionKey StreamSOCKSProxyConfiguration = "kCFStreamPropertySOCKSVersion"

type StreamSOCKSProxyVersion string

const StreamSOCKSProxyVersion4 StreamSOCKSProxyVersion = "kCFStreamSocketSOCKSVersion4"
const StreamSOCKSProxyVersion5 StreamSOCKSProxyVersion = "kCFStreamSocketSOCKSVersion5"

type StreamSocketSecurityLevel string

const StreamSocketSecurityLevelNegotiatedSSL StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelNegotiatedSSL"
const StreamSocketSecurityLevelNone StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelNone"
const StreamSocketSecurityLevelSSLv2 StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelSSLv2"
const StreamSocketSecurityLevelSSLv3 StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelSSLv3"
const StreamSocketSecurityLevelTLSv1 StreamSocketSecurityLevel = "kCFStreamSocketSecurityLevelTLSv1"

type StreamStatus uint

const StreamStatusAtEnd StreamStatus = 5
const StreamStatusClosed StreamStatus = 6
const StreamStatusError StreamStatus = 7
const StreamStatusNotOpen StreamStatus = 0
const StreamStatusOpen StreamStatus = 2
const StreamStatusOpening StreamStatus = 1
const StreamStatusReading StreamStatus = 3
const StreamStatusWriting StreamStatus = 4

type StringCompareOptions uint

const AnchoredSearch StringCompareOptions = 8
const BackwardsSearch StringCompareOptions = 4
const CaseInsensitiveSearch StringCompareOptions = 1
const DiacriticInsensitiveSearch StringCompareOptions = 128
const ForcedOrderingSearch StringCompareOptions = 512
const LiteralSearch StringCompareOptions = 2
const NumericSearch StringCompareOptions = 64
const RegularExpressionSearch StringCompareOptions = 1024
const WidthInsensitiveSearch StringCompareOptions = 256

type StringEncoding uint

type StringEncodingConversionOptions uint

const StringEncodingConversionAllowLossy StringEncodingConversionOptions = 1
const StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 2

type StringEncodingDetectionOptionsKey string

const StringEncodingDetectionAllowLossyKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionAllowLossyKey"
const StringEncodingDetectionDisallowedEncodingsKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionDisallowedEncodingsKey"
const StringEncodingDetectionFromWindowsKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionFromWindowsKey"
const StringEncodingDetectionLikelyLanguageKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionLikelyLanguageKey"
const StringEncodingDetectionLossySubstitutionKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionLossySubstitutionKey"
const StringEncodingDetectionSuggestedEncodingsKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionSuggestedEncodingsKey"
const StringEncodingDetectionUseOnlySuggestedEncodingsKey StringEncodingDetectionOptionsKey = "NSStringEncodingDetectionUseOnlySuggestedEncodingsKey"

type StringEnumerationOptions uint

const StringEnumerationByCaretPositions StringEnumerationOptions = 5
const StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
const StringEnumerationByDeletionClusters StringEnumerationOptions = 6
const StringEnumerationByLines StringEnumerationOptions = 0
const StringEnumerationByParagraphs StringEnumerationOptions = 1
const StringEnumerationBySentences StringEnumerationOptions = 4
const StringEnumerationByWords StringEnumerationOptions = 3
const StringEnumerationLocalized StringEnumerationOptions = 1024
const StringEnumerationReverse StringEnumerationOptions = 256
const StringEnumerationSubstringNotRequired StringEnumerationOptions = 512

type StringTransform string

const StringTransformFullwidthToHalfwidth StringTransform = ")kCFStringTransformFullwidthHalfwidth"
const StringTransformHiraganaToKatakana StringTransform = ")kCFStringTransformHiraganaKatakana"
const StringTransformLatinToArabic StringTransform = ")kCFStringTransformLatinArabic"
const StringTransformLatinToCyrillic StringTransform = ")kCFStringTransformLatinCyrillic"
const StringTransformLatinToGreek StringTransform = ")kCFStringTransformLatinGreek"
const StringTransformLatinToHangul StringTransform = ")kCFStringTransformLatinHangul"
const StringTransformLatinToHebrew StringTransform = ")kCFStringTransformLatinHebrew"
const StringTransformLatinToHiragana StringTransform = ")kCFStringTransformLatinHiragana"
const StringTransformLatinToKatakana StringTransform = ")kCFStringTransformLatinKatakana"
const StringTransformLatinToThai StringTransform = ")kCFStringTransformLatinThai"
const StringTransformMandarinToLatin StringTransform = ")kCFStringTransformMandarinLatin"
const StringTransformStripCombiningMarks StringTransform = ")kCFStringTransformStripCombiningMarks"
const StringTransformStripDiacritics StringTransform = ")kCFStringTransformStripDiacritics"
const StringTransformToLatin StringTransform = ")kCFStringTransformToLatin"
const StringTransformToUnicodeName StringTransform = ")kCFStringTransformToUnicodeName"
const StringTransformToXMLHex StringTransform = ")kCFStringTransformToXMLHex"

type TaskTerminationReason int

const TaskTerminationReasonExit TaskTerminationReason = 1
const TaskTerminationReasonUncaughtSignal TaskTerminationReason = 2

type TestComparisonOperation uint

const BeginsWithComparison TestComparisonOperation = 5
const ContainsComparison TestComparisonOperation = 7
const EndsWithComparison TestComparisonOperation = 6
const EqualToComparison TestComparisonOperation = 0
const GreaterThanComparison TestComparisonOperation = 4
const GreaterThanOrEqualToComparison TestComparisonOperation = 3
const LessThanComparison TestComparisonOperation = 2
const LessThanOrEqualToComparison TestComparisonOperation = 1

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

const TextCheckingTypeAddress TextCheckingType = 16
const TextCheckingTypeCorrection TextCheckingType = 512
const TextCheckingTypeDash TextCheckingType = 128
const TextCheckingTypeDate TextCheckingType = 8
const TextCheckingTypeGrammar TextCheckingType = 4
const TextCheckingTypeLink TextCheckingType = 32
const TextCheckingTypeOrthography TextCheckingType = 1
const TextCheckingTypePhoneNumber TextCheckingType = 2048
const TextCheckingTypeQuote TextCheckingType = 64
const TextCheckingTypeRegularExpression TextCheckingType = 1024
const TextCheckingTypeReplacement TextCheckingType = 256
const TextCheckingTypeSpelling TextCheckingType = 2
const TextCheckingTypeTransitInformation TextCheckingType = 4096

type TextCheckingTypes uint64

type TimeInterval float64

type TimeZoneNameStyle int

const TimeZoneNameStyleDaylightSaving TimeZoneNameStyle = 2
const TimeZoneNameStyleGeneric TimeZoneNameStyle = 4
const TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
const TimeZoneNameStyleShortGeneric TimeZoneNameStyle = 5
const TimeZoneNameStyleShortStandard TimeZoneNameStyle = 1
const TimeZoneNameStyleStandard TimeZoneNameStyle = 0

type URLBookmarkCreationOptions uint

const URLBookmarkCreationMinimalBookmark URLBookmarkCreationOptions = 512
const URLBookmarkCreationPreferFileIDResolution URLBookmarkCreationOptions = 256
const URLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 4096
const URLBookmarkCreationSuitableForBookmarkFile URLBookmarkCreationOptions = 1024
const URLBookmarkCreationWithSecurityScope URLBookmarkCreationOptions = 2048
const URLBookmarkCreationWithoutImplicitSecurityScope URLBookmarkCreationOptions = 536870912

type URLBookmarkFileCreationOptions uint

type URLBookmarkResolutionOptions uint

const URLBookmarkResolutionWithSecurityScope URLBookmarkResolutionOptions = 1024
const URLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 32768
const URLBookmarkResolutionWithoutMounting URLBookmarkResolutionOptions = 512
const URLBookmarkResolutionWithoutUI URLBookmarkResolutionOptions = 256

type URLCacheStoragePolicy uint

const URLCacheStorageAllowed URLCacheStoragePolicy = 0
const URLCacheStorageAllowedInMemoryOnly URLCacheStoragePolicy = 1
const URLCacheStorageNotAllowed URLCacheStoragePolicy = 2

type URLCredentialPersistence uint

const URLCredentialPersistenceForSession URLCredentialPersistence = 1
const URLCredentialPersistenceNone URLCredentialPersistence = 0
const URLCredentialPersistencePermanent URLCredentialPersistence = 2
const URLCredentialPersistenceSynchronizable URLCredentialPersistence = 3

type URLErrorNetworkUnavailableReason int

const URLErrorNetworkUnavailableReasonCellular URLErrorNetworkUnavailableReason = 0
const URLErrorNetworkUnavailableReasonConstrained URLErrorNetworkUnavailableReason = 2
const URLErrorNetworkUnavailableReasonExpensive URLErrorNetworkUnavailableReason = 1

type URLFileProtectionType string

const URLFileProtectionComplete URLFileProtectionType = "NSURLFileProtectionComplete"
const URLFileProtectionCompleteUnlessOpen URLFileProtectionType = "NSURLFileProtectionCompleteUnlessOpen"
const URLFileProtectionCompleteUntilFirstUserAuthentication URLFileProtectionType = "NSURLFileProtectionCompleteUntilFirstUserAuthentication"
const URLFileProtectionNone URLFileProtectionType = "NSURLFileProtectionNone"

type URLFileResourceType string

const URLFileResourceTypeBlockSpecial URLFileResourceType = "NSURLFileResourceTypeBlockSpecial"
const URLFileResourceTypeCharacterSpecial URLFileResourceType = "NSURLFileResourceTypeCharacterSpecial"
const URLFileResourceTypeDirectory URLFileResourceType = "NSURLFileResourceTypeDirectory"
const URLFileResourceTypeNamedPipe URLFileResourceType = "NSURLFileResourceTypeNamedPipe"
const URLFileResourceTypeRegular URLFileResourceType = "NSURLFileResourceTypeRegular"
const URLFileResourceTypeSocket URLFileResourceType = "NSURLFileResourceTypeSocket"
const URLFileResourceTypeSymbolicLink URLFileResourceType = "NSURLFileResourceTypeSymbolicLink"
const URLFileResourceTypeUnknown URLFileResourceType = "NSURLFileResourceTypeUnknown"

type URLHandleStatus uint

const URLHandleLoadFailed URLHandleStatus = 3
const URLHandleLoadInProgress URLHandleStatus = 2
const URLHandleLoadSucceeded URLHandleStatus = 1
const URLHandleNotLoaded URLHandleStatus = 0

type URLRelationship int

const URLRelationshipContains URLRelationship = 0
const URLRelationshipOther URLRelationship = 2
const URLRelationshipSame URLRelationship = 1

type URLRequestAttribution uint

const URLRequestAttributionDeveloper URLRequestAttribution = 0
const URLRequestAttributionUser URLRequestAttribution = 1

type URLRequestCachePolicy uint

const URLRequestReloadIgnoringCacheData URLRequestCachePolicy = 1
const URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
const URLRequestReloadIgnoringLocalCacheData URLRequestCachePolicy = 1
const URLRequestReloadRevalidatingCacheData URLRequestCachePolicy = 5
const URLRequestReturnCacheDataDontLoad URLRequestCachePolicy = 3
const URLRequestReturnCacheDataElseLoad URLRequestCachePolicy = 2
const URLRequestUseProtocolCachePolicy URLRequestCachePolicy = 0

type URLRequestNetworkServiceType uint

const URLNetworkServiceTypeAVStreaming URLRequestNetworkServiceType = 8
const URLNetworkServiceTypeBackground URLRequestNetworkServiceType = 3
const URLNetworkServiceTypeCallSignaling URLRequestNetworkServiceType = 11
const URLNetworkServiceTypeDefault URLRequestNetworkServiceType = 0
const URLNetworkServiceTypeResponsiveAV URLRequestNetworkServiceType = 9
const URLNetworkServiceTypeResponsiveData URLRequestNetworkServiceType = 6
const URLNetworkServiceTypeVideo URLRequestNetworkServiceType = 2
const URLNetworkServiceTypeVoIP URLRequestNetworkServiceType = 1
const URLNetworkServiceTypeVoice URLRequestNetworkServiceType = 4

type URLResourceKey string

const URLAddedToDirectoryDateKey URLResourceKey = "NSURLAddedToDirectoryDateKey"
const URLApplicationIsScriptableKey URLResourceKey = "NSURLApplicationIsScriptableKey"
const URLAttributeModificationDateKey URLResourceKey = "NSURLAttributeModificationDateKey"
const URLCanonicalPathKey URLResourceKey = "NSURLCanonicalPathKey"
const URLContentAccessDateKey URLResourceKey = "NSURLContentAccessDateKey"
const URLContentModificationDateKey URLResourceKey = "NSURLContentModificationDateKey"
const URLContentTypeKey URLResourceKey = "NSURLContentTypeKey"
const URLCreationDateKey URLResourceKey = "NSURLCreationDateKey"
const URLCustomIconKey URLResourceKey = "NSURLCustomIconKey"
const URLDocumentIdentifierKey URLResourceKey = "NSURLDocumentIdentifierKey"
const URLEffectiveIconKey URLResourceKey = "NSURLEffectiveIconKey"
const URLFileAllocatedSizeKey URLResourceKey = "NSURLFileAllocatedSizeKey"
const URLFileContentIdentifierKey URLResourceKey = "NSURLFileContentIdentifierKey"
const URLFileProtectionKey URLResourceKey = "NSURLFileProtectionKey"
const URLFileResourceIdentifierKey URLResourceKey = "NSURLFileResourceIdentifierKey"
const URLFileResourceTypeKey URLResourceKey = "NSURLFileResourceTypeKey"
const URLFileSecurityKey URLResourceKey = "NSURLFileSecurityKey"
const URLFileSizeKey URLResourceKey = "NSURLFileSizeKey"
const URLGenerationIdentifierKey URLResourceKey = "NSURLGenerationIdentifierKey"
const URLHasHiddenExtensionKey URLResourceKey = "NSURLHasHiddenExtensionKey"
const URLIsAliasFileKey URLResourceKey = "NSURLIsAliasFileKey"
const URLIsApplicationKey URLResourceKey = "_NSURLIsApplicationKey"
const URLIsDirectoryKey URLResourceKey = "NSURLIsDirectoryKey"
const URLIsExcludedFromBackupKey URLResourceKey = "NSURLIsExcludedFromBackupKey"
const URLIsExecutableKey URLResourceKey = "NSURLIsExecutableKey"
const URLIsHiddenKey URLResourceKey = "NSURLIsHiddenKey"
const URLIsMountTriggerKey URLResourceKey = "NSURLIsMountTriggerKey"
const URLIsPackageKey URLResourceKey = "NSURLIsPackageKey"
const URLIsPurgeableKey URLResourceKey = "NSURLIsPurgeableKey"
const URLIsReadableKey URLResourceKey = "NSURLIsReadableKey"
const URLIsRegularFileKey URLResourceKey = "NSURLIsRegularFileKey"
const URLIsSparseKey URLResourceKey = "NSURLIsSparseKey"
const URLIsSymbolicLinkKey URLResourceKey = "NSURLIsSymbolicLinkKey"
const URLIsSystemImmutableKey URLResourceKey = "NSURLIsSystemImmutableKey"
const URLIsUbiquitousItemKey URLResourceKey = "NSURLIsUbiquitousItemKey"
const URLIsUserImmutableKey URLResourceKey = "NSURLIsUserImmutableKey"
const URLIsVolumeKey URLResourceKey = "NSURLIsVolumeKey"
const URLIsWritableKey URLResourceKey = "NSURLIsWritableKey"
const URLKeysOfUnsetValuesKey URLResourceKey = "NSURLKeysOfUnsetValuesKey"
const URLLabelColorKey URLResourceKey = "NSURLLabelColorKey"
const URLLabelNumberKey URLResourceKey = "NSURLLabelNumberKey"
const URLLinkCountKey URLResourceKey = "NSURLLinkCountKey"
const URLLocalizedLabelKey URLResourceKey = "NSURLLocalizedLabelKey"
const URLLocalizedNameKey URLResourceKey = "NSURLLocalizedNameKey"
const URLLocalizedTypeDescriptionKey URLResourceKey = "NSURLLocalizedTypeDescriptionKey"
const URLMayHaveExtendedAttributesKey URLResourceKey = "NSURLMayHaveExtendedAttributesKey"
const URLMayShareFileContentKey URLResourceKey = "NSURLMayShareFileContentKey"
const URLNameKey URLResourceKey = "NSURLNameKey"
const URLParentDirectoryURLKey URLResourceKey = "NSURLParentDirectoryURLKey"
const URLPathKey URLResourceKey = "_NSURLPathKey"
const URLPreferredIOBlockSizeKey URLResourceKey = "NSURLPreferredIOBlockSizeKey"
const URLQuarantinePropertiesKey URLResourceKey = "NSURLQuarantinePropertiesKey"
const URLTagNamesKey URLResourceKey = "NSURLTagNamesKey"
const URLThumbnailDictionaryKey URLResourceKey = "NSURLThumbnailDictionaryKey"
const URLThumbnailKey URLResourceKey = "NSURLThumbnailKey"
const URLTotalFileAllocatedSizeKey URLResourceKey = "NSURLTotalFileAllocatedSizeKey"
const URLTotalFileSizeKey URLResourceKey = "NSURLTotalFileSizeKey"
const URLTypeIdentifierKey URLResourceKey = "NSURLTypeIdentifierKey"
const URLUbiquitousItemContainerDisplayNameKey URLResourceKey = "NSURLUbiquitousItemContainerDisplayNameKey"
const URLUbiquitousItemDownloadRequestedKey URLResourceKey = "NSURLUbiquitousItemDownloadRequestedKey"
const URLUbiquitousItemDownloadingErrorKey URLResourceKey = "NSURLUbiquitousItemDownloadingErrorKey"
const URLUbiquitousItemDownloadingStatusKey URLResourceKey = "NSURLUbiquitousItemDownloadingStatusKey"
const URLUbiquitousItemHasUnresolvedConflictsKey URLResourceKey = "NSURLUbiquitousItemHasUnresolvedConflictsKey"
const URLUbiquitousItemIsDownloadedKey URLResourceKey = "NSURLUbiquitousItemIsDownloadedKey"
const URLUbiquitousItemIsDownloadingKey URLResourceKey = "NSURLUbiquitousItemIsDownloadingKey"
const URLUbiquitousItemIsExcludedFromSyncKey URLResourceKey = "NSURLUbiquitousItemIsExcludedFromSyncKey"
const URLUbiquitousItemIsSharedKey URLResourceKey = "NSURLUbiquitousItemIsSharedKey"
const URLUbiquitousItemIsUploadedKey URLResourceKey = "NSURLUbiquitousItemIsUploadedKey"
const URLUbiquitousItemIsUploadingKey URLResourceKey = "NSURLUbiquitousItemIsUploadingKey"
const URLUbiquitousItemPercentDownloadedKey URLResourceKey = "NSURLUbiquitousItemPercentDownloadedKey"
const URLUbiquitousItemPercentUploadedKey URLResourceKey = "NSURLUbiquitousItemPercentUploadedKey"
const URLUbiquitousItemUploadingErrorKey URLResourceKey = "NSURLUbiquitousItemUploadingErrorKey"
const URLUbiquitousSharedItemCurrentUserPermissionsKey URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserPermissionsKey"
const URLUbiquitousSharedItemCurrentUserRoleKey URLResourceKey = "NSURLUbiquitousSharedItemCurrentUserRoleKey"
const URLUbiquitousSharedItemMostRecentEditorNameComponentsKey URLResourceKey = "NSURLUbiquitousSharedItemMostRecentEditorNameComponentsKey"
const URLUbiquitousSharedItemOwnerNameComponentsKey URLResourceKey = "NSURLUbiquitousSharedItemOwnerNameComponentsKey"
const URLVolumeAvailableCapacityForImportantUsageKey URLResourceKey = "NSURLVolumeAvailableCapacityForImportantUsageKey"
const URLVolumeAvailableCapacityForOpportunisticUsageKey URLResourceKey = "NSURLVolumeAvailableCapacityForOpportunisticUsageKey"
const URLVolumeAvailableCapacityKey URLResourceKey = "NSURLVolumeAvailableCapacityKey"
const URLVolumeCreationDateKey URLResourceKey = "NSURLVolumeCreationDateKey"
const URLVolumeIdentifierKey URLResourceKey = "NSURLVolumeIdentifierKey"
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
const URLVolumeSupportsFileProtectionKey URLResourceKey = "NSURLVolumeSupportsFileProtectionKey"
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
const URLVolumeTotalCapacityKey URLResourceKey = "NSURLVolumeTotalCapacityKey"
const URLVolumeURLForRemountingKey URLResourceKey = "NSURLVolumeURLForRemountingKey"
const URLVolumeURLKey URLResourceKey = "NSURLVolumeURLKey"
const URLVolumeUUIDStringKey URLResourceKey = "NSURLVolumeUUIDStringKey"

type URLSessionAuthChallengeDisposition int

const URLSessionAuthChallengeCancelAuthenticationChallenge URLSessionAuthChallengeDisposition = 2
const URLSessionAuthChallengePerformDefaultHandling URLSessionAuthChallengeDisposition = 1
const URLSessionAuthChallengeRejectProtectionSpace URLSessionAuthChallengeDisposition = 3
const URLSessionAuthChallengeUseCredential URLSessionAuthChallengeDisposition = 0

type URLSessionDelayedRequestDisposition int

const URLSessionDelayedRequestCancel URLSessionDelayedRequestDisposition = 2
const URLSessionDelayedRequestContinueLoading URLSessionDelayedRequestDisposition = 0
const URLSessionDelayedRequestUseNewRequest URLSessionDelayedRequestDisposition = 1

type URLSessionResponseDisposition int

const URLSessionResponseAllow URLSessionResponseDisposition = 1
const URLSessionResponseBecomeDownload URLSessionResponseDisposition = 2
const URLSessionResponseBecomeStream URLSessionResponseDisposition = 3
const URLSessionResponseCancel URLSessionResponseDisposition = 0

type URLSessionTaskMetricsDomainResolutionProtocol int

const URLSessionTaskMetricsDomainResolutionProtocolHTTPS URLSessionTaskMetricsDomainResolutionProtocol = 4
const URLSessionTaskMetricsDomainResolutionProtocolTCP URLSessionTaskMetricsDomainResolutionProtocol = 2
const URLSessionTaskMetricsDomainResolutionProtocolTLS URLSessionTaskMetricsDomainResolutionProtocol = 3
const URLSessionTaskMetricsDomainResolutionProtocolUDP URLSessionTaskMetricsDomainResolutionProtocol = 1
const URLSessionTaskMetricsDomainResolutionProtocolUnknown URLSessionTaskMetricsDomainResolutionProtocol = 0

type URLSessionTaskMetricsResourceFetchType int

const URLSessionTaskMetricsResourceFetchTypeLocalCache URLSessionTaskMetricsResourceFetchType = 3
const URLSessionTaskMetricsResourceFetchTypeNetworkLoad URLSessionTaskMetricsResourceFetchType = 1
const URLSessionTaskMetricsResourceFetchTypeServerPush URLSessionTaskMetricsResourceFetchType = 2
const URLSessionTaskMetricsResourceFetchTypeUnknown URLSessionTaskMetricsResourceFetchType = 0

type URLSessionTaskState int

const URLSessionTaskStateCanceling URLSessionTaskState = 2
const URLSessionTaskStateCompleted URLSessionTaskState = 3
const URLSessionTaskStateRunning URLSessionTaskState = 0
const URLSessionTaskStateSuspended URLSessionTaskState = 1

type URLSessionWebSocketCloseCode int

const URLSessionWebSocketCloseCodeAbnormalClosure URLSessionWebSocketCloseCode = 1006
const URLSessionWebSocketCloseCodeGoingAway URLSessionWebSocketCloseCode = 1001
const URLSessionWebSocketCloseCodeInternalServerError URLSessionWebSocketCloseCode = 1011
const URLSessionWebSocketCloseCodeInvalid URLSessionWebSocketCloseCode = 0
const URLSessionWebSocketCloseCodeInvalidFramePayloadData URLSessionWebSocketCloseCode = 1007
const URLSessionWebSocketCloseCodeMandatoryExtensionMissing URLSessionWebSocketCloseCode = 1010
const URLSessionWebSocketCloseCodeMessageTooBig URLSessionWebSocketCloseCode = 1009
const URLSessionWebSocketCloseCodeNoStatusReceived URLSessionWebSocketCloseCode = 1005
const URLSessionWebSocketCloseCodeNormalClosure URLSessionWebSocketCloseCode = 1000
const URLSessionWebSocketCloseCodePolicyViolation URLSessionWebSocketCloseCode = 1008
const URLSessionWebSocketCloseCodeProtocolError URLSessionWebSocketCloseCode = 1002
const URLSessionWebSocketCloseCodeTLSHandshakeFailure URLSessionWebSocketCloseCode = 1015
const URLSessionWebSocketCloseCodeUnsupportedData URLSessionWebSocketCloseCode = 1003

type URLSessionWebSocketMessageType int

const URLSessionWebSocketMessageTypeData URLSessionWebSocketMessageType = 0
const URLSessionWebSocketMessageTypeString URLSessionWebSocketMessageType = 1

type URLThumbnailDictionaryItem string

const Thumbnail1024x1024SizeKey URLThumbnailDictionaryItem = "NSThumbnail1024x1024SizeKey"

type URLUbiquitousItemDownloadingStatus string

const URLUbiquitousItemDownloadingStatusCurrent URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusCurrent"
const URLUbiquitousItemDownloadingStatusDownloaded URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusDownloaded"
const URLUbiquitousItemDownloadingStatusNotDownloaded URLUbiquitousItemDownloadingStatus = "NSURLUbiquitousItemDownloadingStatusNotDownloaded"

type URLUbiquitousSharedItemPermissions string

const URLUbiquitousSharedItemPermissionsReadOnly URLUbiquitousSharedItemPermissions = "NSURLUbiquitousSharedItemPermissionsReadOnly"
const URLUbiquitousSharedItemPermissionsReadWrite URLUbiquitousSharedItemPermissions = "NSURLUbiquitousSharedItemPermissionsReadWrite"

type URLUbiquitousSharedItemRole string

const URLUbiquitousSharedItemRoleOwner URLUbiquitousSharedItemRole = "NSURLUbiquitousSharedItemRoleOwner"
const URLUbiquitousSharedItemRoleParticipant URLUbiquitousSharedItemRole = "NSURLUbiquitousSharedItemRoleParticipant"

type UserActivityPersistentIdentifier string

type UserNotificationActivationType int

const UserNotificationActivationTypeActionButtonClicked UserNotificationActivationType = 2
const UserNotificationActivationTypeAdditionalActionClicked UserNotificationActivationType = 4
const UserNotificationActivationTypeContentsClicked UserNotificationActivationType = 1
const UserNotificationActivationTypeNone UserNotificationActivationType = 0
const UserNotificationActivationTypeReplied UserNotificationActivationType = 3

type ValueTransformerName string

const IsNilTransformerName ValueTransformerName = "NSIsNil"
const IsNotNilTransformerName ValueTransformerName = "NSIsNotNil"
const KeyedUnarchiveFromDataTransformerName ValueTransformerName = "NSKeyedUnarchiveFromData"
const NegateBooleanTransformerName ValueTransformerName = "NSNegateBoolean"
const SecureUnarchiveFromDataTransformerName ValueTransformerName = "NSSecureUnarchiveFromData"
const UnarchiveFromDataTransformerName ValueTransformerName = "NSUnarchiveFromData"

type VolumeEnumerationOptions uint

const VolumeEnumerationProduceFileReferenceURLs VolumeEnumerationOptions = 4
const VolumeEnumerationSkipHiddenVolumes VolumeEnumerationOptions = 2

type WhoseSubelementIdentifier uint

const EverySubelement WhoseSubelementIdentifier = 1
const IndexSubelement WhoseSubelementIdentifier = 0
const MiddleSubelement WhoseSubelementIdentifier = 2
const NoSubelement WhoseSubelementIdentifier = 4
const RandomSubelement WhoseSubelementIdentifier = 3

type XMLDTDNodeKind uint

const XMLAttributeCDATAKind XMLDTDNodeKind = 6
const XMLAttributeEntitiesKind XMLDTDNodeKind = 11
const XMLAttributeEntityKind XMLDTDNodeKind = 10
const XMLAttributeEnumerationKind XMLDTDNodeKind = 14
const XMLAttributeIDKind XMLDTDNodeKind = 7
const XMLAttributeIDRefKind XMLDTDNodeKind = 8
const XMLAttributeIDRefsKind XMLDTDNodeKind = 9
const XMLAttributeNMTokenKind XMLDTDNodeKind = 12
const XMLAttributeNMTokensKind XMLDTDNodeKind = 13
const XMLAttributeNotationKind XMLDTDNodeKind = 15
const XMLElementDeclarationAnyKind XMLDTDNodeKind = 18
const XMLElementDeclarationElementKind XMLDTDNodeKind = 20
const XMLElementDeclarationEmptyKind XMLDTDNodeKind = 17
const XMLElementDeclarationMixedKind XMLDTDNodeKind = 19
const XMLElementDeclarationUndefinedKind XMLDTDNodeKind = 16
const XMLEntityGeneralKind XMLDTDNodeKind = 1
const XMLEntityParameterKind XMLDTDNodeKind = 4
const XMLEntityParsedKind XMLDTDNodeKind = 2
const XMLEntityPredefined XMLDTDNodeKind = 5
const XMLEntityUnparsedKind XMLDTDNodeKind = 3

type XMLDocumentContentKind uint

const XMLDocumentHTMLKind XMLDocumentContentKind = 2
const XMLDocumentTextKind XMLDocumentContentKind = 3
const XMLDocumentXHTMLKind XMLDocumentContentKind = 1
const XMLDocumentXMLKind XMLDocumentContentKind = 0

type XMLNodeKind uint

const XMLAttributeDeclarationKind XMLNodeKind = 10
const XMLAttributeKind XMLNodeKind = 3
const XMLCommentKind XMLNodeKind = 6
const XMLDTDKind XMLNodeKind = 8
const XMLDocumentKind XMLNodeKind = 1
const XMLElementDeclarationKind XMLNodeKind = 11
const XMLElementKind XMLNodeKind = 2
const XMLEntityDeclarationKind XMLNodeKind = 9
const XMLInvalidKind XMLNodeKind = 0
const XMLNamespaceKind XMLNodeKind = 4
const XMLNotationDeclarationKind XMLNodeKind = 12
const XMLProcessingInstructionKind XMLNodeKind = 5
const XMLTextKind XMLNodeKind = 7

type XMLNodeOptions uint

const XMLDocumentIncludeContentTypeDeclaration XMLNodeOptions = 262144
const XMLDocumentTidyHTML XMLNodeOptions = 512
const XMLDocumentTidyXML XMLNodeOptions = 1024
const XMLDocumentValidate XMLNodeOptions = 8192
const XMLDocumentXInclude XMLNodeOptions = 65536
const XMLNodeCompactEmptyElement XMLNodeOptions = 4
const XMLNodeExpandEmptyElement XMLNodeOptions = 2
const XMLNodeIsCDATA XMLNodeOptions = 1
const XMLNodeLoadExternalEntitiesAlways XMLNodeOptions = 16384
const XMLNodeLoadExternalEntitiesNever XMLNodeOptions = 524288
const XMLNodeLoadExternalEntitiesSameOriginOnly XMLNodeOptions = 32768
const XMLNodeNeverEscapeContents XMLNodeOptions = 32
const XMLNodeOptionsNone XMLNodeOptions = 0
const XMLNodePreserveAll XMLNodeOptions = 4293918750
const XMLNodePreserveAttributeOrder XMLNodeOptions = 2097152
const XMLNodePreserveCDATA XMLNodeOptions = 16777216
const XMLNodePreserveCharacterReferences XMLNodeOptions = 134217728
const XMLNodePreserveDTD XMLNodeOptions = 67108864
const XMLNodePreserveEmptyElements XMLNodeOptions = 6
const XMLNodePreserveEntities XMLNodeOptions = 4194304
const XMLNodePreserveNamespaceOrder XMLNodeOptions = 1048576
const XMLNodePreservePrefixes XMLNodeOptions = 8388608
const XMLNodePreserveQuotes XMLNodeOptions = 24
const XMLNodePreserveWhitespace XMLNodeOptions = 33554432
const XMLNodePrettyPrint XMLNodeOptions = 131072
const XMLNodePromoteSignificantWhitespace XMLNodeOptions = 268435456
const XMLNodeUseDoubleQuotes XMLNodeOptions = 16
const XMLNodeUseSingleQuotes XMLNodeOptions = 8

type XMLParserError int

const XMLParserAttributeHasNoValueError XMLParserError = 41
const XMLParserAttributeListNotFinishedError XMLParserError = 51
const XMLParserAttributeListNotStartedError XMLParserError = 50
const XMLParserAttributeNotFinishedError XMLParserError = 40
const XMLParserAttributeNotStartedError XMLParserError = 39
const XMLParserAttributeRedefinedError XMLParserError = 42
const XMLParserCDATANotFinishedError XMLParserError = 63
const XMLParserCharacterRefAtEOFError XMLParserError = 10
const XMLParserCharacterRefInDTDError XMLParserError = 13
const XMLParserCharacterRefInEpilogError XMLParserError = 12
const XMLParserCharacterRefInPrologError XMLParserError = 11
const XMLParserCommentContainsDoubleHyphenError XMLParserError = 80
const XMLParserCommentNotFinishedError XMLParserError = 45
const XMLParserConditionalSectionNotFinishedError XMLParserError = 59
const XMLParserConditionalSectionNotStartedError XMLParserError = 58
const XMLParserDOCTYPEDeclNotFinishedError XMLParserError = 61
const XMLParserDelegateAbortedParseError XMLParserError = 512
const XMLParserDocumentStartError XMLParserError = 3
const XMLParserElementContentDeclNotFinishedError XMLParserError = 55
const XMLParserElementContentDeclNotStartedError XMLParserError = 54
const XMLParserEmptyDocumentError XMLParserError = 4
const XMLParserEncodingNotSupportedError XMLParserError = 32
const XMLParserEntityBoundaryError XMLParserError = 90
const XMLParserEntityIsExternalError XMLParserError = 29
const XMLParserEntityIsParameterError XMLParserError = 30
const XMLParserEntityNotFinishedError XMLParserError = 37
const XMLParserEntityNotStartedError XMLParserError = 36
const XMLParserEntityRefAtEOFError XMLParserError = 14
const XMLParserEntityRefInDTDError XMLParserError = 17
const XMLParserEntityRefInEpilogError XMLParserError = 16
const XMLParserEntityRefInPrologError XMLParserError = 15
const XMLParserEntityRefLoopError XMLParserError = 89
const XMLParserEntityReferenceMissingSemiError XMLParserError = 23
const XMLParserEntityReferenceWithoutNameError XMLParserError = 22
const XMLParserEntityValueRequiredError XMLParserError = 84
const XMLParserEqualExpectedError XMLParserError = 75
const XMLParserExternalStandaloneEntityError XMLParserError = 82
const XMLParserExternalSubsetNotFinishedError XMLParserError = 60
const XMLParserExtraContentError XMLParserError = 86
const XMLParserGTRequiredError XMLParserError = 73
const XMLParserInternalError XMLParserError = 1
const XMLParserInvalidCharacterError XMLParserError = 9
const XMLParserInvalidCharacterInEntityError XMLParserError = 87
const XMLParserInvalidCharacterRefError XMLParserError = 8
const XMLParserInvalidConditionalSectionError XMLParserError = 83
const XMLParserInvalidDecimalCharacterRefError XMLParserError = 7
const XMLParserInvalidEncodingError XMLParserError = 81
const XMLParserInvalidEncodingNameError XMLParserError = 79
const XMLParserInvalidHexCharacterRefError XMLParserError = 6
const XMLParserInvalidURIError XMLParserError = 91
const XMLParserLTRequiredError XMLParserError = 72
const XMLParserLTSlashRequiredError XMLParserError = 74
const XMLParserLessThanSymbolInAttributeError XMLParserError = 38
const XMLParserLiteralNotFinishedError XMLParserError = 44
const XMLParserLiteralNotStartedError XMLParserError = 43
const XMLParserMisplacedCDATAEndStringError XMLParserError = 62
const XMLParserMisplacedXMLDeclarationError XMLParserError = 64
const XMLParserMixedContentDeclNotFinishedError XMLParserError = 53
const XMLParserMixedContentDeclNotStartedError XMLParserError = 52
const XMLParserNAMERequiredError XMLParserError = 68
const XMLParserNMTOKENRequiredError XMLParserError = 67
const XMLParserNamespaceDeclarationError XMLParserError = 35
const XMLParserNoDTDError XMLParserError = 94
const XMLParserNotWellBalancedError XMLParserError = 85
const XMLParserNotationNotFinishedError XMLParserError = 49
const XMLParserNotationNotStartedError XMLParserError = 48
const XMLParserOutOfMemoryError XMLParserError = 2
const XMLParserPCDATARequiredError XMLParserError = 69
const XMLParserParsedEntityRefAtEOFError XMLParserError = 18
const XMLParserParsedEntityRefInEpilogError XMLParserError = 20
const XMLParserParsedEntityRefInInternalError XMLParserError = 88
const XMLParserParsedEntityRefInInternalSubsetError XMLParserError = 21
const XMLParserParsedEntityRefInPrologError XMLParserError = 19
const XMLParserParsedEntityRefMissingSemiError XMLParserError = 25
const XMLParserParsedEntityRefNoNameError XMLParserError = 24
const XMLParserPrematureDocumentEndError XMLParserError = 5
const XMLParserProcessingInstructionNotFinishedError XMLParserError = 47
const XMLParserProcessingInstructionNotStartedError XMLParserError = 46
const XMLParserPublicIdentifierRequiredError XMLParserError = 71
const XMLParserSeparatorRequiredError XMLParserError = 66
const XMLParserSpaceRequiredError XMLParserError = 65
const XMLParserStandaloneValueError XMLParserError = 78
const XMLParserStringNotClosedError XMLParserError = 34
const XMLParserStringNotStartedError XMLParserError = 33
const XMLParserTagNameMismatchError XMLParserError = 76
const XMLParserURIFragmentError XMLParserError = 92
const XMLParserURIRequiredError XMLParserError = 70
const XMLParserUndeclaredEntityError XMLParserError = 26
const XMLParserUnfinishedTagError XMLParserError = 77
const XMLParserUnknownEncodingError XMLParserError = 31
const XMLParserUnparsedEntityError XMLParserError = 28
const XMLParserXMLDeclNotFinishedError XMLParserError = 57
const XMLParserXMLDeclNotStartedError XMLParserError = 56

type XMLParserExternalEntityResolvingPolicy uint

const XMLParserResolveExternalEntitiesAlways XMLParserExternalEntityResolvingPolicy = 3
const XMLParserResolveExternalEntitiesNever XMLParserExternalEntityResolvingPolicy = 0
const XMLParserResolveExternalEntitiesNoNetwork XMLParserExternalEntityResolvingPolicy = 1
const XMLParserResolveExternalEntitiesSameOriginOnly XMLParserExternalEntityResolvingPolicy = 2

type XPCConnectionOptions uint

const XPCConnectionPrivileged XPCConnectionOptions = 4096

type unichar int
