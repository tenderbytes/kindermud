VERSION_FILE=VERSION
SRC_PKGS=$(shell go list ./... | grep -v vendor)
ifeq ($(strip $(NO_REV)),)
	REV=$(shell git rev-parse --short HEAD)
endif

ifeq ($(BUILD_VERSION),)
	VERSION=$(shell cat $(VERSION_FILE))
	BUILD_VERSION_NO_REV=$(VERSION)
	ifeq ($(strip $(REV)),)
		BUILD_VERSION=$(VERSION)
	else
		BUILD_VERSION=$(VERSION)-$(REV)
	endif
endif

IMAGE_REPO=kindermud
ifeq ($(IMAGE_NAME),)
	IMAGE_NAME=$(IMAGE_REPO):$(BUILD_VERSION_NO_REV)
endif

.PHONY: clean image test

all: compile

clean:
	go clean ./...
	rm -rf dist

compile:
	CGO_ENABLED=0 go build -ldflags "-X main.appVersion=$(BUILD_VERSION)" .

dist:
	CGO_ENABLED=0 go build -ldflags "-X main.appVersion=$(BUILD_VERSION)" -o dist/kindermud
	cp config.default.yml dist/kindermud.config.yaml

image:
	docker build -t $(IMAGE_NAME) -t $(IMAGE_REPO):latest .

test:
	set -e;
	for pkg in $(SRC_PKGS); \
	do \
		go test -v $$pkg; \
	done
