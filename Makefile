CURRENT_DIR := $(shell pwd)
APP := template
APP_CMD_DIR := ./cmd


swag-gen:
	swag init -g api/router.go -o api/docs