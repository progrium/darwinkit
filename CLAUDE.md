# DarwinKit Development Guidelines

## Build Commands
- Test all packages: `go test ./...`
- Test a specific package: `go test ./macos/appkit`
- Test a specific test function: `go test ./macos/appkit -run TestSpecificFunction`
- Generate bindings: `go generate ./...`
- Generate for a framework: `go generate ./macos/appkit`
- Run linting: `go vet ./...`
- Regenerate all frameworks: `./generate/tools/regen.sh macos`
- Clobber generated files: `go run ./generate/tools/clobbergen.go ./macos/[framework]`
- Generate framework enums: `go run ./generate/tools/enumexport.go [framework] > ./generate/modules/enums/macos/[framework]/enums.go`

## Code Style
- Use Go idiomatic code patterns when possible
- Follow standard Go formatting with `gofmt`
- Import ordering: standard library, then third-party, then internal packages
- Error handling: Always check errors returned by functions
- When generating bindings, ensure unique method signatures
- Follow existing naming conventions for generated code
- Memory management: Use `objc.Retain()` for objects that need to live beyond event loops
- Wrap code outside event loops in `objc.WithAutoreleasePool()`
- Follow binding API conventions (see docs/bindings.md):
  - Symbol prefixes are removed (NSWindow → Window)
  - Selector names converted to PascalCase (setFrame:display: → SetFrameDisplay)
  - Class methods get function variants with class name prefix

## Foundation and Objective-C Usage
- Create Foundation objects using factory methods, not struct literals:
  - Use `foundation.ArrayWithObjects([]objc.Object{})` rather than `foundation.Array{}`
  - Use `foundation.DictionaryWithCapacity(0)` rather than `foundation.Dictionary{}`
  - Use `foundation.StringWithString("key")` rather than `foundation.String{"key"}`
  - Use `foundation.DataClass.DataWithBytesLength(dataBytes, dataLength)` for byte arrays
- For protocols, use the concrete type implementation (`DrawableObject` not `Drawable`)
- Convert raw Objective-C objects using explicit type casting: `metal.DrawableObject{objcObject}`
- Use proper enum naming from the generated files, which keeps Apple's prefixes:
  - Example: `NEVPNStatusConnected` instead of `VPNStatusConnected`

## Objective-C Method Calling Conventions
- Always use `objc.Call` for method calls, not the deprecated `Send` method:
  - Use `objc.Call[objc.Void](object, objc.Sel("methodName:"), param)` for methods with no return value
  - Use `objc.Call[ReturnType](object, objc.Sel("methodWithReturn"), param)` for methods with return values
- For initialization, separate the alloc and init calls:
  ```go
  // Old style (deprecated):
  obj := objc.Call[objc.Object](SomeClass, objc.Sel("alloc")).Send(objc.Sel("init"))
  
  // New style:
  alloc := objc.Call[objc.Object](SomeClass, objc.Sel("alloc"))
  obj := objc.Call[objc.Object](alloc, objc.Sel("init"))
  ```
- When working with MutableDictionary, use SetObjectForKeyObject method:
  ```go
  // Create a mutable dictionary
  dict := foundation.MutableDictionaryClass.Dictionary()
  
  // Set key-value pairs
  dict.SetObjectForKeyObject(value, key)
  ```

## Generation Pipeline
DarwinKit uses a multi-stage code generation pipeline to create Go bindings for Apple frameworks:

1. **symbolsdb**: The core database of Apple framework symbols (JSON files in a zip archive)
2. **declparse**: Parses Objective-C declarations into structured Go representations
3. **modules**: Defines framework metadata and dependency relationships
4. **typing**: Maps Objective-C types to Go types
5. **codegen**: Generates Go code from parsed declarations and type mappings
6. **enums**: Exports constants and enums from frameworks to Go code

Generation often requires an iterative approach to handle framework-specific quirks:
1. Generate → encounter panic for unknown situation
2. Update code to handle the situation (add to skip/abstract lists or fix parsing)
3. Generate again until successful

## Adding New Frameworks
Follow these detailed steps to add a new framework:

