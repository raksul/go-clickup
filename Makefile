.PHONY: test
test: 
	go test -race -v ./...

.PHONY: vet
vet: 
	go vet ./...

.PHONY: fmt
fmt: 
	gofmt -d -s .