// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [AccessibilityCustomRotorItemResult] class.
var AccessibilityCustomRotorItemResultClass = _AccessibilityCustomRotorItemResultClass{objc.GetClass("NSAccessibilityCustomRotorItemResult")}

type _AccessibilityCustomRotorItemResultClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomRotorItemResult] class.
type IAccessibilityCustomRotorItemResult interface {
	objc.IObject
	CustomLabel() string
	SetCustomLabel(value string)
	TargetElement() AccessibilityElementObject
	TargetRange() foundation.Range
	SetTargetRange(value foundation.Range)
	ItemLoadingToken() AccessibilityLoadingToken
}

// A target accessibility element that a custom rotor references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult?language=objc
type AccessibilityCustomRotorItemResult struct {
	objc.Object
}

func AccessibilityCustomRotorItemResultFrom(ptr unsafe.Pointer) AccessibilityCustomRotorItemResult {
	return AccessibilityCustomRotorItemResult{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AccessibilityCustomRotorItemResult) InitWithTargetElement(targetElement PAccessibilityElement) AccessibilityCustomRotorItemResult {
	po0 := objc.WrapAsProtocol("NSAccessibilityElement", targetElement)
	rv := objc.Call[AccessibilityCustomRotorItemResult](a_, objc.Sel("initWithTargetElement:"), po0)
	return rv
}

// Creates an item result with the specified target element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876308-initwithtargetelement?language=objc
func NewAccessibilityCustomRotorItemResultWithTargetElement(targetElement PAccessibilityElement) AccessibilityCustomRotorItemResult {
	instance := AccessibilityCustomRotorItemResultClass.Alloc().InitWithTargetElement(targetElement)
	instance.Autorelease()
	return instance
}

func (a_ AccessibilityCustomRotorItemResult) InitWithItemLoadingTokenCustomLabel(itemLoadingToken AccessibilityLoadingToken, customLabel string) AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](a_, objc.Sel("initWithItemLoadingToken:customLabel:"), itemLoadingToken, customLabel)
	return rv
}

// Creates an item result with the specified item load token and custom label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2890782-initwithitemloadingtoken?language=objc
func NewAccessibilityCustomRotorItemResultWithItemLoadingTokenCustomLabel(itemLoadingToken AccessibilityLoadingToken, customLabel string) AccessibilityCustomRotorItemResult {
	instance := AccessibilityCustomRotorItemResultClass.Alloc().InitWithItemLoadingTokenCustomLabel(itemLoadingToken, customLabel)
	instance.Autorelease()
	return instance
}

func (ac _AccessibilityCustomRotorItemResultClass) Alloc() AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AccessibilityCustomRotorItemResultClass) New() AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccessibilityCustomRotorItemResult() AccessibilityCustomRotorItemResult {
	return AccessibilityCustomRotorItemResultClass.New()
}

func (a_ AccessibilityCustomRotorItemResult) Init() AccessibilityCustomRotorItemResult {
	rv := objc.Call[AccessibilityCustomRotorItemResult](a_, objc.Sel("init"))
	return rv
}

// A localized label to use instead of the default item label to describe the item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876298-customlabel?language=objc
func (a_ AccessibilityCustomRotorItemResult) CustomLabel() string {
	rv := objc.Call[string](a_, objc.Sel("customLabel"))
	return rv
}

// A localized label to use instead of the default item label to describe the item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876298-customlabel?language=objc
func (a_ AccessibilityCustomRotorItemResult) SetCustomLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setCustomLabel:"), value)
}

// A target element that references an element to message for accessibility properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876335-targetelement?language=objc
func (a_ AccessibilityCustomRotorItemResult) TargetElement() AccessibilityElementObject {
	rv := objc.Call[AccessibilityElementObject](a_, objc.Sel("targetElement"))
	return rv
}

// A range that specifies the area of interest for text-based elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876323-targetrange?language=objc
func (a_ AccessibilityCustomRotorItemResult) TargetRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("targetRange"))
	return rv
}

// A range that specifies the area of interest for text-based elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2876323-targetrange?language=objc
func (a_ AccessibilityCustomRotorItemResult) SetTargetRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setTargetRange:"), value)
}

// A token to determine which item to return. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotoritemresult/2890781-itemloadingtoken?language=objc
func (a_ AccessibilityCustomRotorItemResult) ItemLoadingToken() AccessibilityLoadingToken {
	rv := objc.Call[AccessibilityLoadingToken](a_, objc.Sel("itemLoadingToken"))
	return rv
}
