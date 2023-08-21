// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserScript] class.
var UserScriptClass = _UserScriptClass{objc.GetClass("WKUserScript")}

type _UserScriptClass struct {
	objc.Class
}

// An interface definition for the [UserScript] class.
type IUserScript interface {
	objc.IObject
	Source() string
	IsForMainFrameOnly() bool
	InjectionTime() UserScriptInjectionTime
}

// A script that the web view injects into a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript?language=objc
type UserScript struct {
	objc.Object
}

func UserScriptFrom(ptr unsafe.Pointer) UserScript {
	return UserScript{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserScript) InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	rv := objc.Call[UserScript](u_, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:"), source, injectionTime, forMainFrameOnly)
	return rv
}

// Creates a user script object that contains the specified source code and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/1537750-initwithsource?language=objc
func NewUserScriptWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	instance := UserScriptClass.Alloc().InitWithSourceInjectionTimeForMainFrameOnly(source, injectionTime, forMainFrameOnly)
	instance.Autorelease()
	return instance
}

func (uc _UserScriptClass) Alloc() UserScript {
	rv := objc.Call[UserScript](uc, objc.Sel("alloc"))
	return rv
}

func UserScript_Alloc() UserScript {
	return UserScriptClass.Alloc()
}

func (uc _UserScriptClass) New() UserScript {
	rv := objc.Call[UserScript](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserScript() UserScript {
	return UserScriptClass.New()
}

func (u_ UserScript) Init() UserScript {
	rv := objc.Call[UserScript](u_, objc.Sel("init"))
	return rv
}

// The script’s source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/1537787-source?language=objc
func (u_ UserScript) Source() string {
	rv := objc.Call[string](u_, objc.Sel("source"))
	return rv
}

// A Boolean value that indicates whether to inject the script into the main frame or all frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/1537856-formainframeonly?language=objc
func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := objc.Call[bool](u_, objc.Sel("isForMainFrameOnly"))
	return rv
}

// The time at which to inject the script into the webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/1536492-injectiontime?language=objc
func (u_ UserScript) InjectionTime() UserScriptInjectionTime {
	rv := objc.Call[UserScriptInjectionTime](u_, objc.Sel("injectionTime"))
	return rv
}
