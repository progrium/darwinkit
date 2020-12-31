package core

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation -framework CoreFoundation -framework WebKit
#include <dispatch/dispatch.h>
#include <Foundation/Foundation.h>


void dispatchCb_cgo(void *arg) {
	void dispatchCb();
	dispatchCb();
}

void dispatch(void *f) {
	dispatch_async_f(dispatch_get_main_queue(), f,
		(dispatch_function_t)dispatchCb_cgo
	);
}

*/
import "C"
