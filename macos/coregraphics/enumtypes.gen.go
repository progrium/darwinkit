// AUTO-GENERATED CODE, DO NOT MODIFY
package coregraphics

type BitmapInfo uint32

const KBitmapAlphaInfoMask BitmapInfo = 31
const KBitmapByteOrder16Big BitmapInfo = 12288
const KBitmapByteOrder16Little BitmapInfo = 4096
const KBitmapByteOrder32Big BitmapInfo = 16384
const KBitmapByteOrder32Little BitmapInfo = 8192
const KBitmapByteOrderDefault BitmapInfo = 0
const KBitmapByteOrderMask BitmapInfo = 28672
const KBitmapFloatComponents BitmapInfo = 256
const KBitmapFloatInfoMask BitmapInfo = 3840

type BlendMode int32

const KBlendModeClear BlendMode = 16
const KBlendModeColor BlendMode = 14
const KBlendModeColorBurn BlendMode = 7
const KBlendModeColorDodge BlendMode = 6
const KBlendModeCopy BlendMode = 17
const KBlendModeDarken BlendMode = 4
const KBlendModeDestinationAtop BlendMode = 24
const KBlendModeDestinationIn BlendMode = 22
const KBlendModeDestinationOut BlendMode = 23
const KBlendModeDestinationOver BlendMode = 21
const KBlendModeDifference BlendMode = 10
const KBlendModeExclusion BlendMode = 11
const KBlendModeHardLight BlendMode = 9
const KBlendModeHue BlendMode = 12
const KBlendModeLighten BlendMode = 5
const KBlendModeLuminosity BlendMode = 15
const KBlendModeMultiply BlendMode = 1
const KBlendModeNormal BlendMode = 0
const KBlendModeOverlay BlendMode = 3
const KBlendModePlusDarker BlendMode = 26
const KBlendModePlusLighter BlendMode = 27
const KBlendModeSaturation BlendMode = 13
const KBlendModeScreen BlendMode = 2
const KBlendModeSoftLight BlendMode = 8
const KBlendModeSourceAtop BlendMode = 20
const KBlendModeSourceIn BlendMode = 18
const KBlendModeSourceOut BlendMode = 19
const KBlendModeXOR BlendMode = 25

type ButtonCount uint32

type CaptureOptions uint32

const KCaptureNoFill CaptureOptions = 1
const KCaptureNoOptions CaptureOptions = 0

type CharCode uint16

type ColorConversionInfoTransformType uint32

const KColorConversionTransformApplySpace ColorConversionInfoTransformType = 2
const KColorConversionTransformFromSpace ColorConversionInfoTransformType = 0
const KColorConversionTransformToSpace ColorConversionInfoTransformType = 1

type ColorRenderingIntent int32

const KRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
const KRenderingIntentDefault ColorRenderingIntent = 0
const KRenderingIntentPerceptual ColorRenderingIntent = 3
const KRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
const KRenderingIntentSaturation ColorRenderingIntent = 4

type ColorSpaceModel int32

const KColorSpaceModelCMYK ColorSpaceModel = 2
const KColorSpaceModelDeviceN ColorSpaceModel = 4
const KColorSpaceModelIndexed ColorSpaceModel = 5
const KColorSpaceModelLab ColorSpaceModel = 3
const KColorSpaceModelMonochrome ColorSpaceModel = 0
const KColorSpaceModelPattern ColorSpaceModel = 6
const KColorSpaceModelRGB ColorSpaceModel = 1
const KColorSpaceModelUnknown ColorSpaceModel = -1
const KColorSpaceModelXYZ ColorSpaceModel = 7

type ConfigureOption uint32

const KConfigureForAppOnly ConfigureOption = 0
const KConfigureForSession ConfigureOption = 1
const KConfigurePermanently ConfigureOption = 2

type DirectDisplayID uint32

type DisplayBlendFraction float32

type DisplayChangeSummaryFlags uint32

const KDisplayAddFlag DisplayChangeSummaryFlags = 16
const KDisplayBeginConfigurationFlag DisplayChangeSummaryFlags = 1
const KDisplayDesktopShapeChangedFlag DisplayChangeSummaryFlags = 4096
const KDisplayDisabledFlag DisplayChangeSummaryFlags = 512
const KDisplayEnabledFlag DisplayChangeSummaryFlags = 256
const KDisplayMirrorFlag DisplayChangeSummaryFlags = 1024
const KDisplayMovedFlag DisplayChangeSummaryFlags = 2
const KDisplayRemoveFlag DisplayChangeSummaryFlags = 32
const KDisplaySetMainFlag DisplayChangeSummaryFlags = 4
const KDisplaySetModeFlag DisplayChangeSummaryFlags = 8
const KDisplayUnMirrorFlag DisplayChangeSummaryFlags = 2048

