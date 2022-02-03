FROM golang:1.16 as build

WORKDIR /mvc

COPY . .

RUN go build -o app .

FROM gcr.io/distroless/base

WORKDIR /mvc

COPY --from=build /mvc /mvc

CMD ["./app"]