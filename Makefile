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