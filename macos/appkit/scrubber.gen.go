// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Scrubber] class.
var ScrubberClass = _ScrubberClass{objc.GetClass("NSScrubber")}

type _ScrubberClass struct {
	objc.Class
}

// An interface definition for the [Scrubber] class.
type IScrubber interface {
	IView
	RegisterNibForItemIdentifier(nib INib, itemIdentifier UserInterfaceItemIdentifier)
	MakeItemWithIdentifierOwner(itemIdentifier UserInterfaceItemIdentifier, owner objc.IObject) ScrubberItemView
	MoveItemAtIndexToIndex(oldIndex int, newIndex int)
	RegisterClassForItemIdentifier(itemViewClass objc.IClass, itemIdentifier UserInterfaceItemIdentifier)
	ScrollItemAtIndexToAlignment(index int, alignment ScrubberAlignment)
	ItemViewForItemAtIndex(index int) ScrubberItemView
	RemoveItemsAtIndexes(indexes foundation.IIndexSet)
	InsertItemsAtIndexes(indexes foundation.IIndexSet)
	ReloadItemsAtIndexes(indexes foundation.IIndexSet)
	PerformSequentialBatchUpdates(updateBlock func())
	ReloadData()
	DataSource() ScrubberDataSourceWrapper
	SetDataSource(value PScrubberDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	BackgroundView() View
	SetBackgroundView(value IView)
	IsContinuous() bool
	SetContinuous(value bool)
	HighlightedIndex() int
	Delegate() ScrubberDelegateWrapper
	SetDelegate(value PScrubberDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ShowsAdditionalContentIndicators() bool
	SetShowsAdditionalContentIndicators(value bool)
	SelectionBackgroundStyle() ScrubberSelectionStyle
	SetSelectionBackgroundStyle(value IScrubberSelectionStyle)
	ScrubberLayout() ScrubberLayout
	SetScrubberLayout(value IScrubberLayout)
	NumberOfItems() int
	SelectionOverlayStyle() ScrubberSelectionStyle
	SetSelectionOverlayStyle(value IScrubberSelectionStyle)
	Mode() ScrubberMode
	SetMode(value ScrubberMode)
	ItemAlignment() ScrubberAlignment
	SetItemAlignment(value ScrubberAlignment)
	FloatsSelectionViews() bool
	SetFloatsSelectionViews(value bool)
	SelectedIndex() int
	SetSelectedIndex(value int)
	ShowsArrowButtons() bool
	SetShowsArrowButtons(value bool)
}

// A customizable item picker control for the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber?language=objc
type Scrubber struct {
	View
}

func ScrubberFrom(ptr unsafe.Pointer) Scrubber {
	return Scrubber{
		View: ViewFrom(ptr),
	}
}

func (s_ Scrubber) InitWithFrame(frameRect foundation.Rect) Scrubber {
	rv := objc.Call[Scrubber](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated scrubber object with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588263-initwithframe?language=objc
func Scrubber_InitWithFrame(frameRect foundation.Rect) Scrubber {
	return ScrubberClass.Alloc().InitWithFrame(frameRect)
}

func (sc _ScrubberClass) Alloc() Scrubber {
	rv := objc.Call[Scrubber](sc, objc.Sel("alloc"))
	return rv
}

func Scrubber_Alloc() Scrubber {
	return ScrubberClass.Alloc()
}

func (sc _ScrubberClass) New() Scrubber {
	rv := objc.Call[Scrubber](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubber() Scrubber {
	return ScrubberClass.New()
}

func (s_ Scrubber) Init() Scrubber {
	rv := objc.Call[Scrubber](s_, objc.Sel("init"))
	return rv
}

// Registers a nib file for the scrubber to use when it creates new items in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2646903-registernib?language=objc
func (s_ Scrubber) RegisterNibForItemIdentifier(nib INib, itemIdentifier UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](s_, objc.Sel("registerNib:forItemIdentifier:"), objc.Ptr(nib), itemIdentifier)
}

// Creates or returns a reusable item object with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544776-makeitemwithidentifier?language=objc
func (s_ Scrubber) MakeItemWithIdentifierOwner(itemIdentifier UserInterfaceItemIdentifier, owner objc.IObject) ScrubberItemView {
	rv := objc.Call[ScrubberItemView](s_, objc.Sel("makeItemWithIdentifier:owner:"), itemIdentifier, owner)
	return rv
}

// Moves an item from one index to another in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544864-moveitematindex?language=objc
func (s_ Scrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("moveItemAtIndex:toIndex:"), oldIndex, newIndex)
}

// Registers a class for the scrubber to use when it creates new items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544773-registerclass?language=objc
func (s_ Scrubber) RegisterClassForItemIdentifier(itemViewClass objc.IClass, itemIdentifier UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](s_, objc.Sel("registerClass:forItemIdentifier:"), objc.Ptr(itemViewClass), itemIdentifier)
}

// Scrolls an item to a specified alignment within the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544724-scrollitematindex?language=objc
func (s_ Scrubber) ScrollItemAtIndexToAlignment(index int, alignment ScrubberAlignment) {
	objc.Call[objc.Void](s_, objc.Sel("scrollItemAtIndex:toAlignment:"), index, alignment)
}

// Returns the view for the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544862-itemviewforitematindex?language=objc
func (s_ Scrubber) ItemViewForItemAtIndex(index int) ScrubberItemView {
	rv := objc.Call[ScrubberItemView](s_, objc.Sel("itemViewForItemAtIndex:"), index)
	return rv
}

// Removes the items at the specified indexes from the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544679-removeitemsatindexes?language=objc
func (s_ Scrubber) RemoveItemsAtIndexes(indexes foundation.IIndexSet) {
	objc.Call[objc.Void](s_, objc.Sel("removeItemsAtIndexes:"), objc.Ptr(indexes))
}

// Inserts new items at the specified indexes into the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544867-insertitemsatindexes?language=objc
func (s_ Scrubber) InsertItemsAtIndexes(indexes foundation.IIndexSet) {
	objc.Call[objc.Void](s_, objc.Sel("insertItemsAtIndexes:"), objc.Ptr(indexes))
}

// Reloads the items at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544874-reloaditemsatindexes?language=objc
func (s_ Scrubber) ReloadItemsAtIndexes(indexes foundation.IIndexSet) {
	objc.Call[objc.Void](s_, objc.Sel("reloadItemsAtIndexes:"), objc.Ptr(indexes))
}

// Combines multiple scrubber content updates into a single action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2646906-performsequentialbatchupdates?language=objc
func (s_ Scrubber) PerformSequentialBatchUpdates(updateBlock func()) {
	objc.Call[objc.Void](s_, objc.Sel("performSequentialBatchUpdates:"), updateBlock)
}

// Reloads the content of the entire scrubber, and deselects the currently selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544712-reloaddata?language=objc
func (s_ Scrubber) ReloadData() {
	objc.Call[objc.Void](s_, objc.Sel("reloadData"))
}

// The object that provides the data for the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544804-datasource?language=objc
func (s_ Scrubber) DataSource() ScrubberDataSourceWrapper {
	rv := objc.Call[ScrubberDataSourceWrapper](s_, objc.Sel("dataSource"))
	return rv
}

// The object that provides the data for the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544804-datasource?language=objc
func (s_ Scrubber) SetDataSource(value PScrubberDataSource) {
	po0 := objc.WrapAsProtocol("NSScrubberDataSource", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDataSource:"), po0)
}

// The object that provides the data for the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544804-datasource?language=objc
func (s_ Scrubber) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// A view that is displayed behind the scrubber content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544847-backgroundview?language=objc
func (s_ Scrubber) BackgroundView() View {
	rv := objc.Call[View](s_, objc.Sel("backgroundView"))
	return rv
}

// A view that is displayed behind the scrubber content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544847-backgroundview?language=objc
func (s_ Scrubber) SetBackgroundView(value IView) {
	objc.Call[objc.Void](s_, objc.Sel("setBackgroundView:"), objc.Ptr(value))
}

// A Boolean value that, together with the mode property, determines scrubber interaction style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544694-continuous?language=objc
func (s_ Scrubber) IsContinuous() bool {
	rv := objc.Call[bool](s_, objc.Sel("isContinuous"))
	return rv
}

// A Boolean value that, together with the mode property, determines scrubber interaction style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544694-continuous?language=objc
func (s_ Scrubber) SetContinuous(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setContinuous:"), value)
}

// The index of the highlighted item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588262-highlightedindex?language=objc
func (s_ Scrubber) HighlightedIndex() int {
	rv := objc.Call[int](s_, objc.Sel("highlightedIndex"))
	return rv
}

// The object that acts as the delegate of the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544738-delegate?language=objc
func (s_ Scrubber) Delegate() ScrubberDelegateWrapper {
	rv := objc.Call[ScrubberDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The object that acts as the delegate of the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544738-delegate?language=objc
func (s_ Scrubber) SetDelegate(value PScrubberDelegate) {
	po0 := objc.WrapAsProtocol("NSScrubberDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The object that acts as the delegate of the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544738-delegate?language=objc
func (s_ Scrubber) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The color displayed behind the scrubber content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544869-backgroundcolor?language=objc
func (s_ Scrubber) BackgroundColor() Color {
	rv := objc.Call[Color](s_, objc.Sel("backgroundColor"))
	return rv
}

// The color displayed behind the scrubber content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544869-backgroundcolor?language=objc
func (s_ Scrubber) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544827-showsadditionalcontentindicators?language=objc
func (s_ Scrubber) ShowsAdditionalContentIndicators() bool {
	rv := objc.Call[bool](s_, objc.Sel("showsAdditionalContentIndicators"))
	return rv
}

// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544827-showsadditionalcontentindicators?language=objc
func (s_ Scrubber) SetShowsAdditionalContentIndicators(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setShowsAdditionalContentIndicators:"), value)
}

// The style applied to the background of selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588270-selectionbackgroundstyle?language=objc
func (s_ Scrubber) SelectionBackgroundStyle() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](s_, objc.Sel("selectionBackgroundStyle"))
	return rv
}

// The style applied to the background of selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588270-selectionbackgroundstyle?language=objc
func (s_ Scrubber) SetSelectionBackgroundStyle(value IScrubberSelectionStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectionBackgroundStyle:"), objc.Ptr(value))
}

// An object used to describe the layout of items within the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544850-scrubberlayout?language=objc
func (s_ Scrubber) ScrubberLayout() ScrubberLayout {
	rv := objc.Call[ScrubberLayout](s_, objc.Sel("scrubberLayout"))
	return rv
}

// An object used to describe the layout of items within the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544850-scrubberlayout?language=objc
func (s_ Scrubber) SetScrubberLayout(value IScrubberLayout) {
	objc.Call[objc.Void](s_, objc.Sel("setScrubberLayout:"), objc.Ptr(value))
}

// The number of items represented by the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544836-numberofitems?language=objc
func (s_ Scrubber) NumberOfItems() int {
	rv := objc.Call[int](s_, objc.Sel("numberOfItems"))
	return rv
}

// The style overlaid on selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588265-selectionoverlaystyle?language=objc
func (s_ Scrubber) SelectionOverlayStyle() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](s_, objc.Sel("selectionOverlayStyle"))
	return rv
}

// The style overlaid on selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588265-selectionoverlaystyle?language=objc
func (s_ Scrubber) SetSelectionOverlayStyle(value IScrubberSelectionStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectionOverlayStyle:"), objc.Ptr(value))
}

// A setting that determines whether interaction with the scrubber is fixed or free. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544687-mode?language=objc
func (s_ Scrubber) Mode() ScrubberMode {
	rv := objc.Call[ScrubberMode](s_, objc.Sel("mode"))
	return rv
}

// A setting that determines whether interaction with the scrubber is fixed or free. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544687-mode?language=objc
func (s_ Scrubber) SetMode(value ScrubberMode) {
	objc.Call[objc.Void](s_, objc.Sel("setMode:"), value)
}

// A setting that specifies the snapping behavior of items in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544782-itemalignment?language=objc
func (s_ Scrubber) ItemAlignment() ScrubberAlignment {
	rv := objc.Call[ScrubberAlignment](s_, objc.Sel("itemAlignment"))
	return rv
}

// A setting that specifies the snapping behavior of items in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544782-itemalignment?language=objc
func (s_ Scrubber) SetItemAlignment(value ScrubberAlignment) {
	objc.Call[objc.Void](s_, objc.Sel("setItemAlignment:"), value)
}

// A Boolean value that determines the behavior of the item selection decorations as the scrubber's selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2646904-floatsselectionviews?language=objc
func (s_ Scrubber) FloatsSelectionViews() bool {
	rv := objc.Call[bool](s_, objc.Sel("floatsSelectionViews"))
	return rv
}

// A Boolean value that determines the behavior of the item selection decorations as the scrubber's selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2646904-floatsselectionviews?language=objc
func (s_ Scrubber) SetFloatsSelectionViews(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setFloatsSelectionViews:"), value)
}

// The index of the selected item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588260-selectedindex?language=objc
func (s_ Scrubber) SelectedIndex() int {
	rv := objc.Call[int](s_, objc.Sel("selectedIndex"))
	return rv
}

// The index of the selected item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2588260-selectedindex?language=objc
func (s_ Scrubber) SetSelectedIndex(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectedIndex:"), value)
}

// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544865-showsarrowbuttons?language=objc
func (s_ Scrubber) ShowsArrowButtons() bool {
	rv := objc.Call[bool](s_, objc.Sel("showsArrowButtons"))
	return rv
}

// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/2544865-showsarrowbuttons?language=objc
func (s_ Scrubber) SetShowsArrowButtons(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setShowsArrowButtons:"), value)
}
