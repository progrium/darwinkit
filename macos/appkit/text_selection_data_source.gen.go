// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that objects implement to provide data for, and manage text selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource?language=objc
type PTextSelectionDataSource interface {
	// optional
	BaseWritingDirectionAtLocation(location TextLocationWrapper) TextSelectionNavigationWritingDirection
	HasBaseWritingDirectionAtLocation() bool

	// optional
	LineFragmentRangeForPointInContainerAtLocation(point coregraphics.Point, location TextLocationWrapper) ITextRange
	HasLineFragmentRangeForPointInContainerAtLocation() bool

	// optional
	EnumerateContainerBoundariesFromLocationReverseUsingBlock(location TextLocationWrapper, reverse bool, block func(boundaryLocation TextLocationWrapper, stop *bool))
	HasEnumerateContainerBoundariesFromLocationReverseUsingBlock() bool

	// optional
	EnumerateSubstringsFromLocationOptionsUsingBlock(location TextLocationWrapper, options foundation.StringEnumerationOptions, block func(substring string, substringRange TextRange, enclosingRange TextRange, stop *bool))
	HasEnumerateSubstringsFromLocationOptionsUsingBlock() bool

	// optional
	EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location TextLocationWrapper, block func(caretOffset float64, location TextLocationWrapper, leadingEdge bool, stop *bool))
	HasEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock() bool

	// optional
	TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity TextSelectionGranularity, location TextLocationWrapper) ITextRange
	HasTextRangeForSelectionGranularityEnclosingLocation() bool

	// optional
	OffsetFromLocationToLocation(from TextLocationWrapper, to TextLocationWrapper) int
	HasOffsetFromLocationToLocation() bool

	// optional
	LocationFromLocationWithOffset(location TextLocationWrapper, offset int) PTextLocation
	HasLocationFromLocationWithOffset() bool

	// optional
	TextLayoutOrientationAtLocation(location TextLocationWrapper) TextSelectionNavigationLayoutOrientation
	HasTextLayoutOrientationAtLocation() bool

	// optional
	DocumentRange() ITextRange
	HasDocumentRange() bool
}

// A concrete type wrapper for the [PTextSelectionDataSource] protocol.
type TextSelectionDataSourceWrapper struct {
	objc.Object
}

func (t_ TextSelectionDataSourceWrapper) HasBaseWritingDirectionAtLocation() bool {
	return t_.RespondsToSelector(objc.Sel("baseWritingDirectionAtLocation:"))
}

// Returns the base writing direction at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801839-basewritingdirectionatlocation?language=objc
func (t_ TextSelectionDataSourceWrapper) BaseWritingDirectionAtLocation(location PTextLocation) TextSelectionNavigationWritingDirection {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextSelectionNavigationWritingDirection](t_, objc.Sel("baseWritingDirectionAtLocation:"), po0)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasLineFragmentRangeForPointInContainerAtLocation() bool {
	return t_.RespondsToSelector(objc.Sel("lineFragmentRangeForPoint:inContainerAtLocation:"))
}

// Returns the range of the line fragment that contains the point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801845-linefragmentrangeforpoint?language=objc
func (t_ TextSelectionDataSourceWrapper) LineFragmentRangeForPointInContainerAtLocation(point coregraphics.Point, location PTextLocation) TextRange {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextRange](t_, objc.Sel("lineFragmentRangeForPoint:inContainerAtLocation:"), point, po1)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasEnumerateContainerBoundariesFromLocationReverseUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("enumerateContainerBoundariesFromLocation:reverse:usingBlock:"))
}

// Enumerates all the container boundaries starting from the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801842-enumeratecontainerboundariesfrom?language=objc
func (t_ TextSelectionDataSourceWrapper) EnumerateContainerBoundariesFromLocationReverseUsingBlock(location PTextLocation, reverse bool, block func(boundaryLocation TextLocationWrapper, stop *bool)) {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	objc.Call[objc.Void](t_, objc.Sel("enumerateContainerBoundariesFromLocation:reverse:usingBlock:"), po0, reverse, block)
}

