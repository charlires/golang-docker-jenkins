BASE_TAG=$(shell git rev-parse --short HEAD)
# image name for docker
IMAGE_NAME=charlires/webapp
# app name for go packages
APP_NAME=webapp
# base golang image tag
GOLANG_TAG=1.9.2-alpine
# build args for Dockerfile's
BUILD_BASE_ARGS=--build-arg APP_NAME=$(APP_NAME) --build-arg GOLANG_TAG=$(GOLANG_TAG)
BUILD_TEST_ARGS=--build-arg IMAGE_NAME=$(IMAGE_NAME) --build-arg BASE_TAG=$(BASE_TAG)
BUILD_ARGS=--build-arg IMAGE_NAME=$(IMAGE_NAME) --build-arg APP_NAME=$(APP_NAME) --build-arg BASE_TAG=$(BASE_TAG)
# net port for webapp
PORT=8080

publish:
	@echo ":::build publish image"
	docker push $(IMAGE_NAME):$(BASE_TAG)

build-base:
	@echo ":::building base image"
	docker build --rm -f Base.Dockerfile $(BUILD_BASE_ARGS) -t $(IMAGE_NAME)-base:$(BASE_TAG) .

build:
	@echo ":::building image"
	docker build --rm -f Build.Dockerfile $(BUILD_ARGS) -t $(IMAGE_NAME):$(BASE_TAG) .

build-test:
	@echo ":::building test image"
	docker build --rm -f Test.Dockerfile $(BUILD_TEST_ARGS) -t $(IMAGE_NAME)-test:$(BASE_TAG) .

test-unit:
	@echo ":::running unit tests"
	docker run --rm -i -v $(shell pwd)/report:/go/src/${APP_NAME}/report $(IMAGE_NAME)-test:$(BASE_TAG)

run:
	@echo ":::running dev environment"
	docker run --rm -it \
		-p $(PORT):80 \
		-v `pwd`:/go/src/$(APP_NAME) \
		-w /go/src/$(APP_NAME) \
		golang:$(GOLANG_TAG) go run app/main.go
