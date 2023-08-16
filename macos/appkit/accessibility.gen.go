// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The complete list of properties and methods for accessible elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility?language=objc
type PAccessibility interface {
	// optional
	AccessibilityAttributedStringForRange(range_ foundation.Range) foundation.IAttributedString
	HasAccessibilityAttributedStringForRange() bool

	// optional
	AccessibilityFrameForRange(range_ foundation.Range) foundation.Rect
	HasAccessibilityFrameForRange() bool

	// optional
	AccessibilityRTFForRange(range_ foundation.Range) []byte
	HasAccessibilityRTFForRange() bool

	// optional
	AccessibilityRangeForLine(line int) foundation.Range
	HasAccessibilityRangeForLine() bool

	// optional
	AccessibilityLayoutSizeForScreenSize(size foundation.Size) foundation.Size
	HasAccessibilityLayoutSizeForScreenSize() bool

	// optional
	AccessibilityPerformShowDefaultUI() bool
	HasAccessibilityPerformShowDefaultUI() bool

	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool

	// optional
	AccessibilityPerformShowAlternateUI() bool
	HasAccessibilityPerformShowAlternateUI() bool

	// optional
	AccessibilityLineForIndex(index int) int
	HasAccessibilityLineForIndex() bool

	// optional
	AccessibilityPerformPress() bool
	HasAccessibilityPerformPress() bool

	// optional
	AccessibilityScreenPointForLayoutPoint(point foundation.Point) foundation.Point
	HasAccessibilityScreenPointForLayoutPoint() bool

	// optional
	AccessibilityPerformDelete() bool
	HasAccessibilityPerformDelete() bool

	// optional
	AccessibilityRangeForPosition(point foundation.Point) foundation.Range
	HasAccessibilityRangeForPosition() bool

	// optional
	AccessibilityPerformShowMenu() bool
	HasAccessibilityPerformShowMenu() bool

	// optional
	AccessibilityPerformConfirm() bool
	HasAccessibilityPerformConfirm() bool

	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityStyleRangeForIndex(index int) foundation.Range
	HasAccessibilityStyleRangeForIndex() bool

	// optional
	AccessibilityRangeForIndex(index int) foundation.Range
	HasAccessibilityRangeForIndex() bool

	// optional
	AccessibilityPerformRaise() bool
	HasAccessibilityPerformRaise() bool

	// optional
	AccessibilityPerformCancel() bool
	HasAccessibilityPerformCancel() bool

	// optional
	AccessibilityStringForRange(range_ foundation.Range) string
	HasAccessibilityStringForRange() bool

	// optional
	IsAccessibilitySelectorAllowed(selector objc.Selector) bool
	HasIsAccessibilitySelectorAllowed() bool

	// optional
	AccessibilityCellForColumnRow(column int, row int) objc.IObject
	HasAccessibilityCellForColumnRow() bool

	// optional
	AccessibilityPerformPick() bool
	HasAccessibilityPerformPick() bool

	// optional
	AccessibilityScreenSizeForLayoutSize(size foundation.Size) foundation.Size
	HasAccessibilityScreenSizeForLayoutSize() bool

	// optional
	AccessibilityLayoutPointForScreenPoint(point foundation.Point) foundation.Point
	HasAccessibilityLayoutPointForScreenPoint() bool

	// optional
	SetAccessibilityClearButton(value objc.Object)
	HasSetAccessibilityClearButton() bool

	// optional
	AccessibilityClearButton() objc.IObject
	HasAccessibilityClearButton() bool

	// optional
	SetAccessibilityOverflowButton(value objc.Object)
	HasSetAccessibilityOverflowButton() bool

	// optional
	AccessibilityOverflowButton() objc.IObject
	HasAccessibilityOverflowButton() bool

	// optional
	SetAccessibilityChildren(value []objc.Object)
	HasSetAccessibilityChildren() bool

	// optional
	AccessibilityChildren() []objc.IObject
	HasAccessibilityChildren() bool

	// optional
	SetAccessibilityColumnHeaderUIElements(value []objc.Object)
	HasSetAccessibilityColumnHeaderUIElements() bool

	// optional
	AccessibilityColumnHeaderUIElements() []objc.IObject
	HasAccessibilityColumnHeaderUIElements() bool

	// optional
	SetAccessibilitySelectedTextRanges(value []foundation.Value)
	HasSetAccessibilitySelectedTextRanges() bool

	// optional
	AccessibilitySelectedTextRanges() []foundation.IValue
	HasAccessibilitySelectedTextRanges() bool

	// optional
	SetAccessibilityCriticalValue(value objc.Object)
	HasSetAccessibilityCriticalValue() bool

	// optional
	AccessibilityCriticalValue() objc.IObject
	HasAccessibilityCriticalValue() bool

	// optional
	SetAccessibilityHeader(value objc.Object)
	HasSetAccessibilityHeader() bool

	// optional
	AccessibilityHeader() objc.IObject
	HasAccessibilityHeader() bool

	// optional
	SetAccessibilitySharedCharacterRange(value foundation.Range)
	HasSetAccessibilitySharedCharacterRange() bool

	// optional
	AccessibilitySharedCharacterRange() foundation.Range
	HasAccessibilitySharedCharacterRange() bool

	// optional
	SetAccessibilityMarkerGroupUIElement(value objc.Object)
	HasSetAccessibilityMarkerGroupUIElement() bool

	// optional
	AccessibilityMarkerGroupUIElement() objc.IObject
	HasAccessibilityMarkerGroupUIElement() bool

	// optional
	SetAccessibilityRequired(value bool)
	HasSetAccessibilityRequired() bool

	// optional
	IsAccessibilityRequired() bool
	HasIsAccessibilityRequired() bool

	// optional
	SetAccessibilitySearchButton(value objc.Object)
	HasSetAccessibilitySearchButton() bool

	// optional
	AccessibilitySearchButton() objc.IObject
	HasAccessibilitySearchButton() bool

	// optional
	SetAccessibilityParent(value objc.Object)
	HasSetAccessibilityParent() bool

	// optional
	AccessibilityParent() objc.IObject
	HasAccessibilityParent() bool

	// optional
	SetAccessibilityNumberOfCharacters(value int)
	HasSetAccessibilityNumberOfCharacters() bool

	// optional
	AccessibilityNumberOfCharacters() int
	HasAccessibilityNumberOfCharacters() bool

	// optional
	SetAccessibilityDefaultButton(value objc.Object)
	HasSetAccessibilityDefaultButton() bool

	// optional
	AccessibilityDefaultButton() objc.IObject
	HasAccessibilityDefaultButton() bool

	// optional
	SetAccessibilityTitle(value string)
	HasSetAccessibilityTitle() bool

	// optional
	AccessibilityTitle() string
	HasAccessibilityTitle() bool

	// optional
	SetAccessibilityRowHeaderUIElements(value []objc.Object)
	HasSetAccessibilityRowHeaderUIElements() bool

	// optional
	AccessibilityRowHeaderUIElements() []objc.IObject
	HasAccessibilityRowHeaderUIElements() bool

	// optional
	SetAccessibilityVisibleCharacterRange(value foundation.Range)
	HasSetAccessibilityVisibleCharacterRange() bool

	// optional
	AccessibilityVisibleCharacterRange() foundation.Range
	HasAccessibilityVisibleCharacterRange() bool

	// optional
	SetAccessibilityFrame(value foundation.Rect)
	HasSetAccessibilityFrame() bool

	// optional
	AccessibilityFrame() foundation.Rect
	HasAccessibilityFrame() bool

	// optional
	SetAccessibilityHidden(value bool)
	HasSetAccessibilityHidden() bool

	// optional
	IsAccessibilityHidden() bool
	HasIsAccessibilityHidden() bool

	// optional
	SetAccessibilityMarkerTypeDescription(value string)
	HasSetAccessibilityMarkerTypeDescription() bool

	// optional
	AccessibilityMarkerTypeDescription() string
	HasAccessibilityMarkerTypeDescription() bool

	// optional
	SetAccessibilityRowIndexRange(value foundation.Range)
	HasSetAccessibilityRowIndexRange() bool

	// optional
	AccessibilityRowIndexRange() foundation.Range
	HasAccessibilityRowIndexRange() bool

	// optional
	SetAccessibilityVerticalUnits(value AccessibilityUnits)
	HasSetAccessibilityVerticalUnits() bool

	// optional
	AccessibilityVerticalUnits() AccessibilityUnits
	HasAccessibilityVerticalUnits() bool

	// optional
	SetAccessibilityVisibleRows(value []objc.Object)
	HasSetAccessibilityVisibleRows() bool

	// optional
	AccessibilityVisibleRows() []objc.IObject
	HasAccessibilityVisibleRows() bool

	// optional
	SetAccessibilitySelectedCells(value []objc.Object)
	HasSetAccessibilitySelectedCells() bool

	// optional
	AccessibilitySelectedCells() []objc.IObject
	HasAccessibilitySelectedCells() bool

	// optional
	SetAccessibilityVisibleColumns(value []objc.Object)
	HasSetAccessibilityVisibleColumns() bool

	// optional
	AccessibilityVisibleColumns() []objc.IObject
	HasAccessibilityVisibleColumns() bool

	// optional
	SetAccessibilityMain(value bool)
	HasSetAccessibilityMain() bool

	// optional
	IsAccessibilityMain() bool
	HasIsAccessibilityMain() bool

	// optional
	SetAccessibilityMarkerUIElements(value []objc.Object)
	HasSetAccessibilityMarkerUIElements() bool

	// optional
	AccessibilityMarkerUIElements() []objc.IObject
	HasAccessibilityMarkerUIElements() bool

	// optional
	SetAccessibilityWarningValue(value objc.Object)
	HasSetAccessibilityWarningValue() bool

	// optional
	AccessibilityWarningValue() objc.IObject
	HasAccessibilityWarningValue() bool

	// optional
	SetAccessibilityContents(value []objc.Object)
	HasSetAccessibilityContents() bool

	// optional
	AccessibilityContents() []objc.IObject
	HasAccessibilityContents() bool

	// optional
	SetAccessibilityServesAsTitleForUIElements(value []objc.Object)
	HasSetAccessibilityServesAsTitleForUIElements() bool

	// optional
	AccessibilityServesAsTitleForUIElements() []objc.IObject
	HasAccessibilityServesAsTitleForUIElements() bool

	// optional
	SetAccessibilityVerticalUnitDescription(value string)
	HasSetAccessibilityVerticalUnitDescription() bool

	// optional
	AccessibilityVerticalUnitDescription() string
	HasAccessibilityVerticalUnitDescription() bool

	// optional
	SetAccessibilityIndex(value int)
	HasSetAccessibilityIndex() bool

	// optional
	AccessibilityIndex() int
	HasAccessibilityIndex() bool

	// optional
	SetAccessibilityLabel(value string)
	HasSetAccessibilityLabel() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	SetAccessibilityFocusedWindow(value objc.Object)
	HasSetAccessibilityFocusedWindow() bool

	// optional
	AccessibilityFocusedWindow() objc.IObject
	HasAccessibilityFocusedWindow() bool

	// optional
	SetAccessibilityToolbarButton(value objc.Object)
	HasSetAccessibilityToolbarButton() bool

	// optional
	AccessibilityToolbarButton() objc.IObject
	HasAccessibilityToolbarButton() bool

	// optional
	SetAccessibilityShownMenu(value objc.Object)
	HasSetAccessibilityShownMenu() bool

	// optional
	AccessibilityShownMenu() objc.IObject
	HasAccessibilityShownMenu() bool

	// optional
	SetAccessibilitySplitters(value []objc.Object)
	HasSetAccessibilitySplitters() bool

	// optional
	AccessibilitySplitters() []objc.IObject
	HasAccessibilitySplitters() bool

	// optional
	SetAccessibilityMainWindow(value objc.Object)
	HasSetAccessibilityMainWindow() bool

	// optional
	AccessibilityMainWindow() objc.IObject
	HasAccessibilityMainWindow() bool

	// optional
	SetAccessibilityOrientation(value AccessibilityOrientation)
	HasSetAccessibilityOrientation() bool

	// optional
	AccessibilityOrientation() AccessibilityOrientation
	HasAccessibilityOrientation() bool

	// optional
	SetAccessibilityRole(value AccessibilityRole)
	HasSetAccessibilityRole() bool

	// optional
	AccessibilityRole() AccessibilityRole
	HasAccessibilityRole() bool

	// optional
	SetAccessibilityAllowedValues(value []foundation.Number)
	HasSetAccessibilityAllowedValues() bool

	// optional
	AccessibilityAllowedValues() []foundation.INumber
	HasAccessibilityAllowedValues() bool

	// optional
	SetAccessibilityCloseButton(value objc.Object)
	HasSetAccessibilityCloseButton() bool

	// optional
	AccessibilityCloseButton() objc.IObject
	HasAccessibilityCloseButton() bool

	// optional
	SetAccessibilityCancelButton(value objc.Object)
	HasSetAccessibilityCancelButton() bool

	// optional
	AccessibilityCancelButton() objc.IObject
	HasAccessibilityCancelButton() bool

	// optional
	SetAccessibilityFocused(value bool)
	HasSetAccessibilityFocused() bool

	// optional
	IsAccessibilityFocused() bool
	HasIsAccessibilityFocused() bool

	// optional
	SetAccessibilityCustomRotors(value []AccessibilityCustomRotor)
	HasSetAccessibilityCustomRotors() bool

	// optional
	AccessibilityCustomRotors() []IAccessibilityCustomRotor
	HasAccessibilityCustomRotors() bool

	// optional
	SetAccessibilityOrderedByRow(value bool)
	HasSetAccessibilityOrderedByRow() bool

	// optional
	IsAccessibilityOrderedByRow() bool
	HasIsAccessibilityOrderedByRow() bool

	// optional
	SetAccessibilitySubrole(value AccessibilitySubrole)
	HasSetAccessibilitySubrole() bool

	// optional
	AccessibilitySubrole() AccessibilitySubrole
	HasAccessibilitySubrole() bool

	// optional
	SetAccessibilityTitleUIElement(value objc.Object)
	HasSetAccessibilityTitleUIElement() bool

	// optional
	AccessibilityTitleUIElement() objc.IObject
	HasAccessibilityTitleUIElement() bool

	// optional
	SetAccessibilityChildrenInNavigationOrder(value []objc.Object)
	HasSetAccessibilityChildrenInNavigationOrder() bool

	// optional
	AccessibilityChildrenInNavigationOrder() []objc.IObject
	HasAccessibilityChildrenInNavigationOrder() bool

	// optional
	SetAccessibilityAlternateUIVisible(value bool)
	HasSetAccessibilityAlternateUIVisible() bool

	// optional
	IsAccessibilityAlternateUIVisible() bool
	HasIsAccessibilityAlternateUIVisible() bool

	// optional
	SetAccessibilitySelectedTextRange(value foundation.Range)
	HasSetAccessibilitySelectedTextRange() bool

	// optional
	AccessibilitySelectedTextRange() foundation.Range
	HasAccessibilitySelectedTextRange() bool

	// optional
	SetAccessibilitySharedFocusElements(value []objc.Object)
	HasSetAccessibilitySharedFocusElements() bool

	// optional
	AccessibilitySharedFocusElements() []objc.IObject
	HasAccessibilitySharedFocusElements() bool

	// optional
	SetAccessibilityMaxValue(value objc.Object)
	HasSetAccessibilityMaxValue() bool

	// optional
	AccessibilityMaxValue() objc.IObject
	HasAccessibilityMaxValue() bool

	// optional
	SetAccessibilityDocument(value string)
	HasSetAccessibilityDocument() bool

	// optional
	AccessibilityDocument() string
	HasAccessibilityDocument() bool

	// optional
	SetAccessibilityCustomActions(value []AccessibilityCustomAction)
	HasSetAccessibilityCustomActions() bool

	// optional
	AccessibilityCustomActions() []IAccessibilityCustomAction
	HasAccessibilityCustomActions() bool

	// optional
	SetAccessibilityDecrementButton(value objc.Object)
	HasSetAccessibilityDecrementButton() bool

	// optional
	AccessibilityDecrementButton() objc.IObject
	HasAccessibilityDecrementButton() bool

	// optional
	SetAccessibilitySharedTextUIElements(value []objc.Object)
	HasSetAccessibilitySharedTextUIElements() bool

	// optional
	AccessibilitySharedTextUIElements() []objc.IObject
	HasAccessibilitySharedTextUIElements() bool

	// optional
	SetAccessibilitySelected(value bool)
	HasSetAccessibilitySelected() bool

	// optional
	IsAccessibilitySelected() bool
	HasIsAccessibilitySelected() bool

	// optional
	SetAccessibilityColumns(value []objc.Object)
	HasSetAccessibilityColumns() bool

	// optional
	AccessibilityColumns() []objc.IObject
	HasAccessibilityColumns() bool

	// optional
	SetAccessibilityRoleDescription(value string)
	HasSetAccessibilityRoleDescription() bool

	// optional
	AccessibilityRoleDescription() string
	HasAccessibilityRoleDescription() bool

	// optional
	SetAccessibilityActivationPoint(value foundation.Point)
	HasSetAccessibilityActivationPoint() bool

	// optional
	AccessibilityActivationPoint() foundation.Point
	HasAccessibilityActivationPoint() bool

	// optional
	SetAccessibilityIdentifier(value string)
	HasSetAccessibilityIdentifier() bool

	// optional
	AccessibilityIdentifier() string
	HasAccessibilityIdentifier() bool

	// optional
	SetAccessibilityTopLevelUIElement(value objc.Object)
	HasSetAccessibilityTopLevelUIElement() bool

	// optional
	AccessibilityTopLevelUIElement() objc.IObject
	HasAccessibilityTopLevelUIElement() bool

	// optional
	SetAccessibilityMenuBar(value objc.Object)
	HasSetAccessibilityMenuBar() bool

	// optional
	AccessibilityMenuBar() objc.IObject
	HasAccessibilityMenuBar() bool

	// optional
	SetAccessibilityColumnCount(value int)
	HasSetAccessibilityColumnCount() bool

	// optional
	AccessibilityColumnCount() int
	HasAccessibilityColumnCount() bool

	// optional
	SetAccessibilityWindows(value []objc.Object)
	HasSetAccessibilityWindows() bool

	// optional
	AccessibilityWindows() []objc.IObject
	HasAccessibilityWindows() bool

	// optional
	SetAccessibilityURL(value foundation.URL)
	HasSetAccessibilityURL() bool

	// optional
	AccessibilityURL() foundation.IURL
	HasAccessibilityURL() bool

	// optional
	SetAccessibilityMinimizeButton(value objc.Object)
	HasSetAccessibilityMinimizeButton() bool

	// optional
	AccessibilityMinimizeButton() objc.IObject
	HasAccessibilityMinimizeButton() bool

	// optional
	SetAccessibilityDisclosed(value bool)
	HasSetAccessibilityDisclosed() bool

	// optional
	IsAccessibilityDisclosed() bool
	HasIsAccessibilityDisclosed() bool

	// optional
	SetAccessibilityMinValue(value objc.Object)
	HasSetAccessibilityMinValue() bool

	// optional
	AccessibilityMinValue() objc.IObject
	HasAccessibilityMinValue() bool

	// optional
	SetAccessibilityRows(value []objc.Object)
	HasSetAccessibilityRows() bool

	// optional
	AccessibilityRows() []objc.IObject
	HasAccessibilityRows() bool

	// optional
	SetAccessibilityMarkerValues(value objc.Object)
	HasSetAccessibilityMarkerValues() bool

	// optional
	AccessibilityMarkerValues() objc.IObject
	HasAccessibilityMarkerValues() bool

	// optional
	SetAccessibilityModal(value bool)
	HasSetAccessibilityModal() bool

	// optional
	IsAccessibilityModal() bool
	HasIsAccessibilityModal() bool

	// optional
	SetAccessibilityVisibleChildren(value []objc.Object)
	HasSetAccessibilityVisibleChildren() bool

	// optional
	AccessibilityVisibleChildren() []objc.IObject
	HasAccessibilityVisibleChildren() bool

	// optional
	SetAccessibilityVisibleCells(value []objc.Object)
	HasSetAccessibilityVisibleCells() bool

	// optional
	AccessibilityVisibleCells() []objc.IObject
	HasAccessibilityVisibleCells() bool

	// optional
	SetAccessibilityLabelValue(value float64)
	HasSetAccessibilityLabelValue() bool

	// optional
	AccessibilityLabelValue() float64
	HasAccessibilityLabelValue() bool

	// optional
	SetAccessibilityExpanded(value bool)
	HasSetAccessibilityExpanded() bool

	// optional
	IsAccessibilityExpanded() bool
	HasIsAccessibilityExpanded() bool

	// optional
	SetAccessibilityRulerMarkerType(value AccessibilityRulerMarkerType)
	HasSetAccessibilityRulerMarkerType() bool

	// optional
	AccessibilityRulerMarkerType() AccessibilityRulerMarkerType
	HasAccessibilityRulerMarkerType() bool

	// optional
	SetAccessibilityProtectedContent(value bool)
	HasSetAccessibilityProtectedContent() bool

	// optional
	IsAccessibilityProtectedContent() bool
	HasIsAccessibilityProtectedContent() bool

	// optional
	SetAccessibilityValue(value objc.Object)
	HasSetAccessibilityValue() bool

	// optional
	AccessibilityValue() objc.IObject
	HasAccessibilityValue() bool

	// optional
	SetAccessibilityEdited(value bool)
	HasSetAccessibilityEdited() bool

	// optional
	IsAccessibilityEdited() bool
	HasIsAccessibilityEdited() bool

	// optional
	SetAccessibilityHelp(value string)
	HasSetAccessibilityHelp() bool

	// optional
	AccessibilityHelp() string
	HasAccessibilityHelp() bool

	// optional
	SetAccessibilityRowCount(value int)
	HasSetAccessibilityRowCount() bool

	// optional
	AccessibilityRowCount() int
	HasAccessibilityRowCount() bool

	// optional
	SetAccessibilityColumnIndexRange(value foundation.Range)
	HasSetAccessibilityColumnIndexRange() bool

	// optional
	AccessibilityColumnIndexRange() foundation.Range
	HasAccessibilityColumnIndexRange() bool

	// optional
	SetAccessibilitySortDirection(value AccessibilitySortDirection)
	HasSetAccessibilitySortDirection() bool

	// optional
	AccessibilitySortDirection() AccessibilitySortDirection
	HasAccessibilitySortDirection() bool

	// optional
	SetAccessibilityValueDescription(value string)
	HasSetAccessibilityValueDescription() bool

	// optional
	AccessibilityValueDescription() string
	HasAccessibilityValueDescription() bool

	// optional
	SetAccessibilityLinkedUIElements(value []objc.Object)
	HasSetAccessibilityLinkedUIElements() bool

	// optional
	AccessibilityLinkedUIElements() []objc.IObject
	HasAccessibilityLinkedUIElements() bool

	// optional
	SetAccessibilityExtrasMenuBar(value objc.Object)
	HasSetAccessibilityExtrasMenuBar() bool

	// optional
	AccessibilityExtrasMenuBar() objc.IObject
	HasAccessibilityExtrasMenuBar() bool

	// optional
	SetAccessibilityWindow(value objc.Object)
	HasSetAccessibilityWindow() bool

	// optional
	AccessibilityWindow() objc.IObject
	HasAccessibilityWindow() bool

	// optional
	SetAccessibilityDisclosedByRow(value objc.Object)
	HasSetAccessibilityDisclosedByRow() bool

	// optional
	AccessibilityDisclosedByRow() objc.IObject
	HasAccessibilityDisclosedByRow() bool

	// optional
	SetAccessibilityHorizontalUnitDescription(value string)
	HasSetAccessibilityHorizontalUnitDescription() bool

	// optional
	AccessibilityHorizontalUnitDescription() string
	HasAccessibilityHorizontalUnitDescription() bool

	// optional
	SetAccessibilityZoomButton(value objc.Object)
	HasSetAccessibilityZoomButton() bool

	// optional
	AccessibilityZoomButton() objc.IObject
	HasAccessibilityZoomButton() bool

	// optional
	SetAccessibilityColumnTitles(value []objc.Object)
	HasSetAccessibilityColumnTitles() bool

	// optional
	AccessibilityColumnTitles() []objc.IObject
	HasAccessibilityColumnTitles() bool

	// optional
	SetAccessibilityVerticalScrollBar(value objc.Object)
	HasSetAccessibilityVerticalScrollBar() bool

	// optional
	AccessibilityVerticalScrollBar() objc.IObject
	HasAccessibilityVerticalScrollBar() bool

	// optional
	SetAccessibilityIncrementButton(value objc.Object)
	HasSetAccessibilityIncrementButton() bool

	// optional
	AccessibilityIncrementButton() objc.IObject
	HasAccessibilityIncrementButton() bool

	// optional
	SetAccessibilityGrowArea(value objc.Object)
	HasSetAccessibilityGrowArea() bool

	// optional
	AccessibilityGrowArea() objc.IObject
	HasAccessibilityGrowArea() bool

	// optional
	SetAccessibilityNextContents(value []objc.Object)
	HasSetAccessibilityNextContents() bool

	// optional
	AccessibilityNextContents() []objc.IObject
	HasAccessibilityNextContents() bool

	// optional
	SetAccessibilityHorizontalScrollBar(value objc.Object)
	HasSetAccessibilityHorizontalScrollBar() bool

	// optional
	AccessibilityHorizontalScrollBar() objc.IObject
	HasAccessibilityHorizontalScrollBar() bool

	// optional
	SetAccessibilityUnitDescription(value string)
	HasSetAccessibilityUnitDescription() bool

	// optional
	AccessibilityUnitDescription() string
	HasAccessibilityUnitDescription() bool

	// optional
	SetAccessibilitySelectedColumns(value []objc.Object)
	HasSetAccessibilitySelectedColumns() bool

	// optional
	AccessibilitySelectedColumns() []objc.IObject
	HasAccessibilitySelectedColumns() bool

	// optional
	SetAccessibilityPreviousContents(value []objc.Object)
	HasSetAccessibilityPreviousContents() bool

	// optional
	AccessibilityPreviousContents() []objc.IObject
	HasAccessibilityPreviousContents() bool

	// optional
	SetAccessibilitySelectedText(value string)
	HasSetAccessibilitySelectedText() bool

	// optional
	AccessibilitySelectedText() string
	HasAccessibilitySelectedText() bool

	// optional
	SetAccessibilityHorizontalUnits(value AccessibilityUnits)
	HasSetAccessibilityHorizontalUnits() bool

	// optional
	AccessibilityHorizontalUnits() AccessibilityUnits
	HasAccessibilityHorizontalUnits() bool

	// optional
	SetAccessibilityEnabled(value bool)
	HasSetAccessibilityEnabled() bool

	// optional
	IsAccessibilityEnabled() bool
	HasIsAccessibilityEnabled() bool

	// optional
	SetAccessibilityTabs(value []objc.Object)
	HasSetAccessibilityTabs() bool

	// optional
	AccessibilityTabs() []objc.IObject
	HasAccessibilityTabs() bool

	// optional
	SetAccessibilityInsertionPointLineNumber(value int)
	HasSetAccessibilityInsertionPointLineNumber() bool

	// optional
	AccessibilityInsertionPointLineNumber() int
	HasAccessibilityInsertionPointLineNumber() bool

	// optional
	SetAccessibilityProxy(value objc.Object)
	HasSetAccessibilityProxy() bool

	// optional
	AccessibilityProxy() objc.IObject
	HasAccessibilityProxy() bool

	// optional
	SetAccessibilitySearchMenu(value objc.Object)
	HasSetAccessibilitySearchMenu() bool

	// optional
	AccessibilitySearchMenu() objc.IObject
	HasAccessibilitySearchMenu() bool

	// optional
	SetAccessibilityFilename(value string)
	HasSetAccessibilityFilename() bool

	// optional
	AccessibilityFilename() string
	HasAccessibilityFilename() bool

	// optional
	SetAccessibilitySelectedChildren(value []objc.Object)
	HasSetAccessibilitySelectedChildren() bool

	// optional
	AccessibilitySelectedChildren() []objc.IObject
	HasAccessibilitySelectedChildren() bool

	// optional
	SetAccessibilityApplicationFocusedUIElement(value objc.Object)
	HasSetAccessibilityApplicationFocusedUIElement() bool

	// optional
	AccessibilityApplicationFocusedUIElement() objc.IObject
	HasAccessibilityApplicationFocusedUIElement() bool

	// optional
	SetAccessibilityMinimized(value bool)
	HasSetAccessibilityMinimized() bool

	// optional
	IsAccessibilityMinimized() bool
	HasIsAccessibilityMinimized() bool

	// optional
	SetAccessibilityElement(value bool)
	HasSetAccessibilityElement() bool

	// optional
	IsAccessibilityElement() bool
	HasIsAccessibilityElement() bool

	// optional
	SetAccessibilityDisclosedRows(value objc.Object)
	HasSetAccessibilityDisclosedRows() bool

	// optional
	AccessibilityDisclosedRows() objc.IObject
	HasAccessibilityDisclosedRows() bool

	// optional
	SetAccessibilityFullScreenButton(value objc.Object)
	HasSetAccessibilityFullScreenButton() bool

	// optional
	AccessibilityFullScreenButton() objc.IObject
	HasAccessibilityFullScreenButton() bool

	// optional
	SetAccessibilityFrontmost(value bool)
	HasSetAccessibilityFrontmost() bool

	// optional
	IsAccessibilityFrontmost() bool
	HasIsAccessibilityFrontmost() bool

	// optional
	SetAccessibilityLabelUIElements(value []objc.Object)
	HasSetAccessibilityLabelUIElements() bool

	// optional
	AccessibilityLabelUIElements() []objc.IObject
	HasAccessibilityLabelUIElements() bool

	// optional
	SetAccessibilityDisclosureLevel(value int)
	HasSetAccessibilityDisclosureLevel() bool

	// optional
	AccessibilityDisclosureLevel() int
	HasAccessibilityDisclosureLevel() bool

	// optional
	SetAccessibilityPlaceholderValue(value string)
	HasSetAccessibilityPlaceholderValue() bool

	// optional
	AccessibilityPlaceholderValue() string
	HasAccessibilityPlaceholderValue() bool

	// optional
	SetAccessibilityUnits(value AccessibilityUnits)
	HasSetAccessibilityUnits() bool

	// optional
	AccessibilityUnits() AccessibilityUnits
	HasAccessibilityUnits() bool

	// optional
	SetAccessibilityHandles(value []objc.Object)
	HasSetAccessibilityHandles() bool

	// optional
	AccessibilityHandles() []objc.IObject
	HasAccessibilityHandles() bool

	// optional
	SetAccessibilitySelectedRows(value []objc.Object)
	HasSetAccessibilitySelectedRows() bool

	// optional
	AccessibilitySelectedRows() []objc.IObject
	HasAccessibilitySelectedRows() bool
}

