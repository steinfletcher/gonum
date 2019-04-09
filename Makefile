test:
	go build gonum.go
	go generate
	go test .
