// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLDocument] class.
var XMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}

type _XMLDocumentClass struct {
	objc.Class
}

// An interface definition for the [XMLDocument] class.
type IXMLDocument interface {
	IXMLNode
	ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments map[string]string, error IError) objc.Object
	RootElement() XMLElement
	RemoveChildAtIndex(index uint)
	ValidateAndReturnError(error IError) bool
	InsertChildrenAtIndex(children []IXMLNode, index uint)
	ObjectByApplyingXSLTArgumentsError(xslt []byte, arguments map[string]string, error IError) objc.Object
	XMLDataWithOptions(options XMLNodeOptions) []byte
	SetRootElement(root IXMLElement)
	InsertChildAtIndex(child IXMLNode, index uint)
	ObjectByApplyingXSLTAtURLArgumentsError(xsltURL IURL, argument map[string]string, error IError) objc.Object
	SetChildren(children []IXMLNode)
	AddChild(child IXMLNode)
	ReplaceChildAtIndexWithNode(index uint, node IXMLNode)
	MIMEType() string
	SetMIMEType(value string)
	Version() string
	SetVersion(value string)
	CharacterEncoding() string
	SetCharacterEncoding(value string)
	DocumentContentKind() XMLDocumentContentKind
	SetDocumentContentKind(value XMLDocumentContentKind)
	IsStandalone() bool
	SetStandalone(value bool)
	DTD() XMLDTD
	SetDTD(value IXMLDTD)
	XMLData() []byte
}

// An XML document as internalized into a logical tree structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument?language=objc
type XMLDocument struct {
	XMLNode
}

func XMLDocumentFrom(ptr unsafe.Pointer) XMLDocument {
	return XMLDocument{
		XMLNode: XMLNodeFrom(ptr),
	}
}

func (x_ XMLDocument) InitWithDataOptionsError(data []byte, mask XMLNodeOptions, error IError) XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("initWithData:options:error:"), data, mask, objc.Ptr(error))
	return rv
}

// Initializes and returns an NSXMLDocument object created from an NSData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1413086-initwithdata?language=objc
func NewXMLDocumentWithDataOptionsError(data []byte, mask XMLNodeOptions, error IError) XMLDocument {
	instance := XMLDocumentClass.Alloc().InitWithDataOptionsError(data, mask, error)
	instance.Autorelease()
	return instance
}

func (x_ XMLDocument) InitWithContentsOfURLOptionsError(url IURL, mask XMLNodeOptions, error IError) XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("initWithContentsOfURL:options:error:"), objc.Ptr(url), mask, objc.Ptr(error))
	return rv
}

// Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1418467-initwithcontentsofurl?language=objc
func NewXMLDocumentWithContentsOfURLOptionsError(url IURL, mask XMLNodeOptions, error IError) XMLDocument {
	instance := XMLDocumentClass.Alloc().InitWithContentsOfURLOptionsError(url, mask, error)
	instance.Autorelease()
	return instance
}

func (x_ XMLDocument) InitWithXMLStringOptionsError(string_ string, mask XMLNodeOptions, error IError) XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("initWithXMLString:options:error:"), string_, mask, objc.Ptr(error))
	return rv
}

// Initializes and returns an NSXMLDocument object created from a string containing XML markup text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1416228-initwithxmlstring?language=objc
func NewXMLDocumentWithXMLStringOptionsError(string_ string, mask XMLNodeOptions, error IError) XMLDocument {
	instance := XMLDocumentClass.Alloc().InitWithXMLStringOptionsError(string_, mask, error)
	instance.Autorelease()
	return instance
}

func (x_ XMLDocument) InitWithRootElement(element IXMLElement) XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("initWithRootElement:"), objc.Ptr(element))
	return rv
}

// Returns an NSXMLDocument object initialized with a single child, the root element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409062-initwithrootelement?language=objc
func NewXMLDocumentWithRootElement(element IXMLElement) XMLDocument {
	instance := XMLDocumentClass.Alloc().InitWithRootElement(element)
	instance.Autorelease()
	return instance
}

func (x_ XMLDocument) Init() XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("init"))
	return rv
}

