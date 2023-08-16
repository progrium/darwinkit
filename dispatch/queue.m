#import <stdlib.h>
#import <stdint.h>
#import <dispatch/dispatch.h>
#include "_cgo_export.h"
#import <objc/runtime.h>


void* Dispatch_Get_Main_Queue() {
    return dispatch_get_main_queue();
}

void* Dispatch_Get_Global_Queue(intptr_t identifier, uintptr_t flags) {
    return dispatch_get_global_queue(identifier, flags);
}

void Dispatch_Retain(void* queue) {
    dispatch_retain((dispatch_object_t)queue);
}

void Dispatch_Release(void* queue) {
    dispatch_release((dispatch_object_t)queue);
}

void Dispatch_Async(void* queue, uintptr_t task) {
    dispatch_async((dispatch_queue_t)queue, ^{
        runQueueTask(task);
    });
}

void Dispatch_Sync(void* queue, uintptr_t task) {
    dispatch_sync((dispatch_queue_t)queue, ^{
        runQueueTask(task);
    });
}