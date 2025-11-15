start:
	ls
test:
	cd good-practices && docker compose exec api go test ./test -v
lint:
	cd good-practices && docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint golangci-lint run
ci:
	start test lint