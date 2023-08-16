// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLDTD] class.
var XMLDTDClass = _XMLDTDClass{objc.GetClass("NSXMLDTD")}

type _XMLDTDClass struct {
	objc.Class
}

// An interface definition for the [XMLDTD] class.
type IXMLDTD interface {
	IXMLNode
	NotationDeclarationForName(name string) XMLDTDNode
	RemoveChildAtIndex(index uint)
	AttributeDeclarationForNameElementName(name string, elementName string) XMLDTDNode
	EntityDeclarationForName(name string) XMLDTDNode
	ElementDeclarationForName(name string) XMLDTDNode
	InsertChildrenAtIndex(children []IXMLNode, index uint)
	InsertChildAtIndex(child IXMLNode, index uint)
	SetChildren(children []IXMLNode)
	AddChild(child IXMLNode)
	ReplaceChildAtIndexWithNode(index uint, node IXMLNode)
	SystemID() string
	SetSystemID(value string)
	PublicID() string
	SetPublicID(value string)
}

// A representation of a Document Type Definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd?language=objc
type XMLDTD struct {
	XMLNode
}

func XMLDTDFrom(ptr unsafe.Pointer) XMLDTD {
	return XMLDTD{
		XMLNode: XMLNodeFrom(ptr),
	}
}

func (x_ XMLDTD) InitWithDataOptionsError(data []byte, mask XMLNodeOptions, error IError) XMLDTD {
	rv := objc.Call[XMLDTD](x_, objc.Sel("initWithData:options:error:"), data, mask, objc.Ptr(error))
	return rv
}

// Initializes and returns an NSXMLDTD object created from the DTD declarations encapsulated in an NSData object [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1412807-initwithdata?language=objc
func XMLDTD_InitWithDataOptionsError(data []byte, mask XMLNodeOptions, error IError) XMLDTD {
	return XMLDTDClass.Alloc().InitWithDataOptionsError(data, mask, error)
}

func (x_ XMLDTD) InitWithContentsOfURLOptionsError(url IURL, mask XMLNodeOptions, error IError) XMLDTD {
	rv := objc.Call[XMLDTD](x_, objc.Sel("initWithContentsOfURL:options:error:"), objc.Ptr(url), mask, objc.Ptr(error))
	return rv
}

// Initializes and returns an NSXMLDTD object created from the DTD declarations in a URL-referenced source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1410482-initwithcontentsofurl?language=objc
func XMLDTD_InitWithContentsOfURLOptionsError(url IURL, mask XMLNodeOptions, error IError) XMLDTD {
	return XMLDTDClass.Alloc().InitWithContentsOfURLOptionsError(url, mask, error)
}

func (x_ XMLDTD) Init() XMLDTD {
	rv := objc.Call[XMLDTD](x_, objc.Sel("init"))
	return rv
}

func (xc _XMLDTDClass) Alloc() XMLDTD {
	rv := objc.Call[XMLDTD](xc, objc.Sel("alloc"))
	return rv
}

func XMLDTD_Alloc() XMLDTD {
	return XMLDTDClass.Alloc()
}

func (xc _XMLDTDClass) New() XMLDTD {
	rv := objc.Call[XMLDTD](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLDTD() XMLDTD {
	return XMLDTDClass.New()
}

func (x_ XMLDTD) InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDTD {
	rv := objc.Call[XMLDTD](x_, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

// Returns an NSXMLNode instance initialized with the constant indicating node kind and one or more initialization options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409747-initwithkind?language=objc
func XMLDTD_InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDTD {
	return XMLDTDClass.Alloc().InitWithKindOptions(kind, options)
}

// Returns the DTD node representing the notation declaration identified by the specified notation name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1410533-notationdeclarationforname?language=objc
func (x_ XMLDTD) NotationDeclarationForName(name string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("notationDeclarationForName:"), name)
	return rv
}

// Removes the child node at a particular location in the receiver’s list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1412262-removechildatindex?language=objc
func (x_ XMLDTD) RemoveChildAtIndex(index uint) {
	objc.Call[objc.Void](x_, objc.Sel("removeChildAtIndex:"), index)
}

// Returns the DTD node representing an attribute-list declaration for a given attribute and its element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1411397-attributedeclarationforname?language=objc
func (x_ XMLDTD) AttributeDeclarationForNameElementName(name string, elementName string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("attributeDeclarationForName:elementName:"), name, elementName)
	return rv
}

// Returns the DTD node representing the entity declaration for a specified entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1412325-entitydeclarationforname?language=objc
func (x_ XMLDTD) EntityDeclarationForName(name string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("entityDeclarationForName:"), name)
	return rv
}

// Returns the DTD node representing an element declaration for a specified element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1416067-elementdeclarationforname?language=objc
func (x_ XMLDTD) ElementDeclarationForName(name string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("elementDeclarationForName:"), name)
	return rv
}

// Inserts an array of child nodes at a specified location in the receiver’s list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1413709-insertchildren?language=objc
func (x_ XMLDTD) InsertChildrenAtIndex(children []IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChildren:atIndex:"), children, index)
}

// Inserts a child node in the receiver’s list of children at a specific location in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1411759-insertchild?language=objc
func (x_ XMLDTD) InsertChildAtIndex(child IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChild:atIndex:"), objc.Ptr(child), index)
}

// Removes all existing children of the receiver and replaces them with an array of new child nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1415145-setchildren?language=objc
func (x_ XMLDTD) SetChildren(children []IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("setChildren:"), children)
}

// Adds a child node to the end of the list of existing children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1412988-addchild?language=objc
func (x_ XMLDTD) AddChild(child IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("addChild:"), objc.Ptr(child))
}

// Returns a DTD node representing the predefined entity declaration with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1415434-predefinedentitydeclarationforna?language=objc
func (xc _XMLDTDClass) PredefinedEntityDeclarationForName(name string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](xc, objc.Sel("predefinedEntityDeclarationForName:"), name)
	return rv
}

// Returns a DTD node representing the predefined entity declaration with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1415434-predefinedentitydeclarationforna?language=objc
func XMLDTD_PredefinedEntityDeclarationForName(name string) XMLDTDNode {
	return XMLDTDClass.PredefinedEntityDeclarationForName(name)
}

// Replaces a child at a particular index with another child. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1413890-replacechildatindex?language=objc
func (x_ XMLDTD) ReplaceChildAtIndexWithNode(index uint, node IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("replaceChildAtIndex:withNode:"), index, objc.Ptr(node))
}

// Returns the receiver’s system identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1410949-systemid?language=objc
func (x_ XMLDTD) SystemID() string {
	rv := objc.Call[string](x_, objc.Sel("systemID"))
	return rv
}

// Returns the receiver’s system identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1410949-systemid?language=objc
func (x_ XMLDTD) SetSystemID(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setSystemID:"), value)
}

// Returns the receiver’s public identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1408524-publicid?language=objc
func (x_ XMLDTD) PublicID() string {
	rv := objc.Call[string](x_, objc.Sel("publicID"))
	return rv
}

// Returns the receiver’s public identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtd/1408524-publicid?language=objc
func (x_ XMLDTD) SetPublicID(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setPublicID:"), value)
}
