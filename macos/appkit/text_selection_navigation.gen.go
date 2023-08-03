// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}

type _TextSelectionNavigationClass struct {
	objc.Class
}

type ITextSelectionNavigation interface {
	objc.IObject
	TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection
	DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection
	FlushLayoutCache()
	DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
}

type TextSelectionNavigation struct {
	objc.Object
}

func MakeTextSelectionNavigation(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](tc, objc.GetSelector("alloc"))
	return rv
}

func TextSelectionNavigation_Alloc() TextSelectionNavigation {
	return TextSelectionNavigationClass.Alloc()
}

func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextSelectionNavigation() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

func TextSelectionNavigation_New() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](t_, objc.GetSelector("init"))
	return rv
}

func TextSelectionNavigation_Init() TextSelectionNavigation {
	return TextSelectionNavigationClass.Alloc().Init()
}

func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("textSelectionForSelectionGranularity:enclosingTextSelection:"), selectionGranularity, objc.ExtractPtr(textSelection))
	return rv
}

func (t_ TextSelectionNavigation) DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("destinationSelectionForTextSelection:direction:destination:extending:confined:"), objc.ExtractPtr(textSelection), direction, destination, extending, confined)
	return rv
}

func (t_ TextSelectionNavigation) FlushLayoutCache() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("flushLayoutCache"))
}

func (t_ TextSelectionNavigation) DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange {
	rv := objc.CallMethod[[]TextRange](t_, objc.GetSelector("deletionRangesForTextSelection:direction:destination:allowsDecomposition:"), objc.ExtractPtr(textSelection), direction, destination, allowsDecomposition)
	// TODO: convert slice items...
	return rv
}

func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsNonContiguousRanges"))
	return rv
}

func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsNonContiguousRanges:"), value)
}

func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("rotatesCoordinateSystemForLayoutOrientation"))
	return rv
}

func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRotatesCoordinateSystemForLayoutOrientation:"), value)
}
