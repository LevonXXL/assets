GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go get github.com/jackc/puddle/v2@v2.2.1

# make migration args=up (применить +1 миграцию)
# make migration args=down (отменить 1 миграцию)
.PHONY: migration
migration:
	@goose -dir=migrations postgres $(PG_DB_DSN) $(args)

.PHONY: build
build:
	go build -o assets-app

.PHONY: test
test: # Run all tests.
	@go test -count=1 -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | grep ^total | tr -s '\t'