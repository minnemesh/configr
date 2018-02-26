build: dependencies
	go build node/main.go

dependencies:
	which dep || go get -u github.com/golang/dep/cmd/dep
	dep ensure

test: dependencies
	go test ./node/...
