// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentItem] class.
var ContentItemClass = _ContentItemClass{objc.GetClass("MPContentItem")}

type _ContentItemClass struct {
	objc.Class
}

// An interface definition for the [ContentItem] class.
type IContentItem interface {
	objc.IObject
	Subtitle() string
	SetSubtitle(value string)
	PlaybackProgress() float64
	SetPlaybackProgress(value float64)
	Artwork() MediaItemArtwork
	SetArtwork(value IMediaItemArtwork)
	IsStreamingContent() bool
	SetStreamingContent(value bool)
	IsContainer() bool
	SetContainer(value bool)
	Title() string
	SetTitle(value string)
	IsExplicitContent() bool
	SetExplicitContent(value bool)
	Identifier() string
	IsPlayable() bool
	SetPlayable(value bool)
}

// An object that contains the information for a displayed media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem?language=objc
type ContentItem struct {
	objc.Object
}

func ContentItemFrom(ptr unsafe.Pointer) ContentItem {
	return ContentItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ContentItem) InitWithIdentifier(identifier string) ContentItem {
	rv := objc.Call[ContentItem](c_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Sets the identifier for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620152-initwithidentifier?language=objc
func NewContentItemWithIdentifier(identifier string) ContentItem {
	instance := ContentItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

func (cc _ContentItemClass) Alloc() ContentItem {
	rv := objc.Call[ContentItem](cc, objc.Sel("alloc"))
	return rv
}

func ContentItem_Alloc() ContentItem {
	return ContentItemClass.Alloc()
}

func (cc _ContentItemClass) New() ContentItem {
	rv := objc.Call[ContentItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentItem() ContentItem {
	return ContentItemClass.New()
}

func (c_ ContentItem) Init() ContentItem {
	rv := objc.Call[ContentItem](c_, objc.Sel("init"))
	return rv
}

// A secondary designator for the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620155-subtitle?language=objc
func (c_ ContentItem) Subtitle() string {
	rv := objc.Call[string](c_, objc.Sel("subtitle"))
	return rv
}

// A secondary designator for the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620155-subtitle?language=objc
func (c_ ContentItem) SetSubtitle(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setSubtitle:"), value)
}

// The amount of content played for the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620153-playbackprogress?language=objc
func (c_ ContentItem) PlaybackProgress() float64 {
	rv := objc.Call[float64](c_, objc.Sel("playbackProgress"))
	return rv
}

// The amount of content played for the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620153-playbackprogress?language=objc
func (c_ ContentItem) SetPlaybackProgress(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPlaybackProgress:"), value)
}

// A single image that’s associated with the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620160-artwork?language=objc
func (c_ ContentItem) Artwork() MediaItemArtwork {
	rv := objc.Call[MediaItemArtwork](c_, objc.Sel("artwork"))
	return rv
}

// A single image that’s associated with the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620160-artwork?language=objc
func (c_ ContentItem) SetArtwork(value IMediaItemArtwork) {
	objc.Call[objc.Void](c_, objc.Sel("setArtwork:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the content item is streaming content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1771745-streamingcontent?language=objc
func (c_ ContentItem) IsStreamingContent() bool {
	rv := objc.Call[bool](c_, objc.Sel("isStreamingContent"))
	return rv
}

// A Boolean value that indicates whether the content item is streaming content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1771745-streamingcontent?language=objc
func (c_ ContentItem) SetStreamingContent(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setStreamingContent:"), value)
}

// A Boolean value that indicates whether a media item is container of other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620154-container?language=objc
func (c_ ContentItem) IsContainer() bool {
	rv := objc.Call[bool](c_, objc.Sel("isContainer"))
	return rv
}

// A Boolean value that indicates whether a media item is container of other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620154-container?language=objc
func (c_ ContentItem) SetContainer(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setContainer:"), value)
}

// The public name of the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620156-title?language=objc
func (c_ ContentItem) Title() string {
	rv := objc.Call[string](c_, objc.Sel("title"))
	return rv
}

// The public name of the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620156-title?language=objc
func (c_ ContentItem) SetTitle(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setTitle:"), value)
}

// A Boolean value that indicates whether the media item contains explicit content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1771744-explicitcontent?language=objc
func (c_ ContentItem) IsExplicitContent() bool {
	rv := objc.Call[bool](c_, objc.Sel("isExplicitContent"))
	return rv
}

// A Boolean value that indicates whether the media item contains explicit content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1771744-explicitcontent?language=objc
func (c_ ContentItem) SetExplicitContent(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setExplicitContent:"), value)
}

// The unique identifier for the media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620157-identifier?language=objc
func (c_ ContentItem) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether a media item is able to be played. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620158-playable?language=objc
func (c_ ContentItem) IsPlayable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isPlayable"))
	return rv
}

// A Boolean value that indicates whether a media item is able to be played. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/1620158-playable?language=objc
func (c_ ContentItem) SetPlayable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setPlayable:"), value)
}