type DisplayCount uint32

type DisplayFadeInterval float32

type DisplayFadeReservationToken uint32

type DisplayReservationInterval float32

type DisplayStreamFrameStatus int32

const KDisplayStreamFrameStatusFrameBlank DisplayStreamFrameStatus = 2
const KDisplayStreamFrameStatusFrameComplete DisplayStreamFrameStatus = 0
const KDisplayStreamFrameStatusFrameIdle DisplayStreamFrameStatus = 1
const KDisplayStreamFrameStatusStopped DisplayStreamFrameStatus = 3

type DisplayStreamUpdateRectType int32

const KDisplayStreamUpdateDirtyRects DisplayStreamUpdateRectType = 2
const KDisplayStreamUpdateMovedRects DisplayStreamUpdateRectType = 1
const KDisplayStreamUpdateReducedDirtyRects DisplayStreamUpdateRectType = 3
const KDisplayStreamUpdateRefreshedRects DisplayStreamUpdateRectType = 0

type Error int32

const KErrorCannotComplete Error = 1004
const KErrorFailure Error = 1000
const KErrorIllegalArgument Error = 1001
const KErrorInvalidConnection Error = 1002
const KErrorInvalidContext Error = 1003
const KErrorInvalidOperation Error = 1010
const KErrorNoneAvailable Error = 1011
const KErrorNotImplemented Error = 1006
const KErrorRangeCheck Error = 1007
const KErrorSuccess Error = 0
const KErrorTypeCheck Error = 1008

type EventField uint32

const KEventSourceGroupID EventField = 44
const KEventSourceStateID EventField = 45
const KEventSourceUnixProcessID EventField = 41
const KEventSourceUserData EventField = 42
const KEventSourceUserID EventField = 43
const KEventTargetProcessSerialNumber EventField = 39
const KEventTargetUnixProcessID EventField = 40
const KEventUnacceleratedPointerMovementX EventField = 170
const KEventUnacceleratedPointerMovementY EventField = 171
const KKeyboardEventAutorepeat EventField = 8
const KKeyboardEventKeyboardType EventField = 10
const KKeyboardEventKeycode EventField = 9
const KMouseEventButtonNumber EventField = 3
const KMouseEventClickState EventField = 1
const KMouseEventDeltaX EventField = 4
const KMouseEventDeltaY EventField = 5
const KMouseEventInstantMouser EventField = 6
const KMouseEventNumber EventField = 0
const KMouseEventPressure EventField = 2
const KMouseEventSubtype EventField = 7
const KMouseEventWindowUnderMousePointer EventField = 91
const KMouseEventWindowUnderMousePointerThatCanHandleThisEvent EventField = 92
const KScrollWheelEventDeltaAxis1 EventField = 11
const KScrollWheelEventDeltaAxis2 EventField = 12
const KScrollWheelEventDeltaAxis3 EventField = 13
const KScrollWheelEventFixedPtDeltaAxis1 EventField = 93
const KScrollWheelEventFixedPtDeltaAxis2 EventField = 94
const KScrollWheelEventFixedPtDeltaAxis3 EventField = 95
const KScrollWheelEventInstantMouser EventField = 14
const KScrollWheelEventIsContinuous EventField = 88
const KScrollWheelEventMomentumPhase EventField = 123
const KScrollWheelEventPointDeltaAxis1 EventField = 96
const KScrollWheelEventPointDeltaAxis2 EventField = 97
const KScrollWheelEventPointDeltaAxis3 EventField = 98
const KScrollWheelEventScrollCount EventField = 100
const KScrollWheelEventScrollPhase EventField = 99
const KTabletEventDeviceID EventField = 24
const KTabletEventPointButtons EventField = 18
const KTabletEventPointPressure EventField = 19
const KTabletEventPointX EventField = 15
const KTabletEventPointY EventField = 16
const KTabletEventPointZ EventField = 17
const KTabletEventRotation EventField = 22
const KTabletEventTangentialPressure EventField = 23
const KTabletEventTiltX EventField = 20
const KTabletEventTiltY EventField = 21
const KTabletEventVendor1 EventField = 25
const KTabletEventVendor2 EventField = 26
const KTabletEventVendor3 EventField = 27
const KTabletProximityEventCapabilityMask EventField = 36
const KTabletProximityEventDeviceID EventField = 31
const KTabletProximityEventEnterProximity EventField = 38
const KTabletProximityEventPointerID EventField = 30
const KTabletProximityEventPointerType EventField = 37
const KTabletProximityEventSystemTabletID EventField = 32
const KTabletProximityEventTabletID EventField = 29
const KTabletProximityEventVendorID EventField = 28
const KTabletProximityEventVendorPointerSerialNumber EventField = 34
const KTabletProximityEventVendorPointerType EventField = 33
const KTabletProximityEventVendorUniqueID EventField = 35

