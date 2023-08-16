// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextContainer] class.
var TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}

type _TextContainerClass struct {
	objc.Class
}

// An interface definition for the [TextContainer] class.
type ITextContainer interface {
	objc.IObject
	ReplaceLayoutManager(newLayoutManager ILayoutManager)
	TextLayoutManager() TextLayoutManager
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	TextView() TextView
	SetTextView(value ITextView)
	IsSimpleRectangularTextContainer() bool
	LayoutManager() LayoutManager
	SetLayoutManager(value ILayoutManager)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	ExclusionPaths() []BezierPath
	SetExclusionPaths(value []IBezierPath)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	Size() coregraphics.Size
	SetSize(value coregraphics.Size)
}

// A region where text layout occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer?language=objc
type TextContainer struct {
	objc.Object
}

func TextContainerFrom(ptr unsafe.Pointer) TextContainer {
	return TextContainer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextContainer) InitWithSize(size coregraphics.Size) TextContainer {
	rv := objc.Call[TextContainer](t_, objc.Sel("initWithSize:"), size)
	return rv
}

// Initializes a text container with a specified bounding rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444529-initwithsize?language=objc
func TextContainer_InitWithSize(size coregraphics.Size) TextContainer {
	return TextContainerClass.Alloc().InitWithSize(size)
}

func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.Call[TextContainer](tc, objc.Sel("alloc"))
	return rv
}

func TextContainer_Alloc() TextContainer {
	return TextContainerClass.Alloc()
}

func (tc _TextContainerClass) New() TextContainer {
	rv := objc.Call[TextContainer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

func (t_ TextContainer) Init() TextContainer {
	rv := objc.Call[TextContainer](t_, objc.Sel("init"))
	return rv
}

// Replaces the layout manager for the group of text system objects that contains the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444545-replacelayoutmanager?language=objc
func (t_ TextContainer) ReplaceLayoutManager(newLayoutManager ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("replaceLayoutManager:"), objc.Ptr(newLayoutManager))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/3809914-textlayoutmanager?language=objc
func (t_ TextContainer) TextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("textLayoutManager"))
	return rv
}

// The maximum number of lines that the text container can store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc
func (t_ TextContainer) MaximumNumberOfLines() uint {
	rv := objc.Call[uint](t_, objc.Sel("maximumNumberOfLines"))
	return rv
}

// The maximum number of lines that the text container can store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc
func (t_ TextContainer) SetMaximumNumberOfLines(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setMaximumNumberOfLines:"), value)
}

// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc
func (t_ TextContainer) HeightTracksTextView() bool {
	rv := objc.Call[bool](t_, objc.Sel("heightTracksTextView"))
	return rv
}

// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc
func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setHeightTracksTextView:"), value)
}

// The text container’s text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc
func (t_ TextContainer) TextView() TextView {
	rv := objc.Call[TextView](t_, objc.Sel("textView"))
	return rv
}

// The text container’s text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc
func (t_ TextContainer) SetTextView(value ITextView) {
	objc.Call[objc.Void](t_, objc.Sel("setTextView:"), objc.Ptr(value))
}

// A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view's coordinate system axes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444525-simplerectangulartextcontainer?language=objc
func (t_ TextContainer) IsSimpleRectangularTextContainer() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSimpleRectangularTextContainer"))
	return rv
}

// The text container’s layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc
func (t_ TextContainer) LayoutManager() LayoutManager {
	rv := objc.Call[LayoutManager](t_, objc.Sel("layoutManager"))
	return rv
}

// The text container’s layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc
func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutManager:"), objc.Ptr(value))
}

// The value for the text inset within line fragment rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc
func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := objc.Call[float64](t_, objc.Sel("lineFragmentPadding"))
	return rv
}

// The value for the text inset within line fragment rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc
func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setLineFragmentPadding:"), value)
}

// An array of path objects that represents the regions where text doesn’t display in the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc
func (t_ TextContainer) ExclusionPaths() []BezierPath {
	rv := objc.Call[[]BezierPath](t_, objc.Sel("exclusionPaths"))
	return rv
}

// An array of path objects that represents the regions where text doesn’t display in the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc
func (t_ TextContainer) SetExclusionPaths(value []IBezierPath) {
	objc.Call[objc.Void](t_, objc.Sel("setExclusionPaths:"), value)
}

// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc
func (t_ TextContainer) WidthTracksTextView() bool {
	rv := objc.Call[bool](t_, objc.Sel("widthTracksTextView"))
	return rv
}

// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc
func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setWidthTracksTextView:"), value)
}

// The size of the text container’s bounding rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc
func (t_ TextContainer) Size() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](t_, objc.Sel("size"))
	return rv
}

// The size of the text container’s bounding rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc
func (t_ TextContainer) SetSize(value coregraphics.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setSize:"), value)
}
