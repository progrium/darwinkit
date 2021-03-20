package core

import (
	"github.com/progrium/macdriver/objc"
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
