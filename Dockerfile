FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .


FROM alpine:3.21
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/main ./main
RUN chmod 555 /app/main
USER appuser
CMD ["./main"]

