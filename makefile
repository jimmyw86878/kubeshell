SHELL=/bin/bash
name=kubeshell
tag=latest
all: build_docker_image

build_docker_image:
	@echo Start build image
	docker build -t=$(name):$(tag) -f build/package/Dockerfile ..
	docker image prune --filter label=stage=build -f
