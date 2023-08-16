// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileManager] class.
var FileManagerClass = _FileManagerClass{objc.GetClass("NSFileManager")}

type _FileManagerClass struct {
	objc.Class
}

// An interface definition for the [FileManager] class.
type IFileManager interface {
	objc.IObject
	RemoveItemAtPathError(path string, error IError) bool
	IsExecutableFileAtPath(path string) bool
	SetUbiquitousItemAtURLDestinationURLError(flag bool, url IURL, destinationURL IURL, error IError) bool
	IsUbiquitousItemAtURL(url IURL) bool
	URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool, error IError) URL
	CreateFileAtPathContentsAttributes(path string, data []byte, attr map[FileAttributeKey]objc.IObject) bool
	EnumeratorAtPath(path string) DirectoryEnumerator
	ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []URLResourceKey, mask DirectoryEnumerationOptions, error IError) []URL
	MoveItemAtPathToPathError(srcPath string, dstPath string, error IError) bool
	CopyItemAtPathToPathError(srcPath string, dstPath string, error IError) bool
	GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship *URLRelationship, directory SearchPathDirectory, domainMask SearchPathDomainMask, url IURL, error IError) bool
	AttributesOfItemAtPathError(path string, error IError) map[FileAttributeKey]objc.Object
	URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL
	EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []URLResourceKey, mask DirectoryEnumerationOptions, handler func(url URL, error Error) bool) DirectoryEnumerator
	CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url IURL, createIntermediates bool, attributes map[FileAttributeKey]objc.IObject, error IError) bool
	LinkItemAtPathToPathError(srcPath string, dstPath string, error IError) bool
	URLForUbiquityContainerIdentifier(containerIdentifier string) URL
	RemoveItemAtURLError(URL IURL, error IError) bool
	IsReadableFileAtPath(path string) bool
	ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) URL
	CopyItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool
	TrashItemAtURLResultingItemURLError(url IURL, outResultingURL IURL, error IError) bool
	MoveItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool
	ComponentsToDisplayForPath(path string) []string
	StartDownloadingUbiquitousItemAtURLError(url IURL, error IError) bool
	IsDeletableFileAtPath(path string) bool
	ContentsAtPath(path string) []byte
	EvictUbiquitousItemAtURLError(url IURL, error IError) bool
	CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error IError) bool
	ChangeCurrentDirectoryPath(path string) bool
	FileSystemRepresentationWithPath(path string) *uint8
	MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []URLResourceKey, options VolumeEnumerationOptions) []URL
	StringWithFileSystemRepresentationLength(str *uint8, len uint) string
	UnmountVolumeAtURLOptionsCompletionHandler(url IURL, mask FileManagerUnmountOptions, completionHandler func(errorOrNil Error))
	LinkItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool
	IsWritableFileAtPath(path string) bool
	AttributesOfFileSystemForPathError(path string, error IError) map[FileAttributeKey]objc.Object
	DisplayNameAtPath(path string) string
	FileExistsAtPath(path string) bool
	SetAttributesOfItemAtPathError(attributes map[FileAttributeKey]objc.IObject, path string, error IError) bool
	GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler func(services map[FileProviderServiceName]FileProviderService, error Error))
	ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName string, options FileManagerItemReplacementOptions, resultingURL IURL, error IError) bool
	ContentsOfDirectoryAtPathError(path string, error IError) []string
	SubpathsAtPath(path string) []string
	DestinationOfSymbolicLinkAtPathError(path string, error IError) string
	SubpathsOfDirectoryAtPathError(path string, error IError) []string
	HomeDirectoryForUser(userName string) URL
	URLForPublishingUbiquitousItemAtURLExpirationDateError(url IURL, outDate IDate, error IError) URL
	ContentsEqualAtPathAndPath(path1 string, path2 string) bool
	HomeDirectoryForCurrentUser() URL
	CurrentDirectoryPath() string
	TemporaryDirectory() URL
	Delegate() FileManagerDelegateWrapper
	SetDelegate(value PFileManagerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	UbiquityIdentityToken() objc.Object
}

// A convenient interface to the contents of the file system, and the primary means of interacting with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager?language=objc
type FileManager struct {
	objc.Object
}

