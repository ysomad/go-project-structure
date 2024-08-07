include .env
export

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=${POSTGRES_URL}

.PHONY: run-server
run-server:
	go mod tidy && go mod download && \
	pplog go run ./cmd/server/main.go

.PHONY: run-server-migrate
run-server-migrate:
	go mod tidy && go mod download && \
	go run ./cmd/server/main.go -migrate


.PHONY: gen-server
gen-server:
	rm -rf ./internal/gen/server/* && \
	swagger generate server \
		--spec=api/swagger2.yaml \
		--target=internal/gen/server \
		--model-package=model \
		--strict-responders \
		--exclude-main

.PHONY: test
test:
	go test -v -cover -race -count 1 ./internal/...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: goose-new
goose-new:
	@read -p "Enter the name of the new migration: " name; \
	goose -dir migrations create $${name// /_} sql

.PHONY: goose-up
goose-up:
	@echo "Running all new database migrations..."
	goose -dir migrations validate
	goose -dir migrations up

.PHONY: goose-down
goose-down:
	@echo "Running all down database migrations..."
	goose -dir migrations down

.PHONY: goose-reset
goose-reset:
	@echo "Dropping everything in database..."
	goose -dir migrations reset

.PHONY: goose-status
goose-status:
	goose -dir migrations status
