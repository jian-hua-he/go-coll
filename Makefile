DOCKER_VOLUME_PATH = /go/src/go-coll
DOCKER_IMAGE = golang:1.11-rc

.PHONY: bash
bash:
	docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm $(DOCKER_IMAGE) bash

.PHONY: go-version
go-version:
	docker run -v $(PWD):$(DOCKER_VOLUME_PATH) --rm $(DOCKER_IMAGE) go version

.PHONY: go-test
go-test:
	docker run -v $(PWD):$(DOCKER_VOLUME_PATH) --rm $(DOCKER_IMAGE) bash -c "cd $(DOCKER_VOLUME_PATH)/coll && go test"