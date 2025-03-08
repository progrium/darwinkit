package networkextension

import (
	"fmt"

	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// VPNManager manages VPN configurations and connections
type VPNManager struct {
	objc.Object
}

// VPNManagerClass is the class instance for VPNManager
var VPNManagerClass = objc.GetClass("NEVPNManager")

// SharedManager returns the shared VPN manager instance
func SharedManager() VPNManager {
	return VPNManager{objc.Call[objc.Object](VPNManagerClass, objc.Sel("sharedManager"))}
}

// LoadFromPreferencesWithCompletionHandler loads the manager's preferences
func (vm VPNManager) LoadFromPreferencesWithCompletionHandler(completionHandler objc.IObject) {
	objc.Call[objc.Void](vm, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}

// SaveToPreferencesWithCompletionHandler saves the manager's preferences
func (vm VPNManager) SaveToPreferencesWithCompletionHandler(completionHandler objc.IObject) {
	objc.Call[objc.Void](vm, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}

// Protocol returns the VPN protocol configuration
func (vm VPNManager) Protocol() VPNProtocol {
	return VPNProtocol{objc.Call[objc.Object](vm, objc.Sel("protocol"))}
}

// SetProtocol sets the VPN protocol configuration
func (vm VPNManager) SetProtocol(protocol VPNProtocol) {
	objc.Call[objc.Void](vm, objc.Sel("setProtocol:"), protocol)
}

// Connection returns the VPN connection
func (vm VPNManager) Connection() VPNConnection {
	return VPNConnection{objc.Call[objc.Object](vm, objc.Sel("connection"))}
}

// Enabled returns whether the VPN configuration is enabled
func (vm VPNManager) Enabled() bool {
	return objc.Call[bool](vm, objc.Sel("enabled"))
}

// SetEnabled sets whether the VPN configuration is enabled
func (vm VPNManager) SetEnabled(enabled bool) {
	objc.Call[objc.Void](vm, objc.Sel("setEnabled:"), enabled)
}

// OnDemandEnabled returns whether VPN on demand is enabled
func (vm VPNManager) OnDemandEnabled() bool {
	return objc.Call[bool](vm, objc.Sel("onDemandEnabled"))
}

// SetOnDemandEnabled sets whether VPN on demand is enabled
func (vm VPNManager) SetOnDemandEnabled(onDemandEnabled bool) {
	objc.Call[objc.Void](vm, objc.Sel("setOnDemandEnabled:"), onDemandEnabled)
}

// OnDemandRules returns the VPN on demand rules
func (vm VPNManager) OnDemandRules() foundation.Array {
	return objc.Call[foundation.Array](vm, objc.Sel("onDemandRules"))
}

// SetOnDemandRules sets the VPN on demand rules
func (vm VPNManager) SetOnDemandRules(onDemandRules foundation.Array) {
	objc.Call[objc.Void](vm, objc.Sel("setOnDemandRules:"), onDemandRules)
}

// LocalizedDescription returns the localized description of the VPN configuration
func (vm VPNManager) LocalizedDescription() foundation.String {
	return objc.Call[foundation.String](vm, objc.Sel("localizedDescription"))
}

// SetLocalizedDescription sets the localized description of the VPN configuration
func (vm VPNManager) SetLocalizedDescription(localizedDescription foundation.String) {
	objc.Call[objc.Void](vm, objc.Sel("setLocalizedDescription:"), localizedDescription)
}

// VPNProtocol is a base protocol for VPN protocols
type VPNProtocol struct {
	objc.Object
}

// VPNProtocolClass is the class instance for VPNProtocol
var VPNProtocolClass = objc.GetClass("NEVPNProtocol")

// ServerAddress returns the VPN server address
func (vp VPNProtocol) ServerAddress() foundation.String {
	return objc.Call[foundation.String](vp, objc.Sel("serverAddress"))
}

// SetServerAddress sets the VPN server address
func (vp VPNProtocol) SetServerAddress(serverAddress foundation.String) {
	objc.Call[objc.Void](vp, objc.Sel("setServerAddress:"), serverAddress)
}

// Username returns the VPN username
func (vp VPNProtocol) Username() foundation.String {
	return objc.Call[foundation.String](vp, objc.Sel("username"))
}

// SetUsername sets the VPN username
func (vp VPNProtocol) SetUsername(username foundation.String) {
	objc.Call[objc.Void](vp, objc.Sel("setUsername:"), username)
}

// PasswordReference returns the VPN password keychain reference
func (vp VPNProtocol) PasswordReference() foundation.Data {
	return objc.Call[foundation.Data](vp, objc.Sel("passwordReference"))
}

// SetPasswordReference sets the VPN password keychain reference
func (vp VPNProtocol) SetPasswordReference(passwordReference foundation.Data) {
	objc.Call[objc.Void](vp, objc.Sel("setPasswordReference:"), passwordReference)
}

// DisconnectOnSleep returns whether VPN disconnects on sleep
func (vp VPNProtocol) DisconnectOnSleep() bool {
	return objc.Call[bool](vp, objc.Sel("disconnectOnSleep"))
}

// SetDisconnectOnSleep sets whether VPN disconnects on sleep
func (vp VPNProtocol) SetDisconnectOnSleep(disconnectOnSleep bool) {
	objc.Call[objc.Void](vp, objc.Sel("setDisconnectOnSleep:"), disconnectOnSleep)
}

// IPSecProtocol represents an IPSec VPN protocol configuration
type IPSecProtocol struct {
	VPNProtocol
}

// IPSecProtocolClass is the class instance for IPSecProtocol
var IPSecProtocolClass = objc.GetClass("NEVPNProtocolIPSec")

// NewIPSecProtocol creates a new IPSec protocol configuration
func NewIPSecProtocol() IPSecProtocol {
	alloc := objc.Call[objc.Object](IPSecProtocolClass, objc.Sel("alloc"))
	return IPSecProtocol{VPNProtocol{objc.Call[objc.Object](alloc, objc.Sel("init"))}}
}

// AuthenticationMethod returns the IPSec authentication method
func (ip IPSecProtocol) AuthenticationMethod() IPSecAuthenticationMethod {
	return IPSecAuthenticationMethod(objc.Call[int](ip, objc.Sel("authenticationMethod")))
}

// SetAuthenticationMethod sets the IPSec authentication method
func (ip IPSecProtocol) SetAuthenticationMethod(authenticationMethod IPSecAuthenticationMethod) {
	objc.Call[objc.Void](ip, objc.Sel("setAuthenticationMethod:"), authenticationMethod)
}

// UseExtendedAuthentication returns whether IPSec uses extended authentication
func (ip IPSecProtocol) UseExtendedAuthentication() bool {
	return objc.Call[bool](ip, objc.Sel("useExtendedAuthentication"))
}

// SetUseExtendedAuthentication sets whether IPSec uses extended authentication
func (ip IPSecProtocol) SetUseExtendedAuthentication(useExtendedAuthentication bool) {
	objc.Call[objc.Void](ip, objc.Sel("setUseExtendedAuthentication:"), useExtendedAuthentication)
}

// SharedSecretReference returns the IPSec shared secret keychain reference
func (ip IPSecProtocol) SharedSecretReference() foundation.Data {
	return objc.Call[foundation.Data](ip, objc.Sel("sharedSecretReference"))
}

// SetSharedSecretReference sets the IPSec shared secret keychain reference
func (ip IPSecProtocol) SetSharedSecretReference(sharedSecretReference foundation.Data) {
	objc.Call[objc.Void](ip, objc.Sel("setSharedSecretReference:"), sharedSecretReference)
}

// LocalIdentifier returns the IPSec local identifier
func (ip IPSecProtocol) LocalIdentifier() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("localIdentifier"))
}

// SetLocalIdentifier sets the IPSec local identifier
func (ip IPSecProtocol) SetLocalIdentifier(localIdentifier foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setLocalIdentifier:"), localIdentifier)
}

// RemoteIdentifier returns the IPSec remote identifier
func (ip IPSecProtocol) RemoteIdentifier() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("remoteIdentifier"))
}

