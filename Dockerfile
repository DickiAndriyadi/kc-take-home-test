FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main", "-host=0.0.0.0", "-port=8080"]