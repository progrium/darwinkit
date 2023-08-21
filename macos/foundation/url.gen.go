// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URL] class.
var URLClass = _URLClass{objc.GetClass("NSURL")}

type _URLClass struct {
	objc.Class
}

// An interface definition for the [URL] class.
type IURL interface {
	objc.IObject
	CheckPromisedItemIsReachableAndReturnError(error IError) bool
	GetPromisedItemResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool
	IsFileReferenceURL() bool
	WriteToPasteboard(pasteBoard objc.IObject)
	FileReferenceURL() URL
	StopAccessingSecurityScopedResource()
	CheckResourceIsReachableAndReturnError(error IError) bool
	RemoveAllCachedResourceValues()
	GetFileSystemRepresentationMaxLength(buffer *uint8, maxBufferLength uint) bool
	URLByAppendingPathComponentConformingToType(partialName string, contentType objc.IObject) URL
	RemoveCachedResourceValueForKey(key URLResourceKey)
	SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool
	URLByAppendingPathExtensionForType(contentType objc.IObject) URL
	GetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error IError) []byte
	StartAccessingSecurityScopedResource() bool
	SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey)
	SetResourceValuesError(keyedValues map[URLResourceKey]objc.IObject, error IError) bool
	URLByAppendingPathExtension(pathExtension string) URL
	PromisedItemResourceValuesForKeysError(keys []URLResourceKey, error IError) map[URLResourceKey]objc.Object
	Password() string
	Path() string
	BaseURL() URL
	ResourceSpecifier() string
	Host() string
	Scheme() string
	PathComponents() []string
	AbsoluteString() string
	IsFileURL() bool
	Query() string
	URLByResolvingSymlinksInPath() URL
	HasDirectoryPath() bool
	Fragment() string
	RelativeString() string
	User() string
	URLByStandardizingPath() URL
	Port() Number
	DataRepresentation() []byte
	URLByDeletingLastPathComponent() URL
	PathExtension() string
	URLByDeletingPathExtension() URL
	AbsoluteURL() URL
	StandardizedURL() URL
	LastPathComponent() string
	FileSystemRepresentation() *uint8
	FilePathURL() URL
	RelativePath() string
}

// An object that represents the location of a resource, such as an item on a remote server or the path to a local file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl?language=objc
type URL struct {
	objc.Object
}

func URLFrom(ptr unsafe.Pointer) URL {
	return URL{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URL) InitAbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410750-initabsoluteurlwithdatarepresent?language=objc
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	instance := URLClass.Alloc().InitAbsoluteURLWithDataRepresentationRelativeToURL(data, baseURL)
	instance.Autorelease()
	return instance
}

func (u_ URL) InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *uint8, isDir bool, baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.Ptr(baseURL))
	return rv
}

// Initializes a URL object with a C string representing a local file system path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411210-initfileurlwithfilesystemreprese?language=objc
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *uint8, isDir bool, baseURL IURL) URL {
	instance := URLClass.Alloc().InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path, isDir, baseURL)
	instance.Autorelease()
	return instance
}

func (u_ URL) InitWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1416851-initwithdatarepresentation?language=objc
func NewURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	instance := URLClass.Alloc().InitWithDataRepresentationRelativeToURL(data, baseURL)
	instance.Autorelease()
	return instance
}

func (u_ URL) InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error IError) URL {
	rv := objc.Call[URL](u_, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.Ptr(relativeURL), isStale, objc.Ptr(error))
	return rv
}

// Initializes a newly created NSURL that points to a location specified by resolving bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413475-initbyresolvingbookmarkdata?language=objc
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error IError) URL {
	instance := URLClass.Alloc().InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData, options, relativeURL, isStale, error)
	instance.Autorelease()
	return instance
}

func (u_ URL) InitWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("initWithString:relativeToURL:"), URLString, objc.Ptr(baseURL))
	return rv
}

