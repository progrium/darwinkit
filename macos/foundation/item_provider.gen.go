// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ItemProvider] class.
var ItemProviderClass = _ItemProviderClass{objc.GetClass("NSItemProvider")}

type _ItemProviderClass struct {
	objc.Class
}

// An interface definition for the [ItemProvider] class.
type IItemProvider interface {
	objc.IObject
	LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(url URL, isInPlace bool, error Error)) Progress
	LoadObjectOfClassCompletionHandler(aClass objc.IClass, completionHandler func(object objc.Object, error Error)) Progress
	RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions ItemProviderFileOptions, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(url URL, coordinated bool, error Error)) Progress)
	LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(data []byte, error Error)) Progress
	LoadPreviewImageWithOptionsCompletionHandler(options Dictionary, completionHandler ItemProviderCompletionHandler)
	RegisterObjectOfClassVisibilityLoadHandler(aClass objc.IClass, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(object ItemProviderWritingWrapper, error Error)) Progress)
	LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options Dictionary, completionHandler ItemProviderCompletionHandler)
	LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(url URL, error Error)) Progress
	CanLoadObjectOfClass(aClass objc.IClass) bool
	HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions ItemProviderFileOptions) bool
	RegisterObjectVisibility(object PItemProviderWriting, visibility ItemProviderRepresentationVisibility)
	RegisterObjectObjectVisibility(objectObject objc.IObject, visibility ItemProviderRepresentationVisibility)
	RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler ItemProviderLoadHandler)
	RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(data []byte, error Error)) Progress)
	HasItemConformingToTypeIdentifier(typeIdentifier string) bool
	RegisteredTypeIdentifiersWithFileOptions(fileOptions ItemProviderFileOptions) []string
	PreferredPresentationSize() coregraphics.Size
	SetPreferredPresentationSize(value coregraphics.Size)
	PreviewImageHandler() ItemProviderLoadHandler
	SetPreviewImageHandler(value ItemProviderLoadHandler)
	ContainerFrame() Rect
	SuggestedName() string
	SetSuggestedName(value string)
	SourceFrame() Rect
	RegisteredTypeIdentifiers() []string
}

// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider?language=objc
type ItemProvider struct {
	objc.Object
}

func ItemProviderFrom(ptr unsafe.Pointer) ItemProvider {
	return ItemProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ ItemProvider) InitWithObject(object PItemProviderWriting) ItemProvider {
	po0 := objc.WrapAsProtocol("NSItemProviderWriting", object)
	rv := objc.Call[ItemProvider](i_, objc.Sel("initWithObject:"), po0)
	return rv
}

// Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888328-initwithobject?language=objc
func NewItemProviderWithObject(object PItemProviderWriting) ItemProvider {
	instance := ItemProviderClass.Alloc().InitWithObject(object)
	instance.Autorelease()
	return instance
}

func (i_ ItemProvider) InitWithContentsOfURL(fileURL IURL) ItemProvider {
	rv := objc.Call[ItemProvider](i_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(fileURL))
	return rv
}

// Provides data-backed content from an existing file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403911-initwithcontentsofurl?language=objc
func NewItemProviderWithContentsOfURL(fileURL IURL) ItemProvider {
	instance := ItemProviderClass.Alloc().InitWithContentsOfURL(fileURL)
	instance.Autorelease()
	return instance
}

func (i_ ItemProvider) InitWithItemTypeIdentifier(item objc.IObject, typeIdentifier string) ItemProvider {
	rv := objc.Call[ItemProvider](i_, objc.Sel("initWithItem:typeIdentifier:"), item, typeIdentifier)
	return rv
}

// Creates an item provider with an object, according to the item provider type coercion policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403933-initwithitem?language=objc
func NewItemProviderWithItemTypeIdentifier(item objc.IObject, typeIdentifier string) ItemProvider {
	instance := ItemProviderClass.Alloc().InitWithItemTypeIdentifier(item, typeIdentifier)
	instance.Autorelease()
	return instance
}

func (i_ ItemProvider) Init() ItemProvider {
	rv := objc.Call[ItemProvider](i_, objc.Sel("init"))
	return rv
}

func (ic _ItemProviderClass) Alloc() ItemProvider {
	rv := objc.Call[ItemProvider](ic, objc.Sel("alloc"))
	return rv
}

