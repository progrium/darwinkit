// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileWrapper] class.
var FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	objc.Class
}

// An interface definition for the [FileWrapper] class.
type IFileWrapper interface {
	objc.IObject
	AddRegularFileWithContentsPreferredFilename(data []byte, fileName string) string
	KeyForFileWrapper(child IFileWrapper) string
	MatchesContentsOfURL(url IURL) bool
	RemoveFileWrapper(child IFileWrapper)
	ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) bool
	WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError IError) bool
	AddFileWrapper(child IFileWrapper) string
	Filename() string
	SetFilename(value string)
	IsSymbolicLink() bool
	FileWrappers() map[string]FileWrapper
	Icon() objc.Object
	SetIcon(value objc.IObject)
	IsDirectory() bool
	PreferredFilename() string
	SetPreferredFilename(value string)
	IsRegularFile() bool
	SymbolicLinkDestinationURL() URL
	FileAttributes() map[string]objc.Object
	SetFileAttributes(value map[string]objc.IObject)
	SerializedRepresentation() []byte
	RegularFileContents() []byte
}

// A representation of a node (a file, directory, or symbolic link) in the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper?language=objc
type FileWrapper struct {
	objc.Object
}

func FileWrapperFrom(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("initRegularFileWithContents:"), contents)
	return rv
}

// Initializes the receiver as a regular-file file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1409508-initregularfilewithcontents?language=objc
func FileWrapper_InitRegularFileWithContents(contents []byte) FileWrapper {
	return FileWrapperClass.Alloc().InitRegularFileWithContents(contents)
}

func (f_ FileWrapper) InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("initSymbolicLinkWithDestinationURL:"), objc.Ptr(url))
	return rv
}

// Initializes the receiver as a symbolic-link file wrapper that links to a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415098-initsymboliclinkwithdestinationu?language=objc
func FileWrapper_InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	return FileWrapperClass.Alloc().InitSymbolicLinkWithDestinationURL(url)
}

func (f_ FileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("initWithSerializedRepresentation:"), serializeRepresentation)
	return rv
}

// Initializes the receiver as a regular-file file wrapper from given serialized data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1407515-initwithserializedrepresentation?language=objc
func FileWrapper_InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	return FileWrapperClass.Alloc().InitWithSerializedRepresentation(serializeRepresentation)
}

func (f_ FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return rv
}

// Initializes the receiver as a directory file wrapper, with a given file-wrapper list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415121-initdirectorywithfilewrappers?language=objc
func FileWrapper_InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	return FileWrapperClass.Alloc().InitDirectoryWithFileWrappers(childrenByPreferredName)
}

func (f_ FileWrapper) InitWithURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("initWithURL:options:error:"), objc.Ptr(url), options, objc.Ptr(outError))
	return rv
}

// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415658-initwithurl?language=objc
func FileWrapper_InitWithURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) FileWrapper {
	return FileWrapperClass.Alloc().InitWithURLOptionsError(url, options, outError)
}

func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.Call[FileWrapper](fc, objc.Sel("alloc"))
	return rv
}

func FileWrapper_Alloc() FileWrapper {
	return FileWrapperClass.Alloc()
}

func (fc _FileWrapperClass) New() FileWrapper {
	rv := objc.Call[FileWrapper](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileWrapper() FileWrapper {
	return FileWrapperClass.New()
}

func (f_ FileWrapper) Init() FileWrapper {
	rv := objc.Call[FileWrapper](f_, objc.Sel("init"))
	return rv
}

// Creates a regular-file file wrapper with the given contents and adds it to the receiver, which must be a directory file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1418374-addregularfilewithcontents?language=objc
func (f_ FileWrapper) AddRegularFileWithContentsPreferredFilename(data []byte, fileName string) string {
	rv := objc.Call[string](f_, objc.Sel("addRegularFileWithContents:preferredFilename:"), data, fileName)
	return rv
}

// Returns the dictionary key used by a directory to identify a given file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1407541-keyforfilewrapper?language=objc
func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) string {
	rv := objc.Call[string](f_, objc.Sel("keyForFileWrapper:"), objc.Ptr(child))
	return rv
}

// Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1408360-matchescontentsofurl?language=objc
func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.Call[bool](f_, objc.Sel("matchesContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Removes a child file wrapper from the receiver, which must be a directory file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1417343-removefilewrapper?language=objc
func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	objc.Call[objc.Void](f_, objc.Sel("removeFileWrapper:"), objc.Ptr(child))
}

// Recursively rereads the entire contents of a file wrapper from the specified location on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1411645-readfromurl?language=objc
func (f_ FileWrapper) ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("readFromURL:options:error:"), objc.Ptr(url), options, objc.Ptr(outError))
	return rv
}