func (t_ TextSelectionDataSourceWrapper) HasEnumerateSubstringsFromLocationOptionsUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("enumerateSubstringsFromLocation:options:usingBlock:"))
}

// Enumerates the textual segment boundaries starting at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801843-enumeratesubstringsfromlocation?language=objc
func (t_ TextSelectionDataSourceWrapper) EnumerateSubstringsFromLocationOptionsUsingBlock(location PTextLocation, options foundation.StringEnumerationOptions, block func(substring string, substringRange TextRange, enclosingRange TextRange, stop *bool)) {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	objc.Call[objc.Void](t_, objc.Sel("enumerateSubstringsFromLocation:options:usingBlock:"), po0, options, block)
}

func (t_ TextSelectionDataSourceWrapper) HasEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("enumerateCaretOffsetsInLineFragmentAtLocation:usingBlock:"))
}

// Enumerates all the insertion point caret offsets from left to right in visual order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801841-enumeratecaretoffsetsinlinefragm?language=objc
func (t_ TextSelectionDataSourceWrapper) EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location PTextLocation, block func(caretOffset float64, location TextLocationWrapper, leadingEdge bool, stop *bool)) {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	objc.Call[objc.Void](t_, objc.Sel("enumerateCaretOffsetsInLineFragmentAtLocation:usingBlock:"), po0, block)
}

func (t_ TextSelectionDataSourceWrapper) HasTextRangeForSelectionGranularityEnclosingLocation() bool {
	return t_.RespondsToSelector(objc.Sel("textRangeForSelectionGranularity:enclosingLocation:"))
}

// Returns a text range that corresponds to selection granularity of the enclosing location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801848-textrangeforselectiongranularity?language=objc
func (t_ TextSelectionDataSourceWrapper) TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity TextSelectionGranularity, location PTextLocation) TextRange {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextRange](t_, objc.Sel("textRangeForSelectionGranularity:enclosingLocation:"), selectionGranularity, po1)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasOffsetFromLocationToLocation() bool {
	return t_.RespondsToSelector(objc.Sel("offsetFromLocation:toLocation:"))
}

// Returns the offset between the two locations you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801847-offsetfromlocation?language=objc
func (t_ TextSelectionDataSourceWrapper) OffsetFromLocationToLocation(from PTextLocation, to PTextLocation) int {
	po0 := objc.WrapAsProtocol("NSTextLocation", from)
	po1 := objc.WrapAsProtocol("NSTextLocation", to)
	rv := objc.Call[int](t_, objc.Sel("offsetFromLocation:toLocation:"), po0, po1)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasLocationFromLocationWithOffset() bool {
	return t_.RespondsToSelector(objc.Sel("locationFromLocation:withOffset:"))
}

// Returns a new location using the location and offset you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801846-locationfromlocation?language=objc
func (t_ TextSelectionDataSourceWrapper) LocationFromLocationWithOffset(location PTextLocation, offset int) TextLocationWrapper {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("locationFromLocation:withOffset:"), po0, offset)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasTextLayoutOrientationAtLocation() bool {
	return t_.RespondsToSelector(objc.Sel("textLayoutOrientationAtLocation:"))
}

// Returns the layout orientation at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3852580-textlayoutorientationatlocation?language=objc
func (t_ TextSelectionDataSourceWrapper) TextLayoutOrientationAtLocation(location PTextLocation) TextSelectionNavigationLayoutOrientation {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextSelectionNavigationLayoutOrientation](t_, objc.Sel("textLayoutOrientationAtLocation:"), po0)
	return rv
}

func (t_ TextSelectionDataSourceWrapper) HasDocumentRange() bool {
	return t_.RespondsToSelector(objc.Sel("documentRange"))
}

// Returns the starting and ending locations for the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselectiondatasource/3801840-documentrange?language=objc
func (t_ TextSelectionDataSourceWrapper) DocumentRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("documentRange"))
	return rv
}
