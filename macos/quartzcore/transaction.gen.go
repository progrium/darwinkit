// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Transaction] class.
var TransactionClass = _TransactionClass{objc.GetClass("CATransaction")}

type _TransactionClass struct {
	objc.Class
}

// An interface definition for the [Transaction] class.
type ITransaction interface {
	objc.IObject
}

// A mechanism for grouping multiple layer-tree operations into atomic updates to the render tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction?language=objc
type Transaction struct {
	objc.Object
}

func TransactionFrom(ptr unsafe.Pointer) Transaction {
	return Transaction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TransactionClass) Alloc() Transaction {
	rv := objc.Call[Transaction](tc, objc.Sel("alloc"))
	return rv
}

func Transaction_Alloc() Transaction {
	return TransactionClass.Alloc()
}

func (tc _TransactionClass) New() Transaction {
	rv := objc.Call[Transaction](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTransaction() Transaction {
	return TransactionClass.New()
}

func (t_ Transaction) Init() Transaction {
	rv := objc.Call[Transaction](t_, objc.Sel("init"))
	return rv
}

// Returns the timing function used for all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448269-animationtimingfunction?language=objc
func (tc _TransactionClass) AnimationTimingFunction() MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](tc, objc.Sel("animationTimingFunction"))
	return rv
}

// Returns the timing function used for all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448269-animationtimingfunction?language=objc
func Transaction_AnimationTimingFunction() MediaTimingFunction {
	return TransactionClass.AnimationTimingFunction()
}

// Sets the completion block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448281-setcompletionblock?language=objc
func (tc _TransactionClass) SetCompletionBlock(block func()) {
	objc.Call[objc.Void](tc, objc.Sel("setCompletionBlock:"), block)
}

// Sets the completion block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448281-setcompletionblock?language=objc
func Transaction_SetCompletionBlock(block func()) {
	TransactionClass.SetCompletionBlock(block)
}

// Begin a new transaction for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448282-begin?language=objc
func (tc _TransactionClass) Begin() {
	objc.Call[objc.Void](tc, objc.Sel("begin"))
}

// Begin a new transaction for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448282-begin?language=objc
func Transaction_Begin() {
	TransactionClass.Begin()
}

// Attempts to acquire a recursive spin-lock lock, ensuring that returned layer values are valid until unlocked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448267-lock?language=objc
func (tc _TransactionClass) Lock() {
	objc.Call[objc.Void](tc, objc.Sel("lock"))
}

// Attempts to acquire a recursive spin-lock lock, ensuring that returned layer values are valid until unlocked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448267-lock?language=objc
func Transaction_Lock() {
	TransactionClass.Lock()
}

// Returns whether actions triggered as a result of property changes made within this transaction group are suppressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448276-disableactions?language=objc
func (tc _TransactionClass) DisableActions() bool {
	rv := objc.Call[bool](tc, objc.Sel("disableActions"))
	return rv
}

// Returns whether actions triggered as a result of property changes made within this transaction group are suppressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448276-disableactions?language=objc
func Transaction_DisableActions() bool {
	return TransactionClass.DisableActions()
}

// Commit all changes made during the current transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448255-commit?language=objc
func (tc _TransactionClass) Commit() {
	objc.Call[objc.Void](tc, objc.Sel("commit"))
}

// Commit all changes made during the current transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448255-commit?language=objc
func Transaction_Commit() {
	TransactionClass.Commit()
}

// Flushes any extant implicit transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448270-flush?language=objc
func (tc _TransactionClass) Flush() {
	objc.Call[objc.Void](tc, objc.Sel("flush"))
}

// Flushes any extant implicit transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448270-flush?language=objc
func Transaction_Flush() {
	TransactionClass.Flush()
}

// Sets the timing function used for all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448279-setanimationtimingfunction?language=objc
func (tc _TransactionClass) SetAnimationTimingFunction(function IMediaTimingFunction) {
	objc.Call[objc.Void](tc, objc.Sel("setAnimationTimingFunction:"), objc.Ptr(function))
}

// Sets the timing function used for all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448279-setanimationtimingfunction?language=objc
func Transaction_SetAnimationTimingFunction(function IMediaTimingFunction) {
	TransactionClass.SetAnimationTimingFunction(function)
}

// Sets the animation duration used by all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448283-setanimationduration?language=objc
func (tc _TransactionClass) SetAnimationDuration(dur corefoundation.TimeInterval) {
	objc.Call[objc.Void](tc, objc.Sel("setAnimationDuration:"), dur)
}

// Sets the animation duration used by all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448283-setanimationduration?language=objc
func Transaction_SetAnimationDuration(dur corefoundation.TimeInterval) {
	TransactionClass.SetAnimationDuration(dur)
}

// Sets whether actions triggered as a result of property changes made within this transaction group are suppressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448261-setdisableactions?language=objc
func (tc _TransactionClass) SetDisableActions(flag bool) {
	objc.Call[objc.Void](tc, objc.Sel("setDisableActions:"), flag)
}

// Sets whether actions triggered as a result of property changes made within this transaction group are suppressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448261-setdisableactions?language=objc
func Transaction_SetDisableActions(flag bool) {
	TransactionClass.SetDisableActions(flag)
}

// Returns the completion block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448280-completionblock?language=objc
func (tc _TransactionClass) CompletionBlock() func() {
	rv := objc.Call[func()](tc, objc.Sel("completionBlock"))
	return rv
}

// Returns the completion block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448280-completionblock?language=objc
func Transaction_CompletionBlock() func() {
	return TransactionClass.CompletionBlock()
}

// Relinquishes a previously acquired transaction lock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448285-unlock?language=objc
func (tc _TransactionClass) Unlock() {
	objc.Call[objc.Void](tc, objc.Sel("unlock"))
}

// Relinquishes a previously acquired transaction lock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448285-unlock?language=objc
func Transaction_Unlock() {
	TransactionClass.Unlock()
}

// Sets the arbitrary keyed-data for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448278-setvalue?language=objc
func (tc _TransactionClass) SetValueForKey(anObject objc.IObject, key string) {
	objc.Call[objc.Void](tc, objc.Sel("setValue:forKey:"), anObject, key)
}

// Sets the arbitrary keyed-data for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448278-setvalue?language=objc
func Transaction_SetValueForKey(anObject objc.IObject, key string) {
	TransactionClass.SetValueForKey(anObject, key)
}

// Returns the animation duration used by all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448263-animationduration?language=objc
func (tc _TransactionClass) AnimationDuration() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](tc, objc.Sel("animationDuration"))
	return rv
}

// Returns the animation duration used by all animations within this transaction group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448263-animationduration?language=objc
func Transaction_AnimationDuration() corefoundation.TimeInterval {
	return TransactionClass.AnimationDuration()
}

// Returns the arbitrary keyed-data specified by the given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448259-valueforkey?language=objc
func (tc _TransactionClass) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](tc, objc.Sel("valueForKey:"), key)
	return rv
}

// Returns the arbitrary keyed-data specified by the given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransaction/1448259-valueforkey?language=objc
func Transaction_ValueForKey(key string) objc.Object {
	return TransactionClass.ValueForKey(key)
}
