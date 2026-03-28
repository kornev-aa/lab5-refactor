.PHONY: run-cli run-http run-gui

run-cli:
	go run ./cmd/linux/cli/main.go

run-http:
	go run ./cmd/linux/http/main.go

run-gui:
	go run ./cmd/linux/gui/main.go
