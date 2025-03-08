package corebluetooth

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// CentralManager manages discovered or connected remote peripheral devices
type CentralManager struct {
	objc.Object
}

// CentralManagerClass is the class instance for CentralManager
var CentralManagerClass = objc.GetClass("CBCentralManager")

// NewCentralManager creates a new central manager with the specified delegate and dispatch queue
func NewCentralManager(delegate objc.Object, queue objc.Object) CentralManager {
	alloc := objc.Call[objc.Object](CentralManagerClass, objc.Sel("alloc"))
	return CentralManager{objc.Call[objc.Object](alloc, objc.Sel("initWithDelegate:queue:"), delegate, queue)}
}

// State returns the current state of the central manager
func (cm CentralManager) State() State {
	return State(objc.Call[int](cm, objc.Sel("state")))
}

// ScanForPeripheralsWithServices starts scanning for peripherals advertising the specified services
func (cm CentralManager) ScanForPeripheralsWithServices(serviceUUIDs foundation.Array, options foundation.Dictionary) {
	objc.Call[objc.Void](cm, objc.Sel("scanForPeripheralsWithServices:options:"), serviceUUIDs, options)
}

// StopScan stops scanning for peripherals
func (cm CentralManager) StopScan() {
	objc.Call[objc.Void](cm, objc.Sel("stopScan"))
}

// ConnectPeripheral connects to a peripheral
func (cm CentralManager) ConnectPeripheral(peripheral Peripheral, options foundation.Dictionary) {
	objc.Call[objc.Void](cm, objc.Sel("connectPeripheral:options:"), peripheral, options)
}

// CancelPeripheralConnection cancels a peripheral connection
func (cm CentralManager) CancelPeripheralConnection(peripheral Peripheral) {
	objc.Call[objc.Void](cm, objc.Sel("cancelPeripheralConnection:"), peripheral)
}

// RetrievePeripheralsWithIdentifiers retrieves peripherals by their identifiers
func (cm CentralManager) RetrievePeripheralsWithIdentifiers(identifiers foundation.Array) foundation.Array {
	return objc.Call[foundation.Array](cm, objc.Sel("retrievePeripheralsWithIdentifiers:"), identifiers)
}

// Peripheral represents a remote peripheral device
type Peripheral struct {
	objc.Object
}

// PeripheralClass is the class instance for Peripheral
var PeripheralClass = objc.GetClass("CBPeripheral")

// Identifier returns the peripheral's identifier
func (p Peripheral) Identifier() foundation.UUID {
	return objc.Call[foundation.UUID](p, objc.Sel("identifier"))
}

// Name returns the peripheral's name
func (p Peripheral) Name() foundation.String {
	return objc.Call[foundation.String](p, objc.Sel("name"))
}

// State returns the peripheral's connection state
func (p Peripheral) State() PeripheralState {
	return PeripheralState(objc.Call[int](p, objc.Sel("state")))
}

// Services returns the peripheral's services
func (p Peripheral) Services() foundation.Array {
	return objc.Call[foundation.Array](p, objc.Sel("services"))
}

// DiscoverServices discovers services on the peripheral
func (p Peripheral) DiscoverServices(serviceUUIDs foundation.Array) {
	objc.Call[objc.Void](p, objc.Sel("discoverServices:"), serviceUUIDs)
}

// DiscoverCharacteristicsForService discovers characteristics for a service
func (p Peripheral) DiscoverCharacteristicsForService(characteristicUUIDs foundation.Array, service Service) {
	objc.Call[objc.Void](p, objc.Sel("discoverCharacteristics:forService:"), characteristicUUIDs, service)
}

// ReadValueForCharacteristic reads the value for a characteristic
func (p Peripheral) ReadValueForCharacteristic(characteristic Characteristic) {
	objc.Call[objc.Void](p, objc.Sel("readValueForCharacteristic:"), characteristic)
}

// WriteValueForCharacteristic writes a value for a characteristic
func (p Peripheral) WriteValueForCharacteristic(data foundation.Data, characteristic Characteristic, writeType WriteType) {
	objc.Call[objc.Void](p, objc.Sel("writeValue:forCharacteristic:type:"), data, characteristic, writeType)
}

// SetNotifyValueForCharacteristic sets the notify value for a characteristic
func (p Peripheral) SetNotifyValueForCharacteristic(enabled bool, characteristic Characteristic) {
	objc.Call[objc.Void](p, objc.Sel("setNotifyValue:forCharacteristic:"), enabled, characteristic)
}

// Service represents a peripheral's service
type Service struct {
	objc.Object
}

// ServiceClass is the class instance for Service
var ServiceClass = objc.GetClass("CBService")

