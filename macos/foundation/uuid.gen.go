// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UUID] class.
var UUIDClass = _UUIDClass{objc.GetClass("NSUUID")}

type _UUIDClass struct {
	objc.Class
}

// An interface definition for the [UUID] class.
type IUUID interface {
	objc.IObject
	GetUUIDBytes(uuid *uint8)
	Compare(otherUUID IUUID) ComparisonResult
	UUIDString() string
}

// A universally unique value that can be used to identify types, interfaces, and other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid?language=objc
type UUID struct {
	objc.Object
}

func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UUID) InitWithUUIDBytes(bytes *uint8) UUID {
	rv := objc.Call[UUID](u_, objc.Sel("initWithUUIDBytes:"), bytes)
	return rv
}

// Initializes a new UUID with the given bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/1417039-initwithuuidbytes?language=objc
func UUID_InitWithUUIDBytes(bytes *uint8) UUID {
	return UUIDClass.Alloc().InitWithUUIDBytes(bytes)
}

func (uc _UUIDClass) UUID() UUID {
	rv := objc.Call[UUID](uc, objc.Sel("UUID"))
	return rv
}

// Create and returns a new UUID with RFC 4122 version 4 random bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/1574571-uuid?language=objc
func UUID_UUID() UUID {
	return UUIDClass.UUID()
}

func (u_ UUID) Init() UUID {
	rv := objc.Call[UUID](u_, objc.Sel("init"))
	return rv
}

func (u_ UUID) InitWithUUIDString(string_ string) UUID {
	rv := objc.Call[UUID](u_, objc.Sel("initWithUUIDString:"), string_)
	return rv
}

// Initializes a new UUID with the formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/1411662-initwithuuidstring?language=objc
func UUID_InitWithUUIDString(string_ string) UUID {
	return UUIDClass.Alloc().InitWithUUIDString(string_)
}

func (uc _UUIDClass) Alloc() UUID {
	rv := objc.Call[UUID](uc, objc.Sel("alloc"))
	return rv
}

func UUID_Alloc() UUID {
	return UUIDClass.Alloc()
}

func (uc _UUIDClass) New() UUID {
	rv := objc.Call[UUID](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUUID() UUID {
	return UUIDClass.New()
}

// Returns the UUID as bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/1411420-getuuidbytes?language=objc
func (u_ UUID) GetUUIDBytes(uuid *uint8) {
	objc.Call[objc.Void](u_, objc.Sel("getUUIDBytes:"), uuid)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/3746978-compare?language=objc
func (u_ UUID) Compare(otherUUID IUUID) ComparisonResult {
	rv := objc.Call[ComparisonResult](u_, objc.Sel("compare:"), objc.Ptr(otherUUID))
	return rv
}

// The UUID as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuuid/1416585-uuidstring?language=objc
func (u_ UUID) UUIDString() string {
	rv := objc.Call[string](u_, objc.Sel("UUIDString"))
	return rv
}
