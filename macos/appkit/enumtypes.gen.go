// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import "unsafe"

type AboutPanelOptionKey string

const AboutPanelOptionApplicationIcon AboutPanelOptionKey = "ApplicationIcon"
const AboutPanelOptionApplicationName AboutPanelOptionKey = "ApplicationName"
const AboutPanelOptionApplicationVersion AboutPanelOptionKey = "ApplicationVersion"
const AboutPanelOptionCredits AboutPanelOptionKey = "Credits"
const AboutPanelOptionVersion AboutPanelOptionKey = "Version"

type AlertStyle uint

const AlertStyleCritical AlertStyle = 2
const AlertStyleInformational AlertStyle = 1
const AlertStyleWarning AlertStyle = 0

type AnimationBlockingMode uint

const AnimationBlocking AnimationBlockingMode = 0
const AnimationNonblocking AnimationBlockingMode = 1
const AnimationNonblockingThreaded AnimationBlockingMode = 2

type AnimationCurve uint

const AnimationEaseInOut AnimationCurve = 0
const AnimationEaseIn AnimationCurve = 1
const AnimationEaseOut AnimationCurve = 2
const AnimationLinear AnimationCurve = 3

type AnimationProgress float32

type AppearanceName string

const AppearanceNameAqua AppearanceName = "NSAppearanceNameAqua"
const AppearanceNameDarkAqua AppearanceName = "NSAppearanceNameDarkAqua"
const AppearanceNameVibrantLight AppearanceName = "NSAppearanceNameVibrantLight"
const AppearanceNameVibrantDark AppearanceName = "NSAppearanceNameVibrantDark"
const AppearanceNameAccessibilityHighContrastAqua AppearanceName = "NSAppearanceNameAccessibilityAqua"
const AppearanceNameAccessibilityHighContrastDarkAqua AppearanceName = "NSAppearanceNameAccessibilityDarkAqua"
const AppearanceNameAccessibilityHighContrastVibrantLight AppearanceName = "NSAppearanceNameAccessibilityVibrantLight"
const AppearanceNameAccessibilityHighContrastVibrantDark AppearanceName = "NSAppearanceNameAccessibilityVibrantDark"

type ApplicationActivationOptions uint

const ApplicationActivateAllWindows ApplicationActivationOptions = 1
const ApplicationActivateIgnoringOtherApps ApplicationActivationOptions = 2

type ApplicationActivationPolicy int

const ApplicationActivationPolicyRegular ApplicationActivationPolicy = 0
const ApplicationActivationPolicyAccessory ApplicationActivationPolicy = 1
const ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 2

type ApplicationDelegateReply uint

const ApplicationDelegateReplySuccess ApplicationDelegateReply = 0
const ApplicationDelegateReplyCancel ApplicationDelegateReply = 1
const ApplicationDelegateReplyFailure ApplicationDelegateReply = 2

type ApplicationOcclusionState uint

const ApplicationOcclusionStateVisible ApplicationOcclusionState = 2

type ApplicationPresentationOptions uint

const ApplicationPresentationDefault ApplicationPresentationOptions = 0
const ApplicationPresentationAutoHideDock ApplicationPresentationOptions = 1
const ApplicationPresentationHideDock ApplicationPresentationOptions = 2
const ApplicationPresentationAutoHideMenuBar ApplicationPresentationOptions = 4
const ApplicationPresentationHideMenuBar ApplicationPresentationOptions = 8
const ApplicationPresentationDisableAppleMenu ApplicationPresentationOptions = 16
const ApplicationPresentationDisableProcessSwitching ApplicationPresentationOptions = 32
const ApplicationPresentationDisableForceQuit ApplicationPresentationOptions = 64
const ApplicationPresentationDisableSessionTermination ApplicationPresentationOptions = 128
const ApplicationPresentationDisableHideApplication ApplicationPresentationOptions = 256
const ApplicationPresentationDisableMenuBarTransparency ApplicationPresentationOptions = 512
const ApplicationPresentationFullScreen ApplicationPresentationOptions = 1024
const ApplicationPresentationAutoHideToolbar ApplicationPresentationOptions = 2048
const ApplicationPresentationDisableCursorLocationAssistance ApplicationPresentationOptions = 4096

type ApplicationPrintReply uint

const PrintingCancelled ApplicationPrintReply = 0
const PrintingSuccess ApplicationPrintReply = 1
const PrintingFailure ApplicationPrintReply = 3
const PrintingReplyLater ApplicationPrintReply = 2

type ApplicationTerminateReply uint

const TerminateNow ApplicationTerminateReply = 1
const TerminateCancel ApplicationTerminateReply = 0
const TerminateLater ApplicationTerminateReply = 2

type AutoresizingMaskOptions uint

const ViewNotSizable AutoresizingMaskOptions = 0
const ViewMinXMargin AutoresizingMaskOptions = 1
const ViewWidthSizable AutoresizingMaskOptions = 2
const ViewMaxXMargin AutoresizingMaskOptions = 4
const ViewMinYMargin AutoresizingMaskOptions = 8
const ViewHeightSizable AutoresizingMaskOptions = 16
const ViewMaxYMargin AutoresizingMaskOptions = 32

type BackgroundStyle int

const BackgroundStyleNormal BackgroundStyle = 0
const BackgroundStyleEmphasized BackgroundStyle = 1
const BackgroundStyleRaised BackgroundStyle = 2
const BackgroundStyleLowered BackgroundStyle = 3

type BackingStoreType uint

const BackingStoreBuffered BackingStoreType = 2

type BezelStyle uint

const BezelStyleRounded BezelStyle = 1
const BezelStyleRegularSquare BezelStyle = 2
const BezelStyleShadowlessSquare BezelStyle = 6
const BezelStyleSmallSquare BezelStyle = 10
const BezelStyleRoundRect BezelStyle = 12
const BezelStyleInline BezelStyle = 15
const BezelStyleRecessed BezelStyle = 13
const BezelStyleDisclosure BezelStyle = 5
const BezelStyleRoundedDisclosure BezelStyle = 14
const BezelStyleCircular BezelStyle = 7
const BezelStyleHelpButton BezelStyle = 9
const BezelStyleTexturedRounded BezelStyle = 11
const BezelStyleTexturedSquare BezelStyle = 8

type BezierPathElement uint

const BezierPathElementMoveTo BezierPathElement = 0
const BezierPathElementLineTo BezierPathElement = 1
const BezierPathElementCurveTo BezierPathElement = 2
const BezierPathElementClosePath BezierPathElement = 3

type BitmapFormat uint

const BitmapFormatAlphaFirst BitmapFormat = 1
const BitmapFormatAlphaNonpremultiplied BitmapFormat = 2
const BitmapFormatFloatingPointSamples BitmapFormat = 4
const BitmapFormatSixteenBitBigEndian BitmapFormat = 1024
const BitmapFormatSixteenBitLittleEndian BitmapFormat = 256
const BitmapFormatThirtyTwoBitBigEndian BitmapFormat = 2048
const BitmapFormatThirtyTwoBitLittleEndian BitmapFormat = 512

type BitmapImageFileType uint

const BitmapImageFileTypeTIFF BitmapImageFileType = 0
const BitmapImageFileTypeBMP BitmapImageFileType = 1
const BitmapImageFileTypeGIF BitmapImageFileType = 2
const BitmapImageFileTypeJPEG BitmapImageFileType = 3
const BitmapImageFileTypePNG BitmapImageFileType = 4
const BitmapImageFileTypeJPEG2000 BitmapImageFileType = 5

type BitmapImageRepPropertyKey string

const ImageColorSyncProfileData BitmapImageRepPropertyKey = "NSImageColorSyncProfileData"
const ImageCompressionFactor BitmapImageRepPropertyKey = "NSImageCompressionFactor"
const ImageCompressionMethod BitmapImageRepPropertyKey = "NSImageCompressionMethod"
const ImageCurrentFrame BitmapImageRepPropertyKey = "NSImageCurrentFrame"
const ImageCurrentFrameDuration BitmapImageRepPropertyKey = "NSImageCurrentFrameDuration"
const ImageDitherTransparency BitmapImageRepPropertyKey = "NSImageDitherTransparency"
const ImageEXIFData BitmapImageRepPropertyKey = "NSImageEXIFData"
const ImageFallbackBackgroundColor BitmapImageRepPropertyKey = "NSImageFallbackBackgroundColor"
const ImageFrameCount BitmapImageRepPropertyKey = "NSImageFrameCount"
const ImageGamma BitmapImageRepPropertyKey = "NSImageGamma"
const ImageInterlaced BitmapImageRepPropertyKey = "NSImageInterlaced"
const ImageLoopCount BitmapImageRepPropertyKey = "NSImageLoopCount"
const ImageProgressive BitmapImageRepPropertyKey = "NSImageProgressive"
const ImageRGBColorTable BitmapImageRepPropertyKey = "NSImageRGBColorTable"
const ImageIPTCData BitmapImageRepPropertyKey = "NSImageIPTCData"

type BorderType uint

const BezelBorder BorderType = 2
const GrooveBorder BorderType = 3
const LineBorder BorderType = 1
const NoBorder BorderType = 0

type BoxType uint

const BoxPrimary BoxType = 0
const BoxSeparator BoxType = 2
const BoxCustom BoxType = 4

type BrowserColumnResizingType uint

const BrowserNoColumnResizing BrowserColumnResizingType = 0
const BrowserAutoColumnResizing BrowserColumnResizingType = 1
const BrowserUserColumnResizing BrowserColumnResizingType = 2

type BrowserColumnsAutosaveName string

type BrowserDropOperation uint

const BrowserDropOn BrowserDropOperation = 0
const BrowserDropAbove BrowserDropOperation = 1

type ButtonType uint

const ButtonTypeMomentaryPushIn ButtonType = 7
const ButtonTypeMomentaryLight ButtonType = 0
const ButtonTypeMomentaryChange ButtonType = 5
const ButtonTypePushOnPushOff ButtonType = 1
const ButtonTypeOnOff ButtonType = 6
const ButtonTypeToggle ButtonType = 2
const ButtonTypeSwitch ButtonType = 3
const ButtonTypeRadio ButtonType = 4
const ButtonTypeAccelerator ButtonType = 8
const ButtonTypeMultiLevelAccelerator ButtonType = 9

type CellAttribute uint

const CellAllowsMixedState CellAttribute = 16
const ChangeBackgroundCell CellAttribute = 8
const CellChangesContents CellAttribute = 14
const ChangeGrayCell CellAttribute = 4
const CellDisabled CellAttribute = 0
const CellEditable CellAttribute = 3
const CellHasImageHorizontal CellAttribute = 12
const CellHasImageOnLeftOrBottom CellAttribute = 13
const CellHasOverlappingImage CellAttribute = 11
const CellHighlighted CellAttribute = 5
const CellIsBordered CellAttribute = 10
const CellIsInsetButton CellAttribute = 15
const CellLightsByBackground CellAttribute = 9
const CellLightsByContents CellAttribute = 6
const CellLightsByGray CellAttribute = 7
const PushInCell CellAttribute = 2
const CellState CellAttribute = 1

type CellHitResult uint

const CellHitNone CellHitResult = 0
const CellHitContentArea CellHitResult = 1
const CellHitEditableTextArea CellHitResult = 2
const CellHitTrackableArea CellHitResult = 4

type CellImagePosition uint

const NoImage CellImagePosition = 0
const ImageOnly CellImagePosition = 1
const ImageLeading CellImagePosition = 7
const ImageTrailing CellImagePosition = 8
const ImageLeft CellImagePosition = 2
const ImageRight CellImagePosition = 3
const ImageBelow CellImagePosition = 4
const ImageAbove CellImagePosition = 5
const ImageOverlaps CellImagePosition = 6

type CellStyleMask uint

const NoCellMask CellStyleMask = 0
const PushInCellMask CellStyleMask = 2
const ContentsCellMask CellStyleMask = 1
const ChangeGrayCellMask CellStyleMask = 4
const ChangeBackgroundCellMask CellStyleMask = 8

type CellType uint

const NullCellType CellType = 0
const TextCellType CellType = 1
const ImageCellType CellType = 2

type CollectionElementCategory int

const CollectionElementCategoryItem CollectionElementCategory = 0
const CollectionElementCategorySupplementaryView CollectionElementCategory = 1
const CollectionElementCategoryDecorationView CollectionElementCategory = 2
const CollectionElementCategoryInterItemGap CollectionElementCategory = 3

type CollectionLayoutSectionOrthogonalScrollingBehavior int

const CollectionLayoutSectionOrthogonalScrollingBehaviorNone CollectionLayoutSectionOrthogonalScrollingBehavior = 0
const CollectionLayoutSectionOrthogonalScrollingBehaviorContinuous CollectionLayoutSectionOrthogonalScrollingBehavior = 1
const CollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary CollectionLayoutSectionOrthogonalScrollingBehavior = 2
const CollectionLayoutSectionOrthogonalScrollingBehaviorPaging CollectionLayoutSectionOrthogonalScrollingBehavior = 3
const CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging CollectionLayoutSectionOrthogonalScrollingBehavior = 4
const CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered CollectionLayoutSectionOrthogonalScrollingBehavior = 5

type CollectionUpdateAction int

const CollectionUpdateActionInsert CollectionUpdateAction = 0
const CollectionUpdateActionDelete CollectionUpdateAction = 1
const CollectionUpdateActionReload CollectionUpdateAction = 2
const CollectionUpdateActionMove CollectionUpdateAction = 3
const CollectionUpdateActionNone CollectionUpdateAction = 4

type CollectionViewDecorationElementKind string

type CollectionViewDropOperation int

const CollectionViewDropOn CollectionViewDropOperation = 0
const CollectionViewDropBefore CollectionViewDropOperation = 1

type CollectionViewItemHighlightState int

const CollectionViewItemHighlightNone CollectionViewItemHighlightState = 0
const CollectionViewItemHighlightForSelection CollectionViewItemHighlightState = 1
const CollectionViewItemHighlightForDeselection CollectionViewItemHighlightState = 2
const CollectionViewItemHighlightAsDropTarget CollectionViewItemHighlightState = 3

type CollectionViewScrollDirection int

const CollectionViewScrollDirectionVertical CollectionViewScrollDirection = 0
const CollectionViewScrollDirectionHorizontal CollectionViewScrollDirection = 1

type CollectionViewScrollPosition uint

