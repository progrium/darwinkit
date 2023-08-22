// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
)

// A block the system calls to indicate it has enabled or disabled a profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilechangedblock?language=objc
type CIProfileChangedBlock = func(session CISession, channel ChannelNumber, profile CIProfile, enabled bool)

// A block the system calls when a MIDI-CI node discovery request completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoveryresponseblock?language=objc
type CIDiscoveryResponseBlock = func(discoveredNodes []CIDiscoveredNode)

// A function the system calls after it has completed sending a System Exclusive (SysEx) event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicompletionproc?language=objc
type CompletionProc = func(request *SysexSendRequest)

// A block the system calls when a MIDI-CI session disconnects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisessiondisconnectblock?language=objc
type CISessionDisconnectBlock = func(session CISession, error foundation.Error)

// A block the system calls when a MIDI-CI session or responder receives profile-specific data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilespecificdatablock?language=objc
type CIProfileSpecificDataBlock = func(session CISession, channel ChannelNumber, profile CIProfile, profileSpecificData []byte)

// A block receiving MIDI input that includes the incoming messages and a refCon to identify the source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midireceiveblock?language=objc
type ReceiveBlock = func(evtlist *EventList, srcConnRefCon unsafe.Pointer)

// A block receiving MIDI input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midireadblock?language=objc
type ReadBlock = func(pktlist *PacketList, srcConnRefCon unsafe.Pointer)

// A function receiving MIDI input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midireadproc?language=objc
type ReadProc = func(pktlist *PacketList, readProcRefCon unsafe.Pointer, srcConnRefCon unsafe.Pointer)

// A callback block for notifying clients of state changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinotifyblock?language=objc
type NotifyBlock = func(message *Notification)

// A callback function for notifying clients of state changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinotifyproc?language=objc
type NotifyProc = func(message *Notification, refCon unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midieventvisitor?language=objc
type EventVisitor = func(context unsafe.Pointer, timeStamp TimeStamp, message UniversalMessage)
