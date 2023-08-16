// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BlockOperation] class.
var BlockOperationClass = _BlockOperationClass{objc.GetClass("NSBlockOperation")}

type _BlockOperationClass struct {
	objc.Class
}

// An interface definition for the [BlockOperation] class.
type IBlockOperation interface {
	IOperation
	AddExecutionBlock(block func())
	ExecutionBlocks() []func()
}

// An operation that manages the concurrent execution of one or more blocks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsblockoperation?language=objc
type BlockOperation struct {
	Operation
}

func BlockOperationFrom(ptr unsafe.Pointer) BlockOperation {
	return BlockOperation{
		Operation: OperationFrom(ptr),
	}
}

func (bc _BlockOperationClass) BlockOperationWithBlock(block func()) BlockOperation {
	rv := objc.Call[BlockOperation](bc, objc.Sel("blockOperationWithBlock:"), block)
	return rv
}

// Creates and returns an NSBlockOperation object and adds the specified block to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsblockoperation/1412757-blockoperationwithblock?language=objc
func BlockOperation_BlockOperationWithBlock(block func()) BlockOperation {
	return BlockOperationClass.BlockOperationWithBlock(block)
}

func (bc _BlockOperationClass) Alloc() BlockOperation {
	rv := objc.Call[BlockOperation](bc, objc.Sel("alloc"))
	return rv
}

func BlockOperation_Alloc() BlockOperation {
	return BlockOperationClass.Alloc()
}

func (bc _BlockOperationClass) New() BlockOperation {
	rv := objc.Call[BlockOperation](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBlockOperation() BlockOperation {
	return BlockOperationClass.New()
}

func (b_ BlockOperation) Init() BlockOperation {
	rv := objc.Call[BlockOperation](b_, objc.Sel("init"))
	return rv
}

// Adds the specified block to the receiver’s list of blocks to perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsblockoperation/1414623-addexecutionblock?language=objc
func (b_ BlockOperation) AddExecutionBlock(block func()) {
	objc.Call[objc.Void](b_, objc.Sel("addExecutionBlock:"), block)
}

// The blocks associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsblockoperation/1416555-executionblocks?language=objc
func (b_ BlockOperation) ExecutionBlocks() []func() {
	rv := objc.Call[[]func()](b_, objc.Sel("executionBlocks"))
	return rv
}
