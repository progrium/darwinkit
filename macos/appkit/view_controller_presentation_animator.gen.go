// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that let you define animations to play when transitioning between two view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator?language=objc
type PViewControllerPresentationAnimator interface {
	// optional
	AnimatePresentationOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
	HasAnimatePresentationOfViewControllerFromViewController() bool

	// optional
	AnimateDismissalOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
	HasAnimateDismissalOfViewControllerFromViewController() bool
}

// A concrete type wrapper for the [PViewControllerPresentationAnimator] protocol.
type ViewControllerPresentationAnimatorWrapper struct {
	objc.Object
}

func (v_ ViewControllerPresentationAnimatorWrapper) HasAnimatePresentationOfViewControllerFromViewController() bool {
	return v_.RespondsToSelector(objc.Sel("animatePresentationOfViewController:fromViewController:"))
}

// Called when the specified view controller is about to be presented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator/1434396-animatepresentationofviewcontrol?language=objc
func (v_ ViewControllerPresentationAnimatorWrapper) AnimatePresentationOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("animatePresentationOfViewController:fromViewController:"), objc.Ptr(viewController), objc.Ptr(fromViewController))
}

func (v_ ViewControllerPresentationAnimatorWrapper) HasAnimateDismissalOfViewControllerFromViewController() bool {
	return v_.RespondsToSelector(objc.Sel("animateDismissalOfViewController:fromViewController:"))
}

// Called when a previously-presented view controller is about to be dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator/1434458-animatedismissalofviewcontroller?language=objc
func (v_ ViewControllerPresentationAnimatorWrapper) AnimateDismissalOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("animateDismissalOfViewController:fromViewController:"), objc.Ptr(viewController), objc.Ptr(fromViewController))
}
