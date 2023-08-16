// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionDescriptor] class.
var FunctionDescriptorClass = _FunctionDescriptorClass{objc.GetClass("MTLFunctionDescriptor")}

type _FunctionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [FunctionDescriptor] class.
type IFunctionDescriptor interface {
	objc.IObject
	ConstantValues() FunctionConstantValues
	SetConstantValues(value IFunctionConstantValues)
	Options() FunctionOptions
	SetOptions(value FunctionOptions)
	Name() string
	SetName(value string)
	SpecializedName() string
	SetSpecializedName(value string)
	BinaryArchives() []BinaryArchiveWrapper
	SetBinaryArchives(value []PBinaryArchive)
}

// A description of a function object to create. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor?language=objc
type FunctionDescriptor struct {
	objc.Object
}

func FunctionDescriptorFrom(ptr unsafe.Pointer) FunctionDescriptor {
	return FunctionDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FunctionDescriptorClass) Alloc() FunctionDescriptor {
	rv := objc.Call[FunctionDescriptor](fc, objc.Sel("alloc"))
	return rv
}

func FunctionDescriptor_Alloc() FunctionDescriptor {
	return FunctionDescriptorClass.Alloc()
}

func (fc _FunctionDescriptorClass) New() FunctionDescriptor {
	rv := objc.Call[FunctionDescriptor](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionDescriptor() FunctionDescriptor {
	return FunctionDescriptorClass.New()
}

func (f_ FunctionDescriptor) Init() FunctionDescriptor {
	rv := objc.Call[FunctionDescriptor](f_, objc.Sel("init"))
	return rv
}

// Creates a default function descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3578278-functiondescriptor?language=objc
func (fc _FunctionDescriptorClass) FunctionDescriptor() FunctionDescriptor {
	rv := objc.Call[FunctionDescriptor](fc, objc.Sel("functionDescriptor"))
	return rv
}

// Creates a default function descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3578278-functiondescriptor?language=objc
func FunctionDescriptor_FunctionDescriptor() FunctionDescriptor {
	return FunctionDescriptorClass.FunctionDescriptor()
}

// The set of constant values assigned to the function constants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553993-constantvalues?language=objc
func (f_ FunctionDescriptor) ConstantValues() FunctionConstantValues {
	rv := objc.Call[FunctionConstantValues](f_, objc.Sel("constantValues"))
	return rv
}

// The set of constant values assigned to the function constants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553993-constantvalues?language=objc
func (f_ FunctionDescriptor) SetConstantValues(value IFunctionConstantValues) {
	objc.Call[objc.Void](f_, objc.Sel("setConstantValues:"), objc.Ptr(value))
}

// Flags specifying how Metal should create the new function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553996-options?language=objc
func (f_ FunctionDescriptor) Options() FunctionOptions {
	rv := objc.Call[FunctionOptions](f_, objc.Sel("options"))
	return rv
}

// Flags specifying how Metal should create the new function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553996-options?language=objc
func (f_ FunctionDescriptor) SetOptions(value FunctionOptions) {
	objc.Call[objc.Void](f_, objc.Sel("setOptions:"), value)
}

// The name of the function to fetch from the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553995-name?language=objc
func (f_ FunctionDescriptor) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// The name of the function to fetch from the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553995-name?language=objc
func (f_ FunctionDescriptor) SetName(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setName:"), value)
}

// A new name for the created function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553997-specializedname?language=objc
func (f_ FunctionDescriptor) SpecializedName() string {
	rv := objc.Call[string](f_, objc.Sel("specializedName"))
	return rv
}

// A new name for the created function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553997-specializedname?language=objc
func (f_ FunctionDescriptor) SetSpecializedName(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setSpecializedName:"), value)
}

// The binary archives to search for a previously-compiled version of this function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553992-binaryarchives?language=objc
func (f_ FunctionDescriptor) BinaryArchives() []BinaryArchiveWrapper {
	rv := objc.Call[[]BinaryArchiveWrapper](f_, objc.Sel("binaryArchives"))
	return rv
}

// The binary archives to search for a previously-compiled version of this function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/3553992-binaryarchives?language=objc
func (f_ FunctionDescriptor) SetBinaryArchives(value []PBinaryArchive) {
	objc.Call[objc.Void](f_, objc.Sel("setBinaryArchives:"), value)
}
