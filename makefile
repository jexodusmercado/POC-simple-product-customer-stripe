.PHONY: start migrate
CHECK_FILES?=./...
FLAGS?= -buildvcs=false

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## Start the server
	@go run $(FLAGS) ./cmd/main.go

deps: ## Install dependencies
	@go mod tidy

migrate: ## run the migrations
	cd schema && for file in *.sql; do \
		PGPASSWORD=root psql -v ON_ERROR_STOP=1 -h localhost -U postgres -a -d postgres -f "$$file"; \
	done