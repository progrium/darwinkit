// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchQuery] class.
var SearchQueryClass = _SearchQueryClass{objc.GetClass("CSSearchQuery")}

type _SearchQueryClass struct {
	objc.Class
}

// An interface definition for the [SearchQuery] class.
type ISearchQuery interface {
	objc.IObject
	Start()
	Cancel()
	IsCancelled() bool
	FoundItemCount() uint
	CompletionHandler() func(error foundation.Error)
	SetCompletionHandler(value func(error foundation.Error))
	FoundItemsHandler() func(items []SearchableItem)
	SetFoundItemsHandler(value func(items []SearchableItem))
	ProtectionClasses() []foundation.FileProtectionType
	SetProtectionClasses(value []foundation.FileProtectionType)
}

// The criteria to apply when searching previously indexed app content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery?language=objc
type SearchQuery struct {
	objc.Object
}

func SearchQueryFrom(ptr unsafe.Pointer) SearchQuery {
	return SearchQuery{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SearchQuery) InitWithQueryStringAttributes(queryString string, attributes []string) SearchQuery {
	rv := objc.Call[SearchQuery](s_, objc.Sel("initWithQueryString:attributes:"), queryString, attributes)
	return rv
}

// Initializes and returns a query object with the specified query string and item attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649308-initwithquerystring?language=objc
func NewSearchQueryWithQueryStringAttributes(queryString string, attributes []string) SearchQuery {
	instance := SearchQueryClass.Alloc().InitWithQueryStringAttributes(queryString, attributes)
	instance.Autorelease()
	return instance
}

func (sc _SearchQueryClass) Alloc() SearchQuery {
	rv := objc.Call[SearchQuery](sc, objc.Sel("alloc"))
	return rv
}

func SearchQuery_Alloc() SearchQuery {
	return SearchQueryClass.Alloc()
}

func (sc _SearchQueryClass) New() SearchQuery {
	rv := objc.Call[SearchQuery](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchQuery() SearchQuery {
	return SearchQueryClass.New()
}

func (s_ SearchQuery) Init() SearchQuery {
	rv := objc.Call[SearchQuery](s_, objc.Sel("init"))
	return rv
}

// Queries the index for items that match the query object’s specifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649296-start?language=objc
func (s_ SearchQuery) Start() {
	objc.Call[objc.Void](s_, objc.Sel("start"))
}

// Cancels a query operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649309-cancel?language=objc
func (s_ SearchQuery) Cancel() {
	objc.Call[objc.Void](s_, objc.Sel("cancel"))
}

// A Boolean value that indicates if the system has canceled the query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649294-cancelled?language=objc
func (s_ SearchQuery) IsCancelled() bool {
	rv := objc.Call[bool](s_, objc.Sel("isCancelled"))
	return rv
}

// The number of items found so far. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649300-founditemcount?language=objc
func (s_ SearchQuery) FoundItemCount() uint {
	rv := objc.Call[uint](s_, objc.Sel("foundItemCount"))
	return rv
}

// The block to execute when the query completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649312-completionhandler?language=objc
func (s_ SearchQuery) CompletionHandler() func(error foundation.Error) {
	rv := objc.Call[func(error foundation.Error)](s_, objc.Sel("completionHandler"))
	return rv
}

// The block to execute when the query completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649312-completionhandler?language=objc
func (s_ SearchQuery) SetCompletionHandler(value func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("setCompletionHandler:"), value)
}

// The block to execute when the query finds a new batch of matching items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649306-founditemshandler?language=objc
func (s_ SearchQuery) FoundItemsHandler() func(items []SearchableItem) {
	rv := objc.Call[func(items []SearchableItem)](s_, objc.Sel("foundItemsHandler"))
	return rv
}

// The block to execute when the query finds a new batch of matching items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649306-founditemshandler?language=objc
func (s_ SearchQuery) SetFoundItemsHandler(value func(items []SearchableItem)) {
	objc.Call[objc.Void](s_, objc.Sel("setFoundItemsHandler:"), value)
}

// An array of data protection classes that correspond to the protection classes associated with the indexed items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649311-protectionclasses?language=objc
func (s_ SearchQuery) ProtectionClasses() []foundation.FileProtectionType {
	rv := objc.Call[[]foundation.FileProtectionType](s_, objc.Sel("protectionClasses"))
	return rv
}

// An array of data protection classes that correspond to the protection classes associated with the indexed items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/1649311-protectionclasses?language=objc
func (s_ SearchQuery) SetProtectionClasses(value []foundation.FileProtectionType) {
	objc.Call[objc.Void](s_, objc.Sel("setProtectionClasses:"), value)
}