// A concrete type wrapper for the [PAccessibility] protocol.
type AccessibilityWrapper struct {
	objc.Object
}

func (a_ AccessibilityWrapper) HasAccessibilityAttributedStringForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityAttributedStringForRange:"))
}

// Returns the attributed substring for the specified range of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1532250-accessibilityattributedstringfor?language=objc
func (a_ AccessibilityWrapper) AccessibilityAttributedStringForRange(range_ foundation.Range) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](a_, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityFrameForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFrameForRange:"))
}

// Returns the rectangle that encloses the specified range of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1526088-accessibilityframeforrange?language=objc
func (a_ AccessibilityWrapper) AccessibilityFrameForRange(range_ foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](a_, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityRTFForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRTFForRange:"))
}

// Returns the rich text format (RTF) data that describes the specified range of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1529273-accessibilityrtfforrange?language=objc
func (a_ AccessibilityWrapper) AccessibilityRTFForRange(range_ foundation.Range) []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("accessibilityRTFForRange:"), range_)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityRangeForLine() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeForLine:"))
}

// Returns the range of characters in the specified line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1528813-accessibilityrangeforline?language=objc
func (a_ AccessibilityWrapper) AccessibilityRangeForLine(line int) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityLayoutSizeForScreenSize() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLayoutSizeForScreenSize:"))
}

