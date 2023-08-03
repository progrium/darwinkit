#import <stdint.h>
#import <objc/runtime.h>
#import <Foundation/NSObject.h>
#include "_cgo_export.h"

@interface ActionHandler : NSObject
@property (assign) uintptr_t goID;
@end

@implementation ActionHandler

- (IBAction)onAction:(id)sender {
	return callAction([self goID], sender);
}
- (void)dealloc {
    deleteActionHandle([self goID]);
    [super dealloc];
}
@end

void* C_NewAction(uintptr_t id) {
    ActionHandler* handler = [[ActionHandler alloc] init];
	[handler setGoID:id];
	return handler;
}