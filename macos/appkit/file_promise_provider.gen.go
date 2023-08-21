// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilePromiseProvider] class.
var FilePromiseProviderClass = _FilePromiseProviderClass{objc.GetClass("NSFilePromiseProvider")}

type _FilePromiseProviderClass struct {
	objc.Class
}

// An interface definition for the [FilePromiseProvider] class.
type IFilePromiseProvider interface {
	objc.IObject
	UserInfo() objc.Object
	SetUserInfo(value objc.IObject)
	Delegate() FilePromiseProviderDelegateWrapper
	SetDelegate(value PFilePromiseProviderDelegate)
	SetDelegateObject(valueObject objc.IObject)
	FileType() string
	SetFileType(value string)
}

// An object that provides a promise for the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider?language=objc
type FilePromiseProvider struct {
	objc.Object
}

func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FilePromiseProvider) InitWithFileTypeDelegate(fileType string, delegate PFilePromiseProviderDelegate) FilePromiseProvider {
	po1 := objc.WrapAsProtocol("NSFilePromiseProviderDelegate", delegate)
	rv := objc.Call[FilePromiseProvider](f_, objc.Sel("initWithFileType:delegate:"), fileType, po1)
	return rv
}

// Initializes a file promise provider for a certain file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644594-initwithfiletype?language=objc
func NewFilePromiseProviderWithFileTypeDelegate(fileType string, delegate PFilePromiseProviderDelegate) FilePromiseProvider {
	instance := FilePromiseProviderClass.Alloc().InitWithFileTypeDelegate(fileType, delegate)
	instance.Autorelease()
	return instance
}

func (f_ FilePromiseProvider) Init() FilePromiseProvider {
	rv := objc.Call[FilePromiseProvider](f_, objc.Sel("init"))
	return rv
}

func (fc _FilePromiseProviderClass) Alloc() FilePromiseProvider {
	rv := objc.Call[FilePromiseProvider](fc, objc.Sel("alloc"))
	return rv
}

func FilePromiseProvider_Alloc() FilePromiseProvider {
	return FilePromiseProviderClass.Alloc()
}

func (fc _FilePromiseProviderClass) New() FilePromiseProvider {
	rv := objc.Call[FilePromiseProvider](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilePromiseProvider() FilePromiseProvider {
	return FilePromiseProviderClass.New()
}

// Optional user information to pass to the file promise provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644255-userinfo?language=objc
func (f_ FilePromiseProvider) UserInfo() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("userInfo"))
	return rv
}

// Optional user information to pass to the file promise provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644255-userinfo?language=objc
func (f_ FilePromiseProvider) SetUserInfo(value objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setUserInfo:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644726-delegate?language=objc
func (f_ FilePromiseProvider) Delegate() FilePromiseProviderDelegateWrapper {
	rv := objc.Call[FilePromiseProviderDelegateWrapper](f_, objc.Sel("delegate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644726-delegate?language=objc
func (f_ FilePromiseProvider) SetDelegate(value PFilePromiseProviderDelegate) {
	po0 := objc.WrapAsProtocol("NSFilePromiseProviderDelegate", value)
	objc.SetAssociatedObject(f_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644726-delegate?language=objc
func (f_ FilePromiseProvider) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The file type of the file promise provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644738-filetype?language=objc
func (f_ FilePromiseProvider) FileType() string {
	rv := objc.Call[string](f_, objc.Sel("fileType"))
	return rv
}

// The file type of the file promise provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseprovider/1644738-filetype?language=objc
func (f_ FilePromiseProvider) SetFileType(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setFileType:"), value)
}
