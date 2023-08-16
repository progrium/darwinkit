// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileVersion] class.
var FileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}

type _FileVersionClass struct {
	objc.Class
}

// An interface definition for the [FileVersion] class.
type IFileVersion interface {
	objc.IObject
	ReplaceItemAtURLOptionsError(url IURL, options FileVersionReplacingOptions, error IError) URL
	RemoveAndReturnError(outError IError) bool
	IsResolved() bool
	SetResolved(value bool)
	IsDiscardable() bool
	SetDiscardable(value bool)
	ModificationDate() Date
	OriginatorNameComponents() PersonNameComponents
	LocalizedName() string
	LocalizedNameOfSavingComputer() string
	URL() URL
	IsConflict() bool
	PersistentIdentifier() CodingWrapper
	HasThumbnail() bool
	HasLocalContents() bool
}

// A snapshot of a file at a specific point in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion?language=objc
type FileVersion struct {
	objc.Object
}

func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileVersionClass) Alloc() FileVersion {
	rv := objc.Call[FileVersion](fc, objc.Sel("alloc"))
	return rv
}

func FileVersion_Alloc() FileVersion {
	return FileVersionClass.Alloc()
}

func (fc _FileVersionClass) New() FileVersion {
	rv := objc.Call[FileVersion](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileVersion() FileVersion {
	return FileVersionClass.New()
}

func (f_ FileVersion) Init() FileVersion {
	rv := objc.Call[FileVersion](f_, objc.Sel("init"))
	return rv
}

// Returns the version of the file that has the specified persistent ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1415443-versionofitematurl?language=objc
func (fc _FileVersionClass) VersionOfItemAtURLForPersistentIdentifier(url IURL, persistentIdentifier objc.IObject) FileVersion {
	rv := objc.Call[FileVersion](fc, objc.Sel("versionOfItemAtURL:forPersistentIdentifier:"), objc.Ptr(url), persistentIdentifier)
	return rv
}

// Returns the version of the file that has the specified persistent ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1415443-versionofitematurl?language=objc
func FileVersion_VersionOfItemAtURLForPersistentIdentifier(url IURL, persistentIdentifier objc.IObject) FileVersion {
	return FileVersionClass.VersionOfItemAtURLForPersistentIdentifier(url, persistentIdentifier)
}

// Returns the most recent version object for the file at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1412963-currentversionofitematurl?language=objc
func (fc _FileVersionClass) CurrentVersionOfItemAtURL(url IURL) FileVersion {
	rv := objc.Call[FileVersion](fc, objc.Sel("currentVersionOfItemAtURL:"), objc.Ptr(url))
	return rv
}

// Returns the most recent version object for the file at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1412963-currentversionofitematurl?language=objc
func FileVersion_CurrentVersionOfItemAtURL(url IURL) FileVersion {
	return FileVersionClass.CurrentVersionOfItemAtURL(url)
}

// Removes all versions of a file, except the current one, from the version store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1411537-removeotherversionsofitematurl?language=objc
func (fc _FileVersionClass) RemoveOtherVersionsOfItemAtURLError(url IURL, outError IError) bool {
	rv := objc.Call[bool](fc, objc.Sel("removeOtherVersionsOfItemAtURL:error:"), objc.Ptr(url), objc.Ptr(outError))
	return rv
}

// Removes all versions of a file, except the current one, from the version store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1411537-removeotherversionsofitematurl?language=objc
func FileVersion_RemoveOtherVersionsOfItemAtURLError(url IURL, outError IError) bool {
	return FileVersionClass.RemoveOtherVersionsOfItemAtURLError(url, outError)
}

// Creates and returns a temporary directory to use for saving the contents of the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1411906-temporarydirectoryurlfornewversi?language=objc
func (fc _FileVersionClass) TemporaryDirectoryURLForNewVersionOfItemAtURL(url IURL) URL {
	rv := objc.Call[URL](fc, objc.Sel("temporaryDirectoryURLForNewVersionOfItemAtURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns a temporary directory to use for saving the contents of the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1411906-temporarydirectoryurlfornewversi?language=objc
func FileVersion_TemporaryDirectoryURLForNewVersionOfItemAtURL(url IURL) URL {
	return FileVersionClass.TemporaryDirectoryURLForNewVersionOfItemAtURL(url)
}

// Returns all versions of the specified file except the current version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1418163-otherversionsofitematurl?language=objc
func (fc _FileVersionClass) OtherVersionsOfItemAtURL(url IURL) []FileVersion {
	rv := objc.Call[[]FileVersion](fc, objc.Sel("otherVersionsOfItemAtURL:"), objc.Ptr(url))
	return rv
}

// Returns all versions of the specified file except the current version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1418163-otherversionsofitematurl?language=objc
func FileVersion_OtherVersionsOfItemAtURL(url IURL) []FileVersion {
	return FileVersionClass.OtherVersionsOfItemAtURL(url)
}

// Replace the contents of the specified file with the contents of the current version’s file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1412297-replaceitematurl?language=objc
func (f_ FileVersion) ReplaceItemAtURLOptionsError(url IURL, options FileVersionReplacingOptions, error IError) URL {
	rv := objc.Call[URL](f_, objc.Sel("replaceItemAtURL:options:error:"), objc.Ptr(url), options, objc.Ptr(error))
	return rv
}

// Creates a version of the file at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1413326-addversionofitematurl?language=objc
func (fc _FileVersionClass) AddVersionOfItemAtURLWithContentsOfURLOptionsError(url IURL, contentsURL IURL, options FileVersionAddingOptions, outError IError) FileVersion {
	rv := objc.Call[FileVersion](fc, objc.Sel("addVersionOfItemAtURL:withContentsOfURL:options:error:"), objc.Ptr(url), objc.Ptr(contentsURL), options, objc.Ptr(outError))
	return rv
}

// Creates a version of the file at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1413326-addversionofitematurl?language=objc
func FileVersion_AddVersionOfItemAtURLWithContentsOfURLOptionsError(url IURL, contentsURL IURL, options FileVersionAddingOptions, outError IError) FileVersion {
	return FileVersionClass.AddVersionOfItemAtURLWithContentsOfURLOptionsError(url, contentsURL, options, outError)
}

// Remove this version object and its associated file from the version store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1407486-removeandreturnerror?language=objc
func (f_ FileVersion) RemoveAndReturnError(outError IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("removeAndReturnError:"), objc.Ptr(outError))
	return rv
}

// Returns an array of version objects that are currently in conflict for the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1417854-unresolvedconflictversionsofitem?language=objc
func (fc _FileVersionClass) UnresolvedConflictVersionsOfItemAtURL(url IURL) []FileVersion {
	rv := objc.Call[[]FileVersion](fc, objc.Sel("unresolvedConflictVersionsOfItemAtURL:"), objc.Ptr(url))
	return rv
}

// Returns an array of version objects that are currently in conflict for the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1417854-unresolvedconflictversionsofitem?language=objc
func FileVersion_UnresolvedConflictVersionsOfItemAtURL(url IURL) []FileVersion {
	return FileVersionClass.UnresolvedConflictVersionsOfItemAtURL(url)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1416051-getnonlocalversionsofitematurl?language=objc
func (fc _FileVersionClass) GetNonlocalVersionsOfItemAtURLCompletionHandler(url IURL, completionHandler func(nonlocalFileVersions []FileVersion, error Error)) {
	objc.Call[objc.Void](fc, objc.Sel("getNonlocalVersionsOfItemAtURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1416051-getnonlocalversionsofitematurl?language=objc
func FileVersion_GetNonlocalVersionsOfItemAtURLCompletionHandler(url IURL, completionHandler func(nonlocalFileVersions []FileVersion, error Error)) {
	FileVersionClass.GetNonlocalVersionsOfItemAtURLCompletionHandler(url, completionHandler)
}

// A Boolean value that indicates the version object is not in conflict (YES) or is in conflict (NO). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1414906-resolved?language=objc
func (f_ FileVersion) IsResolved() bool {
	rv := objc.Call[bool](f_, objc.Sel("isResolved"))
	return rv
}

// A Boolean value that indicates the version object is not in conflict (YES) or is in conflict (NO). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1414906-resolved?language=objc
func (f_ FileVersion) SetResolved(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setResolved:"), value)
}

// A Boolean value that specifies whether the system can delete the associated file at some future time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1409988-discardable?language=objc
func (f_ FileVersion) IsDiscardable() bool {
	rv := objc.Call[bool](f_, objc.Sel("isDiscardable"))
	return rv
}

// A Boolean value that specifies whether the system can delete the associated file at some future time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1409988-discardable?language=objc
func (f_ FileVersion) SetDiscardable(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setDiscardable:"), value)
}

// The modification date of the version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1411506-modificationdate?language=objc
func (f_ FileVersion) ModificationDate() Date {
	rv := objc.Call[Date](f_, objc.Sel("modificationDate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1643271-originatornamecomponents?language=objc
func (f_ FileVersion) OriginatorNameComponents() PersonNameComponents {
	rv := objc.Call[PersonNameComponents](f_, objc.Sel("originatorNameComponents"))
	return rv
}

// The string containing the user-presentable name of the file version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1413855-localizedname?language=objc
func (f_ FileVersion) LocalizedName() string {
	rv := objc.Call[string](f_, objc.Sel("localizedName"))
	return rv
}

// The user-presentable name of the computer on which the revision was saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1408866-localizednameofsavingcomputer?language=objc
func (f_ FileVersion) LocalizedNameOfSavingComputer() string {
	rv := objc.Call[string](f_, objc.Sel("localizedNameOfSavingComputer"))
	return rv
}

// The URL identifying the location of the file associated with the file version object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1418131-url?language=objc
func (f_ FileVersion) URL() URL {
	rv := objc.Call[URL](f_, objc.Sel("URL"))
	return rv
}

// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1409277-conflict?language=objc
func (f_ FileVersion) IsConflict() bool {
	rv := objc.Call[bool](f_, objc.Sel("isConflict"))
	return rv
}

// The identifier for this version of the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1407948-persistentidentifier?language=objc
func (f_ FileVersion) PersistentIdentifier() CodingWrapper {
	rv := objc.Call[CodingWrapper](f_, objc.Sel("persistentIdentifier"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1409471-hasthumbnail?language=objc
func (f_ FileVersion) HasThumbnail() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasThumbnail"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/1412014-haslocalcontents?language=objc
func (f_ FileVersion) HasLocalContents() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasLocalContents"))
	return rv
}
