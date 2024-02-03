run:
	go run main.go

test:
	go test -coverprofile coverage.out -v ./...

## Dependencies

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor
	pre-commit autoupdate
