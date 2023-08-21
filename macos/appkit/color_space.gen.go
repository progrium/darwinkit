// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorSpace] class.
var ColorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}

type _ColorSpaceClass struct {
	objc.Class
}

// An interface definition for the [ColorSpace] class.
type IColorSpace interface {
	objc.IObject
	ICCProfileData() []byte
	ColorSpaceModel() ColorSpaceModel
	LocalizedName() string
	NumberOfColorComponents() int
	ColorSyncProfile() unsafe.Pointer
	CGColorSpace() coregraphics.ColorSpaceRef
}

// An object that represents a custom color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace?language=objc
type ColorSpace struct {
	objc.Object
}

func ColorSpaceFrom(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	rv := objc.Call[ColorSpace](c_, objc.Sel("initWithCGColorSpace:"), cgColorSpace)
	return rv
}

// Initializes and returns a color space object initialized from a Core Graphics color-space object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412059-initwithcgcolorspace?language=objc
func NewColorSpaceWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	instance := ColorSpaceClass.Alloc().InitWithCGColorSpace(cgColorSpace)
	instance.Autorelease()
	return instance
}

func (c_ ColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	rv := objc.Call[ColorSpace](c_, objc.Sel("initWithICCProfileData:"), iccData)
	return rv
}

// Initializes and returns a color space object from the specified ICC profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412094-initwithiccprofiledata?language=objc
func NewColorSpaceWithICCProfileData(iccData []byte) ColorSpace {
	instance := ColorSpaceClass.Alloc().InitWithICCProfileData(iccData)
	instance.Autorelease()
	return instance
}

func (c_ ColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	rv := objc.Call[ColorSpace](c_, objc.Sel("initWithColorSyncProfile:"), prof)
	return rv
}

// Initializes and returns a color space object from the specified ColorSync profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412062-initwithcolorsyncprofile?language=objc
func NewColorSpaceWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	instance := ColorSpaceClass.Alloc().InitWithColorSyncProfile(prof)
	instance.Autorelease()
	return instance
}

func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("alloc"))
	return rv
}

func ColorSpace_Alloc() ColorSpace {
	return ColorSpaceClass.Alloc()
}

func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.Call[ColorSpace](c_, objc.Sel("init"))
	return rv
}

// Returns the list of color spaces available on the system that are displayed in the color panel, in the order they are displayed in the color panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412086-availablecolorspaceswithmodel?language=objc
func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	rv := objc.Call[[]ColorSpace](cc, objc.Sel("availableColorSpacesWithModel:"), model)
	return rv
}

// Returns the list of color spaces available on the system that are displayed in the color panel, in the order they are displayed in the color panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412086-availablecolorspaceswithmodel?language=objc
func ColorSpace_AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	return ColorSpaceClass.AvailableColorSpacesWithModel(model)
}

// The ICC profile data from which the color space was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412078-iccprofiledata?language=objc
func (c_ ColorSpace) ICCProfileData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("ICCProfileData"))
	return rv
}

// A color space object that represents an extended gray color space with a gamma value of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644177-extendedgenericgamma22graycolors?language=objc
func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return rv
}

// A color space object that represents an extended gray color space with a gamma value of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644177-extendedgenericgamma22graycolors?language=objc
func ColorSpace_ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	return ColorSpaceClass.ExtendedGenericGamma22GrayColorSpace()
}

// The model on which the color space is based. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412095-colorspacemodel?language=objc
func (c_ ColorSpace) ColorSpaceModel() ColorSpaceModel {
	rv := objc.Call[ColorSpaceModel](c_, objc.Sel("colorSpaceModel"))
	return rv
}

// A color space object that represents a device-independent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412082-genericrgbcolorspace?language=objc
func (cc _ColorSpaceClass) GenericRGBColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("genericRGBColorSpace"))
	return rv
}

// A color space object that represents a device-independent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412082-genericrgbcolorspace?language=objc
func ColorSpace_GenericRGBColorSpace() ColorSpace {
	return ColorSpaceClass.GenericRGBColorSpace()
}

// A color space object that represents a calibrated or device-dependent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412066-devicergbcolorspace?language=objc
func (cc _ColorSpaceClass) DeviceRGBColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("deviceRGBColorSpace"))
	return rv
}

// A color space object that represents a calibrated or device-dependent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412066-devicergbcolorspace?language=objc
func ColorSpace_DeviceRGBColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceRGBColorSpace()
}

