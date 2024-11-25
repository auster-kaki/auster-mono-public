lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --timeout 5m

test:
	go test -v ./...

fmt:
	go fmt ./... && go run golang.org/x/tools/cmd/goimports@latest -w .

build:
	# フラグの詳細：https://qiita.com/ssc-ynakamura/items/da37856f7f217d708a07
	CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o bin/api cmd/api/main.go