func FileManagerFrom(ptr unsafe.Pointer) FileManager {
	return FileManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileManagerClass) FileManagerWithAuthorization(authorization objc.IObject) FileManager {
	rv := objc.Call[FileManager](fc, objc.Sel("fileManagerWithAuthorization:"), objc.Ptr(authorization))
	return rv
}

// Initializes a file manager object that is authorized to perform privileged file system operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/3025773-filemanagerwithauthorization?language=objc
func FileManager_FileManagerWithAuthorization(authorization objc.IObject) FileManager {
	return FileManagerClass.FileManagerWithAuthorization(authorization)
}

func (fc _FileManagerClass) Alloc() FileManager {
	rv := objc.Call[FileManager](fc, objc.Sel("alloc"))
	return rv
}

func FileManager_Alloc() FileManager {
	return FileManagerClass.Alloc()
}

func (fc _FileManagerClass) New() FileManager {
	rv := objc.Call[FileManager](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileManager() FileManager {
	return FileManagerClass.New()
}

func (f_ FileManager) Init() FileManager {
	rv := objc.Call[FileManager](f_, objc.Sel("init"))
	return rv
}

// Removes the file or directory at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1408573-removeitematpath?language=objc
func (f_ FileManager) RemoveItemAtPathError(path string, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("removeItemAtPath:error:"), path, objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414159-isexecutablefileatpath?language=objc
func (f_ FileManager) IsExecutableFileAtPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isExecutableFileAtPath:"), path)
	return rv
}

// Indicates whether the item at the specified URL should be stored in iCloud. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413989-setubiquitous?language=objc
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url IURL, destinationURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, objc.Ptr(url), objc.Ptr(destinationURL), objc.Ptr(error))
	return rv
}

// Returns a Boolean indicating whether the item is targeted for storage in iCloud. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1410218-isubiquitousitematurl?language=objc
func (f_ FileManager) IsUbiquitousItemAtURL(url IURL) bool {
	rv := objc.Call[bool](f_, objc.Sel("isUbiquitousItemAtURL:"), objc.Ptr(url))
	return rv
}

// Locates and optionally creates the specified common directory in a domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1407693-urlfordirectory?language=objc
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool, error IError) URL {
	rv := objc.Call[URL](f_, objc.Sel("URLForDirectory:inDomain:appropriateForURL:create:error:"), directory, domain, objc.Ptr(url), shouldCreate, objc.Ptr(error))
	return rv
}

// Creates a file with the specified content and attributes at the given location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1410695-createfileatpath?language=objc
func (f_ FileManager) CreateFileAtPathContentsAttributes(path string, data []byte, attr map[FileAttributeKey]objc.IObject) bool {
	rv := objc.Call[bool](f_, objc.Sel("createFileAtPath:contents:attributes:"), path, data, attr)
	return rv
}

// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1408726-enumeratoratpath?language=objc
func (f_ FileManager) EnumeratorAtPath(path string) DirectoryEnumerator {
	rv := objc.Call[DirectoryEnumerator](f_, objc.Sel("enumeratorAtPath:"), path)
	return rv
}

// Performs a shallow search of the specified directory and returns URLs for the contained items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413768-contentsofdirectoryaturl?language=objc
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []URLResourceKey, mask DirectoryEnumerationOptions, error IError) []URL {
	rv := objc.Call[[]URL](f_, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), objc.Ptr(url), keys, mask, objc.Ptr(error))
	return rv
}

// Moves the file or directory at the specified path to a new location synchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413529-moveitematpath?language=objc
func (f_ FileManager) MoveItemAtPathToPathError(srcPath string, dstPath string, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("moveItemAtPath:toPath:error:"), srcPath, dstPath, objc.Ptr(error))
	return rv
}

// Copies the item at the specified path to a new location synchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1407903-copyitematpath?language=objc
func (f_ FileManager) CopyItemAtPathToPathError(srcPath string, dstPath string, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("copyItemAtPath:toPath:error:"), srcPath, dstPath, objc.Ptr(error))
	return rv
}

