.PHONY: dev largetype weboverlay

dev:
	go run ./cmd/macdriver/client
largetype:
	go run ./cmd/largetype Hello world

weboverlay:
	go run ./cmd/weboverlay