// UUID returns the service's UUID
func (s Service) UUID() foundation.UUID {
	return objc.Call[foundation.UUID](s, objc.Sel("UUID"))
}

// IsPrimary returns whether the service is primary
func (s Service) IsPrimary() bool {
	return objc.Call[bool](s, objc.Sel("isPrimary"))
}

// Characteristics returns the service's characteristics
func (s Service) Characteristics() foundation.Array {
	return objc.Call[foundation.Array](s, objc.Sel("characteristics"))
}

// IncludedServices returns the service's included services
func (s Service) IncludedServices() foundation.Array {
	return objc.Call[foundation.Array](s, objc.Sel("includedServices"))
}

// Characteristic represents a service's characteristic
type Characteristic struct {
	objc.Object
}

// CharacteristicClass is the class instance for Characteristic
var CharacteristicClass = objc.GetClass("CBCharacteristic")

// UUID returns the characteristic's UUID
func (c Characteristic) UUID() foundation.UUID {
	return objc.Call[foundation.UUID](c, objc.Sel("UUID"))
}

// Value returns the characteristic's value
func (c Characteristic) Value() foundation.Data {
	return objc.Call[foundation.Data](c, objc.Sel("value"))
}

// Properties returns the characteristic's properties
func (c Characteristic) Properties() CharacteristicProperties {
	return CharacteristicProperties(objc.Call[int](c, objc.Sel("properties")))
}

// IsNotifying returns whether the characteristic is notifying
func (c Characteristic) IsNotifying() bool {
	return objc.Call[bool](c, objc.Sel("isNotifying"))
}

// Descriptors returns the characteristic's descriptors
func (c Characteristic) Descriptors() foundation.Array {
	return objc.Call[foundation.Array](c, objc.Sel("descriptors"))
}

// Descriptor represents a characteristic's descriptor
type Descriptor struct {
	objc.Object
}

// DescriptorClass is the class instance for Descriptor
var DescriptorClass = objc.GetClass("CBDescriptor")

// UUID returns the descriptor's UUID
func (d Descriptor) UUID() foundation.UUID {
	return objc.Call[foundation.UUID](d, objc.Sel("UUID"))
}

// Value returns the descriptor's value
func (d Descriptor) Value() objc.Object {
	return objc.Call[objc.Object](d, objc.Sel("value"))
}

// PeripheralManager manages local peripheral devices
type PeripheralManager struct {
	objc.Object
}

// PeripheralManagerClass is the class instance for PeripheralManager
var PeripheralManagerClass = objc.GetClass("CBPeripheralManager")

// NewPeripheralManager creates a new peripheral manager with the specified delegate and dispatch queue
func NewPeripheralManager(delegate objc.Object, queue objc.Object) PeripheralManager {
	alloc := objc.Call[objc.Object](PeripheralManagerClass, objc.Sel("alloc"))
	return PeripheralManager{objc.Call[objc.Object](alloc, objc.Sel("initWithDelegate:queue:"), delegate, queue)}
}

// State returns the current state of the peripheral manager
func (pm PeripheralManager) State() State {
	return State(objc.Call[int](pm, objc.Sel("state")))
}

// AddService adds a service to the peripheral manager
func (pm PeripheralManager) AddService(service objc.Object) {
	objc.Call[objc.Void](pm, objc.Sel("addService:"), service)
}

// RemoveService removes a service from the peripheral manager
func (pm PeripheralManager) RemoveService(service objc.Object) {
	objc.Call[objc.Void](pm, objc.Sel("removeService:"), service)
}

// RemoveAllServices removes all services from the peripheral manager
func (pm PeripheralManager) RemoveAllServices() {
	objc.Call[objc.Void](pm, objc.Sel("removeAllServices"))
}

// StartAdvertising starts advertising peripheral data
func (pm PeripheralManager) StartAdvertising(advertisementData foundation.Dictionary) {
	objc.Call[objc.Void](pm, objc.Sel("startAdvertising:"), advertisementData)
}

// StopAdvertising stops advertising peripheral data
func (pm PeripheralManager) StopAdvertising() {
	objc.Call[objc.Void](pm, objc.Sel("stopAdvertising"))
}

// MutableService represents a mutable service
type MutableService struct {
	Service
}

// MutableServiceClass is the class instance for MutableService
var MutableServiceClass = objc.GetClass("CBMutableService")

// NewMutableService creates a new mutable service with the specified UUID and primary status
func NewMutableService(UUID foundation.UUID, isPrimary bool) MutableService {
	alloc := objc.Call[objc.Object](MutableServiceClass, objc.Sel("alloc"))
	return MutableService{Service{objc.Call[objc.Object](alloc, objc.Sel("initWithType:primary:"), UUID, isPrimary)}}
}

