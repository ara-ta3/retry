install/test:
	go get -u github.com/stretchr/testify

test:
	go test -v ./...
