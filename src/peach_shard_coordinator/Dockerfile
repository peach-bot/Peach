FROM golang:1.13.8 as builder
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /app -a -ldflags '-extldflags "-static"' .

FROM alpine:3.11
CMD ./app
COPY --from=builder /app .
EXPOSE 8080