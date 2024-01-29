package main

import (
	"errors"
	"fmt"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/screencapturekit"
)

func main() {
	captureHandler := &screenCaptureHandler{}
	captureHandler.refreshCapturableWindows()

	sc := screencapturekit.NewStreamConfiguration()
	//cf := screencapturekit.NewContentFilter()

	dispatch.MainQueue().DispatchAsync(func() {
	cf := captureHandler.GetContentFilter()
	s := screencapturekit.NewStreamWithFilterConfigurationDelegate(cf, sc, captureHandler)

	var dispatchQueue dispatch.Queue
	//dispatchQueue = dispatch.CreateQueue("com.example.queue", dispatch.QueueTypeSerial)
	dispatchQueue = dispatch.MainQueue()
	err := foundation.Error{}
	ok := s.AddStreamOutputTypeSampleHandlerQueueError(captureHandler, screencapturekit.StreamOutputTypeScreen, dispatchQueue, err)
	if !ok {
		fmt.Println("s.AddStreamOutputTypeSampleHandlerQueueError", err)
	}

	fmt.Println("s.StartCaptureWithCompletionHandler")
	s.StartCaptureWithCompletionHandler(func(err foundation.Error) {
		fmt.Println("s.StartCaptureWithCompletionHandler", err)
	})

	dispatch.Main()
}

type CaptureType int

const (
	CaptureTypeDisplay CaptureType = iota
	CaptureTypeWindow  CaptureType = iota
)

func (h *screenCaptureHandler) refreshCapturableWindows() {
	fmt.Println("listing sharable content")
	ch := make(chan struct{})
	screencapturekit.ShareableContentClass.GetShareableContentWithCompletionHandler(func(sc screencapturekit.ShareableContent, err foundation.Error) {
		defer close(ch)
		h.availabileDisplays = sc.Displays()
		if h.selectedDisplay == nil && len(h.availabileDisplays) > 0 {
			h.selectedDisplay = &h.availabileDisplays[0]
		}
		h.availbleWindows = sc.Windows()
		if h.selectedWindow == nil && len(h.availbleWindows) > 0 {
			h.selectedWindow = &h.availbleWindows[0]
		}
		for _, app := range sc.Applications() {
			fmt.Println("app", app.ApplicationName())
		}
		fmt.Println("done listing sharable content")
	})
	<-ch
}

type screenCaptureHandler struct {
	availabileDisplays []screencapturekit.Display
	availbleWindows    []screencapturekit.Window

	selectedDisplay *screencapturekit.Display
	selectedWindow  *screencapturekit.Window

	captureType CaptureType
}

func (sh *screenCaptureHandler) GetContentFilter() screencapturekit.ContentFilter {
	filter := screencapturekit.NewContentFilter()

	switch sh.captureType {
	case CaptureTypeDisplay:
		display := sh.selectedDisplay
		windows := []screencapturekit.IWindow{}
		for _, w := range sh.availbleWindows {
			windows = append(windows, w)
		}
		fmt.Println("display:", display)
		fmt.Println("windows:", windows)
		filter = screencapturekit.NewContentFilterWithDisplayIncludingWindows(display, windows)
	case CaptureTypeWindow:
	}
	return filter
}

var _ screencapturekit.PStreamOutput = (*screenCaptureHandler)(nil)
var _ screencapturekit.PStreamDelegate = (*screenCaptureHandler)(nil)

// StreamOutput methods

func (sh *screenCaptureHandler) HasStreamDidOutputSampleBufferOfType() bool {
	panic(errors.New("*streamHandler.HasStreamDidOutputSampleBufferOfType not implemented"))
}

func (sh *screenCaptureHandler) StreamDidOutputSampleBufferOfType(s screencapturekit.Stream, buf coremedia.SampleBufferRef, out screencapturekit.StreamOutputType) {
	panic(errors.New("*streamHandler.StreamDidOutputSampleBufferOfType not implemented"))
}

// StreamDelegate methods

func (sh *screenCaptureHandler) StreamDidStopWithError(s screencapturekit.Stream, err foundation.Error) {
	fmt.Println("StreamDidStopWithError", err)
}
func (sh *screenCaptureHandler) HasStreamDidStopWithError() bool {
	fmt.Println("HasStreamDidStopWithError")
	return true
}
