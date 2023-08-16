// AUTO-GENERATED CODE, DO NOT MODIFY

package corefoundation

import (
	"unsafe"
)

// Callback function that is called, if present, just before a plug-in's code is unloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfpluginunloadfunction?language=objc
type PlugInUnloadFunction = func(plugIn unsafe.Pointer)

// Prototype of a callback function used to determine if two values or keys in a dictionary are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryequalcallback?language=objc
type DictionaryEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) bool

// Prototype of a callback function used to get a description of a value in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetcopydescriptioncallback?language=objc
type SetCopyDescriptionCallBack = func(value unsafe.Pointer) StringRef

// Callback invoked when certain types of activity takes place on a readable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfreadstreamclientcallback?language=objc
type ReadStreamClientCallBack = func(stream ReadStreamRef, type_ StreamEventType, clientCallBackInfo unsafe.Pointer)

// Prototype of a callback function called to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsethashcallback?language=objc
type SetHashCallBack = func(value unsafe.Pointer) HashCode

// A prototype for a function callback that allocates memory of a requested size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorallocatecallback?language=objc
type AllocatorAllocateCallBack = func(allocSize Index, hint OptionFlags, info unsafe.Pointer) unsafe.Pointer

// Prototype of a callback function used to determine if two values in a bag are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagequalcallback?language=objc
type BagEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) bool

// Prototype of a callback function used to retain a value being added to a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetretaincallback?language=objc
type SetRetainCallBack = func(allocator AllocatorRef, value unsafe.Pointer) unsafe.Pointer

// Callback function invoked by the parser to notify your application of parent/child relationships between XML structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparseraddchildcallback?language=objc
type XMLParserAddChildCallBack = func(parser unsafe.Pointer, parent unsafe.Pointer, child unsafe.Pointer, info unsafe.Pointer)

// Prototype of a callback function used to get a description of a value in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarraycopydescriptioncallback?language=objc
type ArrayCopyDescriptionCallBack = func(value unsafe.Pointer) StringRef

// A prototype for a function callback that retains the given data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorretaincallback?language=objc
type AllocatorRetainCallBack = func(info unsafe.Pointer) unsafe.Pointer

// Prototype of a callback function that may be applied to every value in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetapplierfunction?language=objc
type SetApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// Prototype of a callback function used to release a value before it’s removed from a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagreleasecallback?language=objc
type BagReleaseCallBack = func(allocator AllocatorRef, value unsafe.Pointer)

// Callback invoked when a CFRunLoopObserver object is fired. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopobservercallback?language=objc
type RunLoopObserverCallBack = func(observer RunLoopObserverRef, activity RunLoopActivity, info unsafe.Pointer)

// A prototype for a function callback that deallocates a block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatordeallocatecallback?language=objc
type AllocatorDeallocateCallBack = func(ptr unsafe.Pointer, info unsafe.Pointer)

// A prototype for a function callback that gives the size of memory likely to be allocated, given a certain request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorpreferredsizecallback?language=objc
type AllocatorPreferredSizeCallBack = func(size Index, hint OptionFlags, info unsafe.Pointer) Index

// Callback function invoked by the parser when it needs another reference to the information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserretaincallback?language=objc
type XMLParserRetainCallBack = func(info unsafe.Pointer) unsafe.Pointer

// Prototype of a callback function used to determine if two values in an array are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarrayequalcallback?language=objc
type ArrayEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) bool

// Prototype of a callback function used to get a description of a value in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagcopydescriptioncallback?language=objc
type BagCopyDescriptionCallBack = func(value unsafe.Pointer) StringRef

// Callback invoked when a CFMachPort object is invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmachportinvalidationcallback?language=objc
type MachPortInvalidationCallBack = func(port MachPortRef, info unsafe.Pointer)

// A callback which provides a plug-in the opportunity to dynamically register its types with a host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfplugindynamicregisterfunction?language=objc
type PlugInDynamicRegisterFunction = func(plugIn unsafe.Pointer)

