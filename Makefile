run:
	go run cmd/main.go

swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swag-gen:
	@if ! command -v swag >/dev/null 2>&1; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g api/router.go -o api/docs --parseVendor; \
	else \
		$(shell go env GOPATH)/bin/swag init -g api/router.go -o api/docs --parseVendor; \
	fi