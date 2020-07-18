FROM golang:1.13 as builder
WORKDIR $GOPATH/src/peach
COPY . .
RUN go get golang.org/x/tools/cmd/stringer
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go generate ./...
RUN CGO_ENABLED=0 go build -o /app -a -ldflags '-extldflags "-static"' .

FROM alpine:3.11
CMD ./app -sharded=true -log=info
COPY --from=builder /app .