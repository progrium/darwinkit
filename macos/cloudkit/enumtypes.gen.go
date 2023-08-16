// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

// Constants that indicate the availability of the user’s iCloud account. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckaccountstatus?language=objc
type AccountStatus int

const (
	AccountStatusAvailable              AccountStatus = 1
	AccountStatusCouldNotDetermine      AccountStatus = 0
	AccountStatusNoAccount              AccountStatus = 3
	AccountStatusRestricted             AccountStatus = 2
	AccountStatusTemporarilyUnavailable AccountStatus = 4
)

// Constants that represent the status of a permission. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckapplicationpermissionstatus?language=objc
type ApplicationPermissionStatus int

const (
	ApplicationPermissionStatusCouldNotComplete ApplicationPermissionStatus = 1
	ApplicationPermissionStatusDenied           ApplicationPermissionStatus = 2
	ApplicationPermissionStatusGranted          ApplicationPermissionStatus = 3
	ApplicationPermissionStatusInitialState     ApplicationPermissionStatus = 0
)

// Constants that represent the permissions that a user grants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckapplicationpermissions?language=objc
type ApplicationPermissions uint

const (
	ApplicationPermissionUserDiscoverability ApplicationPermissions = 1
)

// Constants that represent the scope of a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasescope?language=objc
type DatabaseScope int

const (
	DatabaseScopePrivate DatabaseScope = 2
	DatabaseScopePublic  DatabaseScope = 1
	DatabaseScopeShared  DatabaseScope = 3
)

// The error codes that CloudKit returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckerrorcode?language=objc
type ErrorCode int

const (
	ErrorAccountTemporarilyUnavailable  ErrorCode = 36
	ErrorAlreadyShared                  ErrorCode = 30
	ErrorAssetFileModified              ErrorCode = 17
	ErrorAssetFileNotFound              ErrorCode = 16
	ErrorAssetNotAvailable              ErrorCode = 35
	ErrorBadContainer                   ErrorCode = 5
	ErrorBadDatabase                    ErrorCode = 24
	ErrorBatchRequestFailed             ErrorCode = 22
	ErrorChangeTokenExpired             ErrorCode = 21
	ErrorConstraintViolation            ErrorCode = 19
	ErrorIncompatibleVersion            ErrorCode = 18
	ErrorInternalError                  ErrorCode = 1
	ErrorInvalidArguments               ErrorCode = 12
	ErrorLimitExceeded                  ErrorCode = 27
	ErrorManagedAccountRestricted       ErrorCode = 32
	ErrorMissingEntitlement             ErrorCode = 8
	ErrorNetworkFailure                 ErrorCode = 4
	ErrorNetworkUnavailable             ErrorCode = 3
	ErrorNotAuthenticated               ErrorCode = 9
	ErrorOperationCancelled             ErrorCode = 20
	ErrorPartialFailure                 ErrorCode = 2
	ErrorParticipantMayNeedVerification ErrorCode = 33
	ErrorPermissionFailure              ErrorCode = 10
	ErrorQuotaExceeded                  ErrorCode = 25
	ErrorReferenceViolation             ErrorCode = 31
	ErrorRequestRateLimited             ErrorCode = 7
	ErrorResultsTruncated               ErrorCode = 13
	ErrorServerRecordChanged            ErrorCode = 14
	ErrorServerRejectedRequest          ErrorCode = 15
	ErrorServerResponseLost             ErrorCode = 34
	ErrorServiceUnavailable             ErrorCode = 6
	ErrorTooManyParticipants            ErrorCode = 29
	ErrorUnknownItem                    ErrorCode = 11
	ErrorUserDeletedZone                ErrorCode = 28
	ErrorZoneBusy                       ErrorCode = 23
	ErrorZoneNotFound                   ErrorCode = 26
)

// Constants that indicate the type of event that generates the push notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationtype?language=objc
type NotificationType int

const (
	NotificationTypeDatabase         NotificationType = 4
	NotificationTypeQuery            NotificationType = 1
	NotificationTypeReadNotification NotificationType = 3
	NotificationTypeRecordZone       NotificationType = 2
)

// Constants that represent possible data transfer sizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgrouptransfersize?language=objc
type OperationGroupTransferSize int

const (
	OperationGroupTransferSizeGigabytes           OperationGroupTransferSize = 5
	OperationGroupTransferSizeHundredsOfGigabytes OperationGroupTransferSize = 7
	OperationGroupTransferSizeHundredsOfMegabytes OperationGroupTransferSize = 4
	OperationGroupTransferSizeKilobytes           OperationGroupTransferSize = 1
	OperationGroupTransferSizeMegabytes           OperationGroupTransferSize = 2
	OperationGroupTransferSizeTensOfGigabytes     OperationGroupTransferSize = 6
	OperationGroupTransferSizeTensOfMegabytes     OperationGroupTransferSize = 3
	OperationGroupTransferSizeUnknown             OperationGroupTransferSize = 0
)

