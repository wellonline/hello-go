run:
	go run main.go

test:
	go test -coverprofile coverage.out -v ./...