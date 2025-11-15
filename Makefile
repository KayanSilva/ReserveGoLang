open:
	cd good-practices
start:
	docker compose up -d
test:
	docker compose exec api go test ./test -v
lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint golangci-lint run
ci:
	open start test lint