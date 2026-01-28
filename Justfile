EXE := "gophinvaders"

default:
	@just --list

test: unittest lint fmt-check gosec tidy build

unittest:
	go test ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found, falling back to go vet"; \
		echo "To install golangci-lint, run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		go vet ./...; \
	fi

gosec:
	@echo "Running security scanner..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec -quiet -fmt=text ./...; \
	else \
		echo "gosec not found, skipping security scan"; \
		echo "To install gosec, run: go install github.com/securego/gosec/v2/cmd/gosec@latest"; \
	fi

fmt:
	go fmt ./cmd/... ./pkg/...

# Check formatting without modifying files
fmt-check:
	./tools/bin/go-fmt-check

tidy:
	go mod tidy

build:
	mkdir -p bin
	go build -o bin/{{EXE}} ./cmd/{{EXE}}
# Build with version info embedded, when appropriate
#	go build -ldflags "-X github.com/sfkleach/pathman/pkg/commands.Source=https://github.com/sfkleach/pathman" -o bin/{{EXE}} ./cmd/{{EXE}}


install:
	go install ./cmd/{{EXE}}

clean:
	rm -rf bin
	rm -rf dist
	rm -rf _build


# Initialize decision records
# init-decisions:
#    python3 scripts/decisions.py --init

# Add a new decision record
# add-decision TOPIC:
#    python3 scripts/decisions.py --add "{{TOPIC}}"


