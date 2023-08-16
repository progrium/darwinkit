GOEXE ?= go

generate:
	$(GOEXE) generate ./...
.PHONY: generate

clobber:
	$(GOEXE) run ./generate/tools/clobbergen.go ./macos
.PHONY: clobber

example:
	$(GOEXE) run ./macos/_examples/widgets/main.go
.PHONY: example

generate/symbols.zip:
	cd generate && wget https://github.com/mactypes/symbolsdb/releases/download/1.0/symbols.zip