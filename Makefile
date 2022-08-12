.PHONY: all
all: test vet fmt

.PHONY: test
test:
	@go test ./twitter -cover

.PHONY: vet
vet:
	@go vet -all ./twitter

.PHONY: fmt
fmt:
	@test -z $$(go fmt ./...)

