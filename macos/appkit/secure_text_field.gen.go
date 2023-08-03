// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}

type _SecureTextFieldClass struct {
	objc.Class
}

type ISecureTextField interface {
	ITextField
}

type SecureTextField struct {
	TextField
}

func MakeSecureTextField(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: MakeTextField(ptr),
	}
}

func (sc _SecureTextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func SecureTextField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	return SecureTextFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (sc _SecureTextFieldClass) LabelWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("labelWithString:"), stringValue)
	return rv
}

func SecureTextField_LabelWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.LabelWithString(stringValue)
}

func (sc _SecureTextFieldClass) TextFieldWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("textFieldWithString:"), stringValue)
	return rv
}

func SecureTextField_TextFieldWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.TextFieldWithString(stringValue)
}

func (sc _SecureTextFieldClass) WrappingLabelWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("wrappingLabelWithString:"), stringValue)
	return rv
}

func SecureTextField_WrappingLabelWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.WrappingLabelWithString(stringValue)
}

func (s_ SecureTextField) InitWithFrame(frameRect foundation.Rect) SecureTextField {
	rv := objc.CallMethod[SecureTextField](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func SecureTextField_InitWithFrame(frameRect foundation.Rect) SecureTextField {
	return SecureTextFieldClass.Alloc().InitWithFrame(frameRect)
}

func (s_ SecureTextField) Init() SecureTextField {
	rv := objc.CallMethod[SecureTextField](s_, objc.GetSelector("init"))
	return rv
}

func SecureTextField_Init() SecureTextField {
	return SecureTextFieldClass.Alloc().Init()
}

func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("alloc"))
	return rv
}

func SecureTextField_Alloc() SecureTextField {
	return SecureTextFieldClass.Alloc()
}

func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSecureTextField() SecureTextField {
	return SecureTextFieldClass.New()
}

func SecureTextField_New() SecureTextField {
	return SecureTextFieldClass.New()
}
