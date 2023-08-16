// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/contacts"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Geocoder] class.
var GeocoderClass = _GeocoderClass{objc.GetClass("CLGeocoder")}

type _GeocoderClass struct {
	objc.Class
}

// An interface definition for the [Geocoder] class.
type IGeocoder interface {
	objc.IObject
	ReverseGeocodeLocationCompletionHandler(location ILocation, completionHandler GeocodeCompletionHandler)
	GeocodeAddressStringCompletionHandler(addressString string, completionHandler GeocodeCompletionHandler)
	CancelGeocode()
	GeocodePostalAddressCompletionHandler(postalAddress contacts.IPostalAddress, completionHandler GeocodeCompletionHandler)
	IsGeocoding() bool
}

// An interface for converting between geographic coordinates and place names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder?language=objc
type Geocoder struct {
	objc.Object
}

func GeocoderFrom(ptr unsafe.Pointer) Geocoder {
	return Geocoder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GeocoderClass) Alloc() Geocoder {
	rv := objc.Call[Geocoder](gc, objc.Sel("alloc"))
	return rv
}

func Geocoder_Alloc() Geocoder {
	return GeocoderClass.Alloc()
}

func (gc _GeocoderClass) New() Geocoder {
	rv := objc.Call[Geocoder](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGeocoder() Geocoder {
	return GeocoderClass.New()
}

func (g_ Geocoder) Init() Geocoder {
	rv := objc.Call[Geocoder](g_, objc.Sel("init"))
	return rv
}

// Submits a reverse-geocoding request for the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423621-reversegeocodelocation?language=objc
func (g_ Geocoder) ReverseGeocodeLocationCompletionHandler(location ILocation, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("reverseGeocodeLocation:completionHandler:"), objc.Ptr(location), completionHandler)
}

// Submits a forward-geocoding request using the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423509-geocodeaddressstring?language=objc
func (g_ Geocoder) GeocodeAddressStringCompletionHandler(addressString string, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodeAddressString:completionHandler:"), addressString, completionHandler)
}

// Cancels a pending geocoding request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423562-cancelgeocode?language=objc
func (g_ Geocoder) CancelGeocode() {
	objc.Call[objc.Void](g_, objc.Sel("cancelGeocode"))
}

// Submits a forward-geocoding requesting using the specified Contacts framework information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/2890752-geocodepostaladdress?language=objc
func (g_ Geocoder) GeocodePostalAddressCompletionHandler(postalAddress contacts.IPostalAddress, completionHandler GeocodeCompletionHandler) {
	objc.Call[objc.Void](g_, objc.Sel("geocodePostalAddress:completionHandler:"), objc.Ptr(postalAddress), completionHandler)
}

// A Boolean value indicating whether the receiver is in the middle of geocoding its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/1423765-geocoding?language=objc
func (g_ Geocoder) IsGeocoding() bool {
	rv := objc.Call[bool](g_, objc.Sel("isGeocoding"))
	return rv
}
