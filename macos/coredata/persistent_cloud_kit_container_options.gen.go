// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/cloudkit"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentCloudKitContainerOptions] class.
var PersistentCloudKitContainerOptionsClass = _PersistentCloudKitContainerOptionsClass{objc.GetClass("NSPersistentCloudKitContainerOptions")}

type _PersistentCloudKitContainerOptionsClass struct {
	objc.Class
}

// An interface definition for the [PersistentCloudKitContainerOptions] class.
type IPersistentCloudKitContainerOptions interface {
	objc.IObject
	DatabaseScope() cloudkit.DatabaseScope
	SetDatabaseScope(value cloudkit.DatabaseScope)
	ContainerIdentifier() string
}

// An object that customizes how a store description aligns with a CloudKit database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions?language=objc
type PersistentCloudKitContainerOptions struct {
	objc.Object
}

func PersistentCloudKitContainerOptionsFrom(ptr unsafe.Pointer) PersistentCloudKitContainerOptions {
	return PersistentCloudKitContainerOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentCloudKitContainerOptions) InitWithContainerIdentifier(containerIdentifier string) PersistentCloudKitContainerOptions {
	rv := objc.Call[PersistentCloudKitContainerOptions](p_, objc.Sel("initWithContainerIdentifier:"), containerIdentifier)
	return rv
}

// Initializes container options using the given CloudKit container identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions/3141675-initwithcontaineridentifier?language=objc
func NewPersistentCloudKitContainerOptionsWithContainerIdentifier(containerIdentifier string) PersistentCloudKitContainerOptions {
	instance := PersistentCloudKitContainerOptionsClass.Alloc().InitWithContainerIdentifier(containerIdentifier)
	instance.Autorelease()
	return instance
}

func (pc _PersistentCloudKitContainerOptionsClass) Alloc() PersistentCloudKitContainerOptions {
	rv := objc.Call[PersistentCloudKitContainerOptions](pc, objc.Sel("alloc"))
	return rv
}

func PersistentCloudKitContainerOptions_Alloc() PersistentCloudKitContainerOptions {
	return PersistentCloudKitContainerOptionsClass.Alloc()
}

func (pc _PersistentCloudKitContainerOptionsClass) New() PersistentCloudKitContainerOptions {
	rv := objc.Call[PersistentCloudKitContainerOptions](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentCloudKitContainerOptions() PersistentCloudKitContainerOptions {
	return PersistentCloudKitContainerOptionsClass.New()
}

func (p_ PersistentCloudKitContainerOptions) Init() PersistentCloudKitContainerOptions {
	rv := objc.Call[PersistentCloudKitContainerOptions](p_, objc.Sel("init"))
	return rv
}

// The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions/3580372-databasescope?language=objc
func (p_ PersistentCloudKitContainerOptions) DatabaseScope() cloudkit.DatabaseScope {
	rv := objc.Call[cloudkit.DatabaseScope](p_, objc.Sel("databaseScope"))
	return rv
}

// The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions/3580372-databasescope?language=objc
func (p_ PersistentCloudKitContainerOptions) SetDatabaseScope(value cloudkit.DatabaseScope) {
	objc.Call[objc.Void](p_, objc.Sel("setDatabaseScope:"), value)
}

// The identifier of the CloudKit container associated with a given store description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions/3141674-containeridentifier?language=objc
func (p_ PersistentCloudKitContainerOptions) ContainerIdentifier() string {
	rv := objc.Call[string](p_, objc.Sel("containerIdentifier"))
	return rv
}
