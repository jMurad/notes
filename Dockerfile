FROM golang:1.23-alpine3.20 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/notes/main.go

FROM alpine AS production

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs/notes.toml ./configs/notes.toml

EXPOSE 8080

CMD [ "./server" ]