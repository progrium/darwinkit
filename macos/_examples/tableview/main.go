package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

const (
	RowHeight = 20
)

type App struct {
	app          appkit.Application
	windowClosed bool
	doStop       func()
}

func (a *App) Stop() {
	a.doStop()
}

func newApp(app appkit.Application) *App {
	return &App{
		app: app,
	}
}

func main() {
	runtime.LockOSThread()

	app := appkit.Application_SharedApplication()
	tvApp := newApp(app)
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		tvApp.mainWindow()
	})
	delegate.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		tvApp.setMainMenu()
	})
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	delegate.SetApplicationWillTerminate(func(foundation.Notification) {
		tvApp.Stop()
	})

	app.SetDelegate(delegate)
	app.Run()
}

type TableViewDataSource struct {
	values        chan [][2]objc.Object
	data          [][2]objc.Object
	delegate      *TableViewDataSourceDelegate
	wg            *sync.WaitGroup
	maxRows       int
	overflowValue int
}

func (t *TableViewDataSource) Wait() {
	t.wg.Wait()
}

func (t *TableViewDataSource) Start(ctx context.Context, tableView appkit.TableView) {
	t.wg.Add(1)
	go valueGenerator(ctx, t.wg, t.values, tableView, t.maxRows, t.overflowValue)
}

func valueGenerator(ctx context.Context, wg *sync.WaitGroup, values chan<- [][2]objc.Object, tableView appkit.TableView, maxRows int, overflowValue int) {
	defer wg.Done()
	currentValues := make([][2]int, maxRows)
	sendToUI := func() {
		snapshot := make([][2]objc.Object, len(currentValues))
		for i, v := range currentValues {
			snapshot[i] = [2]objc.Object{foundation.String_StringWithString(fmt.Sprintf("%d", v[0])).Object, foundation.String_StringWithString(fmt.Sprintf("%d", v[1])).Object}
		}
		select {
		case <-ctx.Done():
			return
		case values <- snapshot:
			dispatch.MainQueue().DispatchAsync(func() {
				tableView.SetNeedsDisplay(true)
			})
		}
	}
	counter := 0
	row := 0
	for {
		row = counter / 2
		if row >= maxRows {
			row = 0
			break
		}
		currentValues[row][counter%2] = counter % overflowValue
		counter++
	}
	sendToUI()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(10 * time.Millisecond)
			row = (counter / 2) % maxRows
			currentValues[row][counter%2] = counter % overflowValue
			counter++
			sendToUI()
		}
	}

}

func (t *TableViewDataSource) getLatest(ctx context.Context) {
	switch t.data {
	case nil:
		select {
		case t.data = <-t.values:
		case <-ctx.Done():
			t.data = nil
			return
		}
	default:
		select {
		case t.data = <-t.values:
		case <-ctx.Done():
			t.data = nil
			return
		default:
		}
	}
}

func NewTableViewDataSource(ctx context.Context, maxRows int, overflowValue int) *TableViewDataSource {
	model := &TableViewDataSource{
		maxRows:       maxRows,
		overflowValue: overflowValue,
		values:        make(chan [][2]objc.Object, 1),
		data:          nil,
		delegate:      &TableViewDataSourceDelegate{},
		wg:            &sync.WaitGroup{},
	}
	model.delegate.SetTableViewObjectValueForTableColumnRow(func(tableView appkit.TableView, tableColumn appkit.TableColumn, row int) objc.Object {
		model.getLatest(ctx)
		if model.data == nil {
			return objc.ObjectFrom(nil)
		}
		switch tableColumn.Identifier() {
		case "Column1":
			return model.data[row][0]
		case "Column2":
			return model.data[row][1]
		}
		panic("unknown column")
	})
	model.delegate.SetNumberOfRowsInTableView(func(tableView appkit.TableView) int {
		model.getLatest(ctx)
		return len(model.data)
	})
	return model
}

func (a *App) mainWindow() {
	w := appkit.NewWindowWithSize(300, 400)
	objc.Retain(&w)
	a.windowClosed = false

	wd := &appkit.WindowDelegate{}
	wd.SetWindowWillClose(func(notification foundation.Notification) {
		a.windowClosed = true
	})
	w.SetDelegate(wd)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	dataSource := NewTableViewDataSource(ctx, 40, 13)
	a.doStop = func() {
		cancel()
		dataSource.Wait()
		if !a.windowClosed {
			w.Close()
		}
	}

	w.SetTitle("Test table view")

	tableView := appkit.NewTableView()
	go dataSource.Start(ctx, tableView)
	tableView.SetRowHeight(RowHeight)
	tableView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tableView.SetHeaderView(appkit.NewTableHeaderViewWithFrame(rectOf(0, 0, 0, RowHeight)))
	tableView.SetGridStyleMask(appkit.TableViewSolidVerticalGridLineMask | appkit.TableViewSolidHorizontalGridLineMask)
	tableView.SetStyle(appkit.TableViewStylePlain)
	tableView.SetRowSizeStyle(appkit.TableViewRowSizeStyleDefault)
	tableView.SetColumnAutoresizingStyle(appkit.TableViewUniformColumnAutoresizingStyle)
	tableView.SetUsesAlternatingRowBackgroundColors(true)
	tableView.SetStyle(appkit.TableViewStyleFullWidth)
	tableView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tableColumn1 := appkit.NewTableColumn().InitWithIdentifier("Column1")
	tableColumn1.SetTitle("Test 1")
	tableColumn1.SetWidth(100)
	tableView.AddTableColumn(tableColumn1)
	tableColumn2 := appkit.NewTableColumn().InitWithIdentifier("Column2")
	tableColumn2.SetTitle("Test 2")
	tableColumn2.SetWidth(100)
	tableView.AddTableColumn(tableColumn2)
	tableView.SetDataSource(dataSource.delegate)
	tableView.SetAllowsColumnSelection(true)
	tableView.SetAutoresizingMask(appkit.ViewWidthSizable | appkit.ViewHeightSizable)

	tsv := appkit.NewScrollView()
	tsv.SetFrameSize(foundation.Size{Width: w.ContentView().Frame().Size.Width, Height: w.ContentView().Frame().Size.Height})
	tsv.SetFrameOrigin(foundation.Point{X: w.ContentView().Frame().Origin.X, Y: w.ContentView().Frame().Origin.Y})
	tsv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tsv.SetDocumentView(tableView)
	tsv.SetHasHorizontalScroller(true)
	tsv.SetHasVerticalScroller(true)
	tsv.SetAutohidesScrollers(true)
	tsv.SetTranslatesAutoresizingMaskIntoConstraints(true)
	tsv.SetAutoresizingMask(appkit.ViewWidthSizable | appkit.ViewHeightSizable)
	w.ContentView().AddSubview(tsv)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchorConstant(tsv.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchorConstant(tsv.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchorConstant(tsv.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchorConstant(tsv.BottomAnchor(), 10).SetActive(true)

	w.Center()
	w.MakeKeyAndOrderFront(nil)

	a.app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	a.app.ActivateIgnoringOtherApps(true)
}

func (a *App) setMainMenu() {
	menuBar := appkit.NewMenuWithTitle("Table View")
	a.app.SetMainMenu(menuBar)

	appMenuItem := appkit.NewMenuItemWithSelector("Table View", "", objc.Selector{})
	appMenu := appkit.NewMenuWithTitle("Table View")
	appMenu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", func(sender objc.Object) { a.app.Terminate(nil) }))
	appMenuItem.SetSubmenu(appMenu)
	menuBar.AddItem(appMenuItem)
}

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
