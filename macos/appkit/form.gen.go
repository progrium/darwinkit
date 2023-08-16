// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Form] class.
var FormClass = _FormClass{objc.GetClass("NSForm")}

type _FormClass struct {
	objc.Class
}

// An interface definition for the [Form] class.
type IForm interface {
	IMatrix
}

// An NSForm object is a vertical matrix of NSFormCell objects to implement the fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsform?language=objc
type Form struct {
	Matrix
}

func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		Matrix: MatrixFrom(ptr),
	}
}

func (fc _FormClass) Alloc() Form {
	rv := objc.Call[Form](fc, objc.Sel("alloc"))
	return rv
}

func Form_Alloc() Form {
	return FormClass.Alloc()
}

func (fc _FormClass) New() Form {
	rv := objc.Call[Form](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewForm() Form {
	return FormClass.New()
}

func (f_ Form) Init() Form {
	rv := objc.Call[Form](f_, objc.Sel("init"))
	return rv
}

func (f_ Form) InitWithFrame(frameRect foundation.Rect) Form {
	rv := objc.Call[Form](f_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a newly allocated matrix with the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436428-initwithframe?language=objc
func Form_InitWithFrame(frameRect foundation.Rect) Form {
	return FormClass.Alloc().InitWithFrame(frameRect)
}
