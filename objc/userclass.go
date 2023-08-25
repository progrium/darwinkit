package objc

import "reflect"

// UserClass is a generic wrapper around Class returned by NewClass.
type UserClass[T IObject] struct {
	Class
}

// New creates an instance of the class then calls init and autorelease before returning.
func (c UserClass[T]) New() T {
	var o T
	oo := c.CreateInstance(0).PerformSelector(Sel("init"))
	setPtr(&o, oo.Ptr())
	o.Autorelease()
	return o
}

// NewClass will allocate a new class using the name of the type passed
// and NSObject as the superclass unless the passed type has a struct tag
// with the key "objc" specifying the name of the superclass to use.
// The returned class will still need to be registered.
func NewClass[T IObject](selectors ...Selector) UserClass[T] {
	var o T
	typ := reflect.TypeOf(o)
	var super string
	for i := 0; i < typ.NumField(); i++ {
		super = typ.Field(i).Tag.Get("objc")
		if super != "" {
			break
		}
	}
	if super == "" {
		super = "NSObject"
	}
	cls := AllocateClass(GetClass(super), typ.Name(), 0)
	for _, sel := range selectors {
		m, ok := typ.MethodByName(selectorToGoName(sel.Name()))
		if !ok {
			panic("allocating class from struct without method for selector: " + sel.Name())
		}
		AddMethod(cls, sel, m.Func.Interface())
	}
	return UserClass[T]{
		Class: cls,
	}
}
