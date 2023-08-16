// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FindResult] class.
var FindResultClass = _FindResultClass{objc.GetClass("WKFindResult")}

type _FindResultClass struct {
	objc.Class
}

// An interface definition for the [FindResult] class.
type IFindResult interface {
	objc.IObject
	MatchFound() bool
}

// An object that contains the results of searching the web view’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindresult?language=objc
type FindResult struct {
	objc.Object
}

func FindResultFrom(ptr unsafe.Pointer) FindResult {
	return FindResult{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FindResultClass) Alloc() FindResult {
	rv := objc.Call[FindResult](fc, objc.Sel("alloc"))
	return rv
}

func FindResult_Alloc() FindResult {
	return FindResultClass.Alloc()
}

func (fc _FindResultClass) New() FindResult {
	rv := objc.Call[FindResult](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFindResult() FindResult {
	return FindResultClass.New()
}

func (f_ FindResult) Init() FindResult {
	rv := objc.Call[FindResult](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the web view found a match during the search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindresult/3516858-matchfound?language=objc
func (f_ FindResult) MatchFound() bool {
	rv := objc.Call[bool](f_, objc.Sel("matchFound"))
	return rv
}
