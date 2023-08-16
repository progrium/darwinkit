// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A collection of Metal shader functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary?language=objc
type PLibrary interface {
	// optional
	NewFunctionWithDescriptorCompletionHandler(descriptor FunctionDescriptor, completionHandler func(function FunctionWrapper, error foundation.Error))
	HasNewFunctionWithDescriptorCompletionHandler() bool

	// optional
	NewIntersectionFunctionWithDescriptorError(descriptor IntersectionFunctionDescriptor, error foundation.Error) PFunction
	HasNewIntersectionFunctionWithDescriptorError() bool

	// optional
	NewFunctionWithNameConstantValuesError(name string, constantValues FunctionConstantValues, error foundation.Error) PFunction
	HasNewFunctionWithNameConstantValuesError() bool

	// optional
	InstallName() string
	HasInstallName() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	Type() LibraryType
	HasType() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	FunctionNames() []string
	HasFunctionNames() bool
}

// A concrete type wrapper for the [PLibrary] protocol.
type LibraryWrapper struct {
	objc.Object
}

func (l_ LibraryWrapper) HasNewFunctionWithDescriptorCompletionHandler() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithDescriptor:completionHandler:"))
}

// Asynchronously creates an object representing a shader function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554040-newfunctionwithdescriptor?language=objc
func (l_ LibraryWrapper) NewFunctionWithDescriptorCompletionHandler(descriptor IFunctionDescriptor, completionHandler func(function FunctionWrapper, error foundation.Error)) {
	objc.Call[objc.Void](l_, objc.Sel("newFunctionWithDescriptor:completionHandler:"), objc.Ptr(descriptor), completionHandler)
}

func (l_ LibraryWrapper) HasNewIntersectionFunctionWithDescriptorError() bool {
	return l_.RespondsToSelector(objc.Sel("newIntersectionFunctionWithDescriptor:error:"))
}

// Synchronously creates an object representing a ray-tracing intersection function, using the specified descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3578135-newintersectionfunctionwithdescr?language=objc
func (l_ LibraryWrapper) NewIntersectionFunctionWithDescriptorError(descriptor IIntersectionFunctionDescriptor, error foundation.IError) FunctionWrapper {
	rv := objc.Call[FunctionWrapper](l_, objc.Sel("newIntersectionFunctionWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (l_ LibraryWrapper) HasNewFunctionWithNameConstantValuesError() bool {
	return l_.RespondsToSelector(objc.Sel("newFunctionWithName:constantValues:error:"))
}

// Synchronously creates a specialized shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1640020-newfunctionwithname?language=objc
func (l_ LibraryWrapper) NewFunctionWithNameConstantValuesError(name string, constantValues IFunctionConstantValues, error foundation.IError) FunctionWrapper {
	rv := objc.Call[FunctionWrapper](l_, objc.Sel("newFunctionWithName:constantValues:error:"), name, objc.Ptr(constantValues), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (l_ LibraryWrapper) HasInstallName() bool {
	return l_.RespondsToSelector(objc.Sel("installName"))
}

// The installation name for a dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554039-installname?language=objc
func (l_ LibraryWrapper) InstallName() string {
	rv := objc.Call[string](l_, objc.Sel("installName"))
	return rv
}

func (l_ LibraryWrapper) HasDevice() bool {
	return l_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1515661-device?language=objc
func (l_ LibraryWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](l_, objc.Sel("device"))
	return rv
}

func (l_ LibraryWrapper) HasType() bool {
	return l_.RespondsToSelector(objc.Sel("type"))
}

// The library’s basic type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/3554042-type?language=objc
func (l_ LibraryWrapper) Type() LibraryType {
	rv := objc.Call[LibraryType](l_, objc.Sel("type"))
	return rv
}

func (l_ LibraryWrapper) HasSetLabel() bool {
	return l_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1516253-label?language=objc
func (l_ LibraryWrapper) SetLabel(value string) {
	objc.Call[objc.Void](l_, objc.Sel("setLabel:"), value)
}

func (l_ LibraryWrapper) HasLabel() bool {
	return l_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1516253-label?language=objc
func (l_ LibraryWrapper) Label() string {
	rv := objc.Call[string](l_, objc.Sel("label"))
	return rv
}

func (l_ LibraryWrapper) HasFunctionNames() bool {
	return l_.RespondsToSelector(objc.Sel("functionNames"))
}

// The names of all public functions in the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrary/1515651-functionnames?language=objc
func (l_ LibraryWrapper) FunctionNames() []string {
	rv := objc.Call[[]string](l_, objc.Sel("functionNames"))
	return rv
}
