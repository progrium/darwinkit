// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var URLClass = _URLClass{objc.GetClass("NSURL")}

type _URLClass struct {
	objc.Class
}

type IURL interface {
	objc.IObject
	GetFileSystemRepresentationMaxLength(buffer *byte, maxBufferLength uint) bool
	CheckResourceIsReachableAndReturnError(error *Error) bool
	IsFileReferenceURL() bool
	ResourceValuesForKeysError(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object
	SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error *Error) bool
	SetResourceValuesError(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey)
	SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey)
	FileReferenceURL() URL
	URLByAppendingPathComponent(pathComponent string) URL
	URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) URL
	URLByAppendingPathExtension(pathExtension string) URL
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	CheckPromisedItemIsReachableAndReturnError(error *Error) bool
	GetPromisedItemResourceValueForKeyError(value *objc.Object, key URLResourceKey, error *Error) bool
	PromisedItemResourceValuesForKeysError(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object
	DataRepresentation() []byte
	IsFileURL() bool
	AbsoluteString() string
	AbsoluteURL() URL
	BaseURL() URL
	FileSystemRepresentation() *byte
	Fragment() string
	Host() string
	LastPathComponent() string
	Password() string
	Path() string
	PathComponents() []string
	PathExtension() string
	Port() Number
	Query() string
	RelativePath() string
	RelativeString() string
	ResourceSpecifier() string
	Scheme() string
	StandardizedURL() URL
	User() string
	FilePathURL() URL
	URLByDeletingLastPathComponent() URL
	URLByDeletingPathExtension() URL
	URLByResolvingSymlinksInPath() URL
	URLByStandardizingPath() URL
	HasDirectoryPath() bool
}

type URL struct {
	objc.Object
}

func MakeURL(ptr unsafe.Pointer) URL {
	return URL{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _URLClass) URLWithString(URLString string) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithString:"), URLString)
	return rv
}

func URL_URLWithString(URLString string) URL {
	return URLClass.URLWithString(URLString)
}

func (u_ URL) InitWithString(URLString string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithString:"), URLString)
	return rv
}

func URL_InitWithString(URLString string) URL {
	return URLClass.Alloc().InitWithString(URLString)
}

func (uc _URLClass) URLWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithString:relativeToURL:"), URLString, objc.ExtractPtr(baseURL))
	return rv
}

func URL_URLWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	return URLClass.URLWithStringRelativeToURL(URLString, baseURL)
}

func (u_ URL) InitWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithString:relativeToURL:"), URLString, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	return URLClass.Alloc().InitWithStringRelativeToURL(URLString, baseURL)
}

func (u_ URL) InitFileURLWithPathIsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:isDirectory:"), path, isDir)
	return rv
}

func URL_InitFileURLWithPathIsDirectory(path string, isDir bool) URL {
	return URLClass.Alloc().InitFileURLWithPathIsDirectory(path, isDir)
}

func (u_ URL) InitFileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:relativeToURL:"), path, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitFileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	return URLClass.Alloc().InitFileURLWithPathRelativeToURL(path, baseURL)
}

func (u_ URL) InitFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	return URLClass.Alloc().InitFileURLWithPathIsDirectoryRelativeToURL(path, isDir, baseURL)
}

func (u_ URL) InitFileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:"), path)
	return rv
}

func URL_InitFileURLWithPath(path string) URL {
	return URLClass.Alloc().InitFileURLWithPath(path)
}

func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error *Error) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLByResolvingAliasFileAtURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(error))
	return rv
}

func URL_URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error *Error) URL {
	return URLClass.URLByResolvingAliasFileAtURLOptionsError(url, options, error)
}

func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.ExtractPtr(relativeURL), isStale, unsafe.Pointer(error))
	return rv
}

func URL_URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	return URLClass.URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData, options, relativeURL, isStale, error)
}

