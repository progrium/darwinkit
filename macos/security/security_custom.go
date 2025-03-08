package security

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// KeychainItem represents an item in the keychain
type KeychainItem struct {
	objc.Object
}

// Methods to interact with the keychain
func AddGenericPassword(serviceName string, accountName string, passwordData []byte, itemRef *KeychainItem) OSStatus {
	service := foundation.StringClass.StringWithString(serviceName)
	account := foundation.StringClass.StringWithString(accountName)
	
	// Create data from bytes
	dataBytes := unsafe.Pointer(&passwordData[0])
	dataLength := uint(len(passwordData))
	password := foundation.DataClass.DataWithBytesLength(dataBytes, dataLength)
	
	return OSStatus(int(objc.Call[int](objc.GetClass("Security"), objc.Sel("SecKeychainAddGenericPassword:account:passwordData:itemRef:"), 
		service, account, password, itemRef)))
}

func FindGenericPassword(serviceName string, accountName string, passwordLength *uint32, passwordData *[]byte, itemRef *KeychainItem) OSStatus {
	service := foundation.StringClass.StringWithString(serviceName)
	account := foundation.StringClass.StringWithString(accountName)
	
	var data foundation.Data
	status := OSStatus(int(objc.Call[int](objc.GetClass("Security"), objc.Sel("SecKeychainFindGenericPassword:account:passwordLength:passwordData:itemRef:"), 
		service, account, passwordLength, &data, itemRef)))
	
	if status == ErrSecSuccess && !data.IsNil() {
		length := data.Length()
		buffer := make([]byte, length)
		
		// Copy bytes from NSData to Go slice
		bytePtr := data.Bytes()
		if bytePtr != nil {
			for i := uint(0); i < length; i++ {
				ptr := unsafe.Pointer(uintptr(bytePtr) + uintptr(i))
				buffer[i] = *(*byte)(ptr)
			}
			*passwordData = buffer
		}
	}
	
	return status
}

func DeleteKeychainItem(item KeychainItem) OSStatus {
	return OSStatus(int(objc.Call[int](objc.GetClass("Security"), objc.Sel("SecKeychainItemDelete:"), item)))
}

// Certificate represents a certificate
type Certificate struct {
	objc.Object
}

// CertificateClass is the class instance for Certificate
var CertificateClass = objc.GetClass("SecCertificate")

// CreateCertificateFromData creates a certificate from data
func CreateCertificateFromData(data foundation.Data) (Certificate, OSStatus) {
	var cert Certificate
	status := OSStatus(int(objc.Call[int](CertificateClass, objc.Sel("SecCertificateCreateWithData:certificate:"), data, &cert)))
	return cert, status
}

// GetSummary returns a summary of the certificate
func (c Certificate) GetSummary() foundation.String {
	return objc.Call[foundation.String](c, objc.Sel("SecCertificateCopySubjectSummary:"), c)
}

// Identity represents an identity (certificate + private key)
type Identity struct {
	objc.Object
}

// IdentityClass is the class instance for Identity
var IdentityClass = objc.GetClass("SecIdentity")

// CreateIdentityWithCertificateAndPrivateKey creates an identity with a certificate and private key
func CreateIdentityWithCertificateAndPrivateKey(certificate Certificate, privateKey objc.Object) (Identity, OSStatus) {
	var identity Identity
	status := OSStatus(int(objc.Call[int](IdentityClass, objc.Sel("SecIdentityCreateWithCertificate:privateKey:identity:"), certificate, privateKey, &identity)))
	return identity, status
}

// GetCertificate returns the certificate for the identity
func (i Identity) GetCertificate() (Certificate, OSStatus) {
	var cert Certificate
	status := OSStatus(int(objc.Call[int](i, objc.Sel("SecIdentityCopyCertificate:certificate:"), i, &cert)))
	return cert, status
}

// Trust represents a trust evaluation
type Trust struct {
	objc.Object
}

// TrustClass is the class instance for Trust
var TrustClass = objc.GetClass("SecTrust")

// CreateTrustWithCertificates creates a trust evaluation with certificates
func CreateTrustWithCertificates(certificates foundation.Array, policies foundation.Array) (Trust, OSStatus) {
	var trust Trust
	status := OSStatus(int(objc.Call[int](TrustClass, objc.Sel("SecTrustCreateWithCertificates:policies:trust:"), certificates, policies, &trust)))
	return trust, status
}

// Evaluate evaluates the trust
func (t Trust) Evaluate() (TrustResultType, OSStatus) {
	var result TrustResultType
	status := OSStatus(int(objc.Call[int](t, objc.Sel("SecTrustEvaluate:result:"), t, &result)))
	return result, status
}

// Policy represents a security policy
type Policy struct {
	objc.Object
}

// PolicyClass is the class instance for Policy
var PolicyClass = objc.GetClass("SecPolicy")

// CreatePolicy creates a policy with properties
func CreatePolicy(policyIdentifier foundation.String, properties foundation.Dictionary) (Policy, OSStatus) {
	var policy Policy
	status := OSStatus(int(objc.Call[int](PolicyClass, objc.Sel("SecPolicyCreateWithProperties:properties:policy:"), policyIdentifier, properties, &policy)))
	return policy, status
}

// OSStatus represents a result code for Security framework operations
type OSStatus int32

// TrustResultType represents a trust evaluation result
type TrustResultType int32

// Common OSStatus result codes
const (
	ErrSecSuccess           OSStatus = 0
	ErrSecUnimplemented     OSStatus = -4
	ErrSecParam             OSStatus = -50
	ErrSecAllocate          OSStatus = -108
	ErrSecNotAvailable      OSStatus = -25291
	ErrSecDuplicateItem     OSStatus = -25299
	ErrSecItemNotFound      OSStatus = -25300
	ErrSecInteractionNotAllowed OSStatus = -25308
	ErrSecDecode            OSStatus = -26275
	ErrSecAuthFailed        OSStatus = -25293
)

// Trust evaluation result types
const (
	TrustResultInvalid      TrustResultType = 0
	TrustResultProceed      TrustResultType = 1
	TrustResultConfirm      TrustResultType = 2
	TrustResultDeny         TrustResultType = 3
	TrustResultUnspecified  TrustResultType = 4
	TrustResultRecoverableTrustFailure TrustResultType = 5
	TrustResultFatalTrustFailure TrustResultType = 6
	TrustResultOtherError   TrustResultType = 7
)

// Common constants
const (
	kSecClass                 = "kSecClass"
	kSecClassGenericPassword  = "kSecClassGenericPassword"
	kSecClassInternetPassword = "kSecClassInternetPassword"
	kSecClassCertificate      = "kSecClassCertificate"
	kSecClassKey              = "kSecClassKey"
	kSecClassIdentity         = "kSecClassIdentity"
	
	kSecAttrService          = "kSecAttrService"
	kSecAttrAccount          = "kSecAttrAccount"
	kSecAttrLabel            = "kSecAttrLabel"
	kSecValueData            = "kSecValueData"
	kSecReturnData           = "kSecReturnData"
	kSecReturnAttributes     = "kSecReturnAttributes"
	kSecReturnRef            = "kSecReturnRef"
	kSecMatchLimit           = "kSecMatchLimit"
	kSecMatchLimitOne        = "kSecMatchLimitOne"
	kSecMatchLimitAll        = "kSecMatchLimitAll"
)