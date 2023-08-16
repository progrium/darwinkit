// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A delegate for a custom rotor that finds the next item result after performing a search with the specified search parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate?language=objc
type PAccessibilityCustomRotorItemSearchDelegate interface {
	// optional
	RotorResultForSearchParameters(rotor AccessibilityCustomRotor, searchParameters AccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult
	HasRotorResultForSearchParameters() bool
}

// A delegate implementation builder for the [PAccessibilityCustomRotorItemSearchDelegate] protocol.
type AccessibilityCustomRotorItemSearchDelegate struct {
	_RotorResultForSearchParameters func(rotor AccessibilityCustomRotor, searchParameters AccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult
}

func (di *AccessibilityCustomRotorItemSearchDelegate) HasRotorResultForSearchParameters() bool {
	return di._RotorResultForSearchParameters != nil
}

// Performs a search with the specified search parameters and returns the item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate/2876324-rotor?language=objc
func (di *AccessibilityCustomRotorItemSearchDelegate) SetRotorResultForSearchParameters(f func(rotor AccessibilityCustomRotor, searchParameters AccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult) {
	di._RotorResultForSearchParameters = f
}

// Performs a search with the specified search parameters and returns the item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate/2876324-rotor?language=objc
func (di *AccessibilityCustomRotorItemSearchDelegate) RotorResultForSearchParameters(rotor AccessibilityCustomRotor, searchParameters AccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult {
	return di._RotorResultForSearchParameters(rotor, searchParameters)
}

// A concrete type wrapper for the [PAccessibilityCustomRotorItemSearchDelegate] protocol.
type AccessibilityCustomRotorItemSearchDelegateWrapper struct {
	objc.Object
}

func (a_ AccessibilityCustomRotorItemSearchDelegateWrapper) HasRotorResultForSearchParameters() bool {
	return a_.RespondsToSelector(objc.Sel("rotor:resultForSearchParameters:"))
}

// Performs a search with the specified search parameters and returns the item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemsearchdelegate/2876324-rotor?language=objc
func (a_ AccessibilityCustomRotorItemSearchDelegateWrapper) RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](a_, objc.Sel("rotor:resultForSearchParameters:"), objc.Ptr(rotor), objc.Ptr(searchParameters))
	return rv
}
