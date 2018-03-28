FROM golang:1.10

# Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers
COPY . .
RUN dep ensure -v
RUN go build -o bin/server

EXPOSE 8080

CMD ["./bin/server"]
