// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeLanguageOptionCommandEvent] class.
var ChangeLanguageOptionCommandEventClass = _ChangeLanguageOptionCommandEventClass{objc.GetClass("MPChangeLanguageOptionCommandEvent")}

type _ChangeLanguageOptionCommandEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeLanguageOptionCommandEvent] class.
type IChangeLanguageOptionCommandEvent interface {
	IRemoteCommandEvent
	Setting() ChangeLanguageOptionSetting
}

// An event requesting a change in the language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent?language=objc
type ChangeLanguageOptionCommandEvent struct {
	RemoteCommandEvent
}

func ChangeLanguageOptionCommandEventFrom(ptr unsafe.Pointer) ChangeLanguageOptionCommandEvent {
	return ChangeLanguageOptionCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (cc _ChangeLanguageOptionCommandEventClass) Alloc() ChangeLanguageOptionCommandEvent {
	rv := objc.Call[ChangeLanguageOptionCommandEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeLanguageOptionCommandEvent_Alloc() ChangeLanguageOptionCommandEvent {
	return ChangeLanguageOptionCommandEventClass.Alloc()
}

func (cc _ChangeLanguageOptionCommandEventClass) New() ChangeLanguageOptionCommandEvent {
	rv := objc.Call[ChangeLanguageOptionCommandEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeLanguageOptionCommandEvent() ChangeLanguageOptionCommandEvent {
	return ChangeLanguageOptionCommandEventClass.New()
}

func (c_ ChangeLanguageOptionCommandEvent) Init() ChangeLanguageOptionCommandEvent {
	rv := objc.Call[ChangeLanguageOptionCommandEvent](c_, objc.Sel("init"))
	return rv
}

// The extent of the language setting change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent/1649697-setting?language=objc
func (c_ ChangeLanguageOptionCommandEvent) Setting() ChangeLanguageOptionSetting {
	rv := objc.Call[ChangeLanguageOptionSetting](c_, objc.Sel("setting"))
	return rv
}
