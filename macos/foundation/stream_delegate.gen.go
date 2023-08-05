// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"github.com/progrium/macdriver/objc"
)

type IStreamDelegate interface {
	ImplementsStreamHandleEvent() bool
	// optional
	StreamHandleEvent(aStream Stream, eventCode StreamEvent)
}

type StreamDelegate struct {
	_StreamHandleEvent func(aStream Stream, eventCode StreamEvent)
}

func (di *StreamDelegate) ImplementsStreamHandleEvent() bool {
	return di._StreamHandleEvent != nil
}

func (di *StreamDelegate) SetStreamHandleEvent(f func(aStream Stream, eventCode StreamEvent)) {
	di._StreamHandleEvent = f
}

func (di *StreamDelegate) StreamHandleEvent(aStream Stream, eventCode StreamEvent) {
	di._StreamHandleEvent(aStream, eventCode)
}

type StreamDelegateWrapper struct {
	objc.Object
}

func (s_ StreamDelegateWrapper) ImplementsStreamHandleEvent() bool {
	return s_.RespondsToSelector(objc.GetSelector("stream:handleEvent:"))
}

func (s_ StreamDelegateWrapper) StreamHandleEvent(aStream IStream, eventCode StreamEvent) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stream:handleEvent:"), objc.ExtractPtr(aStream), eventCode)
}
