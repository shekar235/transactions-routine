run:
	go run main.go

test:
	go test ./...

mockgen:
	mockgen -source=services/transaction_service.go -destination=mocks/mock_transaction_service.go -package=mocks
	mockgen -source=services/account_service.go -destination=mocks/mock_account_service.go -package=mocks

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

docker-build:
	docker build -t transactions-routine .

docker-run:
	docker run -p 8080:8080 transactions-routine
