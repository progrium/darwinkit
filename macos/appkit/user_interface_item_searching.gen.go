// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods an app can implement to provide Spotlight for Help for its own custom help data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching?language=objc
type PUserInterfaceItemSearching interface {
	// optional
	LocalizedTitlesForItem(item objc.Object) []string
	HasLocalizedTitlesForItem() bool

	// optional
	PerformActionForItem(item objc.Object)
	HasPerformActionForItem() bool

	// optional
	SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object))
	HasSearchForItemsWithSearchStringResultLimitMatchedItemHandler() bool

	// optional
	ShowAllHelpTopicsForSearchString(searchString string)
	HasShowAllHelpTopicsForSearchString() bool
}

// A concrete type wrapper for the [PUserInterfaceItemSearching] protocol.
type UserInterfaceItemSearchingWrapper struct {
	objc.Object
}

func (u_ UserInterfaceItemSearchingWrapper) HasLocalizedTitlesForItem() bool {
	return u_.RespondsToSelector(objc.Sel("localizedTitlesForItem:"))
}

// Returns an array of localized strings that will form the help menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420814-localizedtitlesforitem?language=objc
func (u_ UserInterfaceItemSearchingWrapper) LocalizedTitlesForItem(item objc.IObject) []string {
	rv := objc.Call[[]string](u_, objc.Sel("localizedTitlesForItem:"), item)
	return rv
}

func (u_ UserInterfaceItemSearchingWrapper) HasPerformActionForItem() bool {
	return u_.RespondsToSelector(objc.Sel("performActionForItem:"))
}

// Invoked when the user selects a search result in Help menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420812-performactionforitem?language=objc
func (u_ UserInterfaceItemSearchingWrapper) PerformActionForItem(item objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("performActionForItem:"), item)
}

func (u_ UserInterfaceItemSearchingWrapper) HasSearchForItemsWithSearchStringResultLimitMatchedItemHandler() bool {
	return u_.RespondsToSelector(objc.Sel("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"))
}

// Search for the specified items, with the result limit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420816-searchforitemswithsearchstring?language=objc
func (u_ UserInterfaceItemSearchingWrapper) SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object)) {
	objc.Call[objc.Void](u_, objc.Sel("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"), searchString, resultLimit, handleMatchedItems)
}

func (u_ UserInterfaceItemSearchingWrapper) HasShowAllHelpTopicsForSearchString() bool {
	return u_.RespondsToSelector(objc.Sel("showAllHelpTopicsForSearchString:"))
}

// If this method is implemented, a "Show All Help Topics" item will appear in the menu and this method is called when the user selects it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420806-showallhelptopicsforsearchstring?language=objc
func (u_ UserInterfaceItemSearchingWrapper) ShowAllHelpTopicsForSearchString(searchString string) {
	objc.Call[objc.Void](u_, objc.Sel("showAllHelpTopicsForSearchString:"), searchString)
}
