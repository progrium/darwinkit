// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IUserInterfaceItemSearching interface {
	// required
	LocalizedTitlesForItem(item objc.Object) []string
	ImplementsShowAllHelpTopicsForSearchString() bool
	// optional
	ShowAllHelpTopicsForSearchString(searchString string)
	// required
	SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object))
	ImplementsPerformActionForItem() bool
	// optional
	PerformActionForItem(item objc.Object)
}

type UserInterfaceItemSearchingWrapper struct {
	objc.Object
}

func (u_ UserInterfaceItemSearchingWrapper) LocalizedTitlesForItem(item objc.IObject) []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("localizedTitlesForItem:"), objc.ExtractPtr(item))
	return rv
}

func (u_ UserInterfaceItemSearchingWrapper) ImplementsShowAllHelpTopicsForSearchString() bool {
	return u_.RespondsToSelector(objc.GetSelector("showAllHelpTopicsForSearchString:"))
}

func (u_ UserInterfaceItemSearchingWrapper) ShowAllHelpTopicsForSearchString(searchString string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("showAllHelpTopicsForSearchString:"), searchString)
}

func (u_ UserInterfaceItemSearchingWrapper) SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"), searchString, resultLimit, handleMatchedItems)
}

func (u_ UserInterfaceItemSearchingWrapper) ImplementsPerformActionForItem() bool {
	return u_.RespondsToSelector(objc.GetSelector("performActionForItem:"))
}

func (u_ UserInterfaceItemSearchingWrapper) PerformActionForItem(item objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("performActionForItem:"), objc.ExtractPtr(item))
}
