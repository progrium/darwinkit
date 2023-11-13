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
	streamOutputHandler := &streamHandler{}

	sc := screencapturekit.NewStreamConfiguration()
	cf := screencapturekit.NewContentFilter()
	s := screencapturekit.NewStreamWithFilterConfigurationDelegate(cf, sc, streamOutputHandler)

	var dispatchQueue dispatch.Queue
	//dispatchQueue = dispatch.CreateQueue("com.example.queue", dispatch.QueueTypeSerial)
	dispatchQueue = dispatch.MainQueue()
	err := foundation.Error{}
	ok := s.AddStreamOutputTypeSampleHandlerQueueError(streamOutputHandler, screencapturekit.StreamOutputTypeScreen, dispatchQueue, err)
	if !ok {
		fmt.Println("s.AddStreamOutputTypeSampleHandlerQueueError", err)
	}

	fmt.Println("s.StartCaptureWithCompletionHandler")
	s.StartCaptureWithCompletionHandler(func(err foundation.Error) {
		fmt.Println("s.StartCaptureWithCompletionHandler", err)
	})

	dispatch.Main()
}

type streamHandler struct{}

var _ screencapturekit.PStreamOutput = (*streamHandler)(nil)
var _ screencapturekit.PStreamDelegate = (*streamHandler)(nil)

// StreamOutput methods

func (sh *streamHandler) HasStreamDidOutputSampleBufferOfType() bool {
	panic(errors.New("*streamHandler.HasStreamDidOutputSampleBufferOfType not implemented"))
}

func (sh *streamHandler) StreamDidOutputSampleBufferOfType(s screencapturekit.Stream, buf coremedia.SampleBufferRef, out screencapturekit.StreamOutputType) {
	panic(errors.New("*streamHandler.StreamDidOutputSampleBufferOfType not implemented"))
}

// StreamDelegate methods

func (sh *streamHandler) StreamDidStopWithError(s screencapturekit.Stream, err foundation.Error) {
	fmt.Println("StreamDidStopWithError", err)
}
func (sh *streamHandler) HasStreamDidStopWithError() bool {
	fmt.Println("HasStreamDidStopWithError")
	return true
}
