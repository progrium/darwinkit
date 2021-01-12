
GOBIN ?= go

help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
.PHONY: help

largetype: ## build and run largetype example
	$(GOBIN) run ./examples/largetype -font="Monaco" "Hello world"
.PHONY: largetype

topframe: ## build and run topframe example
	$(GOBIN) run ./examples/_topframe
.PHONY: topframe

pomodoro: ## build and run pomodoro example
	$(GOBIN) run ./examples/pomodoro
.PHONY: pomodoro