// SetRemoteIdentifier sets the IPSec remote identifier
func (ip IPSecProtocol) SetRemoteIdentifier(remoteIdentifier foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setRemoteIdentifier:"), remoteIdentifier)
}

// IKEv2Protocol represents an IKEv2 VPN protocol configuration
type IKEv2Protocol struct {
	VPNProtocol
}

// IKEv2ProtocolClass is the class instance for IKEv2Protocol
var IKEv2ProtocolClass = objc.GetClass("NEVPNProtocolIKEv2")

// NewIKEv2Protocol creates a new IKEv2 protocol configuration
func NewIKEv2Protocol() IKEv2Protocol {
	alloc := objc.Call[objc.Object](IKEv2ProtocolClass, objc.Sel("alloc"))
	return IKEv2Protocol{VPNProtocol{objc.Call[objc.Object](alloc, objc.Sel("init"))}}
}

// RemoteIdentifier returns the IKEv2 remote identifier
func (ip IKEv2Protocol) RemoteIdentifier() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("remoteIdentifier"))
}

// SetRemoteIdentifier sets the IKEv2 remote identifier
func (ip IKEv2Protocol) SetRemoteIdentifier(remoteIdentifier foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setRemoteIdentifier:"), remoteIdentifier)
}

// LocalIdentifier returns the IKEv2 local identifier
func (ip IKEv2Protocol) LocalIdentifier() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("localIdentifier"))
}

