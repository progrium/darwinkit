// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import "type_convertion.h"
#import <Foundation/NSString.h>
#import <Foundation/NSData.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* to_ns_string(const char* str) {
    return [NSString stringWithUTF8String:str];
}

const char* to_c_string(void* ptr) {
    return [((NSString*)ptr) UTF8String];
}

void* to_ns_data(data d) {
    if (d.len <= 0) {
        return [NSData data];
    }
    return [NSData dataWithBytes:(Byte *)d.data length:d.len];
}

data to_c_bytes(void* ptr) {
    NSData* nsData = (NSData*)ptr;
    data array = {
        .data = [nsData bytes],
        .len = nsData.length
    };
    return array;
}

void* to_ns_array(array array) {
    NSMutableArray* nsArray = [NSMutableArray arrayWithCapacity:array.len];
    if (array.len > 0) {
        void** arrayData = (void**)array.data;
        for (int i = 0; i < array.len; i++) {
            void* p = arrayData[i];
            [nsArray addObject:(NSObject*)p];
        }
    }
    return nsArray;
}

array to_c_array(void* ptr) {
    NSArray* result_ = (NSArray*)ptr;
    array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
        void** result_Data = malloc(result_count * sizeof(void*));
        for (int i = 0; i < result_count; i++) {
             void* p = [result_ objectAtIndex:i];
             result_Data[i] = p;
        }
        result_Array.data = result_Data;
        result_Array.len = result_count;
    }
    return result_Array;
}

dict to_c_items(void* ptr) {
    NSDictionary* result_ = (NSDictionary*)ptr;
    dict c_dict;
    NSArray * keys = [result_ allKeys];
    int size = [keys count];
    c_dict.len = size;
    if (size > 0) {
        void** key_data = malloc(size * sizeof(void*));
        void** value_data = malloc(size * sizeof(void*));
        for (int i = 0; i < size; i++) {
            NSString* kp = [keys objectAtIndex:i];
            id vp = result_[kp];
            key_data[i] = kp;
            value_data[i] = vp;
        }
        c_dict.key_data = key_data;
        c_dict.value_data = value_data;
    }
    return c_dict;
}

void* to_ns_dict(dict cDict) {
    NSMutableDictionary* dict = [NSMutableDictionary dictionaryWithCapacity:cDict.len];
    if (cDict.len > 0) {
        void** key_data = (void**)cDict.key_data;
        void** value_data = (void**)cDict.value_data;
        for (int i = 0; i < cDict.len; i++) {
            void* kp = key_data[i];
            void* vp = value_data[i];
            [dict setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
        }
    }
    return dict;
}
