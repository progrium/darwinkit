// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLDTDNode] class.
var XMLDTDNodeClass = _XMLDTDNodeClass{objc.GetClass("NSXMLDTDNode")}

type _XMLDTDNodeClass struct {
	objc.Class
}

// An interface definition for the [XMLDTDNode] class.
type IXMLDTDNode interface {
	IXMLNode
	DTDKind() XMLDTDNodeKind
	SetDTDKind(value XMLDTDNodeKind)
	SystemID() string
	SetSystemID(value string)
	PublicID() string
	SetPublicID(value string)
	NotationName() string
	SetNotationName(value string)
	IsExternal() bool
}

// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode?language=objc
type XMLDTDNode struct {
	XMLNode
}

func XMLDTDNodeFrom(ptr unsafe.Pointer) XMLDTDNode {
	return XMLDTDNode{
		XMLNode: XMLNodeFrom(ptr),
	}
}

func (x_ XMLDTDNode) InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1408553-initwithkind?language=objc
func XMLDTDNode_InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDTDNode {
	return XMLDTDNodeClass.Alloc().InitWithKindOptions(kind, options)
}

func (x_ XMLDTDNode) InitWithXMLString(string_ string) XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("initWithXMLString:"), string_)
	return rv
}

// Returns an NSXMLDTDNode object initialized with the DTD declaration in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1409605-initwithxmlstring?language=objc
func XMLDTDNode_InitWithXMLString(string_ string) XMLDTDNode {
	return XMLDTDNodeClass.Alloc().InitWithXMLString(string_)
}

func (x_ XMLDTDNode) Init() XMLDTDNode {
	rv := objc.Call[XMLDTDNode](x_, objc.Sel("init"))
	return rv
}

func (xc _XMLDTDNodeClass) Alloc() XMLDTDNode {
	rv := objc.Call[XMLDTDNode](xc, objc.Sel("alloc"))
	return rv
}

func XMLDTDNode_Alloc() XMLDTDNode {
	return XMLDTDNodeClass.Alloc()
}

func (xc _XMLDTDNodeClass) New() XMLDTDNode {
	rv := objc.Call[XMLDTDNode](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLDTDNode() XMLDTDNode {
	return XMLDTDNodeClass.New()
}

// Returns the receiver’s DTD kind. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1408902-dtdkind?language=objc
func (x_ XMLDTDNode) DTDKind() XMLDTDNodeKind {
	rv := objc.Call[XMLDTDNodeKind](x_, objc.Sel("DTDKind"))
	return rv
}

// Returns the receiver’s DTD kind. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1408902-dtdkind?language=objc
func (x_ XMLDTDNode) SetDTDKind(value XMLDTDNodeKind) {
	objc.Call[objc.Void](x_, objc.Sel("setDTDKind:"), value)
}

// Returns the system identifier associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1410930-systemid?language=objc
func (x_ XMLDTDNode) SystemID() string {
	rv := objc.Call[string](x_, objc.Sel("systemID"))
	return rv
}

// Returns the system identifier associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1410930-systemid?language=objc
func (x_ XMLDTDNode) SetSystemID(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setSystemID:"), value)
}

// Returns the public identifier associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1415631-publicid?language=objc
func (x_ XMLDTDNode) PublicID() string {
	rv := objc.Call[string](x_, objc.Sel("publicID"))
	return rv
}

// Returns the public identifier associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1415631-publicid?language=objc
func (x_ XMLDTDNode) SetPublicID(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setPublicID:"), value)
}

// Returns the name of the notation associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1407292-notationname?language=objc
func (x_ XMLDTDNode) NotationName() string {
	rv := objc.Call[string](x_, objc.Sel("notationName"))
	return rv
}

// Returns the name of the notation associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1407292-notationname?language=objc
func (x_ XMLDTDNode) SetNotationName(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setNotationName:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldtdnode/1409467-external?language=objc
func (x_ XMLDTDNode) IsExternal() bool {
	rv := objc.Call[bool](x_, objc.Sel("isExternal"))
	return rv
}