// Determines the type of relationship that exists between a system directory and the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1411439-getrelationship?language=objc
func (f_ FileManager) GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship *URLRelationship, directory SearchPathDirectory, domainMask SearchPathDomainMask, url IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("getRelationship:ofDirectory:inDomain:toItemAtURL:error:"), outRelationship, directory, domainMask, objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns the attributes of the item at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1410452-attributesofitematpath?language=objc
func (f_ FileManager) AttributesOfItemAtPathError(path string, error IError) map[FileAttributeKey]objc.Object {
	rv := objc.Call[map[FileAttributeKey]objc.Object](f_, objc.Sel("attributesOfItemAtPath:error:"), path, objc.Ptr(error))
	return rv
}

// Returns an array of URLs for the specified common directory in the requested domains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1407726-urlsfordirectory?language=objc
func (f_ FileManager) URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL {
	rv := objc.Call[[]URL](f_, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return rv
}

// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409571-enumeratoraturl?language=objc
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []URLResourceKey, mask DirectoryEnumerationOptions, handler func(url URL, error Error) bool) DirectoryEnumerator {
	rv := objc.Call[DirectoryEnumerator](f_, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), objc.Ptr(url), keys, mask, handler)
	return rv
}

// Creates a directory with the given attributes at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415371-createdirectoryaturl?language=objc
func (f_ FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url IURL, createIntermediates bool, attributes map[FileAttributeKey]objc.IObject, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("createDirectoryAtURL:withIntermediateDirectories:attributes:error:"), objc.Ptr(url), createIntermediates, attributes, objc.Ptr(error))
	return rv
}

// Creates a hard link between the items at the specified paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1411206-linkitematpath?language=objc
func (f_ FileManager) LinkItemAtPathToPathError(srcPath string, dstPath string, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("linkItemAtPath:toPath:error:"), srcPath, dstPath, objc.Ptr(error))
	return rv
}

// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1411653-urlforubiquitycontaineridentifie?language=objc
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string) URL {
	rv := objc.Call[URL](f_, objc.Sel("URLForUbiquityContainerIdentifier:"), containerIdentifier)
	return rv
}

// Removes the file or directory at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413590-removeitematurl?language=objc
func (f_ FileManager) RemoveItemAtURLError(URL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("removeItemAtURL:error:"), objc.Ptr(URL), objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1418292-isreadablefileatpath?language=objc
func (f_ FileManager) IsReadableFileAtPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isReadableFileAtPath:"), path)
	return rv
}

// Returns the container directory associated with the specified security application group identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1412643-containerurlforsecurityapplicati?language=objc
func (f_ FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) URL {
	rv := objc.Call[URL](f_, objc.Sel("containerURLForSecurityApplicationGroupIdentifier:"), groupIdentifier)
	return rv
}

// Copies the file at the specified URL to a new location synchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1412957-copyitematurl?language=objc
func (f_ FileManager) CopyItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("copyItemAtURL:toURL:error:"), objc.Ptr(srcURL), objc.Ptr(dstURL), objc.Ptr(error))
	return rv
}

// Moves an item to the trash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414306-trashitematurl?language=objc
func (f_ FileManager) TrashItemAtURLResultingItemURLError(url IURL, outResultingURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("trashItemAtURL:resultingItemURL:error:"), objc.Ptr(url), objc.Ptr(outResultingURL), objc.Ptr(error))
	return rv
}

// Moves the file or directory at the specified URL to a new location synchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414750-moveitematurl?language=objc
func (f_ FileManager) MoveItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("moveItemAtURL:toURL:error:"), objc.Ptr(srcURL), objc.Ptr(dstURL), objc.Ptr(error))
	return rv
}

// Returns an array of strings representing the user-visible components of a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413929-componentstodisplayforpath?language=objc
func (f_ FileManager) ComponentsToDisplayForPath(path string) []string {
	rv := objc.Call[[]string](f_, objc.Sel("componentsToDisplayForPath:"), path)
	return rv
}

// Starts downloading (if necessary) the specified item to the local system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1410377-startdownloadingubiquitousitemat?language=objc
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1408087-isdeletablefileatpath?language=objc
func (f_ FileManager) IsDeletableFileAtPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isDeletableFileAtPath:"), path)
	return rv
}

