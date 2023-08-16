// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a Print panel object can use to get information from a printing accessory controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanelaccessorizing?language=objc
type PPrintPanelAccessorizing interface {
	// optional
	KeyPathsForValuesAffectingPreview() foundation.ISet
	HasKeyPathsForValuesAffectingPreview() bool

	// optional
	LocalizedSummaryItems() []map[PrintPanelAccessorySummaryKey]string
	HasLocalizedSummaryItems() bool
}

// A concrete type wrapper for the [PPrintPanelAccessorizing] protocol.
type PrintPanelAccessorizingWrapper struct {
	objc.Object
}

func (p_ PrintPanelAccessorizingWrapper) HasKeyPathsForValuesAffectingPreview() bool {
	return p_.RespondsToSelector(objc.Sel("keyPathsForValuesAffectingPreview"))
}

// Returns a set of strings identifying the key paths for any properties that might affect the built-in print preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanelaccessorizing/1490516-keypathsforvaluesaffectingprevie?language=objc
func (p_ PrintPanelAccessorizingWrapper) KeyPathsForValuesAffectingPreview() foundation.Set {
	rv := objc.Call[foundation.Set](p_, objc.Sel("keyPathsForValuesAffectingPreview"))
	return rv
}

func (p_ PrintPanelAccessorizingWrapper) HasLocalizedSummaryItems() bool {
	return p_.RespondsToSelector(objc.Sel("localizedSummaryItems"))
}

// Returns an array of dictionaries containing the localized user setting summary strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanelaccessorizing/1490521-localizedsummaryitems?language=objc
func (p_ PrintPanelAccessorizingWrapper) LocalizedSummaryItems() []map[PrintPanelAccessorySummaryKey]string {
	rv := objc.Call[[]map[PrintPanelAccessorySummaryKey]string](p_, objc.Sel("localizedSummaryItems"))
	return rv
}