func (xc _XMLDocumentClass) Alloc() XMLDocument {
	rv := objc.Call[XMLDocument](xc, objc.Sel("alloc"))
	return rv
}

func XMLDocument_Alloc() XMLDocument {
	return XMLDocumentClass.Alloc()
}

func (xc _XMLDocumentClass) New() XMLDocument {
	rv := objc.Call[XMLDocument](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLDocument() XMLDocument {
	return XMLDocumentClass.New()
}

func (x_ XMLDocument) InitWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDocument {
	rv := objc.Call[XMLDocument](x_, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

// Returns an NSXMLNode instance initialized with the constant indicating node kind and one or more initialization options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlnode/1409747-initwithkind?language=objc
func NewXMLDocumentWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDocument {
	instance := XMLDocumentClass.Alloc().InitWithKindOptions(kind, options)
	instance.Autorelease()
	return instance
}

// Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409737-objectbyapplyingxsltstring?language=objc
func (x_ XMLDocument) ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments map[string]string, error IError) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("objectByApplyingXSLTString:arguments:error:"), xslt, arguments, objc.Ptr(error))
	return rv
}

// Returns the root element of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1411693-rootelement?language=objc
func (x_ XMLDocument) RootElement() XMLElement {
	rv := objc.Call[XMLElement](x_, objc.Sel("rootElement"))
	return rv
}

// Removes the child node of the receiver located at a specified position in its array of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1410060-removechildatindex?language=objc
func (x_ XMLDocument) RemoveChildAtIndex(index uint) {
	objc.Call[objc.Void](x_, objc.Sel("removeChildAtIndex:"), index)
}

// Validates the document against the governing schema and returns whether the document conforms to the schema. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1408561-validateandreturnerror?language=objc
func (x_ XMLDocument) ValidateAndReturnError(error IError) bool {
	rv := objc.Call[bool](x_, objc.Sel("validateAndReturnError:"), objc.Ptr(error))
	return rv
}

// Inserts an array of children at a specified position in the receiver’s array of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1415612-insertchildren?language=objc
func (x_ XMLDocument) InsertChildrenAtIndex(children []IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChildren:atIndex:"), children, index)
}

// Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1415843-objectbyapplyingxslt?language=objc
func (x_ XMLDocument) ObjectByApplyingXSLTArgumentsError(xslt []byte, arguments map[string]string, error IError) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("objectByApplyingXSLT:arguments:error:"), xslt, arguments, objc.Ptr(error))
	return rv
}

// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1412424-xmldatawithoptions?language=objc
func (x_ XMLDocument) XMLDataWithOptions(options XMLNodeOptions) []byte {
	rv := objc.Call[[]byte](x_, objc.Sel("XMLDataWithOptions:"), options)
	return rv
}

// Set the root element of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1415610-setrootelement?language=objc
func (x_ XMLDocument) SetRootElement(root IXMLElement) {
	objc.Call[objc.Void](x_, objc.Sel("setRootElement:"), objc.Ptr(root))
}

// Inserts a node object at specified position in the receiver’s array of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1413912-insertchild?language=objc
func (x_ XMLDocument) InsertChildAtIndex(child IXMLNode, index uint) {
	objc.Call[objc.Void](x_, objc.Sel("insertChild:atIndex:"), objc.Ptr(child), index)
}

// Applies the XSLT pattern rules and templates located at a specified URL to the receiver and returns a document object containing transformed XML markup or an NSData object containing plain text, RTF text, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1410606-objectbyapplyingxsltaturl?language=objc
func (x_ XMLDocument) ObjectByApplyingXSLTAtURLArgumentsError(xsltURL IURL, argument map[string]string, error IError) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("objectByApplyingXSLTAtURL:arguments:error:"), objc.Ptr(xsltURL), argument, objc.Ptr(error))
	return rv
}

// Sets the child nodes of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1407550-setchildren?language=objc
func (x_ XMLDocument) SetChildren(children []IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("setChildren:"), children)
}

