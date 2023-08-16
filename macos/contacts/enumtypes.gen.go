// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

// An authorization status the user can grant for an app to access the specified entity type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnauthorizationstatus?language=objc
type AuthorizationStatus int

const (
	AuthorizationStatusAuthorized    AuthorizationStatus = 3
	AuthorizationStatusDenied        AuthorizationStatus = 2
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	AuthorizationStatusRestricted    AuthorizationStatus = 1
)

// The formatting orders for contact names component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactdisplaynameorder?language=objc
type ContactDisplayNameOrder int

const (
	ContactDisplayNameOrderFamilyNameFirst ContactDisplayNameOrder = 2
	ContactDisplayNameOrderGivenNameFirst  ContactDisplayNameOrder = 1
	ContactDisplayNameOrderUserDefault     ContactDisplayNameOrder = 0
)

// The formatting styles for contact names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatterstyle?language=objc
type ContactFormatterStyle int

const (
	ContactFormatterStyleFullName         ContactFormatterStyle = 0
	ContactFormatterStylePhoneticFullName ContactFormatterStyle = 1
)

// Indicates the sorting order for contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactsortorder?language=objc
type ContactSortOrder int

const (
	ContactSortOrderFamilyName  ContactSortOrder = 3
	ContactSortOrderGivenName   ContactSortOrder = 2
	ContactSortOrderNone        ContactSortOrder = 0
	ContactSortOrderUserDefault ContactSortOrder = 1
)

// The types a contact can be. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontacttype?language=objc
type ContactType int

const (
	ContactTypeOrganization ContactType = 1
	ContactTypePerson       ContactType = 0
)

// The container may be local on the device or associated with a server account that has contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainertype?language=objc
type ContainerType int

const (
	ContainerTypeCardDAV    ContainerType = 3
	ContainerTypeExchange   ContainerType = 2
	ContainerTypeLocal      ContainerType = 1
	ContainerTypeUnassigned ContainerType = 0
)

// The entities the user can grant access to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnentitytype?language=objc
type EntityType int

const (
	EntityTypeContacts EntityType = 0
)

// Error codes that may be returned when using the methods of the Contacts framework. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnerrorcode?language=objc
type ErrorCode int

const (
	ErrorCodeAuthorizationDenied              ErrorCode = 100
	ErrorCodeChangeHistoryExpired             ErrorCode = 603
	ErrorCodeChangeHistoryInvalidAnchor       ErrorCode = 604
	ErrorCodeChangeHistoryInvalidFetchRequest ErrorCode = 605
	ErrorCodeClientIdentifierCollision        ErrorCode = 602
	ErrorCodeClientIdentifierDoesNotExist     ErrorCode = 601
	ErrorCodeClientIdentifierInvalid          ErrorCode = 600
	ErrorCodeCommunicationError               ErrorCode = 1
	ErrorCodeContainmentCycle                 ErrorCode = 202
	ErrorCodeContainmentScope                 ErrorCode = 203
	ErrorCodeDataAccessError                  ErrorCode = 2
	ErrorCodeFeatureDisabledByUser            ErrorCode = 103
	ErrorCodeInsertedRecordAlreadyExists      ErrorCode = 201
	ErrorCodeNoAccessableWritableContainers   ErrorCode = 101
	ErrorCodeParentContainerNotWritable       ErrorCode = 207
	ErrorCodeParentRecordDoesNotExist         ErrorCode = 204
	ErrorCodePolicyViolation                  ErrorCode = 500
	ErrorCodePredicateInvalid                 ErrorCode = 400
	ErrorCodeRecordDoesNotExist               ErrorCode = 200
	ErrorCodeRecordIdentifierInvalid          ErrorCode = 205
	ErrorCodeRecordNotWritable                ErrorCode = 206
	ErrorCodeUnauthorizedKeys                 ErrorCode = 102
	ErrorCodeVCardMalformed                   ErrorCode = 700
	ErrorCodeVCardSummarizationError          ErrorCode = 701
	ErrorCodeValidationConfigurationError     ErrorCode = 302
	ErrorCodeValidationMultipleErrors         ErrorCode = 300
	ErrorCodeValidationTypeMismatch           ErrorCode = 301
)

// Constants for postal formatting styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatterstyle?language=objc
type PostalAddressFormatterStyle int

const (
	PostalAddressFormatterStyleMailingAddress PostalAddressFormatterStyle = 0
)
