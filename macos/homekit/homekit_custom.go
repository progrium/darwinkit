package homekit

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// HomeManager provides an interface for managing homes and accessories
type HomeManager struct {
	objc.Object
}

// HomeManagerClass is the class instance for HomeManager
var HomeManagerClass = objc.GetClass("HMHomeManager")

// NewHomeManager creates a new home manager instance
func NewHomeManager() HomeManager {
	return HomeManager{objc.Call[objc.Object](HomeManagerClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetDelegate sets the delegate for the home manager
func (h HomeManager) SetDelegate(delegate objc.Object) {
	h.Send(objc.Sel("setDelegate:"), delegate)
}

// Primary home methods
func (h HomeManager) PrimaryHome() Home {
	return Home{objc.Call[objc.Object](h, objc.Sel("primaryHome"))}
}

// Homes returns an array of homes
func (h HomeManager) Homes() foundation.Array {
	return objc.Call[foundation.Array](h, objc.Sel("homes"))
}

// AddHomeWithNameCompletionHandler creates a new home
func (h HomeManager) AddHomeWithNameCompletionHandler(name foundation.String, completion foundation.CompletionHandler) {
	h.Send(objc.Sel("addHomeWithName:completionHandler:"), name, completion)
}

// Home represents a collection of rooms and accessories
type Home struct {
	objc.Object
}

// HomeClass is the class instance for Home
var HomeClass = objc.GetClass("HMHome")

// Name returns the name of the home
func (h Home) Name() foundation.String {
	return objc.Call[foundation.String](h, objc.Sel("name"))
}

// SetName sets the name of the home
func (h Home) SetName(name foundation.String, completion foundation.CompletionHandler) {
	h.Send(objc.Sel("setName:completionHandler:"), name, completion)
}

// Rooms returns the rooms in the home
func (h Home) Rooms() foundation.Array {
	return objc.Call[foundation.Array](h, objc.Sel("rooms"))
}

// AddRoomWithNameCompletionHandler adds a new room to the home
func (h Home) AddRoomWithNameCompletionHandler(name foundation.String, completion foundation.CompletionHandler) {
	h.Send(objc.Sel("addRoomWithName:completionHandler:"), name, completion)
}

// Accessories returns the accessories in the home
func (h Home) Accessories() foundation.Array {
	return objc.Call[foundation.Array](h, objc.Sel("accessories"))
}

// Room represents a room in a home
type Room struct {
	objc.Object
}

// RoomClass is the class instance for Room
var RoomClass = objc.GetClass("HMRoom")

// Name returns the name of the room
func (r Room) Name() foundation.String {
	return objc.Call[foundation.String](r, objc.Sel("name"))
}

// SetName sets the name of the room
func (r Room) SetName(name foundation.String, completion foundation.CompletionHandler) {
	r.Send(objc.Sel("setName:completionHandler:"), name, completion)
}

// Accessories returns the accessories in the room
func (r Room) Accessories() foundation.Array {
	return objc.Call[foundation.Array](r, objc.Sel("accessories"))
}

// Accessory represents a HomeKit accessory
type Accessory struct {
	objc.Object
}

// AccessoryClass is the class instance for Accessory
var AccessoryClass = objc.GetClass("HMAccessory")

// Name returns the name of the accessory
func (a Accessory) Name() foundation.String {
	return objc.Call[foundation.String](a, objc.Sel("name"))
}

// SetName sets the name of the accessory
func (a Accessory) SetName(name foundation.String, completion foundation.CompletionHandler) {
	a.Send(objc.Sel("setName:completionHandler:"), name, completion)
}

// Room returns the room that contains the accessory
func (a Accessory) Room() Room {
	return Room{objc.Call[objc.Object](a, objc.Sel("room"))}
}

// Services returns the services provided by the accessory
func (a Accessory) Services() foundation.Array {
	return objc.Call[foundation.Array](a, objc.Sel("services"))
}

// Service represents a service provided by an accessory
type Service struct {
	objc.Object
}

// ServiceClass is the class instance for Service
var ServiceClass = objc.GetClass("HMService")

// Name returns the name of the service
func (s Service) Name() foundation.String {
	return objc.Call[foundation.String](s, objc.Sel("name"))
}

// ServiceType returns the type of the service
func (s Service) ServiceType() foundation.String {
	return objc.Call[foundation.String](s, objc.Sel("serviceType"))
}

// Characteristics returns the characteristics of the service
func (s Service) Characteristics() foundation.Array {
	return objc.Call[foundation.Array](s, objc.Sel("characteristics"))
}

// Characteristic represents a characteristic of a service
type Characteristic struct {
	objc.Object
}

// CharacteristicClass is the class instance for Characteristic
var CharacteristicClass = objc.GetClass("HMCharacteristic")

// CharacteristicType returns the type of the characteristic
func (c Characteristic) CharacteristicType() foundation.String {
	return objc.Call[foundation.String](c, objc.Sel("characteristicType"))
}

// Value returns the current value of the characteristic
func (c Characteristic) Value() objc.Object {
	return objc.Call[objc.Object](c, objc.Sel("value"))
}

// WriteValueCompletionHandler writes a new value to the characteristic
func (c Characteristic) WriteValueCompletionHandler(value objc.Object, completion foundation.CompletionHandler) {
	c.Send(objc.Sel("writeValue:completionHandler:"), value, completion)
}

// ReadValueCompletionHandler reads the current value from the characteristic
func (c Characteristic) ReadValueCompletionHandler(completion foundation.CompletionHandler) {
	c.Send(objc.Sel("readValueWithCompletionHandler:"), completion)
}

// Common service types
const (
	ServiceTypeLightbulb = "00000043-0000-1000-8000-0026BB765291"
	ServiceTypeSwitch    = "00000049-0000-1000-8000-0026BB765291"
	ServiceTypeThermostat = "0000004A-0000-1000-8000-0026BB765291"
	ServiceTypeFan       = "00000040-0000-1000-8000-0026BB765291"
	ServiceTypeOutlet    = "00000047-0000-1000-8000-0026BB765291"
	ServiceTypeLockMechanism = "00000045-0000-1000-8000-0026BB765291"
	ServiceTypeGarageDoorOpener = "00000041-0000-1000-8000-0026BB765291"
)

// Common characteristic types
const (
	CharacteristicTypePowerState = "00000025-0000-1000-8000-0026BB765291"
	CharacteristicTypeBrightness = "00000008-0000-1000-8000-0026BB765291"
	CharacteristicTypeHue        = "00000013-0000-1000-8000-0026BB765291"
	CharacteristicTypeSaturation = "0000002F-0000-1000-8000-0026BB765291"
	CharacteristicTypeTemperature = "00000011-0000-1000-8000-0026BB765291"
	CharacteristicTypeTargetTemperature = "00000035-0000-1000-8000-0026BB765291"
	CharacteristicTypeLockCurrentState = "0000001D-0000-1000-8000-0026BB765291"
	CharacteristicTypeLockTargetState = "0000001E-0000-1000-8000-0026BB765291"
)