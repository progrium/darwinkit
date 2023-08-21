// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLNode] class.
var XMLNodeClass = _XMLNodeClass{objc.GetClass("NSXMLNode")}

type _XMLNodeClass struct {
	objc.Class
}

// An interface definition for the [XMLNode] class.
type IXMLNode interface {
	objc.IObject
	CanonicalXMLStringPreservingComments(comments bool) string
	NodesForXPathError(xpath string, error IError) []XMLNode
	Detach()
	XMLStringWithOptions(options XMLNodeOptions) string
	SetStringValueResolvingEntities(string_ string, resolve bool)
	ChildAtIndex(index uint) XMLNode
	ObjectsForXQueryError(xquery string, error IError) []objc.Object
	RootDocument() XMLDocument
	NextNode() XMLNode
	Level() uint
	URI() string
	SetURI(value string)
	Parent() XMLNode
	Name() string
	SetName(value string)
	LocalName() string
	StringValue() string
	SetStringValue(value string)
	XMLString() string
	PreviousSibling() XMLNode
	Description() string
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	Kind() XMLNodeKind
	XPath() string
	NextSibling() XMLNode
	Prefix() string
	Index() uint
	ChildCount() uint
	PreviousNode() XMLNode
	Children() []XMLNode
}

// The nodes in the abstract, logical tree structure that represents an XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode?language=objc
type XMLNode struct {
	objc.Object
}

func XMLNodeFrom(ptr unsafe.Pointer) XMLNode {
	return XMLNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (x_ XMLNode) InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

// Returns an NSXMLNode instance initialized with the constant indicating node kind and one or more initialization options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409747-initwithkind?language=objc
func NewXMLNodeWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLNode {
	instance := XMLNodeClass.Alloc().InitWithKindOptions(kind, options)
	instance.Autorelease()
	return instance
}

func (x_ XMLNode) Init() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("init"))
	return rv
}

func (xc _XMLNodeClass) Alloc() XMLNode {
	rv := objc.Call[XMLNode](xc, objc.Sel("alloc"))
	return rv
}

func XMLNode_Alloc() XMLNode {
	return XMLNodeClass.Alloc()
}

func (xc _XMLNodeClass) New() XMLNode {
	rv := objc.Call[XMLNode](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLNode() XMLNode {
	return XMLNodeClass.New()
}

// Returns an NSXMLElement object with a given tag identifier, or name [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409760-elementwithname?language=objc
func (xc _XMLNodeClass) ElementWithName(name string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("elementWithName:"), name)
	return rv
}

// Returns an NSXMLElement object with a given tag identifier, or name [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409760-elementwithname?language=objc
func XMLNode_ElementWithName(name string) objc.Object {
	return XMLNodeClass.ElementWithName(name)
}

// Returns a string object encapsulating the receiver’s XML in canonical form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409817-canonicalxmlstringpreservingcomm?language=objc
func (x_ XMLNode) CanonicalXMLStringPreservingComments(comments bool) string {
	rv := objc.Call[string](x_, objc.Sel("canonicalXMLStringPreservingComments:"), comments)
	return rv
}

// Returns an NSXMLNode object representing a processing instruction with a specified name and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409784-processinginstructionwithname?language=objc
func (xc _XMLNodeClass) ProcessingInstructionWithNameStringValue(name string, stringValue string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("processingInstructionWithName:stringValue:"), name, stringValue)
	return rv
}

// Returns an NSXMLNode object representing a processing instruction with a specified name and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409784-processinginstructionwithname?language=objc
func XMLNode_ProcessingInstructionWithNameStringValue(name string, stringValue string) objc.Object {
	return XMLNodeClass.ProcessingInstructionWithNameStringValue(name, stringValue)
}

// Returns the nodes resulting from executing an XPath query upon the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409813-nodesforxpath?language=objc
func (x_ XMLNode) NodesForXPathError(xpath string, error IError) []XMLNode {
	rv := objc.Call[[]XMLNode](x_, objc.Sel("nodesForXPath:error:"), xpath, objc.Ptr(error))
	return rv
}

// Returns the prefix from the specified qualified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1407393-prefixforname?language=objc
func (xc _XMLNodeClass) PrefixForName(name string) string {
	rv := objc.Call[string](xc, objc.Sel("prefixForName:"), name)
	return rv
}

// Returns the prefix from the specified qualified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1407393-prefixforname?language=objc
func XMLNode_PrefixForName(name string) string {
	return XMLNodeClass.PrefixForName(name)
}

// Returns a NSXMLDTDNode object representing the DTD declaration for an element, attribute, entity, or notation based on a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409145-dtdnodewithxmlstring?language=objc
func (xc _XMLNodeClass) DTDNodeWithXMLString(string_ string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("DTDNodeWithXMLString:"), string_)
	return rv
}

// Returns a NSXMLDTDNode object representing the DTD declaration for an element, attribute, entity, or notation based on a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409145-dtdnodewithxmlstring?language=objc
func XMLNode_DTDNodeWithXMLString(string_ string) objc.Object {
	return XMLNodeClass.DTDNodeWithXMLString(string_)
}

