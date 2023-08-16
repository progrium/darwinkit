// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactsUserDefaults] class.
var ContactsUserDefaultsClass = _ContactsUserDefaultsClass{objc.GetClass("CNContactsUserDefaults")}

type _ContactsUserDefaultsClass struct {
	objc.Class
}

// An interface definition for the [ContactsUserDefaults] class.
type IContactsUserDefaults interface {
	objc.IObject
	CountryCode() string
	SortOrder() ContactSortOrder
}

// An object that defines the default options to use when displaying contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactsuserdefaults?language=objc
type ContactsUserDefaults struct {
	objc.Object
}

func ContactsUserDefaultsFrom(ptr unsafe.Pointer) ContactsUserDefaults {
	return ContactsUserDefaults{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactsUserDefaultsClass) SharedDefaults() ContactsUserDefaults {
	rv := objc.Call[ContactsUserDefaults](cc, objc.Sel("sharedDefaults"))
	return rv
}

// The singleton contacts user defaults object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactsuserdefaults/1402827-shareddefaults?language=objc
func ContactsUserDefaults_SharedDefaults() ContactsUserDefaults {
	return ContactsUserDefaultsClass.SharedDefaults()
}

func (cc _ContactsUserDefaultsClass) Alloc() ContactsUserDefaults {
	rv := objc.Call[ContactsUserDefaults](cc, objc.Sel("alloc"))
	return rv
}

func ContactsUserDefaults_Alloc() ContactsUserDefaults {
	return ContactsUserDefaultsClass.Alloc()
}

func (cc _ContactsUserDefaultsClass) New() ContactsUserDefaults {
	rv := objc.Call[ContactsUserDefaults](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactsUserDefaults() ContactsUserDefaults {
	return ContactsUserDefaultsClass.New()
}

func (c_ ContactsUserDefaults) Init() ContactsUserDefaults {
	rv := objc.Call[ContactsUserDefaults](c_, objc.Sel("init"))
	return rv
}

// An ISO country code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactsuserdefaults/1403200-countrycode?language=objc
func (c_ ContactsUserDefaults) CountryCode() string {
	rv := objc.Call[string](c_, objc.Sel("countryCode"))
	return rv
}

// Default sorting order by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactsuserdefaults/1402779-sortorder?language=objc
func (c_ ContactsUserDefaults) SortOrder() ContactSortOrder {
	rv := objc.Call[ContactSortOrder](c_, objc.Sel("sortOrder"))
	return rv
}
