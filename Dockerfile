FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.42.0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./tmp/main ./cmd/main.go
 
EXPOSE 6476

CMD ["air"]
