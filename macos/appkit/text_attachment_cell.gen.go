// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextAttachmentCell] class.
var TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}

type _TextAttachmentCellClass struct {
	objc.Class
}

// An interface definition for the [TextAttachmentCell] class.
type ITextAttachmentCell interface {
	ICell
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent IEvent, cellFrame foundation.Rect, controlView IView, charIndex uint) bool
	WantsToTrackMouse() bool
	DrawWithFrameInViewCharacterIndex(cellFrame foundation.Rect, controlView IView, charIndex uint)
	CellBaselineOffset() foundation.Point
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag foundation.Rect, position foundation.Point, charIndex uint) foundation.Rect
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent IEvent, cellFrame foundation.Rect, controlView IView, charIndex uint, flag bool) bool
	Attachment() TextAttachment
	SetAttachment(value ITextAttachment)
}

// An object that implements the functionality of the text attachment cell protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentcell?language=objc
type TextAttachmentCell struct {
	Cell
}

func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: CellFrom(ptr),
	}
}

func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.Call[TextAttachmentCell](tc, objc.Sel("alloc"))
	return rv
}

func TextAttachmentCell_Alloc() TextAttachmentCell {
	return TextAttachmentCellClass.Alloc()
}

func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.Call[TextAttachmentCell](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachmentCell() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}

func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.Call[TextAttachmentCell](t_, objc.Sel("init"))
	return rv
}

func (t_ TextAttachmentCell) InitImageCell(image IImage) TextAttachmentCell {
	rv := objc.Call[TextAttachmentCell](t_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func TextAttachmentCell_InitImageCell(image IImage) TextAttachmentCell {
	return TextAttachmentCellClass.Alloc().InitImageCell(image)
}

func (t_ TextAttachmentCell) InitTextCell(string_ string) TextAttachmentCell {
	rv := objc.Call[TextAttachmentCell](t_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func TextAttachmentCell_InitTextCell(string_ string) TextAttachmentCell {
	return TextAttachmentCellClass.Alloc().InitTextCell(string_)
}

// Allows an attachment to specify the events for which it tracks the mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508399-wantstotrackmouseforevent?language=objc
func (t_ TextAttachmentCell) WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent IEvent, cellFrame foundation.Rect, controlView IView, charIndex uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("wantsToTrackMouseForEvent:inRect:ofView:atCharacterIndex:"), objc.Ptr(theEvent), cellFrame, objc.Ptr(controlView), charIndex)
	return rv
}

// Returns a Boolean value that indicates whether the attachment handles mouse events occurring over its image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508415-wantstotrackmouse?language=objc
func (t_ TextAttachmentCell) WantsToTrackMouse() bool {
	rv := objc.Call[bool](t_, objc.Sel("wantsToTrackMouse"))
	return rv
}

// Draws the cell's image at the specified index point in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508412-drawwithframe?language=objc
func (t_ TextAttachmentCell) DrawWithFrameInViewCharacterIndex(cellFrame foundation.Rect, controlView IView, charIndex uint) {
	objc.Call[objc.Void](t_, objc.Sel("drawWithFrame:inView:characterIndex:"), cellFrame, objc.Ptr(controlView), charIndex)
}

// Returns the text position where you draw the attachment cell’s image, relative to the current point established in the glyph layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508420-cellbaselineoffset?language=objc
func (t_ TextAttachmentCell) CellBaselineOffset() foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("cellBaselineOffset"))
	return rv
}

// Returns the frame of the cell to draw at the specified position in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508402-cellframefortextcontainer?language=objc
func (t_ TextAttachmentCell) CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag foundation.Rect, position foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("cellFrameForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), objc.Ptr(textContainer), lineFrag, position, charIndex)
	return rv
}

// Handles a mouse-down event on the image at the specified character position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508380-trackmouse?language=objc
func (t_ TextAttachmentCell) TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent IEvent, cellFrame foundation.Rect, controlView IView, charIndex uint, flag bool) bool {
	rv := objc.Call[bool](t_, objc.Sel("trackMouse:inRect:ofView:atCharacterIndex:untilMouseUp:"), objc.Ptr(theEvent), cellFrame, objc.Ptr(controlView), charIndex, flag)
	return rv
}

// Returns the text attachment object that owns the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508396-attachment?language=objc
func (t_ TextAttachmentCell) Attachment() TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("attachment"))
	return rv
}

// Returns the text attachment object that owns the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1508388-nstextattachmentcell/1508396-attachment?language=objc
func (t_ TextAttachmentCell) SetAttachment(value ITextAttachment) {
	objc.Call[objc.Void](t_, objc.Sel("setAttachment:"), objc.Ptr(value))
}