const CollectionViewScrollPositionNone CollectionViewScrollPosition = 0
const CollectionViewScrollPositionTop CollectionViewScrollPosition = 1
const CollectionViewScrollPositionCenteredVertically CollectionViewScrollPosition = 2
const CollectionViewScrollPositionBottom CollectionViewScrollPosition = 4
const CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 512
const CollectionViewScrollPositionLeft CollectionViewScrollPosition = 8
const CollectionViewScrollPositionCenteredHorizontally CollectionViewScrollPosition = 16
const CollectionViewScrollPositionRight CollectionViewScrollPosition = 32
const CollectionViewScrollPositionLeadingEdge CollectionViewScrollPosition = 64
const CollectionViewScrollPositionTrailingEdge CollectionViewScrollPosition = 128
const CollectionViewScrollPositionNearestVerticalEdge CollectionViewScrollPosition = 256

type CollectionViewSupplementaryElementKind string

type CollectionViewTransitionLayoutAnimatedKey string

type ColorListName string

type ColorName string

type ColorPanelMode int

const ColorPanelModeNone ColorPanelMode = -1
const ColorPanelModeGray ColorPanelMode = 0
const ColorPanelModeRGB ColorPanelMode = 1
const ColorPanelModeCMYK ColorPanelMode = 2
const ColorPanelModeHSB ColorPanelMode = 3
const ColorPanelModeCustomPalette ColorPanelMode = 4
const ColorPanelModeColorList ColorPanelMode = 5
const ColorPanelModeWheel ColorPanelMode = 6
const ColorPanelModeCrayon ColorPanelMode = 7

type ColorPanelOptions uint

const ColorPanelGrayModeMask ColorPanelOptions = 1
const ColorPanelRGBModeMask ColorPanelOptions = 2
const ColorPanelCMYKModeMask ColorPanelOptions = 4
const ColorPanelHSBModeMask ColorPanelOptions = 8
const ColorPanelCustomPaletteModeMask ColorPanelOptions = 16
const ColorPanelColorListModeMask ColorPanelOptions = 32
const ColorPanelWheelModeMask ColorPanelOptions = 64
const ColorPanelCrayonModeMask ColorPanelOptions = 128
const ColorPanelAllModesMask ColorPanelOptions = 65535

type ColorRenderingIntent int

const ColorRenderingIntentDefault ColorRenderingIntent = 0
const ColorRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
const ColorRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
const ColorRenderingIntentPerceptual ColorRenderingIntent = 3
const ColorRenderingIntentSaturation ColorRenderingIntent = 4

type ColorSpaceModel int

const ColorSpaceModelUnknown ColorSpaceModel = -1
const ColorSpaceModelGray ColorSpaceModel = 0
const ColorSpaceModelRGB ColorSpaceModel = 1
const ColorSpaceModelCMYK ColorSpaceModel = 2
const ColorSpaceModelLAB ColorSpaceModel = 3
const ColorSpaceModelDeviceN ColorSpaceModel = 4
const ColorSpaceModelIndexed ColorSpaceModel = 5
const ColorSpaceModelPatterned ColorSpaceModel = 6

type ColorSpaceName string

const CalibratedRGBColorSpace ColorSpaceName = "NSCalibratedRGBColorSpace"
const CalibratedWhiteColorSpace ColorSpaceName = "NSCalibratedWhiteColorSpace"
const CustomColorSpace ColorSpaceName = "NSCustomColorSpace"
const DeviceCMYKColorSpace ColorSpaceName = "NSDeviceCMYKColorSpace"
const DeviceRGBColorSpace ColorSpaceName = "NSDeviceRGBColorSpace"
const DeviceWhiteColorSpace ColorSpaceName = "NSDeviceWhiteColorSpace"
const NamedColorSpace ColorSpaceName = "NSNamedColorSpace"
const PatternColorSpace ColorSpaceName = "NSPatternColorSpace"

type ColorSystemEffect int

const ColorSystemEffectNone ColorSystemEffect = 0
const ColorSystemEffectPressed ColorSystemEffect = 1
const ColorSystemEffectDeepPressed ColorSystemEffect = 2
const ColorSystemEffectDisabled ColorSystemEffect = 3
const ColorSystemEffectRollover ColorSystemEffect = 4

type ColorType int

const ColorTypeComponentBased ColorType = 0
const ColorTypePattern ColorType = 1
const ColorTypeCatalog ColorType = 2

type ColorWellStyle int

const ColorWellStyleDefault ColorWellStyle = 0
const ColorWellStyleMinimal ColorWellStyle = 1
const ColorWellStyleExpanded ColorWellStyle = 2

type ComboButtonStyle int

const ComboButtonStyleSplit ComboButtonStyle = 0
const ComboButtonStyleUnified ComboButtonStyle = 1

type CompositingOperation uint

const CompositingOperationClear CompositingOperation = 0
const CompositingOperationCopy CompositingOperation = 1
const CompositingOperationSourceOver CompositingOperation = 2
const CompositingOperationSourceIn CompositingOperation = 3
const CompositingOperationSourceOut CompositingOperation = 4
const CompositingOperationSourceAtop CompositingOperation = 5
const CompositingOperationDestinationOver CompositingOperation = 6
const CompositingOperationDestinationIn CompositingOperation = 7
const CompositingOperationDestinationOut CompositingOperation = 8
const CompositingOperationDestinationAtop CompositingOperation = 9
const CompositingOperationXOR CompositingOperation = 10
const CompositingOperationPlusDarker CompositingOperation = 11
const CompositingOperationPlusLighter CompositingOperation = 13
const CompositingOperationMultiply CompositingOperation = 14
const CompositingOperationScreen CompositingOperation = 15
const CompositingOperationOverlay CompositingOperation = 16
const CompositingOperationDarken CompositingOperation = 17
const CompositingOperationLighten CompositingOperation = 18
const CompositingOperationColorDodge CompositingOperation = 19
const CompositingOperationColorBurn CompositingOperation = 20
const CompositingOperationSoftLight CompositingOperation = 21
const CompositingOperationHardLight CompositingOperation = 22
const CompositingOperationDifference CompositingOperation = 23
const CompositingOperationExclusion CompositingOperation = 24
const CompositingOperationHue CompositingOperation = 25
const CompositingOperationSaturation CompositingOperation = 26
const CompositingOperationColor CompositingOperation = 27
const CompositingOperationLuminosity CompositingOperation = 28

type ControlCharacterAction int

const ControlCharacterActionContainerBreak ControlCharacterAction = 32
const ControlCharacterActionHorizontalTab ControlCharacterAction = 4
const ControlCharacterActionLineBreak ControlCharacterAction = 8
const ControlCharacterActionParagraphBreak ControlCharacterAction = 16
const ControlCharacterActionWhitespace ControlCharacterAction = 2
const ControlCharacterActionZeroAdvancement ControlCharacterAction = 1

type ControlSize uint

const ControlSizeMini ControlSize = 2
const ControlSizeSmall ControlSize = 1
const ControlSizeRegular ControlSize = 0
const ControlSizeLarge ControlSize = 3

type ControlStateValue int

const ControlStateValueOn ControlStateValue = 1
const ControlStateValueOff ControlStateValue = 0
const ControlStateValueMixed ControlStateValue = -1

type ControlTint uint

const DefaultControlTint ControlTint = 0
const BlueControlTint ControlTint = 1
const GraphiteControlTint ControlTint = 6
const ClearControlTint ControlTint = 7

type DatePickerElementFlags uint

const DatePickerElementFlagEra DatePickerElementFlags = 256
const DatePickerElementFlagHourMinute DatePickerElementFlags = 12
const DatePickerElementFlagHourMinuteSecond DatePickerElementFlags = 14
const DatePickerElementFlagTimeZone DatePickerElementFlags = 16
const DatePickerElementFlagYearMonth DatePickerElementFlags = 192
const DatePickerElementFlagYearMonthDay DatePickerElementFlags = 224

type DatePickerMode uint

const DatePickerModeRange DatePickerMode = 1
const DatePickerModeSingle DatePickerMode = 0

type DatePickerStyle uint

const DatePickerStyleClockAndCalendar DatePickerStyle = 1
const DatePickerStyleTextField DatePickerStyle = 2
const DatePickerStyleTextFieldAndStepper DatePickerStyle = 0

type DefinitionOptionKey string

const DefinitionPresentationTypeKey DefinitionOptionKey = "NSDefinitionPresentationTypeKey"

type DeviceDescriptionKey string

const DeviceBitsPerSample DeviceDescriptionKey = "NSDeviceBitsPerSample"
const DeviceColorSpaceName DeviceDescriptionKey = "NSDeviceColorSpaceName"
const DeviceIsPrinter DeviceDescriptionKey = "NSDeviceIsPrinter"
const DeviceIsScreen DeviceDescriptionKey = "NSDeviceIsScreen"
const DeviceResolution DeviceDescriptionKey = "NSDeviceResolution"
const DeviceSize DeviceDescriptionKey = "NSDeviceSize"

type DirectionalRectEdge uint

const DirectionalRectEdgeNone DirectionalRectEdge = 0
const DirectionalRectEdgeTop DirectionalRectEdge = 1
const DirectionalRectEdgeLeading DirectionalRectEdge = 2
const DirectionalRectEdgeBottom DirectionalRectEdge = 4
const DirectionalRectEdgeTrailing DirectionalRectEdge = 8
const DirectionalRectEdgeAll DirectionalRectEdge = 15

type DisplayGamut int

const DisplayGamutP3 DisplayGamut = 2
const DisplayGamutSRGB DisplayGamut = 1

type DocumentChangeType uint

const ChangeDone DocumentChangeType = 0
const ChangeUndone DocumentChangeType = 1
const ChangeCleared DocumentChangeType = 2
const ChangeReadOtherContents DocumentChangeType = 3
const ChangeAutosaved DocumentChangeType = 4
const ChangeRedone DocumentChangeType = 5
const ChangeDiscardable DocumentChangeType = 256

type DragOperation uint

const DragOperationCopy DragOperation = 1
const DragOperationLink DragOperation = 2
const DragOperationGeneric DragOperation = 4
const DragOperationPrivate DragOperation = 8
const DragOperationMove DragOperation = 16
const DragOperationDelete DragOperation = 32
const DragOperationEvery DragOperation = 18446744073709551615
const DragOperationNone DragOperation = 0

type DraggingContext int

const DraggingContextOutsideApplication DraggingContext = 0
const DraggingContextWithinApplication DraggingContext = 1

type DraggingFormation int

const DraggingFormationDefault DraggingFormation = 0
const DraggingFormationNone DraggingFormation = 1
const DraggingFormationPile DraggingFormation = 2
const DraggingFormationList DraggingFormation = 3
const DraggingFormationStack DraggingFormation = 4

type DraggingImageComponentKey string

const DraggingImageComponentIconKey DraggingImageComponentKey = "icon"
const DraggingImageComponentLabelKey DraggingImageComponentKey = "label"

type DraggingItemEnumerationOptions uint

const DraggingItemEnumerationConcurrent DraggingItemEnumerationOptions = 1
const DraggingItemEnumerationClearNonenumeratedImages DraggingItemEnumerationOptions = 65536

type EventButtonMask uint

const EventButtonMaskPenTip EventButtonMask = 1
const EventButtonMaskPenLowerSide EventButtonMask = 2
const EventButtonMaskPenUpperSide EventButtonMask = 4

type EventGestureAxis int

const EventGestureAxisNone EventGestureAxis = 0
const EventGestureAxisHorizontal EventGestureAxis = 1
const EventGestureAxisVertical EventGestureAxis = 2

type EventMask uint64

const EventMaskAny EventMask = 18446744073709551615
const EventMaskLeftMouseDown EventMask = 2
const EventMaskLeftMouseDragged EventMask = 64
const EventMaskLeftMouseUp EventMask = 4
const EventMaskRightMouseDown EventMask = 8
const EventMaskRightMouseDragged EventMask = 128
const EventMaskRightMouseUp EventMask = 16
const EventMaskOtherMouseDown EventMask = 33554432
const EventMaskOtherMouseDragged EventMask = 134217728
const EventMaskOtherMouseUp EventMask = 67108864
const EventMaskMouseEntered EventMask = 256
const EventMaskMouseMoved EventMask = 32
const EventMaskMouseExited EventMask = 512
const EventMaskKeyDown EventMask = 1024
const EventMaskKeyUp EventMask = 2048
const EventMaskBeginGesture EventMask = 524288
const EventMaskEndGesture EventMask = 1048576
const EventMaskMagnify EventMask = 1073741824
const EventMaskSmartMagnify EventMask = 4294967296
const EventMaskSwipe EventMask = 2147483648
const EventMaskRotate EventMask = 262144
const EventMaskGesture EventMask = 536870912
const EventMaskDirectTouch EventMask = 137438953472
const EventMaskTabletPoint EventMask = 8388608
const EventMaskTabletProximity EventMask = 16777216
const EventMaskPressure EventMask = 17179869184
const EventMaskScrollWheel EventMask = 4194304
const EventMaskChangeMode EventMask = 274877906944
const EventMaskAppKitDefined EventMask = 8192
const EventMaskApplicationDefined EventMask = 32768
const EventMaskCursorUpdate EventMask = 131072
const EventMaskFlagsChanged EventMask = 4096
const EventMaskPeriodic EventMask = 65536
const EventMaskSystemDefined EventMask = 16384

type EventModifierFlags uint

const EventModifierFlagCapsLock EventModifierFlags = 65536
const EventModifierFlagShift EventModifierFlags = 131072
const EventModifierFlagControl EventModifierFlags = 262144
const EventModifierFlagOption EventModifierFlags = 524288
const EventModifierFlagCommand EventModifierFlags = 1048576
const EventModifierFlagNumericPad EventModifierFlags = 2097152
const EventModifierFlagHelp EventModifierFlags = 4194304
const EventModifierFlagFunction EventModifierFlags = 8388608
const EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 4294901760

type EventPhase uint

const EventPhaseNone EventPhase = 0
const EventPhaseBegan EventPhase = 1
const EventPhaseStationary EventPhase = 2
const EventPhaseChanged EventPhase = 4
const EventPhaseEnded EventPhase = 8
const EventPhaseCancelled EventPhase = 16
const EventPhaseMayBegin EventPhase = 32

type EventSubtype int16

