.PHONY: all check format-check vet staticcheck test build run

all: build

check: format-check vet staticcheck test

format-check:
	@files="$$(gofmt -l $$(find . -type f -name '*.go' -not -path './vendor/*'))"; \
	if [ -n "$$files" ]; then \
		printf 'Go files need formatting:\n%s\n' "$$files"; \
		exit 1; \
	fi

vet:
	go vet ./...

staticcheck:
	staticcheck ./...

test:
	go test -race ./...

build: check
	go build ./...

run: check
	go run ./cmd/petes-pottery-paradise
