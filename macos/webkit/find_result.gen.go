// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var FindResultClass = _FindResultClass{objc.GetClass("WKFindResult")}

type _FindResultClass struct {
	objc.Class
}

type IFindResult interface {
	objc.IObject
	MatchFound() bool
}

type FindResult struct {
	objc.Object
}

func MakeFindResult(ptr unsafe.Pointer) FindResult {
	return FindResult{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FindResultClass) Alloc() FindResult {
	rv := objc.CallMethod[FindResult](fc, objc.GetSelector("alloc"))
	return rv
}

func FindResult_Alloc() FindResult {
	return FindResultClass.Alloc()
}

func (fc _FindResultClass) New() FindResult {
	rv := objc.CallMethod[FindResult](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFindResult() FindResult {
	return FindResultClass.New()
}

func FindResult_New() FindResult {
	return FindResultClass.New()
}

func (f_ FindResult) Init() FindResult {
	rv := objc.CallMethod[FindResult](f_, objc.GetSelector("init"))
	return rv
}

func FindResult_Init() FindResult {
	return FindResultClass.Alloc().Init()
}

func (f_ FindResult) MatchFound() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("matchFound"))
	return rv
}
