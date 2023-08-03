// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>
#import <objc/runtime.h>


void* Objc_GetClass(const char* name) {
    return objc_getClass(name);
}

void* Objc_AllocateClassPair(void* superClass, const char* name, size_t extraBytes) {
    return objc_allocateClassPair((Class)superClass, name, extraBytes);
}

void Objc_RegisterClassPair(void* class) {
    objc_registerClassPair((Class)class);
}

void Objc_DisposeClassPair(void* class) {
    objc_disposeClassPair((Class)class);
}


void* Class_CreateInstance(void* cls, unsigned idxIvars) {
    return class_createInstance((Class)cls, idxIvars);
}

const char* Class_GetName(void *cls) {
    return class_getName((Class)cls);
}

void Class_SetVersion(void* cls, int version) {
    return class_setVersion((Class)cls, version);
}
int Class_GetVersion(void* cls) {
    return class_getVersion((Class)cls);
}

void* Class_GetSuperClass(void* cls) {
    return class_getSuperclass((Class)cls);
}

bool Class_RespondsToSelectorAddMethod(void* cls, void* sel) {
    return class_respondsToSelector((Class)cls, (SEL)sel);
}

bool Class_AddMethod(void* cls, void* sel, void* imp, const char* types) {
    return class_addMethod((Class)cls, (SEL)sel, (IMP)imp, types);
}

void* Class_ReplaceMethod(void* cls, void* sel, void* imp, const char* types) {
    return class_replaceMethod((Class)cls, (SEL)sel, (IMP)imp, types);
}

void* Class_GetMethodImplementation(void* cls, void* sel) {
    return class_getMethodImplementation((Class)cls, (SEL)sel);
}

void* Class_GetMethodImplementationStret(void* cls, void* sel) {
#ifndef __arm64__
    return class_getMethodImplementation_stret((Class)cls, (SEL)sel);
#else
    return class_getMethodImplementation((Class)cls, (SEL)sel);
#endif
}

void* Class_GetInstanceMethod(void* cls, void* sel) {
    return class_getInstanceMethod((Class)cls, (SEL)sel);
}

void* Class_GetClassMethod(void* cls, void* sel) {
    return class_getClassMethod((Class)cls, (SEL)sel);
}

void* Class_CopyMethodList(void* cls, unsigned int *outCount) {
    return class_copyMethodList((Class)cls, outCount);
}

void* Class_GetProperty(void* cls, const char *name) {
    return class_getProperty((Class)cls, name);
}

void* Class_CopyPropertyList(void* cls, unsigned int *outCount) {
    return class_copyPropertyList((Class)cls, outCount);
}

 bool Class_AddProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount) {
    return class_addProperty((Class)cls, name, (const objc_property_attribute_t *)attributes, attributeCount);
 }

 void Class_ReplaceProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount) {
    class_replaceProperty((Class)cls, name, (const objc_property_attribute_t *)attributes, attributeCount);
 }

bool Class_AddProtocol(void* cls, void* protocol) {
    return class_addProtocol((Class)cls, (Protocol*)protocol);
}