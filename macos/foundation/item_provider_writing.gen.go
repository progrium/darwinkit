// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The protocol for implementing a class to allow an item provider to retrieve data from an instance of the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderwriting?language=objc
type PItemProviderWriting interface {
	// optional
	ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) ItemProviderRepresentationVisibility
	HasItemProviderVisibilityForRepresentationWithTypeIdentifier() bool

	// optional
	LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler func(data []byte, error Error)) IProgress
	HasLoadDataWithTypeIdentifierForItemProviderCompletionHandler() bool

	// optional
	WritableTypeIdentifiersForItemProvider() []string
	HasWritableTypeIdentifiersForItemProvider() bool
}

// A concrete type wrapper for the [PItemProviderWriting] protocol.
type ItemProviderWritingWrapper struct {
	objc.Object
}

func (i_ ItemProviderWritingWrapper) HasItemProviderVisibilityForRepresentationWithTypeIdentifier() bool {
	return i_.RespondsToSelector(objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"))
}

// Asks the item provider for the representation visibility specification for the given UTI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderwriting/2921254-itemprovidervisibilityforreprese?language=objc
func (i_ ItemProviderWritingWrapper) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) ItemProviderRepresentationVisibility {
	rv := objc.Call[ItemProviderRepresentationVisibility](i_, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), typeIdentifier)
	return rv
}

func (i_ ItemProviderWritingWrapper) HasLoadDataWithTypeIdentifierForItemProviderCompletionHandler() bool {
	return i_.RespondsToSelector(objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"))
}

// Loads data of a particular type, identified by the given UTI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderwriting/2888302-loaddatawithtypeidentifier?language=objc
func (i_ ItemProviderWritingWrapper) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler func(data []byte, error Error)) Progress {
	rv := objc.Call[Progress](i_, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), typeIdentifier, completionHandler)
	return rv
}

func (i_ ItemProviderWritingWrapper) HasWritableTypeIdentifiersForItemProvider() bool {
	return i_.RespondsToSelector(objc.Sel("writableTypeIdentifiersForItemProvider"))
}

// An array of UTI strings representing the types of data that can be loaded for an item provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderwriting/2888299-writabletypeidentifiersforitempr?language=objc
func (i_ ItemProviderWritingWrapper) WritableTypeIdentifiersForItemProvider() []string {
	rv := objc.Call[[]string](i_, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return rv
}
