// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SocketPort] class.
var SocketPortClass = _SocketPortClass{objc.GetClass("NSSocketPort")}

type _SocketPortClass struct {
	objc.Class
}

// An interface definition for the [SocketPort] class.
type ISocketPort interface {
	IPort
	Socket() SocketNativeHandle
	Address() []byte
	ProtocolFamily() int
	SocketType() int
	Protocol() int
}

// A port that represents a BSD socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport?language=objc
type SocketPort struct {
	Port
}

func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		Port: PortFrom(ptr),
	}
}

func (s_ SocketPort) InitWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol int, sock SocketNativeHandle) SocketPort {
	rv := objc.Call[SocketPort](s_, objc.Sel("initWithProtocolFamily:socketType:protocol:socket:"), family, type_, protocol, sock)
	return rv
}

// Initializes the receiver with a previously created local socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399484-initwithprotocolfamily?language=objc
func SocketPort_InitWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol int, sock SocketNativeHandle) SocketPort {
	return SocketPortClass.Alloc().InitWithProtocolFamilySocketTypeProtocolSocket(family, type_, protocol, sock)
}

func (s_ SocketPort) InitRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol int, address []byte) SocketPort {
	rv := objc.Call[SocketPort](s_, objc.Sel("initRemoteWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol, address)
	return rv
}

// Initializes the receiver as a remote socket with the provided arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399535-initremotewithprotocolfamily?language=objc
func SocketPort_InitRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol int, address []byte) SocketPort {
	return SocketPortClass.Alloc().InitRemoteWithProtocolFamilySocketTypeProtocolAddress(family, type_, protocol, address)
}

func (s_ SocketPort) InitWithTCPPort(port int) SocketPort {
	rv := objc.Call[SocketPort](s_, objc.Sel("initWithTCPPort:"), port)
	return rv
}

// Initializes the receiver as a local TCP/IP socket of type SOCK_STREAM, listening on a specified port number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399488-initwithtcpport?language=objc
func SocketPort_InitWithTCPPort(port int) SocketPort {
	return SocketPortClass.Alloc().InitWithTCPPort(port)
}

func (s_ SocketPort) Init() SocketPort {
	rv := objc.Call[SocketPort](s_, objc.Sel("init"))
	return rv
}

func (s_ SocketPort) InitRemoteWithTCPPortHost(port int, hostName string) SocketPort {
	rv := objc.Call[SocketPort](s_, objc.Sel("initRemoteWithTCPPort:host:"), port, hostName)
	return rv
}

// Initializes the receiver as a TCP/IP socket of type SOCK_STREAM that can connect to a remote host on a specified port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399474-initremotewithtcpport?language=objc
func SocketPort_InitRemoteWithTCPPortHost(port int, hostName string) SocketPort {
	return SocketPortClass.Alloc().InitRemoteWithTCPPortHost(port, hostName)
}

func (sc _SocketPortClass) Alloc() SocketPort {
	rv := objc.Call[SocketPort](sc, objc.Sel("alloc"))
	return rv
}

func SocketPort_Alloc() SocketPort {
	return SocketPortClass.Alloc()
}

func (sc _SocketPortClass) New() SocketPort {
	rv := objc.Call[SocketPort](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSocketPort() SocketPort {
	return SocketPortClass.New()
}

// The receiver’s native socket identifier on the platform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399492-socket?language=objc
func (s_ SocketPort) Socket() SocketNativeHandle {
	rv := objc.Call[SocketNativeHandle](s_, objc.Sel("socket"))
	return rv
}

// The receiver’s socket address structure stored inside an NSData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399480-address?language=objc
func (s_ SocketPort) Address() []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("address"))
	return rv
}

// The protocol family that the receiver uses for communication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399543-protocolfamily?language=objc
func (s_ SocketPort) ProtocolFamily() int {
	rv := objc.Call[int](s_, objc.Sel("protocolFamily"))
	return rv
}

// The receiver’s socket type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399565-sockettype?language=objc
func (s_ SocketPort) SocketType() int {
	rv := objc.Call[int](s_, objc.Sel("socketType"))
	return rv
}

// The protocol that the receiver uses for communication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketport/1399557-protocol?language=objc
func (s_ SocketPort) Protocol() int {
	rv := objc.Call[int](s_, objc.Sel("protocol"))
	return rv
}