const EventSubtypeApplicationActivated EventSubtype = 1
const EventSubtypeApplicationDeactivated EventSubtype = 2
const EventSubtypeScreenChanged EventSubtype = 8
const EventSubtypeWindowExposed EventSubtype = 0
const EventSubtypeWindowMoved EventSubtype = 4
const EventSubtypePowerOff EventSubtype = 1
const EventSubtypeMouseEvent EventSubtype = 0
const EventSubtypeTabletPoint EventSubtype = 1
const EventSubtypeTabletProximity EventSubtype = 2
const EventSubtypeTouch EventSubtype = 3

type EventSwipeTrackingOptions uint

const EventSwipeTrackingLockDirection EventSwipeTrackingOptions = 1
const EventSwipeTrackingClampGestureAmount EventSwipeTrackingOptions = 2

type EventType uint

const EventTypeLeftMouseDown EventType = 1
const EventTypeLeftMouseDragged EventType = 6
const EventTypeLeftMouseUp EventType = 2
const EventTypeRightMouseDown EventType = 3
const EventTypeRightMouseUp EventType = 4
const EventTypeRightMouseDragged EventType = 7
const EventTypeOtherMouseDown EventType = 25
const EventTypeOtherMouseDragged EventType = 27
const EventTypeOtherMouseUp EventType = 26
const EventTypeMouseMoved EventType = 5
const EventTypeMouseEntered EventType = 8
const EventTypeMouseExited EventType = 9
const EventTypeKeyDown EventType = 10
const EventTypeKeyUp EventType = 11
const EventTypeBeginGesture EventType = 19
const EventTypeEndGesture EventType = 20
const EventTypeMagnify EventType = 30
const EventTypeSmartMagnify EventType = 32
const EventTypeSwipe EventType = 31
const EventTypeRotate EventType = 18
const EventTypeGesture EventType = 29
const EventTypeDirectTouch EventType = 37
const EventTypeTabletPoint EventType = 23
const EventTypeTabletProximity EventType = 24
const EventTypePressure EventType = 34
const EventTypeScrollWheel EventType = 22
const EventTypeChangeMode EventType = 38
const EventTypeAppKitDefined EventType = 13
const EventTypeApplicationDefined EventType = 15
const EventTypeCursorUpdate EventType = 17
const EventTypeFlagsChanged EventType = 12
const EventTypePeriodic EventType = 16
const EventTypeQuickLook EventType = 33
const EventTypeSystemDefined EventType = 14

type FocusRingType uint

const FocusRingTypeDefault FocusRingType = 0
const FocusRingTypeNone FocusRingType = 1
const FocusRingTypeExterior FocusRingType = 2

type FontAction uint

const NoFontChangeAction FontAction = 0
const ViaPanelFontAction FontAction = 1
const AddTraitFontAction FontAction = 2
const SizeUpFontAction FontAction = 3
const SizeDownFontAction FontAction = 4
const HeavierFontAction FontAction = 5
const LighterFontAction FontAction = 6
const RemoveTraitFontAction FontAction = 7

type FontCollectionOptions uint

const FontCollectionApplicationOnlyMask FontCollectionOptions = 1

type FontDescriptorAttributeName string

const FontFamilyAttribute FontDescriptorAttributeName = "NSFontFamilyAttribute"
const FontNameAttribute FontDescriptorAttributeName = "NSFontNameAttribute"
const FontFaceAttribute FontDescriptorAttributeName = "NSFontFaceAttribute"
const FontSizeAttribute FontDescriptorAttributeName = "NSFontSizeAttribute"
const FontVisibleNameAttribute FontDescriptorAttributeName = "NSFontVisibleNameAttribute"
const FontMatrixAttribute FontDescriptorAttributeName = "NSFontMatrixAttribute"
const FontVariationAttribute FontDescriptorAttributeName = "NSCTFontVariationAttribute"
const FontCharacterSetAttribute FontDescriptorAttributeName = "NSCTFontCharacterSetAttribute"
const FontCascadeListAttribute FontDescriptorAttributeName = "NSCTFontCascadeListAttribute"
const FontTraitsAttribute FontDescriptorAttributeName = "NSCTFontTraitsAttribute"
const FontFixedAdvanceAttribute FontDescriptorAttributeName = "NSCTFontFixedAdvanceAttribute"
const FontFeatureSettingsAttribute FontDescriptorAttributeName = "NSCTFontFeatureSettingsAttribute"

type FontDescriptorSymbolicTraits uint32

const FontDescriptorTraitItalic FontDescriptorSymbolicTraits = 1
const FontDescriptorTraitBold FontDescriptorSymbolicTraits = 2
const FontDescriptorTraitExpanded FontDescriptorSymbolicTraits = 32
const FontDescriptorTraitCondensed FontDescriptorSymbolicTraits = 64
const FontDescriptorTraitMonoSpace FontDescriptorSymbolicTraits = 1024
const FontDescriptorTraitVertical FontDescriptorSymbolicTraits = 2048
const FontDescriptorTraitUIOptimized FontDescriptorSymbolicTraits = 4096
const FontDescriptorTraitTightLeading FontDescriptorSymbolicTraits = 32768
const FontDescriptorTraitLooseLeading FontDescriptorSymbolicTraits = 65536
const FontDescriptorClassMask FontDescriptorSymbolicTraits = 4026531840
const FontDescriptorClassUnknown FontDescriptorSymbolicTraits = 0
const FontDescriptorClassOldStyleSerifs FontDescriptorSymbolicTraits = 268435456
const FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 536870912
const FontDescriptorClassModernSerifs FontDescriptorSymbolicTraits = 805306368
const FontDescriptorClassClarendonSerifs FontDescriptorSymbolicTraits = 1073741824
const FontDescriptorClassSlabSerifs FontDescriptorSymbolicTraits = 1342177280
const FontDescriptorClassFreeformSerifs FontDescriptorSymbolicTraits = 1879048192
const FontDescriptorClassSansSerif FontDescriptorSymbolicTraits = 2147483648
const FontDescriptorClassOrnamentals FontDescriptorSymbolicTraits = 2415919104
const FontDescriptorClassScripts FontDescriptorSymbolicTraits = 2684354560
const FontDescriptorClassSymbolic FontDescriptorSymbolicTraits = 3221225472

type FontDescriptorSystemDesign string

const FontDescriptorSystemDesignDefault FontDescriptorSystemDesign = "NSCTFontUIFontDesignDefault"
const FontDescriptorSystemDesignMonospaced FontDescriptorSystemDesign = "NSCTFontUIFontDesignMonospaced"
const FontDescriptorSystemDesignRounded FontDescriptorSystemDesign = "NSCTFontUIFontDesignRounded"
const FontDescriptorSystemDesignSerif FontDescriptorSystemDesign = "NSCTFontUIFontDesignSerif"

type FontPanelModeMask uint

const FontPanelModeMaskAllEffects FontPanelModeMask = 1048320
const FontPanelModesMaskAllModes FontPanelModeMask = 4294967295
const FontPanelModeMaskCollection FontPanelModeMask = 4
const FontPanelModeMaskDocumentColorEffect FontPanelModeMask = 2048
const FontPanelModeMaskFace FontPanelModeMask = 1
const FontPanelModeMaskShadowEffect FontPanelModeMask = 4096
const FontPanelModeMaskSize FontPanelModeMask = 2
const FontPanelModesMaskStandardModes FontPanelModeMask = 65535
const FontPanelModeMaskStrikethroughEffect FontPanelModeMask = 512
const FontPanelModeMaskTextColorEffect FontPanelModeMask = 1024
const FontPanelModeMaskUnderlineEffect FontPanelModeMask = 256

type FontRenderingMode uint

const FontDefaultRenderingMode FontRenderingMode = 0
const FontAntialiasedRenderingMode FontRenderingMode = 1
const FontIntegerAdvancementsRenderingMode FontRenderingMode = 2
const FontAntialiasedIntegerAdvancementsRenderingMode FontRenderingMode = 3

type FontTextStyle string

const FontTextStyleBody FontTextStyle = "UICTFontTextStyleBody"
const FontTextStyleCallout FontTextStyle = "UICTFontTextStyleCallout"
const FontTextStyleCaption1 FontTextStyle = "UICTFontTextStyleCaption1"
const FontTextStyleCaption2 FontTextStyle = "UICTFontTextStyleCaption2"
const FontTextStyleFootnote FontTextStyle = "UICTFontTextStyleFootnote"
const FontTextStyleHeadline FontTextStyle = "UICTFontTextStyleHeadline"
const FontTextStyleSubheadline FontTextStyle = "UICTFontTextStyleSubhead"
const FontTextStyleLargeTitle FontTextStyle = "UICTFontTextStyleTitle0"
const FontTextStyleTitle1 FontTextStyle = "UICTFontTextStyleTitle1"
const FontTextStyleTitle2 FontTextStyle = "UICTFontTextStyleTitle2"
const FontTextStyleTitle3 FontTextStyle = "UICTFontTextStyleTitle3"

type FontTextStyleOptionKey string

type FontTraitMask uint

const BoldFontMask FontTraitMask = 2
const CompressedFontMask FontTraitMask = 512
const CondensedFontMask FontTraitMask = 64
const ExpandedFontMask FontTraitMask = 32
const FixedPitchFontMask FontTraitMask = 1024
const ItalicFontMask FontTraitMask = 1
const NarrowFontMask FontTraitMask = 16
const NonStandardCharacterSetFontMask FontTraitMask = 8
const PosterFontMask FontTraitMask = 256
const SmallCapsFontMask FontTraitMask = 128
const UnboldFontMask FontTraitMask = 4
const UnitalicFontMask FontTraitMask = 16777216

type FontWeight float64

const FontWeightUltraLight FontWeight = -0.800000011920929
const FontWeightThin FontWeight = -0.6000000238418579
const FontWeightLight FontWeight = -0.4000000059604645
const FontWeightRegular FontWeight = 0.0
const FontWeightMedium FontWeight = 0.23000000417232513
const FontWeightSemibold FontWeight = 0.30000001192092896
const FontWeightBold FontWeight = 0.4000000059604645
const FontWeightHeavy FontWeight = 0.5600000023841858
const FontWeightBlack FontWeight = 0.6200000047683716

type FontWidth float64

const FontWidthCompressed FontWidth = -0.3
const FontWidthCondensed FontWidth = -0.2
const FontWidthExpanded FontWidth = 0.2
const FontWidthStandard FontWidth = 0.0

type GestureRecognizerState int

const GestureRecognizerStatePossible GestureRecognizerState = 0
const GestureRecognizerStateBegan GestureRecognizerState = 1
const GestureRecognizerStateChanged GestureRecognizerState = 2
const GestureRecognizerStateEnded GestureRecognizerState = 3
const GestureRecognizerStateCancelled GestureRecognizerState = 4
const GestureRecognizerStateFailed GestureRecognizerState = 5
const GestureRecognizerStateRecognized GestureRecognizerState = 3

type Glyph uint32

type GlyphInscription uint

type GlyphProperty int

const GlyphPropertyNull GlyphProperty = 1
const GlyphPropertyControlCharacter GlyphProperty = 2
const GlyphPropertyElastic GlyphProperty = 4
const GlyphPropertyNonBaseCharacter GlyphProperty = 8

type GradientType uint

type GraphicsContextAttributeKey string

const GraphicsContextDestinationAttributeName GraphicsContextAttributeKey = "NSGraphicsContextDestinationAttributeName"
const GraphicsContextRepresentationFormatAttributeName GraphicsContextAttributeKey = "NSGraphicsContextRepresentationFormatAttributeName"

type GridCellPlacement int

const GridCellPlacementTop GridCellPlacement = 2
const GridCellPlacementBottom GridCellPlacement = 3
const GridCellPlacementCenter GridCellPlacement = 4
const GridCellPlacementFill GridCellPlacement = 5
const GridCellPlacementInherited GridCellPlacement = 0
const GridCellPlacementLeading GridCellPlacement = 2
const GridCellPlacementNone GridCellPlacement = 1
const GridCellPlacementTrailing GridCellPlacement = 3

type GridRowAlignment int

const GridRowAlignmentFirstBaseline GridRowAlignment = 2
const GridRowAlignmentInherited GridRowAlignment = 0
const GridRowAlignmentLastBaseline GridRowAlignment = 3
const GridRowAlignmentNone GridRowAlignment = 1

type HelpAnchorName string

type ImageAlignment uint

const ImageAlignCenter ImageAlignment = 0
const ImageAlignTop ImageAlignment = 1
const ImageAlignTopLeft ImageAlignment = 2
const ImageAlignTopRight ImageAlignment = 3
const ImageAlignLeft ImageAlignment = 4
const ImageAlignBottom ImageAlignment = 5
const ImageAlignBottomLeft ImageAlignment = 6
const ImageAlignBottomRight ImageAlignment = 7
const ImageAlignRight ImageAlignment = 8

type ImageCacheMode uint

const ImageCacheDefault ImageCacheMode = 0
const ImageCacheAlways ImageCacheMode = 1
const ImageCacheBySize ImageCacheMode = 2
const ImageCacheNever ImageCacheMode = 3

type ImageFrameStyle uint

const ImageFrameNone ImageFrameStyle = 0
const ImageFramePhoto ImageFrameStyle = 1
const ImageFrameGrayBezel ImageFrameStyle = 2
const ImageFrameGroove ImageFrameStyle = 3
const ImageFrameButton ImageFrameStyle = 4

type ImageHintKey string

const ImageHintCTM ImageHintKey = "NSImageHintCTM"
const ImageHintInterpolation ImageHintKey = "NSImageHintInterpolation"
const ImageHintUserInterfaceLayoutDirection ImageHintKey = "NSImageHintUserInterfaceLayoutDirection"

type ImageInterpolation uint

const ImageInterpolationDefault ImageInterpolation = 0
const ImageInterpolationNone ImageInterpolation = 1
const ImageInterpolationLow ImageInterpolation = 2
const ImageInterpolationMedium ImageInterpolation = 4
const ImageInterpolationHigh ImageInterpolation = 3

type ImageLayoutDirection int

const ImageLayoutDirectionUnspecified ImageLayoutDirection = -1
const ImageLayoutDirectionLeftToRight ImageLayoutDirection = 2
const ImageLayoutDirectionRightToLeft ImageLayoutDirection = 3

type ImageLoadStatus uint

const ImageLoadStatusCompleted ImageLoadStatus = 0
const ImageLoadStatusCancelled ImageLoadStatus = 1
const ImageLoadStatusInvalidData ImageLoadStatus = 2
const ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
const ImageLoadStatusReadError ImageLoadStatus = 4

