// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}

type _TreeNodeClass struct {
	objc.Class
}

type ITreeNode interface {
	objc.IObject
	DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode
	SortWithSortDescriptorsRecursively(sortDescriptors []foundation.ISortDescriptor, recursively bool)
	RepresentedObject() objc.Object
	IndexPath() foundation.IndexPath
	IsLeaf() bool
	ChildNodes() []TreeNode
	ParentNode() TreeNode
}

type TreeNode struct {
	objc.Object
}

func MakeTreeNode(ptr unsafe.Pointer) TreeNode {
	return TreeNode{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := objc.CallMethod[TreeNode](tc, objc.GetSelector("treeNodeWithRepresentedObject:"), objc.ExtractPtr(modelObject))
	return rv
}

func TreeNode_TreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	return TreeNodeClass.TreeNodeWithRepresentedObject(modelObject)
}

func (t_ TreeNode) InitWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := objc.CallMethod[TreeNode](t_, objc.GetSelector("initWithRepresentedObject:"), objc.ExtractPtr(modelObject))
	return rv
}

func TreeNode_InitWithRepresentedObject(modelObject objc.IObject) TreeNode {
	return TreeNodeClass.Alloc().InitWithRepresentedObject(modelObject)
}

func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := objc.CallMethod[TreeNode](tc, objc.GetSelector("alloc"))
	return rv
}

func TreeNode_Alloc() TreeNode {
	return TreeNodeClass.Alloc()
}

func (tc _TreeNodeClass) New() TreeNode {
	rv := objc.CallMethod[TreeNode](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTreeNode() TreeNode {
	return TreeNodeClass.New()
}

func TreeNode_New() TreeNode {
	return TreeNodeClass.New()
}

func (t_ TreeNode) Init() TreeNode {
	rv := objc.CallMethod[TreeNode](t_, objc.GetSelector("init"))
	return rv
}

func TreeNode_Init() TreeNode {
	return TreeNodeClass.Alloc().Init()
}

func (t_ TreeNode) DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode {
	rv := objc.CallMethod[TreeNode](t_, objc.GetSelector("descendantNodeAtIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (t_ TreeNode) SortWithSortDescriptorsRecursively(sortDescriptors []foundation.ISortDescriptor, recursively bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("sortWithSortDescriptors:recursively:"), sortDescriptors, recursively)
}

func (t_ TreeNode) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("representedObject"))
	return rv
}

func (t_ TreeNode) IndexPath() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](t_, objc.GetSelector("indexPath"))
	return rv
}

func (t_ TreeNode) IsLeaf() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isLeaf"))
	return rv
}

func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := objc.CallMethod[[]TreeNode](t_, objc.GetSelector("childNodes"))
	return rv
}

func (t_ TreeNode) ParentNode() TreeNode {
	rv := objc.CallMethod[TreeNode](t_, objc.GetSelector("parentNode"))
	return rv
}
