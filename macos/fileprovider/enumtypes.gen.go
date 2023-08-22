// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

// Options for creating items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidercreateitemoptions?language=objc
type FileProviderCreateItemOptions uint

const (
	FileProviderCreateItemDeletionConflicted FileProviderCreateItemOptions = 2
	FileProviderCreateItemMayAlreadyExist    FileProviderCreateItemOptions = 1
)

// Options for deleting items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdeleteitemoptions?language=objc
type FileProviderDeleteItemOptions uint

const (
	FileProviderDeleteItemRecursive FileProviderDeleteItemOptions = 1
)

// A unique identifier for a file provider's domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainidentifier?language=objc
type FileProviderDomainIdentifier string

// A mode indicating how the system handles user data when removing a domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainremovalmode?language=objc
type FileProviderDomainRemovalMode int

const (
	FileProviderDomainRemovalModePreserveDirtyUserData      FileProviderDomainRemovalMode = 1
	FileProviderDomainRemovalModePreserveDownloadedUserData FileProviderDomainRemovalMode = 2
	FileProviderDomainRemovalModeRemoveAll                  FileProviderDomainRemovalMode = 0
)

// Modes that modify the system’s behavior while testing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomaintestingmodes?language=objc
type FileProviderDomainTestingModes uint

const (
	FileProviderDomainTestingModeAlwaysEnabled FileProviderDomainTestingModes = 1
	FileProviderDomainTestingModeInteractive   FileProviderDomainTestingModes = 2
)

// The error codes for the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidererrorcode?language=objc
type FileProviderErrorCode int

const (
	FileProviderErrorCannotSynchronize            FileProviderErrorCode = -2005
	FileProviderErrorDeletionRejected             FileProviderErrorCode = -1006
	FileProviderErrorDirectoryNotEmpty            FileProviderErrorCode = -1007
	FileProviderErrorFilenameCollision            FileProviderErrorCode = -1001
	FileProviderErrorInsufficientQuota            FileProviderErrorCode = -1003
	FileProviderErrorNewerExtensionVersionFound   FileProviderErrorCode = -2004
	FileProviderErrorNoSuchItem                   FileProviderErrorCode = -1005
	FileProviderErrorNonEvictable                 FileProviderErrorCode = -2008
	FileProviderErrorNonEvictableChildren         FileProviderErrorCode = -2006
	FileProviderErrorNotAuthenticated             FileProviderErrorCode = -1000
	FileProviderErrorOlderExtensionVersionRunning FileProviderErrorCode = -2003
	FileProviderErrorPageExpired                  FileProviderErrorCode = -1002
	FileProviderErrorProviderNotFound             FileProviderErrorCode = -2001
	FileProviderErrorProviderTranslocated         FileProviderErrorCode = -2002
	FileProviderErrorServerUnreachable            FileProviderErrorCode = -1004
	FileProviderErrorSyncAnchorExpired            FileProviderErrorCode = -1002
	FileProviderErrorUnsyncedEdits                FileProviderErrorCode = -2007
)

// An identifier for custom actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderextensionactionidentifier?language=objc
type FileProviderExtensionActionIdentifier string

// Options for fetching a range of data from a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderfetchcontentsoptions?language=objc
type FileProviderFetchContentsOptions uint

// Flags that define an item’s on-disk properties and its appearance in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderfilesystemflags?language=objc
type FileProviderFileSystemFlags uint

const (
	FileProviderFileSystemHidden              FileProviderFileSystemFlags = 8
	FileProviderFileSystemPathExtensionHidden FileProviderFileSystemFlags = 16
	FileProviderFileSystemUserExecutable      FileProviderFileSystemFlags = 1
	FileProviderFileSystemUserReadable        FileProviderFileSystemFlags = 2
	FileProviderFileSystemUserWritable        FileProviderFileSystemFlags = 4
)

// An item’s capabilities, which define the actions that the user can perform in the document browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemcapabilities?language=objc
type FileProviderItemCapabilities uint

