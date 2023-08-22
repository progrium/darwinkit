// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RoutePickerView] class.
var RoutePickerViewClass = _RoutePickerViewClass{objc.GetClass("AVRoutePickerView")}

type _RoutePickerViewClass struct {
	objc.Class
}

// An interface definition for the [RoutePickerView] class.
type IRoutePickerView interface {
	appkit.IView
	SetRoutePickerButtonColorForState(color appkit.IColor, state RoutePickerViewButtonState)
	RoutePickerButtonColorForState(state RoutePickerViewButtonState) appkit.Color
	Player() avfoundation.Player
	SetPlayer(value avfoundation.IPlayer)
	Delegate() RoutePickerViewDelegateWrapper
	SetDelegate(value PRoutePickerViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsRoutePickerButtonBordered() bool
	SetRoutePickerButtonBordered(value bool)
}

// A view that presents a list of nearby media receivers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview?language=objc
type RoutePickerView struct {
	appkit.View
}

func RoutePickerViewFrom(ptr unsafe.Pointer) RoutePickerView {
	return RoutePickerView{
		View: appkit.ViewFrom(ptr),
	}
}

func (rc _RoutePickerViewClass) Alloc() RoutePickerView {
	rv := objc.Call[RoutePickerView](rc, objc.Sel("alloc"))
	return rv
}

func RoutePickerView_Alloc() RoutePickerView {
	return RoutePickerViewClass.Alloc()
}

func (rc _RoutePickerViewClass) New() RoutePickerView {
	rv := objc.Call[RoutePickerView](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRoutePickerView() RoutePickerView {
	return RoutePickerViewClass.New()
}

func (r_ RoutePickerView) Init() RoutePickerView {
	rv := objc.Call[RoutePickerView](r_, objc.Sel("init"))
	return rv
}

func (r_ RoutePickerView) InitWithFrame(frameRect foundation.Rect) RoutePickerView {
	rv := objc.Call[RoutePickerView](r_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewRoutePickerViewWithFrame(frameRect foundation.Rect) RoutePickerView {
	instance := RoutePickerViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Sets the route picker button color for the specified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915792-setroutepickerbuttoncolor?language=objc
func (r_ RoutePickerView) SetRoutePickerButtonColorForState(color appkit.IColor, state RoutePickerViewButtonState) {
	objc.Call[objc.Void](r_, objc.Sel("setRoutePickerButtonColor:forState:"), objc.Ptr(color), state)
}

// Returns the color of the picker button for the specified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915793-routepickerbuttoncolorforstate?language=objc
func (r_ RoutePickerView) RoutePickerButtonColorForState(state RoutePickerViewButtonState) appkit.Color {
	rv := objc.Call[appkit.Color](r_, objc.Sel("routePickerButtonColorForState:"), state)
	return rv
}

// The player object to perform routing operations for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/3201361-player?language=objc
func (r_ RoutePickerView) Player() avfoundation.Player {
	rv := objc.Call[avfoundation.Player](r_, objc.Sel("player"))
	return rv
}

// The player object to perform routing operations for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/3201361-player?language=objc
func (r_ RoutePickerView) SetPlayer(value avfoundation.IPlayer) {
	objc.Call[objc.Void](r_, objc.Sel("setPlayer:"), objc.Ptr(value))
}

// The delegate object for the route picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915799-delegate?language=objc
func (r_ RoutePickerView) Delegate() RoutePickerViewDelegateWrapper {
	rv := objc.Call[RoutePickerViewDelegateWrapper](r_, objc.Sel("delegate"))
	return rv
}

// The delegate object for the route picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915799-delegate?language=objc
func (r_ RoutePickerView) SetDelegate(value PRoutePickerViewDelegate) {
	po0 := objc.WrapAsProtocol("AVRoutePickerViewDelegate", value)
	objc.SetAssociatedObject(r_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](r_, objc.Sel("setDelegate:"), po0)
}

// The delegate object for the route picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915799-delegate?language=objc
func (r_ RoutePickerView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether the route picker button has a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915795-routepickerbuttonbordered?language=objc
func (r_ RoutePickerView) IsRoutePickerButtonBordered() bool {
	rv := objc.Call[bool](r_, objc.Sel("isRoutePickerButtonBordered"))
	return rv
}

// A Boolean value that indicates whether the route picker button has a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/2915795-routepickerbuttonbordered?language=objc
func (r_ RoutePickerView) SetRoutePickerButtonBordered(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRoutePickerButtonBordered:"), value)
}
