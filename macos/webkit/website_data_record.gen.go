// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var WebsiteDataRecordClass = _WebsiteDataRecordClass{objc.GetClass("WKWebsiteDataRecord")}

type _WebsiteDataRecordClass struct {
	objc.Class
}

type IWebsiteDataRecord interface {
	objc.IObject
	DisplayName() string
	DataTypes() foundation.Set
}

type WebsiteDataRecord struct {
	objc.Object
}

func MakeWebsiteDataRecord(ptr unsafe.Pointer) WebsiteDataRecord {
	return WebsiteDataRecord{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebsiteDataRecordClass) Alloc() WebsiteDataRecord {
	rv := objc.CallMethod[WebsiteDataRecord](wc, objc.GetSelector("alloc"))
	return rv
}

func WebsiteDataRecord_Alloc() WebsiteDataRecord {
	return WebsiteDataRecordClass.Alloc()
}

func (wc _WebsiteDataRecordClass) New() WebsiteDataRecord {
	rv := objc.CallMethod[WebsiteDataRecord](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataRecord() WebsiteDataRecord {
	return WebsiteDataRecordClass.New()
}

func WebsiteDataRecord_New() WebsiteDataRecord {
	return WebsiteDataRecordClass.New()
}

func (w_ WebsiteDataRecord) Init() WebsiteDataRecord {
	rv := objc.CallMethod[WebsiteDataRecord](w_, objc.GetSelector("init"))
	return rv
}

func WebsiteDataRecord_Init() WebsiteDataRecord {
	return WebsiteDataRecordClass.Alloc().Init()
}

func (w_ WebsiteDataRecord) DisplayName() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("displayName"))
	return rv
}

func (w_ WebsiteDataRecord) DataTypes() foundation.Set {
	rv := objc.CallMethod[foundation.Set](w_, objc.GetSelector("dataTypes"))
	return rv
}
