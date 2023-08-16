// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import "math"

// Keys to include in the options dictionary when displaying an About panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaboutpaneloptionkey?language=objc
type AboutPanelOptionKey string

const (
	AboutPanelOptionApplicationIcon    AboutPanelOptionKey = "ApplicationIcon"
	AboutPanelOptionApplicationName    AboutPanelOptionKey = "ApplicationName"
	AboutPanelOptionApplicationVersion AboutPanelOptionKey = "ApplicationVersion"
	AboutPanelOptionCredits            AboutPanelOptionKey = "Credits"
	AboutPanelOptionVersion            AboutPanelOptionKey = "Version"
)

// Constants that describe types of actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityactionname?language=objc
type AccessibilityActionName string

const (
	AccessibilityCancelAction          AccessibilityActionName = "AXCancel"
	AccessibilityConfirmAction         AccessibilityActionName = "AXConfirm"
	AccessibilityDecrementAction       AccessibilityActionName = "AXDecrement"
	AccessibilityDeleteAction          AccessibilityActionName = "AXDelete"
	AccessibilityIncrementAction       AccessibilityActionName = "AXIncrement"
	AccessibilityPickAction            AccessibilityActionName = "AXPick"
	AccessibilityPressAction           AccessibilityActionName = "AXPress"
	AccessibilityRaiseAction           AccessibilityActionName = "AXRaise"
	AccessibilityShowAlternateUIAction AccessibilityActionName = "AXShowAlternateUI"
	AccessibilityShowDefaultUIAction   AccessibilityActionName = "AXShowDefaultUI"
	AccessibilityShowMenuAction        AccessibilityActionName = "AXShowMenu"
)

// Keys for annotation attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityannotationattributekey?language=objc
type AccessibilityAnnotationAttributeKey string

const (
	AccessibilityAnnotationElement  AccessibilityAnnotationAttributeKey = "AXAnnotationElement"
	AccessibilityAnnotationLabel    AccessibilityAnnotationAttributeKey = "AXAnnotationLabel"
	AccessibilityAnnotationLocation AccessibilityAnnotationAttributeKey = "AXAnnotationLocation"
)

// Constants that specify the position where the annotation applies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityannotationposition?language=objc
type AccessibilityAnnotationPosition int

const (
	AccessibilityAnnotationPositionEnd       AccessibilityAnnotationPosition = 2
	AccessibilityAnnotationPositionFullRange AccessibilityAnnotationPosition = 0
	AccessibilityAnnotationPositionStart     AccessibilityAnnotationPosition = 1
)

// Constants that describe attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityattributename?language=objc
type AccessibilityAttributeName string

const (
	AccessibilityActivationPointAttribute            AccessibilityAttributeName = "AXActivationPoint"
	AccessibilityAllowedValuesAttribute              AccessibilityAttributeName = "AXAllowedValues"
	AccessibilityAlternateUIVisibleAttribute         AccessibilityAttributeName = "AXAlternateUIVisible"
	AccessibilityCancelButtonAttribute               AccessibilityAttributeName = "AXCancelButton"
	AccessibilityChildrenAttribute                   AccessibilityAttributeName = "AXChildren"
	AccessibilityClearButtonAttribute                AccessibilityAttributeName = "AXClearButton"
	AccessibilityCloseButtonAttribute                AccessibilityAttributeName = "AXCloseButton"
	AccessibilityColumnCountAttribute                AccessibilityAttributeName = "AXColumnCount"
	AccessibilityColumnHeaderUIElementsAttribute     AccessibilityAttributeName = "AXColumnHeaderUIElements"
	AccessibilityColumnIndexRangeAttribute           AccessibilityAttributeName = "AXColumnIndexRange"
	AccessibilityColumnTitlesAttribute               AccessibilityAttributeName = "AXColumnTitles"
	AccessibilityColumnsAttribute                    AccessibilityAttributeName = "AXColumns"
	AccessibilityContainsProtectedContentAttribute   AccessibilityAttributeName = "AXContainsProtectedContent"
	AccessibilityContentsAttribute                   AccessibilityAttributeName = "AXContents"
	AccessibilityCriticalValueAttribute              AccessibilityAttributeName = "AXCriticalValue"
	AccessibilityDecrementButtonAttribute            AccessibilityAttributeName = "AXDecrementButton"
	AccessibilityDefaultButtonAttribute              AccessibilityAttributeName = "AXDefaultButton"
	AccessibilityDescriptionAttribute                AccessibilityAttributeName = "AXDescription"
	AccessibilityDisclosedByRowAttribute             AccessibilityAttributeName = "AXDisclosedByRow"
	AccessibilityDisclosedRowsAttribute              AccessibilityAttributeName = "AXDisclosedRows"
	AccessibilityDisclosingAttribute                 AccessibilityAttributeName = "AXDisclosing"
	AccessibilityDisclosureLevelAttribute            AccessibilityAttributeName = "AXDisclosureLevel"
	AccessibilityDocumentAttribute                   AccessibilityAttributeName = "AXDocument"
	AccessibilityEditedAttribute                     AccessibilityAttributeName = "AXEdited"
	AccessibilityEnabledAttribute                    AccessibilityAttributeName = "AXEnabled"
	AccessibilityExpandedAttribute                   AccessibilityAttributeName = "AXExpanded"
	AccessibilityExtrasMenuBarAttribute              AccessibilityAttributeName = "AXExtrasMenuBar"
	AccessibilityFilenameAttribute                   AccessibilityAttributeName = "AXFilename"
	AccessibilityFocusedAttribute                    AccessibilityAttributeName = "AXFocused"
	AccessibilityFocusedUIElementAttribute           AccessibilityAttributeName = "AXFocusedUIElement"
	AccessibilityFocusedWindowAttribute              AccessibilityAttributeName = "AXFocusedWindow"
	AccessibilityFrontmostAttribute                  AccessibilityAttributeName = "AXFrontmost"
	AccessibilityFullScreenButtonAttribute           AccessibilityAttributeName = "AXFullScreenButton"
	AccessibilityGrowAreaAttribute                   AccessibilityAttributeName = "AXGrowArea"
	AccessibilityHandlesAttribute                    AccessibilityAttributeName = "AXHandles"
	AccessibilityHeaderAttribute                     AccessibilityAttributeName = "AXHeader"
	AccessibilityHelpAttribute                       AccessibilityAttributeName = "AXHelp"
	AccessibilityHiddenAttribute                     AccessibilityAttributeName = "AXHidden"
	AccessibilityHorizontalScrollBarAttribute        AccessibilityAttributeName = "AXHorizontalScrollBar"
	AccessibilityHorizontalUnitDescriptionAttribute  AccessibilityAttributeName = "AXHorizontalUnitDescription"
	AccessibilityHorizontalUnitsAttribute            AccessibilityAttributeName = "AXHorizontalUnits"
	AccessibilityIdentifierAttribute                 AccessibilityAttributeName = "AXIdentifier"
	AccessibilityIncrementButtonAttribute            AccessibilityAttributeName = "AXIncrementButton"
	AccessibilityIndexAttribute                      AccessibilityAttributeName = "AXIndex"
	AccessibilityInsertionPointLineNumberAttribute   AccessibilityAttributeName = "AXInsertionPointLineNumber"
	AccessibilityLabelUIElementsAttribute            AccessibilityAttributeName = "AXLabelUIElements"
	AccessibilityLabelValueAttribute                 AccessibilityAttributeName = "AXLabelValue"
	AccessibilityLinkedUIElementsAttribute           AccessibilityAttributeName = "AXLinkedUIElements"
	AccessibilityMainAttribute                       AccessibilityAttributeName = "AXMain"
	AccessibilityMainWindowAttribute                 AccessibilityAttributeName = "AXMainWindow"
	AccessibilityMarkerGroupUIElementAttribute       AccessibilityAttributeName = "AXMarkerGroupUIElement"
	AccessibilityMarkerTypeAttribute                 AccessibilityAttributeName = "AXMarkerType"
	AccessibilityMarkerTypeDescriptionAttribute      AccessibilityAttributeName = "AXMarkerTypeDescription"
	AccessibilityMarkerUIElementsAttribute           AccessibilityAttributeName = "AXMarkerUIElements"
	AccessibilityMarkerValuesAttribute               AccessibilityAttributeName = "AXMarkerValues"
	AccessibilityMatteContentUIElementAttribute      AccessibilityAttributeName = "AXMatteContentUIElement"
	AccessibilityMatteHoleAttribute                  AccessibilityAttributeName = "AXMatteHole"
	AccessibilityMaxValueAttribute                   AccessibilityAttributeName = "AXMaxValue"
	AccessibilityMenuBarAttribute                    AccessibilityAttributeName = "AXMenuBar"
	AccessibilityMinValueAttribute                   AccessibilityAttributeName = "AXMinValue"
	AccessibilityMinimizeButtonAttribute             AccessibilityAttributeName = "AXMinimizeButton"
	AccessibilityMinimizedAttribute                  AccessibilityAttributeName = "AXMinimized"
	AccessibilityModalAttribute                      AccessibilityAttributeName = "AXModal"
	AccessibilityNextContentsAttribute               AccessibilityAttributeName = "AXNextContents"
	AccessibilityNumberOfCharactersAttribute         AccessibilityAttributeName = "AXNumberOfCharacters"
	AccessibilityOrderedByRowAttribute               AccessibilityAttributeName = "AXOrderedByRow"
	AccessibilityOrientationAttribute                AccessibilityAttributeName = "AXOrientation"
	AccessibilityOverflowButtonAttribute             AccessibilityAttributeName = "AXOverflowButton"
	AccessibilityParentAttribute                     AccessibilityAttributeName = "AXParent"
	AccessibilityPlaceholderValueAttribute           AccessibilityAttributeName = "AXPlaceholderValue"
	AccessibilityPositionAttribute                   AccessibilityAttributeName = "AXPosition"
	AccessibilityPreviousContentsAttribute           AccessibilityAttributeName = "AXPreviousContents"
	AccessibilityProxyAttribute                      AccessibilityAttributeName = "AXProxy"
	AccessibilityRequiredAttribute                   AccessibilityAttributeName = "AXRequired"
	AccessibilityRoleAttribute                       AccessibilityAttributeName = "AXRole"
	AccessibilityRoleDescriptionAttribute            AccessibilityAttributeName = "AXRoleDescription"
	AccessibilityRowCountAttribute                   AccessibilityAttributeName = "AXRowCount"
	AccessibilityRowHeaderUIElementsAttribute        AccessibilityAttributeName = "AXRowHeaderUIElements"
	AccessibilityRowIndexRangeAttribute              AccessibilityAttributeName = "AXRowIndexRange"
	AccessibilityRowsAttribute                       AccessibilityAttributeName = "AXRows"
	AccessibilitySearchButtonAttribute               AccessibilityAttributeName = "AXSearchButton"
	AccessibilitySearchMenuAttribute                 AccessibilityAttributeName = "AXSearchMenu"
	AccessibilitySelectedAttribute                   AccessibilityAttributeName = "AXSelected"
	AccessibilitySelectedCellsAttribute              AccessibilityAttributeName = "AXSelectedCells"
	AccessibilitySelectedChildrenAttribute           AccessibilityAttributeName = "AXSelectedChildren"
	AccessibilitySelectedColumnsAttribute            AccessibilityAttributeName = "AXSelectedColumns"
	AccessibilitySelectedRowsAttribute               AccessibilityAttributeName = "AXSelectedRows"
	AccessibilitySelectedTextAttribute               AccessibilityAttributeName = "AXSelectedText"
	AccessibilitySelectedTextRangeAttribute          AccessibilityAttributeName = "AXSelectedTextRange"
	AccessibilitySelectedTextRangesAttribute         AccessibilityAttributeName = "AXSelectedTextRanges"
	AccessibilityServesAsTitleForUIElementsAttribute AccessibilityAttributeName = "AXServesAsTitleForUIElements"
	AccessibilitySharedCharacterRangeAttribute       AccessibilityAttributeName = "AXSharedCharacterRange"
	AccessibilitySharedFocusElementsAttribute        AccessibilityAttributeName = "AXSharedFocusElements"
	AccessibilitySharedTextUIElementsAttribute       AccessibilityAttributeName = "AXSharedTextUIElements"
	AccessibilityShownMenuAttribute                  AccessibilityAttributeName = "AXShownMenuUIElement"
	AccessibilitySizeAttribute                       AccessibilityAttributeName = "AXSize"
	AccessibilitySortDirectionAttribute              AccessibilityAttributeName = "AXSortDirection"
	AccessibilitySplittersAttribute                  AccessibilityAttributeName = "AXSplitters"
	AccessibilitySubroleAttribute                    AccessibilityAttributeName = "AXSubrole"
	AccessibilityTabsAttribute                       AccessibilityAttributeName = "AXTabs"
	AccessibilityTitleAttribute                      AccessibilityAttributeName = "AXTitle"
	AccessibilityTitleUIElementAttribute             AccessibilityAttributeName = "AXTitleUIElement"
	AccessibilityToolbarButtonAttribute              AccessibilityAttributeName = "AXToolbarButton"
	AccessibilityTopLevelUIElementAttribute          AccessibilityAttributeName = "AXTopLevelUIElement"
	AccessibilityURLAttribute                        AccessibilityAttributeName = "AXURL"
	AccessibilityUnitDescriptionAttribute            AccessibilityAttributeName = "AXUnitDescription"
	AccessibilityUnitsAttribute                      AccessibilityAttributeName = "AXUnits"
	AccessibilityValueAttribute                      AccessibilityAttributeName = "AXValue"
	AccessibilityValueDescriptionAttribute           AccessibilityAttributeName = "AXValueDescription"
	AccessibilityVerticalScrollBarAttribute          AccessibilityAttributeName = "AXVerticalScrollBar"
	AccessibilityVerticalUnitDescriptionAttribute    AccessibilityAttributeName = "AXVerticalUnitDescription"
	AccessibilityVerticalUnitsAttribute              AccessibilityAttributeName = "AXVerticalUnits"
	AccessibilityVisibleCellsAttribute               AccessibilityAttributeName = "AXVisibleCells"
	AccessibilityVisibleCharacterRangeAttribute      AccessibilityAttributeName = "AXVisibleCharacterRange"
	AccessibilityVisibleChildrenAttribute            AccessibilityAttributeName = "AXVisibleChildren"
	AccessibilityVisibleColumnsAttribute             AccessibilityAttributeName = "AXVisibleColumns"
	AccessibilityVisibleRowsAttribute                AccessibilityAttributeName = "AXVisibleRows"
	AccessibilityWarningValueAttribute               AccessibilityAttributeName = "AXWarningValue"
	AccessibilityWindowAttribute                     AccessibilityAttributeName = "AXWindow"
	AccessibilityWindowsAttribute                    AccessibilityAttributeName = "AXWindows"
	AccessibilityZoomButtonAttribute                 AccessibilityAttributeName = "AXZoomButton"
)

// Constants that describe the direction to search for an item result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotorsearchdirection?language=objc
type AccessibilityCustomRotorSearchDirection int

const (
	AccessibilityCustomRotorSearchDirectionNext     AccessibilityCustomRotorSearchDirection = 1
	AccessibilityCustomRotorSearchDirectionPrevious AccessibilityCustomRotorSearchDirection = 0
)

// Constants that indicate the type of content that the rotor represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotortype?language=objc
type AccessibilityCustomRotorType int

const (
	AccessibilityCustomRotorTypeAnnotation     AccessibilityCustomRotorType = 2
	AccessibilityCustomRotorTypeAny            AccessibilityCustomRotorType = 1
	AccessibilityCustomRotorTypeAudiograph     AccessibilityCustomRotorType = 21
	AccessibilityCustomRotorTypeBoldText       AccessibilityCustomRotorType = 3
	AccessibilityCustomRotorTypeCustom         AccessibilityCustomRotorType = 0
	AccessibilityCustomRotorTypeHeading        AccessibilityCustomRotorType = 4
	AccessibilityCustomRotorTypeHeadingLevel1  AccessibilityCustomRotorType = 5
	AccessibilityCustomRotorTypeHeadingLevel2  AccessibilityCustomRotorType = 6
	AccessibilityCustomRotorTypeHeadingLevel3  AccessibilityCustomRotorType = 7
	AccessibilityCustomRotorTypeHeadingLevel4  AccessibilityCustomRotorType = 8
	AccessibilityCustomRotorTypeHeadingLevel5  AccessibilityCustomRotorType = 9
	AccessibilityCustomRotorTypeHeadingLevel6  AccessibilityCustomRotorType = 10
	AccessibilityCustomRotorTypeImage          AccessibilityCustomRotorType = 11
	AccessibilityCustomRotorTypeItalicText     AccessibilityCustomRotorType = 12
	AccessibilityCustomRotorTypeLandmark       AccessibilityCustomRotorType = 13
	AccessibilityCustomRotorTypeLink           AccessibilityCustomRotorType = 14
	AccessibilityCustomRotorTypeList           AccessibilityCustomRotorType = 15
	AccessibilityCustomRotorTypeMisspelledWord AccessibilityCustomRotorType = 16
	AccessibilityCustomRotorTypeTable          AccessibilityCustomRotorType = 17
	AccessibilityCustomRotorTypeTextField      AccessibilityCustomRotorType = 18
	AccessibilityCustomRotorTypeUnderlinedText AccessibilityCustomRotorType = 19
	AccessibilityCustomRotorTypeVisitedLink    AccessibilityCustomRotorType = 20
)

// Keys for font attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityfontattributekey?language=objc
type AccessibilityFontAttributeKey string

const (
	AccessibilityFontFamilyKey  AccessibilityFontAttributeKey = "AXFontFamily"
	AccessibilityFontNameKey    AccessibilityFontAttributeKey = "AXFontName"
	AccessibilityFontSizeKey    AccessibilityFontAttributeKey = "AXFontSize"
	AccessibilityVisibleNameKey AccessibilityFontAttributeKey = "AXVisibleName"
)

// The name of the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynotificationname?language=objc
type AccessibilityNotificationName string

const (
	AccessibilityAnnouncementRequestedNotification   AccessibilityNotificationName = "AXAnnouncementRequested"
	AccessibilityApplicationActivatedNotification    AccessibilityNotificationName = "AXApplicationActivated"
	AccessibilityApplicationDeactivatedNotification  AccessibilityNotificationName = "AXApplicationDeactivated"
	AccessibilityApplicationHiddenNotification       AccessibilityNotificationName = "AXApplicationHidden"
	AccessibilityApplicationShownNotification        AccessibilityNotificationName = "AXApplicationShown"
	AccessibilityCreatedNotification                 AccessibilityNotificationName = "AXCreated"
	AccessibilityDrawerCreatedNotification           AccessibilityNotificationName = "AXDrawerCreated"
	AccessibilityFocusedUIElementChangedNotification AccessibilityNotificationName = "AXFocusedUIElementChanged"
	AccessibilityFocusedWindowChangedNotification    AccessibilityNotificationName = "AXFocusedWindowChanged"
	AccessibilityHelpTagCreatedNotification          AccessibilityNotificationName = "AXHelpTagCreated"
	AccessibilityLayoutChangedNotification           AccessibilityNotificationName = "AXLayoutChanged"
	AccessibilityMainWindowChangedNotification       AccessibilityNotificationName = "AXMainWindowChanged"
	AccessibilityMovedNotification                   AccessibilityNotificationName = "AXMoved"
	AccessibilityResizedNotification                 AccessibilityNotificationName = "AXResized"
	AccessibilityRowCollapsedNotification            AccessibilityNotificationName = "AXRowCollapsed"
	AccessibilityRowCountChangedNotification         AccessibilityNotificationName = "AXRowCountChanged"
	AccessibilityRowExpandedNotification             AccessibilityNotificationName = "AXRowExpanded"
	AccessibilitySelectedCellsChangedNotification    AccessibilityNotificationName = "AXSelectedCellsChanged"
	AccessibilitySelectedChildrenChangedNotification AccessibilityNotificationName = "AXSelectedChildrenChanged"
	AccessibilitySelectedChildrenMovedNotification   AccessibilityNotificationName = "AXSelectedChildrenMoved"
	AccessibilitySelectedColumnsChangedNotification  AccessibilityNotificationName = "AXSelectedColumnsChanged"
	AccessibilitySelectedRowsChangedNotification     AccessibilityNotificationName = "AXSelectedRowsChanged"
	AccessibilitySelectedTextChangedNotification     AccessibilityNotificationName = "AXSelectedTextChanged"
	AccessibilitySheetCreatedNotification            AccessibilityNotificationName = "AXSheetCreated"
	AccessibilityTitleChangedNotification            AccessibilityNotificationName = "AXTitleChanged"
	AccessibilityUIElementDestroyedNotification      AccessibilityNotificationName = "AXUIElementDestroyed"
	AccessibilityUnitsChangedNotification            AccessibilityNotificationName = "AXUnitsChanged"
	AccessibilityValueChangedNotification            AccessibilityNotificationName = "AXValueChanged"
	AccessibilityWindowCreatedNotification           AccessibilityNotificationName = "AXWindowCreated"
	AccessibilityWindowDeminiaturizedNotification    AccessibilityNotificationName = "AXWindowDeminiaturized"
	AccessibilityWindowMiniaturizedNotification      AccessibilityNotificationName = "AXWindowMiniaturized"
	AccessibilityWindowMovedNotification             AccessibilityNotificationName = "AXWindowMoved"
	AccessibilityWindowResizedNotification           AccessibilityNotificationName = "AXWindowResized"
)

// The key in the user info dictionary for a notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynotificationuserinfokey?language=objc
type AccessibilityNotificationUserInfoKey string

const (
	AccessibilityAnnouncementKey AccessibilityNotificationUserInfoKey = "AXAnnouncementKey"
	AccessibilityPriorityKey     AccessibilityNotificationUserInfoKey = "AXPriorityKey"
	AccessibilityUIElementsKey   AccessibilityNotificationUserInfoKey = "AXUIElementsKey"
)

// Values that indicate the orientation of accessibility elements, such as scroll bars and split views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityorientation?language=objc
type AccessibilityOrientation int

const (
	AccessibilityOrientationHorizontal AccessibilityOrientation = 2
	AccessibilityOrientationUnknown    AccessibilityOrientation = 0
	AccessibilityOrientationVertical   AccessibilityOrientation = 1
)

// Values that indicate the orientation of user interface elements, such as scroll bars and split views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityorientationvalue?language=objc
type AccessibilityOrientationValue string

const (
	AccessibilityHorizontalOrientationValue AccessibilityOrientationValue = "AXHorizontalOrientation"
	AccessibilityUnknownOrientationValue    AccessibilityOrientationValue = "AXUnknownOrientation"
	AccessibilityVerticalOrientationValue   AccessibilityOrientationValue = "AXVerticalOrientation"
)

// Values that describe parameterized attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityparameterizedattributename?language=objc
type AccessibilityParameterizedAttributeName string

const (
	AccessibilityAttributedStringForRangeParameterizedAttribute  AccessibilityParameterizedAttributeName = "AXAttributedStringForRange"
	AccessibilityBoundsForRangeParameterizedAttribute            AccessibilityParameterizedAttributeName = "AXBoundsForRange"
	AccessibilityCellForColumnAndRowParameterizedAttribute       AccessibilityParameterizedAttributeName = "AXCellForColumnAndRow"
	AccessibilityLayoutPointForScreenPointParameterizedAttribute AccessibilityParameterizedAttributeName = "AXLayoutPointForScreenPoint"
	AccessibilityLayoutSizeForScreenSizeParameterizedAttribute   AccessibilityParameterizedAttributeName = "AXLayoutSizeForScreenSize"
	AccessibilityLineForIndexParameterizedAttribute              AccessibilityParameterizedAttributeName = "AXLineForIndex"
	AccessibilityRTFForRangeParameterizedAttribute               AccessibilityParameterizedAttributeName = "AXRTFForRange"
	AccessibilityRangeForIndexParameterizedAttribute             AccessibilityParameterizedAttributeName = "AXRangeForIndex"
	AccessibilityRangeForLineParameterizedAttribute              AccessibilityParameterizedAttributeName = "AXRangeForLine"
	AccessibilityRangeForPositionParameterizedAttribute          AccessibilityParameterizedAttributeName = "AXRangeForPosition"
	AccessibilityScreenPointForLayoutPointParameterizedAttribute AccessibilityParameterizedAttributeName = "AXScreenPointForLayoutPoint"
	AccessibilityScreenSizeForLayoutSizeParameterizedAttribute   AccessibilityParameterizedAttributeName = "AXScreenSizeForLayoutSize"
	AccessibilityStringForRangeParameterizedAttribute            AccessibilityParameterizedAttributeName = "AXStringForRange"
	AccessibilityStyleRangeForIndexParameterizedAttribute        AccessibilityParameterizedAttributeName = "AXStyleRangeForIndex"
)

// A data type for notification priority levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityprioritylevel?language=objc
type AccessibilityPriorityLevel int

const (
	AccessibilityPriorityHigh   AccessibilityPriorityLevel = 90
	AccessibilityPriorityLow    AccessibilityPriorityLevel = 10
	AccessibilityPriorityMedium AccessibilityPriorityLevel = 50
)

// Values that describe types of objects that accessibility elements represent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrole?language=objc
type AccessibilityRole string

const (
	AccessibilityApplicationRole        AccessibilityRole = "AXApplication"
	AccessibilityBrowserRole            AccessibilityRole = "AXBrowser"
	AccessibilityBusyIndicatorRole      AccessibilityRole = "AXBusyIndicator"
	AccessibilityButtonRole             AccessibilityRole = "AXButton"
	AccessibilityCellRole               AccessibilityRole = "AXCell"
	AccessibilityCheckBoxRole           AccessibilityRole = "AXCheckBox"
	AccessibilityColorWellRole          AccessibilityRole = "AXColorWell"
	AccessibilityColumnRole             AccessibilityRole = "AXColumn"
	AccessibilityComboBoxRole           AccessibilityRole = "AXComboBox"
	AccessibilityDisclosureTriangleRole AccessibilityRole = "AXDisclosureTriangle"
	AccessibilityDrawerRole             AccessibilityRole = "AXDrawer"
	AccessibilityGridRole               AccessibilityRole = "AXGrid"
	AccessibilityGroupRole              AccessibilityRole = "AXGroup"
	AccessibilityGrowAreaRole           AccessibilityRole = "AXGrowArea"
	AccessibilityHandleRole             AccessibilityRole = "AXHandle"
	AccessibilityHelpTagRole            AccessibilityRole = "AXHelpTag"
	AccessibilityImageRole              AccessibilityRole = "AXImage"
	AccessibilityIncrementorRole        AccessibilityRole = "AXIncrementor"
	AccessibilityLayoutAreaRole         AccessibilityRole = "AXLayoutArea"
	AccessibilityLayoutItemRole         AccessibilityRole = "AXLayoutItem"
	AccessibilityLevelIndicatorRole     AccessibilityRole = "AXLevelIndicator"
	AccessibilityLinkRole               AccessibilityRole = "AXLink"
	AccessibilityListRole               AccessibilityRole = "AXList"
	AccessibilityMatteRole              AccessibilityRole = "AXMatte"
	AccessibilityMenuBarItemRole        AccessibilityRole = "AXMenuBarItem"
	AccessibilityMenuBarRole            AccessibilityRole = "AXMenuBar"
	AccessibilityMenuButtonRole         AccessibilityRole = "AXMenuButton"
	AccessibilityMenuItemRole           AccessibilityRole = "AXMenuItem"
	AccessibilityMenuRole               AccessibilityRole = "AXMenu"
	AccessibilityOutlineRole            AccessibilityRole = "AXOutline"
	AccessibilityPageRole               AccessibilityRole = "AXPage"
	AccessibilityPopUpButtonRole        AccessibilityRole = "AXPopUpButton"
	AccessibilityPopoverRole            AccessibilityRole = "AXPopover"
	AccessibilityProgressIndicatorRole  AccessibilityRole = "AXProgressIndicator"
	AccessibilityRadioButtonRole        AccessibilityRole = "AXRadioButton"
	AccessibilityRadioGroupRole         AccessibilityRole = "AXRadioGroup"
	AccessibilityRelevanceIndicatorRole AccessibilityRole = "AXRelevanceIndicator"
	AccessibilityRowRole                AccessibilityRole = "AXRow"
	AccessibilityRulerMarkerRole        AccessibilityRole = "AXRulerMarker"
	AccessibilityRulerRole              AccessibilityRole = "AXRuler"
	AccessibilityScrollAreaRole         AccessibilityRole = "AXScrollArea"
	AccessibilityScrollBarRole          AccessibilityRole = "AXScrollBar"
	AccessibilitySheetRole              AccessibilityRole = "AXSheet"
	AccessibilitySliderRole             AccessibilityRole = "AXSlider"
	AccessibilitySortButtonRole         AccessibilityRole = "AXSortButton"
	AccessibilitySplitGroupRole         AccessibilityRole = "AXSplitGroup"
	AccessibilitySplitterRole           AccessibilityRole = "AXSplitter"
	AccessibilityStaticTextRole         AccessibilityRole = "AXStaticText"
	AccessibilitySystemWideRole         AccessibilityRole = "AXSystemWide"
	AccessibilityTabGroupRole           AccessibilityRole = "AXTabGroup"
	AccessibilityTableRole              AccessibilityRole = "AXTable"
	AccessibilityTextAreaRole           AccessibilityRole = "AXTextArea"
	AccessibilityTextFieldRole          AccessibilityRole = "AXTextField"
	AccessibilityToolbarRole            AccessibilityRole = "AXToolbar"
	AccessibilityUnknownRole            AccessibilityRole = "AXUnknown"
	AccessibilityValueIndicatorRole     AccessibilityRole = "AXValueIndicator"
	AccessibilityWindowRole             AccessibilityRole = "AXWindow"
)

// Values that indicate the marker type of an accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrulermarkertype?language=objc
type AccessibilityRulerMarkerType int

const (
	AccessibilityRulerMarkerTypeIndentFirstLine AccessibilityRulerMarkerType = 7
	AccessibilityRulerMarkerTypeIndentHead      AccessibilityRulerMarkerType = 5
	AccessibilityRulerMarkerTypeIndentTail      AccessibilityRulerMarkerType = 6
	AccessibilityRulerMarkerTypeTabStopCenter   AccessibilityRulerMarkerType = 3
	AccessibilityRulerMarkerTypeTabStopDecimal  AccessibilityRulerMarkerType = 4
	AccessibilityRulerMarkerTypeTabStopLeft     AccessibilityRulerMarkerType = 1
	AccessibilityRulerMarkerTypeTabStopRight    AccessibilityRulerMarkerType = 2
	AccessibilityRulerMarkerTypeUnknown         AccessibilityRulerMarkerType = 0
)

// Values that describe ruler marker types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrulermarkertypevalue?language=objc
type AccessibilityRulerMarkerTypeValue string

const (
	AccessibilityCenterTabStopMarkerTypeValue   AccessibilityRulerMarkerTypeValue = "AXCenterTabStopMarkerType"
	AccessibilityDecimalTabStopMarkerTypeValue  AccessibilityRulerMarkerTypeValue = "AXDecimalTabStopMarkerType"
	AccessibilityFirstLineIndentMarkerTypeValue AccessibilityRulerMarkerTypeValue = "AXFirstLineIndentMarkerType"
	AccessibilityHeadIndentMarkerTypeValue      AccessibilityRulerMarkerTypeValue = "AXHeadIndentMarkerType"
	AccessibilityLeftTabStopMarkerTypeValue     AccessibilityRulerMarkerTypeValue = "AXLeftTabStopMarkerType"
	AccessibilityRightTabStopMarkerTypeValue    AccessibilityRulerMarkerTypeValue = "AXRightTabStopMarkerType"
	AccessibilityTailIndentMarkerTypeValue      AccessibilityRulerMarkerTypeValue = "AXTailIndentMarkerType"
	AccessibilityUnknownMarkerTypeValue         AccessibilityRulerMarkerTypeValue = "AXUnknownMarkerType"
)

// Values that indicate the unit values of a ruler or layout area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrulerunitvalue?language=objc
type AccessibilityRulerUnitValue string

