// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationcalculationmode?language=objc
type AnimationCalculationMode string

const (
	KAnimationCubic      AnimationCalculationMode = "cubic"
	KAnimationCubicPaced AnimationCalculationMode = "cubicPaced"
	KAnimationDiscrete   AnimationCalculationMode = "discrete"
	KAnimationLinear     AnimationCalculationMode = "linear"
	KAnimationPaced      AnimationCalculationMode = "paced"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationrotationmode?language=objc
type AnimationRotationMode string

const (
	KAnimationRotateAuto        AnimationRotationMode = "auto"
	KAnimationRotateAutoReverse AnimationRotationMode = "autoReverse"
)

// These constants are used by the autoresizingMask property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caautoresizingmask?language=objc
type AutoresizingMask int

const (
	KLayerHeightSizable AutoresizingMask = 16
	KLayerMaxXMargin    AutoresizingMask = 4
	KLayerMaxYMargin    AutoresizingMask = 32
	KLayerMinXMargin    AutoresizingMask = 1
	KLayerMinYMargin    AutoresizingMask = 8
	KLayerNotSizable    AutoresizingMask = 0
	KLayerWidthSizable  AutoresizingMask = 2
)

// The constraint attribute type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraintattribute?language=objc
type ConstraintAttribute int

const (
	KConstraintHeight ConstraintAttribute = 7
	KConstraintMaxX   ConstraintAttribute = 2
	KConstraintMaxY   ConstraintAttribute = 6
	KConstraintMidX   ConstraintAttribute = 1
	KConstraintMidY   ConstraintAttribute = 5
	KConstraintMinX   ConstraintAttribute = 0
	KConstraintMinY   ConstraintAttribute = 4
	KConstraintWidth  ConstraintAttribute = 3
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cacornermask?language=objc
type CornerMask uint

const (
	KLayerMaxXMaxYCorner CornerMask = 8
	KLayerMaxXMinYCorner CornerMask = 2
	KLayerMinXMaxYCorner CornerMask = 4
	KLayerMinXMinYCorner CornerMask = 1
)

// This mask is used by the edgeAntialiasingMask property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caedgeantialiasingmask?language=objc
type EdgeAntialiasingMask int

const (
	KLayerBottomEdge EdgeAntialiasingMask = 4
	KLayerLeftEdge   EdgeAntialiasingMask = 1
	KLayerRightEdge  EdgeAntialiasingMask = 2
	KLayerTopEdge    EdgeAntialiasingMask = 8
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayeremittermode?language=objc
type EmitterLayerEmitterMode string

const (
	KEmitterLayerOutline EmitterLayerEmitterMode = "outline"
	KEmitterLayerPoints  EmitterLayerEmitterMode = "points"
	KEmitterLayerSurface EmitterLayerEmitterMode = "surface"
	KEmitterLayerVolume  EmitterLayerEmitterMode = "volume"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayeremittershape?language=objc
type EmitterLayerEmitterShape string

const (
	KEmitterLayerCircle    EmitterLayerEmitterShape = "circle"
	KEmitterLayerCuboid    EmitterLayerEmitterShape = "cuboid"
	KEmitterLayerLine      EmitterLayerEmitterShape = "line"
	KEmitterLayerPoint     EmitterLayerEmitterShape = "point"
	KEmitterLayerRectangle EmitterLayerEmitterShape = "rectangle"
	KEmitterLayerSphere    EmitterLayerEmitterShape = "sphere"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayerrendermode?language=objc
type EmitterLayerRenderMode string

const (
	KEmitterLayerAdditive    EmitterLayerRenderMode = "additive"
	KEmitterLayerBackToFront EmitterLayerRenderMode = "backToFront"
	KEmitterLayerOldestFirst EmitterLayerRenderMode = "oldestFirst"
	KEmitterLayerOldestLast  EmitterLayerRenderMode = "oldestLast"
	KEmitterLayerUnordered   EmitterLayerRenderMode = "unordered"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayertype?language=objc
type GradientLayerType string

const (
	KGradientLayerAxial  GradientLayerType = "axial"
	KGradientLayerConic  GradientLayerType = "conic"
	KGradientLayerRadial GradientLayerType = "radial"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayercontentsfilter?language=objc
type LayerContentsFilter string

const (
	KFilterLinear    LayerContentsFilter = "linear"
	KFilterNearest   LayerContentsFilter = "nearest"
	KFilterTrilinear LayerContentsFilter = "trilinear"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayercontentsformat?language=objc
type LayerContentsFormat string

const (
	KContentsFormatGray8Uint   LayerContentsFormat = "Gray8"
	KContentsFormatRGBA16Float LayerContentsFormat = "RGBAh"
	KContentsFormatRGBA8Uint   LayerContentsFormat = "RGBA8"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayercontentsgravity?language=objc
type LayerContentsGravity string

const (
	KGravityBottom           LayerContentsGravity = "bottom"
	KGravityBottomLeft       LayerContentsGravity = "bottomLeft"
	KGravityBottomRight      LayerContentsGravity = "bottomRight"
	KGravityCenter           LayerContentsGravity = "center"
	KGravityLeft             LayerContentsGravity = "left"
	KGravityResize           LayerContentsGravity = "resize"
	KGravityResizeAspect     LayerContentsGravity = "resizeAspect"
	KGravityResizeAspectFill LayerContentsGravity = "resizeAspectFill"
	KGravityRight            LayerContentsGravity = "right"
	KGravityTop              LayerContentsGravity = "top"
	KGravityTopLeft          LayerContentsGravity = "topLeft"
	KGravityTopRight         LayerContentsGravity = "topRight"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayercornercurve?language=objc
type LayerCornerCurve string

const (
	KCornerCurveCircular   LayerCornerCurve = "circular"
	KCornerCurveContinuous LayerCornerCurve = "continuous"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfillmode?language=objc
type MediaTimingFillMode string

const (
	KFillModeBackwards MediaTimingFillMode = "backwards"
	KFillModeBoth      MediaTimingFillMode = "both"
	KFillModeForwards  MediaTimingFillMode = "forwards"
	KFillModeRemoved   MediaTimingFillMode = "removed"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatimingfunctionname?language=objc
type MediaTimingFunctionName string

const (
	KMediaTimingFunctionDefault       MediaTimingFunctionName = "default"
	KMediaTimingFunctionEaseIn        MediaTimingFunctionName = "easeIn"
	KMediaTimingFunctionEaseInEaseOut MediaTimingFunctionName = "easeInEaseOut"
	KMediaTimingFunctionEaseOut       MediaTimingFunctionName = "easeOut"
	KMediaTimingFunctionLinear        MediaTimingFunctionName = "linear"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayerscrollmode?language=objc
type ScrollLayerScrollMode string

const (
	KScrollBoth         ScrollLayerScrollMode = "both"
	KScrollHorizontally ScrollLayerScrollMode = "horizontally"
	KScrollNone         ScrollLayerScrollMode = "none"
	KScrollVertically   ScrollLayerScrollMode = "vertically"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayerfillrule?language=objc
type ShapeLayerFillRule string

const (
	KFillRuleEvenOdd ShapeLayerFillRule = "even-odd"
	KFillRuleNonZero ShapeLayerFillRule = "non-zero"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayerlinecap?language=objc
type ShapeLayerLineCap string

const (
	KLineCapButt   ShapeLayerLineCap = "butt"
	KLineCapRound  ShapeLayerLineCap = "round"
	KLineCapSquare ShapeLayerLineCap = "square"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayerlinejoin?language=objc
type ShapeLayerLineJoin string

const (
	KLineJoinBevel ShapeLayerLineJoin = "bevel"
	KLineJoinMiter ShapeLayerLineJoin = "miter"
	KLineJoinRound ShapeLayerLineJoin = "round"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayeralignmentmode?language=objc
type TextLayerAlignmentMode string

const (
	KAlignmentCenter    TextLayerAlignmentMode = "center"
	KAlignmentJustified TextLayerAlignmentMode = "justified"
	KAlignmentLeft      TextLayerAlignmentMode = "left"
	KAlignmentNatural   TextLayerAlignmentMode = "natural"
	KAlignmentRight     TextLayerAlignmentMode = "right"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayertruncationmode?language=objc
type TextLayerTruncationMode string

const (
	KTruncationEnd    TextLayerTruncationMode = "end"
	KTruncationMiddle TextLayerTruncationMode = "middle"
	KTruncationNone   TextLayerTruncationMode = "none"
	KTruncationStart  TextLayerTruncationMode = "start"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransitionsubtype?language=objc
type TransitionSubtype string

const (
	KTransitionFromBottom TransitionSubtype = "fromBottom"
	KTransitionFromLeft   TransitionSubtype = "fromLeft"
	KTransitionFromRight  TransitionSubtype = "fromRight"
	KTransitionFromTop    TransitionSubtype = "fromTop"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransitiontype?language=objc
type TransitionType string

const (
	KTransitionFade   TransitionType = "fade"
	KTransitionMoveIn TransitionType = "moveIn"
	KTransitionPush   TransitionType = "push"
	KTransitionReveal TransitionType = "reveal"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cavaluefunctionname?language=objc
type ValueFunctionName string

const (
	KValueFunctionRotateX    ValueFunctionName = "rotateX"
	KValueFunctionRotateY    ValueFunctionName = "rotateY"
	KValueFunctionRotateZ    ValueFunctionName = "rotateZ"
	KValueFunctionScale      ValueFunctionName = "scale"
	KValueFunctionScaleX     ValueFunctionName = "scaleX"
	KValueFunctionScaleY     ValueFunctionName = "scaleY"
	KValueFunctionScaleZ     ValueFunctionName = "scaleZ"
	KValueFunctionTranslate  ValueFunctionName = "translate"
	KValueFunctionTranslateX ValueFunctionName = "translateX"
	KValueFunctionTranslateY ValueFunctionName = "translateY"
	KValueFunctionTranslateZ ValueFunctionName = "translateZ"
)
