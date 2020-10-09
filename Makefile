.PHONY: build
build: 
	@go build -o policy-hub

.PHONY: test
test: 
	@go test -v ./...

.PHONY: check-fmt
check-fmt:
	@test -z $$(gofmt -l .) || echo $$(gofmt -l .)

.PHONY: check-vet
check-vet:
	@go vet ./...

.PHONY: check-lint
check-lint:
	@golint -set_exit_status ./...

.PHONY: check
check: check-fmt check-vet check-lint