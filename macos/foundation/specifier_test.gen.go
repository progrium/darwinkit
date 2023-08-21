// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpecifierTest] class.
var SpecifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}

type _SpecifierTestClass struct {
	objc.Class
}

// An interface definition for the [SpecifierTest] class.
type ISpecifierTest interface {
	IScriptWhoseTest
}

// A comparison between an object specifier and a test object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspecifiertest?language=objc
type SpecifierTest struct {
	ScriptWhoseTest
}

func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}

func (s_ SpecifierTest) InitWithObjectSpecifierComparisonOperatorTestObject(obj1 IScriptObjectSpecifier, compOp TestComparisonOperation, obj2 objc.IObject) SpecifierTest {
	rv := objc.Call[SpecifierTest](s_, objc.Sel("initWithObjectSpecifier:comparisonOperator:testObject:"), objc.Ptr(obj1), compOp, obj2)
	return rv
}

// Returns a specifier test initialized to evaluate a test object against an object specified by an object specifier using a given comparison operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspecifiertest/1393833-initwithobjectspecifier?language=objc
func NewSpecifierTestWithObjectSpecifierComparisonOperatorTestObject(obj1 IScriptObjectSpecifier, compOp TestComparisonOperation, obj2 objc.IObject) SpecifierTest {
	instance := SpecifierTestClass.Alloc().InitWithObjectSpecifierComparisonOperatorTestObject(obj1, compOp, obj2)
	instance.Autorelease()
	return instance
}

func (sc _SpecifierTestClass) Alloc() SpecifierTest {
	rv := objc.Call[SpecifierTest](sc, objc.Sel("alloc"))
	return rv
}

func SpecifierTest_Alloc() SpecifierTest {
	return SpecifierTestClass.Alloc()
}

func (sc _SpecifierTestClass) New() SpecifierTest {
	rv := objc.Call[SpecifierTest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpecifierTest() SpecifierTest {
	return SpecifierTestClass.New()
}

func (s_ SpecifierTest) Init() SpecifierTest {
	rv := objc.Call[SpecifierTest](s_, objc.Sel("init"))
	return rv
}
