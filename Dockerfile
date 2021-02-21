FROM golang:1.15 AS build
WORKDIR /go/src

COPY specification.yaml .

COPY LICENSE .
COPY main.go .

COPY go.* ./

COPY pkg ./pkg

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/openapi ./
