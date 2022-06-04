package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation -framework CoreFoundation -framework WebKit -framework AppKit
#include <dispatch/dispatch.h>
#include <Foundation/Foundation.h>
#include <AppKit/AppKit.h>

void monitorGlobal(unsigned long long int mask) {
	[NSEvent addGlobalMonitorForEventsMatchingMask:mask
            handler:^(NSEvent *incomingEvent) {
		void monitorGlobalReentry(NSEvent *event);
		monitorGlobalReentry(incomingEvent);
        return;
    }];
}

void monitorLocal(unsigned long long int mask) {
	[NSEvent addLocalMonitorForEventsMatchingMask:mask
            handler:^(NSEvent *incomingEvent) {
		NSEvent* monitorLocalReentry(NSEvent *event);
		return monitorLocalReentry(incomingEvent);
    }];
}

*/
import "C"
import (
	"errors"
)

const (
	NSEventTypeLeftMouseDown      = 1
	NSEventTypeLeftMouseUp        = 2
	NSEventTypeRightMouseDown     = 3
	NSEventTypeRightMouseUp       = 4
	NSEventTypeMouseMoved         = 5
	NSEventTypeLeftMouseDragged   = 6
	NSEventTypeRightMouseDragged  = 7
	NSEventTypeMouseEntered       = 8
	NSEventTypeMouseExited        = 9
	NSEventTypeKeyDown            = 10
	NSEventTypeKeyUp              = 11
	NSEventTypeFlagsChanged       = 12
	NSEventTypeAppKitDefined      = 13
	NSEventTypeSystemDefined      = 14
	NSEventTypeApplicationDefined = 15
	NSEventTypePeriodic           = 16
	NSEventTypeCursorUpdate       = 17
	NSEventTypeRotate             = 18
	NSEventTypeBeginGesture       = 19
	NSEventTypeEndGesture         = 20
	NSEventTypeScrollWheel        = 22
	NSEventTypeTabletPoint        = 23
	NSEventTypeTabletProximity    = 24
	NSEventTypeOtherMouseDown     = 25
	NSEventTypeOtherMouseUp       = 26
	NSEventTypeOtherMouseDragged  = 27
	NSEventTypeGesture            = 29
	NSEventTypeMagnify            = 30
	NSEventTypeSwipe              = 31
	NSEventTypeSmartMagnify       = 32
	NSEventTypeQuickLook          = 33
	NSEventTypePressure           = 34
	NSEventTypeDirectTouch        = 37
	NSEventTypeChangeMode         = 38

	NSEventMaskLeftMouseDown      uint64 = 1 << NSEventTypeLeftMouseDown
	NSEventMaskLeftMouseUp        uint64 = 1 << NSEventTypeLeftMouseUp
	NSEventMaskRightMouseDown     uint64 = 1 << NSEventTypeRightMouseDown
	NSEventMaskRightMouseUp       uint64 = 1 << NSEventTypeRightMouseUp
	NSEventMaskMouseMoved         uint64 = 1 << NSEventTypeMouseMoved
	NSEventMaskLeftMouseDragged   uint64 = 1 << NSEventTypeLeftMouseDragged
	NSEventMaskRightMouseDragged  uint64 = 1 << NSEventTypeRightMouseDragged
	NSEventMaskMouseEntered       uint64 = 1 << NSEventTypeMouseEntered
	NSEventMaskMouseExited        uint64 = 1 << NSEventTypeMouseExited
	NSEventMaskKeyDown            uint64 = 1 << NSEventTypeKeyDown
	NSEventMaskKeyUp              uint64 = 1 << NSEventTypeKeyUp
	NSEventMaskFlagsChanged       uint64 = 1 << NSEventTypeFlagsChanged
	NSEventMaskAppKitDefined      uint64 = 1 << NSEventTypeAppKitDefined
	NSEventMaskSystemDefined      uint64 = 1 << NSEventTypeSystemDefined
	NSEventMaskApplicationDefined uint64 = 1 << NSEventTypeApplicationDefined
	NSEventMaskPeriodic           uint64 = 1 << NSEventTypePeriodic
	NSEventMaskCursorUpdate       uint64 = 1 << NSEventTypeCursorUpdate
	NSEventMaskRotate             uint64 = 1 << NSEventTypeRotate
	NSEventMaskBeginGesture       uint64 = 1 << NSEventTypeBeginGesture
	NSEventMaskEndGesture         uint64 = 1 << NSEventTypeEndGesture
	NSEventMaskScrollWheel        uint64 = 1 << NSEventTypeScrollWheel
	NSEventMaskTabletPoint        uint64 = 1 << NSEventTypeTabletPoint
	NSEventMaskTabletProximity    uint64 = 1 << NSEventTypeTabletProximity
	NSEventMaskOtherMouseDown     uint64 = 1 << NSEventTypeOtherMouseDown
	NSEventMaskOtherMouseUp       uint64 = 1 << NSEventTypeOtherMouseUp
	NSEventMaskOtherMouseDragged  uint64 = 1 << NSEventTypeOtherMouseDragged
	NSEventMaskGesture            uint64 = 1 << NSEventTypeGesture
	NSEventMaskMagnify            uint64 = 1 << NSEventTypeMagnify
	NSEventMaskSwipe              uint64 = 1 << NSEventTypeSwipe
	NSEventMaskSmartMagnify       uint64 = 1 << NSEventTypeSmartMagnify
	NSEventMaskQuickLook          uint64 = 1 << NSEventTypeQuickLook
	NSEventMaskPressure           uint64 = 1 << NSEventTypePressure
	NSEventMaskDirectTouch        uint64 = 1 << NSEventTypeDirectTouch
	NSEventMaskChangeMode         uint64 = 1 << NSEventTypeChangeMode

	NSEventMaskAny uint64 = 18446744073709551615
)

type NSEvent struct {
	gen_NSEvent
}

func NSEvent_GlobalMonitorMatchingMask(mask uint64, ch chan NSEvent) {
	monitorCh = ch
	C.monitorGlobal(C.ulonglong(mask))
}

func NSEvent_LocalMonitorMatchingMask(mask uint64, ch chan NSEvent) {
	monitorCh = ch
	C.monitorLocal(C.ulonglong(mask))
}

func (e NSEvent) KeyCode() (int64, error) {
	eventType := e.Type()
	if eventType != NSEventTypeKeyDown && eventType != NSEventTypeKeyUp {
		return 0, errors.New("event does not contain a keycode")
	}
	return e.Get("keyCode").Int(), nil
}
func (e NSEvent) Type() int64 {
	return e.Get("type").Int()
}
func (e NSEvent) Characters() (string, error) {
	eventType := e.Type()
	if eventType != NSEventTypeKeyDown && eventType != NSEventTypeKeyUp {
		return "", errors.New("event does not contain characters")
	}
	return e.gen_NSEvent.Characters().String(), nil
}
