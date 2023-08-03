// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var UserInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}

type _UserInterfaceCompressionOptionsClass struct {
	objc.Class
}

type IUserInterfaceCompressionOptions interface {
	objc.IObject
	ContainsOptions(options IUserInterfaceCompressionOptions) bool
	IntersectsOptions(options IUserInterfaceCompressionOptions) bool
	OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	IsEmpty() bool
}

type UserInterfaceCompressionOptions struct {
	objc.Object
}

func MakeUserInterfaceCompressionOptions(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("init"))
	return rv
}

func UserInterfaceCompressionOptions_Init() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.Alloc().Init()
}

func (u_ UserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("initWithCompressionOptions:"), objc.ExtractPtr(options))
	return rv
}

func UserInterfaceCompressionOptions_InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.Alloc().InitWithCompressionOptions(options)
}

func (u_ UserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func UserInterfaceCompressionOptions_InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.Alloc().InitWithIdentifier(identifier)
}

func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("alloc"))
	return rv
}

func UserInterfaceCompressionOptions_Alloc() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.Alloc()
}

func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

func UserInterfaceCompressionOptions_New() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("containsOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("intersectsOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("optionsByAddingOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("optionsByRemovingOptions:"), objc.ExtractPtr(options))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("hideImagesOption"))
	return rv
}

func UserInterfaceCompressionOptions_HideImagesOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.HideImagesOption()
}

func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("hideTextOption"))
	return rv
}

func UserInterfaceCompressionOptions_HideTextOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.HideTextOption()
}

func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("reduceMetricsOption"))
	return rv
}

func UserInterfaceCompressionOptions_ReduceMetricsOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.ReduceMetricsOption()
}

func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("breakEqualWidthsOption"))
	return rv
}

func UserInterfaceCompressionOptions_BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.BreakEqualWidthsOption()
}

func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("standardOptions"))
	return rv
}

func UserInterfaceCompressionOptions_StandardOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.StandardOptions()
}

func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isEmpty"))
	return rv
}
