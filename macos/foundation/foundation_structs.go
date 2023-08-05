package foundation

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// struct define should be synced with <Foundation/NSRange.h> <Foundation/NSAffineTransform.h> <Foundation/NSDecimal.h>

type Point = coregraphics.Point
type Size = coregraphics.Size
type Rect = coregraphics.Rect

type EdgeInsets struct {
	Top    coregraphics.Float
	Left   coregraphics.Float
	Bottom coregraphics.Float
	Right  coregraphics.Float
}

type Range struct {
	Location objc.UInteger
	Length   objc.UInteger
}

type AffineTransformStruct struct {
	M11 coregraphics.Float
	M12 coregraphics.Float
	M21 coregraphics.Float
	M22 coregraphics.Float
	TX  coregraphics.Float
	TY  coregraphics.Float
}

/*
   signed   int _exponent:8;
   unsigned int _length:4;     // length == 0 && isNegative -> NaN
   unsigned int _isNegative:1;
   unsigned int _isCompact:1;
   unsigned int _reserved:18;
   unsigned short _mantissa[NSDecimalMaxSize];//NSDecimalMaxSize=8
*/

type Decimal struct {
	_         [4]byte
	_mantissa [8]uint16
}
