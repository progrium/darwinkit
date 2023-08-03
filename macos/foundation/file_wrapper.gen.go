// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	objc.Class
}

type IFileWrapper interface {
	objc.IObject
	AddFileWrapper(child IFileWrapper) string
	RemoveFileWrapper(child IFileWrapper)
	AddRegularFileWithContentsPreferredFilename(data []byte, fileName string) string
	KeyForFileWrapper(child IFileWrapper) string
	MatchesContentsOfURL(url IURL) bool
	ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError *Error) bool
	WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool
	IsRegularFile() bool
	IsDirectory() bool
	IsSymbolicLink() bool
	FileWrappers() map[string]FileWrapper
	SymbolicLinkDestinationURL() URL
	SerializedRepresentation() []byte
	Filename() string
	SetFilename(value string)
	PreferredFilename() string
	SetPreferredFilename(value string)
	FileAttributes() map[string]objc.Object
	SetFileAttributes(value map[string]objc.IObject)
	RegularFileContents() []byte
}

type FileWrapper struct {
	objc.Object
}

func MakeFileWrapper(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		Object: objc.MakeObject(ptr),
	}
}

func (f_ FileWrapper) InitWithURLOptionsError(url IURL, options FileWrapperReadingOptions, outError *Error) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initWithURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(outError))
	return rv
}

func FileWrapper_InitWithURLOptionsError(url IURL, options FileWrapperReadingOptions, outError *Error) FileWrapper {
	return FileWrapperClass.Alloc().InitWithURLOptionsError(url, options, outError)
}

func (f_ FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return rv
}

func FileWrapper_InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	return FileWrapperClass.Alloc().InitDirectoryWithFileWrappers(childrenByPreferredName)
}

func (f_ FileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initRegularFileWithContents:"), contents)
	return rv
}

func FileWrapper_InitRegularFileWithContents(contents []byte) FileWrapper {
	return FileWrapperClass.Alloc().InitRegularFileWithContents(contents)
}

func (f_ FileWrapper) InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initSymbolicLinkWithDestinationURL:"), objc.ExtractPtr(url))
	return rv
}

func FileWrapper_InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	return FileWrapperClass.Alloc().InitSymbolicLinkWithDestinationURL(url)
}

func (f_ FileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initWithSerializedRepresentation:"), serializeRepresentation)
	return rv
}

func FileWrapper_InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	return FileWrapperClass.Alloc().InitWithSerializedRepresentation(serializeRepresentation)
}

func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.GetSelector("alloc"))
	return rv
}

func FileWrapper_Alloc() FileWrapper {
	return FileWrapperClass.Alloc()
}

func (fc _FileWrapperClass) New() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFileWrapper() FileWrapper {
	return FileWrapperClass.New()
}

func FileWrapper_New() FileWrapper {
	return FileWrapperClass.New()
}

func (f_ FileWrapper) Init() FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("init"))
	return rv
}

func FileWrapper_Init() FileWrapper {
	return FileWrapperClass.Alloc().Init()
}

func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addFileWrapper:"), objc.ExtractPtr(child))
	return rv
}

func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("removeFileWrapper:"), objc.ExtractPtr(child))
}

func (f_ FileWrapper) AddRegularFileWithContentsPreferredFilename(data []byte, fileName string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addRegularFileWithContents:preferredFilename:"), data, fileName)
	return rv
}

func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("keyForFileWrapper:"), objc.ExtractPtr(child))
	return rv
}

func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("matchesContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

func (f_ FileWrapper) ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("readFromURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("writeToURL:options:originalContentsURL:error:"), objc.ExtractPtr(url), options, objc.ExtractPtr(originalContentsURL), unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isRegularFile"))
	return rv
}

func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isDirectory"))
	return rv
}

func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isSymbolicLink"))
	return rv
}

func (f_ FileWrapper) FileWrappers() map[string]FileWrapper {
	rv := objc.CallMethod[map[string]FileWrapper](f_, objc.GetSelector("fileWrappers"))
	return rv
}

func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := objc.CallMethod[URL](f_, objc.GetSelector("symbolicLinkDestinationURL"))
	return rv
}

func (f_ FileWrapper) SerializedRepresentation() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.GetSelector("serializedRepresentation"))
	return rv
}

func (f_ FileWrapper) Filename() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("filename"))
	return rv
}

func (f_ FileWrapper) SetFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFilename:"), value)
}

func (f_ FileWrapper) PreferredFilename() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("preferredFilename"))
	return rv
}

func (f_ FileWrapper) SetPreferredFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setPreferredFilename:"), value)
}

func (f_ FileWrapper) FileAttributes() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, objc.GetSelector("fileAttributes"))
	return rv
}

func (f_ FileWrapper) SetFileAttributes(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFileAttributes:"), value)
}

func (f_ FileWrapper) RegularFileContents() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.GetSelector("regularFileContents"))
	return rv
}
