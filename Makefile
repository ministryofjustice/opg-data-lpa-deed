SHELL = '/bin/bash'
export AWS_ACCESS_KEY_ID ?= X
export AWS_SECRET_ACCESS_KEY ?= X

help:
	@grep --no-filename -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build containers
	docker compose build --parallel lambda-create lambda-update lambda-get apigw

up: ## Start application
	docker compose up -d apigw
	make create-tables

down: ## Stop application
	docker compose down

test-api: ## Test the API endpoints
test-api: URL ?= http://localhost:9000
test-api:
	go build -o ./signer/test-api ./signer && \
	chmod +x ./signer/test-api && \
	./signer/test-api PUT $(URL)/lpas/M-AL9A-7EY3-075D '{"version":"1"}' && \
	./signer/test-api POST $(URL)/lpas/M-AL9A-7EY3-075D/updates '{"type":"BUMP_VERSION","changes":[{"key":"/version","old":"1","new":"2"}]}' && \
	./signer/test-api GET $(URL)/lpas/M-AL9A-7EY3-075D '' | grep '"version":"2"' \

create-tables:
	docker compose run --rm aws dynamodb describe-table --table-name deeds || \
	docker compose run --rm aws dynamodb create-table \
		--table-name deeds \
		--attribute-definitions AttributeName=uid,AttributeType=S \
		--key-schema AttributeName=uid,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST

	docker compose run --rm aws dynamodb describe-table --table-name events || \
	docker compose run --rm aws dynamodb create-table \
		--table-name events \
		--attribute-definitions AttributeName=uid,AttributeType=S AttributeName=created,AttributeType=S \
		--key-schema AttributeName=uid,KeyType=HASH AttributeName=created,KeyType=RANGE \
		--billing-mode PAY_PER_REQUEST

run-structurizr:
	docker pull structurizr/lite
	docker run -it --rm -p 4080:8080 -v $(PWD)/docs/architecture/dsl/local:/usr/local/structurizr structurizr/lite

run-structurizr-export:
	docker pull structurizr/cli:latest
	docker run --rm -v $(PWD)/docs/architecture/dsl/local:/usr/local/structurizr structurizr/cli \
	export -workspace /usr/local/structurizr/workspace.dsl -format mermaid

go-lint: ## Lint Go code
	docker compose run --rm go-lint

gosec: ## Scan Go code for security flaws
	docker compose run --rm gosec