func (u_ URL) InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.ExtractPtr(relativeURL), isStale, unsafe.Pointer(error))
	return rv
}

func URL_InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	return URLClass.Alloc().InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData, options, relativeURL, isStale, error)
}

func (u_ URL) InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	return URLClass.Alloc().InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path, isDir, baseURL)
}

func (u_ URL) InitAbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitAbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.Alloc().InitAbsoluteURLWithDataRepresentationRelativeToURL(data, baseURL)
}

func (u_ URL) InitWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func URL_InitWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.Alloc().InitWithDataRepresentationRelativeToURL(data, baseURL)
}

func (u_ URL) InitWithSchemeHostPath(scheme string, host string, path string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithScheme:host:path:"), scheme, host, path)
	return rv
}

func URL_InitWithSchemeHostPath(scheme string, host string, path string) URL {
	return URLClass.Alloc().InitWithSchemeHostPath(scheme, host, path)
}

func (uc _URLClass) Alloc() URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("alloc"))
	return rv
}

func URL_Alloc() URL {
	return URLClass.Alloc()
}

func (uc _URLClass) New() URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewURL() URL {
	return URLClass.New()
}

func URL_New() URL {
	return URLClass.New()
}

func (u_ URL) Init() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("init"))
	return rv
}

func URL_Init() URL {
	return URLClass.Alloc().Init()
}

func (uc _URLClass) FileURLWithPathIsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:isDirectory:"), path, isDir)
	return rv
}

func URL_FileURLWithPathIsDirectory(path string, isDir bool) URL {
	return URLClass.FileURLWithPathIsDirectory(path, isDir)
}

func (uc _URLClass) FileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:relativeToURL:"), path, objc.ExtractPtr(baseURL))
	return rv
}

func URL_FileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	return URLClass.FileURLWithPathRelativeToURL(path, baseURL)
}

func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func URL_FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	return URLClass.FileURLWithPathIsDirectoryRelativeToURL(path, isDir, baseURL)
}

func (uc _URLClass) FileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:"), path)
	return rv
}

func URL_FileURLWithPath(path string) URL {
	return URLClass.FileURLWithPath(path)
}

func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPathComponents:"), components)
	return rv
}

func URL_FileURLWithPathComponents(components []string) URL {
	return URLClass.FileURLWithPathComponents(components)
}

func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func URL_FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	return URLClass.FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path, isDir, baseURL)
}

func (u_ URL) GetFileSystemRepresentationMaxLength(buffer *byte, maxBufferLength uint) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}

func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("absoluteURLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func URL_AbsoluteURLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.AbsoluteURLWithDataRepresentationRelativeToURL(data, baseURL)
}

func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func URL_URLWithDataRepresentationRelativeToURL(data []byte, baseURL IURL) URL {
	return URLClass.URLWithDataRepresentationRelativeToURL(data, baseURL)
}

func (u_ URL) CheckResourceIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("checkResourceIsReachableAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isFileReferenceURL"))
	return rv
}

func (u_ URL) ResourceValuesForKeysError(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, objc.GetSelector("resourceValuesForKeys:error:"), keys, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setResourceValue:forKey:error:"), objc.ExtractPtr(value), key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValuesError(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setResourceValues:error:"), keyedValues, unsafe.Pointer(error))
	return rv
}

func (u_ URL) RemoveAllCachedResourceValues() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllCachedResourceValues"))
}

func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeCachedResourceValueForKey:"), key)
}

func (u_ URL) SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setTemporaryResourceValue:forKey:"), objc.ExtractPtr(value), key)
}

func (u_ URL) FileReferenceURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("fileReferenceURL"))
	return rv
}

func (u_ URL) URLByAppendingPathComponent(pathComponent string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathComponent:"), pathComponent)
	return rv
}

func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathComponent:isDirectory:"), pathComponent, isDirectory)
	return rv
}

func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathExtension:"), pathExtension)
	return rv
}

