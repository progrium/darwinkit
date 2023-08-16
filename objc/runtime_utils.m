// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <stdint.h>
#import <objc/runtime.h>
#import <objc/NSObject.h>
#include "_cgo_export.h"


@interface DeallocListener : NSObject
@property (assign) uintptr_t goID;
@end

@implementation DeallocListener

- (void)dealloc {
    runTaskAndDeleteHandle([self goID]);
    [super dealloc];
}
@end

void* C_NewDeallocListener(uintptr_t id) {
    DeallocListener* listener = [[DeallocListener alloc] init];
	[listener setGoID:id];
	return listener;
}

void Run_WithAutoreleasePool(uintptr_t ptr) {
    @autoreleasepool {
        runTaskAndDeleteHandle(ptr);
    }
}