// A type that represents the ID of an operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationid?language=objc
type OperationID string

// Constants that indicate the event that triggers the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotificationreason?language=objc
type QueryNotificationReason int

const (
	QueryNotificationReasonRecordCreated QueryNotificationReason = 1
	QueryNotificationReasonRecordDeleted QueryNotificationReason = 3
	QueryNotificationReasonRecordUpdated QueryNotificationReason = 2
)

// Configuration options for a query subscription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscriptionoptions?language=objc
type QuerySubscriptionOptions uint

const (
	QuerySubscriptionOptionsFiresOnRecordCreation QuerySubscriptionOptions = 1
	QuerySubscriptionOptionsFiresOnRecordDeletion QuerySubscriptionOptions = 4
	QuerySubscriptionOptionsFiresOnRecordUpdate   QuerySubscriptionOptions = 2
	QuerySubscriptionOptionsFiresOnce             QuerySubscriptionOptions = 8
)

// A data type that CloudKit requires for record field names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordfieldkey?language=objc
type RecordFieldKey string

const (
	RecordParentKey            RecordFieldKey = "___parent"
	RecordShareKey             RecordFieldKey = "___share"
	ShareThumbnailImageDataKey RecordFieldKey = "cloudkit.thumbnailImageData"
	ShareTitleKey              RecordFieldKey = "cloudkit.title"
	ShareTypeKey               RecordFieldKey = "cloudkit.type"
)

// Constants that indicate which policy to apply when saving records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordsavepolicy?language=objc
type RecordSavePolicy int

const (
	RecordSaveAllKeys                 RecordSavePolicy = 2
	RecordSaveChangedKeys             RecordSavePolicy = 1
	RecordSaveIfServerRecordUnchanged RecordSavePolicy = 0
)

// A data type that CloudKit requires for record types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordtype?language=objc
type RecordType string

const (
	RecordTypeShare      RecordType = "cloudkit.share"
	RecordTypeUserRecord RecordType = "Users"
)

// The capabilities that a record zone supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonecapabilities?language=objc
type RecordZoneCapabilities uint

const (
	RecordZoneCapabilityAtomic          RecordZoneCapabilities = 2
	RecordZoneCapabilityFetchChanges    RecordZoneCapabilities = 1
	RecordZoneCapabilitySharing         RecordZoneCapabilities = 4
	RecordZoneCapabilityZoneWideSharing RecordZoneCapabilities = 8
)

// Constants that indicate the behavior when deleting a referenced record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreferenceaction?language=objc
type ReferenceAction uint

const (
	ReferenceActionDeleteSelf ReferenceAction = 1
	ReferenceActionNone       ReferenceAction = 0
)

// Constants that represent the status of a participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipantacceptancestatus?language=objc
type ShareParticipantAcceptanceStatus int

const (
	ShareParticipantAcceptanceStatusAccepted ShareParticipantAcceptanceStatus = 2
	ShareParticipantAcceptanceStatusPending  ShareParticipantAcceptanceStatus = 1
	ShareParticipantAcceptanceStatusRemoved  ShareParticipantAcceptanceStatus = 3
	ShareParticipantAcceptanceStatusUnknown  ShareParticipantAcceptanceStatus = 0
)

// Constants that represent the permissions to grant to a share participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipantpermission?language=objc
type ShareParticipantPermission int

const (
	ShareParticipantPermissionNone      ShareParticipantPermission = 1
	ShareParticipantPermissionReadOnly  ShareParticipantPermission = 2
	ShareParticipantPermissionReadWrite ShareParticipantPermission = 3
	ShareParticipantPermissionUnknown   ShareParticipantPermission = 0
)

// Constants that represent the role of a share’s participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipantrole?language=objc
type ShareParticipantRole int

const (
	ShareParticipantRoleOwner       ShareParticipantRole = 1
	ShareParticipantRolePrivateUser ShareParticipantRole = 3
	ShareParticipantRolePublicUser  ShareParticipantRole = 4
	ShareParticipantRoleUnknown     ShareParticipantRole = 0
)

// The role of a participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipanttype?language=objc
type ShareParticipantType int

const (
	ShareParticipantTypeOwner       ShareParticipantType = 1
	ShareParticipantTypePrivateUser ShareParticipantType = 3
	ShareParticipantTypePublicUser  ShareParticipantType = 4
	ShareParticipantTypeUnknown     ShareParticipantType = 0
)

// A type that represents a subscription’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscriptionid?language=objc
type SubscriptionID string

// Constants that identify a subscription’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscriptiontype?language=objc
type SubscriptionType int

const (
	SubscriptionTypeDatabase   SubscriptionType = 3
	SubscriptionTypeQuery      SubscriptionType = 1
	SubscriptionTypeRecordZone SubscriptionType = 2
)
