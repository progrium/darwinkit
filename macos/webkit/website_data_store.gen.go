// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var WebsiteDataStoreClass = _WebsiteDataStoreClass{objc.GetClass("WKWebsiteDataStore")}

type _WebsiteDataStoreClass struct {
	objc.Class
}

type IWebsiteDataStore interface {
	objc.IObject
	FetchDataRecordsOfTypesCompletionHandler(dataTypes foundation.ISet, completionHandler func(param1 []WebsiteDataRecord))
	RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes foundation.ISet, dataRecords []IWebsiteDataRecord, completionHandler func())
	RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func())
	IsPersistent() bool
	HttpCookieStore() HTTPCookieStore
}

type WebsiteDataStore struct {
	objc.Object
}

func MakeWebsiteDataStore(ptr unsafe.Pointer) WebsiteDataStore {
	return WebsiteDataStore{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebsiteDataStoreClass) Alloc() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.GetSelector("alloc"))
	return rv
}

func WebsiteDataStore_Alloc() WebsiteDataStore {
	return WebsiteDataStoreClass.Alloc()
}

func (wc _WebsiteDataStoreClass) New() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.New()
}

func WebsiteDataStore_New() WebsiteDataStore {
	return WebsiteDataStoreClass.New()
}

func (w_ WebsiteDataStore) Init() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](w_, objc.GetSelector("init"))
	return rv
}

func WebsiteDataStore_Init() WebsiteDataStore {
	return WebsiteDataStoreClass.Alloc().Init()
}

func (wc _WebsiteDataStoreClass) DefaultDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.GetSelector("defaultDataStore"))
	return rv
}

func WebsiteDataStore_DefaultDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.DefaultDataStore()
}

func (wc _WebsiteDataStoreClass) NonPersistentDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.GetSelector("nonPersistentDataStore"))
	return rv
}

func WebsiteDataStore_NonPersistentDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.NonPersistentDataStore()
}

func (w_ WebsiteDataStore) FetchDataRecordsOfTypesCompletionHandler(dataTypes foundation.ISet, completionHandler func(param1 []WebsiteDataRecord)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("fetchDataRecordsOfTypes:completionHandler:"), objc.ExtractPtr(dataTypes), completionHandler)
}

func (wc _WebsiteDataStoreClass) AllWebsiteDataTypes() foundation.Set {
	rv := objc.CallMethod[foundation.Set](wc, objc.GetSelector("allWebsiteDataTypes"))
	return rv
}

func WebsiteDataStore_AllWebsiteDataTypes() foundation.Set {
	return WebsiteDataStoreClass.AllWebsiteDataTypes()
}

func (w_ WebsiteDataStore) RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes foundation.ISet, dataRecords []IWebsiteDataRecord, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("removeDataOfTypes:forDataRecords:completionHandler:"), objc.ExtractPtr(dataTypes), dataRecords, completionHandler)
}

func (w_ WebsiteDataStore) RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("removeDataOfTypes:modifiedSince:completionHandler:"), objc.ExtractPtr(dataTypes), objc.ExtractPtr(date), completionHandler)
}

func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isPersistent"))
	return rv
}

func (w_ WebsiteDataStore) HttpCookieStore() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](w_, objc.GetSelector("httpCookieStore"))
	return rv
}
