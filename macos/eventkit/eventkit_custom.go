package eventkit

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// EventStore provides an interface for accessing and manipulating calendar events and reminders
type EventStore struct {
	objc.Object
}

// EventStoreClass is the class instance for EventStore
var EventStoreClass = objc.GetClass("EKEventStore")

// NewEventStore creates a new EKEventStore instance
func NewEventStore() EventStore {
	return EventStore{objc.Call[objc.Object](EventStoreClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// RequestAccessToEntityType requests access to either calendar events or reminders
func (e EventStore) RequestAccessToEntityType(entityType int, completion foundation.CompletionHandler) {
	e.Send(objc.Sel("requestAccessToEntityType:completion:"), entityType, completion)
}

// SaveEvent saves the specified event to the event store
func (e EventStore) SaveEvent(event Event, span int, error *foundation.Error) bool {
	return objc.Call[bool](e, objc.Sel("saveEvent:span:error:"), event, span, error)
}

// CalendarsForEntityType returns an array of calendars for the specified entity type
func (e EventStore) CalendarsForEntityType(entityType int) foundation.Array {
	return objc.Call[foundation.Array](e, objc.Sel("calendarsForEntityType:"), entityType)
}

// Event represents a calendar event
type Event struct {
	objc.Object
}

// EventClass is the class instance for Event
var EventClass = objc.GetClass("EKEvent")

// NewEventWithEventStore creates a new EKEvent instance with the specified event store
func NewEventWithEventStore(eventStore EventStore) Event {
	return Event{objc.Call[objc.Object](EventClass, objc.Sel("eventWithEventStore:"), eventStore)}
}

// SetTitle sets the title of the event
func (e Event) SetTitle(title foundation.String) {
	e.Send(objc.Sel("setTitle:"), title)
}

// SetNotes sets the notes of the event
func (e Event) SetNotes(notes foundation.String) {
	e.Send(objc.Sel("setNotes:"), notes)
}

// SetStartDate sets the start date of the event
func (e Event) SetStartDate(startDate foundation.Date) {
	e.Send(objc.Sel("setStartDate:"), startDate)
}

// SetEndDate sets the end date of the event
func (e Event) SetEndDate(endDate foundation.Date) {
	e.Send(objc.Sel("setEndDate:"), endDate)
}

// SetCalendar sets the calendar of the event
func (e Event) SetCalendar(calendar objc.Object) {
	e.Send(objc.Sel("setCalendar:"), calendar)
}

// Calendar represents a calendar in the event store
type Calendar struct {
	objc.Object
}

// CalendarClass is the class instance for Calendar
var CalendarClass = objc.GetClass("EKCalendar")

// NewCalendarForEntityType creates a new calendar for the specified entity type
func NewCalendarForEntityType(entityType int, eventStore EventStore) Calendar {
	return Calendar{objc.Call[objc.Object](CalendarClass, objc.Sel("calendarForEntityType:eventStore:"), entityType, eventStore)}
}

// EntityTypeEvent represents the calendar events entity type
const EntityTypeEvent = 0

// EntityTypeReminder represents the reminders entity type
const EntityTypeReminder = 1

// SpanThisEvent specifies that a change applies only to this event
const SpanThisEvent = 0

// SpanFutureEvents specifies that a change applies to this and future events
const SpanFutureEvents = 1

// SpanAllEvents specifies that a change applies to all events
const SpanAllEvents = 2
