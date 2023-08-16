// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LogicalTest] class.
var LogicalTestClass = _LogicalTestClass{objc.GetClass("NSLogicalTest")}

type _LogicalTestClass struct {
	objc.Class
}

// An interface definition for the [LogicalTest] class.
type ILogicalTest interface {
	IScriptWhoseTest
}

// The logical combination of one or more specifier tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslogicaltest?language=objc
type LogicalTest struct {
	ScriptWhoseTest
}

func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}

func (l_ LogicalTest) InitAndTestWithTests(subTests []ISpecifierTest) LogicalTest {
	rv := objc.Call[LogicalTest](l_, objc.Sel("initAndTestWithTests:"), subTests)
	return rv
}

// Returns an NSLogicalTest object initialized to perform an AND operation with the NSSpecifierTest objects in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslogicaltest/1393854-initandtestwithtests?language=objc
func LogicalTest_InitAndTestWithTests(subTests []ISpecifierTest) LogicalTest {
	return LogicalTestClass.Alloc().InitAndTestWithTests(subTests)
}

func (l_ LogicalTest) InitOrTestWithTests(subTests []ISpecifierTest) LogicalTest {
	rv := objc.Call[LogicalTest](l_, objc.Sel("initOrTestWithTests:"), subTests)
	return rv
}

// Returns an NSLogicalTest object initialized to perform an OR operation with the NSSpecifierTest objects in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslogicaltest/1393875-initortestwithtests?language=objc
func LogicalTest_InitOrTestWithTests(subTests []ISpecifierTest) LogicalTest {
	return LogicalTestClass.Alloc().InitOrTestWithTests(subTests)
}

func (l_ LogicalTest) InitNotTestWithTest(subTest IScriptWhoseTest) LogicalTest {
	rv := objc.Call[LogicalTest](l_, objc.Sel("initNotTestWithTest:"), objc.Ptr(subTest))
	return rv
}

// Returns an NSLogicalTest object initialized to perform a NOT operation on the given NSScriptWhoseTest object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslogicaltest/1393879-initnottestwithtest?language=objc
func LogicalTest_InitNotTestWithTest(subTest IScriptWhoseTest) LogicalTest {
	return LogicalTestClass.Alloc().InitNotTestWithTest(subTest)
}

func (lc _LogicalTestClass) Alloc() LogicalTest {
	rv := objc.Call[LogicalTest](lc, objc.Sel("alloc"))
	return rv
}

func LogicalTest_Alloc() LogicalTest {
	return LogicalTestClass.Alloc()
}

func (lc _LogicalTestClass) New() LogicalTest {
	rv := objc.Call[LogicalTest](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLogicalTest() LogicalTest {
	return LogicalTestClass.New()
}

func (l_ LogicalTest) Init() LogicalTest {
	rv := objc.Call[LogicalTest](l_, objc.Sel("init"))
	return rv
}
