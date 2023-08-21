// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XMLParser] class.
var XMLParserClass = _XMLParserClass{objc.GetClass("NSXMLParser")}

type _XMLParserClass struct {
	objc.Class
}

// An interface definition for the [XMLParser] class.
type IXMLParser interface {
	objc.IObject
	AbortParsing()
	Parse() bool
	ColumnNumber() int
	ExternalEntityResolvingPolicy() XMLParserExternalEntityResolvingPolicy
	SetExternalEntityResolvingPolicy(value XMLParserExternalEntityResolvingPolicy)
	SystemID() string
	ShouldProcessNamespaces() bool
	SetShouldProcessNamespaces(value bool)
	ShouldReportNamespacePrefixes() bool
	SetShouldReportNamespacePrefixes(value bool)
	Delegate() XMLParserDelegateWrapper
	SetDelegate(value PXMLParserDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PublicID() string
	AllowedExternalEntityURLs() Set
	SetAllowedExternalEntityURLs(value ISet)
	ParserError() Error
	ShouldResolveExternalEntities() bool
	SetShouldResolveExternalEntities(value bool)
	LineNumber() int
}

// An event driven parser of XML documents (including DTD declarations). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser?language=objc
type XMLParser struct {
	objc.Object
}

func XMLParserFrom(ptr unsafe.Pointer) XMLParser {
	return XMLParser{
		Object: objc.ObjectFrom(ptr),
	}
}

func (x_ XMLParser) InitWithData(data []byte) XMLParser {
	rv := objc.Call[XMLParser](x_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a parser with the XML contents encapsulated in a given data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1418103-initwithdata?language=objc
func NewXMLParserWithData(data []byte) XMLParser {
	instance := XMLParserClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (x_ XMLParser) InitWithContentsOfURL(url IURL) XMLParser {
	rv := objc.Call[XMLParser](x_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes a parser with the XML content referenced by the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1415575-initwithcontentsofurl?language=objc
func NewXMLParserWithContentsOfURL(url IURL) XMLParser {
	instance := XMLParserClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (x_ XMLParser) InitWithStream(stream IInputStream) XMLParser {
	rv := objc.Call[XMLParser](x_, objc.Sel("initWithStream:"), objc.Ptr(stream))
	return rv
}

// Initializes a parser with the XML contents from the specified stream and parses it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1415904-initwithstream?language=objc
func NewXMLParserWithStream(stream IInputStream) XMLParser {
	instance := XMLParserClass.Alloc().InitWithStream(stream)
	instance.Autorelease()
	return instance
}

func (xc _XMLParserClass) Alloc() XMLParser {
	rv := objc.Call[XMLParser](xc, objc.Sel("alloc"))
	return rv
}

func XMLParser_Alloc() XMLParser {
	return XMLParserClass.Alloc()
}

func (xc _XMLParserClass) New() XMLParser {
	rv := objc.Call[XMLParser](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXMLParser() XMLParser {
	return XMLParserClass.New()
}

func (x_ XMLParser) Init() XMLParser {
	rv := objc.Call[XMLParser](x_, objc.Sel("init"))
	return rv
}

// Stops the parser object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1410083-abortparsing?language=objc
func (x_ XMLParser) AbortParsing() {
	objc.Call[objc.Void](x_, objc.Sel("abortParsing"))
}

// Starts the event-driven parsing operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1411778-parse?language=objc
func (x_ XMLParser) Parse() bool {
	rv := objc.Call[bool](x_, objc.Sel("parse"))
	return rv
}

// The column number of the XML document being processed by the parser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1416983-columnnumber?language=objc
func (x_ XMLParser) ColumnNumber() int {
	rv := objc.Call[int](x_, objc.Sel("columnNumber"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1407399-externalentityresolvingpolicy?language=objc
func (x_ XMLParser) ExternalEntityResolvingPolicy() XMLParserExternalEntityResolvingPolicy {
	rv := objc.Call[XMLParserExternalEntityResolvingPolicy](x_, objc.Sel("externalEntityResolvingPolicy"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1407399-externalentityresolvingpolicy?language=objc
func (x_ XMLParser) SetExternalEntityResolvingPolicy(value XMLParserExternalEntityResolvingPolicy) {
	objc.Call[objc.Void](x_, objc.Sel("setExternalEntityResolvingPolicy:"), value)
}

// The system identifier of the external entity referenced in the XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1411917-systemid?language=objc
func (x_ XMLParser) SystemID() string {
	rv := objc.Call[string](x_, objc.Sel("systemID"))
	return rv
}

// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1418380-shouldprocessnamespaces?language=objc
func (x_ XMLParser) ShouldProcessNamespaces() bool {
	rv := objc.Call[bool](x_, objc.Sel("shouldProcessNamespaces"))
	return rv
}

// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1418380-shouldprocessnamespaces?language=objc
func (x_ XMLParser) SetShouldProcessNamespaces(value bool) {
	objc.Call[objc.Void](x_, objc.Sel("setShouldProcessNamespaces:"), value)
}

// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1410809-shouldreportnamespaceprefixes?language=objc
func (x_ XMLParser) ShouldReportNamespacePrefixes() bool {
	rv := objc.Call[bool](x_, objc.Sel("shouldReportNamespacePrefixes"))
	return rv
}

// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1410809-shouldreportnamespaceprefixes?language=objc
func (x_ XMLParser) SetShouldReportNamespacePrefixes(value bool) {
	objc.Call[objc.Void](x_, objc.Sel("setShouldReportNamespacePrefixes:"), value)
}

// A delegate object that receives messages about the parsing process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1416209-delegate?language=objc
func (x_ XMLParser) Delegate() XMLParserDelegateWrapper {
	rv := objc.Call[XMLParserDelegateWrapper](x_, objc.Sel("delegate"))
	return rv
}

// A delegate object that receives messages about the parsing process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1416209-delegate?language=objc
func (x_ XMLParser) SetDelegate(value PXMLParserDelegate) {
	po0 := objc.WrapAsProtocol("NSXMLParserDelegate", value)
	objc.Call[objc.Void](x_, objc.Sel("setDelegate:"), po0)
}

// A delegate object that receives messages about the parsing process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1416209-delegate?language=objc
func (x_ XMLParser) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](x_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The public identifier of the external entity referenced in the XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1414516-publicid?language=objc
func (x_ XMLParser) PublicID() string {
	rv := objc.Call[string](x_, objc.Sel("publicID"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1412380-allowedexternalentityurls?language=objc
func (x_ XMLParser) AllowedExternalEntityURLs() Set {
	rv := objc.Call[Set](x_, objc.Sel("allowedExternalEntityURLs"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1412380-allowedexternalentityurls?language=objc
func (x_ XMLParser) SetAllowedExternalEntityURLs(value ISet) {
	objc.Call[objc.Void](x_, objc.Sel("setAllowedExternalEntityURLs:"), objc.Ptr(value))
}

// An NSError object from which you can obtain information about a parsing error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1417446-parsererror?language=objc
func (x_ XMLParser) ParserError() Error {
	rv := objc.Call[Error](x_, objc.Sel("parserError"))
	return rv
}

// A Boolean value that determines whether the parser reports declarations of external entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1414143-shouldresolveexternalentities?language=objc
func (x_ XMLParser) ShouldResolveExternalEntities() bool {
	rv := objc.Call[bool](x_, objc.Sel("shouldResolveExternalEntities"))
	return rv
}

// A Boolean value that determines whether the parser reports declarations of external entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1414143-shouldresolveexternalentities?language=objc
func (x_ XMLParser) SetShouldResolveExternalEntities(value bool) {
	objc.Call[objc.Void](x_, objc.Sel("setShouldResolveExternalEntities:"), value)
}

// The line number of the XML document being processed by the parser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparser/1413404-linenumber?language=objc
func (x_ XMLParser) LineNumber() int {
	rv := objc.Call[int](x_, objc.Sel("lineNumber"))
	return rv
}
