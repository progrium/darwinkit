package corebluetooth

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// CentralManagerDelegate protocol implementation
type CentralManagerDelegate struct {
	objc.Object
	// Callback functions
	DidUpdateState                   func(central CentralManager)
	DidDiscoverPeripheral            func(central CentralManager, peripheral Peripheral, advertisementData foundation.Dictionary, RSSI foundation.Number)
	DidConnectPeripheral             func(central CentralManager, peripheral Peripheral)
	DidFailToConnectPeripheral       func(central CentralManager, peripheral Peripheral, error foundation.Error)
	DidDisconnectPeripheral          func(central CentralManager, peripheral Peripheral, error foundation.Error)
	WillRestoreState                 func(central CentralManager, state foundation.Dictionary)
}

// PeripheralDelegate protocol implementation
type PeripheralDelegate struct {
	objc.Object
	// Callback functions
	DidDiscoverServices               func(peripheral Peripheral, error foundation.Error)
	DidDiscoverCharacteristicsForService func(peripheral Peripheral, service Service, error foundation.Error)
	DidUpdateValueForCharacteristic   func(peripheral Peripheral, characteristic Characteristic, error foundation.Error)
	DidWriteValueForCharacteristic    func(peripheral Peripheral, characteristic Characteristic, error foundation.Error)
	DidUpdateNotificationState        func(peripheral Peripheral, characteristic Characteristic, error foundation.Error)
	DidDiscoverDescriptorsForCharacteristic func(peripheral Peripheral, characteristic Characteristic, error foundation.Error)
	DidUpdateValueForDescriptor       func(peripheral Peripheral, descriptor Descriptor, error foundation.Error)
}

// PeripheralManagerDelegate protocol implementation
type PeripheralManagerDelegate struct {
	objc.Object
	// Callback functions
	DidUpdateState                   func(peripheral PeripheralManager)
	DidAddService                    func(peripheral PeripheralManager, service Service, error foundation.Error)
	DidStartAdvertising              func(peripheral PeripheralManager, error foundation.Error)
	DidReceiveReadRequest            func(peripheral PeripheralManager, request objc.Object)
	DidReceiveWriteRequests          func(peripheral PeripheralManager, requests foundation.Array)
	IsReadyToUpdateSubscribers       func(peripheral PeripheralManager)
}

// This is a simplified implementation that does not support the full protocol.
// It returns an NSObject instance, not a proper delegate. For proper implementation of
// delegates, we would need to properly implement protocol methods.
// TODO: Update with a proper delegate implementation
func NewCentralManagerDelegate() CentralManagerDelegate {
	obj := objc.Call[objc.Object](objc.GetClass("NSObject"), objc.Sel("alloc"))
	return CentralManagerDelegate{
		Object: objc.Call[objc.Object](obj, objc.Sel("init")),
	}
}

func NewPeripheralDelegate() PeripheralDelegate {
	obj := objc.Call[objc.Object](objc.GetClass("NSObject"), objc.Sel("alloc"))
	return PeripheralDelegate{
		Object: objc.Call[objc.Object](obj, objc.Sel("init")),
	}
}

func NewPeripheralManagerDelegate() PeripheralManagerDelegate {
	obj := objc.Call[objc.Object](objc.GetClass("NSObject"), objc.Sel("alloc"))
	return PeripheralManagerDelegate{
		Object: objc.Call[objc.Object](obj, objc.Sel("init")),
	}
}