// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptWhoseTest] class.
var ScriptWhoseTestClass = _ScriptWhoseTestClass{objc.GetClass("NSScriptWhoseTest")}

type _ScriptWhoseTestClass struct {
	objc.Class
}

// An interface definition for the [ScriptWhoseTest] class.
type IScriptWhoseTest interface {
	objc.IObject
	IsTrue() bool
}

// An abstract class that provides the basis for testing specifiers one at a time or in groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptwhosetest?language=objc
type ScriptWhoseTest struct {
	objc.Object
}

func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScriptWhoseTest) Init() ScriptWhoseTest {
	rv := objc.Call[ScriptWhoseTest](s_, objc.Sel("init"))
	return rv
}

func (sc _ScriptWhoseTestClass) Alloc() ScriptWhoseTest {
	rv := objc.Call[ScriptWhoseTest](sc, objc.Sel("alloc"))
	return rv
}

func ScriptWhoseTest_Alloc() ScriptWhoseTest {
	return ScriptWhoseTestClass.Alloc()
}

func (sc _ScriptWhoseTestClass) New() ScriptWhoseTest {
	rv := objc.Call[ScriptWhoseTest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptWhoseTest() ScriptWhoseTest {
	return ScriptWhoseTestClass.New()
}

// Returns a Boolean value that indicates whether the test represented by the receiver evaluates to YES. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptwhosetest/1393877-istrue?language=objc
func (s_ ScriptWhoseTest) IsTrue() bool {
	rv := objc.Call[bool](s_, objc.Sel("isTrue"))
	return rv
}
