// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IndexExtensionRequestHandler] class.
var IndexExtensionRequestHandlerClass = _IndexExtensionRequestHandlerClass{objc.GetClass("CSIndexExtensionRequestHandler")}

type _IndexExtensionRequestHandlerClass struct {
	objc.Class
}

// An interface definition for the [IndexExtensionRequestHandler] class.
type IIndexExtensionRequestHandler interface {
	objc.IObject
}

// An interface that implements an index-maintenance app extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csindexextensionrequesthandler?language=objc
type IndexExtensionRequestHandler struct {
	objc.Object
}

func IndexExtensionRequestHandlerFrom(ptr unsafe.Pointer) IndexExtensionRequestHandler {
	return IndexExtensionRequestHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _IndexExtensionRequestHandlerClass) Alloc() IndexExtensionRequestHandler {
	rv := objc.Call[IndexExtensionRequestHandler](ic, objc.Sel("alloc"))
	return rv
}

func IndexExtensionRequestHandler_Alloc() IndexExtensionRequestHandler {
	return IndexExtensionRequestHandlerClass.Alloc()
}

func (ic _IndexExtensionRequestHandlerClass) New() IndexExtensionRequestHandler {
	rv := objc.Call[IndexExtensionRequestHandler](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIndexExtensionRequestHandler() IndexExtensionRequestHandler {
	return IndexExtensionRequestHandlerClass.New()
}

func (i_ IndexExtensionRequestHandler) Init() IndexExtensionRequestHandler {
	rv := objc.Call[IndexExtensionRequestHandler](i_, objc.Sel("init"))
	return rv
}