type EventFilterMask uint32

const KEventFilterMaskPermitLocalKeyboardEvents EventFilterMask = 2
const KEventFilterMaskPermitLocalMouseEvents EventFilterMask = 1
const KEventFilterMaskPermitSystemDefinedEvents EventFilterMask = 4

type EventFlags uint64

const KEventFlagMaskAlphaShift EventFlags = 65536
const KEventFlagMaskAlternate EventFlags = 524288
const KEventFlagMaskCommand EventFlags = 1048576
const KEventFlagMaskControl EventFlags = 262144
const KEventFlagMaskHelp EventFlags = 4194304
const KEventFlagMaskNonCoalesced EventFlags = 256
const KEventFlagMaskNumericPad EventFlags = 2097152
const KEventFlagMaskSecondaryFn EventFlags = 8388608
const KEventFlagMaskShift EventFlags = 131072

type EventMask uint64

type EventMouseSubtype uint32

const KEventMouseSubtypeDefault EventMouseSubtype = 0
const KEventMouseSubtypeTabletPoint EventMouseSubtype = 1
const KEventMouseSubtypeTabletProximity EventMouseSubtype = 2

type EventSourceKeyboardType uint32

type EventSourceStateID int32

const KEventSourceStateCombinedSessionState EventSourceStateID = 0
const KEventSourceStateHIDSystemState EventSourceStateID = 1
const KEventSourceStatePrivate EventSourceStateID = -1

type EventSuppressionState uint32

const KEventSuppressionStateRemoteMouseDrag EventSuppressionState = 1
const KEventSuppressionStateSuppressionInterval EventSuppressionState = 0
const KNumberOfEventSuppressionStates EventSuppressionState = 2

type EventTapLocation uint32

const KAnnotatedSessionEventTap EventTapLocation = 2
const KHIDEventTap EventTapLocation = 0
const KSessionEventTap EventTapLocation = 1

type EventTapOptions uint32

const KEventTapOptionDefault EventTapOptions = 0
const KEventTapOptionListenOnly EventTapOptions = 1

type EventTapPlacement uint32

const KHeadInsertEventTap EventTapPlacement = 0
const KTailAppendEventTap EventTapPlacement = 1

type EventTimestamp uint64

type EventType uint32

const KEventFlagsChanged EventType = 12
const KEventKeyDown EventType = 10
const KEventKeyUp EventType = 11
const KEventLeftMouseDown EventType = 1
const KEventLeftMouseDragged EventType = 6
const KEventLeftMouseUp EventType = 2
const KEventMouseMoved EventType = 5
const KEventNull EventType = 0
const KEventOtherMouseDown EventType = 25
const KEventOtherMouseDragged EventType = 27
const KEventOtherMouseUp EventType = 26
const KEventRightMouseDown EventType = 3
const KEventRightMouseDragged EventType = 7
const KEventRightMouseUp EventType = 4
const KEventScrollWheel EventType = 22
const KEventTabletPointer EventType = 23
const KEventTabletProximity EventType = 24
const KEventTapDisabledByTimeout EventType = 4294967294
const KEventTapDisabledByUserInput EventType = 4294967295

type Float = float64

type FontIndex int

const KFontIndexInvalid FontIndex = 65535
const KFontIndexMax FontIndex = 65534
const KGlyphMax FontIndex = 65534

type FontPostScriptFormat int32

const KFontPostScriptFormatType1 FontPostScriptFormat = 1
const KFontPostScriptFormatType3 FontPostScriptFormat = 3
const KFontPostScriptFormatType42 FontPostScriptFormat = 42

type GammaValue float32

type GesturePhase uint32

