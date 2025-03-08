package networkextension

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// FilterManager manages content filters
type FilterManager struct {
	objc.Object
}

// FilterManagerClass is the class instance for FilterManager
var FilterManagerClass = objc.GetClass("NEFilterManager")

// SharedManager returns the shared filter manager instance
func FilterSharedManager() FilterManager {
	return FilterManager{objc.Call[objc.Object](FilterManagerClass, objc.Sel("sharedManager"))}
}

// LoadFromPreferencesWithCompletionHandler loads the manager's preferences
func (fm FilterManager) LoadFromPreferencesWithCompletionHandler(completionHandler objc.IObject) {
	objc.Call[objc.Void](fm, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}

// SaveToPreferencesWithCompletionHandler saves the manager's preferences
func (fm FilterManager) SaveToPreferencesWithCompletionHandler(completionHandler objc.IObject) {
	objc.Call[objc.Void](fm, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}

// Enabled returns whether the filter configuration is enabled
func (fm FilterManager) Enabled() bool {
	return objc.Call[bool](fm, objc.Sel("enabled"))
}

// SetEnabled sets whether the filter configuration is enabled
func (fm FilterManager) SetEnabled(enabled bool) {
	objc.Call[objc.Void](fm, objc.Sel("setEnabled:"), enabled)
}

// LocalizedDescription returns the localized description of the filter configuration
func (fm FilterManager) LocalizedDescription() foundation.String {
	return objc.Call[foundation.String](fm, objc.Sel("localizedDescription"))
}

// SetLocalizedDescription sets the localized description of the filter configuration
func (fm FilterManager) SetLocalizedDescription(localizedDescription foundation.String) {
	objc.Call[objc.Void](fm, objc.Sel("setLocalizedDescription:"), localizedDescription)
}

// FilterProvider provides content filtering functionality
type FilterProvider struct {
	objc.Object
}

// FilterProviderClass is the class instance for FilterProvider
var FilterProviderClass = objc.GetClass("NEFilterProvider")

// FilterFlow represents a flow of data to be filtered
type FilterFlow struct {
	objc.Object
}

// FilterFlowClass is the class instance for FilterFlow
var FilterFlowClass = objc.GetClass("NEFilterFlow")

// Direction returns the direction of the flow
func (ff FilterFlow) Direction() NEFilterDataDirection {
	return NEFilterDataDirection(objc.Call[int](ff, objc.Sel("direction")))
}

// SourceAppIdentifier returns the source app identifier of the flow
func (ff FilterFlow) SourceAppIdentifier() foundation.String {
	return objc.Call[foundation.String](ff, objc.Sel("sourceAppIdentifier"))
}

// URL returns the URL of the flow
func (ff FilterFlow) URL() foundation.URL {
	return objc.Call[foundation.URL](ff, objc.Sel("URL"))
}

// SocketFlow represents a socket flow to be filtered
type SocketFlow struct {
	FilterFlow
}

// SocketFlowClass is the class instance for SocketFlow
var SocketFlowClass = objc.GetClass("NEFilterSocketFlow")

// RemoteEndpoint returns the remote endpoint of the socket flow
func (sf SocketFlow) RemoteEndpoint() objc.Object {
	return objc.Call[objc.Object](sf, objc.Sel("remoteEndpoint"))
}

// LocalEndpoint returns the local endpoint of the socket flow
func (sf SocketFlow) LocalEndpoint() objc.Object {
	return objc.Call[objc.Object](sf, objc.Sel("localEndpoint"))
}

// SocketProtocol returns the socket protocol of the socket flow
func (sf SocketFlow) SocketProtocol() SocketProtocol {
	return SocketProtocol(objc.Call[int](sf, objc.Sel("socketProtocol")))
}

// SocketProtocol represents a socket protocol
type SocketProtocol int

const (
	SocketProtocolTCP SocketProtocol = 6
	SocketProtocolUDP SocketProtocol = 17
)

// BrowserFlow represents a browser flow to be filtered
type BrowserFlow struct {
	FilterFlow
}

// BrowserFlowClass is the class instance for BrowserFlow
var BrowserFlowClass = objc.GetClass("NEFilterBrowserFlow")

// Request returns the request of the browser flow
func (bf BrowserFlow) Request() foundation.URLRequest {
	return objc.Call[foundation.URLRequest](bf, objc.Sel("request"))
}

// Response returns the response of the browser flow
func (bf BrowserFlow) Response() foundation.URLResponse {
	return objc.Call[foundation.URLResponse](bf, objc.Sel("response"))
}

// FilterVerdict represents a verdict for a filter flow
type FilterVerdict struct {
	objc.Object
}

// FilterVerdictClass is the class instance for FilterVerdict
var FilterVerdictClass = objc.GetClass("NEFilterVerdict")

// AllowVerdict returns an allow verdict
func AllowVerdict() FilterVerdict {
	return FilterVerdict{objc.Call[objc.Object](FilterVerdictClass, objc.Sel("allowVerdict"))}
}

// DropVerdict returns a drop verdict
func DropVerdict() FilterVerdict {
	return FilterVerdict{objc.Call[objc.Object](FilterVerdictClass, objc.Sel("dropVerdict"))}
}

// RedirectVerdict returns a redirect verdict
func RedirectVerdict(url foundation.URL) FilterVerdict {
	return FilterVerdict{objc.Call[objc.Object](FilterVerdictClass, objc.Sel("redirectWithURL:"), url)}
}

// FilterControlProvider provides control for content filtering
type FilterControlProvider struct {
	FilterProvider
}

// FilterControlProviderClass is the class instance for FilterControlProvider
var FilterControlProviderClass = objc.GetClass("NEFilterControlProvider")

// NewFilterControlProvider creates a new filter control provider
func NewFilterControlProvider() FilterControlProvider {
	alloc := objc.Call[objc.Object](FilterControlProviderClass, objc.Sel("alloc"))
	return FilterControlProvider{FilterProvider{objc.Call[objc.Object](alloc, objc.Sel("init"))}}
}

// HandleNewFlow handles a new flow
func (fcp FilterControlProvider) HandleNewFlow(flow FilterFlow, completionHandler objc.IObject) {
	objc.Call[objc.Void](fcp, objc.Sel("handleNewFlow:completionHandler:"), flow, completionHandler)
}

// HandleRemediationForFlow handles remediation for a flow
func (fcp FilterControlProvider) HandleRemediationForFlow(flow FilterFlow, completionHandler objc.IObject) {
	objc.Call[objc.Void](fcp, objc.Sel("handleRemediationForFlow:completionHandler:"), flow, completionHandler)
}

// FilterDataProvider provides data for content filtering
type FilterDataProvider struct {
	FilterProvider
}

// FilterDataProviderClass is the class instance for FilterDataProvider
var FilterDataProviderClass = objc.GetClass("NEFilterDataProvider")

// NewFilterDataProvider creates a new filter data provider
func NewFilterDataProvider() FilterDataProvider {
	alloc := objc.Call[objc.Object](FilterDataProviderClass, objc.Sel("alloc"))
	return FilterDataProvider{FilterProvider{objc.Call[objc.Object](alloc, objc.Sel("init"))}}
}

// HandleNewFlow handles a new flow
func (fdp FilterDataProvider) HandleNewFlow(flow FilterFlow, completionHandler objc.IObject) {
	objc.Call[objc.Void](fdp, objc.Sel("handleNewFlow:completionHandler:"), flow, completionHandler)
}

// ApplyRemediationForFlow applies remediation for a flow
func (fdp FilterDataProvider) ApplyRemediationForFlow(flow FilterFlow, completionHandler objc.IObject) {
	objc.Call[objc.Void](fdp, objc.Sel("applyRemediationForFlow:completionHandler:"), flow, completionHandler)
}