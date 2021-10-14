package core

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// #include <stdlib.h>
import "C"

type NSStringEncoding NSUInteger

const (
	NSUTF8StringEncoding NSStringEncoding = 4
)

// Wrapper for NSString
// https://developer.apple.com/documentation/foundation/nsstring?language=occ
type NSString struct {
	gen_NSString
}

// NSString_FromString returns an initialized NSString object containing a given number of bytes
// from a given buffer of bytes interpreted in a given encoding.
// The buffer of bytes is coming from a the provided string.
// It's preferred to use the shorthand function String.
// https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=occ
func NSString_FromString(s string) NSString {
	b := []byte(s)
	c := C.CBytes(b)
	defer C.free(unsafe.Pointer(c))
	ret := NSString_alloc().InitWithBytes_length_encoding__asNSString(c, NSUInteger(len(b)), NSUTF8StringEncoding)
	return ret
}

func NSString_FromObject(obj objc.Object) NSString {
	return NSString_fromRef(obj)
}

func (s NSString) SizeWithAttributes(attrs NSDictionary) NSSize {
	return s.gen_NSString.SizeWithAttributes_(attrs)
}
