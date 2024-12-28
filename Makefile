build:
	go build -o hot-coffee .

run:
	go run main.go --port 8080

gofumpt:
	gofumpt -l -w .

clean:
	rm -f hot-coffee
