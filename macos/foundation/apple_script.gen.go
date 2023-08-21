// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AppleScript] class.
var AppleScriptClass = _AppleScriptClass{objc.GetClass("NSAppleScript")}

type _AppleScriptClass struct {
	objc.Class
}

// An interface definition for the [AppleScript] class.
type IAppleScript interface {
	objc.IObject
	CompileAndReturnError(errorInfo map[string]objc.IObject) bool
	ExecuteAndReturnError(errorInfo map[string]objc.IObject) AppleEventDescriptor
	ExecuteAppleEventError(event IAppleEventDescriptor, errorInfo map[string]objc.IObject) AppleEventDescriptor
	RichTextSource() AttributedString
	Source() string
	IsCompiled() bool
}

// An object that provides the ability to load, compile, and execute scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript?language=objc
type AppleScript struct {
	objc.Object
}

func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AppleScript) InitWithContentsOfURLError(url IURL, errorInfo map[string]objc.IObject) AppleScript {
	rv := objc.Call[AppleScript](a_, objc.Sel("initWithContentsOfURL:error:"), objc.Ptr(url), errorInfo)
	return rv
}

// Initializes a newly allocated script instance from the source identified by the passed URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1412508-initwithcontentsofurl?language=objc
func NewAppleScriptWithContentsOfURLError(url IURL, errorInfo map[string]objc.IObject) AppleScript {
	instance := AppleScriptClass.Alloc().InitWithContentsOfURLError(url, errorInfo)
	instance.Autorelease()
	return instance
}

func (a_ AppleScript) InitWithSource(source string) AppleScript {
	rv := objc.Call[AppleScript](a_, objc.Sel("initWithSource:"), source)
	return rv
}

// Initializes a newly allocated script instance from the passed source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1414313-initwithsource?language=objc
func NewAppleScriptWithSource(source string) AppleScript {
	instance := AppleScriptClass.Alloc().InitWithSource(source)
	instance.Autorelease()
	return instance
}

func (ac _AppleScriptClass) Alloc() AppleScript {
	rv := objc.Call[AppleScript](ac, objc.Sel("alloc"))
	return rv
}

func AppleScript_Alloc() AppleScript {
	return AppleScriptClass.Alloc()
}

func (ac _AppleScriptClass) New() AppleScript {
	rv := objc.Call[AppleScript](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAppleScript() AppleScript {
	return AppleScriptClass.New()
}

func (a_ AppleScript) Init() AppleScript {
	rv := objc.Call[AppleScript](a_, objc.Sel("init"))
	return rv
}

// Compiles the receiver, if it is not already compiled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1407582-compileandreturnerror?language=objc
func (a_ AppleScript) CompileAndReturnError(errorInfo map[string]objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("compileAndReturnError:"), errorInfo)
	return rv
}

// Executes the receiver, compiling it first if it is not already compiled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1410034-executeandreturnerror?language=objc
func (a_ AppleScript) ExecuteAndReturnError(errorInfo map[string]objc.IObject) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("executeAndReturnError:"), errorInfo)
	return rv
}

// Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1410807-executeappleevent?language=objc
func (a_ AppleScript) ExecuteAppleEventError(event IAppleEventDescriptor, errorInfo map[string]objc.IObject) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("executeAppleEvent:error:"), objc.Ptr(event), errorInfo)
	return rv
}

// Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1387495-richtextsource?language=objc
func (a_ AppleScript) RichTextSource() AttributedString {
	rv := objc.Call[AttributedString](a_, objc.Sel("richTextSource"))
	return rv
}

// The script source for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1408453-source?language=objc
func (a_ AppleScript) Source() string {
	rv := objc.Call[string](a_, objc.Sel("source"))
	return rv
}

// A Boolean value that indicates whether the receiver's script has been compiled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/1410407-compiled?language=objc
func (a_ AppleScript) IsCompiled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isCompiled"))
	return rv
}
