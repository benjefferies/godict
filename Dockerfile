FROM golang as builder

COPY godict.go /go/src/github.com/benjefferies/godict/godict.go
WORKDIR /go/src/github.com/benjefferies/godict

ENV CGO_ENABLED=0

RUN go get -v ./...
RUN go build -o godict -v ./...

FROM alpine

COPY --from=builder /go/src/github.com/benjefferies/godict/godict /usr/local/bin/godict

ENTRYPOINT ["/usr/local/bin/godict"]
