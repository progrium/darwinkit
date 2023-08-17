# Bindings API

* Frameworks have their own packages
  * Lowercase framework name (AppKit => appkit)
  * Lowercase prefix if super long (uniformtypeidentifiers => uti)
* Symbol prefixes are removed (CGPoint => Point)
* Constants and enums are 1:1
  * Extra k prefix is kept (kCGImageStatusInvalidData => KImageStatusInvalidData)
* Classes (ex: NSWindow)
  * Unexported struct type for class (_WindowClass)
  * Variable for class singleton (WindowClass)
  * Interface for class prefixed with I (IWindow)
    * Embeds superclass interface (IView)
  * Struct for class (Window)
    * Embeds superclass struct (View)
  * Instance methods are 1:1
    * Selector names (setFrame:display: => SetFrameDisplay)
    * Arguments with protocols get alt methods with argument as object
      * ...
    * Longer overlapping selector gets _ suffix
      * reload => Reload
      * reload: => Reload_
  * New function (NewWindow)
    * alloc/init/autorelease
  * Init instance methods get New function variants
    * initWithFrame: => NewWindowWithFrame(...)
    * autorelease
  * Class methods are 1:1 on class type
  * Class methods get function variants
    * windowNumbersWithOptions: => Window_WindowNumbersWithOptions(...)
    * removeFrameUsingName: => Window_RemoveFrameUsingName(...)
* Protocols (NSWindowDelegate)
  * ...