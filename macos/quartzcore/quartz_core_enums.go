// auto-generated code, do not modify
package quartzcore

type AnimationCalculationMode string

const AnimationCubic AnimationCalculationMode = "cubic"
const AnimationCubicPaced AnimationCalculationMode = "cubicPaced"
const AnimationDiscrete AnimationCalculationMode = "discrete"
const AnimationLinear AnimationCalculationMode = "linear"
const AnimationPaced AnimationCalculationMode = "paced"

type AnimationRotationMode string

const AnimationRotateAuto AnimationRotationMode = "auto"
const AnimationRotateAutoReverse AnimationRotationMode = "autoReverse"

type AutoresizingMask uint32

const LayerNotSizable AutoresizingMask = 0
const LayerMinXMargin AutoresizingMask = 1
const LayerWidthSizable AutoresizingMask = 2
const LayerMaxXMargin AutoresizingMask = 4
const LayerMinYMargin AutoresizingMask = 8
const LayerHeightSizable AutoresizingMask = 16
const LayerMaxYMargin AutoresizingMask = 32

type ConstraintAttribute int32

const ConstraintHeight ConstraintAttribute = 7
const ConstraintMaxX ConstraintAttribute = 2
const ConstraintMaxY ConstraintAttribute = 6
const ConstraintMidX ConstraintAttribute = 1
const ConstraintMidY ConstraintAttribute = 5
const ConstraintMinX ConstraintAttribute = 0
const ConstraintMinY ConstraintAttribute = 4
const ConstraintWidth ConstraintAttribute = 3

type CornerMask uint

const LayerMaxXMaxYCorner CornerMask = 8
const LayerMaxXMinYCorner CornerMask = 2
const LayerMinXMaxYCorner CornerMask = 4
const LayerMinXMinYCorner CornerMask = 1

type EdgeAntialiasingMask uint32

const LayerLeftEdge EdgeAntialiasingMask = 1
const LayerRightEdge EdgeAntialiasingMask = 2
const LayerBottomEdge EdgeAntialiasingMask = 4
const LayerTopEdge EdgeAntialiasingMask = 8

type GradientLayerType string

const GradientLayerAxial GradientLayerType = "axial"
const GradientLayerConic GradientLayerType = "conic"
const GradientLayerRadial GradientLayerType = "radial"

type LayerContentsFilter string

const FilterLinear LayerContentsFilter = "linear"
const FilterNearest LayerContentsFilter = "nearest"
const FilterTrilinear LayerContentsFilter = "trilinear"

type LayerContentsFormat string

const ContentsFormatRGBA16Float LayerContentsFormat = "RGBAh"
const ContentsFormatRGBA8Uint LayerContentsFormat = "RGBA8"
const ContentsFormatGray8Uint LayerContentsFormat = "Gray8"

type LayerContentsGravity string

const GravityBottom LayerContentsGravity = "bottom"
const GravityBottomLeft LayerContentsGravity = "bottomLeft"
const GravityBottomRight LayerContentsGravity = "bottomRight"
const GravityCenter LayerContentsGravity = "center"
const GravityLeft LayerContentsGravity = "left"
const GravityResize LayerContentsGravity = "resize"
const GravityResizeAspect LayerContentsGravity = "resizeAspect"
const GravityResizeAspectFill LayerContentsGravity = "resizeAspectFill"
const GravityRight LayerContentsGravity = "right"
const GravityTop LayerContentsGravity = "top"
const GravityTopLeft LayerContentsGravity = "topLeft"
const GravityTopRight LayerContentsGravity = "topRight"

type LayerCornerCurve string

const CornerCurveCircular LayerCornerCurve = "circular"
const CornerCurveContinuous LayerCornerCurve = "continuous"

type ScrollLayerScrollMode string

const ScrollBoth ScrollLayerScrollMode = "both"
const ScrollHorizontally ScrollLayerScrollMode = "horizontally"
const ScrollNone ScrollLayerScrollMode = "none"
const ScrollVertically ScrollLayerScrollMode = "vertically"

type ShapeLayerFillRule string

const FillRuleEvenOdd ShapeLayerFillRule = "even-odd"
const FillRuleNonZero ShapeLayerFillRule = "non-zero"

type ShapeLayerLineCap string

const LineCapButt ShapeLayerLineCap = "butt"
const LineCapRound ShapeLayerLineCap = "round"
const LineCapSquare ShapeLayerLineCap = "square"

type ShapeLayerLineJoin string

const LineJoinBevel ShapeLayerLineJoin = "bevel"
const LineJoinMiter ShapeLayerLineJoin = "miter"
const LineJoinRound ShapeLayerLineJoin = "round"

type TextLayerAlignmentMode string

const AlignmentCenter TextLayerAlignmentMode = "center"
const AlignmentJustified TextLayerAlignmentMode = "justified"
const AlignmentLeft TextLayerAlignmentMode = "left"
const AlignmentNatural TextLayerAlignmentMode = "natural"
const AlignmentRight TextLayerAlignmentMode = "right"

type TextLayerTruncationMode string

const TruncationEnd TextLayerTruncationMode = "end"
const TruncationMiddle TextLayerTruncationMode = "middle"
const TruncationNone TextLayerTruncationMode = "none"
const TruncationStart TextLayerTruncationMode = "start"

type TransitionSubtype string

const TransitionFromBottom TransitionSubtype = "fromBottom"
const TransitionFromLeft TransitionSubtype = "fromLeft"
const TransitionFromRight TransitionSubtype = "fromRight"
const TransitionFromTop TransitionSubtype = "fromTop"

type TransitionType string

const TransitionFade TransitionType = "fade"
const TransitionMoveIn TransitionType = "moveIn"
const TransitionPush TransitionType = "push"
const TransitionReveal TransitionType = "reveal"

type ValueFunctionName string

const ValueFunctionRotateX ValueFunctionName = "rotateX"
const ValueFunctionRotateY ValueFunctionName = "rotateY"
const ValueFunctionRotateZ ValueFunctionName = "rotateZ"
const ValueFunctionScale ValueFunctionName = "scale"
const ValueFunctionScaleX ValueFunctionName = "scaleX"
const ValueFunctionScaleY ValueFunctionName = "scaleY"
const ValueFunctionScaleZ ValueFunctionName = "scaleZ"
const ValueFunctionTranslate ValueFunctionName = "translate"
const ValueFunctionTranslateX ValueFunctionName = "translateX"
const ValueFunctionTranslateY ValueFunctionName = "translateY"
const ValueFunctionTranslateZ ValueFunctionName = "translateZ"