1. **Add to modules.go**:
   ```go
   {"FrameworkName", "Framework Display Name", "packagename", "FrameworkName/FrameworkName.h", []string{"NS", "Other prefixes"}},
   ```

2. **Handle Dependencies**:
   - Update coupling maps in modules.go if needed:
     - `CanAbstractModuleCoupling`: Makes types become generic interfaces
     - `CanSkipModuleCoupling`: Skips methods/properties using certain modules
     - `CanIgnoreNotFound`: List of modules that can be ignored if not found

3. **Export Constants**:
   ```
   go run ./generate/tools/enumexport.go [framework] > ./generate/modules/enums/macos/[framework]
   ```

4. **Initialize Framework Package**:
   ```
   go run ./generate/tools/initmod.go macos [framework]
   ```

5. **Generate Structs**:
   ```
   go run ./generate/tools/structs.go [framework] > ./macos/[framework]/[framework]_structs.go
   ```
   - Check the output file for any missing structs
   - Manually handle structs with `_Ctype_struct_` prefix by commenting out or replacing

6. **Generate Bindings**:
   ```
   go generate ./macos/[framework]
   ```
   - Handle panics by adding types to appropriate lists
   - Common issues include unknown types or dependency conflicts

7. **Test the Framework**:
   ```
   go test ./macos/[framework]
   ```
   - Fix compilation errors by handling missing types
   - Most often struct types from other frameworks

8. **Handle Circular Imports**:
   - Use `CanAbstractModuleCoupling` or `CanSkipModuleCoupling` to manage dependencies
   - For circular dependencies between frameworks, make one depend on the other as an abstract type

## Framework Coverage
- See API_ENHANCED_COVERAGE.md for detailed framework coverage analysis and statistics
- DarwinKit currently implements **41 frameworks** (approximately 45-50% of critical Apple frameworks)

### Implementation Status By Category
| Category | Implementation Level | Key Frameworks |
|----------|---------------------|----------------|
| Core | ~70% | Foundation (60%), CoreFoundation (95%) |
| UI | ~30% | AppKit (20%) |
| Graphics & Media | ~35% | CoreGraphics (50%), CoreImage, AVFoundation |
| Data & Storage | ~40% | CoreData, CoreSpotlight |
| Device Features | ~25% | CoreLocation, MapKit, HealthKit |
| Specialized APIs | ~15% | GameKit, JavaScriptCore, Security |

### Recently Added Frameworks (March 2025)
| Framework | Core Classes | Methods | Status |
|-----------|--------------|---------|--------|
| EventKit  | EventStore, Event, Calendar | 10+ | Basic Implementation |
| MapKit    | MapView, PointAnnotation | 10+ | Basic Implementation |
| HealthKit | HealthStore, QuantityType, Workout | 15+ | Basic Implementation |
| HomeKit   | HomeManager, Home, Room, Accessory | 20+ | Basic Implementation |
| GameKit   | Player, Leaderboard, Achievement | 25+ | Basic Implementation |
| Security  | Certificate, Identity, Trust, Policy | 15+ | Basic Implementation |
| JavaScriptCore | Context, Value, VirtualMachine | 25+ | Basic Implementation |
| UserNotifications | NotificationCenter, NotificationContent, NotificationTrigger | 20+ | Basic Implementation |

### Recommended Next Frameworks
1. **Core Services** - Critical for file system operations
2. **Metal** - Modern graphics API (high priority)
3. **CoreBluetooth** - Device connectivity
4. **NetworkExtension** - Advanced networking capabilities
5. **PDFKit** - PDF rendering and manipulation

## Tools Reference

### Generation Tools
- List framework constants: `go run ./generate/tools/constant.go macos [framework] [constant]`
- Verify declaration parsing: `go run ./generate/tools/declcheck.go [framework]`
- Check parsing coverage %: `go run ./generate/tools/declcheck.go [framework] | grep "Total coverage"`
- Check imports: `./generate/tools/imports.sh ./macos/[framework]`
- Search symbolsdb: `go run ./generate/tools/lookup.go [prefix]`
- Find symbol types: `go run ./generate/tools/type.go [typename]`
- View framework stats: `./generate/tools/stats.sh`
- Regenerate all frameworks: `./generate/tools/regen.sh macos`

