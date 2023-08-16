// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ShareMetadata] class.
var ShareMetadataClass = _ShareMetadataClass{objc.GetClass("CKShareMetadata")}

type _ShareMetadataClass struct {
	objc.Class
}

// An interface definition for the [ShareMetadata] class.
type IShareMetadata interface {
	objc.IObject
	ParticipantPermission() ShareParticipantPermission
	Share() Share
	ParticipantRole() ShareParticipantRole
	OwnerIdentity() UserIdentity
	ParticipantStatus() ShareParticipantAcceptanceStatus
	RootRecord() Record
	ContainerIdentifier() string
	HierarchicalRootRecordID() RecordID
}

// An object that describes a shared record’s metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata?language=objc
type ShareMetadata struct {
	objc.Object
}

func ShareMetadataFrom(ptr unsafe.Pointer) ShareMetadata {
	return ShareMetadata{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ShareMetadataClass) Alloc() ShareMetadata {
	rv := objc.Call[ShareMetadata](sc, objc.Sel("alloc"))
	return rv
}

func ShareMetadata_Alloc() ShareMetadata {
	return ShareMetadataClass.Alloc()
}

func (sc _ShareMetadataClass) New() ShareMetadata {
	rv := objc.Call[ShareMetadata](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShareMetadata() ShareMetadata {
	return ShareMetadataClass.New()
}

func (s_ ShareMetadata) Init() ShareMetadata {
	rv := objc.Call[ShareMetadata](s_, objc.Sel("init"))
	return rv
}

// The share’s permissions for the user who retrieves the metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640483-participantpermission?language=objc
func (s_ ShareMetadata) ParticipantPermission() ShareParticipantPermission {
	rv := objc.Call[ShareParticipantPermission](s_, objc.Sel("participantPermission"))
	return rv
}

// The share that owns the metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640412-share?language=objc
func (s_ ShareMetadata) Share() Share {
	rv := objc.Call[Share](s_, objc.Sel("share"))
	return rv
}

// The share’s participant role for the user who retrieves the metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/2980666-participantrole?language=objc
func (s_ ShareMetadata) ParticipantRole() ShareParticipantRole {
	rv := objc.Call[ShareParticipantRole](s_, objc.Sel("participantRole"))
	return rv
}

// The identity of the share’s owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640498-owneridentity?language=objc
func (s_ ShareMetadata) OwnerIdentity() UserIdentity {
	rv := objc.Call[UserIdentity](s_, objc.Sel("ownerIdentity"))
	return rv
}

// The share’s participation status for the user who retrieves the metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640420-participantstatus?language=objc
func (s_ ShareMetadata) ParticipantStatus() ShareParticipantAcceptanceStatus {
	rv := objc.Call[ShareParticipantAcceptanceStatus](s_, objc.Sel("participantStatus"))
	return rv
}

// The share’s root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640366-rootrecord?language=objc
func (s_ ShareMetadata) RootRecord() Record {
	rv := objc.Call[Record](s_, objc.Sel("rootRecord"))
	return rv
}

// The ID of the share’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/1640400-containeridentifier?language=objc
func (s_ ShareMetadata) ContainerIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("containerIdentifier"))
	return rv
}

// The record ID of the shared hierarchy’s root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharemetadata/3793709-hierarchicalrootrecordid?language=objc
func (s_ ShareMetadata) HierarchicalRootRecordID() RecordID {
	rv := objc.Call[RecordID](s_, objc.Sel("hierarchicalRootRecordID"))
	return rv
}
