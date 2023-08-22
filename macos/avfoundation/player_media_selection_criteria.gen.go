// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerMediaSelectionCriteria] class.
var PlayerMediaSelectionCriteriaClass = _PlayerMediaSelectionCriteriaClass{objc.GetClass("AVPlayerMediaSelectionCriteria")}

type _PlayerMediaSelectionCriteriaClass struct {
	objc.Class
}

// An interface definition for the [PlayerMediaSelectionCriteria] class.
type IPlayerMediaSelectionCriteria interface {
	objc.IObject
	PreferredLanguages() []string
	PrincipalMediaCharacteristics() []MediaCharacteristic
	PreferredMediaCharacteristics() []MediaCharacteristic
}

// An object that specifies the preferred languages and media characteristics for a player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria?language=objc
type PlayerMediaSelectionCriteria struct {
	objc.Object
}

func PlayerMediaSelectionCriteriaFrom(ptr unsafe.Pointer) PlayerMediaSelectionCriteria {
	return PlayerMediaSelectionCriteria{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PlayerMediaSelectionCriteria) InitWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []MediaCharacteristic) PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](p_, objc.Sel("initWithPreferredLanguages:preferredMediaCharacteristics:"), preferredLanguages, preferredMediaCharacteristics)
	return rv
}

// Creates media selection criteria with the preferred languages and media characteristics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/1387627-initwithpreferredlanguages?language=objc
func NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []MediaCharacteristic) PlayerMediaSelectionCriteria {
	instance := PlayerMediaSelectionCriteriaClass.Alloc().InitWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages, preferredMediaCharacteristics)
	instance.Autorelease()
	return instance
}

func (p_ PlayerMediaSelectionCriteria) InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []MediaCharacteristic, preferredLanguages []string, preferredMediaCharacteristics []MediaCharacteristic) PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](p_, objc.Sel("initWithPrincipalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:"), principalMediaCharacteristics, preferredLanguages, preferredMediaCharacteristics)
	return rv
}

// Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/3042657-initwithprincipalmediacharacteri?language=objc
func NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []MediaCharacteristic, preferredLanguages []string, preferredMediaCharacteristics []MediaCharacteristic) PlayerMediaSelectionCriteria {
	instance := PlayerMediaSelectionCriteriaClass.Alloc().InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics, preferredLanguages, preferredMediaCharacteristics)
	instance.Autorelease()
	return instance
}

func (pc _PlayerMediaSelectionCriteriaClass) Alloc() PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](pc, objc.Sel("alloc"))
	return rv
}

func PlayerMediaSelectionCriteria_Alloc() PlayerMediaSelectionCriteria {
	return PlayerMediaSelectionCriteriaClass.Alloc()
}

func (pc _PlayerMediaSelectionCriteriaClass) New() PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerMediaSelectionCriteria() PlayerMediaSelectionCriteria {
	return PlayerMediaSelectionCriteriaClass.New()
}

func (p_ PlayerMediaSelectionCriteria) Init() PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](p_, objc.Sel("init"))
	return rv
}

// An array of language identifiers in preferred order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/1388559-preferredlanguages?language=objc
func (p_ PlayerMediaSelectionCriteria) PreferredLanguages() []string {
	rv := objc.Call[[]string](p_, objc.Sel("preferredLanguages"))
	return rv
}

// An array of media characteristics that are essential to select when choosing media with a particular characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/3042658-principalmediacharacteristics?language=objc
func (p_ PlayerMediaSelectionCriteria) PrincipalMediaCharacteristics() []MediaCharacteristic {
	rv := objc.Call[[]MediaCharacteristic](p_, objc.Sel("principalMediaCharacteristics"))
	return rv
}

// An array of media characteristics in preferred order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/1385734-preferredmediacharacteristics?language=objc
func (p_ PlayerMediaSelectionCriteria) PreferredMediaCharacteristics() []MediaCharacteristic {
	rv := objc.Call[[]MediaCharacteristic](p_, objc.Sel("preferredMediaCharacteristics"))
	return rv
}