// Callback function invoked by the parser to notify your application that an error has occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserhandleerrorcallback?language=objc
type XMLParserHandleErrorCallBack = func(parser unsafe.Pointer, error XMLParserStatusCode, info unsafe.Pointer) bool

// Callback function invoked for each observer of a notification when the notification is posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnotificationcallback?language=objc
type NotificationCallback = func(center NotificationCenterRef, observer unsafe.Pointer, name NotificationName, object unsafe.Pointer, userInfo DictionaryRef)

// Callback function used to retain a program-defined information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreeretaincallback?language=objc
type TreeRetainCallBack = func(info unsafe.Pointer) unsafe.Pointer

// Prototype of a callback function used to determine if two values in a set are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetequalcallback?language=objc
type SetEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) bool

// Callback function used to apply a function to all members of a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbinaryheapapplierfunction?language=objc
type BinaryHeapApplierFunction = func(val unsafe.Pointer, context unsafe.Pointer)

// Prototype of a callback function used to release a value before it’s removed from a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsetreleasecallback?language=objc
type SetReleaseCallBack = func(allocator AllocatorRef, value unsafe.Pointer)

// Prototype of a callback function that may be applied to every value in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagapplierfunction?language=objc
type BagApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// Callback function invoked when the parser encounters an XML open tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparsercreatexmlstructurecallback?language=objc
type XMLParserCreateXMLStructureCallBack = func(parser unsafe.Pointer, nodeDesc unsafe.Pointer, info unsafe.Pointer) unsafe.Pointer

// Prototype of a callback function that may be applied to every key-value pair in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryapplierfunction?language=objc
type DictionaryApplierFunction = func(key unsafe.Pointer, value unsafe.Pointer, context unsafe.Pointer)

// Callback invoked to process a message received on a CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmessageportcallback?language=objc
type MessagePortCallBack = func(local MessagePortRef, msgid int32, data DataRef, info unsafe.Pointer) DataRef

// Callback function invoked by the parser to notify your application that an XML structure (and all its children) have been completely parsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserendxmlstructurecallback?language=objc
type XMLParserEndXMLStructureCallBack = func(parser unsafe.Pointer, xmlType unsafe.Pointer, info unsafe.Pointer)

// Callback invoked when certain types of activity takes place on a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketcallback?language=objc
type SocketCallBack = func(s SocketRef, type_ SocketCallBackType, address DataRef, data unsafe.Pointer, info unsafe.Pointer)

// Callback function invoked by the parser when handling the information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparsercopydescriptioncallback?language=objc
type XMLParserCopyDescriptionCallBack = func(info unsafe.Pointer) StringRef

// Callback invoked to process a message received on a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmachportcallback?language=objc
type MachPortCallBack = func(port MachPortRef, msg unsafe.Pointer, size Index, info unsafe.Pointer)

// Callback invoked when certain types of activity takes place on a writable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfwritestreamclientcallback?language=objc
type WriteStreamClientCallBack = func(stream WriteStreamRef, type_ StreamEventType, clientCallBackInfo unsafe.Pointer)

// Defines a structure for a callback for a CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffiledescriptorcallback?language=objc
type FileDescriptorCallBack = func(f FileDescriptorRef, callBackTypes OptionFlags, info unsafe.Pointer)

// Prototype of a callback function used to retain a value being added to an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarrayretaincallback?language=objc
type ArrayRetainCallBack = func(allocator AllocatorRef, value unsafe.Pointer) unsafe.Pointer

// A prototype for a function callback that releases the given data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorreleasecallback?language=objc
type AllocatorReleaseCallBack = func(info unsafe.Pointer)

// Callback function invoked by the parser when it wants to release a reference to the information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserreleasecallback?language=objc
type XMLParserReleaseCallBack = func(info unsafe.Pointer)

// Prototype of a callback function invoked to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbaghashcallback?language=objc
type BagHashCallBack = func(value unsafe.Pointer) HashCode