### Analysis Tools
- Generate enhanced coverage report: `./cmd/tools/analyze_technologies.sh`
- Analyze Apple's frameworks: `go run cmd/tools/tech_analyzer/main.go --outdir="./analysis" --count-symbols --list-macos`
- Generate module entries: `go run cmd/tools/tech_analyzer/main.go --frameworks="[framework1],[framework2]" --gen-modules`
- Analyze implementation coverage: `go run cmd/tools/report_generator/main.go --coverage="$OUTPUT_DIR/analysis/coverage_report.json" --output="API_COVERAGE.md"`

## Implementation Strategy

When implementing a new framework or extending an existing one, follow these guidelines:

### Implementation Depth
- Focus on implementing the most commonly used classes and methods first
- Create "basic implementations" that include 10-25 core methods per class
- Prioritize symbols that are fundamental to the framework's functionality
- Implementation quality is more important than quantity
- Be comprehensive in delegate implementation - implement all key delegate methods

### Implementation Approach
1. **Staged Implementation**:
   - Start with core functionality that unlocks the framework's primary features
   - Add example code that demonstrates practical usage
   - Add more advanced features in subsequent iterations

2. **Dependency Handling**:
   - Favor abstract types over forcing implementation of dependent frameworks
   - Use CanAbstractModuleCoupling for optional dependencies
   - Use CanSkipModuleCoupling for rarely used dependencies

3. **Proper Binding Usage**:
   - Use factory methods instead of struct literals for creating Foundation objects
   - Use generated enums and constants with correct prefixes (NEVPNStatus* not VPNStatus*)
   - Properly handle protocol implementations using concrete types (DrawableObject vs Drawable)
   - Always validate bindings with working examples before considering them complete

4. **Testing Strategy**:
   - Create simple tests that verify binding compilation
   - Create example applications that demonstrate real-world usage
   - Test integrations with other frameworks
   - Always validate example code works before committing

## Common Issues and Troubleshooting

1. **Class or method redeclaration**:
   When you get errors about redeclared types or methods, it usually means a custom implementation already exists in a *_custom.go file.
   - Check for duplicate definitions in generated vs. custom code
   - Make sure generated code doesn't overlap with custom implementations
   - For Metal framework, be especially careful as it has many custom implementations that might clash with generated code

2. **Cannot find method/property**:
   When a method or property appears to be missing:
   - Check if it's declared in a protocol rather than a class
   - Ensure all framework dependencies are properly imported
   - Look for platform-specific availability (macOS vs. iOS)

3. **Type conversion errors**:
   When getting type conversion errors:
   - Use the proper type conversion method (foundation.Data_ToBytes, etc.)
   - Make sure to handle memory management correctly (Retain/Release)
   - Check if the type is actually a generic id or another protocol type
   - For byte arrays, use `unsafe.Pointer(&byteSlice[0])` to get a pointer and `uint(len(byteSlice))` for the length

4. **Framework dependency conflicts**:
   For import cycles or dependency issues:
   - Update the CanAbstractModuleCoupling or CanSkipModuleCoupling maps
   - Create wrapper types for the circular dependencies

5. **Send method not found**:
   If you get errors about 'Send' method being undefined:
   - Replace all instances of the deprecated `Send` method with `objc.Call`
   - See the section on "Objective-C Method Calling Conventions" for examples
   
6. **Missing SetObjectForKey method**:
   When working with dictionaries:
   - MutableDictionary should be used for modifying dictionaries, not Dictionary
   - Use the proper method `SetObjectForKeyObject` (not SetObjectForKey)
   - Ensure you're using the correct concrete implementation class
   
7. **Framework test failures due to Metal**:
   When testing frameworks that have Metal dependencies:
   - Use specific test flags to test only certain functions: `go test -run TestSpecificFunction`
   - Consider refactoring Metal framework to prevent duplicate declarations
   - For temporary fixes, isolate the dependency chains

## Documentation
- Add good comments for custom functionality
- Document memory management expectations for custom methods
- Reference Apple documentation URLs when appropriate
- Update API_ENHANCED_COVERAGE.md when adding new frameworks