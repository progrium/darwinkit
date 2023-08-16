// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebsiteDataRecord] class.
var WebsiteDataRecordClass = _WebsiteDataRecordClass{objc.GetClass("WKWebsiteDataRecord")}

type _WebsiteDataRecordClass struct {
	objc.Class
}

// An interface definition for the [WebsiteDataRecord] class.
type IWebsiteDataRecord interface {
	objc.IObject
	DataTypes() foundation.Set
	DisplayName() string
}

// A record of the data that a particular website stores persistently. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatarecord?language=objc
type WebsiteDataRecord struct {
	objc.Object
}

func WebsiteDataRecordFrom(ptr unsafe.Pointer) WebsiteDataRecord {
	return WebsiteDataRecord{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebsiteDataRecordClass) Alloc() WebsiteDataRecord {
	rv := objc.Call[WebsiteDataRecord](wc, objc.Sel("alloc"))
	return rv
}

func WebsiteDataRecord_Alloc() WebsiteDataRecord {
	return WebsiteDataRecordClass.Alloc()
}

func (wc _WebsiteDataRecordClass) New() WebsiteDataRecord {
	rv := objc.Call[WebsiteDataRecord](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataRecord() WebsiteDataRecord {
	return WebsiteDataRecordClass.New()
}

func (w_ WebsiteDataRecord) Init() WebsiteDataRecord {
	rv := objc.Call[WebsiteDataRecord](w_, objc.Sel("init"))
	return rv
}

// The types of data associated with the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatarecord/1538007-datatypes?language=objc
func (w_ WebsiteDataRecord) DataTypes() foundation.Set {
	rv := objc.Call[foundation.Set](w_, objc.Sel("dataTypes"))
	return rv
}

// The display name for the data record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatarecord/1537733-displayname?language=objc
func (w_ WebsiteDataRecord) DisplayName() string {
	rv := objc.Call[string](w_, objc.Sel("displayName"))
	return rv
}