// Recursively writes the entire contents of a file wrapper to a given file-system URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415981-writetourl?language=objc
func (f_ FileWrapper) WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("writeToURL:options:originalContentsURL:error:"), objc.Ptr(url), options, objc.Ptr(originalContentsURL), objc.Ptr(outError))
	return rv
}

// Adds a child file wrapper to the receiver, which must be a directory file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415067-addfilewrapper?language=objc
func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) string {
	rv := objc.Call[string](f_, objc.Sel("addFileWrapper:"), objc.Ptr(child))
	return rv
}

// The filename of the file wrapper object [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1416684-filename?language=objc
func (f_ FileWrapper) Filename() string {
	rv := objc.Call[string](f_, objc.Sel("filename"))
	return rv
}

// The filename of the file wrapper object [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1416684-filename?language=objc
func (f_ FileWrapper) SetFilename(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setFilename:"), value)
}

// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1408125-symboliclink?language=objc
func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.Call[bool](f_, objc.Sel("isSymbolicLink"))
	return rv
}

// The file wrappers contained by a directory file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1409437-filewrappers?language=objc
func (f_ FileWrapper) FileWrappers() map[string]FileWrapper {
	rv := objc.Call[map[string]FileWrapper](f_, objc.Sel("fileWrappers"))
	return rv
}

// The icon that represents the file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1413123-icon?language=objc
func (f_ FileWrapper) Icon() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("icon"))
	return rv
}

// The icon that represents the file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1413123-icon?language=objc
func (f_ FileWrapper) SetIcon(value objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setIcon:"), objc.Ptr(value))
}

// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1409030-directory?language=objc
func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.Call[bool](f_, objc.Sel("isDirectory"))
	return rv
}

// The preferred filename for the file wrapper object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1409368-preferredfilename?language=objc
func (f_ FileWrapper) PreferredFilename() string {
	rv := objc.Call[string](f_, objc.Sel("preferredFilename"))
	return rv
}

// The preferred filename for the file wrapper object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1409368-preferredfilename?language=objc
func (f_ FileWrapper) SetPreferredFilename(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setPreferredFilename:"), value)
}

// This property contains a boolean value that indicates whether the file wrapper object is a regular-file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1415680-regularfile?language=objc
func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.Call[bool](f_, objc.Sel("isRegularFile"))
	return rv
}

// The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1408364-symboliclinkdestinationurl?language=objc
func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := objc.Call[URL](f_, objc.Sel("symbolicLinkDestinationURL"))
	return rv
}

// A dictionary of file attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1412745-fileattributes?language=objc
func (f_ FileWrapper) FileAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](f_, objc.Sel("fileAttributes"))
	return rv
}

// A dictionary of file attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1412745-fileattributes?language=objc
func (f_ FileWrapper) SetFileAttributes(value map[string]objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setFileAttributes:"), value)
}

// The contents of the file wrapper as an opaque data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1412119-serializedrepresentation?language=objc
func (f_ FileWrapper) SerializedRepresentation() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("serializedRepresentation"))
	return rv
}

// The contents of the file-system node associated with a regular-file file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilewrapper/1410178-regularfilecontents?language=objc
func (f_ FileWrapper) RegularFileContents() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("regularFileContents"))
	return rv
}
