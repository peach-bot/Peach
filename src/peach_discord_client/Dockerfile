FROM golang:1.13 as builder
WORKDIR $GOPATH/src/Peach
COPY . .
RUN go get golang.org/x/tools/cmd/stringer
RUN go generate ./...
RUN go get github.com/stretchr/testify
RUN go get -d ./...
RUN CGO_ENABLED=0 go build -o /app -a -ldflags '-extldflags "-static"' .

FROM alpine:3.11
CMD ./app
COPY --from=builder /app .