// Initializes an NSURL object with a base URL and a relative string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417949-initwithstring?language=objc
func NewURLWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	instance := URLClass.Alloc().InitWithStringRelativeToURL(URLString, baseURL)
	instance.Autorelease()
	return instance
}

func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error IError) URL {
	rv := objc.Call[URL](uc, objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.Ptr(relativeURL), isStale, objc.Ptr(error))
	return rv
}

// Returns a new URL made by resolving bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1572035-urlbyresolvingbookmarkdata?language=objc
func URL_URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error IError) URL {
	return URLClass.URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData, options, relativeURL, isStale, error)
}

func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error IError) URL {
	rv := objc.Call[URL](uc, objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), objc.Ptr(url), options, objc.Ptr(error))
	return rv
}

// Returns a new URL made by resolving the alias file at url. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1416404-urlbyresolvingaliasfileaturl?language=objc
func URL_URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error IError) URL {
	return URLClass.URLByResolvingAliasFileAtURLOptionsError(url, options, error)
}

func (uc _URLClass) URLWithString(URLString string) URL {
	rv := objc.Call[URL](uc, objc.Sel("URLWithString:"), URLString)
	return rv
}

// Creates and returns an NSURL object initialized with a provided URL string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1572047-urlwithstring?language=objc
func URL_URLWithString(URLString string) URL {
	return URLClass.URLWithString(URLString)
}

func (u_ URL) InitFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), path, isDir, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417932-initfileurlwithpath?language=objc
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	instance := URLClass.Alloc().InitFileURLWithPathIsDirectoryRelativeToURL(path, isDir, baseURL)
	instance.Autorelease()
	return instance
}

func (uc _URLClass) Alloc() URL {
	rv := objc.Call[URL](uc, objc.Sel("alloc"))
	return rv
}

func URL_Alloc() URL {
	return URLClass.Alloc()
}

func (uc _URLClass) New() URL {
	rv := objc.Call[URL](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURL() URL {
	return URLClass.New()
}

func (u_ URL) Init() URL {
	rv := objc.Call[URL](u_, objc.Sel("init"))
	return rv
}

// Creates an alias file on disk at a specified location with specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408532-writebookmarkdata?language=objc
func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error IError) bool {
	rv := objc.Call[bool](uc, objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, objc.Ptr(bookmarkFileURL), options, objc.Ptr(error))
	return rv
}

// Creates an alias file on disk at a specified location with specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408532-writebookmarkdata?language=objc
func URL_WriteBookmarkDataToURLOptionsError(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error IError) bool {
	return URLClass.WriteBookmarkDataToURLOptionsError(bookmarkData, bookmarkFileURL, options, error)
}

// Returns whether the promised item can be reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410411-checkpromiseditemisreachableandr?language=objc
func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), objc.Ptr(error))
	return rv
}

// Returns a new URL object initialized with a C string representing a local file system path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411492-fileurlwithfilesystemrepresentat?language=objc
func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *uint8, isDir bool, baseURL IURL) URL {
	rv := objc.Call[URL](uc, objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.Ptr(baseURL))
	return rv
}

// Returns a new URL object initialized with a C string representing a local file system path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411492-fileurlwithfilesystemrepresentat?language=objc
func URL_FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *uint8, isDir bool, baseURL IURL) URL {
	return URLClass.FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path, isDir, baseURL)
}

// Returns the value of the resource property for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1414238-getpromiseditemresourcevalue?language=objc
func (u_ URL) GetPromisedItemResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("getPromisedItemResourceValue:forKey:error:"), value, key, objc.Ptr(error))
	return rv
}

// Returns whether the URL is a file reference URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408507-isfilereferenceurl?language=objc
func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.Call[bool](u_, objc.Sel("isFileReferenceURL"))
	return rv
}

// Writes the URL to the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1532980-writetopasteboard?language=objc
func (u_ URL) WriteToPasteboard(pasteBoard objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("writeToPasteboard:"), objc.Ptr(pasteBoard))
}

// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1418097-resourcevaluesforkeys?language=objc
func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	rv := objc.Call[map[URLResourceKey]objc.Object](uc, objc.Sel("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}

// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1418097-resourcevaluesforkeys?language=objc
func URL_ResourceValuesForKeysFromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	return URLClass.ResourceValuesForKeysFromBookmarkData(keys, bookmarkData)
}

// Returns a new file reference URL that points to the same resource as the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408631-filereferenceurl?language=objc
func (u_ URL) FileReferenceURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("fileReferenceURL"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413201-fileurlwithpath?language=objc
func (uc _URLClass) FileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	rv := objc.Call[URL](uc, objc.Sel("fileURLWithPath:relativeToURL:"), path, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413201-fileurlwithpath?language=objc
func URL_FileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	return URLClass.FileURLWithPathRelativeToURL(path, baseURL)
}

// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413736-stopaccessingsecurityscopedresou?language=objc
func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.Call[objc.Void](u_, objc.Sel("stopAccessingSecurityScopedResource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412404-absoluteurlwithdatarepresentatio?language=objc
func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.Call[URL](uc, objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412404-absoluteurlwithdatarepresentatio?language=objc
func URL_AbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.AbsoluteURLWithDataRepresentationRelativeToURL(data, baseURL)
}

// Returns whether the resource pointed to by a file URL can be reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410597-checkresourceisreachableandretur?language=objc
func (u_ URL) CheckResourceIsReachableAndReturnError(error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("checkResourceIsReachableAndReturnError:"), objc.Ptr(error))
	return rv
}

// Removes all cached resource values and temporary resource values from the URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417078-removeallcachedresourcevalues?language=objc
func (u_ URL) RemoveAllCachedResourceValues() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllCachedResourceValues"))
}

// Fills the provided buffer with a C string representing a local file system path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1415117-getfilesystemrepresentation?language=objc
func (u_ URL) GetFileSystemRepresentationMaxLength(buffer *uint8, maxBufferLength uint) bool {
	rv := objc.Call[bool](u_, objc.Sel("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}

// Returns a URL by appending the specified path component with the file extension for a uniform type identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/3564810-urlbyappendingpathcomponent?language=objc
func (u_ URL) URLByAppendingPathComponentConformingToType(partialName string, contentType objc.IObject) URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByAppendingPathComponent:conformingToType:"), partialName, objc.Ptr(contentType))
	return rv
}

// Removes the cached resource value identified by a given key from the URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410758-removecachedresourcevalueforkey?language=objc
func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.Call[objc.Void](u_, objc.Sel("removeCachedResourceValueForKey:"), key)
}

// Sets the URL’s resource property for a given key to a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413819-setresourcevalue?language=objc
func (u_ URL) SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("setResourceValue:forKey:error:"), value, key, objc.Ptr(error))
	return rv
}

// Returns a URL by appending the path extension for a uniform type identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/3584837-urlbyappendingpathextensionforty?language=objc
func (u_ URL) URLByAppendingPathExtensionForType(contentType objc.IObject) URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByAppendingPathExtensionForType:"), objc.Ptr(contentType))
	return rv
}

// Returns the value of the resource property for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408874-getresourcevalue?language=objc
func (u_ URL) GetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("getResourceValue:forKey:error:"), value, key, objc.Ptr(error))
	return rv
}

// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408344-bookmarkdatawithcontentsofurl?language=objc
func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error IError) []byte {
	rv := objc.Call[[]byte](uc, objc.Sel("bookmarkDataWithContentsOfURL:error:"), objc.Ptr(bookmarkFileURL), objc.Ptr(error))
	return rv
}

// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408344-bookmarkdatawithcontentsofurl?language=objc
func URL_BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error IError) []byte {
	return URLClass.BookmarkDataWithContentsOfURLError(bookmarkFileURL, error)
}

// Reads an NSURL object off of the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1525106-urlfrompasteboard?language=objc
func (uc _URLClass) URLFromPasteboard(pasteBoard objc.IObject) URL {
	rv := objc.Call[URL](uc, objc.Sel("URLFromPasteboard:"), objc.Ptr(pasteBoard))
	return rv
}

