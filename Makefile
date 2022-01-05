LINTER_VERSION=v1.43.0
MOQ_VERSION=v0.2.3
GOTEST_VERSION=v0.0.6
GOTEST=$$GOPATH/bin/gotest

## setup: set up all dependencies a developer needs
setup: setup-ci
	go install github.com/matryer/moq@$(MOQ_VERSION)

setup-ci:
	go mod vendor
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)
	go install github.com/rakyll/gotest@$(GOTEST_VERSION)

ci:
	@git pull -r && make test && git push

t: test
test: lint unit-test

start: start-ci
start-ci:
	go run ./cmd/githubstatus

l: lint
lint:
	golangci-lint run ./...

lf: lintfix
lintfix:
	golangci-lint run ./... --fix

rm: regenerate-mocks
regenerate-mocks:
	find . -iname '*_moq.go' -exec rm {} \;
	go generate ./...

##========= INFRASTRUCTURE-Y THINGS =========##
ut: unit-test
unit-test: unit-test-ci
unit-test-ci:
	@clear; printf '\033[3J'
	@echo "========UNIT TESTS========"
	@$(GOTEST) -race -shuffle=on --tags=unit  ./...
	@echo '✅  UNIT TESTS'
