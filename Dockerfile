FROM golang:1.23-alpine as builder

WORKDIR /app

# Cài các công cụ cần thiết
RUN apk add --no-cache curl

# Copy code vào
COPY . .

# Build ứng dụng
RUN go build -o main main.go

# Download và cài migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz \
  | tar xvz && chmod +x migrate

# ----------------------
# Image nhẹ để chạy
FROM alpine:3.18

WORKDIR /app

# Copy file đã build
COPY --from=builder /app/main .
COPY --from=builder /app/migrate /usr/local/bin/migrate

# Copy các file cần thiết
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
COPY .env .

# Cấp quyền thực thi
RUN chmod +x start.sh /app/wait-for.sh

EXPOSE 8080

ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]