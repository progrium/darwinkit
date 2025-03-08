package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/javascriptcore"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create a window
		frame := foundation.Rect{Size: foundation.Size{600, 400}}
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&window)
		window.SetTitle(foundation.String_StringWithString("JavaScriptCore Example"))
		window.Center()

		// Create a text view
		textView := appkit.NewTextView()
		scrollView := appkit.NewScrollView()
		objc.Retain(&textView)
		objc.Retain(&scrollView)

		scrollView.SetDocumentView(textView)
		window.SetContentView(scrollView)

		// Show the window
		window.MakeKeyAndOrderFront(window)

		// Create a JavaScript context
		jsContext := javascriptcore.NewContext()
		objc.Retain(&jsContext)

		// Set up exception handler
		exceptionBlock := foundation.NewBlockWithVoidObject(func(exception objc.Object) {
			exceptionValue := javascriptcore.Value{exception}
			textView.SetString(foundation.String_StringWithString("JavaScript Error: " + exceptionValue.ToString().UTF8String()))
		})
		
		jsContext.SetExceptionHandler(exceptionBlock)

		// Display initial message
		textView.SetString(foundation.String_StringWithString("JavaScriptCore Demo\n\nExecuting JavaScript...\n"))

		// Run some JavaScript
		script := foundation.String_StringWithString(`
			// Define a simple function
			function calculateSum(a, b) {
				return a + b;
			}
			
			// Call the function with different values
			const results = [];
			results.push("2 + 3 = " + calculateSum(2, 3));
			results.push("5.5 + 4.5 = " + calculateSum(5.5, 4.5));
			results.push("'Hello, ' + 'World!' = " + calculateSum('Hello, ', 'World!'));
			
			// Create an object
			const person = {
				name: 'John Doe',
				age: 30,
				isActive: true,
				greet: function() { return 'Hello, my name is ' + this.name; }
			};
			
			results.push("\\nObject properties:");
			results.push("person.name = " + person.name);
			results.push("person.age = " + person.age);
			results.push("person.isActive = " + person.isActive);
			results.push("person.greet() = " + person.greet());
			
			// Return all results
			results.join('\\n');
		`)
		
		var exception javascriptcore.Value
		result := jsContext.EvaluateScript(script, &exception)
		
		if exception.Pointer() == nil {
			// Add the result to the output
			output := textView.String().UTF8String() + "\n" + result.ToString().UTF8String()
			textView.SetString(foundation.String_StringWithString(output))
		}

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}