const (
	AccessibilityCentimetersUnitValue AccessibilityRulerUnitValue = "AXCentimentersUnit"
	AccessibilityInchesUnitValue      AccessibilityRulerUnitValue = "AXInchesUnit"
	AccessibilityPicasUnitValue       AccessibilityRulerUnitValue = "AXPicasUnit"
	AccessibilityPointsUnitValue      AccessibilityRulerUnitValue = "AXPointsUnit"
	AccessibilityUnknownUnitValue     AccessibilityRulerUnitValue = "AXUnknownUnit"
)

// Values that indicate the sort direction of a column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitysortdirection?language=objc
type AccessibilitySortDirection int

const (
	AccessibilitySortDirectionAscending  AccessibilitySortDirection = 1
	AccessibilitySortDirectionDescending AccessibilitySortDirection = 2
	AccessibilitySortDirectionUnknown    AccessibilitySortDirection = 0
)

// Values that indicate the sort direction of a column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitysortdirectionvalue?language=objc
type AccessibilitySortDirectionValue string

const (
	AccessibilityAscendingSortDirectionValue  AccessibilitySortDirectionValue = "AXAscendingSortDirection"
	AccessibilityDescendingSortDirectionValue AccessibilitySortDirectionValue = "AXDescendingSortDirection"
	AccessibilityUnknownSortDirectionValue    AccessibilitySortDirectionValue = "AXUnknownSortDirection"
)

// Values that describe specialized object subtypes that accessibility elements represent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitysubrole?language=objc
type AccessibilitySubrole string

const (
	AccessibilityCloseButtonSubrole          AccessibilitySubrole = "AXCloseButton"
	AccessibilityCollectionListSubrole       AccessibilitySubrole = "AXCollectionList"
	AccessibilityContentListSubrole          AccessibilitySubrole = "AXContentList"
	AccessibilityDecrementArrowSubrole       AccessibilitySubrole = "AXDecrementArrow"
	AccessibilityDecrementPageSubrole        AccessibilitySubrole = "AXDecrementPage"
	AccessibilityDefinitionListSubrole       AccessibilitySubrole = "AXDefinitionList"
	AccessibilityDescriptionListSubrole      AccessibilitySubrole = "AXDescriptionList"
	AccessibilityDialogSubrole               AccessibilitySubrole = "AXDialog"
	AccessibilityFloatingWindowSubrole       AccessibilitySubrole = "AXFloatingWindow"
	AccessibilityFullScreenButtonSubrole     AccessibilitySubrole = "AXFullScreenButton"
	AccessibilityIncrementArrowSubrole       AccessibilitySubrole = "AXIncrementArrow"
	AccessibilityIncrementPageSubrole        AccessibilitySubrole = "AXIncrementPage"
	AccessibilityMinimizeButtonSubrole       AccessibilitySubrole = "AXMinimizeButton"
	AccessibilityOutlineRowSubrole           AccessibilitySubrole = "AXOutlineRow"
	AccessibilityRatingIndicatorSubrole      AccessibilitySubrole = "AXRatingIndicator"
	AccessibilitySearchFieldSubrole          AccessibilitySubrole = "AXSearchField"
	AccessibilitySectionListSubrole          AccessibilitySubrole = "AXSectionList"
	AccessibilitySecureTextFieldSubrole      AccessibilitySubrole = "AXSecureTextField"
	AccessibilitySortButtonSubrole           AccessibilitySubrole = "AXSortButton"
	AccessibilityStandardWindowSubrole       AccessibilitySubrole = "AXStandardWindow"
	AccessibilitySwitchSubrole               AccessibilitySubrole = "AXSwitch"
	AccessibilitySystemDialogSubrole         AccessibilitySubrole = "AXSystemDialog"
	AccessibilitySystemFloatingWindowSubrole AccessibilitySubrole = "AXSystemFloatingWindow"
	AccessibilityTabButtonSubrole            AccessibilitySubrole = "AXTabButton"
	AccessibilityTableRowSubrole             AccessibilitySubrole = "AXTableRow"
	AccessibilityTextAttachmentSubrole       AccessibilitySubrole = "AXTextAttachment"
	AccessibilityTextLinkSubrole             AccessibilitySubrole = "AXTextLink"
	AccessibilityTimelineSubrole             AccessibilitySubrole = "AXTimeline"
	AccessibilityToggleSubrole               AccessibilitySubrole = "AXToggle"
	AccessibilityToolbarButtonSubrole        AccessibilitySubrole = "AXToolbarButton"
	AccessibilityUnknownSubrole              AccessibilitySubrole = "AXUnknown"
	AccessibilityZoomButtonSubrole           AccessibilitySubrole = "AXZoomButton"
)

// Values that indicate the unit values of a ruler or layout area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityunits?language=objc
type AccessibilityUnits int

const (
	AccessibilityUnitsCentimeters AccessibilityUnits = 2
	AccessibilityUnitsInches      AccessibilityUnits = 1
	AccessibilityUnitsPicas       AccessibilityUnits = 4
	AccessibilityUnitsPoints      AccessibilityUnits = 3
	AccessibilityUnitsUnknown     AccessibilityUnits = 0
)

// The NSAlert class defines the alert styles used by the alertStyle property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalertstyle?language=objc
type AlertStyle uint

