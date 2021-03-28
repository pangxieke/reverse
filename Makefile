IMAGE=reverse
TAG1=$(shell git describe --always --tags)
TAG=${TAG1}_v1_20210326

REGISTRY=registry.cn-shenzhen.aliyuncs.com/pangxieke

default: all

fmt:
	go fmt ./...

all: fmt
	docker build -t ${REGISTRY}/${IMAGE}:${TAG} .

push:
	docker push ${REGISTRY}/${IMAGE}:${TAG}

publish: push
	docker tag ${REGISTRY}/${IMAGE}:${TAG} ${REGISTRY}/${IMAGE}:latest
	docker push ${REGISTRY}/${IMAGE}:latest

builder:
	docker build -t go-builder:latest . -f builder.dockerfile
	docker push go-builder:latest

test:
	gotest -v ./... || go test -v ./...

lint:
	ls -l | grep '^d' | awk '{print $$NF}' | grep -v vender | xargs golint

count:
	cloc --progress=1 ./ --exclude-dir=vendor,doc,pb

.PHONY: test pb stages

