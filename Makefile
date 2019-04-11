test:
	go build gonum.go
	go generate
	go test .

release:
	rm -rf dist
	goreleaser
