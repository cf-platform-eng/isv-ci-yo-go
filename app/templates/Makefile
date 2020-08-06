SHELL = /bin/bash
GO-VER = go1.14

default: build

# #### GO Binary Management ####

deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$(go version)" && \
	 	go version | grep $(GO-VER) > /dev/null

HAS_GO_IMPORTS := $(shell command -v goimports;)

deps-goimports: deps-go-binary
ifndef HAS_GO_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif


# #### CLEAN ####

clean: deps-go-binary 
	rm -rf build/*
	go clean --modcache

# #### DEPS ####

deps-modules: deps-goimports deps-go-binary
	go mod download

deps-counterfeiter: deps-modules
	command -v counterfeiter >/dev/null 2>&1 || go get -u github.com/maxbrunsfeld/counterfeiter/v6

deps: deps-modules deps-counterfeiter


# #### BUILD ####

SRC = $(shell find . -name "*.go" | grep -v "_test\." )
VERSION := $(or $(VERSION), "dev")
LDFLAGS="-X github.com/cf-platform-eng/<%= appName %>/version.Version=$(VERSION)"

build/<%= appName %>: $(SRC)
	go build -o build/<%= appName %> -ldflags ${LDFLAGS} ./cmd/<%= appName %>/main.go

build: deps build/<%= appName %>

build/<%= appName %>-linux: $(SRC)
	GOARCH=amd64 GOOS=linux go build -o build/<%= appName %>-linux -ldflags ${LDFLAGS} ./cmd/<%= appName %>/main.go

build-linux: deps build/<%= appName %>-linux

build/<%= appName %>-darwin: $(SRC) $(GENERATE_ARTIFACTS)
	GOARCH=amd64 GOOS=darwin go build -o build/<%= appName %>-darwin -ldflags ${LDFLAGS} ./cmd/<%= appName %>/main.go

build-darwin: deps build/<%= appName %>-darwin

build-all: build-linux build-darwin


# #### TESTS ####

units: deps 
	ginkgo -r -skipPackage features .

features: deps $(GENERATE_ARTIFACTS)
	ginkgo -r -tags=feature features

test: deps lint units features

lint: deps-goimports
	git ls-files | grep '.go$$' | xargs goimports -l -w
