// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import "math"

// The states that determine when language option changes take effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptionsetting?language=objc
type ChangeLanguageOptionSetting int

const (
	ChangeLanguageOptionSettingNone               ChangeLanguageOptionSetting = 0
	ChangeLanguageOptionSettingNowPlayingItemOnly ChangeLanguageOptionSetting = 1
	ChangeLanguageOptionSettingPermanent          ChangeLanguageOptionSetting = 2
)

// An enumeration that represents error codes for framework operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mperrorcode?language=objc
type ErrorCode int

const (
	ErrorCancelled                     ErrorCode = 6
	ErrorCloudServiceCapabilityMissing ErrorCode = 2
	ErrorNetworkConnectionFailed       ErrorCode = 3
	ErrorNotFound                      ErrorCode = 4
	ErrorNotSupported                  ErrorCode = 5
	ErrorPermissionDenied              ErrorCode = 1
	ErrorRequestTimedOut               ErrorCode = 7
	ErrorUnknown                       ErrorCode = 0
)

// Defines the type for storing a persistent identifier to a particular entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaentitypersistentid?language=objc
type MediaEntityPersistentID uint64

// The properties for defining the type for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediatype?language=objc
type MediaType uint

const (
	MediaTypeAny          MediaType = math.MaxUint
	MediaTypeAnyAudio     MediaType = 255
	MediaTypeAnyVideo     MediaType = 65280
	MediaTypeAudioBook    MediaType = 4
	MediaTypeAudioITunesU MediaType = 8
	MediaTypeHomeVideo    MediaType = 8192
	MediaTypeMovie        MediaType = 256
	MediaTypeMusic        MediaType = 1
	MediaTypeMusicVideo   MediaType = 2048
	MediaTypePodcast      MediaType = 2
	MediaTypeTVShow       MediaType = 512
	MediaTypeVideoITunesU MediaType = 4096
	MediaTypeVideoPodcast MediaType = 1024
)

// The language option type to use for the Now Playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoptiontype?language=objc
type NowPlayingInfoLanguageOptionType uint

const (
	NowPlayingInfoLanguageOptionTypeAudible NowPlayingInfoLanguageOptionType = 0
	NowPlayingInfoLanguageOptionTypeLegible NowPlayingInfoLanguageOptionType = 1
)

// The type of media currently playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfomediatype?language=objc
type NowPlayingInfoMediaType uint

const (
	NowPlayingInfoMediaTypeAudio NowPlayingInfoMediaType = 1
	NowPlayingInfoMediaTypeNone  NowPlayingInfoMediaType = 0
	NowPlayingInfoMediaTypeVideo NowPlayingInfoMediaType = 2
)

// The playback state of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayingplaybackstate?language=objc
type NowPlayingPlaybackState uint

const (
	NowPlayingPlaybackStateInterrupted NowPlayingPlaybackState = 4
	NowPlayingPlaybackStatePaused      NowPlayingPlaybackState = 2
	NowPlayingPlaybackStatePlaying     NowPlayingPlaybackState = 1
	NowPlayingPlaybackStateStopped     NowPlayingPlaybackState = 3
	NowPlayingPlaybackStateUnknown     NowPlayingPlaybackState = 0
)

// Constants indicating the status of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandhandlerstatus?language=objc
type RemoteCommandHandlerStatus int

const (
	RemoteCommandHandlerStatusCommandFailed              RemoteCommandHandlerStatus = 200
	RemoteCommandHandlerStatusDeviceNotFound             RemoteCommandHandlerStatus = 120
	RemoteCommandHandlerStatusNoActionableNowPlayingItem RemoteCommandHandlerStatus = 110
	RemoteCommandHandlerStatusNoSuchContent              RemoteCommandHandlerStatus = 100
	RemoteCommandHandlerStatusSuccess                    RemoteCommandHandlerStatus = 0
)

// Indicates which items to play repeatedly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mprepeattype?language=objc
type RepeatType int

const (
	RepeatTypeAll RepeatType = 2
	RepeatTypeOff RepeatType = 0
	RepeatTypeOne RepeatType = 1
)

// Defines the beginning and ending of seek events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpseekcommandeventtype?language=objc
type SeekCommandEventType uint

const (
	SeekCommandEventTypeBeginSeeking SeekCommandEventType = 0
	SeekCommandEventTypeEndSeeking   SeekCommandEventType = 1
)

// Indicates which item types to shuffle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpshuffletype?language=objc
type ShuffleType int

const (
	ShuffleTypeCollections ShuffleType = 2
	ShuffleTypeItems       ShuffleType = 1
	ShuffleTypeOff         ShuffleType = 0
)
