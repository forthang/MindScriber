FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod tidy && go build -o myApp ./cmd/api-service
FROM alpine
WORKDIR /app
COPY --from=builder /build/myApp .
CMD ["./myApp"]