// Detaches the receiver from its parent node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409806-detach?language=objc
func (x_ XMLNode) Detach() {
	objc.Call[objc.Void](x_, objc.Sel("detach"))
}

// Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409778-xmlstringwithoptions?language=objc
func (x_ XMLNode) XMLStringWithOptions(options XMLNodeOptions) string {
	rv := objc.Call[string](x_, objc.Sel("XMLStringWithOptions:"), options)
	return rv
}

// Returns an NSXMLNode object representing a namespace with a specified name and URI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409738-namespacewithname?language=objc
func (xc _XMLNodeClass) NamespaceWithNameStringValue(name string, stringValue string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("namespaceWithName:stringValue:"), name, stringValue)
	return rv
}

// Returns an NSXMLNode object representing a namespace with a specified name and URI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409738-namespacewithname?language=objc
func XMLNode_NamespaceWithNameStringValue(name string, stringValue string) objc.Object {
	return XMLNodeClass.NamespaceWithNameStringValue(name, stringValue)
}

// Returns an NSXMLNode object representing a comment node containing given text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409831-commentwithstringvalue?language=objc
func (xc _XMLNodeClass) CommentWithStringValue(stringValue string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("commentWithStringValue:"), stringValue)
	return rv
}

// Returns an NSXMLNode object representing a comment node containing given text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409831-commentwithstringvalue?language=objc
func XMLNode_CommentWithStringValue(stringValue string) objc.Object {
	return XMLNodeClass.CommentWithStringValue(stringValue)
}

// Returns an empty document node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409821-document?language=objc
func (xc _XMLNodeClass) Document() objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("document"))
	return rv
}

// Returns an empty document node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409821-document?language=objc
func XMLNode_Document() objc.Object {
	return XMLNodeClass.Document()
}

// Returns an NSXMLNode object representing a text node with specified content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409754-textwithstringvalue?language=objc
func (xc _XMLNodeClass) TextWithStringValue(stringValue string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("textWithStringValue:"), stringValue)
	return rv
}

// Returns an NSXMLNode object representing a text node with specified content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409754-textwithstringvalue?language=objc
func XMLNode_TextWithStringValue(stringValue string) objc.Object {
	return XMLNodeClass.TextWithStringValue(stringValue)
}

// Returns an NSXMLNode object representing one of the predefined namespaces with the specified prefix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409745-predefinednamespaceforprefix?language=objc
func (xc _XMLNodeClass) PredefinedNamespaceForPrefix(name string) XMLNode {
	rv := objc.Call[XMLNode](xc, objc.Sel("predefinedNamespaceForPrefix:"), name)
	return rv
}

// Returns an NSXMLNode object representing one of the predefined namespaces with the specified prefix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409745-predefinednamespaceforprefix?language=objc
func XMLNode_PredefinedNamespaceForPrefix(name string) XMLNode {
	return XMLNodeClass.PredefinedNamespaceForPrefix(name)
}

// Returns the local name from the specified qualified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409646-localnameforname?language=objc
func (xc _XMLNodeClass) LocalNameForName(name string) string {
	rv := objc.Call[string](xc, objc.Sel("localNameForName:"), name)
	return rv
}

// Returns the local name from the specified qualified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409646-localnameforname?language=objc
func XMLNode_LocalNameForName(name string) string {
	return XMLNodeClass.LocalNameForName(name)
}

// Returns an NSXMLDocument object initialized with a given root element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409802-documentwithrootelement?language=objc
func (xc _XMLNodeClass) DocumentWithRootElement(element IXMLElement) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("documentWithRootElement:"), objc.Ptr(element))
	return rv
}

// Returns an NSXMLDocument object initialized with a given root element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409802-documentwithrootelement?language=objc
func XMLNode_DocumentWithRootElement(element IXMLElement) objc.Object {
	return XMLNodeClass.DocumentWithRootElement(element)
}

// Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409758-setstringvalue?language=objc
func (x_ XMLNode) SetStringValueResolvingEntities(string_ string, resolve bool) {
	objc.Call[objc.Void](x_, objc.Sel("setStringValue:resolvingEntities:"), string_, resolve)
}

// Returns an NSXMLNode object representing an attribute node with a given name and string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409795-attributewithname?language=objc
func (xc _XMLNodeClass) AttributeWithNameStringValue(name string, stringValue string) objc.Object {
	rv := objc.Call[objc.Object](xc, objc.Sel("attributeWithName:stringValue:"), name, stringValue)
	return rv
}

// Returns an NSXMLNode object representing an attribute node with a given name and string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409795-attributewithname?language=objc
func XMLNode_AttributeWithNameStringValue(name string, stringValue string) objc.Object {
	return XMLNodeClass.AttributeWithNameStringValue(name, stringValue)
}

// Returns the child node of the receiver at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409811-childatindex?language=objc
func (x_ XMLNode) ChildAtIndex(index uint) XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("childAtIndex:"), index)
	return rv
}

