// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TokenFieldCell] class.
var TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}

type _TokenFieldCellClass struct {
	objc.Class
}

// An interface definition for the [TokenFieldCell] class.
type ITokenFieldCell interface {
	ITextFieldCell
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
	Delegate() TokenFieldCellDelegateWrapper
	SetDelegate(value PTokenFieldCellDelegate)
	SetDelegateObject(valueObject objc.IObject)
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
}

// A text field cell subclass that enables tokenized editing of an array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell?language=objc
type TokenFieldCell struct {
	TextFieldCell
}

func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.Call[TokenFieldCell](tc, objc.Sel("alloc"))
	return rv
}

func TokenFieldCell_Alloc() TokenFieldCell {
	return TokenFieldCellClass.Alloc()
}

func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.Call[TokenFieldCell](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.Call[TokenFieldCell](t_, objc.Sel("init"))
	return rv
}

func (t_ TokenFieldCell) InitTextCell(string_ string) TokenFieldCell {
	rv := objc.Call[TokenFieldCell](t_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func NewTokenFieldCellTextCell(string_ string) TokenFieldCell {
	instance := TokenFieldCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

func (t_ TokenFieldCell) InitImageCell(image IImage) TokenFieldCell {
	rv := objc.Call[TokenFieldCell](t_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewTokenFieldCellImageCell(image IImage) TokenFieldCell {
	instance := TokenFieldCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

// The receiver’s tokenizing character set to a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523822-tokenizingcharacterset?language=objc
func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](t_, objc.Sel("tokenizingCharacterSet"))
	return rv
}

// The receiver’s tokenizing character set to a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523822-tokenizingcharacterset?language=objc
func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	objc.Call[objc.Void](t_, objc.Sel("setTokenizingCharacterSet:"), objc.Ptr(value))
}

// Returns the default completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523817-defaultcompletiondelay?language=objc
func (tc _TokenFieldCellClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](tc, objc.Sel("defaultCompletionDelay"))
	return rv
}

// Returns the default completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523817-defaultcompletiondelay?language=objc
func TokenFieldCell_DefaultCompletionDelay() foundation.TimeInterval {
	return TokenFieldCellClass.DefaultCompletionDelay()
}

// The receiver’s completion delay to a given delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523806-completiondelay?language=objc
func (t_ TokenFieldCell) CompletionDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](t_, objc.Sel("completionDelay"))
	return rv
}

// The receiver’s completion delay to a given delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523806-completiondelay?language=objc
func (t_ TokenFieldCell) SetCompletionDelay(value foundation.TimeInterval) {
	objc.Call[objc.Void](t_, objc.Sel("setCompletionDelay:"), value)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523813-delegate?language=objc
func (t_ TokenFieldCell) Delegate() TokenFieldCellDelegateWrapper {
	rv := objc.Call[TokenFieldCellDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523813-delegate?language=objc
func (t_ TokenFieldCell) SetDelegate(value PTokenFieldCellDelegate) {
	po0 := objc.WrapAsProtocol("NSTokenFieldCellDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523813-delegate?language=objc
func (t_ TokenFieldCell) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Returns the default tokenizing character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523808-defaulttokenizingcharacterset?language=objc
func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](tc, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}

// Returns the default tokenizing character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523808-defaulttokenizingcharacterset?language=objc
func TokenFieldCell_DefaultTokenizingCharacterSet() foundation.CharacterSet {
	return TokenFieldCellClass.DefaultTokenizingCharacterSet()
}

// The token style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523827-tokenstyle?language=objc
func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := objc.Call[TokenStyle](t_, objc.Sel("tokenStyle"))
	return rv
}

// The token style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcell/1523827-tokenstyle?language=objc
func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setTokenStyle:"), value)
}
