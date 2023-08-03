// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var UserContentControllerClass = _UserContentControllerClass{objc.GetClass("WKUserContentController")}

type _UserContentControllerClass struct {
	objc.Class
}

type IUserContentController interface {
	objc.IObject
	AddUserScript(userScript IUserScript)
	RemoveAllUserScripts()
	AddScriptMessageHandlerName(scriptMessageHandler IScriptMessageHandler, name string)
	AddScriptMessageHandler0Name(scriptMessageHandler objc.IObject, name string)
	AddScriptMessageHandlerContentWorldName(scriptMessageHandler IScriptMessageHandler, world IContentWorld, name string)
	AddScriptMessageHandler0ContentWorldName(scriptMessageHandler objc.IObject, world IContentWorld, name string)
	AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply IScriptMessageHandlerWithReply, contentWorld IContentWorld, name string)
	AddScriptMessageHandlerWithReply0ContentWorldName(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string)
	RemoveScriptMessageHandlerForName(name string)
	RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IContentWorld)
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld)
	RemoveAllScriptMessageHandlers()
	AddContentRuleList(contentRuleList IContentRuleList)
	RemoveContentRuleList(contentRuleList IContentRuleList)
	RemoveAllContentRuleLists()
	UserScripts() []UserScript
}

type UserContentController struct {
	objc.Object
}

func MakeUserContentController(ptr unsafe.Pointer) UserContentController {
	return UserContentController{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _UserContentControllerClass) Alloc() UserContentController {
	rv := objc.CallMethod[UserContentController](uc, objc.GetSelector("alloc"))
	return rv
}

func UserContentController_Alloc() UserContentController {
	return UserContentControllerClass.Alloc()
}

func (uc _UserContentControllerClass) New() UserContentController {
	rv := objc.CallMethod[UserContentController](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserContentController() UserContentController {
	return UserContentControllerClass.New()
}

func UserContentController_New() UserContentController {
	return UserContentControllerClass.New()
}

func (u_ UserContentController) Init() UserContentController {
	rv := objc.CallMethod[UserContentController](u_, objc.GetSelector("init"))
	return rv
}

func UserContentController_Init() UserContentController {
	return UserContentControllerClass.Alloc().Init()
}

func (u_ UserContentController) AddUserScript(userScript IUserScript) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addUserScript:"), objc.ExtractPtr(userScript))
}

func (u_ UserContentController) RemoveAllUserScripts() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllUserScripts"))
}

func (u_ UserContentController) AddScriptMessageHandlerName(scriptMessageHandler IScriptMessageHandler, name string) {
	po := objc.WrapAsProtocol("WKScriptMessageHandler", scriptMessageHandler)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:name:"), po, name)
}

func (u_ UserContentController) AddScriptMessageHandler0Name(scriptMessageHandler objc.IObject, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:name:"), objc.ExtractPtr(scriptMessageHandler), name)
}

func (u_ UserContentController) AddScriptMessageHandlerContentWorldName(scriptMessageHandler IScriptMessageHandler, world IContentWorld, name string) {
	po := objc.WrapAsProtocol("WKScriptMessageHandler", scriptMessageHandler)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:contentWorld:name:"), po, objc.ExtractPtr(world), name)
}

func (u_ UserContentController) AddScriptMessageHandler0ContentWorldName(scriptMessageHandler objc.IObject, world IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:contentWorld:name:"), objc.ExtractPtr(scriptMessageHandler), objc.ExtractPtr(world), name)
}

func (u_ UserContentController) AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply IScriptMessageHandlerWithReply, contentWorld IContentWorld, name string) {
	po := objc.WrapAsProtocol("WKScriptMessageHandlerWithReply", scriptMessageHandlerWithReply)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandlerWithReply:contentWorld:name:"), po, objc.ExtractPtr(contentWorld), name)
}

func (u_ UserContentController) AddScriptMessageHandlerWithReply0ContentWorldName(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandlerWithReply:contentWorld:name:"), objc.ExtractPtr(scriptMessageHandlerWithReply), objc.ExtractPtr(contentWorld), name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForName(name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeScriptMessageHandlerForName:"), name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeScriptMessageHandlerForName:contentWorld:"), name, objc.ExtractPtr(contentWorld))
}

func (u_ UserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllScriptMessageHandlersFromContentWorld:"), objc.ExtractPtr(contentWorld))
}

func (u_ UserContentController) RemoveAllScriptMessageHandlers() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllScriptMessageHandlers"))
}

func (u_ UserContentController) AddContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addContentRuleList:"), objc.ExtractPtr(contentRuleList))
}

func (u_ UserContentController) RemoveContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeContentRuleList:"), objc.ExtractPtr(contentRuleList))
}

func (u_ UserContentController) RemoveAllContentRuleLists() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllContentRuleLists"))
}

func (u_ UserContentController) UserScripts() []UserScript {
	rv := objc.CallMethod[[]UserScript](u_, objc.GetSelector("userScripts"))
	// TODO: convert slice items...
	return rv
}
