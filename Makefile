.DEFAULT_GOAL:=all
ifdef VERBOSE
V = -v
else
.SILENT:
endif

GO ?= GO111MODULE=on go

include $(wildcard *.mk)

COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
	SERVICE=circle-ci
endif

.PHONY: all
all: generate fmt build test cover

.PHONY: benchmarks
benchmarks: generate fmt build
	./benchmark.sh

build: generate
	if [ ! -d bin ]; then mkdir bin; fi
	$(GO) build -v -o bin/gencheck ./gencheck
fmt:
	gofmt -l -w -s $$(find . -type f -name '*.go' -not -path "./vendor/*")
test: generate gen-test
	if [ ! -d coverage ]; then mkdir coverage; fi
	$(GO) test -v ./... -race -cover -coverprofile=$(COVERAGEDIR)/coverage.out

cover:
	$(GO) tool cover -html=$(COVERAGEDIR)/coverage.out -o $(COVERAGEDIR)/coverage.html

tc: test cover
coveralls: $(GOVERALLS)
	$(GOVERALLS) -coverprofile=$(COVERAGEDIR)/coverage.out -service=$(SERVICE) -repotoken=$(COVERALLS_TOKEN)

clean: cleandeps
	$(GO) clean
	rm -f bin/gencheck
	rm -rf coverage/

# go-bindata will take all of the template files and create readable assets from them in the executable.
# This way the templates are readable in atom (or another text editor with golang template language support)
generate: deps
	$(GOBINDATA) -o generator/assets.go -nometadata -pkg=generator template/*.tmpl

gen-test: build
	$(GO) generate ./...

install:
	$(GO) install ./gencheck

phony: clean tc build
