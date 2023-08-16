// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Connection] class.
var ConnectionClass = _ConnectionClass{objc.GetClass("NSConnection")}

type _ConnectionClass struct {
	objc.Class
}

// An interface definition for the [Connection] class.
type IConnection interface {
	objc.IObject
}

// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconnection?language=objc
type Connection struct {
	objc.Object
}

func ConnectionFrom(ptr unsafe.Pointer) Connection {
	return Connection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ConnectionClass) Alloc() Connection {
	rv := objc.Call[Connection](cc, objc.Sel("alloc"))
	return rv
}

func Connection_Alloc() Connection {
	return ConnectionClass.Alloc()
}

func (cc _ConnectionClass) New() Connection {
	rv := objc.Call[Connection](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConnection() Connection {
	return ConnectionClass.New()
}

func (c_ Connection) Init() Connection {
	rv := objc.Call[Connection](c_, objc.Sel("init"))
	return rv
}
