// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccessibilityCustomRotor] class.
var AccessibilityCustomRotorClass = _AccessibilityCustomRotorClass{objc.GetClass("NSAccessibilityCustomRotor")}

type _AccessibilityCustomRotorClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomRotor] class.
type IAccessibilityCustomRotor interface {
	objc.IObject
	ItemSearchDelegate() AccessibilityCustomRotorItemSearchDelegateWrapper
	SetItemSearchDelegate(value PAccessibilityCustomRotorItemSearchDelegate)
	SetItemSearchDelegateObject(valueObject objc.IObject)
	ItemLoadingDelegate() AccessibilityElementLoadingWrapper
	SetItemLoadingDelegate(value PAccessibilityElementLoading)
	SetItemLoadingDelegateObject(valueObject objc.IObject)
	Type() AccessibilityCustomRotorType
	SetType(value AccessibilityCustomRotorType)
	Label() string
	SetLabel(value string)
}

// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor?language=objc
type AccessibilityCustomRotor struct {
	objc.Object
}

func AccessibilityCustomRotorFrom(ptr unsafe.Pointer) AccessibilityCustomRotor {
	return AccessibilityCustomRotor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AccessibilityCustomRotor) InitWithRotorTypeItemSearchDelegate(rotorType AccessibilityCustomRotorType, itemSearchDelegate PAccessibilityCustomRotorItemSearchDelegate) AccessibilityCustomRotor {
	po1 := objc.WrapAsProtocol("NSAccessibilityCustomRotorItemSearchDelegate", itemSearchDelegate)
	rv := objc.Call[AccessibilityCustomRotor](a_, objc.Sel("initWithRotorType:itemSearchDelegate:"), rotorType, po1)
	return rv
}

// Creates a custom rotor with the specified rotor type and item search delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876299-initwithrotortype?language=objc
func NewAccessibilityCustomRotorWithRotorTypeItemSearchDelegate(rotorType AccessibilityCustomRotorType, itemSearchDelegate PAccessibilityCustomRotorItemSearchDelegate) AccessibilityCustomRotor {
	instance := AccessibilityCustomRotorClass.Alloc().InitWithRotorTypeItemSearchDelegate(rotorType, itemSearchDelegate)
	instance.Autorelease()
	return instance
}

func (a_ AccessibilityCustomRotor) InitWithLabelItemSearchDelegate(label string, itemSearchDelegate PAccessibilityCustomRotorItemSearchDelegate) AccessibilityCustomRotor {
	po1 := objc.WrapAsProtocol("NSAccessibilityCustomRotorItemSearchDelegate", itemSearchDelegate)
	rv := objc.Call[AccessibilityCustomRotor](a_, objc.Sel("initWithLabel:itemSearchDelegate:"), label, po1)
	return rv
}

// Creates a custom rotor with the specified label and item search delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876333-initwithlabel?language=objc
func NewAccessibilityCustomRotorWithLabelItemSearchDelegate(label string, itemSearchDelegate PAccessibilityCustomRotorItemSearchDelegate) AccessibilityCustomRotor {
	instance := AccessibilityCustomRotorClass.Alloc().InitWithLabelItemSearchDelegate(label, itemSearchDelegate)
	instance.Autorelease()
	return instance
}

func (ac _AccessibilityCustomRotorClass) Alloc() AccessibilityCustomRotor {
	rv := objc.Call[AccessibilityCustomRotor](ac, objc.Sel("alloc"))
	return rv
}

func AccessibilityCustomRotor_Alloc() AccessibilityCustomRotor {
	return AccessibilityCustomRotorClass.Alloc()
}

func (ac _AccessibilityCustomRotorClass) New() AccessibilityCustomRotor {
	rv := objc.Call[AccessibilityCustomRotor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccessibilityCustomRotor() AccessibilityCustomRotor {
	return AccessibilityCustomRotorClass.New()
}

func (a_ AccessibilityCustomRotor) Init() AccessibilityCustomRotor {
	rv := objc.Call[AccessibilityCustomRotor](a_, objc.Sel("init"))
	return rv
}

// The delegate for finding the next item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876307-itemsearchdelegate?language=objc
func (a_ AccessibilityCustomRotor) ItemSearchDelegate() AccessibilityCustomRotorItemSearchDelegateWrapper {
	rv := objc.Call[AccessibilityCustomRotorItemSearchDelegateWrapper](a_, objc.Sel("itemSearchDelegate"))
	return rv
}

// The delegate for finding the next item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876307-itemsearchdelegate?language=objc
func (a_ AccessibilityCustomRotor) SetItemSearchDelegate(value PAccessibilityCustomRotorItemSearchDelegate) {
	po0 := objc.WrapAsProtocol("NSAccessibilityCustomRotorItemSearchDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setItemSearchDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setItemSearchDelegate:"), po0)
}

// The delegate for finding the next item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876307-itemsearchdelegate?language=objc
func (a_ AccessibilityCustomRotor) SetItemSearchDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setItemSearchDelegate:"), objc.Ptr(valueObject))
}

// The delegate for loading item results that don’t have a backing UI element at loading time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2890783-itemloadingdelegate?language=objc
func (a_ AccessibilityCustomRotor) ItemLoadingDelegate() AccessibilityElementLoadingWrapper {
	rv := objc.Call[AccessibilityElementLoadingWrapper](a_, objc.Sel("itemLoadingDelegate"))
	return rv
}

// The delegate for loading item results that don’t have a backing UI element at loading time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2890783-itemloadingdelegate?language=objc
func (a_ AccessibilityCustomRotor) SetItemLoadingDelegate(value PAccessibilityElementLoading) {
	po0 := objc.WrapAsProtocol("NSAccessibilityElementLoading", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setItemLoadingDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setItemLoadingDelegate:"), po0)
}

// The delegate for loading item results that don’t have a backing UI element at loading time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2890783-itemloadingdelegate?language=objc
func (a_ AccessibilityCustomRotor) SetItemLoadingDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setItemLoadingDelegate:"), objc.Ptr(valueObject))
}

// The type of content that the rotor represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876321-type?language=objc
func (a_ AccessibilityCustomRotor) Type() AccessibilityCustomRotorType {
	rv := objc.Call[AccessibilityCustomRotorType](a_, objc.Sel("type"))
	return rv
}

// The type of content that the rotor represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876321-type?language=objc
func (a_ AccessibilityCustomRotor) SetType(value AccessibilityCustomRotorType) {
	objc.Call[objc.Void](a_, objc.Sel("setType:"), value)
}

// The localized label that assistive apps use to describe the custom rotor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876331-label?language=objc
func (a_ AccessibilityCustomRotor) Label() string {
	rv := objc.Call[string](a_, objc.Sel("label"))
	return rv
}

// The localized label that assistive apps use to describe the custom rotor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/2876331-label?language=objc
func (a_ AccessibilityCustomRotor) SetLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLabel:"), value)
}
