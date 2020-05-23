SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

.PHONY: setup
setup:
	go get -u github.com/rakyll/gotest
	go mod download
	go generate -v ./...

.PHONY: test
test:
	DEBUG=true gotest $(TEST_OPTIONS) -v -failfast $(SOURCE_FILES) -run $(TEST_PATTERN) -covermode=atomic -coverprofile=coverage.out -timeout=30s
	go tool cover -html=coverage.out -o coverage.html