const (
	AlertStyleCritical      AlertStyle = 2
	AlertStyleInformational AlertStyle = 1
	AlertStyleWarning       AlertStyle = 0
	CriticalAlertStyle      AlertStyle = 2
	InformationalAlertStyle AlertStyle = 1
	WarningAlertStyle       AlertStyle = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertykey?language=objc
type AnimatablePropertyKey string

const (
	AnimationTriggerOrderIn  AnimatablePropertyKey = "NSAnimationTriggerOrderIn"
	AnimationTriggerOrderOut AnimatablePropertyKey = "NSAnimationTriggerOrderOut"
)

// These constants indicate the blocking mode of an NSAnimation object when it is running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationblockingmode?language=objc
type AnimationBlockingMode uint

const (
	AnimationBlocking            AnimationBlockingMode = 0
	AnimationNonblocking         AnimationBlockingMode = 1
	AnimationNonblockingThreaded AnimationBlockingMode = 2
)

// These constants describe the curve of an animation—that is, the relative speed of an animation from start to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcurve?language=objc
type AnimationCurve uint

const (
	AnimationEaseIn    AnimationCurve = 1
	AnimationEaseInOut AnimationCurve = 0
	AnimationEaseOut   AnimationCurve = 2
	AnimationLinear    AnimationCurve = 3
)

// The type for standard system animation effects, which include both display and sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationeffect?language=objc
type AnimationEffect uint

const (
	AnimationEffectDisappearingItemDefault AnimationEffect = 0
	AnimationEffectPoof                    AnimationEffect = 10
)

// The animation progress, as a floating-point number between 0.0 and 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationprogress?language=objc
type AnimationProgress float64

// Constants for determining which version of AppKit is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappkitversion?language=objc
type AppKitVersion float64

const (
	AppKitVersionNumber                                 AppKitVersion = 2113.600098
	AppKitVersionNumber10_0                             AppKitVersion = 577.000000
	AppKitVersionNumber10_1                             AppKitVersion = 620.000000
	AppKitVersionNumber10_10                            AppKitVersion = 1343.000000
	AppKitVersionNumber10_10_2                          AppKitVersion = 1344.000000
	AppKitVersionNumber10_10_3                          AppKitVersion = 1347.000000
	AppKitVersionNumber10_10_4                          AppKitVersion = 1348.000000
	AppKitVersionNumber10_10_5                          AppKitVersion = 1348.000000
	AppKitVersionNumber10_10_Max                        AppKitVersion = 1349.000000
	AppKitVersionNumber10_11                            AppKitVersion = 1404.000000
	AppKitVersionNumber10_11_1                          AppKitVersion = 1404.130005
	AppKitVersionNumber10_11_2                          AppKitVersion = 1404.339966
	AppKitVersionNumber10_11_3                          AppKitVersion = 1404.339966
	AppKitVersionNumber10_12                            AppKitVersion = 1504.000000
	AppKitVersionNumber10_12_1                          AppKitVersion = 1504.599976
	AppKitVersionNumber10_12_2                          AppKitVersion = 1504.760010
	AppKitVersionNumber10_13                            AppKitVersion = 1561.000000
	AppKitVersionNumber10_13_1                          AppKitVersion = 1561.099976
	AppKitVersionNumber10_13_2                          AppKitVersion = 1561.199951
	AppKitVersionNumber10_13_4                          AppKitVersion = 1561.400024
	AppKitVersionNumber10_14                            AppKitVersion = 1671.000000
	AppKitVersionNumber10_14_1                          AppKitVersion = 1671.099976
	AppKitVersionNumber10_14_2                          AppKitVersion = 1671.199951
	AppKitVersionNumber10_14_3                          AppKitVersion = 1671.300049
	AppKitVersionNumber10_14_4                          AppKitVersion = 1671.400024
	AppKitVersionNumber10_14_5                          AppKitVersion = 1671.500000
	AppKitVersionNumber10_15                            AppKitVersion = 1894.000000
	AppKitVersionNumber10_15_1                          AppKitVersion = 1894.099976
	AppKitVersionNumber10_15_2                          AppKitVersion = 1894.199951
	AppKitVersionNumber10_15_3                          AppKitVersion = 1894.300049
	AppKitVersionNumber10_15_4                          AppKitVersion = 1894.400024
	AppKitVersionNumber10_15_5                          AppKitVersion = 1894.500000
	AppKitVersionNumber10_15_6                          AppKitVersion = 1894.599976
	AppKitVersionNumber10_2                             AppKitVersion = 663.000000
	AppKitVersionNumber10_2_3                           AppKitVersion = 663.599976
	AppKitVersionNumber10_3                             AppKitVersion = 743.000000
	AppKitVersionNumber10_3_2                           AppKitVersion = 743.140015
	AppKitVersionNumber10_3_3                           AppKitVersion = 743.200012
	AppKitVersionNumber10_3_5                           AppKitVersion = 743.239990
	AppKitVersionNumber10_3_7                           AppKitVersion = 743.330017
	AppKitVersionNumber10_3_9                           AppKitVersion = 743.359985
	AppKitVersionNumber10_4                             AppKitVersion = 824.000000
	AppKitVersionNumber10_4_1                           AppKitVersion = 824.099976
	AppKitVersionNumber10_4_3                           AppKitVersion = 824.229980
	AppKitVersionNumber10_4_4                           AppKitVersion = 824.330017
	AppKitVersionNumber10_4_7                           AppKitVersion = 824.409973
	AppKitVersionNumber10_5                             AppKitVersion = 949.000000
	AppKitVersionNumber10_5_2                           AppKitVersion = 949.270020
	AppKitVersionNumber10_5_3                           AppKitVersion = 949.330017
	AppKitVersionNumber10_6                             AppKitVersion = 1038.000000
	AppKitVersionNumber10_7                             AppKitVersion = 1138.000000
	AppKitVersionNumber10_7_2                           AppKitVersion = 1138.229980
	AppKitVersionNumber10_7_3                           AppKitVersion = 1138.319946
	AppKitVersionNumber10_7_4                           AppKitVersion = 1138.469971
	AppKitVersionNumber10_8                             AppKitVersion = 1187.000000
	AppKitVersionNumber10_9                             AppKitVersion = 1265.000000
	AppKitVersionNumber11_0                             AppKitVersion = 2022.000000
	AppKitVersionNumber11_1                             AppKitVersion = 2022.199951
	AppKitVersionNumber11_2                             AppKitVersion = 2022.300049
	AppKitVersionNumber11_3                             AppKitVersion = 2022.400024
	AppKitVersionNumber11_4                             AppKitVersion = 2022.500000
	AppKitVersionNumberWithColumnResizingBrowser        AppKitVersion = 685.000000
	AppKitVersionNumberWithContinuousScrollingBrowser   AppKitVersion = 680.000000
	AppKitVersionNumberWithCursorSizeSupport            AppKitVersion = 682.000000
	AppKitVersionNumberWithCustomSheetPosition          AppKitVersion = 686.000000
	AppKitVersionNumberWithDeferredWindowDisplaySupport AppKitVersion = 1019.000000
	AppKitVersionNumberWithDirectionalTabs              AppKitVersion = 631.000000
	AppKitVersionNumberWithDockTilePlugInSupport        AppKitVersion = 1001.000000
	AppKitVersionNumberWithPatternColorLeakFix          AppKitVersion = 641.000000
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancename?language=objc
type AppearanceName string

const (
	AppearanceNameAccessibilityHighContrastAqua         AppearanceName = "NSAppearanceNameAccessibilityAqua"
	AppearanceNameAccessibilityHighContrastDarkAqua     AppearanceName = "NSAppearanceNameAccessibilityDarkAqua"
	AppearanceNameAccessibilityHighContrastVibrantDark  AppearanceName = "NSAppearanceNameAccessibilityVibrantDark"
	AppearanceNameAccessibilityHighContrastVibrantLight AppearanceName = "NSAppearanceNameAccessibilityVibrantLight"
	AppearanceNameAqua                                  AppearanceName = "NSAppearanceNameAqua"
	AppearanceNameDarkAqua                              AppearanceName = "NSAppearanceNameDarkAqua"
	AppearanceNameLightContent                          AppearanceName = "NSAppearanceNameLightContent"
	AppearanceNameVibrantDark                           AppearanceName = "NSAppearanceNameVibrantDark"
	AppearanceNameVibrantLight                          AppearanceName = "NSAppearanceNameVibrantLight"
)

// The following flags are for activateWithOptions:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationactivationoptions?language=objc
type ApplicationActivationOptions uint

const (
	ApplicationActivateAllWindows        ApplicationActivationOptions = 1
	ApplicationActivateIgnoringOtherApps ApplicationActivationOptions = 2
)

// Activation policies (used by activationPolicy) that control whether and how an app may be activated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationactivationpolicy?language=objc
type ApplicationActivationPolicy int

const (
	ApplicationActivationPolicyAccessory  ApplicationActivationPolicy = 1
	ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 2
	ApplicationActivationPolicyRegular    ApplicationActivationPolicy = 0
)

// Constants that indicate whether a copy or print operation was successful, was canceled, or failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegatereply?language=objc
type ApplicationDelegateReply uint

const (
	ApplicationDelegateReplyCancel  ApplicationDelegateReply = 1
	ApplicationDelegateReplyFailure ApplicationDelegateReply = 2
	ApplicationDelegateReplySuccess ApplicationDelegateReply = 0
)

// This constant indicates whether at least part of any window owned by this app is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationocclusionstate?language=objc
type ApplicationOcclusionState uint

const (
	ApplicationOcclusionStateVisible ApplicationOcclusionState = 2
)

// Constants that control the presentation of the app, typically for fullscreen apps such as games or kiosks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationpresentationoptions?language=objc
type ApplicationPresentationOptions uint

const (
	ApplicationPresentationAutoHideDock                    ApplicationPresentationOptions = 1
	ApplicationPresentationAutoHideMenuBar                 ApplicationPresentationOptions = 4
	ApplicationPresentationAutoHideToolbar                 ApplicationPresentationOptions = 2048
	ApplicationPresentationDefault                         ApplicationPresentationOptions = 0
	ApplicationPresentationDisableAppleMenu                ApplicationPresentationOptions = 16
	ApplicationPresentationDisableCursorLocationAssistance ApplicationPresentationOptions = 4096
	ApplicationPresentationDisableForceQuit                ApplicationPresentationOptions = 64
	ApplicationPresentationDisableHideApplication          ApplicationPresentationOptions = 256
	ApplicationPresentationDisableMenuBarTransparency      ApplicationPresentationOptions = 512
	ApplicationPresentationDisableProcessSwitching         ApplicationPresentationOptions = 32
	ApplicationPresentationDisableSessionTermination       ApplicationPresentationOptions = 128
	ApplicationPresentationFullScreen                      ApplicationPresentationOptions = 1024
	ApplicationPresentationHideDock                        ApplicationPresentationOptions = 2
	ApplicationPresentationHideMenuBar                     ApplicationPresentationOptions = 8
)

// Constants that indicate the outcome of a print request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationprintreply?language=objc
type ApplicationPrintReply uint

const (
	PrintingCancelled  ApplicationPrintReply = 0
	PrintingFailure    ApplicationPrintReply = 3
	PrintingReplyLater ApplicationPrintReply = 2
	PrintingSuccess    ApplicationPrintReply = 1
)

// Constants that determine whether an app should terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationterminatereply?language=objc
type ApplicationTerminateReply uint

const (
	TerminateCancel ApplicationTerminateReply = 0
	TerminateLater  ApplicationTerminateReply = 2
	TerminateNow    ApplicationTerminateReply = 1
)

// Attributes that apply to a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsattributedstringdocumentattributekey?language=objc
type AttributedStringDocumentAttributeKey string

const (
	AppearanceDocumentAttribute         AttributedStringDocumentAttributeKey = "NSAppearanceDocumentAttribute"
	AuthorDocumentAttribute             AttributedStringDocumentAttributeKey = "NSAuthorDocumentAttribute"
	BackgroundColorDocumentAttribute    AttributedStringDocumentAttributeKey = "BackgroundColor"
	BottomMarginDocumentAttribute       AttributedStringDocumentAttributeKey = "BottomMargin"
	CategoryDocumentAttribute           AttributedStringDocumentAttributeKey = "NSCategoryDocumentAttribute"
	CharacterEncodingDocumentAttribute  AttributedStringDocumentAttributeKey = "CharacterEncoding"
	CocoaVersionDocumentAttribute       AttributedStringDocumentAttributeKey = "CocoaRTFVersion"
	CommentDocumentAttribute            AttributedStringDocumentAttributeKey = "NSCommentDocumentAttribute"
	CompanyDocumentAttribute            AttributedStringDocumentAttributeKey = "NSCompanyDocumentAttribute"
	ConvertedDocumentAttribute          AttributedStringDocumentAttributeKey = "Converted"
	CopyrightDocumentAttribute          AttributedStringDocumentAttributeKey = "NSCopyrightDocumentAttribute"
	CreationTimeDocumentAttribute       AttributedStringDocumentAttributeKey = "NSCreationTimeDocumentAttribute"
	DefaultAttributesDocumentAttribute  AttributedStringDocumentAttributeKey = "DefaultAttributes"
	DefaultTabIntervalDocumentAttribute AttributedStringDocumentAttributeKey = "DefaultTabInterval"
	DocumentTypeDocumentAttribute       AttributedStringDocumentAttributeKey = "DocumentType"
	EditorDocumentAttribute             AttributedStringDocumentAttributeKey = "NSEditorDocumentAttribute"
	ExcludedElementsDocumentAttribute   AttributedStringDocumentAttributeKey = "ExcludedElements"
	FileTypeDocumentAttribute           AttributedStringDocumentAttributeKey = "UTI"
	HyphenationFactorDocumentAttribute  AttributedStringDocumentAttributeKey = "HyphenationFactor"
	KeywordsDocumentAttribute           AttributedStringDocumentAttributeKey = "NSKeywordsDocumentAttribute"
	LeftMarginDocumentAttribute         AttributedStringDocumentAttributeKey = "LeftMargin"
	ManagerDocumentAttribute            AttributedStringDocumentAttributeKey = "NSManagerDocumentAttribute"
	ModificationTimeDocumentAttribute   AttributedStringDocumentAttributeKey = "NSModificationTimeDocumentAttribute"
	PaperSizeDocumentAttribute          AttributedStringDocumentAttributeKey = "PaperSize"
	PrefixSpacesDocumentAttribute       AttributedStringDocumentAttributeKey = "PrefixSpaces"
	ReadOnlyDocumentAttribute           AttributedStringDocumentAttributeKey = "ReadOnly"
	RightMarginDocumentAttribute        AttributedStringDocumentAttributeKey = "RightMargin"
	SourceTextScalingDocumentAttribute  AttributedStringDocumentAttributeKey = "SourceTextScaling"
	SubjectDocumentAttribute            AttributedStringDocumentAttributeKey = "NSSubjectDocumentAttribute"
	TextEncodingNameDocumentAttribute   AttributedStringDocumentAttributeKey = "TextEncodingName"
	TextLayoutSectionsAttribute         AttributedStringDocumentAttributeKey = "NSTextLayoutSectionsAttribute"
	TextScalingDocumentAttribute        AttributedStringDocumentAttributeKey = "TextScaling"
	TitleDocumentAttribute              AttributedStringDocumentAttributeKey = "NSTitleDocumentAttribute"
	TopMarginDocumentAttribute          AttributedStringDocumentAttributeKey = "TopMargin"
	ViewModeDocumentAttribute           AttributedStringDocumentAttributeKey = "ViewMode"
	ViewSizeDocumentAttribute           AttributedStringDocumentAttributeKey = "ViewSize"
	ViewZoomDocumentAttribute           AttributedStringDocumentAttributeKey = "ViewZoom"
)

// Options for importing documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsattributedstringdocumentreadingoptionkey?language=objc
type AttributedStringDocumentReadingOptionKey string

const (
	BaseURLDocumentOption                 AttributedStringDocumentReadingOptionKey = "BaseURL"
	CharacterEncodingDocumentOption       AttributedStringDocumentReadingOptionKey = "CharacterEncoding"
	DefaultAttributesDocumentOption       AttributedStringDocumentReadingOptionKey = "DefaultAttributes"
	DocumentTypeDocumentOption            AttributedStringDocumentReadingOptionKey = "DocumentType"
	FileTypeDocumentOption                AttributedStringDocumentReadingOptionKey = "UTI"
	SourceTextScalingDocumentOption       AttributedStringDocumentReadingOptionKey = "SourceTextScaling"
	TargetTextScalingDocumentOption       AttributedStringDocumentReadingOptionKey = "TargetTextScaling"
	TextEncodingNameDocumentOption        AttributedStringDocumentReadingOptionKey = "TextEncodingName"
	TextSizeMultiplierDocumentOption      AttributedStringDocumentReadingOptionKey = "TextSizeMultiplier"
	TimeoutDocumentOption                 AttributedStringDocumentReadingOptionKey = "Timeout"
	WebPreferencesDocumentOption          AttributedStringDocumentReadingOptionKey = "WebPreferences"
	WebResourceLoadDelegateDocumentOption AttributedStringDocumentReadingOptionKey = "WebResourceLoadDelegate"
)

// Constants for the document type document attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsattributedstringdocumenttype?language=objc
type AttributedStringDocumentType string

const (
	DocFormatTextDocumentType     AttributedStringDocumentType = "NSDocFormat"
	HTMLTextDocumentType          AttributedStringDocumentType = "NSHTML"
	MacSimpleTextDocumentType     AttributedStringDocumentType = "NSMacSimpleText"
	OfficeOpenXMLTextDocumentType AttributedStringDocumentType = "NSOfficeOpenXML"
	OpenDocumentTextDocumentType  AttributedStringDocumentType = "NSOpenDocument"
	PlainTextDocumentType         AttributedStringDocumentType = "NSPlainText"
	RTFDTextDocumentType          AttributedStringDocumentType = "NSRTFD"
	RTFTextDocumentType           AttributedStringDocumentType = "NSRTF"
	WebArchiveTextDocumentType    AttributedStringDocumentType = "NSWebArchive"
	WordMLTextDocumentType        AttributedStringDocumentType = "NSWordML"
)

// Constants that specify the autoresizing behaviors for views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsautoresizingmaskoptions?language=objc
type AutoresizingMaskOptions uint

const (
	ViewHeightSizable AutoresizingMaskOptions = 16
	ViewMaxXMargin    AutoresizingMaskOptions = 4
	ViewMaxYMargin    AutoresizingMaskOptions = 32
	ViewMinXMargin    AutoresizingMaskOptions = 1
	ViewMinYMargin    AutoresizingMaskOptions = 8
	ViewNotSizable    AutoresizingMaskOptions = 0
	ViewWidthSizable  AutoresizingMaskOptions = 2
)

// Background styles to apply to a view’s cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbackgroundstyle?language=objc
type BackgroundStyle int

const (
	BackgroundStyleDark       BackgroundStyle = 1
	BackgroundStyleEmphasized BackgroundStyle = 1
	BackgroundStyleLight      BackgroundStyle = 0
	BackgroundStyleLowered    BackgroundStyle = 3
	BackgroundStyleNormal     BackgroundStyle = 0
	BackgroundStyleRaised     BackgroundStyle = 2
)

// Constants that specify how the window device buffers the drawing done in a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbackingstoretype?language=objc
type BackingStoreType uint

const (
	BackingStoreBuffered    BackingStoreType = 2
	BackingStoreNonretained BackingStoreType = 1
	BackingStoreRetained    BackingStoreType = 0
)

// Bezel styles used by the bezelStyle property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezelstyle?language=objc
type BezelStyle uint

const (
	BezelStyleCircular          BezelStyle = 7
	BezelStyleDisclosure        BezelStyle = 5
	BezelStyleHelpButton        BezelStyle = 9
	BezelStyleInline            BezelStyle = 15
	BezelStyleRecessed          BezelStyle = 13
	BezelStyleRegularSquare     BezelStyle = 2
	BezelStyleRoundRect         BezelStyle = 12
	BezelStyleRounded           BezelStyle = 1
	BezelStyleRoundedDisclosure BezelStyle = 14
	BezelStyleShadowlessSquare  BezelStyle = 6
	BezelStyleSmallSquare       BezelStyle = 10
	BezelStyleTexturedRounded   BezelStyle = 11
	BezelStyleTexturedSquare    BezelStyle = 8
	CircularBezelStyle          BezelStyle = 7
	DisclosureBezelStyle        BezelStyle = 5
	HelpButtonBezelStyle        BezelStyle = 9
	InlineBezelStyle            BezelStyle = 15
	RecessedBezelStyle          BezelStyle = 13
	RegularSquareBezelStyle     BezelStyle = 2
	RoundRectBezelStyle         BezelStyle = 12
	RoundedBezelStyle           BezelStyle = 1
	RoundedDisclosureBezelStyle BezelStyle = 14
	ShadowlessSquareBezelStyle  BezelStyle = 6
	SmallIconButtonBezelStyle   BezelStyle = 2
	SmallSquareBezelStyle       BezelStyle = 10
	TexturedRoundedBezelStyle   BezelStyle = 11
	TexturedSquareBezelStyle    BezelStyle = 8
	ThickSquareBezelStyle       BezelStyle = 3
	ThickerSquareBezelStyle     BezelStyle = 4
)

// Constants that specify basic path element commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpathelement?language=objc
type BezierPathElement uint

const (
	BezierPathElementClosePath BezierPathElement = 3
	BezierPathElementCurveTo   BezierPathElement = 2
	BezierPathElementLineTo    BezierPathElement = 1
	BezierPathElementMoveTo    BezierPathElement = 0
	ClosePathBezierPathElement BezierPathElement = 3
	CurveToBezierPathElement   BezierPathElement = 2
	LineToBezierPathElement    BezierPathElement = 1
	MoveToBezierPathElement    BezierPathElement = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindinginfokey?language=objc
type BindingInfoKey string

const (
	ObservedKeyPathKey BindingInfoKey = "NSObservedKeyPath"
	ObservedObjectKey  BindingInfoKey = "NSObservedObject"
	OptionsKey         BindingInfoKey = "NSOptions"
)

// Values that specify a binding for certain methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingname?language=objc
type BindingName string

const (
	AlignmentBinding                        BindingName = "alignment"
	AlternateImageBinding                   BindingName = "alternateImage"
	AlternateTitleBinding                   BindingName = "alternateTitle"
	AnimateBinding                          BindingName = "animate"
	AnimationDelayBinding                   BindingName = "animationDelay"
	ArgumentBinding                         BindingName = "argument"
	AttributedStringBinding                 BindingName = "attributedString"
	ContentArrayBinding                     BindingName = "contentArray"
	ContentArrayForMultipleSelectionBinding BindingName = "contentArrayForMultipleSelection"
	ContentBinding                          BindingName = "content"
	ContentDictionaryBinding                BindingName = "contentDictionary"
	ContentHeightBinding                    BindingName = "contentHeight"
	ContentObjectBinding                    BindingName = "contentObject"
	ContentObjectsBinding                   BindingName = "contentObjects"
	ContentSetBinding                       BindingName = "contentSet"
	ContentValuesBinding                    BindingName = "contentValues"
	ContentWidthBinding                     BindingName = "contentWidth"
	CriticalValueBinding                    BindingName = "criticalValue"
	DataBinding                             BindingName = "data"
	DisplayPatternTitleBinding              BindingName = "displayPatternTitle"
	DisplayPatternValueBinding              BindingName = "displayPatternValue"
	DocumentEditedBinding                   BindingName = "documentEdited"
	DoubleClickArgumentBinding              BindingName = "doubleClickArgument"
	DoubleClickTargetBinding                BindingName = "doubleClickTarget"
	EditableBinding                         BindingName = "editable"
	EnabledBinding                          BindingName = "enabled"
	ExcludedKeysBinding                     BindingName = "excludedKeys"
	FilterPredicateBinding                  BindingName = "filterPredicate"
	FontBinding                             BindingName = "font"
	FontBoldBinding                         BindingName = "fontBold"
	FontFamilyNameBinding                   BindingName = "fontFamilyName"
	FontItalicBinding                       BindingName = "fontItalic"
	FontNameBinding                         BindingName = "fontName"
	FontSizeBinding                         BindingName = "fontSize"
	HeaderTitleBinding                      BindingName = "headerTitle"
	HiddenBinding                           BindingName = "hidden"
	ImageBinding                            BindingName = "image"
	IncludedKeysBinding                     BindingName = "includedKeys"
	InitialKeyBinding                       BindingName = "initialKey"
	InitialValueBinding                     BindingName = "initialValue"
	IsIndeterminateBinding                  BindingName = "isIndeterminate"
	LabelBinding                            BindingName = "label"
	LocalizedKeyDictionaryBinding           BindingName = "localizedKeyDictionary"
	ManagedObjectContextBinding             BindingName = "managedObjectContext"
	MaxValueBinding                         BindingName = "maxValue"
	MaxWidthBinding                         BindingName = "maxWidth"
	MaximumRecentsBinding                   BindingName = "maximumRecents"
	MinValueBinding                         BindingName = "minValue"
	MinWidthBinding                         BindingName = "minWidth"
	MixedStateImageBinding                  BindingName = "mixedStateImage"
	OffStateImageBinding                    BindingName = "offStateImage"
	OnStateImageBinding                     BindingName = "onStateImage"
	PositioningRectBinding                  BindingName = "positioningRect"
	PredicateBinding                        BindingName = "predicate"
	RecentSearchesBinding                   BindingName = "recentSearches"
	RepresentedFilenameBinding              BindingName = "representedFilename"
	RowHeightBinding                        BindingName = "rowHeight"
	SelectedIdentifierBinding               BindingName = "selectedIdentifier"
	SelectedIndexBinding                    BindingName = "selectedIndex"
	SelectedLabelBinding                    BindingName = "selectedLabel"
	SelectedObjectBinding                   BindingName = "selectedObject"
	SelectedObjectsBinding                  BindingName = "selectedObjects"
	SelectedTagBinding                      BindingName = "selectedTag"
	SelectedValueBinding                    BindingName = "selectedValue"
	SelectedValuesBinding                   BindingName = "selectedValues"
	SelectionIndexPathsBinding              BindingName = "selectionIndexPaths"
	SelectionIndexesBinding                 BindingName = "selectionIndexes"
	SortDescriptorsBinding                  BindingName = "sortDescriptors"
	TargetBinding                           BindingName = "target"
	TextColorBinding                        BindingName = "textColor"
	TitleBinding                            BindingName = "title"
	ToolTipBinding                          BindingName = "toolTip"
	TransparentBinding                      BindingName = "transparent"
	ValueBinding                            BindingName = "value"
	ValuePathBinding                        BindingName = "valuePath"
	ValueURLBinding                         BindingName = "valueURL"
	VisibleBinding                          BindingName = "visible"
	WarningValueBinding                     BindingName = "warningValue"
	WidthBinding                            BindingName = "width"
)

// Values that are used as keys in the options dictionary passed to the bind:toObject:withKeyPath:options: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingoption?language=objc
type BindingOption string

const (
	AllowsEditingMultipleValuesSelectionBindingOption BindingOption = "NSAllowsEditingMultipleValuesSelection"
	AllowsNullArgumentBindingOption                   BindingOption = "NSAllowsNullArgument"
	AlwaysPresentsApplicationModalAlertsBindingOption BindingOption = "NSAlwaysPresentsApplicationModalAlerts"
	ConditionallySetsEditableBindingOption            BindingOption = "NSConditionallySetsEditable"
	ConditionallySetsEnabledBindingOption             BindingOption = "NSConditionallySetsEnabled"
	ConditionallySetsHiddenBindingOption              BindingOption = "NSConditionallySetsHidden"
	ContentPlacementTagBindingOption                  BindingOption = "NSContentPlacementTag"
	ContinuouslyUpdatesValueBindingOption             BindingOption = "NSContinuouslyUpdatesValue"
	CreatesSortDescriptorBindingOption                BindingOption = "NSCreatesSortDescriptor"
	DeletesObjectsOnRemoveBindingsOption              BindingOption = "NSDeletesObjectsOnRemove"
	DisplayNameBindingOption                          BindingOption = "NSDisplayName"
	DisplayPatternBindingOption                       BindingOption = "NSDisplayPattern"
	HandlesContentAsCompoundValueBindingOption        BindingOption = "NSHandlesContentAsCompoundValue"
	InsertsNullPlaceholderBindingOption               BindingOption = "NSInsertsNullPlaceholder"
	InvokesSeparatelyWithArrayObjectsBindingOption    BindingOption = "NSInvokesSeparatelyWithArrayObjects"
	MultipleValuesPlaceholderBindingOption            BindingOption = "NSMultipleValuesPlaceholder"
	NoSelectionPlaceholderBindingOption               BindingOption = "NSNoSelectionPlaceholder"
	NotApplicablePlaceholderBindingOption             BindingOption = "NSNotApplicablePlaceholder"
	NullPlaceholderBindingOption                      BindingOption = "NSNullPlaceholder"
	PredicateFormatBindingOption                      BindingOption = "NSPredicateFormat"
	RaisesForNotApplicableKeysBindingOption           BindingOption = "NSRaisesForNotApplicableKeys"
	SelectorNameBindingOption                         BindingOption = "NSSelectorName"
	SelectsAllWhenSettingContentBindingOption         BindingOption = "NSSelectsAllWhenSettingContent"
	ValidatesImmediatelyBindingOption                 BindingOption = "NSValidatesImmediately"
	ValueTransformerBindingOption                     BindingOption = "NSValueTransformer"
	ValueTransformerNameBindingOption                 BindingOption = "NSValueTransformerName"
)

// Constants that represent bitmap component formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapformat?language=objc
type BitmapFormat uint

const (
	AlphaFirstBitmapFormat               BitmapFormat = 1
	AlphaNonpremultipliedBitmapFormat    BitmapFormat = 2
	BitmapFormatAlphaFirst               BitmapFormat = 1
	BitmapFormatAlphaNonpremultiplied    BitmapFormat = 2
	BitmapFormatFloatingPointSamples     BitmapFormat = 4
	BitmapFormatSixteenBitBigEndian      BitmapFormat = 1024
	BitmapFormatSixteenBitLittleEndian   BitmapFormat = 256
	BitmapFormatThirtyTwoBitBigEndian    BitmapFormat = 2048
	BitmapFormatThirtyTwoBitLittleEndian BitmapFormat = 512
	FloatingPointSamplesBitmapFormat     BitmapFormat = 4
)

// Constants that specify bitmap file types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagefiletype?language=objc
type BitmapImageFileType uint

const (
	BMPFileType                 BitmapImageFileType = 1
	BitmapImageFileTypeBMP      BitmapImageFileType = 1
	BitmapImageFileTypeGIF      BitmapImageFileType = 2
	BitmapImageFileTypeJPEG     BitmapImageFileType = 3
	BitmapImageFileTypeJPEG2000 BitmapImageFileType = 5
	BitmapImageFileTypePNG      BitmapImageFileType = 4
	BitmapImageFileTypeTIFF     BitmapImageFileType = 0
	GIFFileType                 BitmapImageFileType = 2
	JPEG2000FileType            BitmapImageFileType = 5
	JPEGFileType                BitmapImageFileType = 3
	PNGFileType                 BitmapImageFileType = 4
	TIFFFileType                BitmapImageFileType = 0
)

// Constants that identify bitmap image representation properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagereppropertykey?language=objc
type BitmapImageRepPropertyKey string

const (
	ImageColorSyncProfileData    BitmapImageRepPropertyKey = "NSImageColorSyncProfileData"
	ImageCompressionFactor       BitmapImageRepPropertyKey = "NSImageCompressionFactor"
	ImageCompressionMethod       BitmapImageRepPropertyKey = "NSImageCompressionMethod"
	ImageCurrentFrame            BitmapImageRepPropertyKey = "NSImageCurrentFrame"
	ImageCurrentFrameDuration    BitmapImageRepPropertyKey = "NSImageCurrentFrameDuration"
	ImageDitherTransparency      BitmapImageRepPropertyKey = "NSImageDitherTransparency"
	ImageEXIFData                BitmapImageRepPropertyKey = "NSImageEXIFData"
	ImageFallbackBackgroundColor BitmapImageRepPropertyKey = "NSImageFallbackBackgroundColor"
	ImageFrameCount              BitmapImageRepPropertyKey = "NSImageFrameCount"
	ImageGamma                   BitmapImageRepPropertyKey = "NSImageGamma"
	ImageIPTCData                BitmapImageRepPropertyKey = "NSImageIPTCData"
	ImageInterlaced              BitmapImageRepPropertyKey = "NSImageInterlaced"
	ImageLoopCount               BitmapImageRepPropertyKey = "NSImageLoopCount"
	ImageProgressive             BitmapImageRepPropertyKey = "NSImageProgressive"
	ImageRGBColorTable           BitmapImageRepPropertyKey = "NSImageRGBColorTable"
)

// These constants specify the type of a view’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbordertype?language=objc
type BorderType uint

const (
	BezelBorder  BorderType = 2
	GrooveBorder BorderType = 3
	LineBorder   BorderType = 1
	NoBorder     BorderType = 0
)

// These constants and data type identifies box types, which, in conjunction with a box's border type, define the appearance of the box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsboxtype?language=objc
type BoxType uint

const (
	BoxCustom    BoxType = 4
	BoxOldStyle  BoxType = 3
	BoxPrimary   BoxType = 0
	BoxSecondary BoxType = 1
	BoxSeparator BoxType = 2
)

// Types of browser column resizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercolumnresizingtype?language=objc
type BrowserColumnResizingType uint

const (
	BrowserAutoColumnResizing BrowserColumnResizingType = 1
	BrowserNoColumnResizing   BrowserColumnResizingType = 0
	BrowserUserColumnResizing BrowserColumnResizingType = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercolumnsautosavename?language=objc
type BrowserColumnsAutosaveName string

// The type used to specify the drop type of a drag-and-drop operation. See browser:objectValueForItem: for more information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdropoperation?language=objc
type BrowserDropOperation uint

const (
	BrowserDropAbove BrowserDropOperation = 1
	BrowserDropOn    BrowserDropOperation = 0
)

// Button types that you can specify using setButtonType:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontype?language=objc
type ButtonType uint

const (
	AcceleratorButton               ButtonType = 8
	ButtonTypeAccelerator           ButtonType = 8
	ButtonTypeMomentaryChange       ButtonType = 5
	ButtonTypeMomentaryLight        ButtonType = 0
	ButtonTypeMomentaryPushIn       ButtonType = 7
	ButtonTypeMultiLevelAccelerator ButtonType = 9
	ButtonTypeOnOff                 ButtonType = 6
	ButtonTypePushOnPushOff         ButtonType = 1
	ButtonTypeRadio                 ButtonType = 4
	ButtonTypeSwitch                ButtonType = 3
	ButtonTypeToggle                ButtonType = 2
	MomentaryChangeButton           ButtonType = 5
	MomentaryLight                  ButtonType = 7
	MomentaryLightButton            ButtonType = 0
	MomentaryPushButton             ButtonType = 0
	MomentaryPushInButton           ButtonType = 7
	MultiLevelAcceleratorButton     ButtonType = 9
	OnOffButton                     ButtonType = 6
	PushOnPushOffButton             ButtonType = 1
	RadioButton                     ButtonType = 4
	SwitchButton                    ButtonType = 3
	ToggleButton                    ButtonType = 2
)

// Constants for specifying how a button behaves when pressed and how it displays its state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscellattribute?language=objc
type CellAttribute uint

const (
	CellAllowsMixedState       CellAttribute = 16
	CellChangesContents        CellAttribute = 14
	CellDisabled               CellAttribute = 0
	CellEditable               CellAttribute = 3
	CellHasImageHorizontal     CellAttribute = 12
	CellHasImageOnLeftOrBottom CellAttribute = 13
	CellHasOverlappingImage    CellAttribute = 11
	CellHighlighted            CellAttribute = 5
	CellIsBordered             CellAttribute = 10
	CellIsInsetButton          CellAttribute = 15
	CellLightsByBackground     CellAttribute = 9
	CellLightsByContents       CellAttribute = 6
	CellLightsByGray           CellAttribute = 7
	CellState                  CellAttribute = 1
	ChangeBackgroundCell       CellAttribute = 8
	ChangeGrayCell             CellAttribute = 4
	PushInCell                 CellAttribute = 2
)

// Constants used by the hitTestForEvent:inRect:ofView: method to determine the effect of an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscellhitresult?language=objc
type CellHitResult uint

const (
	CellHitContentArea      CellHitResult = 1
	CellHitEditableTextArea CellHitResult = 2
	CellHitNone             CellHitResult = 0
	CellHitTrackableArea    CellHitResult = 4
)

// A constant for specifying the position of a button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscellimageposition?language=objc
type CellImagePosition uint

const (
	ImageAbove    CellImagePosition = 5
	ImageBelow    CellImagePosition = 4
	ImageLeading  CellImagePosition = 7
	ImageLeft     CellImagePosition = 2
	ImageOnly     CellImagePosition = 1
	ImageOverlaps CellImagePosition = 6
	ImageRight    CellImagePosition = 3
	ImageTrailing CellImagePosition = 8
	NoImage       CellImagePosition = 0
)

// Constants for specifying a cell’s state and are used mostly for buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscellstatevalue?language=objc
type CellStateValue ControlStateValue

// Constants for specifying what happens when a button is pressed or is displaying its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscellstylemask?language=objc
type CellStyleMask uint

const (
	ChangeBackgroundCellMask CellStyleMask = 8
	ChangeGrayCellMask       CellStyleMask = 4
	ContentsCellMask         CellStyleMask = 1
	NoCellMask               CellStyleMask = 0
	PushInCellMask           CellStyleMask = 2
)

// Constants for specifying how a cell represents its data (as text or as an image). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscelltype?language=objc
type CellType uint

const (
	ImageCellType CellType = 2
	NullCellType  CellType = 0
	TextCellType  CellType = 1
)

// Values that map character identifiers to glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscharactercollection?language=objc
type CharacterCollection uint

const (
	AdobeCNS1CharacterCollection       CharacterCollection = 1
	AdobeGB1CharacterCollection        CharacterCollection = 2
	AdobeJapan1CharacterCollection     CharacterCollection = 3
	AdobeJapan2CharacterCollection     CharacterCollection = 4
	AdobeKorea1CharacterCollection     CharacterCollection = 5
	IdentityMappingCharacterCollection CharacterCollection = 0
)

// Constants that describe how a participant can configure a CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudkitsharingserviceoptions?language=objc
type CloudKitSharingServiceOptions uint

const (
	CloudKitSharingServiceAllowPrivate   CloudKitSharingServiceOptions = 2
	CloudKitSharingServiceAllowPublic    CloudKitSharingServiceOptions = 1
	CloudKitSharingServiceAllowReadOnly  CloudKitSharingServiceOptions = 16
	CloudKitSharingServiceAllowReadWrite CloudKitSharingServiceOptions = 32
	CloudKitSharingServiceStandard       CloudKitSharingServiceOptions = 0
)

// Constants specifying the type of element in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionelementcategory?language=objc
type CollectionElementCategory int

const (
	CollectionElementCategoryDecorationView    CollectionElementCategory = 2
	CollectionElementCategoryInterItemGap      CollectionElementCategory = 3
	CollectionElementCategoryItem              CollectionElementCategory = 0
	CollectionElementCategorySupplementaryView CollectionElementCategory = 1
)

// The scrolling behavior of the layout's sections in relation to the main layout axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutsectionorthogonalscrollingbehavior?language=objc
type CollectionLayoutSectionOrthogonalScrollingBehavior int

const (
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuous                     CollectionLayoutSectionOrthogonalScrollingBehavior = 1
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary CollectionLayoutSectionOrthogonalScrollingBehavior = 2
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging                    CollectionLayoutSectionOrthogonalScrollingBehavior = 4
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered            CollectionLayoutSectionOrthogonalScrollingBehavior = 5
	CollectionLayoutSectionOrthogonalScrollingBehaviorNone                           CollectionLayoutSectionOrthogonalScrollingBehavior = 0
	CollectionLayoutSectionOrthogonalScrollingBehaviorPaging                         CollectionLayoutSectionOrthogonalScrollingBehavior = 3
)

// Constants indicating the type of action being performed on an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionupdateaction?language=objc
type CollectionUpdateAction int

const (
	CollectionUpdateActionDelete CollectionUpdateAction = 1
	CollectionUpdateActionInsert CollectionUpdateAction = 0
	CollectionUpdateActionMove   CollectionUpdateAction = 3
	CollectionUpdateActionNone   CollectionUpdateAction = 4
	CollectionUpdateActionReload CollectionUpdateAction = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdecorationelementkind?language=objc
type CollectionViewDecorationElementKind string

// These constants specify if acceptance of a drop should be at the item it is dropped on or before the item. These constants are used by the  collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath: and collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath: methods in NSCollectionViewDelegate [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdropoperation?language=objc
type CollectionViewDropOperation int

const (
	CollectionViewDropBefore CollectionViewDropOperation = 1
	CollectionViewDropOn     CollectionViewDropOperation = 0
)

// Constants indicating the type of highlight applied to an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitemhighlightstate?language=objc
type CollectionViewItemHighlightState int

const (
	CollectionViewItemHighlightAsDropTarget   CollectionViewItemHighlightState = 3
	CollectionViewItemHighlightForDeselection CollectionViewItemHighlightState = 2
	CollectionViewItemHighlightForSelection   CollectionViewItemHighlightState = 1
	CollectionViewItemHighlightNone           CollectionViewItemHighlightState = 0
)

// Constants indicating the scrolling direction for the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewscrolldirection?language=objc
type CollectionViewScrollDirection int

const (
	CollectionViewScrollDirectionHorizontal CollectionViewScrollDirection = 1
	CollectionViewScrollDirectionVertical   CollectionViewScrollDirection = 0
)

// Constants indicating the options for scrolling the collection view’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewscrollposition?language=objc
type CollectionViewScrollPosition uint

const (
	CollectionViewScrollPositionBottom                CollectionViewScrollPosition = 4
	CollectionViewScrollPositionCenteredHorizontally  CollectionViewScrollPosition = 16
	CollectionViewScrollPositionCenteredVertically    CollectionViewScrollPosition = 2
	CollectionViewScrollPositionLeadingEdge           CollectionViewScrollPosition = 64
	CollectionViewScrollPositionLeft                  CollectionViewScrollPosition = 8
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 512
	CollectionViewScrollPositionNearestVerticalEdge   CollectionViewScrollPosition = 256
	CollectionViewScrollPositionNone                  CollectionViewScrollPosition = 0
	CollectionViewScrollPositionRight                 CollectionViewScrollPosition = 32
	CollectionViewScrollPositionTop                   CollectionViewScrollPosition = 1
	CollectionViewScrollPositionTrailingEdge          CollectionViewScrollPosition = 128
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsupplementaryelementkind?language=objc
type CollectionViewSupplementaryElementKind string

const (
	CollectionElementKindInterItemGapIndicator CollectionViewSupplementaryElementKind = "NSCollectionElementKindInterItemGapIndicator"
	CollectionElementKindSectionFooter         CollectionViewSupplementaryElementKind = "UICollectionElementKindSectionFooter"
	CollectionElementKindSectionHeader         CollectionViewSupplementaryElementKind = "UICollectionElementKindSectionHeader"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewtransitionlayoutanimatedkey?language=objc
type CollectionViewTransitionLayoutAnimatedKey string

// The name assigned to a color list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlistname?language=objc
type ColorListName string

// The name of a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorname?language=objc
type ColorName string

// A type defined for the enum constants specifying color panel modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanelmode?language=objc
type ColorPanelMode int

const (
	CMYKModeColorPanel          ColorPanelMode = 2
	ColorListModeColorPanel     ColorPanelMode = 5
	ColorPanelModeCMYK          ColorPanelMode = 2
	ColorPanelModeColorList     ColorPanelMode = 5
	ColorPanelModeCrayon        ColorPanelMode = 7
	ColorPanelModeCustomPalette ColorPanelMode = 4
	ColorPanelModeGray          ColorPanelMode = 0
	ColorPanelModeHSB           ColorPanelMode = 3
	ColorPanelModeNone          ColorPanelMode = -1
	ColorPanelModeRGB           ColorPanelMode = 1
	ColorPanelModeWheel         ColorPanelMode = 6
	CrayonModeColorPanel        ColorPanelMode = 7
	CustomPaletteModeColorPanel ColorPanelMode = 4
	GrayModeColorPanel          ColorPanelMode = 0
	HSBModeColorPanel           ColorPanelMode = 3
	NoModeColorPanel            ColorPanelMode = -1
	RGBModeColorPanel           ColorPanelMode = 1
	WheelModeColorPanel         ColorPanelMode = 6
)

// The color modes that are enabled for a color panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpaneloptions?language=objc
type ColorPanelOptions uint

const (
	ColorPanelAllModesMask          ColorPanelOptions = 65535
	ColorPanelCMYKModeMask          ColorPanelOptions = 4
	ColorPanelColorListModeMask     ColorPanelOptions = 32
	ColorPanelCrayonModeMask        ColorPanelOptions = 128
	ColorPanelCustomPaletteModeMask ColorPanelOptions = 16
	ColorPanelGrayModeMask          ColorPanelOptions = 1
	ColorPanelHSBModeMask           ColorPanelOptions = 8
	ColorPanelRGBModeMask           ColorPanelOptions = 2
	ColorPanelWheelModeMask         ColorPanelOptions = 64
)

// Constants that specify how Cocoa should handle colors that are not located within the destination color space of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorrenderingintent?language=objc
type ColorRenderingIntent int

const (
	ColorRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
	ColorRenderingIntentDefault              ColorRenderingIntent = 0
	ColorRenderingIntentPerceptual           ColorRenderingIntent = 3
	ColorRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
	ColorRenderingIntentSaturation           ColorRenderingIntent = 4
)

// Constants that describe the abstract model on which color space objects are based. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspacemodel?language=objc
type ColorSpaceModel int

const (
	CMYKColorSpaceModel      ColorSpaceModel = 2
	ColorSpaceModelCMYK      ColorSpaceModel = 2
	ColorSpaceModelDeviceN   ColorSpaceModel = 4
	ColorSpaceModelGray      ColorSpaceModel = 0
	ColorSpaceModelIndexed   ColorSpaceModel = 5
	ColorSpaceModelLAB       ColorSpaceModel = 3
	ColorSpaceModelPatterned ColorSpaceModel = 6
	ColorSpaceModelRGB       ColorSpaceModel = 1
	ColorSpaceModelUnknown   ColorSpaceModel = -1
	DeviceNColorSpaceModel   ColorSpaceModel = 4
	GrayColorSpaceModel      ColorSpaceModel = 0
	IndexedColorSpaceModel   ColorSpaceModel = 5
	LABColorSpaceModel       ColorSpaceModel = 3
	PatternColorSpaceModel   ColorSpaceModel = 6
	RGBColorSpaceModel       ColorSpaceModel = 1
	UnknownColorSpaceModel   ColorSpaceModel = -1
)

// Constants that specify color space names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspacename?language=objc
type ColorSpaceName string

const (
	CalibratedBlackColorSpace ColorSpaceName = "NSCalibratedBlackColorSpace"
	CalibratedRGBColorSpace   ColorSpaceName = "NSCalibratedRGBColorSpace"
	CalibratedWhiteColorSpace ColorSpaceName = "NSCalibratedWhiteColorSpace"
	CustomColorSpace          ColorSpaceName = "NSCustomColorSpace"
	DeviceBlackColorSpace     ColorSpaceName = "NSDeviceBlackColorSpace"
	DeviceCMYKColorSpace      ColorSpaceName = "NSDeviceCMYKColorSpace"
	DeviceRGBColorSpace       ColorSpaceName = "NSDeviceRGBColorSpace"
	DeviceWhiteColorSpace     ColorSpaceName = "NSDeviceWhiteColorSpace"
	NamedColorSpace           ColorSpaceName = "NSNamedColorSpace"
	PatternColorSpace         ColorSpaceName = "NSPatternColorSpace"
)

// Constants for user interactions that change the appearance of a view or control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorsystemeffect?language=objc
type ColorSystemEffect int

const (
	ColorSystemEffectDeepPressed ColorSystemEffect = 2
	ColorSystemEffectDisabled    ColorSystemEffect = 3
	ColorSystemEffectNone        ColorSystemEffect = 0
	ColorSystemEffectPressed     ColorSystemEffect = 1
	ColorSystemEffectRollover    ColorSystemEffect = 4
)

// Constants that indicate the color's type, and which methods may be called on the color object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolortype?language=objc
type ColorType int

const (
	ColorTypeCatalog        ColorType = 2
	ColorTypeComponentBased ColorType = 0
	ColorTypePattern        ColorType = 1
)

// Constants that describe compositing operators in terms of source and destination images, each having an opaque and transparent region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscompositingoperation?language=objc
type CompositingOperation uint

const (
	CompositeClear                      CompositingOperation = 0
	CompositeColor                      CompositingOperation = 27
	CompositeColorBurn                  CompositingOperation = 20
	CompositeColorDodge                 CompositingOperation = 19
	CompositeCopy                       CompositingOperation = 1
	CompositeDarken                     CompositingOperation = 17
	CompositeDestinationAtop            CompositingOperation = 9
	CompositeDestinationIn              CompositingOperation = 7
	CompositeDestinationOut             CompositingOperation = 8
	CompositeDestinationOver            CompositingOperation = 6
	CompositeDifference                 CompositingOperation = 23
	CompositeExclusion                  CompositingOperation = 24
	CompositeHardLight                  CompositingOperation = 22
	CompositeHighlight                  CompositingOperation = 12
	CompositeHue                        CompositingOperation = 25
	CompositeLighten                    CompositingOperation = 18
	CompositeLuminosity                 CompositingOperation = 28
	CompositeMultiply                   CompositingOperation = 14
	CompositeOverlay                    CompositingOperation = 16
	CompositePlusDarker                 CompositingOperation = 11
	CompositePlusLighter                CompositingOperation = 13
	CompositeSaturation                 CompositingOperation = 26
	CompositeScreen                     CompositingOperation = 15
	CompositeSoftLight                  CompositingOperation = 21
	CompositeSourceAtop                 CompositingOperation = 5
	CompositeSourceIn                   CompositingOperation = 3
	CompositeSourceOut                  CompositingOperation = 4
	CompositeSourceOver                 CompositingOperation = 2
	CompositeXOR                        CompositingOperation = 10
	CompositingOperationClear           CompositingOperation = 0
	CompositingOperationColor           CompositingOperation = 27
	CompositingOperationColorBurn       CompositingOperation = 20
	CompositingOperationColorDodge      CompositingOperation = 19
	CompositingOperationCopy            CompositingOperation = 1
	CompositingOperationDarken          CompositingOperation = 17
	CompositingOperationDestinationAtop CompositingOperation = 9
	CompositingOperationDestinationIn   CompositingOperation = 7
	CompositingOperationDestinationOut  CompositingOperation = 8
	CompositingOperationDestinationOver CompositingOperation = 6
	CompositingOperationDifference      CompositingOperation = 23
	CompositingOperationExclusion       CompositingOperation = 24
	CompositingOperationHardLight       CompositingOperation = 22
	CompositingOperationHighlight       CompositingOperation = 12
	CompositingOperationHue             CompositingOperation = 25
	CompositingOperationLighten         CompositingOperation = 18
	CompositingOperationLuminosity      CompositingOperation = 28
	CompositingOperationMultiply        CompositingOperation = 14
	CompositingOperationOverlay         CompositingOperation = 16
	CompositingOperationPlusDarker      CompositingOperation = 11
	CompositingOperationPlusLighter     CompositingOperation = 13
	CompositingOperationSaturation      CompositingOperation = 26
	CompositingOperationScreen          CompositingOperation = 15
	CompositingOperationSoftLight       CompositingOperation = 21
	CompositingOperationSourceAtop      CompositingOperation = 5
	CompositingOperationSourceIn        CompositingOperation = 3
	CompositingOperationSourceOut       CompositingOperation = 4
	CompositingOperationSourceOver      CompositingOperation = 2
	CompositingOperationXOR             CompositingOperation = 10
)

// Constants that describe actions for control characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscontrolcharacteraction?language=objc
type ControlCharacterAction int

const (
	ControlCharacterActionContainerBreak  ControlCharacterAction = 32
	ControlCharacterActionHorizontalTab   ControlCharacterAction = 4
	ControlCharacterActionLineBreak       ControlCharacterAction = 8
	ControlCharacterActionParagraphBreak  ControlCharacterAction = 16
	ControlCharacterActionWhitespace      ControlCharacterAction = 2
	ControlCharacterActionZeroAdvancement ControlCharacterAction = 1
)

// A constant for specifying a cell’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrolsize?language=objc
type ControlSize uint

const (
	ControlSizeLarge   ControlSize = 3
	ControlSizeMini    ControlSize = 2
	ControlSizeRegular ControlSize = 0
	ControlSizeSmall   ControlSize = 1
	MiniControlSize    ControlSize = 2
	RegularControlSize ControlSize = 0
	SmallControlSize   ControlSize = 1
)

// A constant that indicates whether a control is on, off, or in a mixed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrolstatevalue?language=objc
type ControlStateValue int

const (
	ControlStateValueMixed ControlStateValue = -1
	ControlStateValueOff   ControlStateValue = 0
	ControlStateValueOn    ControlStateValue = 1
	MixedState             ControlStateValue = -1
	OffState               ControlStateValue = 0
	OnState                ControlStateValue = 1
)

// Constants for specifying a cell’s tint color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltint?language=objc
type ControlTint uint

const (
	BlueControlTint     ControlTint = 1
	ClearControlTint    ControlTint = 7
	DefaultControlTint  ControlTint = 0
	GraphiteControlTint ControlTint = 6
)

// Constants that allow an app to specify the correction indicator type displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscorrectionindicatortype?language=objc
type CorrectionIndicatorType int

const (
	CorrectionIndicatorTypeDefault   CorrectionIndicatorType = 0
	CorrectionIndicatorTypeGuesses   CorrectionIndicatorType = 2
	CorrectionIndicatorTypeReversion CorrectionIndicatorType = 1
)

// The correction response passed to therecordResponse:toCorrection:forWord:language:inSpellDocumentWithTag: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscorrectionresponse?language=objc
type CorrectionResponse int

const (
	CorrectionResponseAccepted CorrectionResponse = 1
	CorrectionResponseEdited   CorrectionResponse = 4
	CorrectionResponseIgnored  CorrectionResponse = 3
	CorrectionResponseNone     CorrectionResponse = 0
	CorrectionResponseRejected CorrectionResponse = 2
	CorrectionResponseReverted CorrectionResponse = 5
)

// The name of a data asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdataassetname?language=objc
type DataAssetName string

// Constants that specify the date and time elements displayed by the picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickerelementflags?language=objc
type DatePickerElementFlags uint

const (
	DatePickerElementFlagEra              DatePickerElementFlags = 256
	DatePickerElementFlagHourMinute       DatePickerElementFlags = 12
	DatePickerElementFlagHourMinuteSecond DatePickerElementFlags = 14
	DatePickerElementFlagTimeZone         DatePickerElementFlags = 16
	DatePickerElementFlagYearMonth        DatePickerElementFlags = 192
	DatePickerElementFlagYearMonthDay     DatePickerElementFlags = 224
	EraDatePickerElementFlag              DatePickerElementFlags = 256
	HourMinuteDatePickerElementFlag       DatePickerElementFlags = 12
	HourMinuteSecondDatePickerElementFlag DatePickerElementFlags = 14
	TimeZoneDatePickerElementFlag         DatePickerElementFlags = 16
	YearMonthDatePickerElementFlag        DatePickerElementFlags = 192
	YearMonthDayDatePickerElementFlag     DatePickerElementFlags = 224
)

// Constants that define whether the picker provides a single date, or a range of dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickermode?language=objc
type DatePickerMode uint

const (
	DatePickerModeRange  DatePickerMode = 1
	DatePickerModeSingle DatePickerMode = 0
	RangeDateMode        DatePickerMode = 1
	SingleDateMode       DatePickerMode = 0
)

// Constants that define the visual appearance of the date picker cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickerstyle?language=objc
type DatePickerStyle uint

const (
	ClockAndCalendarDatePickerStyle    DatePickerStyle = 1
	DatePickerStyleClockAndCalendar    DatePickerStyle = 1
	DatePickerStyleTextField           DatePickerStyle = 2
	DatePickerStyleTextFieldAndStepper DatePickerStyle = 0
	TextFieldAndStepperDatePickerStyle DatePickerStyle = 0
	TextFieldDatePickerStyle           DatePickerStyle = 2
)

// Keys to include in your definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdefinitionoptionkey?language=objc
type DefinitionOptionKey string

const (
	DefinitionPresentationTypeKey DefinitionOptionKey = "NSDefinitionPresentationTypeKey"
)

// Presentation options for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdefinitionpresentationtype?language=objc
type DefinitionPresentationType string

const (
	DefinitionPresentationTypeDictionaryApplication DefinitionPresentationType = "NSDefinitionPresentationTypeDictionaryApplication"
	DefinitionPresentationTypeOverlay               DefinitionPresentationType = "NSDefinitionPresentationTypeOverlay"
)

// These constants are the keys for device description dictionaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdevicedescriptionkey?language=objc
type DeviceDescriptionKey string

const (
	DeviceBitsPerSample  DeviceDescriptionKey = "NSDeviceBitsPerSample"
	DeviceColorSpaceName DeviceDescriptionKey = "NSDeviceColorSpaceName"
	DeviceIsPrinter      DeviceDescriptionKey = "NSDeviceIsPrinter"
	DeviceIsScreen       DeviceDescriptionKey = "NSDeviceIsScreen"
	DeviceResolution     DeviceDescriptionKey = "NSDeviceResolution"
	DeviceSize           DeviceDescriptionKey = "NSDeviceSize"
)

// Constants that specify an edge or a set of edges, taking the user interface layout direction into account. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdirectionalrectedge?language=objc
type DirectionalRectEdge uint

const (
	DirectionalRectEdgeAll      DirectionalRectEdge = 15
	DirectionalRectEdgeBottom   DirectionalRectEdge = 4
	DirectionalRectEdgeLeading  DirectionalRectEdge = 2
	DirectionalRectEdgeNone     DirectionalRectEdge = 0
	DirectionalRectEdgeTop      DirectionalRectEdge = 1
	DirectionalRectEdgeTrailing DirectionalRectEdge = 8
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdisplaygamut?language=objc
type DisplayGamut int

const (
	DisplayGamutP3   DisplayGamut = 2
	DisplayGamutSRGB DisplayGamut = 1
)

// Values that indicate a document’s edit status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentchangetype?language=objc
type DocumentChangeType uint

const (
	ChangeAutosaved         DocumentChangeType = 4
	ChangeCleared           DocumentChangeType = 2
	ChangeDiscardable       DocumentChangeType = 256
	ChangeDone              DocumentChangeType = 0
	ChangeReadOtherContents DocumentChangeType = 3
	ChangeRedone            DocumentChangeType = 5
	ChangeUndone            DocumentChangeType = 1
)

// A group of constants that represent which operations the dragging source can perform on dragging items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragoperation?language=objc
type DragOperation uint

const (
	DragOperationAll_Obsolete DragOperation = 15
	DragOperationCopy         DragOperation = 1
	DragOperationDelete       DragOperation = 32
	DragOperationEvery        DragOperation = math.MaxUint
	DragOperationGeneric      DragOperation = 4
	DragOperationLink         DragOperation = 2
	DragOperationMove         DragOperation = 16
	DragOperationNone         DragOperation = 0
	DragOperationPrivate      DragOperation = 8
)

// Constants that specify whether a drag terminates within or outside the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingcontext?language=objc
type DraggingContext int

const (
	DraggingContextOutsideApplication DraggingContext = 0
	DraggingContextWithinApplication  DraggingContext = 1
)

// Constants that control the visual format of multiple dragging items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingformation?language=objc
type DraggingFormation int

const (
	DraggingFormationDefault DraggingFormation = 0
	DraggingFormationList    DraggingFormation = 3
	DraggingFormationNone    DraggingFormation = 1
	DraggingFormationPile    DraggingFormation = 2
	DraggingFormationStack   DraggingFormation = 4
)

// Keys that identify components of a dragging image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponentkey?language=objc
type DraggingImageComponentKey string

const (
	DraggingImageComponentIconKey  DraggingImageComponentKey = "icon"
	DraggingImageComponentLabelKey DraggingImageComponentKey = "label"
)

// A group of constants that specify options to use when enumerating dragging items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitemenumerationoptions?language=objc
type DraggingItemEnumerationOptions uint

const (
	DraggingItemEnumerationClearNonenumeratedImages DraggingItemEnumerationOptions = 65536
	DraggingItemEnumerationConcurrent               DraggingItemEnumerationOptions = 1
)

// These constants specify the possible states of a drawer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawerstate?language=objc
type DrawerState uint

const (
	DrawerClosedState  DrawerState = 0
	DrawerClosingState DrawerState = 3
	DrawerOpenState    DrawerState = 2
	DrawerOpeningState DrawerState = 1
)

// Constants you use to identify the activated tablet buttons in an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventbuttonmask?language=objc
type EventButtonMask uint

const (
	EventButtonMaskPenLowerSide EventButtonMask = 2
	EventButtonMaskPenTip       EventButtonMask = 1
	EventButtonMaskPenUpperSide EventButtonMask = 4
	PenLowerSideMask            EventButtonMask = 2
	PenTipMask                  EventButtonMask = 1
	PenUpperSideMask            EventButtonMask = 4
)

// Constants that specify the direction of travel for a gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventgestureaxis?language=objc
type EventGestureAxis int

const (
	EventGestureAxisHorizontal EventGestureAxis = 1
	EventGestureAxisNone       EventGestureAxis = 0
	EventGestureAxisVertical   EventGestureAxis = 2
)

// Constants that you use to filter out specific event types from the stream of incoming events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventmask?language=objc
type EventMask int64

const (
	AnyEventMask                EventMask = -1
	AppKitDefinedMask           EventMask = 8192
	ApplicationDefinedMask      EventMask = 32768
	CursorUpdateMask            EventMask = 131072
	EventMaskAny                EventMask = -1
	EventMaskAppKitDefined      EventMask = 8192
	EventMaskApplicationDefined EventMask = 32768
	EventMaskBeginGesture       EventMask = 524288
	EventMaskChangeMode         EventMask = 274877906944
	EventMaskCursorUpdate       EventMask = 131072
	EventMaskDirectTouch        EventMask = 137438953472
	EventMaskEndGesture         EventMask = 1048576
	EventMaskFlagsChanged       EventMask = 4096
	EventMaskGesture            EventMask = 536870912
	EventMaskKeyDown            EventMask = 1024
	EventMaskKeyUp              EventMask = 2048
	EventMaskLeftMouseDown      EventMask = 2
	EventMaskLeftMouseDragged   EventMask = 64
	EventMaskLeftMouseUp        EventMask = 4
	EventMaskMagnify            EventMask = 1073741824
	EventMaskMouseEntered       EventMask = 256
	EventMaskMouseExited        EventMask = 512
	EventMaskMouseMoved         EventMask = 32
	EventMaskOtherMouseDown     EventMask = 33554432
	EventMaskOtherMouseDragged  EventMask = 134217728
	EventMaskOtherMouseUp       EventMask = 67108864
	EventMaskPeriodic           EventMask = 65536
	EventMaskPressure           EventMask = 17179869184
	EventMaskRightMouseDown     EventMask = 8
	EventMaskRightMouseDragged  EventMask = 128
	EventMaskRightMouseUp       EventMask = 16
	EventMaskRotate             EventMask = 262144
	EventMaskScrollWheel        EventMask = 4194304
	EventMaskSmartMagnify       EventMask = 4294967296
	EventMaskSwipe              EventMask = 2147483648
	EventMaskSystemDefined      EventMask = 16384
	EventMaskTabletPoint        EventMask = 8388608
	EventMaskTabletProximity    EventMask = 16777216
	FlagsChangedMask            EventMask = 4096
	KeyDownMask                 EventMask = 1024
	KeyUpMask                   EventMask = 2048
	LeftMouseDownMask           EventMask = 2
	LeftMouseDraggedMask        EventMask = 64
	LeftMouseUpMask             EventMask = 4
	MouseEnteredMask            EventMask = 256
	MouseExitedMask             EventMask = 512
	MouseMovedMask              EventMask = 32
	OtherMouseDownMask          EventMask = 33554432
	OtherMouseDraggedMask       EventMask = 134217728
	OtherMouseUpMask            EventMask = 67108864
	PeriodicMask                EventMask = 65536
	RightMouseDownMask          EventMask = 8
	RightMouseDraggedMask       EventMask = 128
	RightMouseUpMask            EventMask = 16
	ScrollWheelMask             EventMask = 4194304
	SystemDefinedMask           EventMask = 16384
	TabletPointMask             EventMask = 8388608
	TabletProximityMask         EventMask = 16777216
)

// Flags that represent key states in an event object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventmodifierflags?language=objc
type EventModifierFlags uint

const (
	AlphaShiftKeyMask                           EventModifierFlags = 65536
	AlternateKeyMask                            EventModifierFlags = 524288
	CommandKeyMask                              EventModifierFlags = 1048576
	ControlKeyMask                              EventModifierFlags = 262144
	DeviceIndependentModifierFlagsMask          EventModifierFlags = 4294901760
	EventModifierFlagCapsLock                   EventModifierFlags = 65536
	EventModifierFlagCommand                    EventModifierFlags = 1048576
	EventModifierFlagControl                    EventModifierFlags = 262144
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 4294901760
	EventModifierFlagFunction                   EventModifierFlags = 8388608
	EventModifierFlagHelp                       EventModifierFlags = 4194304
	EventModifierFlagNumericPad                 EventModifierFlags = 2097152
	EventModifierFlagOption                     EventModifierFlags = 524288
	EventModifierFlagShift                      EventModifierFlags = 131072
	FunctionKeyMask                             EventModifierFlags = 8388608
	HelpKeyMask                                 EventModifierFlags = 4194304
	NumericPadKeyMask                           EventModifierFlags = 2097152
	ShiftKeyMask                                EventModifierFlags = 131072
)

// Constants that represent the possible phases during an event phase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventphase?language=objc
type EventPhase uint

const (
	EventPhaseBegan      EventPhase = 1
	EventPhaseCancelled  EventPhase = 16
	EventPhaseChanged    EventPhase = 4
	EventPhaseEnded      EventPhase = 8
	EventPhaseMayBegin   EventPhase = 32
	EventPhaseNone       EventPhase = 0
	EventPhaseStationary EventPhase = 2
)

// Subtypes for various types of events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventsubtype?language=objc
type EventSubtype int

const (
	AWTEventType                       EventSubtype = 16
	ApplicationActivatedEventType      EventSubtype = 1
	ApplicationDeactivatedEventType    EventSubtype = 2
	EventSubtypeApplicationActivated   EventSubtype = 1
	EventSubtypeApplicationDeactivated EventSubtype = 2
	EventSubtypeMouseEvent             EventSubtype = 0
	EventSubtypePowerOff               EventSubtype = 1
	EventSubtypeScreenChanged          EventSubtype = 8
	EventSubtypeTabletPoint            EventSubtype = 1
	EventSubtypeTabletProximity        EventSubtype = 2
	EventSubtypeTouch                  EventSubtype = 3
	EventSubtypeWindowExposed          EventSubtype = 0
	EventSubtypeWindowMoved            EventSubtype = 4
	MouseEventSubtype                  EventSubtype = 0
	PowerOffEventType                  EventSubtype = 1
	ScreenChangedEventType             EventSubtype = 8
	TabletPointEventSubtype            EventSubtype = 1
	TabletProximityEventSubtype        EventSubtype = 2
	TouchEventSubtype                  EventSubtype = 3
	WindowExposedEventType             EventSubtype = 0
	WindowMovedEventType               EventSubtype = 4
)

// Constants that specify swipe-tracking options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventswipetrackingoptions?language=objc
type EventSwipeTrackingOptions uint

const (
	EventSwipeTrackingClampGestureAmount EventSwipeTrackingOptions = 2
	EventSwipeTrackingLockDirection      EventSwipeTrackingOptions = 1
)

// Constants for the types of events that responder objects can handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseventtype?language=objc
type EventType uint

const (
	AppKitDefined               EventType = 13
	ApplicationDefined          EventType = 15
	CursorUpdate                EventType = 17
	EventTypeAppKitDefined      EventType = 13
	EventTypeApplicationDefined EventType = 15
	EventTypeBeginGesture       EventType = 19
	EventTypeChangeMode         EventType = 38
	EventTypeCursorUpdate       EventType = 17
	EventTypeDirectTouch        EventType = 37
	EventTypeEndGesture         EventType = 20
	EventTypeFlagsChanged       EventType = 12
	EventTypeGesture            EventType = 29
	EventTypeKeyDown            EventType = 10
	EventTypeKeyUp              EventType = 11
	EventTypeLeftMouseDown      EventType = 1
	EventTypeLeftMouseDragged   EventType = 6
	EventTypeLeftMouseUp        EventType = 2
	EventTypeMagnify            EventType = 30
	EventTypeMouseEntered       EventType = 8
	EventTypeMouseExited        EventType = 9
	EventTypeMouseMoved         EventType = 5
	EventTypeOtherMouseDown     EventType = 25
	EventTypeOtherMouseDragged  EventType = 27
	EventTypeOtherMouseUp       EventType = 26
	EventTypePeriodic           EventType = 16
	EventTypePressure           EventType = 34
	EventTypeQuickLook          EventType = 33
	EventTypeRightMouseDown     EventType = 3
	EventTypeRightMouseDragged  EventType = 7
	EventTypeRightMouseUp       EventType = 4
	EventTypeRotate             EventType = 18
	EventTypeScrollWheel        EventType = 22
	EventTypeSmartMagnify       EventType = 32
	EventTypeSwipe              EventType = 31
	EventTypeSystemDefined      EventType = 14
	EventTypeTabletPoint        EventType = 23
	EventTypeTabletProximity    EventType = 24
	FlagsChanged                EventType = 12
	KeyDown                     EventType = 10
	KeyUp                       EventType = 11
	LeftMouseDown               EventType = 1
	LeftMouseDragged            EventType = 6
	LeftMouseUp                 EventType = 2
	MouseEntered                EventType = 8
	MouseExited                 EventType = 9
	MouseMoved                  EventType = 5
	OtherMouseDown              EventType = 25
	OtherMouseDragged           EventType = 27
	OtherMouseUp                EventType = 26
	Periodic                    EventType = 16
	RightMouseDown              EventType = 3
	RightMouseDragged           EventType = 7
	RightMouseUp                EventType = 4
	ScrollWheel                 EventType = 22
	SystemDefined               EventType = 14
	TabletPoint                 EventType = 23
	TabletProximity             EventType = 24
)

// These constants define the tags for performFindPanelAction:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfindpanelaction?language=objc
type FindPanelAction uint

const (
	FindPanelActionNext                  FindPanelAction = 2
	FindPanelActionPrevious              FindPanelAction = 3
	FindPanelActionReplace               FindPanelAction = 5
	FindPanelActionReplaceAll            FindPanelAction = 4
	FindPanelActionReplaceAllInSelection FindPanelAction = 8
	FindPanelActionReplaceAndFind        FindPanelAction = 6
	FindPanelActionSelectAll             FindPanelAction = 9
	FindPanelActionSelectAllInSelection  FindPanelAction = 10
	FindPanelActionSetFindString         FindPanelAction = 7
	FindPanelActionShowFindPanel         FindPanelAction = 1
)

// The type of substring matching used by the Find panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfindpanelsubstringmatchtype?language=objc
type FindPanelSubstringMatchType uint

const (
	FindPanelSubstringMatchTypeContains   FindPanelSubstringMatchType = 0
	FindPanelSubstringMatchTypeEndsWith   FindPanelSubstringMatchType = 3
	FindPanelSubstringMatchTypeFullWord   FindPanelSubstringMatchType = 2
	FindPanelSubstringMatchTypeStartsWith FindPanelSubstringMatchType = 1
)

// Constants that indicate how the system draws the focus ring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfocusringplacement?language=objc
type FocusRingPlacement uint

const (
	FocusRingAbove FocusRingPlacement = 2
	FocusRingBelow FocusRingPlacement = 1
	FocusRingOnly  FocusRingPlacement = 0
)

// Constants that describe the style of the focus ring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfocusringtype?language=objc
type FocusRingType uint

const (
	FocusRingTypeDefault  FocusRingType = 0
	FocusRingTypeExterior FocusRingType = 2
	FocusRingTypeNone     FocusRingType = 1
)

// Actions that modify a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontaction?language=objc
type FontAction uint

const (
	AddTraitFontAction    FontAction = 2
	HeavierFontAction     FontAction = 5
	LighterFontAction     FontAction = 6
	NoFontChangeAction    FontAction = 0
	RemoveTraitFontAction FontAction = 7
	SizeDownFontAction    FontAction = 4
	SizeUpFontAction      FontAction = 3
	ViaPanelFontAction    FontAction = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequestoptions?language=objc
type FontAssetRequestOptions uint

const (
	FontAssetRequestOptionUsesStandardUI FontAssetRequestOptions = 1
)

// The following actions are possible values of the NSFontCollectionActionKey in the NSFontCollectionDidChangeNotification userInfo method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionactiontypekey?language=objc
type FontCollectionActionTypeKey string

const (
	FontCollectionWasHidden  FontCollectionActionTypeKey = "NSFontCollectionWasHidden"
	FontCollectionWasRenamed FontCollectionActionTypeKey = "NSFontCollectionWasRenamed"
	FontCollectionWasShown   FontCollectionActionTypeKey = "NSFontCollectionWasShown"
)

// These constants are used by the matchingDescriptorsWithOptions: and matchingDescriptorsForFamily: options dictionary parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionmatchingoptionkey?language=objc
type FontCollectionMatchingOptionKey string

const (
	FontCollectionDisallowAutoActivationOption FontCollectionMatchingOptionKey = "NSCTFontCollectionDisallowAutoActivationOption"
	FontCollectionIncludeDisabledFontsOption   FontCollectionMatchingOptionKey = "NSCTFontCollectionIncludeDisabledFontsOption"
	FontCollectionRemoveDuplicatesOption       FontCollectionMatchingOptionKey = "NSCTFontCollectionRemoveDuplicatesOption"
)

// The constants represent the standard mutable collection names—these names are included in the list of allFontCollectionNames--they have special meaning to the Cocoa font system and should not be hidden or renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionname?language=objc
type FontCollectionName string

const (
	FontCollectionAllFonts     FontCollectionName = "com.apple.AllFonts"
	FontCollectionFavorites    FontCollectionName = "com.apple.Favorites"
	FontCollectionRecentlyUsed FontCollectionName = "com.apple.Recents"
	FontCollectionUser         FontCollectionName = "com.apple.UserFonts"
)

// Constants that support font collection management. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionoptions?language=objc
type FontCollectionOptions uint

const (
	FontCollectionApplicationOnlyMask FontCollectionOptions = 1
)

// These constants are used as keys in the NSFontCollectionDidChangeNotification userInfo dictionary to indicate the changes that have taken place. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionuserinfokey?language=objc
type FontCollectionUserInfoKey string

const (
	FontCollectionActionKey     FontCollectionUserInfoKey = "NSFontCollectionAction"
	FontCollectionNameKey       FontCollectionUserInfoKey = "NSFontCollectionName"
	FontCollectionOldNameKey    FontCollectionUserInfoKey = "NSFontCollectionOldName"
	FontCollectionVisibilityKey FontCollectionUserInfoKey = "NSFontCollectionVisibility"
)

// Constants that specify the visibility of font collections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollectionvisibility?language=objc
type FontCollectionVisibility uint

const (
	FontCollectionVisibilityComputer FontCollectionVisibility = 4
	FontCollectionVisibilityProcess  FontCollectionVisibility = 1
	FontCollectionVisibilityUser     FontCollectionVisibility = 2
)

// Constants for the names of font attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptorattributename?language=objc
type FontDescriptorAttributeName string

const (
	FontCascadeListAttribute     FontDescriptorAttributeName = "NSCTFontCascadeListAttribute"
	FontCharacterSetAttribute    FontDescriptorAttributeName = "NSCTFontCharacterSetAttribute"
	FontFaceAttribute            FontDescriptorAttributeName = "NSFontFaceAttribute"
	FontFamilyAttribute          FontDescriptorAttributeName = "NSFontFamilyAttribute"
	FontFeatureSettingsAttribute FontDescriptorAttributeName = "NSCTFontFeatureSettingsAttribute"
	FontFixedAdvanceAttribute    FontDescriptorAttributeName = "NSCTFontFixedAdvanceAttribute"
	FontMatrixAttribute          FontDescriptorAttributeName = "NSFontMatrixAttribute"
	FontNameAttribute            FontDescriptorAttributeName = "NSFontNameAttribute"
	FontSizeAttribute            FontDescriptorAttributeName = "NSFontSizeAttribute"
	FontTraitsAttribute          FontDescriptorAttributeName = "NSCTFontTraitsAttribute"
	FontVariationAttribute       FontDescriptorAttributeName = "NSCTFontVariationAttribute"
	FontVisibleNameAttribute     FontDescriptorAttributeName = "NSFontVisibleNameAttribute"
)

// Constants to use as keys to retrieve information about a font descriptor from its feature dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptorfeaturekey?language=objc
type FontDescriptorFeatureKey string

const (
	FontFeatureSelectorIdentifierKey FontDescriptorFeatureKey = "CTFeatureSelectorIdentifier"
	FontFeatureTypeIdentifierKey     FontDescriptorFeatureKey = "CTFeatureTypeIdentifier"
)

// A symbolic description of the stylistic aspects of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptorsymbolictraits?language=objc
type FontDescriptorSymbolicTraits uint32

const (
	FontDescriptorClassClarendonSerifs    FontDescriptorSymbolicTraits = 1073741824
	FontDescriptorClassFreeformSerifs     FontDescriptorSymbolicTraits = 1879048192
	FontDescriptorClassMask               FontDescriptorSymbolicTraits = 4026531840
	FontDescriptorClassModernSerifs       FontDescriptorSymbolicTraits = 805306368
	FontDescriptorClassOldStyleSerifs     FontDescriptorSymbolicTraits = 268435456
	FontDescriptorClassOrnamentals        FontDescriptorSymbolicTraits = 2415919104
	FontDescriptorClassSansSerif          FontDescriptorSymbolicTraits = 2147483648
	FontDescriptorClassScripts            FontDescriptorSymbolicTraits = 2684354560
	FontDescriptorClassSlabSerifs         FontDescriptorSymbolicTraits = 1342177280
	FontDescriptorClassSymbolic           FontDescriptorSymbolicTraits = 3221225472
	FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 536870912
	FontDescriptorClassUnknown            FontDescriptorSymbolicTraits = 0
	FontDescriptorTraitBold               FontDescriptorSymbolicTraits = 2
	FontDescriptorTraitCondensed          FontDescriptorSymbolicTraits = 64
	FontDescriptorTraitExpanded           FontDescriptorSymbolicTraits = 32
	FontDescriptorTraitItalic             FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitLooseLeading       FontDescriptorSymbolicTraits = 65536
	FontDescriptorTraitMonoSpace          FontDescriptorSymbolicTraits = 1024
	FontDescriptorTraitTightLeading       FontDescriptorSymbolicTraits = 32768
	FontDescriptorTraitUIOptimized        FontDescriptorSymbolicTraits = 4096
	FontDescriptorTraitVertical           FontDescriptorSymbolicTraits = 2048
)

// Constants for font designs, such as monospace, rounded, and serif. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptorsystemdesign?language=objc
type FontDescriptorSystemDesign string

const (
	FontDescriptorSystemDesignDefault    FontDescriptorSystemDesign = "NSCTFontUIFontDesignDefault"
	FontDescriptorSystemDesignMonospaced FontDescriptorSystemDesign = "NSCTFontUIFontDesignMonospaced"
	FontDescriptorSystemDesignRounded    FontDescriptorSystemDesign = "NSCTFontUIFontDesignRounded"
	FontDescriptorSystemDesignSerif      FontDescriptorSystemDesign = "NSCTFontUIFontDesignSerif"
)

// Constants that can be used as keys to retrieve information about a font descriptor from its trait dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptortraitkey?language=objc
type FontDescriptorTraitKey string

const (
	FontSlantTrait    FontDescriptorTraitKey = "NSCTFontSlantTrait"
	FontSymbolicTrait FontDescriptorTraitKey = "NSCTFontSymbolicTrait"
	FontWeightTrait   FontDescriptorTraitKey = "NSCTFontWeightTrait"
	FontWidthTrait    FontDescriptorTraitKey = "NSCTFontProportionTrait"
)

// Constants that can be used as keys to retrieve information about a font descriptor from its variation axis dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptorvariationkey?language=objc
type FontDescriptorVariationKey string

const (
	FontVariationAxisDefaultValueKey FontDescriptorVariationKey = "NSCTVariationAxisDefaultValue"
	FontVariationAxisIdentifierKey   FontDescriptorVariationKey = "NSCTVariationAxisIdentifier"
	FontVariationAxisMaximumValueKey FontDescriptorVariationKey = "NSCTVariationAxisMaximumValue"
	FontVariationAxisMinimumValueKey FontDescriptorVariationKey = "NSCTVariationAxisMinimumValue"
	FontVariationAxisNameKey         FontDescriptorVariationKey = "NSCTVariationAxisName"
)

// Constants that classify certain stylistic qualities of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontfamilyclass?language=objc
type FontFamilyClass uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanelmodemask?language=objc
type FontPanelModeMask uint

const (
	FontPanelModeMaskAllEffects          FontPanelModeMask = 1048320
	FontPanelModeMaskCollection          FontPanelModeMask = 4
	FontPanelModeMaskDocumentColorEffect FontPanelModeMask = 2048
	FontPanelModeMaskFace                FontPanelModeMask = 1
	FontPanelModeMaskShadowEffect        FontPanelModeMask = 4096
	FontPanelModeMaskSize                FontPanelModeMask = 2
	FontPanelModeMaskStrikethroughEffect FontPanelModeMask = 512
	FontPanelModeMaskTextColorEffect     FontPanelModeMask = 1024
	FontPanelModeMaskUnderlineEffect     FontPanelModeMask = 256
	FontPanelModesMaskAllModes           FontPanelModeMask = 4294967295
	FontPanelModesMaskStandardModes      FontPanelModeMask = 65535
)

// The font rendering mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontrenderingmode?language=objc
type FontRenderingMode uint

const (
	FontAntialiasedIntegerAdvancementsRenderingMode FontRenderingMode = 3
	FontAntialiasedRenderingMode                    FontRenderingMode = 1
	FontDefaultRenderingMode                        FontRenderingMode = 0
	FontIntegerAdvancementsRenderingMode            FontRenderingMode = 2
)

// A symbolic description of stylistic aspects of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontsymbolictraits?language=objc
type FontSymbolicTraits uint32

// Constants that specify the preferred text styles you use with fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfonttextstyle?language=objc
type FontTextStyle string

const (
	FontTextStyleBody        FontTextStyle = "UICTFontTextStyleBody"
	FontTextStyleCallout     FontTextStyle = "UICTFontTextStyleCallout"
	FontTextStyleCaption1    FontTextStyle = "UICTFontTextStyleCaption1"
	FontTextStyleCaption2    FontTextStyle = "UICTFontTextStyleCaption2"
	FontTextStyleFootnote    FontTextStyle = "UICTFontTextStyleFootnote"
	FontTextStyleHeadline    FontTextStyle = "UICTFontTextStyleHeadline"
	FontTextStyleLargeTitle  FontTextStyle = "UICTFontTextStyleTitle0"
	FontTextStyleSubheadline FontTextStyle = "UICTFontTextStyleSubhead"
	FontTextStyleTitle1      FontTextStyle = "UICTFontTextStyleTitle1"
	FontTextStyleTitle2      FontTextStyle = "UICTFontTextStyleTitle2"
	FontTextStyleTitle3      FontTextStyle = "UICTFontTextStyleTitle3"
)

// The options that you apply when requesting the font or font descriptor of a preferred text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfonttextstyleoptionkey?language=objc
type FontTextStyleOptionKey string

// Constants for isolating specific traits of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfonttraitmask?language=objc
type FontTraitMask uint

const (
	BoldFontMask                    FontTraitMask = 2
	CompressedFontMask              FontTraitMask = 512
	CondensedFontMask               FontTraitMask = 64
	ExpandedFontMask                FontTraitMask = 32
	FixedPitchFontMask              FontTraitMask = 1024
	ItalicFontMask                  FontTraitMask = 1
	NarrowFontMask                  FontTraitMask = 16
	NonStandardCharacterSetFontMask FontTraitMask = 8
	PosterFontMask                  FontTraitMask = 256
	SmallCapsFontMask               FontTraitMask = 128
	UnboldFontMask                  FontTraitMask = 4
	UnitalicFontMask                FontTraitMask = 16777216
)

// System-defined font-weight values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontweight?language=objc
type FontWeight float64

const (
	FontWeightBlack      FontWeight = 0.620000
	FontWeightBold       FontWeight = 0.400000
	FontWeightHeavy      FontWeight = 0.560000
	FontWeightLight      FontWeight = -0.400000
	FontWeightMedium     FontWeight = 0.230000
	FontWeightRegular    FontWeight = 0.000000
	FontWeightSemibold   FontWeight = 0.300000
	FontWeightThin       FontWeight = -0.600000
	FontWeightUltraLight FontWeight = -0.800000
)

// The current state of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerstate?language=objc
type GestureRecognizerState int

const (
	GestureRecognizerStateBegan      GestureRecognizerState = 1
	GestureRecognizerStateCancelled  GestureRecognizerState = 4
	GestureRecognizerStateChanged    GestureRecognizerState = 2
	GestureRecognizerStateEnded      GestureRecognizerState = 3
	GestureRecognizerStateFailed     GestureRecognizerState = 5
	GestureRecognizerStatePossible   GestureRecognizerState = 0
	GestureRecognizerStateRecognized GestureRecognizerState = 3
)

// The type used to specify glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyph?language=objc
type Glyph int

// Constants that specify how a glyph is laid out relative to the previous glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinscription?language=objc
type GlyphInscription uint

const (
	GlyphInscribeAbove      GlyphInscription = 2
	GlyphInscribeBase       GlyphInscription = 0
	GlyphInscribeBelow      GlyphInscription = 1
	GlyphInscribeOverBelow  GlyphInscription = 4
	GlyphInscribeOverstrike GlyphInscription = 3
)

// Glyph properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsglyphproperty?language=objc
type GlyphProperty int

const (
	GlyphPropertyControlCharacter GlyphProperty = 2
	GlyphPropertyElastic          GlyphProperty = 4
	GlyphPropertyNonBaseCharacter GlyphProperty = 8
	GlyphPropertyNull             GlyphProperty = 1
)

// Constants that specify gradient drawing options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradientdrawingoptions?language=objc
type GradientDrawingOptions uint

const (
	GradientDrawsAfterEndingLocation    GradientDrawingOptions = 2
	GradientDrawsBeforeStartingLocation GradientDrawingOptions = 1
)

// Specify the gradients used by the gradientType property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradienttype?language=objc
type GradientType uint

const (
	GradientConcaveStrong GradientType = 2
	GradientConcaveWeak   GradientType = 1
	GradientConvexStrong  GradientType = 4
	GradientConvexWeak    GradientType = 3
	GradientNone          GradientType = 0
)

// Constants that specify the dictionary keys for the attributes of the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontextattributekey?language=objc
type GraphicsContextAttributeKey string

const (
	GraphicsContextDestinationAttributeName          GraphicsContextAttributeKey = "NSGraphicsContextDestinationAttributeName"
	GraphicsContextRepresentationFormatAttributeName GraphicsContextAttributeKey = "NSGraphicsContextRepresentationFormatAttributeName"
)

// Constants that specify values for the representation format name key in a graphic context’s attributes dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontextrepresentationformatname?language=objc
type GraphicsContextRepresentationFormatName string

const (
	GraphicsContextPDFFormat GraphicsContextRepresentationFormatName = "NSGraphicsContextPDFFormat"
	GraphicsContextPSFormat  GraphicsContextRepresentationFormatName = "NSGraphicsContextPSFormat"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcellplacement?language=objc
type GridCellPlacement int

const (
	GridCellPlacementBottom    GridCellPlacement = 3
	GridCellPlacementCenter    GridCellPlacement = 4
	GridCellPlacementFill      GridCellPlacement = 5
	GridCellPlacementInherited GridCellPlacement = 0
	GridCellPlacementLeading   GridCellPlacement = 2
	GridCellPlacementNone      GridCellPlacement = 1
	GridCellPlacementTop       GridCellPlacement = 2
	GridCellPlacementTrailing  GridCellPlacement = 3
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrowalignment?language=objc
type GridRowAlignment int

const (
	GridRowAlignmentFirstBaseline GridRowAlignment = 2
	GridRowAlignmentInherited     GridRowAlignment = 0
	GridRowAlignmentLastBaseline  GridRowAlignment = 3
	GridRowAlignmentNone          GridRowAlignment = 1
)

// A pattern of haptic feedback to be provided to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackpattern?language=objc
type HapticFeedbackPattern int

const (
	HapticFeedbackPatternAlignment   HapticFeedbackPattern = 1
	HapticFeedbackPatternGeneric     HapticFeedbackPattern = 0
	HapticFeedbackPatternLevelChange HapticFeedbackPattern = 2
)

// A time at which to provide haptic feedback to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshapticfeedbackperformancetime?language=objc
type HapticFeedbackPerformanceTime uint

const (
	HapticFeedbackPerformanceTimeDefault       HapticFeedbackPerformanceTime = 0
	HapticFeedbackPerformanceTimeDrawCompleted HapticFeedbackPerformanceTime = 2
	HapticFeedbackPerformanceTimeNow           HapticFeedbackPerformanceTime = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpanchorname?language=objc
type HelpAnchorName string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpbookname?language=objc
type HelpBookName string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nshelpmanagercontexthelpkey?language=objc
type HelpManagerContextHelpKey string

// Constants used by imageAlignment that allow you to specify the location of the image in the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagealignment?language=objc
type ImageAlignment uint

const (
	ImageAlignBottom      ImageAlignment = 5
	ImageAlignBottomLeft  ImageAlignment = 6
	ImageAlignBottomRight ImageAlignment = 7
	ImageAlignCenter      ImageAlignment = 0
	ImageAlignLeft        ImageAlignment = 4
	ImageAlignRight       ImageAlignment = 8
	ImageAlignTop         ImageAlignment = 1
	ImageAlignTopLeft     ImageAlignment = 2
	ImageAlignTopRight    ImageAlignment = 3
)

// Constants that specify the caching policy on a per-image basis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecachemode?language=objc
type ImageCacheMode uint

const (
	ImageCacheAlways  ImageCacheMode = 1
	ImageCacheBySize  ImageCacheMode = 2
	ImageCacheDefault ImageCacheMode = 0
	ImageCacheNever   ImageCacheMode = 3
)

// Constants that allow you to specify the kind of frame bordering the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageframestyle?language=objc
type ImageFrameStyle uint

const (
	ImageFrameButton    ImageFrameStyle = 4
	ImageFrameGrayBezel ImageFrameStyle = 2
	ImageFrameGroove    ImageFrameStyle = 3
	ImageFrameNone      ImageFrameStyle = 0
	ImageFramePhoto     ImageFrameStyle = 1
)

// Constants for the keys to include in a hints dictionary when drawing the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagehintkey?language=objc
type ImageHintKey string

const (
	ImageHintCTM                          ImageHintKey = "NSImageHintCTM"
	ImageHintInterpolation                ImageHintKey = "NSImageHintInterpolation"
	ImageHintUserInterfaceLayoutDirection ImageHintKey = "NSImageHintUserInterfaceLayoutDirection"
)

// Constants that specify the interpolation, or image smoothing, behavior used by the image interpolation property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageinterpolation?language=objc
type ImageInterpolation uint

const (
	ImageInterpolationDefault ImageInterpolation = 0
	ImageInterpolationHigh    ImageInterpolation = 3
	ImageInterpolationLow     ImageInterpolation = 2
	ImageInterpolationMedium  ImageInterpolation = 4
	ImageInterpolationNone    ImageInterpolation = 1
)

// Constants that describe the layout direction for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagelayoutdirection?language=objc
type ImageLayoutDirection int

const (
	ImageLayoutDirectionLeftToRight ImageLayoutDirection = 2
	ImageLayoutDirectionRightToLeft ImageLayoutDirection = 3
	ImageLayoutDirectionUnspecified ImageLayoutDirection = -1
)

// Status values for incremental image loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageloadstatus?language=objc
type ImageLoadStatus uint

const (
	ImageLoadStatusCancelled     ImageLoadStatus = 1
	ImageLoadStatusCompleted     ImageLoadStatus = 0
	ImageLoadStatusInvalidData   ImageLoadStatus = 2
	ImageLoadStatusReadError     ImageLoadStatus = 4
	ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
)

// Named images, defined by the system or you, for use in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagename?language=objc
type ImageName string

const (
	ImageNameActionTemplate                          ImageName = "NSActionTemplate"
	ImageNameAddTemplate                             ImageName = "NSAddTemplate"
	ImageNameAdvanced                                ImageName = "NSAdvanced"
	ImageNameApplicationIcon                         ImageName = "NSApplicationIcon"
	ImageNameBluetoothTemplate                       ImageName = "NSBluetoothTemplate"
	ImageNameBonjour                                 ImageName = "NSBonjour"
	ImageNameBookmarksTemplate                       ImageName = "NSBookmarksTemplate"
	ImageNameCaution                                 ImageName = "NSCaution"
	ImageNameColorPanel                              ImageName = "NSColorPanel"
	ImageNameColumnViewTemplate                      ImageName = "NSColumnViewTemplate"
	ImageNameComputer                                ImageName = "NSComputer"
	ImageNameDotMac                                  ImageName = "NSDotMac"
	ImageNameEnterFullScreenTemplate                 ImageName = "NSEnterFullScreenTemplate"
	ImageNameEveryone                                ImageName = "NSEveryone"
	ImageNameExitFullScreenTemplate                  ImageName = "NSExitFullScreenTemplate"
	ImageNameFlowViewTemplate                        ImageName = "NSFlowViewTemplate"
	ImageNameFolder                                  ImageName = "NSFolder"
	ImageNameFolderBurnable                          ImageName = "NSFolderBurnable"
	ImageNameFolderSmart                             ImageName = "NSFolderSmart"
	ImageNameFollowLinkFreestandingTemplate          ImageName = "NSFollowLinkFreestandingTemplate"
	ImageNameFontPanel                               ImageName = "NSFontPanel"
	ImageNameGoBackTemplate                          ImageName = "NSGoBackTemplate"
	ImageNameGoForwardTemplate                       ImageName = "NSGoForwardTemplate"
	ImageNameGoLeftTemplate                          ImageName = "NSGoLeftTemplate"
	ImageNameGoRightTemplate                         ImageName = "NSGoRightTemplate"
	ImageNameHomeTemplate                            ImageName = "NSHomeTemplate"
	ImageNameIChatTheaterTemplate                    ImageName = "NSIChatTheaterTemplate"
	ImageNameIconViewTemplate                        ImageName = "NSIconViewTemplate"
	ImageNameInfo                                    ImageName = "NSInfo"
	ImageNameInvalidDataFreestandingTemplate         ImageName = "NSInvalidDataFreestandingTemplate"
	ImageNameLeftFacingTriangleTemplate              ImageName = "NSLeftFacingTriangleTemplate"
	ImageNameListViewTemplate                        ImageName = "NSListViewTemplate"
	ImageNameLockLockedTemplate                      ImageName = "NSLockLockedTemplate"
	ImageNameLockUnlockedTemplate                    ImageName = "NSLockUnlockedTemplate"
	ImageNameMenuMixedStateTemplate                  ImageName = "NSMenuMixedStateTemplate"
	ImageNameMenuOnStateTemplate                     ImageName = "NSMenuOnStateTemplate"
	ImageNameMobileMe                                ImageName = "NSMobileMe"
	ImageNameMultipleDocuments                       ImageName = "NSMultipleDocuments"
	ImageNameNetwork                                 ImageName = "NSNetwork"
	ImageNamePathTemplate                            ImageName = "NSPathTemplate"
	ImageNamePreferencesGeneral                      ImageName = "NSPreferencesGeneral"
	ImageNameQuickLookTemplate                       ImageName = "NSQuickLookTemplate"
	ImageNameRefreshFreestandingTemplate             ImageName = "NSRefreshFreestandingTemplate"
	ImageNameRefreshTemplate                         ImageName = "NSRefreshTemplate"
	ImageNameRemoveTemplate                          ImageName = "NSRemoveTemplate"
	ImageNameRevealFreestandingTemplate              ImageName = "NSRevealFreestandingTemplate"
	ImageNameRightFacingTriangleTemplate             ImageName = "NSRightFacingTriangleTemplate"
	ImageNameShareTemplate                           ImageName = "NSShareTemplate"
	ImageNameSlideshowTemplate                       ImageName = "NSSlideshowTemplate"
	ImageNameSmartBadgeTemplate                      ImageName = "NSSmartBadgeTemplate"
	ImageNameStatusAvailable                         ImageName = "NSStatusAvailable"
	ImageNameStatusNone                              ImageName = "NSStatusNone"
	ImageNameStatusPartiallyAvailable                ImageName = "NSStatusPartiallyAvailable"
	ImageNameStatusUnavailable                       ImageName = "NSStatusUnavailable"
	ImageNameStopProgressFreestandingTemplate        ImageName = "NSStopProgressFreestandingTemplate"
	ImageNameStopProgressTemplate                    ImageName = "NSStopProgressTemplate"
	ImageNameTouchBarAddDetailTemplate               ImageName = "NSTouchBarAddDetailTemplate"
	ImageNameTouchBarAddTemplate                     ImageName = "NSTouchBarAddTemplate"
	ImageNameTouchBarAlarmTemplate                   ImageName = "NSTouchBarAlarmTemplate"
	ImageNameTouchBarAudioInputMuteTemplate          ImageName = "NSTouchBarAudioInputMuteTemplate"
	ImageNameTouchBarAudioInputTemplate              ImageName = "NSTouchBarAudioInputTemplate"
	ImageNameTouchBarAudioOutputMuteTemplate         ImageName = "NSTouchBarAudioOutputMuteTemplate"
	ImageNameTouchBarAudioOutputVolumeHighTemplate   ImageName = "NSTouchBarAudioOutputVolumeHighTemplate"
	ImageNameTouchBarAudioOutputVolumeLowTemplate    ImageName = "NSTouchBarAudioOutputVolumeLowTemplate"
	ImageNameTouchBarAudioOutputVolumeMediumTemplate ImageName = "NSTouchBarAudioOutputVolumeMediumTemplate"
	ImageNameTouchBarAudioOutputVolumeOffTemplate    ImageName = "NSTouchBarAudioOutputVolumeOffTemplate"
	ImageNameTouchBarBookmarksTemplate               ImageName = "NSTouchBarBookmarksTemplate"
	ImageNameTouchBarColorPickerFill                 ImageName = "NSTouchBarColorPickerFill"
	ImageNameTouchBarColorPickerFont                 ImageName = "NSTouchBarColorPickerFont"
	ImageNameTouchBarColorPickerStroke               ImageName = "NSTouchBarColorPickerStroke"
	ImageNameTouchBarCommunicationAudioTemplate      ImageName = "NSTouchBarCommunicationAudioTemplate"
	ImageNameTouchBarCommunicationVideoTemplate      ImageName = "NSTouchBarCommunicationVideoTemplate"
	ImageNameTouchBarComposeTemplate                 ImageName = "NSTouchBarComposeTemplate"
	ImageNameTouchBarDeleteTemplate                  ImageName = "NSTouchBarDeleteTemplate"
	ImageNameTouchBarDownloadTemplate                ImageName = "NSTouchBarDownloadTemplate"
	ImageNameTouchBarEnterFullScreenTemplate         ImageName = "NSTouchBarEnterFullScreenTemplate"
	ImageNameTouchBarExitFullScreenTemplate          ImageName = "NSTouchBarExitFullScreenTemplate"
	ImageNameTouchBarFastForwardTemplate             ImageName = "NSTouchBarFastForwardTemplate"
	ImageNameTouchBarFolderCopyToTemplate            ImageName = "NSTouchBarFolderCopyToTemplate"
	ImageNameTouchBarFolderMoveToTemplate            ImageName = "NSTouchBarFolderMoveToTemplate"
	ImageNameTouchBarFolderTemplate                  ImageName = "NSTouchBarFolderTemplate"
	ImageNameTouchBarGetInfoTemplate                 ImageName = "NSTouchBarGetInfoTemplate"
	ImageNameTouchBarGoBackTemplate                  ImageName = "NSTouchBarGoBackTemplate"
	ImageNameTouchBarGoDownTemplate                  ImageName = "NSTouchBarGoDownTemplate"
	ImageNameTouchBarGoForwardTemplate               ImageName = "NSTouchBarGoForwardTemplate"
	ImageNameTouchBarGoUpTemplate                    ImageName = "NSTouchBarGoUpTemplate"
	ImageNameTouchBarHistoryTemplate                 ImageName = "NSTouchBarHistoryTemplate"
	ImageNameTouchBarIconViewTemplate                ImageName = "NSTouchBarIconViewTemplate"
	ImageNameTouchBarListViewTemplate                ImageName = "NSTouchBarListViewTemplate"
	ImageNameTouchBarMailTemplate                    ImageName = "NSTouchBarMailTemplate"
	ImageNameTouchBarNewFolderTemplate               ImageName = "NSTouchBarNewFolderTemplate"
	ImageNameTouchBarNewMessageTemplate              ImageName = "NSTouchBarNewMessageTemplate"
	ImageNameTouchBarOpenInBrowserTemplate           ImageName = "NSTouchBarOpenInBrowserTemplate"
	ImageNameTouchBarPauseTemplate                   ImageName = "NSTouchBarPauseTemplate"
	ImageNameTouchBarPlayPauseTemplate               ImageName = "NSTouchBarPlayPauseTemplate"
	ImageNameTouchBarPlayTemplate                    ImageName = "NSTouchBarPlayTemplate"
	ImageNameTouchBarPlayheadTemplate                ImageName = "NSTouchBarPlayheadTemplate"
	ImageNameTouchBarQuickLookTemplate               ImageName = "NSTouchBarQuickLookTemplate"
	ImageNameTouchBarRecordStartTemplate             ImageName = "NSTouchBarRecordStartTemplate"
	ImageNameTouchBarRecordStopTemplate              ImageName = "NSTouchBarRecordStopTemplate"
	ImageNameTouchBarRefreshTemplate                 ImageName = "NSTouchBarRefreshTemplate"
	ImageNameTouchBarRemoveTemplate                  ImageName = "NSTouchBarRemoveTemplate"
	ImageNameTouchBarRewindTemplate                  ImageName = "NSTouchBarRewindTemplate"
	ImageNameTouchBarRotateLeftTemplate              ImageName = "NSTouchBarRotateLeftTemplate"
	ImageNameTouchBarRotateRightTemplate             ImageName = "NSTouchBarRotateRightTemplate"
	ImageNameTouchBarSearchTemplate                  ImageName = "NSTouchBarSearchTemplate"
	ImageNameTouchBarShareTemplate                   ImageName = "NSTouchBarShareTemplate"
	ImageNameTouchBarSidebarTemplate                 ImageName = "NSTouchBarSidebarTemplate"
	ImageNameTouchBarSkipAhead15SecondsTemplate      ImageName = "NSTouchBarSkipAhead15SecondsTemplate"
	ImageNameTouchBarSkipAhead30SecondsTemplate      ImageName = "NSTouchBarSkipAhead30SecondsTemplate"
	ImageNameTouchBarSkipAheadTemplate               ImageName = "NSTouchBarSkipAheadTemplate"
	ImageNameTouchBarSkipBack15SecondsTemplate       ImageName = "NSTouchBarSkipBack15SecondsTemplate"
	ImageNameTouchBarSkipBack30SecondsTemplate       ImageName = "NSTouchBarSkipBack30SecondsTemplate"
	ImageNameTouchBarSkipBackTemplate                ImageName = "NSTouchBarSkipBackTemplate"
	ImageNameTouchBarSkipToEndTemplate               ImageName = "NSTouchBarSkipToEndTemplate"
	ImageNameTouchBarSkipToStartTemplate             ImageName = "NSTouchBarSkipToStartTemplate"
	ImageNameTouchBarSlideshowTemplate               ImageName = "NSTouchBarSlideshowTemplate"
	ImageNameTouchBarTagIconTemplate                 ImageName = "NSTouchBarTagIconTemplate"
	ImageNameTouchBarTextBoldTemplate                ImageName = "NSTouchBarTextBoldTemplate"
	ImageNameTouchBarTextBoxTemplate                 ImageName = "NSTouchBarTextBoxTemplate"
	ImageNameTouchBarTextCenterAlignTemplate         ImageName = "NSTouchBarTextCenterAlignTemplate"
	ImageNameTouchBarTextItalicTemplate              ImageName = "NSTouchBarTextItalicTemplate"
	ImageNameTouchBarTextJustifiedAlignTemplate      ImageName = "NSTouchBarTextJustifiedAlignTemplate"
	ImageNameTouchBarTextLeftAlignTemplate           ImageName = "NSTouchBarTextLeftAlignTemplate"
	ImageNameTouchBarTextListTemplate                ImageName = "NSTouchBarTextListTemplate"
	ImageNameTouchBarTextRightAlignTemplate          ImageName = "NSTouchBarTextRightAlignTemplate"
	ImageNameTouchBarTextStrikethroughTemplate       ImageName = "NSTouchBarTextStrikethroughTemplate"
	ImageNameTouchBarTextUnderlineTemplate           ImageName = "NSTouchBarTextUnderlineTemplate"
	ImageNameTouchBarUserAddTemplate                 ImageName = "NSTouchBarUserAddTemplate"
	ImageNameTouchBarUserGroupTemplate               ImageName = "NSTouchBarUserGroupTemplate"
	ImageNameTouchBarUserTemplate                    ImageName = "NSTouchBarUserTemplate"
	ImageNameTouchBarVolumeDownTemplate              ImageName = "NSTouchBarVolumeDownTemplate"
	ImageNameTouchBarVolumeUpTemplate                ImageName = "NSTouchBarVolumeUpTemplate"
	ImageNameTrashEmpty                              ImageName = "NSTrashEmpty"
	ImageNameTrashFull                               ImageName = "NSTrashFull"
	ImageNameUser                                    ImageName = "NSUser"
	ImageNameUserAccounts                            ImageName = "NSUserAccounts"
	ImageNameUserGroup                               ImageName = "NSUserGroup"
	ImageNameUserGuest                               ImageName = "NSUserGuest"
)

// Constants that identify the loading status of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagereploadstatus?language=objc
type ImageRepLoadStatus int

const (
	ImageRepLoadStatusCompleted       ImageRepLoadStatus = -6
	ImageRepLoadStatusInvalidData     ImageRepLoadStatus = -4
	ImageRepLoadStatusReadingHeader   ImageRepLoadStatus = -2
	ImageRepLoadStatusUnexpectedEOF   ImageRepLoadStatus = -5
	ImageRepLoadStatusUnknownType     ImageRepLoadStatus = -1
	ImageRepLoadStatusWillNeedAllData ImageRepLoadStatus = -3
)

// Constants that describe the resizing mode for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageresizingmode?language=objc
type ImageResizingMode int

const ()

// Constants that specify a cell’s image scaling behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagescaling?language=objc
type ImageScaling uint

const (
	ImageScaleAxesIndependently      ImageScaling = 1
	ImageScaleNone                   ImageScaling = 2
	ImageScaleProportionallyDown     ImageScaling = 0
	ImageScaleProportionallyUpOrDown ImageScaling = 3
	ScaleNone                        ImageScaling = 2
	ScaleProportionally              ImageScaling = 0
	ScaleToFit                       ImageScaling = 1
)

// Constants that specify which scale variant of a symbol image to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolscale?language=objc
type ImageSymbolScale int

const (
	ImageSymbolScaleLarge  ImageSymbolScale = 3
	ImageSymbolScaleMedium ImageSymbolScale = 2
	ImageSymbolScaleSmall  ImageSymbolScale = 1
)

// These constants are used in NSResponder’s interfaceStyle method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsinterfacestyle?language=objc
type InterfaceStyle uint

// The part of the object’s visual representation that should be used to get the value for the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutattribute?language=objc
type LayoutAttribute int

const (
	LayoutAttributeBaseline       LayoutAttribute = 11
	LayoutAttributeBottom         LayoutAttribute = 4
	LayoutAttributeCenterX        LayoutAttribute = 9
	LayoutAttributeCenterY        LayoutAttribute = 10
	LayoutAttributeFirstBaseline  LayoutAttribute = 12
	LayoutAttributeHeight         LayoutAttribute = 8
	LayoutAttributeLastBaseline   LayoutAttribute = 11
	LayoutAttributeLeading        LayoutAttribute = 5
	LayoutAttributeLeft           LayoutAttribute = 1
	LayoutAttributeNotAnAttribute LayoutAttribute = 0
	LayoutAttributeRight          LayoutAttribute = 2
	LayoutAttributeTop            LayoutAttribute = 3
	LayoutAttributeTrailing       LayoutAttribute = 6
	LayoutAttributeWidth          LayoutAttribute = 7
)

// The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutconstraintorientation?language=objc
type LayoutConstraintOrientation int

const (
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	LayoutConstraintOrientationVertical   LayoutConstraintOrientation = 1
)

// A bit mask that specifies both a part of an interface element to align and a direction for the alignment between two interface elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutformatoptions?language=objc
type LayoutFormatOptions uint

const (
	LayoutFormatAlignAllBaseline           LayoutFormatOptions = 2048
	LayoutFormatAlignAllBottom             LayoutFormatOptions = 16
	LayoutFormatAlignAllCenterX            LayoutFormatOptions = 512
	LayoutFormatAlignAllCenterY            LayoutFormatOptions = 1024
	LayoutFormatAlignAllFirstBaseline      LayoutFormatOptions = 4096
	LayoutFormatAlignAllLastBaseline       LayoutFormatOptions = 2048
	LayoutFormatAlignAllLeading            LayoutFormatOptions = 32
	LayoutFormatAlignAllLeft               LayoutFormatOptions = 2
	LayoutFormatAlignAllRight              LayoutFormatOptions = 4
	LayoutFormatAlignAllTop                LayoutFormatOptions = 8
	LayoutFormatAlignAllTrailing           LayoutFormatOptions = 64
	LayoutFormatAlignmentMask              LayoutFormatOptions = 65535
	LayoutFormatDirectionLeadingToTrailing LayoutFormatOptions = 0
	LayoutFormatDirectionLeftToRight       LayoutFormatOptions = 65536
	LayoutFormatDirectionMask              LayoutFormatOptions = 196608
	LayoutFormatDirectionRightToLeft       LayoutFormatOptions = 131072
)

// Layout priority used to indicate the relative importance of constraints, allowing Auto Layout to make appropriate tradeoffs when satisfying the constraints of the system as a whole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutpriority?language=objc
type LayoutPriority float64

const (
	LayoutPriorityDefaultHigh                LayoutPriority = 750.000000
	LayoutPriorityDefaultLow                 LayoutPriority = 250.000000
	LayoutPriorityDragThatCanResizeWindow    LayoutPriority = 510.000000
	LayoutPriorityDragThatCannotResizeWindow LayoutPriority = 490.000000
	LayoutPriorityFittingSizeCompression     LayoutPriority = 50.000000
	LayoutPriorityRequired                   LayoutPriority = 1000.000000
	LayoutPriorityWindowSizeStayPut          LayoutPriority = 500.000000
)

// The relation between the first attribute and the modified second attribute in a constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutrelation?language=objc
type LayoutRelation int

const (
	LayoutRelationEqual              LayoutRelation = 0
	LayoutRelationGreaterThanOrEqual LayoutRelation = 1
	LayoutRelationLessThanOrEqual    LayoutRelation = -1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorplaceholdervisibility?language=objc
type LevelIndicatorPlaceholderVisibility int

const (
	LevelIndicatorPlaceholderVisibilityAlways       LevelIndicatorPlaceholderVisibility = 1
	LevelIndicatorPlaceholderVisibilityAutomatic    LevelIndicatorPlaceholderVisibility = 0
	LevelIndicatorPlaceholderVisibilityWhileEditing LevelIndicatorPlaceholderVisibility = 2
)

// Constants that specify a level indicator's appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorstyle?language=objc
type LevelIndicatorStyle uint

const (
	ContinuousCapacityLevelIndicatorStyle LevelIndicatorStyle = 1
	DiscreteCapacityLevelIndicatorStyle   LevelIndicatorStyle = 2
	LevelIndicatorStyleContinuousCapacity LevelIndicatorStyle = 1
	LevelIndicatorStyleDiscreteCapacity   LevelIndicatorStyle = 2
	LevelIndicatorStyleRating             LevelIndicatorStyle = 3
	LevelIndicatorStyleRelevancy          LevelIndicatorStyle = 0
	RatingLevelIndicatorStyle             LevelIndicatorStyle = 3
	RelevancyLevelIndicatorStyle          LevelIndicatorStyle = 0
)

// Constants that specify what happens when a line is too long for a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslinebreakmode?language=objc
type LineBreakMode int

const (
	LineBreakByCharWrapping     LineBreakMode = 1
	LineBreakByClipping         LineBreakMode = 2
	LineBreakByTruncatingHead   LineBreakMode = 3
	LineBreakByTruncatingMiddle LineBreakMode = 5
	LineBreakByTruncatingTail   LineBreakMode = 4
	LineBreakByWordWrapping     LineBreakMode = 0
)

// Constants that specify how the text system breaks lines while laying out paragraphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslinebreakstrategy?language=objc
type LineBreakStrategy uint

const (
	LineBreakStrategyHangulWordPriority LineBreakStrategy = 2
	LineBreakStrategyNone               LineBreakStrategy = 0
	LineBreakStrategyPushOut            LineBreakStrategy = 1
	LineBreakStrategyStandard           LineBreakStrategy = 65535
)

// Constants that specify the shape of endpoints for an open path when it is stroked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslinecapstyle?language=objc
type LineCapStyle uint

const (
	ButtLineCapStyle   LineCapStyle = 0
	LineCapStyleButt   LineCapStyle = 0
	LineCapStyleRound  LineCapStyle = 1
	LineCapStyleSquare LineCapStyle = 2
	RoundLineCapStyle  LineCapStyle = 1
	SquareLineCapStyle LineCapStyle = 2
)

// Constants that specify the shape of the joins between connected segments of a stroked path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslinejoinstyle?language=objc
type LineJoinStyle uint

const (
	BevelLineJoinStyle LineJoinStyle = 2
	LineJoinStyleBevel LineJoinStyle = 2
	LineJoinStyleMiter LineJoinStyle = 0
	LineJoinStyleRound LineJoinStyle = 1
	MiterLineJoinStyle LineJoinStyle = 0
	RoundLineJoinStyle LineJoinStyle = 1
)

// The direction in which a line moves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslinemovementdirection?language=objc
type LineMovementDirection uint

const (
	LineDoesntMove LineMovementDirection = 0
	LineMovesDown  LineMovementDirection = 3
	LineMovesLeft  LineMovementDirection = 1
	LineMovesRight LineMovementDirection = 2
	LineMovesUp    LineMovementDirection = 4
)

// Values that describe the progression of text on a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslinesweepdirection?language=objc
type LineSweepDirection uint

const (
	LineSweepDown  LineSweepDirection = 2
	LineSweepLeft  LineSweepDirection = 0
	LineSweepRight LineSweepDirection = 1
	LineSweepUp    LineSweepDirection = 3
)

// These constants determine how NSCell objects behave when an NSMatrix object is tracking the mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrixmode?language=objc
type MatrixMode uint

const (
	HighlightModeMatrix MatrixMode = 1
	ListModeMatrix      MatrixMode = 2
	RadioModeMatrix     MatrixMode = 0
	TrackModeMatrix     MatrixMode = 3
)

// These constants are masks used to configure a Media Library Browser to display specific types of media. Combined masks are not yet supported.  In other words, only one nonzero mask value is supported at a time.  If masks are combined, the lowest mask value is used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrary?language=objc
type MediaLibrary uint

const (
	MediaLibraryAudio MediaLibrary = 1
	MediaLibraryImage MediaLibrary = 2
	MediaLibraryMovie MediaLibrary = 4
)

// These constants are used as a bitmask for specifying a set of menu or menu item properties, and are contained by the propertiesToUpdate property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuproperties?language=objc
type MenuProperties uint

const (
	MenuPropertyItemAccessibilityDescription MenuProperties = 32
	MenuPropertyItemAttributedTitle          MenuProperties = 2
	MenuPropertyItemEnabled                  MenuProperties = 16
	MenuPropertyItemImage                    MenuProperties = 8
	MenuPropertyItemKeyEquivalent            MenuProperties = 4
	MenuPropertyItemTitle                    MenuProperties = 1
)

// A set of button return values for modal dialogs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmodalresponse?language=objc
type ModalResponse int

const (
	AlertFirstButtonReturn  ModalResponse = 1000
	AlertSecondButtonReturn ModalResponse = 1001
	AlertThirdButtonReturn  ModalResponse = 1002
	ModalResponseAbort      ModalResponse = -1001
	ModalResponseCancel     ModalResponse = 0
	ModalResponseContinue   ModalResponse = -1002
	ModalResponseOK         ModalResponse = 1
	ModalResponseStop       ModalResponse = -1000
)

// A constant for glyph packing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmultibyteglyphpacking?language=objc
type MultibyteGlyphPacking uint

const (
	NativeShortGlyphPacking MultibyteGlyphPacking = 5
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibname?language=objc
type NibName string

// Constants that specify context parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglcontextparameter?language=objc
type OpenGLContextParameter int

const (
	OpenGLCPCurrentRendererID                    OpenGLContextParameter = 309
	OpenGLCPGPUFragmentProcessing                OpenGLContextParameter = 311
	OpenGLCPGPUVertexProcessing                  OpenGLContextParameter = 310
	OpenGLCPHasDrawable                          OpenGLContextParameter = 314
	OpenGLCPMPSwapsInFlight                      OpenGLContextParameter = 315
	OpenGLCPRasterizationEnable                  OpenGLContextParameter = 221
	OpenGLCPReclaimResources                     OpenGLContextParameter = 308
	OpenGLCPStateValidation                      OpenGLContextParameter = 301
	OpenGLCPSurfaceBackingSize                   OpenGLContextParameter = 304
	OpenGLCPSurfaceOpacity                       OpenGLContextParameter = 236
	OpenGLCPSurfaceOrder                         OpenGLContextParameter = 235
	OpenGLCPSurfaceSurfaceVolatile               OpenGLContextParameter = 306
	OpenGLCPSwapInterval                         OpenGLContextParameter = 222
	OpenGLCPSwapRectangle                        OpenGLContextParameter = 200
	OpenGLCPSwapRectangleEnable                  OpenGLContextParameter = 201
	OpenGLContextParameterCurrentRendererID      OpenGLContextParameter = 309
	OpenGLContextParameterGPUFragmentProcessing  OpenGLContextParameter = 311
	OpenGLContextParameterGPUVertexProcessing    OpenGLContextParameter = 310
	OpenGLContextParameterHasDrawable            OpenGLContextParameter = 314
	OpenGLContextParameterMPSwapsInFlight        OpenGLContextParameter = 315
	OpenGLContextParameterRasterizationEnable    OpenGLContextParameter = 221
	OpenGLContextParameterReclaimResources       OpenGLContextParameter = 308
	OpenGLContextParameterStateValidation        OpenGLContextParameter = 301
	OpenGLContextParameterSurfaceBackingSize     OpenGLContextParameter = 304
	OpenGLContextParameterSurfaceOpacity         OpenGLContextParameter = 236
	OpenGLContextParameterSurfaceOrder           OpenGLContextParameter = 235
	OpenGLContextParameterSurfaceSurfaceVolatile OpenGLContextParameter = 306
	OpenGLContextParameterSwapInterval           OpenGLContextParameter = 222
	OpenGLContextParameterSwapRectangle          OpenGLContextParameter = 200
	OpenGLContextParameterSwapRectangleEnable    OpenGLContextParameter = 201
)

// Constants that specify OpenGL options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglglobaloption?language=objc
type OpenGLGlobalOption uint32

const (
	OpenGLGOClearFormatCache OpenGLGlobalOption = 502
	OpenGLGOFormatCacheSize  OpenGLGlobalOption = 501
	OpenGLGOResetLibrary     OpenGLGlobalOption = 504
	OpenGLGORetainRenderers  OpenGLGlobalOption = 503
	OpenGLGOUseBuildCache    OpenGLGlobalOption = 506
)

// Pixel format attributes for OpenGL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglpixelformatattribute?language=objc
type OpenGLPixelFormatAttribute uint32

// Constants used to configure the contents of a PDF panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpaneloptions?language=objc
type PDFPanelOptions int

const (
	PDFPanelRequestsParentDirectory PDFPanelOptions = 16777216
	PDFPanelShowsOrientation        PDFPanelOptions = 8
	PDFPanelShowsPaperSize          PDFPanelOptions = 4
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerobjectidentifier?language=objc
type PageControllerObjectIdentifier string

// These constants control the transition style of the page controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollertransitionstyle?language=objc
type PageControllerTransitionStyle int

const (
	PageControllerTransitionStyleHorizontalStrip PageControllerTransitionStyle = 2
	PageControllerTransitionStyleStackBook       PageControllerTransitionStyle = 1
	PageControllerTransitionStyleStackHistory    PageControllerTransitionStyle = 0
)

// Constants that describe the orientation of printing on a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspaperorientation?language=objc
type PaperOrientation int

const (
	PaperOrientationLandscape PaperOrientation = 1
	PaperOrientationPortrait  PaperOrientation = 0
)

// Options for preparing the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardcontentsoptions?language=objc
type PasteboardContentsOptions uint

const (
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
)

// Constants that represent the standard pasteboard names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardname?language=objc
type PasteboardName string

const (
	DragPboard            PasteboardName = "Apple CFPasteboard drag"
	FindPboard            PasteboardName = "Apple CFPasteboard find"
	FontPboard            PasteboardName = "Apple CFPasteboard font"
	GeneralPboard         PasteboardName = "Apple CFPasteboard general"
	PasteboardNameDrag    PasteboardName = "Apple CFPasteboard drag"
	PasteboardNameFind    PasteboardName = "Apple CFPasteboard find"
	PasteboardNameFont    PasteboardName = "Apple CFPasteboard font"
	PasteboardNameGeneral PasteboardName = "Apple CFPasteboard general"
	PasteboardNameRuler   PasteboardName = "Apple CFPasteboard ruler"
	RulerPboard           PasteboardName = "Apple CFPasteboard ruler"
)

// Options for reading pasteboard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreadingoptionkey?language=objc
type PasteboardReadingOptionKey string

const (
	PasteboardURLReadingContentsConformToTypesKey PasteboardReadingOptionKey = "NSPasteboardURLReadingContentsConformToTypesKey"
	PasteboardURLReadingFileURLsOnlyKey           PasteboardReadingOptionKey = "NSPasteboardURLReadingFileURLsOnlyKey"
)

// Options that specify how to interpret data on the pasteboard when initializing pasteboard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreadingoptions?language=objc
type PasteboardReadingOptions uint

const (
	PasteboardReadingAsData         PasteboardReadingOptions = 0
	PasteboardReadingAsKeyedArchive PasteboardReadingOptions = 4
	PasteboardReadingAsPropertyList PasteboardReadingOptions = 2
	PasteboardReadingAsString       PasteboardReadingOptions = 1
)

// The supported pasteboard types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtype?language=objc
type PasteboardType string

const (
	ColorPboardType                      PasteboardType = "NSColor pasteboard type"
	FileContentsPboardType               PasteboardType = "NXFileContentsPboardType"
	FilenamesPboardType                  PasteboardType = "NSFilenamesPboardType"
	FilesPromisePboardType               PasteboardType = "Apple files promise pasteboard type"
	FindPanelSearchOptionsPboardType     PasteboardType = "NSFindPanel search options pasteboard type"
	FontPboardType                       PasteboardType = "NeXT font pasteboard type"
	HTMLPboardType                       PasteboardType = "Apple HTML pasteboard type"
	InkTextPboardType                    PasteboardType = "Apple InkText pasteboard type"
	MultipleTextSelectionPboardType      PasteboardType = "Apple multiple text selection pasteboard type"
	PDFPboardType                        PasteboardType = "Apple PDF pasteboard type"
	PICTPboardType                       PasteboardType = "Apple PICT pasteboard type"
	PasteboardTypeColor                  PasteboardType = "com.apple.cocoa.pasteboard.color"
	PasteboardTypeFileURL                PasteboardType = "public.file-url"
	PasteboardTypeFindPanelSearchOptions PasteboardType = "com.apple.cocoa.pasteboard.find-panel-search-options"
	PasteboardTypeFont                   PasteboardType = "com.apple.cocoa.pasteboard.character-formatting"
	PasteboardTypeHTML                   PasteboardType = "public.html"
	PasteboardTypeMultipleTextSelection  PasteboardType = "com.apple.cocoa.pasteboard.multiple-text-selection"
	PasteboardTypePDF                    PasteboardType = "com.adobe.pdf"
	PasteboardTypePNG                    PasteboardType = "public.png"
	PasteboardTypeRTF                    PasteboardType = "public.rtf"
	PasteboardTypeRTFD                   PasteboardType = "com.apple.flat-rtfd"
	PasteboardTypeRuler                  PasteboardType = "com.apple.cocoa.pasteboard.paragraph-formatting"
	PasteboardTypeSound                  PasteboardType = "com.apple.cocoa.pasteboard.sound"
	PasteboardTypeString                 PasteboardType = "public.utf8-plain-text"
	PasteboardTypeTIFF                   PasteboardType = "public.tiff"
	PasteboardTypeTabularText            PasteboardType = "public.utf8-tab-separated-values-text"
	PasteboardTypeTextFinderOptions      PasteboardType = "com.apple.cocoa.pasteboard.find-panel-search-options"
	PasteboardTypeURL                    PasteboardType = "public.url"
	PostScriptPboardType                 PasteboardType = "NeXT Encapsulated PostScript v1.2 pasteboard type"
	RTFDPboardType                       PasteboardType = "NeXT RTFD pasteboard type"
	RTFPboardType                        PasteboardType = "NeXT Rich Text Format v1.0 pasteboard type"
	RulerPboardType                      PasteboardType = "NeXT ruler pasteboard type"
	SoundPboardType                      PasteboardType = "NSSoundPboardType"
	StringPboardType                     PasteboardType = "NSStringPboardType"
	TIFFPboardType                       PasteboardType = "NeXT TIFF v4.0 pasteboard type"
	TabularTextPboardType                PasteboardType = "NeXT tabular text pasteboard type"
	URLPboardType                        PasteboardType = "Apple URL pasteboard type"
	VCardPboardType                      PasteboardType = "Apple VCard pasteboard type"
)

// Search options for the find panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtypefindpanelsearchoptionkey?language=objc
type PasteboardTypeFindPanelSearchOptionKey string

const (
	FindPanelCaseInsensitiveSearch PasteboardTypeFindPanelSearchOptionKey = "NSFindPanelCaseInsensitiveSearch"
	FindPanelSubstringMatch        PasteboardTypeFindPanelSearchOptionKey = "NSFindPanelSubstringMatch"
)

// Search options for text in Finder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtypetextfinderoptionkey?language=objc
type PasteboardTypeTextFinderOptionKey string

const (
	TextFinderCaseInsensitiveKey PasteboardTypeTextFinderOptionKey = "NSFindPanelCaseInsensitiveSearch"
	TextFinderMatchingTypeKey    PasteboardTypeTextFinderOptionKey = "NSFindPanelSubstringMatch"
)

// Type to specify options for writing to a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardwritingoptions?language=objc
type PasteboardWritingOptions uint

const (
	PasteboardWritingPromised PasteboardWritingOptions = 512
)

// NSPathStyle constants represent the different visual and behavioral styles an NSPathControl or NSPathCell object can have. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathstyle?language=objc
type PathStyle int

const (
	PathStyleNavigationBar PathStyle = 1
	PathStylePopUp         PathStyle = 2
	PathStyleStandard      PathStyle = 0
)

// Constants that specify display styles for picker bar items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritemcontrolrepresentation?language=objc
type PickerTouchBarItemControlRepresentation int

const (
	PickerTouchBarItemControlRepresentationAutomatic PickerTouchBarItemControlRepresentation = 0
	PickerTouchBarItemControlRepresentationCollapsed PickerTouchBarItemControlRepresentation = 2
	PickerTouchBarItemControlRepresentationExpanded  PickerTouchBarItemControlRepresentation = 1
)

// Constants that specify selection modes for picker bar items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritemselectionmode?language=objc
type PickerTouchBarItemSelectionMode int

const (
	PickerTouchBarItemSelectionModeMomentary PickerTouchBarItemSelectionMode = 2
	PickerTouchBarItemSelectionModeSelectAny PickerTouchBarItemSelectionMode = 1
	PickerTouchBarItemSelectionModeSelectOne PickerTouchBarItemSelectionMode = 0
)

// The pointing-device types for tablet-proximity events or mouse events with a proximity event subtype. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspointingdevicetype?language=objc
type PointingDeviceType uint

const (
	CursorPointingDevice      PointingDeviceType = 2
	EraserPointingDevice      PointingDeviceType = 3
	PenPointingDevice         PointingDeviceType = 1
	PointingDeviceTypeCursor  PointingDeviceType = 2
	PointingDeviceTypeEraser  PointingDeviceType = 3
	PointingDeviceTypePen     PointingDeviceType = 1
	PointingDeviceTypeUnknown PointingDeviceType = 0
	UnknownPointingDevice     PointingDeviceType = 0
)

// These constants are defined for use with the arrowPosition property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopuparrowposition?language=objc
type PopUpArrowPosition uint

const (
	PopUpArrowAtBottom PopUpArrowPosition = 2
	PopUpArrowAtCenter PopUpArrowPosition = 1
	PopUpNoArrow       PopUpArrowPosition = 0
)

// The set of predefined appearances for a popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverappearance?language=objc
type PopoverAppearance int

const (
	PopoverAppearanceHUD     PopoverAppearance = 1
	PopoverAppearanceMinimal PopoverAppearance = 0
)

// The appearance and disappearance behavior of a popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverbehavior?language=objc
type PopoverBehavior int

const (
	PopoverBehaviorApplicationDefined PopoverBehavior = 0
	PopoverBehaviorSemitransient      PopoverBehavior = 2
	PopoverBehaviorTransient          PopoverBehavior = 1
)

// Values that specify the reason for the NSPopoverWillCloseNotification notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverclosereasonvalue?language=objc
type PopoverCloseReasonValue string

const (
	PopoverCloseReasonDetachToWindow PopoverCloseReasonValue = "NSPopoverCloseReasonDetachToWindow"
	PopoverCloseReasonStandard       PopoverCloseReasonValue = "NSPopoverCloseReasonStandard"
)

// These constants describe the behavior and progression of a pressure gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressurebehavior?language=objc
type PressureBehavior int

const (
	PressureBehaviorPrimaryAccelerator PressureBehavior = 3
	PressureBehaviorPrimaryClick       PressureBehavior = 1
	PressureBehaviorPrimaryDeepClick   PressureBehavior = 5
	PressureBehaviorPrimaryDeepDrag    PressureBehavior = 6
	PressureBehaviorPrimaryDefault     PressureBehavior = 0
	PressureBehaviorPrimaryGeneric     PressureBehavior = 2
	PressureBehaviorUnknown            PressureBehavior = -1
)

// Constants that specify print job attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfoattributekey?language=objc
type PrintInfoAttributeKey string

const (
	PrintAllPages                         PrintInfoAttributeKey = "NSPrintAllPages"
	PrintBottomMargin                     PrintInfoAttributeKey = "NSBottomMargin"
	PrintCopies                           PrintInfoAttributeKey = "NSCopies"
	PrintDetailedErrorReporting           PrintInfoAttributeKey = "NSDetailedErrorReporting"
	PrintFaxNumber                        PrintInfoAttributeKey = "NSFaxNumber"
	PrintFirstPage                        PrintInfoAttributeKey = "NSFirstPage"
	PrintHeaderAndFooter                  PrintInfoAttributeKey = "NSPrintHeaderAndFooter"
	PrintHorizontalPagination             PrintInfoAttributeKey = "NSHorizonalPagination"
	PrintHorizontallyCentered             PrintInfoAttributeKey = "NSHorizontallyCentered"
	PrintJobDisposition                   PrintInfoAttributeKey = "NSJobDisposition"
	PrintJobSavingFileNameExtensionHidden PrintInfoAttributeKey = "NSJobSavingFileNameExtensionHidden"
	PrintJobSavingURL                     PrintInfoAttributeKey = "NSJobSavingURL"
	PrintLastPage                         PrintInfoAttributeKey = "NSLastPage"
	PrintLeftMargin                       PrintInfoAttributeKey = "NSLeftMargin"
	PrintMustCollate                      PrintInfoAttributeKey = "NSMustCollate"
	PrintOrientation                      PrintInfoAttributeKey = "NSOrientation"
	PrintPagesAcross                      PrintInfoAttributeKey = "NSPagesAcross"
	PrintPagesDown                        PrintInfoAttributeKey = "NSPagesDown"
	PrintPaperName                        PrintInfoAttributeKey = "NSPaperName"
	PrintPaperSize                        PrintInfoAttributeKey = "NSPaperSize"
	PrintPrinter                          PrintInfoAttributeKey = "NSPrinter"
	PrintPrinterName                      PrintInfoAttributeKey = "NSPrinterName"
	PrintReversePageOrder                 PrintInfoAttributeKey = "NSReversePageOrder"
	PrintRightMargin                      PrintInfoAttributeKey = "NSRightMargin"
	PrintScalingFactor                    PrintInfoAttributeKey = "NSScalingFactor"
	PrintSelectionOnly                    PrintInfoAttributeKey = "NSPrintSelectionOnly"
	PrintTime                             PrintInfoAttributeKey = "NSPrintTime"
	PrintTopMargin                        PrintInfoAttributeKey = "NSTopMargin"
	PrintVerticalPagination               PrintInfoAttributeKey = "NSVerticalPagination"
	PrintVerticallyCentered               PrintInfoAttributeKey = "NSVerticallyCentered"
)

// The type you use to specify a print info setting key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfosettingkey?language=objc
type PrintInfoSettingKey string

// Constants that specify values for the print job disposition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintjobdispositionvalue?language=objc
type PrintJobDispositionValue string

const (
	PrintCancelJob  PrintJobDispositionValue = "NSPrintCancelJob"
	PrintPreviewJob PrintJobDispositionValue = "NSPrintPreviewJob"
	PrintSaveJob    PrintJobDispositionValue = "NSPrintSaveJob"
	PrintSpoolJob   PrintJobDispositionValue = "NSPrintSpoolJob"
)

// Constants that specify the accessory panel keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpanelaccessorysummarykey?language=objc
type PrintPanelAccessorySummaryKey string

const (
	PrintPanelAccessorySummaryItemDescriptionKey PrintPanelAccessorySummaryKey = "description"
	PrintPanelAccessorySummaryItemNameKey        PrintPanelAccessorySummaryKey = "name"
)

// Constants that specify job style hints for activating the simplified Print panel interface and setting the options to display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpaneljobstylehint?language=objc
type PrintPanelJobStyleHint string

const (
	PrintAllPresetsJobStyleHint PrintPanelJobStyleHint = "All"
	PrintNoPresetsJobStyleHint  PrintPanelJobStyleHint = "None"
	PrintPhotoJobStyleHint      PrintPanelJobStyleHint = "Photo"
)

// Constants that specify options for configuring the contents of the main Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintpaneloptions?language=objc
type PrintPanelOptions uint

const (
	PrintPanelShowsCopies             PrintPanelOptions = 1
	PrintPanelShowsOrientation        PrintPanelOptions = 8
	PrintPanelShowsPageRange          PrintPanelOptions = 2
	PrintPanelShowsPageSetupAccessory PrintPanelOptions = 256
	PrintPanelShowsPaperSize          PrintPanelOptions = 4
	PrintPanelShowsPreview            PrintPanelOptions = 131072
	PrintPanelShowsPrintSelection     PrintPanelOptions = 32
	PrintPanelShowsScaling            PrintPanelOptions = 16
)

// Constants that specify the print quality in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintrenderingquality?language=objc
type PrintRenderingQuality int

const (
	PrintRenderingQualityBest       PrintRenderingQuality = 0
	PrintRenderingQualityResponsive PrintRenderingQuality = 1
)

// The type you use to specify the name of a type of paper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinterpapername?language=objc
type PrinterPaperName string

// Constants that describe the state of a printer information table stored by a printer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintertablestatus?language=objc
type PrinterTableStatus uint

const (
	PrinterTableError    PrinterTableStatus = 2
	PrinterTableNotFound PrinterTableStatus = 1
	PrinterTableOK       PrinterTableStatus = 0
)

// The type you use to describe a printer’s make and model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintertypename?language=objc
type PrinterTypeName string

// Constants that specify page orientations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintingorientation?language=objc
type PrintingOrientation uint

const (
	LandscapeOrientation PrintingOrientation = 1
	PortraitOrientation  PrintingOrientation = 0
)

// Constants that specify the page order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintingpageorder?language=objc
type PrintingPageOrder int

const (
	AscendingPageOrder  PrintingPageOrder = 1
	DescendingPageOrder PrintingPageOrder = -1
	SpecialPageOrder    PrintingPageOrder = 0
	UnknownPageOrder    PrintingPageOrder = 2
)

// Constants that specify the different ways in which an image is divided into pages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintingpaginationmode?language=objc
type PrintingPaginationMode uint

const (
	AutoPagination                  PrintingPaginationMode = 0
	ClipPagination                  PrintingPaginationMode = 2
	FitPagination                   PrintingPaginationMode = 1
	PrintingPaginationModeAutomatic PrintingPaginationMode = 0
	PrintingPaginationModeClip      PrintingPaginationMode = 2
	PrintingPaginationModeFit       PrintingPaginationMode = 1
)

// Constants that specify the progress indicator’s style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicatorstyle?language=objc
type ProgressIndicatorStyle uint

const (
	ProgressIndicatorStyleBar      ProgressIndicatorStyle = 0
	ProgressIndicatorStyleSpinning ProgressIndicatorStyle = 1
)

// Specify the height of a progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprogressindicatorthickness?language=objc
type ProgressIndicatorThickness uint

const (
	ProgressIndicatorPreferredAquaThickness  ProgressIndicatorThickness = 12
	ProgressIndicatorPreferredLargeThickness ProgressIndicatorThickness = 18
	ProgressIndicatorPreferredSmallThickness ProgressIndicatorThickness = 10
	ProgressIndicatorPreferredThickness      ProgressIndicatorThickness = 14
)

// Constants that specify alignment to an edge or a set of edges depending on the user interface layout direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsrectalignment?language=objc
type RectAlignment int

const (
	RectAlignmentBottom         RectAlignment = 5
	RectAlignmentBottomLeading  RectAlignment = 4
	RectAlignmentBottomTrailing RectAlignment = 6
	RectAlignmentLeading        RectAlignment = 3
	RectAlignmentNone           RectAlignment = 0
	RectAlignmentTop            RectAlignment = 1
	RectAlignmentTopLeading     RectAlignment = 2
	RectAlignmentTopTrailing    RectAlignment = 8
	RectAlignmentTrailing       RectAlignment = 7
)

// These constants determine whether apps launched by remote notifications display a badge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsremotenotificationtype?language=objc
type RemoteNotificationType uint

const (
	RemoteNotificationTypeAlert RemoteNotificationType = 4
	RemoteNotificationTypeBadge RemoteNotificationType = 1
	RemoteNotificationTypeNone  RemoteNotificationType = 0
	RemoteNotificationTypeSound RemoteNotificationType = 2
)

// These constants specify the level of severity of a user attention request and are used by cancelUserAttentionRequest: and requestUserAttention:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrequestuserattentiontype?language=objc
type RequestUserAttentionType uint

const (
	CriticalRequest      RequestUserAttentionType = 0
	InformationalRequest RequestUserAttentionType = 10
)

// Specifies a type for nesting modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditornestingmode?language=objc
type RuleEditorNestingMode uint

const (
	RuleEditorNestingModeCompound RuleEditorNestingMode = 2
	RuleEditorNestingModeList     RuleEditorNestingMode = 1
	RuleEditorNestingModeSimple   RuleEditorNestingMode = 3
	RuleEditorNestingModeSingle   RuleEditorNestingMode = 0
)

// These strings are used as keys to the dictionary returned from the delegate’s ruleEditor:numberOfChildrenForCriterion:withRowType: optional method. To construct a valid predicate, the union of the dictionaries for each item in the row must contain the required parts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditorpredicatepartkey?language=objc
type RuleEditorPredicatePartKey string

const (
	RuleEditorPredicateComparisonModifier RuleEditorPredicatePartKey = "NSRuleEditorPredicateComparisonModifier"
	RuleEditorPredicateCompoundType       RuleEditorPredicatePartKey = "NSRuleEditorPredicateCompoundType"
	RuleEditorPredicateCustomSelector     RuleEditorPredicatePartKey = "NSRuleEditorPredicateCustomSelector"
	RuleEditorPredicateLeftExpression     RuleEditorPredicatePartKey = "NSRuleEditorPredicateLeftExpression"
	RuleEditorPredicateOperatorType       RuleEditorPredicatePartKey = "NSRuleEditorPredicateOperatorType"
	RuleEditorPredicateOptions            RuleEditorPredicatePartKey = "NSRuleEditorPredicateOptions"
	RuleEditorPredicateRightExpression    RuleEditorPredicatePartKey = "NSRuleEditorPredicateRightExpression"
)

// Specifies a type for row types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditorrowtype?language=objc
type RuleEditorRowType uint

const (
	RuleEditorRowTypeCompound RuleEditorRowType = 1
	RuleEditorRowTypeSimple   RuleEditorRowType = 0
)

// These constants are defined to specify a ruler’s orientation and are used by orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerorientation?language=objc
type RulerOrientation uint

const (
	HorizontalRuler RulerOrientation = 0
	VerticalRuler   RulerOrientation = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerviewunitname?language=objc
type RulerViewUnitName string

const (
	RulerViewUnitCentimeters RulerViewUnitName = "Centimeters"
	RulerViewUnitInches      RulerViewUnitName = "Inches"
	RulerViewUnitPicas       RulerViewUnitName = "Picas"
	RulerViewUnitPoints      RulerViewUnitName = "Points"
)

// Constants for specifying the type of document-save operation to perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssaveoperationtype?language=objc
type SaveOperationType uint

const (
	AutosaveAsOperation        SaveOperationType = 5
	AutosaveElsewhereOperation SaveOperationType = 3
	AutosaveInPlaceOperation   SaveOperationType = 4
	AutosaveOperation          SaveOperationType = 3
	SaveAsOperation            SaveOperationType = 1
	SaveOperation              SaveOperationType = 0
	SaveToOperation            SaveOperationType = 2
)

// These constants specify where the scroller’s buttons appear and are used by the arrowsPosition property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollarrowposition?language=objc
type ScrollArrowPosition uint

const (
	ScrollerArrowsDefaultSetting ScrollArrowPosition = 0
	ScrollerArrowsMaxEnd         ScrollArrowPosition = 0
	ScrollerArrowsMinEnd         ScrollArrowPosition = 1
	ScrollerArrowsNone           ScrollArrowPosition = 2
)

// These constants determine the elasticity behavior for an axis of the scrollview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollelasticity?language=objc
type ScrollElasticity int

const (
	ScrollElasticityAllowed   ScrollElasticity = 2
	ScrollElasticityAutomatic ScrollElasticity = 0
	ScrollElasticityNone      ScrollElasticity = 1
)

// These constants define the position of the find bar in relation to the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollviewfindbarposition?language=objc
type ScrollViewFindBarPosition int

const (
	ScrollViewFindBarPositionAboveContent         ScrollViewFindBarPosition = 1
	ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
	ScrollViewFindBarPositionBelowContent         ScrollViewFindBarPosition = 2
)

// These constants describe the two scroller buttons and are used by drawArrow:highlight:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollerarrow?language=objc
type ScrollerArrow uint

const (
	ScrollerDecrementArrow ScrollerArrow = 1
	ScrollerIncrementArrow ScrollerArrow = 0
)

// Specify different knob styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollerknobstyle?language=objc
type ScrollerKnobStyle int

const (
	ScrollerKnobStyleDark    ScrollerKnobStyle = 1
	ScrollerKnobStyleDefault ScrollerKnobStyle = 0
	ScrollerKnobStyleLight   ScrollerKnobStyle = 2
)

// These constants specify the different parts of the scroller: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollerpart?language=objc
type ScrollerPart uint

const (
	ScrollerDecrementLine ScrollerPart = 4
	ScrollerDecrementPage ScrollerPart = 1
	ScrollerIncrementLine ScrollerPart = 5
	ScrollerIncrementPage ScrollerPart = 3
	ScrollerKnob          ScrollerPart = 2
	ScrollerKnobSlot      ScrollerPart = 6
	ScrollerNoPart        ScrollerPart = 0
)

// Constants to specify the scroller style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollerstyle?language=objc
type ScrollerStyle int

const (
	ScrollerStyleLegacy  ScrollerStyle = 0
	ScrollerStyleOverlay ScrollerStyle = 1
)

// The specified preferred alignment of items within the scrubber, when they come to rest following a user’s scrolling or paging interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberalignment?language=objc
type ScrubberAlignment int

const (
	ScrubberAlignmentCenter   ScrubberAlignment = 3
	ScrubberAlignmentLeading  ScrubberAlignment = 1
	ScrubberAlignmentNone     ScrubberAlignment = 0
	ScrubberAlignmentTrailing ScrubberAlignment = 2
)

// The scrolling behavior for a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbermode?language=objc
type ScrubberMode int

const (
	ScrubberModeFixed ScrubberMode = 0
	ScrubberModeFree  ScrubberMode = 1
)

// The string that stores the name under which a search field automatically archives a list of recent search strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldrecentsautosavename?language=objc
type SearchFieldRecentsAutosaveName string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentdistribution?language=objc
type SegmentDistribution int

const (
	SegmentDistributionFill               SegmentDistribution = 1
	SegmentDistributionFillEqually        SegmentDistribution = 2
	SegmentDistributionFillProportionally SegmentDistribution = 3
	SegmentDistributionFit                SegmentDistribution = 0
)

// The following constants specify the visual style used to display the segmented control. They are used by segmentStyle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentstyle?language=objc
type SegmentStyle int

const (
	SegmentStyleAutomatic       SegmentStyle = 0
	SegmentStyleCapsule         SegmentStyle = 5
	SegmentStyleRoundRect       SegmentStyle = 3
	SegmentStyleRounded         SegmentStyle = 1
	SegmentStyleSeparated       SegmentStyle = 8
	SegmentStyleSmallSquare     SegmentStyle = 6
	SegmentStyleTexturedRounded SegmentStyle = 2
	SegmentStyleTexturedSquare  SegmentStyle = 4
)

// The following constants specify the type of tracking behavior a segmented control exhibits. They are used by trackingMode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentswitchtracking?language=objc
type SegmentSwitchTracking uint

const (
	SegmentSwitchTrackingMomentary            SegmentSwitchTracking = 2
	SegmentSwitchTrackingMomentaryAccelerator SegmentSwitchTracking = 3
	SegmentSwitchTrackingSelectAny            SegmentSwitchTracking = 1
	SegmentSwitchTrackingSelectOne            SegmentSwitchTracking = 0
)

// These constants specify the preferred direction of selection. They’re used by selectionAffinity and setSelectedRange:affinity:stillSelecting:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsselectionaffinity?language=objc
type SelectionAffinity uint

const (
	SelectionAffinityDownstream SelectionAffinity = 1
	SelectionAffinityUpstream   SelectionAffinity = 0
)

// Constants that specify the direction a window is currently using to change the key view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsselectiondirection?language=objc
type SelectionDirection uint

const (
	DirectSelection   SelectionDirection = 0
	SelectingNext     SelectionDirection = 1
	SelectingPrevious SelectionDirection = 2
)

// These constants specify how much the text view extends the selection when the user drags the mouse. They’re used by selectionGranularity, and selectionRangeForProposedRange:granularity:: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsselectiongranularity?language=objc
type SelectionGranularity uint

const (
	SelectByCharacter SelectionGranularity = 0
	SelectByParagraph SelectionGranularity = 2
	SelectByWord      SelectionGranularity = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsserviceprovidername?language=objc
type ServiceProviderName string

// The sharing scope constants specify the nature of the things you are sharing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingcontentscope?language=objc
type SharingContentScope int

const (
	SharingContentScopeFull    SharingContentScope = 2
	SharingContentScopeItem    SharingContentScope = 0
	SharingContentScopePartial SharingContentScope = 1
)

// Constants that describe the sharing services that macOS supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicename?language=objc
type SharingServiceName string

const (
	SharingServiceNameAddToAperture             SharingServiceName = "com.apple.share.System.add-to-aperture"
	SharingServiceNameAddToIPhoto               SharingServiceName = "com.apple.share.System.add-to-iphoto"
	SharingServiceNameAddToSafariReadingList    SharingServiceName = "com.apple.share.System.add-to-safari-reading-list"
	SharingServiceNameCloudSharing              SharingServiceName = "com.apple.share.CloudSharing"
	SharingServiceNameComposeEmail              SharingServiceName = "com.apple.share.Mail.compose"
	SharingServiceNameComposeMessage            SharingServiceName = "com.apple.messages.ShareExtension"
	SharingServiceNamePostImageOnFlickr         SharingServiceName = "com.apple.share.Video.upload-image-Flickr"
	SharingServiceNamePostOnFacebook            SharingServiceName = "com.apple.share.Facebook.post"
	SharingServiceNamePostOnLinkedIn            SharingServiceName = "com.apple.share.LinkedIn.post"
	SharingServiceNamePostOnSinaWeibo           SharingServiceName = "com.apple.share.SinaWeibo.post"
	SharingServiceNamePostOnTencentWeibo        SharingServiceName = "com.apple.share.TencentWeibo.post"
	SharingServiceNamePostOnTwitter             SharingServiceName = "com.apple.share.Twitter.post"
	SharingServiceNamePostVideoOnTudou          SharingServiceName = "com.apple.share.Video.upload-Tudou"
	SharingServiceNamePostVideoOnVimeo          SharingServiceName = "com.apple.share.Video.upload-Vimeo"
	SharingServiceNamePostVideoOnYouku          SharingServiceName = "com.apple.share.Video.upload-Youku"
	SharingServiceNameSendViaAirDrop            SharingServiceName = "com.apple.share.AirDrop.send"
	SharingServiceNameUseAsDesktopPicture       SharingServiceName = "com.apple.share.System.set-desktop-image"
	SharingServiceNameUseAsFacebookProfileImage SharingServiceName = "com.apple.share.Facebook.set-profile-image"
	SharingServiceNameUseAsLinkedInProfileImage SharingServiceName = "com.apple.share.LinkedIn.set-profile-image"
	SharingServiceNameUseAsTwitterProfileImage  SharingServiceName = "com.apple.share.Twitter.set-profile-image"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorywidth?language=objc
type SliderAccessoryWidth float64

const (
	SliderAccessoryWidthDefault SliderAccessoryWidth = 36.000000
	SliderAccessoryWidthWide    SliderAccessoryWidth = 72.000000
)

// The types of sliders, used by sliderType. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertype?language=objc
type SliderType uint

const (
	CircularSlider     SliderType = 1
	LinearSlider       SliderType = 0
	SliderTypeCircular SliderType = 1
	SliderTypeLinear   SliderType = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssoundname?language=objc
type SoundName string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssoundplaybackdeviceidentifier?language=objc
type SoundPlaybackDeviceIdentifier string

// These constants are used to indicate where speech should be stopped and paused. See pauseSpeakingAtBoundary: and stopSpeakingAtBoundary:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechboundary?language=objc
type SpeechBoundary uint

const (
	SpeechImmediateBoundary SpeechBoundary = 0
	SpeechSentenceBoundary  SpeechBoundary = 2
	SpeechWordBoundary      SpeechBoundary = 1
)

// Keys for the command delimiters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechcommanddelimiterkey?language=objc
type SpeechCommandDelimiterKey string

const (
	SpeechCommandPrefix SpeechCommandDelimiterKey = "Prefix"
	SpeechCommandSuffix SpeechCommandDelimiterKey = "Suffix"
)

// These constants identify key-value pairs used to add vocabulary to the dictionary using addSpeechDictionary:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechdictionarykey?language=objc
type SpeechDictionaryKey string

const (
	SpeechDictionaryAbbreviations    SpeechDictionaryKey = "Abbreviations"
	SpeechDictionaryEntryPhonemes    SpeechDictionaryKey = "Phonemes"
	SpeechDictionaryEntrySpelling    SpeechDictionaryKey = "Spelling"
	SpeechDictionaryLocaleIdentifier SpeechDictionaryKey = "LocaleIdentifier"
	SpeechDictionaryModificationDate SpeechDictionaryKey = "ModificationDate"
	SpeechDictionaryPronunciations   SpeechDictionaryKey = "Pronunciations"
)

// Keys that identify errors that may occur during speech synthesis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeecherrorkey?language=objc
type SpeechErrorKey string

const (
	SpeechErrorCount                 SpeechErrorKey = "Count"
	SpeechErrorNewestCharacterOffset SpeechErrorKey = "NewestCharacterOffset"
	SpeechErrorNewestCode            SpeechErrorKey = "NewestCode"
	SpeechErrorOldestCharacterOffset SpeechErrorKey = "OldestCharacterOffset"
	SpeechErrorOldestCode            SpeechErrorKey = "OldestCode"
)

// Keys for the speaking mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechmode?language=objc
type SpeechMode string

const (
	SpeechModeLiteral SpeechMode = "LTRL"
	SpeechModeNormal  SpeechMode = "NORM"
	SpeechModePhoneme SpeechMode = "PHON"
	SpeechModeText    SpeechMode = "TEXT"
)

// Keys for the speech phoneme information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechphonemeinfokey?language=objc
type SpeechPhonemeInfoKey string

const (
	SpeechPhonemeInfoExample     SpeechPhonemeInfoKey = "Example"
	SpeechPhonemeInfoHiliteEnd   SpeechPhonemeInfoKey = "HiliteEnd"
	SpeechPhonemeInfoHiliteStart SpeechPhonemeInfoKey = "HiliteStart"
	SpeechPhonemeInfoOpcode      SpeechPhonemeInfoKey = "Opcode"
	SpeechPhonemeInfoSymbol      SpeechPhonemeInfoKey = "Symbol"
)

// These constants are used with setObject:forProperty:error: and objectForProperty:error: to get or set the characteristics of a synthesizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechpropertykey?language=objc
type SpeechPropertyKey string

const (
	SpeechCharacterModeProperty    SpeechPropertyKey = "char"
	SpeechCommandDelimiterProperty SpeechPropertyKey = "dlim"
	SpeechCurrentVoiceProperty     SpeechPropertyKey = "cvox"
	SpeechErrorsProperty           SpeechPropertyKey = "erro"
	SpeechInputModeProperty        SpeechPropertyKey = "inpt"
	SpeechNumberModeProperty       SpeechPropertyKey = "nmbr"
	SpeechOutputToFileURLProperty  SpeechPropertyKey = "opaf"
	SpeechPhonemeSymbolsProperty   SpeechPropertyKey = "phsy"
	SpeechPitchBaseProperty        SpeechPropertyKey = "pbas"
	SpeechPitchModProperty         SpeechPropertyKey = "pmod"
	SpeechRateProperty             SpeechPropertyKey = "rate"
	SpeechRecentSyncProperty       SpeechPropertyKey = "sync"
	SpeechResetProperty            SpeechPropertyKey = "rset"
	SpeechStatusProperty           SpeechPropertyKey = "stat"
	SpeechSynthesizerInfoProperty  SpeechPropertyKey = "vers"
	SpeechVolumeProperty           SpeechPropertyKey = "volm"
)

// Keys for the speech synthesizier status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechstatuskey?language=objc
type SpeechStatusKey string

const (
	SpeechStatusNumberOfCharactersLeft SpeechStatusKey = "NumberOfCharactersLeft"
	SpeechStatusOutputBusy             SpeechStatusKey = "OutputBusy"
	SpeechStatusOutputPaused           SpeechStatusKey = "OutputPaused"
	SpeechStatusPhonemeCode            SpeechStatusKey = "PhonemeCode"
)

// Keys for the speech synthesizier information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizerinfokey?language=objc
type SpeechSynthesizerInfoKey string

const (
	SpeechSynthesizerInfoIdentifier SpeechSynthesizerInfoKey = "Identifier"
	SpeechSynthesizerInfoVersion    SpeechSynthesizerInfoKey = "Version"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizervoicename?language=objc
type SpeechSynthesizerVoiceName string

// Constants for the spelling state attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellingstate?language=objc
type SpellingState int

const (
	SpellingStateGrammarFlag  SpellingState = 2
	SpellingStateSpellingFlag SpellingState = 1
)

// The type that specifies the split view’s autosave name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewautosavename?language=objc
type SplitViewAutosaveName string

// Constants that specify the style of the split view’s dividers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdividerstyle?language=objc
type SplitViewDividerStyle int

const (
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
	SplitViewDividerStyleThick        SplitViewDividerStyle = 1
	SplitViewDividerStyleThin         SplitViewDividerStyle = 2
)

// Constants that describe the behavior of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitembehavior?language=objc
type SplitViewItemBehavior int

const (
	SplitViewItemBehaviorContentList SplitViewItemBehavior = 2
	SplitViewItemBehaviorDefault     SplitViewItemBehavior = 0
	SplitViewItemBehaviorSidebar     SplitViewItemBehavior = 1
)

// Constants that describe the split view item’s collapsing behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemcollapsebehavior?language=objc
type SplitViewItemCollapseBehavior int

const (
	SplitViewItemCollapseBehaviorDefault                                  SplitViewItemCollapseBehavior = 0
	SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView SplitViewItemCollapseBehavior = 2
	SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings SplitViewItemCollapseBehavior = 1
	SplitViewItemCollapseBehaviorUseConstraints                           SplitViewItemCollapseBehavior = 3
)

// A group of constants that indicate a highlighting style for your app’s user interface to display during a spring-loading operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadinghighlight?language=objc
type SpringLoadingHighlight int

const (
	SpringLoadingHighlightEmphasized SpringLoadingHighlight = 2
	SpringLoadingHighlightNone       SpringLoadingHighlight = 0
	SpringLoadingHighlightStandard   SpringLoadingHighlight = 1
)

// These constants denote the type of spring-loading behavior configured for the destination object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspringloadingoptions?language=objc
type SpringLoadingOptions uint

const (
	SpringLoadingContinuousActivation SpringLoadingOptions = 2
	SpringLoadingDisabled             SpringLoadingOptions = 0
	SpringLoadingEnabled              SpringLoadingOptions = 1
	SpringLoadingNoHover              SpringLoadingOptions = 8
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewdistribution?language=objc
type StackViewDistribution int

const (
	StackViewDistributionEqualCentering     StackViewDistribution = 4
	StackViewDistributionEqualSpacing       StackViewDistribution = 3
	StackViewDistributionFill               StackViewDistribution = 0
	StackViewDistributionFillEqually        StackViewDistribution = 1
	StackViewDistributionFillProportionally StackViewDistribution = 2
	StackViewDistributionGravityAreas       StackViewDistribution = -1
)

// The gravity areas available in a stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewgravity?language=objc
type StackViewGravity int

const (
	StackViewGravityBottom   StackViewGravity = 3
	StackViewGravityCenter   StackViewGravity = 2
	StackViewGravityLeading  StackViewGravity = 1
	StackViewGravityTop      StackViewGravity = 1
	StackViewGravityTrailing StackViewGravity = 3
)

// The various Auto Layout priorities for a view in the stack view to remain attached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewvisibilitypriority?language=objc
type StackViewVisibilityPriority float64

const (
	StackViewVisibilityPriorityDetachOnlyIfNecessary StackViewVisibilityPriority = 900.000000
	StackViewVisibilityPriorityMustHold              StackViewVisibilityPriority = 1000.000000
	StackViewVisibilityPriorityNotVisible            StackViewVisibilityPriority = 0.000000
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitemautosavename?language=objc
type StatusItemAutosaveName string

// A set of optional status item behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitembehavior?language=objc
type StatusItemBehavior uint

const (
	StatusItemBehaviorRemovalAllowed       StatusItemBehavior = 2
	StatusItemBehaviorTerminationOnRemoval StatusItemBehavior = 4
)

// The name of the storyboard file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardname?language=objc
type StoryboardName string

// A string that uniquely identifies a view controller or window controller in your storyboard file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsceneidentifier?language=objc
type StoryboardSceneIdentifier string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardsegueidentifier?language=objc
type StoryboardSegueIdentifier string

// Constants that specify the rendering options for drawing a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingoptions?language=objc
type StringDrawingOptions int

const (
	StringDrawingDisableScreenFontSubstitution StringDrawingOptions = 4
	StringDrawingOneShot                       StringDrawingOptions = 16
	StringDrawingTruncatesLastVisibleLine      StringDrawingOptions = 32
	StringDrawingUsesDeviceMetrics             StringDrawingOptions = 8
	StringDrawingUsesFontLeading               StringDrawingOptions = 2
	StringDrawingUsesLineFragmentOrigin        StringDrawingOptions = 1
)

// Constants that represent the supported TIFF data-compression schemes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstiffcompression?language=objc
type TIFFCompression uint

const (
	TIFFCompressionCCITTFAX3 TIFFCompression = 3
	TIFFCompressionCCITTFAX4 TIFFCompression = 4
	TIFFCompressionJPEG      TIFFCompression = 6
	TIFFCompressionLZW       TIFFCompression = 5
	TIFFCompressionNEXT      TIFFCompression = 32766
	TIFFCompressionNone      TIFFCompression = 1
	TIFFCompressionOldJPEG   TIFFCompression = 32865
	TIFFCompressionPackBits  TIFFCompression = 32773
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabposition?language=objc
type TabPosition uint

const (
	TabPositionBottom TabPosition = 3
	TabPositionLeft   TabPosition = 2
	TabPositionNone   TabPosition = 0
	TabPositionRight  TabPosition = 4
	TabPositionTop    TabPosition = 1
)

// These constants describe the current display state of a tab: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabstate?language=objc
type TabState uint

const (
	BackgroundTab TabState = 1
	PressedTab    TabState = 2
	SelectedTab   TabState = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewbordertype?language=objc
type TabViewBorderType uint

const (
	TabViewBorderTypeBezel TabViewBorderType = 2
	TabViewBorderTypeLine  TabViewBorderType = 1
	TabViewBorderTypeNone  TabViewBorderType = 0
)

// Tab control style options for a tab view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontrollertabstyle?language=objc
type TabViewControllerTabStyle int

const (
	TabViewControllerTabStyleSegmentedControlOnBottom TabViewControllerTabStyle = 1
	TabViewControllerTabStyleSegmentedControlOnTop    TabViewControllerTabStyle = 0
	TabViewControllerTabStyleToolbar                  TabViewControllerTabStyle = 2
	TabViewControllerTabStyleUnspecified              TabViewControllerTabStyle = -1
)

// These constants specify the tab view’s type as used by the tabViewType property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewtype?language=objc
type TabViewType uint

const (
	BottomTabsBezelBorder TabViewType = 2
	LeftTabsBezelBorder   TabViewType = 1
	NoTabsBezelBorder     TabViewType = 4
	NoTabsLineBorder      TabViewType = 5
	NoTabsNoBorder        TabViewType = 6
	RightTabsBezelBorder  TabViewType = 3
	TopTabsBezelBorder    TabViewType = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumnresizingoptions?language=objc
type TableColumnResizingOptions uint

const (
	TableColumnAutoresizingMask TableColumnResizingOptions = 1
	TableColumnNoResizing       TableColumnResizingOptions = 0
	TableColumnUserResizingMask TableColumnResizingOptions = 2
)

// These constants define table row edges on which row actions are attached. They are used by the tableView:rowActionsForRow:edge: delegate method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowactionedge?language=objc
type TableRowActionEdge int

const (
	TableRowActionEdgeLeading  TableRowActionEdge = 0
	TableRowActionEdgeTrailing TableRowActionEdge = 1
)

// Specifies the animation effects to apply when inserting or removing rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewanimationoptions?language=objc
type TableViewAnimationOptions uint

const (
	TableViewAnimationEffectFade TableViewAnimationOptions = 1
	TableViewAnimationEffectGap  TableViewAnimationOptions = 2
	TableViewAnimationEffectNone TableViewAnimationOptions = 0
	TableViewAnimationSlideDown  TableViewAnimationOptions = 32
	TableViewAnimationSlideLeft  TableViewAnimationOptions = 48
	TableViewAnimationSlideRight TableViewAnimationOptions = 64
	TableViewAnimationSlideUp    TableViewAnimationOptions = 16
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewautosavename?language=objc
type TableViewAutosaveName string

// The following constants specify the autoresizing styles. These constants are used by the  columnAutoresizingStyle property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewcolumnautoresizingstyle?language=objc
type TableViewColumnAutoresizingStyle uint

const (
	TableViewFirstColumnOnlyAutoresizingStyle         TableViewColumnAutoresizingStyle = 5
	TableViewLastColumnOnlyAutoresizingStyle          TableViewColumnAutoresizingStyle = 4
	TableViewNoColumnAutoresizing                     TableViewColumnAutoresizingStyle = 0
	TableViewReverseSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 3
	TableViewSequentialColumnAutoresizingStyle        TableViewColumnAutoresizingStyle = 2
	TableViewUniformColumnAutoresizingStyle           TableViewColumnAutoresizingStyle = 1
)

// These constants specify the drag styles displayed by the table view. They’re used by draggingDestinationFeedbackStyle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdraggingdestinationfeedbackstyle?language=objc
type TableViewDraggingDestinationFeedbackStyle int

const (
	TableViewDraggingDestinationFeedbackStyleGap        TableViewDraggingDestinationFeedbackStyle = 2
	TableViewDraggingDestinationFeedbackStyleNone       TableViewDraggingDestinationFeedbackStyle = -1
	TableViewDraggingDestinationFeedbackStyleRegular    TableViewDraggingDestinationFeedbackStyle = 0
	TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 1
)

// NSTableView defines these constants to specify drop operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdropoperation?language=objc
type TableViewDropOperation uint

const (
	TableViewDropAbove TableViewDropOperation = 1
	TableViewDropOn    TableViewDropOperation = 0
)

// NSTableView defines these constants to specify grid styles. These constants are used by the gridStyleMask property. The mask can be either [appkit/nstableviewgridlinestyle/nstableviewgridnone] or it can contain either or both of the other options combined using the C bitwise OR operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewgridlinestyle?language=objc
type TableViewGridLineStyle uint

const (
	TableViewDashedHorizontalGridLineMask TableViewGridLineStyle = 8
	TableViewGridNone                     TableViewGridLineStyle = 0
	TableViewSolidHorizontalGridLineMask  TableViewGridLineStyle = 2
	TableViewSolidVerticalGridLineMask    TableViewGridLineStyle = 1
)

// Constants that help define the appearance and behavior of action buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowactionstyle?language=objc
type TableViewRowActionStyle int

const (
	TableViewRowActionStyleDestructive TableViewRowActionStyle = 1
	TableViewRowActionStyleRegular     TableViewRowActionStyle = 0
)

// The row size style constants define the size of the rows in the table view. They are used by the effectiveRowSizeStyle and rowSizeStyle properties. You can also query the row size in the NSTableCellView class’ property rowSizeStyle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowsizestyle?language=objc
type TableViewRowSizeStyle int

const (
	TableViewRowSizeStyleCustom  TableViewRowSizeStyle = 0
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
	TableViewRowSizeStyleLarge   TableViewRowSizeStyle = 3
	TableViewRowSizeStyleMedium  TableViewRowSizeStyle = 2
	TableViewRowSizeStyleSmall   TableViewRowSizeStyle = 1
)

// The following constants specify the selection highlight styles. These constants are used by the selectionHighlightStyle property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewselectionhighlightstyle?language=objc
type TableViewSelectionHighlightStyle int

const (
	TableViewSelectionHighlightStyleNone       TableViewSelectionHighlightStyle = -1
	TableViewSelectionHighlightStyleRegular    TableViewSelectionHighlightStyle = 0
	TableViewSelectionHighlightStyleSourceList TableViewSelectionHighlightStyle = 1
)

// Contains the possible style values for a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewstyle?language=objc
type TableViewStyle int

const (
	TableViewStyleAutomatic  TableViewStyle = 0
	TableViewStyleFullWidth  TableViewStyle = 1
	TableViewStyleInset      TableViewStyle = 2
	TableViewStylePlain      TableViewStyle = 4
	TableViewStyleSourceList TableViewStyle = 3
)

// Constants that specify text alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextalignment?language=objc
type TextAlignment int

const (
	CenterTextAlignment    TextAlignment = 2
	JustifiedTextAlignment TextAlignment = 3
	LeftTextAlignment      TextAlignment = 0
	NaturalTextAlignment   TextAlignment = 4
	RightTextAlignment     TextAlignment = 1
	TextAlignmentJustified TextAlignment = 3
	TextAlignmentLeft      TextAlignment = 0
	TextAlignmentNatural   TextAlignment = 4
)

// The following constants specify values used by the methods setValue:type:forDimension:, valueForDimension:, and valueTypeForDimension: to specify text block dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblockdimension?language=objc
type TextBlockDimension uint

const (
	TextBlockHeight        TextBlockDimension = 4
	TextBlockMaximumHeight TextBlockDimension = 6
	TextBlockMaximumWidth  TextBlockDimension = 2
	TextBlockMinimumHeight TextBlockDimension = 5
	TextBlockMinimumWidth  TextBlockDimension = 1
	TextBlockWidth         TextBlockDimension = 0
)

// The following constants specify values used by the properties and methods contentWidthValueType, setWidth:type:forLayer:, setWidth:type:forLayer:, widthForLayer:edge:, and widthValueTypeForLayer:edge: to specify text block layer values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblocklayer?language=objc
type TextBlockLayer int

const (
	TextBlockBorder  TextBlockLayer = 0
	TextBlockMargin  TextBlockLayer = 1
	TextBlockPadding TextBlockLayer = -1
)

// The following constants specify values used by the methods setValue:type:forDimension: and valueTypeForDimension: to specify text block value types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblockvaluetype?language=objc
type TextBlockValueType uint

const (
	TextBlockAbsoluteValueType   TextBlockValueType = 0
	TextBlockPercentageValueType TextBlockValueType = 1
)

// The following constants specify values used by the property verticalAlignment to specify vertical alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblockverticalalignment?language=objc
type TextBlockVerticalAlignment uint

const (
	TextBlockBaselineAlignment TextBlockVerticalAlignment = 3
	TextBlockBottomAlignment   TextBlockVerticalAlignment = 2
	TextBlockMiddleAlignment   TextBlockVerticalAlignment = 1
	TextBlockTopAlignment      TextBlockVerticalAlignment = 0
)

// The constants are optional keys that can be used in the options dictionary parameter of the checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:, requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:, and menuForResult:string:options:atLocation:inView: methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingoptionkey?language=objc
type TextCheckingOptionKey string

const (
	TextCheckingDocumentAuthorKey     TextCheckingOptionKey = "DocumentAuthor"
	TextCheckingDocumentTitleKey      TextCheckingOptionKey = "DocumentTitle"
	TextCheckingDocumentURLKey        TextCheckingOptionKey = "DocumentURL"
	TextCheckingOrthographyKey        TextCheckingOptionKey = "Orthography"
	TextCheckingQuotesKey             TextCheckingOptionKey = "Quotes"
	TextCheckingReferenceDateKey      TextCheckingOptionKey = "ReferenceDate"
	TextCheckingReferenceTimeZoneKey  TextCheckingOptionKey = "ReferenceTimeZone"
	TextCheckingRegularExpressionsKey TextCheckingOptionKey = "RegularExpressions"
	TextCheckingReplacementsKey       TextCheckingOptionKey = "Replacements"
	TextCheckingSelectedRangeKey      TextCheckingOptionKey = "SelectedRange"
)

// Values that control the order in which the framework enumerates text elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanagerenumerationoptions?language=objc
type TextContentManagerEnumerationOptions uint

const (
	TextContentManagerEnumerationOptionsNone    TextContentManagerEnumerationOptions = 0
	TextContentManagerEnumerationOptionsReverse TextContentManagerEnumerationOptions = 1
)

// Constants that identify the semantic meaning for a text-entry area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontenttype?language=objc
type TextContentType string

const (
	TextContentTypeOneTimeCode TextContentType = "one-time-code"
	TextContentTypePassword    TextContentType = "password"
	TextContentTypeUsername    TextContentType = "username"
)

// Constants for the text effect attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexteffectstyle?language=objc
type TextEffectStyle string

const (
	TextEffectLetterpressStyle TextEffectStyle = "_UIKitNewLetterpressStyle"
)

// The style of bezel the text field displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldbezelstyle?language=objc
type TextFieldBezelStyle uint

const (
	TextFieldRoundedBezel TextFieldBezelStyle = 1
	TextFieldSquareBezel  TextFieldBezelStyle = 0
)

// These constants specify the user interface item tags that correspond find action. These constants are passed to the performTextFinderAction: method, the responder will call the appropriate method for the tag. That method will, in turn, determine the desired action and invoke the appropriate method in the NSTextFinder object’s NSTextFinderClient protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderaction?language=objc
type TextFinderAction int

const (
	TextFinderActionHideFindInterface     TextFinderAction = 11
	TextFinderActionHideReplaceInterface  TextFinderAction = 13
	TextFinderActionNextMatch             TextFinderAction = 2
	TextFinderActionPreviousMatch         TextFinderAction = 3
	TextFinderActionReplace               TextFinderAction = 5
	TextFinderActionReplaceAll            TextFinderAction = 4
	TextFinderActionReplaceAllInSelection TextFinderAction = 8
	TextFinderActionReplaceAndFind        TextFinderAction = 6
	TextFinderActionSelectAll             TextFinderAction = 9
	TextFinderActionSelectAllInSelection  TextFinderAction = 10
	TextFinderActionSetSearchString       TextFinderAction = 7
	TextFinderActionShowFindInterface     TextFinderAction = 1
	TextFinderActionShowReplaceInterface  TextFinderAction = 12
)

// The following constants indicate the type of search anchor an action should perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfindermatchingtype?language=objc
type TextFinderMatchingType int

const (
	TextFinderMatchingTypeContains   TextFinderMatchingType = 0
	TextFinderMatchingTypeEndsWith   TextFinderMatchingType = 3
	TextFinderMatchingTypeFullWord   TextFinderMatchingType = 2
	TextFinderMatchingTypeStartsWith TextFinderMatchingType = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputsourceidentifier?language=objc
type TextInputSourceIdentifier string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraittype?language=objc
type TextInputTraitType int

const (
	TextInputTraitTypeDefault TextInputTraitType = 0
	TextInputTraitTypeNo      TextInputTraitType = 1
	TextInputTraitTypeYes     TextInputTraitType = 2
)

// Values that describe options for enumerating text layout fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragmentenumerationoptions?language=objc
type TextLayoutFragmentEnumerationOptions uint

const (
	TextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment TextLayoutFragmentEnumerationOptions = 8
	TextLayoutFragmentEnumerationOptionsEnsuresLayout            TextLayoutFragmentEnumerationOptions = 4
	TextLayoutFragmentEnumerationOptionsEstimatesSize            TextLayoutFragmentEnumerationOptions = 2
	TextLayoutFragmentEnumerationOptionsNone                     TextLayoutFragmentEnumerationOptions = 0
	TextLayoutFragmentEnumerationOptionsReverse                  TextLayoutFragmentEnumerationOptions = 1
)

// Values that describe the possible layout states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragmentstate?language=objc
type TextLayoutFragmentState uint

const (
	TextLayoutFragmentStateCalculatedUsageBounds TextLayoutFragmentState = 2
	TextLayoutFragmentStateEstimatedUsageBounds  TextLayoutFragmentState = 1
	TextLayoutFragmentStateLayoutAvailable       TextLayoutFragmentState = 3
	TextLayoutFragmentStateNone                  TextLayoutFragmentState = 0
)

// Values that describe where and how the framework extends segments of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagersegmentoptions?language=objc
type TextLayoutManagerSegmentOptions uint

const (
	TextLayoutManagerSegmentOptionsHeadSegmentExtended     TextLayoutManagerSegmentOptions = 4
	TextLayoutManagerSegmentOptionsMiddleFragmentsExcluded TextLayoutManagerSegmentOptions = 2
	TextLayoutManagerSegmentOptionsNone                    TextLayoutManagerSegmentOptions = 0
	TextLayoutManagerSegmentOptionsRangeNotRequired        TextLayoutManagerSegmentOptions = 1
	TextLayoutManagerSegmentOptionsTailSegmentExtended     TextLayoutManagerSegmentOptions = 8
	TextLayoutManagerSegmentOptionsUpstreamAffinity        TextLayoutManagerSegmentOptions = 16
)

// Values that describe the rendering of selection boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagersegmenttype?language=objc
type TextLayoutManagerSegmentType int

const (
	TextLayoutManagerSegmentTypeHighlight TextLayoutManagerSegmentType = 2
	TextLayoutManagerSegmentTypeSelection TextLayoutManagerSegmentType = 1
	TextLayoutManagerSegmentTypeStandard  TextLayoutManagerSegmentType = 0
)

// Constants that describe the text layout orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutorientation?language=objc
type TextLayoutOrientation int

const (
	TextLayoutOrientationHorizontal TextLayoutOrientation = 0
	TextLayoutOrientationVertical   TextLayoutOrientation = 1
)

// Constants for the text layout sections document attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutsectionkey?language=objc
type TextLayoutSectionKey string

const (
	TextLayoutSectionOrientation TextLayoutSectionKey = "NSTextLayoutSectionOrientation"
	TextLayoutSectionRange       TextLayoutSectionKey = "NSTextLayoutSectionRange"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistmarkerformat?language=objc
type TextListMarkerFormat string

const (
	TextListMarkerBox                  TextListMarkerFormat = "{box}"
	TextListMarkerCheck                TextListMarkerFormat = "{check}"
	TextListMarkerCircle               TextListMarkerFormat = "{circle}"
	TextListMarkerDecimal              TextListMarkerFormat = "{decimal}"
	TextListMarkerDiamond              TextListMarkerFormat = "{diamond}"
	TextListMarkerDisc                 TextListMarkerFormat = "{disc}"
	TextListMarkerHyphen               TextListMarkerFormat = "{hyphen}"
	TextListMarkerLowercaseAlpha       TextListMarkerFormat = "{lower-alpha}"
	TextListMarkerLowercaseHexadecimal TextListMarkerFormat = "{lower-hexadecimal}"
	TextListMarkerLowercaseLatin       TextListMarkerFormat = "{lower-latin}"
	TextListMarkerLowercaseRoman       TextListMarkerFormat = "{lower-roman}"
	TextListMarkerOctal                TextListMarkerFormat = "{octal}"
	TextListMarkerSquare               TextListMarkerFormat = "{square}"
	TextListMarkerUppercaseAlpha       TextListMarkerFormat = "{upper-alpha}"
	TextListMarkerUppercaseHexadecimal TextListMarkerFormat = "{upper-hexadecimal}"
	TextListMarkerUppercaseLatin       TextListMarkerFormat = "{upper-latin}"
	TextListMarkerUppercaseRoman       TextListMarkerFormat = "{upper-roman}"
)

// Values that available options for text list items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlistoptions?language=objc
type TextListOptions uint

const (
	TextListPrependEnclosingMarker TextListOptions = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextmovement?language=objc
type TextMovement int

const (
	TextMovementBacktab TextMovement = 18
	TextMovementCancel  TextMovement = 23
	TextMovementDown    TextMovement = 22
	TextMovementLeft    TextMovement = 19
	TextMovementOther   TextMovement = 0
	TextMovementReturn  TextMovement = 16
	TextMovementRight   TextMovement = 20
	TextMovementTab     TextMovement = 17
	TextMovementUp      TextMovement = 21
)

// Constants that specify the text scaling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextscalingtype?language=objc
type TextScalingType int

const (
	TextScalingStandard TextScalingType = 0
	TextScalingiOS      TextScalingType = 1
)

// Values that describe the visual location of the text cursor, or the direction of the non-anchored edge of the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionaffinity?language=objc
type TextSelectionAffinity int

const (
	TextSelectionAffinityDownstream TextSelectionAffinity = 1
	TextSelectionAffinityUpstream   TextSelectionAffinity = 0
)

// Values that describe the different granularities available to make a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiongranularity?language=objc
type TextSelectionGranularity int

const (
	TextSelectionGranularityCharacter TextSelectionGranularity = 0
	TextSelectionGranularityLine      TextSelectionGranularity = 3
	TextSelectionGranularityParagraph TextSelectionGranularity = 2
	TextSelectionGranularitySentence  TextSelectionGranularity = 4
	TextSelectionGranularityWord      TextSelectionGranularity = 1
)

// Values that affect how the framework handles navigation across different textual boundaries during a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigationdestination?language=objc
type TextSelectionNavigationDestination int

const (
	TextSelectionNavigationDestinationCharacter TextSelectionNavigationDestination = 0
	TextSelectionNavigationDestinationContainer TextSelectionNavigationDestination = 5
	TextSelectionNavigationDestinationDocument  TextSelectionNavigationDestination = 6
	TextSelectionNavigationDestinationLine      TextSelectionNavigationDestination = 2
	TextSelectionNavigationDestinationParagraph TextSelectionNavigationDestination = 4
	TextSelectionNavigationDestinationSentence  TextSelectionNavigationDestination = 3
	TextSelectionNavigationDestinationWord      TextSelectionNavigationDestination = 1
)

// Values that describe the direction of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigationdirection?language=objc
type TextSelectionNavigationDirection int

const (
	TextSelectionNavigationDirectionBackward TextSelectionNavigationDirection = 1
	TextSelectionNavigationDirectionDown     TextSelectionNavigationDirection = 5
	TextSelectionNavigationDirectionForward  TextSelectionNavigationDirection = 0
	TextSelectionNavigationDirectionLeft     TextSelectionNavigationDirection = 3
	TextSelectionNavigationDirectionRight    TextSelectionNavigationDirection = 2
	TextSelectionNavigationDirectionUp       TextSelectionNavigationDirection = 4
)

// Values that describe the possible layout orientations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigationlayoutorientation?language=objc
type TextSelectionNavigationLayoutOrientation int

const (
	TextSelectionNavigationLayoutOrientationHorizontal TextSelectionNavigationLayoutOrientation = 0
	TextSelectionNavigationLayoutOrientationVertical   TextSelectionNavigationLayoutOrientation = 1
)

// Values that describe how the framework handles different kinds of selection modifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigationmodifier?language=objc
type TextSelectionNavigationModifier uint

const (
	TextSelectionNavigationModifierExtend   TextSelectionNavigationModifier = 1
	TextSelectionNavigationModifierMultiple TextSelectionNavigationModifier = 4
	TextSelectionNavigationModifierVisual   TextSelectionNavigationModifier = 2
)

// Values that describe the writing direction inside a text selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectionnavigationwritingdirection?language=objc
type TextSelectionNavigationWritingDirection int

const (
	TextSelectionNavigationWritingDirectionLeftToRight TextSelectionNavigationWritingDirection = 0
	TextSelectionNavigationWritingDirectionRightToLeft TextSelectionNavigationWritingDirection = 1
)

// Constants that indicate the types of changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageeditactions?language=objc
type TextStorageEditActions uint

const (
	TextStorageEditedAttributes TextStorageEditActions = 1
	TextStorageEditedCharacters TextStorageEditActions = 2
)

// Constants that indicate the types of edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorageeditedoptions?language=objc
type TextStorageEditedOptions uint

// The terminating character for a tab column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttaboptionkey?language=objc
type TextTabOptionKey string

const (
	TabColumnTerminatorsAttributeName TextTabOptionKey = "NSTabColumnTerminatorsAttributeName"
)

// Constants that specify the type of tab stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttabtype?language=objc
type TextTabType uint

const (
	CenterTabStopType  TextTabType = 2
	DecimalTabStopType TextTabType = 3
	LeftTabStopType    TextTabType = 0
	RightTabStopType   TextTabType = 1
)

// These constants, specifying the type of text table layout algorithm, are used with layoutAlgorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttablelayoutalgorithm?language=objc
type TextTableLayoutAlgorithm uint

const (
	TextTableAutomaticLayoutAlgorithm TextTableLayoutAlgorithm = 0
	TextTableFixedLayoutAlgorithm     TextTableLayoutAlgorithm = 1
)

// The position where a linear slider’s tick marks appear (above, below, leading, or trailing). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstickmarkposition?language=objc
type TickMarkPosition uint

const (
	TickMarkAbove            TickMarkPosition = 1
	TickMarkBelow            TickMarkPosition = 0
	TickMarkLeft             TickMarkPosition = 1
	TickMarkPositionAbove    TickMarkPosition = 1
	TickMarkPositionBelow    TickMarkPosition = 0
	TickMarkPositionLeading  TickMarkPosition = 1
	TickMarkPositionTrailing TickMarkPosition = 0
	TickMarkRight            TickMarkPosition = 0
)

// Specify the location of a box’s title with respect to its border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitleposition?language=objc
type TitlePosition uint

const (
	AboveBottom TitlePosition = 4
	AboveTop    TitlePosition = 1
	AtBottom    TitlePosition = 5
	AtTop       TitlePosition = 2
	BelowBottom TitlePosition = 6
	BelowTop    TitlePosition = 3
	NoTitle     TitlePosition = 0
)

// Styles that determine the type of separator displayed between the title bar and content of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebarseparatorstyle?language=objc
type TitlebarSeparatorStyle int

const (
	TitlebarSeparatorStyleAutomatic TitlebarSeparatorStyle = 0
	TitlebarSeparatorStyleLine      TitlebarSeparatorStyle = 2
	TitlebarSeparatorStyleNone      TitlebarSeparatorStyle = 1
	TitlebarSeparatorStyleShadow    TitlebarSeparatorStyle = 3
)

// The NSTokenStyle constants define how tokens are displayed and editable in the NSTokenFieldCell. These values are used by tokenStyle and the delegate method tokenFieldCell:editingStringForRepresentedObject:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenstyle?language=objc
type TokenStyle uint

const (
	DefaultTokenStyle      TokenStyle = 0
	PlainTextTokenStyle    TokenStyle = 1
	RoundedTokenStyle      TokenStyle = 2
	TokenStyleDefault      TokenStyle = 0
	TokenStyleNone         TokenStyle = 1
	TokenStylePlainSquared TokenStyle = 4
	TokenStyleRounded      TokenStyle = 2
	TokenStyleSquared      TokenStyle = 3
)

// This type describes the rectangle used to identify a tooltip rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstooltiptag?language=objc
type ToolTipTag int

// Constants that indicate whether the toolbar displays items using a name, icon, or combination of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardisplaymode?language=objc
type ToolbarDisplayMode uint

const (
	ToolbarDisplayModeDefault      ToolbarDisplayMode = 0
	ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
	ToolbarDisplayModeIconOnly     ToolbarDisplayMode = 2
	ToolbarDisplayModeLabelOnly    ToolbarDisplayMode = 3
)

// A string value that you use to differentiate your app’s toolbars. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaridentifier?language=objc
type ToolbarIdentifier string

// A value that indicates the display style for a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroupcontrolrepresentation?language=objc
type ToolbarItemGroupControlRepresentation int

const (
	ToolbarItemGroupControlRepresentationAutomatic ToolbarItemGroupControlRepresentation = 0
	ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 2
	ToolbarItemGroupControlRepresentationExpanded  ToolbarItemGroupControlRepresentation = 1
)

// A value that indicates how a grouped toolbar item selects its subitems. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroupselectionmode?language=objc
type ToolbarItemGroupSelectionMode int

const (
	ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 2
	ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 1
	ToolbarItemGroupSelectionModeSelectOne ToolbarItemGroupSelectionMode = 0
)

// Constants for the standard toolbar items that the system provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemidentifier?language=objc
type ToolbarItemIdentifier string

const (
	ToolbarCloudSharingItemIdentifier             ToolbarItemIdentifier = "NSToolbarCloudSharingItem"
	ToolbarCustomizeToolbarItemIdentifier         ToolbarItemIdentifier = "NSToolbarCustomizeToolbarItem"
	ToolbarFlexibleSpaceItemIdentifier            ToolbarItemIdentifier = "NSToolbarFlexibleSpaceItem"
	ToolbarPrintItemIdentifier                    ToolbarItemIdentifier = "NSToolbarPrintItem"
	ToolbarSeparatorItemIdentifier                ToolbarItemIdentifier = "NSToolbarSeparatorItem"
	ToolbarShowColorsItemIdentifier               ToolbarItemIdentifier = "NSToolbarShowColorsItem"
	ToolbarShowFontsItemIdentifier                ToolbarItemIdentifier = "NSToolbarShowFontsItem"
	ToolbarSidebarTrackingSeparatorItemIdentifier ToolbarItemIdentifier = "NSToolbarSidebarTrackingSeparatorItemIdentifier"
	ToolbarSpaceItemIdentifier                    ToolbarItemIdentifier = "NSToolbarSpaceItem"
	ToolbarToggleSidebarItemIdentifier            ToolbarItemIdentifier = "NSToolbarToggleSidebarItem"
)

// Constants that indicate which toolbar items to keep in the toolbar when space is limited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemvisibilitypriority?language=objc
type ToolbarItemVisibilityPriority int

const (
	ToolbarItemVisibilityPriorityHigh     ToolbarItemVisibilityPriority = 1000
	ToolbarItemVisibilityPriorityLow      ToolbarItemVisibilityPriority = -1000
	ToolbarItemVisibilityPriorityStandard ToolbarItemVisibilityPriority = 0
	ToolbarItemVisibilityPriorityUser     ToolbarItemVisibilityPriority = 2000
)

// Constants that specify toolbar display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbarsizemode?language=objc
type ToolbarSizeMode uint

const (
	ToolbarSizeModeDefault ToolbarSizeMode = 0
	ToolbarSizeModeRegular ToolbarSizeMode = 1
	ToolbarSizeModeSmall   ToolbarSizeMode = 2
)

// The default type for a Touch Bar customization identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarcustomizationidentifier?language=objc
type TouchBarCustomizationIdentifier string

// An identifier for an item in the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritemidentifier?language=objc
type TouchBarItemIdentifier string

const (
	TouchBarItemIdentifierCandidateList   TouchBarItemIdentifier = "NSTouchBarItemIdentifierCandidateList"
	TouchBarItemIdentifierCharacterPicker TouchBarItemIdentifier = "NSTouchBarItemIdentifierCharacterPicker"
	TouchBarItemIdentifierFixedSpaceLarge TouchBarItemIdentifier = "NSTouchBarItemIdentifierFixedSpaceLarge"
	TouchBarItemIdentifierFixedSpaceSmall TouchBarItemIdentifier = "NSTouchBarItemIdentifierFixedSpaceSmall"
	TouchBarItemIdentifierFlexibleSpace   TouchBarItemIdentifier = "NSTouchBarItemIdentifierFlexibleSpace"
	TouchBarItemIdentifierOtherItemsProxy TouchBarItemIdentifier = "NSTouchBarItemIdentifierOtherItemsProxy"
	TouchBarItemIdentifierTextAlignment   TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextAlignment"
	TouchBarItemIdentifierTextColorPicker TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextColorPicker"
	TouchBarItemIdentifierTextFormat      TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextFormat"
	TouchBarItemIdentifierTextList        TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextList"
	TouchBarItemIdentifierTextStyle       TouchBarItemIdentifier = "NSTouchBarItemIdentifierTextStyle"
)

// Priorities for the visibility of a Touch Bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritempriority?language=objc
type TouchBarItemPriority float64

const (
	TouchBarItemPriorityHigh   TouchBarItemPriority = 1000.000000
	TouchBarItemPriorityLow    TouchBarItemPriority = -1000.000000
	TouchBarItemPriorityNormal TouchBarItemPriority = 0.000000
)

// The possible phases of a touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchphase?language=objc
type TouchPhase uint

const (
	TouchPhaseAny        TouchPhase = math.MaxUint
	TouchPhaseBegan      TouchPhase = 1
	TouchPhaseCancelled  TouchPhase = 16
	TouchPhaseEnded      TouchPhase = 8
	TouchPhaseMoved      TouchPhase = 2
	TouchPhaseStationary TouchPhase = 4
	TouchPhaseTouching   TouchPhase = 7
)

// A bit mask identifying a direct or indirect touch type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchtype?language=objc
type TouchType int

const (
	TouchTypeDirect   TouchType = 0
	TouchTypeIndirect TouchType = 1
)

// A bit mask identifying a direct or indirect touch type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchtypemask?language=objc
type TouchTypeMask uint

const (
	TouchTypeMaskDirect   TouchTypeMask = 1
	TouchTypeMaskIndirect TouchTypeMask = 2
)

// The data type defined for the constants specified in the options parameter of initWithRect:options:owner:userInfo:. These constants are described below; you can specify multiple constants by performing a bitwise-OR operation with them. In particular, you must supply one or more of the tracking-type constants (that is, [appkit/nstrackingareaoptions/nstrackingmouseenteredandexited], [appkit/nstrackingareaoptions/nstrackingmousemoved], and [appkit/nstrackingareaoptions/nstrackingcursorupdate]) and one of the active constants (that is, [appkit/nstrackingareaoptions/nstrackingactivewhenfirstresponder], [appkit/nstrackingareaoptions/nstrackingactiveinkeywindow], [appkit/nstrackingareaoptions/nstrackingactiveinactiveapp], and [appkit/nstrackingareaoptions/nstrackingactivealways]). In addition, you may specify any of the behavior constants (that is, [appkit/nstrackingareaoptions/nstrackingassumeinside], [appkit/nstrackingareaoptions/nstrackinginvisiblerect], and [appkit/nstrackingareaoptions/nstrackingenabledduringmousedrag]). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingareaoptions?language=objc
type TrackingAreaOptions uint

const (
	TrackingActiveAlways             TrackingAreaOptions = 128
	TrackingActiveInActiveApp        TrackingAreaOptions = 64
	TrackingActiveInKeyWindow        TrackingAreaOptions = 32
	TrackingActiveWhenFirstResponder TrackingAreaOptions = 16
	TrackingAssumeInside             TrackingAreaOptions = 256
	TrackingCursorUpdate             TrackingAreaOptions = 4
	TrackingEnabledDuringMouseDrag   TrackingAreaOptions = 1024
	TrackingInVisibleRect            TrackingAreaOptions = 512
	TrackingMouseEnteredAndExited    TrackingAreaOptions = 1
	TrackingMouseMoved               TrackingAreaOptions = 2
)

// This type describes the rectangle used to track the mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingrecttag?language=objc
type TrackingRectTag int

// Constants that determine the layout manager's behavior during layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetterbehavior?language=objc
type TypesetterBehavior int

const (
	TypesetterBehavior_10_2                   TypesetterBehavior = 2
	TypesetterBehavior_10_2_WithCompatibility TypesetterBehavior = 1
	TypesetterBehavior_10_3                   TypesetterBehavior = 3
	TypesetterBehavior_10_4                   TypesetterBehavior = 4
	TypesetterLatestBehavior                  TypesetterBehavior = -1
	TypesetterOriginalBehavior                TypesetterBehavior = 0
)

// The following constants are possible values returned by the actionForControlCharacterAtIndex: method to determine the action associated with a control character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesettercontrolcharacteraction?language=objc
type TypesetterControlCharacterAction uint

const (
	TypesetterContainerBreakAction  TypesetterControlCharacterAction = 32
	TypesetterHorizontalTabAction   TypesetterControlCharacterAction = 4
	TypesetterLineBreakAction       TypesetterControlCharacterAction = 8
	TypesetterParagraphBreakAction  TypesetterControlCharacterAction = 16
	TypesetterWhitespaceAction      TypesetterControlCharacterAction = 2
	TypesetterZeroAdvancementAction TypesetterControlCharacterAction = 1
)

// Constants for the underline style and strikethrough style attribute keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsunderlinestyle?language=objc
type UnderlineStyle int

const (
	UnderlineStyleByWord            UnderlineStyle = 32768
	UnderlineStyleDouble            UnderlineStyle = 9
	UnderlineStyleNone              UnderlineStyle = 0
	UnderlineStylePatternDash       UnderlineStyle = 512
	UnderlineStylePatternDashDot    UnderlineStyle = 768
	UnderlineStylePatternDashDotDot UnderlineStyle = 1024
	UnderlineStylePatternDot        UnderlineStyle = 256
	UnderlineStylePatternSolid      UnderlineStyle = 0
	UnderlineStyleSingle            UnderlineStyle = 1
	UnderlineStyleThick             UnderlineStyle = 2
)

// These constants specify which parts of the scroller are visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsusablescrollerparts?language=objc
type UsableScrollerParts uint

const (
	AllScrollerParts   UsableScrollerParts = 2
	NoScrollerParts    UsableScrollerParts = 0
	OnlyScrollerArrows UsableScrollerParts = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentifier?language=objc
type UserInterfaceItemIdentifier string

const (
	MenuItemImportFromDeviceIdentifier UserInterfaceItemIdentifier = "NSMenuItemImportFromDeviceIdentifier"
	OutlineViewDisclosureButtonKey     UserInterfaceItemIdentifier = "NSOutlineViewDisclosureButtonKey"
	OutlineViewShowHideButtonKey       UserInterfaceItemIdentifier = "NSOutlineViewShowHideButtonKey"
	TableViewRowViewKey                UserInterfaceItemIdentifier = "NSTableViewRowViewKey"
)

// Specifies the directional flow of the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacelayoutdirection?language=objc
type UserInterfaceLayoutDirection int

const (
	UserInterfaceLayoutDirectionLeftToRight UserInterfaceLayoutDirection = 0
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

// The stack view layout directions, and user interface axes for hugging priority and clipping resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacelayoutorientation?language=objc
type UserInterfaceLayoutOrientation int

const (
	UserInterfaceLayoutOrientationHorizontal UserInterfaceLayoutOrientation = 0
	UserInterfaceLayoutOrientationVertical   UserInterfaceLayoutOrientation = 1
)

// The following constants specify the animation effect to apply and are used as values for the animation effect property of the animation view. See the description of NSViewAnimationEffectKey for usage details. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimationeffectname?language=objc
type ViewAnimationEffectName string

const (
	ViewAnimationFadeInEffect  ViewAnimationEffectName = "NSViewAnimationFadeInEffect"
	ViewAnimationFadeOutEffect ViewAnimationEffectName = "NSViewAnimationFadeOutEffect"
)

// The following string constants are keys for the dictionaries in the array passed into initWithViewAnimations: and viewAnimations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimationkey?language=objc
type ViewAnimationKey string

const (
	ViewAnimationEffectKey     ViewAnimationKey = "NSViewAnimationEffectKey"
	ViewAnimationEndFrameKey   ViewAnimationKey = "NSViewAnimationEndFrameKey"
	ViewAnimationStartFrameKey ViewAnimationKey = "NSViewAnimationStartFrameKey"
	ViewAnimationTargetKey     ViewAnimationKey = "NSViewAnimationTargetKey"
)

// Animation options for view transitions in a view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollertransitionoptions?language=objc
type ViewControllerTransitionOptions uint

const (
	ViewControllerTransitionAllowUserInteraction ViewControllerTransitionOptions = 4096
	ViewControllerTransitionCrossfade            ViewControllerTransitionOptions = 1
	ViewControllerTransitionNone                 ViewControllerTransitionOptions = 0
	ViewControllerTransitionSlideBackward        ViewControllerTransitionOptions = 384
	ViewControllerTransitionSlideDown            ViewControllerTransitionOptions = 32
	ViewControllerTransitionSlideForward         ViewControllerTransitionOptions = 320
	ViewControllerTransitionSlideLeft            ViewControllerTransitionOptions = 64
	ViewControllerTransitionSlideRight           ViewControllerTransitionOptions = 128
	ViewControllerTransitionSlideUp              ViewControllerTransitionOptions = 16
)

// These constants are keys that you can use in the options dictionary in enterFullScreenMode:withOptions: and exitFullScreenModeWithOptions:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewfullscreenmodeoptionkey?language=objc
type ViewFullScreenModeOptionKey string

const (
	FullScreenModeAllScreens                     ViewFullScreenModeOptionKey = "NSFullScreenModeAllScreens"
	FullScreenModeApplicationPresentationOptions ViewFullScreenModeOptionKey = "NSFullScreenModeApplicationPresentationOptions"
	FullScreenModeSetting                        ViewFullScreenModeOptionKey = "NSFullScreenModeSetting"
	FullScreenModeWindowLevel                    ViewFullScreenModeOptionKey = "NSFullScreenModeWindowLevel"
)

// These constants specify the location of the layer content when the content is not rerendered in response to view resizing. For more information, see the layerContentsPlacement property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentsplacement?language=objc
type ViewLayerContentsPlacement int

const (
	ViewLayerContentsPlacementBottom                    ViewLayerContentsPlacement = 8
	ViewLayerContentsPlacementBottomLeft                ViewLayerContentsPlacement = 9
	ViewLayerContentsPlacementBottomRight               ViewLayerContentsPlacement = 7
	ViewLayerContentsPlacementCenter                    ViewLayerContentsPlacement = 3
	ViewLayerContentsPlacementLeft                      ViewLayerContentsPlacement = 10
	ViewLayerContentsPlacementRight                     ViewLayerContentsPlacement = 6
	ViewLayerContentsPlacementScaleAxesIndependently    ViewLayerContentsPlacement = 0
	ViewLayerContentsPlacementScaleProportionallyToFill ViewLayerContentsPlacement = 2
	ViewLayerContentsPlacementScaleProportionallyToFit  ViewLayerContentsPlacement = 1
	ViewLayerContentsPlacementTop                       ViewLayerContentsPlacement = 4
	ViewLayerContentsPlacementTopLeft                   ViewLayerContentsPlacement = 11
	ViewLayerContentsPlacementTopRight                  ViewLayerContentsPlacement = 5
)

// Constants that specify how layer resizing is handled when a view is layer-backed or layer-hosting. For more information, see the  layerContentsRedrawPolicy property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentsredrawpolicy?language=objc
type ViewLayerContentsRedrawPolicy int

const (
	ViewLayerContentsRedrawBeforeViewResize  ViewLayerContentsRedrawPolicy = 3
	ViewLayerContentsRedrawCrossfade         ViewLayerContentsRedrawPolicy = 4
	ViewLayerContentsRedrawDuringViewResize  ViewLayerContentsRedrawPolicy = 2
	ViewLayerContentsRedrawNever             ViewLayerContentsRedrawPolicy = 0
	ViewLayerContentsRedrawOnSetNeedsDisplay ViewLayerContentsRedrawPolicy = 1
)

// Constants that specify whether the visual effect view blends with what's either behind or within the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectblendingmode?language=objc
type VisualEffectBlendingMode int

const (
	VisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = 0
	VisualEffectBlendingModeWithinWindow VisualEffectBlendingMode = 1
)

// Constants to specify the material shown by the visual effect view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectmaterial?language=objc
type VisualEffectMaterial int

const (
	VisualEffectMaterialAppearanceBased       VisualEffectMaterial = 0
	VisualEffectMaterialContentBackground     VisualEffectMaterial = 18
	VisualEffectMaterialDark                  VisualEffectMaterial = 2
	VisualEffectMaterialFullScreenUI          VisualEffectMaterial = 15
	VisualEffectMaterialHUDWindow             VisualEffectMaterial = 13
	VisualEffectMaterialHeaderView            VisualEffectMaterial = 10
	VisualEffectMaterialLight                 VisualEffectMaterial = 1
	VisualEffectMaterialMediumLight           VisualEffectMaterial = 8
	VisualEffectMaterialMenu                  VisualEffectMaterial = 5
	VisualEffectMaterialPopover               VisualEffectMaterial = 6
	VisualEffectMaterialSelection             VisualEffectMaterial = 4
	VisualEffectMaterialSheet                 VisualEffectMaterial = 11
	VisualEffectMaterialSidebar               VisualEffectMaterial = 7
	VisualEffectMaterialTitlebar              VisualEffectMaterial = 3
	VisualEffectMaterialToolTip               VisualEffectMaterial = 17
	VisualEffectMaterialUltraDark             VisualEffectMaterial = 9
	VisualEffectMaterialUnderPageBackground   VisualEffectMaterial = 22
	VisualEffectMaterialUnderWindowBackground VisualEffectMaterial = 21
	VisualEffectMaterialWindowBackground      VisualEffectMaterial = 12
)

// Constants to specify how the material appearance should reflect window activity state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectstate?language=objc
type VisualEffectState int

const (
	VisualEffectStateActive                   VisualEffectState = 1
	VisualEffectStateFollowsWindowActiveState VisualEffectState = 0
	VisualEffectStateInactive                 VisualEffectState = 2
)

// The following constants are keys for the dictionary returned by attributesForVoice:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvoiceattributekey?language=objc
type VoiceAttributeKey string

const (
	VoiceAge                          VoiceAttributeKey = "VoiceAge"
	VoiceDemoText                     VoiceAttributeKey = "VoiceDemoText"
	VoiceGender                       VoiceAttributeKey = "VoiceGender"
	VoiceIdentifier                   VoiceAttributeKey = "VoiceIdentifier"
	VoiceIndividuallySpokenCharacters VoiceAttributeKey = "VoiceIndividuallySpokenCharacters"
	VoiceLanguage                     VoiceAttributeKey = "VoiceLanguage"
	VoiceLocaleIdentifier             VoiceAttributeKey = "VoiceLocaleIdentifier"
	VoiceName                         VoiceAttributeKey = "VoiceName"
	VoiceSupportedCharacters          VoiceAttributeKey = "VoiceSupportedCharacters"
)

// The following constants define voice gender attributes, which are the allowable values of the NSVoiceGender key returned by attributesForVoice:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvoicegendername?language=objc
type VoiceGenderName string

const (
	VoiceGenderFemale VoiceGenderName = "VoiceGenderFemale"
	VoiceGenderMale   VoiceGenderName = "VoiceGenderMale"
	VoiceGenderNeuter VoiceGenderName = "VoiceGenderNeuter"
)

// Constants that specify the winding rule a Bézier path uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindingrule?language=objc
type WindingRule uint

const (
	EvenOddWindingRule WindingRule = 1
	NonZeroWindingRule WindingRule = 0
	WindingRuleEvenOdd WindingRule = 1
	WindingRuleNonZero WindingRule = 0
)

// Constants that control the automatic window animation behavior windows use when ordering to the front or out of view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowanimationbehavior?language=objc
type WindowAnimationBehavior int

const (
	WindowAnimationBehaviorAlertPanel     WindowAnimationBehavior = 5
	WindowAnimationBehaviorDefault        WindowAnimationBehavior = 0
	WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 3
	WindowAnimationBehaviorNone           WindowAnimationBehavior = 2
	WindowAnimationBehaviorUtilityWindow  WindowAnimationBehavior = 4
)

// The following constants and the related data type represent a window’s possible backing locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowbackinglocation?language=objc
type WindowBackingLocation uint

const (
	WindowBackingLocationDefault     WindowBackingLocation = 0
	WindowBackingLocationMainMemory  WindowBackingLocation = 2
	WindowBackingLocationVideoMemory WindowBackingLocation = 1
)

// Constants that provide a way to access standard title bar buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowbutton?language=objc
type WindowButton uint

const (
	WindowCloseButton            WindowButton = 0
	WindowDocumentIconButton     WindowButton = 4
	WindowDocumentVersionsButton WindowButton = 6
	WindowFullScreenButton       WindowButton = 7
	WindowMiniaturizeButton      WindowButton = 1
	WindowToolbarButton          WindowButton = 3
	WindowZoomButton             WindowButton = 2
)

// Window collection behaviors related to Mission Control, Spaces, and Stage Manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowcollectionbehavior?language=objc
type WindowCollectionBehavior uint

const (
	WindowCollectionBehaviorCanJoinAllSpaces          WindowCollectionBehavior = 1
	WindowCollectionBehaviorDefault                   WindowCollectionBehavior = 0
	WindowCollectionBehaviorFullScreenAllowsTiling    WindowCollectionBehavior = 2048
	WindowCollectionBehaviorFullScreenAuxiliary       WindowCollectionBehavior = 256
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 4096
	WindowCollectionBehaviorFullScreenNone            WindowCollectionBehavior = 512
	WindowCollectionBehaviorFullScreenPrimary         WindowCollectionBehavior = 128
	WindowCollectionBehaviorIgnoresCycle              WindowCollectionBehavior = 64
	WindowCollectionBehaviorManaged                   WindowCollectionBehavior = 4
	WindowCollectionBehaviorMoveToActiveSpace         WindowCollectionBehavior = 2
	WindowCollectionBehaviorParticipatesInCycle       WindowCollectionBehavior = 32
	WindowCollectionBehaviorStationary                WindowCollectionBehavior = 16
	WindowCollectionBehaviorTransient                 WindowCollectionBehavior = 8
)

// A type that represents the depth, or amount of memory, for a single pixel in a window or screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdepth?language=objc
type WindowDepth int32

const (
	WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 544
	WindowDepthSixtyfourBitRGB             WindowDepth = 528
	WindowDepthTwentyfourBitRGB            WindowDepth = 520
)

// The type of a window’s frame autosave name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowframeautosavename?language=objc
type WindowFrameAutosaveName string

// The standard window levels in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowlevel?language=objc
type WindowLevel int

const (
	DockWindowLevel        WindowLevel = 20
	FloatingWindowLevel    WindowLevel = 3
	MainMenuWindowLevel    WindowLevel = 24
	ModalPanelWindowLevel  WindowLevel = 8
	NormalWindowLevel      WindowLevel = 0
	PopUpMenuWindowLevel   WindowLevel = 101
	ScreenSaverWindowLevel WindowLevel = 1000
	StatusWindowLevel      WindowLevel = 25
	SubmenuWindowLevel     WindowLevel = 3
	TornOffMenuWindowLevel WindowLevel = 3
)

// This constant indicates a window ordering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowlistoptions?language=objc
type WindowListOptions int

const (
	WindowListOrderedFrontToBack WindowListOptions = 1
)

// Options to use when retrieving window numbers from the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindownumberlistoptions?language=objc
type WindowNumberListOptions uint

const (
	WindowNumberListAllApplications WindowNumberListOptions = 1
	WindowNumberListAllSpaces       WindowNumberListOptions = 16
)

// Specifies whether the window is occluded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowocclusionstate?language=objc
type WindowOcclusionState uint

const (
	WindowOcclusionStateVisible WindowOcclusionState = 2
)

// Constants that let you specify how a window is ordered relative to another window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindoworderingmode?language=objc
type WindowOrderingMode int

const (
	WindowAbove WindowOrderingMode = 1
	WindowBelow WindowOrderingMode = -1
	WindowOut   WindowOrderingMode = 0
)

// The type of a window’s frame descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowpersistableframedescriptor?language=objc
type WindowPersistableFrameDescriptor string

// Constants that represent the access levels other processes can have to a window’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowsharingtype?language=objc
type WindowSharingType uint

const (
	WindowSharingNone      WindowSharingType = 0
	WindowSharingReadOnly  WindowSharingType = 1
	WindowSharingReadWrite WindowSharingType = 2
)

// Constants that specify the style of a window, and that you can combine with the C bitwise OR operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowstylemask?language=objc
type WindowStyleMask uint

const (
	BorderlessWindowMask                  WindowStyleMask = 0
	ClosableWindowMask                    WindowStyleMask = 2
	DocModalWindowMask                    WindowStyleMask = 64
	FullScreenWindowMask                  WindowStyleMask = 16384
	FullSizeContentViewWindowMask         WindowStyleMask = 32768
	HUDWindowMask                         WindowStyleMask = 8192
	MiniaturizableWindowMask              WindowStyleMask = 4
	NonactivatingPanelMask                WindowStyleMask = 128
	ResizableWindowMask                   WindowStyleMask = 8
	TexturedBackgroundWindowMask          WindowStyleMask = 256
	TitledWindowMask                      WindowStyleMask = 1
	UnifiedTitleAndToolbarWindowMask      WindowStyleMask = 4096
	UtilityWindowMask                     WindowStyleMask = 16
	WindowStyleMaskBorderless             WindowStyleMask = 0
	WindowStyleMaskClosable               WindowStyleMask = 2
	WindowStyleMaskDocModalWindow         WindowStyleMask = 64
	WindowStyleMaskFullScreen             WindowStyleMask = 16384
	WindowStyleMaskFullSizeContentView    WindowStyleMask = 32768
	WindowStyleMaskHUDWindow              WindowStyleMask = 8192
	WindowStyleMaskMiniaturizable         WindowStyleMask = 4
	WindowStyleMaskNonactivatingPanel     WindowStyleMask = 128
	WindowStyleMaskResizable              WindowStyleMask = 8
	WindowStyleMaskTexturedBackground     WindowStyleMask = 256
	WindowStyleMaskTitled                 WindowStyleMask = 1
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 4096
	WindowStyleMaskUtilityWindow          WindowStyleMask = 16
)

// A value that allows a group of related windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabbingidentifier?language=objc
type WindowTabbingIdentifier string

// The preferred tabbing behavior of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabbingmode?language=objc
type WindowTabbingMode int

const (
	WindowTabbingModeAutomatic  WindowTabbingMode = 0
	WindowTabbingModeDisallowed WindowTabbingMode = 2
	WindowTabbingModePreferred  WindowTabbingMode = 1
)

// Specifies the appearance of the window’s title bar area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtitlevisibility?language=objc
type WindowTitleVisibility int

const (
	WindowTitleHidden  WindowTitleVisibility = 1
	WindowTitleVisible WindowTitleVisibility = 0
)

// Styles that determine the appearance and location of the toolbar in relation to the title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtoolbarstyle?language=objc
type WindowToolbarStyle int

const (
	WindowToolbarStyleAutomatic      WindowToolbarStyle = 0
	WindowToolbarStyleExpanded       WindowToolbarStyle = 1
	WindowToolbarStylePreference     WindowToolbarStyle = 2
	WindowToolbarStyleUnified        WindowToolbarStyle = 3
	WindowToolbarStyleUnifiedCompact WindowToolbarStyle = 4
)

// A value that indicates the user’s preference for window tabbing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowusertabbingpreference?language=objc
type WindowUserTabbingPreference int

const (
	WindowUserTabbingPreferenceAlways       WindowUserTabbingPreference = 1
	WindowUserTabbingPreferenceInFullScreen WindowUserTabbingPreference = 2
	WindowUserTabbingPreferenceManual       WindowUserTabbingPreference = 0
)

// The types of privileged file operations that can be authorized by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceauthorizationtype?language=objc
type WorkspaceAuthorizationType int

const (
	WorkspaceAuthorizationTypeCreateSymbolicLink WorkspaceAuthorizationType = 0
	WorkspaceAuthorizationTypeReplaceFile        WorkspaceAuthorizationType = 2
	WorkspaceAuthorizationTypeSetAttributes      WorkspaceAuthorizationType = 1
)

// Keys that indicate how to display a new desktop image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspacedesktopimageoptionkey?language=objc
type WorkspaceDesktopImageOptionKey string

const (
	WorkspaceDesktopImageAllowClippingKey WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageAllowClippingKey"
	WorkspaceDesktopImageFillColorKey     WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageFillColorKey"
	WorkspaceDesktopImageScalingKey       WorkspaceDesktopImageOptionKey = "NSWorkspaceDesktopImageScalingKey"
)

// These constants specify different types of file operations used by performFileOperation:source:destination:files:tag:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspacefileoperationname?language=objc
type WorkspaceFileOperationName string

const (
	WorkspaceCompressOperation   WorkspaceFileOperationName = "compress"
	WorkspaceCopyOperation       WorkspaceFileOperationName = "copy"
	WorkspaceDecompressOperation WorkspaceFileOperationName = "decompress"
	WorkspaceDecryptOperation    WorkspaceFileOperationName = "decrypt"
	WorkspaceDestroyOperation    WorkspaceFileOperationName = "destroy"
	WorkspaceDuplicateOperation  WorkspaceFileOperationName = "duplicate"
	WorkspaceEncryptOperation    WorkspaceFileOperationName = "encrypt"
	WorkspaceLinkOperation       WorkspaceFileOperationName = "link"
	WorkspaceMoveOperation       WorkspaceFileOperationName = "move"
	WorkspaceRecycleOperation    WorkspaceFileOperationName = "recycle"
)

// Constants that describe options for creating icons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceiconcreationoptions?language=objc
type WorkspaceIconCreationOptions uint

const (
	Exclude10_4ElementsIconCreationOption      WorkspaceIconCreationOptions = 4
	ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 2
)

// The following keys can be used in the configuration dictionary of the launchApplicationAtURL:options:configuration:error: method.  Each key is optional, and if omitted, default behavior is applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspacelaunchconfigurationkey?language=objc
type WorkspaceLaunchConfigurationKey string

const (
	WorkspaceLaunchConfigurationAppleEvent   WorkspaceLaunchConfigurationKey = "NSWorkspaceLaunchConfigurationAppleEvent"
	WorkspaceLaunchConfigurationArchitecture WorkspaceLaunchConfigurationKey = "NSWorkspaceLaunchConfigurationArchitecture"
	WorkspaceLaunchConfigurationArguments    WorkspaceLaunchConfigurationKey = "NSWorkspaceLaunchConfigurationArguments"
	WorkspaceLaunchConfigurationEnvironment  WorkspaceLaunchConfigurationKey = "NSWorkspaceLaunchConfigurationEnvironment"
)

// Constants specifying how you want to launch an app [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspacelaunchoptions?language=objc
type WorkspaceLaunchOptions uint

const (
	WorkspaceLaunchAllowingClassicStartup   WorkspaceLaunchOptions = 131072
	WorkspaceLaunchAndHide                  WorkspaceLaunchOptions = 1048576
	WorkspaceLaunchAndHideOthers            WorkspaceLaunchOptions = 2097152
	WorkspaceLaunchAndPrint                 WorkspaceLaunchOptions = 2
	WorkspaceLaunchAsync                    WorkspaceLaunchOptions = 65536
	WorkspaceLaunchDefault                  WorkspaceLaunchOptions = 65536
	WorkspaceLaunchInhibitingBackgroundOnly WorkspaceLaunchOptions = 128
	WorkspaceLaunchNewInstance              WorkspaceLaunchOptions = 524288
	WorkspaceLaunchPreferringClassic        WorkspaceLaunchOptions = 262144
	WorkspaceLaunchWithErrorPresentation    WorkspaceLaunchOptions = 64
	WorkspaceLaunchWithoutActivation        WorkspaceLaunchOptions = 512
	WorkspaceLaunchWithoutAddingToRecents   WorkspaceLaunchOptions = 256
)

// Constants that specify the writing direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nswritingdirection?language=objc
type WritingDirection int

const (
	WritingDirectionLeftToRight WritingDirection = 0
	WritingDirectionNatural     WritingDirection = -1
	WritingDirectionRightToLeft WritingDirection = 1
)

// Constants for the writing direction attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nswritingdirectionformattype?language=objc
type WritingDirectionFormatType int

const (
	WritingDirectionEmbedding WritingDirectionFormatType = 0
	WritingDirectionOverride  WritingDirectionFormatType = 2
)
