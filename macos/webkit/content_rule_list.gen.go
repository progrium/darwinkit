// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentRuleList] class.
var ContentRuleListClass = _ContentRuleListClass{objc.GetClass("WKContentRuleList")}

type _ContentRuleListClass struct {
	objc.Class
}

// An interface definition for the [ContentRuleList] class.
type IContentRuleList interface {
	objc.IObject
	Identifier() string
}

// A compiled list of rules to apply to web content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentrulelist?language=objc
type ContentRuleList struct {
	objc.Object
}

func ContentRuleListFrom(ptr unsafe.Pointer) ContentRuleList {
	return ContentRuleList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentRuleListClass) Alloc() ContentRuleList {
	rv := objc.Call[ContentRuleList](cc, objc.Sel("alloc"))
	return rv
}

func ContentRuleList_Alloc() ContentRuleList {
	return ContentRuleListClass.Alloc()
}

func (cc _ContentRuleListClass) New() ContentRuleList {
	rv := objc.Call[ContentRuleList](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentRuleList() ContentRuleList {
	return ContentRuleListClass.New()
}

func (c_ ContentRuleList) Init() ContentRuleList {
	rv := objc.Call[ContentRuleList](c_, objc.Sel("init"))
	return rv
}

// The identifier for the rule list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentrulelist/2902755-identifier?language=objc
func (c_ ContentRuleList) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}