const KGesturePhaseBegan GesturePhase = 1
const KGesturePhaseCancelled GesturePhase = 8
const KGesturePhaseChanged GesturePhase = 2
const KGesturePhaseEnded GesturePhase = 4
const KGesturePhaseMayBegin GesturePhase = 128
const KGesturePhaseNone GesturePhase = 0

type Glyph FontIndex

type GlyphDeprecatedEnum int32

type GradientDrawingOptions uint32

const KGradientDrawsAfterEndLocation GradientDrawingOptions = 2
const KGradientDrawsBeforeStartLocation GradientDrawingOptions = 1

type ImageAlphaInfo uint32

const KImageAlphaFirst ImageAlphaInfo = 4
const KImageAlphaLast ImageAlphaInfo = 3
const KImageAlphaNone ImageAlphaInfo = 0
const KImageAlphaNoneSkipFirst ImageAlphaInfo = 6
const KImageAlphaNoneSkipLast ImageAlphaInfo = 5
const KImageAlphaOnly ImageAlphaInfo = 7
const KImageAlphaPremultipliedFirst ImageAlphaInfo = 2
const KImageAlphaPremultipliedLast ImageAlphaInfo = 1

type ImageByteOrderInfo uint32

const KImageByteOrder16Big ImageByteOrderInfo = 12288
const KImageByteOrder16Little ImageByteOrderInfo = 4096
const KImageByteOrder32Big ImageByteOrderInfo = 16384
const KImageByteOrder32Little ImageByteOrderInfo = 8192
const KImageByteOrderDefault ImageByteOrderInfo = 0
const KImageByteOrderMask ImageByteOrderInfo = 28672

type ImagePixelFormatInfo uint32

const KImagePixelFormatMask ImagePixelFormatInfo = 983040
const KImagePixelFormatPacked ImagePixelFormatInfo = 0
const KImagePixelFormatRGB101010 ImagePixelFormatInfo = 196608
const KImagePixelFormatRGB555 ImagePixelFormatInfo = 65536
const KImagePixelFormatRGB565 ImagePixelFormatInfo = 131072
const KImagePixelFormatRGBCIF10 ImagePixelFormatInfo = 262144

type InterpolationQuality int32

const KInterpolationDefault InterpolationQuality = 0
const KInterpolationHigh InterpolationQuality = 3
const KInterpolationLow InterpolationQuality = 2
const KInterpolationMedium InterpolationQuality = 4
const KInterpolationNone InterpolationQuality = 1

type KeyCode uint16

type LineCap int32

const KLineCapButt LineCap = 0
const KLineCapRound LineCap = 1
const KLineCapSquare LineCap = 2

type LineJoin int32

const KLineJoinBevel LineJoin = 2
const KLineJoinMiter LineJoin = 0
const KLineJoinRound LineJoin = 1

type MomentumScrollPhase uint32

const KMomentumScrollPhaseBegin MomentumScrollPhase = 1
const KMomentumScrollPhaseContinue MomentumScrollPhase = 2
const KMomentumScrollPhaseEnd MomentumScrollPhase = 3
const KMomentumScrollPhaseNone MomentumScrollPhase = 0

type MouseButton uint32

const KMouseButtonCenter MouseButton = 2
const KMouseButtonLeft MouseButton = 0
const KMouseButtonRight MouseButton = 1

type OpenGLDisplayMask uint32

type PDFAccessPermissions uint32

const KPDFAllowsCommenting PDFAccessPermissions = 64
const KPDFAllowsContentAccessibility PDFAccessPermissions = 32
const KPDFAllowsContentCopying PDFAccessPermissions = 16
const KPDFAllowsDocumentAssembly PDFAccessPermissions = 8
const KPDFAllowsDocumentChanges PDFAccessPermissions = 4
const KPDFAllowsFormFieldEntry PDFAccessPermissions = 128
const KPDFAllowsHighQualityPrinting PDFAccessPermissions = 2
const KPDFAllowsLowQualityPrinting PDFAccessPermissions = 1

type PDFBoolean byte

type PDFBox int32

const KPDFArtBox PDFBox = 4
const KPDFBleedBox PDFBox = 2
const KPDFCropBox PDFBox = 1
const KPDFMediaBox PDFBox = 0
const KPDFTrimBox PDFBox = 3

type PDFDataFormat int32

const PDFDataFormatJPEG2000 PDFDataFormat = 2
const PDFDataFormatJPEGEncoded PDFDataFormat = 1
const PDFDataFormatRaw PDFDataFormat = 0

