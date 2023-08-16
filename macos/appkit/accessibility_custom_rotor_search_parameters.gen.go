// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccessibilityCustomRotorSearchParameters] class.
var AccessibilityCustomRotorSearchParametersClass = _AccessibilityCustomRotorSearchParametersClass{objc.GetClass("NSAccessibilityCustomRotorSearchParameters")}

type _AccessibilityCustomRotorSearchParametersClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomRotorSearchParameters] class.
type IAccessibilityCustomRotorSearchParameters interface {
	objc.IObject
	SearchDirection() AccessibilityCustomRotorSearchDirection
	SetSearchDirection(value AccessibilityCustomRotorSearchDirection)
	CurrentItem() AccessibilityCustomRotorItemResult
	SetCurrentItem(value IAccessibilityCustomRotorItemResult)
	FilterString() string
	SetFilterString(value string)
}

// Search parameters for a custom rotor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters?language=objc
type AccessibilityCustomRotorSearchParameters struct {
	objc.Object
}

func AccessibilityCustomRotorSearchParametersFrom(ptr unsafe.Pointer) AccessibilityCustomRotorSearchParameters {
	return AccessibilityCustomRotorSearchParameters{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AccessibilityCustomRotorSearchParametersClass) Alloc() AccessibilityCustomRotorSearchParameters {
	rv := objc.Call[AccessibilityCustomRotorSearchParameters](ac, objc.Sel("alloc"))
	return rv
}

func AccessibilityCustomRotorSearchParameters_Alloc() AccessibilityCustomRotorSearchParameters {
	return AccessibilityCustomRotorSearchParametersClass.Alloc()
}

func (ac _AccessibilityCustomRotorSearchParametersClass) New() AccessibilityCustomRotorSearchParameters {
	rv := objc.Call[AccessibilityCustomRotorSearchParameters](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccessibilityCustomRotorSearchParameters() AccessibilityCustomRotorSearchParameters {
	return AccessibilityCustomRotorSearchParametersClass.New()
}

func (a_ AccessibilityCustomRotorSearchParameters) Init() AccessibilityCustomRotorSearchParameters {
	rv := objc.Call[AccessibilityCustomRotorSearchParameters](a_, objc.Sel("init"))
	return rv
}

// The direction to search for an item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876322-searchdirection?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) SearchDirection() AccessibilityCustomRotorSearchDirection {
	rv := objc.Call[AccessibilityCustomRotorSearchDirection](a_, objc.Sel("searchDirection"))
	return rv
}

// The direction to search for an item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876322-searchdirection?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) SetSearchDirection(value AccessibilityCustomRotorSearchDirection) {
	objc.Call[objc.Void](a_, objc.Sel("setSearchDirection:"), value)
}

// The current item that determines where the search starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876332-currentitem?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) CurrentItem() AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](a_, objc.Sel("currentItem"))
	return rv
}

// The current item that determines where the search starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876332-currentitem?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) SetCurrentItem(value IAccessibilityCustomRotorItemResult) {
	objc.Call[objc.Void](a_, objc.Sel("setCurrentItem:"), objc.Ptr(value))
}

// A string of text to filter the results against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876328-filterstring?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) FilterString() string {
	rv := objc.Call[string](a_, objc.Sel("filterString"))
	return rv
}

// A string of text to filter the results against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchparameters/2876328-filterstring?language=objc
func (a_ AccessibilityCustomRotorSearchParameters) SetFilterString(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setFilterString:"), value)
}
