// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TreeController] class.
var TreeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}

type _TreeControllerClass struct {
	objc.Class
}

// An interface definition for the [TreeController] class.
type ITreeController interface {
	IObjectController
	LeafKeyPathForNode(node ITreeNode) string
	RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IIndexPath)
	RemoveSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	SetSelectionIndexPath(indexPath foundation.IIndexPath) bool
	MoveNodeToIndexPath(node ITreeNode, indexPath foundation.IIndexPath)
	AddSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	MoveNodesToIndexPath(nodes []ITreeNode, startingIndexPath foundation.IIndexPath)
	CountKeyPathForNode(node ITreeNode) string
	SetSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	InsertObjectAtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IIndexPath)
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IIndexPath)
	RearrangeObjects()
	InsertChild(sender objc.IObject) objc.Object
	AddChild(sender objc.IObject) objc.Object
	Insert(sender objc.IObject) objc.Object
	InsertObjectsAtArrangedObjectIndexPaths(objects []objc.IObject, indexPaths []foundation.IIndexPath)
	ChildrenKeyPathForNode(node ITreeNode) string
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	ChildrenKeyPath() string
	SetChildrenKeyPath(value string)
	SelectionIndexPath() foundation.IndexPath
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	ArrangedObjects() TreeNode
	SelectionIndexPaths() []foundation.IndexPath
	CanInsertChild() bool
	CanInsert() bool
	LeafKeyPath() string
	SetLeafKeyPath(value string)
	CanAddChild() bool
	CountKeyPath() string
	SetCountKeyPath(value string)
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SelectedNodes() []TreeNode
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
}

// A bindings-compatible controller that manages a tree of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller?language=objc
type TreeController struct {
	ObjectController
}

func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}

func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.Call[TreeController](tc, objc.Sel("alloc"))
	return rv
}

func TreeController_Alloc() TreeController {
	return TreeControllerClass.Alloc()
}

func (tc _TreeControllerClass) New() TreeController {
	rv := objc.Call[TreeController](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTreeController() TreeController {
	return TreeControllerClass.New()
}

func (t_ TreeController) Init() TreeController {
	rv := objc.Call[TreeController](t_, objc.Sel("init"))
	return rv
}

func (t_ TreeController) InitWithContent(content objc.IObject) TreeController {
	rv := objc.Call[TreeController](t_, objc.Sel("initWithContent:"), content)
	return rv
}

// Initializes and returns an NSObjectController object with the given content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1529422-initwithcontent?language=objc
func TreeController_InitWithContent(content objc.IObject) TreeController {
	return TreeControllerClass.Alloc().InitWithContent(content)
}

// Returns the key path that specifies whether the node is a leaf node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1526740-leafkeypathfornode?language=objc
func (t_ TreeController) LeafKeyPathForNode(node ITreeNode) string {
	rv := objc.Call[string](t_, objc.Sel("leafKeyPathForNode:"), objc.Ptr(node))
	return rv
}

// Removes the object at the specified indexPath in the tree controller’s arranged objects from the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1530806-removeobjectatarrangedobjectinde?language=objc
func (t_ TreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("removeObjectAtArrangedObjectIndexPath:"), objc.Ptr(indexPath))
}

// Removes the objects at the specified indexPaths from the tree controller’s current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1532401-removeselectionindexpaths?language=objc
func (t_ TreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.Call[bool](t_, objc.Sel("removeSelectionIndexPaths:"), indexPaths)
	return rv
}

// Sets the tree controller’s current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1527490-setselectionindexpath?language=objc
func (t_ TreeController) SetSelectionIndexPath(indexPath foundation.IIndexPath) bool {
	rv := objc.Call[bool](t_, objc.Sel("setSelectionIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Moves the specified tree node to the new index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1532133-movenode?language=objc
func (t_ TreeController) MoveNodeToIndexPath(node ITreeNode, indexPath foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("moveNode:toIndexPath:"), objc.Ptr(node), objc.Ptr(indexPath))
}

// Adds the objects at the specified indexPaths in the tree controller’s content to the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1535862-addselectionindexpaths?language=objc
func (t_ TreeController) AddSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.Call[bool](t_, objc.Sel("addSelectionIndexPaths:"), indexPaths)
	return rv
}

// Moves the specified tree nodes to the new index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1530337-movenodes?language=objc
func (t_ TreeController) MoveNodesToIndexPath(nodes []ITreeNode, startingIndexPath foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("moveNodes:toIndexPath:"), nodes, objc.Ptr(startingIndexPath))
}

// Returns the key path that provides the number of children for a specified node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1528109-countkeypathfornode?language=objc
func (t_ TreeController) CountKeyPathForNode(node ITreeNode) string {
	rv := objc.Call[string](t_, objc.Sel("countKeyPathForNode:"), objc.Ptr(node))
	return rv
}

// Sets the tree controller’s current selection to the specified index paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1528494-setselectionindexpaths?language=objc
func (t_ TreeController) SetSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.Call[bool](t_, objc.Sel("setSelectionIndexPaths:"), indexPaths)
	return rv
}

// Inserts object into the tree controller’s arranged objects array at the location specified by indexPath, and adds it to the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1530156-insertobject?language=objc
func (t_ TreeController) InsertObjectAtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("insertObject:atArrangedObjectIndexPath:"), object, objc.Ptr(indexPath))
}

// Removes the objects at the specified indexPaths in the tree controller’s arranged objects from the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1525602-removeobjectsatarrangedobjectind?language=objc
func (t_ TreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("removeObjectsAtArrangedObjectIndexPaths:"), indexPaths)
}

// Use this method to trigger reordering of the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1525868-rearrangeobjects?language=objc
func (t_ TreeController) RearrangeObjects() {
	objc.Call[objc.Void](t_, objc.Sel("rearrangeObjects"))
}

