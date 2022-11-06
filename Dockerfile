# Build stage
FROM golang:alpine AS builder
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/go-production

# Run stage
FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/src/app/bin /go/bin
RUN mkdir -p /go/bin/public/uploads
CMD /go/bin/go-production