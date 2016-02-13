default:
	go test -v -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
