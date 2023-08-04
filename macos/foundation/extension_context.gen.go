// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}

type _ExtensionContextClass struct {
	objc.Class
}

type IExtensionContext interface {
	objc.IObject
	CancelRequestWithError(error IError)
	CompleteRequestReturningItemsCompletionHandler(items []objc.IObject, completionHandler func(expired bool))
	OpenURLCompletionHandler(URL IURL, completionHandler func(success bool))
	InputItems() []objc.Object
}

type ExtensionContext struct {
	objc.Object
}

func MakeExtensionContext(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](ec, objc.GetSelector("alloc"))
	return rv
}

func ExtensionContext_Alloc() ExtensionContext {
	return ExtensionContextClass.Alloc()
}

func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionContext() ExtensionContext {
	return ExtensionContextClass.New()
}

func ExtensionContext_New() ExtensionContext {
	return ExtensionContextClass.New()
}

func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](e_, objc.GetSelector("init"))
	return rv
}

func ExtensionContext_Init() ExtensionContext {
	return ExtensionContextClass.Alloc().Init()
}

func (e_ ExtensionContext) CancelRequestWithError(error IError) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("cancelRequestWithError:"), objc.ExtractPtr(error))
}

func (e_ ExtensionContext) CompleteRequestReturningItemsCompletionHandler(items []objc.IObject, completionHandler func(expired bool)) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("completeRequestReturningItems:completionHandler:"), items, completionHandler)
}

func (e_ ExtensionContext) OpenURLCompletionHandler(URL IURL, completionHandler func(success bool)) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("openURL:completionHandler:"), objc.ExtractPtr(URL), completionHandler)
}

func (e_ ExtensionContext) InputItems() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](e_, objc.GetSelector("inputItems"))
	return rv
}
