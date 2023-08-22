// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderDomain] class.
var FileProviderDomainClass = _FileProviderDomainClass{objc.GetClass("NSFileProviderDomain")}

type _FileProviderDomainClass struct {
	objc.Class
}

// An interface definition for the [FileProviderDomain] class.
type IFileProviderDomain interface {
	objc.IObject
	IsHidden() bool
	SetHidden(value bool)
	UserEnabled() bool
	IsDisconnected() bool
	TestingModes() FileProviderDomainTestingModes
	SetTestingModes(value FileProviderDomainTestingModes)
	DisplayName() string
	Identifier() FileProviderDomainIdentifier
	BackingStoreIdentity() []byte
}

// A File Provider extension's domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain?language=objc
type FileProviderDomain struct {
	objc.Object
}

func FileProviderDomainFrom(ptr unsafe.Pointer) FileProviderDomain {
	return FileProviderDomain{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FileProviderDomain) InitWithIdentifierDisplayName(identifier FileProviderDomainIdentifier, displayName string) FileProviderDomain {
	rv := objc.Call[FileProviderDomain](f_, objc.Sel("initWithIdentifier:displayName:"), identifier, displayName)
	return rv
}

// Creates a new file provider domain with the specified identifier and display name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3684871-initwithidentifier?language=objc
func NewFileProviderDomainWithIdentifierDisplayName(identifier FileProviderDomainIdentifier, displayName string) FileProviderDomain {
	instance := FileProviderDomainClass.Alloc().InitWithIdentifierDisplayName(identifier, displayName)
	instance.Autorelease()
	return instance
}

func (fc _FileProviderDomainClass) Alloc() FileProviderDomain {
	rv := objc.Call[FileProviderDomain](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderDomain_Alloc() FileProviderDomain {
	return FileProviderDomainClass.Alloc()
}

func (fc _FileProviderDomainClass) New() FileProviderDomain {
	rv := objc.Call[FileProviderDomain](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderDomain() FileProviderDomain {
	return FileProviderDomainClass.New()
}

func (f_ FileProviderDomain) Init() FileProviderDomain {
	rv := objc.Call[FileProviderDomain](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that determines whether the domain is visible to users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3553283-hidden?language=objc
func (f_ FileProviderDomain) IsHidden() bool {
	rv := objc.Call[bool](f_, objc.Sel("isHidden"))
	return rv
}

// A Boolean value that determines whether the domain is visible to users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3553283-hidden?language=objc
func (f_ FileProviderDomain) SetHidden(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setHidden:"), value)
}

// A Boolean value that indicates whether the user has enabled or disabled the domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3553284-userenabled?language=objc
func (f_ FileProviderDomain) UserEnabled() bool {
	rv := objc.Call[bool](f_, objc.Sel("userEnabled"))
	return rv
}

// A Boolean value indicating that the domain is present, but disconnected from the file extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3294483-disconnected?language=objc
func (f_ FileProviderDomain) IsDisconnected() bool {
	rv := objc.Call[bool](f_, objc.Sel("isDisconnected"))
	return rv
}

// A mode that gives the File Provider extension more control over the system’s behavior during testing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3727798-testingmodes?language=objc
func (f_ FileProviderDomain) TestingModes() FileProviderDomainTestingModes {
	rv := objc.Call[FileProviderDomainTestingModes](f_, objc.Sel("testingModes"))
	return rv
}

// A mode that gives the File Provider extension more control over the system’s behavior during testing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3727798-testingmodes?language=objc
func (f_ FileProviderDomain) SetTestingModes(value FileProviderDomainTestingModes) {
	objc.Call[objc.Void](f_, objc.Sel("setTestingModes:"), value)
}

// The name of the domain displayed in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/2882119-displayname?language=objc
func (f_ FileProviderDomain) DisplayName() string {
	rv := objc.Call[string](f_, objc.Sel("displayName"))
	return rv
}

// The domain's unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/2882117-identifier?language=objc
func (f_ FileProviderDomain) Identifier() FileProviderDomainIdentifier {
	rv := objc.Call[FileProviderDomainIdentifier](f_, objc.Sel("identifier"))
	return rv
}

// A unique identifier for the backing store used by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/3852591-backingstoreidentity?language=objc
func (f_ FileProviderDomain) BackingStoreIdentity() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("backingStoreIdentity"))
	return rv
}
