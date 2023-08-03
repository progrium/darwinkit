// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <objc/NSObject.h>
#include "_cgo_export.h"


@interface ProtocolImplBase : NSObject
@property (assign) uintptr_t goID;
@end

@implementation ProtocolImplBase

- (BOOL)respondsToSelector:(SEL)aSelector {
	return respondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteHandle([self goID]);
	[super dealloc];
}
@end


void* New_ProtocolImpl(void* class, uintptr_t goID) {
    ProtocolImplBase* o =(ProtocolImplBase*)[[[(Class)class alloc] init] autorelease];
    [o setGoID:goID];
    return o;
}


void* Objc_GetProtocol(const char* name) {
    return objc_getProtocol(name);
}

void* Objc_AllocateProtocol(const char* name) {
    return objc_allocateProtocol(name);
}

const char* Protocol_GetName(void *protocol) {
    return protocol_getName((Protocol*)protocol);
}

struct objc_method_description Protocol_GetMethodDescription(void* protocol, void* sel, bool required, bool instanceMethod) {
    return protocol_getMethodDescription((Protocol*)protocol, (SEL)sel, required, instanceMethod);
}

struct objc_method_description* Protocol_CopyMethodDescriptionList(void* protocol, bool required, bool instanceMethod, unsigned int *outCount) {
    return protocol_copyMethodDescriptionList((Protocol*)protocol, required, instanceMethod, outCount);
}

void* Protocol_CopyProtocolList(void* protocol, unsigned int *outCount) {
    return protocol_copyProtocolList((Protocol*)protocol, outCount);
}

void* Protocol_GetProperty(void* protocol, const char *name, bool required, bool isInstanceProperty) {
    return protocol_getProperty((Protocol*)protocol, name, required, isInstanceProperty);
}

void* Protocol_CopyPropertyList(void* protocol, unsigned int *outCount) {
    return protocol_copyPropertyList((Protocol*)protocol, outCount);
}