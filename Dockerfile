FROM golang:1.22-alpine3.18 AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz \
  | tar xvz && chmod +x migrate

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .
COPY start.sh .
COPY wait-for.sh .
COPY --from=builder /app/migrate /usr/local/bin/migrate
COPY db/migration ./migration
COPY .env .

RUN chmod +x start.sh
RUN chmod +x /app/wait-for.sh

EXPOSE 8080

ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]