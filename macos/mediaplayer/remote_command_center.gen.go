// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RemoteCommandCenter] class.
var RemoteCommandCenterClass = _RemoteCommandCenterClass{objc.GetClass("MPRemoteCommandCenter")}

type _RemoteCommandCenterClass struct {
	objc.Class
}

// An interface definition for the [RemoteCommandCenter] class.
type IRemoteCommandCenter interface {
	objc.IObject
	PlayCommand() RemoteCommand
	SkipBackwardCommand() SkipIntervalCommand
	DisableLanguageOptionCommand() RemoteCommand
	SkipForwardCommand() SkipIntervalCommand
	ChangeRepeatModeCommand() ChangeRepeatModeCommand
	LikeCommand() FeedbackCommand
	TogglePlayPauseCommand() RemoteCommand
	NextTrackCommand() RemoteCommand
	ChangePlaybackPositionCommand() ChangePlaybackPositionCommand
	BookmarkCommand() FeedbackCommand
	ChangePlaybackRateCommand() ChangePlaybackRateCommand
	PreviousTrackCommand() RemoteCommand
	EnableLanguageOptionCommand() RemoteCommand
	ChangeShuffleModeCommand() ChangeShuffleModeCommand
	PauseCommand() RemoteCommand
	DislikeCommand() FeedbackCommand
	StopCommand() RemoteCommand
	SeekBackwardCommand() RemoteCommand
	SeekForwardCommand() RemoteCommand
	RatingCommand() RatingCommand
}

// An object that responds to remote control events sent by external accessories and system controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter?language=objc
type RemoteCommandCenter struct {
	objc.Object
}

func RemoteCommandCenterFrom(ptr unsafe.Pointer) RemoteCommandCenter {
	return RemoteCommandCenter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RemoteCommandCenterClass) Alloc() RemoteCommandCenter {
	rv := objc.Call[RemoteCommandCenter](rc, objc.Sel("alloc"))
	return rv
}

func RemoteCommandCenter_Alloc() RemoteCommandCenter {
	return RemoteCommandCenterClass.Alloc()
}

func (rc _RemoteCommandCenterClass) New() RemoteCommandCenter {
	rv := objc.Call[RemoteCommandCenter](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRemoteCommandCenter() RemoteCommandCenter {
	return RemoteCommandCenterClass.New()
}

func (r_ RemoteCommandCenter) Init() RemoteCommandCenter {
	rv := objc.Call[RemoteCommandCenter](r_, objc.Sel("init"))
	return rv
}

// Returns the shared object you use to access the system’s remote command objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618994-sharedcommandcenter?language=objc
func (rc _RemoteCommandCenterClass) SharedCommandCenter() RemoteCommandCenter {
	rv := objc.Call[RemoteCommandCenter](rc, objc.Sel("sharedCommandCenter"))
	return rv
}

// Returns the shared object you use to access the system’s remote command objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618994-sharedcommandcenter?language=objc
func RemoteCommandCenter_SharedCommandCenter() RemoteCommandCenter {
	return RemoteCommandCenterClass.SharedCommandCenter()
}

// The command object for starting playback of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1619000-playcommand?language=objc
func (r_ RemoteCommandCenter) PlayCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("playCommand"))
	return rv
}

// The command object for playing a previous point in a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618996-skipbackwardcommand?language=objc
func (r_ RemoteCommandCenter) SkipBackwardCommand() SkipIntervalCommand {
	rv := objc.Call[SkipIntervalCommand](r_, objc.Sel("skipBackwardCommand"))
	return rv
}

// The command object for disabling a language option [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618988-disablelanguageoptioncommand?language=objc
func (r_ RemoteCommandCenter) DisableLanguageOptionCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("disableLanguageOptionCommand"))
	return rv
}

// The command object for playing a future point in a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618990-skipforwardcommand?language=objc
func (r_ RemoteCommandCenter) SkipForwardCommand() SkipIntervalCommand {
	rv := objc.Call[SkipIntervalCommand](r_, objc.Sel("skipForwardCommand"))
	return rv
}

