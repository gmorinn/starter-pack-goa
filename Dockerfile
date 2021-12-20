FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/app/
WORKDIR /go/src/app
RUN go mod download
COPY . /go/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bucketeer app

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/app/build/bucketeer /usr/bin/bucketeer
EXPOSE 8088 8088
ENTRYPOINT ["/usr/bin/bucketeer"]