func ItemProvider_Alloc() ItemProvider {
	return ItemProviderClass.Alloc()
}

func (ic _ItemProviderClass) New() ItemProvider {
	rv := objc.Call[ItemProvider](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewItemProvider() ItemProvider {
	return ItemProviderClass.New()
}

// Asynchronously opens a file in place, if possible, returning a progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888335-loadinplacefilerepresentationfor?language=objc
func (i_ ItemProvider) LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(url URL, isInPlace bool, error Error)) Progress {
	rv := objc.Call[Progress](i_, objc.Sel("loadInPlaceFileRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}

// Asynchronously loads an object of a specified class to an item provider, returning a progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888336-loadobjectofclass?language=objc
func (i_ ItemProvider) LoadObjectOfClassCompletionHandler(aClass objc.IClass, completionHandler func(object objc.Object, error Error)) Progress {
	rv := objc.Call[Progress](i_, objc.Sel("loadObjectOfClass:completionHandler:"), objc.Ptr(aClass), completionHandler)
	return rv
}

// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888337-registerfilerepresentationfortyp?language=objc
func (i_ ItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions ItemProviderFileOptions, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(url URL, coordinated bool, error Error)) Progress) {
	objc.Call[objc.Void](i_, objc.Sel("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:"), typeIdentifier, fileOptions, visibility, loadHandler)
}

// Asynchronously copies the provided, typed data into a generic data object, returning a progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888331-loaddatarepresentationfortypeide?language=objc
func (i_ ItemProvider) LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(data []byte, error Error)) Progress {
	rv := objc.Call[Progress](i_, objc.Sel("loadDataRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}

// Loads the preview image for the item that the item provider represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403925-loadpreviewimagewithoptions?language=objc
func (i_ ItemProvider) LoadPreviewImageWithOptionsCompletionHandler(options Dictionary, completionHandler ItemProviderCompletionHandler) {
	objc.Call[objc.Void](i_, objc.Sel("loadPreviewImageWithOptions:completionHandler:"), options, completionHandler)
}

// Lazily adds representations of a specified object class to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888329-registerobjectofclass?language=objc
func (i_ ItemProvider) RegisterObjectOfClassVisibilityLoadHandler(aClass objc.IClass, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(object ItemProviderWritingWrapper, error Error)) Progress) {
	objc.Call[objc.Void](i_, objc.Sel("registerObjectOfClass:visibility:loadHandler:"), objc.Ptr(aClass), visibility, loadHandler)
}

// Loads the item’s data and coerces it to the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403900-loaditemfortypeidentifier?language=objc
func (i_ ItemProvider) LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options Dictionary, completionHandler ItemProviderCompletionHandler) {
	objc.Call[objc.Void](i_, objc.Sel("loadItemForTypeIdentifier:options:completionHandler:"), typeIdentifier, options, completionHandler)
}

// Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888338-loadfilerepresentationfortypeide?language=objc
func (i_ ItemProvider) LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler func(url URL, error Error)) Progress {
	rv := objc.Call[Progress](i_, objc.Sel("loadFileRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}

// Returns a Boolean value indicating whether an item provider can load objects of a specified class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888333-canloadobjectofclass?language=objc
func (i_ ItemProvider) CanLoadObjectOfClass(aClass objc.IClass) bool {
	rv := objc.Call[bool](i_, objc.Sel("canLoadObjectOfClass:"), objc.Ptr(aClass))
	return rv
}

// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier and to specified open-in-place behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888320-hasrepresentationconformingtotyp?language=objc
func (i_ ItemProvider) HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions ItemProviderFileOptions) bool {
	rv := objc.Call[bool](i_, objc.Sel("hasRepresentationConformingToTypeIdentifier:fileOptions:"), typeIdentifier, fileOptions)
	return rv
}

// Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888321-registerobject?language=objc
func (i_ ItemProvider) RegisterObjectVisibility(object PItemProviderWriting, visibility ItemProviderRepresentationVisibility) {
	po0 := objc.WrapAsProtocol("NSItemProviderWriting", object)
	objc.Call[objc.Void](i_, objc.Sel("registerObject:visibility:"), po0, visibility)
}

// Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888321-registerobject?language=objc
func (i_ ItemProvider) RegisterObjectObjectVisibility(objectObject objc.IObject, visibility ItemProviderRepresentationVisibility) {
	objc.Call[objc.Void](i_, objc.Sel("registerObject:visibility:"), objc.Ptr(objectObject), visibility)
}

// Lazily registers an item, according to the item provider type coercion policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403917-registeritemfortypeidentifier?language=objc
func (i_ ItemProvider) RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler ItemProviderLoadHandler) {
	objc.Call[objc.Void](i_, objc.Sel("registerItemForTypeIdentifier:loadHandler:"), typeIdentifier, loadHandler)
}

// Registers a data-backed representation for an item, specifiying item visibility and a load handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888319-registerdatarepresentationfortyp?language=objc
func (i_ ItemProvider) RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility ItemProviderRepresentationVisibility, loadHandler func(arg0 func(data []byte, error Error)) Progress) {
	objc.Call[objc.Void](i_, objc.Sel("registerDataRepresentationForTypeIdentifier:visibility:loadHandler:"), typeIdentifier, visibility, loadHandler)
}

// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier file options parameter with a value of zero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403921-hasitemconformingtotypeidentifie?language=objc
func (i_ ItemProvider) HasItemConformingToTypeIdentifier(typeIdentifier string) bool {
	rv := objc.Call[bool](i_, objc.Sel("hasItemConformingToTypeIdentifier:"), typeIdentifier)
	return rv
}

// Returns an array with a subset of type identifiers for the item provider, according to the specified file options, in the same order they were registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2888334-registeredtypeidentifierswithfil?language=objc
func (i_ ItemProvider) RegisteredTypeIdentifiersWithFileOptions(fileOptions ItemProviderFileOptions) []string {
	rv := objc.Call[[]string](i_, objc.Sel("registeredTypeIdentifiersWithFileOptions:"), fileOptions)
	return rv
}

// The ideal presentation size of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1528574-preferredpresentationsize?language=objc
func (i_ ItemProvider) PreferredPresentationSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](i_, objc.Sel("preferredPresentationSize"))
	return rv
}

