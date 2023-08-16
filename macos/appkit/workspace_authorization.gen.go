// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WorkspaceAuthorization] class.
var WorkspaceAuthorizationClass = _WorkspaceAuthorizationClass{objc.GetClass("NSWorkspaceAuthorization")}

type _WorkspaceAuthorizationClass struct {
	objc.Class
}

// An interface definition for the [WorkspaceAuthorization] class.
type IWorkspaceAuthorization interface {
	objc.IObject
}

// The authorization granted to the app by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceauthorization?language=objc
type WorkspaceAuthorization struct {
	objc.Object
}

func WorkspaceAuthorizationFrom(ptr unsafe.Pointer) WorkspaceAuthorization {
	return WorkspaceAuthorization{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WorkspaceAuthorizationClass) Alloc() WorkspaceAuthorization {
	rv := objc.Call[WorkspaceAuthorization](wc, objc.Sel("alloc"))
	return rv
}

func WorkspaceAuthorization_Alloc() WorkspaceAuthorization {
	return WorkspaceAuthorizationClass.Alloc()
}

func (wc _WorkspaceAuthorizationClass) New() WorkspaceAuthorization {
	rv := objc.Call[WorkspaceAuthorization](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWorkspaceAuthorization() WorkspaceAuthorization {
	return WorkspaceAuthorizationClass.New()
}

func (w_ WorkspaceAuthorization) Init() WorkspaceAuthorization {
	rv := objc.Call[WorkspaceAuthorization](w_, objc.Sel("init"))
	return rv
}
