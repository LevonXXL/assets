GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

# make migration args=up (применить +1 миграцию)
# make migration args=down (отменить 1 миграцию)
.PHONY: migration
migration:
	@goose -dir=migrations postgres $(PG_DB_DSN) $(args)

.PHONY: build
build:
	go build -o assets-app