type ImageName string

const ImageNameActionTemplate ImageName = "NSActionTemplate"
const ImageNameAddTemplate ImageName = "NSAddTemplate"
const ImageNameAdvanced ImageName = "NSAdvanced"
const ImageNameApplicationIcon ImageName = "NSApplicationIcon"
const ImageNameBluetoothTemplate ImageName = "NSBluetoothTemplate"
const ImageNameBonjour ImageName = "NSBonjour"
const ImageNameBookmarksTemplate ImageName = "NSBookmarksTemplate"
const ImageNameCaution ImageName = "NSCaution"
const ImageNameColorPanel ImageName = "NSColorPanel"
const ImageNameColumnViewTemplate ImageName = "NSColumnViewTemplate"
const ImageNameComputer ImageName = "NSComputer"
const ImageNameEnterFullScreenTemplate ImageName = "NSEnterFullScreenTemplate"
const ImageNameEveryone ImageName = "NSEveryone"
const ImageNameExitFullScreenTemplate ImageName = "NSExitFullScreenTemplate"
const ImageNameFlowViewTemplate ImageName = "NSFlowViewTemplate"
const ImageNameFolder ImageName = "NSFolder"
const ImageNameFolderBurnable ImageName = "NSFolderBurnable"
const ImageNameFolderSmart ImageName = "NSFolderSmart"
const ImageNameFollowLinkFreestandingTemplate ImageName = "NSFollowLinkFreestandingTemplate"
const ImageNameFontPanel ImageName = "NSFontPanel"
const ImageNameGoBackTemplate ImageName = "NSGoBackTemplate"
const ImageNameGoForwardTemplate ImageName = "NSGoForwardTemplate"
const ImageNameGoLeftTemplate ImageName = "NSGoLeftTemplate"
const ImageNameGoRightTemplate ImageName = "NSGoRightTemplate"
const ImageNameHomeTemplate ImageName = "NSHomeTemplate"
const ImageNameIChatTheaterTemplate ImageName = "NSIChatTheaterTemplate"
const ImageNameIconViewTemplate ImageName = "NSIconViewTemplate"
const ImageNameInfo ImageName = "NSInfo"
const ImageNameInvalidDataFreestandingTemplate ImageName = "NSInvalidDataFreestandingTemplate"
const ImageNameLeftFacingTriangleTemplate ImageName = "NSLeftFacingTriangleTemplate"
const ImageNameListViewTemplate ImageName = "NSListViewTemplate"
const ImageNameLockLockedTemplate ImageName = "NSLockLockedTemplate"
const ImageNameLockUnlockedTemplate ImageName = "NSLockUnlockedTemplate"
const ImageNameMenuMixedStateTemplate ImageName = "NSMenuMixedStateTemplate"
const ImageNameMenuOnStateTemplate ImageName = "NSMenuOnStateTemplate"
const ImageNameMobileMe ImageName = "NSMobileMe"
const ImageNameMultipleDocuments ImageName = "NSMultipleDocuments"
const ImageNameNetwork ImageName = "NSNetwork"
const ImageNamePathTemplate ImageName = "NSPathTemplate"
const ImageNamePreferencesGeneral ImageName = "NSPreferencesGeneral"
const ImageNameQuickLookTemplate ImageName = "NSQuickLookTemplate"
const ImageNameRefreshFreestandingTemplate ImageName = "NSRefreshFreestandingTemplate"
const ImageNameRefreshTemplate ImageName = "NSRefreshTemplate"
const ImageNameRemoveTemplate ImageName = "NSRemoveTemplate"
const ImageNameRevealFreestandingTemplate ImageName = "NSRevealFreestandingTemplate"
const ImageNameRightFacingTriangleTemplate ImageName = "NSRightFacingTriangleTemplate"
const ImageNameShareTemplate ImageName = "NSShareTemplate"
const ImageNameSlideshowTemplate ImageName = "NSSlideshowTemplate"
const ImageNameSmartBadgeTemplate ImageName = "NSSmartBadgeTemplate"
const ImageNameStatusAvailable ImageName = "NSStatusAvailable"
const ImageNameStatusNone ImageName = "NSStatusNone"
const ImageNameStatusPartiallyAvailable ImageName = "NSStatusPartiallyAvailable"
const ImageNameStatusUnavailable ImageName = "NSStatusUnavailable"
const ImageNameStopProgressFreestandingTemplate ImageName = "NSStopProgressFreestandingTemplate"
const ImageNameStopProgressTemplate ImageName = "NSStopProgressTemplate"
const ImageNameTouchBarAddDetailTemplate ImageName = "NSTouchBarAddDetailTemplate"
const ImageNameTouchBarAddTemplate ImageName = "NSTouchBarAddTemplate"
const ImageNameTouchBarAlarmTemplate ImageName = "NSTouchBarAlarmTemplate"
const ImageNameTouchBarAudioInputMuteTemplate ImageName = "NSTouchBarAudioInputMuteTemplate"
const ImageNameTouchBarAudioInputTemplate ImageName = "NSTouchBarAudioInputTemplate"
const ImageNameTouchBarAudioOutputMuteTemplate ImageName = "NSTouchBarAudioOutputMuteTemplate"
const ImageNameTouchBarAudioOutputVolumeHighTemplate ImageName = "NSTouchBarAudioOutputVolumeHighTemplate"
const ImageNameTouchBarAudioOutputVolumeLowTemplate ImageName = "NSTouchBarAudioOutputVolumeLowTemplate"
const ImageNameTouchBarAudioOutputVolumeMediumTemplate ImageName = "NSTouchBarAudioOutputVolumeMediumTemplate"
const ImageNameTouchBarAudioOutputVolumeOffTemplate ImageName = "NSTouchBarAudioOutputVolumeOffTemplate"
const ImageNameTouchBarBookmarksTemplate ImageName = "NSTouchBarBookmarksTemplate"
const ImageNameTouchBarColorPickerFill ImageName = "NSTouchBarColorPickerFill"
const ImageNameTouchBarColorPickerFont ImageName = "NSTouchBarColorPickerFont"
const ImageNameTouchBarColorPickerStroke ImageName = "NSTouchBarColorPickerStroke"
const ImageNameTouchBarCommunicationAudioTemplate ImageName = "NSTouchBarCommunicationAudioTemplate"
const ImageNameTouchBarCommunicationVideoTemplate ImageName = "NSTouchBarCommunicationVideoTemplate"
const ImageNameTouchBarComposeTemplate ImageName = "NSTouchBarComposeTemplate"
const ImageNameTouchBarDeleteTemplate ImageName = "NSTouchBarDeleteTemplate"
const ImageNameTouchBarDownloadTemplate ImageName = "NSTouchBarDownloadTemplate"
const ImageNameTouchBarEnterFullScreenTemplate ImageName = "NSTouchBarEnterFullScreenTemplate"
const ImageNameTouchBarExitFullScreenTemplate ImageName = "NSTouchBarExitFullScreenTemplate"
const ImageNameTouchBarFastForwardTemplate ImageName = "NSTouchBarFastForwardTemplate"
const ImageNameTouchBarFolderCopyToTemplate ImageName = "NSTouchBarFolderCopyToTemplate"
const ImageNameTouchBarFolderMoveToTemplate ImageName = "NSTouchBarFolderMoveToTemplate"
const ImageNameTouchBarFolderTemplate ImageName = "NSTouchBarFolderTemplate"
const ImageNameTouchBarGetInfoTemplate ImageName = "NSTouchBarGetInfoTemplate"
const ImageNameTouchBarGoBackTemplate ImageName = "NSTouchBarGoBackTemplate"
const ImageNameTouchBarGoDownTemplate ImageName = "NSTouchBarGoDownTemplate"
const ImageNameTouchBarGoForwardTemplate ImageName = "NSTouchBarGoForwardTemplate"
const ImageNameTouchBarGoUpTemplate ImageName = "NSTouchBarGoUpTemplate"
const ImageNameTouchBarHistoryTemplate ImageName = "NSTouchBarHistoryTemplate"
const ImageNameTouchBarIconViewTemplate ImageName = "NSTouchBarIconViewTemplate"
const ImageNameTouchBarListViewTemplate ImageName = "NSTouchBarListViewTemplate"
const ImageNameTouchBarMailTemplate ImageName = "NSTouchBarMailTemplate"
const ImageNameTouchBarNewFolderTemplate ImageName = "NSTouchBarNewFolderTemplate"
const ImageNameTouchBarNewMessageTemplate ImageName = "NSTouchBarNewMessageTemplate"
const ImageNameTouchBarOpenInBrowserTemplate ImageName = "NSTouchBarOpenInBrowserTemplate"
const ImageNameTouchBarPauseTemplate ImageName = "NSTouchBarPauseTemplate"
const ImageNameTouchBarPlayPauseTemplate ImageName = "NSTouchBarPlayPauseTemplate"
const ImageNameTouchBarPlayTemplate ImageName = "NSTouchBarPlayTemplate"
const ImageNameTouchBarPlayheadTemplate ImageName = "NSTouchBarPlayheadTemplate"
const ImageNameTouchBarQuickLookTemplate ImageName = "NSTouchBarQuickLookTemplate"
const ImageNameTouchBarRecordStartTemplate ImageName = "NSTouchBarRecordStartTemplate"
const ImageNameTouchBarRecordStopTemplate ImageName = "NSTouchBarRecordStopTemplate"
const ImageNameTouchBarRefreshTemplate ImageName = "NSTouchBarRefreshTemplate"
const ImageNameTouchBarRemoveTemplate ImageName = "NSTouchBarRemoveTemplate"
const ImageNameTouchBarRewindTemplate ImageName = "NSTouchBarRewindTemplate"
const ImageNameTouchBarRotateLeftTemplate ImageName = "NSTouchBarRotateLeftTemplate"
const ImageNameTouchBarRotateRightTemplate ImageName = "NSTouchBarRotateRightTemplate"
const ImageNameTouchBarSearchTemplate ImageName = "NSTouchBarSearchTemplate"
const ImageNameTouchBarShareTemplate ImageName = "NSTouchBarShareTemplate"
const ImageNameTouchBarSidebarTemplate ImageName = "NSTouchBarSidebarTemplate"
const ImageNameTouchBarSkipAhead15SecondsTemplate ImageName = "NSTouchBarSkipAhead15SecondsTemplate"
const ImageNameTouchBarSkipAhead30SecondsTemplate ImageName = "NSTouchBarSkipAhead30SecondsTemplate"
const ImageNameTouchBarSkipAheadTemplate ImageName = "NSTouchBarSkipAheadTemplate"
const ImageNameTouchBarSkipBack15SecondsTemplate ImageName = "NSTouchBarSkipBack15SecondsTemplate"
const ImageNameTouchBarSkipBack30SecondsTemplate ImageName = "NSTouchBarSkipBack30SecondsTemplate"
const ImageNameTouchBarSkipBackTemplate ImageName = "NSTouchBarSkipBackTemplate"
const ImageNameTouchBarSkipToEndTemplate ImageName = "NSTouchBarSkipToEndTemplate"
const ImageNameTouchBarSkipToStartTemplate ImageName = "NSTouchBarSkipToStartTemplate"
const ImageNameTouchBarSlideshowTemplate ImageName = "NSTouchBarSlideshowTemplate"
const ImageNameTouchBarTagIconTemplate ImageName = "NSTouchBarTagIconTemplate"
const ImageNameTouchBarTextBoldTemplate ImageName = "NSTouchBarTextBoldTemplate"
const ImageNameTouchBarTextBoxTemplate ImageName = "NSTouchBarTextBoxTemplate"
const ImageNameTouchBarTextCenterAlignTemplate ImageName = "NSTouchBarTextCenterAlignTemplate"
const ImageNameTouchBarTextItalicTemplate ImageName = "NSTouchBarTextItalicTemplate"
const ImageNameTouchBarTextJustifiedAlignTemplate ImageName = "NSTouchBarTextJustifiedAlignTemplate"
const ImageNameTouchBarTextLeftAlignTemplate ImageName = "NSTouchBarTextLeftAlignTemplate"
const ImageNameTouchBarTextListTemplate ImageName = "NSTouchBarTextListTemplate"
const ImageNameTouchBarTextRightAlignTemplate ImageName = "NSTouchBarTextRightAlignTemplate"
const ImageNameTouchBarTextStrikethroughTemplate ImageName = "NSTouchBarTextStrikethroughTemplate"
const ImageNameTouchBarTextUnderlineTemplate ImageName = "NSTouchBarTextUnderlineTemplate"
const ImageNameTouchBarUserAddTemplate ImageName = "NSTouchBarUserAddTemplate"
const ImageNameTouchBarUserGroupTemplate ImageName = "NSTouchBarUserGroupTemplate"
const ImageNameTouchBarUserTemplate ImageName = "NSTouchBarUserTemplate"
const ImageNameTouchBarVolumeDownTemplate ImageName = "NSTouchBarVolumeDownTemplate"
const ImageNameTouchBarVolumeUpTemplate ImageName = "NSTouchBarVolumeUpTemplate"
const ImageNameTrashEmpty ImageName = "NSTrashEmpty"
const ImageNameTrashFull ImageName = "NSTrashFull"
const ImageNameUser ImageName = "NSUser"
const ImageNameUserAccounts ImageName = "NSUserAccounts"
const ImageNameUserGroup ImageName = "NSUserGroup"
const ImageNameUserGuest ImageName = "NSUserGuest"

type ImageResizingMode int

type ImageScaling uint

const ImageScaleProportionallyDown ImageScaling = 0
const ImageScaleAxesIndependently ImageScaling = 1
const ImageScaleNone ImageScaling = 2
const ImageScaleProportionallyUpOrDown ImageScaling = 3

type ImageSymbolScale int

const ImageSymbolScaleSmall ImageSymbolScale = 1
const ImageSymbolScaleMedium ImageSymbolScale = 2
const ImageSymbolScaleLarge ImageSymbolScale = 3

type InterfaceStyle uint

type LayoutAttribute int

const LayoutAttributeLeft LayoutAttribute = 1
const LayoutAttributeRight LayoutAttribute = 2
const LayoutAttributeTop LayoutAttribute = 3
const LayoutAttributeBottom LayoutAttribute = 4
const LayoutAttributeLeading LayoutAttribute = 5
const LayoutAttributeTrailing LayoutAttribute = 6
const LayoutAttributeWidth LayoutAttribute = 7
const LayoutAttributeHeight LayoutAttribute = 8
const LayoutAttributeCenterX LayoutAttribute = 9
const LayoutAttributeCenterY LayoutAttribute = 10
const LayoutAttributeBaseline LayoutAttribute = 11
const LayoutAttributeLastBaseline LayoutAttribute = 11
const LayoutAttributeFirstBaseline LayoutAttribute = 12
const LayoutAttributeNotAnAttribute LayoutAttribute = 0

