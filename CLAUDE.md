# DarwinKit Development Notes

## Useful Commands

### Build & Test
```bash
# Build the entire project
go build ./...

# Run tests
go test ./...

# Run specific tests
go test ./macos/appkit
```

### Example Applications
```bash
# Run the screenshot example
cd macos/_examples/screenshot && go run main.go

# Run the webview example
cd macos/_examples/webview && go run main.go
```

## Framework Information

### ScreenCaptureKit
The ScreenCaptureKit implementation supports taking screenshots with robust TCC permission handling:

```go
// Basic screenshot
manager := screencapturekit.SCScreenshotManager_SharedManager()
manager.CaptureScreenshotWithCompletion(func(image appkit.Image, err screencapturekit.SCError) {
    // Process image or handle error
})

// With retry for TCC permission issues
manager.CaptureScreenshotWithRetry(3, time.Second, func(image appkit.Image, err screencapturekit.SCError, success bool) {
    // Handle result with automatic retry for permission issues
})
```

When using ScreenCaptureKit, be aware that:
1. Users must grant screen recording permission via System Preferences
2. Permission issues are common and should be handled gracefully
3. The application may need to be restarted after permission is granted

### Naming Patterns
- Class method calls: `ClassName_MethodName()`
- Instance methods: `instance.MethodName()`
- Create objects: `ClassName_New()` or similar constructor patterns

## Code Style Preferences

### Imports
Order imports as follows:
1. Standard library imports
2. External library imports
3. DarwinKit imports, with `objc` last

```go
import (
    "fmt"
    "os"
    
    "github.com/progrium/darwinkit/macos/appkit"
    "github.com/progrium/darwinkit/macos/foundation"
    "github.com/progrium/darwinkit/objc"
)
```

### Error Handling
For Objective-C errors:
- Check `IsNil()` on error objects before using them
- Use descriptive error messages when possible
- Implement the `Error()` interface on custom error types

## Memory Management
Remember to handle memory appropriately:
- Use `objc.CreateMallocBlock` for callbacks
- Check for `nil` objects before accessing them
- Be aware of ownership rules when working with Objective-C objects
