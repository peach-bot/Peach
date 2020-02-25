FROM golang:1.13.8-alpine3.11 as builder
COPY . .
RUN go build -o /app .

FROM alpine:3.11
CMD ["./app"]
COPY --from=builder /app .