.PHONY: dev largetype topframe

dev:
	go run ./cmd/macdriver-host

largetype:
	go run ./cmd/largetype Hello world

topframe:
	go run ./cmd/topframe