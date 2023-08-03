// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>
#include "_cgo_export.h"

void* IMP_ImplementationWithBlock(void* bp) {
    return imp_implementationWithBlock((id)bp);
}

void* IMP_GetBlock(void* ptr) {
    return imp_getBlock((IMP)ptr);
}

bool IMP_RemoveBlock(void* ptr) {
    return imp_removeBlock((IMP)ptr);
}

void* Method_GetName(void* method) {
    return method_getName((Method)method);
}

const char* Method_GetTypeEncoding(void* method) {
    return method_getTypeEncoding((Method)method);
}

void* Method_GetImplementation(void* method) {
    return method_getImplementation((Method)method);
}

void* Method_SetImplementation(void* method, void* imp) {
    return method_setImplementation((Method)method, (IMP)imp);
}


const char* Property_GetName(void* property) {
    return property_getName((objc_property_t)property);
}

const char* Property_GetAttributes(void* property) {
    return property_getAttributes((objc_property_t)property);
}

char* Property_CopyAttributeValue(void* property, const char* name) {
    return property_copyAttributeValue((objc_property_t)property, name);
}

objc_property_attribute_t * Property_CopyAttributeList(void* property, unsigned int *outCount) {
    return  property_copyAttributeList((objc_property_t)property, outCount);
}

void* Selector_SEL_RegisterName(const char* name) {
    return (void*)sel_registerName(name);
}

const char* Selector_SEL_GetName(void* ptr) {
    return sel_getName((SEL)ptr);
}

void Objc_SetAssociatedObject(void* ptr, const void* key, void* valuePtr, uintptr_t policy) {
    objc_setAssociatedObject((id)ptr, key, (id)valuePtr, policy);
}

void* Objc_GetAssociatedObject(void* ptr, const void* key) {
    return objc_getAssociatedObject((id)ptr, key);
}

void Objc_RemoveAssociatedObjects(void* ptr) {
    return objc_removeAssociatedObjects((id)ptr);
}