// Converts the provided size in screen coordinates to a size in the layout area’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535016-accessibilitylayoutsizeforscreen?language=objc
func (a_ AccessibilityWrapper) AccessibilityLayoutSizeForScreenSize(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](a_, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformShowDefaultUI() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformShowDefaultUI"))
}

// Returns to the accessibility element’s original UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1531207-accessibilityperformshowdefaultu?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1526626-accessibilityperformdecrement?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformShowAlternateUI() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformShowAlternateUI"))
}

// Displays the accessibility element’s alternative UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1533983-accessibilityperformshowalternat?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityLineForIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLineForIndex:"))
}

// Returns the line number for the line that contains the specified character index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1525305-accessibilitylineforindex?language=objc
func (a_ AccessibilityWrapper) AccessibilityLineForIndex(index int) int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformPress() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformPress"))
}

// Simulates clicking the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1526358-accessibilityperformpress?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformPress() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformPress"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityScreenPointForLayoutPoint() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityScreenPointForLayoutPoint:"))
}

// Converts the provided point in the layout area’s coordinates to a point in the screen’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1524668-accessibilityscreenpointforlayou?language=objc
func (a_ AccessibilityWrapper) AccessibilityScreenPointForLayoutPoint(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](a_, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformDelete() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDelete"))
}