// Creates a new object of the class specified by objectClass and inserts it into the tree controller’s content as a child of the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1524323-insertchild?language=objc
func (t_ TreeController) InsertChild(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("insertChild:"), sender)
	return rv
}

// Adds a child object to the currently selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1527676-addchild?language=objc
func (t_ TreeController) AddChild(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("addChild:"), sender)
	return rv
}

// Creates a new object of the class specified by objectClass and inserts it into the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1533671-insert?language=objc
func (t_ TreeController) Insert(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("insert:"), sender)
	return rv
}

// Inserts objects into the tree controller’s arranged objects array at the locations specified in indexPaths, and adds them to the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1530531-insertobjects?language=objc
func (t_ TreeController) InsertObjectsAtArrangedObjectIndexPaths(objects []objc.IObject, indexPaths []foundation.IIndexPath) {
	objc.Call[objc.Void](t_, objc.Sel("insertObjects:atArrangedObjectIndexPaths:"), objects, indexPaths)
}

// Returns the key path used to find the children in the specified tree node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1525412-childrenkeypathfornode?language=objc
func (t_ TreeController) ChildrenKeyPathForNode(node ITreeNode) string {
	rv := objc.Call[string](t_, objc.Sel("childrenKeyPathForNode:"), objc.Ptr(node))
	return rv
}

// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1524473-preservesselection?language=objc
func (t_ TreeController) PreservesSelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("preservesSelection"))
	return rv
}

// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1524473-preservesselection?language=objc
func (t_ TreeController) SetPreservesSelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setPreservesSelection:"), value)
}

// The key path used to find the children in the tree controller’s objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1528721-childrenkeypath?language=objc
func (t_ TreeController) ChildrenKeyPath() string {
	rv := objc.Call[string](t_, objc.Sel("childrenKeyPath"))
	return rv
}

// The key path used to find the children in the tree controller’s objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1528721-childrenkeypath?language=objc
func (t_ TreeController) SetChildrenKeyPath(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setChildrenKeyPath:"), value)
}

// The index path of the first selected object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1533951-selectionindexpath?language=objc
func (t_ TreeController) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](t_, objc.Sel("selectionIndexPath"))
	return rv
}

// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1529530-alwaysusesmultiplevaluesmarker?language=objc
func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Call[bool](t_, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}

// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1529530-alwaysusesmultiplevaluesmarker?language=objc
func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}

// An array containing the sort descriptors used to arrange the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1526827-sortdescriptors?language=objc
func (t_ TreeController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Call[[]foundation.SortDescriptor](t_, objc.Sel("sortDescriptors"))
	return rv
}

// An array containing the sort descriptors used to arrange the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1526827-sortdescriptors?language=objc
func (t_ TreeController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.Call[objc.Void](t_, objc.Sel("setSortDescriptors:"), value)
}

// The tree controller’s sorted content objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1527465-arrangedobjects?language=objc
func (t_ TreeController) ArrangedObjects() TreeNode {
	rv := objc.Call[TreeNode](t_, objc.Sel("arrangedObjects"))
	return rv
}

// An array containing the index paths of the currently selected objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1534861-selectionindexpaths?language=objc
func (t_ TreeController) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.Call[[]foundation.IndexPath](t_, objc.Sel("selectionIndexPaths"))
	return rv
}

// A Boolean value that indicates if a child object can be inserted into the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1524647-caninsertchild?language=objc
func (t_ TreeController) CanInsertChild() bool {
	rv := objc.Call[bool](t_, objc.Sel("canInsertChild"))
	return rv
}

// A Boolean value that indicates if an object can be inserted into the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1534180-caninsert?language=objc
func (t_ TreeController) CanInsert() bool {
	rv := objc.Call[bool](t_, objc.Sel("canInsert"))
	return rv
}

// The key path used by the tree controller to determine if a node is a leaf key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1532164-leafkeypath?language=objc
func (t_ TreeController) LeafKeyPath() string {
	rv := objc.Call[string](t_, objc.Sel("leafKeyPath"))
	return rv
}

// The key path used by the tree controller to determine if a node is a leaf key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1532164-leafkeypath?language=objc
func (t_ TreeController) SetLeafKeyPath(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setLeafKeyPath:"), value)
}

// A Boolean value that indicates if a child object can be added to the tree controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1525790-canaddchild?language=objc
func (t_ TreeController) CanAddChild() bool {
	rv := objc.Call[bool](t_, objc.Sel("canAddChild"))
	return rv
}

// The key path used to find the number of children for a node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1529127-countkeypath?language=objc
func (t_ TreeController) CountKeyPath() string {
	rv := objc.Call[string](t_, objc.Sel("countKeyPath"))
	return rv
}

// The key path used to find the number of children for a node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1529127-countkeypath?language=objc
func (t_ TreeController) SetCountKeyPath(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setCountKeyPath:"), value)
}

// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1534200-selectsinsertedobjects?language=objc
func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := objc.Call[bool](t_, objc.Sel("selectsInsertedObjects"))
	return rv
}

// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1534200-selectsinsertedobjects?language=objc
func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectsInsertedObjects:"), value)
}

// An array containing the tree controller’s selected tree nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1534151-selectednodes?language=objc
func (t_ TreeController) SelectedNodes() []TreeNode {
	rv := objc.Call[[]TreeNode](t_, objc.Sel("selectedNodes"))
	return rv
}

// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1526188-avoidsemptyselection?language=objc
func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("avoidsEmptySelection"))
	return rv
}

// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/1526188-avoidsemptyselection?language=objc
func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAvoidsEmptySelection:"), value)
}
