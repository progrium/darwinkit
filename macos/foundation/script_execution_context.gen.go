// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptExecutionContext] class.
var ScriptExecutionContextClass = _ScriptExecutionContextClass{objc.GetClass("NSScriptExecutionContext")}

type _ScriptExecutionContextClass struct {
	objc.Class
}

// An interface definition for the [ScriptExecutionContext] class.
type IScriptExecutionContext interface {
	objc.IObject
	TopLevelObject() objc.Object
	SetTopLevelObject(value objc.IObject)
	ObjectBeingTested() objc.Object
	SetObjectBeingTested(value objc.IObject)
	RangeContainerObject() objc.Object
	SetRangeContainerObject(value objc.IObject)
}

// The context in which the current script command is executed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext?language=objc
type ScriptExecutionContext struct {
	objc.Object
}

func ScriptExecutionContextFrom(ptr unsafe.Pointer) ScriptExecutionContext {
	return ScriptExecutionContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScriptExecutionContextClass) Alloc() ScriptExecutionContext {
	rv := objc.Call[ScriptExecutionContext](sc, objc.Sel("alloc"))
	return rv
}

func ScriptExecutionContext_Alloc() ScriptExecutionContext {
	return ScriptExecutionContextClass.Alloc()
}

func (sc _ScriptExecutionContextClass) New() ScriptExecutionContext {
	rv := objc.Call[ScriptExecutionContext](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptExecutionContext() ScriptExecutionContext {
	return ScriptExecutionContextClass.New()
}

func (s_ ScriptExecutionContext) Init() ScriptExecutionContext {
	rv := objc.Call[ScriptExecutionContext](s_, objc.Sel("init"))
	return rv
}

// Returns the shared NSScriptExecutionContext instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1408607-sharedscriptexecutioncontext?language=objc
func (sc _ScriptExecutionContextClass) SharedScriptExecutionContext() ScriptExecutionContext {
	rv := objc.Call[ScriptExecutionContext](sc, objc.Sel("sharedScriptExecutionContext"))
	return rv
}

// Returns the shared NSScriptExecutionContext instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1408607-sharedscriptexecutioncontext?language=objc
func ScriptExecutionContext_SharedScriptExecutionContext() ScriptExecutionContext {
	return ScriptExecutionContextClass.SharedScriptExecutionContext()
}

// Sets the top-level object for an object-specifier evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1415288-toplevelobject?language=objc
func (s_ ScriptExecutionContext) TopLevelObject() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("topLevelObject"))
	return rv
}

// Sets the top-level object for an object-specifier evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1415288-toplevelobject?language=objc
func (s_ ScriptExecutionContext) SetTopLevelObject(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setTopLevelObject:"), value)
}

// Sets the top-level container object currently being tested in a “whose” qualifier to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1412411-objectbeingtested?language=objc
func (s_ ScriptExecutionContext) ObjectBeingTested() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("objectBeingTested"))
	return rv
}

// Sets the top-level container object currently being tested in a “whose” qualifier to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1412411-objectbeingtested?language=objc
func (s_ ScriptExecutionContext) SetObjectBeingTested(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setObjectBeingTested:"), value)
}

// Sets the top-level container object for a range-specifier evaluation to a give object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1416391-rangecontainerobject?language=objc
func (s_ ScriptExecutionContext) RangeContainerObject() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("rangeContainerObject"))
	return rv
}

// Sets the top-level container object for a range-specifier evaluation to a give object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/1416391-rangecontainerobject?language=objc
func (s_ ScriptExecutionContext) SetRangeContainerObject(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setRangeContainerObject:"), value)
}
