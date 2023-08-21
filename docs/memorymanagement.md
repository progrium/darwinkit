# Memory Management

Working with Objective-C from Go requires understanding how [Objective-C memory management works](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/MemoryMgmt.html#//apple_ref/doc/uid/10000011-SW1) and how DarwinKit facilitates the co-existance of the two memory management systems. 

### Quick Version

ARC is a feature of the Objective-C compiler, which does not compile Go code. Therefore memory management of Objective-C objects in Go is done with MMR. DarwinKit continues the policy of "the code that allocates is the code responsible for releasing" so unless you explicitly call the `Alloc()` method, you can assume `Autorelease()` has been called on the object unless documented otherwise. 

This applies to the Go-style constructors made for classes, which are the functions prefixed with "New". This also means all code should be run in an autorelease pool. This is already the case for most delegates and callbacks since these are called from the AppKit event loop, which has an autorelease pool for every cycle of the loop. Code outside this loop (such as code run before appkit.Application is run, or code in a goroutine) should be wrapped in [objc.WithAutoreleasePool](https://pkg.go.dev/github.com/progrium/macdriver@main/objc#WithAutoreleasePool). 

Objects that you want to retain should be passed by reference to [objc.Retain](https://pkg.go.dev/github.com/progrium/macdriver@main/objc#Retain). Using this instead of the object's `Retain()` method directly will let the Go GC know it should release the object before cleaning it up. Objects that only need to exist within the current event cycle don't need to be retained. Objects passed to other objects that need to be retained by them should be retained by the receiving object. If you get an unexplained segfault, chances are an object needed to be retained. A common situation for this are appkit.Windows you create. 

### Full Explanation

Objective-C uses reference counting or a "retain and release" model where objects are deallocated when their internal retain count reaches zero. When an object is allocated, its retain count is set to 1. It is the responsibility of the allocating code to release the object when finished, decrementing its retain count by 1. This would deallocate the object unless other code has retained the object, incrementing the retain count by 1. Every alloc or retain must have a subsequent release.

It should be noted that modern Objective-C and Swift code don't have to do this manual retain and release process because of Automatic Reference Counting, or ARC, which is a feature of their compiler that figures out where to insert retains and releases for you. Our code is compiled by Go so we don't get to take advantage of this, but we do have access to the underlying retain and release methods.

That said, even without ARC there is a way to ease retain and release by using autorelease pools. Along with retain and release methods on all objects, there is also an autorelease method. This marks the object for a deferred release, somewhat like using Go defer to cleanup a resource. Autoreleased objects will be released when the last created autorelease pool on the stack is drained. 

In DarwinKit we have `objc.WithAutoreleasePool()` which takes a function that is immediately executed with a new autorelease pool that is drained when the function is finished. In other words, any objects that have autorelease called from the given function will get released after the function returns. This is equivalent to the `@autoreleasepool` [block syntax](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/mmAutoreleasePools.html#//apple_ref/doc/uid/20000047-CJBFBEDI) in Objective-C. 

Luckily, most of the code we write using Apple frameworks lives in delegate methods or callbacks that are called from an event loop managed by the AppKit framework, and each iteration of the loop has its own autorelease pool. So the common scenario for objects is simply making sure autorelease is called on them after allocating so they will be released at the end of the current cycle of the event loop.

The Objective-C [memory management policy](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/mmRules.html#//apple_ref/doc/uid/20000994-BAJHFBGH) specifies basic rules to know how to use retain and release. The basic idea is that if you allocate an object (using `alloc`, `new`, `copy`, or `mutableCopy`), you own it and are responsible for releasing or autoreleasing it. When an object that you did not explicitly allocate is returned by a function, you do not need to release it unless you take ownership of it by retaining it. This also applies to objects returned by class methods that create new objects like `NSString#stringWithFormat`. Since they are allocating we can assume they are calling autorelease before returning. 

As it turns out, DarwinKit generates Go idiomatic New functions for all classes that does an `alloc` and `init` followed by an `autorelease` before returning. You can always do your own explicit allocation if this is not the desired behavior, but this means in the common case where you are writing code in a delegate or callback that will be called from the application event loop, or are otherwise in an autorelease pool, you can basically write Go code as usual and not have to do anything special. UNLESS you *do* need to take ownership of an object.

If you need the object to live longer than this event loop iteration, you will need to take ownership. So if you assign it to a variable declared outside the event loop, like a global variable, or assign to a struct field or append to a slice that was declared outside the loop, or pass by reference to anything that will need it to stick around, you will probably want to retain the object and have the Go garbage collector be responsible for releasing the object. Another situation you will need to take ownership like this is using the object in a new goroutine.

DarwinKit provides `objc.Retain()`, which calls retain on the object and creates a [Go finalizer](https://pkg.go.dev/runtime#SetFinalizer) that will release the object when Go garbage collects it. We recommend this instead of calling retain and release directly on the object, but you have that option if you know what you're doing. 

Long story short, there's only one special thing you have to do in your Go code to accomodate the Objective-C memory management system, which is use `objc.Retain()` when you need to keep an object around or use it from a goroutine. If you work with objects outside the application event loop, either in programs that don't start it or in the code before starting it, you can just wrap it all with an `objc.WithAutoreleasePool()`. 

That's it. The majority of the concern is dealing with Objective-C objects in "Go space". Go values in "Objective-C space" are not so much a problem because values are either passed by value, are already Objective-C objects, or are delegates or callbacks that are kept from Go garbage collection by the `objc` package.



