package sysconfig

// The reference to an object that represents a network protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkprotocolref?language=objc
type NetworkProtocolRef uintptr

// The handle to an open preferences session for accessing system configuration preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scpreferencesref?language=objc
type PreferencesRef uintptr

// The handle to manage a connection-oriented service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectionref?language=objc
type NetworkConnectionRef uintptr

// The reference to an object that represents a network service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkserviceref?language=objc
type NetworkServiceRef uintptr

// Structure containing user-specified data and callbacks for a dynamic store session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scdynamicstorecontext?language=objc
type DynamicStoreContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// The handle to a network address or name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkreachabilityref?language=objc
type NetworkReachabilityRef uintptr

// The reference to an object that represents the status of an Ethernet bond interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scbondstatusref?language=objc
type BondStatusRef uintptr

// A structure containing user-specified data and callbacks for accessing system configuration preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scpreferencescontext?language=objc
type PreferencesContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// The reference to an object that represents a network interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkinterfaceref?language=objc
type NetworkInterfaceRef uintptr

// The handle to an open dynamic store session with the system configuration daemon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scdynamicstoreref?language=objc
type DynamicStoreRef uintptr

// The reference to an object that represents a network set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworksetref?language=objc
type NetworkSetRef uintptr

// Structure containing user-specified data and callbacks used with [systemconfiguration/scnetworkreachabilitysetcallback]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkreachabilitycontext?language=objc
type NetworkReachabilityContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}

// A structure containing user-specified data and callbacks for a network connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectioncontext?language=objc
type NetworkConnectionContext struct {
	Version         int64
	Info            uintptr
	Retain          uintptr
	Release         uintptr
	CopyDescription uintptr
}