// Reads an NSURL object off of the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1525106-urlfrompasteboard?language=objc
func URL_URLFromPasteboard(pasteBoard objc.IObject) URL {
	return URLClass.URLFromPasteboard(pasteBoard)
}

// Returns a bookmark for the URL, created with specified options and resource values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417795-bookmarkdatawithoptions?language=objc
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error IError) []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, objc.Ptr(relativeURL), objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1572042-urlwithdatarepresentation?language=objc
func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.Call[URL](uc, objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, objc.Ptr(baseURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1572042-urlwithdatarepresentation?language=objc
func URL_URLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.URLWithDataRepresentationRelativeToURL(data, baseURL)
}

// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417051-startaccessingsecurityscopedreso?language=objc
func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.Call[bool](u_, objc.Sel("startAccessingSecurityScopedResource"))
	return rv
}

// Sets a temporary resource value on the URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411094-settemporaryresourcevalue?language=objc
func (u_ URL) SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey) {
	objc.Call[objc.Void](u_, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}

// Sets the URL’s resource properties for a given set of keys to a given set of values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408208-setresourcevalues?language=objc
func (u_ URL) SetResourceValuesError(keyedValues map[URLResourceKey]objc.IObject, error IError) bool {
	rv := objc.Call[bool](u_, objc.Sel("setResourceValues:error:"), keyedValues, objc.Ptr(error))
	return rv
}

// Initializes and returns a newly created NSURL object as a file URL with specified path components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1414206-fileurlwithpathcomponents?language=objc
func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := objc.Call[URL](uc, objc.Sel("fileURLWithPathComponents:"), components)
	return rv
}

// Initializes and returns a newly created NSURL object as a file URL with specified path components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1414206-fileurlwithpathcomponents?language=objc
func URL_FileURLWithPathComponents(components []string) URL {
	return URLClass.FileURLWithPathComponents(components)
}

// Returns a new URL by appending a path extension to the original URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417082-urlbyappendingpathextension?language=objc
func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByAppendingPathExtension:"), pathExtension)
	return rv
}

// Returns the resource values for the properties identified by specified array of keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1407746-promiseditemresourcevaluesforkey?language=objc
func (u_ URL) PromisedItemResourceValuesForKeysError(keys []URLResourceKey, error IError) map[URLResourceKey]objc.Object {
	rv := objc.Call[map[URLResourceKey]objc.Object](u_, objc.Sel("promisedItemResourceValuesForKeys:error:"), keys, objc.Ptr(error))
	return rv
}

// The password conforming to RFC 1808. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412096-password?language=objc
func (u_ URL) Password() string {
	rv := objc.Call[string](u_, objc.Sel("password"))
	return rv
}

// The path, conforming to RFC 1808. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408809-path?language=objc
func (u_ URL) Path() string {
	rv := objc.Call[string](u_, objc.Sel("path"))
	return rv
}

// The base URL. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412311-baseurl?language=objc
func (u_ URL) BaseURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("baseURL"))
	return rv
}

// The resource specifier. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1415309-resourcespecifier?language=objc
func (u_ URL) ResourceSpecifier() string {
	rv := objc.Call[string](u_, objc.Sel("resourceSpecifier"))
	return rv
}

// The host, conforming to RFC 1808. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413640-host?language=objc
func (u_ URL) Host() string {
	rv := objc.Call[string](u_, objc.Sel("host"))
	return rv
}

// The scheme. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413437-scheme?language=objc
func (u_ URL) Scheme() string {
	rv := objc.Call[string](u_, objc.Sel("scheme"))
	return rv
}

// An array containing the  path components. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1407365-pathcomponents?language=objc
func (u_ URL) PathComponents() []string {
	rv := objc.Call[[]string](u_, objc.Sel("pathComponents"))
	return rv
}

