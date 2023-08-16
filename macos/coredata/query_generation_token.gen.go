// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QueryGenerationToken] class.
var QueryGenerationTokenClass = _QueryGenerationTokenClass{objc.GetClass("NSQueryGenerationToken")}

type _QueryGenerationTokenClass struct {
	objc.Class
}

// An interface definition for the [QueryGenerationToken] class.
type IQueryGenerationToken interface {
	objc.IObject
}

// A token that indicates which generation of the persistent store is being accessed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsquerygenerationtoken?language=objc
type QueryGenerationToken struct {
	objc.Object
}

func QueryGenerationTokenFrom(ptr unsafe.Pointer) QueryGenerationToken {
	return QueryGenerationToken{
		Object: objc.ObjectFrom(ptr),
	}
}

func (qc _QueryGenerationTokenClass) Alloc() QueryGenerationToken {
	rv := objc.Call[QueryGenerationToken](qc, objc.Sel("alloc"))
	return rv
}

func QueryGenerationToken_Alloc() QueryGenerationToken {
	return QueryGenerationTokenClass.Alloc()
}

func (qc _QueryGenerationTokenClass) New() QueryGenerationToken {
	rv := objc.Call[QueryGenerationToken](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQueryGenerationToken() QueryGenerationToken {
	return QueryGenerationTokenClass.New()
}

func (q_ QueryGenerationToken) Init() QueryGenerationToken {
	rv := objc.Call[QueryGenerationToken](q_, objc.Sel("init"))
	return rv
}

// A token that informs a context to use the current generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsquerygenerationtoken/1640578-currentquerygenerationtoken?language=objc
func (qc _QueryGenerationTokenClass) CurrentQueryGenerationToken() QueryGenerationToken {
	rv := objc.Call[QueryGenerationToken](qc, objc.Sel("currentQueryGenerationToken"))
	return rv
}

// A token that informs a context to use the current generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsquerygenerationtoken/1640578-currentquerygenerationtoken?language=objc
func QueryGenerationToken_CurrentQueryGenerationToken() QueryGenerationToken {
	return QueryGenerationTokenClass.CurrentQueryGenerationToken()
}
