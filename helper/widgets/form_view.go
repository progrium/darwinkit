package widgets

import (
	"github.com/progrium/macdriver/macos/appkit"
)

type LabelAlignment int

const (
	LabelAlignmentLeading  LabelAlignment = 0
	LabelAlignmentTrailing LabelAlignment = 1
	LabelAlignmentCenter   LabelAlignment = 2
)

// FormView ia an appkit.View that arrange form field name and controls.
type FormView struct {
	appkit.GridView

	labelFont appkit.IFont
}

// NewFormView create new form view
func NewFormView() *FormView {
	gv := appkit.GridViewClass.GridViewWithNumberOfColumnsRows(2, 0)
	gv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	gv.SetContentHuggingPriorityForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationHorizontal)
	gv.SetContentHuggingPriorityForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationVertical)

	return &FormView{
		GridView: gv,
	}
}

// AddExpandRow add a row, expand to fill parent view height.
func (f *FormView) AddExpandRow() {
	v1 := appkit.NewView()
	v2 := appkit.NewView()
	f.AddRowWithViews([]appkit.IView{v1, v2})
}

// AddRow add a new form row
func (f *FormView) AddRow(name string, control appkit.IControl) {
	label := f.newLabel(name)
	control.SetContentCompressionResistancePriorityForOrientation(appkit.LayoutPriorityRequired, appkit.LayoutConstraintOrientationVertical)
	f.AddRowWithViews([]appkit.IView{label, control})
}

// InsertRow insert a new form row at specific location
func (f *FormView) InsertRow(index int, name string, control appkit.IControl) {
	if index > f.NumberOfRows() {
		panic("index out of row range")
	}
	label := f.newLabel(name)
	f.InsertRowAtIndexWithViews(index, []appkit.IView{label, control})
}

// SetRowSpacing set spacing between rows.
func (f *FormView) SetRowSpacing(spacing float64) {
	f.SetRowSpacing(spacing)
}

// SetLabelWidth set width for labels
func (f *FormView) SetLabelWidth(width float64) {
	f.ColumnAtIndex(0).SetWidth(width)
}

// SetLabelAlignment set label text alignment
func (f *FormView) SetLabelAlignment(alignment LabelAlignment) {
	switch alignment {
	case LabelAlignmentLeading:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementLeading)
	case LabelAlignmentTrailing:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementTrailing)
	case LabelAlignmentCenter:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementCenter)
	}
}

// SetLabelFont set label font
func (f *FormView) SetLabelFont(font appkit.IFont) {
	f.labelFont = font
	labelColumn := f.ColumnAtIndex(0)
	size := labelColumn.NumberOfCells()
	for i := 0; i < size; i++ {
		v := labelColumn.CellAtIndex(i).ContentView()
		label := appkit.MakeTextField(v.Ptr())
		label.SetFont(font)
	}
}

// SetLabelControlSpacing set spacing between label and control
func (f *FormView) SetLabelControlSpacing(spacing float64) {
	f.SetColumnSpacing(spacing)
}

func (f *FormView) newLabel(name string) appkit.ITextField {
	label := appkit.NewLabel(name)
	if f.labelFont != nil && f.labelFont.Ptr() != nil {
		label.SetFont(f.labelFont)
	}

	return label
}
