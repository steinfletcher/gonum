test:
	go build -ldflags "-X main.version=`git describe --tags`" gonum.go
	go generate
	go test .

release:
	rm -rf dist
	goreleaser
