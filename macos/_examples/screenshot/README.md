# ScreenCaptureKit Screenshot Example

This example demonstrates how to use DarwinKit's ScreenCaptureKit implementation to take screenshots with proper permission handling.

## Features Demonstrated

- Taking screenshots with SCScreenshotManager
- Handling TCC permission issues with retry logic
- Providing user guidance when permissions are denied
- Saving screenshots to disk as PNG files

## Running the Example

```bash
go run main.go
```

## Code Overview

1. Initializes SCScreenshotManager
2. Sets up a handler to process the captured image
3. Uses retry logic to handle temporary permission issues
4. Provides guidance if permissions are persistently denied
5. Converts the captured image to PNG format
6. Saves the image to the desktop

## Permission Handling

The example demonstrates best practices for handling TCC permissions:

- Automatic retry for temporary issues
- Clear instructions for the user
- Integration with System Preferences
- Graceful failure with helpful messaging
