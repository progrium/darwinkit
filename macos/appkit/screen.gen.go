// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ScreenClass = _ScreenClass{objc.GetClass("NSScreen")}

type _ScreenClass struct {
	objc.Class
}

type IScreen interface {
	objc.IObject
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	Depth() WindowDepth
	Frame() foundation.Rect
	SupportedWindowDepths() *WindowDepth
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	ColorSpace() ColorSpace
	LocalizedName() string
	BackingScaleFactor() float64
	VisibleFrame() foundation.Rect
	SafeAreaInsets() foundation.EdgeInsets
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64
	MaximumExtendedDynamicRangeColorComponentValue() float64
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64
	MaximumFramesPerSecond() int
	MinimumRefreshInterval() foundation.TimeInterval
	MaximumRefreshInterval() foundation.TimeInterval
	DisplayUpdateGranularity() foundation.TimeInterval
	LastDisplayUpdateTimestamp() foundation.TimeInterval
	AuxiliaryTopLeftArea() foundation.Rect
	AuxiliaryTopRightArea() foundation.Rect
}

type Screen struct {
	objc.Object
}

func MakeScreen(ptr unsafe.Pointer) Screen {
	return Screen{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScreenClass) Alloc() Screen {
	rv := objc.CallMethod[Screen](sc, objc.GetSelector("alloc"))
	return rv
}

func Screen_Alloc() Screen {
	return ScreenClass.Alloc()
}

func (sc _ScreenClass) New() Screen {
	rv := objc.CallMethod[Screen](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScreen() Screen {
	return ScreenClass.New()
}

func Screen_New() Screen {
	return ScreenClass.New()
}

func (s_ Screen) Init() Screen {
	rv := objc.CallMethod[Screen](s_, objc.GetSelector("init"))
	return rv
}

func Screen_Init() Screen {
	return ScreenClass.Alloc().Init()
}

func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

func (s_ Screen) BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("backingAlignedRect:options:"), rect, options)
	return rv
}

func (s_ Screen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("convertRectFromBacking:"), rect)
	return rv
}

func (s_ Screen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("convertRectToBacking:"), rect)
	return rv
}

func (sc _ScreenClass) MainScreen() Screen {
	rv := objc.CallMethod[Screen](sc, objc.GetSelector("mainScreen"))
	return rv
}

func Screen_MainScreen() Screen {
	return ScreenClass.MainScreen()
}

func (sc _ScreenClass) DeepestScreen() Screen {
	rv := objc.CallMethod[Screen](sc, objc.GetSelector("deepestScreen"))
	return rv
}

func Screen_DeepestScreen() Screen {
	return ScreenClass.DeepestScreen()
}

func (sc _ScreenClass) Screens() []Screen {
	rv := objc.CallMethod[[]Screen](sc, objc.GetSelector("screens"))
	// TODO: convert slice items...
	return rv
}

func Screen_Screens() []Screen {
	return ScreenClass.Screens()
}

func (s_ Screen) Depth() WindowDepth {
	rv := objc.CallMethod[WindowDepth](s_, objc.GetSelector("depth"))
	return rv
}

func (s_ Screen) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("frame"))
	return rv
}

func (s_ Screen) SupportedWindowDepths() *WindowDepth {
	rv := objc.CallMethod[*WindowDepth](s_, objc.GetSelector("supportedWindowDepths"))
	return rv
}

func (s_ Screen) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](s_, objc.GetSelector("deviceDescription"))
	return rv
}

func (s_ Screen) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](s_, objc.GetSelector("colorSpace"))
	return rv
}

func (s_ Screen) LocalizedName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("localizedName"))
	return rv
}

func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.CallMethod[bool](sc, objc.GetSelector("screensHaveSeparateSpaces"))
	return rv
}

func Screen_ScreensHaveSeparateSpaces() bool {
	return ScreenClass.ScreensHaveSeparateSpaces()
}

func (s_ Screen) BackingScaleFactor() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("backingScaleFactor"))
	return rv
}

func (s_ Screen) VisibleFrame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("visibleFrame"))
	return rv
}

func (s_ Screen) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.GetSelector("safeAreaInsets"))
	return rv
}

func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumFramesPerSecond() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("maximumFramesPerSecond"))
	return rv
}

func (s_ Screen) MinimumRefreshInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("minimumRefreshInterval"))
	return rv
}

func (s_ Screen) MaximumRefreshInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("maximumRefreshInterval"))
	return rv
}

func (s_ Screen) DisplayUpdateGranularity() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("displayUpdateGranularity"))
	return rv
}

func (s_ Screen) LastDisplayUpdateTimestamp() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("lastDisplayUpdateTimestamp"))
	return rv
}

func (s_ Screen) AuxiliaryTopLeftArea() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("auxiliaryTopLeftArea"))
	return rv
}

func (s_ Screen) AuxiliaryTopRightArea() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("auxiliaryTopRightArea"))
	return rv
}