// Returns the objects resulting from executing an XQuery query upon the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409768-objectsforxquery?language=objc
func (x_ XMLNode) ObjectsForXQueryError(xquery string, error IError) []objc.Object {
	rv := objc.Call[[]objc.Object](x_, objc.Sel("objectsForXQuery:error:"), xquery, objc.Ptr(error))
	return rv
}

// Returns the NSXMLDocument object containing the root element and representing the XML document as a whole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409830-rootdocument?language=objc
func (x_ XMLNode) RootDocument() XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("rootDocument"))
	return rv
}

// Returns the next NSXMLNode object in document order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409800-nextnode?language=objc
func (x_ XMLNode) NextNode() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("nextNode"))
	return rv
}

// Returns the nesting level of the receiver within the tree hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1407508-level?language=objc
func (x_ XMLNode) Level() uint {
	rv := objc.Call[uint](x_, objc.Sel("level"))
	return rv
}

// Returns the URI associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409774-uri?language=objc
func (x_ XMLNode) URI() string {
	rv := objc.Call[string](x_, objc.Sel("URI"))
	return rv
}

// Returns the URI associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409774-uri?language=objc
func (x_ XMLNode) SetURI(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setURI:"), value)
}

// Returns the parent node of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409740-parent?language=objc
func (x_ XMLNode) Parent() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("parent"))
	return rv
}

// Returns the name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409770-name?language=objc
func (x_ XMLNode) Name() string {
	rv := objc.Call[string](x_, objc.Sel("name"))
	return rv
}

// Returns the name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409770-name?language=objc
func (x_ XMLNode) SetName(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setName:"), value)
}

// Returns the local name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409734-localname?language=objc
func (x_ XMLNode) LocalName() string {
	rv := objc.Call[string](x_, objc.Sel("localName"))
	return rv
}

// Returns the content of the receiver as a string value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409818-stringvalue?language=objc
func (x_ XMLNode) StringValue() string {
	rv := objc.Call[string](x_, objc.Sel("stringValue"))
	return rv
}

// Returns the content of the receiver as a string value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409818-stringvalue?language=objc
func (x_ XMLNode) SetStringValue(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setStringValue:"), value)
}

// Returns the string representation of the receiver as it would appear in an XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409772-xmlstring?language=objc
func (x_ XMLNode) XMLString() string {
	rv := objc.Call[string](x_, objc.Sel("XMLString"))
	return rv
}

// Returns the previous NSXMLNode object that is a sibling node to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409764-previoussibling?language=objc
func (x_ XMLNode) PreviousSibling() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("previousSibling"))
	return rv
}

// Returns a description of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409788-description?language=objc
func (x_ XMLNode) Description() string {
	rv := objc.Call[string](x_, objc.Sel("description"))
	return rv
}

// Returns the object value of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409756-objectvalue?language=objc
func (x_ XMLNode) ObjectValue() objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("objectValue"))
	return rv
}

// Returns the object value of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409756-objectvalue?language=objc
func (x_ XMLNode) SetObjectValue(value objc.IObject) {
	objc.Call[objc.Void](x_, objc.Sel("setObjectValue:"), value)
}

// Returns the kind of node the receiver is as a constant of type NSXMLNodeKind. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1408882-kind?language=objc
func (x_ XMLNode) Kind() XMLNodeKind {
	rv := objc.Call[XMLNodeKind](x_, objc.Sel("kind"))
	return rv
}

// Returns the XPath expression identifying the receiver’s location in the document tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409826-xpath?language=objc
func (x_ XMLNode) XPath() string {
	rv := objc.Call[string](x_, objc.Sel("XPath"))
	return rv
}

// Returns the next NSXMLNode object that is a sibling node to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409750-nextsibling?language=objc
func (x_ XMLNode) NextSibling() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("nextSibling"))
	return rv
}

// Returns the prefix of the receiver’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409797-prefix?language=objc
func (x_ XMLNode) Prefix() string {
	rv := objc.Call[string](x_, objc.Sel("prefix"))
	return rv
}

// Returns the index of the receiver identifying its position relative to its sibling nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409743-index?language=objc
func (x_ XMLNode) Index() uint {
	rv := objc.Call[uint](x_, objc.Sel("index"))
	return rv
}

// Returns the number of child nodes the receiver has. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409808-childcount?language=objc
func (x_ XMLNode) ChildCount() uint {
	rv := objc.Call[uint](x_, objc.Sel("childCount"))
	return rv
}

// Returns the previous NSXMLNode object in document order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409736-previousnode?language=objc
func (x_ XMLNode) PreviousNode() XMLNode {
	rv := objc.Call[XMLNode](x_, objc.Sel("previousNode"))
	return rv
}

// Returns an immutable array containing the child nodes of the receiver (as NSXMLNode objects). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409828-children?language=objc
func (x_ XMLNode) Children() []XMLNode {
	rv := objc.Call[[]XMLNode](x_, objc.Sel("children"))
	return rv
}