// SetLocalIdentifier sets the IKEv2 local identifier
func (ip IKEv2Protocol) SetLocalIdentifier(localIdentifier foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setLocalIdentifier:"), localIdentifier)
}

// ServerCertificateIssuerCommonName returns the IKEv2 server certificate issuer common name
func (ip IKEv2Protocol) ServerCertificateIssuerCommonName() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("serverCertificateIssuerCommonName"))
}

// SetServerCertificateIssuerCommonName sets the IKEv2 server certificate issuer common name
func (ip IKEv2Protocol) SetServerCertificateIssuerCommonName(serverCertificateIssuerCommonName foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setServerCertificateIssuerCommonName:"), serverCertificateIssuerCommonName)
}

// ServerCertificateCommonName returns the IKEv2 server certificate common name
func (ip IKEv2Protocol) ServerCertificateCommonName() foundation.String {
	return objc.Call[foundation.String](ip, objc.Sel("serverCertificateCommonName"))
}

// SetServerCertificateCommonName sets the IKEv2 server certificate common name
func (ip IKEv2Protocol) SetServerCertificateCommonName(serverCertificateCommonName foundation.String) {
	objc.Call[objc.Void](ip, objc.Sel("setServerCertificateCommonName:"), serverCertificateCommonName)
}

// VPNConnection represents a VPN connection
type VPNConnection struct {
	objc.Object
}

// VPNConnectionClass is the class instance for VPNConnection
var VPNConnectionClass = objc.GetClass("NEVPNConnection")

// StartVPNTunnelAndReturnError starts the VPN tunnel
func (vc VPNConnection) StartVPNTunnelAndReturnError() (bool, foundation.Error) {
	var error foundation.Error
	success := objc.Call[bool](vc, objc.Sel("startVPNTunnelAndReturnError:"), &error)
	return success, error
}

// StartVPNTunnelWithOptionsAndReturnError starts the VPN tunnel with options
func (vc VPNConnection) StartVPNTunnelWithOptionsAndReturnError(options foundation.Dictionary) (bool, foundation.Error) {
	var error foundation.Error
	success := objc.Call[bool](vc, objc.Sel("startVPNTunnelWithOptions:andReturnError:"), options, &error)
	return success, error
}

// StopVPNTunnel stops the VPN tunnel
func (vc VPNConnection) StopVPNTunnel() {
	objc.Call[objc.Void](vc, objc.Sel("stopVPNTunnel"))
}

// Status returns the VPN connection status
func (vc VPNConnection) Status() VPNStatus {
	return VPNStatus(objc.Call[int](vc, objc.Sel("status")))
}

// OnDemandRule represents a VPN on-demand rule
type OnDemandRule struct {
	objc.Object
}

// OnDemandRuleClass is the class instance for OnDemandRule
var OnDemandRuleClass = objc.GetClass("NEOnDemandRule")

