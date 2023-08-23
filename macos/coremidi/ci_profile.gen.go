// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CIProfile] class.
var CIProfileClass = _CIProfileClass{objc.GetClass("MIDICIProfile")}

type _CIProfileClass struct {
	objc.Class
}

// An interface definition for the [CIProfile] class.
type ICIProfile interface {
	objc.IObject
	Name() string
	ProfileID() []byte
}

// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofile?language=objc
type CIProfile struct {
	objc.Object
}

func CIProfileFrom(ptr unsafe.Pointer) CIProfile {
	return CIProfile{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CIProfile) InitWithData(data []byte) CIProfile {
	rv := objc.Call[CIProfile](c_, objc.Sel("initWithData:"), data)
	return rv
}

// Creates a MIDI profile for the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofile/3580326-initwithdata?language=objc
func NewCIProfileWithData(data []byte) CIProfile {
	instance := CIProfileClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (cc _CIProfileClass) Alloc() CIProfile {
	rv := objc.Call[CIProfile](cc, objc.Sel("alloc"))
	return rv
}

func CIProfile_Alloc() CIProfile {
	return CIProfileClass.Alloc()
}

func (cc _CIProfileClass) New() CIProfile {
	rv := objc.Call[CIProfile](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIProfile() CIProfile {
	return CIProfileClass.New()
}

func (c_ CIProfile) Init() CIProfile {
	rv := objc.Call[CIProfile](c_, objc.Sel("init"))
	return rv
}

// A string that describes the profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofile/2977105-name?language=objc
func (c_ CIProfile) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

// The unique five-byte profile identifier that represents the profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofile/2977106-profileid?language=objc
func (c_ CIProfile) ProfileID() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("profileID"))
	return rv
}
