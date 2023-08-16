// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The URLAuthenticationChallengeSender protocol represents the interface that the sender of an authentication challenge must implement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender?language=objc
type PURLAuthenticationChallengeSender interface {
	// optional
	CancelAuthenticationChallenge(challenge URLAuthenticationChallenge)
	HasCancelAuthenticationChallenge() bool

	// optional
	ContinueWithoutCredentialForAuthenticationChallenge(challenge URLAuthenticationChallenge)
	HasContinueWithoutCredentialForAuthenticationChallenge() bool

	// optional
	PerformDefaultHandlingForAuthenticationChallenge(challenge URLAuthenticationChallenge)
	HasPerformDefaultHandlingForAuthenticationChallenge() bool

	// optional
	RejectProtectionSpaceAndContinueWithChallenge(challenge URLAuthenticationChallenge)
	HasRejectProtectionSpaceAndContinueWithChallenge() bool

	// optional
	UseCredentialForAuthenticationChallenge(credential URLCredential, challenge URLAuthenticationChallenge)
	HasUseCredentialForAuthenticationChallenge() bool
}

// A concrete type wrapper for the [PURLAuthenticationChallengeSender] protocol.
type URLAuthenticationChallengeSenderWrapper struct {
	objc.Object
}

func (u_ URLAuthenticationChallengeSenderWrapper) HasCancelAuthenticationChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("cancelAuthenticationChallenge:"))
}

// Cancels a given authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender/1414474-cancelauthenticationchallenge?language=objc
func (u_ URLAuthenticationChallengeSenderWrapper) CancelAuthenticationChallenge(challenge IURLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("cancelAuthenticationChallenge:"), objc.Ptr(challenge))
}

func (u_ URLAuthenticationChallengeSenderWrapper) HasContinueWithoutCredentialForAuthenticationChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("continueWithoutCredentialForAuthenticationChallenge:"))
}

// Attempt to continue downloading a request without providing a credential for a given challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender/1413016-continuewithoutcredentialforauth?language=objc
func (u_ URLAuthenticationChallengeSenderWrapper) ContinueWithoutCredentialForAuthenticationChallenge(challenge IURLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("continueWithoutCredentialForAuthenticationChallenge:"), objc.Ptr(challenge))
}

func (u_ URLAuthenticationChallengeSenderWrapper) HasPerformDefaultHandlingForAuthenticationChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("performDefaultHandlingForAuthenticationChallenge:"))
}

// Causes the system-provided default behavior to be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender/1414590-performdefaulthandlingforauthent?language=objc
func (u_ URLAuthenticationChallengeSenderWrapper) PerformDefaultHandlingForAuthenticationChallenge(challenge IURLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("performDefaultHandlingForAuthenticationChallenge:"), objc.Ptr(challenge))
}

func (u_ URLAuthenticationChallengeSenderWrapper) HasRejectProtectionSpaceAndContinueWithChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("rejectProtectionSpaceAndContinueWithChallenge:"))
}

// Rejects the currently supplied protection space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender/1417331-rejectprotectionspaceandcontinue?language=objc
func (u_ URLAuthenticationChallengeSenderWrapper) RejectProtectionSpaceAndContinueWithChallenge(challenge IURLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("rejectProtectionSpaceAndContinueWithChallenge:"), objc.Ptr(challenge))
}

func (u_ URLAuthenticationChallengeSenderWrapper) HasUseCredentialForAuthenticationChallenge() bool {
	return u_.RespondsToSelector(objc.Sel("useCredential:forAuthenticationChallenge:"))
}

// Attempt to use a given credential for a given authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlauthenticationchallengesender/1411062-usecredential?language=objc
func (u_ URLAuthenticationChallengeSenderWrapper) UseCredentialForAuthenticationChallenge(credential IURLCredential, challenge IURLAuthenticationChallenge) {
	objc.Call[objc.Void](u_, objc.Sel("useCredential:forAuthenticationChallenge:"), objc.Ptr(credential), objc.Ptr(challenge))
}