// InterfaceTypeMatch returns the interface type match
func (odr OnDemandRule) InterfaceTypeMatch() InterfaceType {
	return InterfaceType(objc.Call[int](odr, objc.Sel("interfaceTypeMatch")))
}

// SetInterfaceTypeMatch sets the interface type match
func (odr OnDemandRule) SetInterfaceTypeMatch(interfaceTypeMatch InterfaceType) {
	objc.Call[objc.Void](odr, objc.Sel("setInterfaceTypeMatch:"), interfaceTypeMatch)
}

// Action returns the rule action
func (odr OnDemandRule) Action() OnDemandRuleAction {
	return OnDemandRuleAction(objc.Call[int](odr, objc.Sel("action")))
}

// SetAction sets the rule action
func (odr OnDemandRule) SetAction(action OnDemandRuleAction) {
	objc.Call[objc.Void](odr, objc.Sel("setAction:"), action)
}

// DNSSearchDomainMatch returns the DNS search domain match
func (odr OnDemandRule) DNSSearchDomainMatch() foundation.Array {
	return objc.Call[foundation.Array](odr, objc.Sel("DNSSearchDomainMatch"))
}

// SetDNSSearchDomainMatch sets the DNS search domain match
func (odr OnDemandRule) SetDNSSearchDomainMatch(DNSSearchDomainMatch foundation.Array) {
	objc.Call[objc.Void](odr, objc.Sel("setDNSSearchDomainMatch:"), DNSSearchDomainMatch)
}

// DNSServerAddressMatch returns the DNS server address match
func (odr OnDemandRule) DNSServerAddressMatch() foundation.Array {
	return objc.Call[foundation.Array](odr, objc.Sel("DNSServerAddressMatch"))
}

// SetDNSServerAddressMatch sets the DNS server address match
func (odr OnDemandRule) SetDNSServerAddressMatch(DNSServerAddressMatch foundation.Array) {
	objc.Call[objc.Void](odr, objc.Sel("setDNSServerAddressMatch:"), DNSServerAddressMatch)
}

// VPNStatus represents the status of a VPN connection
type VPNStatus int

const (
	VPNStatusInvalid      VPNStatus = 0
	VPNStatusDisconnected VPNStatus = 1
	VPNStatusConnecting   VPNStatus = 2
	VPNStatusConnected    VPNStatus = 3
	VPNStatusReasserting  VPNStatus = 4
	VPNStatusDisconnecting VPNStatus = 5
)

// IPSecAuthenticationMethod represents the authentication method for IPSec
type IPSecAuthenticationMethod int

const (
	IPSecAuthenticationMethodNone      IPSecAuthenticationMethod = 0
	IPSecAuthenticationMethodCertificate IPSecAuthenticationMethod = 1
	IPSecAuthenticationMethodSharedSecret IPSecAuthenticationMethod = 2
)

// InterfaceType represents the type of network interface
type InterfaceType int

const (
	InterfaceTypeOther       InterfaceType = 0
	InterfaceTypeWiFi        InterfaceType = 1
	InterfaceTypeCellular    InterfaceType = 2
	InterfaceTypeWiredEthernet InterfaceType = 3
)

// OnDemandRuleAction represents the action for a VPN on-demand rule
type OnDemandRuleAction int

const (
	OnDemandRuleActionConnect    OnDemandRuleAction = 1
	OnDemandRuleActionDisconnect OnDemandRuleAction = 2
	OnDemandRuleActionEvaluateConnection OnDemandRuleAction = 3
	OnDemandRuleActionIgnore     OnDemandRuleAction = 4
)

// ContentFilter provides content filtering functionality
type ContentFilter struct {
	objc.Object
}

// ContentFilterClass is the class instance for ContentFilter
var ContentFilterClass = objc.GetClass("NEFilterFlow")

// CreateVPNPasswordWithUsername creates a keychain item for VPN credentials
// This functionality is no longer implemented in this file due to issues with Dictionary methods.
// If needed, it should be reimplemented using the low-level Security framework C API directly.
func CreateVPNPasswordWithUsername(username string, password string, service string) (foundation.Data, error) {
	return foundation.Data{}, fmt.Errorf("not implemented - use Security framework C API directly")
}