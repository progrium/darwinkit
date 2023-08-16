// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLDocument] class.
var DOMHTMLDocumentClass = _DOMHTMLDocumentClass{objc.GetClass("DOMHTMLDocument")}

type _DOMHTMLDocumentClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLDocument] class.
type IDOMHTMLDocument interface {
	IDOMDocument
	CreateDocumentFragmentWithText(text string) DOMDocumentFragment
	CreateDocumentFragmentWithMarkupStringBaseURL(markupString string, baseURL foundation.IURL) DOMDocumentFragment
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldocument?language=objc
type DOMHTMLDocument struct {
	DOMDocument
}

func DOMHTMLDocumentFrom(ptr unsafe.Pointer) DOMHTMLDocument {
	return DOMHTMLDocument{
		DOMDocument: DOMDocumentFrom(ptr),
	}
}

func (dc _DOMHTMLDocumentClass) Alloc() DOMHTMLDocument {
	rv := objc.Call[DOMHTMLDocument](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLDocument_Alloc() DOMHTMLDocument {
	return DOMHTMLDocumentClass.Alloc()
}

func (dc _DOMHTMLDocumentClass) New() DOMHTMLDocument {
	rv := objc.Call[DOMHTMLDocument](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLDocument() DOMHTMLDocument {
	return DOMHTMLDocumentClass.New()
}

func (d_ DOMHTMLDocument) Init() DOMHTMLDocument {
	rv := objc.Call[DOMHTMLDocument](d_, objc.Sel("init"))
	return rv
}

// Creates a document fragment containing the given text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldocument/1537472-createdocumentfragmentwithtext?language=objc
func (d_ DOMHTMLDocument) CreateDocumentFragmentWithText(text string) DOMDocumentFragment {
	rv := objc.Call[DOMDocumentFragment](d_, objc.Sel("createDocumentFragmentWithText:"), text)
	return rv
}

// Creates a document fragment containing the given HTML markup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldocument/1538060-createdocumentfragmentwithmarkup?language=objc
func (d_ DOMHTMLDocument) CreateDocumentFragmentWithMarkupStringBaseURL(markupString string, baseURL foundation.IURL) DOMDocumentFragment {
	rv := objc.Call[DOMDocumentFragment](d_, objc.Sel("createDocumentFragmentWithMarkupString:baseURL:"), markupString, objc.Ptr(baseURL))
	return rv
}
