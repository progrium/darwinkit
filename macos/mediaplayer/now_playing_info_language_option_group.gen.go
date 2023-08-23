// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NowPlayingInfoLanguageOptionGroup] class.
var NowPlayingInfoLanguageOptionGroupClass = _NowPlayingInfoLanguageOptionGroupClass{objc.GetClass("MPNowPlayingInfoLanguageOptionGroup")}

type _NowPlayingInfoLanguageOptionGroupClass struct {
	objc.Class
}

// An interface definition for the [NowPlayingInfoLanguageOptionGroup] class.
type INowPlayingInfoLanguageOptionGroup interface {
	objc.IObject
	AllowEmptySelection() bool
}

// A grouped set of language options where only a single language option can be active at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoptiongroup?language=objc
type NowPlayingInfoLanguageOptionGroup struct {
	objc.Object
}

func NowPlayingInfoLanguageOptionGroupFrom(ptr unsafe.Pointer) NowPlayingInfoLanguageOptionGroup {
	return NowPlayingInfoLanguageOptionGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NowPlayingInfoLanguageOptionGroupClass) Alloc() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Call[NowPlayingInfoLanguageOptionGroup](nc, objc.Sel("alloc"))
	return rv
}

func NowPlayingInfoLanguageOptionGroup_Alloc() NowPlayingInfoLanguageOptionGroup {
	return NowPlayingInfoLanguageOptionGroupClass.Alloc()
}

func (nc _NowPlayingInfoLanguageOptionGroupClass) New() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Call[NowPlayingInfoLanguageOptionGroup](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNowPlayingInfoLanguageOptionGroup() NowPlayingInfoLanguageOptionGroup {
	return NowPlayingInfoLanguageOptionGroupClass.New()
}

func (n_ NowPlayingInfoLanguageOptionGroup) Init() NowPlayingInfoLanguageOptionGroup {
	rv := objc.Call[NowPlayingInfoLanguageOptionGroup](n_, objc.Sel("init"))
	return rv
}

// A Boolean that indicates whether the system requires a selection for the language option group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoptiongroup/1623161-allowemptyselection?language=objc
func (n_ NowPlayingInfoLanguageOptionGroup) AllowEmptySelection() bool {
	rv := objc.Call[bool](n_, objc.Sel("allowEmptySelection"))
	return rv
}
