// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserContentController] class.
var UserContentControllerClass = _UserContentControllerClass{objc.GetClass("WKUserContentController")}

type _UserContentControllerClass struct {
	objc.Class
}

// An interface definition for the [UserContentController] class.
type IUserContentController interface {
	objc.IObject
	RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IContentWorld)
	AddScriptMessageHandlerName(scriptMessageHandler PScriptMessageHandler, name string)
	AddScriptMessageHandlerObjectName(scriptMessageHandlerObject objc.IObject, name string)
	RemoveContentRuleList(contentRuleList IContentRuleList)
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld)
	RemoveAllContentRuleLists()
	RemoveAllScriptMessageHandlers()
	AddContentRuleList(contentRuleList IContentRuleList)
	RemoveAllUserScripts()
	AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply PScriptMessageHandlerWithReply, contentWorld IContentWorld, name string)
	AddScriptMessageHandlerWithReplyObjectContentWorldName(scriptMessageHandlerWithReplyObject objc.IObject, contentWorld IContentWorld, name string)
	AddUserScript(userScript IUserScript)
	UserScripts() []UserScript
}

// An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller?language=objc
type UserContentController struct {
	objc.Object
}

func UserContentControllerFrom(ptr unsafe.Pointer) UserContentController {
	return UserContentController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UserContentControllerClass) Alloc() UserContentController {
	rv := objc.Call[UserContentController](uc, objc.Sel("alloc"))
	return rv
}

func UserContentController_Alloc() UserContentController {
	return UserContentControllerClass.Alloc()
}

func (uc _UserContentControllerClass) New() UserContentController {
	rv := objc.Call[UserContentController](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserContentController() UserContentController {
	return UserContentControllerClass.New()
}

func (u_ UserContentController) Init() UserContentController {
	rv := objc.Call[UserContentController](u_, objc.Sel("init"))
	return rv
}

// Uninstalls a custom message handler from the specified content world in your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/3585116-removescriptmessagehandlerfornam?language=objc
func (u_ UserContentController) RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IContentWorld) {
	objc.Call[objc.Void](u_, objc.Sel("removeScriptMessageHandlerForName:contentWorld:"), name, objc.Ptr(contentWorld))
}

// Installs a message handler that you can call from your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/1537172-addscriptmessagehandler?language=objc
func (u_ UserContentController) AddScriptMessageHandlerName(scriptMessageHandler PScriptMessageHandler, name string) {
	po0 := objc.WrapAsProtocol("WKScriptMessageHandler", scriptMessageHandler)
	objc.Call[objc.Void](u_, objc.Sel("addScriptMessageHandler:name:"), po0, name)
}

// Installs a message handler that you can call from your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/1537172-addscriptmessagehandler?language=objc
func (u_ UserContentController) AddScriptMessageHandlerObjectName(scriptMessageHandlerObject objc.IObject, name string) {
	objc.Call[objc.Void](u_, objc.Sel("addScriptMessageHandler:name:"), objc.Ptr(scriptMessageHandlerObject), name)
}

// Removes the specified rule list from the content controller object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/2902763-removecontentrulelist?language=objc
func (u_ UserContentController) RemoveContentRuleList(contentRuleList IContentRuleList) {
	objc.Call[objc.Void](u_, objc.Sel("removeContentRuleList:"), objc.Ptr(contentRuleList))
}

// Uninstalls all custom message handlers from the specified content world in your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/3585115-removeallscriptmessagehandlersfr?language=objc
func (u_ UserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld) {
	objc.Call[objc.Void](u_, objc.Sel("removeAllScriptMessageHandlersFromContentWorld:"), objc.Ptr(contentWorld))
}

// Removes all rules lists from the content controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/2902748-removeallcontentrulelists?language=objc
func (u_ UserContentController) RemoveAllContentRuleLists() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllContentRuleLists"))
}

// Uninstalls all custom message handlers associated with the user content controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/3585114-removeallscriptmessagehandlers?language=objc
func (u_ UserContentController) RemoveAllScriptMessageHandlers() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllScriptMessageHandlers"))
}

// Adds the specified content rule list to the content controller object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/2902756-addcontentrulelist?language=objc
func (u_ UserContentController) AddContentRuleList(contentRuleList IContentRuleList) {
	objc.Call[objc.Void](u_, objc.Sel("addContentRuleList:"), objc.Ptr(contentRuleList))
}

// Removes all user scripts from the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/1536540-removealluserscripts?language=objc
func (u_ UserContentController) RemoveAllUserScripts() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllUserScripts"))
}

// Installs a message handler that returns a reply to your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/3585113-addscriptmessagehandlerwithreply?language=objc
func (u_ UserContentController) AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply PScriptMessageHandlerWithReply, contentWorld IContentWorld, name string) {
	po0 := objc.WrapAsProtocol("WKScriptMessageHandlerWithReply", scriptMessageHandlerWithReply)
	objc.Call[objc.Void](u_, objc.Sel("addScriptMessageHandlerWithReply:contentWorld:name:"), po0, objc.Ptr(contentWorld), name)
}

// Installs a message handler that returns a reply to your JavaScript code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/3585113-addscriptmessagehandlerwithreply?language=objc
func (u_ UserContentController) AddScriptMessageHandlerWithReplyObjectContentWorldName(scriptMessageHandlerWithReplyObject objc.IObject, contentWorld IContentWorld, name string) {
	objc.Call[objc.Void](u_, objc.Sel("addScriptMessageHandlerWithReply:contentWorld:name:"), objc.Ptr(scriptMessageHandlerWithReplyObject), objc.Ptr(contentWorld), name)
}

// Injects the specified script into the webpage’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/1537448-adduserscript?language=objc
func (u_ UserContentController) AddUserScript(userScript IUserScript) {
	objc.Call[objc.Void](u_, objc.Sel("addUserScript:"), objc.Ptr(userScript))
}

// The user scripts associated with the user content controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/1538046-userscripts?language=objc
func (u_ UserContentController) UserScripts() []UserScript {
	rv := objc.Call[[]UserScript](u_, objc.Sel("userScripts"))
	return rv
}
