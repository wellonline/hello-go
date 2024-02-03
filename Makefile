# Makefile

## Dependencies...

BASE_IMAGE_NAME := wellonline
SERVICE_NAME    := hello-go
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)

run:
	go run main.go

test:
	go test -coverprofile coverage.out -v ./...

## Dependencies

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor
	pre-commit autoupdate

# ==============================================================================
# Building containers

service:
	docker build \
		-f zarf/docker/dockerfile.service \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
