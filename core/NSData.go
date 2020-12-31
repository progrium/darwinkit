package core

import (
	"github.com/progrium/macdriver/objc"
)

import "C"

type NSData struct {
	objc.Object
}

func NSData_WithBytes(b []byte, length uint64) NSData {
	return NSData{objc.Get("NSData").Send("dataWithBytes:length:", uintptr(C.CBytes(b)), length)}
}
