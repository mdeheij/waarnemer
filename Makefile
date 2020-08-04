all: install

install:
	@echo "Installing.."
	@go install -v ./...

run: install
	@waarnemer server -v

cockroachdb:
	@docker run -d --rm --name=cockroachdb --hostname=cockroachdb -p 26257:26257 -p 8080 -v "${PWD}/.database:/cockroach/cockroach-data" cockroachdb/cockroach:v20.1.4 start --insecure

.PHONY: help run
