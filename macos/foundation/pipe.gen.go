// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Pipe] class.
var PipeClass = _PipeClass{objc.GetClass("NSPipe")}

type _PipeClass struct {
	objc.Class
}

// An interface definition for the [Pipe] class.
type IPipe interface {
	objc.IObject
	FileHandleForReading() FileHandle
	FileHandleForWriting() FileHandle
}

// A one-way communications channel between related processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspipe?language=objc
type Pipe struct {
	objc.Object
}

func PipeFrom(ptr unsafe.Pointer) Pipe {
	return Pipe{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PipeClass) Alloc() Pipe {
	rv := objc.Call[Pipe](pc, objc.Sel("alloc"))
	return rv
}

func Pipe_Alloc() Pipe {
	return PipeClass.Alloc()
}

func (pc _PipeClass) New() Pipe {
	rv := objc.Call[Pipe](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPipe() Pipe {
	return PipeClass.New()
}

func (p_ Pipe) Init() Pipe {
	rv := objc.Call[Pipe](p_, objc.Sel("init"))
	return rv
}

// Returns an NSPipe object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspipe/1580418-pipe?language=objc
func (pc _PipeClass) Pipe() Pipe {
	rv := objc.Call[Pipe](pc, objc.Sel("pipe"))
	return rv
}

// Returns an NSPipe object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspipe/1580418-pipe?language=objc
func Pipe_Pipe() Pipe {
	return PipeClass.Pipe()
}

// The receiver's read file handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspipe/1414352-filehandleforreading?language=objc
func (p_ Pipe) FileHandleForReading() FileHandle {
	rv := objc.Call[FileHandle](p_, objc.Sel("fileHandleForReading"))
	return rv
}

// The receiver's write file handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspipe/1412889-filehandleforwriting?language=objc
func (p_ Pipe) FileHandleForWriting() FileHandle {
	rv := objc.Call[FileHandle](p_, objc.Sel("fileHandleForWriting"))
	return rv
}
