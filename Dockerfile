FROM golang:1.16 AS build
WORKDIR /go/src

COPY LICENSE .

COPY pkg ./pkg
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o goals-service .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/goals-service ./
EXPOSE 8080/tcp
ENTRYPOINT ["./goals-service"]
