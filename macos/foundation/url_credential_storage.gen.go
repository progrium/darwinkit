// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLCredentialStorage] class.
var URLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}

type _URLCredentialStorageClass struct {
	objc.Class
}

// An interface definition for the [URLCredentialStorage] class.
type IURLCredentialStorage interface {
	objc.IObject
	DefaultCredentialForProtectionSpace(space IURLProtectionSpace) URLCredential
	SetDefaultCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space IURLProtectionSpace, task IURLSessionTask, completionHandler func(credential URLCredential))
	RemoveCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	CredentialsForProtectionSpace(space IURLProtectionSpace) map[string]URLCredential
	SetCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	GetCredentialsForProtectionSpaceTaskCompletionHandler(protectionSpace IURLProtectionSpace, task IURLSessionTask, completionHandler func(credentials map[string]URLCredential))
	AllCredentials() Dictionary
}

// The manager of a shared credentials cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage?language=objc
type URLCredentialStorage struct {
	objc.Object
}

func URLCredentialStorageFrom(ptr unsafe.Pointer) URLCredentialStorage {
	return URLCredentialStorage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLCredentialStorageClass) Alloc() URLCredentialStorage {
	rv := objc.Call[URLCredentialStorage](uc, objc.Sel("alloc"))
	return rv
}

func URLCredentialStorage_Alloc() URLCredentialStorage {
	return URLCredentialStorageClass.Alloc()
}

func (uc _URLCredentialStorageClass) New() URLCredentialStorage {
	rv := objc.Call[URLCredentialStorage](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLCredentialStorage() URLCredentialStorage {
	return URLCredentialStorageClass.New()
}

func (u_ URLCredentialStorage) Init() URLCredentialStorage {
	rv := objc.Call[URLCredentialStorage](u_, objc.Sel("init"))
	return rv
}

// Returns the default credential for the specified protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1412650-defaultcredentialforprotectionsp?language=objc
func (u_ URLCredentialStorage) DefaultCredentialForProtectionSpace(space IURLProtectionSpace) URLCredential {
	rv := objc.Call[URLCredential](u_, objc.Sel("defaultCredentialForProtectionSpace:"), objc.Ptr(space))
	return rv
}

// Sets the default credential for a specified protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1416502-setdefaultcredential?language=objc
func (u_ URLCredentialStorage) SetDefaultCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Call[objc.Void](u_, objc.Sel("setDefaultCredential:forProtectionSpace:"), objc.Ptr(credential), objc.Ptr(space))
}

// Gets the default credential for the specified protection space, which is being accessed by the given task, and passes it to the provided completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1411794-getdefaultcredentialforprotectio?language=objc
func (u_ URLCredentialStorage) GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space IURLProtectionSpace, task IURLSessionTask, completionHandler func(credential URLCredential)) {
	objc.Call[objc.Void](u_, objc.Sel("getDefaultCredentialForProtectionSpace:task:completionHandler:"), objc.Ptr(space), objc.Ptr(task), completionHandler)
}

// Removes the specified credential from the credential storage for the specified protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1408664-removecredential?language=objc
func (u_ URLCredentialStorage) RemoveCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Call[objc.Void](u_, objc.Sel("removeCredential:forProtectionSpace:"), objc.Ptr(credential), objc.Ptr(space))
}

// Returns a dictionary containing the credentials for the specified protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1413910-credentialsforprotectionspace?language=objc
func (u_ URLCredentialStorage) CredentialsForProtectionSpace(space IURLProtectionSpace) map[string]URLCredential {
	rv := objc.Call[map[string]URLCredential](u_, objc.Sel("credentialsForProtectionSpace:"), objc.Ptr(space))
	return rv
}

// Adds a credential to the credential storage for the specified protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1407227-setcredential?language=objc
func (u_ URLCredentialStorage) SetCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Call[objc.Void](u_, objc.Sel("setCredential:forProtectionSpace:"), objc.Ptr(credential), objc.Ptr(space))
}

// Gets a dictionary containing the credentials for the specified protection space, on behalf of the given task, and passes the dictionary to the provided completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1418119-getcredentialsforprotectionspace?language=objc
func (u_ URLCredentialStorage) GetCredentialsForProtectionSpaceTaskCompletionHandler(protectionSpace IURLProtectionSpace, task IURLSessionTask, completionHandler func(credentials map[string]URLCredential)) {
	objc.Call[objc.Void](u_, objc.Sel("getCredentialsForProtectionSpace:task:completionHandler:"), objc.Ptr(protectionSpace), objc.Ptr(task), completionHandler)
}

// The shared URL credential storage instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1412355-sharedcredentialstorage?language=objc
func (uc _URLCredentialStorageClass) SharedCredentialStorage() URLCredentialStorage {
	rv := objc.Call[URLCredentialStorage](uc, objc.Sel("sharedCredentialStorage"))
	return rv
}

// The shared URL credential storage instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1412355-sharedcredentialstorage?language=objc
func URLCredentialStorage_SharedCredentialStorage() URLCredentialStorage {
	return URLCredentialStorageClass.SharedCredentialStorage()
}

// The credentials for all available protection spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcredentialstorage/1413859-allcredentials?language=objc
func (u_ URLCredentialStorage) AllCredentials() Dictionary {
	rv := objc.Call[Dictionary](u_, objc.Sel("allCredentials"))
	return rv
}
