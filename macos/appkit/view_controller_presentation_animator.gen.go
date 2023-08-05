// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IViewControllerPresentationAnimator interface {
	// required
	AnimatePresentationOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
	// required
	AnimateDismissalOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
}

type ViewControllerPresentationAnimatorWrapper struct {
	objc.Object
}

func (v_ ViewControllerPresentationAnimatorWrapper) AnimatePresentationOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("animatePresentationOfViewController:fromViewController:"), objc.ExtractPtr(viewController), objc.ExtractPtr(fromViewController))
}

func (v_ ViewControllerPresentationAnimatorWrapper) AnimateDismissalOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("animateDismissalOfViewController:fromViewController:"), objc.ExtractPtr(viewController), objc.ExtractPtr(fromViewController))
}
