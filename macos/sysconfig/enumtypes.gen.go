// AUTO-GENERATED CODE, DO NOT MODIFY

package sysconfig

// Flags that indicate whether the specified network node name or address is reachable, whether a connection is required, and whether some user intervention may be required when establishing a connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectionflags?language=objc
type EtworkConnectionFlags uint32

// The PPP-specific status of the network connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectionpppstatus?language=objc
type EtworkConnectionPPPStatus int32

const (
	KNetworkConnectionPPPAuthenticating     EtworkConnectionPPPStatus = 5
	KNetworkConnectionPPPConnected          EtworkConnectionPPPStatus = 8
	KNetworkConnectionPPPConnectingLink     EtworkConnectionPPPStatus = 2
	KNetworkConnectionPPPDialOnTraffic      EtworkConnectionPPPStatus = 3
	KNetworkConnectionPPPDisconnected       EtworkConnectionPPPStatus = 0
	KNetworkConnectionPPPDisconnectingLink  EtworkConnectionPPPStatus = 10
	KNetworkConnectionPPPHoldingLinkOff     EtworkConnectionPPPStatus = 11
	KNetworkConnectionPPPInitializing       EtworkConnectionPPPStatus = 1
	KNetworkConnectionPPPNegotiatingLink    EtworkConnectionPPPStatus = 4
	KNetworkConnectionPPPNegotiatingNetwork EtworkConnectionPPPStatus = 7
	KNetworkConnectionPPPSuspended          EtworkConnectionPPPStatus = 12
	KNetworkConnectionPPPTerminating        EtworkConnectionPPPStatus = 9
	KNetworkConnectionPPPWaitingForCallBack EtworkConnectionPPPStatus = 6
	KNetworkConnectionPPPWaitingForRedial   EtworkConnectionPPPStatus = 13
)

// The current status of the network connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectionstatus?language=objc
type EtworkConnectionStatus int32

const (
	KNetworkConnectionConnected     EtworkConnectionStatus = 2
	KNetworkConnectionConnecting    EtworkConnectionStatus = 1
	KNetworkConnectionDisconnected  EtworkConnectionStatus = 0
	KNetworkConnectionDisconnecting EtworkConnectionStatus = 3
	KNetworkConnectionInvalid       EtworkConnectionStatus = -1
)

// Flags that indicate the reachability of a network node name or address, including whether a connection is required, and whether some user intervention might be required when establishing a connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkreachabilityflags?language=objc
type EtworkReachabilityFlags uint32

const (
	KNetworkReachabilityFlagsConnectionAutomatic  EtworkReachabilityFlags = 8
	KNetworkReachabilityFlagsConnectionOnDemand   EtworkReachabilityFlags = 32
	KNetworkReachabilityFlagsConnectionOnTraffic  EtworkReachabilityFlags = 8
	KNetworkReachabilityFlagsConnectionRequired   EtworkReachabilityFlags = 4
	KNetworkReachabilityFlagsInterventionRequired EtworkReachabilityFlags = 16
	KNetworkReachabilityFlagsIsDirect             EtworkReachabilityFlags = 131072
	KNetworkReachabilityFlagsIsLocalAddress       EtworkReachabilityFlags = 65536
	KNetworkReachabilityFlagsReachable            EtworkReachabilityFlags = 2
	KNetworkReachabilityFlagsTransientConnection  EtworkReachabilityFlags = 1
)

// The type of notification (used with the SCPreferencesCallBack callback). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scpreferencesnotification?language=objc
type PreferencesNotification uint32

const (
	KPreferencesNotificationApply  PreferencesNotification = 2
	KPreferencesNotificationCommit PreferencesNotification = 1
)