// Deletes the accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1524609-accessibilityperformdelete?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformDelete() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDelete"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityRangeForPosition() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeForPosition:"))
}

// Returns the range of characters for the glyph at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1531615-accessibilityrangeforposition?language=objc
func (a_ AccessibilityWrapper) AccessibilityRangeForPosition(point foundation.Point) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformShowMenu() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformShowMenu"))
}

// Displays the menu accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1532774-accessibilityperformshowmenu?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformShowMenu() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformConfirm() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformConfirm"))
}

// Simulates pressing Return in the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534952-accessibilityperformconfirm?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformConfirm() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformConfirm"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1525705-accessibilityperformincrement?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityStyleRangeForIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityStyleRangeForIndex:"))
}

// Returns a range of characters that all have the same style as the specified character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1530474-accessibilitystylerangeforindex?language=objc
func (a_ AccessibilityWrapper) AccessibilityStyleRangeForIndex(index int) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityRangeForIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeForIndex:"))
}

// Returns the range of characters for the glyph that includes the specified character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1525329-accessibilityrangeforindex?language=objc
func (a_ AccessibilityWrapper) AccessibilityRangeForIndex(index int) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformRaise() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformRaise"))
}

// Brings the window to the front. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1530545-accessibilityperformraise?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformRaise() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformRaise"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformCancel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformCancel"))
}

// Cancels the current operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1528679-accessibilityperformcancel?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformCancel() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformCancel"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityStringForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityStringForRange:"))
}

// Returns the substring for the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534940-accessibilitystringforrange?language=objc
func (a_ AccessibilityWrapper) AccessibilityStringForRange(range_ foundation.Range) string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityStringForRange:"), range_)
	return rv
}

func (a_ AccessibilityWrapper) HasIsAccessibilitySelectorAllowed() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilitySelectorAllowed:"))
}

// Returns a Boolean value that indicates whether assistive apps can invoke the specified selector on the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1524956-isaccessibilityselectorallowed?language=objc
func (a_ AccessibilityWrapper) IsAccessibilitySelectorAllowed(selector objc.Selector) bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityCellForColumnRow() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCellForColumn:row:"))
}

// Returns the cell at the specified column and row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1532709-accessibilitycellforcolumn?language=objc
func (a_ AccessibilityWrapper) AccessibilityCellForColumnRow(column int, row int) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityPerformPick() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformPick"))
}

// Selects the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535130-accessibilityperformpick?language=objc
func (a_ AccessibilityWrapper) AccessibilityPerformPick() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformPick"))
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityScreenSizeForLayoutSize() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityScreenSizeForLayoutSize:"))
}

// Converts the provided size in the layout area’s coordinates to a size in the screen’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1526114-accessibilityscreensizeforlayout?language=objc
func (a_ AccessibilityWrapper) AccessibilityScreenSizeForLayoutSize(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](a_, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

func (a_ AccessibilityWrapper) HasAccessibilityLayoutPointForScreenPoint() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLayoutPointForScreenPoint:"))
}

// Converts the provided point in screen coordinates to a point in the layout area’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1526401-accessibilitylayoutpointforscree?language=objc
func (a_ AccessibilityWrapper) AccessibilityLayoutPointForScreenPoint(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](a_, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityClearButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityClearButton:"))
}

// The clear button for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534949-accessibilityclearbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityClearButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityClearButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityClearButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityClearButton"))
}

// The clear button for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534949-accessibilityclearbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityClearButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityClearButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityOverflowButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityOverflowButton:"))
}

// The overflow button for the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534943-accessibilityoverflowbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityOverflowButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityOverflowButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityOverflowButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityOverflowButton"))
}

// The overflow button for the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534943-accessibilityoverflowbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityOverflowButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityOverflowButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityChildren() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityChildren:"))
}

// The child accessibility elements in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535018-accessibilitychildren?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityChildren(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityChildren:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityChildren() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityChildren"))
}

// The child accessibility elements in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535018-accessibilitychildren?language=objc
func (a_ AccessibilityWrapper) AccessibilityChildren() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityChildren"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityColumnHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityColumnHeaderUIElements:"))
}

// The column header accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534988-accessibilitycolumnheaderuieleme?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityColumnHeaderUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityColumnHeaderUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityColumnHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumnHeaderUIElements"))
}

// The column header accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534988-accessibilitycolumnheaderuieleme?language=objc
func (a_ AccessibilityWrapper) AccessibilityColumnHeaderUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityColumnHeaderUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedTextRanges() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedTextRanges:"))
}

// An array of ranges for the currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535133-accessibilityselectedtextranges?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedTextRanges(value []foundation.IValue) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedTextRanges:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedTextRanges() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedTextRanges"))
}

// An array of ranges for the currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535133-accessibilityselectedtextranges?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedTextRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](a_, objc.Sel("accessibilitySelectedTextRanges"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityCriticalValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityCriticalValue:"))
}

// The critical value for the level indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534973-accessibilitycriticalvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityCriticalValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityCriticalValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityCriticalValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCriticalValue"))
}

// The critical value for the level indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534973-accessibilitycriticalvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityCriticalValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityCriticalValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHeader() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHeader:"))
}

// The header for the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534938-accessibilityheader?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHeader(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHeader:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHeader() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHeader"))
}

// The header for the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534938-accessibilityheader?language=objc
func (a_ AccessibilityWrapper) AccessibilityHeader() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityHeader"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySharedCharacterRange() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySharedCharacterRange:"))
}

// The range of characters that the accessibility element displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535069-accessibilitysharedcharacterrang?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySharedCharacterRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySharedCharacterRange:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySharedCharacterRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySharedCharacterRange"))
}

// The range of characters that the accessibility element displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535069-accessibilitysharedcharacterrang?language=objc
func (a_ AccessibilityWrapper) AccessibilitySharedCharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMarkerGroupUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMarkerGroupUIElement:"))
}

// The user interface element that functions as a marker group for the ruler accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535081-accessibilitymarkergroupuielemen?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMarkerGroupUIElement(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMarkerGroupUIElement:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMarkerGroupUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMarkerGroupUIElement"))
}

// The user interface element that functions as a marker group for the ruler accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535081-accessibilitymarkergroupuielemen?language=objc
func (a_ AccessibilityWrapper) AccessibilityMarkerGroupUIElement() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMarkerGroupUIElement"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRequired() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRequired:"))
}

// A Boolean value that determines whether the accessibility element must have content for successful submission of a form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1646618-accessibilityrequired?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRequired(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRequired:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityRequired() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityRequired"))
}

// A Boolean value that determines whether the accessibility element must have content for successful submission of a form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1646618-accessibilityrequired?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityRequired() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityRequired"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySearchButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySearchButton:"))
}

// The search button for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535105-accessibilitysearchbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySearchButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySearchButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySearchButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySearchButton"))
}

// The search button for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535105-accessibilitysearchbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilitySearchButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilitySearchButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityParent() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityParent:"))
}

// The accessibility element’s parent in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535040-accessibilityparent?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityParent(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityParent:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityParent() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityParent"))
}

// The accessibility element’s parent in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535040-accessibilityparent?language=objc
func (a_ AccessibilityWrapper) AccessibilityParent() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityParent"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityNumberOfCharacters() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityNumberOfCharacters:"))
}

// The number of characters in the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534982-accessibilitynumberofcharacters?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityNumberOfCharacters(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityNumberOfCharacters:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityNumberOfCharacters() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityNumberOfCharacters"))
}

