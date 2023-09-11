FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN go build goApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]