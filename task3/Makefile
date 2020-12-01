GOPATH = $(shell go env GOPATH)

.PHONY: help test lint benchmark build

help: Makefile
	@echo " Доступные команды "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

$(GOPATH)/bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.32.2

## test: Запуск тестов
test:
	go test -v ./...

## lint: Запуск линтера
lint: $(GOPATH)/bin/golangci-lint
	$(GOPATH)/bin/golangci-lint run

## benchmark: Запуск бенчмарков
benchmark:
	go test -bench=. -benchmem ./...

## build: Сборка приложения
build:
	go build -o ./topcnt cmd/topcnt/main.go
	@echo App location: ${PWD}/topcnt