// The number of characters in the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534982-accessibilitynumberofcharacters?language=objc
func (a_ AccessibilityWrapper) AccessibilityNumberOfCharacters() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDefaultButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDefaultButton:"))
}

// The child accessibility element that represents the window’s default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534957-accessibilitydefaultbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDefaultButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDefaultButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDefaultButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDefaultButton"))
}

// The child accessibility element that represents the window’s default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534957-accessibilitydefaultbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityDefaultButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityDefaultButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityTitle() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityTitle:"))
}

// The title of the accessibility element—for example, a button’s visible text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535033-accessibilitytitle?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityTitle(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityTitle:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityTitle() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityTitle"))
}

// The title of the accessibility element—for example, a button’s visible text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535033-accessibilitytitle?language=objc
func (a_ AccessibilityWrapper) AccessibilityTitle() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityTitle"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRowHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRowHeaderUIElements:"))
}

// The row header accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535014-accessibilityrowheaderuielements?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRowHeaderUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRowHeaderUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRowHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRowHeaderUIElements"))
}

// The row header accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535014-accessibilityrowheaderuielements?language=objc
func (a_ AccessibilityWrapper) AccessibilityRowHeaderUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityRowHeaderUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVisibleCharacterRange() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVisibleCharacterRange:"))
}

// The range of visible characters in the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535058-accessibilityvisiblecharacterran?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVisibleCharacterRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVisibleCharacterRange:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVisibleCharacterRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleCharacterRange"))
}

// The range of visible characters in the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535058-accessibilityvisiblecharacterran?language=objc
func (a_ AccessibilityWrapper) AccessibilityVisibleCharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFrame() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFrame:"))
}

// The accessibility element’s frame in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534939-accessibilityframe?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFrame(value foundation.Rect) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFrame:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityFrame() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFrame"))
}

// The accessibility element’s frame in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534939-accessibilityframe?language=objc
func (a_ AccessibilityWrapper) AccessibilityFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](a_, objc.Sel("accessibilityFrame"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHidden() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHidden:"))
}

// A Boolean value that determines whether the app is in a hidden state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534961-accessibilityhidden?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHidden(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHidden:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityHidden() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityHidden"))
}

// A Boolean value that determines whether the app is in a hidden state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534961-accessibilityhidden?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityHidden() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityHidden"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMarkerTypeDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMarkerTypeDescription:"))
}

// A human-readable description of the marker type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534968-accessibilitymarkertypedescripti?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMarkerTypeDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMarkerTypeDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMarkerTypeDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMarkerTypeDescription"))
}

// A human-readable description of the marker type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534968-accessibilitymarkertypedescripti?language=objc
func (a_ AccessibilityWrapper) AccessibilityMarkerTypeDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityMarkerTypeDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRowIndexRange() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRowIndexRange:"))
}

// The row index range of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535153-accessibilityrowindexrange?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRowIndexRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRowIndexRange:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRowIndexRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRowIndexRange"))
}

// The row index range of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535153-accessibilityrowindexrange?language=objc
func (a_ AccessibilityWrapper) AccessibilityRowIndexRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRowIndexRange"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVerticalUnits() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVerticalUnits:"))
}

// The units that the layout area uses for vertical values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535011-accessibilityverticalunits?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVerticalUnits(value AccessibilityUnits) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVerticalUnits:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVerticalUnits() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVerticalUnits"))
}

// The units that the layout area uses for vertical values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535011-accessibilityverticalunits?language=objc
func (a_ AccessibilityWrapper) AccessibilityVerticalUnits() AccessibilityUnits {
	rv := objc.Call[AccessibilityUnits](a_, objc.Sel("accessibilityVerticalUnits"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVisibleRows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVisibleRows:"))
}

// The visible rows for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535004-accessibilityvisiblerows?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVisibleRows(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVisibleRows:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVisibleRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleRows"))
}

// The visible rows for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535004-accessibilityvisiblerows?language=objc
func (a_ AccessibilityWrapper) AccessibilityVisibleRows() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleRows"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedCells() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedCells:"))
}

// The currently selected cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535101-accessibilityselectedcells?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedCells(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedCells:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedCells() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedCells"))
}

// The currently selected cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535101-accessibilityselectedcells?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedCells() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedCells"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVisibleColumns() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVisibleColumns:"))
}

// The visible columns for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535150-accessibilityvisiblecolumns?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVisibleColumns(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVisibleColumns:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVisibleColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleColumns"))
}

// The visible columns for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535150-accessibilityvisiblecolumns?language=objc
func (a_ AccessibilityWrapper) AccessibilityVisibleColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleColumns"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMain() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMain:"))
}

// A Boolean value that determines whether the window is the app’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534936-accessibilitymain?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMain(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMain:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityMain() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityMain"))
}

// A Boolean value that determines whether the window is the app’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534936-accessibilitymain?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityMain() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityMain"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMarkerUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMarkerUIElements:"))
}

// An array of marker accessibility elements for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535037-accessibilitymarkeruielements?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMarkerUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMarkerUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMarkerUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMarkerUIElements"))
}

// An array of marker accessibility elements for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535037-accessibilitymarkeruielements?language=objc
func (a_ AccessibilityWrapper) AccessibilityMarkerUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityMarkerUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityWarningValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityWarningValue:"))
}

// The warning value for the level indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535032-accessibilitywarningvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityWarningValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityWarningValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityWarningValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityWarningValue"))
}

// The warning value for the level indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535032-accessibilitywarningvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityWarningValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityWarningValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityContents() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityContents:"))
}

// The contents of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535026-accessibilitycontents?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityContents(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityContents:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityContents() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityContents"))
}

// The contents of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535026-accessibilitycontents?language=objc
func (a_ AccessibilityWrapper) AccessibilityContents() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityContents"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityServesAsTitleForUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityServesAsTitleForUIElements:"))
}

// The list of elements that the accessibility element is a title for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535122-accessibilityservesastitleforuie?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityServesAsTitleForUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityServesAsTitleForUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityServesAsTitleForUIElements"))
}

// The list of elements that the accessibility element is a title for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535122-accessibilityservesastitleforuie?language=objc
func (a_ AccessibilityWrapper) AccessibilityServesAsTitleForUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVerticalUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVerticalUnitDescription:"))
}

// A description of the layout area’s vertical units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535065-accessibilityverticalunitdescrip?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVerticalUnitDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVerticalUnitDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVerticalUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVerticalUnitDescription"))
}

// A description of the layout area’s vertical units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535065-accessibilityverticalunitdescrip?language=objc
func (a_ AccessibilityWrapper) AccessibilityVerticalUnitDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityVerticalUnitDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityIndex:"))
}

// The index of the row or column that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535067-accessibilityindex?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityIndex(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityIndex:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityIndex"))
}

// The index of the row or column that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535067-accessibilityindex?language=objc
func (a_ AccessibilityWrapper) AccessibilityIndex() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityIndex"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityLabel:"))
}

// A short description of the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534976-accessibilitylabel?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityLabel:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// A short description of the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534976-accessibilitylabel?language=objc
func (a_ AccessibilityWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFocusedWindow() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFocusedWindow:"))
}

// The child window with the current focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534986-accessibilityfocusedwindow?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFocusedWindow(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFocusedWindow:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityFocusedWindow() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFocusedWindow"))
}

// The child window with the current focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534986-accessibilityfocusedwindow?language=objc
func (a_ AccessibilityWrapper) AccessibilityFocusedWindow() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityFocusedWindow"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityToolbarButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityToolbarButton:"))
}

// The child accessibility element that represents the window’s toolbar button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535075-accessibilitytoolbarbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityToolbarButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityToolbarButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityToolbarButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityToolbarButton"))
}

// The child accessibility element that represents the window’s toolbar button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535075-accessibilitytoolbarbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityToolbarButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityToolbarButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityShownMenu() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityShownMenu:"))
}

// The menu currently displaying for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534983-accessibilityshownmenu?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityShownMenu(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityShownMenu:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityShownMenu() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityShownMenu"))
}

// The menu currently displaying for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534983-accessibilityshownmenu?language=objc
func (a_ AccessibilityWrapper) AccessibilityShownMenu() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityShownMenu"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySplitters() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySplitters:"))
}

// An array that contains the views and splitter bar from the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535088-accessibilitysplitters?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySplitters(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySplitters:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySplitters() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySplitters"))
}

// An array that contains the views and splitter bar from the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535088-accessibilitysplitters?language=objc
func (a_ AccessibilityWrapper) AccessibilitySplitters() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySplitters"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMainWindow() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMainWindow:"))
}

// The app’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535138-accessibilitymainwindow?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMainWindow(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMainWindow:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMainWindow() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMainWindow"))
}

// The app’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535138-accessibilitymainwindow?language=objc
func (a_ AccessibilityWrapper) AccessibilityMainWindow() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMainWindow"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityOrientation() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityOrientation:"))
}

// The orientation of the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535106-accessibilityorientation?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityOrientation(value AccessibilityOrientation) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityOrientation:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityOrientation() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityOrientation"))
}

// The orientation of the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535106-accessibilityorientation?language=objc
func (a_ AccessibilityWrapper) AccessibilityOrientation() AccessibilityOrientation {
	rv := objc.Call[AccessibilityOrientation](a_, objc.Sel("accessibilityOrientation"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRole() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRole:"))
}

// The type of interface element that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535005-accessibilityrole?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRole(value AccessibilityRole) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRole:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRole() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRole"))
}

// The type of interface element that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535005-accessibilityrole?language=objc
func (a_ AccessibilityWrapper) AccessibilityRole() AccessibilityRole {
	rv := objc.Call[AccessibilityRole](a_, objc.Sel("accessibilityRole"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityAllowedValues() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityAllowedValues:"))
}

// The allowed values for the slider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534941-accessibilityallowedvalues?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityAllowedValues(value []foundation.INumber) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityAllowedValues:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityAllowedValues() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityAllowedValues"))
}

// The allowed values for the slider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534941-accessibilityallowedvalues?language=objc
func (a_ AccessibilityWrapper) AccessibilityAllowedValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("accessibilityAllowedValues"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityCloseButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityCloseButton:"))
}

// The child accessibility element that represents the window’s close button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535022-accessibilityclosebutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityCloseButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityCloseButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityCloseButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCloseButton"))
}

// The child accessibility element that represents the window’s close button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535022-accessibilityclosebutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityCloseButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityCloseButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityCancelButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityCancelButton:"))
}

