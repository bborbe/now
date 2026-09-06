REGISTRY ?= docker.io
IMAGE ?= bborbe/now
ifeq ($(VERSION),)
	VERSION := $(shell git fetch --tags; git describe --tags `git rev-list --tags --max-count=1`)
endif

deps:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kisielk/errcheck
	go get -u github.com/haya14busa/goverage
	go get -u github.com/schrej/godacov

clean:
	docker rmi $(REGISTRY)/$(IMAGE):$(VERSION)

.PHONY: check-version-tag
check-version-tag:
	@if [ -n "$(ALLOW_UNTAGGED_BUILD)" ]; then \
		echo "ALLOW_UNTAGGED_BUILD set — skipping version/tag check"; \
	else \
		head_tag=$$(git describe --tags --exact-match HEAD 2>/dev/null); \
		if [ "$$head_tag" != "$(VERSION)" ]; then \
			echo "ERROR: refusing to build $(VERSION) from this tree." >&2; \
			echo "  HEAD is at tag: $${head_tag:-<untagged>}" >&2; \
			echo "  building as:    $(VERSION)" >&2; \
			echo "  An image stamped vX.Y.Z must be built from the vX.Y.Z tag." >&2; \
			echo "  Fix: git checkout $(VERSION)   (or set ALLOW_UNTAGGED_BUILD=1 for a scratch build)" >&2; \
			exit 1; \
		fi; \
	fi

build: check-version-tag
	docker build --no-cache --rm=true -t $(REGISTRY)/$(IMAGE):$(VERSION) -f ./Dockerfile .

upload:
	docker push $(REGISTRY)/$(IMAGE):$(VERSION)

run:
	docker run \
	-e PORT=8080 \
	-p 8080:8080 \
	$(REGISTRY)/$(IMAGE):$(VERSION)

install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/now/*.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/now-server/*.go

precommit: ensure format test check addlicense
	@echo "ready to commit"

ensure:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure

format:
	@go get golang.org/x/tools/cmd/goimports
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

test:
	go test -cover -race $(shell go list ./... | grep -v /vendor/)

check: lint vet errcheck

lint:
	@go get golang.org/x/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

errcheck:
	@go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write|Fprint)' $(shell go list ./... | grep -v /vendor/)

addlicense:
	go get github.com/google/addlicense
	addlicense -c "Benjamin Borbe" -y 2018 -l bsd ./*.go ./cmd/*/*.go
