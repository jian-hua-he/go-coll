.PHONY: go-version
go-version:
	docker run -v ${PWD}:/code --rm golang:1.11-rc go version