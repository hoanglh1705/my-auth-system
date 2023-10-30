.DEFAULT_GOAL := help

# HOST is only used for API specs generation
HOST ?= localhost:8081

depends: ## Install & build dependencies
	go get ./...
	go build ./...
	go mod tidy

mod.clean:
	go clean -modcache

mod: 
	go mod tidy && go mod vendor

run: 
	@go run cmd/server/main.go