// Overridden by subclasses to substitute a custom class for an NSXML class that the parser uses to create node instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409660-replacementclassforclass?language=objc
func (xc _XMLDocumentClass) ReplacementClassForClass(cls objc.IClass) objc.Class {
	rv := objc.Call[objc.Class](xc, objc.Sel("replacementClassForClass:"), objc.Ptr(cls))
	return rv
}

// Overridden by subclasses to substitute a custom class for an NSXML class that the parser uses to create node instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409660-replacementclassforclass?language=objc
func XMLDocument_ReplacementClassForClass(cls objc.IClass) objc.Class {
	return XMLDocumentClass.ReplacementClassForClass(cls)
}

// Adds a child node after the last of the receiver’s existing children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1415740-addchild?language=objc
func (x_ XMLDocument) AddChild(child IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("addChild:"), objc.Ptr(child))
}

// Replaces the child node of the receiver located at a specified position in its array of children with another node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1416857-replacechildatindex?language=objc
func (x_ XMLDocument) ReplaceChildAtIndexWithNode(index uint, node IXMLNode) {
	objc.Call[objc.Void](x_, objc.Sel("replaceChildAtIndex:withNode:"), index, objc.Ptr(node))
}

// Returns the MIME type for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1408633-mimetype?language=objc
func (x_ XMLDocument) MIMEType() string {
	rv := objc.Call[string](x_, objc.Sel("MIMEType"))
	return rv
}

// Returns the MIME type for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1408633-mimetype?language=objc
func (x_ XMLDocument) SetMIMEType(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setMIMEType:"), value)
}

// Sets the version of the receiver’s XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409066-version?language=objc
func (x_ XMLDocument) Version() string {
	rv := objc.Call[string](x_, objc.Sel("version"))
	return rv
}

// Sets the version of the receiver’s XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1409066-version?language=objc
func (x_ XMLDocument) SetVersion(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setVersion:"), value)
}

// Sets the character encoding of the receiver to encoding, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1410987-characterencoding?language=objc
func (x_ XMLDocument) CharacterEncoding() string {
	rv := objc.Call[string](x_, objc.Sel("characterEncoding"))
	return rv
}

// Sets the character encoding of the receiver to encoding, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1410987-characterencoding?language=objc
func (x_ XMLDocument) SetCharacterEncoding(value string) {
	objc.Call[objc.Void](x_, objc.Sel("setCharacterEncoding:"), value)
}

// Sets the kind of output content for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1407426-documentcontentkind?language=objc
func (x_ XMLDocument) DocumentContentKind() XMLDocumentContentKind {
	rv := objc.Call[XMLDocumentContentKind](x_, objc.Sel("documentContentKind"))
	return rv
}

// Sets the kind of output content for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1407426-documentcontentkind?language=objc
func (x_ XMLDocument) SetDocumentContentKind(value XMLDocumentContentKind) {
	objc.Call[objc.Void](x_, objc.Sel("setDocumentContentKind:"), value)
}

// Sets a Boolean value that specifies whether the receiver represents a standalone XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1413655-standalone?language=objc
func (x_ XMLDocument) IsStandalone() bool {
	rv := objc.Call[bool](x_, objc.Sel("isStandalone"))
	return rv
}

// Sets a Boolean value that specifies whether the receiver represents a standalone XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1413655-standalone?language=objc
func (x_ XMLDocument) SetStandalone(value bool) {
	objc.Call[objc.Void](x_, objc.Sel("setStandalone:"), value)
}

// Returns an NSXMLDTD object representing the internal DTD associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1418474-dtd?language=objc
func (x_ XMLDocument) DTD() XMLDTD {
	rv := objc.Call[XMLDTD](x_, objc.Sel("DTD"))
	return rv
}

// Returns an NSXMLDTD object representing the internal DTD associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1418474-dtd?language=objc
func (x_ XMLDocument) SetDTD(value IXMLDTD) {
	objc.Call[objc.Void](x_, objc.Sel("setDTD:"), objc.Ptr(value))
}

// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmldocument/1411660-xmldata?language=objc
func (x_ XMLDocument) XMLData() []byte {
	rv := objc.Call[[]byte](x_, objc.Sel("XMLData"))
	return rv
}
