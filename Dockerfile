# Build stage
FROM golang:alpine AS builder
RUN apk --no-cache add gcc g++ make git
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o start

# Run stage
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /app/start .
COPY --from=builder /app/favicon.ico .

RUN mkdir -p /app/uploads

EXPOSE 8088
CMD [ "/app/start" ]