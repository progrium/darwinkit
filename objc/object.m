// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>
#import <objc/NSObject.h>
#include "_cgo_export.h"
#import <objc/runtime.h>

void* C_NSObject_NewObject() {
    NSObject* result_ = [NSObject new];
    return result_;
}

bool Object_IsKindOfClass(void* ptr, void* classPtr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj isKindOfClass:(Class)classPtr];
}

bool Object_IsMemberOfClass(void* ptr, void* classPtr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj isMemberOfClass:(Class)classPtr];
}

bool Object_RespondsToSelector(void* ptr, void* sel) {
    NSObject* obj = (NSObject*)ptr;
    return [obj respondsToSelector:(SEL)sel];
}

bool Object_ConformsToProtocol(void* ptr, void* protocolPtr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj conformsToProtocol:(Protocol*)protocolPtr];
}

void* Object_GetClass(void* ptr) {
    return object_getClass((id)ptr);
}

bool Object_IsProxy(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj isProxy];
}

int Object_RetainCount(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj retainCount];
}

void* Object_Retain(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj retain];
}

void Object_Release(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj release];
}

void* Object_Autorelease(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj autorelease];
}

void* C_NSObject_Copy(void* ptr) {
    NSObject* nSObject = (NSObject*)ptr;
    id result_ = [nSObject copy];
    return result_;
}

void* C_NSObject_MutableCopy(void* ptr) {
    NSObject* nSObject = (NSObject*)ptr;
    id result_ = [nSObject mutableCopy];
    return result_;
}

void Object_Dealloc(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj dealloc];
}

void* Object_PerformSelector(void* ptr, void* sel_p) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p];
}

void* Object_PerformSelector_WithObject(void* ptr, void* sel_p, void* param) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p withObject:(NSObject*)param];
}

void* Object_PerformSelector_WithObject_WithObject(void* ptr, void* sel_p, void* param1, void* param2) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p withObject:(NSObject*)param1  withObject:(NSObject*)param2];
}

const char* Object_Description(void* ptr) {
    return (const char*)[[((NSObject*)ptr) description] performSelector:@selector(UTF8String)];
}