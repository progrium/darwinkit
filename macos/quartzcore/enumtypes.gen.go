// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

type AnimationCalculationMode string

const KAnimationCubic AnimationCalculationMode = "cubic"
const KAnimationCubicPaced AnimationCalculationMode = "cubicPaced"
const KAnimationDiscrete AnimationCalculationMode = "discrete"
const KAnimationLinear AnimationCalculationMode = "linear"
const KAnimationPaced AnimationCalculationMode = "paced"

type AnimationRotationMode string

const KAnimationRotateAuto AnimationRotationMode = "auto"
const KAnimationRotateAutoReverse AnimationRotationMode = "autoReverse"

type AutoresizingMask int

const KLayerHeightSizable AutoresizingMask = 16
const KLayerMaxXMargin AutoresizingMask = 4
const KLayerMaxYMargin AutoresizingMask = 32
const KLayerMinXMargin AutoresizingMask = 1
const KLayerMinYMargin AutoresizingMask = 8
const KLayerNotSizable AutoresizingMask = 0
const KLayerWidthSizable AutoresizingMask = 2

type ConstraintAttribute int

const KConstraintHeight ConstraintAttribute = 7
const KConstraintMaxX ConstraintAttribute = 2
const KConstraintMaxY ConstraintAttribute = 6
const KConstraintMidX ConstraintAttribute = 1
const KConstraintMidY ConstraintAttribute = 5
const KConstraintMinX ConstraintAttribute = 0
const KConstraintMinY ConstraintAttribute = 4
const KConstraintWidth ConstraintAttribute = 3

type CornerMask uint

const KLayerMaxXMaxYCorner CornerMask = 8
const KLayerMaxXMinYCorner CornerMask = 2
const KLayerMinXMaxYCorner CornerMask = 4
const KLayerMinXMinYCorner CornerMask = 1

type EdgeAntialiasingMask int

const KLayerBottomEdge EdgeAntialiasingMask = 4
const KLayerLeftEdge EdgeAntialiasingMask = 1
const KLayerRightEdge EdgeAntialiasingMask = 2
const KLayerTopEdge EdgeAntialiasingMask = 8

type EmitterLayerEmitterMode string

const KEmitterLayerOutline EmitterLayerEmitterMode = "outline"
const KEmitterLayerPoints EmitterLayerEmitterMode = "points"
const KEmitterLayerSurface EmitterLayerEmitterMode = "surface"
const KEmitterLayerVolume EmitterLayerEmitterMode = "volume"

type EmitterLayerEmitterShape string

const KEmitterLayerCircle EmitterLayerEmitterShape = "circle"
const KEmitterLayerCuboid EmitterLayerEmitterShape = "cuboid"
const KEmitterLayerLine EmitterLayerEmitterShape = "line"
const KEmitterLayerPoint EmitterLayerEmitterShape = "point"
const KEmitterLayerRectangle EmitterLayerEmitterShape = "rectangle"
const KEmitterLayerSphere EmitterLayerEmitterShape = "sphere"

type EmitterLayerRenderMode string

const KEmitterLayerAdditive EmitterLayerRenderMode = "additive"
const KEmitterLayerBackToFront EmitterLayerRenderMode = "backToFront"
const KEmitterLayerOldestFirst EmitterLayerRenderMode = "oldestFirst"
const KEmitterLayerOldestLast EmitterLayerRenderMode = "oldestLast"
const KEmitterLayerUnordered EmitterLayerRenderMode = "unordered"

type GradientLayerType string

const KGradientLayerAxial GradientLayerType = "axial"
const KGradientLayerConic GradientLayerType = "conic"
const KGradientLayerRadial GradientLayerType = "radial"

type LayerContentsFilter string

const KFilterLinear LayerContentsFilter = "linear"
const KFilterNearest LayerContentsFilter = "nearest"
const KFilterTrilinear LayerContentsFilter = "trilinear"

type LayerContentsFormat string

const KContentsFormatGray8Uint LayerContentsFormat = "Gray8"
const KContentsFormatRGBA16Float LayerContentsFormat = "RGBAh"
const KContentsFormatRGBA8Uint LayerContentsFormat = "RGBA8"

type LayerContentsGravity string

