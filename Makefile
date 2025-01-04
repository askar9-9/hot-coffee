build:
	go build -o ./cmd/hot-coffee ./cmd/main.go

run:
	go run ./cmd/main.go --port 8080

gofumpt:
	gofumpt -l -w .

clean:
	rm -f ./cmd/hot-coffee
