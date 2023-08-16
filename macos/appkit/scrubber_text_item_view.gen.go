// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberTextItemView] class.
var ScrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}

type _ScrubberTextItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	IScrubberItemView
	TextField() TextField
	Title() string
	SetTitle(value string)
}

// A concrete view subclass for displaying text for an item in a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview?language=objc
type ScrubberTextItemView struct {
	ScrubberItemView
}

func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}

func (sc _ScrubberTextItemViewClass) Alloc() ScrubberTextItemView {
	rv := objc.Call[ScrubberTextItemView](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberTextItemView_Alloc() ScrubberTextItemView {
	return ScrubberTextItemViewClass.Alloc()
}

func (sc _ScrubberTextItemViewClass) New() ScrubberTextItemView {
	rv := objc.Call[ScrubberTextItemView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberTextItemView() ScrubberTextItemView {
	return ScrubberTextItemViewClass.New()
}

func (s_ ScrubberTextItemView) Init() ScrubberTextItemView {
	rv := objc.Call[ScrubberTextItemView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScrubberTextItemView) InitWithFrame(frameRect foundation.Rect) ScrubberTextItemView {
	rv := objc.Call[ScrubberTextItemView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func ScrubberTextItemView_InitWithFrame(frameRect foundation.Rect) ScrubberTextItemView {
	return ScrubberTextItemViewClass.Alloc().InitWithFrame(frameRect)
}

// The text field that the scrubber item uses to display its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/2544844-textfield?language=objc
func (s_ ScrubberTextItemView) TextField() TextField {
	rv := objc.Call[TextField](s_, objc.Sel("textField"))
	return rv
}

// The text displayed for the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/2544795-title?language=objc
func (s_ ScrubberTextItemView) Title() string {
	rv := objc.Call[string](s_, objc.Sel("title"))
	return rv
}

// The text displayed for the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/2544795-title?language=objc
func (s_ ScrubberTextItemView) SetTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTitle:"), value)
}
