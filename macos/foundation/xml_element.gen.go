// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLElement] class.
var XMLElementClass = _XMLElementClass{objc.GetClass("NSXMLElement")}

type _XMLElementClass struct {
	objc.Class
}

// An interface definition for the [XMLElement] class.
type IXMLElement interface {
	IXMLNode
	RemoveNamespaceForPrefix(name string)
	NamespaceForPrefix(name string) XMLNode
	AttributeForName(name string) XMLNode
	AddNamespace(aNamespace IXMLNode)
	RemoveChildAtIndex(index uint)
	ElementsForLocalNameURI(localName string, URI string) []XMLElement
	AttributeForLocalNameURI(localName string, URI string) XMLNode
	InsertChildrenAtIndex(children []IXMLNode, index uint)
	RemoveAttributeForName(name string)
	ResolvePrefixForNamespaceURI(namespaceURI string) string
	InsertChildAtIndex(child IXMLNode, index uint)
	ResolveNamespaceForName(name string) XMLNode
	SetChildren(children []IXMLNode)
	AddAttribute(attribute IXMLNode)
	SetAttributesWithDictionary(attributes map[string]string)
	AddChild(child IXMLNode)
	NormalizeAdjacentTextNodesPreservingCDATA(preserve bool)
	ElementsForName(name string) []XMLElement
	ReplaceChildAtIndexWithNode(index uint, node IXMLNode)
	Namespaces() []XMLNode
	SetNamespaces(value []IXMLNode)
	Attributes() []XMLNode
	SetAttributes(value []IXMLNode)
}

// The element nodes in an XML tree structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement?language=objc
type XMLElement struct {
	XMLNode
}

func XMLElementFrom(ptr unsafe.Pointer) XMLElement {
	return XMLElement{
		XMLNode: XMLNodeFrom(ptr),
	}
}

func (x_ XMLElement) InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLElement {
	rv := objc.Call[XMLElement](x_, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388323-initwithkind?language=objc
func NewXMLElementWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLElement {
	instance := XMLElementClass.Alloc().InitWithKindOptions(kind, options)
	instance.Autorelease()
	return instance
}

func (x_ XMLElement) InitWithNameStringValue(name string, string_ string) XMLElement {
	rv := objc.Call[XMLElement](x_, objc.Sel("initWithName:stringValue:"), name, string_)
	return rv
}

// Returns an NSXMLElement object initialized with a specified name and a single text-node child containing a specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388356-initwithname?language=objc
func NewXMLElementWithNameStringValue(name string, string_ string) XMLElement {
	instance := XMLElementClass.Alloc().InitWithNameStringValue(name, string_)
	instance.Autorelease()
	return instance
}

func (x_ XMLElement) InitWithXMLStringError(string_ string, error IError) XMLElement {
	rv := objc.Call[XMLElement](x_, objc.Sel("initWithXMLString:error:"), string_, objc.Ptr(error))
	return rv
}

// Returns an NSXMLElement object created from a specified string containing XML markup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388325-initwithxmlstring?language=objc
func NewXMLElementWithXMLStringError(string_ string, error IError) XMLElement {
	instance := XMLElementClass.Alloc().InitWithXMLStringError(string_, error)
	instance.Autorelease()
	return instance
}

func (xc _XMLElementClass) Alloc() XMLElement {
	rv := objc.Call[XMLElement](xc, objc.Sel("alloc"))
	return rv
}

func XMLElement_Alloc() XMLElement {
	return XMLElementClass.Alloc()
}

func (xc _XMLElementClass) New() XMLElement {
	rv := objc.Call[XMLElement](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLElement() XMLElement {
	return XMLElementClass.New()
}

func (x_ XMLElement) Init() XMLElement {
	rv := objc.Call[XMLElement](x_, objc.Sel("init"))
	return rv
}

// Removes a namespace node that is identified by a given prefix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388327-removenamespaceforprefix?language=objc
func (x_ XMLElement) RemoveNamespaceForPrefix(name string) {
	objc.Call[objc.Void](x_, objc.Sel("removeNamespaceForPrefix:"), name)
}

// Returns the namespace node with a specified prefix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388307-namespaceforprefix?language=objc
func (x_ XMLElement) NamespaceForPrefix(name string) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("namespaceForPrefix:"), name)
	return rv
}

// Returns the attribute node of the receiver with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388305-attributeforname?language=objc
func (x_ XMLElement) AttributeForName(name string) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("attributeForName:"), name)
	return rv
}

// Adds a namespace node to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388358-addnamespace?language=objc
func (x_ XMLElement) AddNamespace(aNamespace IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("addNamespace:"), objc.Ptr(aNamespace))
}

