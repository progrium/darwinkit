# Generation

## How Generation Works

At a high level, given an entry in [modules.go](../generate/modules/modules.go) and a package directory for a framework, you can use `go generate` to generate or re-generate the majority of the source files for a framework based on the [bindings API spec](./bindings.md). Below is an overview of the parts and packages involved so you know where to go when you inevitably deal with hangups [adding a new framework](#adding-new-frameworks).

### modules package

This contains [modules.go](../generate/modules/modules.go) which is where the list of known frameworks is defined, as well as some helpers in determining what to do in several scenarios. 

One is finding a type that belongs to a module that is not known. Generate will panic in this case so you can decide what to do: add the framework, or add it to the ignore list. Then you can run generate again. That actually sums up the process well: Generate, get panic for unknown situation, write code to handle situation, generate again.

There are also two lists that say what to do when a known module is encounted in the context of generating a module. This is to help decouple modules to avoid circular imports. You have two options right now: skip generating the method/property/whatever that involves the type from a module, or abstract the type. The latter means instead of using the appropriate bindings type or interface, it will instead use a generic Object type or interface. I have improvements of this in mind (like recreating types and interfaces in a sub `deps` package for the module), but for now this should work. Most decouplings aren't huge dependencies, they're the one off random method on NSString depending on GameplayKit for some reason kind. 

### modules/enums

While we have a database of symbols (next section), we don't have a consistent database of values for constants. So we have a tool called [enumexport](#enumexport-framework) that will create a file for each platform's framework containing every constant and its value. Then we reference this when generating `enumtypes.gen.go`.

## symbolsdb

This "database" is a zip file full of JSON documents for every symbol across the Apple frameworks. It was produced using the Apple docset of a documentation app, and in this project is treated as an immutable external resource. Be sure to run `make generate/symbols.zip` to download it. Leave it in the generate package directory and use some of the tools below to query it. You can also unzip it and look through it by hand, but that's usually more tedious. Every JSON file matches this schema:

```golang
type Symbol struct {
	Name string
	Path string
	Kind string

	Description string
	Type        string
	Parent      string
	Modules     []string
	Platforms   []Platform
	Declaration string
	Parameters  []Parameter
	Return      string
	Deprecated  bool
}

type Platform struct {
	Name         string
	IntroducedAt string
	Current      string
	Beta         bool
	Deprecated   bool
	DeprecatedAt string
}

type Parameter struct {
	Name        string
	Description string
}
```

The kind of symbols it has are:

* classes
* constants
* enums
* frameworks
* functions
* macros
* methods
* properties
* protocols
* structs
* unions
* types (other typedefs)

We don't use all of these yet, but we can get there!

### declparse

The symbols from symbolsdb usually include the raw declaration you'd find in a header file for that symbol. We need to parse this into a structured representation and that's what declparse does. It's a tiny lexer and parser specifically for declarations, it doesn't parse program code. It's used specifically for parsing the declarations from symbols. 

There may be cases it can't parse yet. The tool [declcheck](#declcheck-framework) lets you run through all the symbols for a framework and see the errors it runs into with a report at the end showing the percentage of coverage it has. 

declparse has a test suite, so typically we will throw an instance of a kind of declaration it has a problem with into the test suite and then patch the parse code to make tests pass. When symbols are encountered it can't parse, they are usually skipped. 

### codegen

This is where actual code generation happens so to modify the structure of source files, this is where to look. This package has types it uses for generation that are setup by the core generate package by reading in symbols and usually setting up a typing type.

### typing

This is where types are modeled for conversion from Objective-C types to Go types that's used by codegen. The original generation code we adopted had codegen and typing without any dependencies on their source of metadata, so they've been used or expanded on as is.

### generate

This root package orchestrates it all. It has a wrapper for the symbolsdb for looking up and caching access to it. It also has a couple of significant large switches used for parsing found types into the appropriate typing type. So this is often modified when new types or kinds of types are run into in APIs while generating new frameworks.

The members.go file has a list of type names that if found will skip the property or method using them. Reasons for this might be all kinds of things. Maybe we don't know what to do with that type yet. Maybe that type name overlaps with another type name somehow. Maybe that type is deprecated and isn't actually defined any more. Or maybe the type was added to just get some framework that was giving trouble to generate. Most of these types should be revisited and handled properly.

## Adding New Frameworks

So you want to add a new framework! Great! It might be easy, but probably not. At least it should get easier the more we add and the more we refactor and improve our generation code. And you will probably be doing a little of this in this process. At the very least, adding types to lists to ignore for now. Also be sure you've run `make generate/symbols.zip` before starting.

### Step 1: Add framework to modules.go

There is a list of known modules in [modules.go](../generate/modules/modules.go) and you want to add an entry to it. Order doesn't really matter, though if you're adding two modules with prefixes that share a common prefix, you want the longer prefix first. 

While you're in here, make sure the module isn't in the `CanIgnoreNotFound` list. If your framework depends on a framework that you don't want to add yet, you'll put its name in this list. 

### Step 2: Run enumexport to get constants

The [enumexport](#enumexport-framework) tool uses symbolsdb to get all supported enums and constants, writes a program that will output the values for them, and runs that program. The tool has constants that define the platform and version it focuses on. For now you'd just change those if needed. Ideally, you'd run this on the latest version so it includes all known constants. If not, whatever version you're on will be good enough for now. Somebody else in the community can run it on a newer version if needed, just make an issue for it. 

This is another point you might need to modify code, in this case code of the enumexport tool itself. It has its own list of names and prefixes to ignore for various reasons. Usually its because the constant just isn't defined any more, it was deprecated and removed. Leave a comment as to why you added it to a list if you do. 

Run the tool with your framework as the argument first and inspect the output to make sure it works. You may also discover prefixes you need to add to your module entry. There is also a case where a framework's header file isn't enough to include the constants found in symbolsdb. There are some hand added extra includes for specific frameworks which you may need to replicate. 

Once it generates good output, you can pipe it to the appropriate file under ./generate/modules/enums/macos ... if you run it without a framework, it will regenerate all these files, which is an option as well, but be sure not to include any removals in your PR that might be from generating on an older version of the platform that whats in there. 

### Step 3: Run initmod to make the new framework package

The [initmod](#initmod-platform-framework) tool can now be run to generate the initial files for the package that are not intended to be generated again. At this point you won't have to modify anything in these yet, they're good as is. 


### Step 4: Run go generate for your framework

Now the fun begins. Run `go generate ./macos/your-framework` and it will attempt to generate it. If it panics, that is by design. That means it found a type or module you need to handle. Most panics should have extra logging information before the panic to help you identify where it found the type. But it panics so that if you need more information you can add extra debugging along the stack trace path. Just be sure to remove these when you don't need them.

Without panics, the output should only be telling us the skips its making or choices on fall back types its making. Usually these are fine and intended. 

Once it runs successfully, you have generated code and can check it out. But you might not be done yet! After generation, always run `go test ./macos/your-framework`. There is an empty test made by `initmod` just to make sure the generated code can compile. Here you are most likely to discover types that don't exist. The most common of these are struct types. Structs are not generated yet, so they need to be added manually. Sometimes they'll be structs for *other* frameworks. If they have "Ref" in the name and are pointers, you can usually just make these `unsafe.Pointer` types. 

It very common that you won't be able to generate a new framework without adding another framework. So these sessions might involve generating several frameworks at once. Sometimes these frameworks depend on each other, and so the tests will reveal were circular imports are. 

### Step 5: Decouple to avoid circular imports

This is the hardest to describe best practices because it involves judgement calls and some sleuthing. The basic idea is that you will want to inspect the imports of the framework in question using the [imports](#imports-path) tool. This tool aggregates all the imports across source files in a package/directory so you can see everything the package imports. That's about all the help you get, otherwise you'll be doing a search in your editor to find out what's using what from those imports. 

Then you can add select dependent packages to the appropriate list mentioned above in [module package](#modules-package) to decouple them. Usually these are to avoid pulling in a dependency that has a dependency chain that comes back to you. I guess this is more common when adding a lot of frameworks at once.

The more common situation is you add a framework that seriously depends on another framework, so you decide to generate it as well, only to find they both directly depend on each other. In this case, picking the one to decouple from the other depends. One approach is to pick the framework that is not directly used as much. There are a lot of these kinds of lower-level frameworks. Another approach is to pick the one with the fewest dependency points. Maybe there are other approaches, but it's up to you so good luck! Worst case is somebody will disagree in the PR, so be sure to explain your choices in the PR.


### Step 6: Enjoy but also improve this process

If steps or issues come up not explained here, definitely bring them up on GitHub or Discord. If you definitely think something needs to be added to this document, add it. If you definitely think something was done wrong in the generation code, fix it. That's it! Thank you.


## Tools Overview

> [!IMPORTANT]
> All of these tools are made to be run from the project root.

### clobbergen [path]

Removes all files under the path that have the auto-generated banner

`go run ./generate/tools/clobbergen.go ./macos/appkit`

### constant [platform] [framework] [constant]

Shows the value for a constant from ./generate/modules/enums

`go run ./generate/tools/constant.go macos appkit NSWindowBelow`

### declcheck [framework]

Runs declparse against symbol declarations for a given framework in symbolsdb.

`go run ./generate/tools/declcheck.go appkit`

### enumexport [framework]

Goes through symbolsdb for constants and enum values that are ints, strings, or floats to generate a program that will output the constant and value as found on the current platform. Though it limits symbols to those on macos-12 right now, this is modifiable in constants in the source.

With a framework argument it will output that frameworks constants to stdout, which could be inspected or piped into a file to update a specific framework in ./generate/modules/enums:

`go run ./generate/tools/enumexport.go appkit > ./generate/modules/enums/macos/appkit`

Without arguments it will generate files in ./generate/modules/enums for all known frameworks.

`go run ./generate/tools/enumexport.go`

### genmod

Runs generation for a framework package and is not meant to be used directly, but by `go generate` via the `go:generate` declaration at the top of the framework package main source file. It can be used for a specific framework or all frameworks:

`go generate ./macos/appkit`

`go generate ./...`

### initmod [platform] [framework]

Makes the framework package directory if it doesn't exist and creates the starting non-generated
files for the framework, which must be a known framework in ./generate/modules.

`go run ./generate/tools/initmod.go macos appkit`

### imports [path]

Shell script that shows all the Go imports for a package path.

`./generate/tools/imports.sh ./macos/appkit`

### lookup [prefix]

Finds all symbols in symbolsdb with a path prefixed with the given prefix. Helpful to inspect symbolsdb, especially in combination with `jq`.

`go run ./generate/tools/lookup.go appkit | jq 'select(.Kind == "Framework")`

### stats

Shell script that gives basic stats on generated frameworks.

`./generate/tools/stats.sh`

### type [symbol]

Finds type symbol(s) in symbolsdb with the given type name. Can be used with `jq`.

`go run ./generate/tools/type.go NSWindow`

### regen [platform]

Re-generates frameworks that have been generated (have .gen.go files).

`./generate/tools/regen.sh macos`
