package layout

import (
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

func AliginAnchors(anchor appkit.ILayoutAnchor, targetAncor appkit.ILayoutAnchor) {
	anchor.ConstraintEqualToAnchor(targetAncor).SetActive(true)
}

func AliginBottom(view appkit.IView, targetView appkit.IView) {
	view.BottomAnchor().ConstraintEqualToAnchor(targetView.BottomAnchor()).SetActive(true)
}

func AliginLeft(view appkit.IView, targetView appkit.IView) {
	view.LeftAnchor().ConstraintEqualToAnchor(targetView.LeftAnchor()).SetActive(true)
}

func AliginTop(view appkit.IView, targetView appkit.IView) {
	view.TopAnchor().ConstraintEqualToAnchor(targetView.TopAnchor()).SetActive(true)
}

func AliginRight(view appkit.IView, targetView appkit.IView) {
	view.RightAnchor().ConstraintEqualToAnchor(targetView.RightAnchor()).SetActive(true)
}

func AliginLeading(view appkit.IView, targetView appkit.IView) {
	view.LeadingAnchor().ConstraintEqualToAnchor(targetView.LeadingAnchor()).SetActive(true)
}

func AliginTrailing(view appkit.IView, targetView appkit.IView) {
	view.TrailingAnchor().ConstraintEqualToAnchor(targetView.TrailingAnchor()).SetActive(true)
}

func AliginFirstBaseline(view appkit.IView, targetView appkit.IView) {
	view.FirstBaselineAnchor().ConstraintEqualToAnchor(targetView.FirstBaselineAnchor()).SetActive(true)
}

func AliginLastBaseline(view appkit.IView, targetView appkit.IView) {
	view.LastBaselineAnchor().ConstraintEqualToAnchor(targetView.LastBaselineAnchor()).SetActive(true)
}

func AliginCenterX(view appkit.IView, targetView appkit.IView) {
	view.CenterXAnchor().ConstraintEqualToAnchor(targetView.CenterXAnchor()).SetActive(true)
}

func AliginCenterY(view appkit.IView, targetView appkit.IView) {
	view.CenterYAnchor().ConstraintEqualToAnchor(targetView.CenterYAnchor()).SetActive(true)
}

func PinAnchorTo(anchor appkit.ILayoutAnchor, targetAncor appkit.ILayoutAnchor, offset float64) {
	anchor.ConstraintEqualToAnchorConstant(targetAncor, offset).SetActive(true)
}

// PinEdgesToSuperView set view's insets to it's super view.
func PinEdgesToSuperView(view appkit.IView, edgeInsets foundation.EdgeInsets) {
	superView := view.Superview()
	if superView.IsNil() {
		return
	}
	view.TopAnchor().ConstraintEqualToAnchorConstant(superView.TopAnchor(), edgeInsets.Top).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchorConstant(superView.RightAnchor(), -edgeInsets.Right).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchorConstant(superView.BottomAnchor(), -edgeInsets.Bottom).SetActive(true)
	view.LeftAnchor().ConstraintEqualToAnchorConstant(superView.LeftAnchor(), edgeInsets.Left).SetActive(true)
}

func SetWidth(view appkit.IView, width float64) {
	view.WidthAnchor().ConstraintEqualToConstant(width).SetActive(true)
}

func SetHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintEqualToConstant(height).SetActive(true)
}

func SetMaxWidth(view appkit.IView, width float64) {
	view.WidthAnchor().ConstraintLessThanOrEqualToConstant(width).SetActive(true)
}

func SetMaxHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintLessThanOrEqualToConstant(height).SetActive(true)
}

func SetMinWidth(view appkit.IView, width float64) {
	view.WidthAnchor().ConstraintGreaterThanOrEqualToConstant(width).SetActive(true)
}

func SetMinHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintGreaterThanOrEqualToConstant(height).SetActive(true)
}

func AlignWidth(view appkit.IView, targetView appkit.IView) {
	view.WidthAnchor().ConstraintEqualToAnchor(targetView.WidthAnchor()).SetActive(true)
}

func AlginHeight(view appkit.IView, targetView appkit.IView) {
	view.HeightAnchor().ConstraintEqualToAnchor(targetView.HeightAnchor()).SetActive(true)
}

// AutoAlignControlsWidth set width anchor for multi controls, using the max width for all controls' width
func AutoAlignControlsWidth(controls ...appkit.IControl) {
	var maxWidth float64
	for _, control := range controls {
		control.SizeToFit()
		w := control.Bounds().Size.Width
		if maxWidth < w {
			maxWidth = w
		}
	}
	for _, control := range controls {
		control.WidthAnchor().ConstraintEqualToConstant(maxWidth).SetActive(true)
	}
}
