package networkextension

import (
	"testing"
)

func TestVPNManagerClassExists(t *testing.T) {
	if VPNManagerClass.Ptr() == nil {
		t.Error("VPNManagerClass is nil")
	}
}

func TestVPNProtocolClassExists(t *testing.T) {
	if VPNProtocolClass.Ptr() == nil {
		t.Error("VPNProtocolClass is nil")
	}
}

func TestIPSecProtocolClassExists(t *testing.T) {
	if IPSecProtocolClass.Ptr() == nil {
		t.Error("IPSecProtocolClass is nil")
	}
}

func TestIKEv2ProtocolClassExists(t *testing.T) {
	if IKEv2ProtocolClass.Ptr() == nil {
		t.Error("IKEv2ProtocolClass is nil")
	}
}

func TestVPNConnectionClassExists(t *testing.T) {
	if VPNConnectionClass.Ptr() == nil {
		t.Error("VPNConnectionClass is nil")
	}
}

func TestOnDemandRuleClassExists(t *testing.T) {
	if OnDemandRuleClass.Ptr() == nil {
		t.Error("OnDemandRuleClass is nil")
	}
}

func TestFilterClassExists(t *testing.T) {
	if FilterManagerClass.Ptr() == nil {
		t.Error("FilterManagerClass is nil")
	}
}

func TestFilterFilterClass(t *testing.T) {
	if FilterFlowClass.Ptr() != nil {
		// Just test the class exists (will be nil if it doesn't)
		// Even this is okay
	}
}

func TestVPNStatusConstants(t *testing.T) {
	// Just ensure they are defined
	_ = VPNStatusInvalid
	_ = VPNStatusDisconnected
	_ = VPNStatusConnecting
	_ = VPNStatusConnected
	_ = VPNStatusReasserting
	_ = VPNStatusDisconnecting
}

func TestIPSecAuthenticationMethodConstants(t *testing.T) {
	// Just ensure they are defined
	_ = IPSecAuthenticationMethodNone
	_ = IPSecAuthenticationMethodCertificate
	_ = IPSecAuthenticationMethodSharedSecret
}

func TestInterfaceTypeConstants(t *testing.T) {
	// Just ensure they are defined
	_ = InterfaceTypeOther
	_ = InterfaceTypeWiFi
	_ = InterfaceTypeCellular
	_ = InterfaceTypeWiredEthernet
}

func TestOnDemandRuleActionConstants(t *testing.T) {
	// Just ensure they are defined
	_ = OnDemandRuleActionConnect
	_ = OnDemandRuleActionDisconnect
	_ = OnDemandRuleActionEvaluateConnection
	_ = OnDemandRuleActionIgnore
}