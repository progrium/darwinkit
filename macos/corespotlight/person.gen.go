// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Person] class.
var PersonClass = _PersonClass{objc.GetClass("CSPerson")}

type _PersonClass struct {
	objc.Class
}

// An interface definition for the [Person] class.
type IPerson interface {
	objc.IObject
	HandleIdentifier() string
	ContactIdentifier() string
	SetContactIdentifier(value string)
	DisplayName() string
	Handles() []string
}

// An object representing a person in the context of search results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson?language=objc
type Person struct {
	objc.Object
}

func PersonFrom(ptr unsafe.Pointer) Person {
	return Person{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ Person) InitWithDisplayNameHandlesHandleIdentifier(displayName string, handles []string, handleIdentifier string) Person {
	rv := objc.Call[Person](p_, objc.Sel("initWithDisplayName:handles:handleIdentifier:"), displayName, handles, handleIdentifier)
	return rv
}

// Returns a new CSPerson object initialized with the specified display name and contact attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618464-initwithdisplayname?language=objc
func NewPersonWithDisplayNameHandlesHandleIdentifier(displayName string, handles []string, handleIdentifier string) Person {
	instance := PersonClass.Alloc().InitWithDisplayNameHandlesHandleIdentifier(displayName, handles, handleIdentifier)
	instance.Autorelease()
	return instance
}

func (pc _PersonClass) Alloc() Person {
	rv := objc.Call[Person](pc, objc.Sel("alloc"))
	return rv
}

func Person_Alloc() Person {
	return PersonClass.Alloc()
}

func (pc _PersonClass) New() Person {
	rv := objc.Call[Person](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPerson() Person {
	return PersonClass.New()
}

func (p_ Person) Init() Person {
	rv := objc.Call[Person](p_, objc.Sel("init"))
	return rv
}

// A key that identifies the type of contact property represented by the person object’s handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618466-handleidentifier?language=objc
func (p_ Person) HandleIdentifier() string {
	rv := objc.Call[string](p_, objc.Sel("handleIdentifier"))
	return rv
}

// The identifier for the contact associated with the person. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618468-contactidentifier?language=objc
func (p_ Person) ContactIdentifier() string {
	rv := objc.Call[string](p_, objc.Sel("contactIdentifier"))
	return rv
}

// The identifier for the contact associated with the person. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618468-contactidentifier?language=objc
func (p_ Person) SetContactIdentifier(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setContactIdentifier:"), value)
}

// A display name for the person. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618469-displayname?language=objc
func (p_ Person) DisplayName() string {
	rv := objc.Call[string](p_, objc.Sel("displayName"))
	return rv
}

// An array of contact handles related to the person. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csperson/1618465-handles?language=objc
func (p_ Person) Handles() []string {
	rv := objc.Call[[]string](p_, objc.Sel("handles"))
	return rv
}
