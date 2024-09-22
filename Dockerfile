FROM golang:1.22.7-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o registro main.go

CMD ["./registro"]