func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](uc, objc.GetSelector("bookmarkDataWithContentsOfURL:error:"), objc.ExtractPtr(bookmarkFileURL), unsafe.Pointer(error))
	return rv
}

func URL_BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error *Error) []byte {
	return URLClass.BookmarkDataWithContentsOfURLError(bookmarkFileURL, error)
}

func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, objc.ExtractPtr(relativeURL), unsafe.Pointer(error))
	return rv
}

func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](uc, objc.GetSelector("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}

func URL_ResourceValuesForKeysFromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	return URLClass.ResourceValuesForKeysFromBookmarkData(keys, bookmarkData)
}

func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error *Error) bool {
	rv := objc.CallMethod[bool](uc, objc.GetSelector("writeBookmarkData:toURL:options:error:"), bookmarkData, objc.ExtractPtr(bookmarkFileURL), options, unsafe.Pointer(error))
	return rv
}

func URL_WriteBookmarkDataToURLOptionsError(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error *Error) bool {
	return URLClass.WriteBookmarkDataToURLOptionsError(bookmarkData, bookmarkFileURL, options, error)
}

func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("startAccessingSecurityScopedResource"))
	return rv
}

func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("stopAccessingSecurityScopedResource"))
}

func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("checkPromisedItemIsReachableAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (u_ URL) GetPromisedItemResourceValueForKeyError(value *objc.Object, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("getPromisedItemResourceValue:forKey:error:"), unsafe.Pointer(value), key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) PromisedItemResourceValuesForKeysError(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, objc.GetSelector("promisedItemResourceValuesForKeys:error:"), keys, unsafe.Pointer(error))
	return rv
}

func (u_ URL) DataRepresentation() []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("dataRepresentation"))
	return rv
}

func (u_ URL) IsFileURL() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isFileURL"))
	return rv
}

func (u_ URL) AbsoluteString() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("absoluteString"))
	return rv
}

func (u_ URL) AbsoluteURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("absoluteURL"))
	return rv
}

func (u_ URL) BaseURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("baseURL"))
	return rv
}

func (u_ URL) FileSystemRepresentation() *byte {
	rv := objc.CallMethod[*byte](u_, objc.GetSelector("fileSystemRepresentation"))
	return rv
}

func (u_ URL) Fragment() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("fragment"))
	return rv
}

func (u_ URL) Host() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("host"))
	return rv
}

func (u_ URL) LastPathComponent() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("lastPathComponent"))
	return rv
}

func (u_ URL) Password() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("password"))
	return rv
}

func (u_ URL) Path() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("path"))
	return rv
}

func (u_ URL) PathComponents() []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("pathComponents"))
	return rv
}

func (u_ URL) PathExtension() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("pathExtension"))
	return rv
}

func (u_ URL) Port() Number {
	rv := objc.CallMethod[Number](u_, objc.GetSelector("port"))
	return rv
}

func (u_ URL) Query() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("query"))
	return rv
}

func (u_ URL) RelativePath() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("relativePath"))
	return rv
}

func (u_ URL) RelativeString() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("relativeString"))
	return rv
}

func (u_ URL) ResourceSpecifier() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("resourceSpecifier"))
	return rv
}

func (u_ URL) Scheme() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("scheme"))
	return rv
}

func (u_ URL) StandardizedURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("standardizedURL"))
	return rv
}

func (u_ URL) User() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("user"))
	return rv
}

func (u_ URL) FilePathURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("filePathURL"))
	return rv
}

func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByDeletingLastPathComponent"))
	return rv
}

func (u_ URL) URLByDeletingPathExtension() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByDeletingPathExtension"))
	return rv
}

func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByResolvingSymlinksInPath"))
	return rv
}

func (u_ URL) URLByStandardizingPath() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByStandardizingPath"))
	return rv
}

func (u_ URL) HasDirectoryPath() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("hasDirectoryPath"))
	return rv
}
