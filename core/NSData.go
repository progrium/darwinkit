package core

import (
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

import "C"

// NSData is a static byte buffer in memory.
// https://developer.apple.com/documentation/foundation/nsdata?language=objc
type NSData struct {
	objc.Object
}

// NSData_WithBytes creates a data object containing a given number of bytes copied from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func NSData_WithBytes(b []byte, length uint64) NSData {
	return NSData{objc.Get("NSData").Send("dataWithBytes:length:", uintptr(C.CBytes(b)), length)}
}

// Length is the number of bytes contained by the data object.
// https://developer.apple.com/documentation/foundation/nsdata/1416769-length?language=objc
func (d NSData) Length() uint64 {
	return d.Get("length").Uint()
}

// Bytes returns a slice of bytes containing the data object's contents.
// https://developer.apple.com/documentation/foundation/nsdata/1410616-bytes?language=objc
func (d NSData) Bytes() []byte {
	ptr := d.Get("bytes").Pointer()
	length := d.Length()
	if length == 0 {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(ptr), C.int(length))
}