const KGravityBottom LayerContentsGravity = "bottom"
const KGravityBottomLeft LayerContentsGravity = "bottomLeft"
const KGravityBottomRight LayerContentsGravity = "bottomRight"
const KGravityCenter LayerContentsGravity = "center"
const KGravityLeft LayerContentsGravity = "left"
const KGravityResize LayerContentsGravity = "resize"
const KGravityResizeAspect LayerContentsGravity = "resizeAspect"
const KGravityResizeAspectFill LayerContentsGravity = "resizeAspectFill"
const KGravityRight LayerContentsGravity = "right"
const KGravityTop LayerContentsGravity = "top"
const KGravityTopLeft LayerContentsGravity = "topLeft"
const KGravityTopRight LayerContentsGravity = "topRight"

type LayerCornerCurve string

const KCornerCurveCircular LayerCornerCurve = "circular"
const KCornerCurveContinuous LayerCornerCurve = "continuous"

type MediaTimingFillMode string

const KFillModeBackwards MediaTimingFillMode = "backwards"
const KFillModeBoth MediaTimingFillMode = "both"
const KFillModeForwards MediaTimingFillMode = "forwards"
const KFillModeRemoved MediaTimingFillMode = "removed"

type MediaTimingFunctionName string

const KMediaTimingFunctionDefault MediaTimingFunctionName = "default"
const KMediaTimingFunctionEaseIn MediaTimingFunctionName = "easeIn"
const KMediaTimingFunctionEaseInEaseOut MediaTimingFunctionName = "easeInEaseOut"
const KMediaTimingFunctionEaseOut MediaTimingFunctionName = "easeOut"
const KMediaTimingFunctionLinear MediaTimingFunctionName = "linear"

type ScrollLayerScrollMode string

const KScrollBoth ScrollLayerScrollMode = "both"
const KScrollHorizontally ScrollLayerScrollMode = "horizontally"
const KScrollNone ScrollLayerScrollMode = "none"
const KScrollVertically ScrollLayerScrollMode = "vertically"

type ShapeLayerFillRule string

const KFillRuleEvenOdd ShapeLayerFillRule = "even-odd"
const KFillRuleNonZero ShapeLayerFillRule = "non-zero"

type ShapeLayerLineCap string

const KLineCapButt ShapeLayerLineCap = "butt"
const KLineCapRound ShapeLayerLineCap = "round"
const KLineCapSquare ShapeLayerLineCap = "square"

type ShapeLayerLineJoin string

const KLineJoinBevel ShapeLayerLineJoin = "bevel"
const KLineJoinMiter ShapeLayerLineJoin = "miter"
const KLineJoinRound ShapeLayerLineJoin = "round"

type TextLayerAlignmentMode string

const KAlignmentCenter TextLayerAlignmentMode = "center"
const KAlignmentJustified TextLayerAlignmentMode = "justified"
const KAlignmentLeft TextLayerAlignmentMode = "left"
const KAlignmentNatural TextLayerAlignmentMode = "natural"
const KAlignmentRight TextLayerAlignmentMode = "right"

type TextLayerTruncationMode string

const KTruncationEnd TextLayerTruncationMode = "end"
const KTruncationMiddle TextLayerTruncationMode = "middle"
const KTruncationNone TextLayerTruncationMode = "none"
const KTruncationStart TextLayerTruncationMode = "start"

type TransitionSubtype string

const KTransitionFromBottom TransitionSubtype = "fromBottom"
const KTransitionFromLeft TransitionSubtype = "fromLeft"
const KTransitionFromRight TransitionSubtype = "fromRight"
const KTransitionFromTop TransitionSubtype = "fromTop"

type TransitionType string

const KTransitionFade TransitionType = "fade"
const KTransitionMoveIn TransitionType = "moveIn"
const KTransitionPush TransitionType = "push"
const KTransitionReveal TransitionType = "reveal"

type ValueFunctionName string

const KValueFunctionRotateX ValueFunctionName = "rotateX"
const KValueFunctionRotateY ValueFunctionName = "rotateY"
const KValueFunctionRotateZ ValueFunctionName = "rotateZ"
const KValueFunctionScale ValueFunctionName = "scale"
const KValueFunctionScaleX ValueFunctionName = "scaleX"
const KValueFunctionScaleY ValueFunctionName = "scaleY"
const KValueFunctionScaleZ ValueFunctionName = "scaleZ"
const KValueFunctionTranslate ValueFunctionName = "translate"
const KValueFunctionTranslateX ValueFunctionName = "translateX"
const KValueFunctionTranslateY ValueFunctionName = "translateY"
const KValueFunctionTranslateZ ValueFunctionName = "translateZ"