type PDFInteger int32

type PDFObjectType int32

const KPDFObjectTypeArray PDFObjectType = 7
const KPDFObjectTypeBoolean PDFObjectType = 2
const KPDFObjectTypeDictionary PDFObjectType = 8
const KPDFObjectTypeInteger PDFObjectType = 3
const KPDFObjectTypeName PDFObjectType = 5
const KPDFObjectTypeNull PDFObjectType = 1
const KPDFObjectTypeReal PDFObjectType = 4
const KPDFObjectTypeStream PDFObjectType = 9
const KPDFObjectTypeString PDFObjectType = 6

type PDFReal float32

type PDFTagType int32

const PDFTagTypeAnnotation PDFTagType = 507
const PDFTagTypeArt PDFTagType = 102
const PDFTagTypeBibliography PDFTagType = 504
const PDFTagTypeBlockQuote PDFTagType = 105
const PDFTagTypeCaption PDFTagType = 106
const PDFTagTypeCode PDFTagType = 505
const PDFTagTypeDiv PDFTagType = 104
const PDFTagTypeDocument PDFTagType = 100
const PDFTagTypeFigure PDFTagType = 700
const PDFTagTypeForm PDFTagType = 702
const PDFTagTypeFormula PDFTagType = 701
const PDFTagTypeHeader PDFTagType = 201
const PDFTagTypeHeader1 PDFTagType = 202
const PDFTagTypeHeader2 PDFTagType = 203
const PDFTagTypeHeader3 PDFTagType = 204
const PDFTagTypeHeader4 PDFTagType = 205
const PDFTagTypeHeader5 PDFTagType = 206
const PDFTagTypeHeader6 PDFTagType = 207
const PDFTagTypeIndex PDFTagType = 109
const PDFTagTypeLabel PDFTagType = 302
const PDFTagTypeLink PDFTagType = 506
const PDFTagTypeList PDFTagType = 300
const PDFTagTypeListBody PDFTagType = 303
const PDFTagTypeListItem PDFTagType = 301
const PDFTagTypeNonStructure PDFTagType = 110
const PDFTagTypeNote PDFTagType = 502
const PDFTagTypeParagraph PDFTagType = 200
const PDFTagTypePart PDFTagType = 101
const PDFTagTypePrivate PDFTagType = 111
const PDFTagTypeQuote PDFTagType = 501
const PDFTagTypeReference PDFTagType = 503
const PDFTagTypeRuby PDFTagType = 600
const PDFTagTypeRubyAnnotationText PDFTagType = 602
const PDFTagTypeRubyBaseText PDFTagType = 601
const PDFTagTypeRubyPunctuation PDFTagType = 603
const PDFTagTypeSection PDFTagType = 103
const PDFTagTypeSpan PDFTagType = 500
const PDFTagTypeTOC PDFTagType = 107
const PDFTagTypeTOCI PDFTagType = 108
const PDFTagTypeTable PDFTagType = 400
const PDFTagTypeTableBody PDFTagType = 405
const PDFTagTypeTableDataCell PDFTagType = 403
const PDFTagTypeTableFooter PDFTagType = 406
const PDFTagTypeTableHeader PDFTagType = 404
const PDFTagTypeTableHeaderCell PDFTagType = 402
const PDFTagTypeTableRow PDFTagType = 401
const PDFTagTypeWarichu PDFTagType = 604
const PDFTagTypeWarichuPunctiation PDFTagType = 606
const PDFTagTypeWarichuText PDFTagType = 605

type PathDrawingMode int32

const KPathEOFill PathDrawingMode = 1
const KPathEOFillStroke PathDrawingMode = 4
const KPathFill PathDrawingMode = 0
const KPathFillStroke PathDrawingMode = 3
const KPathStroke PathDrawingMode = 2

type PathElementType int32

const KPathElementAddCurveToPoint PathElementType = 3
const KPathElementAddLineToPoint PathElementType = 1
const KPathElementAddQuadCurveToPoint PathElementType = 2
const KPathElementCloseSubpath PathElementType = 4
const KPathElementMoveToPoint PathElementType = 0

type PatternTiling int32

const KPatternTilingConstantSpacing PatternTiling = 2
const KPatternTilingConstantSpacingMinimalDistortion PatternTiling = 1
const KPatternTilingNoDistortion PatternTiling = 0

type RectCount uint32

