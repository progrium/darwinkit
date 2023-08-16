// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface an XML parser uses to inform its delegate about the content of the parsed document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate?language=objc
type PXMLParserDelegate interface {
	// optional
	ParserDidStartDocument(parser XMLParser)
	HasParserDidStartDocument() bool

	// optional
	ParserDidEndDocument(parser XMLParser)
	HasParserDidEndDocument() bool

	// optional
	ParserFoundComment(parser XMLParser, comment string)
	HasParserFoundComment() bool
}

// A delegate implementation builder for the [PXMLParserDelegate] protocol.
type XMLParserDelegate struct {
	_ParserDidStartDocument func(parser XMLParser)
	_ParserDidEndDocument   func(parser XMLParser)
	_ParserFoundComment     func(parser XMLParser, comment string)
}

func (di *XMLParserDelegate) HasParserDidStartDocument() bool {
	return di._ParserDidStartDocument != nil
}

// Sent by the parser object to the delegate when it begins parsing a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1412065-parserdidstartdocument?language=objc
func (di *XMLParserDelegate) SetParserDidStartDocument(f func(parser XMLParser)) {
	di._ParserDidStartDocument = f
}

// Sent by the parser object to the delegate when it begins parsing a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1412065-parserdidstartdocument?language=objc
func (di *XMLParserDelegate) ParserDidStartDocument(parser XMLParser) {
	di._ParserDidStartDocument(parser)
}
func (di *XMLParserDelegate) HasParserDidEndDocument() bool {
	return di._ParserDidEndDocument != nil
}

// Sent by the parser object to the delegate when it has successfully completed parsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1418172-parserdidenddocument?language=objc
func (di *XMLParserDelegate) SetParserDidEndDocument(f func(parser XMLParser)) {
	di._ParserDidEndDocument = f
}

// Sent by the parser object to the delegate when it has successfully completed parsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1418172-parserdidenddocument?language=objc
func (di *XMLParserDelegate) ParserDidEndDocument(parser XMLParser) {
	di._ParserDidEndDocument(parser)
}
func (di *XMLParserDelegate) HasParserFoundComment() bool {
	return di._ParserFoundComment != nil
}

// Sent by a parser object to its delegate when it encounters a comment in the XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1417651-parser?language=objc
func (di *XMLParserDelegate) SetParserFoundComment(f func(parser XMLParser, comment string)) {
	di._ParserFoundComment = f
}

// Sent by a parser object to its delegate when it encounters a comment in the XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1417651-parser?language=objc
func (di *XMLParserDelegate) ParserFoundComment(parser XMLParser, comment string) {
	di._ParserFoundComment(parser, comment)
}

// A concrete type wrapper for the [PXMLParserDelegate] protocol.
type XMLParserDelegateWrapper struct {
	objc.Object
}

func (x_ XMLParserDelegateWrapper) HasParserDidStartDocument() bool {
	return x_.RespondsToSelector(objc.Sel("parserDidStartDocument:"))
}

// Sent by the parser object to the delegate when it begins parsing a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1412065-parserdidstartdocument?language=objc
func (x_ XMLParserDelegateWrapper) ParserDidStartDocument(parser IXMLParser) {
	objc.Call[objc.Void](x_, objc.Sel("parserDidStartDocument:"), objc.Ptr(parser))
}

func (x_ XMLParserDelegateWrapper) HasParserDidEndDocument() bool {
	return x_.RespondsToSelector(objc.Sel("parserDidEndDocument:"))
}

// Sent by the parser object to the delegate when it has successfully completed parsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1418172-parserdidenddocument?language=objc
func (x_ XMLParserDelegateWrapper) ParserDidEndDocument(parser IXMLParser) {
	objc.Call[objc.Void](x_, objc.Sel("parserDidEndDocument:"), objc.Ptr(parser))
}

func (x_ XMLParserDelegateWrapper) HasParserFoundComment() bool {
	return x_.RespondsToSelector(objc.Sel("parser:foundComment:"))
}

// Sent by a parser object to its delegate when it encounters a comment in the XML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate/1417651-parser?language=objc
func (x_ XMLParserDelegateWrapper) ParserFoundComment(parser IXMLParser, comment string) {
	objc.Call[objc.Void](x_, objc.Sel("parser:foundComment:"), objc.Ptr(parser), comment)
}
