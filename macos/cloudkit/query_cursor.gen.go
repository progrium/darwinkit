// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QueryCursor] class.
var QueryCursorClass = _QueryCursorClass{objc.GetClass("CKQueryCursor")}

type _QueryCursorClass struct {
	objc.Class
}

// An interface definition for the [QueryCursor] class.
type IQueryCursor interface {
	objc.IObject
}

// An object that marks the stopping point for a query and the starting point for retrieving the remaining results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerycursor?language=objc
type QueryCursor struct {
	objc.Object
}

func QueryCursorFrom(ptr unsafe.Pointer) QueryCursor {
	return QueryCursor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (qc _QueryCursorClass) Alloc() QueryCursor {
	rv := objc.Call[QueryCursor](qc, objc.Sel("alloc"))
	return rv
}

func QueryCursor_Alloc() QueryCursor {
	return QueryCursorClass.Alloc()
}

func (qc _QueryCursorClass) New() QueryCursor {
	rv := objc.Call[QueryCursor](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQueryCursor() QueryCursor {
	return QueryCursorClass.New()
}

func (q_ QueryCursor) Init() QueryCursor {
	rv := objc.Call[QueryCursor](q_, objc.Sel("init"))
	return rv
}
