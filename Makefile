DOCKER_PORT = 80
HOST_PORT = 6060

DOCKER_VOLUME_PATH = /go/src/github.com/jian-hua-he/go-coll
DOCKER_IMAGE = golang:1.11-rc

BASE_COMMAND ?= docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm -p $(HOST_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE)

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-test
go-test:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH)/coll && go test"

.PHONY: go-doc
go-doc:
	$(BASE_COMMAND) godoc -http=0.0.0.0:80

