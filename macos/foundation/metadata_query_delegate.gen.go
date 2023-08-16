// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// An interface that enables the delegate of a metadata query to provide substitute results or attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate?language=objc
type PMetadataQueryDelegate interface {
	// optional
	MetadataQueryReplacementValueForAttributeValue(query MetadataQuery, attrName string, attrValue objc.Object) objc.IObject
	HasMetadataQueryReplacementValueForAttributeValue() bool
}

// A delegate implementation builder for the [PMetadataQueryDelegate] protocol.
type MetadataQueryDelegate struct {
	_MetadataQueryReplacementValueForAttributeValue func(query MetadataQuery, attrName string, attrValue objc.Object) objc.IObject
}

func (di *MetadataQueryDelegate) HasMetadataQueryReplacementValueForAttributeValue() bool {
	return di._MetadataQueryReplacementValueForAttributeValue != nil
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (di *MetadataQueryDelegate) SetMetadataQueryReplacementValueForAttributeValue(f func(query MetadataQuery, attrName string, attrValue objc.Object) objc.IObject) {
	di._MetadataQueryReplacementValueForAttributeValue = f
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (di *MetadataQueryDelegate) MetadataQueryReplacementValueForAttributeValue(query MetadataQuery, attrName string, attrValue objc.Object) objc.IObject {
	return di._MetadataQueryReplacementValueForAttributeValue(query, attrName, attrValue)
}

// A concrete type wrapper for the [PMetadataQueryDelegate] protocol.
type MetadataQueryDelegateWrapper struct {
	objc.Object
}

func (m_ MetadataQueryDelegateWrapper) HasMetadataQueryReplacementValueForAttributeValue() bool {
	return m_.RespondsToSelector(objc.Sel("metadataQuery:replacementValueForAttribute:value:"))
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (m_ MetadataQueryDelegateWrapper) MetadataQueryReplacementValueForAttributeValue(query IMetadataQuery, attrName string, attrValue objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("metadataQuery:replacementValueForAttribute:value:"), objc.Ptr(query), attrName, attrValue)
	return rv
}