// The child accessibility element that represents the window’s cancel button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535060-accessibilitycancelbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityCancelButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityCancelButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityCancelButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCancelButton"))
}

// The child accessibility element that represents the window’s cancel button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535060-accessibilitycancelbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityCancelButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityCancelButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFocused() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFocused:"))
}

// A Boolean value that determines whether the accessibility element has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534994-accessibilityfocused?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFocused(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFocused:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityFocused() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityFocused"))
}

// A Boolean value that determines whether the accessibility element has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534994-accessibilityfocused?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityFocused() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityFocused"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityCustomRotors() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityCustomRotors:"))
}

// The custom rotors of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2876053-accessibilitycustomrotors?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityCustomRotors(value []IAccessibilityCustomRotor) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityCustomRotors:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityCustomRotors() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCustomRotors"))
}

// The custom rotors of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2876053-accessibilitycustomrotors?language=objc
func (a_ AccessibilityWrapper) AccessibilityCustomRotors() []AccessibilityCustomRotor {
	rv := objc.Call[[]AccessibilityCustomRotor](a_, objc.Sel("accessibilityCustomRotors"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityOrderedByRow() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityOrderedByRow:"))
}

// A Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535061-accessibilityorderedbyrow?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityOrderedByRow(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityOrderedByRow:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityOrderedByRow() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityOrderedByRow"))
}

// A Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535061-accessibilityorderedbyrow?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityOrderedByRow() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySubrole() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySubrole:"))
}

// The specialized interface element type that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535070-accessibilitysubrole?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySubrole(value AccessibilitySubrole) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySubrole:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySubrole() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySubrole"))
}

// The specialized interface element type that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535070-accessibilitysubrole?language=objc
func (a_ AccessibilityWrapper) AccessibilitySubrole() AccessibilitySubrole {
	rv := objc.Call[AccessibilitySubrole](a_, objc.Sel("accessibilitySubrole"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityTitleUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityTitleUIElement:"))
}

// A static text element that represents the accessibility element’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535155-accessibilitytitleuielement?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityTitleUIElement(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityTitleUIElement:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityTitleUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityTitleUIElement"))
}

// A static text element that represents the accessibility element’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535155-accessibilitytitleuielement?language=objc
func (a_ AccessibilityWrapper) AccessibilityTitleUIElement() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityTitleUIElement"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityChildrenInNavigationOrder() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityChildrenInNavigationOrder:"))
}

// An array of child accessibility elements in order for linear navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2869552-accessibilitychildreninnavigatio?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityChildrenInNavigationOrder(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityChildrenInNavigationOrder() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityChildrenInNavigationOrder"))
}

// An array of child accessibility elements in order for linear navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2869552-accessibilitychildreninnavigatio?language=objc
func (a_ AccessibilityWrapper) AccessibilityChildrenInNavigationOrder() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityAlternateUIVisible() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityAlternateUIVisible:"))
}

// A Boolean value that determines whether the accessibility element’s alternative UI is currently visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535035-accessibilityalternateuivisible?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityAlternateUIVisible(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityAlternateUIVisible:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityAlternateUIVisible() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityAlternateUIVisible"))
}

// A Boolean value that determines whether the accessibility element’s alternative UI is currently visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535035-accessibilityalternateuivisible?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedTextRange() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedTextRange:"))
}

// The range of the currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534989-accessibilityselectedtextrange?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedTextRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedTextRange:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedTextRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedTextRange"))
}

// The range of the currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534989-accessibilityselectedtextrange?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedTextRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilitySelectedTextRange"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySharedFocusElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySharedFocusElements:"))
}

// An array of elements that shares the keyboard focus with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534990-accessibilitysharedfocuselements?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySharedFocusElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySharedFocusElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySharedFocusElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySharedFocusElements"))
}

// An array of elements that shares the keyboard focus with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534990-accessibilitysharedfocuselements?language=objc
func (a_ AccessibilityWrapper) AccessibilitySharedFocusElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySharedFocusElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMaxValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMaxValue:"))
}

// The maximum value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535078-accessibilitymaxvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMaxValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMaxValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMaxValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMaxValue"))
}

// The maximum value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535078-accessibilitymaxvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityMaxValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMaxValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDocument() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDocument:"))
}

// The URL for the file that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534993-accessibilitydocument?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDocument(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDocument:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDocument() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDocument"))
}

// The URL for the file that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534993-accessibilitydocument?language=objc
func (a_ AccessibilityWrapper) AccessibilityDocument() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityDocument"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityCustomActions() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityCustomActions:"))
}

// The custom actions of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2869551-accessibilitycustomactions?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityCustomActions(value []IAccessibilityCustomAction) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityCustomActions:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityCustomActions() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityCustomActions"))
}

// The custom actions of the current accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/2869551-accessibilitycustomactions?language=objc
func (a_ AccessibilityWrapper) AccessibilityCustomActions() []AccessibilityCustomAction {
	rv := objc.Call[[]AccessibilityCustomAction](a_, objc.Sel("accessibilityCustomActions"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDecrementButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDecrementButton:"))
}

// The decrement button for the stepper accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535097-accessibilitydecrementbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDecrementButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDecrementButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDecrementButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDecrementButton"))
}

// The decrement button for the stepper accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535097-accessibilitydecrementbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityDecrementButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityDecrementButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySharedTextUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySharedTextUIElements:"))
}

// Other elements that share text with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534991-accessibilitysharedtextuielement?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySharedTextUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySharedTextUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySharedTextUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySharedTextUIElements"))
}

// Other elements that share text with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534991-accessibilitysharedtextuielement?language=objc
func (a_ AccessibilityWrapper) AccessibilitySharedTextUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySharedTextUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelected() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelected:"))
}

// A Boolean value that determines whether the accessibility element is currently in a selected state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534981-accessibilityselected?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelected(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelected:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilitySelected() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilitySelected"))
}

// A Boolean value that determines whether the accessibility element is currently in a selected state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534981-accessibilityselected?language=objc
func (a_ AccessibilityWrapper) IsAccessibilitySelected() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilitySelected"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityColumns() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityColumns:"))
}

// The column accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535115-accessibilitycolumns?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityColumns(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityColumns:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumns"))
}

// The column accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535115-accessibilitycolumns?language=objc
func (a_ AccessibilityWrapper) AccessibilityColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityColumns"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRoleDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRoleDescription:"))
}

// A localized, human-intelligible description of the accessibility element’s role, such as radio button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535144-accessibilityroledescription?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRoleDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRoleDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRoleDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRoleDescription"))
}

// A localized, human-intelligible description of the accessibility element’s role, such as radio button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535144-accessibilityroledescription?language=objc
func (a_ AccessibilityWrapper) AccessibilityRoleDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityRoleDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityActivationPoint() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityActivationPoint:"))
}

// The activation point for the user interface element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535149-accessibilityactivationpoint?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityActivationPoint(value foundation.Point) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityActivationPoint:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityActivationPoint() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityActivationPoint"))
}

// The activation point for the user interface element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535149-accessibilityactivationpoint?language=objc
func (a_ AccessibilityWrapper) AccessibilityActivationPoint() foundation.Point {
	rv := objc.Call[foundation.Point](a_, objc.Sel("accessibilityActivationPoint"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityIdentifier() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityIdentifier:"))
}

// The accessibility element’s identity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535023-accessibilityidentifier?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityIdentifier(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityIdentifier:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityIdentifier() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityIdentifier"))
}

// The accessibility element’s identity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535023-accessibilityidentifier?language=objc
func (a_ AccessibilityWrapper) AccessibilityIdentifier() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityIdentifier"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityTopLevelUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityTopLevelUIElement:"))
}

// The top-level element that contains the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535092-accessibilitytopleveluielement?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityTopLevelUIElement(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityTopLevelUIElement:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityTopLevelUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityTopLevelUIElement"))
}

// The top-level element that contains the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535092-accessibilitytopleveluielement?language=objc
func (a_ AccessibilityWrapper) AccessibilityTopLevelUIElement() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityTopLevelUIElement"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMenuBar() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMenuBar:"))
}

// The app’s menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535055-accessibilitymenubar?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMenuBar(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMenuBar:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMenuBar() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMenuBar"))
}

// The app’s menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535055-accessibilitymenubar?language=objc
func (a_ AccessibilityWrapper) AccessibilityMenuBar() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMenuBar"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityColumnCount() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityColumnCount:"))
}

// The number of columns in the accessibility element’s grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534966-accessibilitycolumncount?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityColumnCount(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityColumnCount:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityColumnCount() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumnCount"))
}

// The number of columns in the accessibility element’s grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534966-accessibilitycolumncount?language=objc
func (a_ AccessibilityWrapper) AccessibilityColumnCount() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityColumnCount"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityWindows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityWindows:"))
}

// An array that contains all the app’s windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535117-accessibilitywindows?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityWindows(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityWindows:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityWindows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityWindows"))
}

// An array that contains all the app’s windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535117-accessibilitywindows?language=objc
func (a_ AccessibilityWrapper) AccessibilityWindows() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityWindows"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityURL() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityURL:"))
}

// The URL for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535157-accessibilityurl?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityURL(value foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityURL:"), objc.Ptr(value))
}

func (a_ AccessibilityWrapper) HasAccessibilityURL() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityURL"))
}

// The URL for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535157-accessibilityurl?language=objc
func (a_ AccessibilityWrapper) AccessibilityURL() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("accessibilityURL"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMinimizeButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMinimizeButton:"))
}

// The child accessibility element that represents the window’s minimize button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535052-accessibilityminimizebutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMinimizeButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMinimizeButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMinimizeButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMinimizeButton"))
}

