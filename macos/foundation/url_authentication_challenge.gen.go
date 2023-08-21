// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLAuthenticationChallenge] class.
var URLAuthenticationChallengeClass = _URLAuthenticationChallengeClass{objc.GetClass("NSURLAuthenticationChallenge")}

type _URLAuthenticationChallengeClass struct {
	objc.Class
}

// An interface definition for the [URLAuthenticationChallenge] class.
type IURLAuthenticationChallenge interface {
	objc.IObject
	Error() Error
	FailureResponse() URLResponse
	ProposedCredential() URLCredential
	Sender() URLAuthenticationChallengeSenderWrapper
	ProtectionSpace() URLProtectionSpace
	PreviousFailureCount() int
}

// A challenge from a server requiring authentication from the client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge?language=objc
type URLAuthenticationChallenge struct {
	objc.Object
}

func URLAuthenticationChallengeFrom(ptr unsafe.Pointer) URLAuthenticationChallenge {
	return URLAuthenticationChallenge{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLAuthenticationChallenge) InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space IURLProtectionSpace, credential IURLCredential, previousFailureCount int, response IURLResponse, error IError, sender PURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	po5 := objc.WrapAsProtocol("NSURLAuthenticationChallengeSender", sender)
	rv := objc.Call[URLAuthenticationChallenge](u_, objc.Sel("initWithProtectionSpace:proposedCredential:previousFailureCount:failureResponse:error:sender:"), objc.Ptr(space), objc.Ptr(credential), previousFailureCount, objc.Ptr(response), objc.Ptr(error), po5)
	return rv
}

// Initializes an authentication challenge from parameters you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1416511-initwithprotectionspace?language=objc
func NewURLAuthenticationChallengeWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space IURLProtectionSpace, credential IURLCredential, previousFailureCount int, response IURLResponse, error IError, sender PURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	instance := URLAuthenticationChallengeClass.Alloc().InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space, credential, previousFailureCount, response, error, sender)
	instance.Autorelease()
	return instance
}

func (u_ URLAuthenticationChallenge) InitWithAuthenticationChallengeSender(challenge IURLAuthenticationChallenge, sender PURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	po1 := objc.WrapAsProtocol("NSURLAuthenticationChallengeSender", sender)
	rv := objc.Call[URLAuthenticationChallenge](u_, objc.Sel("initWithAuthenticationChallenge:sender:"), objc.Ptr(challenge), po1)
	return rv
}

// Creates an authentication challenge from an existing challenge instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1411154-initwithauthenticationchallenge?language=objc
func NewURLAuthenticationChallengeWithAuthenticationChallengeSender(challenge IURLAuthenticationChallenge, sender PURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	instance := URLAuthenticationChallengeClass.Alloc().InitWithAuthenticationChallengeSender(challenge, sender)
	instance.Autorelease()
	return instance
}

func (uc _URLAuthenticationChallengeClass) Alloc() URLAuthenticationChallenge {
	rv := objc.Call[URLAuthenticationChallenge](uc, objc.Sel("alloc"))
	return rv
}

func URLAuthenticationChallenge_Alloc() URLAuthenticationChallenge {
	return URLAuthenticationChallengeClass.Alloc()
}

func (uc _URLAuthenticationChallengeClass) New() URLAuthenticationChallenge {
	rv := objc.Call[URLAuthenticationChallenge](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLAuthenticationChallenge() URLAuthenticationChallenge {
	return URLAuthenticationChallengeClass.New()
}

func (u_ URLAuthenticationChallenge) Init() URLAuthenticationChallenge {
	rv := objc.Call[URLAuthenticationChallenge](u_, objc.Sel("init"))
	return rv
}

// The error object representing the last authentication failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1413033-error?language=objc
func (u_ URLAuthenticationChallenge) Error() Error {
	rv := objc.Call[Error](u_, objc.Sel("error"))
	return rv
}

// The URL response object representing the last authentication failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1414978-failureresponse?language=objc
func (u_ URLAuthenticationChallenge) FailureResponse() URLResponse {
	rv := objc.Call[URLResponse](u_, objc.Sel("failureResponse"))
	return rv
}

// The proposed credential for this challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1417749-proposedcredential?language=objc
func (u_ URLAuthenticationChallenge) ProposedCredential() URLCredential {
	rv := objc.Call[URLCredential](u_, objc.Sel("proposedCredential"))
	return rv
}

// The sender of the challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1407533-sender?language=objc
func (u_ URLAuthenticationChallenge) Sender() URLAuthenticationChallengeSenderWrapper {
	rv := objc.Call[URLAuthenticationChallengeSenderWrapper](u_, objc.Sel("sender"))
	return rv
}

// The receiver’s protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1410012-protectionspace?language=objc
func (u_ URLAuthenticationChallenge) ProtectionSpace() URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](u_, objc.Sel("protectionSpace"))
	return rv
}

// The receiver’s count of failed authentication attempts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge/1416522-previousfailurecount?language=objc
func (u_ URLAuthenticationChallenge) PreviousFailureCount() int {
	rv := objc.Call[int](u_, objc.Sel("previousFailureCount"))
	return rv
}