// A color space object that represents a calibrated or device-dependent gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412084-devicegraycolorspace?language=objc
func (cc _ColorSpaceClass) DeviceGrayColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("deviceGrayColorSpace"))
	return rv
}

// A color space object that represents a calibrated or device-dependent gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412084-devicegraycolorspace?language=objc
func ColorSpace_DeviceGrayColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceGrayColorSpace()
}

// A color space object that represents an sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412071-srgbcolorspace?language=objc
func (cc _ColorSpaceClass) SRGBColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("sRGBColorSpace"))
	return rv
}

// A color space object that represents an sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412071-srgbcolorspace?language=objc
func ColorSpace_SRGBColorSpace() ColorSpace {
	return ColorSpaceClass.SRGBColorSpace()
}

// A color space object that represents an extended sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644175-extendedsrgbcolorspace?language=objc
func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("extendedSRGBColorSpace"))
	return rv
}

// A color space object that represents an extended sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644175-extendedsrgbcolorspace?language=objc
func ColorSpace_ExtendedSRGBColorSpace() ColorSpace {
	return ColorSpaceClass.ExtendedSRGBColorSpace()
}

// The localized name of the color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412072-localizedname?language=objc
func (c_ ColorSpace) LocalizedName() string {
	rv := objc.Call[string](c_, objc.Sel("localizedName"))
	return rv
}

// A color space object that represents a calibrated or device-dependent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412090-devicecmykcolorspace?language=objc
func (cc _ColorSpaceClass) DeviceCMYKColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("deviceCMYKColorSpace"))
	return rv
}

// A color space object that represents a calibrated or device-dependent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412090-devicecmykcolorspace?language=objc
func ColorSpace_DeviceCMYKColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceCMYKColorSpace()
}

// The number of components, excluding alpha, the color space supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412099-numberofcolorcomponents?language=objc
func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfColorComponents"))
	return rv
}

// The ColorSync profile from which the color space was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412076-colorsyncprofile?language=objc
func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("colorSyncProfile"))
	return rv
}

// A color space object that represents a device-independent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412088-genericcmykcolorspace?language=objc
func (cc _ColorSpaceClass) GenericCMYKColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("genericCMYKColorSpace"))
	return rv
}

// A color space object that represents a device-independent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412088-genericcmykcolorspace?language=objc
func ColorSpace_GenericCMYKColorSpace() ColorSpace {
	return ColorSpaceClass.GenericCMYKColorSpace()
}

// A color space object that represents an Adobe RGB (1998) color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412064-adobergb1998colorspace?language=objc
func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("adobeRGB1998ColorSpace"))
	return rv
}

// A color space object that represents an Adobe RGB (1998) color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412064-adobergb1998colorspace?language=objc
func ColorSpace_AdobeRGB1998ColorSpace() ColorSpace {
	return ColorSpaceClass.AdobeRGB1998ColorSpace()
}

// A color space object that represents a device-independent gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412060-genericgraycolorspace?language=objc
func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("genericGrayColorSpace"))
	return rv
}

// A color space object that represents a device-independent gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412060-genericgraycolorspace?language=objc
func ColorSpace_GenericGrayColorSpace() ColorSpace {
	return ColorSpaceClass.GenericGrayColorSpace()
}

// The Core Graphics color-space object that represents a color space equivalent to the color space’s. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412073-cgcolorspace?language=objc
func (c_ ColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("CGColorSpace"))
	return rv
}

// A color space object that represents a P3 Display color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644170-displayp3colorspace?language=objc
func (cc _ColorSpaceClass) DisplayP3ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("displayP3ColorSpace"))
	return rv
}

// A color space object that represents a P3 Display color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1644170-displayp3colorspace?language=objc
func ColorSpace_DisplayP3ColorSpace() ColorSpace {
	return ColorSpaceClass.DisplayP3ColorSpace()
}

// A color space object that represents a gray color space with a gamma value of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412103-genericgamma22graycolorspace?language=objc
func (cc _ColorSpaceClass) GenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](cc, objc.Sel("genericGamma22GrayColorSpace"))
	return rv
}

// A color space object that represents a gray color space with a gamma value of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/1412103-genericgamma22graycolorspace?language=objc
func ColorSpace_GenericGamma22GrayColorSpace() ColorSpace {
	return ColorSpaceClass.GenericGamma22GrayColorSpace()
}