// The child accessibility element that represents the window’s minimize button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535052-accessibilityminimizebutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityMinimizeButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMinimizeButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDisclosed() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDisclosed:"))
}

// A Boolean value that determines whether the row is disclosing other rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535124-accessibilitydisclosed?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDisclosed(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDisclosed:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityDisclosed() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityDisclosed"))
}

// A Boolean value that determines whether the row is disclosing other rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535124-accessibilitydisclosed?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityDisclosed() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMinValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMinValue:"))
}

// The minimum value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534995-accessibilityminvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMinValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMinValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMinValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMinValue"))
}

// The minimum value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534995-accessibilityminvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityMinValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMinValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRows:"))
}

// The row accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534945-accessibilityrows?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRows(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRows:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRows"))
}

// The row accessibility elements for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534945-accessibilityrows?language=objc
func (a_ AccessibilityWrapper) AccessibilityRows() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityRows"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMarkerValues() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMarkerValues:"))
}

// The marker values for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535076-accessibilitymarkervalues?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMarkerValues(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMarkerValues:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityMarkerValues() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityMarkerValues"))
}

// The marker values for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535076-accessibilitymarkervalues?language=objc
func (a_ AccessibilityWrapper) AccessibilityMarkerValues() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityMarkerValues"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityModal() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityModal:"))
}

// A Boolean value that determines whether the window is modal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535140-accessibilitymodal?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityModal(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityModal:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityModal() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityModal"))
}

// A Boolean value that determines whether the window is modal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535140-accessibilitymodal?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityModal() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityModal"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVisibleChildren() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVisibleChildren:"))
}

// The accessibility element’s visible child accessibility elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534964-accessibilityvisiblechildren?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVisibleChildren(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVisibleChildren:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVisibleChildren() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleChildren"))
}

// The accessibility element’s visible child accessibility elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534964-accessibilityvisiblechildren?language=objc
func (a_ AccessibilityWrapper) AccessibilityVisibleChildren() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleChildren"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVisibleCells() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVisibleCells:"))
}

// The visible cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535042-accessibilityvisiblecells?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVisibleCells(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVisibleCells:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVisibleCells() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleCells"))
}

// The visible cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535042-accessibilityvisiblecells?language=objc
func (a_ AccessibilityWrapper) AccessibilityVisibleCells() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleCells"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityLabelValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityLabelValue:"))
}

// The value of the label accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535108-accessibilitylabelvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityLabelValue(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityLabelValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityLabelValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabelValue"))
}

// The value of the label accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535108-accessibilitylabelvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityLabelValue() float64 {
	rv := objc.Call[float64](a_, objc.Sel("accessibilityLabelValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityExpanded() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityExpanded:"))
}

// A Boolean value that determines whether the accessibility element is in an expanded state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535045-accessibilityexpanded?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityExpanded(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityExpanded:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityExpanded() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityExpanded"))
}

// A Boolean value that determines whether the accessibility element is in an expanded state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535045-accessibilityexpanded?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityExpanded() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityExpanded"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRulerMarkerType() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRulerMarkerType:"))
}

// The type of markers for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535099-accessibilityrulermarkertype?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRulerMarkerType(value AccessibilityRulerMarkerType) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRulerMarkerType:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRulerMarkerType() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRulerMarkerType"))
}

// The type of markers for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535099-accessibilityrulermarkertype?language=objc
func (a_ AccessibilityWrapper) AccessibilityRulerMarkerType() AccessibilityRulerMarkerType {
	rv := objc.Call[AccessibilityRulerMarkerType](a_, objc.Sel("accessibilityRulerMarkerType"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityProtectedContent() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityProtectedContent:"))
}

// A Boolean value that determines whether the accessibility element contains protected content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535083-accessibilityprotectedcontent?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityProtectedContent(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityProtectedContent:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityProtectedContent() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityProtectedContent"))
}

// A Boolean value that determines whether the accessibility element contains protected content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535083-accessibilityprotectedcontent?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityProtectedContent() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityValue:"))
}

// The accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535103-accessibilityvalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// The accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535103-accessibilityvalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityEdited() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityEdited:"))
}

// A Boolean value that indicates whether the accessibility element is in an edited state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535077-accessibilityedited?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityEdited(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityEdited:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityEdited() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityEdited"))
}

// A Boolean value that indicates whether the accessibility element is in an edited state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535077-accessibilityedited?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityEdited() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityEdited"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHelp() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHelp:"))
}

// The help text for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534974-accessibilityhelp?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHelp(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHelp:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHelp() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHelp"))
}

// The help text for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534974-accessibilityhelp?language=objc
func (a_ AccessibilityWrapper) AccessibilityHelp() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityHelp"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityRowCount() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityRowCount:"))
}

// The number of rows in the accessibility element’s grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535013-accessibilityrowcount?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityRowCount(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityRowCount:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityRowCount() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRowCount"))
}

// The number of rows in the accessibility element’s grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535013-accessibilityrowcount?language=objc
func (a_ AccessibilityWrapper) AccessibilityRowCount() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityRowCount"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityColumnIndexRange() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityColumnIndexRange:"))
}

// The column index range of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534979-accessibilitycolumnindexrange?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityColumnIndexRange(value foundation.Range) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityColumnIndexRange:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityColumnIndexRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumnIndexRange"))
}

// The column index range of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534979-accessibilitycolumnindexrange?language=objc
func (a_ AccessibilityWrapper) AccessibilityColumnIndexRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityColumnIndexRange"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySortDirection() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySortDirection:"))
}

// The accessibility element’s sort direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534962-accessibilitysortdirection?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySortDirection(value AccessibilitySortDirection) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySortDirection:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySortDirection() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySortDirection"))
}

// The accessibility element’s sort direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534962-accessibilitysortdirection?language=objc
func (a_ AccessibilityWrapper) AccessibilitySortDirection() AccessibilitySortDirection {
	rv := objc.Call[AccessibilitySortDirection](a_, objc.Sel("accessibilitySortDirection"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityValueDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityValueDescription:"))
}

// A human-readable description of the accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535113-accessibilityvaluedescription?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityValueDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityValueDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityValueDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValueDescription"))
}

// A human-readable description of the accessibility element’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535113-accessibilityvaluedescription?language=objc
func (a_ AccessibilityWrapper) AccessibilityValueDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityValueDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityLinkedUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityLinkedUIElements:"))
}

// The elements that have links with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534972-accessibilitylinkeduielements?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityLinkedUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityLinkedUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityLinkedUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLinkedUIElements"))
}

// The elements that have links with the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534972-accessibilitylinkeduielements?language=objc
func (a_ AccessibilityWrapper) AccessibilityLinkedUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityLinkedUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityExtrasMenuBar() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityExtrasMenuBar:"))
}

// The icon for the app’s menu bar extra. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534996-accessibilityextrasmenubar?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityExtrasMenuBar(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityExtrasMenuBar:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityExtrasMenuBar() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityExtrasMenuBar"))
}

// The icon for the app’s menu bar extra. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534996-accessibilityextrasmenubar?language=objc
func (a_ AccessibilityWrapper) AccessibilityExtrasMenuBar() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityExtrasMenuBar"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityWindow() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityWindow:"))
}

// The window that contains the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535030-accessibilitywindow?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityWindow(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityWindow:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityWindow() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityWindow"))
}

// The window that contains the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535030-accessibilitywindow?language=objc
func (a_ AccessibilityWrapper) AccessibilityWindow() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityWindow"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDisclosedByRow() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDisclosedByRow:"))
}

// The row disclosing the current row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535146-accessibilitydisclosedbyrow?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDisclosedByRow(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDisclosedByRow:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDisclosedByRow() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDisclosedByRow"))
}

// The row disclosing the current row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535146-accessibilitydisclosedbyrow?language=objc
func (a_ AccessibilityWrapper) AccessibilityDisclosedByRow() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityDisclosedByRow"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHorizontalUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHorizontalUnitDescription:"))
}

// A description of the layout area’s horizontal units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535095-accessibilityhorizontalunitdescr?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHorizontalUnitDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHorizontalUnitDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHorizontalUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHorizontalUnitDescription"))
}

// A description of the layout area’s horizontal units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535095-accessibilityhorizontalunitdescr?language=objc
func (a_ AccessibilityWrapper) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityHorizontalUnitDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityZoomButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityZoomButton:"))
}

// The child accessibility element that represents the window’s zoom button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535090-accessibilityzoombutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityZoomButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityZoomButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityZoomButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityZoomButton"))
}

// The child accessibility element that represents the window’s zoom button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535090-accessibilityzoombutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityZoomButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityZoomButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityColumnTitles() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityColumnTitles:"))
}

// The column titles for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535148-accessibilitycolumntitles?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityColumnTitles(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityColumnTitles:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityColumnTitles() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumnTitles"))
}

// The column titles for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535148-accessibilitycolumntitles?language=objc
func (a_ AccessibilityWrapper) AccessibilityColumnTitles() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityColumnTitles"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityVerticalScrollBar() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityVerticalScrollBar:"))
}

// The vertical scroll bar for the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535053-accessibilityverticalscrollbar?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityVerticalScrollBar(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityVerticalScrollBar:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityVerticalScrollBar() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVerticalScrollBar"))
}

// The vertical scroll bar for the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535053-accessibilityverticalscrollbar?language=objc
func (a_ AccessibilityWrapper) AccessibilityVerticalScrollBar() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityVerticalScrollBar"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityIncrementButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityIncrementButton:"))
}

// The increment button for the stepper accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535007-accessibilityincrementbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityIncrementButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityIncrementButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityIncrementButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityIncrementButton"))
}

// The increment button for the stepper accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535007-accessibilityincrementbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityIncrementButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityIncrementButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityGrowArea() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityGrowArea:"))
}

// The child accessibility element that represents the window’s grow area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535074-accessibilitygrowarea?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityGrowArea(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityGrowArea:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityGrowArea() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityGrowArea"))
}

// The child accessibility element that represents the window’s grow area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535074-accessibilitygrowarea?language=objc
func (a_ AccessibilityWrapper) AccessibilityGrowArea() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityGrowArea"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityNextContents() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityNextContents:"))
}