// SetCharacteristics sets the service's characteristics
func (ms MutableService) SetCharacteristics(characteristics foundation.Array) {
	objc.Call[objc.Void](ms, objc.Sel("setCharacteristics:"), characteristics)
}

// SetIncludedServices sets the service's included services
func (ms MutableService) SetIncludedServices(includedServices foundation.Array) {
	objc.Call[objc.Void](ms, objc.Sel("setIncludedServices:"), includedServices)
}

// MutableCharacteristic represents a mutable characteristic
type MutableCharacteristic struct {
	Characteristic
}

// MutableCharacteristicClass is the class instance for MutableCharacteristic
var MutableCharacteristicClass = objc.GetClass("CBMutableCharacteristic")

// NewMutableCharacteristic creates a new mutable characteristic with the specified UUID, properties, value, and permissions
func NewMutableCharacteristic(UUID foundation.UUID, properties CharacteristicProperties, value foundation.Data, permissions AttributePermissions) MutableCharacteristic {
	alloc := objc.Call[objc.Object](MutableCharacteristicClass, objc.Sel("alloc"))
	return MutableCharacteristic{Characteristic{objc.Call[objc.Object](alloc, objc.Sel("initWithType:properties:value:permissions:"), UUID, properties, value, permissions)}}
}

// SetValue sets the characteristic's value
func (mc MutableCharacteristic) SetValue(value foundation.Data) {
	objc.Call[objc.Void](mc, objc.Sel("setValue:"), value)
}

// SetDescriptors sets the characteristic's descriptors
func (mc MutableCharacteristic) SetDescriptors(descriptors foundation.Array) {
	objc.Call[objc.Void](mc, objc.Sel("setDescriptors:"), descriptors)
}

// CBUUID represents a Bluetooth UUID
type CBUUID struct {
	objc.Object
}

// CBUUIDClass is the class instance for CBUUID
var CBUUIDClass = objc.GetClass("CBUUID")

// UUIDWithString creates a UUID from a string
func UUIDWithString(UUIDString foundation.String) foundation.UUID {
	return objc.Call[foundation.UUID](CBUUIDClass, objc.Sel("UUIDWithString:"), UUIDString)
}

// UUIDWithData creates a UUID from data
func UUIDWithData(data foundation.Data) foundation.UUID {
	return objc.Call[foundation.UUID](CBUUIDClass, objc.Sel("UUIDWithData:"), data)
}

// State represents the state of a manager
type State int

const (
	StateUnknown      State = 0
	StateResetting    State = 1
	StateUnsupported  State = 2
	StateUnauthorized State = 3
	StatePoweredOff   State = 4
	StatePoweredOn    State = 5
)

// PeripheralState represents the state of a peripheral
type PeripheralState int

const (
	PeripheralStateDisconnected PeripheralState = 0
	PeripheralStateConnecting   PeripheralState = 1
	PeripheralStateConnected    PeripheralState = 2
	PeripheralStateDisconnecting PeripheralState = 3
)

// CharacteristicProperties represents the properties of a characteristic
type CharacteristicProperties int

const (
	CharacteristicPropertyBroadcast                      CharacteristicProperties = 0x01
	CharacteristicPropertyRead                           CharacteristicProperties = 0x02
	CharacteristicPropertyWriteWithoutResponse           CharacteristicProperties = 0x04
	CharacteristicPropertyWrite                          CharacteristicProperties = 0x08
	CharacteristicPropertyNotify                         CharacteristicProperties = 0x10
	CharacteristicPropertyIndicate                       CharacteristicProperties = 0x20
	CharacteristicPropertyAuthenticatedSignedWrites      CharacteristicProperties = 0x40
	CharacteristicPropertyExtendedProperties             CharacteristicProperties = 0x80
	CharacteristicPropertyNotifyEncryptionRequired       CharacteristicProperties = 0x100
	CharacteristicPropertyIndicateEncryptionRequired     CharacteristicProperties = 0x200
)

// AttributePermissions represents the permissions of an attribute
type AttributePermissions int

const (
	AttributePermissionsReadable                       AttributePermissions = 0x01
	AttributePermissionsWriteable                      AttributePermissions = 0x02
	AttributePermissionsReadEncryptionRequired         AttributePermissions = 0x04
	AttributePermissionsWriteEncryptionRequired        AttributePermissions = 0x08
)

// WriteType represents the type of write operation
type WriteType int

const (
	WriteWithResponse      WriteType = 0
	WriteWithoutResponse   WriteType = 1
)