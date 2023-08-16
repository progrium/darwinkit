// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HelpManager] class.
var HelpManagerClass = _HelpManagerClass{objc.GetClass("NSHelpManager")}

type _HelpManagerClass struct {
	objc.Class
}

// An interface definition for the [HelpManager] class.
type IHelpManager interface {
	objc.IObject
	RemoveContextHelpForObject(object objc.IObject)
	OpenHelpAnchorInBook(anchor HelpAnchorName, book HelpBookName)
	ContextHelpForObject(object objc.IObject) foundation.AttributedString
	ShowContextHelpForObjectLocationHint(object objc.IObject, pt foundation.Point) bool
	RegisterBooksInBundle(bundle foundation.IBundle) bool
	SetContextHelpForObject(attrString foundation.IAttributedString, object objc.IObject)
	FindStringInBook(query string, book HelpBookName)
}

// An object for displaying online help for an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager?language=objc
type HelpManager struct {
	objc.Object
}

func HelpManagerFrom(ptr unsafe.Pointer) HelpManager {
	return HelpManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HelpManagerClass) Alloc() HelpManager {
	rv := objc.Call[HelpManager](hc, objc.Sel("alloc"))
	return rv
}

func HelpManager_Alloc() HelpManager {
	return HelpManagerClass.Alloc()
}

func (hc _HelpManagerClass) New() HelpManager {
	rv := objc.Call[HelpManager](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHelpManager() HelpManager {
	return HelpManagerClass.New()
}

func (h_ HelpManager) Init() HelpManager {
	rv := objc.Call[HelpManager](h_, objc.Sel("init"))
	return rv
}

// Removes the association between an object and its context-sensitive help. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500906-removecontexthelpforobject?language=objc
func (h_ HelpManager) RemoveContextHelpForObject(object objc.IObject) {
	objc.Call[objc.Void](h_, objc.Sel("removeContextHelpForObject:"), object)
}

// Finds and displays the text at the given anchor location in the given book. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500908-openhelpanchor?language=objc
func (h_ HelpManager) OpenHelpAnchorInBook(anchor HelpAnchorName, book HelpBookName) {
	objc.Call[objc.Void](h_, objc.Sel("openHelpAnchor:inBook:"), anchor, book)
}

// Returns context-sensitive help for an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500919-contexthelpforobject?language=objc
func (h_ HelpManager) ContextHelpForObject(object objc.IObject) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](h_, objc.Sel("contextHelpForObject:"), object)
	return rv
}

// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500930-showcontexthelpforobject?language=objc
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objc.IObject, pt foundation.Point) bool {
	rv := objc.Call[bool](h_, objc.Sel("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}

// Registers one or more help books in the given bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500914-registerbooksinbundle?language=objc
func (h_ HelpManager) RegisterBooksInBundle(bundle foundation.IBundle) bool {
	rv := objc.Call[bool](h_, objc.Sel("registerBooksInBundle:"), objc.Ptr(bundle))
	return rv
}

// Associates help content with an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500921-setcontexthelp?language=objc
func (h_ HelpManager) SetContextHelpForObject(attrString foundation.IAttributedString, object objc.IObject) {
	objc.Call[objc.Void](h_, objc.Sel("setContextHelp:forObject:"), objc.Ptr(attrString), object)
}

// Performs a search for the specified string in the specified book. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500904-findstring?language=objc
func (h_ HelpManager) FindStringInBook(query string, book HelpBookName) {
	objc.Call[objc.Void](h_, objc.Sel("findString:inBook:"), query, book)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/2870247-contexthelpmodeactive?language=objc
func (hc _HelpManagerClass) ContextHelpModeActive() bool {
	rv := objc.Call[bool](hc, objc.Sel("contextHelpModeActive"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/2870247-contexthelpmodeactive?language=objc
func HelpManager_ContextHelpModeActive() bool {
	return HelpManagerClass.ContextHelpModeActive()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/2870247-contexthelpmodeactive?language=objc
func (hc _HelpManagerClass) SetContextHelpModeActive(value bool) {
	objc.Call[objc.Void](hc, objc.Sel("setContextHelpModeActive:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/2870247-contexthelpmodeactive?language=objc
func HelpManager_SetContextHelpModeActive(value bool) {
	HelpManagerClass.SetContextHelpModeActive(value)
}

// Returns the shared NSHelpManager instance, creating it if it does not already exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500924-sharedhelpmanager?language=objc
func (hc _HelpManagerClass) SharedHelpManager() HelpManager {
	rv := objc.Call[HelpManager](hc, objc.Sel("sharedHelpManager"))
	return rv
}

// Returns the shared NSHelpManager instance, creating it if it does not already exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanager/1500924-sharedhelpmanager?language=objc
func HelpManager_SharedHelpManager() HelpManager {
	return HelpManagerClass.SharedHelpManager()
}
