FROM golang:1.13 as builder
RUN go get -t github.com/stretchr/testify
COPY . .
RUN go get -d ./...
RUN CGO_ENABLED=0 go build -o /app -a -ldflags '-extldflags "-static"' .

FROM alpine:3.11
CMD ./app
COPY --from=builder /app .
EXPOSE 8080