// The URL string for the receiver as an absolute URL. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1409868-absolutestring?language=objc
func (u_ URL) AbsoluteString() string {
	rv := objc.Call[string](u_, objc.Sel("absoluteString"))
	return rv
}

// A boolean value that determines whether the receiver is a file URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408782-fileurl?language=objc
func (u_ URL) IsFileURL() bool {
	rv := objc.Call[bool](u_, objc.Sel("isFileURL"))
	return rv
}

// The query string, conforming to RFC 1808. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1407543-query?language=objc
func (u_ URL) Query() string {
	rv := objc.Call[string](u_, objc.Sel("query"))
	return rv
}

// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1415965-urlbyresolvingsymlinksinpath?language=objc
func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByResolvingSymlinksInPath"))
	return rv
}

// A Boolean value that indicates whether the URL string’s path represents a directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411475-hasdirectorypath?language=objc
func (u_ URL) HasDirectoryPath() bool {
	rv := objc.Call[bool](u_, objc.Sel("hasDirectoryPath"))
	return rv
}

// The fragment identifier, conforming to RFC 1808. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413775-fragment?language=objc
func (u_ URL) Fragment() string {
	rv := objc.Call[string](u_, objc.Sel("fragment"))
	return rv
}

// A string representation of the relative portion of the URL. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411417-relativestring?language=objc
func (u_ URL) RelativeString() string {
	rv := objc.Call[string](u_, objc.Sel("relativeString"))
	return rv
}

// The user name, conforming to RFC 1808. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1418335-user?language=objc
func (u_ URL) User() string {
	rv := objc.Call[string](u_, objc.Sel("user"))
	return rv
}

// A URL that points to the same resource as the original URL using an absolute path. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1414302-urlbystandardizingpath?language=objc
func (u_ URL) URLByStandardizingPath() URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByStandardizingPath"))
	return rv
}

// The port, conforming to RFC 1808. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1413455-port?language=objc
func (u_ URL) Port() Number {
	rv := objc.Call[Number](u_, objc.Sel("port"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1407656-datarepresentation?language=objc
func (u_ URL) DataRepresentation() []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("dataRepresentation"))
	return rv
}

// A URL you create by removing the last path component from the receiver. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411592-urlbydeletinglastpathcomponent?language=objc
func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByDeletingLastPathComponent"))
	return rv
}

// The path extension. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410208-pathextension?language=objc
func (u_ URL) PathExtension() string {
	rv := objc.Call[string](u_, objc.Sel("pathExtension"))
	return rv
}

// A URL you create by removing the path extension from the receiver, if any. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412357-urlbydeletingpathextension?language=objc
func (u_ URL) URLByDeletingPathExtension() URL {
	rv := objc.Call[URL](u_, objc.Sel("URLByDeletingPathExtension"))
	return rv
}

// An absolute URL that refers to the same resource as the receiver. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1414266-absoluteurl?language=objc
func (u_ URL) AbsoluteURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("absoluteURL"))
	return rv
}

// A copy of the URL with any instances of ".." or "." removed from its path. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1411073-standardizedurl?language=objc
func (u_ URL) StandardizedURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("standardizedURL"))
	return rv
}

// The last path component. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1417444-lastpathcomponent?language=objc
func (u_ URL) LastPathComponent() string {
	rv := objc.Call[string](u_, objc.Sel("lastPathComponent"))
	return rv
}

// A C string containing the URL’s file system path. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1412925-filesystemrepresentation?language=objc
func (u_ URL) FileSystemRepresentation() *uint8 {
	rv := objc.Call[*uint8](u_, objc.Sel("fileSystemRepresentation"))
	return rv
}

// A file path URL that points to the same resource as the URL object. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1408442-filepathurl?language=objc
func (u_ URL) FilePathURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("filePathURL"))
	return rv
}

// The relative path, conforming to RFC 1808. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/1410263-relativepath?language=objc
func (u_ URL) RelativePath() string {
	rv := objc.Call[string](u_, objc.Sel("relativePath"))
	return rv
}
