default: lint tests

tests:
	go test -count=1 ./...

lint:
	staticcheck -checks "all" ./...
