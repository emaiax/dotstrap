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
	gotest $(TEST_OPTIONS) -v -failfast $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
