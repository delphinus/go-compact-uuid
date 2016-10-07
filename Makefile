# ref. http://postd.cc/auto-documented-makefile/
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

glide: ## Install glide
	go get github.com/Masterminds/glide

test: glide ## Run tests only
	go test $$(glide novendor)

install: glide ## Install packages for dependencies
	glide install

update: glide ## Update packages for dependencies
	glide update
