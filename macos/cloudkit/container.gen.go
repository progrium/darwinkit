// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Container] class.
var ContainerClass = _ContainerClass{objc.GetClass("CKContainer")}

type _ContainerClass struct {
	objc.Class
}

// An interface definition for the [Container] class.
type IContainer interface {
	objc.IObject
	DatabaseWithDatabaseScope(databaseScope DatabaseScope) Database
	FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler func(outstandingOperationIDs []OperationID, error foundation.Error))
	FetchShareMetadataWithURLCompletionHandler(url foundation.IURL, completionHandler func(metadata ShareMetadata, error foundation.Error))
	AddOperation(operation IOperation)
	FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler func(shareParticipant ShareParticipant, error foundation.Error))
	FetchUserRecordIDWithCompletionHandler(completionHandler func(recordID RecordID, error foundation.Error))
	FetchLongLivedOperationWithIDCompletionHandler(operationID OperationID, completionHandler func(outstandingOperation Operation, error foundation.Error))
	AcceptShareMetadataCompletionHandler(metadata IShareMetadata, completionHandler func(acceptedShare Share, error foundation.Error))
	FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler func(shareParticipant ShareParticipant, error foundation.Error))
	AccountStatusWithCompletionHandler(completionHandler func(accountStatus AccountStatus, error foundation.Error))
	FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID IRecordID, completionHandler func(shareParticipant ShareParticipant, error foundation.Error))
	SharedCloudDatabase() Database
	PrivateCloudDatabase() Database
	PublicCloudDatabase() Database
	ContainerIdentifier() string
}

// A conduit to your app’s databases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer?language=objc
type Container struct {
	objc.Object
}

func ContainerFrom(ptr unsafe.Pointer) Container {
	return Container{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContainerClass) Alloc() Container {
	rv := objc.Call[Container](cc, objc.Sel("alloc"))
	return rv
}

func Container_Alloc() Container {
	return ContainerClass.Alloc()
}

func (cc _ContainerClass) New() Container {
	rv := objc.Call[Container](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContainer() Container {
	return ContainerClass.New()
}

func (c_ Container) Init() Container {
	rv := objc.Call[Container](c_, objc.Sel("init"))
	return rv
}

// Returns the database with the specified scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1640475-databasewithdatabasescope?language=objc
func (c_ Container) DatabaseWithDatabaseScope(databaseScope DatabaseScope) Database {
	rv := objc.Call[Database](c_, objc.Sel("databaseWithDatabaseScope:"), databaseScope)
	return rv
}

// Fetches the IDs of any long-lived operations that are running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399160-fetchalllonglivedoperationidswit?language=objc
func (c_ Container) FetchAllLongLivedOperationIDsWithCompletionHandler(completionHandler func(outstandingOperationIDs []OperationID, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchAllLongLivedOperationIDsWithCompletionHandler:"), completionHandler)
}

// Fetches the share metadata for the specified share URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/2113666-fetchsharemetadatawithurl?language=objc
func (c_ Container) FetchShareMetadataWithURLCompletionHandler(url foundation.IURL, completionHandler func(metadata ShareMetadata, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchShareMetadataWithURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

// Creates a container for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399193-containerwithidentifier?language=objc
func (cc _ContainerClass) ContainerWithIdentifier(containerIdentifier string) Container {
	rv := objc.Call[Container](cc, objc.Sel("containerWithIdentifier:"), containerIdentifier)
	return rv
}

// Creates a container for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399193-containerwithidentifier?language=objc
func Container_ContainerWithIdentifier(containerIdentifier string) Container {
	return ContainerClass.ContainerWithIdentifier(containerIdentifier)
}

// Adds an operation to the container’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399215-addoperation?language=objc
func (c_ Container) AddOperation(operation IOperation) {
	objc.Call[objc.Void](c_, objc.Sel("addOperation:"), objc.Ptr(operation))
}

// Fetches the share participant with the specified email address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1640489-fetchshareparticipantwithemailad?language=objc
func (c_ Container) FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler func(shareParticipant ShareParticipant, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchShareParticipantWithEmailAddress:completionHandler:"), emailAddress, completionHandler)
}

// Returns the app’s default container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399189-defaultcontainer?language=objc
func (cc _ContainerClass) DefaultContainer() Container {
	rv := objc.Call[Container](cc, objc.Sel("defaultContainer"))
	return rv
}

// Returns the app’s default container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399189-defaultcontainer?language=objc
func Container_DefaultContainer() Container {
	return ContainerClass.DefaultContainer()
}

// Fetches the user record ID of the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399191-fetchuserrecordidwithcompletionh?language=objc
func (c_ Container) FetchUserRecordIDWithCompletionHandler(completionHandler func(recordID RecordID, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchUserRecordIDWithCompletionHandler:"), completionHandler)
}

// Fetches the long-lived operation for the specified operation ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399164-fetchlonglivedoperationwithid?language=objc
func (c_ Container) FetchLongLivedOperationWithIDCompletionHandler(operationID OperationID, completionHandler func(outstandingOperation Operation, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchLongLivedOperationWithID:completionHandler:"), operationID, completionHandler)
}

// Accepts the specified share metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/2113667-acceptsharemetadata?language=objc
func (c_ Container) AcceptShareMetadataCompletionHandler(metadata IShareMetadata, completionHandler func(acceptedShare Share, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("acceptShareMetadata:completionHandler:"), objc.Ptr(metadata), completionHandler)
}

// Fetches the share participant with the specified phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1640493-fetchshareparticipantwithphonenu?language=objc
func (c_ Container) FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler func(shareParticipant ShareParticipant, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchShareParticipantWithPhoneNumber:completionHandler:"), phoneNumber, completionHandler)
}

// Determines whether the system can access the user’s iCloud account. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399180-accountstatuswithcompletionhandl?language=objc
func (c_ Container) AccountStatusWithCompletionHandler(completionHandler func(accountStatus AccountStatus, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("accountStatusWithCompletionHandler:"), completionHandler)
}

// Fetches the share participant with the specified user record ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1640387-fetchshareparticipantwithuserrec?language=objc
func (c_ Container) FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID IRecordID, completionHandler func(shareParticipant ShareParticipant, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("fetchShareParticipantWithUserRecordID:completionHandler:"), objc.Ptr(userRecordID), completionHandler)
}

// The database that contains shared data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1640408-sharedclouddatabase?language=objc
func (c_ Container) SharedCloudDatabase() Database {
	rv := objc.Call[Database](c_, objc.Sel("sharedCloudDatabase"))
	return rv
}

// The user’s private database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399205-privateclouddatabase?language=objc
func (c_ Container) PrivateCloudDatabase() Database {
	rv := objc.Call[Database](c_, objc.Sel("privateCloudDatabase"))
	return rv
}

// The app’s public database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399166-publicclouddatabase?language=objc
func (c_ Container) PublicCloudDatabase() Database {
	rv := objc.Call[Database](c_, objc.Sel("publicCloudDatabase"))
	return rv
}

// The container’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcontainer/1399182-containeridentifier?language=objc
func (c_ Container) ContainerIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("containerIdentifier"))
	return rv
}