// Type of the callback function used by the CFTree apply function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreeapplierfunction?language=objc
type TreeApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// Prototype of a callback function that may be applied to every value in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarrayapplierfunction?language=objc
type ArrayApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// Prototype of a callback function used to retain a value being added to a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbagretaincallback?language=objc
type BagRetainCallBack = func(allocator AllocatorRef, value unsafe.Pointer) unsafe.Pointer

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfplugininstancegetinterfacefunction?language=objc
type PlugInInstanceGetInterfaceFunction = func(instance PlugInInstanceRef, interfaceName StringRef, ftbl unsafe.Pointer) bool

// Prototype of a callback function used to release a key-value pair before it’s removed from a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryreleasecallback?language=objc
type DictionaryReleaseCallBack = func(allocator AllocatorRef, value unsafe.Pointer)

// Prototype of a callback function invoked to compute a hash code for a key. Hash codes are used when key-value pairs are accessed, added, or removed from a collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryhashcallback?language=objc
type DictionaryHashCallBack = func(value unsafe.Pointer) HashCode

// A prototype for a function callback that reallocates memory of a requested size for an existing block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorreallocatecallback?language=objc
type AllocatorReallocateCallBack = func(ptr unsafe.Pointer, newsize Index, hint OptionFlags, info unsafe.Pointer) unsafe.Pointer

// Callback function that a plug-in author must implement to create a plug-in instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfpluginfactoryfunction?language=objc
type PlugInFactoryFunction = func(allocator AllocatorRef, typeUUID UUIDRef) unsafe.Pointer

// Callback function that compares two values. You provide a pointer to this callback in certain Core Foundation sorting functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcomparatorfunction?language=objc
type ComparatorFunction = func(val1 unsafe.Pointer, val2 unsafe.Pointer, context unsafe.Pointer) ComparisonResult

// Prototype of a callback function used to release a value before it’s removed from an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfarrayreleasecallback?language=objc
type ArrayReleaseCallBack = func(allocator AllocatorRef, value unsafe.Pointer)

// Callback invoked when a CFMessagePort object is invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfmessageportinvalidationcallback?language=objc
type MessagePortInvalidationCallBack = func(ms MessagePortRef, info unsafe.Pointer)

// Callback function used to provide a description of the program-defined information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreecopydescriptioncallback?language=objc
type TreeCopyDescriptionCallBack = func(info unsafe.Pointer) StringRef

// Callback invoked when an asynchronous user notification dialog is dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfusernotificationcallback?language=objc
type UserNotificationCallBack = func(userNotification UserNotificationRef, responseFlags OptionFlags)

// Callback invoked when a CFRunLoopTimer object fires. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunlooptimercallback?language=objc
type RunLoopTimerCallBack = func(timer RunLoopTimerRef, info unsafe.Pointer)

// Prototype of a callback function used to get a description of a value or key in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionarycopydescriptioncallback?language=objc
type DictionaryCopyDescriptionCallBack = func(value unsafe.Pointer) StringRef

// Callback function used to release a previously retained program-defined information pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftreereleasecallback?language=objc
type TreeReleaseCallBack = func(info unsafe.Pointer)

// Prototype of a callback function used to retain a value or key being added to a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdictionaryretaincallback?language=objc
type DictionaryRetainCallBack = func(allocator AllocatorRef, value unsafe.Pointer) unsafe.Pointer

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfplugininstancedeallocateinstancedatafunction?language=objc
type PlugInInstanceDeallocateInstanceDataFunction = func(instanceData unsafe.Pointer)

// A prototype for a function callback that provides a description of the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfallocatorcopydescriptioncallback?language=objc
type AllocatorCopyDescriptionCallBack = func(info unsafe.Pointer) StringRef

// Callback function invoked by the parser to notify your application that an external entity has been referenced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserresolveexternalentitycallback?language=objc
type XMLParserResolveExternalEntityCallBack = func(parser unsafe.Pointer, extID *XMLExternalID, info unsafe.Pointer) DataRef
