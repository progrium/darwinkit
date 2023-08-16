// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompileOptions] class.
var CompileOptionsClass = _CompileOptionsClass{objc.GetClass("MTLCompileOptions")}

type _CompileOptionsClass struct {
	objc.Class
}

// An interface definition for the [CompileOptions] class.
type ICompileOptions interface {
	objc.IObject
	InstallName() string
	SetInstallName(value string)
	LibraryType() LibraryType
	SetLibraryType(value LibraryType)
	PreprocessorMacros() map[string]objc.Object
	SetPreprocessorMacros(value map[string]objc.IObject)
	FastMathEnabled() bool
	SetFastMathEnabled(value bool)
	PreserveInvariance() bool
	SetPreserveInvariance(value bool)
	LanguageVersion() LanguageVersion
	SetLanguageVersion(value LanguageVersion)
	Libraries() []DynamicLibraryWrapper
	SetLibraries(value []PDynamicLibrary)
}

// Compilation settings for a Metal shader library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions?language=objc
type CompileOptions struct {
	objc.Object
}

func CompileOptionsFrom(ptr unsafe.Pointer) CompileOptions {
	return CompileOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CompileOptionsClass) Alloc() CompileOptions {
	rv := objc.Call[CompileOptions](cc, objc.Sel("alloc"))
	return rv
}

func CompileOptions_Alloc() CompileOptions {
	return CompileOptionsClass.Alloc()
}

func (cc _CompileOptionsClass) New() CompileOptions {
	rv := objc.Call[CompileOptions](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompileOptions() CompileOptions {
	return CompileOptionsClass.New()
}

func (c_ CompileOptions) Init() CompileOptions {
	rv := objc.Call[CompileOptions](c_, objc.Sel("init"))
	return rv
}

// For a dynamic library, the name to use when installing the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554034-installname?language=objc
func (c_ CompileOptions) InstallName() string {
	rv := objc.Call[string](c_, objc.Sel("installName"))
	return rv
}

// For a dynamic library, the name to use when installing the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554034-installname?language=objc
func (c_ CompileOptions) SetInstallName(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setInstallName:"), value)
}

// The kind of library to create. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554036-librarytype?language=objc
func (c_ CompileOptions) LibraryType() LibraryType {
	rv := objc.Call[LibraryType](c_, objc.Sel("libraryType"))
	return rv
}

// The kind of library to create. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554036-librarytype?language=objc
func (c_ CompileOptions) SetLibraryType(value LibraryType) {
	objc.Call[objc.Void](c_, objc.Sel("setLibraryType:"), value)
}

// A list of preprocessor macros to apply when compiling the library source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1516172-preprocessormacros?language=objc
func (c_ CompileOptions) PreprocessorMacros() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("preprocessorMacros"))
	return rv
}

// A list of preprocessor macros to apply when compiling the library source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1516172-preprocessormacros?language=objc
func (c_ CompileOptions) SetPreprocessorMacros(value map[string]objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setPreprocessorMacros:"), value)
}

// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1515914-fastmathenabled?language=objc
func (c_ CompileOptions) FastMathEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("fastMathEnabled"))
	return rv
}

// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1515914-fastmathenabled?language=objc
func (c_ CompileOptions) SetFastMathEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setFastMathEnabled:"), value)
}

// A Boolean value that indicates whether the compiler should compile vertex shaders conservatively to generate consistent position calculations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3564462-preserveinvariance?language=objc
func (c_ CompileOptions) PreserveInvariance() bool {
	rv := objc.Call[bool](c_, objc.Sel("preserveInvariance"))
	return rv
}

// A Boolean value that indicates whether the compiler should compile vertex shaders conservatively to generate consistent position calculations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3564462-preserveinvariance?language=objc
func (c_ CompileOptions) SetPreserveInvariance(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setPreserveInvariance:"), value)
}

// The language version used to interpret the library source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1515494-languageversion?language=objc
func (c_ CompileOptions) LanguageVersion() LanguageVersion {
	rv := objc.Call[LanguageVersion](c_, objc.Sel("languageVersion"))
	return rv
}

// The language version used to interpret the library source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/1515494-languageversion?language=objc
func (c_ CompileOptions) SetLanguageVersion(value LanguageVersion) {
	objc.Call[objc.Void](c_, objc.Sel("setLanguageVersion:"), value)
}

// An array of dynamic libraries the Metal compiler links against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554035-libraries?language=objc
func (c_ CompileOptions) Libraries() []DynamicLibraryWrapper {
	rv := objc.Call[[]DynamicLibraryWrapper](c_, objc.Sel("libraries"))
	return rv
}

// An array of dynamic libraries the Metal compiler links against. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/3554035-libraries?language=objc
func (c_ CompileOptions) SetLibraries(value []PDynamicLibrary) {
	objc.Call[objc.Void](c_, objc.Sel("setLibraries:"), value)
}
