package corebluetooth

import (
	"testing"
)

func TestCentralManagerClassExists(t *testing.T) {
	if CentralManagerClass.Ptr() == nil {
		t.Error("CentralManagerClass is nil")
	}
}

func TestPeripheralClassExists(t *testing.T) {
	if PeripheralClass.Ptr() == nil {
		t.Error("PeripheralClass is nil")
	}
}

func TestServiceClassExists(t *testing.T) {
	if ServiceClass.Ptr() == nil {
		t.Error("ServiceClass is nil")
	}
}

func TestCharacteristicClassExists(t *testing.T) {
	if CharacteristicClass.Ptr() == nil {
		t.Error("CharacteristicClass is nil")
	}
}

func TestMutableServiceClassExists(t *testing.T) {
	if MutableServiceClass.Ptr() == nil {
		t.Error("MutableServiceClass is nil")
	}
}

func TestMutableCharacteristicClassExists(t *testing.T) {
	if MutableCharacteristicClass.Ptr() == nil {
		t.Error("MutableCharacteristicClass is nil")
	}
}

func TestCBUUIDClassExists(t *testing.T) {
	if CBUUIDClass.Ptr() == nil {
		t.Error("CBUUIDClass is nil")
	}
}

func TestStateConstants(t *testing.T) {
	// Just ensure they are defined without errors
	_ = StateUnknown
	_ = StateResetting
	_ = StateUnsupported
	_ = StateUnauthorized
	_ = StatePoweredOff
	_ = StatePoweredOn
}

func TestPeripheralStateConstants(t *testing.T) {
	// Just ensure they are defined without errors
	_ = PeripheralStateDisconnected
	_ = PeripheralStateConnecting
	_ = PeripheralStateConnected
	_ = PeripheralStateDisconnecting
}

func TestCharacteristicPropertiesConstants(t *testing.T) {
	// Just ensure they are defined without errors
	_ = CharacteristicPropertyBroadcast
	_ = CharacteristicPropertyRead
	_ = CharacteristicPropertyWriteWithoutResponse
	_ = CharacteristicPropertyWrite
	_ = CharacteristicPropertyNotify
	_ = CharacteristicPropertyIndicate
	_ = CharacteristicPropertyAuthenticatedSignedWrites
	_ = CharacteristicPropertyExtendedProperties
	_ = CharacteristicPropertyNotifyEncryptionRequired
	_ = CharacteristicPropertyIndicateEncryptionRequired
}

func TestAttributePermissionsConstants(t *testing.T) {
	// Just ensure they are defined without errors
	_ = AttributePermissionsReadable
	_ = AttributePermissionsWriteable
	_ = AttributePermissionsReadEncryptionRequired
	_ = AttributePermissionsWriteEncryptionRequired
}

func TestWriteTypeConstants(t *testing.T) {
	// Just ensure they are defined without errors
	_ = WriteWithResponse
	_ = WriteWithoutResponse
}