// The contents that follow the divider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535034-accessibilitynextcontents?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityNextContents(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityNextContents:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityNextContents() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityNextContents"))
}

// The contents that follow the divider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535034-accessibilitynextcontents?language=objc
func (a_ AccessibilityWrapper) AccessibilityNextContents() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityNextContents"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHorizontalScrollBar() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHorizontalScrollBar:"))
}

// The horizontal scroll bar for the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534942-accessibilityhorizontalscrollbar?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHorizontalScrollBar(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHorizontalScrollBar:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHorizontalScrollBar() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHorizontalScrollBar"))
}

// The horizontal scroll bar for the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534942-accessibilityhorizontalscrollbar?language=objc
func (a_ AccessibilityWrapper) AccessibilityHorizontalScrollBar() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityHorizontalScrollBar"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityUnitDescription:"))
}

// A human-readable description of the ruler’s units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535094-accessibilityunitdescription?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityUnitDescription(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityUnitDescription:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityUnitDescription() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityUnitDescription"))
}

// A human-readable description of the ruler’s units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535094-accessibilityunitdescription?language=objc
func (a_ AccessibilityWrapper) AccessibilityUnitDescription() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityUnitDescription"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedColumns() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedColumns:"))
}

// The currently selected columns for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534978-accessibilityselectedcolumns?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedColumns(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedColumns:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedColumns"))
}

// The currently selected columns for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534978-accessibilityselectedcolumns?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedColumns"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityPreviousContents() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityPreviousContents:"))
}

// The contents that precede the divider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534950-accessibilitypreviouscontents?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityPreviousContents(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityPreviousContents:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityPreviousContents() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPreviousContents"))
}

// The contents that precede the divider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534950-accessibilitypreviouscontents?language=objc
func (a_ AccessibilityWrapper) AccessibilityPreviousContents() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityPreviousContents"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedText() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedText:"))
}

// The currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535038-accessibilityselectedtext?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedText(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedText:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedText() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedText"))
}

// The currently selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535038-accessibilityselectedtext?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedText() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilitySelectedText"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHorizontalUnits() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHorizontalUnits:"))
}

// The units that the layout area uses for horizontal values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535154-accessibilityhorizontalunits?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHorizontalUnits(value AccessibilityUnits) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHorizontalUnits:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHorizontalUnits() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHorizontalUnits"))
}

// The units that the layout area uses for horizontal values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535154-accessibilityhorizontalunits?language=objc
func (a_ AccessibilityWrapper) AccessibilityHorizontalUnits() AccessibilityUnits {
	rv := objc.Call[AccessibilityUnits](a_, objc.Sel("accessibilityHorizontalUnits"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityEnabled() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityEnabled:"))
}

// A Boolean value that determines whether the accessibility element responds to user events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535024-accessibilityenabled?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityEnabled(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityEnabled:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityEnabled() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityEnabled"))
}

// A Boolean value that determines whether the accessibility element responds to user events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535024-accessibilityenabled?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityEnabled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityEnabled"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityTabs() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityTabs:"))
}

// The tab accessibility elements for the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535044-accessibilitytabs?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityTabs(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityTabs:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityTabs() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityTabs"))
}

// The tab accessibility elements for the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535044-accessibilitytabs?language=objc
func (a_ AccessibilityWrapper) AccessibilityTabs() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityTabs"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityInsertionPointLineNumber() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityInsertionPointLineNumber:"))
}

// The line number that contains the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535050-accessibilityinsertionpointlinen?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityInsertionPointLineNumber(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityInsertionPointLineNumber:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityInsertionPointLineNumber() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityInsertionPointLineNumber"))
}

// The line number that contains the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535050-accessibilityinsertionpointlinen?language=objc
func (a_ AccessibilityWrapper) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityProxy() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityProxy:"))
}

// The child accessibility element that represents the window’s proxy icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535143-accessibilityproxy?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityProxy(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityProxy:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityProxy() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityProxy"))
}

// The child accessibility element that represents the window’s proxy icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535143-accessibilityproxy?language=objc
func (a_ AccessibilityWrapper) AccessibilityProxy() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityProxy"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySearchMenu() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySearchMenu:"))
}

// The search menu for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535015-accessibilitysearchmenu?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySearchMenu(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySearchMenu:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySearchMenu() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySearchMenu"))
}

// The search menu for the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535015-accessibilitysearchmenu?language=objc
func (a_ AccessibilityWrapper) AccessibilitySearchMenu() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilitySearchMenu"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFilename() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFilename:"))
}

// The filename for the file that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535068-accessibilityfilename?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFilename(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFilename:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityFilename() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFilename"))
}

// The filename for the file that the accessibility element represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535068-accessibilityfilename?language=objc
func (a_ AccessibilityWrapper) AccessibilityFilename() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityFilename"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedChildren() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedChildren:"))
}

// The accessibility element’s currently selected children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534970-accessibilityselectedchildren?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedChildren(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedChildren:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedChildren() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedChildren"))
}

// The accessibility element’s currently selected children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534970-accessibilityselectedchildren?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedChildren() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedChildren"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityApplicationFocusedUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityApplicationFocusedUIElement:"))
}

// The child accessibility element with the current focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535139-accessibilityapplicationfocusedu?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityApplicationFocusedUIElement(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityApplicationFocusedUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityApplicationFocusedUIElement"))
}

// The child accessibility element with the current focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535139-accessibilityapplicationfocusedu?language=objc
func (a_ AccessibilityWrapper) AccessibilityApplicationFocusedUIElement() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityMinimized() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityMinimized:"))
}

// A Boolean value that determines whether this window is in a minimized state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535028-accessibilityminimized?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityMinimized(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityMinimized:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityMinimized() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityMinimized"))
}

// A Boolean value that determines whether this window is in a minimized state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535028-accessibilityminimized?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityMinimized() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityMinimized"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityElement() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityElement:"))
}

// A Boolean value that determines whether the accessibility element participates in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535002-accessibilityelement?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityElement(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityElement:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityElement() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityElement"))
}

// A Boolean value that determines whether the accessibility element participates in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535002-accessibilityelement?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityElement() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityElement"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDisclosedRows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDisclosedRows:"))
}

// The rows that the current row discloses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535008-accessibilitydisclosedrows?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDisclosedRows(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDisclosedRows:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDisclosedRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDisclosedRows"))
}

// The rows that the current row discloses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535008-accessibilitydisclosedrows?language=objc
func (a_ AccessibilityWrapper) AccessibilityDisclosedRows() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityDisclosedRows"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFullScreenButton() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFullScreenButton:"))
}

// The child accessibility element that represents the window’s full-screen button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534999-accessibilityfullscreenbutton?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFullScreenButton(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFullScreenButton:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityFullScreenButton() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFullScreenButton"))
}

// The child accessibility element that represents the window’s full-screen button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534999-accessibilityfullscreenbutton?language=objc
func (a_ AccessibilityWrapper) AccessibilityFullScreenButton() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityFullScreenButton"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityFrontmost() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFrontmost:"))
}

// A Boolean value that determines whether the app is the frontmost app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535073-accessibilityfrontmost?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityFrontmost(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFrontmost:"), value)
}

func (a_ AccessibilityWrapper) HasIsAccessibilityFrontmost() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityFrontmost"))
}

// A Boolean value that determines whether the app is the frontmost app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535073-accessibilityfrontmost?language=objc
func (a_ AccessibilityWrapper) IsAccessibilityFrontmost() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityLabelUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityLabelUIElements:"))
}

// The child label elements for the slider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534954-accessibilitylabeluielements?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityLabelUIElements(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityLabelUIElements:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityLabelUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabelUIElements"))
}

// The child label elements for the slider accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1534954-accessibilitylabeluielements?language=objc
func (a_ AccessibilityWrapper) AccessibilityLabelUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityLabelUIElements"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityDisclosureLevel() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityDisclosureLevel:"))
}

// The indention level for the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535111-accessibilitydisclosurelevel?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityDisclosureLevel(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityDisclosureLevel:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityDisclosureLevel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDisclosureLevel"))
}

// The indention level for the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535111-accessibilitydisclosurelevel?language=objc
func (a_ AccessibilityWrapper) AccessibilityDisclosureLevel() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityPlaceholderValue() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityPlaceholderValue:"))
}

// The placeholder value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535063-accessibilityplaceholdervalue?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityPlaceholderValue(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityPlaceholderValue:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityPlaceholderValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPlaceholderValue"))
}

// The placeholder value for the accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535063-accessibilityplaceholdervalue?language=objc
func (a_ AccessibilityWrapper) AccessibilityPlaceholderValue() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityPlaceholderValue"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityUnits() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityUnits:"))
}

// The units for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535029-accessibilityunits?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityUnits(value AccessibilityUnits) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityUnits:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityUnits() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityUnits"))
}

// The units for the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535029-accessibilityunits?language=objc
func (a_ AccessibilityWrapper) AccessibilityUnits() AccessibilityUnits {
	rv := objc.Call[AccessibilityUnits](a_, objc.Sel("accessibilityUnits"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilityHandles() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityHandles:"))
}

// The drag handle accessibility elements for the layout item element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535085-accessibilityhandles?language=objc
func (a_ AccessibilityWrapper) SetAccessibilityHandles(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityHandles:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilityHandles() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityHandles"))
}

// The drag handle accessibility elements for the layout item element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535085-accessibilityhandles?language=objc
func (a_ AccessibilityWrapper) AccessibilityHandles() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityHandles"))
	return rv
}

func (a_ AccessibilityWrapper) HasSetAccessibilitySelectedRows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedRows:"))
}

// The currently selected rows for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535125-accessibilityselectedrows?language=objc
func (a_ AccessibilityWrapper) SetAccessibilitySelectedRows(value []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedRows:"), value)
}

func (a_ AccessibilityWrapper) HasAccessibilitySelectedRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedRows"))
}

// The currently selected rows for the table or outline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibility/1535125-accessibilityselectedrows?language=objc
func (a_ AccessibilityWrapper) AccessibilitySelectedRows() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedRows"))
	return rv
}
