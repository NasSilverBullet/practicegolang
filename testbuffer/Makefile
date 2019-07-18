# const
BINARYNAME=practicegolang
MAINDIR=

# enable go module
export GO111MODULE=on

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: help build test clean remove deploy analyse run share

# command
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## go build
	go build -v -o ${BINARYNAME} ./${MAINDIR}

test: ## go test
	go test -v -cover ./...
clean: ## go clean
	go clean -cache -testcache ./...

analyze: ## do static code analysis
	goimports -l -w .
	go vet ./...
	golint ./...

remove: ## remove binary
	rm -f ./${BINARYNAME}

deploy: remove clean test analyze build ## run 'build' with 'remove', 'clean', 'test', 'analyze' and 'build'

run: ## mytail [args]
	./${BINARYNAME} ${RUN_ARGS}

share: ## generate The Go Playground URL
	find . ! -path '*/\.*' -a ! -name 'mytail' -a ! -name 'piyo.txt' -type f | xargs gp share
