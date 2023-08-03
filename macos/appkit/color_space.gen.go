// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

var ColorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}

type _ColorSpaceClass struct {
	objc.Class
}

type IColorSpace interface {
	objc.IObject
	CGColorSpace() coregraphics.ColorSpaceRef
	ColorSpaceModel() ColorSpaceModel
	ColorSyncProfile() unsafe.Pointer
	ICCProfileData() []byte
	LocalizedName() string
	NumberOfColorComponents() int
}

type ColorSpace struct {
	objc.Object
}

func MakeColorSpace(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ ColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithCGColorSpace:"), cgColorSpace)
	return rv
}

func ColorSpace_InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	return ColorSpaceClass.Alloc().InitWithCGColorSpace(cgColorSpace)
}

func (c_ ColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithColorSyncProfile:"), prof)
	return rv
}

func ColorSpace_InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	return ColorSpaceClass.Alloc().InitWithColorSyncProfile(prof)
}

func (c_ ColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithICCProfileData:"), iccData)
	return rv
}

func ColorSpace_InitWithICCProfileData(iccData []byte) ColorSpace {
	return ColorSpaceClass.Alloc().InitWithICCProfileData(iccData)
}

func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("alloc"))
	return rv
}

func ColorSpace_Alloc() ColorSpace {
	return ColorSpaceClass.Alloc()
}

func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

func ColorSpace_New() ColorSpace {
	return ColorSpaceClass.New()
}

func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("init"))
	return rv
}

func ColorSpace_Init() ColorSpace {
	return ColorSpaceClass.Alloc().Init()
}

func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	rv := objc.CallMethod[[]ColorSpace](cc, objc.GetSelector("availableColorSpacesWithModel:"), model)
	// TODO: convert slice items...
	return rv
}

func ColorSpace_AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	return ColorSpaceClass.AvailableColorSpacesWithModel(model)
}

func (cc _ColorSpaceClass) DeviceRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceRGBColorSpace"))
	return rv
}

func ColorSpace_DeviceRGBColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceRGBColorSpace()
}

func (cc _ColorSpaceClass) GenericRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericRGBColorSpace"))
	return rv
}

func ColorSpace_GenericRGBColorSpace() ColorSpace {
	return ColorSpaceClass.GenericRGBColorSpace()
}

func (cc _ColorSpaceClass) DeviceCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceCMYKColorSpace"))
	return rv
}

func ColorSpace_DeviceCMYKColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceCMYKColorSpace()
}

func (cc _ColorSpaceClass) GenericCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericCMYKColorSpace"))
	return rv
}

func ColorSpace_GenericCMYKColorSpace() ColorSpace {
	return ColorSpaceClass.GenericCMYKColorSpace()
}

func (cc _ColorSpaceClass) DeviceGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceGrayColorSpace"))
	return rv
}

func ColorSpace_DeviceGrayColorSpace() ColorSpace {
	return ColorSpaceClass.DeviceGrayColorSpace()
}

func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericGrayColorSpace"))
	return rv
}

func ColorSpace_GenericGrayColorSpace() ColorSpace {
	return ColorSpaceClass.GenericGrayColorSpace()
}

func (cc _ColorSpaceClass) SRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("sRGBColorSpace"))
	return rv
}

func ColorSpace_SRGBColorSpace() ColorSpace {
	return ColorSpaceClass.SRGBColorSpace()
}

func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("extendedSRGBColorSpace"))
	return rv
}

func ColorSpace_ExtendedSRGBColorSpace() ColorSpace {
	return ColorSpaceClass.ExtendedSRGBColorSpace()
}

func (cc _ColorSpaceClass) DisplayP3ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("displayP3ColorSpace"))
	return rv
}

func ColorSpace_DisplayP3ColorSpace() ColorSpace {
	return ColorSpaceClass.DisplayP3ColorSpace()
}

func (cc _ColorSpaceClass) GenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericGamma22GrayColorSpace"))
	return rv
}

func ColorSpace_GenericGamma22GrayColorSpace() ColorSpace {
	return ColorSpaceClass.GenericGamma22GrayColorSpace()
}

func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("extendedGenericGamma22GrayColorSpace"))
	return rv
}

func ColorSpace_ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	return ColorSpaceClass.ExtendedGenericGamma22GrayColorSpace()
}

func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("adobeRGB1998ColorSpace"))
	return rv
}

func ColorSpace_AdobeRGB1998ColorSpace() ColorSpace {
	return ColorSpaceClass.AdobeRGB1998ColorSpace()
}

func (c_ ColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.CallMethod[coregraphics.ColorSpaceRef](c_, objc.GetSelector("CGColorSpace"))
	return rv
}

func (c_ ColorSpace) ColorSpaceModel() ColorSpaceModel {
	rv := objc.CallMethod[ColorSpaceModel](c_, objc.GetSelector("colorSpaceModel"))
	return rv
}

func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](c_, objc.GetSelector("colorSyncProfile"))
	return rv
}

func (c_ ColorSpace) ICCProfileData() []byte {
	rv := objc.CallMethod[[]byte](c_, objc.GetSelector("ICCProfileData"))
	return rv
}

func (c_ ColorSpace) LocalizedName() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedName"))
	return rv
}

func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfColorComponents"))
	return rv
}
