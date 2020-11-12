.PHONY: dev largetype

dev:
	go run ./cmd/macdriver/client/client.go

largetype:
	go run ./cmd/largetype Hello world
