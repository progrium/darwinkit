package types

import (
	"github.com/progrium/macdriver/bridge"
	"github.com/progrium/macdriver/bridge/resource"
)

type Window struct {
	*resource.Handle

	Title        string
	Position     bridge.Point
	Size         bridge.Size
	Closable     bool
	Minimizable  bool
	Resizable    bool
	Background   *bridge.Color
	Borderless   bool
	CornerRadius float64
	AlwaysOnTop  bool
	IgnoreMouse  bool
	Center       bool
	URL          string
	Image        string
}

type Indicator struct {
	*resource.Handle

	Icon string
	Text string
	Menu *Menu
}

type Menu struct {
	*resource.Handle

	Icon    string
	Title   string
	Tooltip string
	Items   []MenuItem
}

type MenuItem struct {
	*resource.Handle

	Title     string
	Icon      string
	Tooltip   string
	Separator bool
	Enabled   bool
	Checked   bool
}