// Returns the contents of the file at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1407347-contentsatpath?language=objc
func (f_ FileManager) ContentsAtPath(path string) []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("contentsAtPath:"), path)
	return rv
}

// Removes the local copy of the specified item that’s stored in iCloud. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409696-evictubiquitousitematurl?language=objc
func (f_ FileManager) EvictUbiquitousItemAtURLError(url IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("evictUbiquitousItemAtURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Creates a symbolic link at the specified URL that points to an item at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414652-createsymboliclinkaturl?language=objc
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), objc.Ptr(url), objc.Ptr(destURL), objc.Ptr(error))
	return rv
}

// Changes the path of the current working directory to the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1412020-changecurrentdirectorypath?language=objc
func (f_ FileManager) ChangeCurrentDirectoryPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("changeCurrentDirectoryPath:"), path)
	return rv
}

// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1417057-filesystemrepresentationwithpath?language=objc
func (f_ FileManager) FileSystemRepresentationWithPath(path string) *uint8 {
	rv := objc.Call[*uint8](f_, objc.Sel("fileSystemRepresentationWithPath:"), path)
	return rv
}

// Returns an array of URLs that identify the mounted volumes available on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409626-mountedvolumeurlsincludingresour?language=objc
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []URLResourceKey, options VolumeEnumerationOptions) []URL {
	rv := objc.Call[[]URL](f_, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), propertyKeys, options)
	return rv
}

// Returns an NSString object whose contents are derived from the specified C-string path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409640-stringwithfilesystemrepresentati?language=objc
func (f_ FileManager) StringWithFileSystemRepresentationLength(str *uint8, len uint) string {
	rv := objc.Call[string](f_, objc.Sel("stringWithFileSystemRepresentation:length:"), str, len)
	return rv
}

// Starts the process of unmounting the specified volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409917-unmountvolumeaturl?language=objc
func (f_ FileManager) UnmountVolumeAtURLOptionsCompletionHandler(url IURL, mask FileManagerUnmountOptions, completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("unmountVolumeAtURL:options:completionHandler:"), objc.Ptr(url), mask, completionHandler)
}

// Creates a hard link between the items at the specified URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414456-linkitematurl?language=objc
func (f_ FileManager) LinkItemAtURLToURLError(srcURL IURL, dstURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("linkItemAtURL:toURL:error:"), objc.Ptr(srcURL), objc.Ptr(dstURL), objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1416680-iswritablefileatpath?language=objc
func (f_ FileManager) IsWritableFileAtPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isWritableFileAtPath:"), path)
	return rv
}

// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1411896-attributesoffilesystemforpath?language=objc
func (f_ FileManager) AttributesOfFileSystemForPathError(path string, error IError) map[FileAttributeKey]objc.Object {
	rv := objc.Call[map[FileAttributeKey]objc.Object](f_, objc.Sel("attributesOfFileSystemForPath:error:"), path, objc.Ptr(error))
	return rv
}

// Returns the display name of the file or directory at a specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409751-displaynameatpath?language=objc
func (f_ FileManager) DisplayNameAtPath(path string) string {
	rv := objc.Call[string](f_, objc.Sel("displayNameAtPath:"), path)
	return rv
}

// Returns a Boolean value that indicates whether a file or directory exists at a specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415645-fileexistsatpath?language=objc
func (f_ FileManager) FileExistsAtPath(path string) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileExistsAtPath:"), path)
	return rv
}

// Sets the attributes of the specified file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413667-setattributes?language=objc
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes map[FileAttributeKey]objc.IObject, path string, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, path, objc.Ptr(error))
	return rv
}

// Returns the services provided by the File Provider extension that manages the item at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/2915262-getfileproviderservicesforitemat?language=objc
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler func(services map[FileProviderServiceName]FileProviderService, error Error)) {
	objc.Call[objc.Void](f_, objc.Sel("getFileProviderServicesForItemAtURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1412432-replaceitematurl?language=objc
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName string, options FileManagerItemReplacementOptions, resultingURL IURL, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), objc.Ptr(originalItemURL), objc.Ptr(newItemURL), backupItemName, options, objc.Ptr(resultingURL), objc.Ptr(error))
	return rv
}

