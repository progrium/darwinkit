.PHONY: dev largetype topframe

dev:
	go run ./cmd/macdriver/client
largetype:
	go run ./cmd/largetype Hello world

topframe:
	go run ./cmd/topframe