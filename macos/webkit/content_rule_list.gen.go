// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ContentRuleListClass = _ContentRuleListClass{objc.GetClass("WKContentRuleList")}

type _ContentRuleListClass struct {
	objc.Class
}

type IContentRuleList interface {
	objc.IObject
	Identifier() string
}

type ContentRuleList struct {
	objc.Object
}

func MakeContentRuleList(ptr unsafe.Pointer) ContentRuleList {
	return ContentRuleList{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ContentRuleListClass) Alloc() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](cc, objc.GetSelector("alloc"))
	return rv
}

func ContentRuleList_Alloc() ContentRuleList {
	return ContentRuleListClass.Alloc()
}

func (cc _ContentRuleListClass) New() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewContentRuleList() ContentRuleList {
	return ContentRuleListClass.New()
}

func ContentRuleList_New() ContentRuleList {
	return ContentRuleListClass.New()
}

func (c_ ContentRuleList) Init() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](c_, objc.GetSelector("init"))
	return rv
}

func ContentRuleList_Init() ContentRuleList {
	return ContentRuleListClass.Alloc().Init()
}

func (c_ ContentRuleList) Identifier() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("identifier"))
	return rv
}
