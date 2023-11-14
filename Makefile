.PHONY: generate test clobber example examples

GOEXE ?= go
EXAMPLES := $(wildcard macos/_examples/*)

generate:
	$(GOEXE) generate ./...

test:
	$(GOEXE) test ./...

clobber:
	$(GOEXE) run ./generate/tools/clobbergen.go ./macos

example:
	$(GOEXE) run ./macos/_examples/helloworld/main.go

examples: _local/bin
	@for dir in $(EXAMPLES); do \
		$(GOEXE) build -o ./_local/bin/$$(basename $$dir) ./$$dir; \
	done

_local/bin:
	mkdir -p _local/bin

generate/symbols.zip:
	cd generate && wget https://github.com/mactypes/symbolsdb/releases/download/1.1/symbols.zip