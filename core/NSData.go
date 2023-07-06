package core

// #include <stdlib.h>
import "C"

// NSData is a static byte buffer in memory.
// https://developer.apple.com/documentation/foundation/nsdata?language=objc
type NSData struct {
	gen_NSData
}

// NSData_WithBytes creates a data object containing a given number of bytes copied from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func NSData_WithBytes(b []byte, length uint64) NSData {
	buf := C.CBytes(b)
	defer C.free(buf)
	return NSData_DataWithBytesLength(buf, NSUInteger(length))
}

// Length is the number of bytes contained by the data object.
// https://developer.apple.com/documentation/foundation/nsdata/1416769-length?language=objc
func (d NSData) Length() uint64 {
	return uint64(d.gen_NSData.Length())
}

// Bytes returns a slice of bytes containing the data object's contents.
// https://developer.apple.com/documentation/foundation/nsdata/1410616-bytes?language=objc
func (d NSData) Bytes() []byte {
	length := d.Length()
	if length == 0 {
		return nil
	}
	ptr := d.gen_NSData.Bytes()
	return C.GoBytes(ptr, C.int(length))
}
