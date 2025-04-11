FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
RUN chmod 555 /app/main

FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --chown=appuser:appgroup --from=builder /app/main ./main
USER appuser
CMD ["./main"]

