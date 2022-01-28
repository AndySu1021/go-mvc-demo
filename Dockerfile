FROM golang:1.17-buster

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

WORKDIR /mvc

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]