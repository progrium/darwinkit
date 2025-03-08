package main

import (
	"fmt"
	"time"

	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/eventkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create a window
		frame := foundation.Rect{Size: foundation.Size{600, 400}}
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&window)
		window.SetTitle(foundation.String_StringWithString("EventKit Example"))
		window.Center()

		// Create a text view
		textView := appkit.NewTextView()
		scrollView := appkit.NewScrollView()
		objc.Retain(&textView)
		objc.Retain(&scrollView)

		scrollView.SetDocumentView(textView)
		window.SetContentView(scrollView)

		// Create event store
		eventStore := eventkit.NewEventStore()
		objc.Retain(&eventStore)

		// Show the window
		window.MakeKeyAndOrderFront(window)

		// Request calendar access
		completion := foundation.NewBlockWithVoidBoolError(func(granted bool, err foundation.Error) {
			if !granted {
				textView.SetString(foundation.String_StringWithString("Calendar access not granted"))
				return
			}

			textView.SetString(foundation.String_StringWithString("Calendar access granted!\n\n"))

			// Create a new event
			event := eventkit.NewEventWithEventStore(eventStore)
			objc.Retain(&event)

			// Set event properties
			event.SetTitle(foundation.String_StringWithString("Test Event from Go"))
			event.SetNotes(foundation.String_StringWithString("Created using DarwinKit"))

			// Set event times
			now := time.Now()
			startDate := foundation.Date_DateWithTimeIntervalSinceNow(float64(60 * 60)) // 1 hour from now
			endDate := foundation.Date_DateWithTimeIntervalSinceNow(float64(2 * 60 * 60)) // 2 hours from now
			event.SetStartDate(startDate)
			event.SetEndDate(endDate)

			// Get default calendar
			calendars := eventStore.CalendarsForEntityType(eventkit.EntityTypeEvent)
			if calendars.Count() == 0 {
				textView.SetString(foundation.String_StringWithString("No calendars found"))
				return
			}

			calendar := calendars.ObjectAtIndex(0)
			event.SetCalendar(calendar)

			// Save the event
			var error foundation.Error
			success := eventStore.SaveEvent(event, eventkit.SpanThisEvent, &error)

			if success {
				message := fmt.Sprintf("Event created successfully!\n\nTitle: Test Event from Go\nStart: %s\nEnd: %s", 
					now.Add(time.Hour).Format(time.RFC1123),
					now.Add(2*time.Hour).Format(time.RFC1123))
				textView.SetString(foundation.String_StringWithString(message))
			} else {
				errorMessage := fmt.Sprintf("Failed to create event: %s", error.LocalizedDescription().UTF8String())
				textView.SetString(foundation.String_StringWithString(errorMessage))
			}
		})

		eventStore.RequestAccessToEntityType(eventkit.EntityTypeEvent, completion)

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}