// The ideal presentation size of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1528574-preferredpresentationsize?language=objc
func (i_ ItemProvider) SetPreferredPresentationSize(value coregraphics.Size) {
	objc.Call[objc.Void](i_, objc.Sel("setPreferredPresentationSize:"), value)
}

// The custom preview image handler block for the item provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403904-previewimagehandler?language=objc
func (i_ ItemProvider) PreviewImageHandler() ItemProviderLoadHandler {
	rv := objc.Call[ItemProviderLoadHandler](i_, objc.Sel("previewImageHandler"))
	return rv
}

// The custom preview image handler block for the item provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403904-previewimagehandler?language=objc
func (i_ ItemProvider) SetPreviewImageHandler(value ItemProviderLoadHandler) {
	objc.Call[objc.Void](i_, objc.Sel("setPreviewImageHandler:"), value)
}

// The rectangle of the item’s visible content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1528571-containerframe?language=objc
func (i_ ItemProvider) ContainerFrame() Rect {
	rv := objc.Call[Rect](i_, objc.Sel("containerFrame"))
	return rv
}

// The filename to use when writing the provided data to a file on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2890244-suggestedname?language=objc
func (i_ ItemProvider) SuggestedName() string {
	rv := objc.Call[string](i_, objc.Sel("suggestedName"))
	return rv
}

// The filename to use when writing the provided data to a file on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/2890244-suggestedname?language=objc
func (i_ ItemProvider) SetSuggestedName(value string) {
	objc.Call[objc.Void](i_, objc.Sel("setSuggestedName:"), value)
}

// The rectangle that the item occupies in the host app’s source window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1528569-sourceframe?language=objc
func (i_ ItemProvider) SourceFrame() Rect {
	rv := objc.Call[Rect](i_, objc.Sel("sourceFrame"))
	return rv
}

// Returns the array of type identifiers for the item provider, in the same order they were registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/1403923-registeredtypeidentifiers?language=objc
func (i_ ItemProvider) RegisteredTypeIdentifiers() []string {
	rv := objc.Call[[]string](i_, objc.Sel("registeredTypeIdentifiers"))
	return rv
}
