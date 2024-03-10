FROM golang:latest

WORKDIR /app

COPY ./src .

RUN go mod download

RUN go build -o main .

EXPOSE 8888

CMD ["./main"]