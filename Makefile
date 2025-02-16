GOARCH = amd64
PKG_VERSION != git describe --tags
COMMIT != git rev-parse --short HEAD
APP_NAME ?= RENAMEME
PACKAGE_NAME ?= github.com/kormiltsev/be

build:
	go build -o $(APP_NAME)-$(GOARCH) \
	-X $(PACKAGE_NAME)/version.Version=$(PKG_VERSION) \
	-X $(PACKAGE_NAME)/version.GitCommit=$(COMMIT)" .

build_i386:
	@GOARCH=386 GO386=387 make build

build_amd64:
	@GOARCH=amd64 make build

build_arm64:
	@GOARCH=arm64 make build

.PHONY: build build_i386 build_amd64 build_arm64