type LayoutConstraintOrientation int

const LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
const LayoutConstraintOrientationVertical LayoutConstraintOrientation = 1

type LayoutFormatOptions uint

const LayoutFormatAlignAllLeft LayoutFormatOptions = 2
const LayoutFormatAlignAllRight LayoutFormatOptions = 4
const LayoutFormatAlignAllTop LayoutFormatOptions = 8
const LayoutFormatAlignAllBottom LayoutFormatOptions = 16
const LayoutFormatAlignAllLeading LayoutFormatOptions = 32
const LayoutFormatAlignAllTrailing LayoutFormatOptions = 64
const LayoutFormatAlignAllCenterX LayoutFormatOptions = 512
const LayoutFormatAlignAllCenterY LayoutFormatOptions = 1024
const LayoutFormatAlignAllBaseline LayoutFormatOptions = 2048
const LayoutFormatAlignAllLastBaseline LayoutFormatOptions = 2048
const LayoutFormatAlignAllFirstBaseline LayoutFormatOptions = 4096
const LayoutFormatAlignmentMask LayoutFormatOptions = 65535
const LayoutFormatDirectionLeadingToTrailing LayoutFormatOptions = 0
const LayoutFormatDirectionLeftToRight LayoutFormatOptions = 65536
const LayoutFormatDirectionRightToLeft LayoutFormatOptions = 131072
const LayoutFormatDirectionMask LayoutFormatOptions = 196608

type LayoutPriority float32

const LayoutPriorityRequired LayoutPriority = 1000
const LayoutPriorityDefaultHigh LayoutPriority = 750
const LayoutPriorityDragThatCanResizeWindow LayoutPriority = 510
const LayoutPriorityWindowSizeStayPut LayoutPriority = 500
const LayoutPriorityDragThatCannotResizeWindow LayoutPriority = 490
const LayoutPriorityDefaultLow LayoutPriority = 250
const LayoutPriorityFittingSizeCompression LayoutPriority = 50

type LayoutRelation int

const LayoutRelationLessThanOrEqual LayoutRelation = -1
const LayoutRelationEqual LayoutRelation = 0
const LayoutRelationGreaterThanOrEqual LayoutRelation = 1

type LevelIndicatorPlaceholderVisibility int

const LevelIndicatorPlaceholderVisibilityAutomatic LevelIndicatorPlaceholderVisibility = 0
const LevelIndicatorPlaceholderVisibilityAlways LevelIndicatorPlaceholderVisibility = 1
const LevelIndicatorPlaceholderVisibilityWhileEditing LevelIndicatorPlaceholderVisibility = 2

type LevelIndicatorStyle uint

const LevelIndicatorStyleRelevancy LevelIndicatorStyle = 0
const LevelIndicatorStyleContinuousCapacity LevelIndicatorStyle = 1
const LevelIndicatorStyleDiscreteCapacity LevelIndicatorStyle = 2
const LevelIndicatorStyleRating LevelIndicatorStyle = 3

type LineBreakMode uint

const LineBreakByWordWrapping LineBreakMode = 0
const LineBreakByCharWrapping LineBreakMode = 1
const LineBreakByClipping LineBreakMode = 2
const LineBreakByTruncatingHead LineBreakMode = 3
const LineBreakByTruncatingTail LineBreakMode = 4
const LineBreakByTruncatingMiddle LineBreakMode = 5

type LineBreakStrategy uint

const LineBreakStrategyNone LineBreakStrategy = 0
const LineBreakStrategyPushOut LineBreakStrategy = 1
const LineBreakStrategyHangulWordPriority LineBreakStrategy = 2
const LineBreakStrategyStandard LineBreakStrategy = 65535

type LineCapStyle uint

const LineCapStyleButt LineCapStyle = 0
const LineCapStyleRound LineCapStyle = 1
const LineCapStyleSquare LineCapStyle = 2

type LineJoinStyle uint

const LineJoinStyleMiter LineJoinStyle = 0
const LineJoinStyleRound LineJoinStyle = 1
const LineJoinStyleBevel LineJoinStyle = 2

type LineMovementDirection uint

const LineMovesLeft LineMovementDirection = 1
const LineMovesRight LineMovementDirection = 2
const LineMovesDown LineMovementDirection = 3
const LineMovesUp LineMovementDirection = 4
const LineDoesntMove LineMovementDirection = 0

type LineSweepDirection uint

const LineSweepLeft LineSweepDirection = 0
const LineSweepRight LineSweepDirection = 1
const LineSweepDown LineSweepDirection = 2
const LineSweepUp LineSweepDirection = 3

type MatrixMode uint

const TrackModeMatrix MatrixMode = 3
const HighlightModeMatrix MatrixMode = 1
const RadioModeMatrix MatrixMode = 0
const ListModeMatrix MatrixMode = 2

type MenuProperties uint

const MenuPropertyItemTitle MenuProperties = 1
const MenuPropertyItemAttributedTitle MenuProperties = 2
const MenuPropertyItemKeyEquivalent MenuProperties = 4
const MenuPropertyItemImage MenuProperties = 8
const MenuPropertyItemEnabled MenuProperties = 16
const MenuPropertyItemAccessibilityDescription MenuProperties = 32

type ModalResponse int

const ModalResponseOK ModalResponse = 1
const ModalResponseCancel ModalResponse = 0
const ModalResponseContinue ModalResponse = -1002
const ModalResponseStop ModalResponse = -1000
const ModalResponseAbort ModalResponse = -1001
const AlertFirstButtonReturn ModalResponse = 1000
const AlertSecondButtonReturn ModalResponse = 1001
const AlertThirdButtonReturn ModalResponse = 1002

type ModalSession unsafe.Pointer

type NibName string

type PDFPanelOptions int

const PDFPanelShowsPaperSize PDFPanelOptions = 4
const PDFPanelShowsOrientation PDFPanelOptions = 8
const PDFPanelRequestsParentDirectory PDFPanelOptions = 16777216

type PaperOrientation int

const PaperOrientationPortrait PaperOrientation = 0
const PaperOrientationLandscape PaperOrientation = 1

type PasteboardContentsOptions uint

const PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1

type PasteboardName string

const PasteboardNameDrag PasteboardName = "Apple CFPasteboard drag"
const PasteboardNameFind PasteboardName = "Apple CFPasteboard find"
const PasteboardNameFont PasteboardName = "Apple CFPasteboard font"
const PasteboardNameGeneral PasteboardName = "Apple CFPasteboard general"
const PasteboardNameRuler PasteboardName = "Apple CFPasteboard ruler"

type PasteboardReadingOptionKey string

const PasteboardURLReadingContentsConformToTypesKey PasteboardReadingOptionKey = "NSPasteboardURLReadingContentsConformToTypesKey"
const PasteboardURLReadingFileURLsOnlyKey PasteboardReadingOptionKey = "NSPasteboardURLReadingFileURLsOnlyKey"

type PasteboardType string

const PasteboardTypeURL PasteboardType = "public.url"
const PasteboardTypeColor PasteboardType = "com.apple.cocoa.pasteboard.color"
const FileContentsPboardType PasteboardType = "NXFileContentsPboardType"
const PasteboardTypeFileURL PasteboardType = "public.file-url"
const FindPanelSearchOptionsPboardType PasteboardType = "NSFindPanel search options pasteboard type"
const PasteboardTypeFont PasteboardType = "com.apple.cocoa.pasteboard.character-formatting"
const PasteboardTypeHTML PasteboardType = "public.html"
const PasteboardTypeMultipleTextSelection PasteboardType = "com.apple.cocoa.pasteboard.multiple-text-selection"
const PasteboardTypePDF PasteboardType = "com.adobe.pdf"
const PasteboardTypePNG PasteboardType = "public.png"
const PasteboardTypeRTF PasteboardType = "public.rtf"
const PasteboardTypeRTFD PasteboardType = "com.apple.flat-rtfd"
const PasteboardTypeRuler PasteboardType = "com.apple.cocoa.pasteboard.paragraph-formatting"
const PasteboardTypeSound PasteboardType = "com.apple.cocoa.pasteboard.sound"
const PasteboardTypeString PasteboardType = "public.utf8-plain-text"
const PasteboardTypeTabularText PasteboardType = "public.utf8-tab-separated-values-text"
const PasteboardTypeTextFinderOptions PasteboardType = "com.apple.cocoa.pasteboard.find-panel-search-options"
const PasteboardTypeTIFF PasteboardType = "public.tiff"

type PasteboardWritingOptions uint

const PasteboardWritingPromised PasteboardWritingOptions = 512

type PathStyle int

const PathStyleStandard PathStyle = 0
const PathStylePopUp PathStyle = 2

type PointingDeviceType uint

const PointingDeviceTypeCursor PointingDeviceType = 2
const PointingDeviceTypeEraser PointingDeviceType = 3
const PointingDeviceTypePen PointingDeviceType = 1
const PointingDeviceTypeUnknown PointingDeviceType = 0

type PopoverBehavior int

const PopoverBehaviorApplicationDefined PopoverBehavior = 0
const PopoverBehaviorTransient PopoverBehavior = 1
const PopoverBehaviorSemitransient PopoverBehavior = 2

type PressureBehavior int

const PressureBehaviorUnknown PressureBehavior = -1
const PressureBehaviorPrimaryDefault PressureBehavior = 0
const PressureBehaviorPrimaryClick PressureBehavior = 1
const PressureBehaviorPrimaryGeneric PressureBehavior = 2
const PressureBehaviorPrimaryAccelerator PressureBehavior = 3
const PressureBehaviorPrimaryDeepClick PressureBehavior = 5
const PressureBehaviorPrimaryDeepDrag PressureBehavior = 6

type PrintInfoAttributeKey string

const PrintPaperName PrintInfoAttributeKey = "NSPaperName"
const PrintPaperSize PrintInfoAttributeKey = "NSPaperSize"
const PrintOrientation PrintInfoAttributeKey = "NSOrientation"
const PrintScalingFactor PrintInfoAttributeKey = "NSScalingFactor"
const PrintLeftMargin PrintInfoAttributeKey = "NSLeftMargin"
const PrintRightMargin PrintInfoAttributeKey = "NSRightMargin"
const PrintTopMargin PrintInfoAttributeKey = "NSTopMargin"
const PrintBottomMargin PrintInfoAttributeKey = "NSBottomMargin"
const PrintHorizontallyCentered PrintInfoAttributeKey = "NSHorizontallyCentered"
const PrintVerticallyCentered PrintInfoAttributeKey = "NSVerticallyCentered"
const PrintHorizontalPagination PrintInfoAttributeKey = "NSHorizonalPagination"
const PrintVerticalPagination PrintInfoAttributeKey = "NSVerticalPagination"
const PrintAllPages PrintInfoAttributeKey = "NSPrintAllPages"
const PrintCopies PrintInfoAttributeKey = "NSCopies"
const PrintDetailedErrorReporting PrintInfoAttributeKey = "NSDetailedErrorReporting"
const PrintFaxNumber PrintInfoAttributeKey = "NSFaxNumber"
const PrintFirstPage PrintInfoAttributeKey = "NSFirstPage"
const PrintHeaderAndFooter PrintInfoAttributeKey = "NSPrintHeaderAndFooter"
const PrintJobDisposition PrintInfoAttributeKey = "NSJobDisposition"
const PrintJobSavingFileNameExtensionHidden PrintInfoAttributeKey = "NSJobSavingFileNameExtensionHidden"
const PrintJobSavingURL PrintInfoAttributeKey = "NSJobSavingURL"
const PrintLastPage PrintInfoAttributeKey = "NSLastPage"
const PrintMustCollate PrintInfoAttributeKey = "NSMustCollate"
const PrintPagesAcross PrintInfoAttributeKey = "NSPagesAcross"
const PrintPagesDown PrintInfoAttributeKey = "NSPagesDown"
const PrintPrinter PrintInfoAttributeKey = "NSPrinter"
const PrintPrinterName PrintInfoAttributeKey = "NSPrinterName"
const PrintReversePageOrder PrintInfoAttributeKey = "NSReversePageOrder"
const PrintSelectionOnly PrintInfoAttributeKey = "NSPrintSelectionOnly"
const PrintTime PrintInfoAttributeKey = "NSPrintTime"

type PrintJobDispositionValue string

const PrintSpoolJob PrintJobDispositionValue = "NSPrintSpoolJob"
const PrintPreviewJob PrintJobDispositionValue = "NSPrintPreviewJob"
const PrintSaveJob PrintJobDispositionValue = "NSPrintSaveJob"
const PrintCancelJob PrintJobDispositionValue = "NSPrintCancelJob"

type PrintPanelJobStyleHint string

const PrintPhotoJobStyleHint PrintPanelJobStyleHint = "Photo"
const PrintAllPresetsJobStyleHint PrintPanelJobStyleHint = "All"
const PrintNoPresetsJobStyleHint PrintPanelJobStyleHint = "None"

type PrintPanelOptions uint

const PrintPanelShowsCopies PrintPanelOptions = 1
const PrintPanelShowsPageRange PrintPanelOptions = 2
const PrintPanelShowsPaperSize PrintPanelOptions = 4
const PrintPanelShowsOrientation PrintPanelOptions = 8
const PrintPanelShowsScaling PrintPanelOptions = 16
const PrintPanelShowsPrintSelection PrintPanelOptions = 32
const PrintPanelShowsPageSetupAccessory PrintPanelOptions = 256
const PrintPanelShowsPreview PrintPanelOptions = 131072

type PrintRenderingQuality int

const PrintRenderingQualityBest PrintRenderingQuality = 0
const PrintRenderingQualityResponsive PrintRenderingQuality = 1

type PrinterPaperName string

type PrinterTableStatus uint

const PrinterTableOK PrinterTableStatus = 0
const PrinterTableNotFound PrinterTableStatus = 1
const PrinterTableError PrinterTableStatus = 2

type PrinterTypeName string

type PrintingPageOrder int

const AscendingPageOrder PrintingPageOrder = 1
const DescendingPageOrder PrintingPageOrder = -1
const SpecialPageOrder PrintingPageOrder = 0
const UnknownPageOrder PrintingPageOrder = 2

