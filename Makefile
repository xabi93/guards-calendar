default: lint tests

tests:
	go test -count=1 ./... -cover

lint:
	staticcheck -checks "all" ./...
