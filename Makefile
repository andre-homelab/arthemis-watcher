SWAG ?= swag

.PHONY: docs main api

docs:
	$(SWAG) init -g main.go -o docs --dir api,internal/handlers,internal/models --parseDependency || \
		go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g main.go -o docs --dir api,internal/handlers,internal/models --parseDependency

api:
	go run api/main.go

dev: docs api
