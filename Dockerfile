FROM golang:alpine as builder

RUN mkdir -p /url-shortener
WORKDIR /url-shortener
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener

FROM alpine:latest

RUN mkdir -p /cmd
WORKDIR /cmd

COPY --from=builder /url-shortener/url-shortener .

CMD ["./url-shortener"]