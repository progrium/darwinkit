// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var UserScriptClass = _UserScriptClass{objc.GetClass("WKUserScript")}

type _UserScriptClass struct {
	objc.Class
}

type IUserScript interface {
	objc.IObject
	Source() string
	InjectionTime() UserScriptInjectionTime
	IsForMainFrameOnly() bool
}

type UserScript struct {
	objc.Object
}

func MakeUserScript(ptr unsafe.Pointer) UserScript {
	return UserScript{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserScript) InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.GetSelector("initWithSource:injectionTime:forMainFrameOnly:"), source, injectionTime, forMainFrameOnly)
	return rv
}

func UserScript_InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	return UserScriptClass.Alloc().InitWithSourceInjectionTimeForMainFrameOnly(source, injectionTime, forMainFrameOnly)
}

func (u_ UserScript) InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld IContentWorld) UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.GetSelector("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), source, injectionTime, forMainFrameOnly, objc.ExtractPtr(contentWorld))
	return rv
}

func UserScript_InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld IContentWorld) UserScript {
	return UserScriptClass.Alloc().InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source, injectionTime, forMainFrameOnly, contentWorld)
}

func (uc _UserScriptClass) Alloc() UserScript {
	rv := objc.CallMethod[UserScript](uc, objc.GetSelector("alloc"))
	return rv
}

func UserScript_Alloc() UserScript {
	return UserScriptClass.Alloc()
}

func (uc _UserScriptClass) New() UserScript {
	rv := objc.CallMethod[UserScript](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserScript() UserScript {
	return UserScriptClass.New()
}

func UserScript_New() UserScript {
	return UserScriptClass.New()
}

func (u_ UserScript) Init() UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.GetSelector("init"))
	return rv
}

func UserScript_Init() UserScript {
	return UserScriptClass.Alloc().Init()
}

func (u_ UserScript) Source() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("source"))
	return rv
}

func (u_ UserScript) InjectionTime() UserScriptInjectionTime {
	rv := objc.CallMethod[UserScriptInjectionTime](u_, objc.GetSelector("injectionTime"))
	return rv
}

func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isForMainFrameOnly"))
	return rv
}