type PrintingPaginationMode uint

const PrintingPaginationModeAutomatic PrintingPaginationMode = 0
const PrintingPaginationModeFit PrintingPaginationMode = 1
const PrintingPaginationModeClip PrintingPaginationMode = 2

type ProgressIndicatorStyle uint

type RectAlignment int

const RectAlignmentNone RectAlignment = 0
const RectAlignmentTop RectAlignment = 1
const RectAlignmentTopLeading RectAlignment = 2
const RectAlignmentLeading RectAlignment = 3
const RectAlignmentBottomLeading RectAlignment = 4
const RectAlignmentBottom RectAlignment = 5
const RectAlignmentBottomTrailing RectAlignment = 6
const RectAlignmentTrailing RectAlignment = 7
const RectAlignmentTopTrailing RectAlignment = 8

type RemoteNotificationType uint

const RemoteNotificationTypeNone RemoteNotificationType = 0
const RemoteNotificationTypeBadge RemoteNotificationType = 1
const RemoteNotificationTypeSound RemoteNotificationType = 2
const RemoteNotificationTypeAlert RemoteNotificationType = 4

type RequestUserAttentionType uint

const CriticalRequest RequestUserAttentionType = 0
const InformationalRequest RequestUserAttentionType = 10

type RuleEditorNestingMode uint

const RuleEditorNestingModeCompound RuleEditorNestingMode = 2
const RuleEditorNestingModeList RuleEditorNestingMode = 1
const RuleEditorNestingModeSimple RuleEditorNestingMode = 3
const RuleEditorNestingModeSingle RuleEditorNestingMode = 0

type RuleEditorPredicatePartKey string

const RuleEditorPredicateComparisonModifier RuleEditorPredicatePartKey = "NSRuleEditorPredicateComparisonModifier"
const RuleEditorPredicateCompoundType RuleEditorPredicatePartKey = "NSRuleEditorPredicateCompoundType"
const RuleEditorPredicateCustomSelector RuleEditorPredicatePartKey = "NSRuleEditorPredicateCustomSelector"
const RuleEditorPredicateLeftExpression RuleEditorPredicatePartKey = "NSRuleEditorPredicateLeftExpression"
const RuleEditorPredicateOperatorType RuleEditorPredicatePartKey = "NSRuleEditorPredicateOperatorType"
const RuleEditorPredicateOptions RuleEditorPredicatePartKey = "NSRuleEditorPredicateOptions"
const RuleEditorPredicateRightExpression RuleEditorPredicatePartKey = "NSRuleEditorPredicateRightExpression"

type RuleEditorRowType uint

const RuleEditorRowTypeCompound RuleEditorRowType = 1
const RuleEditorRowTypeSimple RuleEditorRowType = 0

type RulerOrientation uint

const HorizontalRuler RulerOrientation = 0
const VerticalRuler RulerOrientation = 1

type RulerViewUnitName string

const RulerViewUnitCentimeters RulerViewUnitName = "Centimeters"
const RulerViewUnitInches RulerViewUnitName = "Inches"
const RulerViewUnitPicas RulerViewUnitName = "Picas"
const RulerViewUnitPoints RulerViewUnitName = "Points"

type SaveOperationType uint

const SaveOperation SaveOperationType = 0
const SaveAsOperation SaveOperationType = 1
const SaveToOperation SaveOperationType = 2
const AutosaveElsewhereOperation SaveOperationType = 3
const AutosaveInPlaceOperation SaveOperationType = 4
const AutosaveAsOperation SaveOperationType = 5

type ScrollArrowPosition uint

type ScrollElasticity int

const ScrollElasticityAutomatic ScrollElasticity = 0
const ScrollElasticityNone ScrollElasticity = 1
const ScrollElasticityAllowed ScrollElasticity = 2

type ScrollViewFindBarPosition int

const ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
const ScrollViewFindBarPositionAboveContent ScrollViewFindBarPosition = 1
const ScrollViewFindBarPositionBelowContent ScrollViewFindBarPosition = 2

type ScrollerArrow uint

type ScrollerKnobStyle int

const ScrollerKnobStyleDefault ScrollerKnobStyle = 0
const ScrollerKnobStyleDark ScrollerKnobStyle = 1
const ScrollerKnobStyleLight ScrollerKnobStyle = 2

type ScrollerPart uint

const ScrollerKnob ScrollerPart = 2
const ScrollerKnobSlot ScrollerPart = 6
const ScrollerDecrementPage ScrollerPart = 1
const ScrollerIncrementPage ScrollerPart = 3
const ScrollerNoPart ScrollerPart = 0

type ScrollerStyle int

const ScrollerStyleLegacy ScrollerStyle = 0
const ScrollerStyleOverlay ScrollerStyle = 1

type SearchFieldRecentsAutosaveName string

type SegmentDistribution int

const SegmentDistributionFit SegmentDistribution = 0
const SegmentDistributionFill SegmentDistribution = 1
const SegmentDistributionFillEqually SegmentDistribution = 2
const SegmentDistributionFillProportionally SegmentDistribution = 3

type SegmentStyle int

const SegmentStyleAutomatic SegmentStyle = 0
const SegmentStyleRounded SegmentStyle = 1
const SegmentStyleTexturedRounded SegmentStyle = 2
const SegmentStyleRoundRect SegmentStyle = 3
const SegmentStyleTexturedSquare SegmentStyle = 4
const SegmentStyleCapsule SegmentStyle = 5
const SegmentStyleSmallSquare SegmentStyle = 6
const SegmentStyleSeparated SegmentStyle = 8

type SegmentSwitchTracking uint

const SegmentSwitchTrackingSelectOne SegmentSwitchTracking = 0
const SegmentSwitchTrackingSelectAny SegmentSwitchTracking = 1
const SegmentSwitchTrackingMomentary SegmentSwitchTracking = 2
const SegmentSwitchTrackingMomentaryAccelerator SegmentSwitchTracking = 3

type SelectionAffinity uint

const SelectionAffinityUpstream SelectionAffinity = 0
const SelectionAffinityDownstream SelectionAffinity = 1

type SelectionDirection uint

const DirectSelection SelectionDirection = 0
const SelectingNext SelectionDirection = 1
const SelectingPrevious SelectionDirection = 2

type SelectionGranularity uint

const SelectByCharacter SelectionGranularity = 0
const SelectByWord SelectionGranularity = 1
const SelectByParagraph SelectionGranularity = 2

type SharingContentScope int

const SharingContentScopeItem SharingContentScope = 0
const SharingContentScopePartial SharingContentScope = 1
const SharingContentScopeFull SharingContentScope = 2

type SharingServiceName string

const SharingServiceNameAddToAperture SharingServiceName = "com.apple.share.System.add-to-aperture"
const SharingServiceNameAddToIPhoto SharingServiceName = "com.apple.share.System.add-to-iphoto"
const SharingServiceNameAddToSafariReadingList SharingServiceName = "com.apple.share.System.add-to-safari-reading-list"
const SharingServiceNameCloudSharing SharingServiceName = "com.apple.share.CloudSharing"
const SharingServiceNameComposeEmail SharingServiceName = "com.apple.share.Mail.compose"
const SharingServiceNameComposeMessage SharingServiceName = "com.apple.messages.ShareExtension"
const SharingServiceNameSendViaAirDrop SharingServiceName = "com.apple.share.AirDrop.send"
const SharingServiceNameUseAsDesktopPicture SharingServiceName = "com.apple.share.System.set-desktop-image"

type SliderType uint

const SliderTypeCircular SliderType = 1
const SliderTypeLinear SliderType = 0

type SoundName string

type SoundPlaybackDeviceIdentifier string

type SplitViewAutosaveName string

type SplitViewDividerStyle int

const SplitViewDividerStyleThick SplitViewDividerStyle = 1
const SplitViewDividerStyleThin SplitViewDividerStyle = 2
const SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3

type SpringLoadingHighlight int

const SpringLoadingHighlightNone SpringLoadingHighlight = 0
const SpringLoadingHighlightStandard SpringLoadingHighlight = 1
const SpringLoadingHighlightEmphasized SpringLoadingHighlight = 2

type StackViewDistribution int

const StackViewDistributionEqualCentering StackViewDistribution = 4
const StackViewDistributionEqualSpacing StackViewDistribution = 3
const StackViewDistributionFill StackViewDistribution = 0
const StackViewDistributionFillEqually StackViewDistribution = 1
const StackViewDistributionFillProportionally StackViewDistribution = 2
const StackViewDistributionGravityAreas StackViewDistribution = -1

type StackViewGravity int

const StackViewGravityTop StackViewGravity = 1
const StackViewGravityLeading StackViewGravity = 1
const StackViewGravityCenter StackViewGravity = 2
const StackViewGravityBottom StackViewGravity = 3
const StackViewGravityTrailing StackViewGravity = 3

type StackViewVisibilityPriority float32

const StackViewVisibilityPriorityMustHold StackViewVisibilityPriority = 1000
const StackViewVisibilityPriorityDetachOnlyIfNecessary StackViewVisibilityPriority = 900
const StackViewVisibilityPriorityNotVisible StackViewVisibilityPriority = 0

type StatusItemAutosaveName string

type StatusItemBehavior uint

const StatusItemBehaviorRemovalAllowed StatusItemBehavior = 2
const StatusItemBehaviorTerminationOnRemoval StatusItemBehavior = 4

type StoryboardName string

type StoryboardSceneIdentifier string

type TIFFCompression uint

const TIFFCompressionNone TIFFCompression = 1
const TIFFCompressionCCITTFAX3 TIFFCompression = 3
const TIFFCompressionCCITTFAX4 TIFFCompression = 4
const TIFFCompressionLZW TIFFCompression = 5
const TIFFCompressionJPEG TIFFCompression = 6
const TIFFCompressionNEXT TIFFCompression = 32766
const TIFFCompressionPackBits TIFFCompression = 32773
const TIFFCompressionOldJPEG TIFFCompression = 32865

type TabPosition uint

const TabPositionBottom TabPosition = 3
const TabPositionLeft TabPosition = 2
const TabPositionNone TabPosition = 0
const TabPositionRight TabPosition = 4
const TabPositionTop TabPosition = 1

type TabState uint

const BackgroundTab TabState = 1
const PressedTab TabState = 2
const SelectedTab TabState = 0

type TabViewBorderType uint

const TabViewBorderTypeBezel TabViewBorderType = 2
const TabViewBorderTypeLine TabViewBorderType = 1
const TabViewBorderTypeNone TabViewBorderType = 0

type TabViewType uint

const TopTabsBezelBorder TabViewType = 0
const NoTabsBezelBorder TabViewType = 4
const NoTabsLineBorder TabViewType = 5
const NoTabsNoBorder TabViewType = 6
const BottomTabsBezelBorder TabViewType = 2
const LeftTabsBezelBorder TabViewType = 1
const RightTabsBezelBorder TabViewType = 3

type TableColumnResizingOptions uint

const TableColumnAutoresizingMask TableColumnResizingOptions = 1
const TableColumnNoResizing TableColumnResizingOptions = 0
const TableColumnUserResizingMask TableColumnResizingOptions = 2

type TableRowActionEdge int

const TableRowActionEdgeLeading TableRowActionEdge = 0
const TableRowActionEdgeTrailing TableRowActionEdge = 1

type TableViewAnimationOptions uint

const TableViewAnimationEffectNone TableViewAnimationOptions = 0
const TableViewAnimationEffectFade TableViewAnimationOptions = 1
const TableViewAnimationEffectGap TableViewAnimationOptions = 2
const TableViewAnimationSlideUp TableViewAnimationOptions = 16
const TableViewAnimationSlideDown TableViewAnimationOptions = 32
const TableViewAnimationSlideLeft TableViewAnimationOptions = 48
const TableViewAnimationSlideRight TableViewAnimationOptions = 64

type TableViewAutosaveName string

type TableViewColumnAutoresizingStyle uint

const TableViewNoColumnAutoresizing TableViewColumnAutoresizingStyle = 0
const TableViewUniformColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 1
const TableViewSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 2
const TableViewReverseSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 3
const TableViewLastColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 4
const TableViewFirstColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 5

type TableViewDraggingDestinationFeedbackStyle int

const TableViewDraggingDestinationFeedbackStyleNone TableViewDraggingDestinationFeedbackStyle = -1
const TableViewDraggingDestinationFeedbackStyleRegular TableViewDraggingDestinationFeedbackStyle = 0
const TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 1
const TableViewDraggingDestinationFeedbackStyleGap TableViewDraggingDestinationFeedbackStyle = 2

type TableViewDropOperation uint

const TableViewDropOn TableViewDropOperation = 0
const TableViewDropAbove TableViewDropOperation = 1

type TableViewGridLineStyle uint

const TableViewGridNone TableViewGridLineStyle = 0
const TableViewSolidVerticalGridLineMask TableViewGridLineStyle = 1
const TableViewSolidHorizontalGridLineMask TableViewGridLineStyle = 2
const TableViewDashedHorizontalGridLineMask TableViewGridLineStyle = 8

type TableViewRowActionStyle int

const TableViewRowActionStyleRegular TableViewRowActionStyle = 0
const TableViewRowActionStyleDestructive TableViewRowActionStyle = 1

type TableViewRowSizeStyle int

const TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
const TableViewRowSizeStyleCustom TableViewRowSizeStyle = 0
const TableViewRowSizeStyleSmall TableViewRowSizeStyle = 1
const TableViewRowSizeStyleMedium TableViewRowSizeStyle = 2
const TableViewRowSizeStyleLarge TableViewRowSizeStyle = 3

type TableViewSelectionHighlightStyle int

const TableViewSelectionHighlightStyleNone TableViewSelectionHighlightStyle = -1
const TableViewSelectionHighlightStyleRegular TableViewSelectionHighlightStyle = 0

type TableViewStyle int

const TableViewStyleAutomatic TableViewStyle = 0
const TableViewStyleFullWidth TableViewStyle = 1
const TableViewStyleInset TableViewStyle = 2
const TableViewStyleSourceList TableViewStyle = 3
const TableViewStylePlain TableViewStyle = 4

type TextAlignment int

const TextAlignmentLeft TextAlignment = 0
const TextAlignmentJustified TextAlignment = 3
const TextAlignmentNatural TextAlignment = 4

type TextBlockDimension uint

