.PHONY: all
all: test vet lint fmt

.PHONY: test
test:
	@go test ./twitter -cover

.PHONY: vet
vet:
	@go vet -all ./twitter

.PHONY: lint
lint:
	@golint -set_exit_status ./...

.PHONY: fmt
fmt:
	@test -z $$(go fmt ./...)