const (
	FileProviderItemCapabilitiesAllowsAddingSubItems     FileProviderItemCapabilities = 2
	FileProviderItemCapabilitiesAllowsAll                FileProviderItemCapabilities = 63
	FileProviderItemCapabilitiesAllowsContentEnumerating FileProviderItemCapabilities = 1
	FileProviderItemCapabilitiesAllowsDeleting           FileProviderItemCapabilities = 32
	FileProviderItemCapabilitiesAllowsEvicting           FileProviderItemCapabilities = 64
	FileProviderItemCapabilitiesAllowsExcludingFromSync  FileProviderItemCapabilities = 128
	FileProviderItemCapabilitiesAllowsReading            FileProviderItemCapabilities = 1
	FileProviderItemCapabilitiesAllowsRenaming           FileProviderItemCapabilities = 8
	FileProviderItemCapabilitiesAllowsReparenting        FileProviderItemCapabilities = 4
	FileProviderItemCapabilitiesAllowsTrashing           FileProviderItemCapabilities = 16
	FileProviderItemCapabilitiesAllowsWriting            FileProviderItemCapabilities = 2
)

// A decoration identifier defined in the File Provider extension’s information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemdecorationidentifier?language=objc
type FileProviderItemDecorationIdentifier string

// Fields that specify which of the item’s properties have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemfields?language=objc
type FileProviderItemFields uint

const (
	FileProviderItemContentModificationDate FileProviderItemFields = 128
	FileProviderItemContents                FileProviderItemFields = 1
	FileProviderItemCreationDate            FileProviderItemFields = 64
	FileProviderItemExtendedAttributes      FileProviderItemFields = 512
	FileProviderItemFavoriteRank            FileProviderItemFields = 32
	FileProviderItemFileSystemFlags         FileProviderItemFields = 256
	FileProviderItemFilename                FileProviderItemFields = 2
	FileProviderItemLastUsedDate            FileProviderItemFields = 8
	FileProviderItemParentItemIdentifier    FileProviderItemFields = 4
	FileProviderItemTagData                 FileProviderItemFields = 16
	FileProviderItemTypeAndCreator          FileProviderItemFields = 1024
)

// A unique identifier for an item managed by the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemidentifier?language=objc
type FileProviderItemIdentifier string

const (
	FileProviderRootContainerItemIdentifier       FileProviderItemIdentifier = "NSFileProviderRootContainerItemIdentifier"
	FileProviderTrashContainerItemIdentifier      FileProviderItemIdentifier = "NSFileProviderTrashContainerItemIdentifier"
	FileProviderWorkingSetContainerItemIdentifier FileProviderItemIdentifier = "NSFileProviderWorkingSetContainerItemIdentifier"
)

// Options for disconnecting a domain from the extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanagerdisconnectionoptions?language=objc
type FileProviderManagerDisconnectionOptions uint

const (
	FileProviderManagerDisconnectionOptionsTemporary FileProviderManagerDisconnectionOptions = 1
)

// Flags that provides additional information about the provided content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermaterializationflags?language=objc
type FileProviderMaterializationFlags uint

// Options for modifying items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermodifyitemoptions?language=objc
type FileProviderModifyItemOptions uint

const (
	FileProviderModifyItemMayAlreadyExist FileProviderModifyItemOptions = 1
)

// The location where the operation takes place. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperationside?language=objc
type FileProviderTestingOperationSide uint

const (
	FileProviderTestingOperationSideDisk         FileProviderTestingOperationSide = 0
	FileProviderTestingOperationSideFileProvider FileProviderTestingOperationSide = 1
)

// The action that an operation performs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperationtype?language=objc
type FileProviderTestingOperationType int

const (
	FileProviderTestingOperationTypeChildrenEnumeration FileProviderTestingOperationType = 6
	FileProviderTestingOperationTypeCollisionResolution FileProviderTestingOperationType = 7
	FileProviderTestingOperationTypeContentFetch        FileProviderTestingOperationType = 5
	FileProviderTestingOperationTypeCreation            FileProviderTestingOperationType = 2
	FileProviderTestingOperationTypeDeletion            FileProviderTestingOperationType = 4
	FileProviderTestingOperationTypeIngestion           FileProviderTestingOperationType = 0
	FileProviderTestingOperationTypeLookup              FileProviderTestingOperationType = 1
	FileProviderTestingOperationTypeModification        FileProviderTestingOperationType = 3
)
