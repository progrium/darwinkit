// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TreeNode] class.
var TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}

type _TreeNodeClass struct {
	objc.Class
}

// An interface definition for the [TreeNode] class.
type ITreeNode interface {
	objc.IObject
	SortWithSortDescriptorsRecursively(sortDescriptors []foundation.ISortDescriptor, recursively bool)
	DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode
	ParentNode() TreeNode
	IsLeaf() bool
	MutableChildNodes() foundation.MutableArray
	ChildNodes() []TreeNode
	IndexPath() foundation.IndexPath
	RepresentedObject() objc.Object
}

// A node in a tree of nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode?language=objc
type TreeNode struct {
	objc.Object
}

func TreeNodeFrom(ptr unsafe.Pointer) TreeNode {
	return TreeNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := objc.Call[TreeNode](tc, objc.Sel("treeNodeWithRepresentedObject:"), modelObject)
	return rv
}

// Creates and returns a tree node that represents the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1537056-treenodewithrepresentedobject?language=objc
func TreeNode_TreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	return TreeNodeClass.TreeNodeWithRepresentedObject(modelObject)
}

func (t_ TreeNode) InitWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := objc.Call[TreeNode](t_, objc.Sel("initWithRepresentedObject:"), modelObject)
	return rv
}

// Initializes a newly allocated tree node that represents the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1533294-initwithrepresentedobject?language=objc
func TreeNode_InitWithRepresentedObject(modelObject objc.IObject) TreeNode {
	return TreeNodeClass.Alloc().InitWithRepresentedObject(modelObject)
}

func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := objc.Call[TreeNode](tc, objc.Sel("alloc"))
	return rv
}

func TreeNode_Alloc() TreeNode {
	return TreeNodeClass.Alloc()
}

func (tc _TreeNodeClass) New() TreeNode {
	rv := objc.Call[TreeNode](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTreeNode() TreeNode {
	return TreeNodeClass.New()
}

func (t_ TreeNode) Init() TreeNode {
	rv := objc.Call[TreeNode](t_, objc.Sel("init"))
	return rv
}

// Sorts the receiver’s subtree using the values of the represented objects with the specified sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1525615-sortwithsortdescriptors?language=objc
func (t_ TreeNode) SortWithSortDescriptorsRecursively(sortDescriptors []foundation.ISortDescriptor, recursively bool) {
	objc.Call[objc.Void](t_, objc.Sel("sortWithSortDescriptors:recursively:"), sortDescriptors, recursively)
}

// Returns the receiver’s descendant at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1534417-descendantnodeatindexpath?language=objc
func (t_ TreeNode) DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode {
	rv := objc.Call[TreeNode](t_, objc.Sel("descendantNodeAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// The receiver’s parent node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1530728-parentnode?language=objc
func (t_ TreeNode) ParentNode() TreeNode {
	rv := objc.Call[TreeNode](t_, objc.Sel("parentNode"))
	return rv
}

// A Boolean that indicates whether the receiver is a leaf node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1532729-leaf?language=objc
func (t_ TreeNode) IsLeaf() bool {
	rv := objc.Call[bool](t_, objc.Sel("isLeaf"))
	return rv
}

// A mutable array that provides read-write access to the receiver’s child nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1527238-mutablechildnodes?language=objc
func (t_ TreeNode) MutableChildNodes() foundation.MutableArray {
	rv := objc.Call[foundation.MutableArray](t_, objc.Sel("mutableChildNodes"))
	return rv
}

// An array containing receiver’s child nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1525285-childnodes?language=objc
func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := objc.Call[[]TreeNode](t_, objc.Sel("childNodes"))
	return rv
}

// The position of the receiver relative to its root parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1532255-indexpath?language=objc
func (t_ TreeNode) IndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](t_, objc.Sel("indexPath"))
	return rv
}

// The object the tree node represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/1531596-representedobject?language=objc
func (t_ TreeNode) RepresentedObject() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("representedObject"))
	return rv
}
