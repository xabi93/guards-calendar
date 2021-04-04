default: test lint


test:
	go test -count=1 ./...

lint:
	staticcheck -checks "all" ./...