const TextBlockWidth TextBlockDimension = 0
const TextBlockMinimumWidth TextBlockDimension = 1
const TextBlockMaximumWidth TextBlockDimension = 2
const TextBlockHeight TextBlockDimension = 4
const TextBlockMinimumHeight TextBlockDimension = 5
const TextBlockMaximumHeight TextBlockDimension = 6

type TextBlockLayer int

const TextBlockPadding TextBlockLayer = -1
const TextBlockBorder TextBlockLayer = 0
const TextBlockMargin TextBlockLayer = 1

type TextBlockValueType uint

const TextBlockAbsoluteValueType TextBlockValueType = 0
const TextBlockPercentageValueType TextBlockValueType = 1

type TextBlockVerticalAlignment uint

const TextBlockTopAlignment TextBlockVerticalAlignment = 0
const TextBlockMiddleAlignment TextBlockVerticalAlignment = 1
const TextBlockBottomAlignment TextBlockVerticalAlignment = 2
const TextBlockBaselineAlignment TextBlockVerticalAlignment = 3

type TextCheckingOptionKey string

const TextCheckingDocumentAuthorKey TextCheckingOptionKey = "DocumentAuthor"
const TextCheckingDocumentTitleKey TextCheckingOptionKey = "DocumentTitle"
const TextCheckingDocumentURLKey TextCheckingOptionKey = "DocumentURL"
const TextCheckingOrthographyKey TextCheckingOptionKey = "Orthography"
const TextCheckingQuotesKey TextCheckingOptionKey = "Quotes"
const TextCheckingReferenceDateKey TextCheckingOptionKey = "ReferenceDate"
const TextCheckingReferenceTimeZoneKey TextCheckingOptionKey = "ReferenceTimeZone"
const TextCheckingRegularExpressionsKey TextCheckingOptionKey = "RegularExpressions"
const TextCheckingReplacementsKey TextCheckingOptionKey = "Replacements"
const TextCheckingSelectedRangeKey TextCheckingOptionKey = "SelectedRange"

type TextFieldBezelStyle uint

const TextFieldSquareBezel TextFieldBezelStyle = 0
const TextFieldRoundedBezel TextFieldBezelStyle = 1

type TextInputSourceIdentifier string

type TextLayoutFragmentEnumerationOptions uint

const TextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment TextLayoutFragmentEnumerationOptions = 8
const TextLayoutFragmentEnumerationOptionsEnsuresLayout TextLayoutFragmentEnumerationOptions = 4
const TextLayoutFragmentEnumerationOptionsEstimatesSize TextLayoutFragmentEnumerationOptions = 2
const TextLayoutFragmentEnumerationOptionsReverse TextLayoutFragmentEnumerationOptions = 1
const TextLayoutFragmentEnumerationOptionsNone TextLayoutFragmentEnumerationOptions = 0

type TextLayoutFragmentState uint

const TextLayoutFragmentStateCalculatedUsageBounds TextLayoutFragmentState = 2
const TextLayoutFragmentStateEstimatedUsageBounds TextLayoutFragmentState = 1
const TextLayoutFragmentStateLayoutAvailable TextLayoutFragmentState = 3
const TextLayoutFragmentStateNone TextLayoutFragmentState = 0

type TextLayoutManagerSegmentOptions uint

const TextLayoutManagerSegmentOptionsHeadSegmentExtended TextLayoutManagerSegmentOptions = 4
const TextLayoutManagerSegmentOptionsMiddleFragmentsExcluded TextLayoutManagerSegmentOptions = 2
const TextLayoutManagerSegmentOptionsRangeNotRequired TextLayoutManagerSegmentOptions = 1
const TextLayoutManagerSegmentOptionsTailSegmentExtended TextLayoutManagerSegmentOptions = 8
const TextLayoutManagerSegmentOptionsUpstreamAffinity TextLayoutManagerSegmentOptions = 16
const TextLayoutManagerSegmentOptionsNone TextLayoutManagerSegmentOptions = 0

type TextLayoutManagerSegmentType int

const TextLayoutManagerSegmentTypeHighlight TextLayoutManagerSegmentType = 2
const TextLayoutManagerSegmentTypeSelection TextLayoutManagerSegmentType = 1
const TextLayoutManagerSegmentTypeStandard TextLayoutManagerSegmentType = 0

type TextLayoutOrientation int

const TextLayoutOrientationHorizontal TextLayoutOrientation = 0
const TextLayoutOrientationVertical TextLayoutOrientation = 1

type TextListMarkerFormat string

const TextListMarkerBox TextListMarkerFormat = "{box}"
const TextListMarkerCheck TextListMarkerFormat = "{check}"
const TextListMarkerCircle TextListMarkerFormat = "{circle}"
const TextListMarkerDecimal TextListMarkerFormat = "{decimal}"
const TextListMarkerDiamond TextListMarkerFormat = "{diamond}"
const TextListMarkerDisc TextListMarkerFormat = "{disc}"
const TextListMarkerHyphen TextListMarkerFormat = "{hyphen}"
const TextListMarkerLowercaseAlpha TextListMarkerFormat = "{lower-alpha}"
const TextListMarkerLowercaseHexadecimal TextListMarkerFormat = "{lower-hexadecimal}"
const TextListMarkerLowercaseLatin TextListMarkerFormat = "{lower-latin}"
const TextListMarkerLowercaseRoman TextListMarkerFormat = "{lower-roman}"
const TextListMarkerOctal TextListMarkerFormat = "{octal}"
const TextListMarkerSquare TextListMarkerFormat = "{square}"
const TextListMarkerUppercaseAlpha TextListMarkerFormat = "{upper-alpha}"
const TextListMarkerUppercaseHexadecimal TextListMarkerFormat = "{upper-hexadecimal}"
const TextListMarkerUppercaseLatin TextListMarkerFormat = "{upper-latin}"
const TextListMarkerUppercaseRoman TextListMarkerFormat = "{upper-roman}"

type TextListOptions uint

const TextListPrependEnclosingMarker TextListOptions = 1

type TextSelectionAffinity int

const TextSelectionAffinityDownstream TextSelectionAffinity = 1
const TextSelectionAffinityUpstream TextSelectionAffinity = 0

type TextSelectionGranularity int

const TextSelectionGranularityCharacter TextSelectionGranularity = 0
const TextSelectionGranularityWord TextSelectionGranularity = 1
const TextSelectionGranularityParagraph TextSelectionGranularity = 2
const TextSelectionGranularityLine TextSelectionGranularity = 3
const TextSelectionGranularitySentence TextSelectionGranularity = 4

type TextSelectionNavigationDestination int

const TextSelectionNavigationDestinationCharacter TextSelectionNavigationDestination = 0
const TextSelectionNavigationDestinationWord TextSelectionNavigationDestination = 1
const TextSelectionNavigationDestinationLine TextSelectionNavigationDestination = 2
const TextSelectionNavigationDestinationSentence TextSelectionNavigationDestination = 3
const TextSelectionNavigationDestinationParagraph TextSelectionNavigationDestination = 4
const TextSelectionNavigationDestinationContainer TextSelectionNavigationDestination = 5
const TextSelectionNavigationDestinationDocument TextSelectionNavigationDestination = 6

type TextSelectionNavigationDirection int

const TextSelectionNavigationDirectionForward TextSelectionNavigationDirection = 0
const TextSelectionNavigationDirectionBackward TextSelectionNavigationDirection = 1
const TextSelectionNavigationDirectionLeft TextSelectionNavigationDirection = 3
const TextSelectionNavigationDirectionRight TextSelectionNavigationDirection = 2
const TextSelectionNavigationDirectionUp TextSelectionNavigationDirection = 4
const TextSelectionNavigationDirectionDown TextSelectionNavigationDirection = 5

type TextSelectionNavigationModifier uint

const TextSelectionNavigationModifierExtend TextSelectionNavigationModifier = 1
const TextSelectionNavigationModifierMultiple TextSelectionNavigationModifier = 4
const TextSelectionNavigationModifierVisual TextSelectionNavigationModifier = 2

type TextSelectionNavigationWritingDirection int

const TextSelectionNavigationWritingDirectionLeftToRight TextSelectionNavigationWritingDirection = 0
const TextSelectionNavigationWritingDirectionRightToLeft TextSelectionNavigationWritingDirection = 1

type TextStorageEditActions uint

const TextStorageEditedAttributes TextStorageEditActions = 1
const TextStorageEditedCharacters TextStorageEditActions = 2

type TextStorageEditedOptions uint

type TextTabOptionKey string

const TabColumnTerminatorsAttributeName TextTabOptionKey = "NSTabColumnTerminatorsAttributeName"

type TextTabType uint

const LeftTabStopType TextTabType = 0
const RightTabStopType TextTabType = 1
const CenterTabStopType TextTabType = 2
const DecimalTabStopType TextTabType = 3

type TextTableLayoutAlgorithm uint

const TextTableAutomaticLayoutAlgorithm TextTableLayoutAlgorithm = 0
const TextTableFixedLayoutAlgorithm TextTableLayoutAlgorithm = 1

type TickMarkPosition uint

const TickMarkPositionBelow TickMarkPosition = 0
const TickMarkPositionAbove TickMarkPosition = 1
const TickMarkPositionLeading TickMarkPosition = 1
const TickMarkPositionTrailing TickMarkPosition = 0

type TitlePosition uint

const NoTitle TitlePosition = 0
const AboveTop TitlePosition = 1
const AtTop TitlePosition = 2
const BelowTop TitlePosition = 3
const AboveBottom TitlePosition = 4
const AtBottom TitlePosition = 5
const BelowBottom TitlePosition = 6

type TitlebarSeparatorStyle int

const TitlebarSeparatorStyleAutomatic TitlebarSeparatorStyle = 0
const TitlebarSeparatorStyleLine TitlebarSeparatorStyle = 2
const TitlebarSeparatorStyleNone TitlebarSeparatorStyle = 1
const TitlebarSeparatorStyleShadow TitlebarSeparatorStyle = 3

type TokenStyle uint

const TokenStyleDefault TokenStyle = 0
const TokenStyleNone TokenStyle = 1
const TokenStylePlainSquared TokenStyle = 4
const TokenStyleRounded TokenStyle = 2
const TokenStyleSquared TokenStyle = 3

type ToolTipTag int

type ToolbarDisplayMode uint

const ToolbarDisplayModeDefault ToolbarDisplayMode = 0
const ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
const ToolbarDisplayModeIconOnly ToolbarDisplayMode = 2
const ToolbarDisplayModeLabelOnly ToolbarDisplayMode = 3

type ToolbarIdentifier string

type ToolbarItemGroupControlRepresentation int

const ToolbarItemGroupControlRepresentationAutomatic ToolbarItemGroupControlRepresentation = 0
const ToolbarItemGroupControlRepresentationExpanded ToolbarItemGroupControlRepresentation = 1
const ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 2

type ToolbarItemGroupSelectionMode int

const ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 2
const ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 1
const ToolbarItemGroupSelectionModeSelectOne ToolbarItemGroupSelectionMode = 0

type ToolbarItemIdentifier string

const ToolbarSpaceItemIdentifier ToolbarItemIdentifier = "NSToolbarSpaceItem"
const ToolbarFlexibleSpaceItemIdentifier ToolbarItemIdentifier = "NSToolbarFlexibleSpaceItem"
const ToolbarCloudSharingItemIdentifier ToolbarItemIdentifier = "NSToolbarCloudSharingItem"
const ToolbarPrintItemIdentifier ToolbarItemIdentifier = "NSToolbarPrintItem"
const ToolbarShowColorsItemIdentifier ToolbarItemIdentifier = "NSToolbarShowColorsItem"
const ToolbarShowFontsItemIdentifier ToolbarItemIdentifier = "NSToolbarShowFontsItem"
const ToolbarToggleSidebarItemIdentifier ToolbarItemIdentifier = "NSToolbarToggleSidebarItem"
const ToolbarSidebarTrackingSeparatorItemIdentifier ToolbarItemIdentifier = "NSToolbarSidebarTrackingSeparatorItemIdentifier"

type ToolbarItemVisibilityPriority int

const ToolbarItemVisibilityPriorityStandard ToolbarItemVisibilityPriority = 0
const ToolbarItemVisibilityPriorityLow ToolbarItemVisibilityPriority = -1000
const ToolbarItemVisibilityPriorityHigh ToolbarItemVisibilityPriority = 1000
const ToolbarItemVisibilityPriorityUser ToolbarItemVisibilityPriority = 2000

type ToolbarSizeMode uint

type TouchBarCustomizationIdentifier string

type TouchBarItemIdentifier string

const TouchBarItemIdentifierFixedSpaceSmall TouchBarItemIdentifier = "NSTouchBarItemIdentifierFixedSpaceSmall"
const TouchBarItemIdentifierFixedSpaceLarge TouchBarItemIdentifier = "NSTouchBarItemIdentifierFixedSpaceLarge"
const TouchBarItemIdentifierFlexibleSpace TouchBarItemIdentifier = "NSTouchBarItemIdentifierFlexibleSpace"
const TouchBarItemIdentifierOtherItemsProxy TouchBarItemIdentifier = "NSTouchBarItemIdentifierOtherItemsProxy"
const TouchBarItemIdentifierCandidateList TouchBarItemIdentifier = "NSTouchBarItemIdentifierCandidateList"
const TouchBarItemIdentifierCharacterPicker TouchBarItemIdentifier = "NSTouchBarItemIdentifierCharacterPicker"
const TouchBarItemIdentifierTextFormat TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextFormat"
const TouchBarItemIdentifierTextAlignment TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextAlignment"
const TouchBarItemIdentifierTextColorPicker TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextColorPicker"
const TouchBarItemIdentifierTextList TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextList"
const TouchBarItemIdentifierTextStyle TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextStyle"

type TouchBarItemPriority float32

const TouchBarItemPriorityLow TouchBarItemPriority = -1000.0
const TouchBarItemPriorityNormal TouchBarItemPriority = 0.0
const TouchBarItemPriorityHigh TouchBarItemPriority = 1000.0

type TouchPhase uint

const TouchPhaseBegan TouchPhase = 1
const TouchPhaseMoved TouchPhase = 2
const TouchPhaseStationary TouchPhase = 4
const TouchPhaseEnded TouchPhase = 8
const TouchPhaseCancelled TouchPhase = 16
const TouchPhaseTouching TouchPhase = 7
const TouchPhaseAny TouchPhase = 18446744073709551615

type TouchType int

