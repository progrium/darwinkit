package objc

import (
	"fmt"
)

type notAnObject struct {
	value interface{}
}

func (notAnObject) Pointer() uintptr                                      { panic("not an object") }
func (notAnObject) Send(selector string, args ...interface{}) Object      { panic("not an object") }
func (notAnObject) SendSuper(selector string, args ...interface{}) Object { panic("not an object") }
func (notAnObject) Class() Class                                          { panic("not an object") }
func (notAnObject) Alloc() Object                                         { panic("not an object") }
func (notAnObject) Init() Object                                          { panic("not an object") }
func (notAnObject) Retain() Object                                        { panic("not an object") }
func (notAnObject) Release() Object                                       { panic("not an object") }
func (notAnObject) Autorelease() Object                                   { panic("not an object") }
func (notAnObject) Copy() Object                                          { panic("not an object") }
func (notAnObject) Equals(o Object) bool                                  { panic("not an object") }
func (notAnObject) Set(setter string, args ...interface{})                { panic("not an object") }
func (notAnObject) Get(getter string) Object                              { panic("not an object") }
func (notAnObject) GetSt(getter string, ret interface{})                  { panic("not an object") }

func (o notAnObject) String() string  { return fmt.Sprint(o.value) }
func (o notAnObject) Bool() bool      { return o.value.(bool) }
func (o notAnObject) CString() string { return o.value.(string) }
func (o notAnObject) Uint() uint64    { return asUint(o.value) }
func (o notAnObject) Int() int64      { return asInt(o.value) }
func (o notAnObject) Float() float64  { return asFloat(o.value) }
