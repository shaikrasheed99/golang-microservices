FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . . 

ENTRYPOINT ["go", "run"]

CMD ["main.go"]