// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QueuePlayer] class.
var QueuePlayerClass = _QueuePlayerClass{objc.GetClass("AVQueuePlayer")}

type _QueuePlayerClass struct {
	objc.Class
}

// An interface definition for the [QueuePlayer] class.
type IQueuePlayer interface {
	IPlayer
	Items() []PlayerItem
	RemoveAllItems()
	RemoveItem(item IPlayerItem)
	InitWithItems(items []IPlayerItem) QueuePlayer
	CanInsertItemAfterItem(item IPlayerItem, afterItem IPlayerItem) bool
	AdvanceToNextItem()
	InsertItemAfterItem(item IPlayerItem, afterItem IPlayerItem)
}

// An object that plays a sequence of player items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer?language=objc
type QueuePlayer struct {
	Player
}

func QueuePlayerFrom(ptr unsafe.Pointer) QueuePlayer {
	return QueuePlayer{
		Player: PlayerFrom(ptr),
	}
}

func (qc _QueuePlayerClass) QueuePlayerWithItems(items []IPlayerItem) QueuePlayer {
	rv := objc.Call[QueuePlayer](qc, objc.Sel("queuePlayerWithItems:"), items)
	return rv
}

// Returns an object that plays a queue of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1538384-queueplayerwithitems?language=objc
func QueuePlayer_QueuePlayerWithItems(items []IPlayerItem) QueuePlayer {
	return QueuePlayerClass.QueuePlayerWithItems(items)
}

func (qc _QueuePlayerClass) Alloc() QueuePlayer {
	rv := objc.Call[QueuePlayer](qc, objc.Sel("alloc"))
	return rv
}

func QueuePlayer_Alloc() QueuePlayer {
	return QueuePlayerClass.Alloc()
}

func (qc _QueuePlayerClass) New() QueuePlayer {
	rv := objc.Call[QueuePlayer](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQueuePlayer() QueuePlayer {
	return QueuePlayerClass.New()
}

func (q_ QueuePlayer) Init() QueuePlayer {
	rv := objc.Call[QueuePlayer](q_, objc.Sel("init"))
	return rv
}

func (qc _QueuePlayerClass) PlayerWithPlayerItem(item IPlayerItem) QueuePlayer {
	rv := objc.Call[QueuePlayer](qc, objc.Sel("playerWithPlayerItem:"), objc.Ptr(item))
	return rv
}

// Returns a new player initialized to play the specified player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1538390-playerwithplayeritem?language=objc
func QueuePlayer_PlayerWithPlayerItem(item IPlayerItem) QueuePlayer {
	return QueuePlayerClass.PlayerWithPlayerItem(item)
}

func (q_ QueuePlayer) InitWithURL(URL foundation.IURL) QueuePlayer {
	rv := objc.Call[QueuePlayer](q_, objc.Sel("initWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1385706-initwithurl?language=objc
func NewQueuePlayerWithURL(URL foundation.IURL) QueuePlayer {
	instance := QueuePlayerClass.Alloc().InitWithURL(URL)
	instance.Autorelease()
	return instance
}

func (q_ QueuePlayer) InitWithPlayerItem(item IPlayerItem) QueuePlayer {
	rv := objc.Call[QueuePlayer](q_, objc.Sel("initWithPlayerItem:"), objc.Ptr(item))
	return rv
}

// Creates a new player to play the specified player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387104-initwithplayeritem?language=objc
func NewQueuePlayerWithPlayerItem(item IPlayerItem) QueuePlayer {
	instance := QueuePlayerClass.Alloc().InitWithPlayerItem(item)
	instance.Autorelease()
	return instance
}

func (qc _QueuePlayerClass) PlayerWithURL(URL foundation.IURL) QueuePlayer {
	rv := objc.Call[QueuePlayer](qc, objc.Sel("playerWithURL:"), objc.Ptr(URL))
	return rv
}

// Returns a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1538409-playerwithurl?language=objc
func QueuePlayer_PlayerWithURL(URL foundation.IURL) QueuePlayer {
	return QueuePlayerClass.PlayerWithURL(URL)
}

// Returns an array of the currently enqueued items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1390539-items?language=objc
func (q_ QueuePlayer) Items() []PlayerItem {
	rv := objc.Call[[]PlayerItem](q_, objc.Sel("items"))
	return rv
}

// Removes all player items from the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1385788-removeallitems?language=objc
func (q_ QueuePlayer) RemoveAllItems() {
	objc.Call[objc.Void](q_, objc.Sel("removeAllItems"))
}

// Removes a given player item from the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1387400-removeitem?language=objc
func (q_ QueuePlayer) RemoveItem(item IPlayerItem) {
	objc.Call[objc.Void](q_, objc.Sel("removeItem:"), objc.Ptr(item))
}

// Creates an object that plays a queue of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1389345-initwithitems?language=objc
func (q_ QueuePlayer) InitWithItems(items []IPlayerItem) QueuePlayer {
	rv := objc.Call[QueuePlayer](q_, objc.Sel("initWithItems:"), items)
	return rv
}

// Returns a Boolean value that indicates whether you can insert a player item into the player’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1387289-caninsertitem?language=objc
func (q_ QueuePlayer) CanInsertItemAfterItem(item IPlayerItem, afterItem IPlayerItem) bool {
	rv := objc.Call[bool](q_, objc.Sel("canInsertItem:afterItem:"), objc.Ptr(item), objc.Ptr(afterItem))
	return rv
}

// Ends playback of the current item and starts playback of the next item in the player’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1389318-advancetonextitem?language=objc
func (q_ QueuePlayer) AdvanceToNextItem() {
	objc.Call[objc.Void](q_, objc.Sel("advanceToNextItem"))
}

// Inserts a player item after another player item in the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueueplayer/1388543-insertitem?language=objc
func (q_ QueuePlayer) InsertItemAfterItem(item IPlayerItem, afterItem IPlayerItem) {
	objc.Call[objc.Void](q_, objc.Sel("insertItem:afterItem:"), objc.Ptr(item), objc.Ptr(afterItem))
}
