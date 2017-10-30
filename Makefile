.PHONY: all
all: setup lint test

.PHONY: test
test: setup
	go test ./...

.PHONY: cover
cover: setup
	mkdir -p coverage
	gocov test ./... | gocov-html > coverage/coverage.html

sources = $(shell find . -name '*.go' -not -path './vendor/*')
.PHONY: goimports
goimports: setup
	goimports -w $(sources)

.PHONY: lint
lint: setup
	gometalinter ./... --enable=goimports --disable=golint --vendor -t

.PHONY: check
check: setup
	gometalinter ./... --disable-all --enable=vet --enable=vetshadow --enable=goimports --vendor -t

.PHONY: ci
ci: setup check test

.PHONY: install
install: setup
	go install

BIN_DIR := $(GOPATH)/bin
GOIMPORTS := $(BIN_DIR)/goimports
GOMETALINTER := $(BIN_DIR)/gometalinter
DEP := $(BIN_DIR)/dep
GOCOV := $(BIN_DIR)/gocov
GOCOV_HTML := $(BIN_DIR)/gocov-html

$(GOIMPORTS):
	go get -u golang.org/x/tools/cmd/goimports

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

$(GOCOV):
	go get -u github.com/axw/gocov/gocov

$(GOCOV_HTML):
	go get -u gopkg.in/matm/v1/gocov-html

$(DEP):
	go get -u github.com/golang/dep/cmd/dep

tools: $(GOIMPORTS) $(GOMETALINTER) $(GOCOV) $(GOCOV_HTML) $(DEP)

vendor: $(DEP)
	dep ensure

setup: tools vendor

updatedeps:
	dep ensure -update

BINARY := locksmith
VERSION ?= latest
PLATFORMS := darwin/amd64 linux/amd64 windows/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.PHONY: $(PLATFORMS)
$(PLATFORMS): setup
	mkdir -p $(CURDIR)/release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build -ldflags="-X main.version=$(VERSION)" \
	-o release/$(BINARY)-v$(VERSION)-$(os)-$(arch)

.PHONY: release
release: $(PLATFORMS)