// Removes the child node of the receiver identified by a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388311-removechildatindex?language=objc
func (x_ XMLElement) RemoveChildAtIndex(index uint) {
	objc.Call[objc.Void](x_, objc.Sel("removeChildAtIndex:"), index)
}

// Returns the child element nodes (as NSXMLElement objects) of the receiver that are matched with the specified local name and URI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388344-elementsforlocalname?language=objc
func (x_ XMLElement) ElementsForLocalNameURI(localName string, URI string) []XMLElement {
	rv := objc.Call[[]XMLElement](x_, objc.Sel("elementsForLocalName:URI:"), localName, URI)
	return rv
}

// Returns the attribute node of the receiver that is identified by a local name and URI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388315-attributeforlocalname?language=objc
func (x_ XMLElement) AttributeForLocalNameURI(localName string, URI string) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("attributeForLocalName:URI:"), localName, URI)
	return rv
}

// Inserts an array of child nodes at a specified location in the receiver’s list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388331-insertchildren?language=objc
func (x_ XMLElement) InsertChildrenAtIndex(children []IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChildren:atIndex:"), children, index)
}

// Removes an attribute node identified by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388338-removeattributeforname?language=objc
func (x_ XMLElement) RemoveAttributeForName(name string) {
	objc.Call[objc.Void](x_, objc.Sel("removeAttributeForName:"), name)
}

// Returns the prefix associated with the specified URI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388309-resolveprefixfornamespaceuri?language=objc
func (x_ XMLElement) ResolvePrefixForNamespaceURI(namespaceURI string) string {
	rv := objc.Call[string](x_, objc.Sel("resolvePrefixForNamespaceURI:"), namespaceURI)
	return rv
}

// Inserts a new child node at a specified location in the receiver’s list of child nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388332-insertchild?language=objc
func (x_ XMLElement) InsertChildAtIndex(child IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChild:atIndex:"), objc.Ptr(child), index)
}

// Returns the namespace node with the prefix matching the given qualified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388334-resolvenamespaceforname?language=objc
func (x_ XMLElement) ResolveNamespaceForName(name string) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("resolveNamespaceForName:"), name)
	return rv
}

// Sets all child nodes of the receiver at once, replacing any existing children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388350-setchildren?language=objc
func (x_ XMLElement) SetChildren(children []IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("setChildren:"), children)
}

// Adds an attribute node to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388336-addattribute?language=objc
func (x_ XMLElement) AddAttribute(attribute IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("addAttribute:"), objc.Ptr(attribute))
}

// Sets the attributes of the receiver based on the key-value pairs specified in the passed dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388354-setattributeswithdictionary?language=objc
func (x_ XMLElement) SetAttributesWithDictionary(attributes map[string]string) {
	objc.Call[objc.Void](x_, objc.Sel("setAttributesWithDictionary:"), attributes)
}

// Adds a child node at the end of the receiver’s current list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388340-addchild?language=objc
func (x_ XMLElement) AddChild(child IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("addChild:"), objc.Ptr(child))
}

// Coalesces adjacent text nodes of the receiver that you have explicitly added, optionally including CDATA sections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388346-normalizeadjacenttextnodespreser?language=objc
func (x_ XMLElement) NormalizeAdjacentTextNodesPreservingCDATA(preserve bool) {
	objc.Call[objc.Void](x_, objc.Sel("normalizeAdjacentTextNodesPreservingCDATA:"), preserve)
}

// Returns the child element nodes (as NSXMLElement objects) of the receiver that have a specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388329-elementsforname?language=objc
func (x_ XMLElement) ElementsForName(name string) []XMLElement {
	rv := objc.Call[[]XMLElement](x_, objc.Sel("elementsForName:"), name)
	return rv
}

// Replaces a child node at a specified location with another child node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388317-replacechildatindex?language=objc
func (x_ XMLElement) ReplaceChildAtIndexWithNode(index uint, node IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("replaceChildAtIndex:withNode:"), index, objc.Ptr(node))
}

// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388342-namespaces?language=objc
func (x_ XMLElement) Namespaces() []XMLNode {
	rv := objc.Call[[]XMLNode](x_, objc.Sel("namespaces"))
	return rv
}

// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388342-namespaces?language=objc
func (x_ XMLElement) SetNamespaces(value []IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("setNamespaces:"), value)
}

// Sets all attributes of the receiver at once, replacing any existing attribute nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388321-attributes?language=objc
func (x_ XMLElement) Attributes() []XMLNode {
	rv := objc.Call[[]XMLNode](x_, objc.Sel("attributes"))
	return rv
}

// Sets all attributes of the receiver at once, replacing any existing attribute nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlelement/1388321-attributes?language=objc
func (x_ XMLElement) SetAttributes(value []IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("setAttributes:"), value)
}
