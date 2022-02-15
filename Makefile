FONT_ESC := $(shell printf '\033')
FONT_BOLD := ${FONT_ESC}[1m
FONT_NC := ${FONT_ESC}[0m #No color

all:
	@echo "Use a specific goal. To list all goals, type 'make help'"

.PHONY: build # Builds for current OS binary
build:
	@pkger -include /assets
	@go build
	@mv sdk-cli sx

.PHONY: install # installs the project
install:
	@go install

.PHONY: test # Runs unit tests
test:
	@go test -v ./...

.PHONY: help # Generates list of goals with description
help:
	@echo "Available goals:\n"
	@grep '^.PHONY: .* #' Makefile | sed "s/\.PHONY"