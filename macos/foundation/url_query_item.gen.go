// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLQueryItem] class.
var URLQueryItemClass = _URLQueryItemClass{objc.GetClass("NSURLQueryItem")}

type _URLQueryItemClass struct {
	objc.Class
}

// An interface definition for the [URLQueryItem] class.
type IURLQueryItem interface {
	objc.IObject
	Value() string
	Name() string
}

// An object representing a single name/value pair for an item in the query portion of a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem?language=objc
type URLQueryItem struct {
	objc.Object
}

func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLQueryItem) InitWithNameValue(name string, value string) URLQueryItem {
	rv := objc.Call[URLQueryItem](u_, objc.Sel("initWithName:value:"), name, value)
	return rv
}

// Initializes a newly allocated query item with the specified name and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/1410963-initwithname?language=objc
func NewURLQueryItemWithNameValue(name string, value string) URLQueryItem {
	instance := URLQueryItemClass.Alloc().InitWithNameValue(name, value)
	instance.Autorelease()
	return instance
}

func (uc _URLQueryItemClass) QueryItemWithNameValue(name string, value string) URLQueryItem {
	rv := objc.Call[URLQueryItem](uc, objc.Sel("queryItemWithName:value:"), name, value)
	return rv
}

// Creates a new query item with the specified name and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/1572045-queryitemwithname?language=objc
func URLQueryItem_QueryItemWithNameValue(name string, value string) URLQueryItem {
	return URLQueryItemClass.QueryItemWithNameValue(name, value)
}

func (uc _URLQueryItemClass) Alloc() URLQueryItem {
	rv := objc.Call[URLQueryItem](uc, objc.Sel("alloc"))
	return rv
}

func URLQueryItem_Alloc() URLQueryItem {
	return URLQueryItemClass.Alloc()
}

func (uc _URLQueryItemClass) New() URLQueryItem {
	rv := objc.Call[URLQueryItem](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLQueryItem() URLQueryItem {
	return URLQueryItemClass.New()
}

func (u_ URLQueryItem) Init() URLQueryItem {
	rv := objc.Call[URLQueryItem](u_, objc.Sel("init"))
	return rv
}

// The value for the query item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/1412041-value?language=objc
func (u_ URLQueryItem) Value() string {
	rv := objc.Call[string](u_, objc.Sel("value"))
	return rv
}

// The name of the query item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/1407785-name?language=objc
func (u_ URLQueryItem) Name() string {
	rv := objc.Call[string](u_, objc.Sel("name"))
	return rv
}
