// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Share] class.
var ShareClass = _ShareClass{objc.GetClass("CKShare")}

type _ShareClass struct {
	objc.Class
}

// An interface definition for the [Share] class.
type IShare interface {
	IRecord
	AddParticipant(participant IShareParticipant)
	RemoveParticipant(participant IShareParticipant)
	CurrentUserParticipant() ShareParticipant
	Owner() ShareParticipant
	PublicPermission() ShareParticipantPermission
	SetPublicPermission(value ShareParticipantPermission)
	URL() foundation.URL
	Participants() []ShareParticipant
}

// A specialized record type that manages a collection of shared records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare?language=objc
type Share struct {
	Record
}

func ShareFrom(ptr unsafe.Pointer) Share {
	return Share{
		Record: RecordFrom(ptr),
	}
}

func (s_ Share) InitWithRootRecord(rootRecord IRecord) Share {
	rv := objc.Call[Share](s_, objc.Sel("initWithRootRecord:"), objc.Ptr(rootRecord))
	return rv
}

// Creates a new share for the specified record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640448-initwithrootrecord?language=objc
func NewShareWithRootRecord(rootRecord IRecord) Share {
	instance := ShareClass.Alloc().InitWithRootRecord(rootRecord)
	instance.Autorelease()
	return instance
}

func (s_ Share) InitWithRecordZoneID(recordZoneID IRecordZoneID) Share {
	rv := objc.Call[Share](s_, objc.Sel("initWithRecordZoneID:"), objc.Ptr(recordZoneID))
	return rv
}

// Creates a new share for the specified record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/3746825-initwithrecordzoneid?language=objc
func NewShareWithRecordZoneID(recordZoneID IRecordZoneID) Share {
	instance := ShareClass.Alloc().InitWithRecordZoneID(recordZoneID)
	instance.Autorelease()
	return instance
}

func (sc _ShareClass) Alloc() Share {
	rv := objc.Call[Share](sc, objc.Sel("alloc"))
	return rv
}

func Share_Alloc() Share {
	return ShareClass.Alloc()
}

func (sc _ShareClass) New() Share {
	rv := objc.Call[Share](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShare() Share {
	return ShareClass.New()
}

func (s_ Share) Init() Share {
	rv := objc.Call[Share](s_, objc.Sel("init"))
	return rv
}

func (s_ Share) InitWithRecordTypeZoneID(recordType RecordType, zoneID IRecordZoneID) Share {
	rv := objc.Call[Share](s_, objc.Sel("initWithRecordType:zoneID:"), recordType, objc.Ptr(zoneID))
	return rv
}

// Creates a record in the specified zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462202-initwithrecordtype?language=objc
func NewShareWithRecordTypeZoneID(recordType RecordType, zoneID IRecordZoneID) Share {
	instance := ShareClass.Alloc().InitWithRecordTypeZoneID(recordType, zoneID)
	instance.Autorelease()
	return instance
}

// Adds a participant to the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640443-addparticipant?language=objc
func (s_ Share) AddParticipant(participant IShareParticipant) {
	objc.Call[objc.Void](s_, objc.Sel("addParticipant:"), objc.Ptr(participant))
}

// Removes a participant from the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640523-removeparticipant?language=objc
func (s_ Share) RemoveParticipant(participant IShareParticipant) {
	objc.Call[objc.Void](s_, objc.Sel("removeParticipant:"), objc.Ptr(participant))
}

// The participant that represents the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640441-currentuserparticipant?language=objc
func (s_ Share) CurrentUserParticipant() ShareParticipant {
	rv := objc.Call[ShareParticipant](s_, objc.Sel("currentUserParticipant"))
	return rv
}

// The participant that represents the share’s owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640503-owner?language=objc
func (s_ Share) Owner() ShareParticipant {
	rv := objc.Call[ShareParticipant](s_, objc.Sel("owner"))
	return rv
}

// The permission for anyone with access to the share’s URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640494-publicpermission?language=objc
func (s_ Share) PublicPermission() ShareParticipantPermission {
	rv := objc.Call[ShareParticipantPermission](s_, objc.Sel("publicPermission"))
	return rv
}

// The permission for anyone with access to the share’s URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640494-publicpermission?language=objc
func (s_ Share) SetPublicPermission(value ShareParticipantPermission) {
	objc.Call[objc.Void](s_, objc.Sel("setPublicPermission:"), value)
}

// The URL for inviting participants to the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640465-url?language=objc
func (s_ Share) URL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("URL"))
	return rv
}

// An array that contains the share’s participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/1640453-participants?language=objc
func (s_ Share) Participants() []ShareParticipant {
	rv := objc.Call[[]ShareParticipant](s_, objc.Sel("participants"))
	return rv
}