// The command object for changing the repeat mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1649694-changerepeatmodecommand?language=objc
func (r_ RemoteCommandCenter) ChangeRepeatModeCommand() ChangeRepeatModeCommand {
	rv := objc.Call[ChangeRepeatModeCommand](r_, objc.Sel("changeRepeatModeCommand"))
	return rv
}

// The command object for indicating that a user likes what is currently playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618998-likecommand?language=objc
func (r_ RemoteCommandCenter) LikeCommand() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](r_, objc.Sel("likeCommand"))
	return rv
}

// The command object for toggling between playing and pausing the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618992-toggleplaypausecommand?language=objc
func (r_ RemoteCommandCenter) TogglePlayPauseCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("togglePlayPauseCommand"))
	return rv
}

// The command object for selecting the next track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618999-nexttrackcommand?language=objc
func (r_ RemoteCommandCenter) NextTrackCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("nextTrackCommand"))
	return rv
}

// The command object for changing the playback position in a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618997-changeplaybackpositioncommand?language=objc
func (r_ RemoteCommandCenter) ChangePlaybackPositionCommand() ChangePlaybackPositionCommand {
	rv := objc.Call[ChangePlaybackPositionCommand](r_, objc.Sel("changePlaybackPositionCommand"))
	return rv
}

// The command object for indicating that a user wants to remember a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1619002-bookmarkcommand?language=objc
func (r_ RemoteCommandCenter) BookmarkCommand() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](r_, objc.Sel("bookmarkCommand"))
	return rv
}

// The command object for changing the playback rate of the current media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618991-changeplaybackratecommand?language=objc
func (r_ RemoteCommandCenter) ChangePlaybackRateCommand() ChangePlaybackRateCommand {
	rv := objc.Call[ChangePlaybackRateCommand](r_, objc.Sel("changePlaybackRateCommand"))
	return rv
}

// The command object for selecting the previous track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618978-previoustrackcommand?language=objc
func (r_ RemoteCommandCenter) PreviousTrackCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("previousTrackCommand"))
	return rv
}

// The command object for enabling a language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618980-enablelanguageoptioncommand?language=objc
func (r_ RemoteCommandCenter) EnableLanguageOptionCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("enableLanguageOptionCommand"))
	return rv
}

// The command object for changing the shuffle mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1649692-changeshufflemodecommand?language=objc
func (r_ RemoteCommandCenter) ChangeShuffleModeCommand() ChangeShuffleModeCommand {
	rv := objc.Call[ChangeShuffleModeCommand](r_, objc.Sel("changeShuffleModeCommand"))
	return rv
}

// The command object for pausing playback of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618979-pausecommand?language=objc
func (r_ RemoteCommandCenter) PauseCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("pauseCommand"))
	return rv
}

// The command object for indicating that a user dislikes what is currently playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618995-dislikecommand?language=objc
func (r_ RemoteCommandCenter) DislikeCommand() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](r_, objc.Sel("dislikeCommand"))
	return rv
}

// The command object for stopping playback of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618984-stopcommand?language=objc
func (r_ RemoteCommandCenter) StopCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("stopCommand"))
	return rv
}

// The command object for seeking backward through a single media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618982-seekbackwardcommand?language=objc
func (r_ RemoteCommandCenter) SeekBackwardCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("seekBackwardCommand"))
	return rv
}

// The command object for seeking forward through a single media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618981-seekforwardcommand?language=objc
func (r_ RemoteCommandCenter) SeekForwardCommand() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("seekForwardCommand"))
	return rv
}

// The command object for rating a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/1618986-ratingcommand?language=objc
func (r_ RemoteCommandCenter) RatingCommand() RatingCommand {
	rv := objc.Call[RatingCommand](r_, objc.Sel("ratingCommand"))
	return rv
}
