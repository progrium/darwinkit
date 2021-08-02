package objctest

import "unsafe"

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>

void * numWithBool      (bool v)               { return [NSNumber numberWithBool:v];             }
void * numWithChar      (signed char v)        { return [NSNumber numberWithChar:v];             }
void * numWithShort     (signed short v)       { return [NSNumber numberWithShort:v];            }
void * numWithInt       (signed int v)         { return [NSNumber numberWithInt:v];              }
void * numWithLongLong  (signed long long v)   { return [NSNumber numberWithLongLong:v];         }
void * numWithUChar     (unsigned char v)      { return [NSNumber numberWithUnsignedChar:v];     }
void * numWithUShort    (unsigned short v)     { return [NSNumber numberWithUnsignedShort:v];    }
void * numWithUInt      (unsigned int v)       { return [NSNumber numberWithUnsignedInt:v];      }
void * numWithULongLong (unsigned long long v) { return [NSNumber numberWithUnsignedLongLong:v]; }
void * numWithFloat     (float v)              { return [NSNumber numberWithFloat:v];            }
void * numWithDouble    (double v)             { return [NSNumber numberWithDouble:v];           }

bool               numBoolValue      (void * num) { return [(NSNumber *)num boolValue];             }
signed char        numCharValue      (void * num) { return [(NSNumber *)num charValue];             }
signed short       numShortValue     (void * num) { return [(NSNumber *)num shortValue];            }
signed int         numIntValue       (void * num) { return [(NSNumber *)num intValue];              }
signed long long   numLongLongValue  (void * num) { return [(NSNumber *)num longLongValue];         }
unsigned char      numUCharValue     (void * num) { return [(NSNumber *)num unsignedCharValue];     }
unsigned short     numUShortValue    (void * num) { return [(NSNumber *)num unsignedShortValue];    }
unsigned int       numUIntValue      (void * num) { return [(NSNumber *)num unsignedIntValue];      }
unsigned long long numULongLongValue (void * num) { return [(NSNumber *)num unsignedLongLongValue]; }
float              numFloatValue     (void * num) { return [(NSNumber *)num floatValue];            }
double             numDoubleValue    (void * num) { return [(NSNumber *)num doubleValue];           }

void * stringWith(const char * c) { return [NSString stringWithUTF8String:c]; }
const char * stringValue(void * str) { return [(NSString *)str UTF8String]; }
*/
import "C"

func NSNumberWithBool(v bool) unsafe.Pointer       { return C.numWithBool(C.bool(v)) }
func NSNumberWithInt8(v int8) unsafe.Pointer       { return C.numWithChar(C.schar(v)) }
func NSNumberWithInt16(v int16) unsafe.Pointer     { return C.numWithShort(C.short(v)) }
func NSNumberWithInt32(v int32) unsafe.Pointer     { return C.numWithInt(C.int(v)) }
func NSNumberWithInt64(v int64) unsafe.Pointer     { return C.numWithLongLong(C.longlong(v)) }
func NSNumberWithUint8(v uint8) unsafe.Pointer     { return C.numWithUChar(C.uchar(v)) }
func NSNumberWithUint16(v uint16) unsafe.Pointer   { return C.numWithUShort(C.ushort(v)) }
func NSNumberWithUint32(v uint32) unsafe.Pointer   { return C.numWithUInt(C.uint(v)) }
func NSNumberWithUint64(v uint64) unsafe.Pointer   { return C.numWithULongLong(C.ulonglong(v)) }
func NSNumberWithFloat32(v float32) unsafe.Pointer { return C.numWithFloat(C.float(v)) }
func NSNumberWithFloat64(v float64) unsafe.Pointer { return C.numWithDouble(C.double(v)) }

func NSNumberBoolValue(v unsafe.Pointer) bool       { return bool(C.numBoolValue(v)) }
func NSNumberInt8Value(v unsafe.Pointer) int8       { return int8(C.numCharValue(v)) }
func NSNumberInt16Value(v unsafe.Pointer) int16     { return int16(C.numShortValue(v)) }
func NSNumberInt32Value(v unsafe.Pointer) int32     { return int32(C.numIntValue(v)) }
func NSNumberInt64Value(v unsafe.Pointer) int64     { return int64(C.numLongLongValue(v)) }
func NSNumberUint8Value(v unsafe.Pointer) uint8     { return uint8(C.numUCharValue(v)) }
func NSNumberUint16Value(v unsafe.Pointer) uint16   { return uint16(C.numUShortValue(v)) }
func NSNumberUint32Value(v unsafe.Pointer) uint32   { return uint32(C.numUIntValue(v)) }
func NSNumberUint64Value(v unsafe.Pointer) uint64   { return uint64(C.numULongLongValue(v)) }
func NSNumberFloat32Value(v unsafe.Pointer) float32 { return float32(C.numFloatValue(v)) }
func NSNumberFloat64Value(v unsafe.Pointer) float64 { return float64(C.numDoubleValue(v)) }

func NSStringWith(v string) unsafe.Pointer {
	c := C.CString(v)
	defer C.free(unsafe.Pointer(c))
	return C.stringWith(c)
}

func NSStringValue(v unsafe.Pointer) string {
	return C.GoString(C.stringValue(v))
}
