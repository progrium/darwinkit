// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableGroup] class.
var MutableGroupClass = _MutableGroupClass{objc.GetClass("CNMutableGroup")}

type _MutableGroupClass struct {
	objc.Class
}

// An interface definition for the [MutableGroup] class.
type IMutableGroup interface {
	IGroup
	SetName(value string)
}

// A mutable object that represents a group of contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablegroup?language=objc
type MutableGroup struct {
	Group
}

func MutableGroupFrom(ptr unsafe.Pointer) MutableGroup {
	return MutableGroup{
		Group: GroupFrom(ptr),
	}
}

func (mc _MutableGroupClass) Alloc() MutableGroup {
	rv := objc.Call[MutableGroup](mc, objc.Sel("alloc"))
	return rv
}

func MutableGroup_Alloc() MutableGroup {
	return MutableGroupClass.Alloc()
}

func (mc _MutableGroupClass) New() MutableGroup {
	rv := objc.Call[MutableGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableGroup() MutableGroup {
	return MutableGroupClass.New()
}

func (m_ MutableGroup) Init() MutableGroup {
	rv := objc.Call[MutableGroup](m_, objc.Sel("init"))
	return rv
}

// The name of the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablegroup/1402928-name?language=objc
func (m_ MutableGroup) SetName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setName:"), value)
}
