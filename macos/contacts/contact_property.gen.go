// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactProperty] class.
var ContactPropertyClass = _ContactPropertyClass{objc.GetClass("CNContactProperty")}

type _ContactPropertyClass struct {
	objc.Class
}

// An interface definition for the [ContactProperty] class.
type IContactProperty interface {
	objc.IObject
	Key() string
	Value() objc.Object
	Contact() Contact
	Label() string
	Identifier() string
}

// An object that represents a property of a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty?language=objc
type ContactProperty struct {
	objc.Object
}

func ContactPropertyFrom(ptr unsafe.Pointer) ContactProperty {
	return ContactProperty{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactPropertyClass) Alloc() ContactProperty {
	rv := objc.Call[ContactProperty](cc, objc.Sel("alloc"))
	return rv
}

func ContactProperty_Alloc() ContactProperty {
	return ContactPropertyClass.Alloc()
}

func (cc _ContactPropertyClass) New() ContactProperty {
	rv := objc.Call[ContactProperty](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactProperty() ContactProperty {
	return ContactPropertyClass.New()
}

func (c_ ContactProperty) Init() ContactProperty {
	rv := objc.Call[ContactProperty](c_, objc.Sel("init"))
	return rv
}

// The key of the contact property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/1403092-key?language=objc
func (c_ ContactProperty) Key() string {
	rv := objc.Call[string](c_, objc.Sel("key"))
	return rv
}

// The value of the property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/1403204-value?language=objc
func (c_ ContactProperty) Value() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("value"))
	return rv
}

// The associated contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/1403022-contact?language=objc
func (c_ ContactProperty) Contact() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("contact"))
	return rv
}

// The label of the labeled value of the property array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/1403379-label?language=objc
func (c_ ContactProperty) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

// The identifier of the labeled value in the array of labeled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/1403307-identifier?language=objc
func (c_ ContactProperty) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}