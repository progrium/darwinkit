// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var FontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}

type _FontDescriptorClass struct {
	objc.Class
}

type IFontDescriptor interface {
	objc.IObject
	FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor
	FontDescriptorWithFace(newFace string) FontDescriptor
	FontDescriptorWithFamily(newFamily string) FontDescriptor
	FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor
	FontDescriptorWithSize(newPointSize float64) FontDescriptor
	FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor
	ObjectForKey(attribute FontDescriptorAttributeName) objc.Object
	FontAttributes() map[FontDescriptorAttributeName]objc.Object
	Matrix() foundation.AffineTransform
	PointSize() float64
	PostscriptName() string
	SymbolicTraits() FontDescriptorSymbolicTraits
	RequiresFontAssetRequest() bool
}

type FontDescriptor struct {
	objc.Object
}

func MakeFontDescriptor(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (f_ FontDescriptor) InitWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("initWithFontAttributes:"), attributes)
	return rv
}

func FontDescriptor_InitWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	return FontDescriptorClass.Alloc().InitWithFontAttributes(attributes)
}

func (f_ FontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithDesign:"), design)
	return rv
}

func FontDescriptor_FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	return FontDescriptorClass.Alloc().FontDescriptorWithDesign(design)
}

func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("alloc"))
	return rv
}

func FontDescriptor_Alloc() FontDescriptor {
	return FontDescriptorClass.Alloc()
}

func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

func FontDescriptor_New() FontDescriptor {
	return FontDescriptorClass.New()
}

func (f_ FontDescriptor) Init() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("init"))
	return rv
}

func FontDescriptor_Init() FontDescriptor {
	return FontDescriptorClass.Alloc().Init()
}

func (fc _FontDescriptorClass) PreferredFontDescriptorForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("preferredFontDescriptorForTextStyle:options:"), style, options)
	return rv
}

func FontDescriptor_PreferredFontDescriptorForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	return FontDescriptorClass.PreferredFontDescriptorForTextStyleOptions(style, options)
}

func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}

func FontDescriptor_FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	return FontDescriptorClass.FontDescriptorWithFontAttributes(attributes)
}

func (fc _FontDescriptorClass) FontDescriptorWithNameMatrix(fontName string, matrix foundation.IAffineTransform) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("fontDescriptorWithName:matrix:"), fontName, objc.ExtractPtr(matrix))
	return rv
}

func FontDescriptor_FontDescriptorWithNameMatrix(fontName string, matrix foundation.IAffineTransform) FontDescriptor {
	return FontDescriptorClass.FontDescriptorWithNameMatrix(fontName, matrix)
}

func (fc _FontDescriptorClass) FontDescriptorWithNameSize(fontName string, size float64) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.GetSelector("fontDescriptorWithName:size:"), fontName, size)
	return rv
}

func FontDescriptor_FontDescriptorWithNameSize(fontName string, size float64) FontDescriptor {
	return FontDescriptorClass.FontDescriptorWithNameSize(fontName, size)
}

func (f_ FontDescriptor) FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorByAddingAttributes:"), attributes)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithFace:"), newFace)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithFamily:"), newFamily)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithMatrix:"), objc.ExtractPtr(matrix))
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSize(newPointSize float64) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithSize:"), newPointSize)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptorWithSymbolicTraits:"), symbolicTraits)
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor {
	rv := objc.CallMethod[[]FontDescriptor](f_, objc.GetSelector("matchingFontDescriptorsWithMandatoryKeys:"), objc.ExtractPtr(mandatoryKeys))
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("matchingFontDescriptorWithMandatoryKeys:"), objc.ExtractPtr(mandatoryKeys))
	return rv
}

func (f_ FontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("objectForKey:"), attribute)
	return rv
}

func (f_ FontDescriptor) FontAttributes() map[FontDescriptorAttributeName]objc.Object {
	rv := objc.CallMethod[map[FontDescriptorAttributeName]objc.Object](f_, objc.GetSelector("fontAttributes"))
	return rv
}

func (f_ FontDescriptor) Matrix() foundation.AffineTransform {
	rv := objc.CallMethod[foundation.AffineTransform](f_, objc.GetSelector("matrix"))
	return rv
}

func (f_ FontDescriptor) PointSize() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("pointSize"))
	return rv
}

func (f_ FontDescriptor) PostscriptName() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("postscriptName"))
	return rv
}

func (f_ FontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	rv := objc.CallMethod[FontDescriptorSymbolicTraits](f_, objc.GetSelector("symbolicTraits"))
	return rv
}

func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("requiresFontAssetRequest"))
	return rv
}