const TouchTypeDirect TouchType = 0
const TouchTypeIndirect TouchType = 1

type TouchTypeMask uint

const TouchTypeMaskDirect TouchTypeMask = 1
const TouchTypeMaskIndirect TouchTypeMask = 2

type TrackingAreaOptions uint

const TrackingMouseEnteredAndExited TrackingAreaOptions = 1
const TrackingMouseMoved TrackingAreaOptions = 2
const TrackingCursorUpdate TrackingAreaOptions = 4
const TrackingActiveWhenFirstResponder TrackingAreaOptions = 16
const TrackingActiveInKeyWindow TrackingAreaOptions = 32
const TrackingActiveInActiveApp TrackingAreaOptions = 64
const TrackingActiveAlways TrackingAreaOptions = 128
const TrackingAssumeInside TrackingAreaOptions = 256
const TrackingInVisibleRect TrackingAreaOptions = 512
const TrackingEnabledDuringMouseDrag TrackingAreaOptions = 1024

type TrackingRectTag int

type TypesetterBehavior int

const TypesetterLatestBehavior TypesetterBehavior = -1
const TypesetterOriginalBehavior TypesetterBehavior = 0
const TypesetterBehavior_10_2_WithCompatibility TypesetterBehavior = 1
const TypesetterBehavior_10_2 TypesetterBehavior = 2
const TypesetterBehavior_10_3 TypesetterBehavior = 3
const TypesetterBehavior_10_4 TypesetterBehavior = 4

type TypesetterControlCharacterAction uint

const TypesetterZeroAdvancementAction TypesetterControlCharacterAction = 1
const TypesetterWhitespaceAction TypesetterControlCharacterAction = 2
const TypesetterHorizontalTabAction TypesetterControlCharacterAction = 4
const TypesetterLineBreakAction TypesetterControlCharacterAction = 8
const TypesetterParagraphBreakAction TypesetterControlCharacterAction = 16
const TypesetterContainerBreakAction TypesetterControlCharacterAction = 32

type UnderlineStyle int

const UnderlineStyleNone UnderlineStyle = 0
const UnderlineStyleSingle UnderlineStyle = 1
const UnderlineStyleThick UnderlineStyle = 2
const UnderlineStyleDouble UnderlineStyle = 9
const UnderlineStylePatternSolid UnderlineStyle = 0
const UnderlineStylePatternDot UnderlineStyle = 256
const UnderlineStylePatternDash UnderlineStyle = 512
const UnderlineStylePatternDashDot UnderlineStyle = 768
const UnderlineStylePatternDashDotDot UnderlineStyle = 1024
const UnderlineStyleByWord UnderlineStyle = 32768

type UsableScrollerParts uint

const NoScrollerParts UsableScrollerParts = 0
const AllScrollerParts UsableScrollerParts = 2

type UserInterfaceItemIdentifier string

type UserInterfaceLayoutDirection int

const UserInterfaceLayoutDirectionLeftToRight UserInterfaceLayoutDirection = 0
const UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1

type UserInterfaceLayoutOrientation int

const UserInterfaceLayoutOrientationHorizontal UserInterfaceLayoutOrientation = 0
const UserInterfaceLayoutOrientationVertical UserInterfaceLayoutOrientation = 1

type ViewAnimationKey string

const ViewAnimationEffectKey ViewAnimationKey = "NSViewAnimationEffectKey"
const ViewAnimationEndFrameKey ViewAnimationKey = "NSViewAnimationEndFrameKey"
const ViewAnimationStartFrameKey ViewAnimationKey = "NSViewAnimationStartFrameKey"
const ViewAnimationTargetKey ViewAnimationKey = "NSViewAnimationTargetKey"

type ViewControllerTransitionOptions uint

const ViewControllerTransitionNone ViewControllerTransitionOptions = 0
const ViewControllerTransitionCrossfade ViewControllerTransitionOptions = 1
const ViewControllerTransitionSlideUp ViewControllerTransitionOptions = 16
const ViewControllerTransitionSlideDown ViewControllerTransitionOptions = 32
const ViewControllerTransitionSlideLeft ViewControllerTransitionOptions = 64
const ViewControllerTransitionSlideRight ViewControllerTransitionOptions = 128
const ViewControllerTransitionSlideForward ViewControllerTransitionOptions = 320
const ViewControllerTransitionSlideBackward ViewControllerTransitionOptions = 384
const ViewControllerTransitionAllowUserInteraction ViewControllerTransitionOptions = 4096

type ViewFullScreenModeOptionKey string

const FullScreenModeAllScreens ViewFullScreenModeOptionKey = "NSFullScreenModeAllScreens"
const FullScreenModeApplicationPresentationOptions ViewFullScreenModeOptionKey = "NSFullScreenModeApplicationPresentationOptions"
const FullScreenModeSetting ViewFullScreenModeOptionKey = "NSFullScreenModeSetting"
const FullScreenModeWindowLevel ViewFullScreenModeOptionKey = "NSFullScreenModeWindowLevel"

type ViewLayerContentsPlacement int

const ViewLayerContentsPlacementScaleAxesIndependently ViewLayerContentsPlacement = 0
const ViewLayerContentsPlacementScaleProportionallyToFit ViewLayerContentsPlacement = 1
const ViewLayerContentsPlacementScaleProportionallyToFill ViewLayerContentsPlacement = 2
const ViewLayerContentsPlacementCenter ViewLayerContentsPlacement = 3
const ViewLayerContentsPlacementTop ViewLayerContentsPlacement = 4
const ViewLayerContentsPlacementTopRight ViewLayerContentsPlacement = 5
const ViewLayerContentsPlacementRight ViewLayerContentsPlacement = 6
const ViewLayerContentsPlacementBottomRight ViewLayerContentsPlacement = 7
const ViewLayerContentsPlacementBottom ViewLayerContentsPlacement = 8
const ViewLayerContentsPlacementBottomLeft ViewLayerContentsPlacement = 9
const ViewLayerContentsPlacementLeft ViewLayerContentsPlacement = 10
const ViewLayerContentsPlacementTopLeft ViewLayerContentsPlacement = 11

type ViewLayerContentsRedrawPolicy int

const ViewLayerContentsRedrawNever ViewLayerContentsRedrawPolicy = 0
const ViewLayerContentsRedrawOnSetNeedsDisplay ViewLayerContentsRedrawPolicy = 1
const ViewLayerContentsRedrawDuringViewResize ViewLayerContentsRedrawPolicy = 2
const ViewLayerContentsRedrawBeforeViewResize ViewLayerContentsRedrawPolicy = 3
const ViewLayerContentsRedrawCrossfade ViewLayerContentsRedrawPolicy = 4

type VisualEffectBlendingMode int

const VisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = 0
const VisualEffectBlendingModeWithinWindow VisualEffectBlendingMode = 1

type VisualEffectMaterial int

const VisualEffectMaterialTitlebar VisualEffectMaterial = 3
const VisualEffectMaterialSelection VisualEffectMaterial = 4
const VisualEffectMaterialMenu VisualEffectMaterial = 5
const VisualEffectMaterialPopover VisualEffectMaterial = 6
const VisualEffectMaterialSidebar VisualEffectMaterial = 7
const VisualEffectMaterialHeaderView VisualEffectMaterial = 10
const VisualEffectMaterialSheet VisualEffectMaterial = 11
const VisualEffectMaterialWindowBackground VisualEffectMaterial = 12
const VisualEffectMaterialHUDWindow VisualEffectMaterial = 13
const VisualEffectMaterialFullScreenUI VisualEffectMaterial = 15
const VisualEffectMaterialToolTip VisualEffectMaterial = 17
const VisualEffectMaterialContentBackground VisualEffectMaterial = 18
const VisualEffectMaterialUnderWindowBackground VisualEffectMaterial = 21
const VisualEffectMaterialUnderPageBackground VisualEffectMaterial = 22

type VisualEffectState int

const VisualEffectStateFollowsWindowActiveState VisualEffectState = 0
const VisualEffectStateActive VisualEffectState = 1
const VisualEffectStateInactive VisualEffectState = 2

type WindingRule uint

const WindingRuleNonZero WindingRule = 0
const WindingRuleEvenOdd WindingRule = 1

type WindowAnimationBehavior int

const WindowAnimationBehaviorDefault WindowAnimationBehavior = 0
const WindowAnimationBehaviorNone WindowAnimationBehavior = 2
const WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 3
const WindowAnimationBehaviorUtilityWindow WindowAnimationBehavior = 4
const WindowAnimationBehaviorAlertPanel WindowAnimationBehavior = 5

type WindowBackingLocation uint

type WindowButton uint

const WindowCloseButton WindowButton = 0
const WindowMiniaturizeButton WindowButton = 1
const WindowZoomButton WindowButton = 2
const WindowToolbarButton WindowButton = 3
const WindowDocumentIconButton WindowButton = 4
const WindowDocumentVersionsButton WindowButton = 6

type WindowCollectionBehavior uint

const WindowCollectionBehaviorDefault WindowCollectionBehavior = 0
const WindowCollectionBehaviorCanJoinAllSpaces WindowCollectionBehavior = 1
const WindowCollectionBehaviorMoveToActiveSpace WindowCollectionBehavior = 2
const WindowCollectionBehaviorManaged WindowCollectionBehavior = 4
const WindowCollectionBehaviorTransient WindowCollectionBehavior = 8
const WindowCollectionBehaviorStationary WindowCollectionBehavior = 16
const WindowCollectionBehaviorParticipatesInCycle WindowCollectionBehavior = 32
const WindowCollectionBehaviorIgnoresCycle WindowCollectionBehavior = 64
const WindowCollectionBehaviorFullScreenPrimary WindowCollectionBehavior = 128
const WindowCollectionBehaviorFullScreenAuxiliary WindowCollectionBehavior = 256
const WindowCollectionBehaviorFullScreenNone WindowCollectionBehavior = 512
const WindowCollectionBehaviorFullScreenAllowsTiling WindowCollectionBehavior = 2048
const WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 4096

type WindowDepth int32

const WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 544
const WindowDepthSixtyfourBitRGB WindowDepth = 528
const WindowDepthTwentyfourBitRGB WindowDepth = 520

type WindowFrameAutosaveName string

type WindowLevel int

const FloatingWindowLevel WindowLevel = 3
const MainMenuWindowLevel WindowLevel = 24
const ModalPanelWindowLevel WindowLevel = 8
const NormalWindowLevel WindowLevel = 0
const PopUpMenuWindowLevel WindowLevel = 101
const ScreenSaverWindowLevel WindowLevel = 1000
const StatusWindowLevel WindowLevel = 25
const SubmenuWindowLevel WindowLevel = 3
const TornOffMenuWindowLevel WindowLevel = 3

type WindowListOptions int

const WindowListOrderedFrontToBack WindowListOptions = 1

type WindowNumberListOptions uint

const WindowNumberListAllApplications WindowNumberListOptions = 1
const WindowNumberListAllSpaces WindowNumberListOptions = 16

type WindowOcclusionState uint

const WindowOcclusionStateVisible WindowOcclusionState = 2

type WindowOrderingMode int

const WindowAbove WindowOrderingMode = 1
const WindowBelow WindowOrderingMode = -1
const WindowOut WindowOrderingMode = 0

type WindowPersistableFrameDescriptor string

type WindowSharingType uint

const WindowSharingNone WindowSharingType = 0
const WindowSharingReadOnly WindowSharingType = 1
const WindowSharingReadWrite WindowSharingType = 2

type WindowStyleMask uint

const WindowStyleMaskBorderless WindowStyleMask = 0
const WindowStyleMaskTitled WindowStyleMask = 1
const WindowStyleMaskClosable WindowStyleMask = 2
const WindowStyleMaskMiniaturizable WindowStyleMask = 4
const WindowStyleMaskResizable WindowStyleMask = 8
const WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 4096
const WindowStyleMaskFullScreen WindowStyleMask = 16384
const WindowStyleMaskFullSizeContentView WindowStyleMask = 32768
const WindowStyleMaskUtilityWindow WindowStyleMask = 16
const WindowStyleMaskDocModalWindow WindowStyleMask = 64
const WindowStyleMaskNonactivatingPanel WindowStyleMask = 128
const WindowStyleMaskHUDWindow WindowStyleMask = 8192

type WindowTabbingIdentifier string

type WindowTabbingMode int

const WindowTabbingModeAutomatic WindowTabbingMode = 0
const WindowTabbingModeDisallowed WindowTabbingMode = 2
const WindowTabbingModePreferred WindowTabbingMode = 1

type WindowTitleVisibility int

const WindowTitleVisible WindowTitleVisibility = 0
const WindowTitleHidden WindowTitleVisibility = 1

type WindowToolbarStyle int

const WindowToolbarStyleAutomatic WindowToolbarStyle = 0
const WindowToolbarStyleExpanded WindowToolbarStyle = 1
const WindowToolbarStylePreference WindowToolbarStyle = 2
const WindowToolbarStyleUnified WindowToolbarStyle = 3
const WindowToolbarStyleUnifiedCompact WindowToolbarStyle = 4

type WindowUserTabbingPreference int

const WindowUserTabbingPreferenceAlways WindowUserTabbingPreference = 1
const WindowUserTabbingPreferenceInFullScreen WindowUserTabbingPreference = 2
const WindowUserTabbingPreferenceManual WindowUserTabbingPreference = 0

type WorkspaceAuthorizationType int

const WorkspaceAuthorizationTypeCreateSymbolicLink WorkspaceAuthorizationType = 0
const WorkspaceAuthorizationTypeReplaceFile WorkspaceAuthorizationType = 2
const WorkspaceAuthorizationTypeSetAttributes WorkspaceAuthorizationType = 1

type WorkspaceDesktopImageOptionKey string

const WorkspaceDesktopImageScalingKey WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageScalingKey"
const WorkspaceDesktopImageAllowClippingKey WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageAllowClippingKey"
const WorkspaceDesktopImageFillColorKey WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageFillColorKey"

type WorkspaceFileOperationName string

type WorkspaceIconCreationOptions uint

const ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 2
const Exclude10_4ElementsIconCreationOption WorkspaceIconCreationOptions = 4

type WorkspaceLaunchConfigurationKey string

type WorkspaceLaunchOptions uint

type WritingDirection int

const WritingDirectionNatural WritingDirection = -1
const WritingDirectionLeftToRight WritingDirection = 0
const WritingDirectionRightToLeft WritingDirection = 1
