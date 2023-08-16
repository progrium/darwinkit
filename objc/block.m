// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdint.h>
#import <Block.h>
#include "_cgo_export.h"

enum {
    BLOCK_DEALLOCATING =      (0x0001),  // runtime
    BLOCK_REFCOUNT_MASK =     (0xfffe),  // runtime
    BLOCK_NEEDS_FREE =        (1 << 24), // runtime
    BLOCK_HAS_COPY_DISPOSE =  (1 << 25), // compiler
    BLOCK_HAS_CTOR =          (1 << 26), // compiler: helpers have C++ code
    BLOCK_IS_GC =             (1 << 27), // runtime
    BLOCK_IS_GLOBAL =         (1 << 28), // compiler
    BLOCK_USE_STRET =         (1 << 29), // compiler: undefined if !BLOCK_HAS_SIGNATURE
    BLOCK_HAS_SIGNATURE  =    (1 << 30), // compiler
    BLOCK_HAS_EXTENDED_LAYOUT=(1 << 31)  // compiler
};

struct block_descriptor {
    unsigned long int reserved;
    unsigned long int size;
    void (*copy_helper)(void* dst, void* src);
    void (*dispose_helper)(void* src);
    const char* signature;
};

struct block_literal {
    void* isa;
    int   flags;
    int   reserved;
    void (*invoke)(void*);
    struct block_descriptor* descriptor;
    uintptr_t handle;
};

// only stack block copu use this func, fill this func if we add block_create_stack.
static void oc_copy_helper(void* _dst, void* _src) {
    struct block_literal* dst = (struct block_literal*)_dst;
    struct block_literal* src = (struct block_literal*)_src;
}

static void oc_dispose_helper(void* _src) {
    struct block_literal* block = (struct block_literal*)_src;
    free((void*)(block->descriptor));
    free((void*)(block->descriptor->signature));
    deleteHandle(block->handle);
    // block willbe freed by outer dispose
}

static struct block_descriptor gDescriptorTemplate = { 0, sizeof(struct block_literal), oc_copy_helper,
             oc_dispose_helper, 0};

static struct block_literal gLiteralTemplate = {0, BLOCK_HAS_COPY_DISPOSE, 0, 0, &gDescriptorTemplate};


void* block_get_invoke(void* block) {
    return ((struct block_literal*)block)->invoke;
}

void* block_create_global(const char* signature, void* callable) {
    struct block_literal* block;

    block = malloc(sizeof(struct block_literal));
    *block = gLiteralTemplate;
    block->descriptor = malloc(sizeof(struct block_descriptor));
    *(block->descriptor) = *(gLiteralTemplate.descriptor);

    block->descriptor->signature = signature;
    block->flags |= BLOCK_HAS_SIGNATURE | BLOCK_IS_GLOBAL;
    block->isa = objc_lookUpClass("__NSGlobalBlock__");
    block->invoke = callable;
    return (void*)block;
}

void* block_create_malloc(const char* signature, void* callable, uintptr_t handle) {
    struct block_literal* block;

    block = malloc(sizeof(struct block_literal));
    *block = gLiteralTemplate;
    block->descriptor = malloc(sizeof(struct block_descriptor));
    *(block->descriptor) = *(gLiteralTemplate.descriptor);

    block->descriptor->signature = signature;
    block->flags |= BLOCK_HAS_SIGNATURE | BLOCK_HAS_COPY_DISPOSE | BLOCK_NEEDS_FREE | 2;  //  refcount 1
    block->isa = objc_lookUpClass("__NSMallocBlock__");
    block->invoke = callable;
    block->handle = handle;
    return (void*)block;
}

void* block_copy(void *ptr) {
    return Block_copy(ptr);
}

void block_release(void *ptr) {
    Block_release(ptr);
}

// fost test
void* testBlock() {
    int (^blk)(int) = ^int (int count) {
        return count / 2;
    }; 
    return blk;
}