// Performs a shallow search of the specified directory and returns the paths of any contained items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1414584-contentsofdirectoryatpath?language=objc
func (f_ FileManager) ContentsOfDirectoryAtPathError(path string, error IError) []string {
	rv := objc.Call[[]string](f_, objc.Sel("contentsOfDirectoryAtPath:error:"), path, objc.Ptr(error))
	return rv
}

// Returns an array of strings identifying the paths for all items in the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1413742-subpathsatpath?language=objc
func (f_ FileManager) SubpathsAtPath(path string) []string {
	rv := objc.Call[[]string](f_, objc.Sel("subpathsAtPath:"), path)
	return rv
}

// Returns the path of the item pointed to by a symbolic link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415161-destinationofsymboliclinkatpath?language=objc
func (f_ FileManager) DestinationOfSymbolicLinkAtPathError(path string, error IError) string {
	rv := objc.Call[string](f_, objc.Sel("destinationOfSymbolicLinkAtPath:error:"), path, objc.Ptr(error))
	return rv
}

// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1417353-subpathsofdirectoryatpath?language=objc
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path string, error IError) []string {
	rv := objc.Call[[]string](f_, objc.Sel("subpathsOfDirectoryAtPath:error:"), path, objc.Ptr(error))
	return rv
}

// Returns the home directory for the specified user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1642853-homedirectoryforuser?language=objc
func (f_ FileManager) HomeDirectoryForUser(userName string) URL {
	rv := objc.Call[URL](f_, objc.Sel("homeDirectoryForUser:"), userName)
	return rv
}

// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1411577-urlforpublishingubiquitousitemat?language=objc
func (f_ FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url IURL, outDate IDate, error IError) URL {
	rv := objc.Call[URL](f_, objc.Sel("URLForPublishingUbiquitousItemAtURL:expirationDate:error:"), objc.Ptr(url), objc.Ptr(outDate), objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415275-contentsequalatpath?language=objc
func (f_ FileManager) ContentsEqualAtPathAndPath(path1 string, path2 string) bool {
	rv := objc.Call[bool](f_, objc.Sel("contentsEqualAtPath:andPath:"), path1, path2)
	return rv
}

// The home directory for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1642807-homedirectoryforcurrentuser?language=objc
func (f_ FileManager) HomeDirectoryForCurrentUser() URL {
	rv := objc.Call[URL](f_, objc.Sel("homeDirectoryForCurrentUser"))
	return rv
}

// The path to the program’s current directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1410766-currentdirectorypath?language=objc
func (f_ FileManager) CurrentDirectoryPath() string {
	rv := objc.Call[string](f_, objc.Sel("currentDirectoryPath"))
	return rv
}

// The temporary directory for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1642996-temporarydirectory?language=objc
func (f_ FileManager) TemporaryDirectory() URL {
	rv := objc.Call[URL](f_, objc.Sel("temporaryDirectory"))
	return rv
}

// The delegate of the file manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415163-delegate?language=objc
func (f_ FileManager) Delegate() FileManagerDelegateWrapper {
	rv := objc.Call[FileManagerDelegateWrapper](f_, objc.Sel("delegate"))
	return rv
}

// The delegate of the file manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415163-delegate?language=objc
func (f_ FileManager) SetDelegate(value PFileManagerDelegate) {
	po0 := objc.WrapAsProtocol("NSFileManagerDelegate", value)
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), po0)
}

// The delegate of the file manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1415163-delegate?language=objc
func (f_ FileManager) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// An opaque token that represents the current user’s iCloud Drive Documents identity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1408036-ubiquityidentitytoken?language=objc
func (f_ FileManager) UbiquityIdentityToken() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("ubiquityIdentityToken"))
	return rv
}

// The shared file manager object for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409234-defaultmanager?language=objc
func (fc _FileManagerClass) DefaultManager() FileManager {
	rv := objc.Call[FileManager](fc, objc.Sel("defaultManager"))
	return rv
}

// The shared file manager object for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanager/1409234-defaultmanager?language=objc
func FileManager_DefaultManager() FileManager {
	return FileManagerClass.DefaultManager()
}