type RectEdge uint32

const RectMaxXEdge RectEdge = 2
const RectMaxYEdge RectEdge = 3
const RectMinXEdge RectEdge = 0
const RectMinYEdge RectEdge = 1

type RefreshRate float64

type ScreenUpdateOperation uint32

const KScreenUpdateOperationMove ScreenUpdateOperation = 1
const KScreenUpdateOperationReducedDirtyRectangleCount ScreenUpdateOperation = 2147483648
const KScreenUpdateOperationRefresh ScreenUpdateOperation = 0

type ScrollEventUnit uint32

const KScrollEventUnitLine ScrollEventUnit = 1
const KScrollEventUnitPixel ScrollEventUnit = 0

type ScrollPhase uint32

const KScrollPhaseBegan ScrollPhase = 1
const KScrollPhaseCancelled ScrollPhase = 8
const KScrollPhaseChanged ScrollPhase = 2
const KScrollPhaseEnded ScrollPhase = 4
const KScrollPhaseMayBegin ScrollPhase = 128

type TextDrawingMode int32

const KTextClip TextDrawingMode = 7
const KTextFill TextDrawingMode = 0
const KTextFillClip TextDrawingMode = 4
const KTextFillStroke TextDrawingMode = 2
const KTextFillStrokeClip TextDrawingMode = 6
const KTextInvisible TextDrawingMode = 3
const KTextStroke TextDrawingMode = 1
const KTextStrokeClip TextDrawingMode = 5

type TextEncoding int32

const KEncodingFontSpecific TextEncoding = 0
const KEncodingMacRoman TextEncoding = 1

type WheelCount uint32

type WindowBackingType uint32

const KBackingStoreBuffered WindowBackingType = 2
const KBackingStoreNonretained WindowBackingType = 1
const KBackingStoreRetained WindowBackingType = 0

type WindowID uint32

type WindowImageOption uint32

const KWindowImageBestResolution WindowImageOption = 8
const KWindowImageBoundsIgnoreFraming WindowImageOption = 1
const KWindowImageDefault WindowImageOption = 0
const KWindowImageNominalResolution WindowImageOption = 16
const KWindowImageOnlyShadows WindowImageOption = 4
const KWindowImageShouldBeOpaque WindowImageOption = 2

type WindowLevel int32

type WindowLevelKey int32

const KAssistiveTechHighWindowLevelKey WindowLevelKey = 20
const KBackstopMenuLevelKey WindowLevelKey = 3
const KBaseWindowLevelKey WindowLevelKey = 0
const KCursorWindowLevelKey WindowLevelKey = 19
const KDesktopIconWindowLevelKey WindowLevelKey = 18
const KDesktopWindowLevelKey WindowLevelKey = 2
const KDockWindowLevelKey WindowLevelKey = 7
const KDraggingWindowLevelKey WindowLevelKey = 12
const KFloatingWindowLevelKey WindowLevelKey = 5
const KHelpWindowLevelKey WindowLevelKey = 16
const KMainMenuWindowLevelKey WindowLevelKey = 8
const KMaximumWindowLevelKey WindowLevelKey = 14
const KMinimumWindowLevelKey WindowLevelKey = 1
const KModalPanelWindowLevelKey WindowLevelKey = 10
const KNormalWindowLevelKey WindowLevelKey = 4
const KNumberOfWindowLevelKeys WindowLevelKey = 21
const KOverlayWindowLevelKey WindowLevelKey = 15
const KPopUpMenuWindowLevelKey WindowLevelKey = 11
const KScreenSaverWindowLevelKey WindowLevelKey = 13
const KStatusWindowLevelKey WindowLevelKey = 9
const KTornOffMenuWindowLevelKey WindowLevelKey = 6
const KUtilityWindowLevelKey WindowLevelKey = 17

type WindowListOption uint32

const KWindowListExcludeDesktopElements WindowListOption = 16
const KWindowListOptionAll WindowListOption = 0
const KWindowListOptionIncludingWindow WindowListOption = 8
const KWindowListOptionOnScreenAboveWindow WindowListOption = 2
const KWindowListOptionOnScreenBelowWindow WindowListOption = 4
const KWindowListOptionOnScreenOnly WindowListOption = 1

type WindowSharingType uint32

const KWindowSharingNone WindowSharingType = 0
const KWindowSharingReadOnly WindowSharingType = 1
const KWindowSharingReadWrite WindowSharingType = 2
