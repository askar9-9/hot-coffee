FROM golang:1.23

WORKDIR /app

COPY . .

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]