APP := bin/md-table-gen

GO := go


.PHONY: all build run test clean

all: test build

build:
	@$(GO) build -o $(APP) cmd/*go

run:
	@$(GO) run $$(ls cmd/* | grep -v _test)

test:
	@$(GO) test -v ./...

gen-output: build
	@echo "creating sample files in _temp for testing"
	@mkdir _temp || true
	@cp -a examples/*.md _temp/
	@for i in _temp/*.md ; do $(APP) bin/md-gen-table -f examples/test2.yaml -f examples/test.yaml -r "$${i}" ; done
	
clean:
	@rm -rf $(APP)
	@rm -rf _temp
