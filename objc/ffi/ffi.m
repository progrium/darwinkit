// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdint.h>
#import <ffi/ffi.h>
#include "_cgo_export.h"


ffi_status ffi_prep_cif0(uintptr_t cif, ffi_abi abi, unsigned int nargs, uintptr_t rtype, uintptr_t atypes) {
    return ffi_prep_cif((ffi_cif *)cif, abi, nargs, (ffi_type *)rtype, (ffi_type **)atypes);
}
void ffi_call0(uintptr_t cif, void* fn, uintptr_t rvalue, uintptr_t avalue) {
    return ffi_call((ffi_cif *)cif, fn, (void *)rvalue, (void **)avalue);
}
void *ffi_closure_alloc0(uintptr_t code) {
    return ffi_closure_alloc(sizeof(ffi_closure), (void **)code);
}
ffi_status ffi_prep_closure_loc0(void* closure, uintptr_t cif, void* fn, uintptr_t user_data, void *codeloc) {
	   return ffi_prep_closure_loc((ffi_closure *)closure, (ffi_cif *)cif, fn, (void *)user_data, codeloc);
}

void forward_to_go(ffi_cif *cif, void *ret, void* args[], void * user_data) {
    handleClosure(cif, ret, args, user_data);
}